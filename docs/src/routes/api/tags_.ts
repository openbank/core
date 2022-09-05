import { allTags } from "$lib/content";
import type { Package, Type } from "$lib/packages";
import type { EndpointOutput, RequestEvent } from "@sveltejs/kit";

export async function get({ url }: RequestEvent): Promise<EndpointOutput> {
	const filter = url.searchParams.get("tag");
	const tags = await allTags();

	const pages = tags.map((tag) => ({ name: tag.name, path: tag.path }));

	if (!filter) {
		return {
			body: {
				pages,
				pkgs: [],
			},
		};
	}

	const used = new Set<string>();

	// check recursively adds types to the used set. It does not recurse into
	// references.
	const check = (t?: Type) => {
		if (!t) return;
		switch (t.type) {
			case "message":
				t.fields.forEach((f) => {
					check(f.type);
				});
				break;
			case "ref":
				used.add(t.name);
				break;
			case "map":
				check(t.key);
				check(t.value);
				break;
			case "array":
				check(t.value);
				break;
		}
	};

	// Add used types for the current tag's packages.
	const allPkgs = tags.flatMap((t) => t.packages || []);
	tags
		.filter((t) => t.packages && (t.path == filter || filter == "all"))
		.forEach((tag) =>
			tag.packages.forEach((p) => {
				Object.values(p.types).forEach((t) => {
					check(t);
				});
				p.services
					.flatMap((s) => s.endpoints)
					.forEach((e) => {
						check(e.request);
						check(e.response);
						if (e.body_field != "" && e.body_field != "*") {
							// Recurse into body fields as they are a reference.
							const { type } = e.request.fields.find(
								(f) => f.name == e.body_field,
							);
							if (type.type !== "ref") {
								throw new Error(
									"expected reference for body field, got " +
										type.type +
										"instead",
								);
							}
							const name = type.name;
							used.add(name);
							const pkgID = name.substring(0, name.lastIndexOf("."));
							const pkg = allPkgs.find((p) => p.id === pkgID);
							check(pkg.types[name]);
						}
					});
			}),
		);

	const addTag = (p: Package, tag: string) => {
		p.tag = tag;
		return p;
	};

	const pkgs = [];
	tags.forEach((tag) => {
		const toAdd = tag.packages || [];

		// add all items from the package of the tag itself.
		if (tag.path === filter || filter === "all") {
			pkgs.push(...toAdd.map((p) => addTag(p, tag.path)));
			return;
		}

		// add all other packages, but filter out unused types, and
		// remove the services, not needed
		pkgs.push(
			...toAdd.map((p) => {
				const typs = {};
				Object.keys(p.types).forEach((typ) => {
					if (!used.has(p.id + "." + typ)) {
						return;
					}
					typs[typ] = p.types[typ];
				});
				p.types = typs;
				p.services = [];
				return addTag(p, tag.path);
			}),
		);
		return;
	});

	return {
		body: {
			pages,
			pkgs,
		},
	};
}

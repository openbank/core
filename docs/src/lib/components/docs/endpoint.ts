import type {
	Endpoint,
	Field,
	Message,
	Package,
	Packages,
} from "$lib/packages";

const urlRe = /\/{([^}]+)}/g;

export default class EndpointResolver {
	private readonly pkgs: Packages;
	private readonly pkg: Package;
	private readonly endpoint: Endpoint;

	private readonly reqFields: Field[] = [];

	constructor(pkgs: Packages, pkg: Package, endpoint: Endpoint) {
		this.pkgs = pkgs;
		this.pkg = pkg;
		this.endpoint = endpoint;

		if (this.endpoint.request) {
			this.reqFields = this.endpoint.request.fields;
		}
	}

	public params(): Field[][] {
		if (this.endpoint.body_field === "*") {
			return [[], []];
		}

		const matches = Array.from(this.endpoint.path.matchAll(urlRe)).map(
			(m) => m[1],
		);

		const pathParams = this.reqFields.filter((f) => matches.includes(f.name));
		const queryParams = this.reqFields
			.filter((f) => !matches.includes(f.name))
			.filter((f) => this.endpoint.body_field !== f.name);

		return [pathParams, queryParams];
	}

	public body(): { body: Message; link: string } {
		const bf = this.endpoint.body_field;
		if (bf === "") {
			return { body: null, link: "" };
		}
		if (bf === "*") {
			return { body: this.endpoint.request, link: "" };
		}

		const field = this.reqFields.find((f) => f.name === bf);
		if (field.type.type !== "ref") {
			throw new Error(
				"expected reference for body field, got " + field.type.type + "instead",
			);
		}

		const rs = this.pkgs.resolve(field.type);
		if (rs.type.type !== "message") {
			throw new Error(
				"expected message for body field, got " + rs.type.type + "instead",
			);
		}

		let link = "";
		if (this.pkg.id !== rs.package.id) {
			link = "/" + rs.package.tag;
		}

		return {
			body: rs.type,
			link: link + `#${rs.name}`,
		};
	}

	public response(): { resp: Message; link: string } {
		const r = this.endpoint.response;
		if (!r) {
			return { resp: null, link: "" };
		}
		if (r.type === "message") {
			return { resp: r, link: `#${this.pkg.id}.${r.name}` };
		}
		const rs = this.pkgs.resolve(r);
		if (rs.type.type !== "message") {
			throw new Error(
				"expected message for body field, got " + rs.type.type + "instead",
			);
		}

		let link = "";
		if (this.pkg.id !== rs.package.id) {
			link = "/" + rs.package.tag;
		}
		return {
			resp: rs.type,
			link: link + `#${rs.name}`,
		};
	}
}

<script context="module" lang="ts">
	import type { LoadInput, LoadOutput } from "@sveltejs/kit";

	export async function load({
		params,
		fetch,
	}: LoadInput): Promise<LoadOutput> {
		const rList: Response = await fetch(`/api/tags_?tag=${params.tag}`);
		const list = await rList.json();

		const rTag = await fetch(`/api/${params.tag}`);
		const tag = await rTag.json();

		if (rTag.ok && rList.ok) {
			return {
				props: {
					pages: list.pages,
					pkgs: list.pkgs,
					tag,
				},
			};
		}

		return {
			status: rTag.status,
			error: new Error(`Could not load page`),
		};
	}
</script>

<script lang="ts">
	import { Package, Packages, Page, Tag } from "$lib/packages";
	import Main from "$lib/components/main/Main.svelte";

	export let pages: Page[];
	export let tag: Tag;
	export let pkgs: Package[];
	const packages = new Packages(pkgs);
</script>

<Main {pages} {tag} {packages} docs />

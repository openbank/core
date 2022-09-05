<script lang="ts" context="module">
	import type { LoadInput, LoadOutput } from "@sveltejs/kit";

	export async function load({ fetch }: LoadInput): Promise<LoadOutput> {
		const rList = await fetch(`/api/tags_`);
		const list = await rList.json();

		const rTag = await fetch(`/api/default`);
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
	import Main from "$lib/components/main/Main.svelte";
	import { Package, Packages, Page, Tag } from "$lib/packages";

	export let pages: Page[];
	export let tag: Tag;
	export let pkgs: Package[];
	const packages = new Packages(pkgs);
</script>

<Main {pages} {tag} {packages} />

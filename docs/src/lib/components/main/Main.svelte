<script lang="ts">
	import Sidebar from "$lib/components/sidebar/Sidebar.svelte";
	import Endpoints from "$lib/components/docs/Endpoints.svelte";
	import Overview from "$lib/components/overview/Overview.svelte";
	import Types from "$lib/components/docs/Types.svelte";
	import MD from "$lib/components/md/MD.svelte";

	import type { Packages, Page, Tag } from "$lib/packages";

	export let pages: Page[];
	export let tag: Tag;
	export let packages: Packages;
	export let docs = false;

	let container: HTMLElement;
</script>

<div class="container relative mx-auto flex justify-center">
	<!-- sidebar -->
	<Sidebar {pages} current={tag} />
	<!-- main content -->
	<main
		bind:this={container}
		id="content"
		class="mt-32 mb-24 w-full px-4"
		class:withborder={!docs}
	>
		<!-- preamble -->
		{#if tag && tag.preamble}
			<MD inline={false} src={tag.preamble} />
		{/if}
		<!-- types and endpoints -->
		{#if docs}
			{#if tag && tag.packages}
				{#each tag.packages as pkg}
					<Endpoints pkgs={packages} {pkg} />
				{/each}
			{/if}
			{#if tag && tag.packages}
				{#each tag.packages as pkg}
					<Types pkgs={packages} {pkg} />
				{/each}
			{/if}
		{/if}
	</main>
	<!-- sidebar overview  -->
	<Overview {tag} {container} />
</div>

<style lang="postcss">
	.withborder {
		@apply border-r-4 border-gray-200;
	}
</style>

<script lang="ts">
	import MD from "$lib/components/md/MD.svelte";
	import type { Package, Packages } from "$lib/packages";
	import Endpoint from "./Endpoint.svelte";

	export let pkgs: Packages;
	export let pkg: Package;

	$: endpoints = pkg.services
		.map((s) => s.endpoints.length)
		.reduce((a, b) => a + b, 0);
</script>

{#if endpoints > 0}
	<h2 id={pkg.id}>{pkg.name.toUpperCase()}</h2>
	<p><MD src={pkg.description} /></p>

	{#each pkg.services as service}
		{#each service.endpoints as endpoint}
			<Endpoint {pkgs} {pkg} {service} {endpoint} />
		{/each}
	{/each}
{/if}

<style lang="postcss">
	h2::before {
		content: "";
		display: block;
		height: 100px;
		margin-top: -100px;
		visibility: hidden;
	}
</style>

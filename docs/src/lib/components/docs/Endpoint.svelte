<script lang="ts">
	import MD from "$lib/components/md/MD.svelte";
	import type { Endpoint, Package, Packages, Service } from "$lib/packages";
	import Table from "./table/Table.svelte";
	import EndpointResolver from "./endpoint";
	import Toggle from "../toggle/Toggle.svelte";

	export let pkgs: Packages;
	export let pkg: Package;
	export let service: Service;
	export let endpoint: Endpoint;

	const res = new EndpointResolver(pkgs, pkg, endpoint);
	const [pathParams, queryParams] = res.params();
	const { body, link: bodyLink } = res.body();
	const { resp, link: respLink } = res.response();

	const hasParams = pathParams.length !== 0 || queryParams.length !== 0;
</script>

<h3 id={`${pkg.id}.${service.name}.${endpoint.name}`} class="py-2 text-xl">
	{endpoint.name}
</h3>

<div class="not-prose">
	<code class={`${endpoint.method.toLowerCase()}`}>
		{endpoint.method}
		{endpoint.path}
	</code>
</div>
<p><em><MD src={endpoint.description + "."} /></em></p>

{#if hasParams}
	<strong>Parameters</strong>

	<Table>
		{#if pathParams && pathParams.length !== 0}
			<tr class="border-b-2 border-gray-200 bg-gray-200">
				<th class="w-full p-3" colspan="3">URL Parameters</th>
			</tr>
			{#each pathParams as field}
				<tr class="border-b-2 border-gray-200 hover:bg-white">
					<td class="py-3">{field.name}</td>
					<td class="py-3 px-1"><MD src={field.description} /></td>
					<td class="break-all py-3"><MD src={pkgs.typ(field.type)} /></td>
				</tr>
			{/each}
		{/if}

		{#if queryParams && queryParams.length !== 0}
			<tr class="border-b-2 border-gray-200 bg-gray-200">
				<th class="w-full p-3" colspan="3">Query Parameters</th>
			</tr>
			{#each queryParams as field}
				<tr class="border-b-2 border-gray-200 hover:bg-white">
					<td class="py-3">{field.name}</td>
					<td class="py-3 px-1"><MD src={field.description} /></td>
					<td class="break-all py-3"><MD src={pkgs.typ(field.type)} /></td>
				</tr>
			{/each}
		{/if}
	</Table>
{/if}

{#if body && body.fields.length}
	<div class="mt-8">
		<span class="mt-8">
			{#if bodyLink}
				<strong>Request Body</strong>
				(<a href={bodyLink}>{body.name}</a>)
			{:else}
				<strong>Body</strong> ({body.name})
			{/if}
		</span>

		<Table>
			{#each body.fields as field}
				<tr class="border-b-2 border-gray-200 hover:bg-white">
					<td class="py-3">{field.name}</td>
					<td class="py-3 px-1"><MD src={field.description} /></td>
					<td class="break-all py-3"><MD src={pkgs.typ(field.type)} /></td>
				</tr>
			{/each}
		</Table>
	</div>
{/if}

{#if resp && resp.fields.length}
	<div class="mt-8">
		<Toggle>
			<span slot="title">
				{#if respLink}
					<strong>Response Type</strong>
					(<a href={respLink}>{resp.name}</a>)
				{:else}
					<strong>Response Type</strong> ({resp.name})
				{/if}
			</span>

			<Table>
				{#each resp.fields as field}
					<tr class="border-b-2 border-gray-200 hover:bg-white">
						<td class="py-3">{field.name}</td>
						<td class="py-3 px-1"><MD src={field.description} /></td>
						<td class="break-all py-3"><MD src={pkgs.typ(field.type)} /></td>
					</tr>
				{/each}
			</Table>
		</Toggle>
	</div>
{/if}

<style lang="postcss">
	code {
		@apply px-1.5 py-1;
		@apply rounded;
		@apply text-sm;
	}

	code.get {
		@apply bg-blue-100 text-main-secondary;
	}

	code.post {
		@apply bg-green-100 text-green-500;
	}

	code.delete {
		@apply bg-red-100 text-red-400;
	}

	code.put {
		@apply bg-yellow-100 text-yellow-500;
	}

	code.patch {
		@apply bg-yellow-100 text-yellow-500;
	}
</style>

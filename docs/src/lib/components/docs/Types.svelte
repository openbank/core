<script lang="ts">
	import MD from "$lib/components/md/MD.svelte";
	import type { Package, Packages } from "$lib/packages";
	import Table from "./table/Table.svelte";

	export let pkgs: Packages;
	export let pkg: Package;

	const types = Object.keys(pkg.types).map((t) => {
		const typ = pkg.types[t];
		const data = {
			id: pkg.id + "." + t,
			type: typ.type,
			name: typ.name,
			description: typ.description,
			fields: undefined,
			values: undefined,
		};
		switch (typ.type) {
			case "message":
				data.fields = typ.fields;
				break;
			case "enum":
				data.values = typ.values;
		}
		return data;
	});
</script>

{#each types as t}
	<br />
	<h3 class="no-prose text-lg font-bold" id={t.id}>{t.name}</h3>
	<p><em><MD src={t.description + "."} /></em></p>
	<Table enumType={t.type === "enum"}>
		{#if t.type === "message"}
			{#each t.fields as field}
				<tr class="border-b-2 border-gray-200 hover:bg-white">
					<td class="py-3">{field.name}</td>
					<td class="py-3 px-1"><MD src={field.description} /></td>
					<td class="break-all py-3"><MD src={pkgs.typ(field.type)} /></td>
				</tr>
			{/each}
		{:else if t.type === "enum"}
			{#each t.values as value}
				<tr class="border-b-2 border-gray-200 hover:bg-white">
					<td class="py-3">{value.value}</td>
					<td class="py-3 px-1"><MD src={value.description} /></td>
				</tr>
			{/each}
		{/if}
	</Table>
{/each}

<style lang="postcss">
	h3::before {
		content: "";
		display: block;
		height: 100px;
		margin-top: -100px;
		visibility: hidden;
	}
</style>

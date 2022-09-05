<script lang="ts">
	export let enumType = false;
</script>

<table class="no-prose table-auto">
	{#if enumType}
		<tr class="border-b-2 border-gray-500">
			<th class="w-3/12 pb-4">Name</th>
			<th class="w-9/12 pb-4">Description</th>
		</tr>
	{:else}
		<tr class="border-b-2 border-gray-500">
			<th class="w-2/12 pb-4">Name</th>
			<th class="w-7/12 pb-4">Description</th>
			<th class="w-3/12 pb-4">Type</th>
		</tr>
	{/if}
	<slot />
</table>

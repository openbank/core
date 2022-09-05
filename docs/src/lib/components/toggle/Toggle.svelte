<script lang="ts">
	import { Accordion, AccordionItem } from "svelte-accessible-accordion";

	let expanded = true;
</script>

<Accordion>
	<AccordionItem bind:expanded>
		<div
			class="-ml-2 flex items-center rounded py-1 pl-2 hover:bg-blue-200"
			slot="title"
		>
			<slot name="title" />
			<div class="px-4">
				{#if expanded}
					<!-- prettier-ignore -->
					<svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 15l7-7 7 7" />
          </svg>
				{:else}
					<!-- prettier-ignore -->
					<svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" >
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />  
          </svg>
				{/if}
			</div>
		</div>
		<slot />
	</AccordionItem>
</Accordion>

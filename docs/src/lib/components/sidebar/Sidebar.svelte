<script lang="ts">
	import Logo from "$lib/components/logo/Logo.svelte";

	import type { Tag, Page } from "$lib/packages";

	export let pages: Page[] = [];
	export let current: Tag;

	$: cur = current.name;
</script>

<section
	class="sticky top-0 h-screen border-r-2 border-gray-200 pr-2 md:w-[13em]"
>
	<div class="relative flex max-h-screen flex-col overflow-y-auto pb-32">
		<a
			href="/"
			class="flex h-16 items-center
      justify-end border-r-2 border-transparent text-main-primary md:max-w-[14em]"
		>
			<Logo />
		</a>
		<!-- search btn -->
		<div class="flex w-full flex-col">
			<div
				class="mt-4 flex cursor-pointer items-center gap-2 rounded-lg bg-blue-200
        px-4 py-2 text-sm text-gray-500 hover:brightness-95"
			>
				<!-- prettier-ignore -->
				<svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
					<path fill-rule="evenodd" d="M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z" clip-rule="evenodd" />
				</svg>

				<p class="text-sm">Search</p>
				<span class="ml-auto mr-3 text-right text-xs text-gray-400">/</span>
			</div>
		</div>
		<!-- ref -->
		<div class="mx-4">
			<div class="head"><h2>Reference</h2></div>
			<div class="sec"><h2>Pages</h2></div>
			{#each pages as p}
				{#if p.name === "default"}
					<a href={`/`} class="subsec" class:active={cur === p.name}>Home</a>
				{:else}
					<a href={`/${p.path}`} class="subsec" class:active={cur === p.name}
						>{p.name}</a
					>
				{/if}
			{/each}
		</div>
	</div>
</section>

<style lang="postcss">
	.head {
		@apply pt-4 font-heading text-sm uppercase tracking-wide text-gray-800;
	}

	.sec {
		@apply ml-1 mt-4 font-heading text-xs uppercase tracking-widest text-gray-600;
	}

	.subsec {
		@apply my-1 ml-2 block py-1 pl-2 text-sm capitalize;
		@apply rounded text-main-black hover:brightness-95;
	}

	.active {
		@apply font-bold text-main-primary;
		@apply bg-blue-200 hover:bg-blue-100;
	}

	section {
		scrollbar-width: thin;
	}

	::-webkit-scrollbar {
		width: 4px;
	}

	::-webkit-scrollbar-thumb {
		@apply rounded bg-white transition;
	}

	section:hover ::-webkit-scrollbar-thumb {
		@apply bg-gray-500 transition;
	}
</style>

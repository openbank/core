<blockquote class="not-prose bquote flex gap-2 py-2 pr-4">
	<!-- prettier-ignore -->
	<svg xmlns="http://www.w3.org/2000/svg" class="mt-0.5 h-5 w-5 text-main-primary" fill="none" viewBox="0 0 24 24" stroke="currentColor">
		<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
	</svg>

	<div class="w-[calc(100%-40px)]">
		<slot />
	</div>
</blockquote>

<style lang="postcss">
	:global(.bquote a) {
		@apply underline;
	}

	:global(.bquote code) {
		@apply text-main-secondary;
	}
</style>

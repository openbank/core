<script>
	import { browser } from "$app/env";

	import Logo from "$lib/components/logo/Logo.svelte";

	let show = true;
	let threshold = 0;
	let last = browser ? window.scrollY : 0;
	const onScroll = () => {
		const current = window.scrollY;
		if (current >= last && current > 40) {
			show = false;
			threshold = current - 40;
		} else if (current < threshold) {
			show = true;
		}
		last = current;
	};
</script>

<svelte:window on:scroll={onScroll} />

<nav class:show class="fixed top-0 z-10 h-16 w-full bg-main-primary text-white">
	<div class="container mx-auto flex h-full items-center justify-between">
		<a
			href="/"
			class="flex h-8 w-full
      items-center justify-end border-r-2 border-transparent pr-6 md:max-w-[14em]"
		>
			<Logo />
		</a>
		<div
			class="flex w-full items-center
                gap-4 pl-4 md:max-w-[13em] xl:max-w-xs"
		>
			<button class="btn">Log In</button>
			<button class="btn btn-sec">Sign Up</button>
		</div>
	</div>
</nav>

<style lang="postcss">
	nav {
		transform: translateY(-100%);
		transition: 250ms;
	}
	nav.show {
		transform: translateY(0);
	}
</style>

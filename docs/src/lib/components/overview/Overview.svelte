<script lang="ts">
	import { browser } from "$app/env";

	import type { Tag } from "$lib/packages";
	import { onDestroy } from "svelte";
	import { writable } from "svelte/store";
	export let tag: Tag;
	export let container: HTMLElement;

	let show = true;
	let sec: HTMLElement;
	let active = "";
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

	$: pkgEndpoints = tag.packages
		.flatMap((p) =>
			p.services.map((s) => {
				let contains = false;
				const endpoints = s.endpoints.map((e) => {
					const id = `${p.id}.${s.name}.${e.name}`;
					if (id === active) {
						contains = true;
					}

					return {
						id: id + "-ov",
						href: "#" + id,
						text: e.name,
						active: id === active,
					};
				});

				return {
					name: s.name.replace(/Service$/, ""),
					active: contains,
					endpoints,
				};
			}),
		)
		.filter((pe) => pe.endpoints.length);

	$: pkgTypes = tag.packages
		.map((p) => {
			const ts = p.types;

			let contains = false;
			const types = Object.keys(p.types).map((t) => {
				const type = `${p.id}.${t}`;
				if (t === active) {
					contains = true;
				}

				return {
					id: `${type}-ov`,
					href: `#${type}`,
					text: ts[t].name,
					active: type === active,
				};
			});

			return {
				name: p.name,
				active: contains,
				types,
			};
		})
		.filter((pt) => pt.types.length);

	interface Entry {
		obs: IntersectionObserver;
		el: Element;
		showing: boolean;
	}

	interface Entries {
		[key: string]: Entry;
	}

	const activeList = writable<Entries>({});

	activeList.subscribe((entries) => {
		const ids = Object.keys(entries);
		const first = ids.find((id) => entries[id].showing);
		if (!first) {
			return;
		}
		active = entries[first].el.id;

		const ov = document.getElementById(active + "-ov");
		if (!ov) {
			console.error(active + "-ov", "not found");
			return;
		}

		sec.scroll({ top: ov.offsetTop - 196, behavior: "smooth" });
	});

	const clear = () => {
		Object.keys($activeList).forEach((id) => {
			$activeList[id].obs.disconnect();
		});
		activeList.set({});
	};

	const loadNav = (c: HTMLElement, tag: Tag) => {
		clear();
		const svcIDs = tag.packages.flatMap((p) =>
			p.services.flatMap((s) =>
				s.endpoints.map((e) => `[id="${p.id}.${s.name}.${e.name}"]`),
			),
		);

		const pkgIDs = pkgTypes
			.flatMap((pt) => pt.types.map((t) => t.href))
			.map((id) => `[id="${id.slice(1)}"]`);

		const ids = [...svcIDs, ...pkgIDs];
		if (!ids.length) {
			return;
		}
		c.querySelectorAll(ids.join(", ")).forEach((e) => {
			const obs = new IntersectionObserver(
				(e) => {
					activeList.update((entries) => {
						entries[e[0].target.id] = {
							obs,
							el: e[0].target,
							showing: e[0].isIntersecting,
						};
						return entries;
					});
				},
				{ threshold: 0.5, rootMargin: "-20px" },
			);

			obs.observe(e);
			activeList.update((entries) => {
				entries[e.id] = {
					obs,
					el: e,
					showing: false,
				};
				return entries;
			});
		});
	};

	$: {
		if (container) {
			loadNav(container, tag);
		}
	}

	onDestroy(() => {
		clear();
	});
</script>

<svelte:window on:scroll={onScroll} />

<section
	class:show
	class="sticky top-20 hidden max-h-[55vh] w-[14em] pr-4 xl:block"
>
	<div bind:this={sec} class="h-full w-full overflow-y-auto">
		<div class="-mt-3 flex flex-col gap-2 border-l-4 border-main-primary pl-2">
			{#if tag.packages.length}
				<div class="head">Endpoints</div>

				<!-- endpoints -->

				{#each pkgEndpoints as pe}
					<div class="sec" class:active={pe.active}>
						{pe.name}
					</div>
					{#each pe.endpoints as e}
						<a id={e.id} href={e.href} class="subsec" class:active={e.active}>
							{e.text}
						</a>
					{/each}
				{/each}

				<div class="head">Types</div>

				<!-- types -->

				{#each pkgTypes as pt}
					<div class="sec" class:active={pt.active}>
						{pt.name}
					</div>
					{#each pt.types as t}
						<a id={t.id} href={t.href} class="subsec" class:active={t.active}>
							{t.text}
						</a>
					{/each}
				{/each}
			{/if}
		</div>
	</div>
</section>

<style lang="postcss">
	section {
		transform: translate(0, calc(-4em));
		transition: transform 0.25s;
		overscroll-behavior: none;
	}

	section.show {
		transition: transform 0.25s;
		transform: none;
	}

	.head {
		@apply sticky top-0 bg-gray-100 pt-3 text-sm font-bold uppercase tracking-wide text-black;
	}

	.sec {
		@apply pt-3 text-xs font-bold uppercase tracking-widest text-gray-500;
	}

	.sec.active {
		@apply text-main-secondary;
	}

	.subsec {
		@apply -ml-2 overflow-x-hidden text-ellipsis pl-3 text-xs;
	}

	.subsec.active {
		@apply bg-blue-200 font-bold;
	}

	section {
		scrollbar-width: thin;
	}

	::-webkit-scrollbar {
		width: 4px;
	}

	::-webkit-scrollbar-thumb {
		@apply rounded bg-gray-500 transition;
	}
</style>

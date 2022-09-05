<script lang="ts">
	import SvelteMarkdown from "svelte-markdown";
	import Blockquote from "./Blockquote.svelte";

	export let src = "";
	export let inline = true;

	let source: string;
	$: {
		source = src;
		if (inline) {
			let first = src[0] || "";
			let rest = src.slice(1) || "";
			source = first.toUpperCase() + rest.replace("\n\n", "<br>");
		}
	}
</script>

<SvelteMarkdown
	{source}
	isInline={inline}
	renderers={{ blockquote: Blockquote }}
/>

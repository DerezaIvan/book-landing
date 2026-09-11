<script lang="ts">
	import type { Snippet } from 'svelte';

	type Props = {
		children: Snippet;
		index: string;
		tone?: 'milk' | 'paper' | 'forest' | 'night';
	};

	let { children, index, tone = 'paper' }: Props = $props();
</script>

<div class="chapter chapter--{tone}" data-book-page={index} style="--page-turn:0;">
	<div class="chapter__sheet">
		<span class="chapter__folio" aria-hidden="true">Dereza Stories · {index}</span>
		<div class="chapter__content">
			{@render children()}
		</div>
	</div>
</div>

<style lang="scss">
	@use '@/components/layout/Chapter/Chapter';
</style>

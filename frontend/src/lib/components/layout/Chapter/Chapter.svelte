<script lang="ts">
	import type { Snippet } from 'svelte';

	type Props = {
		children: Snippet;
		id: string;
		index: string;
		tone?: 'milk' | 'paper' | 'forest' | 'night';
	};

	let { children, id, index, tone = 'paper' }: Props = $props();
</script>

<section class="chapter chapter--{tone}" {id} aria-label={`Раздел ${index}`}>
	<div class="chapter__sheet">
		<span class="chapter__folio" aria-hidden="true">Dereza Stories · {index}</span>
		<div class="chapter__content">
			{@render children()}
		</div>
	</div>
</section>

<style lang="scss">
	@use '@/components/layout/Chapter/Chapter';
</style>

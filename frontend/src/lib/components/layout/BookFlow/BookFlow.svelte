<script lang="ts">
	import type { Snippet } from 'svelte';

	type Props = {
		children: Snippet;
	};

	let { children }: Props = $props();
</script>

<div class="book-flow">
	{@render children()}
</div>

<style lang="scss">
	@use '@/components/layout/BookFlow/BookFlow';
</style>

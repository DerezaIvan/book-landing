<script lang="ts">
	import { onMount, type Snippet } from 'svelte';

	type Props = {
		anchors: string[];
		children: Snippet;
		pages: number;
	};

	let { anchors, children, pages }: Props = $props();
	let flow = $state<HTMLElement>();

	onMount(() => {
		if (!flow) return;
		const flowElement = flow;
		const chapters = Array.from(flowElement.querySelectorAll<HTMLElement>('[data-book-page]'));

		const navigateToIndex = (index: number) => {
			const pageIndex = Math.min(pages - 1, Math.max(0, index));
			const flowTop = flowElement.getBoundingClientRect().top + window.scrollY;
			const target = flowTop + pageIndex * window.innerHeight * 1.8;

			window.scrollTo({ top: target, behavior: 'smooth' });
		};

		const navigateToPage = (event: Event) => {
			navigateToIndex(Number((event as CustomEvent<{ index: number }>).detail.index));
		};

		const navigateToHash = () => {
			const index = anchors.indexOf(window.location.hash.slice(1));
			if (index >= 0) navigateToIndex(index);
		};

		window.addEventListener('bookflow:navigate', navigateToPage);
		window.addEventListener('hashchange', navigateToHash);
		const hashFrame = window.requestAnimationFrame(navigateToHash);

		if (window.matchMedia('(prefers-reduced-motion: reduce)').matches) {
			flowElement.style.setProperty('--book-flow-reveal', '1');
			return () => {
				window.removeEventListener('bookflow:navigate', navigateToPage);
				window.removeEventListener('hashchange', navigateToHash);
				window.cancelAnimationFrame(hashFrame);
			};
		}

		const holdRatio = 0.66;
		let frame = 0;

		const update = () => {
			const bounds = flowElement.getBoundingClientRect();
			flowElement.style.setProperty('--book-flow-reveal', bounds.top <= 0 ? '1' : '0');
			const distance = Math.max(1, flowElement.offsetHeight - window.innerHeight);
			const travelled = Math.min(distance, Math.max(0, -bounds.top));
			const segment = distance / Math.max(1, pages - 1);
			const step = travelled / segment;

			chapters.forEach((chapter, index) => {
				const phase = Math.min(1, Math.max(0, step - index));
				const linearTurn = Math.min(1, Math.max(0, (phase - holdRatio) / (1 - holdRatio)));
				const turn = linearTurn * linearTurn * (3 - 2 * linearTurn);

				chapter.style.setProperty('--page-turn', String(turn));
				chapter.style.zIndex = String(pages - index);
				chapter.style.pointerEvents = step >= index && step < index + 1 ? 'auto' : 'none';
			});

			frame = 0;
		};

		const requestUpdate = () => {
			if (!frame) frame = window.requestAnimationFrame(update);
		};

		update();
		window.addEventListener('scroll', requestUpdate, { passive: true });
		window.addEventListener('resize', requestUpdate);

		return () => {
			window.removeEventListener('bookflow:navigate', navigateToPage);
			window.removeEventListener('hashchange', navigateToHash);
			window.removeEventListener('scroll', requestUpdate);
			window.removeEventListener('resize', requestUpdate);
			window.cancelAnimationFrame(hashFrame);
			if (frame) window.cancelAnimationFrame(frame);
		};
	});
</script>

<div
	class="book-flow"
	bind:this={flow}
	style={`--book-flow-height:${100 + Math.max(1, pages - 1) * 180}svh; --book-flow-reveal:0;`}
>
	{#each anchors as anchor, index (anchor)}
		<span
			class="book-flow__anchor"
			id={anchor}
			style={`--book-flow-anchor:${index * 180}svh;`}
			aria-hidden="true"
		></span>
	{/each}
	<div class="book-flow__stage">
		{@render children()}
	</div>
</div>

<style lang="scss">
	@use '@/components/layout/BookFlow/BookFlow';
</style>

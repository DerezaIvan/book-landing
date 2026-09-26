<script lang="ts">
	import { base } from '$app/paths';
	import { onMount } from 'svelte';

	let hero = $state<HTMLElement>();
	let progress = $state(0);
	let staticPresentation = $state(false);
	let pointerX = $state(0);
	let pointerY = $state(0);
	let openProgress = $derived(smoothstep(Math.min(1, Math.max(0, (progress - 0.12) / 0.72))));
	let spreadProgress = $derived(Math.min(1, Math.max(0, (openProgress - 0.28) / 0.48)));
	let leftPageProgress = $derived(Math.min(1, Math.max(0, (openProgress - 0.78) / 0.14)));
	let coverOpacity = $derived(Math.min(1, Math.max(0, (0.86 - openProgress) / 0.14)));
	let coverAngle = $derived(openProgress * -108);
	let stageX = $derived(-50 - (1 - openProgress) * 16);
	let stageTop = $derived(73 - openProgress * 1.5);
	let titleOpacity = $derived(
		staticPresentation ? 1 : 1 - smoothstep(Math.min(1, Math.max(0, (progress - 0.28) / 0.3)))
	);

	function smoothstep(value: number) {
		return value * value * (3 - 2 * value);
	}

	onMount(() => {
		if (
			window.matchMedia(
				'(prefers-reduced-motion: reduce), (max-width: 1024px), (hover: none), (pointer: coarse)'
			).matches
		) {
			staticPresentation = true;
			progress = 1;
			return;
		}

		let frame = 0;
		const update = () => {
			if (!hero) return;
			const rect = hero.getBoundingClientRect();
			progress = Math.min(
				1,
				Math.max(0, -rect.top / Math.max(1, rect.height - window.innerHeight))
			);
			frame = 0;
		};
		const onScroll = () => {
			if (!frame) frame = window.requestAnimationFrame(update);
		};

		update();
		window.addEventListener('scroll', onScroll, { passive: true });
		return () => {
			window.removeEventListener('scroll', onScroll);
			if (frame) window.cancelAnimationFrame(frame);
		};
	});

	function trackPointer(event: PointerEvent) {
		if (event.pointerType === 'touch') return;
		const bounds = (event.currentTarget as HTMLElement).getBoundingClientRect();
		pointerX = (event.clientX - bounds.left) / bounds.width - 0.5;
		pointerY = (event.clientY - bounds.top) / bounds.height - 0.5;
	}
</script>

<section class="hero" bind:this={hero} onpointermove={trackPointer} aria-labelledby="hero-title">
	<div
		class="hero__sticky"
		style={`--progress:${progress}; --open:${openProgress}; --spread:${spreadProgress}; --left-page:${leftPageProgress}; --cover-opacity:${coverOpacity}; --cover-angle:${coverAngle}deg; --stage-x:${stageX}%; --stage-top:${stageTop}%; --title-opacity:${titleOpacity}; --intro-opacity:${staticPresentation ? 1 : 1 - progress}; --tilt-x:${pointerY * -1.4}deg; --tilt-y:${pointerX * 2}deg;`}
	>
		<div class="hero__kicker">Семейная студия персональных книг</div>
		<h1 class="hero__title" id="hero-title">
			<span>Истории, в которых</span>
			<em>ваш ребёнок</em>
			<span>становится главным героем</span>
		</h1>

		<div class="book__stage" aria-hidden="true">
			<div class="book">
				<div class="book__shadow"></div>
				<div class="book__block" aria-hidden="true"></div>
				<div class="book__spread">
					<img
						class="book__artwork"
						src={`${base}/images/mock-story-garden.png`}
						alt=""
						width="1536"
						height="1024"
					/>
					<div class="book__page book__page--left">
						<span class="book__folio">Dereza Stories · глава 01</span>
					</div>
					<div class="book__page book__page--right">
						<span class="book__folio">Тайный сад</span>
						<p>
							Знакомый мир становится началом большого приключения, когда главным героем становится
							ваш ребёнок.
						</p>
					</div>
				</div>
				<div class="book__cover">
					<span class="book__monogram">D</span>
					<strong>История<br />только для тебя</strong>
					<small>Dereza Stories</small>
				</div>
			</div>
		</div>

		<div class="hero__note">Создаём вместе с вами<br />и бережно относимся к каждой детали</div>
		<a class="hero__cta" href="#contact">Обсудить книгу <span aria-hidden="true">↗</span></a>
		<div class="hero__scroll-cue"><span></span> Листайте — книга откроется</div>
	</div>
</section>

<style lang="scss">
	@use '@/components/sections/Hero/Hero';
</style>

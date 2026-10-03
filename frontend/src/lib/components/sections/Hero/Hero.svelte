<script lang="ts">
	import { base } from '$app/paths';
	import { magnetic } from '@/actions/magnetic';
	import { onMount } from 'svelte';

	let hero = $state<HTMLElement>();
	let progress = $state(0);
	let staticPresentation = $state(false);
	let pointerX = $state(0);
	let pointerY = $state(0);
	let openProgress = $derived(smoothstep(Math.min(1, Math.max(0, (progress - 0.12) / 0.72))));
	let spreadProgress = $derived(Math.min(1, Math.max(0, (openProgress - 0.28) / 0.48)));
	let leftPageProgress = $derived(Math.min(1, Math.max(0, (openProgress - 0.86) / 0.1)));
	let coverOpacity = $derived(Math.min(1, Math.max(0, (0.94 - openProgress) / 0.08)));
	let coverAngle = $derived(openProgress * -164);
	let stageX = $derived(-50 - (1 - openProgress) * 25);
	let stageTop = $derived(70 - openProgress * 1.5);
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
	<div class="hero__sky" aria-hidden="true">
		<div class="hero__aurora hero__aurora--one"></div>
		<div class="hero__aurora hero__aurora--two"></div>
		<div class="hero__stars hero__stars--near"></div>
		<div class="hero__stars hero__stars--far"></div>
	</div>
	<div
		class="hero__sticky"
		style={`--progress:${progress}; --open:${openProgress}; --spread:${spreadProgress}; --left-page:${leftPageProgress}; --cover-opacity:${coverOpacity}; --cover-angle:${coverAngle}deg; --stage-x:${stageX}%; --stage-top:${stageTop}%; --title-opacity:${titleOpacity}; --intro-opacity:${staticPresentation ? 1 : 1 - progress}; --tilt-x:${pointerY * -1.4}deg; --tilt-y:${pointerX * 2}deg; --px:${pointerX}; --py:${pointerY};`}
	>
		<div class="hero__fireflies" aria-hidden="true">
			<span></span><span></span><span></span><span></span><span></span>
		</div>
		<div class="hero__kicker">Семейная студия персональных книг</div>
		<h1 class="hero__title" id="hero-title">
			<span>Истории, в которых</span>
			<em>ваш ребёнок</em>
			<span>становится главным героем</span>
		</h1>

		<div class="book__stage" aria-hidden="true">
			<div class="book__glow"></div>
			<div class="book">
				<div class="book__shadow"></div>
				<div class="book__block" aria-hidden="true"></div>
				<div class="book__spread">
					<img
						class="book__artwork"
						src={`${base}/images/emilia-lemur.jpg`}
						alt=""
						width="1600"
						height="1067"
					/>
					<div class="book__page book__page--left">
						<span class="book__folio">Dereza Stories · глава 01</span>
					</div>
					<div class="book__page book__page--right">
						<span class="book__folio">Солнечная бухта</span>
						<p>
							Знакомый мир становится началом большого приключения, когда главным героем становится
							ваш ребёнок.
						</p>
					</div>
				</div>
				<div class="book__cover">
					<div class="book__cover-face book__cover-face--front">
						<span class="book__monogram">D</span>
						<strong>История<br />только для тебя</strong>
						<small>Dereza Stories</small>
					</div>
					<div class="book__cover-face book__cover-face--back" aria-hidden="true">
						<span></span>
					</div>
				</div>
			</div>
		</div>

		<div class="hero__note">Создаём вместе с вами<br />и бережно относимся к каждой детали</div>
		<a class="hero__cta" href="#contact" use:magnetic>
			Обсудить книгу
			<span class="hero__cta-icon" aria-hidden="true">
				<svg viewBox="0 0 24 24" width="16" height="16" fill="currentColor">
					<path
						d="M12 2c.6 4.8 2.4 7.4 4.7 8.9 1.4.9 3.1 1.6 5.3 1.9-4.8.6-7.4 2.4-8.9 4.7-.9 1.4-1.6 3.1-1.9 5.3-.6-4.8-2.4-7.4-4.7-8.9-1.4-.9-3.1-1.6-5.3-1.9 4.8-.6 7.4-2.4 8.9-4.7.9-1.4 1.6-3.1 1.9-5.3z"
					/>
				</svg>
			</span>
		</a>
		<div class="hero__scroll-cue"><span></span> Листайте — книга откроется</div>
	</div>
</section>

<style lang="scss">
	@use '@/components/sections/Hero/Hero';
</style>

<script lang="ts">
	import { onMount } from 'svelte';
	import type { Snippet } from 'svelte';

	type Props = {
		children?: Snippet;
	};

	let { children }: Props = $props();

	let hero = $state<HTMLElement>();
	let progress = $state(0);
	let pointerX = $state(0);
	let pointerY = $state(0);
	let openProgress = $derived(Math.min(1, Math.max(0, (progress - 0.12) / 0.72)));
	let spreadProgress = $derived(Math.min(1, Math.max(0, (openProgress - 0.28) / 0.48)));
	let leftPageProgress = $derived(Math.min(1, Math.max(0, (openProgress - 0.78) / 0.14)));
	let coverOpacity = $derived(Math.min(1, Math.max(0, (0.86 - openProgress) / 0.14)));
	let coverAngle = $derived(openProgress * -108);
	let stageX = $derived(-50 - (1 - openProgress) * 25);
	let stageTop = $derived(80 - openProgress * 18);
	let titleOpacity = $derived(Math.max(0, 1 - progress * 2.1));
	let diveProgress = $derived(Math.min(1, Math.max(0, (openProgress - 0.7) / 0.3)));
	let revealProgress = $derived(Math.min(1, Math.max(0, (diveProgress - 0.16) / 0.84)));

	onMount(() => {
		if (window.matchMedia('(prefers-reduced-motion: reduce)').matches) {
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
		style={`--progress:${progress}; --open:${openProgress}; --spread:${spreadProgress}; --left-page:${leftPageProgress}; --cover-opacity:${coverOpacity}; --cover-angle:${coverAngle}deg; --stage-x:${stageX}%; --stage-top:${stageTop}%; --title-opacity:${titleOpacity}; --dive:${diveProgress}; --reveal:${revealProgress}; --tilt-x:${pointerY * -2}deg; --tilt-y:${pointerX * 3}deg;`}
	>
		{#if children}
			<div class="hero__reveal">
				{@render children()}
			</div>
		{/if}
		<div class="hero__kicker">Семейная студия персональных книг</div>
		<h1 class="hero__title" id="hero-title">
			<span>Истории, в которых</span>
			<em>ваш ребёнок</em>
			<span>становится главным героем</span>
		</h1>

		<div class="book__stage" aria-label="Макет персональной книги, раскрывающийся при прокрутке">
			<div class="book">
				<div class="book__shadow"></div>
				<div class="book__block" aria-hidden="true"></div>
				<div class="book__spread">
					<div class="book__page book__page--left">
						<span class="book__folio">Почему это важно</span>
						<p class="book__preview-lead">
							Ребёнок узнаёт себя <em>на каждой странице.</em>
						</p>
					</div>
					<div class="book__page book__page--right">
						<span class="book__folio">Dereza Stories · 02</span>
						<p>
							В характере героя, знакомых семейных деталях и маленьких победах. Так появляется
							книга, к которой хочется возвращаться вместе - сегодня и много лет спустя.
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
		<div class="hero__scroll-cue"><span></span> Потяните вниз - книга откроется</div>
	</div>
</section>

<style lang="scss">
	@use '@/components/sections/Hero/Hero';
</style>

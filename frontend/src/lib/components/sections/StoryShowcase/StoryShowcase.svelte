<script lang="ts">
	import { base } from '$app/paths';
	import { parallax } from '@/actions/parallax';
	import { reveal } from '@/actions/reveal';
	import { spotlight } from '@/actions/spotlight';
	import { tilt } from '@/actions/tilt';

	let activeSlide = $state(0);
	let direction = $state<'next' | 'previous'>('next');
	let pointerStart = $state<number | null>(null);
	let lightboxImage = $state<{ src: string; alt: string } | null>(null);

	const scenes = [
		{
			src: `${base}/images/eva-walk-spread.jpg`,
			alt: 'Разворот книги: Ева гуляет с мамой, папой и щенком Каспером по берёзовой роще',
			label: 'Лесная прогулка',
			text: 'Знакомые семейные детали становятся началом истории, в которой ребёнок — главный герой.'
		},
		{
			src: `${base}/images/eva-squirrel-spread.jpg`,
			alt: 'Разворот книги: Ева угощает рыжую белочку ягодами рядом с пеньком',
			label: 'Новый друг',
			text: 'Любимые звери и интересы ребёнка вплетаются в сюжет и ведут приключение дальше.'
		},
		{
			src: `${base}/images/emilia-lemur.jpg`,
			alt: 'Иллюстрация: Эмилия угощает лемура печеньем на поляне над Солнечной бухтой',
			label: 'Пикник в бухте',
			text: 'Тёплые сцены, которые хочется перечитывать вместе — сегодня и много лет спустя.'
		}
	];

	function showSlide(index: number) {
		if (index === activeSlide) return;
		direction =
			index > activeSlide || (activeSlide === scenes.length - 1 && index === 0)
				? 'next'
				: 'previous';
		activeSlide = (index + scenes.length) % scenes.length;
	}

	function handlePointerDown(event: PointerEvent) {
		if (event.pointerType === 'mouse') return;
		pointerStart = event.clientX;
	}

	function handlePointerUp(event: PointerEvent) {
		if (pointerStart === null) return;
		const distance = event.clientX - pointerStart;
		pointerStart = null;
		if (Math.abs(distance) < 44) return;
		showSlide(activeSlide + (distance < 0 ? 1 : -1));
	}

	function closeLightbox() {
		lightboxImage = null;
	}

	function handleLightboxKeydown(event: KeyboardEvent) {
		if (event.key === 'Escape') closeLightbox();
	}

	$effect(() => {
		if (!lightboxImage) return;
		const previousOverflow = document.body.style.overflow;
		document.body.style.overflow = 'hidden';
		return () => {
			document.body.style.overflow = previousOverflow;
		};
	});
</script>

<svelte:window onkeydown={handleLightboxKeydown} />

<section class="story-showcase" aria-labelledby="story-showcase-title" use:spotlight>
	<header class="story-showcase__heading" data-reveal="words" use:reveal use:parallax={0.05}>
		<div>
			<span class="story-showcase__eyebrow">Загляните внутрь</span>
			<h2 id="story-showcase-title">Страницы настоящей истории</h2>
		</div>
		<p>
			Книга создаётся вокруг ребёнка — от первой встречи с героем до финала, который хочется
			перечитывать.
		</p>
	</header>

	<div class="story-showcase__reader" use:tilt={2.5}>
		{#key activeSlide}
			<article
				class="story-showcase__spread story-showcase__spread--{direction}"
				aria-label={`${scenes[activeSlide].label}, страница ${activeSlide + 1} из ${scenes.length}`}
			>
				<button
					class="story-showcase__image-button"
					type="button"
					onclick={() => (lightboxImage = scenes[activeSlide])}
					onpointerdown={handlePointerDown}
					onpointerup={handlePointerUp}
					onpointercancel={() => (pointerStart = null)}
					aria-label={`Увеличить иллюстрацию: ${scenes[activeSlide].alt}`}
				>
					<img
						src={scenes[activeSlide].src}
						alt={scenes[activeSlide].alt}
						width="1920"
						height="960"
					/>
					<span class="story-showcase__zoom" aria-hidden="true">Увеличить ↗</span>
				</button>
				<div class="story-showcase__copy" aria-live="polite" aria-atomic="true">
					<span class="story-showcase__page-number"
						>0{activeSlide + 1} <i>/ 0{scenes.length}</i></span
					>
					<span class="story-showcase__scene-label">{scenes[activeSlide].label}</span>
					<p>{scenes[activeSlide].text}</p>
				</div>
			</article>
		{/key}

		<div class="story-showcase__controls" role="group" aria-label="Управление просмотром истории">
			<button
				class="story-showcase__arrow"
				type="button"
				onclick={() => showSlide(activeSlide - 1)}
				aria-label="Предыдущая страница">←</button
			>
			<span class="story-showcase__current" aria-live="polite"
				>0{activeSlide + 1} / 0{scenes.length}</span
			>
			<button
				class="story-showcase__arrow"
				type="button"
				onclick={() => showSlide(activeSlide + 1)}
				aria-label="Следующая страница">→</button
			>
			<span class="story-showcase__hint">Листайте или нажимайте на стрелки</span>
		</div>
		<div class="story-showcase__progress" role="group" aria-label="Выбрать страницу">
			{#each scenes as scene, index (scene.label)}
				<button
					type="button"
					class:active={activeSlide === index}
					onclick={() => showSlide(index)}
					aria-label={`Открыть страницу ${index + 1}: ${scene.label}`}
					aria-current={activeSlide === index ? 'step' : undefined}
				>
					<img src={scene.src} alt="" width="280" height="152" loading="lazy" />
					<span class="story-showcase__scene-number">0{index + 1}</span>
					<strong>{scene.label}</strong>
				</button>
			{/each}
		</div>
	</div>
</section>

{#if lightboxImage}
	<div
		class="story-showcase__lightbox"
		role="dialog"
		aria-modal="true"
		aria-label="Просмотр иллюстрации"
	>
		<button
			class="story-showcase__lightbox-backdrop"
			type="button"
			onclick={closeLightbox}
			aria-label="Закрыть изображение"
		></button>
		<div class="story-showcase__lightbox-content">
			<img src={lightboxImage.src} alt={lightboxImage.alt} />
		</div>
		<button
			class="story-showcase__lightbox-close"
			type="button"
			onclick={closeLightbox}
			aria-label="Закрыть изображение"
		>
			<img src={`${base}/images/icon-close.svg`} alt="" aria-hidden="true" />
		</button>
	</div>
{/if}

<style lang="scss">
	@use '@/components/sections/StoryShowcase/StoryShowcase';
</style>

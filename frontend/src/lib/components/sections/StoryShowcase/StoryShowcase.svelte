<script lang="ts">
	import { base } from '$app/paths';
	import { reveal } from '@/actions/reveal';

	let activeSlide = $state(0);
	let direction = $state<'next' | 'previous'>('next');
	let pointerStart = $state<number | null>(null);
	let lightboxImage = $state<{ src: string; alt: string } | null>(null);

	const scenes = [
		{
			src: `${base}/images/misha-star-ball-cover.jpg`,
			alt: 'Миша и Зайчик Топ на обложке книги «Миша и звёздный мяч»',
			label: 'Обложка',
			text: 'Знакомство с героем и миром, в который захочется возвращаться.'
		},
		{
			src: `${base}/images/misha-star-ball-discovery.jpg`,
			alt: 'Миша и Зайчик Топ находят мяч с жёлтой звездой',
			label: 'Приключение',
			text: 'Любимый интерес ребёнка становится частью сюжета и отправной точкой приключения.'
		},
		{
			src: `${base}/images/misha-star-ball-finale.jpg`,
			alt: 'Миша и Зайчик Топ прощаются на закате',
			label: 'Финал истории',
			text: 'Тёплая последняя страница, которую можно перечитывать вместе.'
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

<section class="story-showcase" aria-labelledby="story-showcase-title">
	<header class="story-showcase__heading" data-reveal="words" use:reveal>
		<div>
			<span class="story-showcase__eyebrow">Загляните внутрь</span>
			<h2 id="story-showcase-title">Три страницы одной истории</h2>
		</div>
		<p>
			Книга создаётся вокруг ребёнка — от первой встречи с героем до финала, который хочется
			перечитывать.
		</p>
	</header>

	<div class="story-showcase__reader">
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
						width="1400"
						height="788"
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

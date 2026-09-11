<script lang="ts">
	import { base } from '$app/paths';
	import { reveal } from '@/actions/reveal';

	let activeSlide = $state(0);
	let lightboxImage = $state<{ src: string; alt: string } | null>(null);

	const examples = [
		{
			id: 'story-01',
			caption:
				'«Миша и звёздный мяч» - один из примеров. Герои, сюжет и иллюстрации каждой книги создаются вокруг вашего ребёнка и его интересов.',
			cards: [
				{
					position: 'cover',
					src: `${base}/images/misha-star-ball-cover.jpg`,
					alt: 'Обложка книги «Миша и звёздный мяч»',
					number: '01',
					label: 'Обложка'
				},
				{
					position: 'finale',
					src: `${base}/images/misha-star-ball-finale.jpg`,
					alt: 'Миша и Зайчик Топ прощаются на закате',
					number: '04',
					label: 'Финал истории'
				},
				{
					position: 'feature',
					src: `${base}/images/misha-star-ball-discovery.jpg`,
					alt: 'Миша и Зайчик Топ находят мяч с жёлтой звездой',
					number: '02',
					label: 'Начало приключения'
				}
			]
		},
		{
			id: 'story-02',
			caption:
				'Каждая история получает собственный визуальный ритм: от первого знакомства с героем до тёплого финала, который хочется перечитывать вместе.',
			cards: [
				{
					position: 'cover',
					src: `${base}/images/misha-star-ball-discovery.jpg`,
					alt: 'Пример сюжетного разворота персональной книги',
					number: '01',
					label: 'Завязка'
				},
				{
					position: 'finale',
					src: `${base}/images/misha-star-ball-cover.jpg`,
					alt: 'Пример оформления персональной книги',
					number: '02',
					label: 'Герои'
				},
				{
					position: 'feature',
					src: `${base}/images/misha-star-ball-finale.jpg`,
					alt: 'Пример финального разворота персональной книги',
					number: '03',
					label: 'Развязка'
				}
			]
		},
		{
			id: 'story-03',
			caption:
				'Это временный пример наполнения галереи. Позже каждый такой слайд можно заменить отдельной готовой книгой без изменения композиции секции.',
			cards: [
				{
					position: 'cover',
					src: `${base}/images/misha-star-ball-finale.jpg`,
					alt: 'Пример атмосферы персональной истории',
					number: '01',
					label: 'Атмосфера'
				},
				{
					position: 'finale',
					src: `${base}/images/misha-star-ball-discovery.jpg`,
					alt: 'Пример приключения в персональной книге',
					number: '02',
					label: 'Приключение'
				},
				{
					position: 'feature',
					src: `${base}/images/misha-star-ball-cover.jpg`,
					alt: 'Пример главной иллюстрации персональной книги',
					number: '03',
					label: 'Главный герой'
				}
			]
		}
	];

	function showSlide(index: number) {
		activeSlide = (index + examples.length) % examples.length;
	}

	function openLightbox(image: { src: string; alt: string }) {
		lightboxImage = image;
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
		<span class="story-showcase__eyebrow">Пример персональной истории</span>
		<h2 id="story-showcase-title">Один разворот - и история оживает</h2>
	</header>
	<div class="story-showcase__slider">
		<div class="story-showcase__track" style={`--active-slide:${activeSlide}`}>
			{#each examples as example, slideIndex (example.id)}
				<article class="story-showcase__slide" aria-label={`Пример ${slideIndex + 1} из ${examples.length}`}>
					<div class="story-showcase__gallery">
						<div class="story-showcase__stack">
							{#each example.cards.slice(0, 2) as card (card.position)}
								<figure class="story-showcase__card story-showcase__card--{card.position}">
									<button
										class="story-showcase__image-button"
										type="button"
										onclick={() => openLightbox(card)}
										aria-label={`Открыть изображение: ${card.alt}`}
									>
										<img src={card.src} alt={card.alt} width="1400" height="788" loading="lazy" />
									</button>
									<figcaption><span>{card.number}</span> {card.label}</figcaption>
								</figure>
							{/each}
						</div>
						<figure class="story-showcase__card story-showcase__card--feature">
							<button
								class="story-showcase__image-button"
								type="button"
								onclick={() => openLightbox(example.cards[2])}
								aria-label={`Открыть изображение: ${example.cards[2].alt}`}
							>
								<img
									src={example.cards[2].src}
									alt={example.cards[2].alt}
									width="1400"
									height="788"
									loading="lazy"
								/>
							</button>
							<figcaption><span>{example.cards[2].number}</span> {example.cards[2].label}</figcaption>
						</figure>
					</div>
					<p class="story-showcase__caption">{example.caption}</p>
				</article>
			{/each}
		</div>

		<div class="story-showcase__controls">
			<div class="story-showcase__dots" aria-label="Выбор примера">
				{#each examples as example, index (example.id)}
					<button
						type="button"
						class:active={activeSlide === index}
						onclick={() => showSlide(index)}
						aria-label={`Показать пример ${index + 1}`}
						aria-current={activeSlide === index ? 'true' : undefined}
					></button>
				{/each}
			</div>
			<span class="story-showcase__counter">0{activeSlide + 1} / 0{examples.length}</span>
			<div class="story-showcase__arrows">
				<button type="button" onclick={() => showSlide(activeSlide - 1)} aria-label="Предыдущий пример">←</button>
				<button type="button" onclick={() => showSlide(activeSlide + 1)} aria-label="Следующий пример">→</button>
			</div>
		</div>
	</div>
</section>

{#if lightboxImage}
	<div class="story-showcase__lightbox" role="dialog" aria-modal="true" aria-label="Просмотр иллюстрации">
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

<script lang="ts">
	import { base } from '$app/paths';
	import { reveal } from '@/actions/reveal';

	let childName = $state('');
	let activeScene = $state<'garden' | 'pond'>('garden');

	const scenes = {
		garden: {
			image: `${base}/images/mock-story-garden.png`,
			alt: 'Иллюстрация тайного сада для примера персональной книги',
			name: 'Тайный сад',
			caption: 'В один удивительный день {name} нашёл дорогу в тайный сад…'
		},
		pond: {
			image: `${base}/images/mock-story-pond.png`,
			alt: 'Иллюстрация пруда для примера персональной книги',
			name: 'Огни у пруда',
			caption: 'Когда у пруда зажглись огни, {name} понял: приключение начинается…'
		}
	};

	let scene = $derived(scenes[activeScene]);
	let heroName = $derived(childName.trim() || 'ваш герой');
	let sampleCaption = $derived(scene.caption.replace('{name}', heroName));
</script>

<section class="manifesto">
	<div class="manifesto__inner">
		<div class="manifesto__story">
			<div class="manifesto__copy" data-reveal="words" use:reveal>
				<span class="manifesto__eyebrow">Почему это важно</span>
				<p class="manifesto__lead">Ребёнок узнаёт себя <span>на каждой странице.</span></p>
				<p class="manifesto__detail">
					В характере героя, знакомых семейных деталях и маленьких победах. Так появляется книга, к
					которой хочется возвращаться вместе — сегодня и много лет спустя.
				</p>
				<a class="manifesto__link" href="#contact"
					>Обсудить свою историю <span aria-hidden="true">↗</span></a
				>
			</div>

			<div class="manifesto__preview">
				<div class="manifesto__preview-image">
					{#key activeScene}
						<img src={scene.image} alt={scene.alt} width="1536" height="1024" />
					{/key}
					<div class="manifesto__preview-cover" aria-live="polite" aria-atomic="true">
						<span>Маленький фрагмент книги</span>
						<strong>{scene.name}</strong>
						<p>{sampleCaption}</p>
					</div>
					<span class="manifesto__preview-index" aria-hidden="true">D · 01</span>
				</div>

				<div class="manifesto__preview-controls">
					<label class="manifesto__name-field">
						<span>Имя главного героя</span>
						<input
							bind:value={childName}
							maxlength="18"
							placeholder="Например, Миша"
							aria-label="Введите имя героя для примера"
						/>
					</label>
					<div class="manifesto__scene-picker" role="group" aria-label="Выберите сюжет примера">
						<button
							type="button"
							class:active={activeScene === 'garden'}
							aria-pressed={activeScene === 'garden'}
							onclick={() => (activeScene = 'garden')}>Тайный сад</button
						>
						<button
							type="button"
							class:active={activeScene === 'pond'}
							aria-pressed={activeScene === 'pond'}
							onclick={() => (activeScene = 'pond')}>У пруда</button
						>
					</div>
				</div>
				<p class="manifesto__preview-note">
					Это пример: настоящую историю мы придумаем вместе с вами.
				</p>
			</div>
		</div>
	</div>
</section>

<style lang="scss">
	@use '@/components/sections/Manifesto/Manifesto';
</style>

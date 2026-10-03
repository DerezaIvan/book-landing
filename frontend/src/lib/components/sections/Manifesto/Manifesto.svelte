<script lang="ts">
	import { base } from '$app/paths';
	import { parallax } from '@/actions/parallax';
	import { reveal } from '@/actions/reveal';
	import { tilt } from '@/actions/tilt';

	let childName = $state('');
	let activeScene = $state<'garden' | 'pond'>('garden');

	const scenes = {
		garden: {
			image: `${base}/images/emilia-lemur.jpg`,
			alt: 'Иллюстрация: Эмилия и лемур делятся печеньем на поляне над бухтой',
			name: 'Солнечная бухта',
			caption: 'В один удивительный день {name} отправился на пикник в Солнечную бухту…'
		},
		pond: {
			image: `${base}/images/emilia-blanket.jpg`,
			alt: 'Иллюстрация: Эмилия расстилает плед для пикника рядом с гепардом и зайчиком',
			name: 'Плед для друзей',
			caption: 'Когда плед был расстелен, {name} понял: приключение начинается…'
		}
	};

	let scene = $derived(scenes[activeScene]);
	let heroName = $derived(childName.trim() || 'ваш герой');
	let sampleCaption = $derived(scene.caption.replace('{name}', heroName));
</script>

<section class="manifesto">
	<div class="manifesto__inner">
		<div class="manifesto__story">
			<div class="manifesto__copy" data-reveal="words" use:reveal use:parallax={0.04}>
				<span class="manifesto__eyebrow">Почему это важно</span>
				<p class="manifesto__lead">Ребёнок узнаёт себя <span>на каждой странице.</span></p>
				<p class="manifesto__detail">
					В характере героя, знакомых семейных деталях и маленьких победах. Так появляется книга, к
					которой хочется возвращаться вместе — сегодня и много лет спустя.
				</p>
				<a class="manifesto__link" href="#contact"
					>Обсудить свою историю <span aria-hidden="true">✦</span></a
				>
			</div>

			<div class="manifesto__preview" use:parallax={0.07} use:tilt={4}>
				<div class="manifesto__preview-image">
					{#key activeScene}
						<img src={scene.image} alt={scene.alt} width="1600" height="1067" />
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
							onclick={() => (activeScene = 'garden')}>Солнечная бухта</button
						>
						<button
							type="button"
							class:active={activeScene === 'pond'}
							aria-pressed={activeScene === 'pond'}
							onclick={() => (activeScene = 'pond')}>Плед для друзей</button
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

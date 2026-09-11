<script lang="ts">
	import { base } from '$app/paths';
	import { reveal } from '@/actions/reveal';
</script>

<section class="studio">
	<figure class="studio__portrait is-visible" data-reveal="curtain" use:reveal>
		<img
			src={`${base}/images/founders.jpg`}
			alt="Иван и Алина Дереза, основатели семейной студии"
			width="1122"
			height="1402"
			loading="lazy"
		/>
		<figcaption>Иван и Алина Дереза</figcaption>
	</figure>
	<div class="studio__copy" data-reveal="words" use:reveal>
		<span class="studio__eyebrow">Семейная студия</span>
		<h2>Мы - Иван и Алина Дереза.</h2>
		<p class="studio__lead">
			Создаём истории как родители и как люди, которые верят: ребёнку важно увидеть себя смелым,
			любопытным и любимым.
		</p>
		<p>
			В каждой книге соединяем личные детали семьи, авторский сюжет и иллюстрации - без конвейера и
			без случайных историй.
		</p>
		<div class="studio__signature">Иван & Алина</div>
	</div>
</section>

<style lang="scss">
	@use '@/components/sections/Studio/Studio';
</style>

<script lang="ts">
	import { base } from '$app/paths';
	import { parallax } from '@/actions/parallax';
	import { reveal } from '@/actions/reveal';
	import { tilt } from '@/actions/tilt';
</script>

<section class="formats">
	<div class="formats__title" data-reveal="slide" use:reveal use:parallax={0.05}>
		<span class="formats__eyebrow">Печатная и цифровая версии</span>
		<h2>История в том<br />формате, который<br />останется с вами</h2>
	</div>
	<div class="formats__showcase">
		<div class="formats__stack" data-reveal="stack" use:reveal>
			<article class="formats__card formats__card--print" use:parallax={0.06} use:tilt={3}>
				<div class="formats__media">
					<img
						src={`${base}/images/format-print-book.jpg`}
						alt="Ребёнок держит раскрытую печатную книгу"
						width="586"
						height="595"
						loading="lazy"
					/>
				</div>
				<div class="formats__content">
					<span class="formats__index">01 / твёрдый переплёт</span>
					<strong>Книга, которую можно<br />поставить на полку</strong>
					<small>Плотная бумага · подарочная упаковка</small>
				</div>
			</article>
			<article class="formats__card formats__card--digital" use:parallax={-0.06} use:tilt={3}>
				<div class="formats__media">
					<img
						src={`${base}/images/format-digital-story.jpg`}
						alt="Иллюстрированный разворот цифровой книги на экране"
						width="1280"
						height="640"
						loading="lazy"
					/>
				</div>
				<div class="formats__content">
					<span class="formats__index">02 / цифровая версия</span>
					<strong>История, которая<br />всегда под рукой</strong>
					<small>PDF · для телефона, планшета и печати</small>
				</div>
			</article>
		</div>
	</div>
</section>

<style lang="scss">
	@use '@/components/sections/Formats/Formats';
</style>

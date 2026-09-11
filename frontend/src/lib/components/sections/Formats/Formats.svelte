<script lang="ts">
	import { base } from '$app/paths';
	import { reveal } from '@/actions/reveal';
</script>

<section class="formats">
	<div class="formats__title" data-reveal="slide" use:reveal>
		<span class="formats__eyebrow">Печатная и цифровая версии</span>
		<h2>История в том<br />формате, который<br />останется с вами</h2>
	</div>
	<div class="formats__showcase">
		<p class="formats__hint">
			<span class="formats__hint-desktop">Наведите на книгу</span>
			<span class="formats__hint-touch">Коснитесь книги</span>
			<i aria-hidden="true">↓</i>
		</p>
		<div class="formats__stack" data-reveal="stack" use:reveal>
			<article class="formats__card formats__card--print" tabindex="0">
			<div class="formats__media" aria-hidden="true">
				<img src={`${base}/images/format-print-book.jpg`} alt="" width="586" height="595" />
				<span class="formats__media-caption">Настоящая книга · в ваших руках</span>
			</div>
			<div class="formats__content">
				<span class="formats__index">01 / твёрдый переплёт</span>
				<strong>Книга, которую можно<br />поставить на полку</strong>
				<small>Плотная бумага · подарочная упаковка</small>
			</div>
			</article>
			<article class="formats__card formats__card--digital" tabindex="0">
			<div class="formats__media" aria-hidden="true">
				<img src={`${base}/images/format-digital-story.jpg`} alt="" width="1280" height="640" />
				<span class="formats__media-caption">Целый разворот · на любом экране</span>
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

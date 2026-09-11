<script lang="ts">
	import { reveal } from '@/actions/reveal';

	const chapters = [
		{
			number: '01',
			title: 'Знакомимся',
			text: 'Вы рассказываете нам о ребёнке: что он любит, о чём мечтает и кого хочет взять в приключение.',
			mark: '○'
		},
		{
			number: '02',
			title: 'Пишем историю',
			text: 'Создаём сюжет, где знакомые черты и важные семейные детали становятся частью настоящей сказки.',
			mark: '✦'
		},
		{
			number: '03',
			title: 'Собираем книгу',
			text: 'Иллюстрируем, верстаем и готовим историю к печати - как вещь, которую хочется хранить годами.',
			mark: '□'
		}
	];
</script>

<section class="process">
	<div class="process__heading" data-reveal="split" use:reveal>
		<span class="process__eyebrow">Как рождается история</span>
		<h2>От первой детали<br />до книги, <em>которая останется</em></h2>
	</div>
	<div class="process__chapters">
		{#each chapters as chapter, index (chapter.number)}
			<article class="process__chapter" data-reveal="chapter" use:reveal style={`--order:${index}`}>
				<span class="process__number">{chapter.number}</span>
				<h3>{chapter.title}</h3>
				<p>{chapter.text}</p>
				<div class="process__mark" aria-hidden="true">{chapter.mark}</div>
			</article>
		{/each}
	</div>
</section>

<style lang="scss">
	@use '@/components/sections/Process/Process';
</style>

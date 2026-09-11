<script lang="ts">
	import { reveal } from '@/actions/reveal';
</script>

<section class="manifesto">
	<div class="manifesto__inner">
		<div class="manifesto__heading" data-reveal="words" use:reveal>
			<span class="manifesto__eyebrow">Почему это важно</span>
			<p class="manifesto__lead">Ребёнок узнаёт себя <span>на каждой странице.</span></p>
		</div>

		<p class="manifesto__detail" data-reveal="words" use:reveal>
			В характере героя, знакомых семейных деталях и маленьких победах. Так появляется книга, к
			которой хочется возвращаться вместе - сегодня и много лет спустя.
		</p>
	</div>
</section>

<style lang="scss">
	@use '@/components/sections/Manifesto/Manifesto';
</style>

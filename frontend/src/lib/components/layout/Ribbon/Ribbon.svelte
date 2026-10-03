<script lang="ts">
	const phrases = [
		'Ваш ребёнок — главный герой',
		'Авторский сюжет',
		'Иллюстрации вручную',
		'Семейная студия',
		'Книга, которую хранят годами'
	];
</script>

<div class="ribbon" aria-hidden="true">
	<div class="ribbon__track">
		{#each [...phrases, ...phrases] as phrase, index (index)}
			<span class="ribbon__item">{phrase} <i>✦</i></span>
		{/each}
	</div>
</div>

<style lang="scss">
	@use '@/styles/variables' as v;

	.ribbon {
		position: relative;
		z-index: 4;
		padding: 1.05rem 0;
		overflow: hidden;
		color: #f7e8ce;
		white-space: nowrap;
		background: linear-gradient(90deg, v.$color-berry-deep, v.$color-berry 50%, v.$color-berry-deep);
		border-block: 1px solid rgba(240, 205, 143, 0.35);
		box-shadow: 0 1rem 3rem rgba(12, 17, 15, 0.25);

		&__track {
			display: inline-flex;
			animation: marquee 36s linear infinite;
		}

		&__item {
			display: inline-flex;
			gap: 1.6rem;
			align-items: center;
			padding-right: 1.6rem;
			font-family: v.$font-serif;
			font-size: clamp(1rem, 1.5vw, 1.35rem);
			letter-spacing: 0.04em;

			i {
				color: v.$color-gold;
				font-size: 0.75em;
				font-style: normal;
			}
		}

		&:hover .ribbon__track {
			animation-play-state: paused;
		}
	}

	@media (prefers-reduced-motion: reduce) {
		.ribbon__track {
			animation: none;
		}
	}
</style>

<script lang="ts">
	let menuOpen = $state(false);

	function navigateToPage(event: MouseEvent, index: number) {
		event.preventDefault();
		menuOpen = false;
		window.dispatchEvent(new CustomEvent('bookflow:navigate', { detail: { index } }));
	}
</script>

<header class="site-header" class:site-header--menu-open={menuOpen}>
	<a class="site-header__brand" href="#top" aria-label="Dereza Stories - на главную">
		<span class="site-header__brand-mark">DS</span>
		<span>Dereza Stories</span>
	</a>
	<button
		class="site-header__menu-button"
		type="button"
		aria-expanded={menuOpen}
		aria-label={menuOpen ? 'Закрыть меню' : 'Открыть меню'}
		onclick={() => (menuOpen = !menuOpen)}
	>
		<span class="site-header__menu-label site-header__menu-label--open" aria-hidden="true">Меню</span>
		<span class="site-header__menu-label site-header__menu-label--close" aria-hidden="true">Закрыть</span>
	</button>
	<nav class="site-header__nav" aria-label="Основная навигация">
		<a href="#approach" onclick={(event) => navigateToPage(event, 2)}>Как мы создаём</a>
		<a href="#stories" onclick={(event) => navigateToPage(event, 0)}>Истории</a>
		<a href="#studio" onclick={(event) => navigateToPage(event, 4)}>О студии</a>
	</nav>
	<a class="site-header__cta" href="#contact" onclick={(event) => navigateToPage(event, 5)}>
		Обсудить книгу <span>↗</span>
	</a>
</header>

<style lang="scss">
	@use '@/components/layout/Header/Header';
</style>

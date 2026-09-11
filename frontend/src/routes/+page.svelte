<script lang="ts">
	import Footer from '@/components/layout/Footer/Footer.svelte';
	import Header from '@/components/layout/Header/Header.svelte';
	import Chapter from '@/components/layout/Chapter/Chapter.svelte';
	import BookFlow from '@/components/layout/BookFlow/BookFlow.svelte';
	import Contact from '@/components/sections/Contact/Contact.svelte';
	import Formats from '@/components/sections/Formats/Formats.svelte';
	import Hero from '@/components/sections/Hero/Hero.svelte';
	import Manifesto from '@/components/sections/Manifesto/Manifesto.svelte';
	import Process from '@/components/sections/Process/Process.svelte';
	import StoryShowcase from '@/components/sections/StoryShowcase/StoryShowcase.svelte';
	import Studio from '@/components/sections/Studio/Studio.svelte';
</script>

<svelte:head>
	<title>Персональные детские книги | Dereza Stories</title>
	<meta
		name="description"
		content="Создаём персональные детские книги, в которых ваш ребёнок становится главным героем. Авторский сюжет, иллюстрации и бережное семейное производство."
	/>
	<meta name="robots" content="index, follow, max-image-preview:large, max-snippet:-1, max-video-preview:-1" />
	<link rel="canonical" href="https://derezaivan.github.io/book-landing/" />
	<link rel="manifest" href="/book-landing/site.webmanifest" />
	<meta name="theme-color" content="#273129" />
	<meta property="og:type" content="website" />
	<meta property="og:locale" content="ru_RU" />
	<meta property="og:site_name" content="Dereza Stories" />
	<meta property="og:title" content="Персональные детские книги | Dereza Stories" />
	<meta property="og:description" content="Истории, в которых ваш ребёнок становится главным героем. Авторский сюжет и иллюстрации создаются вокруг вашей семьи." />
	<meta property="og:url" content="https://derezaivan.github.io/book-landing/" />
	<meta property="og:image" content="https://derezaivan.github.io/book-landing/images/misha-star-ball-cover.jpg" />
	<meta property="og:image:alt" content="Пример персональной детской книги Dereza Stories" />
	<meta name="twitter:card" content="summary_large_image" />
	<meta name="twitter:title" content="Персональные детские книги | Dereza Stories" />
	<meta name="twitter:description" content="Истории, в которых ваш ребёнок становится главным героем." />
	<meta name="twitter:image" content="https://derezaivan.github.io/book-landing/images/misha-star-ball-cover.jpg" />
	<script type="application/ld+json">
		{JSON.stringify({
			'@context': 'https://schema.org',
			'@graph': [
				{
					'@type': 'Organization',
					'@id': 'https://derezaivan.github.io/book-landing/#organization',
					name: 'Dereza Stories',
					url: 'https://derezaivan.github.io/book-landing/',
					description: 'Семейная студия персональных детских книг',
					image: 'https://derezaivan.github.io/book-landing/images/founders.jpg',
					founder: [
						{ '@type': 'Person', name: 'Иван Дереза' },
						{ '@type': 'Person', name: 'Алина Дереза' }
					]
				},
				{
					'@type': 'WebSite',
					'@id': 'https://derezaivan.github.io/book-landing/#website',
					url: 'https://derezaivan.github.io/book-landing/',
					name: 'Dereza Stories',
					inLanguage: 'ru-RU',
					publisher: { '@id': 'https://derezaivan.github.io/book-landing/#organization' }
				},
				{
					'@type': 'Service',
					'@id': 'https://derezaivan.github.io/book-landing/#service',
					name: 'Создание персональных детских книг',
					serviceType: 'Персональные детские книги',
					url: 'https://derezaivan.github.io/book-landing/',
					description: 'Авторская история и иллюстрации, созданные вокруг ребёнка и его семьи.',
					provider: { '@id': 'https://derezaivan.github.io/book-landing/#organization' }
				}
			]
		})}
	</script>
</svelte:head>

<div class="site-shell">
	<Header />
	<main id="top">
		<Hero />
		<BookFlow
			pages={6}
			anchors={['stories', 'story-showcase', 'approach', 'formats', 'studio', 'contact']}
		>
			<Chapter index="02" tone="milk"><Manifesto /></Chapter>
			<Chapter index="03" tone="milk"><StoryShowcase /></Chapter>
			<Chapter index="04" tone="forest"><Process /></Chapter>
			<Chapter index="05" tone="paper"><Formats /></Chapter>
			<Chapter index="06" tone="paper"><Studio /></Chapter>
			<Chapter index="07" tone="night"><Contact /></Chapter>
		</BookFlow>
	</main>
	<Footer />
</div>

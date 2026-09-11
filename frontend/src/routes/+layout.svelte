<script lang="ts">
	import favicon from '@/assets/favicon.svg';
	import '@/styles/global.scss';
	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
	<meta name="theme-color" content="#f2ede3" />
</svelte:head>

{@render children()}

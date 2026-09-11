import adapter from '@sveltejs/adapter-static';
import { sveltekit } from '@sveltejs/kit/vite';
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';
import { fileURLToPath } from 'node:url';
import { defineConfig } from 'vite';

const libPath = fileURLToPath(new URL('./src/lib', import.meta.url));
const base = process.env.BASE_PATH ?? '';

export default defineConfig({
	resolve: {
		alias: {
			'@': libPath
		}
	},
	server: {
		proxy: {
			'/api': 'http://localhost:8080'
		}
	},
	plugins: [
			sveltekit({
			preprocess: vitePreprocess(),
			alias: {
				'@': libPath
			},
			compilerOptions: {
				// Force runes mode for the project, except for libraries. Can be removed in svelte 6.
				runes: ({ filename }) =>
					filename.split(/[/\\]/).includes('node_modules') ? undefined : true
			},

			paths: {
				base
			},
			adapter: adapter({
				pages: 'build',
				assets: 'build',
				fallback: '404.html',
				precompress: true,
				strict: true
			})
		})
	]
});

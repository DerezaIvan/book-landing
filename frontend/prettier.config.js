/** @type {import("prettier").Config} */
const config = {
	useTabs: true,
	tabWidth: 2,
	singleQuote: true,
	semi: true,
	trailingComma: 'none',
	printWidth: 100,
	bracketSameLine: false,
	svelteSortOrder: 'options-scripts-markup-styles',
	svelteAllowShorthand: true,
	plugins: ['prettier-plugin-svelte'],
	overrides: [
		{ files: '*.svelte', options: { parser: 'svelte' } },
		{ files: '*.{scss,css}', options: { parser: 'scss', singleQuote: true } },
		{ files: '*.{ts,js}', options: { parser: 'typescript' } }
	]
};

export default config;

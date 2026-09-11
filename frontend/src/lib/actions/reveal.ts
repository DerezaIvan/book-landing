export function reveal(node: HTMLElement) {
	if (
		window.matchMedia(
			'(prefers-reduced-motion: reduce), (max-width: 1024px), (hover: none), (pointer: coarse)'
		).matches
	) {
		node.classList.add('is-visible');
		return;
	}

	const observer = new IntersectionObserver(
		([entry]) => {
			if (entry.isIntersecting) {
				node.classList.add('is-visible');
				observer.disconnect();
			}
		},
		{ threshold: 0.18 }
	);
	observer.observe(node);

	return { destroy: () => observer.disconnect() };
}

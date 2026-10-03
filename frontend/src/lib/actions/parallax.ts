export function parallax(node: HTMLElement, speed = 0.1) {
	if (
		window.matchMedia('(prefers-reduced-motion: reduce), (hover: none), (pointer: coarse)').matches
	) {
		return;
	}

	let frame = 0;
	const update = () => {
		frame = 0;
		const rect = node.getBoundingClientRect();
		if (rect.bottom < -200 || rect.top > window.innerHeight + 200) return;
		const offset = (rect.top + rect.height / 2 - window.innerHeight / 2) * -speed;
		node.style.setProperty('--py', `${offset.toFixed(1)}px`);
	};
	const onScroll = () => {
		if (!frame) frame = requestAnimationFrame(update);
	};

	update();
	window.addEventListener('scroll', onScroll, { passive: true });
	window.addEventListener('resize', onScroll);
	return {
		destroy() {
			window.removeEventListener('scroll', onScroll);
			window.removeEventListener('resize', onScroll);
			if (frame) cancelAnimationFrame(frame);
		}
	};
}

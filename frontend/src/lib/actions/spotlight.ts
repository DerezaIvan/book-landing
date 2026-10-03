export function spotlight(node: HTMLElement) {
	if (
		window.matchMedia('(prefers-reduced-motion: reduce), (hover: none), (pointer: coarse)').matches
	) {
		return;
	}

	let frame = 0;
	const onMove = (event: PointerEvent) => {
		if (event.pointerType === 'touch') return;
		if (frame) return;
		frame = requestAnimationFrame(() => {
			frame = 0;
			const bounds = node.getBoundingClientRect();
			node.style.setProperty('--mx', `${(event.clientX - bounds.left).toFixed(0)}px`);
			node.style.setProperty('--my', `${(event.clientY - bounds.top).toFixed(0)}px`);
		});
	};

	node.addEventListener('pointermove', onMove);
	return {
		destroy() {
			node.removeEventListener('pointermove', onMove);
			if (frame) cancelAnimationFrame(frame);
		}
	};
}

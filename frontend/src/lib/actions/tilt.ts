export function tilt(node: HTMLElement, max = 5) {
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
			const px = (event.clientX - bounds.left) / bounds.width - 0.5;
			const py = (event.clientY - bounds.top) / bounds.height - 0.5;
			node.style.setProperty('--tilt-rx', `${(-py * max).toFixed(2)}deg`);
			node.style.setProperty('--tilt-ry', `${(px * max).toFixed(2)}deg`);
		});
	};
	const onLeave = () => {
		node.style.setProperty('--tilt-rx', '0deg');
		node.style.setProperty('--tilt-ry', '0deg');
	};

	node.addEventListener('pointermove', onMove);
	node.addEventListener('pointerleave', onLeave);
	return {
		destroy() {
			node.removeEventListener('pointermove', onMove);
			node.removeEventListener('pointerleave', onLeave);
			if (frame) cancelAnimationFrame(frame);
		}
	};
}

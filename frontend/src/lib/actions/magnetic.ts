export function magnetic(node: HTMLElement, strength = 0.28) {
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
			const x = (event.clientX - (bounds.left + bounds.width / 2)) * strength;
			const y = (event.clientY - (bounds.top + bounds.height / 2)) * strength;
			node.style.translate = `${x.toFixed(1)}px ${y.toFixed(1)}px`;
		});
	};
	const onLeave = () => {
		node.style.translate = '0px 0px';
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

<script lang="ts">
	let formStatus = $state<'idle' | 'sending' | 'success' | 'error'>('idle');

	async function submitApplication(event: SubmitEvent) {
		event.preventDefault();
		formStatus = 'sending';
		const form = event.currentTarget as HTMLFormElement;
		const values = new FormData(form);

		try {
			const endpoint = import.meta.env.VITE_CONTACT_ENDPOINT || '/api/applications';
			const response = await fetch(endpoint, {
				method: 'POST',
				headers: { 'content-type': 'application/json' },
				body: JSON.stringify({ name: values.get('name'), contact: values.get('contact') })
			});
			if (!response.ok) throw new Error('request failed');
			form.reset();
			formStatus = 'success';
		} catch {
			formStatus = 'error';
		}
	}
</script>

<section class="contact">
	<div class="contact__orbit" aria-hidden="true"></div>
	<span class="contact__eyebrow">Начнём с разговора</span>
	<h2>Какую историю<br />мы расскажем <em>вместе?</em></h2>
	<p class="contact__intro">
		Оставьте контакт - мы напишем вам, чтобы познакомиться и обсудить будущую книгу.
	</p>
	<div class="contact__panel">
		<div class="contact__panel-copy">
			<span>Первое знакомство</span>
			<p>Расскажите, как с вами связаться. Детали будущей истории обсудим лично.</p>
		</div>
		<div class="contact__form-column">
			<form class="contact__form" onsubmit={submitApplication}>
				<label>
					<span>Как к вам обращаться</span>
					<input name="name" autocomplete="name" placeholder=" " required />
				</label>
				<label>
					<span>Телефон или Telegram</span>
					<input name="contact" autocomplete="tel" placeholder=" " required />
				</label>
				<button type="submit" disabled={formStatus === 'sending'}>
					<span>{formStatus === 'sending' ? 'Отправляем…' : 'Обсудить книгу'}</span>
					<i aria-hidden="true">→</i>
				</button>
			</form>
			{#if formStatus === 'success'}
				<p class="contact__message" role="status">
					Спасибо! Мы получили вашу заявку и скоро свяжемся.
				</p>
			{/if}
			{#if formStatus === 'error'}
				<p class="contact__message contact__message--error" role="alert">
					Не удалось отправить заявку. Попробуйте ещё раз чуть позже.
				</p>
			{/if}
			<small class="contact__legal">
				Нажимая кнопку, вы соглашаетесь на обработку персональных данных.
			</small>
		</div>
	</div>
</section>

<style lang="scss">
	@use '@/components/sections/Contact/Contact';
</style>

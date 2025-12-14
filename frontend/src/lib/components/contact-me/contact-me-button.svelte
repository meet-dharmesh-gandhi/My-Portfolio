<script lang="ts">
	import '$lib/components/contact-me/contact-me-button.css';

    let isHovering = $state(false);
	let hasClicked = $state(false);
	let isHoveringInside = $state(false);
	let isContactVisible = $derived(isHovering || hasClicked || isHoveringInside);
</script>

<div class="fixed flex flex-col justify-end items-end bottom-2 right-2 rounded-lg">
	{#if isContactVisible}
		<aside
			class="bottom-full right-0 w-max grid grid-cols-2 *:w-42 *:truncate bg-(--white) rounded-lg py-2 px-4 text-(--dark-blue) rounded-br-none"
			onmouseenter={() => (isHoveringInside = true)}
			onmouseleave={() => (isHoveringInside = false)}
		>
			<strong>Phone Number:</strong>
			<p>1234567890</p>
			<strong>Email:</strong>
			<a target="_blank" href="mailto:meetgandhi13062005@gmail.com">meetgandhi13062005@gmail.com</a>
			<strong>LinkedIn:</strong>
			<a target="_blank" href="https://linkedin.com/in/meet-d-gandhi"
				>https://linkedin.com/in/meet-d-gandhi</a
			>
			<strong>Github:</strong>
			<a target="_blank" href="https://github.com/meet-dharmesh-gandhi"
				>https://github.com/meet-dharmesh-gandhi</a
			>
		</aside>
	{/if}
	<button
		role="menu"
		tabindex="0"
		class={`contact-me-button relative bg-(--white) rounded-lg py-2 px-4 text-(--dark-blue) font-bold cursor-pointer flex flex-col gap-2 justify-end items-end hover:rounded-t-none hover:rounded-tr-none ${isContactVisible ? '' : 'before:hidden after:hidden'}`}
		onmouseenter={() => (isHovering = true)}
		onmouseleave={() => (isHovering = false)}
		onclick={() => (hasClicked = !hasClicked)}
	>
		Contact Me?
	</button>
</div>

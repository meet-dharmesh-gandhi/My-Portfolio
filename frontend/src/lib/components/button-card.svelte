<script lang="ts">
	import ViewMoreButton from "./home/ViewMoreButton.svelte";
    import { twMerge } from 'tailwind-merge';

    const {
		title,
		redirect,
		cardClass = '',
		contentClass = '',
		buttonText,
		children
	} = $props();
</script>

<article
	class={twMerge(
		'flex flex-col bg-(--blue) text-(--black) rounded-xl p-8 border-4 g-dg-btg-md-border',
		cardClass
	)}
>
	{#if title}
		<h3>{title}</h3>
	{/if}
	<div class={contentClass}>
		{@render children()}
	</div>
	<ViewMoreButton link={redirect ?? '/'} text={buttonText} />
</article>

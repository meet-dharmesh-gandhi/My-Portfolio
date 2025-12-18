<script lang="ts">
    import { twMerge } from 'tailwind-merge';

	const { colorClass = 'bg-(--black)', className = '', showTitle = true, title, children } = $props();
</script>

<section
	id={typeof title === 'string' ? title.replaceAll(' ', '-').toLowerCase() : undefined}
	class={twMerge('flex flex-col justify-center items-center gap-4 py-8', colorClass, className)}
>
	{#if title !== '' && showTitle}
		<h2 class="text-(--green)">{title}</h2>
	{/if}
	{@render children()}
</section>

<script lang="ts">
	const { colorClass = 'bg-(--black)', title, children } = $props();
</script>

<section class={`flex flex-col justify-center items-center gap-4 py-8 ${colorClass}`}>
	<h2 class="text-(--green)">{title}</h2>
	{@render children()}
</section>

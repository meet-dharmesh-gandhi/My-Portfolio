<script lang="ts">
	import ViewMoreButton from "./home/ViewMoreButton.svelte";

    const { title, redirect, cardClass = '', children } = $props();
</script>

<article
	class={`flex flex-col bg-(--blue) text-(--black) rounded-xl p-8 border-4 g-dg-btg-md-border ${cardClass}`}
>
	<h3>{title}</h3>
	{@render children()}
	<ViewMoreButton link={redirect ?? '/'} text="View this project" />
</article>

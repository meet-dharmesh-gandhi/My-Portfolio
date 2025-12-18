<script lang="ts">
	import Section from "./Section.svelte";

    const { colorClass } = $props();
</script>

<Section
	{colorClass}
	title="My Journey"
	className="h-[80dvh] gap-16 mt-auto mb-auto text-center"
	showTitle={false}
>
	<h3 class="text-4xl!">Want to know my human side?</h3>
	<a
		href="/"
		class="revert-a bg-(--lighter-blue) text-(--black) px-16 py-2 rounded-full font-bold border-4 border-(--dark-blue) hover:bg-(--green) hover:text-(--dark-blue) hover:border-(--dark-green) hover:scale-105 transition-all duration-200"
	>
		Click here to view my journey as an individual &rarr;
	</a>
</Section>

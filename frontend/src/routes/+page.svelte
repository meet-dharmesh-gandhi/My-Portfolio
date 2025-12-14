<script>
	import Hero from "$lib/components/home/Hero.svelte";
	import Intro from "$lib/components/home/Intro.svelte";
	import Experience from "$lib/components/home/Experience.svelte";
	import './home.css';
	import Projects from "$lib/components/home/Projects.svelte";

	let sections = [Hero, Intro, Experience, Projects];
</script>

<main>
	{#each sections as Section, index}
		<Section colorClass={'section-gradient'} />
		{#if index !== sections.length - 1}
			<hr />
		{/if}
	{/each}
</main>

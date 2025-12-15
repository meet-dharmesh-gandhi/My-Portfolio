<script>
	import './home.css';
	import Hero from "$lib/components/home/Hero.svelte";
	import Intro from "$lib/components/home/Intro.svelte";
	import Experience from "$lib/components/home/Experience.svelte";
	import Projects from "$lib/components/home/Projects.svelte";
	import Achievements from "$lib/components/home/Achievements.svelte";

	let sections = [Hero, Intro, Experience, Projects, Achievements];
</script>

<main>
	{#each sections as Section, index}
		<Section colorClass={'section-gradient'} />
		{#if index !== sections.length - 1}
			<hr />
		{/if}
	{/each}
</main>

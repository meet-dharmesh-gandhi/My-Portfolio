<script lang="ts">
	import ButtonCard from "../button-card.svelte";
	import Section from "./Section.svelte";

	interface Experience {
		role: string,
		org?: string;
		description?: string;
		redirect?: string;
	}

	const { colorClass = 'bg-(--black)' } = $props();

    let experiences: Experience[] = [
        {
            role: 'Robotics Head',
            org: 'IEEE',
            description: 'Led and organized competitive robotics events in collaboration with multiple IEEE teams and branches.',
			redirect: '/',
		},
        {
            role: 'Web Development Intern',
            org: 'Club Duelz',
            description: 'Contributed feature ideas and implemented them through collaborative brainstorming and development.',
			redirect: '/',
		},
    ];
</script>

<Section title="Experience" {colorClass}>
	<article class="grid grid-cols-2 gap-8 p-8">
		{#each experiences as experience}
			<ButtonCard title={experience.role} redirect={experience.redirect ?? '/'}>
				{#if experience.org}
					<h4>{experience.org}</h4>
				{/if}
				{#if experience.description}
					<p class="mt-4">{experience.description}</p>
				{/if}
				<div aria-hidden="true" class="mt-4"></div>
			</ButtonCard>
		{/each}
	</article>
</Section>

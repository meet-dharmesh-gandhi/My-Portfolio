<script lang="ts">
	import ViewMoreButton from "./ViewMoreButton.svelte";

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

<section class={`flex flex-col justify-center items-center py-8 ${colorClass}`}>
	<h2 class="text-(--green)">Experience</h2>
	<article class="grid grid-cols-2 gap-8 p-8">
		{#each experiences as experience}
			<article
				class="flex flex-col bg-(--blue) rounded-xl w-fit p-8 text-(--black) border-4 g-dg-btg-md-border"
			>
				<h3>{experience.role}</h3>
				{#if experience.org}
					<h4>{experience.org}</h4>
				{/if}
				{#if experience.description}
					<p class="mt-8">{experience.description}</p>
				{/if}
				<div class="mt-8 flex justify-center">
					<ViewMoreButton
						link={experience.redirect ?? '/'}
						text="Know more about this experience"
					/>
				</div>
			</article>
		{/each}
	</article>
</section>

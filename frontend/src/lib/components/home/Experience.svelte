<script lang="ts">
	import "$lib/components/home/experience.css";

	interface Experience {
		role: string,
		org?: string;
		description?: string;
		redirect?: string;
	}

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

<section class="flex flex-col justify-center items-center">
	<h1 class="text-(--green)">Experience</h1>
	<article class="grid grid-cols-2 gap-8 p-8">
		{#each experiences as experience}
			<div
				class="flex flex-col bg-(--blue) rounded-xl w-fit p-8 text-(--black) border-4 g-dg-btg-md-border"
			>
				<h1>{experience.role}</h1>
				{#if experience.org}
					<h2>{experience.org}</h2>
				{/if}
				{#if experience.description}
					<p>{experience.description}</p>
				{/if}
				<a
					href={experience.redirect ?? '/'}
					class="view-experience-btn flex justify-center items-center self-center gap-4 bg-(--dark-blue) rounded-md text-(--white) w-fit px-4 cursor-pointer mt-8"
				>
					Know more about this experience
					<span class="text-2xl"> &rarr; </span>
				</a>
			</div>
		{/each}
	</article>
</section>

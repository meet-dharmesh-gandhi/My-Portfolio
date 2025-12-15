<script lang="ts">
	import Section from "./Section.svelte";
	import ViewMoreButton from "./ViewMoreButton.svelte";

    interface Achievements {
        title: string;
    }

    const { colorClass } = $props();

    const achievements: Achievements[] = [
        {
            title: 'WMC Hackathon Winner'
        },
        {
            title: 'NASA Space Apps Challenge Participant'
        },
        {
            title: 'IEEE AUSB ICVES Conference Volunteer'
        },
    ]
</script>

<Section title="Achievements" {colorClass}>
	<article class="flex flex-wrap gap-8 px-8 justify-center items-center">
		{#each achievements as achievement}
			<h3
				class="bg-(--blue) rounded-full px-4 py-2 border-2 g-dg-btg-md-border-capsule text-(--white)"
			>
				{achievement.title}
			</h3>
		{/each}
	</article>
	<div aria-hidden="true" class="mb-4"></div>
	<ViewMoreButton
		text="View more achievements"
		link="/#achievements"
		className="bg-(--green) text-(--black) rounded-full border-3 border-(--blue) py-1 px-8"
	/>
</Section>

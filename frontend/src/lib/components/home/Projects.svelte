<script lang="ts">
	import ButtonCard from "../button-card.svelte";
import Section from "./Section.svelte";

    interface Projects {
        title: string;
        description?: string;
        tags?: string[];
        madeByMeOnly?: false;
        redirect?: string;
    }

    interface Props {
        colorClass: string;
    }

	const { colorClass = 'bg-(--black)' } = $props();

    let projects: Projects[] = [
        {
            title: 'AnyAuth',
            description: 'This is a NPM library built for easy integration of OAuth into a website and its backend.',
            tags: ['JavaScript', 'NPM', 'Github'],
            redirect: '/'
        },
        {
            title: 'Cyber Tools',
            description: 'This is a platform to execute basic and advanced cyber security applications',
            tags: ['NodeJS', 'NextJS', 'Postman', 'Github', 'ChakraUI', 'Cloudinary'],
            madeByMeOnly: false,
            redirect: '/'
        },
        {
            title: 'BIJOI App',
            description: 'This is an app built for communication between managers and workers',
            tags: ['Flutter', 'NodeJS', 'Postman', 'Github', 'Cloudinary'],
            redirect: '/'
        }
    ]
</script>

<Section title="Projects" {colorClass}>
	<article class="grid grid-cols-3 gap-8 p-8">
		{#each projects as project}
			<ButtonCard
				title={project.title}
				redirect={project.redirect ?? ''}
				cardClass="relative gap-4"
			>
				<p>{project.description}</p>
				<article class="flex flex-wrap gap-x-2 gap-y-1">
					{#each project.tags as tag}
						<p
							class="px-4 py-2 border-2 bg-(--dark-blue) border-(--lighter-blue) text-(--white) rounded-full cursor-default"
						>
							{tag}
						</p>
					{/each}
				</article>
				<p class="absolute right-1.5 top-2 bg-(--green) rounded-full w-fit px-2 py-1 text-xs">
					{project.madeByMeOnly === false ? 'Made In Collaboration' : 'Made By Me'}
				</p>
				<div aria-hidden="true" class="mt-auto mb-8"></div>
			</ButtonCard>
		{/each}
	</article>
</Section>

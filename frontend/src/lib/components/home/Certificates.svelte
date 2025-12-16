<script lang="ts">
	import type { CarouselCard } from "$lib/types/carousel/carousel";
	import Carousel from "../carousel/Carousel.svelte";
    import Section from "./Section.svelte";
	import ViewMoreButton from "./ViewMoreButton.svelte";

    interface Certificates {
        title: string;
        filename: string;
        description?: string;
        redirect?: string;
    }

    interface CarouselExtras {
        title: string;
        description?: string;
        redirect?: string;
    }

    const { colorClass } = $props();

    const basePath = '/src/lib/assets/Certificate-Images/';
    const certificates: Certificates[] = [
        {
            title: 'BIJOI Internship',
            filename: 'BIJOI Certificate.jpg',
            description: 'Summmer internship 2025 certificate at BIJOI.',
            redirect: '/',
        },
        {
            title: 'ICVES Volunteer',
            filename: 'ICVES Conference Volunteering Certificate.jpg',
            description: 'Volunteer at IEEE AUSB hosted ICVES international conference.',
            redirect: '/',
        },
        {
            title: 'WMC Hackathon',
            filename: 'WMC Winning Certificate.jpg',
            description: 'Won the WMC Competition at AU',
            redirect: '/',
        }
    ];

    const carouselCards: CarouselCard<CarouselExtras>[] = certificates.map((certificate) => ({
        image: {
            url: basePath + certificate.filename,
            alt: `${certificate.title} Certificate`,
            className: 'aspect-[2/1] object-cover object-top mb-4'
        },
        className: 'flex flex-col bg-(--blue) justify-start items-start gap-1 text-(--black) rounded-xl overflow-hidden min-w-[50vw]',
        extraInfo: {
            title: certificate.title,
            description: certificate.description,
            redirect: certificate.redirect,
        },
    }));
</script>

<Section title="Certificates" {colorClass}>
	<Carousel {carouselCards} className="w-[60vw]">
		{#snippet children({
			carouselCard
		}: {
			carouselCard: CarouselCard<CarouselExtras> /* prettier-ignore */
		})}
			<div class="p-8">
				<h3>{carouselCard.extraInfo.title}</h3>
				<h4>{carouselCard.extraInfo.description}</h4>
			</div>
			{#if carouselCard.extraInfo.redirect}
				<ViewMoreButton
					text="View this experience"
					link={carouselCard.extraInfo.redirect}
					type="link"
					className="mb-8 mt-auto"
				/>
			{/if}
		{/snippet}
	</Carousel>
</Section>

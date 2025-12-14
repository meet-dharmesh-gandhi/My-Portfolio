<script lang="ts">
	import Section from "./Section.svelte";

	const { colorClass = 'bg-(--black)' } = $props();
</script>

<Section title="Bio" {colorClass}>
	<div class="w-[80vw] py-2 px-2 flex flex-col justify-start items-start gap-3">
		<p>
			Hi, I&rsquo;m Meet Gandhi &mdash; a developer driven by curiosity and a love for building
			real-world systems.
		</p>
		<p>
			I&rsquo;m particularly interested in system design and operating systems, while also being
			comfortable across full-stack web and cross-platform application development.
		</p>
		<p>
			This portfolio captures that journey &mdash; from learning how systems work to building them
			with confidence.
		</p>
	</div>
</Section>

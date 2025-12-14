<script lang="ts">
    import image from '$lib/assets/Meet Resume Photo.jpg';
	import { twMerge } from 'tailwind-merge';

	const { colorClass = 'bg-(--black)' } = $props();
</script>

<section class={twMerge('flex gap-8 justify-center items-center h-dvh', colorClass)}>
	<div class="rounded-full overflow-hidden h-fit">
		<img
			src={image}
			alt="Meet Gandhi Profile"
			class="h-[50vh] w-[50vh] object-cover object-top"
			draggable="false"
		/>
	</div>
	<div class="flex flex-col justify-start items-start h-[40vh] gap-1 pt-8">
		<h1>Meet Gandhi</h1>
		<h3>Building Real World Systems</h3>
		<h4>Flutter &middot; Full Stack &middot; Open Source &middot; System Design</h4>
		<a
			href="/#bio"
			class="revert-a mt-auto bg-(--white) rounded-full px-8 py-2 w-fit font-bold text-(--dark-blue) cursor-pointer hover:bg-(--dark-blue) hover:text-(--white) transition-all duration-200"
		>
			Explore more about me
		</a>
	</div>
</section>

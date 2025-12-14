<script lang="ts">
    import image from '$lib/assets/Meet Resume Photo.jpg';
</script>

<section class="flex gap-8 justify-center items-center h-dvh">
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
		<h2>Building Real World Systems</h2>
		<h3>Flutter &middot; Full Stack &middot; Open Source &middot; System Design</h3>
		<button
			class="mt-auto bg-(--white) rounded-full px-8 py-2 w-fit font-bold text-(--dark-blue) cursor-pointer hover:bg-(--dark-blue) hover:text-(--white) transition-all duration-200"
			>Explore more about me</button
		>
	</div>
</section>

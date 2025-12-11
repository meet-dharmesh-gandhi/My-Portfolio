<script lang="ts">
    let { title }: { title: string } = $props();
</script>

<button
	class="bg-(--blue) border-2 border-solid border-(--white) text-(--white) rounded-full px-4 py-2 transition-all duration-100 ease-in-out hover:bg-(--maroon) hover:border-(--green) hover:text-(--white) focus:bg-(--maroon) active:bg-(--maroon) nav-options text-center w-fit cursor-pointer"
	type="submit"
	aria-label={`${title} page link`}
>
	{title}
</button>

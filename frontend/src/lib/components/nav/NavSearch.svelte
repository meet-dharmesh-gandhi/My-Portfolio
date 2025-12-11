<script lang="ts">
	import SearchIcon from "$lib/assets/search-icon.svg?raw";
</script>

<search>
	<form
		class="flex justify-center items-stretch gap-2 bg-(--blue) rounded-full overflow-hidden pl-5"
		method="dialog"
	>
		<input
			class="my-1 w-full text-(--pure-white)"
			type="text"
			placeholder="Search Me..."
			name="search-button"
		/>
		<button
			class="search-button rounded-full! my-1 mr-1 bg-(--white)! cursor-pointer"
			type="submit"
			aria-label="search-button"
		>
			{@html SearchIcon}
		</button>
	</form>
</search>

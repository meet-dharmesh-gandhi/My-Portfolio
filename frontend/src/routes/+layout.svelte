<script lang="ts">
	import './layout.css';
	import favicon from '$lib/assets/favicon.ico';
	import Navbar from '$lib/components/nav/Navbar.svelte';
	import ContactMeButton from '$lib/components/contact-me/contact-me-button.svelte';

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
	<title>Meet Gandhi Portfolio</title>
	<meta
		name="description"
		content="Home of Meet Gandhi's unique portfolio. Highlights the pages you can explore further and gives breif description about Meet"
	/>
</svelte:head>

<Navbar />
{@render children()}
<ContactMeButton />

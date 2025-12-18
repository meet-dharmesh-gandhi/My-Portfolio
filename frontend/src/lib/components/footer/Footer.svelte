<script lang="ts">
    import image from '$lib/assets/Meet Resume Photo.jpg';
    import LinkedInIcon from '$lib/assets/linkedin.svg?raw';
    import GitHubIcon from '$lib/assets/github.svg?raw';
    import PhoneIcon from '$lib/assets/phone.svg?raw';
    import './footer.css';

    type Link = [string, string, string];
    type PageLink = [string, string];

    const links: Link[] = [
        [LinkedInIcon, 'https://linkedin/in/meet-d-gandhi', 'LinkedIn'],
        [GitHubIcon, 'https://github.com/meet-dharmesh-gandhi', 'GitHub'],
        [PhoneIcon, '1234567890', 'Phone'],
    ];

    const pageLinks: PageLink[] = [
        ['Home', '/'],
        ['Builder', '/builder'],
    ];
</script>

<footer class="footer flex justify-around bg-(--blue) py-4">
	<aside
		class="flex flex-col justify-center items-center flex-1 py-4 border-r-2 border-r-(--black)"
	>
		<img
			src={image}
			alt="Meet Gandhi Profile"
			class="h-[100px] w-[100px] object-cover object-top rounded-full"
			draggable="false"
		/>
		<h2 class="text-(--black)">Meet Gandhi</h2>
	</aside>
	<article
		class="flex flex-col justify-start items-center gap-4 flex-1 py-4 border-r-2 border-r-(--black)"
	>
		<div class="flex flex-col gap-4">
			<h3 class="text-(--black)">Quick Links</h3>
			<ul class="flex flex-col gap-2 w-full">
				{#each pageLinks as pageLink}
					<li class="w-full">
						<a href={pageLink[1]} class="text-(--black)!"><p class="text-md">{pageLink[0]}</p></a>
					</li>
				{/each}
			</ul>
		</div>
	</article>
	<aside class="flex justify-center items-center flex-1 py-4">
		<ul class="flex flex-col justify-between items-stretch gap-2">
			{#each links as link}
				<li class="bg-(--white) py-1 px-2 rounded-lg">
					<a href={link[0]} class="revert-a flex gap-4">
						{@html link[0]}
						<p class="text-(--black)">{link[2]}</p>
					</a>
				</li>
			{/each}
		</ul>
	</aside>
</footer>

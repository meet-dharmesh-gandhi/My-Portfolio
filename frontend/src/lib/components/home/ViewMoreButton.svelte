<script lang="ts">
    import './ViewMoreButton.css';

    interface Props {
        link: string;
        text: string;
    }

    const { link, text }: Props = $props();
</script>

<a
	href={link}
	class="view-experience-btn flex justify-center items-center self-center gap-4 bg-(--dark-blue) rounded-md text-(--white) w-fit px-4 cursor-pointer"
>
	{text}
	<span class="text-2xl"> &rarr; </span>
</a>

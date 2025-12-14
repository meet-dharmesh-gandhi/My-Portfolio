<script lang="ts">
    import './ViewMoreButton.css';
    import { twMerge } from 'tailwind-merge';

    interface Props {
        link: string;
        text: string;
        className?: string;
    }

    const { link, text, className }: Props = $props();
</script>

<a
	href={link}
	class={twMerge(
		'view-experience-btn revert-a flex justify-center items-center self-center gap-4 bg-(--dark-blue) rounded-md text-(--white) w-fit px-4 cursor-pointer',
		className
	)}
>
	{text}
	<span class="text-2xl"> &rarr; </span>
</a>

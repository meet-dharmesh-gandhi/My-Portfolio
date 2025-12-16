<script lang="ts" generics="T">
	import { twMerge } from "tailwind-merge";
    import type { CarouselCard } from "$lib/types/carousel/carousel";
	import LeftIcon from '$lib/assets/left-icon.svg?raw';
	import RightIcon from '$lib/assets/right-icon.svg?raw';
	import './carousel.css';
	import { onMount } from "svelte";

    interface Props<T> {
        carouselCards: CarouselCard<T>[];
        keepArrows?: boolean;
        className?: string;
		carouselTime?: number;
        children: any;
    }

	interface ScrollCarouselProps {
		left?: boolean;
		shouldRestartTimer?: boolean;
	}

    const { carouselCards, keepArrows = true, className, carouselTime = 3000, children }: Props<T> = $props();
	const arrowClassName = "arrow-icon bg-(--white) border-2 border-(--blue) rounded-full p-4 cursor-pointer hover:bg-(--dark-blue) hover:border-(--white)";

    let interval: NodeJS.Timeout | undefined;
	let ulEle: HTMLUListElement | null | undefined;
	let scrollLen: number | undefined, scrollWidth: number | undefined;

	function getNextScrollDistance(reverse: boolean = false) {
		if (ulEle && scrollWidth && scrollLen) {
			if (reverse) {
				console.log('getting next scroll distance in reverse');
				console.log(ulEle.scrollLeft, scrollLen, scrollWidth);
				if (ulEle.scrollLeft < (scrollLen / 4)) {
					return scrollWidth;
				}
				return -scrollLen;
			} else {
				if (scrollWidth - ulEle.scrollLeft < scrollLen)
					return -scrollWidth;
				return scrollLen;
			}
		}
	}

	function defineScrollConstants() {
		if (ulEle) {
			if (scrollLen === undefined) scrollLen = window.innerWidth * 0.7;
			if (scrollWidth === undefined) scrollWidth = ulEle.scrollWidth;
		}
	}

	function restartTimer() {
		clearInterval(interval);
		interval = setInterval(() => scrollCarousel(), carouselTime);
	}

	function pauseTimer() {
		clearInterval(interval);
	}

	function scrollCarousel({left = false, shouldRestartTimer = false}: ScrollCarouselProps = {}) {
		if (ulEle) {
			defineScrollConstants();
			ulEle.scrollBy({ left: getNextScrollDistance(left), behavior: 'smooth' });
			if (shouldRestartTimer)
				restartTimer();
		}
	}

	onMount(restartTimer);
</script>

<article class="relative w-full flex justify-center items-center gap-32">
	<h2 class="absolute w-0 opacity-0 h-0">Certificates carousel containing 3 certificates</h2>
	<button
		class={arrowClassName}
		onclick={() => scrollCarousel({ left: true, shouldRestartTimer: true })}
	>
		{@html LeftIcon}
	</button>
	<ul
		class={twMerge(
			'w-[80vw] flex flex-nowrap overflow-x-scroll gap-8 snap-mandatory snap-x p-2 no-scrollbar',
			className
		)}
		bind:this={ulEle}
	>
		<div aria-hidden="true" class="min-w-[calc(5vw-2rem)]"></div>
		{#each carouselCards as carouselCard}
			<li
				role="group"
				class={twMerge('min-w-[70vw] snap-center', carouselCard.className ?? '')}
				onmouseenter={pauseTimer}
				onmouseleave={restartTimer}
			>
				<img
					class={carouselCard.image.className ?? ''}
					src={carouselCard.image.url}
					alt={carouselCard.image.alt}
				/>
				{@render children({ carouselCard })}
			</li>
		{/each}
		<div aria-hidden="true" class="min-w-[calc(5vw-2rem)]"></div>
	</ul>
	<button class={arrowClassName} onclick={() => scrollCarousel({ shouldRestartTimer: true })}>
		{@html RightIcon}
	</button>
</article>

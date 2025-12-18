export interface CarouselImage {
	file: string;
	alt: string;
	className?: string;
}

export interface CarouselCard<T> {
	image: CarouselImage;
	className?: string;
	extraInfo: T;
}

<script lang="ts">
    let { title }: { title: string } = $props();
	let focusedElement: HTMLButtonElement | null = null;

	/**
	 * This function takes in the event and checks if any arrow key is pressed
	 * According to the key pressed, it focuses on the next or previous sibling
	 * If the extreme children are focused, the children in the other extreme are focused
	 * @param e
	 */
	function onKeyboardClick(e: KeyboardEvent) {
		const { key, target } = e;
		if (target && target instanceof HTMLElement) {
			const { previousElementSibling, parentElement, nextElementSibling } = target;
			if ((key === 'ArrowUp' || key === 'ArrowLeft')) {
				// move to the previous element
				if (previousElementSibling && previousElementSibling instanceof HTMLButtonElement) {
					// focus on previous element
					previousElementSibling.focus();
					focusedElement = previousElementSibling;
				} else if (parentElement?.lastElementChild && parentElement.lastElementChild instanceof HTMLButtonElement) {
					// if there is no previous element, focus on the last one
					parentElement.lastElementChild.focus();
					focusedElement = parentElement.lastElementChild;
				}
			} else if ((key === 'ArrowDown' || key === 'ArrowRight')) {
				// move to the next element
				if (nextElementSibling && nextElementSibling instanceof HTMLButtonElement) {
					// focus on the next element
					nextElementSibling.focus();
					focusedElement = nextElementSibling;
				} else if (parentElement?.firstElementChild && parentElement.firstElementChild instanceof HTMLButtonElement) {
					// if there is no next element, focus on the first element
					parentElement.firstElementChild.focus();
					focusedElement = parentElement.firstElementChild;
				}
			}
		}
	}
</script>

<button
	class="bg-(--blue) border-2 border-solid border-(--white) text-(--white) rounded-full px-4 py-2 transition-all duration-100 ease-in-out hover:bg-(--maroon) hover:border-(--green) hover:text-(--white) focus:bg-(--maroon) active:bg-(--maroon) nav-options text-center w-fit cursor-pointer"
	type="submit"
	aria-label={`${title} page link`}
	aria-expanded="true"
	role="menuitem"
	onkeydown={onKeyboardClick}
>
	{title}
</button>

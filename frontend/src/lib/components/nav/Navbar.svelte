<script lang="ts">
    import "./Navbar.css"
	import NavbarOption from "./NavbarOption.svelte";
	import NavSearch from "./NavSearch.svelte";

	type tabNames = 'Explore' | 'About Me' | 'Fun';
	type validHovers = tabNames | 'dialog';

    let hoveringAt: validHovers | null = $state(null);
	const tabs: Record<tabNames, string[]> = {
		'Explore': ['Open Source', 'Projects', 'Experience', 'Certificates', 'Blogs', 'Achievements'].sort((a, b) => a.length < b.length ? -1 : 1),
		'About Me': ['Contact', 'CV', 'Journey Timeline'].sort((a, b) => a.length < b.length ? -1 : 1),
		'Fun': ['Engineering Playground', 'Builder'].sort((a, b) => a.length < b.length ? -1 : 1),
	};
	let wantsToView: tabNames | null = $state(null);
	let dialog: HTMLDialogElement | null = $state(null);

	$effect(() => {
		if (dialog instanceof HTMLDialogElement) {
			if (hoveringAt && wantsToView && wantsToView in tabs) {
				dialog.show();
				dialog.querySelector('input')?.focus();
			// }
			} else {
				dialog.close();
			}
		}
	})

    function onHoverStart(hover: validHovers) {
		if (hover !== 'dialog' && hover) {
			wantsToView = hover;
		}
        hoveringAt = hover;
	}

    function onHoverEnd(hover: validHovers) {
        if (hoveringAt === hover) {
            hoveringAt = null;
        }
    }

	function onButtonClick() {
		if (dialog instanceof HTMLDialogElement) {
			dialog.show();
		}
	}

	function onKeyDownDialog(e: KeyboardEvent) {
		if (e.key === 'Tab') {
			const dialogChildren = [...(dialog?.querySelectorAll('button:not([disabled]), input:not([disabled])') ?? [])];
			if (!dialogChildren) return;
			const firstChild = dialogChildren.at(0);
			const lastChild = dialogChildren.at(-1);
			if (!(firstChild instanceof HTMLElement && lastChild instanceof HTMLElement)) {
				return;
			}
			if (e.shiftKey && firstChild === e.target && lastChild) {
				e.preventDefault();
				lastChild.focus();
			} else if (!e.shiftKey && dialogChildren.at(-1) === e.target && firstChild) {
				e.preventDefault();
				firstChild.focus();
			}
		}
	}
</script>

<nav
	class="flex z-10 items-center fixed w-full justify-between pl-8 bg-(--glass-white) text-(--yellow) rounded-b-4xl rounded-bl-4xl backdrop-blur-sm border border-solid border-(--white) border-t-0 transition-all duration-200 hover:rounded-none hover:bg-(--white) hover:backdrop-blur-none hover:text-(--dark-blue) hover:border-none"
>
	<h2>Meet Gandhi</h2>
	<ul class="flex">
		{#each Object.keys(tabs) as ele}
			<li>
				<button
					type="button"
					class="nav-tab"
					onmouseenter={() => onHoverStart(ele as tabNames)}
					onmouseleave={() => onHoverEnd(ele as tabNames)}
					onclick={onButtonClick}
				>
					{ele}
				</button>
			</li>
		{/each}
	</ul>
	<dialog
		class="mega-menu bg-transparent"
		onmouseenter={() => onHoverStart('dialog')}
		onmouseleave={() => onHoverEnd('dialog')}
		onkeydown={onKeyDownDialog}
		bind:this={dialog}
		closedby="any"
		role="menu"
	>
		<div class="grid grid-cols-2 justify-between gap-10 p-10 rounded-b-2xl wrapper bg-(--white)">
			<section>
				<NavSearch />
			</section>
			<section class="flex justify-end">
				<form class="flex flex-wrap gap-2 justify-end">
					{#if wantsToView}
						{#each tabs[wantsToView] as ele}
							<NavbarOption title={ele} />
						{/each}
					{/if}
				</form>
			</section>
		</div>
	</dialog>
</nav>

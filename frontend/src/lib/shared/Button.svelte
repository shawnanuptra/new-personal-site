<script lang="ts">
	import type { HTMLButtonAttributes } from 'svelte/elements';

	interface ButtonProps extends HTMLButtonAttributes {
		label: string;
		variant: 'primary' | 'secondary';
	}

	let { label, variant = 'primary', ...rest }: ButtonProps = $props();
</script>

<button class="button button--{variant}" {...rest}>
	{label}
</button>

<style>
	.button {
		padding: 0.5rem 1rem;
		border-radius: 0.5rem;
		border: 1px solid transparent;
		cursor: pointer;
		transition: all 150ms ease-in-out;
	}

	/* Variants */
	.button--primary {
		border-color: var(--black);
		background-color: var(--black);
		color: var(--white);

		&:hover,
		&:focus {
			background-color: var(--white);
			color: var(--black);

			transform: translateX(-1px) translateY(-1px);
			box-shadow: var(--black) -5px 5px 0px;
			outline: none;
		}
	}

	.button--secondary {
		border: 1px dashed var(--black);
		background-color: var(--white);
		color: var(--black);

		&:hover,
		&:focus {
			border: 1px solid var(--black);
			background-color: var(--white);
			color: var(--black);

			transform: translateX(-1px) translateY(-1px);
			box-shadow: var(--black) -5px 5px 0px;
			outline: none;
		}
	}
</style>

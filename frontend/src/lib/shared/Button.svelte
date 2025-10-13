<script lang="ts">
	import type { HTMLButtonAttributes } from 'svelte/elements';

	interface ButtonProps extends HTMLButtonAttributes {
		label: string;
		to?: string;
		variant?: 'primary' | 'secondary';
	}

	let { label, to, variant = 'primary', ...rest }: ButtonProps = $props();
</script>

{#if !to}
	<button class="button button--{variant}" {...rest}>
		{label}
	</button>
{:else}
	<a href={to} class="button button--{variant}">{label}</a>
{/if}

<style>
	.button {
		font-size: 1rem;
		display: inline-block;
		padding: 0.5rem 1rem;
		border-radius: var(--border-radius-sm);
		border: 1px solid transparent;
		cursor: pointer;
		transition: all 150ms ease-in-out;
		text-decoration: none;
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

			transform: translateX(2px) translateY(-1px);
			box-shadow: var(--box-shadow);
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

			transform: translateX(2px) translateY(-1px);
			box-shadow: var(--box-shadow);
			outline: none;
		}
	}
</style>

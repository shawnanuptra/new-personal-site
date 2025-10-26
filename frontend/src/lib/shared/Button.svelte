<script lang="ts">
	import type { HTMLButtonAttributes } from 'svelte/elements';

	interface ButtonProps extends HTMLButtonAttributes {
		label: string;
		to?: string;
		variant?: 'primary' | 'secondary';
		size?: 'large' | 'small';
	}

	let { label, to, variant = 'primary', size = 'small', ...rest }: ButtonProps = $props();
</script>

{#if !to}
	<button class="button button--{variant} size-{size}" {...rest}>
		{label}
	</button>
{:else}
	<!-- eslint-disable-next-line svelte/no-navigation-without-resolve -->
	<a href={to} class="button button--{variant} size-{size}">{label}</a>
{/if}

<style>
	.button {
		display: block;
		width: fit-content;
		padding: 0.5rem 1rem;
		border-radius: var(--border-radius-sm);
		border: 1px solid transparent;
		cursor: pointer;
		transition: all 150ms ease-in-out;
		text-decoration: none;

		&.size-large {
			border: 2px solid transparent;
			font-size: 1.5rem;
			font-weight: bold;
		}
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

			&.size-large {
				border: 2px solid var(--black);
			}

			transform: translateX(5px) translateY(-5px);
			box-shadow: var(--box-shadow);
			outline: none;
		}
	}

	.button--secondary {
		border: 1px dashed var(--black);
		background-color: var(--white);
		color: var(--black);

		&.size-large {
			border: 2px dashed var(--black);
		}

		&:hover,
		&:focus {
			border: 1px solid var(--black);
			background-color: var(--white);
			color: var(--black);

			&.size-large {
				border: 2px solid var(--black);
			}

			transform: translateX(5px) translateY(-5px);
			box-shadow: var(--box-shadow);
			outline: none;
		}
	}
</style>

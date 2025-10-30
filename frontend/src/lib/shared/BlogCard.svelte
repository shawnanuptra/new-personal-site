<script lang="ts">
	import type { Blog } from '$lib/types/api';

	const { blog }: { blog: Blog } = $props();
</script>

<article class="card">
	<div class="card-title-wrapper">
		<a class="card-title" href={'blogs/' + blog.slug}>{blog.title}</a>
	</div>
	<p class="card-content">{blog.description}</p>
</article>

<style>
	.card {
		position: relative;
		display: flex;
		flex-direction: column;
		gap: 1rem;
		background-color: var(--white);

		border: 2px dashed var(--black);
		border-radius: var(--border-radius-md);
		padding: 2rem;
		transition: all 150ms ease-in-out;

		&:hover,
		&:focus-within {
			border: 2px solid var(--black);
			box-shadow: var(--box-shadow-lg);
			transform: translateX(10px) translateY(-10px);
		}
	}

	.card-content {
		order: 3;
	}

	.card-title-wrapper {
		order: 2;
	}

	.card-title {
		font-size: 2.5rem;
		font-weight: bold;
		text-decoration-style: dashed;
		color: var(--black);
		transition: color 150ms 300ms ease-in-out;

		/* stretch the hitbox */
		&::after {
			position: absolute;
			content: '';
			inset: 0;
			z-index: 1;
		}

		margin: 1rem 0;

		&:hover,
		&:focus,
		&:focus-visible {
			outline: none;
			box-shadow: none;
			background: linear-gradient(to bottom, var(--white) 0%, var(--white) 63%, #fdff32 63%);
		}
	}
</style>

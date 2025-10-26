<script lang="ts">
	import type { Project } from '$lib/types/api';

	const { project }: { project: Project } = $props();
</script>

<article class="card">
	<div class="card-title-wrapper">
		<a class="card-title" href={'project/' + project.slug}>{project.title}</a>
	</div>
	<p class="card-content">{project.description}</p>
	<img
		class="card-thumbnail"
		src={project.thumbnailUrl + '?w=750&fit=max'}
		aria-hidden="true"
		alt=""
	/>
</article>

<style>
	.card {
		position: relative;
		display: grid;
		grid-row: 3 span;
		grid-template-rows: subgrid;
		gap: 1rem;

		border: 2px dashed var(--black);
		border-radius: var(--border-radius-md);
		padding: 2rem 1rem;
		transition: all 150ms ease-in-out;

		&:hover,
		&:focus-within {
			border: 2px solid var(--black);
			box-shadow: var(--box-shadow-lg);
			transform: translateX(10px) translateY(-10px);
		}
	}

	.card-thumbnail {
		order: 1;
		max-height: 15rem;
		margin: auto;
	}

	.card-content {
		order: 3;
		padding: 0 0.5rem;
	}

	.card-title-wrapper {
		order: 2;
		display: inline-block;
		transition: all 150ms 150ms ease-in-out;
		border: 2px solid transparent;
		border-radius: var(--border-radius-sm);
		color: var(--black);
		margin: auto 0;
		padding: 0.25rem 0.5rem;

		&:has(> .card-title:hover) {
			border: 2px solid var(--black);
			text-decoration-style: solid;

			box-shadow: var(--box-shadow-sm);
		}
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
		display: inline;

		&:hover,
		&:focus,
		&:focus-visible {
			outline: none;
			box-shadow: none;
		}
	}
</style>

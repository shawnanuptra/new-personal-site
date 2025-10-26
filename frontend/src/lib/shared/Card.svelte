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
		background-color: var(--white);

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

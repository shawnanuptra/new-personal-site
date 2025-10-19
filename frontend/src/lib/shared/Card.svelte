<script lang="ts">
	import type { Project } from '$lib/types/api';

	const { project }: { project: Project } = $props();
</script>

<article class="card">
	<div>
		<div class="card-title-wrapper">
			<a class="card-title" href={'project/' + project.slug}>{project.title}</a>
		</div>
		<p>{project.description}</p>
	</div>
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
		width: 100%;
		display: flex;
		flex-direction: column-reverse;
		border: 2px dashed var(--black);
		border-radius: var(--border-radius-md);
		padding: 1rem 1.25rem;
		transition: all 150ms ease-in-out;
		justify-content: space-between;

		&:hover,
		&:focus-within {
			border: 2px solid var(--black);
			box-shadow: var(--box-shadow-lg);
			transform: translateX(10px) translateY(-10px);
		}
	}

	.card-title-wrapper {
		display: inline-block;
		transition: all 150ms 300ms ease-in-out;
		border: 2px solid transparent;
		border-radius: var(--border-radius-md);
		padding: 0.5rem;

		&:has(> .card-title:hover) {
			background-color: orange;

			border: 2px solid var(--black);
			text-decoration-style: solid;

			margin-bottom: 1rem;
			box-shadow: var(--box-shadow);
		}
	}

	.card-title {
		font-size: 2.5rem;
		font-weight: bold;
		text-decoration-style: dashed;

		/* stretch the hitbox */
		&::after {
			position: absolute;
			content: '';
			inset: 0;
			z-index: 1;
		}

		padding: 0.25rem;
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

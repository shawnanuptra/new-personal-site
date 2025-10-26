<script lang="ts">
	import Card from '$lib/shared/Card.svelte';

	const { data } = $props();
	const { projects } = data;
</script>

<main>
	<section>
		<h2>Projects</h2>
		<div class="cards">
			{#each projects as project (project.slug)}
				<Card {project} />
			{/each}
		</div>
	</section>

	<section>
		<h2>Blogs</h2>
		<div class="cards">
			{#each projects as project (project.slug)}
				<p>{project.title}</p>
			{/each}
		</div>
	</section>
</main>

<style>
	main {
		display: flex;
		flex-direction: column;
		gap: 2rem;
		margin: 0 auto;
		max-width: var(--max-width-xlg);
	}

	section {
		display: flex;
		flex-direction: column;
		flex-shrink: 1;
		width: 100%;
		padding: 2rem;
		border: 2px solid black;
		box-shadow: var(--box-shadow);
		border-radius: var(--border-radius-md);
		background-color: var(--white);
		gap: 1.5rem;

		h2 {
			font-size: 3rem;
		}
	}

	.cards {
		display: grid;
		grid-template-columns: repeat(auto-fit, minmax(400px, 1fr));
		gap: 1.5rem;

		grid-template-rows: auto 1fr auto;
	}
</style>

<script lang="ts">
	const { data } = $props();
	const { projects } = data;
</script>

<main>
	<section>
		<h1>Projects</h1>
		{#each projects as project (project.slug)}
			<p>{project.title}</p>
		{/each}
	</section>
</main>

<style>
	main {
		margin: 0 auto;
		border: 1px solid black;
		display: flex;
		max-width: var(--max-width-xlg);
	}
</style>

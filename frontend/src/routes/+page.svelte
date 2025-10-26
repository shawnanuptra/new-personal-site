<script lang="ts">
	import BlogCard from '$lib/shared/BlogCard.svelte';
	import Button from '$lib/shared/Button.svelte';
	import Card from '$lib/shared/Card.svelte';

	const { data } = $props();
	const { projects, blogs } = data;
</script>

<main>
	<section>
		<h2 id="projects">Projects</h2>
		<div class="cards">
			{#each projects as project (project.slug)}
				<Card {project} />
			{/each}
		</div>

		<div class="section-cta">
			<Button label="See more projects &#8594;" to="/projects" size="large" />
		</div>
	</section>

	<section>
		<h2 id="blogs">Blogs</h2>
		<div class="blog-cards">
			{#each blogs as blog (blog.slug)}
				<BlogCard {blog} />
			{/each}
		</div>

		<div class="section-cta">
			<Button label="See more blogs &#8594;" to="/blogs" size="large" />
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
		border: 3px solid var(--black);
		box-shadow: var(--box-shadow);
		border-radius: var(--border-radius-md);
		background-color: var(--white);
		gap: 1.5rem;

		h2 {
			font-size: 3rem;
			scroll-margin-top: 10rem;
		}
	}

	.cards {
		display: grid;
		grid-template-columns: repeat(auto-fit, minmax(400px, 1fr));
		gap: 1.5rem;

		grid-template-rows: auto 1fr auto;
	}

	.section-cta {
		margin: 3rem auto;
	}

	.blog-cards {
		display: grid;
		grid-template-columns: 1fr;
		gap: 1.5rem;
	}
</style>

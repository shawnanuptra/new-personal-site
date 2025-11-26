<script lang="ts">
	const { data } = $props();

	const { project } = data;
</script>

<main>
	<section class="container">
		{@html project.markdownContent}
	</section>
</main>

<style>
	main {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
	}

	.container {
		max-width: var(--max-width-xlg);
	}

	section {
		display: flex;
		flex-direction: column;
		flex-shrink: 1;
		width: 100%;
		padding: 4rem 4rem;
		border: 3px solid var(--black);
		box-shadow: var(--box-shadow);
		border-radius: var(--border-radius-md);
		background-color: var(--white);

		:global(pre:has(code)) {
			padding: 2rem;
			border-radius: var(--border-radius-md);
		}

		:global(p > code) {
			font-family: 'JetBrains Mono', monospace;
			font-optical-sizing: auto;
			font-weight: 400;
			font-style: italic;
			font-size: 0.9rem;
			background-color: #2d353b;
			padding: 0.25ch 0.5ch;
			word-break: break-word;
			color: #d3c6aa;
			border-radius: var(--border-radius-xs);
		}

		:global(ol) {
			list-style: number inside;
			padding-inline-start: 1em;
		}

		:global(li) {
			margin-bottom: 0;
		}

		:global(*) {
			margin-bottom: 1rem;
		}

		:global(h1) {
			font-size: 4.5rem;
		}

		:global(h2) {
			font-size: 3rem;
		}
	}
</style>

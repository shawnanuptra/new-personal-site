<script lang="ts">
	import Button from '$lib/shared/Button.svelte';
</script>

<header class="header">
	<nav class="navbar">
		<ul class="navbar__list">
			<li class="navbar__list-item">
				<Button label="Home" to="/" variant="secondary" />
			</li>
			<li class="navbar__list-item">
				<Button label="Projects" to="/projects" variant="secondary" />
			</li>
			<li class="navbar__list-item">
				<Button label="Blogs" to="/blogs" variant="secondary" />
			</li>
			<li class="navbar__list-item">
				<Button label="Today I Learned" to="/today-i-learned" variant="secondary" />
			</li>
		</ul>
	</nav>
</header>

<style>
	.header {
		margin: 2rem auto;
		max-width: var(--max-width-xlg);
		position: sticky;
		top: 2rem;
	}

	.navbar {
		border: 1px solid var(--black);
		background: var(--white);
		padding: 1rem 1.5rem;
		border-radius: var(--border-radius-md);
		box-shadow: var(--box-shadow);
	}

	.navbar__list {
		display: flex;
		flex-direction: row;
		list-style: none;
		padding: 0;
		margin: 0;
		gap: 1rem;
	}

	.navbar__list-item:first-of-type {
		margin-right: auto;
	}
</style>

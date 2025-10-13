<script lang="ts">
	import favicon from '$lib/assets/favicon.svg';
	import '$lib/styles/globals.css';
	import Header from '$lib/shared/Header.svelte';

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<Header />
{@render children?.()}

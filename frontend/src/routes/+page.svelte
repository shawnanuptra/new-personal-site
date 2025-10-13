<script lang="ts">
	import type { PageProps } from './$types';
	import Button from '$lib/shared/Button.svelte';

	let { data }: PageProps = $props();
</script>

<Button label="hello" variant="secondary" />

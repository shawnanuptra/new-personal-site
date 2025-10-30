<script lang="ts">
	const { data } = $props();

	const { blog } = data;
</script>

<h1>{JSON.stringify(blog)}</h1>

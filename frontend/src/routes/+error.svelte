<script lang="ts">
	import { page } from '$app/state';
</script>

<h1>{page.status}</h1>
<p><strong>{page.error?.message}</strong></p>

<script lang="ts">
	interface propType {
		label: string;
		variant: "primary" | "secondary" | "tertiary";
	}

	let { label, variant = "primary" }: propType = $props();
</script>

<button class="button button--{variant}">
	{label}
</button>

<style>
	.button {
		padding: 0.5rem 1rem;
		border-radius: 0.5rem;
		border: 1px solid transparent;
		cursor: pointer;
	}

	/* Variants */
	.button--primary {
		border-color: black;
		background-color: black;
		color: white;

		&:hover {
			background-color: white;
			color: black;
		}
	}

	.button--secondary {
		border-color: black;
		background-color: white;
		color: black;

		&:hover {
			background-color: black;
			color: white;
		}
	}

	.button--tertiary {
		border-color: white;
		background-color: white;
		color: black;

		&:hover {
			background-color: white;
			color: black;
		}
	}
</style>

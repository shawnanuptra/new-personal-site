<script module>
	import { defineMeta } from '@storybook/addon-svelte-csf';
	import Card from './Card.svelte';
	import { fn } from 'storybook/test';

	const mockProject = {
		title: 'Project Title',
		slug: 'project-title',
		description: "This is the card's description",
		thumbnailUrl:
			'https://thunderdungeon.com/wp-content/uploads/2025/04/the-lion-memes-18-20250415.jpg',
		publishedAt: '2025-08-03',
		content: "This shouldn't be here!",
	};

	const { Story } = defineMeta({
		title: 'Shared/Card',
		component: Card,
		tags: ['autodocs'],
		argTypes: {
			project: { control: 'object' },
		},
		args: {
			onclick: fn(),
		},
	});
</script>

<Story name="Single Card" args={{ project: mockProject }} />

<Story name="Grid of Cards" asChild>
	<div style="display: grid; grid-template-columns: repeat(3, 1fr); gap: 20px;">
		<Card
			project={{
				title: 'Project Title',
				slug: 'project-title',
				description: "This is the card's description",
				thumbnailUrl:
					'https://thunderdungeon.com/wp-content/uploads/2025/04/the-lion-memes-18-20250415.jpg',
				publishedAt: '2025-08-03',
				content: "This shouldn't be here!",
			}}
		/>
		<Card
			project={{
				title: 'Project Title Longer Longer',
				slug: 'project-title',
				description:
					"This is the card's descriptionThis shouldn't be here!This shouldn't be here!This shouldn't be here!This shouldn't be here!This shouldn't be here!This shouldn't be here!This shouldn't be here!",
				thumbnailUrl:
					'https://thunderdungeon.com/wp-content/uploads/2025/04/the-lion-memes-18-20250415.jpg',
				publishedAt: '2025-08-03',
				content: "This shouldn't be here!",
			}}
		/>
	</div>
</Story>

<script module>
	import { defineMeta } from '@storybook/addon-svelte-csf';
	import BlogCard from './BlogCard.svelte';
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
		title: 'Shared/BlogCard',
		component: BlogCard,
		tags: ['autodocs'],
		argTypes: {
			project: { control: 'object' },
		},
		args: {
			onclick: fn(),
		},
	});
</script>

<Story name="Single Card" args={{ blog: mockProject }} />

<Story name="Grid of Cards" asChild>
	<div style="display: grid; grid-template-columns: repeat(3, 1fr); gap: 20px;">
		<BlogCard
			blog={{
				title: 'Project Title',
				slug: 'project-title',
				description: "This is the card's description",
				publishedAt: '2025-08-03',
				content: "This shouldn't be here!",
			}}
		/>
		<BlogCard
			blog={{
				title: 'Project Title Longer Longer',
				slug: 'project-title',
				series: 'aweseom paweseom',
				entry: 1,
				description:
					"This is the card's descriptionThis shouldn't be here!This shouldn't be here!This shouldn't be here!This shouldn't be here!This shouldn't be here!This shouldn't be here!This shouldn't be here!",
				publishedAt: '2025-08-03',
				content: "This shouldn't be here!",
			}}
		/>
	</div>
</Story>

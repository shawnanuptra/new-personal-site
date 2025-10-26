<script module>
	import { defineMeta } from '@storybook/addon-svelte-csf';
	import Card from './Card.svelte';
	import { fn } from 'storybook/test';

	// More on how to set up stories at: https://storybook.js.org/docs/writing-stories
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

<Story
	name="Project"
	args={{
		project: {
			title: 'Project Title',
			slug: 'project-title',
			description: "This is the card's description",
			thumbnailUrl:
				'https://thunderdungeon.com/wp-content/uploads/2025/04/the-lion-memes-18-20250415.jpg',
			publishedAt: '2025-08-03',
			series: 'project-test',
			entry: 1,
			content: "This shouldn't be here!",
		},
	}}
/>

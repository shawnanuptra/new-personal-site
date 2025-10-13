<script module>
	import { defineMeta } from '@storybook/addon-svelte-csf';
	import Header from './Header.svelte';
	import { fn } from 'storybook/test';

	// More on how to set up stories at: https://storybook.js.org/docs/writing-stories
	const { Story } = defineMeta({
		title: 'Shared/Header',
		component: Header,
		tags: ['autodocs'],
		argTypes: {
			label: { control: 'text' },
		},
		args: {
			onclick: fn(),
		},
	});
</script>

<Story name="Header" />

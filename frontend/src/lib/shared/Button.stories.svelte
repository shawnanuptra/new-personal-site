<script module>
	import { defineMeta } from '@storybook/addon-svelte-csf';
	import Button from './Button.svelte';
	import { fn } from 'storybook/test';

	// More on how to set up stories at: https://storybook.js.org/docs/writing-stories
	const { Story } = defineMeta({
		title: 'Shared/Button',
		component: Button,
		tags: ['autodocs'],
		argTypes: {
			label: { control: 'text' },
			variant: { options: ['primary', 'secondary', 'tertiary'] },
		},
		args: {
			onclick: fn(),
		},
	});
</script>

<Story name="Primary" args={{ label: 'Button', variant: 'primary' }} />
<Story name="Secondary" args={{ label: 'Button', variant: 'secondary' }} />

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
			variant: { options: ['primary', 'secondary'] },
			size: { options: ['small', 'large'] },
		},
		args: {
			onclick: fn(),
		},
	});
</script>

<Story name="Primary Button" args={{ label: 'Button' }} />
<Story name="Secondary Button" args={{ label: 'Button', variant: 'secondary' }} />

<Story name="Primary Link Button" args={{ label: 'Button', to: 'google.com' }} />

<Story
	name="Secondary Link Button"
	args={{ label: 'Button', variant: 'secondary', to: 'google.com' }}
/>

<Story name="Large Primary Button" args={{ label: 'Button', size: 'large' }} />
<Story
	name="Large Secondary Button"
	args={{ label: 'Button', variant: 'secondary', size: 'large' }}
/>

<Story
	name="Large Primary Link Button"
	args={{ label: 'Button', to: 'google.com', size: 'large' }}
/>

<Story
	name="Large Secondary Link Button"
	args={{ label: 'Button', variant: 'secondary', to: 'google.com', size: 'large' }}
/>

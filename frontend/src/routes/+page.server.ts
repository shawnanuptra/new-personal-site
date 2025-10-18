import type { ApiResponse, Projects } from '$lib/types/api';
import type { PageServerLoad } from './$types';
import { error as errorRedirect } from '@sveltejs/kit'

export const load: PageServerLoad = async ({ fetch }) => {
	const response = await fetch('/api/projects');

	const { data, error }: ApiResponse<Projects> = await response.json();

	if (!error) {
		const { projects } = data;

		return {
			projects,
		};
	}

	// redirect if api call is error
	errorRedirect(500, { message: error })
}

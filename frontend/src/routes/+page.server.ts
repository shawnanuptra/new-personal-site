import type { ApiResponse, Projects } from '$lib/types/api';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch }) => {
	const response = await fetch('api/projects');
	const { projects }: Projects = await response.json();

	return {
		projects,
	};
};

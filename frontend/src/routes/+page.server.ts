import type { ApiResponse, Projects } from '$lib/types/api';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch }) => {
	const response = await fetch('api/projects');
	const { data }: ApiResponse<Projects> = await response.json();

	const { projects } = data;

	return {
		projects,
	};
};

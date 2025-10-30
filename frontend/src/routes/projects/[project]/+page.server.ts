import type { PageServerLoad } from "./$types";
import type { ApiResponse, Project } from "$lib/types/api";
import { error as errorRedirect } from '@sveltejs/kit';

export const load: PageServerLoad = async ({ fetch, params }) => {
	const projectResponse = await fetch('/api/projects/' + params.project)
	const projectJson = await projectResponse.json()

	const { data: project, error: projectErr }: ApiResponse<Project> = projectJson;

	if (!projectErr) {
		return {
			project
		};
	}

	// redirect if api call is error
	errorRedirect(500, { message: projectErr })
}

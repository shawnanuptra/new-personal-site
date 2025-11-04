import type { ApiResponse, ProjectResponse } from "$lib/types/api";
import type { PageServerLoad } from "./$types";
import { error as errorRedirect } from '@sveltejs/kit';
import { getMarked } from "$lib/server/markdown";

export const load: PageServerLoad = async ({ fetch, params }) => {
	const marked = await getMarked();

	const projectResponse = await fetch('/api/projects/' + params.project)
	const projectJson = await projectResponse.json()

	const { data: { project }, error: projectErr }: ApiResponse<ProjectResponse> = projectJson;

	project.markdownContent = await marked.parse(project.markdownContent);

	if (!projectErr) {
		return {
			project,
		};
	}

	// redirect if api call is error
	errorRedirect(500, { message: projectErr })
}

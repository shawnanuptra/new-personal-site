import type { ApiResponse, BlogsResponse, ProjectsResponse } from '$lib/types/api';
import type { PageServerLoad } from './$types';
import { error as errorRedirect } from '@sveltejs/kit'

export const load: PageServerLoad = async ({ fetch }) => {
	const [projectsResponse, blogsResponse] = await Promise.all([
		fetch('/api/projects'),
		fetch('/api/blogs'),
	]);

	const [projectsJson, blogsJson] = await Promise.all([
		projectsResponse.json(),
		blogsResponse.json(),
	])

	const { data: projectsData, error: projectsErr }: ApiResponse<ProjectsResponse> = projectsJson;
	const { data: blogsData, error: blogsErr }: ApiResponse<BlogsResponse> = blogsJson;

	if (!projectsErr && !blogsErr) {
		const { projects } = projectsData;
		const { blogs } = blogsData;

		return {
			projects,
			blogs
		};
	}

	// redirect if api call is error
	errorRedirect(500, { message: projectsErr + '\n' + blogsErr })
}

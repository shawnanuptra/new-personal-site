import type { ApiResponse, Blog, BlogsResponse } from "$lib/types/api";
import type { PageServerLoad } from "./$types";
import { error as errorRedirect } from '@sveltejs/kit'

export const load: PageServerLoad = async ({ fetch, params }) => {
	const blogResponse = await fetch('/api/blogs/' + params.blog)
	const blogJson = await blogResponse.json()

	const { data: blog, error: blogErr }: ApiResponse<Blog> = blogJson;

	if (!blogErr) {
		return {
			blog
		};
	}

	// redirect if api call is error
	errorRedirect(500, { message: blogErr })
}

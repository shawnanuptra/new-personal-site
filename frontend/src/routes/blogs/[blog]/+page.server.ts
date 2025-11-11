import type { ApiResponse, BlogResponse } from "$lib/types/api";
import type { PageServerLoad } from "./$types";
import { error as errorRedirect } from '@sveltejs/kit';
import { getMarked } from "$lib/server/markdown";

export const load: PageServerLoad = async ({ fetch, params }) => {
	const marked = await getMarked();

	const blogResponse = await fetch('/api/blogs/' + params.blog)
	const blogJson = await blogResponse.json()

	const { data: { blog }, error: blogErr }: ApiResponse<BlogResponse> = blogJson;

	blog.markdownContent = await marked.parse(blog.markdownContent);

	if (!blogErr) {
		return {
			blog
		};
	}

	// redirect if api call is error
	errorRedirect(500, { message: blogErr })
}

export type ApiResponse<T> = {
	data: T;
	error?: any;
}

export type ProjectsResponse = {
	projects: Project[];
}

export type ProjectResponse = {
	project: Project;
}

export type Project = {
	title: string;
	slug: string;
	description: string;
	thumbnailUrl: string;
	publishedAt: string;
	content: string;
	markdownContent: string;
}

export type BlogsResponse = {
	blogs: Blog[];
}

export type BlogResponse = {
	blog: Blog;
}

export type Blog = {
	title: string;
	slug: string;
	description: string;
	publishedAt: string;
	series?: string;
	entry?: number;
	markdownContent: string;
}

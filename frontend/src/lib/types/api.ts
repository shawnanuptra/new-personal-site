export type ApiResponse<T> = {
	data: T;
	error?: any;
}

export type Projects = {
	projects: Project[];
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

export type Blogs = {
	blogs: Blog[]
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

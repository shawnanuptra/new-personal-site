export interface ApiResponse<T> {
	data: T;
	error?: any;
}

export interface Projects {
	projects: Project[];
}

export interface Project {
	title: string;
	slug: string;
	description: string;
	thumbnailUrl: string;
	publishedAt: string;
	series: string;
	entry: number;
	content: string;
}

import type { RequestHandler } from './$types';
import { SERVER_HOST } from '$env/static/private';

export const GET: RequestHandler = ({ params, url }) => {
	return fetch(`${SERVER_HOST}/projects${url.search}`);
};

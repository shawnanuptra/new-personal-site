import type { RequestHandler } from './$types';

export const GET: RequestHandler = ({ params, url }) => {
    return fetch('http://localhost:8080/ping');
};

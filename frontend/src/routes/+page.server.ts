import type { PageLoad } from './$types';

export const load: PageLoad = async ({ params }) => {

    const response = await fetch('http://localhost:8080/ping')
    const { message: ping } = await response.json()

    return {
        ping,
    };
};

import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch }) => {

    const response = await fetch('api/ping')
    const { message: ping } = await response.json()

    return {
        ping,
    };
};

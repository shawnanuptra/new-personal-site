import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch }) => {

    const response = await fetch('api/projects')
    const { projects } = await response.json()

    return {
        projects,
    };
};

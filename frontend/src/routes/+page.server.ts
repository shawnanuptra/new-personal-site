import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch }) => {

    const response = await fetch('api/projects')
    const { error } = await response.json()

    console.log({ error })
    // const { projects, error } = await response.json()


    console.log('here')
    return {
        projects: [],
    };
};

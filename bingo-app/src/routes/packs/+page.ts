import type { PageLoad } from './$types';

export const load = (async ({ fetch }) => {
  const res = await fetch('https://taskbingo.com/api/task/getPacks', {
        method: 'GET',
        headers: {'Origin': 'taskbingo.com'},
    })

    const packs = await res.json()

    return { packs }
}) satisfies PageLoad;
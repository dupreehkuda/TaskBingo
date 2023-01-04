import type { PageLoad } from './$types';

export const load = (async ({ fetch }) => {
  const res = await fetch('https://taskbingo.com/api/user/getUserData', {
        method: 'POST',
        headers: {'Origin': 'taskbingo.com'},
    })

    const userInfo = await res.json()

    return {
        userInfo
    }
}) satisfies PageLoad;
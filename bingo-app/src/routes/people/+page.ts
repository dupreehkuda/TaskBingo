import type { PageLoad } from './$types';

export const load = (async ({ fetch }) => {
  const res = await fetch('https://taskbingo.com/api/user/getAllUsers', {
    method: 'GET',
    headers: {'Origin': 'taskbingo.com'},
  })

  const users = await res.json()

  return { users }
}) satisfies PageLoad;
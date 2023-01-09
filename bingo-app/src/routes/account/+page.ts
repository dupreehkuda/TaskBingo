import type { PageLoad } from './$types';
import Account from '../accountStore'

export const load = (async ({ fetch }) => {
  const res = await fetch('https://taskbingo.com/api/user/getUserData', {
        method: 'POST',
        headers: {'Origin': 'taskbingo.com'},
    })

    const userInfo = await res.json()

    Account.set(userInfo)
}) satisfies PageLoad;
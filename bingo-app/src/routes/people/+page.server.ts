import Account from '../accountStore';
import type { PageServerLoad } from './$types';


export const load = (async ({ fetch }) => {
  const userData = await fetch('https://taskbingo.com/api/user/getUserData', {
    method: 'GET',
    headers: {'Origin': 'taskbingo.com'},
  })

  const userInfo = await userData.json()

  Account.set(userInfo)

  const res = await fetch('https://taskbingo.com/api/user/getAllUsers', {
    method: 'GET',
    headers: {'Origin': 'taskbingo.com'},
  })

  const users = await res.json()

  return { users }
}) satisfies PageServerLoad;
import Account from '../accountStore';
import type { PageLoad } from './$types';

const API_URL = import.meta.env.VITE_API_URL;
const WEB_URL = import.meta.env.VITE_WEB_URL;
export const ssr = false

export const load = (async ({ fetch }) => {
  const userData = await fetch(`${API_URL}/api/user/getUserData`, {
    method: 'GET',
    headers: {'Origin': WEB_URL},
    credentials: 'include',
  })

  const userInfo = await userData.json()

  Account.set(userInfo)

  const res = await fetch(`${API_URL}/api/user/getAllUsers`, {
    method: 'GET',
    headers: {'Origin': WEB_URL},
    credentials: 'include',
  })

  const users = await res.json()

  return { users }
}) satisfies PageLoad;
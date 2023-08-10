import Account from '../accountStore';
import { get } from 'svelte/store';
import type { PageLoad } from '../../../.svelte-kit/types/src/routes/game/$types';
import { API_URL, WEB_URL } from '../temporary';
export const ssr = false

export const load = (async ({ fetch }) => {
  const userData = await fetch(`${API_URL}/api/user/getUserData`, {
    method: 'GET',
    headers: {'Origin': WEB_URL},
    credentials: 'include',
  })

  const userInfo = await userData.json()

  Account.set(userInfo)

  const res = await fetch(`${API_URL}/api/task/getRatedPacks`, {
    method: 'GET',
    headers: {'Origin': WEB_URL},
    credentials: 'include',
  })

  const packs = await res.json()

  return { packs }
}) satisfies PageLoad;

export async function _Like(pack: any, liked: boolean) {
  const newReq = {
    id: pack.id,
  } 

  if (liked) {
    const res = await fetch(`${API_URL}/api/user/dislikePack`, {
      method: 'POST',
      headers: {'Origin': WEB_URL},
      body: JSON.stringify(newReq),
      credentials: 'include',
    })

    let account = get(Account)
    account.likedPacks = account.likedPacks.filter(e => e.id !== pack.id)
    Account.set(account)

  } else {
    const res = await fetch(`${API_URL}/api/user/likePack`, {
      method: 'POST',
      headers: {'Origin': WEB_URL},
      body: JSON.stringify(newReq),
      credentials: 'include',
    })

    if (res.ok) {
      let account = get(Account)
      account.likedPacks.push({id: pack.id, pack: pack.pack})
      Account.set(account)
    }
  }
};

export async function _Rate(pack: any, rated: boolean) {
  const newReq = {
    id: pack.id,
  } 

  if (rated) {
    const res = await fetch(`${API_URL}/api/user/unratePack`, {
      method: 'POST',
      headers: {'Origin': WEB_URL},
      body: JSON.stringify(newReq),
      credentials: 'include',
    })

    let account = get(Account)
    account.ratedPacks = account.ratedPacks.filter(e => e !== pack.id)
    Account.set(account)

  } else {
    const res = await fetch(`${API_URL}/api/user/ratePack`, {
      method: 'POST',
      headers: {'Origin': WEB_URL},
      body: JSON.stringify(newReq),
      credentials: 'include',
    })

    let account = get(Account)
    account.ratedPacks.push(pack.id)
    Account.set(account)
  }
};

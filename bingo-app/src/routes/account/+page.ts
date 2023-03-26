import type { PageLoad } from './$types';
import Account from '../accountStore';
import { get } from 'svelte/store';

export const load = (async ({ fetch }) => {
  const res = await fetch('https://taskbingo.com/api/user/getUserData', {
    method: 'GET',
    headers: {'Origin': 'taskbingo.com'},
  })

  const userInfo = await res.json()

  Account.set(userInfo)
}) satisfies PageLoad;

export async function _Like(pack: any, liked: boolean) {
  const newResp = {
    id: pack.id,
  } 

  if (liked) {
    const res = await fetch('https://taskbingo.com/api/user/dislikePack', {
      method: 'POST',
      headers: {'Origin': 'taskbingo.com'},
      body: JSON.stringify(newResp)
    })

    let account = get(Account)
    account.likedPacks = account.likedPacks.filter(e => e.id !== pack.id)
    Account.set(account)

  } else {
    const res = await fetch('https://taskbingo.com/api/user/likePack', {
      method: 'POST',
      headers: {'Origin': 'taskbingo.com'},
      body: JSON.stringify(newResp)
    })

    let account = get(Account)
    account.likedPacks.push({id: pack.id, pack: pack.pack})
    Account.set(account)
  }
};

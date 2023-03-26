import Account from '../accountStore';
import { get } from 'svelte/store';
import type { PageLoad } from './$types';

export const load = (async ({ fetch }) => {
  const userData = await fetch('https://taskbingo.com/api/user/getUserData', {
    method: 'GET',
    headers: {'Origin': 'taskbingo.com'},
  })

  const userInfo = await userData.json()

  Account.set(userInfo)

  const res = await fetch('https://taskbingo.com/api/task/getPacks', {
    method: 'GET',
    headers: {'Origin': 'taskbingo.com'},
  })

  const packs = await res.json()
  console.log(packs)
  return { packs }
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

export async function _Rate(pack: any, rated: boolean) {
  const newResp = {
    id: pack.id,
  } 

  if (rated) {
    const res = await fetch('https://taskbingo.com/api/user/unratePack', {
      method: 'POST',
      headers: {'Origin': 'taskbingo.com'},
      body: JSON.stringify(newResp)
    })

    let account = get(Account)
    account.ratedPacks = account.ratedPacks.filter(e => e !== pack.id)
    Account.set(account)

  } else {
    const res = await fetch('https://taskbingo.com/api/user/ratePack', {
      method: 'POST',
      headers: {'Origin': 'taskbingo.com'},
      body: JSON.stringify(newResp)
    })

    let account = get(Account)
    account.ratedPacks.push(pack.id)
    Account.set(account)
  }
};

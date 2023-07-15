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

export async function _LikePack(pack: any, liked: boolean) {
  const newReq = {
    id: pack.id,
  } 

  if (liked) {
    const res = await fetch('https://taskbingo.com/api/user/dislikePack', {
      method: 'POST',
      headers: {'Origin': 'taskbingo.com'},
      body: JSON.stringify(newReq)
    })

    if (res.ok) {
      let account = get(Account)
      account.likedPacks = account.likedPacks.filter(e => e.id !== pack.id)
      Account.set(account)
    }
  } else {
    const res = await fetch('https://taskbingo.com/api/user/likePack', {
      method: 'POST',
      headers: {'Origin': 'taskbingo.com'},
      body: JSON.stringify(newReq)
    })

    if (res.ok) {
      let account = get(Account)
      account.likedPacks.push({id: pack.id, pack: pack.pack})
      Account.set(account)
    }
  }
};

export async function _DeleteGame(gameID: string) {
  const newReq = { gameID: gameID } 

  const res = await fetch('https://taskbingo.com/api/game/delete', {
      method: 'DELETE',
      headers: {'Origin': 'taskbingo.com'},
      body: JSON.stringify(newReq)
    })

    if (res.ok) {
      let account = get(Account)
      account.games = account.games.filter(e => e.gameId !== gameID)
      Account.set(account)
    }
}
import Account, { type AccountData } from '../accountStore';
import CurrentGame from '../currentGame';
import { get } from 'svelte/store';
import { goto } from '$app/navigation';
import type { PageLoad } from '../../../.svelte-kit/types/src/routes/game/$types';
import { API_URL, WEB_URL } from '../temporary';
import { markLoggedOut } from '$lib/stores/auth';
export const ssr = false

export const load = (async ({ fetch }) => {
  // User data is best-effort: if the session is gone (cookie pointing at a
  // wiped DB, expired JWT, network blip), we still render the page with the
  // public packs so the user is not stuck on a 500.
  try {
    const userData = await fetch(`${API_URL}/api/user/getUserData`, {
      method: 'GET',
      headers: {'Origin': WEB_URL},
      credentials: 'include',
    })

    if (userData.ok) {
      const userInfo: AccountData = await userData.json()
      // Backend may serialise empty slices as null — normalise once.
      userInfo.friends ??= []
      userInfo.likedPacks ??= []
      userInfo.ratedPacks ??= []
      userInfo.packs ??= []
      userInfo.games ??= []
      Account.set(userInfo)
    } else if (userData.status === 401) {
      markLoggedOut()
    }
  } catch {
    // Backend down — keep going with no account info.
  }

  const res = await fetch(`${API_URL}/api/task/getRatedPacks`, {
    method: 'GET',
    headers: {'Origin': WEB_URL},
    credentials: 'include',
  })

  const packs = res.ok ? await res.json() : []
  return { packs: packs ?? [] }
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

export async function _StartSolo(packID: string) {
  const res = await fetch(`${API_URL}/api/game/solo/start`, {
    method: 'POST',
    headers: { 'Origin': WEB_URL, 'Content-Type': 'application/json' },
    body: JSON.stringify({ packID }),
    credentials: 'include',
  })

  if (!res.ok) {
    return
  }

  const { gameID, numbers } = await res.json()
  const account = get(Account)

  CurrentGame.set({
    gameID,
    user1ID: account.userID,
    user2ID: '',
    packID,
    status: 1,
    user1Bingo: 0,
    user2Bingo: 0,
    winner: '',
    numbers,
    user1Numbers: new Array(16).fill(0),
    user2Numbers: new Array(16).fill(0),
    kind: 'solo',
  })

  goto('/game?solo=true')
}

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

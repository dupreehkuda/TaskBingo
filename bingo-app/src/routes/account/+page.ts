import type { PageLoad } from './$types';
import Account, { type AccountData } from '../accountStore';
import CurrentGame from '../currentGame';
import { get } from 'svelte/store';

const API_URL = import.meta.env.VITE_API_URL;
const WEB_URL = import.meta.env.VITE_WEB_URL;

export const load = (async ({ fetch }) => {
  const res = await fetch(`${API_URL}/api/user/getUserData`, {
    method: 'GET',
    headers: {'Origin': WEB_URL, 'Access-Control-Allow-Origin': '*'},
  })

  const userInfo: AccountData = await res.json()

  const newReq = {
    ids: userInfo.games
        .map(obj => obj.packId)
        .filter(packId => !userInfo.likedPacks.some(obj => obj.id === packId)),
  } 

  if (newReq.ids.length !== 0) {
    const neededPacks = await fetch('${API_URL}/api/task/getTaskPacks', {
      method: 'POST',
      headers: {'Origin': WEB_URL, 'Access-Control-Allow-Origin': '*'},
      body: JSON.stringify(newReq)
    })
  
    userInfo.packs = await neededPacks.json()
  }

  Account.set(userInfo)
}) satisfies PageLoad;

export async function _LikePack(pack: any, liked: boolean) {
  const newReq = {
    id: pack.id,
  } 

  if (liked) {
    const res = await fetch('${API_URL}/api/user/dislikePack', {
      method: 'POST',
      headers: {'Origin': WEB_URL},
      body: JSON.stringify(newReq)
    })

    if (res.ok) {
      let account = get(Account)
      account.likedPacks = account.likedPacks.filter(e => e.id !== pack.id)
      Account.set(account)
    }
  } else {
    const res = await fetch('${API_URL}/api/user/likePack', {
      method: 'POST',
      headers: {'Origin': WEB_URL},
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

  const res = await fetch('${API_URL}/api/game/delete', {
      method: 'DELETE',
      headers: {'Origin': WEB_URL},
      body: JSON.stringify(newReq)
    })

    if (res.ok) {
      let account = get(Account)
      account.games = account.games.filter(e => e.gameId !== gameID)
      Account.set(account)
    }
}

export async function _GetGame(gameID: string) {
    let currentGame = get(CurrentGame)
    if (currentGame !== null) {
      if (currentGame?.gameID === gameID) {
        return
      }
    }

    const newReq = {
      gameID: gameID,
    } 

    const res = await fetch('${API_URL}/api/game/get', {
        method: 'POST',
        headers: {'Origin': WEB_URL},
        body: JSON.stringify(newReq)
    })

    if (res.ok) {
      const gameInfo = await res.json()
      CurrentGame.set(gameInfo)
    }
}
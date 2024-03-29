import Account, { type Game } from '../../routes/accountStore';
import { get } from 'svelte/store';
import { API_URL, WEB_URL } from '../../routes/temporary';

export async function CreateGame(opponent: string, pack: string) {
    const newReq = {
        opponent: opponent,
        pack: pack,
    } 
    
    const res = await fetch(`${API_URL}/api/game/create`, {
      method: 'POST',
      headers: {'Origin': WEB_URL},
      body: JSON.stringify(newReq),
      credentials: 'include',
    })

    const gameInfo = await res.json()
    if (res.ok) {
      let account = get(Account)
      account.games.push(gameInfo)
      Account.set(account)
    }
};
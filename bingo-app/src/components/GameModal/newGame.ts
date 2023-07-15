import Account, { type Game } from '../../routes/accountStore';
import { get } from 'svelte/store';

export async function CreateGame(opponent: string, pack: string) {
    const newReq = {
        opponent: opponent,
        pack: pack,
    } 
    
    const res = await fetch('https://taskbingo.com/api/game/create', {
      method: 'POST',
      headers: {'Origin': 'taskbingo.com'},
      body: JSON.stringify(newReq)
    })

    const gameInfo = await res.json()
    if (res.ok) {
      let account = get(Account)
      account.games.push(gameInfo)
      Account.set(account)
    }
};
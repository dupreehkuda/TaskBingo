import Account, { type TaskPack } from '../../routes/accountStore';
import { get } from 'svelte/store';

export function GetTaskPack(packID: string): TaskPack {
    let account = get(Account)

    let idx = account.likedPacks.findIndex(pack => pack.id === packID)

    if (idx !== -1) {
        return account.likedPacks[idx]
    }

    idx = account.packs.findIndex(pack => pack.id === packID)

    return account.packs[idx]
}

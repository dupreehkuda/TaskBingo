import { writable } from "svelte/store"

export interface AccountData {
    login:      string;
    city:       string;
    wins:       number;
    lose:       number;
    bingo: number;
    friends:    Friend[];
    packs:      Pack[];
}

export interface Friend {
    login: string;
    city:  string;
    bingo: number;
}

export interface Pack {
    id:  string;
    tasks: string[];
}

const Account = writable<AccountData>();

export default Account;
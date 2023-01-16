import { writable } from "svelte/store"

export interface AccountData {
    login:      string;
    city:       string;
    wins:       number;
    lose:       number;
    bingo:      number;
    friends:    Friend[];
    likedPacks: Pack[];
    ratedPacks: string[]
}

export interface Friend {
    login: string;
    status: number;
    wins: number;
    loses: number;
}

export interface Pack {
    id:  string;
    tasks: string[];
}

const Account = writable<AccountData>();

export default Account;
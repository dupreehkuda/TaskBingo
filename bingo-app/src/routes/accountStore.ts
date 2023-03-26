import { writable } from "svelte/store"

export interface AccountData {
    userID:     string;
    username:   string;
    city:       string;
    wins:       number;
    lose:       number;
    bingo:      number;
    friends:    Friend[];
    likedPacks: Pack[];
    ratedPacks: string[]
}

export interface Friend {
    userID: string;
    username: string;
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
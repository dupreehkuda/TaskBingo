import { writable } from "svelte/store"

export interface AccountData {
    userID:     string;
    username:   string;
    city:       string;
    wins:       number;
    lose:       number;
    bingo:      number;
    friends:    Friend[];
    likedPacks: TaskPack[];
    ratedPacks: string[];
    games: Game[];
}

export interface Game {
    gameId: string;
    user1Id: string;
    user2Id: string;
    packId: string;
    status: number;
    user1Bingo: number;
    user2Bingo: number;
    winner: string;
}

export interface Friend {
    userID: string;
    username: string;
    status: number;
    wins: number;
    loses: number;
}

export interface TaskPack {
    id: string;
    pack: Pack;
}

export interface Pack {
    title:  string;
    tasks: string[];
}

const Account = writable<AccountData>();

export default Account;
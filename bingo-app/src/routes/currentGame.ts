import { writable } from "svelte/store"

export interface Game {
    gameID: string,
    user1ID: string,
    user2ID: string,
    packID: string,
    status: number,
    user1Bingo: number,
    user2Bingo: number,
    winner: string,
    numbers: number[],
    user1Numbers: number[],
    user2Numbers: number[],
}

const CurrentGame = writable<Game>();

export default CurrentGame;
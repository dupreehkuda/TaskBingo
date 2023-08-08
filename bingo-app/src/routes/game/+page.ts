import Account from '../accountStore';
import CurrentGame from '../currentGame';
import { get } from 'svelte/store';

export interface gameUpdate {
    status: number;
    userID: string;
    bingo: number;
    numbers: number[];
}

let update: gameUpdate;

export function _GameHandler(): { socket: WebSocket, closer: () => void } {
    let account = get(Account)
    let game = get(CurrentGame)
    // document.cookie = 'X-Authorization=' + authToken + '; path=/';
    console.log(game)
    const connection = '?game=' + game.gameID + '&user=' + account.userID
    const socket = new WebSocket('wss://taskbingo.com/api/game/start' + connection);

    socket.onopen = () => {   
        console.log('WebSocket connection established');
        sendInitial(socket)
    };

    socket.onmessage = (event) => {
        update = JSON.parse(event.data);
        processUpdate(update)
    };

    socket.onclose = () => {
        console.log('WebSocket connection closed');
    };

    let closer: () => void = function closer() {
        socket.close();
    }

    return { socket, closer };
}

function sendInitial(socket: WebSocket) {
    let account = get(Account)

    let update = {
        userID: account.userID,
        finished: false,
        numbers:[0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
    }

    if (socket && socket.readyState === WebSocket.OPEN) {
        socket.send(JSON.stringify(update));
    }
    
    return
}

export function _SendUpdate(socket: WebSocket, newNum: number, close: boolean) {
    let account = get(Account)
    let game = get(CurrentGame)

    placeNumber(newNum)

    let update = {
        userID: account.userID,
        finished: close ? true : false,
        numbers: (account.userID === game.user1ID) ? game.user1Numbers : game.user2Numbers,
    }

    if (socket && socket.readyState === WebSocket.OPEN) {
        socket.send(JSON.stringify(update));
    }
    
    return
}

function processUpdate(update: gameUpdate) {
    let game = get(CurrentGame)

    game.status = update.status

    if (update.numbers === null || update.userID === "") {
        return
    }
    
    if (update.userID === game.user1ID) {
        game.user1Numbers = update.numbers
        game.user1Bingo = update.bingo
    } else if (update.userID === game.user2ID) {
        game.user2Numbers = update.numbers
        game.user2Bingo = update.bingo
    }

    CurrentGame.set(game)

    return
}

function placeNumber(num: number) {
    let account = get(Account)
    let game = get(CurrentGame)

    const index = game.numbers.indexOf(num);
    
    if (index === -1) {
      return
    }

    if (account.userID === game.user1ID) {
        game.user1Numbers[index] = num
    } else {
        game.user2Numbers[index] = num
    }

    CurrentGame.set(game)

    return
}


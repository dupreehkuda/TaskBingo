import { goto } from '$app/navigation';
import Account from '../accountStore';
import CurrentGame from '../currentGame';
import { get } from 'svelte/store';

const WS_URL = (import.meta.env.VITE_WS_URL) ? import.meta.env.VITE_WS_URL : 'wss://taskbingo.com';

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
    const connection = '?game=' + game.gameID + '&user=' + account.userID
    const socket = new WebSocket(`${WS_URL}/api/game/start` + connection);

    socket.onopen = () => {   
        sendInitial(socket)
    };

    socket.onmessage = (event) => {
        update = JSON.parse(event.data);
        processUpdate(update)
    };

    socket.onclose = () => {
        _RedirectOnAccount()
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
    placeNumber(newNum)

    let account = get(Account)
    let game = get(CurrentGame)

    if (account.userID === game.user1ID) game.user1Bingo = countBingo(game.user1Numbers);
    else game.user2Bingo = countBingo(game.user2Numbers);

    let update = {
        userID: account.userID,
        finished: close ? true : false,
        numbers: (account.userID === game.user1ID) ? game.user1Numbers : game.user2Numbers,
    }

    CurrentGame.set(game)

    if (socket && socket.readyState === WebSocket.OPEN) {
        socket.send(JSON.stringify(update));
    }
    
    return
}

function processUpdate(update: gameUpdate) {
    let game = get(CurrentGame)

    if (update.status === 6) {
        _RedirectOnAccount()
        return
    }

    game.status = update.status

    if (update.numbers === null || update.userID === "") {
        CurrentGame.set(game)
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

export function _RedirectOnAccount() {
  goto('/account');
}

function placeNumber(num: number) {
    let account = get(Account)
    let game = get(CurrentGame)

    const index = game.numbers.indexOf(num);
    
    if (index === -1) {
      return
    }

    if (account.userID === game.user1ID) {
        if (game.user1Numbers[index] === num) {
            game.user1Numbers[index] = 0
        } else {
            game.user1Numbers[index] = num
        }
    } else {
        if (game.user2Numbers[index] === num) {
            game.user2Numbers[index] = 0
        } else {
            game.user2Numbers[index] = num
        }
    }

    CurrentGame.set(game)

    return
}

// TODO think of doing this on backend for both players. temporary thing
function countBingo(numbers: number[]): number {
    let bingos = 0;
  
    for (let i = 0; i < 4; i++) {
      if (
        numbers[i] !== 0 &&
        numbers[i + 4] !== 0 &&
        numbers[i + 8] !== 0 &&
        numbers[i + 12] !== 0
      ) {
        bingos += 1;
      }
    }
  
    for (let i = 0; i < numbers.length; i += 4) {
      if (
        numbers[i] !== 0 &&
        numbers[i + 1] !== 0 &&
        numbers[i + 2] !== 0 &&
        numbers[i + 3] !== 0
      ) {
        bingos += 1;
      }
    }
  
    if (
      numbers[0] !== 0 &&
      numbers[5] !== 0 &&
      numbers[10] !== 0 &&
      numbers[15] !== 0
    ) {
      bingos += 1;
    }
  
    if (
      numbers[3] !== 0 &&
      numbers[6] !== 0 &&
      numbers[9] !== 0 &&
      numbers[12] !== 0
    ) {
      bingos += 1;
    }
  
    return bingos;
  }

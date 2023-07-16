import Account from '../accountStore';
import CurrentGame from '../currentGame';
import { get } from 'svelte/store';

export interface gameUpdate {
    status: number;
    userID: string;
    numbers: number[];
}

let update: gameUpdate;

export function _GameHandler(): () => void {
    let account = get(Account)
    let game = get(CurrentGame)
    
    // document.cookie = 'X-Authorization=' + authToken + '; path=/';

    const connection = '?game=' + game.gameID + '&user=' + account.userID
    const socket = new WebSocket('wss://taskbingo.com/api/game/start' + connection);

    socket.onopen = () => {   
        console.log('WebSocket connection established');
    };

    socket.onmessage = (event) => {
        update = event.data;
        console.log(update);
    };

    socket.onclose = () => {
        console.log('WebSocket connection closed');
    };

    let closer: () => void = function closer() {
        socket.close();
    }

    socket.send(JSON.stringify(sendUpdate()))

    return closer;
}

function sendUpdate() {
    let account = get(Account)
    let game = get(CurrentGame)

    let update = {
        userID: account.userID,
        finished: false,
        numbers:[0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
    }

    return update
}
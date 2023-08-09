<script lang="ts">
    import {_GameHandler, _SendUpdate} from "./+page"
    import { onDestroy, onMount } from 'svelte';
    import CurrentGame from '../currentGame';
    import {get, type Unsubscriber} from "svelte/store";
    import Grid from '../../components/Grid/Grid.svelte';
	import Account from "../accountStore";
	import Keypad from "../../components/Keypad/Keypad.svelte";
    import { Progressbar } from "flowbite-svelte";
	import TasksCard from "../../components/TasksCard/TasksCard.svelte";

    let gameHandler: { socket: WebSocket, closer: () => void };
    let socket: WebSocket
    let closer: () => void;
    let account = get(Account)
    let game = get(CurrentGame)

    let unsubscribe: Unsubscriber;
    let usersBingo: number;
    let opponentsBingo: number;

    onMount(() => {
        if (!gameHandler) {
            gameHandler = _GameHandler();
        }

        socket = gameHandler.socket;
        closer = gameHandler.closer;

        unsubscribe = CurrentGame.subscribe(value => {
            usersBingo = value.user1ID === account.userID ? value.user1Bingo : value.user2Bingo
            opponentsBingo = value.user1ID === account.userID ? value.user2Bingo : value.user1Bingo
        });
    });
    
    onDestroy(() => {
        unsubscribe();
        closer();
    });

    window.addEventListener('beforeunload', function() {
        closer();
    })

    let submittedValue: number;

    function handleNewNumberSubmit(event: any) {
        submittedValue = event.detail;

        if (submittedValue !== 0 && Number(submittedValue) <= 16) {
            _SendUpdate(socket, submittedValue, false)
        }
    }
</script>

<svelte:head>
    <title>Game</title>
</svelte:head>

<main>
    <!-- {#if $CurrentGame.status < 3}
        <h3 class="fonty">Waiting for the opponent</h3>
    {:else} -->
        <body class="grid grid-cols-3 gap-4">
            <div class="col-span-2 p-5"><Grid numbers={game.numbers}/></div>
            
            <div class="justify-center">
                <div class="top-spacer scrolling-wrapper spacer allWidth">
                    <div class="rectangle flex flex-row">
                        <div class="lefty basis-1/12"></div>
                        <div class="lefty basis-5/12">
                            <span class="text-lg personText font-medium">You</span>
                        </div>
                        <div class="space-x-1 flex flex-row justify-end basis-5/12">
                            <span class="text-lg personText font-medium">{usersBingo}</span>
                        </div>
                        <div class="lefty basis-1/12"></div>
                    </div>
                    <div class="rectangle flex flex-row">
                        <div class="lefty basis-1/12"></div>
                        <div class="lefty basis-5/12">
                            <span class="text-lg personText font-medium">Opponent</span>
                        </div>
                        <div class="space-x-1 flex flex-row justify-end basis-5/12">
                            <span class="text-lg personText font-medium">{opponentsBingo}</span>
                        </div>
                        <div class="lefty basis-1/12"></div>
                    </div>
                </div>
                
                <Keypad on:submit={handleNewNumberSubmit}/>

                <TasksCard packID={game.packID}/>
            </div>
        </body>
    <!-- {/if}     -->
</main>

<style>
    .top-spacer {
        margin-top: 7%; /* adjust the value as needed */
    }
    span {
        text-align: left;
        font-weight: 300;
    }

    .personText {
        color: #112a41
    }

    main {
        text-align: center;
        max-width: 60%;
        margin: 0 auto;
    }

    .rectangle {
        border-radius: 10px; 
        background-color: #ffffff;
        margin-bottom: 0.5em;
        padding: 0.3em;
    }
</style>

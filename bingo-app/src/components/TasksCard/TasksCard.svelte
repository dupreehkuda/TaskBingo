<script lang="ts">
      import { onDestroy, onMount } from "svelte";
      import { GetTaskPack } from "./packData";
      import CurrentGame from "../../routes/currentGame";
      import { get, type Unsubscriber } from "svelte/store";
      import Account from "../../routes/accountStore";

    export let packID = "";

    let usersTasks: number[] = $CurrentGame.user1ID === $Account.userID ? $CurrentGame.user1Numbers : $CurrentGame.user2Numbers;
    let opponentsTasks: number[] = $CurrentGame.user1ID === $Account.userID ? $CurrentGame.user2Numbers : $CurrentGame.user1Numbers;

    let unsubscribe: Unsubscriber;
    let account = get(Account)

    onMount(() => {
        unsubscribe = CurrentGame.subscribe(value => {
            usersTasks = value.user1ID === account.userID ? value.user1Numbers : value.user2Numbers
            opponentsTasks = value.user1ID === account.userID ? value.user2Numbers : value.user1Numbers
        });
    });

    onDestroy(() => {
        unsubscribe();
    });
</script>

<div class="top-spacer flex flex-col w-full bg-white shadow-lg rounded-lg p-6">
    <h1 class="text-2xl text-center cardText mb-4">{GetTaskPack(packID).pack.title}</h1>
    {#each GetTaskPack(packID).pack.tasks as task, i}
        <li class="flex flex-row leftspace">
            <span class="basis-1/12 leading-tight cardText">{i+1}</span>
            <div class="flex basis-2/12">
                {#if usersTasks.includes(i+1)}
                    <div class="w-3 h-3 rounded-full done mr-2"></div>
                {:else}
                    <div class="w-3 h-3 rounded-full notDone mr-2"></div>
                {/if}
                {#if opponentsTasks.includes(i+1)}
                    <div class="w-3 h-3 rounded-full done mr-2"></div>
                {:else}
                    <div class="w-3 h-3 rounded-full notDone mr-2"></div>
                {/if}
            </div>
            <span class="basis-9/12 leading-tight cardText" class:completed="{usersTasks.includes(i+1)}">{task}</span>
        </li>
    {/each}
</div>

<style>
    .completed {
        text-decoration: line-through;
    }

    .top-spacer {
        margin-top: 10%;
    }

    span {
        text-align: left;
        font-weight: 300;
    }

    .cardText {
        font-family: Prompt;
        color: #112a41
    }

    .notDone {
        margin-top: 6%;
        background-color: #acddff;
    }

    .done {
        margin-top: 6%;
        background-color: #db7676;
    }
</style>
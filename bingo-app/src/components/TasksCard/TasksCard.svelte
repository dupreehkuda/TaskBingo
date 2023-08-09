<script lang="ts">
	import { onDestroy, onMount } from "svelte";
	import type { TaskPack } from "../../routes/accountStore";
	import { GetTaskPack } from "./packData";
	import CurrentGame from "../../routes/currentGame";
	import { get, type Unsubscriber } from "svelte/store";
	import Account from "../../routes/accountStore";

    export let packID = "";

    let usersTasks: number[] = $CurrentGame.user1ID === $Account.userID ? $CurrentGame.user1Numbers : $CurrentGame.user2Numbers;;
    let unsubscribe: Unsubscriber;
    let account = get(Account)

    onMount(() => {
        unsubscribe = CurrentGame.subscribe(value => {
            usersTasks = value.user1ID === account.userID ? value.user1Numbers : value.user2Numbers
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
            <span class="basis-1/5 leading-tight cardText">{i+1}</span>
            <span class="basis-4/5 leading-tight cardText" class:completed="{usersTasks.includes(i+1)}">{task}</span>
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
</style>
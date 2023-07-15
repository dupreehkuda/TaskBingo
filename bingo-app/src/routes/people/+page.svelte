<script lang="ts">
    import type { PageData } from './$types';
    import { Button } from "flowbite-svelte";
    import Account from "../accountStore";
	import { DeleteFriend, RequestFriend, AcceptFriend } from '../friendRequests';
    import GameModal from '../../components/GameModal/GameModal.svelte';

    let showModal = false;
    let selectedPackID: string;
    let selectedFriendID: string;

    export let data: PageData;
    const { users } = data
</script>

<svelte:head>
    <title>People</title>
</svelte:head>

<main>
    <div class="scrolling-wrapper spacer allWidth">
            {#each users as user}
                <div class="rectangle flex flex-row">
                    <div class="lefty basis-1/12"><span class="text-lg personText">{user.bingo}</span></div>
                    <div class="lefty basis-7/12"><span class="text-lg personText font-medium">{user.username} <span class="mt-2 mb-4 text-xs">{user.city}</span></span></div>

                    {#if $Account?.friends.some(e => e.userID === user.userID && e.status === 3)}
                        <div class="space-x-1 flex flex-row justify-end basis-4/12">
                            <Button size="xs" on:click={() => (showModal = true, selectedFriendID = user.userID)}>Play</Button>
                            <Button size="xs" color="red" class="dark:!text-white-800" on:click={() => DeleteFriend(user.userID)}>X</Button>
                        </div>
                    {:else if $Account?.friends.some(e => e.userID === user.userID && e.status === 1)}
                        <div class="space-x-1 flex flex-row justify-end basis-4/12">
                            <Button disabled size="xs">Sent</Button>
                            <Button size="xs" color="red" class="dark:!text-white-800" on:click={() => DeleteFriend(user.userID)}>X</Button>
                        </div> 
                    {:else if $Account?.friends.some(e => e.userID === user.userID && e.status === 2)}
                        <div class="space-x-1 flex flex-row justify-end basis-4/12">
                            <Button size="xs" on:click={() => AcceptFriend(user.userID)}>Accept</Button>
                            <Button size="xs" color="red" class="dark:!text-white-800" on:click={() => DeleteFriend(user.userID)}>X</Button>
                        </div>   
                    {:else if user.userID == $Account?.userID}
                        <div class="space-x-1 flex flex-row justify-end basis-4/12">
                            <Button href="/account" size="xs">Account</Button>
                        </div>
                    {:else}
                        <div class="space-x-1 flex flex-row justify-end basis-4/12">
                            <Button size="xs" on:click={() => RequestFriend(user.userID, user.username)}>Add friend</Button>
                        </div>
                    {/if}
                    
                </div>
            {/each}
    </div>


    <GameModal bind:showModal {selectedFriendID} {selectedPackID}/>
</main>

<style>
    span {
        text-align: left;
        font-weight: 300;
    }

    .personText {
        color: #112a41
    }

    main {
        /* min-width: 50%; */
        text-align: center;
        max-width: fit-content;
        margin: 0 auto;
    }

    .rectangle {
        border-radius: 10px; 
        background-color: #e8e8e6;
        margin-bottom: 0.5em;
        padding: 0.6em;
        min-width: 23em; 
    }
</style>
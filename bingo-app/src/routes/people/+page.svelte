<script lang="ts">
    import type { PageData } from './$types';
    import { Button } from "flowbite-svelte";
    import Account from "../accountStore";
	import { DeleteFriend, RequestFriend, AcceptFriend } from './+page';

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
                    <div class="lefty basis-1/12"><span class="text-md whiteBingoText">{user.bingo}</span></div>
                    <div class="lefty basis-8/12"><span class="text-lg font-medium">{user.login} <span class="mt-2 mb-4 text-sm">{user.city}</span></span></div>

                    {#if $Account?.friends.some(e => e.login === user.login && e.status === 3)}
                        <div class="basis-3/12">
                            <Button size="xs">Play</Button>
                            <Button size="xs" color="red" class="dark:!text-white-800" on:click={() => DeleteFriend(user.login)}>X</Button>
                        </div>
                    {:else if $Account?.friends.some(e => e.login === user.login && e.status === 1)}
                        <div class="basis-3/12">
                            <Button disabled size="xs">Requested</Button>
                            <Button size="xs" color="red" class="dark:!text-white-800" on:click={() => DeleteFriend(user.login)}>X</Button>
                        </div> 
                    {:else if $Account?.friends.some(e => e.login === user.login && e.status === 2)}
                        <div class="basis-3/12">
                            <Button size="xs" on:click={() => AcceptFriend(user.login)}>Accept</Button>
                            <Button size="xs" color="red" class="dark:!text-white-800" on:click={() => DeleteFriend(user.login)}>X</Button>
                        </div>   
                    {:else if user.login == $Account?.login}
                        <div class="basis-3/12">
                            <Button href="/account" size="xs">Account</Button>
                        </div>
                    {:else}
                        <div class="basis-3/12">
                            <Button size="xs" on:click={() => RequestFriend(user.login)}>Add friend</Button>
                        </div>
                    {/if}
                    
                </div>
            {/each}
    </div>
</main>

<style>
    span {
        text-align: left;
        font-weight: 300;
    }

    .whiteBingoText {
        color: aliceblue;
    }

    main {
        /* min-width: 50%; */
        text-align: center;
        max-width: fit-content;
        margin: 0 auto;
    }

    .rectangle {
        border-radius: 10px; 
        background-color: #0f4879;
        margin-bottom: 0.5em;
        padding: 0.6em;
        min-width: 23em; 
    }
</style>
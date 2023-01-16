<script lang="ts">
    import { Button } from "flowbite-svelte";
    import { Like } from "./+page"
    import { DeleteFriend, RequestFriend, AcceptFriend } from '../friendRequests';

    import Account from "../accountStore";
</script>

<svelte:head>
    <title>Account</title>
</svelte:head>

<main>
    <div class="leftalign spacer">
        <h1>{$Account.login}
            <span class="leftalign win">{$Account.wins}</span>
            <span>:</span>
            <span class="lose">{$Account.lose}</span>
        </h1>
        
        <h4>{$Account.city}</h4>

        <div class="leftalign">
            {#if $Account.wins === 0 && $Account.lose === 0}
                <span class="leftalign">Here will be your winrate when you play at least one game</span>
            {:else}
                <span class="leftalign">Winrate: {$Account.wins === 0 && $Account.lose === 0 ? "null" : ($Account.wins / ($Account.wins + $Account.lose)) * 100 + "%"}</span>
                <span>Total bingo: {$Account.bingo === 0 ? "null" : $Account.bingo}</span>
            {/if}
        </div>
    </div>

    <div class="leftspaceheading"><h3 class="mb-2">Friends</h3></div>
    <div class="leftspace scrolling-wrapper gap-1.5">
        {#each $Account.friends as friend}
            <div class="friend flex flex-col">
                <div class="basis-1/3">
                    <h5>{friend.login}</h5>
                </div>

                <div class="basis-1/3">
                    <span class="text-xs text-gray-300">{friend.wins}/{friend.loses}</span>
                </div>
                {#if friend.status == 3}
                    <div class="flex flex-row gap-1.5 basis-1/3">
                        <Button class="basis-2/3 fonty" size="xs">Play</Button>
                        <Button class="basis-1/3 dark:!text-white-800" size="xs" color="red" on:click={() => DeleteFriend(friend.login)}>X</Button>
                    </div>
                {:else if friend.status == 2}
                    <div class="flex flex-row gap-1.5 basis-1/3">
                        <Button class="basis-2/3 fonty" size="xs" on:click={() => AcceptFriend(friend.login)}>Accept</Button>
                        <Button class="basis-1/3 dark:!text-white-800" size="xs" color="red" on:click={() => DeleteFriend(friend.login)}>X</Button>
                    </div>
                {:else if friend.status == 1}
                    <div class="flex flex-row gap-1.5 basis-1/3">
                        <Button class="basis-2/3 fonty" disabled size="xs" on:click={() => AcceptFriend(friend.login)}>Sent</Button>
                        <Button class="basis-1/3 dark:!text-white-800" size="xs" color="red" on:click={() => DeleteFriend(friend.login)}>X</Button>
                    </div>
                {/if}
            </div>
        {/each}
    </div>

    <div class="leftspaceheading"><h3 class="mb-2">Packs</h3></div>
    <div class="flex leftalign spacer05">
        <Button href="/newpack" class="fonty">Create new pack</Button>
    </div>
    <div class="scrolling-wrapper spacer">
        {#each $Account.likedPacks as pack}
            <div class="pack flex flex-col justify-between mx-1">
                <h5 class="mb-2 text-xl">{pack.id}</h5>
                <ul class="my-1 space-y-1.5">
                    {#each pack.tasks as task, i}
                        <li class="flex flex-row leftspace">
                            <span class="basis-1/5 leading-tight">{i+1}</span>
                            <span class="basis-4/5 leading-tight text-gray-300">{task}</span>
                        </li>
                    {/each}

                    <div class="flex flex-row gap-2">
                        <Button class="basis-4/5 fonty">Choose pack</Button>
                        <Button class="basis-1/5" color="light" on:click={() => Like(pack, $Account?.likedPacks.some(e => e.id === pack.id))}>
                            {#if $Account?.likedPacks.some(e => e.id === pack.id)}
                                <img src="heart-solid.svg" alt="solid heart"/>
                            {:else}
                                <img src="heart-regular.svg" alt="regular heart"/>
                            {/if}
                        </Button>
                    </div>
                </ul>
            </div>
        {/each}
    </div>
</main>

<style>
    span {
        text-align: left;
        font-weight: 300;
    }

    h1 {
        text-align: left;
        font-family: Unbounded;
        font-weight: 400;
        margin-left: 0.35em;
        font-size: xx-large;
    }

    h3 {
        text-align: left;
        margin-left: 0.35em;
        font-family: Unbounded;
    }

    h4 {
        text-align: left;
        margin-left: 0.70em;
        font-family: Unbounded;
        font-weight: 200;
        margin-bottom: 1em;
    }

    .spacer {
        margin-bottom: 2em;
    }

    .win {
        font-weight: 300;
        color: #36e804;
    }

    .lose {
        font-weight: 300;
        color: #eb1a0f;
    }
    main {
        min-width: 50%;
        text-align: center;
        max-width: max-content;
        margin: 0 auto;
    }

    .pack {
        border-radius: 10px; 
        background-color: #0f4879;
        margin-bottom: 0.5em;
        padding: 0.6em;
        min-width: 22em; 
    }

    .friend {
        border-radius: 10px; 
        background-color: #0f4879;
        margin-bottom: 1em;
        padding: 0.5em;
        padding-top: 1em;
        min-height: 10em;
        min-width: 10em;
    }
</style>

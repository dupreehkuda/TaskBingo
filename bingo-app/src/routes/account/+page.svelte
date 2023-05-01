<script lang="ts">
    import { Button } from "flowbite-svelte";
    import { _Like } from "./+page"
    import { DeleteFriend, RequestFriend, AcceptFriend } from '../friendRequests';

    import Account from "../accountStore";
</script>

<svelte:head>
    <title>Account</title>
</svelte:head>

<main>
    <div class="leftalign spacer">
        <h1>{$Account.username}
            <!-- <span class="leftalign win">{$Account.wins}</span>
            <span>:</span>
            <span class="lose">{$Account.lose}</span> -->
            <span class="leftalign">{$Account.bingo}</span>
        </h1>
        
        <h4>{$Account.city}</h4>

        <!-- <div class="leftalign">
            {#if $Account.wins === 0 && $Account.lose === 0}
                <span class="leftalign">Here will be your winrate when you play at least one game</span>
            {:else}
                <span class="leftalign">Winrate: {$Account.wins === 0 && $Account.lose === 0 ? "null" : ($Account.wins / ($Account.wins + $Account.lose)) * 100 + "%"}</span>
                <span>Total bingo: {$Account.bingo === 0 ? "null" : $Account.bingo}</span>
            {/if}
        </div> -->
    </div>

    <div class="leftspaceheading"><h3 class="mb-2">Friends</h3></div>
    <div class="leftspace scrolling-wrapper gap-1.5">
        {#each $Account.friends as friend}
            <div class="friend flex flex-col">
                <div class="basis-5/12 cardText">
                    <h5>{friend.username}</h5>
                </div>

                <div class="basis-4/12">
                    <span class="text-xs cardText">{friend.wins}/{friend.loses}</span>
                </div>
                <div class="flex flex-row gap-1.5 basis-3/12">
                    {#if friend.status == 3}
                        <Button class="basis-2/3 fonty" size="xs">Play</Button>
                        <Button class="basis-1/3 dark:!text-white-800" size="xs" color="red" on:click={() => DeleteFriend(friend.userID)}>X</Button>
                    {:else if friend.status == 2}
                        <Button class="basis-2/3 fonty" size="xs" on:click={() => AcceptFriend(friend.userID)}>Accept</Button>
                        <Button class="basis-1/3 dark:!text-white-800" size="xs" color="red" on:click={() => DeleteFriend(friend.userID)}>X</Button>
                    {:else if friend.status == 1}
                        <Button class="basis-2/3 fonty" disabled size="xs" on:click={() => AcceptFriend(friend.userID)}>Sent</Button>
                        <Button class="basis-1/3 dark:!text-white-800" size="xs" color="red" on:click={() => DeleteFriend(friend.userID)}>X</Button>
                    {/if}
                </div>
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
                <h5 class="mb-2 text-xl cardText">{pack.pack.title}</h5>
                <ul class="my-1 space-y-1.5">
                    {#each pack.pack.tasks as task, i}
                        <li class="flex flex-row leftspace">
                            <span class="basis-1/5 leading-tight cardText">{i+1}</span>
                            <span class="basis-4/5 leading-tight cardText">{task}</span>
                        </li>
                    {/each}

                    <div class="flex flex-row gap-2">
                        <Button class="basis-4/5 fonty">Choose pack</Button>
                        <Button class="basis-1/5" color="light" on:click={() => _Like(pack, $Account?.likedPacks.some(e => e.id === pack.id))}>
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
        font-family: Prompt;
        font-weight: 400;
        margin-left: 0.35em;
        font-size: xx-large;
    }

    h3 {
        font-weight: 400;
        font-size: large;
        text-align: left;
        margin-left: 0.35em;
        font-family: Prompt;
    }

    h4 {
        text-align: left;
        margin-left: 0.70em;
        font-family: Prompt;
        font-weight: 200;
        margin-bottom: 1em;
    }

    h5 {
        font-weight: 400;
        font-size: large;
    }

    .spacer {
        margin-bottom: 2em;
    }

    .cardText {
        color: #112a41
    }

    /* .win {
        font-weight: 300;
        color: #32ca08;
    }

    .lose {
        font-weight: 300;
        color: #d8170d;
    } */
    main {
        min-width: 50%;
        text-align: center;
        max-width: max-content;
        margin: 0 auto;
    }

    .pack {
        border-radius: 10px; 
        background-color: #e8e8e6;
        margin-bottom: 0.5em;
        padding: 0.6em;
        min-width: 22em; 
    }

    .friend {
        border-radius: 10px; 
        background-color: #e8e8e6;
        margin-bottom: 1em;
        padding: 0.5em;
        padding-top: 1em;
        min-height: 10em;
        min-width: 10em;
    }
</style>

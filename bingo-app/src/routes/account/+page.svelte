<script lang="ts">
    import { Button } from "flowbite-svelte";
    import { _LikePack, _DeleteGame, _GetGame } from "./+page"
    import { DeleteFriend, AcceptFriend } from '../friendRequests';
    import Account from "../accountStore";
    import CurrentGame from "../currentGame";
    import { get } from 'svelte/store';
    import GameModal from '../../components/GameModal/GameModal.svelte';

	let showModal = false;
    let selectedPackID: string;
    let selectedFriendID: string;

    function getOpponentUsername(userId: string) {
        for (const friend of $Account.friends) {
            if (friend.userID === userId) {
                return friend.username
            }
        }
    }

    function getPackTitle(packID: string) {
        let idx = $Account.likedPacks.findIndex(pack => pack.id === packID)

        if (idx !== -1) {
            return $Account.likedPacks[idx].pack.title
        }

        idx = $Account.packs.findIndex(pack => pack.id === packID)

        if (idx !== -1) {
            return $Account.packs[idx].pack.title
        }

        return ""
    }
</script>

<svelte:head>
    <title>Account</title>
</svelte:head>

<main>
    <div class="leftalign spacer">
        <h1>{$Account.username}
            <span class="leftalign">{$Account.bingo}</span>
        </h1>
        
        <h4>{$Account.city}</h4>
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
                        <Button class="basis-2/3 fonty" size="xs" on:click={() => (showModal = true, selectedFriendID = friend.userID)}>Play</Button>
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
                        <Button class="basis-4/5 fonty" on:click={() => (showModal = true, selectedPackID = pack.id)}>Choose pack</Button>
                        <Button class="basis-1/5" color="light" on:click={() => _LikePack(pack, $Account?.likedPacks.some(e => e.id === pack.id))}>
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
    <div class="leftspaceheading"><h3 class="mb-2">Games</h3></div>
    <div class="flex leftalign spacer05">
        <Button class="fonty" on:click={() => (showModal = true)}>Create new game</Button>
    </div>
    <div class="scrolling-wrapper spacer">
        {#each $Account.games as game}
            <div class="game flex flex-col justify-between mx-1">
                {#if $Account.userID === game.user1Id}
                    <h5 class="mb-2 text-xl cardText">{getOpponentUsername(game.user2Id)}</h5>
                {:else}
                    <h5 class="mb-2 text-xl cardText">{getOpponentUsername(game.user1Id)}</h5>
                {/if}
                <span class="text-xs cardText">{getPackTitle(game.packId)}</span>
                <ul class="my-1 space-y-1.5">
                    <div class="flex flex-row gap-2">
                        <Button href="/game" class="basis-4/5 fonty" 
                        on:mouseenter={() => (_GetGame(game.gameId))}>
                            Start
                        </Button>
                        <Button class="basis-1/5 dark:!text-white-800" size="xs" color="red" on:click={() => _DeleteGame(game.gameId)}>X</Button>
                    </div>
                </ul>
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

    .game {
        border-radius: 10px; 
        background-color: #e8e8e6;
        margin-bottom: 0.5em;
        padding: 0.6em;
        min-width: 12em; 
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

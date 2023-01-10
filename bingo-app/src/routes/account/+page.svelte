<script lang="ts">
    import { Button } from "flowbite-svelte";

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
                <span>Scoreboard: {$Account.scoreboard === 0 ? "null" : $Account.scoreboard}</span>
            {/if}
        </div>
    </div>

    <div class="leftspaceheading"><h3 class="mb-2">Friends</h3></div>
    <div class="leftspace scrolling-wrapper">
        {#each $Account.friends as friend}
            <div class="friend flex flex-col">
                <div class="basis-2/3">
                    <h5>{friend.login}</h5>
                    <span class="text-xs text-gray-300">{friend.city}</span>
                </div>
                
                <div class="flex flex-col basis-1/3">
                    <Button class="fonty">Play</Button>
                </div>
            </div>
        {/each}
    </div>

    <div class="leftspaceheading"><h3 class="mb-2">Packs</h3></div>
    <div class="leftspace scrolling-wrapper spacer">
        {#each $Account.packs as pack}
            <div class="pack flex flex-col">
                <h5 class="mb-2 text-xl">{pack.id}</h5>
                <ul class="my-1 space-y-1.5">
                    {#each pack.tasks as task, i}
                        <li class="flex flex-row leftspace">
                            <span class="basis-1/5 leading-tight">{i+1}</span>
                            <span class="basis-4/5 leading-tight text-gray-300">{task}</span>
                        </li>
                    {/each}

                    <div class="flex flex-col">
                        <Button class="fonty">Choose pack</Button>
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
        /* min-width: 23em;  */
    }

    .friend {
        border-radius: 10px; 
        background-color: #0f4879;
        margin-bottom: 1em;
        padding: 0.5em;
        padding-top: 1em;
        min-height: 9em;
        min-width: 10em;
    }
</style>

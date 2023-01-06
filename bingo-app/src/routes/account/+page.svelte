<script lang="ts">
    import { Button, Card } from "flowbite-svelte";
	import { load } from "./+page";
    import { onMount } from "svelte";

    onMount(() => {
        load
    })

    export let data;
    const { userInfo } = data

    console.log(userInfo)
</script>

<svelte:head>
    <title>Account</title>
</svelte:head>

<main>
    <div class="leftspace spacer">
        <h1>{userInfo.login}
            <span class="leftspace win">{userInfo.wins}</span>
            <span>:</span>
            <span class="lose">{userInfo.lose}</span>
        </h1>
        
        <h4>{userInfo.city}</h4>

        <div class="leftspace">
            {#if userInfo.wins === 0 && userInfo.lose === 0}
                <span class="leftspace">Here will be your winrate when you play at least one game</span>
            {:else}
                <span class="leftspace">Winrate: {userInfo.wins === 0 && userInfo.lose === 0 ? "null" : (userInfo.wins / (userInfo.wins + userInfo.lose)) * 100 + "%"}</span>
                <span>Scoreboard: {userInfo.scoreboard === 0 ? "null" : userInfo.scoreboard}</span>
            {/if}
        </div>
    </div>

    <div class="leftspaceheading"><h3 class="mb-2">Friends</h3></div>
    <div class="scrolling-wrapper spacer">
        {#each userInfo.friends as friend}
            <Card padding="sm" class="card">
                <div class="flex flex-col items-center pb-4">
                    <div class="mb-1">
                        <h5>{friend.login}</h5>
                        <span class="text-xs text-gray-300">{friend.city}</span>
                    </div>
                    <span class="text-sm text-gray-300">{friend.score}</span>
                    <div class="flex mt-4 space-x-3 lg:mt-6">
                        <Button class="dark fonty dark:text-white">Play</Button>
                    </div>
                </div>
            </Card>
        {/each}
    </div>

    <div class="leftspaceheading"><h3 class="mb-2">Packs</h3></div>
    <div class="scrolling-wrapper spacer packs">
        {#each userInfo.packs as pack}
          <Card size="md" padding="xl" class="card">
            <h5 class="mb-4 text-xl">{pack.id}</h5>
            <ul class="my-7 space-y-2">
                {#each pack.tasks as task, i}
                    <li class="flex space-x-2">
                        <span class="leading-tight">{i+1}</span>
                        <span class="leading-tight text-gray-300">{task}</span>
                    </li>
                {/each}
            <Button class="fonty dark w-full">Choose pack</Button>
          </Card>
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
</style>

<script lang="ts">
    import type { PageData } from './$types';
    import { Button } from "flowbite-svelte";
    import Account from "../accountStore";

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
                    <div class="lefty basis-3/4"><span class="text-lg font-medium">{user.login} <span class="mt-2 mb-4 text-sm">{user.city}</span></span></div>

                    {#if $Account?.friends.some(e => e.login === user.login)}
                        <div class="basis-1/4">
                            <Button size="xs">Play</Button>
                            <Button size="xs" color="red" class="dark:!text-white-800">X</Button>
                        </div>
                    {:else}
                        <div class="basis-1/4">
                            <Button size="xs">Add friend</Button>
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
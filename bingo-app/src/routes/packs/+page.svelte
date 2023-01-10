<script lang="ts">
    import type { PageData } from './$types';
    import { Button, Card } from "flowbite-svelte";
	import { Like } from "./+page";
    import Account from "../accountStore";

    export let data: PageData;
    const { packs } = data
</script>

<svelte:head>
    <title>Packs</title>
</svelte:head>

<main>
    <div class="scrolling-wrapper spacer">
        {#each packs as pack}
            <div>
                <Card size="sm" padding="sm" class="card">
                    <h5 class="mb-4 text-xl">{pack.id}</h5>
                    <ul class="my-7 space-y-2">
                        {#each pack.tasks as task, i}
                            <li class="flex space-x-2">
                                <span class="leading-tight">{i+1}</span>
                                <span class="leading-tight text-gray-300">{task}</span>
                            </li>
                        {/each}

                        <div>
                            <li style="margin-top:2em" class="flex space-x-2">
                                <Button class="fonty dark w-full">Play</Button>
        
                                <!-- <Button class="fonty dark w-full" on:click={() => Rate(pack.id, rated)} on:click={() => (rated = !rated)}>
                                    {#if rated}
                                        <img src="star-solid.svg" alt="solid star"/>
                                    {:else}
                                        <img src="star-regular.svg" alt="regular star"/>
                                    {/if}
                                </Button> -->
        
                                <Button class="fonty dark w-full" on:click={() => Like(pack, $Account?.packs.some(e => e.id === pack.id))}>
                                    {#if $Account?.packs.some(e => e.id === pack.id)}
                                        <img src="heart-solid.svg" alt="solid heart"/>
                                    {:else}
                                        <img src="heart-regular.svg" alt="regular heart"/>
                                    {/if}
                                </Button>
                            </li>
                        </div>
                </Card>
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
        min-width: 50%;
        text-align: center;
        max-width: max-content;
        margin: 0 auto;
    }

    /* .icons {
        max-block-size: 150%;
    } */

    img {
        size-adjust: 150%;
        block-size: 150%;
    }
</style>
<script lang="ts">
    // import type { PageData } from './$types';
    // import { Button, Card } from "flowbite-svelte";
	// // import { load } from "./+page.server";
    // import { onMount } from "svelte";

    // onMount(() => {
    //     load
    // })

    // export let data: PageData;
    // const { people } = data

</script>

<svelte:head>
    <title>People</title>
</svelte:head>

<main>
    <h5>Nothing here</h5>
    <div class="scrolling-wrapper spacer packs">
        <!-- {#each packs as pack}
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
                    <li style="margin-top:2em" class="flex space-x-2">
                        <Button class="fonty dark w-full">Play</Button>
                        <Button class="fonty dark w-full" on:click={rate}><img src="star-regular.svg" alt="filled star"/></Button>
                        <Button class="fonty dark w-full" on:click={like}><img src="heart-regular.svg" alt="filled star"/></Button>
                    </li>
                </Card>
            </div>
        {/each} -->
    </div>
</main>

<style>
    /* span {
        text-align: left;
        font-weight: 300;
    }
    main {
        min-width: 50%;
        text-align: center;
        max-width: max-content;
        margin: 0 auto;
    }

    img {
        size-adjust: 150%;
        block-size: 150%;
    } */
</style>
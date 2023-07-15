<script lang="ts">
    import type { PageData } from './$types';
    import { Button } from "flowbite-svelte";
	import { _Like, _Rate } from "./+page";
    import Account from "../accountStore";
    import GameModal from '../../components/GameModal/GameModal.svelte';

	let showModal = false;
    let selectedPackID: string;
    let selectedFriendID: string;

    export let data: PageData;
    const { packs } = data
</script>

<svelte:head>
    <title>Packs</title>
</svelte:head>

<main>
    <div class="scrolling-wrapper spacer">
        {#each packs as pack}
            <div class="rectangle flex flex-col justify-between mx-1">
                <h5 class="mb-2 text-xl cardText">{pack.pack.title}</h5>
                <ul class="my-1 space-y-1.5">
                    {#each pack.pack.tasks as task, i}
                        <li class="flex flex-row leftspace">
                            <span class="basis-1/5 leading-tight cardText">{i+1}</span>
                            <span class="basis-4/5 leading-tight cardText">{task}</span>
                        </li>
                    {/each}

                    <div class="flex flex-col">
                        <li class="flex space-x-2">
                            {#if $Account?.likedPacks.some(e => e.id === pack.id)}
                                <Button class="basis-3/5 fonty" on:click={() => (showModal = true, selectedPackID = pack.id)}>Play</Button>
                            {:else}
                                <Button class="basis-3/5 fonty" disabled>Like to Play</Button>
                            {/if}
                            
                            <Button class="basis-1/5" color="light" on:click={() => _Rate(pack, $Account?.ratedPacks.some(e => e === pack.id))}>
                                {#if $Account?.ratedPacks.some(e => e === pack.id)}
                                    <img src="star-solid.svg" alt="solid star"/>
                                {:else}
                                    <img src="star-regular.svg" alt="regular star"/>
                                {/if}
                            </Button>
    
                            <Button class="basis-1/5" color="light" on:click={() => _Like(pack, $Account?.likedPacks.some(e => e.id === pack.id))}>
                                {#if $Account?.likedPacks.some(e => e.id === pack.id)}
                                    <img src="heart-solid.svg" alt="solid heart"/>
                                {:else}
                                    <img src="heart-regular.svg" alt="regular heart"/>
                                {/if}
                            </Button>
                        </li>
                    </div>
                </ul>
            </div>
        {/each}
    </div>

    <GameModal bind:showModal {selectedFriendID} {selectedPackID}/>
</main>

<style>
    h5 {
        font-weight: 400;
    }

    span {
        text-align: left;
        font-weight: 300;
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

    img {
        size-adjust: 150%;
        block-size: 150%;
    }

    .rectangle {
        border-radius: 10px; 
        background-color: #e8e8e6;
        /* margin-bottom: 0.5em; */
        padding: 0.6em;
        min-width: 23em; 
    }
</style>
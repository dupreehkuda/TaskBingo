<script lang="ts">
    import type { PageData } from './$types';
    import Account from '../accountStore';
    import type { TaskPack } from '../accountStore';
    import Page from '$lib/ui/Page.svelte';
    import Stack from '$lib/ui/Stack.svelte';
    import Cluster from '$lib/ui/Cluster.svelte';
    import Paper from '$lib/ui/Paper.svelte';
    import Button from '$lib/ui/Button.svelte';
    import TabBar from '$lib/ui/TabBar.svelte';
    import GameModal from '$lib/feature/GameModal.svelte';
    import { _Like, _Rate, _StartSolo } from './+page';

    export let data: PageData;
    const { packs } = data;

    let tab = 'public';
    let showModal = false;
    let selectedPackID = '';
    let selectedFriendID = '';

    $: publicPacks = (packs ?? []).filter((p: TaskPack) => !p.isPrivate);
    $: liked = (packs ?? []).filter((p: TaskPack) => $Account?.likedPacks.some((e) => e.id === p.id));
    $: mine = (packs ?? []).filter((p: TaskPack) => p.creator === $Account?.userID);

    $: visible = tab === 'liked' ? liked : tab === 'mine' ? mine : publicPacks;

    function play(packID: string) {
        selectedPackID = packID;
        selectedFriendID = '';
        showModal = true;
    }
</script>

<svelte:head><title>Packs · taskbingo</title></svelte:head>

<Page width="normal">
    <Stack gap="l">
        <Cluster gap="m" align="center" justify="between">
            <h1 class="title">Packs</h1>
            <TabBar
                bind:value={tab}
                items={[
                    { value: 'public', label: 'Public' },
                    { value: 'liked', label: 'Liked' },
                    { value: 'mine', label: 'Mine' },
                ]}
                ariaLabel="Pack categories"
            />
        </Cluster>

        <Stack gap="s">
            {#each visible as pack (pack.id)}
                {@const isLiked = $Account?.likedPacks.some((e) => e.id === pack.id) ?? false}
                {@const isRated = $Account?.ratedPacks.some((e) => e === pack.id) ?? false}
                <Paper padding="md">
                    <Stack gap="m">
                        <h3 class="pack-title">{pack.pack.title}</h3>
                        <ol class="tasks">
                            {#each pack.pack.tasks as task, i}
                                <li><span class="num">{i + 1}</span><span>{task}</span></li>
                            {/each}
                        </ol>
                        <Cluster gap="s">
                            {#if isLiked}
                                <Button variant="accent" size="sm" on:click={() => play(pack.id)}>Play</Button>
                            {:else}
                                <Button variant="ghost" size="sm" disabled>Like to play</Button>
                            {/if}
                            <Button variant="ghost" size="sm" on:click={() => _StartSolo(pack.id)}>Solo</Button>
                            <Button variant="ghost" size="sm" on:click={() => _Rate(pack, isRated)} aria-label={isRated ? 'Unrate pack' : 'Rate pack'}>
                                <span aria-hidden="true">{isRated ? '★' : '☆'}</span>
                            </Button>
                            <Button variant="ghost" size="sm" on:click={() => _Like(pack, isLiked)} aria-label={isLiked ? 'Unlike pack' : 'Like pack'}>
                                <span aria-hidden="true">{isLiked ? '♥' : '♡'}</span>
                            </Button>
                        </Cluster>
                    </Stack>
                </Paper>
            {/each}

            {#if visible.length === 0}
                <Paper padding="md">
                    <span class="empty">
                        {#if tab === 'liked'}
                            No liked packs yet — tap the heart on any pack.
                        {:else if tab === 'mine'}
                            <a href="/newpack">Create your first pack →</a>
                        {:else}
                            No packs available.
                        {/if}
                    </span>
                </Paper>
            {/if}
        </Stack>
    </Stack>

    <GameModal bind:showModal {selectedFriendID} {selectedPackID} />
</Page>

<style>
    .title {
        font-family: var(--font-display);
        font-size: 1.6rem;
        font-weight: 400;
        letter-spacing: -0.015em;
    }
    .pack-title {
        font-family: var(--font-display);
        font-size: 1.15rem;
        font-weight: 400;
        letter-spacing: -0.01em;
        color: var(--ink);
        margin: 0;
    }
    .tasks {
        list-style: none;
        margin: 0; padding: 0;
        display: grid;
        grid-template-columns: 1fr 1fr;
        gap: var(--space-1) var(--space-4);
        font-size: 0.85rem;
        color: var(--ink-2);
        line-height: 1.45;
    }
    @media (max-width: 640px) {
        .tasks { grid-template-columns: 1fr; }
    }
    .tasks li { display: flex; gap: var(--space-2); }
    .num { color: var(--ink-4); min-width: 1.5em; }
    .empty { color: var(--ink-2); font-style: italic; font-size: 0.9rem; }
    .empty a { color: var(--ink); border-bottom: 1px solid rgba(40, 55, 95, 0.25); text-decoration: none; }
</style>

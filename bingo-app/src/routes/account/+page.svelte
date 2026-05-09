<script lang="ts">
    import { onMount } from 'svelte';
    import { browser } from '$app/environment';
    import Account from '../accountStore';
    import Page from '$lib/ui/Page.svelte';
    import Stack from '$lib/ui/Stack.svelte';
    import Cluster from '$lib/ui/Cluster.svelte';
    import Paper from '$lib/ui/Paper.svelte';
    import Button from '$lib/ui/Button.svelte';
    import TabBar from '$lib/ui/TabBar.svelte';
    import GameModal from '$lib/feature/GameModal.svelte';
    import StatsBoard from '$lib/feature/StatsBoard.svelte';
    import ContinuePlaying from '$lib/feature/ContinuePlaying.svelte';
    import AccountFriendRow from '$lib/feature/AccountFriendRow.svelte';
    import AccountPackRow from '$lib/feature/AccountPackRow.svelte';
    import AccountGameRow from '$lib/feature/AccountGameRow.svelte';
    import { _LikePack, _DeleteGame, _GetGame } from './+page';
    import { DeleteFriend, AcceptFriend } from '../friendRequests';
    import { _StartSolo } from '../packs/+page';

    type TabKey = 'friends' | 'packs' | 'games';

    let tab: TabKey = 'friends';
    let showModal = false;
    let selectedPackID = '';
    let selectedFriendID = '';
    let showPastGames = false;

    function readHashTab(): TabKey {
        if (!browser) return 'friends';
        const h = window.location.hash.slice(1);
        if (h === 'packs' || h === 'games' || h === 'friends') return h;
        return 'friends';
    }

    function writeHashTab(next: TabKey) {
        if (!browser) return;
        history.replaceState(null, '', `#${next}`);
    }

    onMount(() => {
        tab = readHashTab();
    });

    $: if (browser) writeHashTab(tab);

    function getOpponentUsername(userId: string) {
        for (const f of $Account?.friends ?? []) if (f.userID === userId) return f.username;
        return '—';
    }
    function getPackTitle(packID: string) {
        const liked = $Account?.likedPacks.find((p) => p.id === packID);
        if (liked) return liked.pack.title;
        const owned = $Account?.packs?.find((p) => p.id === packID);
        return owned?.pack.title ?? '—';
    }

    $: friendsAll = (() => {
        const all = $Account?.friends ?? [];
        const byOrder = (s: number) => (s === 2 ? 0 : s === 1 ? 1 : 2);
        return [...all].sort((a, b) => byOrder(a.status) - byOrder(b.status));
    })();
    $: packsAll = $Account?.likedPacks ?? [];
    $: activeGames = ($Account?.games ?? []).filter((g) => g.status !== 3);
    $: pastGames = ($Account?.games ?? []).filter((g) => g.status === 3);
    $: featuredGame = activeGames[0];

    function openWithPack(packID: string) { selectedPackID = packID; selectedFriendID = ''; showModal = true; }
    function openWithFriend(friendID: string) { selectedFriendID = friendID; selectedPackID = ''; showModal = true; }

    function gameLabel(game: { kind: string; user1Id: string; user2Id: string; packId: string }): string {
        const pack = getPackTitle(game.packId);
        if (game.kind === 'solo') return `Solo · ${pack}`;
        const me = $Account?.userID;
        const opp = me === game.user1Id ? game.user2Id : game.user1Id;
        return `vs. ${getOpponentUsername(opp)} · ${pack}`;
    }
</script>

<svelte:head><title>{$Account?.username ?? 'Account'} · taskbingo</title></svelte:head>

<Page width="normal">
    <Stack gap="xl">
        {#if $Account}
            <header class="greet">
                <h1 class="name">{$Account.username}</h1>
                <Cluster gap="m" align="baseline">
                    <span class="score">{$Account.bingo}</span>
                    <span class="score-label">total bingos</span>
                    {#if $Account.soloBingo}
                        <span class="solo">solo: {$Account.soloBingo}</span>
                    {/if}
                </Cluster>
                <p class="city">{$Account.city}</p>
            </header>

            <StatsBoard />

            {#if featuredGame}
                <ContinuePlaying
                    game={featuredGame}
                    opponentName={featuredGame.kind === 'solo'
                        ? ''
                        : getOpponentUsername($Account.userID === featuredGame.user1Id ? featuredGame.user2Id : featuredGame.user1Id)}
                    packTitle={getPackTitle(featuredGame.packId)}
                    on:resume={(e) => _GetGame(e.detail)}
                />
            {/if}

            <section class="tabs-section">
                <Cluster gap="m" align="center" justify="between">
                    <TabBar
                        bind:value={tab}
                        items={[
                            { value: 'friends', label: 'Friends' },
                            { value: 'packs', label: 'Packs' },
                            { value: 'games', label: 'Games' },
                        ]}
                        ariaLabel="Account sections"
                    />
                    {#if tab === 'packs'}
                        <Button variant="ghost" size="sm" href="/newpack">New pack</Button>
                    {:else if tab === 'games'}
                        <Button variant="ghost" size="sm" on:click={() => (showModal = true)}>New game</Button>
                    {/if}
                </Cluster>

                <Paper padding="md">
                    {#if tab === 'friends'}
                        {#if friendsAll.length === 0}
                            <p class="empty">No friends yet — <a href="/people">find people →</a></p>
                        {:else}
                            {#each friendsAll as friend (friend.userID)}
                                <AccountFriendRow
                                    {friend}
                                    on:play={(e) => openWithFriend(e.detail)}
                                    on:accept={(e) => AcceptFriend(e.detail)}
                                    on:remove={(e) => DeleteFriend(e.detail)}
                                />
                            {/each}
                        {/if}
                    {:else if tab === 'packs'}
                        {#if packsAll.length === 0}
                            <p class="empty">No liked packs yet — <a href="/packs">browse public packs →</a></p>
                        {:else}
                            {#each packsAll as pack (pack.id)}
                                <AccountPackRow
                                    {pack}
                                    on:play={(e) => openWithPack(e.detail)}
                                    on:solo={(e) => _StartSolo(e.detail)}
                                    on:unlike={(e) => _LikePack({ id: e.detail }, true)}
                                />
                            {/each}
                        {/if}
                    {:else}
                        {#if activeGames.length === 0 && pastGames.length === 0}
                            <p class="empty">Your first board awaits.</p>
                        {:else}
                            {#each activeGames as game (game.gameId)}
                                <AccountGameRow
                                    {game}
                                    label={gameLabel(game)}
                                    on:resume={(e) => _GetGame(e.detail)}
                                    on:delete={(e) => _DeleteGame(e.detail)}
                                />
                            {/each}
                            {#if pastGames.length > 0}
                                <button class="see-all" on:click={() => (showPastGames = !showPastGames)}>
                                    {showPastGames ? 'Hide past games' : `Show ${pastGames.length} past games`}
                                </button>
                                {#if showPastGames}
                                    {#each pastGames as game (game.gameId)}
                                        <AccountGameRow
                                            {game}
                                            label={gameLabel(game)}
                                            past
                                            on:delete={(e) => _DeleteGame(e.detail)}
                                        />
                                    {/each}
                                {/if}
                            {/if}
                        {/if}
                    {/if}
                </Paper>
            </section>
        {/if}
    </Stack>

    <GameModal bind:showModal {selectedFriendID} {selectedPackID} />
</Page>

<style>
    .greet { display: flex; flex-direction: column; gap: var(--space-2); }
    .name {
        font-family: var(--font-display);
        font-size: 2rem;
        font-weight: 400;
        letter-spacing: -0.02em;
    }
    .score {
        font-family: var(--font-display);
        font-size: 2rem;
        font-weight: 400;
        color: var(--ink);
        line-height: 1;
    }
    .score-label {
        font-size: 0.65rem;
        text-transform: uppercase;
        letter-spacing: 0.16em;
        color: var(--ink-3);
    }
    .solo {
        font-family: var(--font-display);
        font-style: italic;
        font-size: 0.95rem;
        color: var(--ink-3);
    }
    .city { color: var(--ink-3); font-size: 0.85rem; }

    .tabs-section {
        display: flex;
        flex-direction: column;
        gap: var(--space-3);
    }

    .empty {
        font-family: var(--font-display);
        font-style: italic;
        color: var(--ink-3);
        text-align: center;
        margin: var(--space-3) 0;
        font-size: 0.95rem;
    }
    .empty a {
        color: var(--ink);
        border-bottom: 1px solid rgba(40, 55, 95, 0.25);
        text-decoration: none;
    }
    .empty a:hover {
        color: var(--accent-from);
        border-bottom-color: var(--accent-from);
    }

    .see-all {
        background: transparent;
        border: none;
        color: var(--ink-2);
        font-size: 0.78rem;
        cursor: pointer;
        padding: var(--space-3) 0;
        align-self: flex-start;
        text-decoration: underline;
        text-underline-offset: 4px;
        font-family: var(--font-body);
    }
    .see-all:hover { color: var(--ink); }
</style>

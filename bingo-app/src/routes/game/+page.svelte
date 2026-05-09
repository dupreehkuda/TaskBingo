<script lang="ts">
    import { onDestroy, onMount } from 'svelte';
    import { browser } from '$app/environment';
    import { get, type Unsubscriber } from 'svelte/store';
    import Account from '../accountStore';
    import CurrentGame from '../currentGame';
    import Page from '$lib/ui/Page.svelte';
    import Stack from '$lib/ui/Stack.svelte';
    import Cluster from '$lib/ui/Cluster.svelte';
    import Glass from '$lib/ui/Glass.svelte';
    import TabBar from '$lib/ui/TabBar.svelte';
    import Button from '$lib/ui/Button.svelte';
    import BingoGrid from '$lib/feature/BingoGrid.svelte';
    import Keypad from '$lib/feature/Keypad.svelte';
    import TaskList from '$lib/feature/TaskList.svelte';
    import NotesJournal from '$lib/feature/NotesJournal.svelte';
    import {
        _GameHandler,
        _SendUpdate,
        _RedirectOnAccount,
        _PlaceSoloNumber,
        _SoloFinish,
    } from './+page';

    let socket: WebSocket | undefined;
    let closer: () => void = () => {};

    let account = get(Account);
    let game = get(CurrentGame);
    let isSolo = game?.kind === 'solo';
    let finished = false;
    let activeTab: 'mark' | 'tasks' | 'notes' = 'mark';

    onMount(() => {
        if (!game) {
            _RedirectOnAccount();
            return;
        }

        if (!isSolo) {
            const handler = _GameHandler();
            socket = handler.socket;
            closer = handler.closer;
        }

        if (browser) {
            window.addEventListener('beforeunload', () => closer());
        }
    });

    onDestroy(() => {
        closer();
    });

    function handleSubmit(e: CustomEvent<number>) {
        const n = e.detail;
        if (n <= 0 || n > 16) return;
        if (isSolo) _PlaceSoloNumber(n);
        else if (socket) _SendUpdate(socket, n, false);
    }

    async function handleFinish() {
        finished = !finished;
        if (isSolo) {
            if (finished) {
                await _SoloFinish();
                _RedirectOnAccount();
            }
            return;
        }
        if (socket) _SendUpdate(socket, 0, finished);
    }

    $: marked = $CurrentGame
        ? ($CurrentGame.user1ID === account.userID ? $CurrentGame.user1Numbers : $CurrentGame.user2Numbers)
        : [];
    $: oppMarked = $CurrentGame
        ? ($CurrentGame.user1ID === account.userID ? $CurrentGame.user2Numbers : $CurrentGame.user1Numbers)
        : [];

    // The 10 possible 4-in-a-row lines on a 4×4 board.
    const ALL_LINES: number[][] = [
        [0, 1, 2, 3], [4, 5, 6, 7], [8, 9, 10, 11], [12, 13, 14, 15],
        [0, 4, 8, 12], [1, 5, 9, 13], [2, 6, 10, 14], [3, 7, 11, 15],
        [0, 5, 10, 15], [3, 6, 9, 12],
    ];

    function computeWinning(numbers: number[], markedArr: number[]): Set<number> {
        const set = new Set(markedArr.filter((n) => n !== 0));
        const winning = new Set<number>();
        for (const line of ALL_LINES) {
            if (line.every((idx) => set.has(numbers[idx]))) {
                line.forEach((idx) => winning.add(idx));
            }
        }
        return winning;
    }

    function countBingoLines(numbers: number[], markedArr: number[]): number {
        const set = new Set(markedArr.filter((n) => n !== 0));
        let n = 0;
        for (const line of ALL_LINES) {
            if (line.every((idx) => set.has(numbers[idx]))) n++;
        }
        return n;
    }

    $: winningCells = $CurrentGame ? computeWinning($CurrentGame.numbers, marked) : new Set<number>();
    $: usersBingo = $CurrentGame ? countBingoLines($CurrentGame.numbers, marked) : 0;
    $: opponentsBingo = $CurrentGame ? countBingoLines($CurrentGame.numbers, oppMarked) : 0;
</script>

<svelte:head><title>Game · taskbingo</title></svelte:head>

<div class="game-shell">
    {#if !isSolo && $CurrentGame?.status < 3}
        <Page width="narrow">
            <Stack gap="l" align="center">
                <h2 class="waiting">Waiting for the opponent…</h2>
                <span class="waiting-dot"></span>
            </Stack>
        </Page>
    {:else if $CurrentGame}
        <!-- Sticky glass header -->
        <div class="game-header">
            <Glass radius="card" blur="strong">
                <div class="score-row">
                    <div class="scores">
                        <span class="who">You</span>
                        <span class="score">{usersBingo}</span>
                        {#if !isSolo}
                            <span class="sep">·</span>
                            <span class="who">Opp</span>
                            <span class="score">{opponentsBingo}</span>
                        {/if}
                    </div>
                    <Button variant="accent" size="sm" on:click={handleFinish}>
                        {finished ? 'Resume' : 'Finish'}
                    </Button>
                </div>
            </Glass>
        </div>

        <!-- Body -->
        <Page width="normal">
            <div class="game-body" class:disabled={finished}>
                <div class="grid-col">
                    <BingoGrid numbers={game.numbers} marked={marked} {winningCells} />
                </div>
                <div class="side-col">
                    <Stack gap="m">
                        <TabBar
                            bind:value={activeTab}
                            items={[
                                { value: 'mark', label: 'Mark' },
                                { value: 'tasks', label: 'Tasks' },
                                { value: 'notes', label: 'Notes' },
                            ]}
                            ariaLabel="Game panel"
                        />
                        {#if activeTab === 'mark'}
                            <Keypad on:submit={handleSubmit} />
                        {:else if activeTab === 'tasks'}
                            <TaskList
                                packID={game.packID}
                                usersTasks={marked}
                                opponentsTasks={oppMarked}
                                showOpponent={!isSolo}
                            />
                        {:else}
                            <NotesJournal gameID={game.gameID} />
                        {/if}
                    </Stack>
                </div>
            </div>
        </Page>
    {/if}
</div>

<style>
    .game-shell { width: 100%; }

    .waiting {
        font-family: var(--font-display);
        font-style: italic;
        font-weight: 400;
        font-size: 1.4rem;
        color: var(--ink-2);
    }
    .waiting-dot {
        width: 8px; height: 8px; border-radius: var(--radius-pill);
        background: var(--accent);
        animation: pulse 1.4s var(--ease) infinite;
    }
    @keyframes pulse {
        0%, 100% { opacity: 0.4; transform: scale(1); }
        50% { opacity: 1; transform: scale(1.4); }
    }

    .game-header {
        position: sticky;
        top: 0;
        z-index: 5;
        padding: var(--space-3) var(--space-5);
        max-width: 60rem;
        margin: 0 auto;
    }

    .score-row {
        display: flex;
        align-items: center;
        justify-content: space-between;
        gap: var(--space-4);
        padding: var(--space-3) var(--space-4);
    }

    .scores {
        display: flex;
        align-items: baseline;
        gap: var(--space-3);
    }

    .game-body {
        display: flex;
        flex-wrap: wrap;
        gap: var(--space-5);
        justify-content: center;
        align-items: flex-start;
    }
    .game-body.disabled {
        pointer-events: none;
        opacity: 0.55;
    }

    .grid-col {
        flex: 0 0 26rem;
        max-width: 26rem;
        width: 100%;
    }

    .side-col {
        flex: 1 1 20rem;
        max-width: 22rem;
        width: 100%;
        position: sticky;
        top: calc(var(--space-3) + 76px);
    }

    @media (max-width: 1023px) {
        .grid-col { flex-basis: 22rem; max-width: 22rem; }
        .side-col { position: static; flex-basis: 22rem; max-width: 22rem; }
    }

    .who {
        font-size: 0.65rem;
        letter-spacing: 0.16em;
        text-transform: uppercase;
        color: var(--ink-3);
    }
    .score {
        font-family: var(--font-display);
        font-size: 1.5rem;
        line-height: 1;
        color: var(--ink);
        font-variant-numeric: tabular-nums;
    }
    .sep {
        color: var(--ink-4);
        margin: 0 var(--space-2);
    }
</style>

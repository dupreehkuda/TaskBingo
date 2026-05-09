<script lang="ts">
    import { onDestroy, onMount } from 'svelte';
    import { browser } from '$app/environment';
    import BingoCell from './BingoCell.svelte';

    /**
     * HeroDemo — looped animation of a 4×4 board filling in. Each cycle:
     * 1) shuffles numbers 1–16,
     * 2) picks a random "target" line and a random set of extra cells,
     *    then plays the marks in order so the target completes at the end
     *    (additional lines may form by chance);
     * 3) detects every completed 4-in-a-row line from the final marked set;
     * 4) the shared BingoCell turns those cells sage green;
     * 5) fades out and restarts.
     *
     * Stays still under prefers-reduced-motion (renders fully marked).
     */

    const ALL_LINES: number[][] = [
        [0, 1, 2, 3], [4, 5, 6, 7], [8, 9, 10, 11], [12, 13, 14, 15], // rows
        [0, 4, 8, 12], [1, 5, 9, 13], [2, 6, 10, 14], [3, 7, 11, 15], // cols
        [0, 5, 10, 15], [3, 6, 9, 12], // diagonals
    ];

    function shuffle<T>(arr: readonly T[]): T[] {
        const a = arr.slice();
        for (let i = a.length - 1; i > 0; i--) {
            const j = Math.floor(Math.random() * (i + 1));
            [a[i], a[j]] = [a[j], a[i]];
        }
        return a;
    }

    function pickRandom<T>(arr: readonly T[]): T {
        return arr[Math.floor(Math.random() * arr.length)];
    }

    interface Round {
        numbers: number[];
        sequence: number[];
    }

    function generateRound(): Round {
        const numbers = shuffle([1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16]);
        const targetLine = pickRandom(ALL_LINES);
        const targetSet = new Set(targetLine);

        // 4–8 extras alongside the 4 target cells → 8–12 total marks
        const extraCount = 4 + Math.floor(Math.random() * 5);
        const otherIdx = [...Array(16).keys()].filter((i) => !targetSet.has(i));
        const extras = shuffle(otherIdx).slice(0, extraCount);

        const sequence = [
            ...shuffle(extras).map((i) => numbers[i]),
            ...shuffle(targetLine).map((i) => numbers[i]),
        ];

        return { numbers, sequence };
    }

    function computeWinning(numbers: number[], marked: number[]): Set<number> {
        const markedSet = new Set(marked);
        const winning = new Set<number>();
        for (const line of ALL_LINES) {
            if (line.every((idx) => markedSet.has(numbers[idx]))) {
                line.forEach((idx) => winning.add(idx));
            }
        }
        return winning;
    }

    let round: Round = generateRound();
    let marked: number[] = [];
    let winning: Set<number> = new Set();
    let exiting = false;

    let stepTimer: ReturnType<typeof setInterval> | undefined;
    let resetTimer: ReturnType<typeof setTimeout> | undefined;
    let restartTimer: ReturnType<typeof setTimeout> | undefined;

    function tick() {
        if (marked.length < round.sequence.length) {
            marked = [...marked, round.sequence[marked.length]];
            if (marked.length === round.sequence.length) {
                winning = computeWinning(round.numbers, marked);
                if (stepTimer) clearInterval(stepTimer);
                stepTimer = undefined;
                resetTimer = setTimeout(() => {
                    exiting = true;
                    restartTimer = setTimeout(() => {
                        round = generateRound();
                        marked = [];
                        winning = new Set();
                        exiting = false;
                        stepTimer = setInterval(tick, 520);
                    }, 520);
                }, 1700);
            }
        }
    }

    onMount(() => {
        if (!browser) return;
        const reduced = window.matchMedia('(prefers-reduced-motion: reduce)').matches;
        if (reduced) {
            marked = round.sequence;
            winning = computeWinning(round.numbers, marked);
            return;
        }
        restartTimer = setTimeout(() => {
            stepTimer = setInterval(tick, 520);
        }, 500);
    });

    onDestroy(() => {
        if (stepTimer) clearInterval(stepTimer);
        if (resetTimer) clearTimeout(resetTimer);
        if (restartTimer) clearTimeout(restartTimer);
    });
</script>

<div class="demo" class:exiting aria-hidden="true">
    <div class="grid">
        {#each round.numbers as n, idx (idx)}
            <BingoCell {n} on={marked.includes(n)} win={winning.has(idx)} />
        {/each}
    </div>
</div>

<style>
    .demo {
        width: 100%;
        max-width: 22rem;
        transition: opacity 480ms var(--ease);
    }
    .demo.exiting { opacity: 0; }

    .grid {
        display: grid;
        grid-template-columns: repeat(4, 1fr);
        gap: 8px;
        width: 100%;
    }

    @media (prefers-reduced-motion: reduce) {
        .demo { transition: none; }
    }
</style>

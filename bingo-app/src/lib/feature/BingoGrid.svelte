<script lang="ts">
    import BingoCell from './BingoCell.svelte';

    /**
     * BingoGrid — 4×4 board. `numbers` is the underlying random ordering
     * shown in cells; `marked` is the array of values the player has marked
     * (matches the existing user1Numbers/user2Numbers shape: 16-length, 0
     * means empty, otherwise the original number). Cell rendering lives
     * in the shared BingoCell so home-page demo and live game look identical.
     */
    export let numbers: number[] = [];
    export let marked: number[] = [];
    export let winningCells: Set<number> = new Set();
</script>

<div class="grid" role="grid" aria-label="Bingo board">
    {#each numbers as n, idx (n + '-' + idx)}
        <BingoCell {n} on={marked.includes(n)} win={winningCells.has(idx)} />
    {/each}
</div>

<style>
    .grid {
        display: grid;
        grid-template-columns: repeat(4, 1fr);
        gap: 8px;
        width: 100%;
    }
</style>

<script lang="ts">
    /**
     * StatTile — small paper-soft tile: big Fraunces number + uppercase label.
     * Used in the summary row at the top of the StatsBoard.
     */
    export let label: string;
    export let value: number | string;
    export let hint: string | undefined = undefined;
</script>

<div class="tile">
    <div class="value">{value}</div>
    <div class="label">{label}</div>
    {#if hint}<div class="hint">{hint}</div>{/if}
</div>

<style>
    .tile {
        background: var(--paper-soft);
        border: 1px solid var(--hairline);
        border-radius: var(--radius-md);
        padding: var(--space-3) var(--space-4);
        display: flex;
        flex-direction: column;
        gap: 0.15rem;
        min-width: 0;
    }

    .value {
        font-family: var(--font-display);
        font-weight: 400;
        font-size: 1.6rem;
        line-height: 1;
        color: var(--ink);
        font-variant-numeric: tabular-nums;
        letter-spacing: -0.015em;
    }

    .label {
        font-size: 0.6rem;
        letter-spacing: 0.18em;
        text-transform: uppercase;
        color: var(--ink-3);
        margin-top: 0.35rem;
    }

    .hint {
        font-size: 0.7rem;
        color: var(--ink-3);
        font-style: italic;
        font-family: var(--font-display);
    }
</style>

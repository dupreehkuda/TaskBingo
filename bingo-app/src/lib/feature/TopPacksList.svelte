<script lang="ts">
    import type { StatsPackCount } from '../../routes/statsStore';

    export let title: string;
    export let items: StatsPackCount[] = [];

    $: max = Math.max(1, ...items.map((i) => i.count));
</script>

<section class="ranks">
    <h3 class="title">{title}</h3>
    {#if items.length === 0}
        <p class="empty">No packs played yet in this window.</p>
    {:else}
        <ol class="list">
            {#each items as pack, i (pack.id)}
                <li>
                    <div class="row">
                        <span class="rank">{i + 1}</span>
                        <span class="name" title={pack.title}>{pack.title}</span>
                        <span class="count">×{pack.count}</span>
                    </div>
                    <div class="track" aria-hidden="true">
                        <span
                            class="fill"
                            style:width="{Math.round((pack.count / max) * 100)}%"
                        ></span>
                    </div>
                </li>
            {/each}
        </ol>
    {/if}
</section>

<style>
    .ranks {
        display: flex;
        flex-direction: column;
        gap: var(--space-3);
    }

    .title {
        font-family: var(--font-display);
        font-size: 1rem;
        font-weight: 400;
        letter-spacing: -0.01em;
        color: var(--ink);
        margin: 0;
    }

    .empty {
        color: var(--ink-3);
        font-style: italic;
        font-size: 0.85rem;
        margin: 0;
    }

    .list {
        list-style: none;
        margin: 0;
        padding: 0;
        display: flex;
        flex-direction: column;
        gap: var(--space-3);
    }

    .row {
        display: grid;
        grid-template-columns: auto 1fr auto;
        gap: var(--space-3);
        align-items: baseline;
        font-size: 0.85rem;
        color: var(--ink);
    }

    .rank {
        color: var(--ink-4);
        font-variant-numeric: tabular-nums;
        font-size: 0.78rem;
        min-width: 1em;
    }

    .name {
        font-family: var(--font-display);
        font-weight: 400;
        color: var(--ink);
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
    }

    .count {
        color: var(--ink-3);
        font-variant-numeric: tabular-nums;
        font-size: 0.78rem;
    }

    .track {
        margin-top: 0.35rem;
        height: 4px;
        background: var(--paper-soft);
        border-radius: var(--radius-pill);
        overflow: hidden;
    }

    .fill {
        display: block;
        height: 100%;
        background: var(--accent);
        border-radius: var(--radius-pill);
        animation: grow 380ms var(--ease) both;
        transform-origin: left;
    }

    @keyframes grow {
        from { transform: scaleX(0); }
        to   { transform: scaleX(1); }
    }

    @media (prefers-reduced-motion: reduce) {
        .fill { animation: none; }
    }
</style>

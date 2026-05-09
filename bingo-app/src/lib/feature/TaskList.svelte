<script lang="ts">
    import { GetTaskPack } from '$lib/data/packs';

    export let packID = '';
    export let usersTasks: number[] = [];
    export let opponentsTasks: number[] = [];
    export let showOpponent = true;

    $: pack = GetTaskPack(packID);
</script>

<div class="list">
    <h3 class="header">{pack.pack.title}</h3>
    <ol>
        {#each pack.pack.tasks as task, i}
            {@const num = i + 1}
            {@const mineDone = usersTasks.includes(num)}
            {@const oppDone = opponentsTasks.includes(num)}
            <li class:done={mineDone}>
                <span class="num">{num}</span>
                <span class="body">{task}</span>
                <span class="marks" aria-hidden="true">
                    {#if showOpponent}
                        <span class="dot" class:filled={oppDone} title="Opponent"></span>
                    {/if}
                    <span class="dot" class:filled={mineDone} title="You"></span>
                </span>
            </li>
        {/each}
    </ol>
</div>

<style>
    .list { display: flex; flex-direction: column; gap: var(--space-2); }
    .header {
        font-family: var(--font-display);
        font-size: 0.95rem;
        font-weight: 400;
        letter-spacing: -0.01em;
        color: var(--ink-2);
        margin: 0;
    }
    ol { list-style: none; margin: 0; padding: 0; display: flex; flex-direction: column; }
    li {
        display: grid;
        grid-template-columns: 1.4em 1fr auto;
        gap: var(--space-2);
        align-items: center;
        padding: 0.18rem 0;
        font-size: 0.82rem;
        color: var(--ink);
        line-height: 1.35;
    }
    li.done .body { text-decoration: line-through; color: var(--ink-3); }
    .num {
        color: var(--ink-4);
        font-variant-numeric: tabular-nums;
        font-size: 0.7rem;
        text-align: right;
    }
    .body {
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
    }
    .marks { display: inline-flex; gap: 3px; }
    .dot {
        width: 6px; height: 6px;
        border-radius: var(--radius-pill);
        background: var(--paper-soft);
        border: 1px solid var(--hairline);
    }
    .dot.filled { background: var(--accent); border-color: transparent; }
</style>

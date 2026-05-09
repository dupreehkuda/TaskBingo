<script lang="ts">
    import { createEventDispatcher } from 'svelte';
    import Cluster from '$lib/ui/Cluster.svelte';
    import Button from '$lib/ui/Button.svelte';
    import type { TaskPack } from '../../routes/accountStore';

    /**
     * AccountPackRow — one row inside the Packs tab list. Click anywhere on
     * the row that isn't a button toggles full-list expansion (all 16 tasks).
     * Buttons stop propagation so clicks on Play/Solo/heart never trigger
     * expand/collapse.
     */
    export let pack: TaskPack;

    const dispatch = createEventDispatcher<{
        play: string;
        solo: string;
        unlike: string;
    }>();

    let expanded = false;

    function toggle() {
        expanded = !expanded;
    }

    function key(e: KeyboardEvent) {
        if (e.key === 'Enter' || e.key === ' ') {
            e.preventDefault();
            toggle();
        }
    }
</script>

<div
    class="row"
    class:expanded
    role="button"
    tabindex="0"
    aria-expanded={expanded}
    on:click={toggle}
    on:keydown={key}
>
    <div class="head">
        <div class="title-block">
            <span class="title">{pack.pack.title}</span>
            {#if pack.isPrivate}<span class="badge">private</span>{/if}
        </div>
        <span class="chevron" class:rot={expanded} aria-hidden="true">
            <svg viewBox="0 0 20 20" width="14" height="14"><path d="M5 8l5 5 5-5" fill="none" stroke="currentColor" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round"/></svg>
        </span>
    </div>

    {#if expanded}
        <ol class="tasks-full">
            {#each pack.pack.tasks as task, i}
                <li><span class="num">{i + 1}</span><span>{task}</span></li>
            {/each}
        </ol>
    {:else}
        <ul class="tasks-preview">
            {#each pack.pack.tasks.slice(0, 3) as task, i}
                <li><span class="num">{i + 1}</span><span>{task}</span></li>
            {/each}
            {#if pack.pack.tasks.length > 3}
                <li class="more">+ {pack.pack.tasks.length - 3} more</li>
            {/if}
        </ul>
    {/if}

    <div class="actions" on:click|stopPropagation on:keydown|stopPropagation role="presentation">
        <Cluster gap="s">
            <Button variant="accent" size="sm" on:click={() => dispatch('play', pack.id)}>Play</Button>
            <Button variant="ghost" size="sm" on:click={() => dispatch('solo', pack.id)}>Solo</Button>
            <Button variant="ghost" size="sm" on:click={() => dispatch('unlike', pack.id)} aria-label="Unlike pack">
                <span aria-hidden="true">♥</span>
            </Button>
        </Cluster>
    </div>
</div>

<style>
    .row {
        display: flex;
        flex-direction: column;
        gap: var(--space-2);
        padding: var(--space-3) 0;
        border-bottom: 1px solid var(--hairline);
        cursor: pointer;
        outline: none;
    }
    .row:last-child { border-bottom: none; }
    .row:focus-visible {
        outline: 2px solid var(--accent-ring);
        outline-offset: 2px;
        border-radius: var(--radius-sm);
    }

    .head {
        display: flex;
        align-items: center;
        justify-content: space-between;
        gap: var(--space-3);
    }

    .title-block {
        display: flex;
        align-items: baseline;
        gap: var(--space-2);
        min-width: 0;
    }

    .title {
        font-family: var(--font-display);
        font-size: 1.05rem;
        color: var(--ink);
        font-weight: 500;
        letter-spacing: -0.01em;
    }

    .badge {
        font-family: var(--font-body);
        font-size: 0.6rem;
        padding: 0.1rem 0.5rem;
        border-radius: var(--radius-pill);
        background: var(--paper-soft);
        color: var(--ink-2);
        border: 1px solid var(--hairline);
    }

    .chevron {
        color: var(--ink-3);
        transition: transform var(--dur-base) var(--ease);
        line-height: 0;
    }
    .chevron.rot { transform: rotate(180deg); }

    .tasks-preview, .tasks-full {
        list-style: none;
        margin: 0;
        padding: 0;
        display: grid;
        gap: var(--space-1) var(--space-4);
        font-size: 0.85rem;
        color: var(--ink-2);
        line-height: 1.5;
    }
    .tasks-preview { grid-template-columns: 1fr; }
    .tasks-full {
        grid-template-columns: 1fr 1fr;
    }
    @media (max-width: 640px) {
        .tasks-full { grid-template-columns: 1fr; }
    }

    .tasks-preview li, .tasks-full li {
        display: flex;
        gap: var(--space-2);
    }
    .num {
        color: var(--ink-4);
        min-width: 1.5em;
        font-variant-numeric: tabular-nums;
    }
    .more {
        color: var(--ink-3);
        font-style: italic;
        padding-left: calc(1.5em + var(--space-2));
    }

    .actions {
        display: flex;
        justify-content: flex-end;
        margin-top: var(--space-1);
    }
</style>

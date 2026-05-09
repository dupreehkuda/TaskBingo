<script lang="ts">
    import { createEventDispatcher } from 'svelte';
    import Cluster from '$lib/ui/Cluster.svelte';
    import Button from '$lib/ui/Button.svelte';
    import type { Game } from '../../routes/accountStore';

    /**
     * AccountGameRow — one row inside the Games tab list. Used for both
     * active games (resume) and finished games (delete only). The page
     * controls which actions render via the `past` prop.
     */
    export let game: Game;
    export let label: string;
    export let past = false;

    const dispatch = createEventDispatcher<{
        resume: string;
        delete: string;
    }>();
</script>

<div class="row" class:past>
    <div class="meta">
        <span class="title">{label}</span>
        {#if past}
            <span class="sub">finished</span>
        {/if}
    </div>

    <Cluster gap="s">
        {#if !past}
            <Button
                variant="accent"
                size="sm"
                href={game.kind === 'solo' ? '/game?solo=true' : '/game'}
                on:mouseenter={() => dispatch('resume', game.gameId)}
                on:focus={() => dispatch('resume', game.gameId)}
            >
                Resume
            </Button>
        {/if}
        <Button variant="ghost" size="sm" on:click={() => dispatch('delete', game.gameId)}>Delete</Button>
    </Cluster>
</div>

<style>
    .row {
        display: flex;
        align-items: center;
        justify-content: space-between;
        gap: var(--space-3);
        padding: var(--space-3) 0;
        border-bottom: 1px solid var(--hairline);
    }
    .row:last-child { border-bottom: none; }
    .row.past .title { color: var(--ink-3); }

    .meta {
        display: flex;
        flex-direction: column;
        gap: 0.1rem;
        min-width: 0;
    }

    .title {
        font-family: var(--font-display);
        font-weight: 500;
        font-size: 0.95rem;
        color: var(--ink);
    }

    .sub {
        font-size: 0.7rem;
        color: var(--ink-3);
        text-transform: uppercase;
        letter-spacing: 0.16em;
    }
</style>

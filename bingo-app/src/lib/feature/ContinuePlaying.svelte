<script lang="ts">
    import { createEventDispatcher } from 'svelte';
    import Glass from '$lib/ui/Glass.svelte';
    import Button from '$lib/ui/Button.svelte';
    import type { Game } from '../../routes/accountStore';

    /**
     * ContinuePlaying — single glass strip for the user's most recent active
     * game. Renders nothing if no game is supplied; the page is expected to
     * gate render with {#if featured}.
     */
    export let game: Game;
    export let opponentName: string = '';
    export let packTitle: string = '';

    const dispatch = createEventDispatcher<{ resume: string }>();

    function resume() {
        dispatch('resume', game.gameId);
    }
</script>

<Glass radius="card" blur="strong">
    <div class="row">
        <div class="meta">
            <span class="eyebrow">continue playing</span>
            <span class="title">
                {#if game.kind === 'solo'}
                    Solo game
                {:else}
                    vs. {opponentName || '—'}
                {/if}
            </span>
            <span class="sub">{packTitle || '—'}</span>
        </div>
        <Button
            variant="accent"
            size="md"
            href={game.kind === 'solo' ? '/game?solo=true' : '/game'}
            on:mouseenter={resume}
            on:focus={resume}
        >
            Resume
        </Button>
    </div>
</Glass>

<style>
    .row {
        display: flex;
        align-items: center;
        justify-content: space-between;
        gap: var(--space-4);
        padding: var(--space-3) var(--space-4);
    }

    .meta {
        display: flex;
        flex-direction: column;
        gap: 0.15rem;
        min-width: 0;
    }

    .eyebrow {
        font-size: 0.6rem;
        letter-spacing: 0.18em;
        text-transform: uppercase;
        color: var(--ink-3);
    }

    .title {
        font-family: var(--font-display);
        font-size: 1.05rem;
        color: var(--ink);
        letter-spacing: -0.01em;
    }

    .sub {
        font-size: 0.8rem;
        color: var(--ink-3);
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
    }

    @media (max-width: 480px) {
        .row {
            flex-direction: column;
            align-items: stretch;
            gap: var(--space-3);
        }
    }
</style>

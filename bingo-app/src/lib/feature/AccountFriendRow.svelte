<script lang="ts">
    import { createEventDispatcher } from 'svelte';
    import Cluster from '$lib/ui/Cluster.svelte';
    import Button from '$lib/ui/Button.svelte';
    import type { Friend } from '../../routes/accountStore';

    /**
     * AccountFriendRow — one row inside the Friends tab list.
     * Status semantics (per backend):
     *   1 = sent (we requested)
     *   2 = incoming (they requested us)
     *   3 = accepted
     */
    export let friend: Friend;

    const dispatch = createEventDispatcher<{
        play: string;
        accept: string;
        remove: string;
    }>();

    $: incoming = friend.status === 2;
    $: sent = friend.status === 1;
    $: accepted = friend.status === 3;
</script>

<div class="row" class:incoming class:sent>
    <div class="left">
        {#if incoming}<span class="dot" aria-hidden="true"></span>{/if}
        <span class="avatar" aria-hidden="true"></span>
        <div class="who">
            <span class="name">{friend.username}</span>
            {#if accepted}
                <span class="meta">{friend.wins}/{friend.loses}</span>
            {:else if incoming}
                <span class="meta">wants to play</span>
            {:else if sent}
                <span class="meta">request sent</span>
            {/if}
        </div>
    </div>

    <Cluster gap="s" align="center" justify="end">
        {#if accepted}
            <Button variant="accent" size="sm" on:click={() => dispatch('play', friend.userID)}>Play</Button>
            <Button variant="ghost" size="sm" on:click={() => dispatch('remove', friend.userID)}>Remove</Button>
        {:else if incoming}
            <Button variant="accent" size="sm" on:click={() => dispatch('accept', friend.userID)}>Accept</Button>
            <Button variant="ghost" size="sm" on:click={() => dispatch('remove', friend.userID)}>Decline</Button>
        {:else if sent}
            <Button variant="ghost" size="sm" disabled>Sent</Button>
            <Button variant="ghost" size="sm" on:click={() => dispatch('remove', friend.userID)}>Cancel</Button>
        {/if}
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
    .row.sent { opacity: 0.7; }

    .left {
        display: flex;
        align-items: center;
        gap: var(--space-3);
        min-width: 0;
    }

    .dot {
        width: 6px; height: 6px;
        border-radius: var(--radius-pill);
        background: var(--accent);
        flex-shrink: 0;
    }

    .avatar {
        width: 28px; height: 28px;
        border-radius: var(--radius-pill);
        background: var(--accent);
        opacity: 0.85;
        flex-shrink: 0;
    }

    .who {
        display: flex;
        flex-direction: column;
        gap: 0.1rem;
        min-width: 0;
    }

    .name {
        font-family: var(--font-display);
        font-weight: 500;
        color: var(--ink);
        font-size: 0.95rem;
    }

    .meta {
        font-size: 0.75rem;
        color: var(--ink-3);
        font-variant-numeric: tabular-nums;
    }
</style>

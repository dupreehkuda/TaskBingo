<script lang="ts">
    import { onMount } from 'svelte';
    import Stack from '$lib/ui/Stack.svelte';
    import Cluster from '$lib/ui/Cluster.svelte';
    import Textarea from '$lib/ui/Textarea.svelte';
    import Button from '$lib/ui/Button.svelte';
    import {
        _ListComments,
        _AddComment,
        _EditComment,
        _DeleteComment,
        type Comment,
    } from '../../routes/game/comments';

    export let gameID: string;

    let comments: Comment[] = [];
    let draft = '';
    let editingID: string | null = null;
    let editingDraft = '';
    let composerEl: { focus: () => void } | undefined;

    onMount(async () => {
        comments = await _ListComments(gameID);
    });

    async function add() {
        if (!draft.trim()) return;
        const c = await _AddComment(gameID, draft.trim());
        if (c) {
            // newest at top
            comments = [c, ...comments];
            draft = '';
            // keep cursor in the composer so the next note can be typed straight away
            composerEl?.focus();
        }
    }
    function startEdit(c: Comment) { editingID = c.id; editingDraft = c.body; }
    async function saveEdit() {
        if (!editingID || !editingDraft.trim()) return;
        const updated = await _EditComment(editingID, editingDraft.trim());
        if (updated) comments = comments.map((c) => (c.id === updated.id ? updated : c));
        editingID = null;
    }
    async function remove(id: string) {
        if (await _DeleteComment(id)) comments = comments.filter((c) => c.id !== id);
    }
    function fmt(ts: string): string {
        try { return new Date(ts).toLocaleString(); } catch { return ts; }
    }

    // Always render newest first regardless of fetch order.
    $: ordered = [...comments].sort(
        (a, b) => new Date(b.createdAt).getTime() - new Date(a.createdAt).getTime(),
    );
</script>

<div class="journal">
    <h3>Notes</h3>

    <div class="composer">
        <Textarea bind:this={composerEl} label="Leave a note for yourself…" rows={2} bind:value={draft} />
        <Cluster gap="s" justify="end">
            <Button size="sm" variant="accent" on:click={add} disabled={!draft.trim()}>Add</Button>
        </Cluster>
    </div>

    {#if ordered.length === 0}
        <p class="empty">No notes yet.</p>
    {:else}
        <Stack gap="s">
            {#each ordered as c (c.id)}
                <div class="entry">
                    {#if editingID === c.id}
                        <Textarea label="Edit note" rows={2} bind:value={editingDraft} />
                        <Cluster gap="s">
                            <Button size="sm" variant="accent" on:click={saveEdit}>Save</Button>
                            <Button size="sm" variant="ghost" on:click={() => (editingID = null)}>Cancel</Button>
                        </Cluster>
                    {:else}
                        <p class="body">{c.body}</p>
                        <Cluster gap="s" align="baseline">
                            <span class="meta">{fmt(c.createdAt)}</span>
                            <button type="button" class="mini" on:click={() => startEdit(c)}>edit</button>
                            <button type="button" class="mini" on:click={() => remove(c.id)}>delete</button>
                        </Cluster>
                    {/if}
                </div>
            {/each}
        </Stack>
    {/if}
</div>

<style>
    .journal { display: flex; flex-direction: column; gap: var(--space-3); }
    h3 {
        font-family: var(--font-display);
        font-size: 1.1rem;
        font-weight: 400;
        letter-spacing: -0.01em;
        margin: 0;
    }
    .entry {
        padding: var(--space-3) 0;
        border-bottom: 1px solid var(--hairline);
        display: flex; flex-direction: column; gap: var(--space-2);
    }
    .entry:last-of-type { border-bottom: none; }
    .body { white-space: pre-wrap; font-size: 0.9rem; color: var(--ink); margin: 0; }
    .meta { font-size: 0.72rem; color: var(--ink-3); }
    .empty { color: var(--ink-3); font-style: italic; font-size: 0.85rem; }
    .composer {
        display: flex;
        flex-direction: column;
        gap: var(--space-2);
        padding-bottom: var(--space-3);
        border-bottom: 1px solid var(--hairline);
    }
    .mini {
        background: transparent;
        border: none;
        color: var(--ink-2);
        font-size: 0.72rem;
        cursor: pointer;
        text-decoration: underline;
        text-underline-offset: 3px;
        padding: 0;
    }
    .mini:hover { color: var(--ink); }
</style>

<script lang="ts">
    import { onMount } from 'svelte'
    import { Button, Textarea } from 'flowbite-svelte'
    import {
        _ListComments,
        _AddComment,
        _EditComment,
        _DeleteComment,
        type Comment,
    } from '../../routes/game/comments'

    export let gameID: string

    let comments: Comment[] = []
    let draft = ''
    let editingID: string | null = null
    let editingDraft = ''

    onMount(async () => {
        comments = await _ListComments(gameID)
    })

    async function add() {
        if (!draft.trim()) return
        const c = await _AddComment(gameID, draft.trim())
        if (c) {
            comments = [...comments, c]
            draft = ''
        }
    }

    function startEdit(c: Comment) {
        editingID = c.id
        editingDraft = c.body
    }

    async function saveEdit() {
        if (!editingID || !editingDraft.trim()) return
        const updated = await _EditComment(editingID, editingDraft.trim())
        if (updated) {
            comments = comments.map(c => (c.id === updated.id ? updated : c))
        }
        editingID = null
        editingDraft = ''
    }

    async function remove(id: string) {
        if (await _DeleteComment(id)) {
            comments = comments.filter(c => c.id !== id)
        }
    }

    function fmt(ts: string): string {
        try {
            return new Date(ts).toLocaleString()
        } catch {
            return ts
        }
    }
</script>

<section class="journal">
    <h4>Notes</h4>

    <ul class="entries">
        {#each comments as c (c.id)}
            <li class="entry">
                {#if editingID === c.id}
                    <Textarea bind:value={editingDraft} rows={2} />
                    <div class="row">
                        <Button size="xs" on:click={saveEdit}>Save</Button>
                        <Button size="xs" color="alternative" on:click={() => (editingID = null)}>Cancel</Button>
                    </div>
                {:else}
                    <div class="body">{c.body}</div>
                    <div class="meta">
                        <span>{fmt(c.createdAt)}</span>
                        <Button size="xs" color="alternative" on:click={() => startEdit(c)}>edit</Button>
                        <Button size="xs" color="red" on:click={() => remove(c.id)}>delete</Button>
                    </div>
                {/if}
            </li>
        {/each}
        {#if comments.length === 0}
            <li class="empty">No notes yet.</li>
        {/if}
    </ul>

    <div class="composer">
        <Textarea bind:value={draft} placeholder="Leave a note for yourself…" rows={2} />
        <Button size="sm" on:click={add}>Add</Button>
    </div>
</section>

<style>
    .journal {
        margin-top: 1em;
        padding: 0.6em;
        background-color: #ffffff;
        border-radius: 10px;
        color: #112a41;
    }
    h4 {
        font-weight: 500;
        margin-bottom: 0.4em;
    }
    .entries {
        list-style: none;
        padding: 0;
        margin: 0 0 0.6em 0;
    }
    .entry {
        padding: 0.4em 0;
        border-bottom: 1px solid #eee;
    }
    .entry:last-child {
        border-bottom: none;
    }
    .body {
        white-space: pre-wrap;
    }
    .meta,
    .row {
        display: flex;
        gap: 0.4em;
        align-items: center;
        margin-top: 0.3em;
        font-size: 0.8em;
        color: #5f6c7a;
    }
    .empty {
        color: #5f6c7a;
        font-style: italic;
        padding: 0.4em 0;
    }
    .composer {
        display: flex;
        flex-direction: column;
        gap: 0.4em;
    }
</style>

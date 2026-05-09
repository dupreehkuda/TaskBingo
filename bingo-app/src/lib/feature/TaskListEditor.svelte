<script lang="ts">
    import { createEventDispatcher } from 'svelte';
    import Cluster from '$lib/ui/Cluster.svelte';

    const MAX = 16;

    export let tasks: string[] = [];

    const dispatch = createEventDispatcher<{ change: string[] }>();

    let draft = '';
    let editingIndex: number | null = null;
    let editingValue = '';
    let dragIndex: number | null = null;

    function commit(next: string[]) {
        tasks = next;
        dispatch('change', next);
    }

    function add() {
        const v = draft.trim();
        if (!v) return;
        if (tasks.length >= MAX) return;
        commit([...tasks, v]);
        draft = '';
    }

    function onKey(e: KeyboardEvent) {
        if (e.key === 'Enter') { e.preventDefault(); add(); }
    }

    function onPaste(e: ClipboardEvent) {
        const text = e.clipboardData?.getData('text') ?? '';
        if (!text.includes('\n')) return;
        e.preventDefault();
        const lines = text.split(/\r?\n/).map((l) => l.trim()).filter(Boolean);
        const room = MAX - tasks.length;
        if (room <= 0) return;
        commit([...tasks, ...lines.slice(0, room)]);
        draft = '';
    }

    function remove(i: number) {
        commit(tasks.filter((_, idx) => idx !== i));
    }

    function startEdit(i: number) {
        editingIndex = i;
        editingValue = tasks[i];
    }
    function saveEdit() {
        if (editingIndex === null) return;
        const v = editingValue.trim();
        if (!v) return;
        commit(tasks.map((t, idx) => (idx === editingIndex ? v : t)));
        editingIndex = null;
    }
    function cancelEdit() { editingIndex = null; }

    function onDragStart(i: number) { dragIndex = i; }
    function onDragOver(e: DragEvent, i: number) {
        if (dragIndex === null || dragIndex === i) return;
        e.preventDefault();
    }
    function onDrop(i: number) {
        if (dragIndex === null || dragIndex === i) { dragIndex = null; return; }
        const next = [...tasks];
        const [moved] = next.splice(dragIndex, 1);
        next.splice(i, 0, moved);
        commit(next);
        dragIndex = null;
    }
</script>

<div class="editor">
    <div class="add-row">
        <div class="input-wrap">
            <input
                type="text"
                placeholder="Add a task and press Enter…"
                bind:value={draft}
                on:keydown={onKey}
                on:paste={onPaste}
                disabled={tasks.length >= MAX}
                aria-label="New task"
            />
            <span class="counter" class:full={tasks.length === MAX}>{tasks.length} / {MAX}</span>
        </div>
        <button type="button" class="add-btn" on:click={add} disabled={tasks.length >= MAX || !draft.trim()}>Add</button>
    </div>

    <ol class="chips">
        {#each tasks as task, i (i + '-' + task)}
            <li
                class="chip"
                draggable="true"
                on:dragstart={() => onDragStart(i)}
                on:dragover={(e) => onDragOver(e, i)}
                on:drop={() => onDrop(i)}
            >
                <span class="grip" aria-hidden="true">⋮⋮</span>
                <span class="num">{i + 1}</span>
                {#if editingIndex === i}
                    <input
                        class="edit-input"
                        type="text"
                        bind:value={editingValue}
                        on:keydown={(e) => { if (e.key === 'Enter') { e.preventDefault(); saveEdit(); } if (e.key === 'Escape') cancelEdit(); }}
                        autofocus
                    />
                    <Cluster gap="s">
                        <button type="button" class="mini accent" on:click={saveEdit}>save</button>
                        <button type="button" class="mini" on:click={cancelEdit}>cancel</button>
                    </Cluster>
                {:else}
                    <span class="body">{task}</span>
                    <Cluster gap="s">
                        <button type="button" class="mini" on:click={() => startEdit(i)} aria-label="Edit task {i + 1}">edit</button>
                        <button type="button" class="mini" on:click={() => remove(i)} aria-label="Remove task {i + 1}">remove</button>
                    </Cluster>
                {/if}
            </li>
        {/each}
    </ol>

    <p class="hint">Tip: paste 16 lines at once to fill the pack.</p>
</div>

<style>
    .editor { display: flex; flex-direction: column; gap: var(--space-3); }

    .add-row { display: flex; gap: var(--space-2); align-items: stretch; }
    .input-wrap { flex: 1; position: relative; }
    .input-wrap input {
        width: 100%;
        background: var(--paper);
        border: 1px solid var(--hairline);
        border-radius: var(--radius-md);
        padding: 0.85rem 4.5rem 0.85rem 1rem;
        color: var(--ink);
        font-size: 0.95rem;
        outline: none;
        min-height: 48px;
        transition: border-color var(--dur-fast) var(--ease), box-shadow var(--dur-fast) var(--ease);
    }
    .input-wrap input:focus {
        border-color: var(--accent-from);
        box-shadow: 0 0 0 3px var(--accent-ring);
    }
    .counter {
        position: absolute;
        right: 0.85rem; top: 50%; transform: translateY(-50%);
        font-size: 0.78rem; color: var(--ink-3);
        font-variant-numeric: tabular-nums;
    }
    .counter.full { color: #c44e4e; }

    .add-btn {
        background: var(--accent);
        color: var(--on-accent);
        border: 1px solid rgba(255, 255, 255, 0.5);
        border-radius: var(--radius-md);
        font-family: var(--font-body);
        font-weight: 500;
        padding: 0 1.2rem;
        min-width: 80px;
        cursor: pointer;
        font-size: 0.85rem;
        min-height: 48px;
    }
    .add-btn:disabled { opacity: 0.45; cursor: not-allowed; }

    .chips {
        list-style: none; margin: 0; padding: 0;
        display: flex; flex-direction: column;
        gap: var(--space-2);
    }
    .chip {
        display: grid;
        grid-template-columns: auto auto 1fr auto;
        gap: var(--space-3);
        align-items: center;
        background: var(--paper);
        border: 1px solid var(--hairline);
        border-radius: var(--radius-md);
        padding: 0.6rem 0.9rem;
    }
    .grip {
        cursor: grab;
        color: var(--ink-4);
        font-size: 1rem;
        user-select: none;
        line-height: 1;
    }
    .num {
        color: var(--ink-3);
        font-size: 0.78rem;
        font-variant-numeric: tabular-nums;
        min-width: 1.5em;
    }
    .body { font-size: 0.9rem; color: var(--ink); }

    .edit-input {
        background: transparent;
        border: 1px solid var(--accent-from);
        border-radius: var(--radius-sm);
        padding: 0.4rem 0.6rem;
        color: var(--ink);
        font-size: 0.9rem;
        outline: none;
    }

    .mini {
        background: transparent;
        border: 1px solid var(--hairline);
        border-radius: var(--radius-pill);
        padding: 0.25rem 0.6rem;
        font-size: 0.7rem;
        color: var(--ink-2);
        cursor: pointer;
        font-family: var(--font-body);
    }
    .mini:hover { color: var(--ink); }
    .mini.accent {
        background: var(--accent);
        color: var(--on-accent);
        border-color: rgba(255,255,255,0.5);
    }

    .hint {
        font-size: 0.78rem;
        color: var(--ink-3);
        font-style: italic;
    }
</style>

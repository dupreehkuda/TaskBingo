<script lang="ts">
    import { createEventDispatcher } from 'svelte';

    const dispatch = createEventDispatcher<{ submit: number }>();

    let entered = '';

    function press(d: string) {
        if (entered.length >= 2) return;
        entered = entered + d;
    }
    function back() {
        entered = entered.slice(0, -1);
    }
    function submit() {
        const n = Number(entered);
        if (Number.isFinite(n) && n >= 1 && n <= 16) {
            dispatch('submit', n);
        }
        entered = '';
    }

    function onKey(e: KeyboardEvent) {
        if (e.key >= '0' && e.key <= '9') { e.preventDefault(); press(e.key); }
        else if (e.key === 'Backspace') { e.preventDefault(); back(); }
        else if (e.key === 'Enter') { e.preventDefault(); submit(); }
    }
</script>

<div class="pad" on:keydown={onKey} tabindex="0" role="group" aria-label="Number pad">
    <div class="display" aria-live="polite">{entered || '·'}</div>

    <div class="grid">
        {#each ['1','2','3','4','5','6','7','8','9'] as d}
            <button type="button" class="key" on:click={() => press(d)} aria-label={`Number ${d}`}>{d}</button>
        {/each}
        <button type="button" class="key util" on:click={back} aria-label="Backspace">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round" width="20" height="20"><path d="M21 6h-12l-7 6 7 6h12a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2z"/><path d="M18 9l-6 6"/><path d="M12 9l6 6"/></svg>
        </button>
        <button type="button" class="key" on:click={() => press('0')} aria-label="Number 0">0</button>
        <button type="button" class="key submit" on:click={submit} aria-label="Submit">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round" width="20" height="20"><path d="M5 12h14"/><path d="M13 6l6 6-6 6"/></svg>
        </button>
    </div>
</div>

<style>
    .pad {
        display: flex;
        flex-direction: column;
        gap: var(--space-3);
        outline: none;
    }
    .pad:focus-visible { outline: 2px solid var(--accent-ring); outline-offset: 4px; border-radius: var(--radius-lg); }

    .display {
        background: var(--paper-soft);
        border: 1px solid var(--hairline);
        border-radius: var(--radius-md);
        text-align: center;
        font-family: var(--font-display);
        font-size: 1.4rem;
        padding: var(--space-3);
        color: var(--ink);
        font-variant-numeric: tabular-nums;
        min-height: 44px;
    }

    .grid {
        display: grid;
        grid-template-columns: repeat(3, 1fr);
        gap: var(--space-2);
    }

    .key {
        background: var(--paper);
        border: 1px solid var(--hairline);
        border-radius: var(--radius-md);
        padding: var(--space-3);
        font-family: var(--font-display);
        font-size: 1.15rem;
        color: var(--ink);
        cursor: pointer;
        min-height: 48px;
        transition: background var(--dur-fast) var(--ease);
    }
    .key:hover { background: var(--paper-soft); }
    .key:active { transform: translateY(1px); }
    .key:focus-visible { outline: 2px solid var(--accent-ring); outline-offset: 2px; }

    .key.util { color: var(--ink-2); }
    .key.submit {
        background: var(--accent);
        color: var(--on-accent);
        border-color: rgba(255,255,255,0.5);
    }
    .key.submit:hover { filter: brightness(1.04); }
</style>

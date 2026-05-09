<script lang="ts">
    /**
     * TabBar — segmented control. Two-way bind on `value`. Same component
     * is used as a tab strip on desktop and as a segmented control on mobile.
     */
    export let value: string;
    export let items: { value: string; label: string }[] = [];
    export let ariaLabel = 'Sections';
</script>

<div class="tabs" role="tablist" aria-label={ariaLabel}>
    {#each items as item (item.value)}
        <button
            type="button"
            role="tab"
            class="tab"
            class:active={value === item.value}
            aria-selected={value === item.value}
            on:click={() => (value = item.value)}
        >
            {item.label}
        </button>
    {/each}
</div>

<style>
    .tabs {
        display: inline-flex;
        gap: 2px;
        padding: 4px;
        background: var(--paper-soft);
        border-radius: var(--radius-pill);
        border: 1px solid var(--hairline);
    }

    .tab {
        appearance: none;
        background: transparent;
        border: none;
        padding: 0.45rem 0.95rem;
        border-radius: var(--radius-pill);
        font-family: var(--font-body);
        font-size: 0.78rem;
        font-weight: 500;
        color: var(--ink-2);
        cursor: pointer;
        min-height: 36px;
        transition: background var(--dur-fast) var(--ease),
                    color var(--dur-fast) var(--ease);
    }

    .tab:hover { color: var(--ink); }
    .tab.active {
        background: var(--paper);
        color: var(--ink);
        box-shadow: 0 1px 2px rgba(40, 55, 95, 0.06);
    }
    .tab:focus-visible {
        outline: 2px solid var(--accent-ring);
        outline-offset: 2px;
    }
</style>

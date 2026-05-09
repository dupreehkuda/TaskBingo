<script lang="ts">
    type Variant = 'accent' | 'ghost' | 'quiet';
    type Size = 'sm' | 'md' | 'lg';

    export let variant: Variant = 'accent';
    export let size: Size = 'md';
    export let href: string | undefined = undefined;
    export let type: 'button' | 'submit' | 'reset' = 'button';
    export let disabled = false;
    export let fullWidth = false;
</script>

{#if href !== undefined}
    <a
        class="btn btn--{variant} btn--{size}"
        class:full={fullWidth}
        {href}
        role="button"
        on:click
        on:mouseenter
        on:mouseleave
        on:focus
        {...$$restProps}
    >
        <slot />
    </a>
{:else}
    <button
        class="btn btn--{variant} btn--{size}"
        class:full={fullWidth}
        {type}
        {disabled}
        on:click
        on:mouseenter
        on:mouseleave
        on:focus
        {...$$restProps}
    >
        <slot />
    </button>
{/if}

<style>
    .btn {
        display: inline-flex;
        align-items: center;
        justify-content: center;
        gap: var(--space-2);
        border-radius: var(--radius-pill);
        font-family: var(--font-body);
        font-weight: 500;
        letter-spacing: 0.005em;
        cursor: pointer;
        transition: filter var(--dur-base) var(--ease),
                    box-shadow var(--dur-base) var(--ease),
                    transform var(--dur-base) var(--ease),
                    background var(--dur-base) var(--ease),
                    color var(--dur-base) var(--ease),
                    border-color var(--dur-base) var(--ease);
        text-decoration: none;
        border: 1px solid transparent;
        white-space: nowrap;
        will-change: transform;
    }

    .btn:focus-visible {
        outline: 2px solid var(--accent-ring);
        outline-offset: 2px;
    }

    .btn:disabled {
        opacity: 0.5;
        cursor: not-allowed;
    }

    .btn:not(:disabled):active { transform: translateY(1px) scale(0.99); }

    .btn--accent:not(:disabled):hover {
        filter: brightness(1.06);
        transform: translateY(-1px);
        box-shadow: 0 14px 28px -8px var(--accent-shadow);
    }
    .btn--ghost:not(:disabled):hover {
        background: var(--paper-soft);
        border-color: rgba(40, 55, 95, 0.14);
        transform: translateY(-1px);
    }
    .btn--quiet:not(:disabled):hover {
        color: var(--accent-from);
        border-bottom-color: var(--accent-from);
    }
    .btn--quiet:not(:disabled):active { transform: none; }

    .full { width: 100%; }

    /* Sizes */
    .btn--sm { padding: 0.45rem 0.95rem; font-size: 0.78rem; min-height: 32px; }
    .btn--md { padding: 0.62rem 1.2rem;  font-size: 0.85rem; min-height: 44px; }
    .btn--lg { padding: 0.78rem 1.5rem;  font-size: 0.95rem; min-height: 52px; }

    /* Variants */
    .btn--accent {
        background: var(--accent);
        color: var(--on-accent);
        border-color: rgba(255, 255, 255, 0.5);
        box-shadow: 0 8px 18px -6px var(--accent-shadow);
    }

    .btn--ghost {
        background: var(--paper);
        color: var(--ink);
        border-color: var(--hairline);
    }

    .btn--quiet {
        background: transparent;
        color: var(--ink);
        border: none;
        border-bottom: 1px solid rgba(40, 55, 95, 0.25);
        border-radius: 0;
        padding: 0 0 1px 0;
        min-height: auto;
    }
</style>

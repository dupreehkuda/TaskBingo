<script lang="ts">
    /**
     * Paper — solid ivory content surface. The workhorse.
     * Use for cards, list items, forms, modal bodies, the bingo grid container.
     * Never nest paper inside paper — separate via spacing instead.
     */
    type Padding = 'sm' | 'md' | 'lg';

    export let padding: Padding = 'md';
    export let radius: 'md' | 'lg' = 'lg';
</script>

<div class="paper paper--p-{padding} paper--r-{radius}" {...$$restProps}>
    <slot />
</div>

<style>
    .paper {
        background: var(--paper);
        border: 1px solid var(--hairline);
        color: var(--ink);
        box-shadow:
            0 1px 0 rgba(255, 255, 255, 0.6) inset,
            0 2px 6px -2px rgba(40, 55, 95, 0.05),
            0 14px 32px -16px rgba(40, 55, 95, 0.10);
    }

    .paper--p-sm { padding: var(--space-3) var(--space-4); }
    .paper--p-md { padding: var(--space-4) var(--space-5); }
    .paper--p-lg { padding: var(--space-5) var(--space-6); }

    .paper--r-md { border-radius: var(--radius-md); }
    .paper--r-lg { border-radius: var(--radius-lg); }
</style>

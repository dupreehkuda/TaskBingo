<script lang="ts">
    /**
     * Page — outer wrapper for every route. Caps width and pads the gutters.
     * Use as the outermost element on each route's +page.svelte.
     */
    type Width = 'narrow' | 'normal' | 'wide';

    export let width: Width = 'normal';
</script>

<main class="page page--w-{width}" {...$$restProps}>
    <slot />
</main>

<style>
    .page {
        margin: 0 auto;
        padding: var(--space-6) var(--space-5) var(--space-7);
        width: 100%;
        animation: page-in var(--dur-base) var(--ease);
    }
    .page--w-narrow { max-width: 28rem; }
    .page--w-normal { max-width: 60rem; }
    .page--w-wide   { max-width: 80rem; }

    @media (max-width: 640px) {
        .page { padding: var(--space-5) var(--space-4) var(--space-6); }
    }

    @keyframes page-in {
        from { opacity: 0; transform: translateY(8px); }
        to   { opacity: 1; transform: translateY(0); }
    }

    @media (prefers-reduced-motion: reduce) {
        .page { animation: none; }
    }
</style>

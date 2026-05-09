<script lang="ts">
    /**
     * Atmosphere — paints the page canvas. Render once at the layout root.
     * It applies the mesh gradient + noise to <body> via :global on mount and
     * cleans up on destroy, so SvelteKit page navigation never flashes white.
     */
    import { onDestroy, onMount } from 'svelte';
    import { browser } from '$app/environment';

    onMount(() => {
        if (!browser) return;
        document.body.classList.add('atmos');
    });

    onDestroy(() => {
        if (!browser) return;
        document.body.classList.remove('atmos');
    });
</script>

<slot />

<style>
    :global(body.atmos) {
        position: relative;
        background:
            radial-gradient(58% 50% at 18% 22%, var(--atmos-sky) 0%, transparent 64%),
            radial-gradient(48% 48% at 84% 80%, var(--atmos-lilac) 0%, transparent 66%),
            linear-gradient(140deg, var(--atmos-base-from) 0%, var(--atmos-base-to) 100%);
        background-attachment: fixed;
    }

    /* Noise overlay sits in a fixed pseudo-element so it doesn't scroll. */
    :global(body.atmos)::before {
        content: '';
        position: fixed;
        inset: 0;
        pointer-events: none;
        z-index: 0;
        mix-blend-mode: overlay;
        opacity: 0.5;
        background-image: url("data:image/svg+xml;utf8,<svg xmlns='http://www.w3.org/2000/svg' width='160' height='160'><filter id='n'><feTurbulence type='fractalNoise' baseFrequency='0.85' numOctaves='2' stitchTiles='stitch'/><feColorMatrix values='0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.04 0'/></filter><rect width='100%25' height='100%25' filter='url(%23n)'/></svg>");
    }

    /* Page content sits above the noise overlay. */
    :global(body.atmos) > * {
        position: relative;
        z-index: 1;
    }
</style>

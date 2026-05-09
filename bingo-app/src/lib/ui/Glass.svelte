<script lang="ts">
    /**
     * Glass — frosted chrome surface. Rare. Use only for sticky headers
     * (when scrolled), the side panel during /game, and modal overlays.
     * Never nest glass inside glass; never put another glass inside this slot.
     */
    type Radius = 'pill' | 'card';
    type Blur = 'soft' | 'strong';

    export let radius: Radius = 'card';
    export let blur: Blur = 'strong';
</script>

<div class="glass glass--r-{radius} glass--b-{blur}" {...$$restProps}>
    <slot />
</div>

<style>
    .glass {
        background: rgba(252, 253, 255, 0.55);
        border: 1px solid rgba(255, 255, 255, 0.65);
        color: var(--ink);
        box-shadow:
            0 22px 44px -18px rgba(40, 55, 95, 0.16),
            inset 0 1px 0 rgba(255, 255, 255, 0.7);
    }

    .glass--r-pill { border-radius: var(--radius-pill); }
    .glass--r-card { border-radius: var(--radius-lg); }

    .glass--b-soft {
        backdrop-filter: blur(16px) saturate(125%);
        -webkit-backdrop-filter: blur(16px) saturate(125%);
    }
    .glass--b-strong {
        backdrop-filter: blur(28px) saturate(130%);
        -webkit-backdrop-filter: blur(28px) saturate(130%);
    }
</style>

<script lang="ts">
    type Gap = 's' | 'm' | 'l' | 'xl';

    export let gap: Gap = 'm';
    export let align: 'start' | 'center' | 'stretch' = 'stretch';
</script>

<div class="stack stack--g-{gap} stack--a-{align}" {...$$restProps}>
    <slot />
</div>

<style>
    .stack { display: flex; flex-direction: column; }
    .stack--g-s  { gap: var(--space-2); }
    .stack--g-m  { gap: var(--space-4); }
    .stack--g-l  { gap: var(--space-5); }
    .stack--g-xl { gap: var(--space-6); }
    .stack--a-start  { align-items: flex-start; }
    .stack--a-center { align-items: center; }
    .stack--a-stretch { align-items: stretch; }
</style>

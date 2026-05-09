<script lang="ts">
    type Gap = 's' | 'm' | 'l';

    export let gap: Gap = 'm';
    export let align: 'start' | 'center' | 'baseline' = 'center';
    export let justify: 'start' | 'center' | 'end' | 'between' = 'start';
</script>

<div class="cluster cluster--g-{gap} cluster--a-{align} cluster--j-{justify}" {...$$restProps}>
    <slot />
</div>

<style>
    .cluster { display: flex; flex-wrap: wrap; }
    .cluster--g-s { gap: var(--space-2); }
    .cluster--g-m { gap: var(--space-3); }
    .cluster--g-l { gap: var(--space-5); }
    .cluster--a-start    { align-items: flex-start; }
    .cluster--a-center   { align-items: center; }
    .cluster--a-baseline { align-items: baseline; }
    .cluster--j-start   { justify-content: flex-start; }
    .cluster--j-center  { justify-content: center; }
    .cluster--j-end     { justify-content: flex-end; }
    .cluster--j-between { justify-content: space-between; }
</style>

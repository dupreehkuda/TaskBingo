<script lang="ts">
    import { onMount } from 'svelte';
    import Header from '$lib/feature/Header.svelte';
    import Atmosphere from '$lib/ui/Atmosphere.svelte';
    import { bootstrapAuth } from '$lib/stores/auth';
    import '../app.postcss';

    onMount(() => {
        // One-shot ping on app load: detects an existing HttpOnly cookie and
        // syncs the local Auth flag so the header reflects reality.
        bootstrapAuth();
    });
</script>

<svelte:head>
    <link rel="icon" type="image/svg" href="/favicon.svg" />
</svelte:head>

<a href="#main" class="skip-link">Skip to main content</a>
<Atmosphere>
    <Header />
    <div id="main">
        <slot />
    </div>
</Atmosphere>

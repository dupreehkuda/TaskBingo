<script lang="ts">
    import {_GameHandler} from "./+page"
    import { onMount } from 'svelte';

    let closer: () => void;

    onMount(() => {closer = _GameHandler()});
    window.addEventListener('beforeunload', function() {
        closer();
    })
</script>

<svelte:head>
    <title>Game</title>
    <meta name="theme-color" content="#D6E2EC">
</svelte:head>

<body>
    <main>
        <p>hi mom.</p>
    </main>
</body>

<style>
    main {
        text-align: center;
        padding: 1em;
        max-width: 440px;
        margin: 0 auto;
    }
</style>
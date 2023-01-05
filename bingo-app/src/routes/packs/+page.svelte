<script lang="ts">
    import { Button, Card } from "flowbite-svelte";
	import { load } from "./+page";
    import { onMount } from "svelte";

    onMount(() => {
        load
    })

    export let data;
    const { userInfo } = data
</script>

<main>
    <h1>nothing here yet</h1>
</main>


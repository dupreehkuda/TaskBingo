<script lang="ts">
    import { Button } from 'flowbite-svelte';

    const newResp = {
        userID: 'dupreeh'
    }

    let userID = 'dupreeh'
    let result: any = null

    async function doPost() {
        const res = await fetch('https://taskbingo.com/api/user/getUserData', {
            method: 'POST',
            headers: {'Origin': 'taskbingo.com'},
            body: JSON.stringify(newResp)
        })

        const json = await res.json()
        console.log(JSON.stringify(json))
        result = JSON.stringify(json)
    }
</script>

<main>
    <ul style="padding-left:0px">
        <Button gradient color="teal" on:click={doPost}>Get Points</Button>

        <li style="display:inline">Result: {result}</li>
    </ul>
</main>

<style>
    main {
        text-align: center;
        padding: 1em;
        max-width: 240px;
        margin: 0 auto;
    }
</style>


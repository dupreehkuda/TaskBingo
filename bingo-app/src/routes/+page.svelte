<script lang="ts">
    import Header from "./Header.svelte";
    import GetPoints from "./GetPoints.svelte";
</script>

<header>
<!--    <h1>Sveltish Bingo</h1>-->
    <Header/>
    <GetPoints/>
</header>

<style>
    main {
        text-align: center;
        padding: 1em;
        max-width: 240px;
        margin: 0 auto;
    }

    header {
        text-align: center;
        padding: 1em;
        /*max-width: 500px;*/
        margin: 0 auto;
    }

    h3 {
        color: #ff3e00;
        font-size: 2em;
        font-weight: 100;
    }

    h1 {
        color: teal;
        text-transform: uppercase;
        font-size: 4em;
        font-weight: 100;
    }

    @media (min-width: 640px) {
        main {
            max-width: none;
        }
    }
</style>
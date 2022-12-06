<script>
    import { Label, Input, Helper, Button } from 'flowbite-svelte'
    import Header from "../Header.svelte";
</script>

<header>
    <Header/>
</header>
<main>
    <Label class='block mb-3'>Login</Label>
    <div class="mb-4">
        <Input label="Email" id="email" name="email" required placeholder="john@example.com"/>
    </div>

    <div class="mb-4">
        <Input label="Password" id="password" name="password" required placeholder="password"/>
    </div>

<!--    <Helper class='text-sm mt-2'>We will not send you any marketing materials. We promise.</Helper>-->

    <Button href="/register" gradient color="teal">Register</Button>
    <Button gradient color="teal">Login</Button>
</main>

<style>
    main {
        text-align: center;
        padding: 1em;
        max-width: 440px;
        margin: 0 auto;
    }

    header {
        text-align: center;
        padding: 1em;
        /*max-width: 500px;*/
        margin: 0 auto;
    }
</style>
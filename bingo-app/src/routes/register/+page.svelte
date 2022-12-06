<script>
    import Header from "../Header.svelte";
    import {Label, Input, Button} from "flowbite-svelte";
    import {inputClass} from "flowbite-svelte/forms/Radio.svelte";

    const newResp = {
        name: inputClass.name,
        email: inputClass.email,
        password: inputClass.password
    }

    async function register() {
        const res = await fetch('http://localhost:8081/api/register', {
            method: 'POST',
            body: JSON.stringify(newResp)
        })

        const json = await res.json()
        console.log(JSON.stringify(json))
    }
</script>

<header>
    <Header/>
</header>
<main>
    <Label class='block mb-3'>Register</Label>
    <div class="mb-4">
        <Label for="input-group-1" class="block mb-2">Name</Label>
        <Input label="Name" id="name" name="name" required placeholder="John"/>
    </div>

    <div class="mb-4">
        <Label for="input-group-1" class="block mb-2">Email</Label>
        <Input label="Email" id="email" name="email" required placeholder="john@example.com"/>
    </div>

    <div class="mb-4">
        <Label for="input-group-1" class="block mb-2">Password</Label>
        <Input label="Password" id="password" name="password" required placeholder="password"/>
    </div>

    <div class="mb-4">
        <Input label="Password" id="password" name="password" required placeholder="re-enter password"/>
    </div>

    <Button href="/" gradient color="teal" on:click={register}>Register</Button>
</main>

<style>
    header {
        text-align: center;
        padding: 1em;
        margin: 0 auto;
    }

    main {
        text-align: center;
        padding: 1em;
        max-width: 440px;
        margin: 0 auto;
    }
</style>
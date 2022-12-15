<script>
    import Header from "../Header.svelte";
    import {ButtonGroup, InputAddon, Label, Input, Button} from "flowbite-svelte";


    let show = false;

    async function submit(e) {
        e.preventDefault()

        const formData = new FormData(e.target);
        const data = {};
        for (let field of formData) {
            const [key, value] = field;
            data[key] = value;
        }

        const newResp = {
            login: data.name,
            email: data.email,
            password: data.password
        }

        console.log(data)
        console.log(JSON.stringify(newResp))

        const res = await fetch('https://taskbingo.com/api/user/register', {
            method: 'POST',
            headers: {'Origin': 'taskbingo.com'},
            body: JSON.stringify(newResp)
        })

        console.log(res.headers.get("set-cookie"))

        const json = await res.json()
        console.log(JSON.stringify(json))
    }
</script>

<header>
    <Header/>
</header>

<main>
    <Label class='block mb-3'>Register</Label>
    <form on:submit={submit}>
        <div class="mb-4">
            <Label for="input-group-1" class="block mb-2">Name</Label>
            <Input label="Name" id="name" name="name" required placeholder="John"/>
        </div>

        <div class="mb-4">
            <Label for="input-group-1" class="block mb-2">Email</Label>
            <Input label="Email" id="email" name="email" required placeholder="john@example.com"/>
        </div>

        <Label for="password" class="mb-2">Your password</Label>
        <ButtonGroup class="mb-3 w-full">
            <InputAddon>
                <button on:click={() => (show = !show)}>
                    {#if show}
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5"><path stroke-linecap="round" stroke-linejoin="round" d="M2.036 12.322a1.012 1.012 0 010-.639C3.423 7.51 7.36 4.5 12 4.5c4.638 0 8.573 3.007 9.963 7.178.07.207.07.431 0 .639C20.577 16.49 16.64 19.5 12 19.5c-4.638 0-8.573-3.007-9.963-7.178z" /><path stroke-linecap="round" stroke-linejoin="round" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" /></svg>
                    {:else}
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5"><path stroke-linecap="round" stroke-linejoin="round" d="M3.98 8.223A10.477 10.477 0 001.934 12C3.226 16.338 7.244 19.5 12 19.5c.993 0 1.953-.138 2.863-.395M6.228 6.228A10.45 10.45 0 0112 4.5c4.756 0 8.773 3.162 10.065 7.498a10.523 10.523 0 01-4.293 5.774M6.228 6.228L3 3m3.228 3.228l3.65 3.65m7.894 7.894L21 21m-3.228-3.228l-3.65-3.65m0 0a3 3 0 10-4.243-4.243m4.242 4.242L9.88 9.88" /></svg>
                    {/if}
                </button>
            </InputAddon>
            <Input id="password" name="password" type={show ? 'text' : 'password'} placeholder="Password" />
        </ButtonGroup>

        <Button class="mb-3" type="submit" gradient color="teal" on:click={submit()}>Register</Button>
    </form>

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
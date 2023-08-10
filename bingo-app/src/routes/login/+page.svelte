<script lang='ts'>
    import { ButtonGroup, InputAddon, Input, Button } from 'flowbite-svelte'
    import { goto } from '$app/navigation';
    import { API_URL, WEB_URL } from '../temporary';

    let show = false;

    export const ssr = false

    async function submit(e: any) {
        e.preventDefault()

        const formData = new FormData(e.target);
        const data: any = {};
        for (let field of formData) {
            const [key, value] = field;
            data[key] = value;
        }

        const newReq = {
            username: data.username,
            password: data.password
        }

        const res = await fetch(`${API_URL}/api/user/login`, {
            method: 'POST',
            headers: {'Origin': WEB_URL},
            body: JSON.stringify(newReq),
            credentials: 'include',
        })
        
        if (res.ok) { goto('/account'); }
    }
</script>

<title>Login</title>
<body>
    <main>
        <!-- <Label color="disabled" class='text block mb-3 dark'>Login</Label> -->
        <form on:submit={submit}>
            <div class="mb-4">
                <Input label="login" id="username" name="username" required placeholder="Login"/>
            </div>

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

            <div class="mb-4">
                <Button href="/register" color="light">Register</Button>
                <Button type="submit" color="light">Login</Button>
            </div>
        </form>
    </main>
</body>

<style>
    main {
        text-align: center;
        padding: 1em;
        max-width: 440px;
        margin: 0 auto;
        background-color: #07417b;
        color: #07417b;
    }
</style>
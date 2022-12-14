<script>
    import { Label, Input, Button } from 'flowbite-svelte'
    import Header from "../Header.svelte";

    async function submit(e) {
        e.preventDefault()

        const formData = new FormData(e.target);
        const data = {};
        for (let field of formData) {
            const [key, value] = field;
            data[key] = value;
        }

        const newResp = {
            login: data.login,
            password: data.password
        }

        console.log(data)
        console.log(JSON.stringify(newResp))

        const res = await fetch('http://localhost:8082/api/user/login', {
            method: 'POST',
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
    <Label class='block mb-3'>Login</Label>
    <form on:submit={submit}>
        <div class="mb-4">
            <Input label="login" id="login" name="login" required placeholder="Login"/>
        </div>

        <div class="mb-4">
            <Input label="Password" id="password" name="password" required placeholder="password"/>
        </div>

        <Button href="/register" gradient color="teal">Register</Button>
        <Button type="submit" gradient color="teal">Login</Button>
    </form>
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
<script lang='ts'>
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

        const res = await fetch('https://taskbingo.com/api/user/login', {
            method: 'POST',
            headers: {'Origin': 'taskbingo.com'},
            body: JSON.stringify(newResp)
        })

        console.log(res.headers.get("set-cookie"))

        const json = await res.json()
        console.log(JSON.stringify(json))
    }
</script>

<title>Login</title>
<body>
    <main>
        <Label color="disabled" class='text block mb-3 dark'>Login</Label>
        <form on:submit={submit}>
            <div class="dark mb-4">
                <Input label="login" id="login" name="login" required placeholder="Login"/>
            </div>

            <div class="dark mb-4">
                <Input label="Password" id="password" name="password" required placeholder="Password"/>
            </div>

            <div class="dark mb-4">
                <Button href="/register" gradient color="teal">Register</Button>
                <Button type="submit" gradient color="teal">Login</Button>
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
        background-color: #11202d;
        color: rgb(238, 237, 197);
    }
</style>
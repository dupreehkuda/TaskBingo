export async function Submit(event: any): Promise<number> {
    event.preventDefault()
    const data = new FormData(event.target);

    const login = data.get('login') as string;
    const email = data.get('email') as string;
    const city = data.get('city') as string;
    const password = data.get('password') as string;

    const newResp = {
        login: login,
        email: email,
        city: city,
        password: password
    }

    console.log(data)
    console.log(JSON.stringify(newResp))

    const res = await fetch('https://taskbingo.com/api/user/register', {
        method: 'POST',
        headers: {'Origin': 'taskbingo.com'},
        body: JSON.stringify(newResp)
    })

    return res.status
}
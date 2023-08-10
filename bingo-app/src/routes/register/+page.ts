import { goto } from "$app/navigation";

export async function _Submit(event: any): Promise<number> {
    event.preventDefault()
    const data = new FormData(event.target);

    const username = data.get('username') as string;
    const email = data.get('email') as string;
    const city = data.get('city') as string;
    const password = data.get('password') as string;

    const newResp = {
        username: username,
        email: email,
        city: city,
        password: password
    }

    const res = await fetch('https://taskbingo.com/api/user/register', {
        method: 'POST',
        headers: {'Origin': 'taskbingo.com'},
        body: JSON.stringify(newResp)
    })


    if (res.status === 200) {
        redirectOnFinish()
    }
    
    return res.status
}

function redirectOnFinish() {
    goto('/account');
}
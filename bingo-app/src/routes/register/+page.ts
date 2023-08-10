import { goto } from "$app/navigation";
import { API_URL, WEB_URL } from '../temporary';

export const ssr = false

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

    console.log(data)
    console.log(JSON.stringify(newResp))

    const res = await fetch(`${API_URL}/api/user/register`, {
        method: 'POST',
        headers: {'Origin': WEB_URL},
        body: JSON.stringify(newResp),
        credentials: 'include',
    })
    
    if (res.status === 200) {
        redirectOnFinish()
    }
    
    return res.status
}

function redirectOnFinish() {
    goto('/account');
  }
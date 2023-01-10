import Account from '../accountStore';
import { get } from 'svelte/store';
import type { PageLoad } from './$types';

export const load = (async ({ fetch }) => {
  const userData = await fetch('https://taskbingo.com/api/user/getUserData', {
    method: 'GET',
    headers: {'Origin': 'taskbingo.com'},
  })

  const userInfo = await userData.json()

  Account.set(userInfo)

  const res = await fetch('https://taskbingo.com/api/task/getPacks', {
    method: 'GET',
    headers: {'Origin': 'taskbingo.com'},
  })

  const packs = await res.json()

  return { packs }
}) satisfies PageLoad;

export async function Like(pack: any, liked: boolean) {
  const newResp = {
    id: pack.id,
  } 

  if (liked) {
    const res = await fetch('https://taskbingo.com/api/user/dislikePack', {
      method: 'POST',
      headers: {'Origin': 'taskbingo.com'},
      body: JSON.stringify(newResp)
    })

    let account = get(Account)
    account.packs = account.packs.filter(e => e.id !== pack.id)
    Account.set(account)

  } else {
    const res = await fetch('https://taskbingo.com/api/user/likePack', {
      method: 'POST',
      headers: {'Origin': 'taskbingo.com'},
      body: JSON.stringify(newResp)
    })

    let account = get(Account)
    account.packs.push({id: pack.id, tasks: pack.tasks})
    Account.set(account)
  }
};

// export async function Rate(id: any, rated: boolean) {
//   const newResp = {
//     id: id,
//   } 

//   if (rated) {
//     const res = await fetch('https://taskbingo.com/api/user/unratePack', {
//       method: 'POST',
//       headers: {'Origin': 'taskbingo.com'},
//       body: JSON.stringify(newResp)
//     })

//     console.log(res.status)
//   } else {
//     const res = await fetch('https://taskbingo.com/api/user/ratePack', {
//       method: 'POST',
//       headers: {'Origin': 'taskbingo.com'},
//       body: JSON.stringify(newResp)
//     })

//     console.log(res.status)
//   }
// };
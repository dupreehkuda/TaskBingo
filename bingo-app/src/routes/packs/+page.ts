import type { PageLoad } from './$types';

export const load = (async ({ fetch }) => {
  const res = await fetch('https://taskbingo.com/api/task/getPacks', {
    method: 'GET',
    headers: {'Origin': 'taskbingo.com'},
  })

  const packs = await res.json()

  return { packs }
}) satisfies PageLoad;

export async function Like(id: any, liked: boolean) {
  const newResp = {
    id: id,
  } 

  if (liked) {
    const res = await fetch('https://taskbingo.com/api/user/dislikePack', {
      method: 'POST',
      headers: {'Origin': 'taskbingo.com'},
      body: JSON.stringify(newResp)
    })

    console.log(res.status)
  } else {
    const res = await fetch('https://taskbingo.com/api/user/likePack', {
      method: 'POST',
      headers: {'Origin': 'taskbingo.com'},
      body: JSON.stringify(newResp)
    })

    console.log(res.status)
  }
};

export async function Rate(id: any, rated: boolean) {
  const newResp = {
    id: id,
  } 

  if (rated) {
    const res = await fetch('https://taskbingo.com/api/user/unratePack', {
      method: 'POST',
      headers: {'Origin': 'taskbingo.com'},
      body: JSON.stringify(newResp)
    })

    console.log(res.status)
  } else {
    const res = await fetch('https://taskbingo.com/api/user/ratePack', {
      method: 'POST',
      headers: {'Origin': 'taskbingo.com'},
      body: JSON.stringify(newResp)
    })

    console.log(res.status)
  }
};
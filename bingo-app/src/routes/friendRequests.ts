import Account from './accountStore';
import { get } from 'svelte/store';
import { API_URL, WEB_URL } from './temporary';

export async function RequestFriend(personId: string, username: string) {
  const newResp = {
    person: personId,
  } 
  
  const res = await fetch(`${API_URL}/api/user/requestFriend`, {
    method: 'POST',
    headers: {'Origin': WEB_URL},
    body: JSON.stringify(newResp),
    credentials: 'include',
  })

  let account = get(Account)

  account.friends.push({userID: personId, username: username, status: 1, wins: 0, loses: 0})

  Account.set(account)
};

export async function AcceptFriend(personId: string) {
  const newResp = {
    person: personId,
  } 
  
  const res = await fetch(`${API_URL}/api/user/acceptFriend`, {
    method: 'POST',
    headers: {'Origin': WEB_URL},
    body: JSON.stringify(newResp),
    credentials: 'include',
  })

  let account = get(Account)

  const personIDX = account.friends.findIndex((friend => friend.userID == personId))
  account.friends[personIDX].status = 3

  Account.set(account)
};

export async function DeleteFriend(personId: string) {
  const newResp = {
    person: personId,
  } 
  
  const res = await fetch(`${API_URL}/api/user/deleteFriend`, {
    method: 'POST',
    headers: {'Origin': WEB_URL},
    body: JSON.stringify(newResp),
    credentials: 'include',
  })

  let account = get(Account)
  account.friends = account.friends.filter(friend => friend.userID !== personId)

  Account.set(account)
};
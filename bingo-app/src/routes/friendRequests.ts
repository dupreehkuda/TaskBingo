import Account from './accountStore';
import { get } from 'svelte/store';

export async function RequestFriend(personId: string, username: string) {
  const newResp = {
    person: personId,
  } 
  
  const res = await fetch('https://taskbingo.com/api/user/requestFriend', {
    method: 'POST',
    headers: {'Origin': 'taskbingo.com'},
    body: JSON.stringify(newResp)
  })

  let account = get(Account)

  account.friends.push({userID: personId, username: username, status: 1, wins: 0, loses: 0})

  Account.set(account)
};

export async function AcceptFriend(personId: string) {
  const newResp = {
    person: personId,
  } 
  
  const res = await fetch('https://taskbingo.com/api/user/acceptFriend', {
    method: 'POST',
    headers: {'Origin': 'taskbingo.com'},
    body: JSON.stringify(newResp)
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
  
  const res = await fetch('https://taskbingo.com/api/user/deleteFriend', {
    method: 'POST',
    headers: {'Origin': 'taskbingo.com'},
    body: JSON.stringify(newResp)
  })

  let account = get(Account)
  account.friends = account.friends.filter(friend => friend.userID !== personId)

  Account.set(account)
};
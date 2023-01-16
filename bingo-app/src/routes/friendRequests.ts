import Account from './accountStore';
import { get } from 'svelte/store';

export async function RequestFriend(person: string) {
  const newResp = {
    person: person,
  } 
  
  const res = await fetch('https://taskbingo.com/api/user/requestFriend', {
    method: 'POST',
    headers: {'Origin': 'taskbingo.com'},
    body: JSON.stringify(newResp)
  })

  let account = get(Account)

  account.friends.push({login: person, status: 1, wins: 0, loses: 0})

  Account.set(account)
};

export async function AcceptFriend(person: string) {
  const newResp = {
    person: person,
  } 
  
  const res = await fetch('https://taskbingo.com/api/user/acceptFriend', {
    method: 'POST',
    headers: {'Origin': 'taskbingo.com'},
    body: JSON.stringify(newResp)
  })

  let account = get(Account)

  const personIDX = account.friends.findIndex((friend => friend.login == person))
  account.friends[personIDX].status = 3

  Account.set(account)
};

export async function DeleteFriend(person: string) {
  const newResp = {
    person: person,
  } 
  
  const res = await fetch('https://taskbingo.com/api/user/deleteFriend', {
    method: 'POST',
    headers: {'Origin': 'taskbingo.com'},
    body: JSON.stringify(newResp)
  })

  let account = get(Account)
  account.friends = account.friends.filter(friend => friend.login !== person)

  Account.set(account)
};
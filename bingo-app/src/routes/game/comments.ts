import { API_URL, WEB_URL } from '../temporary'

export interface Comment {
    id: string
    gameID: string
    userID: string
    body: string
    createdAt: string
    updatedAt: string
}

export async function _ListComments(gameID: string): Promise<Comment[]> {
    const res = await fetch(`${API_URL}/api/game/${gameID}/comments`, {
        method: 'GET',
        headers: { 'Origin': WEB_URL },
        credentials: 'include',
    })
    if (!res.ok) return []
    return await res.json()
}

export async function _AddComment(gameID: string, body: string): Promise<Comment | null> {
    const res = await fetch(`${API_URL}/api/game/${gameID}/comments`, {
        method: 'POST',
        headers: { 'Origin': WEB_URL, 'Content-Type': 'application/json' },
        body: JSON.stringify({ body }),
        credentials: 'include',
    })
    if (!res.ok) return null
    return await res.json()
}

export async function _EditComment(commentID: string, body: string): Promise<Comment | null> {
    const res = await fetch(`${API_URL}/api/game/comments/${commentID}`, {
        method: 'PATCH',
        headers: { 'Origin': WEB_URL, 'Content-Type': 'application/json' },
        body: JSON.stringify({ body }),
        credentials: 'include',
    })
    if (!res.ok) return null
    return await res.json()
}

export async function _DeleteComment(commentID: string): Promise<boolean> {
    const res = await fetch(`${API_URL}/api/game/comments/${commentID}`, {
        method: 'DELETE',
        headers: { 'Origin': WEB_URL },
        credentials: 'include',
    })
    return res.ok
}

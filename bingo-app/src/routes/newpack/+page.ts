import { API_URL, WEB_URL } from '../temporary';

export const ssr = false;

export interface NewPackPayload {
    name: string;
    tasks: string[];
    isPrivate: boolean;
}

export async function _CreatePack(payload: NewPackPayload): Promise<number> {
    if (payload.tasks.length !== 16) {
        return 422;
    }
    const body = {
        id: '',
        pack: { title: payload.name, tasks: payload.tasks },
        isPrivate: payload.isPrivate,
    };

    const res = await fetch(`${API_URL}/api/task/setTaskPack`, {
        method: 'POST',
        headers: { 'Origin': WEB_URL },
        body: JSON.stringify(body),
        credentials: 'include',
    });

    return res.status;
}

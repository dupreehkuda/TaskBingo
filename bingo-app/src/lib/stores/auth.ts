import { writable } from 'svelte/store';
import { browser } from '$app/environment';
import { API_URL, WEB_URL } from '../../routes/temporary';

/**
 * Auth — the single source of truth for "is the user logged in" on the client.
 *
 * The actual session lives in an HttpOnly JWT cookie set by the backend; that
 * cookie is unreadable to JavaScript by design (XSS protection). For the UI to
 * know whether to show the logged-in nav, we keep a parallel non-secret flag
 * in localStorage. It is set by markAuthed() after a successful login/register
 * and cleared by markLoggedOut(). The flag is purely cosmetic: every protected
 * request still relies on the HttpOnly cookie. If the cookie ever expires
 * while the flag is still around, the next API call will 401 and we can clear
 * the flag from there (handled in `clearIfUnauthorized`).
 */
const STORAGE_KEY = 'tb-authed';

function readStoredAuth(): boolean {
    if (!browser) return false;
    try {
        return localStorage.getItem(STORAGE_KEY) === '1';
    } catch {
        return false;
    }
}

export const Auth = writable<boolean>(readStoredAuth());

export function markAuthed(): void {
    if (browser) {
        try { localStorage.setItem(STORAGE_KEY, '1'); } catch { /* ignore */ }
    }
    Auth.set(true);
}

export function markLoggedOut(): void {
    if (browser) {
        try { localStorage.removeItem(STORAGE_KEY); } catch { /* ignore */ }
    }
    Auth.set(false);
}

/**
 * Helper for fetch wrappers: if a protected response comes back 401 we know
 * the cookie is gone or expired. Drop the local flag so the header swaps back.
 */
export function clearIfUnauthorized(status: number): void {
    if (status === 401) markLoggedOut();
}

/**
 * Bootstrap on app load. Probes a cheap authenticated endpoint to detect a
 * still-valid HttpOnly cookie even when the localStorage flag is missing
 * (e.g., the user logged in before this flag mechanism existed, or in
 * another browser/device). 200 → set the flag. 401 → make sure it's cleared.
 *
 * Safe to call multiple times; idempotent.
 */
export async function bootstrapAuth(): Promise<void> {
    if (!browser) return;
    try {
        const res = await fetch(`${API_URL}/api/user/getUserData`, {
            method: 'GET',
            headers: { Origin: WEB_URL },
            credentials: 'include',
        });
        if (res.ok) markAuthed();
        else if (res.status === 401) markLoggedOut();
    } catch {
        // Network down or backend not running — leave whatever flag we already had.
    }
}

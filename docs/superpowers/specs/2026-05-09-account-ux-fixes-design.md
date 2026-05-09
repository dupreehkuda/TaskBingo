# `/account` UX fixes: tabs, fewer layers, reactive header

**Date:** 2026-05-09
**Status:** Draft (design)

## Goal

Make the logged-in `/account` page usable as data grows and fix the navigation gap that strands authenticated users on a public-only header.

Five concrete pains from user testing:

1. Too many paper cards — every friend / pack / game is its own card, the page reads as a quilt.
2. No way to view a pack in full from `/account` (only 3 tasks + "+ N more").
3. No visible entry to start a solo game once on `/account`.
4. Header keeps showing the logged-out nav after a successful login or registration.
5. The page is one tall vertical scroll: friends, packs, games stacked. Easy to lose orientation as collections grow.

All five hang together as "make the dashboard a snapshot worth coming back to". One spec.

## Non-goals

- No backend changes. The `/api/user/getUserData` and `/api/user/stats` endpoints stay as they are.
- No logout flow. Out of scope (not in feedback).
- No `/games` route. Past games stay behind a "show N past" toggle inside the Games tab.
- No global state library (no Zustand, no Pinia-equivalent). The new auth signal lives in a single Svelte writable.
- No restyle of `/people` or `/packs`. They already use lists of paper cards correctly — each row IS a content unit there.
- No search inside `/account`. Each tab is short-list scoped; full search lives on `/people` and `/packs`.

## Approach

### 1. Auth state goes reactive

**Problem.** `Header.svelte` reads the JWT cookie once in `onMount` and stores `authed` as a local boolean. After a successful `/login` or `/register` the page navigates with `goto('/account')`, but the layout (and therefore Header) is not remounted. `authed` stays `false`, the nav keeps showing only `about | login`.

**Fix.** A new writable store `bingo-app/src/lib/stores/auth.ts`:

```ts
import { writable } from 'svelte/store';
import { browser } from '$app/environment';

function readCookieAuth(): boolean {
    if (!browser) return false;
    try {
        const m = RegExp('auth=[^;]+').exec(document.cookie);
        if (!m) return false;
        const token = decodeURIComponent(m.toString().replace(/^[^=]+./, ''));
        const payload = token.split('.')[1];
        const json = atob(payload.replace(/-/g, '+').replace(/_/g, '/'));
        return !!JSON.parse(json).user;
    } catch {
        return false;
    }
}

export const Auth = writable<boolean>(readCookieAuth());
export function markAuthed() { Auth.set(true); }
export function markLoggedOut() { Auth.set(false); }
```

`Header.svelte` subscribes via `$Auth` instead of holding a local `authed`. Login and register call `markAuthed()` after a successful response, before `goto('/account')`. Cookie remains the source of truth for cross-tab and reload (the store is initialised from the cookie on module load).

This also future-proofs logout — any future logout helper just calls `markLoggedOut()` and the header reacts.

### 2. `/account` becomes a tabbed dashboard with minimal layering

The vertical reading order:

```
[Greeting]            text on atmosphere — username, total bingos, city
[StatsBoard]          tiles + flat charts directly on atmosphere
[Continue playing]    glass strip — only when an active game exists
[TabBar]              segmented control on atmosphere: Friends · Packs · Games
[Tab content]         one paper card; rows separated by hairlines
```

**Layer budget per page** (counting wrapping `<Paper>` / `<Glass>` containers, not the small `--paper-soft` tiles used as data points):

- `<Paper>` containers: exactly **1** at all times (the active tab body).
- `<Glass>` containers: **0–1** depending on whether a featured game exists, plus the Header when scrolled past 64px.
- Everything else lives directly on the atmosphere.

**Greeting.** Already on atmosphere. Stays. No change.

**StatsBoard — strip the wrappers.** The current component renders three paper cards (one for each chart, one for top packs). Drop those wrappers. Tiles row stays (4 paper-soft tiles — small, data-presentation, not "section card"). Bar charts render with their internal `--paper-soft` tracks giving them anchor; hairlines separate the section title from the bars. Top packs renders as a hairline-tracked list inline. The whole stats block reads as a single rhythmic strip on the atmosphere, not a stack of stacked cards.

**Continue playing — glass.** Replace the existing paper card. New shape: a single full-width `<Glass>` strip with `radius="card"`, `blur="strong"`, containing two columns — left: pack title (Fraunces), opponent name or "Solo game" (Plus Jakarta small); right: a `Resume` accent button. Renders only when `activeGames[0]` exists. Hidden completely otherwise.

**TabBar.** The existing `$lib/ui/TabBar.svelte` primitive, on atmosphere, three items: `Friends` / `Packs` / `Games`. Default tab `Friends`. Active tab persists in URL hash (`/account#packs`) so reload keeps the tab and links can deep-link.

**Tab content — one paper card, list of rows.** Each row is `display: grid` with hairline `border-bottom`. The row content varies by tab. No nested papers. No paper-per-row.

### 3. Friends tab

Rows:
- Pending requests at the top, marked with a small accent dot before the username and a sentence-cased status text ("wants to play").
- Sent requests next, dimmed (`--ink-3`).
- Accepted friends below, full opacity.
- Each row: avatar dot, username (Fraunces), `wins/losses` in tabular numerals, action cluster on the right (Play / Accept / Cancel / Remove depending on status).

Empty state: italic Fraunces line "No friends yet —" followed by a quiet ghost link "find people →" to `/people`.

### 4. Packs tab — with inline expand

Rows show:
- Title (Fraunces, semibold-ish weight 500)
- 3-task preview (Plus Jakarta small, `--ink-2`), `+ 13 more` if pack > 3 tasks
- Action cluster: `Play`, `Solo`, heart toggle
- A subtle chevron at the far right indicating expandability

Click anywhere on the row that isn't a button → row expands. Expanded:
- All 16 tasks in a numbered two-column grid (one column on mobile)
- The 3-task preview is replaced by the full list
- Chevron rotates 180°
- Click again or click another row to collapse

Multiple rows can be expanded independently — local state held by the row component, not lifted to page.

Empty state: "No liked packs yet —" + ghost link "browse public packs →" to `/packs`.

### 5. Games tab

Rows:
- In-progress games first. Each row: opponent name (or `Solo`), pack title underneath, Resume button on right.
- Below the active list, a quiet ghost button "Show N past games". Clicking expands a section with finished games (each: pack title + outcome + Delete).

Empty state: "Your first board awaits."

### 6. Solo entry path (resolves issue #3)

After fix #1 lands, the header shows `play | packs | people` for logged-in users, restoring access to `/packs`. **In addition**, the Packs tab on `/account` carries a `Solo` button on every row, so users can start a solo game in two taps from the dashboard without opening `/packs` at all. This is the primary path; `/packs` Solo button stays for users discovering new packs.

### 7. Component decomposition

Rule of thumb: extract a feature component when the inline JSX inside `+page.svelte` would push the file past ~280 lines or duplicate the row pattern. Predicted breakouts:

- `bingo-app/src/lib/feature/AccountFriendRow.svelte` — one row, props `friend`, dispatches `play`, `accept`, `delete`, `remove`.
- `bingo-app/src/lib/feature/AccountPackRow.svelte` — one row, props `pack`, internal expand state, dispatches `play`, `solo`, `like`.
- `bingo-app/src/lib/feature/AccountGameRow.svelte` — one row, props `game`, dispatches `resume`, `delete`.
- `bingo-app/src/lib/feature/ContinuePlaying.svelte` — the glass strip.

If by the end the page sits comfortably under 280 lines without these breakouts, skip them — but the row components are obvious and reusable, so they will likely earn their keep.

## Critical files

**Create:**
- `bingo-app/src/lib/stores/auth.ts` — auth writable + helpers.
- `bingo-app/src/lib/feature/ContinuePlaying.svelte` — glass strip with featured game.
- `bingo-app/src/lib/feature/AccountFriendRow.svelte`
- `bingo-app/src/lib/feature/AccountPackRow.svelte`
- `bingo-app/src/lib/feature/AccountGameRow.svelte`

**Modify:**
- `bingo-app/src/lib/feature/Header.svelte` — replace local `authed` with `$Auth`. Drop the `readAuth` helper (lives in the store now).
- `bingo-app/src/routes/login/+page.svelte` — call `markAuthed()` before `goto('/account')`.
- `bingo-app/src/routes/register/+page.svelte` — same.
- `bingo-app/src/routes/account/+page.svelte` — full rewrite of the layout: greeting → stats → continue playing → tab bar → active tab body. Hash-driven tab persistence. Drop the per-section paper-each-item rendering.
- `bingo-app/src/lib/feature/StatsBoard.svelte` — drop the paper wrappers around `BarChart` and `TopPacksList`. Tiles row stays.
- `bingo-app/src/routes/+page.svelte` (the home) — read `$Auth` instead of inline cookie regex, so the home hero swaps CTAs the moment the user logs in. Small change but keeps the auth source of truth in one place.

**Reuse (do not duplicate):**
- `bingo-app/src/lib/ui/{Paper,Glass,TabBar,Stack,Cluster,Page,Button}.svelte` — existing primitives.
- `bingo-app/src/lib/feature/{StatsBoard,BarChart,TopPacksList,StatTile}.svelte` — existing stats components, only the wrapper paper inside `StatsBoard` changes.
- The friend / pack / game data flows from `accountStore` already; rows just bind to it.

## Verification

1. **Auth reactivity** — Log out (clear cookie via DevTools), reload, header shows `about | login`. Log in. Without page reload, header instantly switches to `play | packs | people | avatar`. Refresh — still authed (cookie source of truth).
2. **Tab persistence** — Open `/account`, switch to Packs tab, hash becomes `#packs`. Refresh page → same tab active. Open `/account#games` directly → Games tab active.
3. **Layer count** — On `/account` with one active game: open DevTools → Inspect → confirm exactly one element with the `paper` class (active tab body) and one with the `glass` class (Continue playing strip), plus the Header glass when scrolled past 64px.
4. **Pack expand** — On Packs tab, click any pack row. All 16 tasks appear; chevron rotates. Click again, collapses. Click a different row, both can be open simultaneously. Mobile (375px): tasks render in single column.
5. **Solo entry** — From `/account`, switch to Packs tab, click Solo on any liked pack → redirected to `/game?solo=true` and a solo game starts.
6. **Empty states** — Fresh account (no friends, no liked packs, no games): each tab body shows the italic Fraunces empty line; tasks like "find people →" are clickable and route correctly.
7. **Stats still work** — StatsBoard renders correctly without its paper wrappers; bar tracks and rank tracks still read clearly on the atmosphere; mobile responsive layout (375px) still sound.
8. **Build & type-check** — `cd bingo-app && npm run check` reports 0 errors; `npm run build` succeeds.

## Open decisions to confirm in review

1. **Tab default.** Currently `Friends`. Could also be "last visited" (persisted in localStorage) but URL hash already lets users deep-link, and "last visited" surprises returning users with whichever tab they happened to be in. Default `Friends` keeps it predictable.
2. **Past games location.** The "Show N past" toggle inside Games tab keeps everything in one tab. The alternative — a dedicated `/games` route — wasn't requested and adds nav weight; staying in-tab.
3. **Auth-store cookie sync.** The store is read once at module load. If a user logs in in another tab, this tab won't update until reload. Acceptable for v1; a `BroadcastChannel` listener could come later if it matters.
4. **ContinuePlaying for solo.** When the featured game is a solo game (kind `solo`), the strip says "Solo game · {pack title}" with the same Resume button. No special-casing needed.

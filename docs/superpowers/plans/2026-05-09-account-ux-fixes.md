# `/account` UX Fixes Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Replace the long-scroll paper-card-quilt `/account` with a tabbed dashboard (one paper card visible at a time, optional glass strip for the active game), make the header reactive to login/logout, and expose pack expansion + solo entry from the dashboard itself.

**Architecture:** Frontend-only. New `Auth` writable in `$lib/stores/auth.ts` is the single source of truth for login state, kept in sync with the JWT cookie. `/account` becomes a TabBar-driven page rendering one of three row-component lists (Friends / Packs / Games) inside a single `<Paper>`. `<ContinuePlaying>` is a thin glass strip that appears only when an active game exists. `<StatsBoard>` drops its inner paper wrappers so its charts and top-packs list float on the atmosphere alongside the tile row.

**Tech Stack:** SvelteKit 1.14 · Svelte 3.53 · TypeScript · existing Aurora Calm tokens. No new dependencies.

---

## Verification Strategy

The project has no test framework (out of scope per the original Aurora Calm spec). Each task verifies via:

- **Type-check:** `cd bingo-app && npm run check` — must report 0 errors.
- **Build:** `cd bingo-app && npm run build` — must succeed.
- **Manual visual:** `cd bingo-app && npm run dev`, open http://localhost:3000, exercise the change in the browser at desktop and mobile (375×844 in DevTools).

For pages that fetch data (`/account`), the local stack must be running: `make compose` from the repo root.

## File Structure

```
bingo-app/src/lib/stores/
└── auth.ts                                  # NEW: writable auth signal + helpers

bingo-app/src/lib/feature/
├── ContinuePlaying.svelte                   # NEW: glass strip with featured game
├── AccountFriendRow.svelte                  # NEW: one friend row
├── AccountPackRow.svelte                    # NEW: one pack row with inline expand
├── AccountGameRow.svelte                    # NEW: one game row
├── Header.svelte                            # MODIFY: subscribe to $Auth
└── StatsBoard.svelte                        # MODIFY: drop inner <Paper> wrappers

bingo-app/src/routes/
├── +page.svelte                             # MODIFY: read $Auth instead of cookie regex
├── login/+page.svelte                       # MODIFY: markAuthed() on success
├── register/+page.svelte                    # MODIFY: markAuthed() on success
└── account/+page.svelte                     # MODIFY: full rewrite with tabs + rows
```

Untouched: `accountStore.ts`, `temporary.ts`, every existing `$lib/ui/*` primitive, all backend code, `/people` and `/packs` routes.

---

## Phase 1 — Auth reactivity

After this phase the header swaps the moment a user logs in or registers, and the home page CTA flips from `Start a game / log in` to `Continue playing / browse packs` without a page reload.

### Task 1.1: Create the auth store

**Files:**
- Create: `bingo-app/src/lib/stores/auth.ts`

- [ ] **Step 1: Write `bingo-app/src/lib/stores/auth.ts`**

```ts
import { writable } from 'svelte/store';
import { browser } from '$app/environment';

/**
 * Auth — the single source of truth for "is the user logged in".
 * Initialised from the JWT cookie on module load (so a refresh keeps the
 * session). Login/register flows call markAuthed() after a successful
 * response; any future logout helper should call markLoggedOut().
 */
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

export function markAuthed(): void {
    Auth.set(true);
}

export function markLoggedOut(): void {
    Auth.set(false);
}
```

- [ ] **Step 2: Type-check**

Run: `cd bingo-app && npm run check`
Expected: 0 errors.

- [ ] **Step 3: Commit**

```bash
git add bingo-app/src/lib/stores/auth.ts
git commit -m "feat(ui): add reactive Auth store"
```

---

### Task 1.2: Wire Header to `$Auth`

**Files:**
- Modify: `bingo-app/src/lib/feature/Header.svelte`

- [ ] **Step 1: Replace the script block of `bingo-app/src/lib/feature/Header.svelte`**

Open the file. Replace the entire `<script lang="ts">…</script>` block with:

```svelte
<script lang="ts">
    import { onDestroy, onMount } from 'svelte';
    import { browser } from '$app/environment';
    import { page } from '$app/stores';
    import { Auth } from '$lib/stores/auth';

    let scrolled = false;

    function onScroll() {
        scrolled = window.scrollY > 64;
    }

    onMount(() => {
        if (!browser) return;
        onScroll();
        window.addEventListener('scroll', onScroll, { passive: true });
    });

    onDestroy(() => {
        if (!browser) return;
        window.removeEventListener('scroll', onScroll);
    });

    $: pathname = $page.url.pathname;
    $: isActive = (href: string) => pathname === href || (href !== '/' && pathname.startsWith(href));
    $: authed = $Auth;
</script>
```

The template body of `Header.svelte` already references `authed` for nav switching — no template changes required.

- [ ] **Step 2: Type-check**

Run: `cd bingo-app && npm run check`
Expected: 0 errors.

- [ ] **Step 3: Commit**

```bash
git add bingo-app/src/lib/feature/Header.svelte
git commit -m "feat(ui): subscribe Header to Auth store"
```

---

### Task 1.3: Mark authed in login, register, and read it on home

**Files:**
- Modify: `bingo-app/src/routes/login/+page.svelte`
- Modify: `bingo-app/src/routes/register/+page.svelte`
- Modify: `bingo-app/src/routes/+page.svelte`

- [ ] **Step 1: Edit `bingo-app/src/routes/login/+page.svelte`**

Find the line:

```ts
import { API_URL, WEB_URL } from '../temporary';
```

Add immediately after it:

```ts
import { markAuthed } from '$lib/stores/auth';
```

Then find the `if (res.ok)` branch inside `submit()`:

```ts
if (res.ok) {
    goto('/account');
}
```

Replace it with:

```ts
if (res.ok) {
    markAuthed();
    goto('/account');
}
```

- [ ] **Step 2: Edit `bingo-app/src/routes/register/+page.svelte`**

Find the line:

```ts
import { _Submit } from './+page';
```

Add immediately after it:

```ts
import { markAuthed } from '$lib/stores/auth';
```

Then find the success branch inside `submit()`:

```ts
if (status === 200) {
    goto('/account');
}
```

Replace it with:

```ts
if (status === 200) {
    markAuthed();
    goto('/account');
}
```

- [ ] **Step 3: Edit `bingo-app/src/routes/+page.svelte` (home)**

Replace this block:

```svelte
<script lang="ts">
    import { browser } from '$app/environment';
    import Page from '$lib/ui/Page.svelte';
    import Cluster from '$lib/ui/Cluster.svelte';
    import Button from '$lib/ui/Button.svelte';
    import HeroDemo from '$lib/feature/HeroDemo.svelte';

    let authed = false;
    if (browser) {
        authed = RegExp('auth=[^;]+').exec(document.cookie) !== null;
    }
</script>
```

With:

```svelte
<script lang="ts">
    import Page from '$lib/ui/Page.svelte';
    import Cluster from '$lib/ui/Cluster.svelte';
    import Button from '$lib/ui/Button.svelte';
    import HeroDemo from '$lib/feature/HeroDemo.svelte';
    import { Auth } from '$lib/stores/auth';

    $: authed = $Auth;
</script>
```

- [ ] **Step 4: Type-check + build**

Run: `cd bingo-app && npm run check && npm run build`
Expected: 0 errors, build succeeds.

- [ ] **Step 5: Manual verify**

Run: `cd bingo-app && npm run dev`
- Log out (DevTools → Application → Cookies → delete `auth`), reload `/`. Header shows `about | login`. Hero CTA reads `Start a game | log in`.
- Log in via `/login`. Without a manual reload: header swaps to `play | packs | people | avatar`. Navigate to `/`. Hero CTA now reads `Continue playing | browse packs`.

- [ ] **Step 6: Commit**

```bash
git add bingo-app/src/routes/login/+page.svelte bingo-app/src/routes/register/+page.svelte bingo-app/src/routes/+page.svelte
git commit -m "feat(ui): hook login/register/home into Auth store"
```

---

## Phase 2 — Row + atmospheric components

Each row component handles one item shape (friend / pack / game) and emits typed CustomEvents. Page-level handlers stay in `account/+page.svelte` and route data flows. Pack row owns its expand state locally.

### Task 2.1: ContinuePlaying glass strip

**Files:**
- Create: `bingo-app/src/lib/feature/ContinuePlaying.svelte`

- [ ] **Step 1: Write `bingo-app/src/lib/feature/ContinuePlaying.svelte`**

```svelte
<script lang="ts">
    import { createEventDispatcher } from 'svelte';
    import Glass from '$lib/ui/Glass.svelte';
    import Button from '$lib/ui/Button.svelte';
    import type { Game } from '../../routes/accountStore';

    /**
     * ContinuePlaying — single glass strip for the user's most recent active
     * game. Renders nothing if no game is supplied; the page is expected to
     * gate render with {#if featured}.
     */
    export let game: Game;
    export let opponentName: string = '';
    export let packTitle: string = '';

    const dispatch = createEventDispatcher<{ resume: string }>();

    function resume() {
        dispatch('resume', game.gameId);
    }
</script>

<Glass radius="card" blur="strong">
    <div class="row">
        <div class="meta">
            <span class="eyebrow">continue playing</span>
            <span class="title">
                {#if game.kind === 'solo'}
                    Solo game
                {:else}
                    vs. {opponentName || '—'}
                {/if}
            </span>
            <span class="sub">{packTitle || '—'}</span>
        </div>
        <Button
            variant="accent"
            size="md"
            href={game.kind === 'solo' ? '/game?solo=true' : '/game'}
            on:click={resume}
        >
            Resume
        </Button>
    </div>
</Glass>

<style>
    .row {
        display: flex;
        align-items: center;
        justify-content: space-between;
        gap: var(--space-4);
        padding: var(--space-3) var(--space-4);
    }

    .meta {
        display: flex;
        flex-direction: column;
        gap: 0.15rem;
        min-width: 0;
    }

    .eyebrow {
        font-size: 0.6rem;
        letter-spacing: 0.18em;
        text-transform: uppercase;
        color: var(--ink-3);
    }

    .title {
        font-family: var(--font-display);
        font-size: 1.05rem;
        color: var(--ink);
        letter-spacing: -0.01em;
    }

    .sub {
        font-size: 0.8rem;
        color: var(--ink-3);
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
    }

    @media (max-width: 480px) {
        .row {
            flex-direction: column;
            align-items: stretch;
            gap: var(--space-3);
        }
    }
</style>
```

- [ ] **Step 2: Type-check + commit**

```bash
cd bingo-app && npm run check
git add bingo-app/src/lib/feature/ContinuePlaying.svelte
git commit -m "feat(ui): add ContinuePlaying glass strip"
```

---

### Task 2.2: AccountFriendRow

**Files:**
- Create: `bingo-app/src/lib/feature/AccountFriendRow.svelte`

- [ ] **Step 1: Write `bingo-app/src/lib/feature/AccountFriendRow.svelte`**

```svelte
<script lang="ts">
    import { createEventDispatcher } from 'svelte';
    import Cluster from '$lib/ui/Cluster.svelte';
    import Button from '$lib/ui/Button.svelte';
    import type { Friend } from '../../routes/accountStore';

    /**
     * AccountFriendRow — one row inside the Friends tab list.
     * Status semantics (per backend):
     *   1 = sent (we requested)
     *   2 = incoming (they requested us)
     *   3 = accepted
     */
    export let friend: Friend;

    const dispatch = createEventDispatcher<{
        play: string;
        accept: string;
        remove: string;
    }>();

    $: incoming = friend.status === 2;
    $: sent = friend.status === 1;
    $: accepted = friend.status === 3;
</script>

<div class="row" class:incoming class:sent>
    <div class="left">
        {#if incoming}<span class="dot" aria-hidden="true"></span>{/if}
        <span class="avatar" aria-hidden="true"></span>
        <div class="who">
            <span class="name">{friend.username}</span>
            {#if accepted}
                <span class="meta">{friend.wins}/{friend.loses}</span>
            {:else if incoming}
                <span class="meta">wants to play</span>
            {:else if sent}
                <span class="meta">request sent</span>
            {/if}
        </div>
    </div>

    <Cluster gap="s" align="center" justify="end">
        {#if accepted}
            <Button variant="accent" size="sm" on:click={() => dispatch('play', friend.userID)}>Play</Button>
            <Button variant="ghost" size="sm" on:click={() => dispatch('remove', friend.userID)}>Remove</Button>
        {:else if incoming}
            <Button variant="accent" size="sm" on:click={() => dispatch('accept', friend.userID)}>Accept</Button>
            <Button variant="ghost" size="sm" on:click={() => dispatch('remove', friend.userID)}>Decline</Button>
        {:else if sent}
            <Button variant="ghost" size="sm" disabled>Sent</Button>
            <Button variant="ghost" size="sm" on:click={() => dispatch('remove', friend.userID)}>Cancel</Button>
        {/if}
    </Cluster>
</div>

<style>
    .row {
        display: flex;
        align-items: center;
        justify-content: space-between;
        gap: var(--space-3);
        padding: var(--space-3) 0;
        border-bottom: 1px solid var(--hairline);
    }
    .row:last-child { border-bottom: none; }
    .row.sent { opacity: 0.7; }

    .left {
        display: flex;
        align-items: center;
        gap: var(--space-3);
        min-width: 0;
    }

    .dot {
        width: 6px; height: 6px;
        border-radius: var(--radius-pill);
        background: var(--accent);
        flex-shrink: 0;
    }

    .avatar {
        width: 28px; height: 28px;
        border-radius: var(--radius-pill);
        background: var(--accent);
        opacity: 0.85;
        flex-shrink: 0;
    }

    .who {
        display: flex;
        flex-direction: column;
        gap: 0.1rem;
        min-width: 0;
    }

    .name {
        font-family: var(--font-display);
        font-weight: 500;
        color: var(--ink);
        font-size: 0.95rem;
    }

    .meta {
        font-size: 0.75rem;
        color: var(--ink-3);
        font-variant-numeric: tabular-nums;
    }
</style>
```

- [ ] **Step 2: Type-check + commit**

```bash
cd bingo-app && npm run check
git add bingo-app/src/lib/feature/AccountFriendRow.svelte
git commit -m "feat(ui): add AccountFriendRow"
```

---

### Task 2.3: AccountPackRow with inline expand

**Files:**
- Create: `bingo-app/src/lib/feature/AccountPackRow.svelte`

- [ ] **Step 1: Write `bingo-app/src/lib/feature/AccountPackRow.svelte`**

```svelte
<script lang="ts">
    import { createEventDispatcher } from 'svelte';
    import Cluster from '$lib/ui/Cluster.svelte';
    import Button from '$lib/ui/Button.svelte';
    import type { TaskPack } from '../../routes/accountStore';

    /**
     * AccountPackRow — one row inside the Packs tab list. Click anywhere on
     * the row that isn't a button toggles full-list expansion (all 16 tasks).
     * Buttons stop propagation so clicks on Play/Solo/heart never trigger
     * expand/collapse.
     */
    export let pack: TaskPack;

    const dispatch = createEventDispatcher<{
        play: string;
        solo: string;
        unlike: string;
    }>();

    let expanded = false;

    function toggle() {
        expanded = !expanded;
    }

    function key(e: KeyboardEvent) {
        if (e.key === 'Enter' || e.key === ' ') {
            e.preventDefault();
            toggle();
        }
    }
</script>

<div
    class="row"
    class:expanded
    role="button"
    tabindex="0"
    aria-expanded={expanded}
    on:click={toggle}
    on:keydown={key}
>
    <div class="head">
        <div class="title-block">
            <span class="title">{pack.pack.title}</span>
            {#if pack.isPrivate}<span class="badge">private</span>{/if}
        </div>
        <span class="chevron" class:rot={expanded} aria-hidden="true">
            <svg viewBox="0 0 20 20" width="14" height="14"><path d="M5 8l5 5 5-5" fill="none" stroke="currentColor" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round"/></svg>
        </span>
    </div>

    {#if expanded}
        <ol class="tasks-full">
            {#each pack.pack.tasks as task, i}
                <li><span class="num">{i + 1}</span><span>{task}</span></li>
            {/each}
        </ol>
    {:else}
        <ul class="tasks-preview">
            {#each pack.pack.tasks.slice(0, 3) as task, i}
                <li><span class="num">{i + 1}</span><span>{task}</span></li>
            {/each}
            {#if pack.pack.tasks.length > 3}
                <li class="more">+ {pack.pack.tasks.length - 3} more</li>
            {/if}
        </ul>
    {/if}

    <div class="actions" on:click|stopPropagation on:keydown|stopPropagation role="presentation">
        <Cluster gap="s">
            <Button variant="accent" size="sm" on:click={() => dispatch('play', pack.id)}>Play</Button>
            <Button variant="ghost" size="sm" on:click={() => dispatch('solo', pack.id)}>Solo</Button>
            <Button variant="ghost" size="sm" on:click={() => dispatch('unlike', pack.id)} aria-label="Unlike pack">
                <span aria-hidden="true">♥</span>
            </Button>
        </Cluster>
    </div>
</div>

<style>
    .row {
        display: flex;
        flex-direction: column;
        gap: var(--space-2);
        padding: var(--space-3) 0;
        border-bottom: 1px solid var(--hairline);
        cursor: pointer;
        outline: none;
    }
    .row:last-child { border-bottom: none; }
    .row:focus-visible {
        outline: 2px solid var(--accent-ring);
        outline-offset: 2px;
        border-radius: var(--radius-sm);
    }

    .head {
        display: flex;
        align-items: center;
        justify-content: space-between;
        gap: var(--space-3);
    }

    .title-block {
        display: flex;
        align-items: baseline;
        gap: var(--space-2);
        min-width: 0;
    }

    .title {
        font-family: var(--font-display);
        font-size: 1.05rem;
        color: var(--ink);
        font-weight: 500;
        letter-spacing: -0.01em;
    }

    .badge {
        font-family: var(--font-body);
        font-size: 0.6rem;
        padding: 0.1rem 0.5rem;
        border-radius: var(--radius-pill);
        background: var(--paper-soft);
        color: var(--ink-2);
        border: 1px solid var(--hairline);
    }

    .chevron {
        color: var(--ink-3);
        transition: transform var(--dur-base) var(--ease);
        line-height: 0;
    }
    .chevron.rot { transform: rotate(180deg); }

    .tasks-preview, .tasks-full {
        list-style: none;
        margin: 0;
        padding: 0;
        display: grid;
        gap: var(--space-1) var(--space-4);
        font-size: 0.85rem;
        color: var(--ink-2);
        line-height: 1.5;
    }
    .tasks-preview { grid-template-columns: 1fr; }
    .tasks-full {
        grid-template-columns: 1fr 1fr;
    }
    @media (max-width: 640px) {
        .tasks-full { grid-template-columns: 1fr; }
    }

    .tasks-preview li, .tasks-full li {
        display: flex;
        gap: var(--space-2);
    }
    .num {
        color: var(--ink-4);
        min-width: 1.5em;
        font-variant-numeric: tabular-nums;
    }
    .more {
        color: var(--ink-3);
        font-style: italic;
        padding-left: calc(1.5em + var(--space-2));
    }

    .actions {
        display: flex;
        justify-content: flex-end;
        margin-top: var(--space-1);
    }
</style>
```

- [ ] **Step 2: Type-check + commit**

```bash
cd bingo-app && npm run check
git add bingo-app/src/lib/feature/AccountPackRow.svelte
git commit -m "feat(ui): add AccountPackRow with inline expand"
```

---

### Task 2.4: AccountGameRow

**Files:**
- Create: `bingo-app/src/lib/feature/AccountGameRow.svelte`

- [ ] **Step 1: Write `bingo-app/src/lib/feature/AccountGameRow.svelte`**

```svelte
<script lang="ts">
    import { createEventDispatcher } from 'svelte';
    import Cluster from '$lib/ui/Cluster.svelte';
    import Button from '$lib/ui/Button.svelte';
    import type { Game } from '../../routes/accountStore';

    /**
     * AccountGameRow — one row inside the Games tab list. Used for both
     * active games (resume) and finished games (delete only). The page
     * controls which actions render via the `past` prop.
     */
    export let game: Game;
    export let label: string;
    export let past = false;

    const dispatch = createEventDispatcher<{
        resume: string;
        delete: string;
    }>();
</script>

<div class="row" class:past>
    <div class="meta">
        <span class="title">{label}</span>
        {#if past}
            <span class="sub">finished</span>
        {/if}
    </div>

    <Cluster gap="s">
        {#if !past}
            <Button
                variant="accent"
                size="sm"
                href={game.kind === 'solo' ? '/game?solo=true' : '/game'}
                on:click={() => dispatch('resume', game.gameId)}
            >
                Resume
            </Button>
        {/if}
        <Button variant="ghost" size="sm" on:click={() => dispatch('delete', game.gameId)}>Delete</Button>
    </Cluster>
</div>

<style>
    .row {
        display: flex;
        align-items: center;
        justify-content: space-between;
        gap: var(--space-3);
        padding: var(--space-3) 0;
        border-bottom: 1px solid var(--hairline);
    }
    .row:last-child { border-bottom: none; }
    .row.past .title { color: var(--ink-3); }

    .meta {
        display: flex;
        flex-direction: column;
        gap: 0.1rem;
        min-width: 0;
    }

    .title {
        font-family: var(--font-display);
        font-weight: 500;
        font-size: 0.95rem;
        color: var(--ink);
    }

    .sub {
        font-size: 0.7rem;
        color: var(--ink-3);
        text-transform: uppercase;
        letter-spacing: 0.16em;
    }
</style>
```

- [ ] **Step 2: Type-check + commit**

```bash
cd bingo-app && npm run check
git add bingo-app/src/lib/feature/AccountGameRow.svelte
git commit -m "feat(ui): add AccountGameRow"
```

---

## Phase 3 — `/account` rewrite

### Task 3.1: Rewrite `/account/+page.svelte` with tabs and rows

**Files:**
- Modify: `bingo-app/src/routes/account/+page.svelte` (replace entirely)

- [ ] **Step 1: Replace `bingo-app/src/routes/account/+page.svelte` with the following**

```svelte
<script lang="ts">
    import { onMount } from 'svelte';
    import { browser } from '$app/environment';
    import Account from '../accountStore';
    import Page from '$lib/ui/Page.svelte';
    import Stack from '$lib/ui/Stack.svelte';
    import Cluster from '$lib/ui/Cluster.svelte';
    import Paper from '$lib/ui/Paper.svelte';
    import Button from '$lib/ui/Button.svelte';
    import TabBar from '$lib/ui/TabBar.svelte';
    import GameModal from '$lib/feature/GameModal.svelte';
    import StatsBoard from '$lib/feature/StatsBoard.svelte';
    import ContinuePlaying from '$lib/feature/ContinuePlaying.svelte';
    import AccountFriendRow from '$lib/feature/AccountFriendRow.svelte';
    import AccountPackRow from '$lib/feature/AccountPackRow.svelte';
    import AccountGameRow from '$lib/feature/AccountGameRow.svelte';
    import { _LikePack, _DeleteGame, _GetGame } from './+page';
    import { DeleteFriend, AcceptFriend } from '../friendRequests';
    import { _StartSolo } from '../packs/+page';

    type TabKey = 'friends' | 'packs' | 'games';

    let tab: TabKey = 'friends';
    let showModal = false;
    let selectedPackID = '';
    let selectedFriendID = '';
    let showPastGames = false;

    function readHashTab(): TabKey {
        if (!browser) return 'friends';
        const h = window.location.hash.slice(1);
        if (h === 'packs' || h === 'games' || h === 'friends') return h;
        return 'friends';
    }

    function writeHashTab(next: TabKey) {
        if (!browser) return;
        history.replaceState(null, '', `#${next}`);
    }

    onMount(() => {
        tab = readHashTab();
    });

    $: if (browser) writeHashTab(tab);

    function getOpponentUsername(userId: string) {
        for (const f of $Account?.friends ?? []) if (f.userID === userId) return f.username;
        return '—';
    }
    function getPackTitle(packID: string) {
        const liked = $Account?.likedPacks.find((p) => p.id === packID);
        if (liked) return liked.pack.title;
        const owned = $Account?.packs?.find((p) => p.id === packID);
        return owned?.pack.title ?? '—';
    }

    $: friendsAll = (() => {
        const all = $Account?.friends ?? [];
        const byOrder = (s: number) => (s === 2 ? 0 : s === 1 ? 1 : 2);
        return [...all].sort((a, b) => byOrder(a.status) - byOrder(b.status));
    })();
    $: packsAll = $Account?.likedPacks ?? [];
    $: activeGames = ($Account?.games ?? []).filter((g) => g.status !== 3);
    $: pastGames = ($Account?.games ?? []).filter((g) => g.status === 3);
    $: featuredGame = activeGames[0];

    function openWithPack(packID: string) { selectedPackID = packID; selectedFriendID = ''; showModal = true; }
    function openWithFriend(friendID: string) { selectedFriendID = friendID; selectedPackID = ''; showModal = true; }

    function gameLabel(game: { kind: string; user1Id: string; user2Id: string; packId: string }): string {
        const pack = getPackTitle(game.packId);
        if (game.kind === 'solo') return `Solo · ${pack}`;
        const me = $Account?.userID;
        const opp = me === game.user1Id ? game.user2Id : game.user1Id;
        return `vs. ${getOpponentUsername(opp)} · ${pack}`;
    }
</script>

<svelte:head><title>{$Account?.username ?? 'Account'} · taskbingo</title></svelte:head>

<Page width="normal">
    <Stack gap="xl">
        {#if $Account}
            <header class="greet">
                <h1 class="name">{$Account.username}</h1>
                <Cluster gap="m" align="baseline">
                    <span class="score">{$Account.bingo}</span>
                    <span class="score-label">total bingos</span>
                    {#if $Account.soloBingo}
                        <span class="solo">solo: {$Account.soloBingo}</span>
                    {/if}
                </Cluster>
                <p class="city">{$Account.city}</p>
            </header>

            <StatsBoard />

            {#if featuredGame}
                <ContinuePlaying
                    game={featuredGame}
                    opponentName={featuredGame.kind === 'solo'
                        ? ''
                        : getOpponentUsername($Account.userID === featuredGame.user1Id ? featuredGame.user2Id : featuredGame.user1Id)}
                    packTitle={getPackTitle(featuredGame.packId)}
                    on:resume={(e) => _GetGame(e.detail)}
                />
            {/if}

            <section class="tabs-section">
                <Cluster gap="m" align="center" justify="between">
                    <TabBar
                        bind:value={tab}
                        items={[
                            { value: 'friends', label: 'Friends' },
                            { value: 'packs', label: 'Packs' },
                            { value: 'games', label: 'Games' },
                        ]}
                        ariaLabel="Account sections"
                    />
                    {#if tab === 'packs'}
                        <Button variant="ghost" size="sm" href="/newpack">New pack</Button>
                    {:else if tab === 'games'}
                        <Button variant="ghost" size="sm" on:click={() => (showModal = true)}>New game</Button>
                    {/if}
                </Cluster>

                <Paper padding="md">
                    {#if tab === 'friends'}
                        {#if friendsAll.length === 0}
                            <p class="empty">No friends yet — <a href="/people">find people →</a></p>
                        {:else}
                            {#each friendsAll as friend (friend.userID)}
                                <AccountFriendRow
                                    {friend}
                                    on:play={(e) => openWithFriend(e.detail)}
                                    on:accept={(e) => AcceptFriend(e.detail)}
                                    on:remove={(e) => DeleteFriend(e.detail)}
                                />
                            {/each}
                        {/if}
                    {:else if tab === 'packs'}
                        {#if packsAll.length === 0}
                            <p class="empty">No liked packs yet — <a href="/packs">browse public packs →</a></p>
                        {:else}
                            {#each packsAll as pack (pack.id)}
                                <AccountPackRow
                                    {pack}
                                    on:play={(e) => openWithPack(e.detail)}
                                    on:solo={(e) => _StartSolo(e.detail)}
                                    on:unlike={(e) => _LikePack({ id: e.detail }, true)}
                                />
                            {/each}
                        {/if}
                    {:else}
                        {#if activeGames.length === 0 && pastGames.length === 0}
                            <p class="empty">Your first board awaits.</p>
                        {:else}
                            {#each activeGames as game (game.gameId)}
                                <AccountGameRow
                                    {game}
                                    label={gameLabel(game)}
                                    on:resume={(e) => _GetGame(e.detail)}
                                    on:delete={(e) => _DeleteGame(e.detail)}
                                />
                            {/each}
                            {#if pastGames.length > 0}
                                <button class="see-all" on:click={() => (showPastGames = !showPastGames)}>
                                    {showPastGames ? 'Hide past games' : `Show ${pastGames.length} past games`}
                                </button>
                                {#if showPastGames}
                                    {#each pastGames as game (game.gameId)}
                                        <AccountGameRow
                                            {game}
                                            label={gameLabel(game)}
                                            past
                                            on:delete={(e) => _DeleteGame(e.detail)}
                                        />
                                    {/each}
                                {/if}
                            {/if}
                        {/if}
                    {/if}
                </Paper>
            </section>
        {/if}
    </Stack>

    <GameModal bind:showModal {selectedFriendID} {selectedPackID} />
</Page>

<style>
    .greet { display: flex; flex-direction: column; gap: var(--space-2); }
    .name {
        font-family: var(--font-display);
        font-size: 2rem;
        font-weight: 400;
        letter-spacing: -0.02em;
    }
    .score {
        font-family: var(--font-display);
        font-size: 2rem;
        font-weight: 400;
        color: var(--ink);
        line-height: 1;
    }
    .score-label {
        font-size: 0.65rem;
        text-transform: uppercase;
        letter-spacing: 0.16em;
        color: var(--ink-3);
    }
    .solo {
        font-family: var(--font-display);
        font-style: italic;
        font-size: 0.95rem;
        color: var(--ink-3);
    }
    .city { color: var(--ink-3); font-size: 0.85rem; }

    .tabs-section {
        display: flex;
        flex-direction: column;
        gap: var(--space-3);
    }

    .empty {
        font-family: var(--font-display);
        font-style: italic;
        color: var(--ink-3);
        text-align: center;
        margin: var(--space-3) 0;
        font-size: 0.95rem;
    }
    .empty a {
        color: var(--ink);
        border-bottom: 1px solid rgba(40, 55, 95, 0.25);
        text-decoration: none;
    }
    .empty a:hover {
        color: var(--accent-from);
        border-bottom-color: var(--accent-from);
    }

    .see-all {
        background: transparent;
        border: none;
        color: var(--ink-2);
        font-size: 0.78rem;
        cursor: pointer;
        padding: var(--space-3) 0;
        align-self: flex-start;
        text-decoration: underline;
        text-underline-offset: 4px;
        font-family: var(--font-body);
    }
    .see-all:hover { color: var(--ink); }
</style>
```

- [ ] **Step 2: Type-check + build**

Run: `cd bingo-app && npm run check && npm run build`
Expected: 0 errors, build succeeds.

- [ ] **Step 3: Manual verify**

Run: `cd bingo-app && npm run dev` (with `make compose` running for the backend).

- Log in, navigate to `/account`. The page renders: greeting → StatsBoard → (if active game) ContinuePlaying glass strip → TabBar → one Paper containing Friends rows.
- Click `Packs` tab, URL hash becomes `#packs`. The same Paper now contains pack rows.
- Click any pack row (not on a button) → expands to show all 16 tasks; chevron rotates. Click again → collapses.
- Click `Solo` button on a pack row → solo game starts and redirects to `/game?solo=true`.
- Click `Games` tab → games list. If past games exist, click "Show N past games" → list expands.
- Reload `/account#packs` directly → opens on Packs tab.
- DevTools mobile preset (375px) → sections stack, pack expanded view becomes single column.
- DevTools Inspect: confirm that exactly one `.paper` element is on the page (active tab body) at any time (plus the StatTile internals which are paper-soft only, not the `<Paper>` component).

- [ ] **Step 4: Commit**

```bash
git add bingo-app/src/routes/account/+page.svelte
git commit -m "feat(ui): tabbed /account dashboard with one paper card and row components"
```

---

## Phase 4 — Strip the StatsBoard wrappers

### Task 4.1: Drop inner `<Paper>` from StatsBoard

**Files:**
- Modify: `bingo-app/src/lib/feature/StatsBoard.svelte`

- [ ] **Step 1: Open `bingo-app/src/lib/feature/StatsBoard.svelte`**

Find the block:

```svelte
        <Paper padding="md">
            <Stack gap="l">
                <BarChart buckets={$Stats.buckets} metric="bingo" title="Bingos" />
            </Stack>
        </Paper>

        <div class="split">
            <Paper padding="md">
                <BarChart buckets={$Stats.buckets} metric="tasks" title="Tasks closed" />
            </Paper>
            <Paper padding="md">
                <TopPacksList title="Most played packs" items={$Stats.topPacks} />
            </Paper>
        </div>
```

Replace it with:

```svelte
        <BarChart buckets={$Stats.buckets} metric="bingo" title="Bingos" />

        <div class="split">
            <BarChart buckets={$Stats.buckets} metric="tasks" title="Tasks closed" />
            <TopPacksList title="Most played packs" items={$Stats.topPacks} />
        </div>
```

- [ ] **Step 2: Drop the now-unused `Paper` and `Stack` imports**

In the same file, near the top of the script block, find:

```ts
    import Paper from '$lib/ui/Paper.svelte';
    import Stack from '$lib/ui/Stack.svelte';
```

Remove the `Paper` import. Keep `Stack` only if it's still referenced elsewhere in the file. Open the file and search for `<Stack` and `<Paper` — if no other references, remove both imports.

- [ ] **Step 3: Type-check + build**

Run: `cd bingo-app && npm run check && npm run build`
Expected: 0 errors, build succeeds.

- [ ] **Step 4: Manual verify**

Run dev server. On `/account`:
- StatsBoard tile row still renders correctly.
- "Bingos" bar chart sits directly on the atmosphere — bars visible on their internal `--paper-soft` tracks.
- "Tasks closed" and "Most played packs" sit side-by-side, also on the atmosphere, no outer card.
- DevTools Inspect: confirm `<Paper>` count on `/account` is exactly **1** (the active tab body) when an active game exists, plus the ContinuePlaying glass.

- [ ] **Step 5: Commit**

```bash
git add bingo-app/src/lib/feature/StatsBoard.svelte
git commit -m "refactor(ui): drop StatsBoard inner paper wrappers"
```

---

## Self-Review Checklist

- [ ] **Spec coverage:**
    - § Auth state goes reactive → Tasks 1.1, 1.2, 1.3.
    - § `/account` becomes tabbed dashboard with minimal layering → Task 3.1.
    - § Friends tab → AccountFriendRow (Task 2.2) + page wiring (Task 3.1).
    - § Packs tab with inline expand → AccountPackRow (Task 2.3) + page wiring.
    - § Games tab → AccountGameRow (Task 2.4) + page wiring with past-games toggle.
    - § Solo entry path (Header + per-row Solo) → Header from Task 1.2 surfaces `play | packs | people`; AccountPackRow has a `Solo` button dispatching to `_StartSolo`.
    - § Component decomposition → Tasks 2.1–2.4 create the four predicted components; the `+page.svelte` rewrite (Task 3.1) is well under the 280-line threshold and uses them.
    - § StatsBoard wrapper-paper removal → Task 4.1.
    - § Pack expand semantics (click anywhere not a button, two rows can be open simultaneously, mobile single-column) → AccountPackRow handles row click + per-component local `expanded` state + responsive grid.
    - § Empty-state copy → Task 3.1 inline empty-state paragraphs for each tab.
    - § Hash-driven tab persistence → Task 3.1 `readHashTab` / `writeHashTab` + reactive sync.
    - § Auth-store cookie sync (open decision: read once at module load) → Task 1.1 reads cookie at module init only.
    - § Tab default `Friends` → Task 3.1 default value of the `tab` variable.
    - § ContinuePlaying covers solo via `kind === 'solo'` branch → Task 2.1.

- [ ] **No placeholders:** No "TBD", no "implement later", no "similar to Task N". Every step contains the exact code or command.

- [ ] **Type consistency:**
    - Event payloads on rows (`play | accept | remove | solo | unlike | resume | delete`) carry a `string` (the relevant userID / packId / gameId), and `+page.svelte` consumes `e.detail` as a string in every handler.
    - Tab values (`'friends' | 'packs' | 'games'`) match between `TabBar items`, `tab` variable type, and `readHashTab` return.
    - `AccountPackRow` consumes `TaskPack` from `accountStore.ts` — the same shape the `Account` store yields in `likedPacks`.
    - `AccountGameRow` and `ContinuePlaying` both consume `Game` from `accountStore.ts`. `ContinuePlaying` accesses only `kind` and `gameId`; safe.
    - `_StartSolo` import in `account/+page.svelte` is from `../packs/+page` — confirmed exported from `bingo-app/src/routes/packs/+page.ts`.

---

## Execution Handoff

Plan complete and saved to `docs/superpowers/plans/2026-05-09-account-ux-fixes.md`. Two execution options:

**1. Subagent-Driven (recommended)** — fresh subagent per task, review between tasks, fast iteration.

**2. Inline Execution** — execute tasks in this session using executing-plans, batch execution with checkpoints.

Which approach?

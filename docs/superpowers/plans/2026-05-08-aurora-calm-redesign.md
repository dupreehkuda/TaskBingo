# Aurora Calm Frontend Redesign Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Replace the entire SvelteKit frontend visual + UX with the "Aurora Calm" design system locked in `docs/superpowers/specs/2026-05-08-aurora-calm-redesign-design.md` — atmospheric pastel canvas, warm ivory paper for content, sparse frosted glass for chrome, custom primitives in place of `flowbite-svelte`, and structural rework of `/account`, `/newpack`, `/game`.

**Architecture:** SvelteKit-only changes. Backend (Go services + Postgres + WebSocket protocol) is untouched. New design system lives in `bingo-app/src/lib/ui/` (primitives) and `bingo-app/src/lib/feature/` (feature components replacing `bingo-app/src/components/`). Six sequential phases — each phase ends with the app running. Mobile-first, equal weight to desktop.

**Tech Stack:** SvelteKit 1.14 · Svelte 3.53 · TypeScript · Tailwind CSS 3.1 · PostCSS · Google Fonts (Fraunces + Plus Jakarta Sans). No new runtime dependencies.

---

## Verification Strategy

The project has **no test framework** (this is explicitly out of scope per the spec). Each task verifies via:

- **Type-check:** `cd bingo-app && npm run check` — must report 0 errors and 0 warnings.
- **Build:** `cd bingo-app && npm run build` — must complete without errors.
- **Manual visual:** `cd bingo-app && npm run dev`, then open http://localhost:5173 (or the port Vite prints) and confirm the change in a real browser. Test desktop ≥1024px AND mobile <768px (use browser devtools device emulation).
- **Lint:** `cd bingo-app && npm run lint` — must pass before each commit.

Backend must be running for any page that fetches data. From repo root: `make compose` (then `make compose-down` to tear down). Static pages (`/`, `/about`, `/login` form, `/register` form) work without the backend.

If a step says "verify in browser", spend 30 seconds doing it. Don't claim a visual change is done from code-reading alone.

## File Structure

New directories:

```
bingo-app/src/lib/                # SvelteKit's $lib alias auto-resolves here
├── tokens.css                    # CSS custom properties — single source of truth
├── ui/                           # Design system primitives
│   ├── Atmosphere.svelte         # Body-level mesh + noise
│   ├── Paper.svelte              # Solid ivory content surface
│   ├── Glass.svelte              # Frosted chrome surface
│   ├── Page.svelte               # Page-level wrapper (max-width, padding)
│   ├── Stack.svelte              # Vertical flex with gap token
│   ├── Cluster.svelte            # Horizontal flex-wrap with gap token
│   ├── Button.svelte             # accent | ghost | quiet variants
│   ├── Input.svelte              # Floating-label text input
│   ├── Textarea.svelte           # Floating-label textarea
│   ├── Toggle.svelte             # Pill switch
│   ├── Select.svelte             # Paper-styled native select
│   ├── TabBar.svelte             # Tabs / segmented control (responsive)
│   └── index.ts                  # Barrel export
└── feature/                      # Feature components replacing src/components/
    ├── Header.svelte             # Brand + nav, glass-on-scroll
    ├── BingoGrid.svelte          # 4×4 grid with winning-row glow
    ├── Keypad.svelte             # Number pad
    ├── TaskList.svelte           # Tasks panel for /game
    ├── NotesJournal.svelte       # Comments CRUD
    ├── GameModal.svelte          # New game dialog
    └── TaskListEditor.svelte     # Stream editor for /newpack
```

Modified existing files:

```
bingo-app/src/app.html            # +Google Fonts <link>
bingo-app/src/app.postcss         # Replace global styles, import tokens.css
bingo-app/src/routes/+layout.svelte    # Wrap in Atmosphere, swap Header
bingo-app/src/routes/+page.svelte      # /
bingo-app/src/routes/about/+page.svelte
bingo-app/src/routes/login/+page.svelte
bingo-app/src/routes/register/+page.svelte
bingo-app/src/routes/account/+page.svelte
bingo-app/src/routes/people/+page.svelte
bingo-app/src/routes/packs/+page.svelte
bingo-app/src/routes/newpack/+page.svelte
bingo-app/src/routes/newpack/+page.ts  # Adjust _Submit to read from chip array
bingo-app/src/routes/game/+page.svelte
bingo-app/tailwind.config.cjs     # Drop flowbite content path & plugin
bingo-app/package.json            # Drop flowbite + flowbite-svelte
```

Deleted at end of Phase 6:

```
bingo-app/src/components/         # Entire dir — superseded by lib/feature/
```

The backend (`game-service/`, `user-data-service/`, `task-data-service/`, `api/`, `migrations/`, etc.) is **not touched** by this plan.

---

## Phase 1 — Foundation

The atmosphere appears on every page. Existing pages still render with flowbite — visual mismatch is expected for one phase. Each task in this phase ends in a working app.

### Task 1.1: Add design tokens

**Files:**
- Create: `bingo-app/src/lib/tokens.css`
- Modify: `bingo-app/src/app.postcss` (replace whole file)

- [ ] **Step 1: Create `bingo-app/src/lib/tokens.css`**

```css
:root {
    /* === Atmosphere base === */
    --atmos-base-from: #f4f6fa;
    --atmos-base-to: #efeaf1;
    --atmos-sky: #d3dfe8;
    --atmos-lilac: #e2d7e3;

    /* === Surfaces === */
    --paper: #fbfaf6;
    --paper-soft: #f1f0eb;
    --hairline: rgba(40, 50, 80, 0.06);

    /* === Accent (use as background-image, not background-color) === */
    --accent: linear-gradient(135deg, #94aacf 0%, #b8a3cd 100%);
    --accent-from: #94aacf;
    --accent-to: #b8a3cd;
    --accent-shadow: rgba(95, 110, 160, 0.4);
    --accent-ring: rgba(148, 170, 207, 0.6);

    /* === Text === */
    --ink: #1c2540;
    --ink-2: #4d5878;
    --ink-3: #7e88a8;
    --ink-4: #97a0bb;
    --on-accent: #ffffff;

    /* === Typography === */
    --font-display: 'Fraunces', 'DM Serif Display', Georgia, serif;
    --font-body: 'Plus Jakarta Sans', system-ui, -apple-system, sans-serif;

    /* === Spacing scale (4-pt) === */
    --space-1: 0.25rem;
    --space-2: 0.5rem;
    --space-3: 0.75rem;
    --space-4: 1rem;
    --space-5: 1.5rem;
    --space-6: 2rem;
    --space-7: 3rem;
    --space-8: 4rem;

    /* === Radii === */
    --radius-sm: 8px;
    --radius-md: 12px;
    --radius-lg: 18px;
    --radius-pill: 999px;

    /* === Motion === */
    --ease: cubic-bezier(0.4, 0, 0.2, 1);
    --dur-fast: 140ms;
    --dur-base: 220ms;
    --dur-slow: 380ms;
}

@media (prefers-reduced-motion: reduce) {
    :root {
        --dur-fast: 0ms;
        --dur-base: 0ms;
        --dur-slow: 0ms;
    }
}
```

- [ ] **Step 2: Replace `bingo-app/src/app.postcss` entirely**

```postcss
@tailwind base;
@tailwind components;
@tailwind utilities;

@import './lib/tokens.css';

html, body {
    background: linear-gradient(140deg, var(--atmos-base-from) 0%, var(--atmos-base-to) 100%);
    color: var(--ink);
    font-family: var(--font-body);
    font-size: 16px;
    line-height: 1.5;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
}

body {
    min-height: 100vh;
}

/* Reset Tailwind's default heading styling — we set type via primitives */
h1, h2, h3, h4, h5, h6 {
    font-family: var(--font-display);
    font-weight: 400;
    letter-spacing: -0.015em;
    color: var(--ink);
}

a {
    color: inherit;
}

button, input, textarea, select {
    font-family: inherit;
    font-size: inherit;
}

/* Skip link for keyboard users */
.skip-link {
    position: absolute;
    top: -100px;
    left: var(--space-4);
    padding: var(--space-2) var(--space-4);
    background: var(--paper);
    color: var(--ink);
    border-radius: var(--radius-pill);
    border: 1px solid var(--hairline);
    z-index: 100;
}
.skip-link:focus {
    top: var(--space-4);
}
```

- [ ] **Step 3: Type-check + build**

```bash
cd bingo-app
npm run check
npm run build
```

Expected: 0 errors. Build succeeds.

- [ ] **Step 4: Commit**

```bash
git add bingo-app/src/lib/tokens.css bingo-app/src/app.postcss
git commit -m "feat(ui): add design tokens and reset global styles"
```

---

### Task 1.2: Add Google Fonts to app.html

**Files:**
- Modify: `bingo-app/src/app.html` (add font preconnect + stylesheet inside `<head>`)

- [ ] **Step 1: Insert font links above `%sveltekit.head%`**

In `bingo-app/src/app.html`, add these three lines immediately before `%sveltekit.head%`:

```html
		<link rel="preconnect" href="https://fonts.googleapis.com">
		<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
		<link href="https://fonts.googleapis.com/css2?family=Fraunces:ital,opsz,wght@0,9..144,400;0,9..144,500;1,9..144,300;1,9..144,400&family=Plus+Jakarta+Sans:wght@400;500;600&display=swap" rel="stylesheet">
```

- [ ] **Step 2: Run dev server and inspect Network tab**

```bash
cd bingo-app
npm run dev
```

Open http://localhost:5173, open DevTools → Network → Font, refresh. Expected: `Fraunces[…].woff2` and `PlusJakartaSans[…].woff2` requests, both 200 status.

- [ ] **Step 3: Commit**

```bash
git add bingo-app/src/app.html
git commit -m "feat(ui): load Fraunces and Plus Jakarta Sans"
```

---

### Task 1.3: Atmosphere primitive

**Files:**
- Create: `bingo-app/src/lib/ui/Atmosphere.svelte`

- [ ] **Step 1: Write `bingo-app/src/lib/ui/Atmosphere.svelte`**

```svelte
<script lang="ts">
    /**
     * Atmosphere — paints the page canvas. Render once at the layout root.
     * It applies the mesh gradient + noise to <body> via :global on mount and
     * cleans up on destroy, so SvelteKit page navigation never flashes white.
     */
    import { onDestroy, onMount } from 'svelte';
    import { browser } from '$app/environment';

    onMount(() => {
        if (!browser) return;
        document.body.classList.add('atmos');
    });

    onDestroy(() => {
        if (!browser) return;
        document.body.classList.remove('atmos');
    });
</script>

<slot />

<style>
    :global(body.atmos) {
        position: relative;
        background:
            radial-gradient(58% 50% at 18% 22%, var(--atmos-sky) 0%, transparent 64%),
            radial-gradient(48% 48% at 84% 80%, var(--atmos-lilac) 0%, transparent 66%),
            linear-gradient(140deg, var(--atmos-base-from) 0%, var(--atmos-base-to) 100%);
        background-attachment: fixed;
    }

    /* Noise overlay sits in a fixed pseudo-element so it doesn't scroll. */
    :global(body.atmos)::before {
        content: '';
        position: fixed;
        inset: 0;
        pointer-events: none;
        z-index: 0;
        mix-blend-mode: overlay;
        opacity: 0.5;
        background-image: url("data:image/svg+xml;utf8,<svg xmlns='http://www.w3.org/2000/svg' width='160' height='160'><filter id='n'><feTurbulence type='fractalNoise' baseFrequency='0.85' numOctaves='2' stitchTiles='stitch'/><feColorMatrix values='0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.04 0'/></filter><rect width='100%25' height='100%25' filter='url(%23n)'/></svg>");
    }

    /* Page content sits above the noise overlay. */
    :global(body.atmos) > * {
        position: relative;
        z-index: 1;
    }
</style>
```

- [ ] **Step 2: Type-check**

```bash
cd bingo-app && npm run check
```

Expected: 0 errors.

- [ ] **Step 3: Commit**

```bash
git add bingo-app/src/lib/ui/Atmosphere.svelte
git commit -m "feat(ui): add Atmosphere primitive"
```

---

### Task 1.4: Paper primitive

**Files:**
- Create: `bingo-app/src/lib/ui/Paper.svelte`

- [ ] **Step 1: Write `bingo-app/src/lib/ui/Paper.svelte`**

```svelte
<script lang="ts">
    /**
     * Paper — solid ivory content surface. The workhorse.
     * Use for cards, list items, forms, modal bodies, the bingo grid container.
     * Never nest paper inside paper — separate via spacing instead.
     */
    type Padding = 'sm' | 'md' | 'lg';

    export let padding: Padding = 'md';
    export let radius: 'md' | 'lg' = 'lg';
</script>

<div class="paper paper--p-{padding} paper--r-{radius}" {...$$restProps}>
    <slot />
</div>

<style>
    .paper {
        background: var(--paper);
        border: 1px solid var(--hairline);
        color: var(--ink);
        box-shadow:
            0 1px 0 rgba(255, 255, 255, 0.6) inset,
            0 2px 6px -2px rgba(40, 55, 95, 0.05),
            0 14px 32px -16px rgba(40, 55, 95, 0.10);
    }

    .paper--p-sm { padding: var(--space-3) var(--space-4); }
    .paper--p-md { padding: var(--space-4) var(--space-5); }
    .paper--p-lg { padding: var(--space-5) var(--space-6); }

    .paper--r-md { border-radius: var(--radius-md); }
    .paper--r-lg { border-radius: var(--radius-lg); }
</style>
```

- [ ] **Step 2: Type-check**

```bash
cd bingo-app && npm run check
```

Expected: 0 errors.

- [ ] **Step 3: Commit**

```bash
git add bingo-app/src/lib/ui/Paper.svelte
git commit -m "feat(ui): add Paper primitive"
```

---

### Task 1.5: Glass primitive

**Files:**
- Create: `bingo-app/src/lib/ui/Glass.svelte`

- [ ] **Step 1: Write `bingo-app/src/lib/ui/Glass.svelte`**

```svelte
<script lang="ts">
    /**
     * Glass — frosted chrome surface. Rare. Use only for sticky headers
     * (when scrolled), the side panel during /game, and modal overlays.
     * Never nest glass inside glass; never put another glass inside this slot.
     */
    type Radius = 'pill' | 'card';
    type Blur = 'soft' | 'strong';

    export let radius: Radius = 'card';
    export let blur: Blur = 'strong';
</script>

<div class="glass glass--r-{radius} glass--b-{blur}" {...$$restProps}>
    <slot />
</div>

<style>
    .glass {
        background: rgba(252, 253, 255, 0.55);
        border: 1px solid rgba(255, 255, 255, 0.65);
        color: var(--ink);
        box-shadow:
            0 22px 44px -18px rgba(40, 55, 95, 0.16),
            inset 0 1px 0 rgba(255, 255, 255, 0.7);
    }

    .glass--r-pill { border-radius: var(--radius-pill); }
    .glass--r-card { border-radius: var(--radius-lg); }

    .glass--b-soft {
        backdrop-filter: blur(16px) saturate(125%);
        -webkit-backdrop-filter: blur(16px) saturate(125%);
    }
    .glass--b-strong {
        backdrop-filter: blur(28px) saturate(130%);
        -webkit-backdrop-filter: blur(28px) saturate(130%);
    }
</style>
```

- [ ] **Step 2: Type-check + commit**

```bash
cd bingo-app && npm run check
git add bingo-app/src/lib/ui/Glass.svelte
git commit -m "feat(ui): add Glass primitive"
```

---

### Task 1.6: Layout primitives (Stack, Cluster, Page)

**Files:**
- Create: `bingo-app/src/lib/ui/Stack.svelte`
- Create: `bingo-app/src/lib/ui/Cluster.svelte`
- Create: `bingo-app/src/lib/ui/Page.svelte`

- [ ] **Step 1: Write `bingo-app/src/lib/ui/Stack.svelte`**

```svelte
<script lang="ts">
    type Gap = 's' | 'm' | 'l' | 'xl';

    export let gap: Gap = 'm';
    export let align: 'start' | 'center' | 'stretch' = 'stretch';
</script>

<div class="stack stack--g-{gap} stack--a-{align}" {...$$restProps}>
    <slot />
</div>

<style>
    .stack { display: flex; flex-direction: column; }
    .stack--g-s  { gap: var(--space-2); }
    .stack--g-m  { gap: var(--space-4); }
    .stack--g-l  { gap: var(--space-5); }
    .stack--g-xl { gap: var(--space-6); }
    .stack--a-start  { align-items: flex-start; }
    .stack--a-center { align-items: center; }
    .stack--a-stretch { align-items: stretch; }
</style>
```

- [ ] **Step 2: Write `bingo-app/src/lib/ui/Cluster.svelte`**

```svelte
<script lang="ts">
    type Gap = 's' | 'm' | 'l';

    export let gap: Gap = 'm';
    export let align: 'start' | 'center' | 'baseline' = 'center';
    export let justify: 'start' | 'center' | 'end' | 'between' = 'start';
</script>

<div class="cluster cluster--g-{gap} cluster--a-{align} cluster--j-{justify}" {...$$restProps}>
    <slot />
</div>

<style>
    .cluster { display: flex; flex-wrap: wrap; }
    .cluster--g-s { gap: var(--space-2); }
    .cluster--g-m { gap: var(--space-3); }
    .cluster--g-l { gap: var(--space-5); }
    .cluster--a-start    { align-items: flex-start; }
    .cluster--a-center   { align-items: center; }
    .cluster--a-baseline { align-items: baseline; }
    .cluster--j-start   { justify-content: flex-start; }
    .cluster--j-center  { justify-content: center; }
    .cluster--j-end     { justify-content: flex-end; }
    .cluster--j-between { justify-content: space-between; }
</style>
```

- [ ] **Step 3: Write `bingo-app/src/lib/ui/Page.svelte`**

```svelte
<script lang="ts">
    /**
     * Page — outer wrapper for every route. Caps width and pads the gutters.
     * Use as the outermost element on each route's +page.svelte.
     */
    type Width = 'narrow' | 'normal' | 'wide';

    export let width: Width = 'normal';
</script>

<main class="page page--w-{width}" {...$$restProps}>
    <slot />
</main>

<style>
    .page {
        margin: 0 auto;
        padding: var(--space-6) var(--space-5) var(--space-7);
        width: 100%;
    }
    .page--w-narrow { max-width: 28rem; }
    .page--w-normal { max-width: 60rem; }
    .page--w-wide   { max-width: 80rem; }

    @media (max-width: 640px) {
        .page { padding: var(--space-5) var(--space-4) var(--space-6); }
    }
</style>
```

- [ ] **Step 4: Type-check + commit**

```bash
cd bingo-app && npm run check
git add bingo-app/src/lib/ui/Stack.svelte bingo-app/src/lib/ui/Cluster.svelte bingo-app/src/lib/ui/Page.svelte
git commit -m "feat(ui): add Stack, Cluster, Page layout primitives"
```

---

### Task 1.7: Button primitive

**Files:**
- Create: `bingo-app/src/lib/ui/Button.svelte`

- [ ] **Step 1: Write `bingo-app/src/lib/ui/Button.svelte`**

```svelte
<script lang="ts">
    type Variant = 'accent' | 'ghost' | 'quiet';
    type Size = 'sm' | 'md' | 'lg';

    export let variant: Variant = 'accent';
    export let size: Size = 'md';
    export let href: string | undefined = undefined;
    export let type: 'button' | 'submit' | 'reset' = 'button';
    export let disabled = false;
    export let fullWidth = false;
</script>

{#if href !== undefined}
    <a class="btn btn--{variant} btn--{size}" class:full={fullWidth} {href} role="button" {...$$restProps}>
        <slot />
    </a>
{:else}
    <button
        class="btn btn--{variant} btn--{size}"
        class:full={fullWidth}
        {type}
        {disabled}
        on:click
        on:mouseenter
        on:mouseleave
        {...$$restProps}
    >
        <slot />
    </button>
{/if}

<style>
    .btn {
        display: inline-flex;
        align-items: center;
        justify-content: center;
        gap: var(--space-2);
        border-radius: var(--radius-pill);
        font-family: var(--font-body);
        font-weight: 500;
        letter-spacing: 0.005em;
        cursor: pointer;
        transition: filter var(--dur-fast) var(--ease),
                    box-shadow var(--dur-fast) var(--ease),
                    transform var(--dur-fast) var(--ease);
        text-decoration: none;
        border: 1px solid transparent;
        white-space: nowrap;
    }

    .btn:focus-visible {
        outline: 2px solid var(--accent-ring);
        outline-offset: 2px;
    }

    .btn:disabled {
        opacity: 0.5;
        cursor: not-allowed;
    }

    .btn:not(:disabled):hover { filter: brightness(1.04); }
    .btn:not(:disabled):active { transform: translateY(1px); }

    .full { width: 100%; }

    /* Sizes */
    .btn--sm { padding: 0.45rem 0.95rem; font-size: 0.78rem; min-height: 32px; }
    .btn--md { padding: 0.62rem 1.2rem;  font-size: 0.85rem; min-height: 44px; }
    .btn--lg { padding: 0.78rem 1.5rem;  font-size: 0.95rem; min-height: 52px; }

    /* Variants */
    .btn--accent {
        background: var(--accent);
        color: var(--on-accent);
        border-color: rgba(255, 255, 255, 0.5);
        box-shadow: 0 8px 18px -6px var(--accent-shadow);
    }

    .btn--ghost {
        background: var(--paper);
        color: var(--ink);
        border-color: var(--hairline);
    }

    .btn--quiet {
        background: transparent;
        color: var(--ink);
        border: none;
        border-bottom: 1px solid rgba(40, 55, 95, 0.25);
        border-radius: 0;
        padding: 0 0 1px 0;
        min-height: auto;
    }
    .btn--quiet:not(:disabled):active { transform: none; }
</style>
```

- [ ] **Step 2: Type-check + commit**

```bash
cd bingo-app && npm run check
git add bingo-app/src/lib/ui/Button.svelte
git commit -m "feat(ui): add Button primitive (accent/ghost/quiet)"
```

---

### Task 1.8: Input primitive

**Files:**
- Create: `bingo-app/src/lib/ui/Input.svelte`

- [ ] **Step 1: Write `bingo-app/src/lib/ui/Input.svelte`**

```svelte
<script lang="ts">
    /**
     * Input — text input with floating label. Label rests inside the input
     * until focus or value, then slides up.
     */
    export let label: string;
    export let value: string = '';
    export let type: 'text' | 'email' | 'password' | 'search' | 'url' = 'text';
    export let name: string | undefined = undefined;
    export let id: string | undefined = undefined;
    export let placeholder: string = '';
    export let required = false;
    export let disabled = false;
    export let error: string | undefined = undefined;
    export let hint: string | undefined = undefined;

    let focused = false;
    $: filled = value !== '' && value !== undefined && value !== null;
    $: lifted = focused || filled;
    $: inputId = id ?? `input-${name ?? Math.random().toString(36).slice(2, 8)}`;

    function onInput(e: Event) {
        const t = e.currentTarget as HTMLInputElement;
        value = t.value;
    }
</script>

<label class="field" class:has-error={error}>
    <span class="label" class:lifted>{label}{required ? ' *' : ''}</span>
    <input
        id={inputId}
        {type}
        {name}
        {placeholder}
        {required}
        {disabled}
        {value}
        aria-invalid={error ? 'true' : 'false'}
        aria-describedby={error || hint ? `${inputId}-msg` : undefined}
        on:input={onInput}
        on:focus={() => (focused = true)}
        on:blur={() => (focused = false)}
        on:keydown
        on:keyup
        on:change
    />
    {#if error}
        <span id="{inputId}-msg" class="msg msg--error">{error}</span>
    {:else if hint}
        <span id="{inputId}-msg" class="msg">{hint}</span>
    {/if}
</label>

<style>
    .field {
        display: flex;
        flex-direction: column;
        position: relative;
    }

    .label {
        position: absolute;
        left: 1rem;
        top: 50%;
        transform: translateY(-50%);
        color: var(--ink-3);
        font-size: 0.95rem;
        pointer-events: none;
        transition: top var(--dur-fast) var(--ease),
                    transform var(--dur-fast) var(--ease),
                    font-size var(--dur-fast) var(--ease),
                    color var(--dur-fast) var(--ease);
    }

    .label.lifted {
        top: 0.4rem;
        transform: translateY(0);
        font-size: 0.7rem;
        color: var(--ink-3);
        letter-spacing: 0.04em;
        text-transform: uppercase;
    }

    input {
        background: var(--paper);
        border: 1px solid var(--hairline);
        border-radius: var(--radius-md);
        padding: 1.5rem 1rem 0.5rem;
        color: var(--ink);
        font-size: 0.95rem;
        min-height: 56px;
        transition: border-color var(--dur-fast) var(--ease),
                    box-shadow var(--dur-fast) var(--ease);
        outline: none;
    }

    input:focus {
        border-color: var(--accent-from);
        box-shadow: 0 0 0 3px var(--accent-ring);
    }

    .has-error input {
        border-color: #c44e4e;
    }
    .has-error input:focus {
        box-shadow: 0 0 0 3px rgba(196, 78, 78, 0.25);
    }

    .msg {
        font-size: 0.75rem;
        color: var(--ink-3);
        margin-top: 0.25rem;
        margin-left: 1rem;
    }
    .msg--error {
        color: #c44e4e;
    }

    input:disabled {
        opacity: 0.6;
        cursor: not-allowed;
    }
</style>
```

- [ ] **Step 2: Type-check + commit**

```bash
cd bingo-app && npm run check
git add bingo-app/src/lib/ui/Input.svelte
git commit -m "feat(ui): add Input primitive with floating label"
```

---

### Task 1.9: Textarea primitive

**Files:**
- Create: `bingo-app/src/lib/ui/Textarea.svelte`

- [ ] **Step 1: Write `bingo-app/src/lib/ui/Textarea.svelte`**

```svelte
<script lang="ts">
    export let label: string;
    export let value: string = '';
    export let name: string | undefined = undefined;
    export let id: string | undefined = undefined;
    export let placeholder: string = '';
    export let rows = 3;
    export let required = false;
    export let disabled = false;
    export let error: string | undefined = undefined;

    let focused = false;
    $: filled = value !== '' && value !== undefined && value !== null;
    $: lifted = focused || filled;
    $: areaId = id ?? `textarea-${name ?? Math.random().toString(36).slice(2, 8)}`;

    function onInput(e: Event) {
        const t = e.currentTarget as HTMLTextAreaElement;
        value = t.value;
    }
</script>

<label class="field" class:has-error={error}>
    <span class="label" class:lifted>{label}{required ? ' *' : ''}</span>
    <textarea
        id={areaId}
        {name}
        {rows}
        {placeholder}
        {required}
        {disabled}
        {value}
        aria-invalid={error ? 'true' : 'false'}
        on:input={onInput}
        on:focus={() => (focused = true)}
        on:blur={() => (focused = false)}
    ></textarea>
    {#if error}<span class="msg msg--error">{error}</span>{/if}
</label>

<style>
    .field { display: flex; flex-direction: column; position: relative; }

    .label {
        position: absolute;
        left: 1rem;
        top: 1rem;
        color: var(--ink-3);
        font-size: 0.95rem;
        pointer-events: none;
        transition: top var(--dur-fast) var(--ease),
                    font-size var(--dur-fast) var(--ease);
    }
    .label.lifted {
        top: 0.4rem;
        font-size: 0.7rem;
        letter-spacing: 0.04em;
        text-transform: uppercase;
    }

    textarea {
        background: var(--paper);
        border: 1px solid var(--hairline);
        border-radius: var(--radius-md);
        padding: 1.5rem 1rem 0.75rem;
        color: var(--ink);
        font-size: 0.95rem;
        font-family: inherit;
        line-height: 1.5;
        resize: vertical;
        min-height: 96px;
        outline: none;
        transition: border-color var(--dur-fast) var(--ease),
                    box-shadow var(--dur-fast) var(--ease);
    }
    textarea:focus {
        border-color: var(--accent-from);
        box-shadow: 0 0 0 3px var(--accent-ring);
    }

    .has-error textarea { border-color: #c44e4e; }

    .msg {
        font-size: 0.75rem;
        color: var(--ink-3);
        margin: 0.25rem 0 0 1rem;
    }
    .msg--error { color: #c44e4e; }
</style>
```

- [ ] **Step 2: Type-check + commit**

```bash
cd bingo-app && npm run check
git add bingo-app/src/lib/ui/Textarea.svelte
git commit -m "feat(ui): add Textarea primitive"
```

---

### Task 1.10: Toggle primitive

**Files:**
- Create: `bingo-app/src/lib/ui/Toggle.svelte`

- [ ] **Step 1: Write `bingo-app/src/lib/ui/Toggle.svelte`**

```svelte
<script lang="ts">
    /**
     * Toggle — pill switch. Uses native <input type="checkbox"> under
     * the hood so forms read it as a normal field.
     */
    export let label: string;
    export let checked = false;
    export let name: string | undefined = undefined;
    export let disabled = false;
</script>

<label class="toggle" class:disabled>
    <input type="checkbox" {name} bind:checked {disabled} on:change />
    <span class="track" aria-hidden="true">
        <span class="thumb"></span>
    </span>
    <span class="label-text">{label}</span>
</label>

<style>
    .toggle {
        display: inline-flex;
        align-items: center;
        gap: var(--space-3);
        cursor: pointer;
        user-select: none;
    }
    .toggle.disabled { cursor: not-allowed; opacity: 0.5; }

    input {
        position: absolute;
        opacity: 0;
        pointer-events: none;
    }

    .track {
        position: relative;
        width: 44px;
        height: 26px;
        border-radius: var(--radius-pill);
        background: var(--paper-soft);
        border: 1px solid var(--hairline);
        transition: background var(--dur-fast) var(--ease);
    }

    .thumb {
        position: absolute;
        left: 2px;
        top: 2px;
        width: 20px;
        height: 20px;
        border-radius: 999px;
        background: white;
        box-shadow: 0 1px 2px rgba(40, 55, 95, 0.18);
        transition: transform var(--dur-fast) var(--ease);
    }

    input:checked + .track {
        background: var(--accent);
        border-color: transparent;
    }
    input:checked + .track .thumb {
        transform: translateX(18px);
    }

    input:focus-visible + .track {
        outline: 2px solid var(--accent-ring);
        outline-offset: 2px;
    }

    .label-text {
        font-size: 0.9rem;
        color: var(--ink-2);
    }
</style>
```

- [ ] **Step 2: Type-check + commit**

```bash
cd bingo-app && npm run check
git add bingo-app/src/lib/ui/Toggle.svelte
git commit -m "feat(ui): add Toggle pill switch"
```

---

### Task 1.11: Select primitive

**Files:**
- Create: `bingo-app/src/lib/ui/Select.svelte`

- [ ] **Step 1: Write `bingo-app/src/lib/ui/Select.svelte`**

```svelte
<script lang="ts">
    /**
     * Select — paper-styled native <select> with floating label and chevron.
     * Sticks to native semantics for free a11y and mobile UX.
     */
    export let label: string;
    export let value: string = '';
    export let items: { value: string; name: string }[] = [];
    export let name: string | undefined = undefined;
    export let id: string | undefined = undefined;
    export let required = false;
    export let disabled = false;

    let focused = false;
    $: filled = value !== '' && value !== undefined;
    $: lifted = focused || filled;
    $: selectId = id ?? `select-${name ?? Math.random().toString(36).slice(2, 8)}`;
</script>

<label class="field">
    <span class="label" class:lifted>{label}{required ? ' *' : ''}</span>
    <select
        id={selectId}
        {name}
        {required}
        {disabled}
        bind:value
        on:focus={() => (focused = true)}
        on:blur={() => (focused = false)}
        on:change
    >
        <option value="" disabled hidden></option>
        {#each items as item (item.value)}
            <option value={item.value}>{item.name}</option>
        {/each}
    </select>
    <svg class="chevron" viewBox="0 0 20 20" aria-hidden="true">
        <path d="M5 8l5 5 5-5" fill="none" stroke="currentColor" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round"/>
    </svg>
</label>

<style>
    .field { position: relative; display: flex; flex-direction: column; }

    .label {
        position: absolute;
        left: 1rem;
        top: 50%;
        transform: translateY(-50%);
        color: var(--ink-3);
        font-size: 0.95rem;
        pointer-events: none;
        transition: top var(--dur-fast) var(--ease),
                    transform var(--dur-fast) var(--ease),
                    font-size var(--dur-fast) var(--ease);
    }
    .label.lifted {
        top: 0.4rem;
        transform: translateY(0);
        font-size: 0.7rem;
        text-transform: uppercase;
        letter-spacing: 0.04em;
    }

    select {
        appearance: none;
        -webkit-appearance: none;
        background: var(--paper);
        border: 1px solid var(--hairline);
        border-radius: var(--radius-md);
        padding: 1.5rem 2.4rem 0.5rem 1rem;
        color: var(--ink);
        font-size: 0.95rem;
        min-height: 56px;
        outline: none;
        cursor: pointer;
        transition: border-color var(--dur-fast) var(--ease),
                    box-shadow var(--dur-fast) var(--ease);
    }
    select:focus {
        border-color: var(--accent-from);
        box-shadow: 0 0 0 3px var(--accent-ring);
    }

    .chevron {
        position: absolute;
        right: 1rem;
        top: 50%;
        transform: translateY(-50%);
        width: 18px;
        height: 18px;
        color: var(--ink-3);
        pointer-events: none;
    }
</style>
```

- [ ] **Step 2: Type-check + commit**

```bash
cd bingo-app && npm run check
git add bingo-app/src/lib/ui/Select.svelte
git commit -m "feat(ui): add Select primitive (paper-styled native)"
```

---

### Task 1.12: TabBar primitive

**Files:**
- Create: `bingo-app/src/lib/ui/TabBar.svelte`

- [ ] **Step 1: Write `bingo-app/src/lib/ui/TabBar.svelte`**

```svelte
<script lang="ts">
    /**
     * TabBar — segmented control. Two-way bind on `value`. Same component
     * is used as a tab strip on desktop and as a segmented control on mobile.
     */
    export let value: string;
    export let items: { value: string; label: string }[] = [];
    export let ariaLabel = 'Sections';
</script>

<div class="tabs" role="tablist" aria-label={ariaLabel}>
    {#each items as item (item.value)}
        <button
            type="button"
            role="tab"
            class="tab"
            class:active={value === item.value}
            aria-selected={value === item.value}
            on:click={() => (value = item.value)}
        >
            {item.label}
        </button>
    {/each}
</div>

<style>
    .tabs {
        display: inline-flex;
        gap: 2px;
        padding: 4px;
        background: var(--paper-soft);
        border-radius: var(--radius-pill);
        border: 1px solid var(--hairline);
    }

    .tab {
        appearance: none;
        background: transparent;
        border: none;
        padding: 0.45rem 0.95rem;
        border-radius: var(--radius-pill);
        font-family: var(--font-body);
        font-size: 0.78rem;
        font-weight: 500;
        color: var(--ink-2);
        cursor: pointer;
        min-height: 36px;
        transition: background var(--dur-fast) var(--ease),
                    color var(--dur-fast) var(--ease);
    }

    .tab:hover { color: var(--ink); }
    .tab.active {
        background: var(--paper);
        color: var(--ink);
        box-shadow: 0 1px 2px rgba(40, 55, 95, 0.06);
    }
    .tab:focus-visible {
        outline: 2px solid var(--accent-ring);
        outline-offset: 2px;
    }
</style>
```

- [ ] **Step 2: Type-check + commit**

```bash
cd bingo-app && npm run check
git add bingo-app/src/lib/ui/TabBar.svelte
git commit -m "feat(ui): add TabBar segmented control"
```

---

### Task 1.13: Wire Atmosphere into the layout

**Files:**
- Modify: `bingo-app/src/routes/+layout.svelte`

- [ ] **Step 1: Replace `bingo-app/src/routes/+layout.svelte` entirely**

```svelte
<script lang="ts">
    import Header from '../components/Header/Header.svelte';
    import Atmosphere from '$lib/ui/Atmosphere.svelte';
    import '../app.postcss';
</script>

<svelte:head>
    <link rel="icon" type="image/svg" href="/favicon.svg" />
</svelte:head>

<a href="#main" class="skip-link">Skip to main content</a>
<Atmosphere>
    <Header />
    <div id="main">
        <slot />
    </div>
</Atmosphere>
```

Note: `Header` is the existing flowbite-based component. It keeps working — Phase 2 swaps it.

- [ ] **Step 2: Type-check + run dev**

```bash
cd bingo-app
npm run check
npm run dev
```

Open http://localhost:5173 and confirm: pastel mesh atmosphere visible, noise grain visible at close inspection, existing flowbite header still renders on top, page content sits over the noise overlay.

- [ ] **Step 3: Commit**

```bash
git add bingo-app/src/routes/+layout.svelte
git commit -m "feat(ui): wrap layout in Atmosphere primitive"
```

---

### Task 1.14: Barrel export `$lib/ui`

**Files:**
- Create: `bingo-app/src/lib/ui/index.ts`

- [ ] **Step 1: Write `bingo-app/src/lib/ui/index.ts`**

```ts
export { default as Atmosphere } from './Atmosphere.svelte';
export { default as Paper } from './Paper.svelte';
export { default as Glass } from './Glass.svelte';
export { default as Page } from './Page.svelte';
export { default as Stack } from './Stack.svelte';
export { default as Cluster } from './Cluster.svelte';
export { default as Button } from './Button.svelte';
export { default as Input } from './Input.svelte';
export { default as Textarea } from './Textarea.svelte';
export { default as Toggle } from './Toggle.svelte';
export { default as Select } from './Select.svelte';
export { default as TabBar } from './TabBar.svelte';
```

- [ ] **Step 2: Type-check + commit**

```bash
cd bingo-app && npm run check
git add bingo-app/src/lib/ui/index.ts
git commit -m "feat(ui): barrel export design-system primitives"
```

---

## Phase 2 — Public pages

Header replaces flowbite's Navbar. Then `/`, `/about`, `/login`, `/register`. After this phase, all unauthenticated routes use the new design system.

### Task 2.1: New Header

**Files:**
- Create: `bingo-app/src/lib/feature/Header.svelte`
- Modify: `bingo-app/src/routes/+layout.svelte` (swap import)

- [ ] **Step 1: Write `bingo-app/src/lib/feature/Header.svelte`**

```svelte
<script lang="ts">
    import { onDestroy, onMount } from 'svelte';
    import { browser } from '$app/environment';
    import { page } from '$app/stores';
    import Glass from '$lib/ui/Glass.svelte';

    let scrolled = false;
    let authed = false;

    function readAuth() {
        if (!browser) return false;
        try {
            const cookieString = RegExp('auth=[^;]+').exec(document.cookie);
            if (!cookieString) return false;
            const token = decodeURIComponent(cookieString.toString().replace(/^[^=]+./, ''));
            const base64Url = token.split('.')[1];
            const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
            const jsonPayload = decodeURIComponent(window.atob(base64).split('').map((c) => '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2)).join(''));
            const data = JSON.parse(jsonPayload);
            return !!data.user;
        } catch {
            return false;
        }
    }

    function onScroll() {
        scrolled = window.scrollY > 64;
    }

    onMount(() => {
        if (!browser) return;
        authed = readAuth();
        onScroll();
        window.addEventListener('scroll', onScroll, { passive: true });
    });

    onDestroy(() => {
        if (!browser) return;
        window.removeEventListener('scroll', onScroll);
    });

    $: pathname = $page.url.pathname;
    $: isActive = (href: string) => pathname === href || (href !== '/' && pathname.startsWith(href));
</script>

<div class="header" class:scrolled>
    {#if scrolled}
        <Glass radius="card" blur="strong" let:dummy>
            <div class="row">
                <a class="brand" href="/">taskbingo<span class="tag">beta</span></a>
                <nav class="nav" aria-label="Primary">
                    {#if authed}
                        <a class:active={isActive('/account')} href="/account" data-sveltekit-prefetch>play</a>
                        <a class:active={isActive('/packs')} href="/packs" data-sveltekit-prefetch>packs</a>
                        <a class:active={isActive('/people')} href="/people" data-sveltekit-prefetch>people</a>
                    {/if}
                    <a class:active={isActive('/about')} href="/about">about</a>
                    {#if authed}
                        <a class="avatar" href="/account" aria-label="Account" data-sveltekit-prefetch></a>
                    {:else}
                        <a class:active={isActive('/login')} href="/login">login</a>
                    {/if}
                </nav>
            </div>
        </Glass>
    {:else}
        <div class="row plain">
            <a class="brand" href="/">taskbingo<span class="tag">beta</span></a>
            <nav class="nav" aria-label="Primary">
                {#if authed}
                    <a class:active={isActive('/account')} href="/account" data-sveltekit-prefetch>play</a>
                    <a class:active={isActive('/packs')} href="/packs" data-sveltekit-prefetch>packs</a>
                    <a class:active={isActive('/people')} href="/people" data-sveltekit-prefetch>people</a>
                {/if}
                <a class:active={isActive('/about')} href="/about">about</a>
                {#if authed}
                    <a class="avatar" href="/account" aria-label="Account" data-sveltekit-prefetch></a>
                {:else}
                    <a class:active={isActive('/login')} href="/login">login</a>
                {/if}
            </nav>
        </div>
    {/if}
</div>

<style>
    .header {
        position: sticky;
        top: 0;
        z-index: 10;
        padding: var(--space-3) var(--space-5);
        transition: padding var(--dur-base) var(--ease);
    }
    .header.scrolled {
        padding: var(--space-2) var(--space-5);
    }
    .row {
        display: flex;
        align-items: center;
        justify-content: space-between;
        gap: var(--space-5);
        max-width: 80rem;
        margin: 0 auto;
        padding: var(--space-2) var(--space-4);
    }
    .row.plain {
        padding: var(--space-3) 0;
    }

    .brand {
        font-family: var(--font-display);
        font-size: 1.15rem;
        color: var(--ink);
        text-decoration: none;
        letter-spacing: -0.012em;
    }
    .tag {
        margin-left: 0.5em;
        font-family: var(--font-body);
        font-size: 0.6rem;
        color: var(--ink-3);
        letter-spacing: 0.16em;
        text-transform: uppercase;
        font-weight: 500;
    }

    .nav {
        display: flex;
        align-items: center;
        gap: var(--space-5);
        font-size: 0.78rem;
        font-weight: 500;
        color: var(--ink-2);
    }
    .nav a {
        text-decoration: none;
        color: inherit;
        position: relative;
        padding-bottom: 2px;
    }
    .nav a:hover { color: var(--ink); }
    .nav a.active {
        color: var(--ink);
    }
    .nav a.active::after {
        content: '';
        position: absolute;
        left: 0; right: 0; bottom: -8px; height: 1px;
        background: var(--accent);
        border-radius: 1px;
    }

    .avatar {
        width: 28px;
        height: 28px;
        border-radius: var(--radius-pill);
        background: var(--accent);
        box-shadow: 0 0 0 1px rgba(255, 255, 255, 0.7);
    }

    @media (max-width: 640px) {
        .row { gap: var(--space-3); padding: var(--space-2); }
        .nav { gap: var(--space-3); font-size: 0.72rem; }
    }
</style>
```

- [ ] **Step 2: Swap the import in `bingo-app/src/routes/+layout.svelte`**

Replace this line:

```svelte
import Header from '../components/Header/Header.svelte';
```

with:

```svelte
import Header from '$lib/feature/Header.svelte';
```

- [ ] **Step 3: Run dev and verify**

```bash
cd bingo-app && npm run dev
```

Open http://localhost:5173 and confirm:
- Header sits as plain text on the atmosphere at the top.
- Scroll the page (any route) — when scrolled past 64px, header transitions to a glass strip with backdrop blur.
- "play / packs / people" appear when authenticated; "login" appears when not.
- Active route has a gradient hairline underline.
- Mobile (DevTools 375px width): nav links wrap or stay inline, header still works.

- [ ] **Step 4: Type-check + commit**

```bash
cd bingo-app && npm run check
git add bingo-app/src/lib/feature/Header.svelte bingo-app/src/routes/+layout.svelte
git commit -m "feat(ui): replace Header with custom glass-on-scroll variant"
```

---

### Task 2.2: Redesign `/`

**Files:**
- Modify: `bingo-app/src/routes/+page.svelte` (replace whole file)

- [ ] **Step 1: Replace `bingo-app/src/routes/+page.svelte`**

```svelte
<script lang="ts">
    import { browser } from '$app/environment';
    import Page from '$lib/ui/Page.svelte';
    import Stack from '$lib/ui/Stack.svelte';
    import Cluster from '$lib/ui/Cluster.svelte';
    import Button from '$lib/ui/Button.svelte';
    import Paper from '$lib/ui/Paper.svelte';

    let authed = false;
    if (browser) {
        authed = RegExp('auth=[^;]+').exec(document.cookie) !== null;
    }
</script>

<svelte:head>
    <title>taskbingo</title>
</svelte:head>

<Page width="wide">
    <div class="hero">
        <div class="hero__text">
            <span class="eyebrow">premise</span>
            <h1 class="display">Sixteen tasks.<br/><em>One quiet board.</em></h1>
            <p class="lead">
                Договоритесь о шестнадцати делах, отметьте сделанное, собирайте линии.
                Удача и труд на одной сетке.
            </p>
            <Cluster gap="m" align="center">
                {#if authed}
                    <Button variant="accent" size="lg" href="/account">Continue playing</Button>
                    <Button variant="quiet" href="/packs">browse packs</Button>
                {:else}
                    <Button variant="accent" size="lg" href="/register">Start a game</Button>
                    <Button variant="quiet" href="/login">log in</Button>
                {/if}
            </Cluster>
        </div>

        <div class="hero__steps">
            <Paper padding="lg">
                <Stack gap="l">
                    <span class="eyebrow">how it works</span>
                    <ol class="steps">
                        <li><span class="step__num">1</span><span>Invite a friend or start solo.</span></li>
                        <li><span class="step__num">2</span><span>Agree on a pack of 16 tasks.</span></li>
                        <li><span class="step__num">3</span><span>Mark them off as you go.</span></li>
                    </ol>
                </Stack>
            </Paper>
        </div>
    </div>
</Page>

<style>
    .hero {
        display: grid;
        grid-template-columns: 1.05fr 0.95fr;
        gap: var(--space-7);
        align-items: center;
        min-height: 60vh;
    }
    @media (max-width: 900px) {
        .hero { grid-template-columns: 1fr; gap: var(--space-6); }
    }

    .hero__text { display: flex; flex-direction: column; gap: var(--space-4); max-width: 30rem; }

    .eyebrow {
        font-size: 0.65rem;
        letter-spacing: 0.18em;
        text-transform: uppercase;
        color: var(--ink-3);
    }

    .display {
        font-family: var(--font-display);
        font-weight: 400;
        font-size: clamp(2.2rem, 4vw, 3rem);
        line-height: 1.05;
        letter-spacing: -0.022em;
        color: var(--ink);
    }
    .display em { font-style: italic; color: var(--ink-3); font-weight: 300; }

    .lead {
        font-size: 0.95rem;
        line-height: 1.6;
        color: var(--ink-2);
        max-width: 36ch;
    }

    .steps {
        list-style: none;
        margin: 0; padding: 0;
        display: flex; flex-direction: column;
        gap: var(--space-3);
    }
    .steps li {
        display: flex;
        align-items: flex-start;
        gap: var(--space-3);
        font-size: 0.9rem;
        color: var(--ink-2);
        line-height: 1.5;
    }
    .step__num {
        display: inline-flex;
        align-items: center; justify-content: center;
        width: 24px; height: 24px;
        background: var(--accent);
        color: var(--on-accent);
        border-radius: var(--radius-pill);
        font-family: var(--font-display);
        font-size: 0.78rem;
        flex-shrink: 0;
    }
</style>
```

- [ ] **Step 2: Run dev and verify**

```bash
cd bingo-app && npm run dev
```

Open http://localhost:5173 and confirm:
- Hero text floats on the atmosphere on the left, paper card with steps on the right.
- Below 900px width, hero stacks vertically.
- "Start a game" / "Continue playing" works; the quiet text link below works.

- [ ] **Step 3: Type-check + commit**

```bash
cd bingo-app && npm run check
git add bingo-app/src/routes/+page.svelte
git commit -m "feat(ui): redesign / hero with paper steps card"
```

---

### Task 2.3: Redesign `/login`

**Files:**
- Modify: `bingo-app/src/routes/login/+page.svelte` (replace whole file)

- [ ] **Step 1: Replace `bingo-app/src/routes/login/+page.svelte`**

```svelte
<script lang="ts">
    import { goto } from '$app/navigation';
    import Page from '$lib/ui/Page.svelte';
    import Stack from '$lib/ui/Stack.svelte';
    import Paper from '$lib/ui/Paper.svelte';
    import Input from '$lib/ui/Input.svelte';
    import Button from '$lib/ui/Button.svelte';
    import { API_URL, WEB_URL } from '../temporary';

    let username = '';
    let password = '';
    let showPassword = false;
    let error = '';
    let submitting = false;

    async function submit(e: Event) {
        e.preventDefault();
        error = '';
        submitting = true;

        const res = await fetch(`${API_URL}/api/user/login`, {
            method: 'POST',
            headers: { 'Origin': WEB_URL },
            body: JSON.stringify({ username, password }),
            credentials: 'include',
        });

        submitting = false;

        if (res.ok) {
            goto('/account');
        } else {
            error = res.status === 401 ? 'Incorrect username or password.' : 'Something went wrong. Try again.';
        }
    }
</script>

<svelte:head><title>Log in · taskbingo</title></svelte:head>

<Page width="narrow">
    <Stack gap="l" align="stretch">
        <h1 class="title">Welcome back.</h1>
        <Paper padding="lg">
            <form on:submit={submit}>
                <Stack gap="m">
                    <Input label="Username" name="username" bind:value={username} required />
                    <div class="pw">
                        <Input
                            label="Password"
                            name="password"
                            type={showPassword ? 'text' : 'password'}
                            bind:value={password}
                            required
                        />
                        <button type="button" class="eye" on:click={() => (showPassword = !showPassword)} aria-label={showPassword ? 'Hide password' : 'Show password'}>
                            {#if showPassword}
                                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round" width="20" height="20"><path d="M2 12s3.5-7 10-7 10 7 10 7-3.5 7-10 7-10-7-10-7z"/><circle cx="12" cy="12" r="3"/></svg>
                            {:else}
                                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round" width="20" height="20"><path d="M3 3l18 18"/><path d="M2 12s3.5-7 10-7c2 0 3.7.6 5.1 1.5"/><path d="M9.5 9.5a3 3 0 0 0 4.2 4.2"/><path d="M22 12s-3.5 7-10 7c-1.4 0-2.6-.3-3.7-.8"/></svg>
                            {/if}
                        </button>
                    </div>
                    {#if error}
                        <span class="error" role="alert">{error}</span>
                    {/if}
                    <Button type="submit" variant="accent" fullWidth disabled={submitting}>
                        {submitting ? 'Logging in…' : 'Log in'}
                    </Button>
                </Stack>
            </form>
        </Paper>
        <p class="alt">New here? <a href="/register">Create an account</a></p>
    </Stack>
</Page>

<style>
    .title {
        font-family: var(--font-display);
        font-size: 1.6rem;
        font-weight: 400;
        letter-spacing: -0.015em;
        color: var(--ink);
        text-align: center;
    }

    .pw { position: relative; }
    .eye {
        position: absolute;
        right: 0.75rem;
        top: 50%;
        transform: translateY(-50%);
        background: transparent;
        border: none;
        color: var(--ink-3);
        cursor: pointer;
        padding: 0.5rem;
        border-radius: var(--radius-md);
    }
    .eye:hover { color: var(--ink); }
    .eye:focus-visible { outline: 2px solid var(--accent-ring); outline-offset: 2px; }

    .error {
        color: #c44e4e;
        font-size: 0.85rem;
        text-align: center;
    }

    .alt {
        text-align: center;
        font-size: 0.85rem;
        color: var(--ink-2);
    }
    .alt a { color: var(--ink); border-bottom: 1px solid rgba(40, 55, 95, 0.25); text-decoration: none; }
</style>
```

- [ ] **Step 2: Run dev and verify**

```bash
cd bingo-app && npm run dev
```

Open http://localhost:5173/login and confirm:
- Centred narrow paper card on the atmosphere.
- Floating labels lift on focus or value.
- Eye icon toggles password visibility.
- Form submits and redirects to /account on success (requires running backend).
- Error message appears on failed login.

- [ ] **Step 3: Type-check + commit**

```bash
cd bingo-app && npm run check
git add bingo-app/src/routes/login/+page.svelte
git commit -m "feat(ui): redesign /login with paper card form"
```

---

### Task 2.4: Redesign `/register`

**Files:**
- Modify: `bingo-app/src/routes/register/+page.svelte` (replace whole file)

- [ ] **Step 1: Replace `bingo-app/src/routes/register/+page.svelte`**

```svelte
<script lang="ts">
    import { goto } from '$app/navigation';
    import Page from '$lib/ui/Page.svelte';
    import Stack from '$lib/ui/Stack.svelte';
    import Paper from '$lib/ui/Paper.svelte';
    import Input from '$lib/ui/Input.svelte';
    import Button from '$lib/ui/Button.svelte';
    import { _Submit } from './+page';

    let username = '';
    let email = '';
    let city = '';
    let password = '';
    let showPassword = false;
    let error = '';
    let submitting = false;

    async function submit(e: Event) {
        e.preventDefault();
        error = '';
        submitting = true;
        const status = await _Submit(e);
        submitting = false;

        if (status === 200) {
            goto('/account');
        } else if (status === 409) {
            error = 'Username or email already in use.';
        } else {
            error = 'Something went wrong. Try again.';
        }
    }
</script>

<svelte:head><title>Register · taskbingo</title></svelte:head>

<Page width="narrow">
    <Stack gap="l" align="stretch">
        <h1 class="title">Create an account.</h1>
        <Paper padding="lg">
            <form on:submit={submit}>
                <Stack gap="m">
                    <Input label="Username" name="username" bind:value={username} required />
                    <Input label="Email" name="email" type="email" bind:value={email} required />
                    <Input label="City" name="city" bind:value={city} required />
                    <div class="pw">
                        <Input
                            label="Password"
                            name="password"
                            type={showPassword ? 'text' : 'password'}
                            bind:value={password}
                            required
                        />
                        <button type="button" class="eye" on:click={() => (showPassword = !showPassword)} aria-label={showPassword ? 'Hide password' : 'Show password'}>
                            {#if showPassword}
                                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round" width="20" height="20"><path d="M2 12s3.5-7 10-7 10 7 10 7-3.5 7-10 7-10-7-10-7z"/><circle cx="12" cy="12" r="3"/></svg>
                            {:else}
                                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round" width="20" height="20"><path d="M3 3l18 18"/><path d="M2 12s3.5-7 10-7c2 0 3.7.6 5.1 1.5"/><path d="M9.5 9.5a3 3 0 0 0 4.2 4.2"/><path d="M22 12s-3.5 7-10 7c-1.4 0-2.6-.3-3.7-.8"/></svg>
                            {/if}
                        </button>
                    </div>
                    {#if error}
                        <span class="error" role="alert">{error}</span>
                    {/if}
                    <Button type="submit" variant="accent" fullWidth disabled={submitting}>
                        {submitting ? 'Creating…' : 'Register'}
                    </Button>
                </Stack>
            </form>
        </Paper>
        <p class="alt">Have an account? <a href="/login">Log in</a></p>
    </Stack>
</Page>

<style>
    .title {
        font-family: var(--font-display);
        font-size: 1.6rem;
        font-weight: 400;
        text-align: center;
    }
    .pw { position: relative; }
    .eye {
        position: absolute; right: 0.75rem; top: 50%; transform: translateY(-50%);
        background: transparent; border: none; color: var(--ink-3); cursor: pointer; padding: 0.5rem; border-radius: var(--radius-md);
    }
    .eye:hover { color: var(--ink); }
    .eye:focus-visible { outline: 2px solid var(--accent-ring); outline-offset: 2px; }

    .error { color: #c44e4e; font-size: 0.85rem; text-align: center; }
    .alt { text-align: center; font-size: 0.85rem; color: var(--ink-2); }
    .alt a { color: var(--ink); border-bottom: 1px solid rgba(40, 55, 95, 0.25); text-decoration: none; }
</style>
```

- [ ] **Step 2: Verify in browser**

Open http://localhost:5173/register and confirm same patterns as /login. Try registering — on conflict the inline error appears; on success redirects to /account.

- [ ] **Step 3: Type-check + commit**

```bash
cd bingo-app && npm run check
git add bingo-app/src/routes/register/+page.svelte
git commit -m "feat(ui): redesign /register with paper card form"
```

---

### Task 2.5: Redesign `/about`

**Files:**
- Modify: `bingo-app/src/routes/about/+page.svelte` (replace whole file)

- [ ] **Step 1: Replace `bingo-app/src/routes/about/+page.svelte`**

```svelte
<script lang="ts">
    import Page from '$lib/ui/Page.svelte';
    import Stack from '$lib/ui/Stack.svelte';

    type Entry = { date: string; title: string; body: string; isFinal?: boolean };

    const entries: Entry[] = [
        { date: 'May 2020', title: 'Genesis', body: "In the throes of Covid lockdowns, when days melded into nights, my productivity levels were so low, they would make a sloth look like an overachiever. But then, amidst the haze of takeout containers and unwashed laundry, the concept of Task Bingo took its first breath." },
        { date: 'June 2020', title: 'Humble Beginnings', body: "Frankly, our initial version wasn't winning any beauty contests. It was a scrappy mess hastily pieced together. The graphics? Primitive. The UX? Don't even ask. But it worked, and that's what mattered." },
        { date: 'August 2020', title: 'Hiatus', body: "As quarantine faded, our lives began to resume some semblance of normality. Energized, productive, and even a tad athletic, we drifted away from our creation. The Task Bingo files were cast aside, lost in the chaos of other fleeting aspirations." },
        { date: 'December 2022', title: 'The Awakening', body: "Fast forward a bit. Motivation? Plummeted. Enthusiasm? Non-existent. As the days blurred, a ghost of an idea whispered, \"Remember Task Bingo?\" It wasn't so much an epiphany, more like stubbing your toe and finding a lost toy. However, this time, the bar was set higher." },
        { date: 'January 2023', title: 'The Grind', body: "Building the new version wasn't a walk in the park. More like trudging through a swamp with a backpack full of bricks. Between the coding hiccups and the occasional bout of good ol' procrastination, the journey had its fair share of bumps. But hey, nothing that a bit of elbow grease couldn't fix." },
        { date: 'August 2023', title: 'Endgame', isFinal: true, body: "Here we are, a tad later than anticipated — the beta version of taskbingo is now a reality. Still rough around the edges, maybe has a bug or twenty, but damn, isn't it a sight to behold? Plans for development are brewing, but for the moment, taskbingo is taking a short intermission. Stay tuned." },
    ];
</script>

<svelte:head><title>About · taskbingo</title></svelte:head>

<Page width="narrow">
    <Stack gap="xl">
        <header class="head">
            <h1 class="display">About</h1>
            <p class="meta">
                Made by <a href="https://t.me/dupreehkuda" target="_blank" rel="noopener noreferrer">@dupreehkuda</a>
                · <a href="https://github.com/dupreehkuda/taskbingo" target="_blank" rel="noopener noreferrer">repo</a>
            </p>
        </header>

        <ol class="timeline">
            {#each entries as e (e.date)}
                <li class="entry">
                    <span class="date">{e.date}</span>
                    <h3 class="title" class:final={e.isFinal}>{e.title}</h3>
                    <p class="body">{e.body}</p>
                </li>
            {/each}
        </ol>
    </Stack>
</Page>

<style>
    .head { text-align: center; }
    .display {
        font-family: var(--font-display);
        font-size: 2rem;
        font-weight: 400;
        letter-spacing: -0.018em;
    }
    .meta { color: var(--ink-3); font-size: 0.85rem; margin-top: var(--space-2); }
    .meta a { color: var(--ink-2); border-bottom: 1px solid rgba(40, 55, 95, 0.18); text-decoration: none; }
    .meta a:hover { color: var(--ink); }

    .timeline {
        list-style: none;
        margin: 0; padding: 0;
        display: flex; flex-direction: column;
        gap: var(--space-7);
        border-left: 1px solid var(--hairline);
        padding-left: var(--space-5);
        position: relative;
    }
    .timeline::before {
        content: '';
        position: absolute;
        left: -1px;
        top: 0;
        bottom: 0;
        width: 1px;
        background: var(--accent);
        opacity: 0.4;
    }

    .entry { display: flex; flex-direction: column; gap: var(--space-2); }
    .date {
        font-family: var(--font-display);
        font-style: italic;
        font-weight: 300;
        font-size: 0.85rem;
        color: var(--ink-3);
    }
    .title {
        font-family: var(--font-display);
        font-weight: 400;
        font-size: 1.1rem;
        color: var(--ink);
        margin: 0;
    }
    .title.final {
        background: var(--accent);
        -webkit-background-clip: text;
        background-clip: text;
        color: transparent;
    }
    .body {
        font-size: 0.9rem;
        line-height: 1.6;
        color: var(--ink-2);
    }
</style>
```

- [ ] **Step 2: Verify in browser**

Open http://localhost:5173/about and confirm:
- Narrow column on atmosphere.
- Hairline gradient on the left edge of the timeline.
- "Endgame" title shows the periwinkle gradient via background-clip.

- [ ] **Step 3: Type-check + commit**

```bash
cd bingo-app && npm run check
git add bingo-app/src/routes/about/+page.svelte
git commit -m "feat(ui): redesign /about as prose timeline"
```

---

## Phase 3 — User dashboard

`/account`, `/people`, `/packs`. Carousels removed; vertical stream layouts.

### Task 3.1: Redesign `/account`

**Files:**
- Modify: `bingo-app/src/routes/account/+page.svelte` (replace whole file)

- [ ] **Step 1: Replace `bingo-app/src/routes/account/+page.svelte`**

```svelte
<script lang="ts">
    import Account from '../accountStore';
    import Page from '$lib/ui/Page.svelte';
    import Stack from '$lib/ui/Stack.svelte';
    import Cluster from '$lib/ui/Cluster.svelte';
    import Paper from '$lib/ui/Paper.svelte';
    import Button from '$lib/ui/Button.svelte';
    import GameModal from '$lib/feature/GameModal.svelte';
    import { _LikePack, _DeleteGame, _GetGame } from './+page';
    import { DeleteFriend, AcceptFriend } from '../friendRequests';

    let showModal = false;
    let selectedPackID = '';
    let selectedFriendID = '';
    let showAllFriends = false;
    let showAllPacks = false;
    let showAllGames = false;

    function openWithPack(packID: string) { selectedPackID = packID; selectedFriendID = ''; showModal = true; }
    function openWithFriend(friendID: string) { selectedFriendID = friendID; selectedPackID = ''; showModal = true; }
    function openEmpty() { selectedFriendID = ''; selectedPackID = ''; showModal = true; }

    function getOpponentUsername(userId: string) {
        for (const f of $Account.friends) if (f.userID === userId) return f.username;
        return '—';
    }
    function getPackTitle(packID: string) {
        const liked = $Account.likedPacks.find((p) => p.id === packID);
        if (liked) return liked.pack.title;
        const owned = $Account.packs?.find((p) => p.id === packID);
        return owned?.pack.title ?? '—';
    }

    $: friendsAccepted = $Account?.friends.filter((f) => f.status === 3) ?? [];
    $: friendsRequests = $Account?.friends.filter((f) => f.status === 2) ?? [];
    $: friendsSent = $Account?.friends.filter((f) => f.status === 1) ?? [];
    $: friendsAll = [...friendsRequests, ...friendsAccepted, ...friendsSent];
    $: friendsVisible = showAllFriends ? friendsAll : friendsAll.slice(0, 5);
    $: packsAll = $Account?.likedPacks ?? [];
    $: packsVisible = showAllPacks ? packsAll : packsAll.slice(0, 5);
    $: activeGames = $Account?.games.filter((g) => g.status !== 3) ?? [];
    $: featuredGame = activeGames[0];
    $: pastGames = $Account?.games.filter((g) => g.status === 3) ?? [];
    $: pastVisible = showAllGames ? pastGames : [];
</script>

<svelte:head><title>{$Account?.username ?? 'Account'} · taskbingo</title></svelte:head>

<Page width="normal">
    <Stack gap="xl">
        <!-- Greeting block on atmosphere -->
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
        {/if}

        <!-- Continue playing -->
        {#if featuredGame}
            <section>
                <h2 class="section-title">Continue playing</h2>
                <Paper padding="lg">
                    <Cluster gap="m" align="center" justify="between">
                        <Stack gap="s">
                            <span class="game-title">
                                {#if featuredGame.kind === 'solo'}
                                    Solo game
                                {:else if $Account.userID === featuredGame.user1Id}
                                    vs. {getOpponentUsername(featuredGame.user2Id)}
                                {:else}
                                    vs. {getOpponentUsername(featuredGame.user1Id)}
                                {/if}
                            </span>
                            <span class="game-sub">{getPackTitle(featuredGame.packId)}</span>
                        </Stack>
                        <Button
                            variant="accent"
                            href={featuredGame.kind === 'solo' ? '/game?solo=true' : '/game'}
                            on:mouseenter={() => _GetGame(featuredGame.gameId)}
                        >
                            {featuredGame.kind === 'solo' ? 'Resume' : 'Start'}
                        </Button>
                    </Cluster>
                </Paper>
            </section>
        {/if}

        <!-- Friends -->
        <section>
            <h2 class="section-title">Friends</h2>
            <Stack gap="s">
                {#each friendsVisible as friend (friend.userID)}
                    <Paper padding="sm">
                        <Cluster gap="m" align="center" justify="between">
                            <Stack gap="s">
                                <span class="friend-name">{friend.username}</span>
                                <span class="friend-meta">{friend.wins}/{friend.loses}</span>
                            </Stack>
                            <Cluster gap="s">
                                {#if friend.status === 3}
                                    <Button variant="accent" size="sm" on:click={() => openWithFriend(friend.userID)}>Play</Button>
                                    <Button variant="ghost" size="sm" on:click={() => DeleteFriend(friend.userID)}>Remove</Button>
                                {:else if friend.status === 2}
                                    <Button variant="accent" size="sm" on:click={() => AcceptFriend(friend.userID)}>Accept</Button>
                                    <Button variant="ghost" size="sm" on:click={() => DeleteFriend(friend.userID)}>Decline</Button>
                                {:else if friend.status === 1}
                                    <Button variant="ghost" size="sm" disabled>Sent</Button>
                                    <Button variant="ghost" size="sm" on:click={() => DeleteFriend(friend.userID)}>Cancel</Button>
                                {/if}
                            </Cluster>
                        </Cluster>
                    </Paper>
                {/each}
                {#if friendsAll.length === 0}
                    <Paper padding="md"><span class="empty">No friends yet — head to <a href="/people">people</a>.</span></Paper>
                {:else if friendsAll.length > 5}
                    <button class="see-all" on:click={() => (showAllFriends = !showAllFriends)}>
                        {showAllFriends ? 'show less' : `see all ${friendsAll.length}`}
                    </button>
                {/if}
            </Stack>
        </section>

        <!-- Packs -->
        <section>
            <Cluster gap="m" align="baseline" justify="between">
                <h2 class="section-title">Packs</h2>
                <Button variant="ghost" size="sm" href="/newpack">Create new pack</Button>
            </Cluster>
            <Stack gap="s">
                {#each packsVisible as pack (pack.id)}
                    <Paper padding="md">
                        <Stack gap="s">
                            <Cluster gap="m" align="baseline" justify="between">
                                <span class="pack-title">
                                    {pack.pack.title}
                                    {#if pack.isPrivate}<span class="badge">private</span>{/if}
                                </span>
                            </Cluster>
                            <ul class="pack-tasks">
                                {#each pack.pack.tasks.slice(0, 3) as task, i}
                                    <li><span class="i">{i + 1}</span><span>{task}</span></li>
                                {/each}
                                {#if pack.pack.tasks.length > 3}
                                    <li class="more">+ {pack.pack.tasks.length - 3} more</li>
                                {/if}
                            </ul>
                            <Cluster gap="s">
                                <Button variant="accent" size="sm" on:click={() => openWithPack(pack.id)}>Choose pack</Button>
                                <Button variant="ghost" size="sm" on:click={() => _LikePack(pack, true)} aria-label="Unlike pack">
                                    <span aria-hidden="true">♥</span>
                                </Button>
                            </Cluster>
                        </Stack>
                    </Paper>
                {/each}
                {#if packsAll.length === 0}
                    <Paper padding="md"><span class="empty">No packs liked yet — tap the heart on any pack in <a href="/packs">packs</a>.</span></Paper>
                {:else if packsAll.length > 5}
                    <button class="see-all" on:click={() => (showAllPacks = !showAllPacks)}>
                        {showAllPacks ? 'show less' : `see all ${packsAll.length}`}
                    </button>
                {/if}
            </Stack>
        </section>

        <!-- Games -->
        <section>
            <Cluster gap="m" align="baseline" justify="between">
                <h2 class="section-title">Games</h2>
                <Button variant="ghost" size="sm" on:click={openEmpty}>Create new game</Button>
            </Cluster>
            <Stack gap="s">
                {#each activeGames as game (game.gameId)}
                    <Paper padding="sm">
                        <Cluster gap="m" align="center" justify="between">
                            <Stack gap="s">
                                <span class="game-title">
                                    {#if game.kind === 'solo'}
                                        Solo · {getPackTitle(game.packId)}
                                    {:else if $Account.userID === game.user1Id}
                                        vs. {getOpponentUsername(game.user2Id)} · {getPackTitle(game.packId)}
                                    {:else}
                                        vs. {getOpponentUsername(game.user1Id)} · {getPackTitle(game.packId)}
                                    {/if}
                                </span>
                            </Stack>
                            <Cluster gap="s">
                                <Button
                                    variant="accent"
                                    size="sm"
                                    href={game.kind === 'solo' ? '/game?solo=true' : '/game'}
                                    on:mouseenter={() => _GetGame(game.gameId)}
                                >
                                    {game.kind === 'solo' ? 'Resume' : 'Start'}
                                </Button>
                                <Button variant="ghost" size="sm" on:click={() => _DeleteGame(game.gameId)}>Delete</Button>
                            </Cluster>
                        </Cluster>
                    </Paper>
                {/each}
                {#if pastGames.length > 0}
                    <button class="see-all" on:click={() => (showAllGames = !showAllGames)}>
                        {showAllGames ? 'hide past games' : `show ${pastGames.length} past games`}
                    </button>
                {/if}
                {#each pastVisible as game (game.gameId)}
                    <Paper padding="sm">
                        <Cluster gap="m" align="center" justify="between">
                            <span class="game-title past">
                                Finished · {getPackTitle(game.packId)}
                            </span>
                            <Button variant="ghost" size="sm" on:click={() => _DeleteGame(game.gameId)}>Delete</Button>
                        </Cluster>
                    </Paper>
                {/each}
            </Stack>
        </section>
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

    .section-title {
        font-family: var(--font-display);
        font-size: 1.1rem;
        font-weight: 400;
        letter-spacing: -0.01em;
        color: var(--ink);
        margin-bottom: var(--space-3);
    }

    .friend-name { font-weight: 500; color: var(--ink); }
    .friend-meta { font-size: 0.75rem; color: var(--ink-3); }

    .pack-title { font-family: var(--font-display); font-size: 1rem; color: var(--ink); }
    .badge {
        display: inline-block;
        font-size: 0.6em;
        padding: 0.15em 0.55em;
        border-radius: var(--radius-pill);
        background: var(--paper-soft);
        color: var(--ink-2);
        border: 1px solid var(--hairline);
        margin-left: 0.4em;
        vertical-align: middle;
        font-family: var(--font-body);
    }
    .pack-tasks {
        list-style: none; margin: 0; padding: 0;
        display: flex; flex-direction: column;
        gap: var(--space-1);
        font-size: 0.85rem;
        color: var(--ink-2);
    }
    .pack-tasks li { display: flex; gap: var(--space-3); }
    .pack-tasks .i { color: var(--ink-4); min-width: 1.5em; }
    .pack-tasks .more { color: var(--ink-3); font-style: italic; padding-left: calc(1.5em + var(--space-3)); }

    .game-title { font-weight: 500; color: var(--ink); }
    .game-title.past { color: var(--ink-3); }
    .game-sub { font-size: 0.78rem; color: var(--ink-3); }

    .empty { color: var(--ink-2); font-size: 0.9rem; }
    .empty a { color: var(--ink); border-bottom: 1px solid rgba(40, 55, 95, 0.25); text-decoration: none; }

    .see-all {
        background: transparent;
        border: none;
        color: var(--ink-2);
        font-size: 0.78rem;
        cursor: pointer;
        padding: var(--space-2);
        align-self: flex-start;
        text-decoration: underline;
        text-underline-offset: 4px;
    }
    .see-all:hover { color: var(--ink); }
    .see-all:focus-visible { outline: 2px solid var(--accent-ring); outline-offset: 2px; border-radius: var(--radius-sm); }
</style>
```

- [ ] **Step 2: Verify in browser**

With backend running (`make compose`), log in and navigate to `/account`. Confirm:
- Greeting block on atmosphere with name, total bingos, optional solo line, city.
- "Continue playing" paper card if there's an active game.
- Friends as vertical paper rows. "see all N" expands.
- Packs as vertical paper rows showing 3 tasks + "+ N more". "see all N" expands.
- Games as vertical paper rows. "show N past games" expands.
- No horizontal carousels anywhere on the page.

- [ ] **Step 3: Type-check + commit**

Note: `GameModal` import will fail until Task 5.5. To unblock this task, temporarily rebuild `GameModal` here OR continue with this commit and accept a broken import that gets fixed in Phase 5. **Recommended:** stub `bingo-app/src/lib/feature/GameModal.svelte` first with a minimal version that keeps the API:

```svelte
<script lang="ts">
    export let showModal: boolean;
    export let selectedPackID: string;
    export let selectedFriendID: string;
</script>

{#if showModal}
    <p>GameModal — to be redesigned in Phase 5</p>
{/if}
```

Commit the stub as a separate change first, then commit the page:

```bash
cd bingo-app && npm run check
git add bingo-app/src/lib/feature/GameModal.svelte
git commit -m "chore(ui): stub GameModal for Phase 3 unblock"
git add bingo-app/src/routes/account/+page.svelte
git commit -m "feat(ui): redesign /account as vertical stream"
```

---

### Task 3.2: Redesign `/people`

**Files:**
- Modify: `bingo-app/src/routes/people/+page.svelte` (replace whole file)

- [ ] **Step 1: Replace `bingo-app/src/routes/people/+page.svelte`**

```svelte
<script lang="ts">
    import type { PageData } from './$types';
    import Account from '../accountStore';
    import Page from '$lib/ui/Page.svelte';
    import Stack from '$lib/ui/Stack.svelte';
    import Cluster from '$lib/ui/Cluster.svelte';
    import Paper from '$lib/ui/Paper.svelte';
    import Input from '$lib/ui/Input.svelte';
    import Button from '$lib/ui/Button.svelte';
    import GameModal from '$lib/feature/GameModal.svelte';
    import { DeleteFriend, RequestFriend, AcceptFriend } from '../friendRequests';

    export let data: PageData;
    const { users } = data;

    let showModal = false;
    let selectedFriendID = '';
    let selectedPackID = '';
    let query = '';

    $: filtered = (users ?? []).filter((u) => {
        const q = query.trim().toLowerCase();
        if (!q) return true;
        return u.username.toLowerCase().includes(q) || (u.city ?? '').toLowerCase().includes(q);
    });

    function statusOf(userID: string) {
        return $Account?.friends.find((f) => f.userID === userID)?.status ?? 0;
    }
</script>

<svelte:head><title>People · taskbingo</title></svelte:head>

<Page width="normal">
    <Stack gap="l">
        <h1 class="title">People</h1>
        <Input label="Search by username or city" bind:value={query} type="search" />
        <Stack gap="s">
            {#each filtered as user (user.userID)}
                {@const s = statusOf(user.userID)}
                <Paper padding="sm">
                    <Cluster gap="m" align="center" justify="between">
                        <Stack gap="s">
                            <Cluster gap="s" align="baseline">
                                <span class="bingo">{user.bingo}</span>
                                <span class="username">{user.username}</span>
                            </Cluster>
                            <span class="city">{user.city}</span>
                        </Stack>
                        <Cluster gap="s">
                            {#if user.userID === $Account?.userID}
                                <Button variant="ghost" size="sm" href="/account">Account</Button>
                            {:else if s === 3}
                                <Button variant="accent" size="sm" on:click={() => { selectedFriendID = user.userID; selectedPackID = ''; showModal = true; }}>Play</Button>
                                <Button variant="ghost" size="sm" on:click={() => DeleteFriend(user.userID)}>Remove</Button>
                            {:else if s === 1}
                                <Button variant="ghost" size="sm" disabled>Sent</Button>
                                <Button variant="ghost" size="sm" on:click={() => DeleteFriend(user.userID)}>Cancel</Button>
                            {:else if s === 2}
                                <Button variant="accent" size="sm" on:click={() => AcceptFriend(user.userID)}>Accept</Button>
                                <Button variant="ghost" size="sm" on:click={() => DeleteFriend(user.userID)}>Decline</Button>
                            {:else}
                                <Button variant="accent" size="sm" on:click={() => RequestFriend(user.userID, user.username)}>Add friend</Button>
                            {/if}
                        </Cluster>
                    </Cluster>
                </Paper>
            {/each}
            {#if filtered.length === 0}
                <Paper padding="md">
                    <span class="empty">
                        {query ? 'No matches.' : 'No people yet.'}
                    </span>
                </Paper>
            {/if}
        </Stack>
    </Stack>

    <GameModal bind:showModal {selectedFriendID} {selectedPackID} />
</Page>

<style>
    .title {
        font-family: var(--font-display);
        font-size: 1.6rem;
        font-weight: 400;
        letter-spacing: -0.015em;
    }
    .username { font-weight: 500; color: var(--ink); }
    .bingo {
        font-family: var(--font-display);
        font-size: 1.05rem;
        color: var(--ink);
        min-width: 1.5em;
        text-align: right;
    }
    .city { font-size: 0.78rem; color: var(--ink-3); font-style: italic; }
    .empty { color: var(--ink-2); font-size: 0.9rem; font-style: italic; }
</style>
```

- [ ] **Step 2: Verify in browser**

Open `/people`. Confirm vertical list of paper cards, search filter works on username and city, action buttons match friendship state.

- [ ] **Step 3: Type-check + commit**

```bash
cd bingo-app && npm run check
git add bingo-app/src/routes/people/+page.svelte
git commit -m "feat(ui): redesign /people with search + paper rows"
```

---

### Task 3.3: Redesign `/packs`

**Files:**
- Modify: `bingo-app/src/routes/packs/+page.svelte` (replace whole file)

- [ ] **Step 1: Replace `bingo-app/src/routes/packs/+page.svelte`**

```svelte
<script lang="ts">
    import type { PageData } from './$types';
    import Account from '../accountStore';
    import Page from '$lib/ui/Page.svelte';
    import Stack from '$lib/ui/Stack.svelte';
    import Cluster from '$lib/ui/Cluster.svelte';
    import Paper from '$lib/ui/Paper.svelte';
    import Button from '$lib/ui/Button.svelte';
    import TabBar from '$lib/ui/TabBar.svelte';
    import GameModal from '$lib/feature/GameModal.svelte';
    import { _Like, _Rate, _StartSolo } from './+page';

    export let data: PageData;
    const { packs } = data;

    let tab = 'public';
    let showModal = false;
    let selectedPackID = '';
    let selectedFriendID = '';

    $: liked = (packs ?? []).filter((p) => $Account?.likedPacks.some((e) => e.id === p.id));
    $: mine = (packs ?? []).filter((p) => $Account?.packs?.some((e) => e.id === p.id));

    $: visible = tab === 'liked' ? liked : tab === 'mine' ? mine : packs ?? [];

    function play(packID: string) {
        selectedPackID = packID;
        selectedFriendID = '';
        showModal = true;
    }
</script>

<svelte:head><title>Packs · taskbingo</title></svelte:head>

<Page width="normal">
    <Stack gap="l">
        <Cluster gap="m" align="center" justify="between">
            <h1 class="title">Packs</h1>
            <TabBar
                bind:value={tab}
                items={[
                    { value: 'public', label: 'Public' },
                    { value: 'liked', label: 'Liked' },
                    { value: 'mine', label: 'Mine' },
                ]}
                ariaLabel="Pack categories"
            />
        </Cluster>

        <Stack gap="s">
            {#each visible as pack (pack.id)}
                {@const isLiked = $Account?.likedPacks.some((e) => e.id === pack.id) ?? false}
                {@const isRated = $Account?.ratedPacks.some((e) => e === pack.id) ?? false}
                <Paper padding="md">
                    <Stack gap="m">
                        <h3 class="pack-title">{pack.pack.title}</h3>
                        <ol class="tasks">
                            {#each pack.pack.tasks as task, i}
                                <li><span class="num">{i + 1}</span><span>{task}</span></li>
                            {/each}
                        </ol>
                        <Cluster gap="s">
                            {#if isLiked}
                                <Button variant="accent" size="sm" on:click={() => play(pack.id)}>Play</Button>
                            {:else}
                                <Button variant="ghost" size="sm" disabled>Like to play</Button>
                            {/if}
                            <Button variant="ghost" size="sm" on:click={() => _StartSolo(pack.id)}>Solo</Button>
                            <Button variant="ghost" size="sm" on:click={() => _Rate(pack, isRated)} aria-label={isRated ? 'Unrate pack' : 'Rate pack'}>
                                <span aria-hidden="true">{isRated ? '★' : '☆'}</span>
                            </Button>
                            <Button variant="ghost" size="sm" on:click={() => _Like(pack, isLiked)} aria-label={isLiked ? 'Unlike pack' : 'Like pack'}>
                                <span aria-hidden="true">{isLiked ? '♥' : '♡'}</span>
                            </Button>
                        </Cluster>
                    </Stack>
                </Paper>
            {/each}

            {#if visible.length === 0}
                <Paper padding="md">
                    <span class="empty">
                        {#if tab === 'liked'}
                            No liked packs yet — tap the heart on any pack.
                        {:else if tab === 'mine'}
                            <a href="/newpack">Create your first pack →</a>
                        {:else}
                            No packs available.
                        {/if}
                    </span>
                </Paper>
            {/if}
        </Stack>
    </Stack>

    <GameModal bind:showModal {selectedFriendID} {selectedPackID} />
</Page>

<style>
    .title {
        font-family: var(--font-display);
        font-size: 1.6rem;
        font-weight: 400;
        letter-spacing: -0.015em;
    }
    .pack-title {
        font-family: var(--font-display);
        font-size: 1.15rem;
        font-weight: 400;
        letter-spacing: -0.01em;
        color: var(--ink);
        margin: 0;
    }
    .tasks {
        list-style: none;
        margin: 0; padding: 0;
        display: grid;
        grid-template-columns: 1fr 1fr;
        gap: var(--space-1) var(--space-4);
        font-size: 0.85rem;
        color: var(--ink-2);
        line-height: 1.45;
    }
    @media (max-width: 640px) {
        .tasks { grid-template-columns: 1fr; }
    }
    .tasks li { display: flex; gap: var(--space-2); }
    .num { color: var(--ink-4); min-width: 1.5em; }
    .empty { color: var(--ink-2); font-style: italic; font-size: 0.9rem; }
    .empty a { color: var(--ink); border-bottom: 1px solid rgba(40, 55, 95, 0.25); text-decoration: none; }
</style>
```

- [ ] **Step 2: Verify in browser**

Open `/packs`. Confirm tabs `Public · Liked · Mine`, vertical paper feed, two-column tasks layout (collapses to one on mobile), action buttons.

- [ ] **Step 3: Type-check + commit**

```bash
cd bingo-app && npm run check
git add bingo-app/src/routes/packs/+page.svelte
git commit -m "feat(ui): redesign /packs with tabs and vertical feed"
```

---

## Phase 4 — `/newpack` stream editor

Replace 16 separate inputs with a single task editor that builds a chip list.

### Task 4.1: TaskListEditor component

**Files:**
- Create: `bingo-app/src/lib/feature/TaskListEditor.svelte`

- [ ] **Step 1: Write `bingo-app/src/lib/feature/TaskListEditor.svelte`**

```svelte
<script lang="ts">
    import { createEventDispatcher } from 'svelte';
    import Paper from '$lib/ui/Paper.svelte';
    import Cluster from '$lib/ui/Cluster.svelte';

    const MAX = 16;

    export let tasks: string[] = [];

    const dispatch = createEventDispatcher<{ change: string[] }>();

    let draft = '';
    let editingIndex: number | null = null;
    let editingValue = '';
    let dragIndex: number | null = null;

    function commit(next: string[]) {
        tasks = next;
        dispatch('change', next);
    }

    function add() {
        const v = draft.trim();
        if (!v) return;
        if (tasks.length >= MAX) return;
        commit([...tasks, v]);
        draft = '';
    }

    function onKey(e: KeyboardEvent) {
        if (e.key === 'Enter') { e.preventDefault(); add(); }
    }

    function onPaste(e: ClipboardEvent) {
        const text = e.clipboardData?.getData('text') ?? '';
        if (!text.includes('\n')) return;
        e.preventDefault();
        const lines = text.split(/\r?\n/).map((l) => l.trim()).filter(Boolean);
        const room = MAX - tasks.length;
        if (room <= 0) return;
        commit([...tasks, ...lines.slice(0, room)]);
        draft = '';
    }

    function remove(i: number) {
        commit(tasks.filter((_, idx) => idx !== i));
    }

    function startEdit(i: number) {
        editingIndex = i;
        editingValue = tasks[i];
    }
    function saveEdit() {
        if (editingIndex === null) return;
        const v = editingValue.trim();
        if (!v) return;
        commit(tasks.map((t, idx) => (idx === editingIndex ? v : t)));
        editingIndex = null;
    }
    function cancelEdit() { editingIndex = null; }

    function onDragStart(i: number) { dragIndex = i; }
    function onDragOver(e: DragEvent, i: number) {
        if (dragIndex === null || dragIndex === i) return;
        e.preventDefault();
    }
    function onDrop(i: number) {
        if (dragIndex === null || dragIndex === i) { dragIndex = null; return; }
        const next = [...tasks];
        const [moved] = next.splice(dragIndex, 1);
        next.splice(i, 0, moved);
        commit(next);
        dragIndex = null;
    }
</script>

<div class="editor">
    <div class="add-row">
        <div class="input-wrap">
            <input
                type="text"
                placeholder="Add a task and press Enter…"
                bind:value={draft}
                on:keydown={onKey}
                on:paste={onPaste}
                disabled={tasks.length >= MAX}
                aria-label="New task"
            />
            <span class="counter" class:full={tasks.length === MAX}>{tasks.length} / {MAX}</span>
        </div>
        <button type="button" class="add-btn" on:click={add} disabled={tasks.length >= MAX || !draft.trim()}>Add</button>
    </div>

    <ol class="chips">
        {#each tasks as task, i (i + '-' + task)}
            <li
                class="chip"
                draggable="true"
                on:dragstart={() => onDragStart(i)}
                on:dragover={(e) => onDragOver(e, i)}
                on:drop={() => onDrop(i)}
            >
                <span class="grip" aria-hidden="true">⋮⋮</span>
                <span class="num">{i + 1}</span>
                {#if editingIndex === i}
                    <input
                        class="edit-input"
                        type="text"
                        bind:value={editingValue}
                        on:keydown={(e) => { if (e.key === 'Enter') { e.preventDefault(); saveEdit(); } if (e.key === 'Escape') cancelEdit(); }}
                        autofocus
                    />
                    <Cluster gap="s">
                        <button type="button" class="mini accent" on:click={saveEdit}>save</button>
                        <button type="button" class="mini" on:click={cancelEdit}>cancel</button>
                    </Cluster>
                {:else}
                    <span class="body">{task}</span>
                    <Cluster gap="s">
                        <button type="button" class="mini" on:click={() => startEdit(i)} aria-label="Edit task {i + 1}">edit</button>
                        <button type="button" class="mini" on:click={() => remove(i)} aria-label="Remove task {i + 1}">remove</button>
                    </Cluster>
                {/if}
            </li>
        {/each}
    </ol>

    <p class="hint">Tip: paste 16 lines at once to fill the pack.</p>
</div>

<style>
    .editor { display: flex; flex-direction: column; gap: var(--space-3); }

    .add-row { display: flex; gap: var(--space-2); align-items: stretch; }
    .input-wrap { flex: 1; position: relative; }
    .input-wrap input {
        width: 100%;
        background: var(--paper);
        border: 1px solid var(--hairline);
        border-radius: var(--radius-md);
        padding: 0.85rem 4.5rem 0.85rem 1rem;
        color: var(--ink);
        font-size: 0.95rem;
        outline: none;
        min-height: 48px;
        transition: border-color var(--dur-fast) var(--ease), box-shadow var(--dur-fast) var(--ease);
    }
    .input-wrap input:focus {
        border-color: var(--accent-from);
        box-shadow: 0 0 0 3px var(--accent-ring);
    }
    .counter {
        position: absolute;
        right: 0.85rem; top: 50%; transform: translateY(-50%);
        font-size: 0.78rem; color: var(--ink-3);
        font-variant-numeric: tabular-nums;
    }
    .counter.full { color: #c44e4e; }

    .add-btn {
        background: var(--accent);
        color: var(--on-accent);
        border: 1px solid rgba(255, 255, 255, 0.5);
        border-radius: var(--radius-md);
        font-family: var(--font-body);
        font-weight: 500;
        padding: 0 1.2rem;
        min-width: 80px;
        cursor: pointer;
        font-size: 0.85rem;
        min-height: 48px;
    }
    .add-btn:disabled { opacity: 0.45; cursor: not-allowed; }

    .chips {
        list-style: none; margin: 0; padding: 0;
        display: flex; flex-direction: column;
        gap: var(--space-2);
    }
    .chip {
        display: grid;
        grid-template-columns: auto auto 1fr auto;
        gap: var(--space-3);
        align-items: center;
        background: var(--paper);
        border: 1px solid var(--hairline);
        border-radius: var(--radius-md);
        padding: 0.6rem 0.9rem;
    }
    .grip {
        cursor: grab;
        color: var(--ink-4);
        font-size: 1rem;
        user-select: none;
        line-height: 1;
    }
    .num {
        color: var(--ink-3);
        font-size: 0.78rem;
        font-variant-numeric: tabular-nums;
        min-width: 1.5em;
    }
    .body { font-size: 0.9rem; color: var(--ink); }

    .edit-input {
        background: transparent;
        border: 1px solid var(--accent-from);
        border-radius: var(--radius-sm);
        padding: 0.4rem 0.6rem;
        color: var(--ink);
        font-size: 0.9rem;
        outline: none;
    }

    .mini {
        background: transparent;
        border: 1px solid var(--hairline);
        border-radius: var(--radius-pill);
        padding: 0.25rem 0.6rem;
        font-size: 0.7rem;
        color: var(--ink-2);
        cursor: pointer;
        font-family: var(--font-body);
    }
    .mini:hover { color: var(--ink); }
    .mini.accent {
        background: var(--accent);
        color: var(--on-accent);
        border-color: rgba(255,255,255,0.5);
    }

    .hint {
        font-size: 0.78rem;
        color: var(--ink-3);
        font-style: italic;
    }
</style>
```

- [ ] **Step 2: Type-check + commit**

```bash
cd bingo-app && npm run check
git add bingo-app/src/lib/feature/TaskListEditor.svelte
git commit -m "feat(ui): add TaskListEditor stream editor for /newpack"
```

---

### Task 4.2: Wire `/newpack` to TaskListEditor

**Files:**
- Modify: `bingo-app/src/routes/newpack/+page.svelte` (replace whole file)
- Modify: `bingo-app/src/routes/newpack/+page.ts` (replace whole file)

- [ ] **Step 1: Replace `bingo-app/src/routes/newpack/+page.ts`**

```ts
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
```

- [ ] **Step 2: Replace `bingo-app/src/routes/newpack/+page.svelte`**

```svelte
<script lang="ts">
    import { goto } from '$app/navigation';
    import Page from '$lib/ui/Page.svelte';
    import Stack from '$lib/ui/Stack.svelte';
    import Cluster from '$lib/ui/Cluster.svelte';
    import Paper from '$lib/ui/Paper.svelte';
    import Input from '$lib/ui/Input.svelte';
    import Toggle from '$lib/ui/Toggle.svelte';
    import Button from '$lib/ui/Button.svelte';
    import TaskListEditor from '$lib/feature/TaskListEditor.svelte';
    import { _CreatePack } from './+page';

    let name = '';
    let tasks: string[] = [];
    let isPrivate = false;
    let submitting = false;
    let error = '';

    $: canSave = name.trim().length > 0 && tasks.length === 16 && !submitting;

    async function submit(e: Event) {
        e.preventDefault();
        if (!canSave) return;
        submitting = true;
        error = '';
        const status = await _CreatePack({ name: name.trim(), tasks, isPrivate });
        submitting = false;
        if (status === 200) {
            goto('/account');
        } else {
            error = 'Could not save the pack. Try again.';
        }
    }
</script>

<svelte:head><title>New pack · taskbingo</title></svelte:head>

<Page width="normal">
    <Stack gap="l">
        <h1 class="title">New pack</h1>

        <Paper padding="lg">
            <form on:submit={submit}>
                <Stack gap="l">
                    <Input label="Pack name" name="name" bind:value={name} required />

                    <div>
                        <span class="section-label">Tasks</span>
                        <TaskListEditor bind:tasks />
                    </div>

                    <Cluster gap="m" align="center" justify="between">
                        <Toggle label="Private (only you can see and play)" bind:checked={isPrivate} />
                        <Button type="submit" variant="accent" disabled={!canSave}>
                            {submitting ? 'Saving…' : 'Create pack'}
                        </Button>
                    </Cluster>

                    {#if error}<span class="error">{error}</span>{/if}
                </Stack>
            </form>
        </Paper>
    </Stack>
</Page>

<style>
    .title {
        font-family: var(--font-display);
        font-size: 1.6rem;
        font-weight: 400;
        letter-spacing: -0.015em;
    }
    .section-label {
        display: block;
        font-size: 0.65rem;
        text-transform: uppercase;
        letter-spacing: 0.16em;
        color: var(--ink-3);
        margin-bottom: var(--space-3);
    }
    .error { color: #c44e4e; font-size: 0.85rem; }
</style>
```

- [ ] **Step 3: Verify in browser**

Open `/newpack`. Confirm:
- Pack name input.
- Single "Add a task and press Enter" input with counter `0 / 16`.
- Typing a task and pressing Enter adds a chip below; counter increments.
- Pasting 16 lines at once fills all chips.
- Edit / remove buttons on each chip work.
- Drag-handle reorders chips.
- Private toggle works.
- Save button disabled until name + 16 tasks; on success redirects to `/account`.

- [ ] **Step 4: Type-check + commit**

```bash
cd bingo-app && npm run check
git add bingo-app/src/routes/newpack/+page.ts bingo-app/src/routes/newpack/+page.svelte
git commit -m "feat(ui): replace /newpack 16-input form with stream editor"
```

---

## Phase 5 — `/game`

Most complex page. New BingoGrid, Keypad, TaskList, NotesJournal, GameModal, then the page itself.

### Task 5.1: New BingoGrid

**Files:**
- Create: `bingo-app/src/lib/feature/BingoGrid.svelte`

- [ ] **Step 1: Write `bingo-app/src/lib/feature/BingoGrid.svelte`**

```svelte
<script lang="ts">
    /**
     * BingoGrid — 4×4 board. `numbers` is the underlying random ordering
     * shown in cells; `marked` is the array of values the player has marked
     * (matches the existing user1Numbers/user2Numbers shape: 16-length, 0
     * means empty, otherwise the original number).
     */
    export let numbers: number[] = [];
    export let marked: number[] = [];
    export let winningCells: Set<number> = new Set();

    function isMarked(n: number): boolean {
        return marked.includes(n);
    }
</script>

<div class="grid" role="grid" aria-label="Bingo board">
    {#each numbers as n, idx (n + '-' + idx)}
        {@const on = isMarked(n)}
        {@const win = winningCells.has(idx)}
        <div class="cell" class:on class:win role="gridcell" aria-label={on ? `Cell ${idx + 1}, marked ${n}` : `Cell ${idx + 1}, empty`}>
            <span class="num">{on ? n : ''}</span>
        </div>
    {/each}
</div>

<style>
    .grid {
        display: grid;
        grid-template-columns: repeat(4, 1fr);
        gap: 8px;
        width: 100%;
    }

    .cell {
        aspect-ratio: 1;
        border-radius: var(--radius-md);
        display: flex;
        align-items: center;
        justify-content: center;
        background: var(--paper-soft);
        color: var(--ink-4);
        position: relative;
        transition: background var(--dur-fast) var(--ease),
                    color var(--dur-fast) var(--ease),
                    transform var(--dur-base) var(--ease);
    }

    .cell.on {
        background: var(--accent);
        color: var(--on-accent);
        box-shadow: 0 4px 10px -3px var(--accent-shadow),
                    inset 0 1px 0 rgba(255,255,255,0.25);
        animation: pop var(--dur-base) var(--ease);
    }

    .cell.win {
        box-shadow: 0 0 0 2px rgba(148, 170, 207, 0.5),
                    0 6px 18px -4px var(--accent-shadow);
        animation: glow 600ms var(--ease);
    }

    .num {
        font-family: var(--font-display);
        font-weight: 500;
        font-size: clamp(1.1rem, 2.4vw, 1.6rem);
        line-height: 1;
        font-variant-numeric: tabular-nums;
    }

    @keyframes pop {
        from { transform: scale(0.92); }
        to { transform: scale(1); }
    }

    @keyframes glow {
        0% { box-shadow: 0 0 0 0 rgba(148, 170, 207, 0); }
        50% { box-shadow: 0 0 0 6px rgba(148, 170, 207, 0.45); }
        100% { box-shadow: 0 0 0 2px rgba(148, 170, 207, 0.5), 0 6px 18px -4px var(--accent-shadow); }
    }
</style>
```

- [ ] **Step 2: Type-check + commit**

```bash
cd bingo-app && npm run check
git add bingo-app/src/lib/feature/BingoGrid.svelte
git commit -m "feat(game): add BingoGrid with mark + winning-row glow"
```

---

### Task 5.2: New Keypad

**Files:**
- Create: `bingo-app/src/lib/feature/Keypad.svelte`

- [ ] **Step 1: Write `bingo-app/src/lib/feature/Keypad.svelte`**

```svelte
<script lang="ts">
    import { createEventDispatcher } from 'svelte';

    const dispatch = createEventDispatcher<{ submit: number }>();

    let entered = '';

    function press(d: string) {
        if (entered.length >= 2) return;
        entered = entered + d;
    }
    function back() {
        entered = entered.slice(0, -1);
    }
    function submit() {
        const n = Number(entered);
        if (Number.isFinite(n) && n >= 1 && n <= 16) {
            dispatch('submit', n);
        }
        entered = '';
    }

    function onKey(e: KeyboardEvent) {
        if (e.key >= '0' && e.key <= '9') { e.preventDefault(); press(e.key); }
        else if (e.key === 'Backspace') { e.preventDefault(); back(); }
        else if (e.key === 'Enter') { e.preventDefault(); submit(); }
    }
</script>

<div class="pad" on:keydown={onKey} tabindex="0" role="group" aria-label="Number pad">
    <div class="display" aria-live="polite">{entered || '·'}</div>

    <div class="grid">
        {#each ['1','2','3','4','5','6','7','8','9'] as d}
            <button type="button" class="key" on:click={() => press(d)} aria-label={`Number ${d}`}>{d}</button>
        {/each}
        <button type="button" class="key util" on:click={back} aria-label="Backspace">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round" width="20" height="20"><path d="M21 6h-12l-7 6 7 6h12a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2z"/><path d="M18 9l-6 6"/><path d="M12 9l6 6"/></svg>
        </button>
        <button type="button" class="key" on:click={() => press('0')} aria-label="Number 0">0</button>
        <button type="button" class="key submit" on:click={submit} aria-label="Submit">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round" width="20" height="20"><path d="M5 12h14"/><path d="M13 6l6 6-6 6"/></svg>
        </button>
    </div>
</div>

<style>
    .pad {
        display: flex;
        flex-direction: column;
        gap: var(--space-3);
        outline: none;
    }
    .pad:focus-visible { outline: 2px solid var(--accent-ring); outline-offset: 4px; border-radius: var(--radius-lg); }

    .display {
        background: var(--paper-soft);
        border: 1px solid var(--hairline);
        border-radius: var(--radius-md);
        text-align: center;
        font-family: var(--font-display);
        font-size: 1.4rem;
        padding: var(--space-3);
        color: var(--ink);
        font-variant-numeric: tabular-nums;
        min-height: 44px;
    }

    .grid {
        display: grid;
        grid-template-columns: repeat(3, 1fr);
        gap: var(--space-2);
    }

    .key {
        background: var(--paper);
        border: 1px solid var(--hairline);
        border-radius: var(--radius-md);
        padding: var(--space-3);
        font-family: var(--font-display);
        font-size: 1.15rem;
        color: var(--ink);
        cursor: pointer;
        min-height: 48px;
        transition: background var(--dur-fast) var(--ease);
    }
    .key:hover { background: var(--paper-soft); }
    .key:active { transform: translateY(1px); }
    .key:focus-visible { outline: 2px solid var(--accent-ring); outline-offset: 2px; }

    .key.util { color: var(--ink-2); }
    .key.submit {
        background: var(--accent);
        color: var(--on-accent);
        border-color: rgba(255,255,255,0.5);
    }
    .key.submit:hover { filter: brightness(1.04); }
</style>
```

- [ ] **Step 2: Type-check + commit**

```bash
cd bingo-app && npm run check
git add bingo-app/src/lib/feature/Keypad.svelte
git commit -m "feat(game): add new Keypad with paper keys + keyboard input"
```

---

### Task 5.3: New TaskList

**Files:**
- Create: `bingo-app/src/lib/feature/TaskList.svelte`

- [ ] **Step 1: Write `bingo-app/src/lib/feature/TaskList.svelte`**

```svelte
<script lang="ts">
    import { GetTaskPack } from '../../components/TasksCard/packData';

    export let packID = '';
    export let usersTasks: number[] = [];
    export let opponentsTasks: number[] = [];
    export let showOpponent = true;

    $: pack = GetTaskPack(packID);
</script>

<div class="list">
    <h3 class="header">{pack.pack.title}</h3>
    <ol>
        {#each pack.pack.tasks as task, i}
            {@const num = i + 1}
            {@const mineDone = usersTasks.includes(num)}
            {@const oppDone = opponentsTasks.includes(num)}
            <li class:done={mineDone}>
                <span class="num">{num}</span>
                <span class="dots" aria-hidden="true">
                    <span class="dot" class:filled={mineDone} title="You"></span>
                    {#if showOpponent}
                        <span class="dot" class:filled={oppDone} title="Opponent"></span>
                    {/if}
                </span>
                <span class="body">{task}</span>
            </li>
        {/each}
    </ol>
</div>

<style>
    .list { display: flex; flex-direction: column; gap: var(--space-2); }
    .header {
        font-family: var(--font-display);
        font-size: 1.1rem;
        font-weight: 400;
        letter-spacing: -0.01em;
        color: var(--ink);
    }
    ol { list-style: none; margin: 0; padding: 0; display: flex; flex-direction: column; gap: var(--space-1); }
    li {
        display: grid;
        grid-template-columns: 1.5em auto 1fr;
        gap: var(--space-2);
        align-items: center;
        padding: var(--space-2) 0;
        font-size: 0.88rem;
        color: var(--ink);
        border-bottom: 1px solid var(--hairline);
    }
    li:last-child { border-bottom: none; }
    li.done .body { text-decoration: line-through; color: var(--ink-3); }
    .num { color: var(--ink-4); font-variant-numeric: tabular-nums; font-size: 0.78rem; }
    .dots { display: inline-flex; gap: 4px; }
    .dot {
        width: 8px; height: 8px;
        border-radius: var(--radius-pill);
        background: var(--paper-soft);
        border: 1px solid var(--hairline);
    }
    .dot.filled { background: var(--accent); border-color: transparent; }
</style>
```

Note: This component imports `GetTaskPack` from the existing `components/TasksCard/packData.ts`. Phase 6 will move that file or inline its logic.

- [ ] **Step 2: Type-check + commit**

```bash
cd bingo-app && npm run check
git add bingo-app/src/lib/feature/TaskList.svelte
git commit -m "feat(game): add new TaskList with strikethrough + dual dots"
```

---

### Task 5.4: New NotesJournal

**Files:**
- Create: `bingo-app/src/lib/feature/NotesJournal.svelte`

- [ ] **Step 1: Write `bingo-app/src/lib/feature/NotesJournal.svelte`**

```svelte
<script lang="ts">
    import { onMount } from 'svelte';
    import Stack from '$lib/ui/Stack.svelte';
    import Cluster from '$lib/ui/Cluster.svelte';
    import Textarea from '$lib/ui/Textarea.svelte';
    import Button from '$lib/ui/Button.svelte';
    import {
        _ListComments,
        _AddComment,
        _EditComment,
        _DeleteComment,
        type Comment,
    } from '../../routes/game/comments';

    export let gameID: string;

    let comments: Comment[] = [];
    let draft = '';
    let editingID: string | null = null;
    let editingDraft = '';

    onMount(async () => {
        comments = await _ListComments(gameID);
    });

    async function add() {
        if (!draft.trim()) return;
        const c = await _AddComment(gameID, draft.trim());
        if (c) {
            comments = [...comments, c];
            draft = '';
        }
    }
    function startEdit(c: Comment) { editingID = c.id; editingDraft = c.body; }
    async function saveEdit() {
        if (!editingID || !editingDraft.trim()) return;
        const updated = await _EditComment(editingID, editingDraft.trim());
        if (updated) comments = comments.map((c) => (c.id === updated.id ? updated : c));
        editingID = null;
    }
    async function remove(id: string) {
        if (await _DeleteComment(id)) comments = comments.filter((c) => c.id !== id);
    }
    function fmt(ts: string): string {
        try { return new Date(ts).toLocaleString(); } catch { return ts; }
    }
</script>

<div class="journal">
    <h3>Notes</h3>

    <Stack gap="s">
        {#each comments as c (c.id)}
            <div class="entry">
                {#if editingID === c.id}
                    <Textarea label="Edit note" rows={2} bind:value={editingDraft} />
                    <Cluster gap="s">
                        <Button size="sm" variant="accent" on:click={saveEdit}>Save</Button>
                        <Button size="sm" variant="ghost" on:click={() => (editingID = null)}>Cancel</Button>
                    </Cluster>
                {:else}
                    <p class="body">{c.body}</p>
                    <Cluster gap="s" align="baseline">
                        <span class="meta">{fmt(c.createdAt)}</span>
                        <button type="button" class="mini" on:click={() => startEdit(c)}>edit</button>
                        <button type="button" class="mini" on:click={() => remove(c.id)}>delete</button>
                    </Cluster>
                {/if}
            </div>
        {/each}

        {#if comments.length === 0}
            <p class="empty">No notes yet.</p>
        {/if}

        <div class="composer">
            <Textarea label="Leave a note for yourself…" rows={2} bind:value={draft} />
            <Cluster gap="s" justify="end">
                <Button size="sm" variant="accent" on:click={add} disabled={!draft.trim()}>Add</Button>
            </Cluster>
        </div>
    </Stack>
</div>

<style>
    .journal { display: flex; flex-direction: column; gap: var(--space-3); }
    h3 {
        font-family: var(--font-display);
        font-size: 1.1rem;
        font-weight: 400;
        letter-spacing: -0.01em;
        margin: 0;
    }
    .entry {
        padding: var(--space-3) 0;
        border-bottom: 1px solid var(--hairline);
        display: flex; flex-direction: column; gap: var(--space-2);
    }
    .entry:last-of-type { border-bottom: none; }
    .body { white-space: pre-wrap; font-size: 0.9rem; color: var(--ink); margin: 0; }
    .meta { font-size: 0.72rem; color: var(--ink-3); }
    .empty { color: var(--ink-3); font-style: italic; font-size: 0.85rem; }
    .composer { padding-top: var(--space-2); }
    .mini {
        background: transparent;
        border: none;
        color: var(--ink-2);
        font-size: 0.72rem;
        cursor: pointer;
        text-decoration: underline;
        text-underline-offset: 3px;
        padding: 0;
    }
    .mini:hover { color: var(--ink); }
</style>
```

- [ ] **Step 2: Type-check + commit**

```bash
cd bingo-app && npm run check
git add bingo-app/src/lib/feature/NotesJournal.svelte
git commit -m "feat(game): add new NotesJournal with paper styling"
```

---

### Task 5.5: New GameModal

**Files:**
- Modify: `bingo-app/src/lib/feature/GameModal.svelte` (replace stub from Task 3.1)

- [ ] **Step 1: Replace `bingo-app/src/lib/feature/GameModal.svelte`**

```svelte
<script lang="ts">
    import Account from '../../routes/accountStore';
    import Stack from '$lib/ui/Stack.svelte';
    import Cluster from '$lib/ui/Cluster.svelte';
    import Paper from '$lib/ui/Paper.svelte';
    import Select from '$lib/ui/Select.svelte';
    import Button from '$lib/ui/Button.svelte';
    import { CreateGame } from '../../components/GameModal/newGame';

    export let showModal: boolean;
    export let selectedPackID = '';
    export let selectedFriendID = '';

    let dialog: HTMLDialogElement;
    let exists = false;

    function close() {
        selectedPackID = '';
        selectedFriendID = '';
        exists = false;
        dialog?.close();
    }

    function checkIfExists(): boolean {
        for (const g of $Account.games) {
            if ((g.user1Id === selectedFriendID || g.user2Id === selectedFriendID) && g.status !== 3) {
                exists = true;
                return true;
            }
        }
        exists = false;
        return false;
    }

    function create(e: Event) {
        e.preventDefault();
        if (checkIfExists()) return;
        if (!selectedFriendID || !selectedPackID) return;
        CreateGame(selectedFriendID, selectedPackID);
        close();
    }

    $: if (dialog && showModal) dialog.showModal();
</script>

<dialog
    bind:this={dialog}
    on:close={() => (showModal = false)}
    on:click|self={close}
    aria-labelledby="gm-title"
>
    <Paper padding="lg">
        <Stack gap="m">
            <h2 id="gm-title">Create a game</h2>
            <form on:submit={create}>
                <Stack gap="m">
                    <Select
                        label="Friend"
                        bind:value={selectedFriendID}
                        items={$Account.friends.filter((f) => f.status === 3).map((f) => ({ value: f.userID, name: f.username }))}
                        on:change={checkIfExists}
                        required
                    />
                    <Select
                        label="Pack"
                        bind:value={selectedPackID}
                        items={$Account.likedPacks.map((p) => ({ value: p.id, name: p.pack.title }))}
                        required
                    />
                    {#if exists}
                        <span class="warn">An active game with this friend already exists.</span>
                    {/if}
                    <Cluster gap="s" justify="end">
                        <Button type="button" variant="ghost" on:click={close}>Cancel</Button>
                        <Button type="submit" variant="accent" disabled={exists || !selectedFriendID || !selectedPackID}>Create</Button>
                    </Cluster>
                </Stack>
            </form>
        </Stack>
    </Paper>
</dialog>

<style>
    dialog {
        max-width: 26rem;
        min-width: 18rem;
        width: calc(100% - 2rem);
        border: none;
        background: transparent;
        padding: 0;
        margin: auto;
    }
    dialog::backdrop {
        background: rgba(28, 37, 64, 0.35);
        backdrop-filter: blur(8px);
        -webkit-backdrop-filter: blur(8px);
    }
    dialog[open] { animation: zoom var(--dur-base) var(--ease); }

    h2 {
        font-family: var(--font-display);
        font-size: 1.25rem;
        font-weight: 400;
        letter-spacing: -0.012em;
        margin: 0;
    }

    .warn {
        font-size: 0.85rem;
        color: #c44e4e;
    }

    @keyframes zoom {
        from { transform: scale(0.96); opacity: 0; }
        to { transform: scale(1); opacity: 1; }
    }
</style>
```

- [ ] **Step 2: Verify in browser**

Open `/account` (logged in), click `Create new game`. Modal opens with paper card, two paper-styled selects, periwinkle Create button. Close on backdrop click works.

- [ ] **Step 3: Type-check + commit**

```bash
cd bingo-app && npm run check
git add bingo-app/src/lib/feature/GameModal.svelte
git commit -m "feat(game): replace GameModal with paper dialog"
```

---

### Task 5.6: Redesign `/game`

**Files:**
- Modify: `bingo-app/src/routes/game/+page.svelte` (replace whole file)

- [ ] **Step 1: Replace `bingo-app/src/routes/game/+page.svelte`**

```svelte
<script lang="ts">
    import { onDestroy, onMount } from 'svelte';
    import { browser } from '$app/environment';
    import { get, type Unsubscriber } from 'svelte/store';
    import Account from '../accountStore';
    import CurrentGame from '../currentGame';
    import Page from '$lib/ui/Page.svelte';
    import Stack from '$lib/ui/Stack.svelte';
    import Cluster from '$lib/ui/Cluster.svelte';
    import Paper from '$lib/ui/Paper.svelte';
    import Glass from '$lib/ui/Glass.svelte';
    import TabBar from '$lib/ui/TabBar.svelte';
    import Button from '$lib/ui/Button.svelte';
    import BingoGrid from '$lib/feature/BingoGrid.svelte';
    import Keypad from '$lib/feature/Keypad.svelte';
    import TaskList from '$lib/feature/TaskList.svelte';
    import NotesJournal from '$lib/feature/NotesJournal.svelte';
    import {
        _GameHandler,
        _SendUpdate,
        _RedirectOnAccount,
        _PlaceSoloNumber,
        _SoloFinish,
    } from './+page';

    let socket: WebSocket | undefined;
    let closer: () => void = () => {};
    let unsubscribe: Unsubscriber | undefined;

    let account = get(Account);
    let game = get(CurrentGame);
    let isSolo = game?.kind === 'solo';
    let finished = false;
    let usersBingo = 0;
    let opponentsBingo = 0;
    let activeTab: 'mark' | 'tasks' | 'notes' = 'mark';

    onMount(() => {
        if (!game) {
            _RedirectOnAccount();
            return;
        }

        if (!isSolo) {
            const handler = _GameHandler();
            socket = handler.socket;
            closer = handler.closer;
        }

        unsubscribe = CurrentGame.subscribe((value) => {
            usersBingo = value.user1ID === account.userID ? value.user1Bingo : value.user2Bingo;
            opponentsBingo = value.user1ID === account.userID ? value.user2Bingo : value.user1Bingo;
        });

        if (browser) {
            window.addEventListener('beforeunload', () => closer());
        }
    });

    onDestroy(() => {
        unsubscribe?.();
        closer();
    });

    function handleSubmit(e: CustomEvent<number>) {
        const n = e.detail;
        if (n <= 0 || n > 16) return;
        if (isSolo) _PlaceSoloNumber(n);
        else if (socket) _SendUpdate(socket, n, false);
    }

    async function handleFinish() {
        finished = !finished;
        if (isSolo) {
            if (finished) {
                await _SoloFinish();
                _RedirectOnAccount();
            }
            return;
        }
        if (socket) _SendUpdate(socket, 0, finished);
    }

    $: marked = $CurrentGame
        ? ($CurrentGame.user1ID === account.userID ? $CurrentGame.user1Numbers : $CurrentGame.user2Numbers)
        : [];
    $: oppMarked = $CurrentGame
        ? ($CurrentGame.user1ID === account.userID ? $CurrentGame.user2Numbers : $CurrentGame.user1Numbers)
        : [];
</script>

<svelte:head><title>Game · taskbingo</title></svelte:head>

<div class="game-shell">
    {#if !isSolo && $CurrentGame?.status < 3}
        <Page width="narrow">
            <Stack gap="l" align="center">
                <h2 class="waiting">Waiting for the opponent…</h2>
                <span class="waiting-dot"></span>
            </Stack>
        </Page>
    {:else if $CurrentGame}
        <!-- Sticky glass header -->
        <div class="game-header">
            <Glass radius="card" blur="strong">
                <Cluster gap="m" align="center" justify="between">
                    <Cluster gap="l" align="baseline">
                        <Stack gap="s">
                            <span class="who">You</span>
                            <span class="score">{usersBingo}</span>
                        </Stack>
                        {#if !isSolo}
                            <span class="vs">·</span>
                            <Stack gap="s">
                                <span class="who">Opponent</span>
                                <span class="score">{opponentsBingo}</span>
                            </Stack>
                        {/if}
                    </Cluster>
                    <Button variant="accent" size="sm" on:click={handleFinish}>
                        {finished ? 'Resume' : 'Finish'}
                    </Button>
                </Cluster>
            </Glass>
        </div>

        <!-- Body -->
        <Page width="wide">
            <div class="game-body" class:disabled={finished}>
                <div class="grid-col">
                    <Paper padding="md">
                        <BingoGrid numbers={game.numbers} marked={marked} />
                    </Paper>
                </div>
                <div class="side-col">
                    <Paper padding="md">
                        <Stack gap="m">
                            <TabBar
                                bind:value={activeTab}
                                items={[
                                    { value: 'mark', label: 'Mark' },
                                    { value: 'tasks', label: 'Tasks' },
                                    { value: 'notes', label: 'Notes' },
                                ]}
                                ariaLabel="Game panel"
                            />
                            {#if activeTab === 'mark'}
                                <Keypad on:submit={handleSubmit} />
                            {:else if activeTab === 'tasks'}
                                <TaskList
                                    packID={game.packID}
                                    usersTasks={marked}
                                    opponentsTasks={oppMarked}
                                    showOpponent={!isSolo}
                                />
                            {:else}
                                <NotesJournal gameID={game.gameID} />
                            {/if}
                        </Stack>
                    </Paper>
                </div>
            </div>
        </Page>
    {/if}
</div>

<style>
    .game-shell { width: 100%; }

    .waiting {
        font-family: var(--font-display);
        font-style: italic;
        font-weight: 400;
        font-size: 1.4rem;
        color: var(--ink-2);
    }
    .waiting-dot {
        width: 8px; height: 8px; border-radius: var(--radius-pill);
        background: var(--accent);
        animation: pulse 1.4s var(--ease) infinite;
    }
    @keyframes pulse {
        0%, 100% { opacity: 0.4; transform: scale(1); }
        50% { opacity: 1; transform: scale(1.4); }
    }

    .game-header {
        position: sticky;
        top: 0;
        z-index: 5;
        padding: var(--space-3) var(--space-5);
        max-width: 80rem;
        margin: 0 auto;
    }

    .game-body {
        display: grid;
        grid-template-columns: 1.4fr 1fr;
        gap: var(--space-5);
        align-items: start;
    }
    .game-body.disabled {
        pointer-events: none;
        opacity: 0.55;
    }

    .side-col {
        position: sticky;
        top: calc(var(--space-3) + 76px);
    }

    @media (max-width: 1023px) {
        .game-body { grid-template-columns: 1fr; }
        .side-col { position: static; }
    }

    .who { font-size: 0.6rem; letter-spacing: 0.16em; text-transform: uppercase; color: var(--ink-3); }
    .score { font-family: var(--font-display); font-size: 1.4rem; line-height: 1; color: var(--ink); font-variant-numeric: tabular-nums; }
    .vs { color: var(--ink-4); }
</style>
```

- [ ] **Step 2: Verify in browser**

Start a game (solo from `/packs` for the simplest test). Open `/game?solo=true`. Confirm:
- Sticky glass header at top with score + Finish button.
- Bingo grid centered in left paper card; tap a number on the keypad → cell marks with pop animation.
- TabBar switches between Mark / Tasks / Notes.
- Below 1024px width the layout stacks: grid on top, panel below.
- Finish button works (solo: redirects to /account; duo: sends WS update).

- [ ] **Step 3: Type-check + commit**

```bash
cd bingo-app && npm run check
git add bingo-app/src/routes/game/+page.svelte
git commit -m "feat(game): redesign /game with sticky header + tabbed side panel"
```

---

## Phase 6 — Cleanup

Remove flowbite, drop unused fonts, delete the old `components/` directory.

### Task 6.1: Move shared helpers out of `components/`

**Files:**
- Modify: `bingo-app/src/lib/feature/TaskList.svelte` (update import)
- Modify: `bingo-app/src/lib/feature/GameModal.svelte` (update import)
- Move: `bingo-app/src/components/TasksCard/packData.ts` → `bingo-app/src/lib/data/packs.ts`
- Move: `bingo-app/src/components/GameModal/newGame.ts` → `bingo-app/src/lib/data/newGame.ts`

- [ ] **Step 1: Move the helpers**

```bash
mkdir -p bingo-app/src/lib/data
mv bingo-app/src/components/TasksCard/packData.ts bingo-app/src/lib/data/packs.ts
mv bingo-app/src/components/GameModal/newGame.ts bingo-app/src/lib/data/newGame.ts
```

- [ ] **Step 2: Update imports in `bingo-app/src/lib/feature/TaskList.svelte`**

Change:
```ts
import { GetTaskPack } from '../../components/TasksCard/packData';
```
to:
```ts
import { GetTaskPack } from '$lib/data/packs';
```

- [ ] **Step 3: Update imports in `bingo-app/src/lib/feature/GameModal.svelte`**

Change:
```ts
import { CreateGame } from '../../components/GameModal/newGame';
```
to:
```ts
import { CreateGame } from '$lib/data/newGame';
```

- [ ] **Step 4: Type-check + verify**

```bash
cd bingo-app && npm run check
```

Expected: 0 errors.

- [ ] **Step 5: Commit**

```bash
git add bingo-app/src/lib/data bingo-app/src/lib/feature/TaskList.svelte bingo-app/src/lib/feature/GameModal.svelte
git rm -r bingo-app/src/components/TasksCard bingo-app/src/components/GameModal
git commit -m "refactor(ui): move pack/newGame helpers to lib/data"
```

---

### Task 6.2: Delete old components directory

**Files:**
- Delete: `bingo-app/src/components/`

After Phase 5, `Header`, `Grid`, `Keypad`, `Comments`, `TasksCard` should no longer be imported anywhere. Verify and delete.

- [ ] **Step 1: Verify no imports remain**

```bash
cd bingo-app
grep -rn "from .*components/Header" src/ || echo "no Header imports"
grep -rn "from .*components/Grid" src/ || echo "no Grid imports"
grep -rn "from .*components/Keypad" src/ || echo "no Keypad imports"
grep -rn "from .*components/Comments" src/ || echo "no Comments imports"
grep -rn "from .*components/TasksCard" src/ || echo "no TasksCard imports"
```

Expected: every line says "no <X> imports". If any actual import is shown, fix it before continuing.

- [ ] **Step 2: Remove the directory**

```bash
git rm -r bingo-app/src/components
```

- [ ] **Step 3: Type-check + build**

```bash
cd bingo-app
npm run check
npm run build
```

Expected: 0 errors. Build succeeds.

- [ ] **Step 4: Commit**

```bash
git commit -m "chore(ui): drop legacy components/ directory"
```

---

### Task 6.3: Remove flowbite dependencies

**Files:**
- Modify: `bingo-app/tailwind.config.cjs`
- Modify: `bingo-app/package.json`
- Run: `npm install` to refresh `package-lock.json`

- [ ] **Step 1: Edit `bingo-app/tailwind.config.cjs`**

Replace the entire file with:

```js
const config = {
    content: ['./src/**/*.{html,js,svelte,ts}'],
    theme: {
        extend: {}
    },
    plugins: [],
    darkMode: 'class',
};

module.exports = config;
```

- [ ] **Step 2: Edit `bingo-app/package.json`**

Remove these two lines from `devDependencies`:

```json
"flowbite": "^1.5.4",
"flowbite-svelte": "^0.28.4",
```

- [ ] **Step 3: Refresh dependencies**

```bash
cd bingo-app
npm install
```

Expected: `package-lock.json` updates; `node_modules/flowbite*` is removed.

- [ ] **Step 4: Type-check + build**

```bash
npm run check
npm run build
```

Expected: 0 errors. Build succeeds. If a stray flowbite import was missed, fix it now.

- [ ] **Step 5: Commit**

```bash
git add bingo-app/tailwind.config.cjs bingo-app/package.json bingo-app/package-lock.json
git commit -m "chore(deps): drop flowbite and flowbite-svelte"
```

---

### Task 6.4: Drop unused Google Fonts and final audit

**Files:**
- Modify: `bingo-app/static/fonts.css` — delete content (Unbounded font-face is unused)

- [ ] **Step 1: Replace `bingo-app/static/fonts.css` with an empty file (or delete it)**

```bash
cd bingo-app
rm static/fonts.css
```

If anything references it, remove the reference. Search:
```bash
grep -rn "fonts.css" src/ static/
```

Expected: no matches.

- [ ] **Step 2: Final visual + a11y audit**

Run dev:
```bash
npm run dev
```

For each route below, in DevTools, set the device emulation to iPhone 13 (390×844) and again to a desktop viewport (1440×900):

- `/`
- `/about`
- `/login`
- `/register`
- `/account`
- `/people`
- `/packs`
- `/newpack`
- `/game?solo=true` (start one from `/packs` first)

Confirm for each:
- Atmosphere visible underneath, no white flashes.
- No flowbite styles leaking.
- Focus rings appear on Tab navigation.
- All buttons ≥44px tap target on mobile.
- No horizontal scrollbars (the carousel pattern is gone).
- Reduced motion: enable in DevTools Rendering panel ("Emulate CSS prefers-reduced-motion: reduce") and verify animations stop.

- [ ] **Step 3: Final commit**

```bash
git add bingo-app/static
git commit -m "chore(ui): drop unused Google Fonts shim"
```

---

## Self-Review Checklist

- [ ] **Spec coverage:** Every spec section is mapped:
    - Atmosphere → Task 1.3
    - Paper / Glass → Tasks 1.4, 1.5
    - Tokens / typography → Tasks 1.1, 1.2
    - Form primitives (Button/Input/Textarea/Toggle/Select/TabBar) → Tasks 1.7–1.12
    - Layout primitives (Page/Stack/Cluster) → Task 1.6
    - Layering rule (1–2 glass per screen) → enforced visually in pages, glass used only in Header (scroll), GameModal backdrop, /game sticky header
    - Per-page UX changes (`/`, `/about`, `/login`, `/register`, `/account`, `/people`, `/packs`, `/newpack`, `/game`) → Tasks 2.2–2.5, 3.1–3.3, 4.2, 5.6
    - Feature components (BingoGrid, Keypad, TaskList, NotesJournal, GameModal, Header, TaskListEditor) → Tasks 5.1–5.5, 2.1, 4.1
    - Motion (focus, hover, cell pop, win glow, reduced-motion) → encoded in tokens.css + per-component animations
    - Empty states → /account, /packs, /people, NotesJournal "no notes yet"
    - Loading states → forms have `submitting` flag with disabled state and label change
    - Error states → /login, /register, /newpack inline error spans
    - Accessibility (focus, contrast, keyboard, reduced motion, skip-link) → tokens + per-component
    - Responsive breakpoints (≥1024 / 768–1023 / <768) → /game, /, /packs, header
    - Phasing → 6 phases, sequential, each ends with working app

- [ ] **No placeholders:** No "TBD", "TODO", or "implement later" in any task. Every step contains the actual code or command.

- [ ] **Type consistency:**
    - `Button` props: `variant`, `size`, `href`, `type`, `disabled`, `fullWidth` — used the same way across pages
    - `Input` props: `label`, `value`, `type`, `name`, `placeholder`, `required`, `disabled`, `error`, `hint` — consistent
    - `Paper` `padding` enum: `'sm' | 'md' | 'lg'` — only these values used
    - `Stack` `gap`: `'s' | 'm' | 'l' | 'xl'`; `Cluster` `gap`: `'s' | 'm' | 'l'` (intentional: Cluster has no `xl` because horizontal clusters rarely need it)
    - `tasks` array shape on `/newpack`: `string[]`, length-checked === 16 in `_CreatePack`
    - `winningCells` on `BingoGrid`: `Set<number>` — defaulted to empty; the page does not yet detect winning cells (out of scope; cells still show `.on` styling)
    - `GameModal` API: `showModal`, `selectedPackID`, `selectedFriendID` — same as the original signature so existing call sites in `/account`, `/people`, `/packs` work

---

## Execution Handoff

Plan complete and saved to `docs/superpowers/plans/2026-05-08-aurora-calm-redesign.md`. Two execution options:

**1. Subagent-Driven (recommended)** — fresh subagent per task, review between tasks, fast iteration.

**2. Inline Execution** — execute tasks in this session using executing-plans, batch execution with checkpoints.

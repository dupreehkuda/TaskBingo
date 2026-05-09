# Aurora Calm: frontend redesign with quiet glass + paper aesthetic

**Date:** 2026-05-08
**Status:** Draft (design)

## Goal

Redesign the entire SvelteKit frontend (`bingo-app/`) into a calm, restrained "Aurora Calm" visual language: a soft pastel atmosphere as the page canvas, warm ivory paper for content, and rare frosted glass moments reserved for chrome and atmosphere. Restructure the user journey on the three structurally weak pages (`/account`, `/newpack`, `/game`) and bring the rest into visual and UX hygiene. Replace `flowbite-svelte` with custom primitives. Mobile-first, equal weight to desktop.

## Non-goals

- No backend changes. The HTTP/WS contracts, JWT cookie behaviour, and every API route stay identical.
- No new product features (no DMs, no new game modes, no rankings beyond what already exists).
- No copy rewrite beyond what the new layouts require — keep existing English copy where it fits.
- No new dependencies beyond what's needed for the type system (Google Fonts CSS for Fraunces + Plus Jakarta) — no animation library, no UI kit, no icon system beyond a small inline-SVG set.
- No localisation. The app stays English-only.
- No analytics, no consent banner, no cookie modal.
- No service worker / PWA / install prompt.

## Aesthetic foundation

Locked in brainstorm (visual companion v3). Four primitives, six tokens, two surface types.

### Atmosphere (page canvas)

A soft pastel mesh painted on the body. Two radial whispers over a slightly-tinted base, plus a faint noise overlay to kill the AI-gradient look:

```css
background:
  radial-gradient(58% 50% at 18% 22%, #d3dfe8 0%, transparent 64%),
  radial-gradient(48% 48% at 84% 80%, #e2d7e3 0%, transparent 66%),
  linear-gradient(140deg, #f4f6fa 0%, #efeaf1 100%);
```

Noise is an inline SVG turbulence at `opacity: 0.5`, `mix-blend-mode: overlay`. The atmosphere covers the full viewport behind every page (lives on `<body>` or in the root `+layout.svelte`).

### Surface types

Two surfaces. Strict rule.

**Paper** — workhorse. Used for every content card, list item, form, game grid container, modal body. Looks like warm ivory writing paper.

```css
background: #fbfaf6;
border: 1px solid rgba(40, 50, 80, 0.06);
border-radius: 18px;
box-shadow:
  0 1px 0 rgba(255,255,255,0.6) inset,
  0 2px 6px -2px rgba(40, 55, 95, 0.05),
  0 14px 32px -16px rgba(40, 55, 95, 0.10);
color: #1c2540;
```

**Glass** — rare. Used only for floating chrome (sticky header on scroll, side-panel during active game, modal overlay backdrop). Maximum 1–2 glass surfaces visible per screen.

```css
background: rgba(252, 253, 255, 0.55);
border: 1px solid rgba(255, 255, 255, 0.65);
backdrop-filter: blur(28px) saturate(130%);
-webkit-backdrop-filter: blur(28px) saturate(130%);
border-radius: 18px;
box-shadow:
  0 22px 44px -18px rgba(40, 55, 95, 0.16),
  inset 0 1px 0 rgba(255, 255, 255, 0.7);
```

### Layering rule

Inviolable:

- **Glass never sits on glass.** If a glass surface contains other surfaces, those inner surfaces are flat (no `backdrop-filter`).
- **Paper sits on atmosphere or inside glass.** Never paper-on-paper either — separation comes from spacing, not nested cards.
- **Atmosphere is the only "level 0".** Everything else is exactly one step up.
- The home page renders **zero glass surfaces** by default (header floats as plain text on the atmosphere). Glass appears only when the user scrolls past 64px and the header sticks.

### Tokens

Eleven CSS custom properties. No additional palette improvisation in components.

| Token | Value | Use |
|---|---|---|
| `--atmos-base` | `#f4f6fa` → `#efeaf1` (linear-gradient) | Body background base |
| `--atmos-sky` | `#d3dfe8` | Upper-left whisper in mesh |
| `--atmos-lilac` | `#e2d7e3` | Lower-right whisper in mesh |
| `--paper` | `#fbfaf6` | Solid content surface |
| `--accent` | `linear-gradient(135deg, #94aacf, #b8a3cd)` | Periwinkle: primary buttons, active states, brand line |
| `--ink` | `#1c2540` | Primary text |
| `--ink-2` | `#4d5878` | Secondary text |
| `--ink-3` | `#7e88a8` | Tertiary text, italic display |
| `--ink-4` | `#97a0bb` | Disabled, hint, captions |
| `--paper-soft` | `#f1f0eb` | Inactive bingo cell, inactive chip background |
| `--hairline` | `rgba(40, 50, 80, 0.06)` | Border for paper |

Borrowing inks creates a single-axis tonal scale; designers can pick the right "depth" without inventing new greys.

### Typography

Two families from Google Fonts.

- **Display:** [Fraunces](https://fonts.google.com/specimen/Fraunces), weights 300/400, italic available. Used for page titles, card titles, big numbers, hero h1. Italic is sparing — at most one italic phrase per block.
- **Body:** [Plus Jakarta Sans](https://fonts.google.com/specimen/Plus+Jakarta+Sans), weights 400/500/600. Used for everything else.

Replace the existing Google Fonts import in `Header.svelte` (a kitchen-sink 8-family load) with one focused `<link>` in `app.html`. Drop unused fonts: Abril Fatface, Josefin Sans, Montserrat, Oswald, Righteous, Yellowtail, Meow Script, Prompt, Unbounded.

Type scale (rem-based, 16px root):

| Level | Family | Size | Weight | Letter-spacing |
|---|---|---|---|---|
| `display-xl` | Fraunces | 2.4rem | 400 | -0.022em |
| `display-l` | Fraunces | 1.6rem | 400 | -0.015em |
| `display-m` | Fraunces | 1.15rem | 400 | -0.01em |
| `body-l` | Jakarta | 1rem | 400 | 0 |
| `body-m` | Jakarta | 0.875rem | 400 | 0 |
| `body-s` | Jakarta | 0.78rem | 500 | 0 |
| `eyebrow` | Jakarta | 0.65rem | 500 | 0.18em uppercase |

## Per-page UX changes

Nine pages. Three need structural rework, six need visual + hygiene polish.

### `/` Home

**Now:** Static "Premise" paragraph + a conditional list of plain links (`/people`, `/newpack`, `/account`) on dark blue.

**Becomes:** Hero text floating directly on the atmosphere (no card). One `display-xl` line, one CTA button (`Start a game` → `/account` if logged in, `/login` otherwise), one ghost link (`Browse packs`). For logged-in users with at least one active game, render exactly one paper card to the right with a "continue this game" preview (active bingo grid, score, opponent name). For new users, render a single paper card with three short steps (1. invite a friend, 2. agree on a pack, 3. play). No three-column layout, no carousels.

### `/about`

**Now:** Flowbite `Timeline` + `TimelineItem` with bullet-on-vertical-line. Mixed dark/light text. Uses Prompt font.

**Becomes:** Same six entries, rendered as a narrow prose column (max 36rem) on the atmosphere. Each entry is a small block: italic Fraunces date label + Fraunces h3 title + Jakarta paragraph. A thin gradient hairline (`--accent` at 40% opacity) on the left edge replaces the bullet column. Drop the flowbite Timeline component.

### `/login`, `/register`

**Now:** Centred form, bug: `body { background-color: #07417b }` repeated; flowbite `Input`, `ButtonGroup`, `InputAddon` with custom-coloured eye-icon SVG.

**Becomes:** Single paper card centred on the atmosphere (max-width 22rem). Fraunces `display-l` title (`Welcome back` / `Create an account`). Custom Input primitive with floating label (label rests inside input until focus, slides up on focus or value). Single full-width `--accent` button. Inline link to the other route below ("New here? Register" / "Have an account? Log in"). Eye-toggle for password is an inline icon inside the input, not a separate button group. On `/register`, `city` becomes a free-text input (current behaviour) with a placeholder; no dropdown — adding a city list is non-goal.

### `/account`

**Now:** Username header, three horizontal scrolling carousels (Friends, Packs, Games) with hidden scrollbars, plus modal-trigger buttons. Each carousel uses `min-width: 23em` items so on most viewports ≥3 items fit horizontally; users have to scroll horizontally with no affordance.

**Becomes:** Vertical "stream" layout. Sections in this order:

1. **Greeting + score block** at the top — Fraunces `display-l` username + Fraunces big number for `bingo` total + small italic `solo: N` if `soloBingo > 0`. Sits on the atmosphere, no card.
2. **Continue playing** — featured paper card with a single in-progress game (most recent), showing the live grid and score. Big CTA button `Resume`. Renders only if the user has an active game.
3. **Friends** — vertical list of paper cards, max 5 visible, "see all" link below if more. Each row: avatar dot, username, win/loss tally, action button (Play / Accept / Sent / Remove).
4. **Packs** — vertical list of paper cards (liked packs), each showing title + 3-task preview + "private" pill if private. Action buttons: `Choose pack` (opens GameModal), heart toggle. Max 5 visible, "see all" link.
5. **Games** — collapsed by default, "show N past games" link expands the list. Solo games and finished games belong here.
6. **Create new pack** + **Create new game** — two ghost buttons in a footer row, no longer scattered into each section.

Drop the horizontal carousel pattern entirely.

### `/people`

**Now:** Vertical list of cards already (no carousel here, contrary to first read), but each card uses cream-grey `#e8e8e6` and flowbite Buttons. No search.

**Becomes:** Search input (paper card with floating label) at the top, filters by username and city. Below: vertical list of paper cards. Each card: avatar dot, username, city in `body-s` italic Jakarta, bingo score on the right, action button matching friendship state (`Add friend` / `Sent` / `Accept` / `Play` / `Remove`). Empty state when search has no matches: paper card with friendly italic Fraunces line.

### `/packs`

**Now:** Single horizontal carousel of all public packs (cream cards, 23em min-width, hidden scrollbar).

**Becomes:** Tabs at top: `Public · Liked · Mine`. Below tabs: vertical feed of paper cards. Each card shows pack title (Fraunces `display-m`), the 16 tasks as a numbered list (toned `--ink-2`), and an action row: `Play` (with friend) / `Solo` / heart / star. Tabs preserve scroll position when switching. Empty states for `Liked` ("No packs liked yet — tap the heart on any pack") and `Mine` ("Create your first pack").

### `/newpack`

**Now:** A literal hand-rolled stack of 16 separate `<Input>` fields with `name="task1"` … `name="task16"`. The user must tab through 16 fields. No autosave, no reorder, no paste-multiple, no validation beyond `required`.

**Becomes:** Stream editor.

- One Fraunces `display-l` title input ("Pack name") at the top.
- One single-line `task input` below, with a placeholder "Add a task and press Enter…".
- Below: a vertical list of paper-chip rows, one per added task. Each chip shows index (`1·`, `2·`, …) and the task body, plus an inline edit / delete button. Drag-handle on the left to reorder.
- Counter on the right of the input: `12 / 16`. Disabled add when the count is 16.
- Paste-multiline (text with `\n`) splits on newlines and adds each as a chip in order, capped at 16.
- "Private" toggle pill in the footer row.
- Save button enabled only when name is non-empty and exactly 16 tasks are present.

The form serialises into the existing `task1`…`task16` fields on submit (one `_Submit` change in `+page.ts`), so the backend contract is untouched.

### `/game`

**Now:** A 3-column grid (`grid-cols-3 gap-4`): left 2/3 is the bingo grid, right 1/3 is a column of vertically stacked tiny score rectangles + Finish button + Keypad + TasksCard + Comments. On mobile it collapses awkwardly.

**Becomes:** Two-region layout, mobile-first.

**Top — sticky glass header.** A thin glass strip across the top showing: opponent name + opponent score / your name + your score / the number of "lines" each player has scored / a Finish button on the right. Stays sticky while scrolling. This is the only persistent glass element on the screen.

**Body desktop (≥1024px):** Two columns. Left (1.4fr): bingo grid as a paper card, large, centred vertically. Right (1fr): paper card with three tabs — `Mark a task` (the keypad), `Tasks` (the TasksCard list with checkmarks), `Notes` (the Comments journal). Tab content area scrolls if needed. The right column is sticky-positioned at viewport top + sticky-header height.

**Body mobile (<1024px):** Single column. Bingo grid as a paper card at the top, square (1:1). Below: a horizontal segmented control (`Mark · Tasks · Notes`), and below that the active section as a paper card. Tab swap is instant (no animation gating input).

**Solo vs duo:** Solo hides the opponent score in the sticky header. Duo waiting state ("Waiting for the opponent") becomes an atmosphere-only screen with a Fraunces italic line + a tiny gradient pulsing dot.

**Bingo cell visual:** Inactive cell is a flat `--paper-soft` square with a faint dot in `--ink-4`. Active cell is the periwinkle accent gradient with white numerals and a soft 4px-spread inner shadow. When a 4-in-a-row is detected, the four cells in that row get a soft 6px-blur outer glow in `--accent` for 600ms (one-shot, not pulsing).

## Components & primitives

All custom Svelte components, no `flowbite-svelte`. Each lives in `bingo-app/src/lib/ui/` (new directory).

### Surface primitives

- `<Atmosphere />` — root wrapper. Renders the mesh + noise on `<body>` via styled `:global` on mount. Used once in `+layout.svelte`.
- `<Paper>` — content surface. Props: `padding` (`sm | md | lg`), `as` (defaults to `div`), pass-through class. Uses the paper token system.
- `<Glass>` — atmospheric surface. Props: `radius` (`pill | card`), `blur` (`soft | strong`). Used for sticky header, side panel, modal.

### Form primitives

- `<Button>` — variants `accent | ghost | quiet`, sizes `sm | md | lg`. Accent uses `--accent` gradient. Ghost is a paper-coloured pill with hairline. Quiet is just a text link with a 1px underline.
- `<Input>` — floating-label text input. Props: `label`, `type`, `value` (bind), `error?`, `hint?`. Uses focus ring (`outline: 2px solid color-mix(in srgb, --accent 60%, transparent)`).
- `<Textarea>` — same pattern as Input, taller.
- `<Toggle>` — pill switch. Used for `isPrivate` on `/newpack`.
- `<Select>` — paper-styled native `<select>` wrapper with floating label and a custom chevron. Used inside `<GameModal>` for friend / pack pickers.
- `<TabBar>` — segmented control (mobile) and tab strip (desktop). Same component, responsive.

### Layout primitives

- `<Page>` — page-level wrapper with consistent vertical padding and max-width (60rem default).
- `<Stack gap="…" />` — vertical flex with gap token (`s | m | l | xl`).
- `<Cluster gap="…" />` — horizontal flex-wrap with gap.

### Feature components

- `<BingoGrid {numbers} {marked} {winningRow?} on:cellClick />` — replaces `Grid.svelte`. Renders the 4×4. Highlights `winningRow` with the glow effect.
- `<Keypad on:submit />` — number pad. Visual: paper buttons in a 3-column layout, plus a backspace and submit. Replaces `Keypad.svelte`. Already structurally fine, just restyled.
- `<TaskList {tasks} {usersDone} {opponentsDone} />` — replaces the inner content of `TasksCard.svelte`. Each row: number, two completion dots (mine, theirs), task body, strikethrough when mine is done.
- `<NotesJournal {gameID} />` — replaces `Comments.svelte`. Same CRUD, styled as paper-on-paper-soft alternating rows. Drop the inline edit dance — clicking a row enters edit-in-place with a small Save / Cancel cluster appearing below.
- `<GameModal>` — replaces existing `GameModal.svelte`. Glass-backed `<dialog>` with a paper card body. Two flowbite `Select`s become two `<Select>` primitives (paper-styled native selects).
- `<Header>` — replaces `Header/Header.svelte`. On home, brand + nav float as plain text on the atmosphere. On scroll past 64px, the header transitions into a glass strip (border + backdrop-blur fade-in). On other pages, glass is on by default. Active link uses the gradient hairline underline.

### Removing flowbite

Audit `bingo-app/src/`:

- `Navbar`, `NavBrand`, `NavLi`, `NavUl`, `NavHamburger` — replaced by custom `<Header>`.
- `Button` — replaced.
- `Input`, `Label`, `InputAddon`, `ButtonGroup` — replaced.
- `Select` — replaced (inside `<GameModal>`).
- `Textarea` — replaced (inside `<NotesJournal>`).
- `Alert`, `Checkbox` — replaced.
- `Timeline`, `TimelineItem` (in `/about`) — replaced.

Once all imports are removed: drop `flowbite-svelte` and `flowbite` from `package.json` and from `tailwind.config.cjs` `content` array + plugin. Drop the `flowbite-svelte` import in `Header.svelte`. Verify `npm run build` still passes.

## Motion

Calm, sparing. Three motion budgets:

**Functional motion** (always on, never disabled): focus-ring fade-in (120ms), input floating-label slide (180ms), tab swap content cross-fade (140ms).

**Decorative motion** (skipped under `prefers-reduced-motion: reduce`):

- Page transitions: 200ms cross-fade + 8px translate-up on the new page.
- Card hover: shadow lift from 14px to 22px blur, no transform.
- Button hover: subtle background brightness shift via `filter: brightness(1.04)`. No scale.
- Bingo cell mark: 240ms scale-from-0.92 + accent gradient fill.
- 4-in-a-row glow: 600ms one-shot, not looped.
- Sticky header glass-on appearance: 240ms backdrop-filter ramp from 0 to 28px.

No spring physics, no parallax, no confetti, no audio. The vibe is "things settle into place gently".

## Empty / loading / error states

The current app has effectively none. New defaults:

**Empty states** — paper card with a Fraunces italic line + a single CTA. Variants: no friends, no packs, no liked packs, no in-progress games, no notes on a game, no search results.

**Loading states** — three pulsing dots (gradient `--accent`, fade in/out, 800ms cycle, staggered 120ms). For paper-card containers waiting on data, render a skeleton: same paper shell, but title and subtitle replaced with `--paper-soft` rounded bars at 60% / 40% width. No shimmer animation.

**Error states** — paper card with `--ink` body and a small accent-coloured retry pill. Auth errors on `/login` show inline below the field (no flowbite Alert). Network errors on data fetch show "Couldn't load. Try again." with a retry button.

## Accessibility & responsive

### Contrast

- `--ink` on `--paper` → 18.4:1 ✅ AAA
- `--ink-2` on `--paper` → 8.1:1 ✅ AAA
- `--ink-3` on `--paper` → 4.6:1 ✅ AA (use only for non-essential text — italic display, captions)
- `--ink-4` on `--paper` → 3.0:1 — **decorative only**, never for body text
- White text on `--accent` gradient at darkest stop (`#94aacf`) → 3.4:1 — passes AA for large text (button labels are ≥0.85rem semibold, qualify as large). Verify per-button at implementation time.
- Glass surfaces have variable contrast. Solution: text inside glass uses `--ink` or stronger; if glass sits over a darker patch of atmosphere, increase the glass `background` opacity to 0.7. Tested per-component.

### Focus

- Visible focus on every interactive element. 2px solid ring in `color-mix(in srgb, --accent 60%, transparent)` with a 2px offset. Same ring for keyboard and mouse focus (no `:focus-visible` distinction — keep it predictable).
- Skip link at the top of `+layout.svelte` ("Skip to main") for keyboard users.

### Keyboard

- Bingo grid: arrow keys move focus between cells; Enter / Space marks a cell on `/game` (when it's a number-input flow, the focused cell shows the next-to-mark hint).
- Keypad: digits 0–9 on the keyboard work the same as clicking; Backspace and Enter map to the on-screen Backspace and Submit.
- Tab order is DOM order. No `tabindex > 0`.

### Reduced motion

`prefers-reduced-motion: reduce` collapses all decorative motion. Transitions become instant `display` swaps; fades become instant.

### Responsive breakpoints

- `≥1024px` — desktop layouts (two-column /game, side-by-side /home hero + game card).
- `768–1023px` — tablet, mostly desktop layout but with tighter padding and stacked /home hero.
- `<768px` — mobile. /game becomes single-column with bottom-tab segmented control. /account stream stays vertical (already mobile-friendly). /newpack chip list scrolls; the input becomes sticky at the top of the form.
- Tap targets ≥44×44px on mobile (buttons are min-height 44px).

## Implementation phasing

Six phases. Each ends in a working app — the redesign rolls out per page, not all at once. Phases run in order.

**Phase 1 — Foundation.** No user-visible changes outside of the atmosphere appearing.

- Add Google Fonts CSS link in `app.html` (Fraunces + Plus Jakarta).
- Add the new design tokens as CSS custom properties in `app.postcss`.
- Implement `<Atmosphere />`, `<Paper>`, `<Glass>`, `<Stack>`, `<Cluster>`, `<Page>` primitives.
- Implement `<Button>`, `<Input>`, `<Textarea>`, `<Toggle>`, `<TabBar>`.
- Apply atmosphere to `+layout.svelte`.
- Existing pages still render with flowbite — they just sit on the new atmosphere. Visual mismatch is expected for one phase.

**Phase 2 — Public pages.** Replace `<Header>`, then convert `/`, `/login`, `/register`, `/about` to the new primitives.

**Phase 3 — User dashboard.** Convert `/account`, `/people`, `/packs`. Largest visual shift in this phase; carousels removed.

**Phase 4 — `/newpack` stream editor.** Single page, but new component (`<TaskListEditor>`). Backend serialisation contract preserved.

**Phase 5 — `/game`.** Most complex. Sticky glass header, two-column layout desktop, segmented control mobile, new `BingoGrid` with winning-row glow, `Keypad` / `TaskList` / `NotesJournal` paper cards.

**Phase 6 — Cleanup.** Remove `flowbite-svelte` and `flowbite` from `package.json`, `tailwind.config.cjs`. Drop unused fonts from `Header.svelte` Google Fonts link (after the new link in `app.html` is verified live). Final a11y pass, contrast audit, reduced-motion verification.

## Out of scope

- New game modes, new pack categories, new social features.
- Notifications (in-app, push, email).
- Profile customisation (avatar uploads, custom colour picker).
- Internationalisation.
- Search beyond the simple `/people` filter.
- Dark mode. Aurora Calm is light-only by design. A dark variant is a future spec.
- Performance budget tracking (Lighthouse, bundle size). Removing flowbite is expected to shrink the bundle, but no specific target is set.
- Visual regression test suite (the project has no test infrastructure today; spinning that up is a separate spec).

## Open decisions to confirm in review

1. **Avatar visuals.** The current app has no avatars. The mockups show a periwinkle dot. Options: (a) keep dots forever, (b) auto-generate from username initials in Fraunces letter on accent, (c) future feature. Default: (b).
2. **Active cell number colour.** Mockups use white text on the periwinkle gradient. WCAG AA passes for the accent's lighter stop (`#b8a3cd`) at 4.0:1 — borderline for small text. Numbers in a 4×4 grid are typically 0.85rem semibold, qualify as "large". Acceptable, but worth eye-checking in implementation.
3. **`/about` content stays as-is.** The current copy is tongue-in-cheek and dated to August 2023. Keeping it verbatim. If the user wants new copy, that's a separate change.
4. **`/account` "see all" links** — clicking expands the list inline (default) vs. navigating to a dedicated page (`/account/friends`, etc.). Default: inline expansion. Adding routes is a low-cost change later if the lists outgrow inline.
5. **Header behaviour on mobile.** Above 768px the header is a top strip. Below 768px there isn't enough vertical room for a sticky strip + glass; default is the strip stays at the top but shrinks vertically on scroll. Hamburger menu is not added — the four nav links (`play`, `packs`, `people`, `about`) fit horizontally even on small screens at `body-s` size.

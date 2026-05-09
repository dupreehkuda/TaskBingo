# Deploying TaskBingo to Timeweb (docker-compose app)

Target: Timeweb Cloud "Apps" service, the docker-compose flavour. External
Postgres already provisioned on Timeweb DBaaS with the existing schema
(migrations `0001_init` already applied through some prior path).

This deploy uses three containers behind a single nginx reverse proxy:

```
[ Timeweb edge / TLS ]
         │
       :80
         │
   ┌──nginx──┐
   │         │
 /api/    everything else
   │         │
:8082    :3000
server     web
   │
   └── DATABASE_DSN ──► Timeweb managed Postgres (external)
```

The browser only ever talks to the public domain — same-origin, no CORS,
no preflight surprises.

---

## 1. Pre-flight on the Postgres side

The Go server runs `golang-migrate` on startup (see
`internal/app/migrate.go`). It tracks applied migrations in a
`schema_migrations` table. Two cases:

### Case A — `schema_migrations` already exists and tracks `1`

You're golden. Booting the new image will apply `0002` and `0003` and
that's it.

Quick check from any psql client:

```sql
SELECT version, dirty FROM schema_migrations ORDER BY version;
```

Expected: one row `1, false` (or `2, false` / `3, false` if newer).

### Case B — `schema_migrations` does not exist

The DB has data but was never managed by `golang-migrate`. Booting the
server would try to apply `0001` from scratch. `0001_init.up.sql` is
written with `CREATE TABLE IF NOT EXISTS` everywhere so re-running is
safe — no duplicate-table errors. But `0002` has one
non-idempotent statement:

```sql
ALTER TABLE games ALTER COLUMN user2_id DROP NOT NULL;
```

If `user2_id` is already nullable, Postgres treats `DROP NOT NULL` as a
no-op (no error). If it's still `NOT NULL`, it succeeds. Safe both ways.

So the SAFE path is:

1. **Backup the DB first** via Timeweb panel ("Создать резервную копию").
2. Manually mark `0001` as applied so the migrator skips it:

   ```sql
   CREATE TABLE IF NOT EXISTS schema_migrations (
     version bigint NOT NULL PRIMARY KEY,
     dirty   boolean NOT NULL
   );
   INSERT INTO schema_migrations (version, dirty)
   VALUES (1, false)
   ON CONFLICT (version) DO NOTHING;
   ```

3. Boot the server — it will apply `0002` and `0003`.

If you'd rather skip the manual seed, just let the server boot with no
`schema_migrations` row; it will harmlessly re-run `0001` (idempotent)
then `0002` then `0003`. Confirm by checking the table afterwards.

### What the new migrations do

- `0002_solo_private_comments` — adds `packs.is_private`, makes
  `games.user2_id` nullable, adds `games.kind`, adds `users.solo_bingo`,
  creates `game_comments` table.
- `0003_stats_indexes` — composite indexes
  `games(user1_id, finished)` and `games(user2_id, finished)` for the
  new `/api/user/stats` endpoint.

Neither destroys data. Backups still recommended.

---

## 2. Files needed in the repo

These are committed and ready:

| File                              | Purpose                                                                   |
| --------------------------------- | ------------------------------------------------------------------------- |
| `Dockerfile`                      | Server image (Go, runs migrations on boot)                                |
| `bingo-app/Dockerfile`            | Web image (SvelteKit + adapter-node, real prod build, not Vite preview)   |
| `docker-compose.yml`              | The compose Timeweb auto-detects from the repo root                       |
| `docker-compose.local.yml`        | Local-only compose (postgres container, host port mappings) — not used in prod |
| `nginx.conf`                      | Reverse proxy: `/api/*` → server, `/ *` → web, WebSocket upgrade for `/api/game/start` |
| `.env.timeweb.example`            | Template for the prod env vars                                            |
| `migrations/postgresql/*.sql`     | Migration files baked into the server image at build                       |

---

## 3. `.env` on the Timeweb host

Copy `.env.timeweb.example` to `.env` (Timeweb's panel exposes an "env
vars" section that you can fill in instead — same effect). Replace every
placeholder:

```
DATABASE_DSN=postgresql://USER:PASSWORD@HOST:5432/DBNAME?sslmode=require
SECRET=<openssl rand -hex 32 — at least 32 chars>
CURRENT_DOMAIN=taskbingo.com
ALLOWED_WS_ORIGINS=https://taskbingo.com
VITE_API_URL=https://taskbingo.com
VITE_WS_URL=wss://taskbingo.com
VITE_WEB_URL=taskbingo.com
GAME_SERVICE_ADDRESS=:8082
```

Notes:
- `SECRET` must be the same across deploys — changing it invalidates
  every active session (every `auth` cookie).
- `CURRENT_DOMAIN` controls the cookie's `Domain=` and the `Secure` flag
  (the backend sets `Secure: domain != "localhost"` in
  `internal/rpc/http/user/register.go:setAuthCookie`).
- `VITE_*` are baked into the web image at build time. Changing them
  requires rebuilding the `web` service.

---

## 4. Deploy on Timeweb

In the Timeweb panel:

1. **Create a new App → Docker Compose**.
2. Connect this Git repo (or upload zipped contents). Timeweb picks up
   `docker-compose.yml` at the repo root automatically — no extra
   config-file selection needed.
3. Paste the env vars from `.env.timeweb.example` into the panel's
   env-vars section, with real values.
4. Map external port **80 → container nginx 80**. (Timeweb's edge
   handles TLS — they expose your service over HTTPS regardless of what
   port your app listens on internally.)
5. Add the domain (`taskbingo.com`) in the App's domain settings. Point
   the DNS A/CNAME record at the value Timeweb shows.
6. Click Deploy.

First build pulls Node 20 + Go 1.25 alpine images, builds the SvelteKit
bundle and the Go binary. Takes ~3–5 minutes.

When the App is healthy:

- `https://taskbingo.com/` should serve the homepage (HeroDemo board
  animating).
- `https://taskbingo.com/api/user/getAllUsers` without auth should
  respond `401 Unauthorized` — confirms the proxy reaches the Go server
  and CheckToken middleware fires.

---

## 5. Verify migrations applied

From any psql client connected to the prod DB:

```sql
SELECT version, dirty FROM schema_migrations ORDER BY version;
-- expect: 1 false, 2 false, 3 false

\d users
-- should include solo_bingo

\d games
-- user2_id should be nullable; should include `kind`

\d packs
-- should include is_private

\d game_comments
-- table should exist

\d games_user1_finished_idx
-- should be a btree on (user1_id, finished)
```

If any expected column / index is missing, check `taskbingo-server`
logs — `MigrateUp` failures are logged before the server exits with
`migrate up failed`.

---

## 6. Subsequent updates

When you push new code:

1. In Timeweb panel, click "Rebuild" on the App.
2. Compose runs: `down → build → up -d`.
3. Server restart triggers `MigrateUp` automatically — any new
   migration files in `migrations/postgresql/` get applied in version
   order. No manual SQL needed.

If you need a fast-path without rebuilding the web image (only backend
changes), via SSH on the host:

```bash
docker compose build server
docker compose up -d server
```

Cookie-based sessions survive the restart since cookies live in the
browser. The localStorage `tb-authed` flag is also browser-side, so
existing sessions just pick up where they left off.

---

## 7. Troubleshooting

**`Failed to load resource: ... Access-Control-Allow-Credentials`** — the
browser is hitting a different origin than where the page was served.
Check that `VITE_API_URL` matches `https://CURRENT_DOMAIN` exactly (with
the protocol). Rebuild the web image to bake the corrected value in.

**`get_data failed: no rows in result set`** — the JWT cookie is signed
with the prod `SECRET` but points at a `userID` that doesn't exist in
the prod DB. Usually because someone copied a dev cookie. Have the user
clear cookies for the domain and re-register.

**Server keeps restarting** — `docker logs taskbingo-server`. Most
common: bad `DATABASE_DSN` (password contains special chars that need
URL-encoding) or the migration step fails on a custom-applied schema.
Restore from backup, fix the migration manually, retry.

**WebSocket disconnects after seconds** — Timeweb edge default idle
timeout might be short. The nginx config already sets
`proxy_read_timeout 3600s`. If disconnects persist, look at the Timeweb
panel for an "idle timeout" or "WebSocket support" toggle.

**Cookie not set / `Secure` missing** — the cookie is `Secure` because
`CURRENT_DOMAIN != "localhost"`. Browsers refuse to set it unless the
page is served over HTTPS. Double-check that Timeweb is actually
terminating TLS and the public URL is `https://`, not `http://`.

---

## 8. Rollback

Worst case: redeploy a previous git commit (Timeweb panel → "Deploy
specific revision" or revert the branch and click Rebuild).

Backend rollback: `MigrateDown` is **not** wired into the server
binary on purpose — automatic down-migrations are too dangerous. To
roll back schema, run the appropriate `*.down.sql` file manually with
`psql`, then deploy the older server image. Order matters: down 0003
before down 0002.

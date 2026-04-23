The Next.js advanced flag scaffolds a modern Next.js frontend alongside your Go backend. It runs the official `create-next-app` CLI with sensible defaults and initializes [shadcn/ui](https://ui.shadcn.com/) so you have production-grade components available from the first commit.

The Next.js flag is mutually exclusive with the React and HTMX/Templ flags — only one frontend option may be selected at a time.

## What it generates

`create-next-app@latest` is invoked with:

```
--ts --app --eslint --src-dir --tailwind --use-npm --import-alias "@/*"
```

Which gives you:

- **App Router** (Next.js 13+ default)
- **TypeScript**
- **Tailwind CSS** (baked in — no separate `--feature tailwind` needed)
- **ESLint**
- **`src/` directory** layout
- **`@/*` import alias**

On top of that, the blueprint:

- Runs `npx shadcn@latest init -d` and adds the `button` and `card` components as a baseline.
- Overwrites `src/app/page.tsx` with a starter page that fetches from the Go backend using `process.env.NEXT_PUBLIC_API_URL`.
- Writes `frontend/.env.local` with `NEXT_PUBLIC_API_URL=http://localhost:<PORT>` (read from the project's root `.env`).
- Writes `frontend/next.config.mjs` with `output: "standalone"` and a `rewrites()` rule proxying `/api/:path*` to the Go backend — useful for server-side fetches where `NEXT_PUBLIC_*` is not appropriate.

## Project structure

```bash
/ (Root)
├── frontend/                     # Next.js advanced flag. Excludes HTMX and React.
│   ├── .env.local                # NEXT_PUBLIC_API_URL pointing at the Go backend.
│   ├── next.config.mjs           # standalone output + /api/* proxy to Go.
│   ├── components.json           # shadcn/ui config.
│   ├── src/
│   │   ├── app/
│   │   │   ├── layout.tsx
│   │   │   └── page.tsx          # Overwritten with a Go-backend demo page.
│   │   ├── components/ui/        # shadcn components (button, card, ...).
│   │   └── lib/utils.ts
│   ├── public/
│   ├── tsconfig.json
│   └── package.json
```

## Usage

```bash
cd frontend
npm install
npm run dev
```

The Next.js dev server runs on `http://localhost:3000` and talks to the Go backend on `http://localhost:8080` (or whatever `PORT` is set to).

## Makefile

`make run` starts the Go server and the Next.js dev server together:

```makefile
run:
    @go run cmd/api/main.go &
    @npm install --prefer-offline --no-fund --prefix ./frontend
    @npm run dev --prefix ./frontend
```

## Docker

Combine with the `docker` feature to get a multi-stage Dockerfile and docker-compose service. The frontend service is exposed on `3000` and points at the `app` service via `NEXT_PUBLIC_API_URL=http://app:${PORT}`. Spin everything up with:

```bash
make docker-run
```

## Environment variables

| Variable | Where | Purpose |
|---|---|---|
| `PORT` | project `.env` | Go backend port |
| `NEXT_PUBLIC_API_URL` | `frontend/.env.local` | Base URL the client uses to call the Go API |

If you change `PORT`, update `NEXT_PUBLIC_API_URL` to match.

## Adding more shadcn components

```bash
cd frontend
npx shadcn@latest add dialog input form
```

## Notes

- The first run downloads `create-next-app` and shadcn dependencies; subsequent runs are much faster thanks to npm's cache.
- `--no-turbopack` is passed for stability; swap it out once Turbopack graduates.

# Golang App
This repository is a Go project designed for practicing and learning how to create a REST API application using Go. 
I'm follow the Ruby on Rails (ROR) folder structure as a base because I'm using ruby as main languange, I think it's fun to create this project using ruby structure folders

# Requirements

| Tool      | Version    | Notes                                                        |
|-----------|------------|--------------------------------------------------------------|
| Go        | `1.23+`    | `go.mod` now declares `go 1.23.0`                            |
| Node.js   | `>= 20`    | required by Vite 6 (`engines` in `package.json`)             |
| `fresh`   | latest     | `go install github.com/pilu/fresh@latest` (hot reload)       |
| PostgreSQL| 14+        | see the Database section                                     |
| Redis     | 6+         | sessions (optional in development)                           |
| Fiber     | `v3`       | see the Fiber v3 section                                     |

# Project Structure (Rails-like)
The project follows the Ruby on Rails folder layout, adapted to Go conventions:

| Rails                 | This project                         | Notes                                              |
|-----------------------|--------------------------------------|----------------------------------------------------|
| `app/controllers`     | `app/controllers`                    | Shell controller + `app/controllers/api` for JSON  |
| `app/models`          | `app/models`                         | ActiveRecord-like models (GORM / PostgreSQL)       |
| `app/views`           | `app/views`                          | Single embedded SPA shell + renderer (`shell.go`)  |
| `app/javascript`      | `app/javascript`                     | Vue 3 SPA: `main.js`, `router.js`, `views/`        |
| —                     | `static/javascript`                  | Reusable components/forms (imported via `@legacy`) |
| `app/helpers`         | `app/helpers`                        | View/presentation helpers                          |
| `app/mailers`         | `app/mailers`                        | Base mailer for outbound email                     |
| `app/jobs`            | `app/jobs`                           | Background job contract                            |
| `app/services`        | `app/services`                       | Service objects (Rails-style PORO/services)        |
| `config/`             | `config/`                            | Application wiring (`application.go`, `routes.go`)  |
| `config/database.yml` | `config/postgres`                    | Primary DB (PostgreSQL + GORM)                     |
| —                     | `config/mongo`                       | Legacy Mongo connection, kept for reference        |
| `config/environments` | `config/environments`                | `development`, `test`, `production`                |
| `config/initializers` | `config/initializers`                | Boot hooks run once at startup                     |
| `config/routes.rb`    | `config/routes.go`                   | Route definitions (HTML shell + `/api/v1`)         |
| `config/locales`      | `config/localize` + `i18n.yml`       | i18n / translations                                |
| `log/production.log`  | `config/logger`                      | Request logger (`Stream` writer in Fiber v3)       |
| `db/migrate`          | `db/migrate`                         | Timestamped SQL migrations + `schema_migrations`   |
| `bin/rails db:migrate`| `db/migrate/cmd`                     | Migration runner (`-direction=up/down/status`)     |
| `db/schema.rb`        | `schema_migrations` table            | Applied versions ledger                            |
| `db/seeds.rb`         | `seeds/`                             | Database seeders (`seeds/seed_example.go`)          |
| `lib/`                | `lib/`                               | Shared utilities (`lib/env.go`, `lib/directory.go`)|
| `public/`             | `static/`                            | Built assets; `static/dist` is embedded in the binary |
| `test/`               | `test/`                              | Unit tests (`test/views` renders the shell)        |
| `Gemfile`             | `go.mod` / `go.sum`                  | Dependency manifest                                |
| `config.ru`           | `main.go`                            | Application entry point                            |
| `Rakefile`            | `Makefile`                           | Task runner (`make help`)                          |
| `jsbundling`/webpacker| `vite.config.js`                     | Asset build pipeline (Vite)                        |
| `tailwind.config.js`  | `static/css/tailwind.css` (`@theme`) | Tailwind CSS v4 is configured in CSS               |

New Rails-style folders added here (`app/helpers`, `app/mailers`, `app/jobs`, `db/migrate`, `config/initializers`) are wired in a Go-friendly way:

- `config/initializers` exposes a registry; add `init()` files and `main.go` runs them once at boot via `initializers.Run()`.
- `db/migrate` holds one timestamped file per change, registered with `migrate.Register` and tracked in `schema_migrations` like Rails.
- `app/mailers`, `app/jobs`, and `app/helpers` provide base types/helpers ready to be extended per feature.
- `Makefile` mirrors the most used `rake` tasks (`make migrate`, `make seed`, `make assets-watch`, ...).

# Web Framework: Fiber v3

The HTTP layer runs on **Fiber v3** (`github.com/gofiber/fiber/v3`). v3 is a breaking release, so
this project was updated together with the upgrade:

| Fiber v2                                   | Fiber v3 (this project)                                  |
|--------------------------------------------|----------------------------------------------------------|
| `func(c *fiber.Ctx) error`                 | `func(c fiber.Ctx) error` (`Ctx` is now an interface)     |
| `c.BodyParser(&v)`                         | `c.Bind().Body(&v)` (`Bind().Query()`, `Bind().URI()`)    |
| `c.Redirect(url, fiber.StatusTemporaryRedirect)` | `c.Redirect().Status(...).To(url)`                  |
| `github.com/gofiber/fiber/v2/middleware/...`| `github.com/gofiber/fiber/v3/middleware/...`            |
| `logger.Config{Output: w}`                 | `logger.Config{Stream: w}`                                |
| `github.com/gofiber/storage/redis`         | removed - `config/session/storage.go` wraps go-redis v9   |
| `github.com/gofiber/template/html/v2`      | removed - the HTML layer is a `html/template` shell       |
| `app.Listen(":3000")` (error ignored)      | `Listen` error is returned, `ShutdownWithTimeout` on SIGINT/SIGTERM |

Where the changes live:

- `config/routes.go` - `fiber.New` config (`ErrorHandler`, `BodyLimit`), session `Locals`, graceful shutdown in `StartServer()`.
- `app/middlewares/*.go` - handlers take `fiber.Ctx`; the JWT middleware now answers **401** exactly once instead of silently continuing.
- `app/controllers/**` - `Bind().Body()` instead of `BodyParser()`, and proper status codes (`400`, `401`, `404`, `409`, `422`).
- `config/session/` - `SessionStoreNew(client)` reuses the single Redis connection from `config/redis` through a small `fiber.Storage` adapter (`Get`/`Set`/`Delete`/`Reset`/`Close`), so no driver pinned to Fiber v2 is pulled in.
- `config/logger/logger.go` - Fiber v3 logger middleware (import paths + `Stream`).
- `app/middlewares/jwt.go` - the signing secret is read lazily (it used to be read in `init()`, before `.env` was loaded, so tokens were signed with an empty key); `exp`/`iat` claims are set and the HMAC algorithm is pinned.

Refresh the module graph once after the upgrade:

````
go mod tidy
````

If `v3.0.0-beta.4` no longer resolves, point the module at the current release:

````
go get github.com/gofiber/fiber/v3@latest
go mod tidy
````

# Views: one embedded SPA shell

There is **no server-side rendering of pages** and no template engine. Every HTML route answers with
the same shell (`app/views/layouts/application.html`); Vue takes over in the browser.

| Rails                                   | This project                                          |
|-----------------------------------------|-------------------------------------------------------|
| `app/views/layouts/application.html.erb`| `app/views/layouts/application.html` (single shell)   |
| `<%= yield %>`                          | `<div id="app"></div>` + Vue router                   |
| `render "wedding/show"`                 | `views.RenderShell(c, views.Props{...})`              |
| `content_for :javascript`               | `<script type="module">` resolved from the manifest   |
| `stylesheet_link_tag` / `javascript_include_tag` | hashed URLs from `.vite/manifest.json`       |
| `gon` / `window.__INITIAL_STATE__`      | `<script type="application/json" id="app-props">`     |
| `config/routes.rb`                      | `config/routes.go` (mirrored by `app/javascript/router.js`) |
| `rescue_from` / `public/500.html`       | `errorHandler` in `config/routes.go`                  |
| `get "*path" => "errors#not_found"`     | `app.Use(spaController.NotFound())`                   |

How it is put together:

- `app/views/assets.go` resolves the bundle: `ModeEmbed` reads `static/dist` from the binary
  (production), `ModeDisk` reads `static/dev` from disk (development/test) and re-reads the manifest
  whenever Vite rewrites it, so `npm run dev` needs no Go restart.
- `app/views/shell.go` parses the layout once (re-parsing by mtime in development), injects the
  hashed CSS/JS, the `<head>` metadata and the props, and always answers `Cache-Control: no-store`.
- `app/views/handler.go` serves the embedded bundle at `/public/*` in production with
  `immutable` caching for the hashed files.
- `app/views/embed.go` + `static/embed.go` hold the `//go:embed` directives (`layouts`, `dist`,
  `styles`).
- `app/controllers/spa_controller.go` renders the shell for `/`, `/wedding/:id` and the 404 fallback.

Nothing secret ever reaches the markup: authentication lives in an httpOnly cookie, and the client
props only carry `route`, `params`, `locale` and `env`.

# Frontend: SPA + Vite + go:embed

Webpack, the standalone `tailwindcss` CLI, PostCSS **and** the Fiber template engine are gone.

````
app/javascript/main.js         -> the only Vite entry (imports static/css/tailwind.css)
app/javascript/router.js       -> client routes, mirrors config/routes.go
app/javascript/lib/api.js      -> JSON client (cookies + CSRF header)
app/javascript/views/*.vue     -> HomeView, WeddingView, NotFoundView
static/javascript/**           -> reusable components, imported through the @legacy alias
static/css/tailwind.css        -> Tailwind v4 entry (@import "tailwindcss" + @theme)
vite.config.js                 -> build config (mode based output directory + manifest)
````

| Mode / command            | Output                | Public URL                  |
|---------------------------|-----------------------|-----------------------------|
| `npm run build:dev`       | `static/dev`          | `/public/dev/...` (disk)    |
| `npm run dev` (watch)     | `static/dev`          | `/public/dev/...` (disk)    |
| `npm run build:prod`      | `static/dist`         | `/public/dist/...` (embedded)|

Filenames are hashed (`javascript/main-<hash>.js`, `assets/tailwind-out-<hash>.css`) and linked
through `.vite/manifest.json`, which is why the Go side never hardcodes an asset path.

Build the assets once:

````
npm install
npm run build:dev
````

Watch them while developing (single process, rewrites `static/dev`):

````
npm run dev
````

Build the minified production assets **before** compiling the binary, because `static/dist` is
embedded into it:

````
npm run build:prod
go build
````

`static/dist/.gitkeep` exists only so `//go:embed all:dist` keeps compiling on a clean checkout;
`static/dist` and `static/dev` are git-ignored.

## Tailwind CSS v3 -> v4 notes

- The dev/prod Tailwind configs are gone: `content` became `@source` and `theme.extend.fontFamily`
  became the `@theme` block (`--font-marsha`, `--font-dosis`, `--font-playfair-display`) inside
  `static/css/tailwind.css`. The `safelist: /./` is no longer needed because Vite compiles once and
  watches the sources (dynamic classes written as plain strings are still detected).
- `@source` must keep covering `app/javascript`, `app/views` **and** `static/javascript`; a missing
  entry silently drops utilities from the compiled CSS.
- Tailwind v4 renamed a few utilities (`shadow-sm` -> `shadow-xs`, `rounded-sm` -> `rounded-xs`,
  `outline-none` -> `outline-hidden`, `bg-opacity-*` -> `bg-black/50`, ...). Check a component if a
  style suddenly looks different.
- If a Vue `<style>` block uses `@apply`, add a reference to the Tailwind entry at the top of that block:

  ````
  @reference "../../css/tailwind.css";
  ````

# Authentication: cookie + CSRF

- `POST /api/v1/users/login` returns the token in the body **and** stores it in an httpOnly cookie
  (`auth_token`, `SameSite=Lax`, `Secure` in production). `app/javascript/lib/api.js` therefore uses
  `withCredentials: true` and never touches the token itself.
- A second, readable cookie (`csrf_token`) is issued at the same time. Every non-GET request must
  echo it in the `X-CSRF-Token` header (double submit, `app/middlewares/auth_cookie.go`).
- Requests with no cookies at all (login, sign-up, curl with a bearer token) are exempt: a
  cross-site attacker cannot forge them.
- `app/middlewares/jwt.go` reads the cookie first and falls back to `Authorization: Bearer <token>`
  so CLI tools and the existing API tests keep working.
- `POST /api/v1/users/logout` clears both cookies.
- `CSRF` is attached to the whole `/api/v1` group in `config/routes.go`.

# Database (PostgreSQL)

The application now talks to **PostgreSQL** through [GORM](https://gorm.io), the closest Go equivalent to ActiveRecord.

- Connection: `config/postgres/postgres.go` (`postgres.NewPostgres()`), configured with the `POSTGRES_*` variables in `.env`.
- Models: `app/models/*.go` — `TableName()` pins Rails-style table names (`users`, `wedding_data`), UUID primary keys are generated in a `BeforeCreate` hook, and nested payloads (bride, groom, gallery, gifts, ...) are stored as `jsonb`.
- Services: `app/services/*.go` — one method per ActiveRecord-style query (`Where`, `First`, `Create`, `Updates`).
- Migrations: `db/migrate/*.go`, tracked in the `schema_migrations` table.
- Seeds: `seeds/seed_example.go`.

The old Mongo code in `config/mongo` (and the `MONGO_*` variables) is **intentionally not deleted**: it still compiles and can be used as a fallback/reference while the migration settles.

Because `gorm.io/gorm` and `gorm.io/driver/postgres` were added to `go.mod` (and the dependencies were upgraded), refresh `go.sum` once:

````
go mod tidy
````

Create the database first (Rails: `bin/rails db:create`):

````
createdb golang_app_development
````

Run the migrations (Rails: `bin/rails db:migrate`):

````
make migrate
# or: go run ./db/migrate/cmd -direction=up
````

Roll back the last migration (Rails: `bin/rails db:rollback`):

````
make rollback
````

Check migration status:

````
make migrate-status
````

Load the seed data (Rails: `bin/rails db:seed`):

````
make seed
````

# Running the Go Application

Run the Go application with hot reload:
````
make watch
# or: fresh
````

Run the Go application:
````
make run
# or: go run .
````

Run the Go application with a specific environment:
````
ENV=development go run .
ENV=production go run .
````

`Ctrl+C` now triggers a graceful shutdown (Fiber v3): in-flight requests get up to 10 seconds to
finish, then the PostgreSQL pool and the Redis client registered in `main.go` are closed.

# Development
Run two processes: the Go app (`make watch`) and the asset watcher (`make dev` / `npm run dev`).
Because Tailwind v4 and Vue are both compiled by Vite, adding new components or new utility classes
no longer requires a separate CSS rebuild step - the watcher picks the change up.

The HTML shell is rendered from disk in development and re-parsed when it changes, so editing
`app/views/layouts/application.html` does not require restarting `go run .` (`.fresh.yaml` keeps
`.html` in `no_rebuild_ext`).

# Running Tests
Run all unit tests:
````
make test
# or: go test ./...
````
Run specific unit tests:
````
go test ./test/views
````
If you want to separate environment variables for production and test environments, create a new file named .env.test.

Run all unit tests with .env.test:
````
ENV=test go test ./...
````

Run specific unit tests with .env.test:
````
ENV=test go test ./test/views
````

`test/views/shell_test.go` builds a fake Vite bundle in a temp directory and asserts that the shell
links the hashed files, injects the props, never leaks a token, and that the CSRF middleware rejects
a cookie request without the `X-CSRF-Token` header.

# Component Guidelines

The components used in this project follow the guidelines from [HyperUI](https://www.hyperui.dev/).

# Running StoryBook

Storybook now uses the Vite builder (`@storybook/vue3-vite`), so it reuses `vite.config.js`
(Vue + Tailwind CSS v4) and receives the compiled stylesheet from `.storybook/preview.js`.

````
make storybook
# or: npm run storybook
````

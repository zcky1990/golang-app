package views_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	mid "golang_app/golangApp/app/middlewares"
	"golang_app/golangApp/app/views"

	"github.com/gofiber/fiber/v3"
)

// writeFakeBundle creates a minimal Vite output (manifest + emitted files) so the
// shell can be rendered without running a real build.
func writeFakeBundle(t *testing.T) string {
	t.Helper()

	root := t.TempDir()
	dev := filepath.Join(root, "dev")

	for _, dir := range []string{filepath.Join(dev, ".vite"), filepath.Join(dev, "javascript"), filepath.Join(dev, "assets")} {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			t.Fatalf("mkdir %s: %v", dir, err)
		}
	}

	manifest := `{"app/javascript/main.js":{"file":"javascript/main-abc123.js","src":"app/javascript/main.js","isEntry":true,"css":["assets/tailwind-out-def456.css"]}}`
	if err := os.WriteFile(filepath.Join(dev, ".vite", "manifest.json"), []byte(manifest), 0o644); err != nil {
		t.Fatalf("write manifest: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dev, "javascript", "main-abc123.js"), []byte("// bundle"), 0o644); err != nil {
		t.Fatalf("write js: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dev, "assets", "tailwind-out-def456.css"), []byte("/* css */"), 0o644); err != nil {
		t.Fatalf("write css: %v", err)
	}

	return root
}

func renderShell(t *testing.T, props views.Props) *http.Response {
	t.Helper()

	views.ConfigureWith(views.ModeDisk, writeFakeBundle(t), "/public")

	app := fiber.New()
	app.Get("/", func(ctx fiber.Ctx) error {
		return views.RenderShell(ctx, props)
	})

	response, err := app.Test(httptest.NewRequest(http.MethodGet, "/", nil))
	if err != nil {
		t.Fatalf("request: %v", err)
	}
	return response
}

func readBody(t *testing.T, response *http.Response) string {
	t.Helper()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	if err := response.Body.Close(); err != nil {
		t.Fatalf("close body: %v", err)
	}
	return string(body)
}

func TestRenderShellLinksHashedAssets(t *testing.T) {
	response := renderShell(t, views.Props{
		Title:       "Wedding Invitation",
		Description: "Open a wedding invitation.",
		Locale:      "en",
		Route:       "wedding",
		Params:      map[string]string{"id": "6f1a1b1c-1111-2222-3333-444455556666"},
	})
	body := readBody(t, response)

	if response.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", response.StatusCode)
	}

	for _, want := range []string{
		`<div id="app"></div>`,
		`<title>Wedding Invitation</title>`,
		`name="description" content="Open a wedding invitation."`,
		`/public/dev/assets/tailwind-out-def456.css`,
		`/public/dev/javascript/main-abc123.js`,
		`type="module"`,
		`id="app-props"`,
		`"route":"wedding"`,
		`"id":"6f1a1b1c-1111-2222-3333-444455556666"`,
		`lang="en"`,
	} {
		if !strings.Contains(body, want) {
			t.Errorf("shell is missing %q", want)
		}
	}

	// The shell must never leak a credential into the markup.
	for _, unwanted := range []string{"auth_token", "csrf_token", "Bearer "} {
		if strings.Contains(body, unwanted) {
			t.Errorf("shell must not contain %q", unwanted)
		}
	}
}

func TestRenderShellHonoursStatus(t *testing.T) {
	response := renderShell(t, views.Props{Title: "Page not found", Route: "not_found", Status: fiber.StatusNotFound})
	_ = readBody(t, response)

	if response.StatusCode != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", response.StatusCode)
	}
	if got := response.Header.Get(fiber.HeaderCacheControl); got != "no-store" {
		t.Errorf("expected the shell to be uncacheable, got %q", got)
	}
}

func TestCSRFRequiresHeaderForCookieRequests(t *testing.T) {
	app := fiber.New()
	app.Post("/api/v1/users/update-user", mid.CSRF(), func(ctx fiber.Ctx) error {
		return ctx.SendString("ok")
	})

	req := httptest.NewRequest(http.MethodPost, "/api/v1/users/update-user", nil)
	req.AddCookie(&http.Cookie{Name: mid.AuthCookieName, Value: "jwt"})
	req.AddCookie(&http.Cookie{Name: mid.CSRFCookieName, Value: "token"})

	response, err := app.Test(req)
	if err != nil {
		t.Fatalf("request: %v", err)
	}
	_ = readBody(t, response)

	if response.StatusCode != http.StatusForbidden {
		t.Fatalf("expected 403 without the CSRF header, got %d", response.StatusCode)
	}
}

func TestCSRFAcceptsMatchingHeader(t *testing.T) {
	app := fiber.New()
	app.Post("/api/v1/users/update-user", mid.CSRF(), func(ctx fiber.Ctx) error {
		return ctx.SendString("ok")
	})

	req := httptest.NewRequest(http.MethodPost, "/api/v1/users/update-user", nil)
	req.AddCookie(&http.Cookie{Name: mid.AuthCookieName, Value: "jwt"})
	req.AddCookie(&http.Cookie{Name: mid.CSRFCookieName, Value: "token"})
	req.Header.Set(mid.CSRFHeaderName, "token")

	response, err := app.Test(req)
	if err != nil {
		t.Fatalf("request: %v", err)
	}
	_ = readBody(t, response)

	if response.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 with the CSRF header, got %d", response.StatusCode)
	}
}

func TestCSRFSkipsPreAuthRequests(t *testing.T) {
	app := fiber.New()
	app.Post("/api/v1/users/login", mid.CSRF(), func(ctx fiber.Ctx) error {
		return ctx.SendString("ok")
	})

	response, err := app.Test(httptest.NewRequest(http.MethodPost, "/api/v1/users/login", nil))
	if err != nil {
		t.Fatalf("request: %v", err)
	}
	_ = readBody(t, response)

	if response.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 for a cookieless request, got %d", response.StatusCode)
	}
}

func TestIsHTMLRequest(t *testing.T) {
	app := fiber.New()
	// Use (not Get) so the probe runs for every path, including /api/*.
	app.Use(func(ctx fiber.Ctx) error {
		if views.IsHTMLRequest(ctx) {
			return ctx.SendString("html")
		}
		return ctx.SendString("json")
	})

	cases := []struct {
		path   string
		accept string
		want   string
	}{
		{"/", "text/html,application/xhtml+xml", "html"},
		{"/wedding/abc", "*/*", "html"},
		{"/api/v1/unknown", "application/json", "json"},
		{"/api/v1/unknown", "", "json"},
		{"/missing", "application/json", "json"},
	}

	for _, tc := range cases {
		req := httptest.NewRequest(http.MethodGet, tc.path, nil)
		if tc.accept != "" {
			req.Header.Set("Accept", tc.accept)
		}
		response, err := app.Test(req)
		if err != nil {
			t.Fatalf("request %s: %v", tc.path, err)
		}
		body := readBody(t, response)
		if body != tc.want {
			t.Errorf("%s (accept %q): expected %s, got %s", tc.path, tc.accept, tc.want, body)
		}
	}
}

package app

import (
	"encoding/json"
	"net/url"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/karalef/shlink/urlservice"
)

// New creates a new app.
// appURL is the URL of the app.
func New(service urlservice.Service, appURL string) (*App, error) {
	u, err := url.Parse(appURL)
	if err != nil {
		return nil, err
	}
	a := &App{
		srv:     fiber.New(),
		Service: service,
		URL:     u,
	}

	a.srv.Get("/r/:id", a.HandleRedirect)
	a.srv.Get("/create", a.HandleCreate)
	a.srv.Get("/data/:id", a.HandleData)

	return a, nil
}

// App represents the app.
type App struct {
	srv *fiber.App

	Service urlservice.Service
	URL     *url.URL
}

// Run starts the listener with the given addr.
func (a *App) Run(addr string) error {
	return a.srv.Listen(addr)
}

// Shutdown gracefully shuts down the app.
func (a *App) Shutdown(timeout ...time.Duration) error {
	if len(timeout) > 0 {
		return a.srv.ShutdownWithTimeout(timeout[0])
	}
	return a.srv.Shutdown()
}

// MakeURL creates a URL from the given id.
func (a *App) MakeURL(id string) string {
	return a.URL.JoinPath("r", id).String()
}

// HandleRedirect handles a request to redirect to the given URL.
func (a *App) HandleRedirect(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	origin, err := a.Service.GetOrigin(ctx.UserContext(), id)
	if err != nil {
		return err
	}
	return ctx.Redirect(origin)
}

// HandleCreate handles a request to create a new URL.
func (a *App) HandleCreate(ctx *fiber.Ctx) error {
	origin := ctx.Query("origin", "")
	if origin == "" {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	short, err := a.Service.CreateShort(ctx.UserContext(), origin)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.SendString(a.MakeURL(short))
}

// HandleData handles a request to view a URL data.
func (a *App) HandleData(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	u, err := a.Service.Get(ctx.UserContext(), id)
	if err != nil {
		return err
	}
	u.Short = a.MakeURL(id)
	return json.NewEncoder(ctx).Encode(u)
}

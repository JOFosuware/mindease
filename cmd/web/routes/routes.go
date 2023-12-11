package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/jofosuware/mindease/cmd/web/middleware"
	"github.com/jofosuware/mindease/internal/config"
	"github.com/jofosuware/mindease/internal/handlers"
)

func Routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(chiMiddleware.Recoverer)
	//mux.Use(middleware.NoSurf)
	mux.Use(middleware.SessionLoad)

	mux.Get("/", handlers.Repo.GetHome)
	mux.Get("/book-session", handlers.Repo.GetBookSession)
	mux.Get("/bipolar", handlers.Repo.GetBipolarExperts)
	mux.Get("/create-account", handlers.Repo.GetClientForm)
	mux.Post("/create-client", handlers.Repo.PostClientInformation)

	mux.Get("/provider", handlers.Repo.GetProviderPage)
	mux.Get("/provider-form", handlers.Repo.GetProviderForm)
	mux.Post("/client-profile", handlers.Repo.PostClientProfile)
	mux.Get("/our-experts", handlers.Repo.GetProviderProfile)
	mux.Get("/login", handlers.Repo.GetProviderLoginForm)
	mux.Post("/login", handlers.Repo.PostLoginProvider)

	//Notification
	mux.Post("/to-provider", handlers.Repo.PostToProvider)

	//Pharmacy
	mux.Get("/request-drugs", handlers.Repo.GetDrugForm)
	mux.Post("/request-drugs", handlers.Repo.PostDrugRequest)

	//App section
	mux.Get("/about", handlers.Repo.GetAboutPage)
	mux.Get("/services", handlers.Repo.GetServicesPage)
	mux.Get("/portfolio", handlers.Repo.GetPortfolioPage)

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	mux.Route("/admin", func(mux chi.Router) {
		//mux.Get("/signup", handlers.Repo.GetProviderForm)
	})

	return mux
}

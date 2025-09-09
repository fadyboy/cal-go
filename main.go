package main

import (
	"fmt"
	"net/http"

	"github.com/fadyboy/lenslocked/controllers"
	"github.com/fadyboy/lenslocked/models"
	"github.com/fadyboy/lenslocked/templates"
	"github.com/fadyboy/lenslocked/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/csrf"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	// setup db connection
	cfg := models.DefaultDBConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userService := models.UserService{
		DB: db,
	}

	sessionService := models.SessionService{
		DB: db,
	}

	usersC := controllers.Users{
		UserService:    &userService,
		SessionService: &sessionService,
	}

	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))

	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))

	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))

	r.Get("/faq", controllers.FAQ(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "about.gohtml", "tailwind.gohtml"))

	r.Get("/about", controllers.StaticHandler(tpl))

	// Signup
	usersC.Templates.New = views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	r.Get("/signup", usersC.New)

	r.Post("/signup", usersC.Create)

	// SignIn
	usersC.Templates.SignIn = views.Must(views.ParseFS(templates.FS, "signin.gohtml", "tailwind.gohtml"))
	r.Get("/signin", usersC.SignIn)
	r.Post("/signin", usersC.ProcessSignIn)

	// signed in user
	usersC.Templates.New = views.Must(views.ParseFS(templates.FS, "current_user.gohtml", "tailwind.gohtml"))
	r.Get("/users/me", usersC.CurrentUser)

	// Signout
	r.Post("/signout", usersC.ProcessSignOut)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})

	// middleware for setting user context
	umw := controllers.UserMiddleWare{
		SessionService: &sessionService,
	}

	// csrf
	csrfKey := "32-byte-long-auth-key"
	csrfMw := csrf.Protect(
		[]byte(csrfKey),
		csrf.Secure(false),
		csrf.TrustedOrigins([]string{"localhost:3000"}),
	)

	fmt.Println("Starting the server on port :3000")
	http.ListenAndServe(":3000", csrfMw(umw.SetUser(r)))
}

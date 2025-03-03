package main

import (
	"net/http"

	"github.com/TobiasGleiter/go-snippetbox/internal/middleware"
	"github.com/TobiasGleiter/go-snippetbox/ui"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	// Use the http.FileServerFS() function to create a HTTP handler which
	// serves the embedded files in ui.Files. It's important to note that our
	// static files are contained in the "static" folder of the ui.Files
	// embedded filesystem. So, for example, our CSS stylesheet is located at
	// "static/css/main.css". This means that we no longer need to strip the
	// prefix from the request URL -- any requests that start with /static/ can
	// just be passed directly to the file server and the corresponding static
	// file will be served (so long as it exists).
	mux.Handle("GET /static/", http.FileServerFS(ui.Files))

	mux.HandleFunc("GET /ping", ping)

	dynamic := middleware.Chain(app.sessionManager.LoadAndSave, app.authenticate)

	mux.Handle("GET /{$}", dynamic(http.HandlerFunc(app.home)))
	mux.Handle("GET /snippet/view/{id}", dynamic(noSurf(http.HandlerFunc(app.snippetView))))
	mux.Handle("GET /user/signup", dynamic(noSurf(http.HandlerFunc(app.userSignup))))
	mux.Handle("POST /user/signup", dynamic(noSurf(http.HandlerFunc(app.userSignupPost))))
	mux.Handle("GET /user/login", dynamic(noSurf(http.HandlerFunc(app.userLogin))))
	mux.Handle("POST /user/login", dynamic(noSurf(http.HandlerFunc(app.userLoginPost))))

	mux.Handle("GET /snippet/create", dynamic(noSurf(app.requireAuthentication(http.HandlerFunc(app.snippetCreate)))))
	mux.Handle("POST /snippet/create", dynamic(noSurf(app.requireAuthentication(http.HandlerFunc(app.snippetCreate)))))
	mux.Handle("POST /user/logout", dynamic(app.authenticate(noSurf(app.requireAuthentication(http.HandlerFunc(app.userLogoutPost))))))

	standard := middleware.Chain(app.recoverPanic, app.logRequest, commonHeaders)

	// Wrap the existing chain with the recoverPanic middleware.
	return standard(mux)
}

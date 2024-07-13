package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.Handle("GET /{$}", app.sessionManager.LoadAndSave(app.authenticate((http.HandlerFunc(app.home)))))
	mux.Handle("GET /snippet/view/{id}", app.sessionManager.LoadAndSave(app.authenticate(noSurf(http.HandlerFunc(app.snippetView)))))
	mux.Handle("GET /user/signup", app.sessionManager.LoadAndSave(app.authenticate(noSurf(http.HandlerFunc(app.userSignup)))))
	mux.Handle("POST /user/signup", app.sessionManager.LoadAndSave(app.authenticate(noSurf(http.HandlerFunc(app.userSignupPost)))))
	mux.Handle("GET /user/login", app.sessionManager.LoadAndSave(app.authenticate(noSurf(http.HandlerFunc(app.userLogin)))))
	mux.Handle("POST /user/login", app.sessionManager.LoadAndSave(app.authenticate(noSurf(http.HandlerFunc(app.userLoginPost)))))

	mux.Handle("GET /snippet/create", app.sessionManager.LoadAndSave(app.authenticate(noSurf(app.requireAuthentication(http.HandlerFunc(app.snippetCreate))))))
	mux.Handle("POST /snippet/create", app.sessionManager.LoadAndSave(app.authenticate(noSurf(app.requireAuthentication(http.HandlerFunc(app.snippetCreate))))))
	mux.Handle("POST /user/logout", app.sessionManager.LoadAndSave(app.authenticate(noSurf(app.requireAuthentication(http.HandlerFunc(app.userLogoutPost))))))

	// Wrap the existing chain with the recoverPanic middleware.
	return app.recoverPanic(app.logRequest(commonHeaders(mux)))
}

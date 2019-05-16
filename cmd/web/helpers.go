package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// serverError helps manage error reporting and adds a stack trace onto it
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack() )
	app.errorLog.Output(2,trace)  // backs up the stack trace to what led to this locale.
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// clientError reports errors from client
func (app *application) clientError(w http.ResponseWriter, status int ) {
	http.Error(w,http.StatusText(status), status)
}

// notFound is a convienence function to report client not found errors.
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}


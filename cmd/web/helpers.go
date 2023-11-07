package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	app.errorLog.Print(err.Error())
	app.errorLog.Output(2, string(debug.Stack()))
	status := http.StatusInternalServerError
	http.Error(w, http.StatusText(status), status)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) respondOk(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	fmt.Fprint(w, http.StatusText(status))
}

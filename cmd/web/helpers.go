package main

import (
	"net/http"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	app.errorLog.Print(err.Error())
	status := http.StatusInternalServerError
	http.Error(w, http.StatusText(status), status)
}

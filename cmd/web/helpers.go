package main

import "net/http"

func (app *application) printServerError(w http.ResponseWriter, err error) {
	app.errorLog.Print(err.Error())
	http.Error(w, "Internal Server Error", 500)
}

package api

import (
	"FirstProject/pkg/olimpiad"
	"FirstProject/pkg/tableOlimpiads"
	"fmt"
	"net/http"
	"strconv"
)

type httpHandlerShawAddedOlimp struct {
	olimpDataTable *tableOlimpiads.TableOlimpiads
}

func NewHandlerShawAddedOlimp(olimpDataTable *tableOlimpiads.TableOlimpiads) httpHandlerShawAddedOlimp {
	return httpHandlerShawAddedOlimp{olimpDataTable: olimpDataTable}
}

func (h httpHandlerShawAddedOlimp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	subject := r.FormValue("subject")
	level := r.FormValue("level")

	levelInt, _ := strconv.Atoi(level)

	olimp := olimpiad.NewOlimpiad(name, subject, levelInt)
	h.olimpDataTable.Add(olimp)

	fmt.Fprintf(w, "We add olimp with name: %s, subject: %s, level: %s", name, subject, level)
}

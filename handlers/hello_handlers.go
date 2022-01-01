package handlers

import (
	"fmt"
	"net/http"
	"time"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Boilerplate Status OK "+time.Now().String())
}

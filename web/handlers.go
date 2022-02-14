package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/misad3v/go-cli-example.git/app"
)

type Server struct {
	Port string
}

func (s Server) SumHandler(w http.ResponseWriter, r *http.Request) {
	a, err := strconv.ParseFloat(r.URL.Query().Get("a"), 64)

	if err != nil {
		a = 0
	}

	b, err := strconv.ParseFloat(r.URL.Query().Get("b"), 64)

	if err != nil {
		b = 0
	}

	calc := app.NewCalc()

	calc.A = a
	calc.B = b

	w.Write([]byte(fmt.Sprintf("Resultado é %f", calc.Sum())))
}

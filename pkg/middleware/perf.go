package middleware

import (
	"net/http"

	"gopkg.in/macaron.v1"

	m "github.com/p0hil/grafana/pkg/models"
)

func MeasureRequestTime() macaron.Handler {
	return func(res http.ResponseWriter, req *http.Request, c *m.ReqContext) {
	}
}

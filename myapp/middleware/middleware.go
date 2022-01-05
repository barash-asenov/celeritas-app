package middleware

import (
	"github.com/barash-asenov/celeritas"
	"myapp/data"
)

type Middleware struct {
	App    *celeritas.Celeritas
	Models data.Models
}

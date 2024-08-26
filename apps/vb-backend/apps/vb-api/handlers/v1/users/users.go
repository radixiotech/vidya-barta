package v1_users_handler

import (
	"math/rand"
	"net/http"

	"github.com/radixiotech/vidya-barta/foundation/web"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Query(w http.ResponseWriter, r *http.Request) {
	random := rand.Intn(100)

	if random < 50 {
		web.Respond(
			w,
			web.Fail(web.ErrorResponse{Message: http.StatusText(http.StatusInternalServerError)}),
			http.StatusInternalServerError,
		)
		return
	}

	web.Respond(w, web.Success(http.StatusText(http.StatusOK)), http.StatusOK)
}

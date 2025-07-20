package authorize

import (
	"encoding/json"
	"errors"
	"net/http"

	dto "VK_test_proect/internal/api"
	"VK_test_proect/internal/api/responses"
	contract "VK_test_proect/internal/service/authorize"
)

type Handler struct {
	service service
}

func New(service service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) PostLogin(w http.ResponseWriter, r *http.Request) {
	var data dto.UserDTOIn

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "error", http.StatusBadRequest)

		return
	}

	in := contract.In{
		Email:    string(data.Email),
		Password: data.Password,
	}

	out, err := h.service.Authorize(r.Context(), in)
	if err != nil {
		if errors.Is(err, contract.ErrUserNotFound) {
			http.Error(w, err.Error(), http.StatusUnauthorized)

			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	res := dto.Token(out.Token)

	responses.OK(w, res, http.StatusOK)
}

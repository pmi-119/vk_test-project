package register

import (
	"encoding/json"
	"net/http"

	dto "VK_test_proect/internal/api"
	"VK_test_proect/internal/api/responses"
	contract "VK_test_proect/internal/service/register"
)

type Handler struct {
	service service
}

func New(service service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) PostRegister(w http.ResponseWriter, r *http.Request) {
	var data dto.UserDTOIn

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	in := contract.In{
		Login:    data.Login,
		Email:    data.Email,
		Password: data.Password,
	}

	user, err := h.service.RegisterUser(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	res := dto.UserDTOOut{
		Id:       user.Id,
		Login:    user.Login,
		Email:    user.Email,
		Password: user.Password,
	}

	responses.OK(w, res, http.StatusCreated)
}

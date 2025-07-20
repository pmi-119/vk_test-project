package add_product

import (
	"encoding/json"
	"log"
	"net/http"

	"VK_test_proect/internal/api"
	"VK_test_proect/internal/api/responses"
	contract "VK_test_proect/internal/service/add_product"
)

type Handler struct {
	service      service
	tokenService tokenService
}

func New(service service, tokenService tokenService) *Handler {
	return &Handler{
		service:      service,
		tokenService: tokenService,
	}
}

func (h *Handler) PostFeed(w http.ResponseWriter, r *http.Request) {
	tokenString, ok := api.GetTokenFromHeader(*r)
	if !ok {
		http.Error(w, "forbidden", http.StatusForbidden)

		return
	}

	userID, err := h.tokenService.Validate(tokenString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)

		return
	}
	var data api.ProductDTOIn

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	in := contract.In{
		Title:       data.Title,
		Description: data.Description,
		ImageURL:    data.ImageUrl,
		Price:       data.Price,
		UserID:      userID,
	}

	out, err := h.service.AddingProduct(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("RegisterUser error:", err)

		return
	}

	res := api.ProductDTOOut{
		Id:          out.Product.Id,
		UserID:      out.Product.UserID,
		Title:       out.Product.Title,
		Description: out.Product.Description,
		ImageUrl:    out.Product.ImageUrl,
		Price:       out.Product.Price,
		CreatedAt:   out.Product.CreatedAt,
	}

	responses.OK(w, res, http.StatusCreated)
}

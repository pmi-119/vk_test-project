package get_feed

import (
	"encoding/json"
	"log"
	"net/http"

	dto "VK_test_proect/internal/api"
	"VK_test_proect/internal/api/responses"
	contract "VK_test_proect/internal/service/get_feed"
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
	var data dto.FilterDTO

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "error", http.StatusBadRequest)

		return
	}

	in := contract.In{}

	if data.Paging != nil {
		in.Paging = &contract.Paging{
			Page:  data.Paging.Page,
			Limit: data.Paging.Limit,
		}
	}

	if data.Sorting != nil {
		in.Sorting = &contract.Sorting{
			Column: data.Sorting.Column,
			Order:  data.Sorting.Order,
		}
	}

	if data.PriceFilter != nil {
		in.PriceFilter = &contract.PriceFilter{
			Max: data.PriceFilter.Max,
			Min: data.PriceFilter.Min,
		}
	}

	tokenString, ok := dto.GetTokenFromHeader(*r)
	if ok {
		userID, err := h.tokenService.Validate(tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)

			return
		}

		in.UserID = &userID
	}

	out, err := h.service.GetFeed(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("RegisterUser error:", err)

		return
	}

	res := make([]dto.ProductDTOInfo, 0, len(out.Items))

	for _, item := range out.Items {
		res = append(res, dto.ProductDTOInfo{
			Title:         item.Title,
			Description:   item.Description,
			ImageUrl:      item.ImageUrl,
			Price:         item.Price,
			UserLogin:     item.UserLogin,
			IsCurrentUser: item.IsCurrentUser,
		})
	}

	responses.OK(w, res, http.StatusOK)
}

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
	service service
}

func New(service service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) PostFeed(w http.ResponseWriter, r *http.Request) {
	var data dto.FilterDTO

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "error", http.StatusBadRequest)

		return
	}

	in := contract.In{
		PriceFilter: &contract.PriceFilter{
			Min: data.PriceFilter.Min,
			Max: data.PriceFilter.Max,
		},
		Sorting: &contract.Sorting{
			SortingByColumn: data.Sorting.SortingByColumn,
			SortingOrder:    data.Sorting.SortingOrder,
		},
		Paging: &contract.Paging{
			Page:  data.Paging.Page,
			Limit: data.Paging.Limit,
		},
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
			Title:       item.Title,
			Description: item.Description,
			ImageUrl:    item.ImageUrl,
			Price:       item.Price,
			UserLogin:   item.UserLogin,
		})
	}

	responses.OK(w, res, http.StatusCreated)
}

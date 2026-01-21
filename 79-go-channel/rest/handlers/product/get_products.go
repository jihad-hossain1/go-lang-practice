package product

import (
	"ecom/domain"
	"ecom/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {

	reqQuery := r.URL.Query()
	pageAsStr := reqQuery.Get("page")
	limitAsStr := reqQuery.Get("limit")

	page, _ := strconv.ParseInt(pageAsStr, 10, 32)
	limit, _ := strconv.ParseInt(limitAsStr, 10, 32)

	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	// products are using go channel
	prdCh := make(chan []*domain.Product)
	go func() {
		productList, err := h.svc.List(page, limit)
		if err != nil {
			util.SendError(w, http.StatusInternalServerError, "Internal server error")
			return
		}

		prdCh <- productList
	}()

	cntCh := make(chan int64)
	go func() {
		cnt, err := h.svc.Count()
		if err != nil {
			util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		cntCh <- cnt

	}()

	pdctList := <-prdCh
	cnt := <-cntCh

	util.SendPage(w, pdctList, page, limit, cnt)
}

package controllers

import (
	"encoding/json"
	meliModels "meli-challenge-proxyapi-golang/meli/models"
	meliService "meli-challenge-proxyapi-golang/meli/services"
	"meli-challenge-proxyapi-golang/models"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

func GetItemsByQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	params := mux.Vars(r)
	q := params["q"]

	// fmt.Println("Query string key value", q)

	// if q == "" {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte("Query invalid")) //CUSTOM ERROR RESPONSE
	// 	return
	// }

	var meliItemsResponse = meliService.GetItemsByQuery(q)

	var items []models.Item
	categories := []string{}

	if len(meliItemsResponse.Filters) > 0 {
		for _, category := range meliItemsResponse.Filters[0].Values[0].PathFromRoot {
			categories = append(categories, category.Name)
		}
	}

	var wg sync.WaitGroup

	for _, result := range meliItemsResponse.Results {
		wg.Add(1)
		go func(r meliModels.Results) {
			defer wg.Done()

			var descriptionResponse = meliService.GetItemDescriptionById(r.Id)
			var currencyResponse = meliService.GetCurrencyById(r.CurrencyId)

			var item = models.Item{
				Id:    r.Id,
				Title: r.Title,
				Price: models.Price{
					Currency: r.CurrencyId,
					Amount:   r.Price,
					Decimals: currencyResponse.DecimalPlaces,
				},
				Picture:      r.Thumbnail,
				Condition:    r.Condition,
				FreeShipping: r.Shipping.FreeShipping,
				SoldQuantity: r.SoldQuantity,
				Description:  descriptionResponse.PlainText,
			}

			items = append(items, item)
		}(result)
	}

	wg.Wait()

	var ItemsResponse = models.ItemsResponse{
		Categories: categories,
		Items:      items,
	}

	j, _ := json.Marshal(ItemsResponse)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(j))
}

func GetItemById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	params := mux.Vars(r)

	var id = params["id"]

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid id")) //CUSTOM ERROR RESPONSE
		return
	}

	var itemResponse = meliService.GetItemById(id)
	var descriptionResponse = meliService.GetItemDescriptionById(id)
	var currencyResponse = meliService.GetCurrencyById(itemResponse.CurrencyID)

	var item = models.Item{
		Id:    itemResponse.Id,
		Title: itemResponse.Title,
		Price: models.Price{
			Currency: itemResponse.CurrencyID,
			Amount:   itemResponse.Price,
			Decimals: currencyResponse.DecimalPlaces,
		},
		Picture:      itemResponse.Thumbnail,
		Condition:    itemResponse.Condition,
		FreeShipping: itemResponse.Shipping.FreeShipping,
		SoldQuantity: itemResponse.SoldQuantity,
		Description:  descriptionResponse.PlainText,
	}

	response := models.ItemResponse{
		Item: item,
	}

	j, _ := json.Marshal(response)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(j))
}

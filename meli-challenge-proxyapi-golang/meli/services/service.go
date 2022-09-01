package services

import (
	"encoding/json"
	"io"
	"log"
	meliModels "meli-challenge-proxyapi-golang/meli/models"
	"net/http"
)

const (
	MELI_API_URL string = "https://api.mercadolibre.com"
)

func GetItemById(pId string) meliModels.MeliItemResponse {
	//defer utils.TimeTrack(time.Now())

	response, error := http.Get(MELI_API_URL + "/items/" + pId)

	if error != nil {
		log.Fatal(error)
	}

	if response.Body != nil {
		defer response.Body.Close()
	}

	data, error := io.ReadAll(response.Body)

	if error != nil {
		log.Fatal(error)
	}

	var item meliModels.MeliItemResponse
	json.Unmarshal(data, &item)

	return item
}

func GetItemDescriptionById(pId string) meliModels.MeliDescriptionResponse {
	//defer utils.TimeTrack(time.Now())

	response, error := http.Get(MELI_API_URL + "/items/" + pId + "/description")

	if error != nil {
		log.Fatal(error)
	}

	if response.Body != nil {
		defer response.Body.Close()
	}

	data, error := io.ReadAll(response.Body)

	if error != nil {
		log.Fatal(error)
	}

	var description meliModels.MeliDescriptionResponse
	json.Unmarshal(data, &description)

	return description
}

func GetCurrencyById(pCurrencyId string) meliModels.MeliCurrencyResponse {
	//defer utils.TimeTrack(time.Now())

	response, error := http.Get(MELI_API_URL + "/currencies/" + pCurrencyId)

	if error != nil {
		log.Fatal(error)
	}

	if response.Body != nil {
		defer response.Body.Close()
	}

	data, error := io.ReadAll(response.Body)

	if error != nil {
		log.Fatal(error)
	}

	var currency meliModels.MeliCurrencyResponse
	json.Unmarshal(data, &currency)

	return currency
}

func GetItemsByQuery(pQuery string) meliModels.MeliItemsResponse {
	//defer utils.TimeTrack(time.Now())

	response, error := http.Get(MELI_API_URL + "/sites/MLU/search?q=" + pQuery + "&limit=5")

	if error != nil {
		log.Fatal(error)
	}

	if response.Body != nil {
		defer response.Body.Close()
	}

	data, error := io.ReadAll(response.Body)

	if error != nil {
		log.Fatal(error)
	}

	var items meliModels.MeliItemsResponse
	json.Unmarshal(data, &items)

	return items
}

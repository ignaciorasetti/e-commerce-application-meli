package models

import "time"

type Results struct {
	Id         string  `json:"id"`
	SiteId     string  `json:"siteId"`
	Price      float64 `json:"price"`
	Title      string  `json:"title"`
	CurrencyId string  `json:"currency_id"`
	Thumbnail  string  `json:"thumbnail"`

	AcceptsMercadopago string `json:"accepts_mercadopago"`
	Address            struct {
		CityId    string `json:"city_id"`
		CityName  string `json:"city_name"`
		StateId   string `json:"state_id"`
		StateName string `json:"state_name"`
	} `json:"address"`
	Attributes        interface{} `json:"attributes"`
	AvailableQuantity int         `json:"available_quantity"`
	SoldQuantity      int         `json:"sold_quantity"`
	Shipping          struct {
		FreeShipping bool        `json:"free_shipping"`
		LogisticType string      `json:"logistic_type"`
		Mode         string      `json:"mode"`
		StorePickUp  bool        `json:"store_pick_up"`
		Tags         interface{} `json:"tags"`
	} `json:"shipping"`
	Condition string `json:"condition"`
}

type PathFromRoot struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type MeliItemsResponse struct {
	SiteId                 string      `json:"site_id"`
	CountryDefaultTimeZone string      `json:"country_default_time_zone"`
	Query                  string      `json:"query"`
	Paging                 interface{} `json:"paging"`
	Results                []Results   `json:"results"`
	Sort                   struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"sort"`
	AvailableSorts interface{} `json:"available_sorts"`
	Filters        []struct {
		Id     string `json:"id"`
		Name   string `json:"name"`
		Type   string `json:"type"`
		Values []struct {
			Id           string         `json:"id"`
			Name         string         `json:"name"`
			PathFromRoot []PathFromRoot `json:"path_from_root"`
		} `json:"values"`
	} `json:"filters"`
	AvailableFilters float64 `json:"available_filters"`
}

type MeliItemResponse struct {
	Id                string      `json:"id"`
	SiteID            string      `json:"site_id"`
	Title             string      `json:"title"`
	Subtitle          interface{} `json:"subtitle"`
	SellerID          int         `json:"seller_id"`
	CategoryID        string      `json:"category_id"`
	OfficialStoreID   int         `json:"official_store_id"`
	Price             float64     `json:"price"`
	BasePrice         float64     `json:"base_price"`
	OriginalPrice     interface{} `json:"original_price"`
	CurrencyID        string      `json:"currency_id"`
	InitialQuantity   int         `json:"initial_quantity"`
	AvailableQuantity int         `json:"available_quantity"`
	SoldQuantity      int         `json:"sold_quantity"`
	SaleTerms         []struct {
		ID          string      `json:"id"`
		Name        string      `json:"name"`
		ValueID     string      `json:"value_id"`
		ValueName   string      `json:"value_name"`
		ValueStruct interface{} `json:"value_struct"`
		Values      []struct {
			ID     string      `json:"id"`
			Name   string      `json:"name"`
			Struct interface{} `json:"struct"`
		} `json:"values"`
	} `json:"sale_terms"`
	BuyingMode      string    `json:"buying_mode"`
	ListingTypeID   string    `json:"listing_type_id"`
	StartTime       time.Time `json:"start_time"`
	StopTime        time.Time `json:"stop_time"`
	Condition       string    `json:"condition"`
	Permalink       string    `json:"permalink"`
	ThumbnailID     string    `json:"thumbnail_id"`
	Thumbnail       string    `json:"thumbnail"`
	SecureThumbnail string    `json:"secure_thumbnail"`
	Pictures        []struct {
		ID        string `json:"id"`
		URL       string `json:"url"`
		SecureURL string `json:"secure_url"`
		Size      string `json:"size"`
		MaxSize   string `json:"max_size"`
		Quality   string `json:"quality"`
	} `json:"pictures"`
	VideoID      interface{} `json:"video_id"`
	Descriptions []struct {
		ID string `json:"id"`
	} `json:"descriptions"`
	AcceptsMercadopago           bool          `json:"accepts_mercadopago"`
	NonMercadoPagoPaymentMethods []interface{} `json:"non_mercado_pago_payment_methods"`
	Shipping                     struct {
		Mode        string `json:"mode"`
		FreeMethods []struct {
			ID   int `json:"id"`
			Rule struct {
				Default          bool        `json:"default"`
				FreeMode         string      `json:"free_mode"`
				FreeShippingFlag bool        `json:"free_shipping_flag"`
				Value            interface{} `json:"value"`
			} `json:"rule"`
		} `json:"free_methods"`
		Tags         []string    `json:"tags"`
		Dimensions   interface{} `json:"dimensions"`
		LocalPickUp  bool        `json:"local_pick_up"`
		FreeShipping bool        `json:"free_shipping"`
		LogisticType string      `json:"logistic_type"`
		StorePickUp  bool        `json:"store_pick_up"`
	} `json:"shipping"`
	InternationalDeliveryMode string `json:"international_delivery_mode"`
	SellerAddress             struct {
		City struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"city"`
		State struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"state"`
		Country struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"country"`
		SearchLocation struct {
			Neighborhood struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"neighborhood"`
			City struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"city"`
			State struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"state"`
		} `json:"search_location"`
		ID int `json:"id"`
	} `json:"seller_address"`
	SellerContact interface{} `json:"seller_contact"`
	Location      struct {
	} `json:"location"`
	CoverageAreas []interface{} `json:"coverage_areas"`
	Attributes    []struct {
		ID          string      `json:"id"`
		Name        string      `json:"name"`
		ValueID     string      `json:"value_id"`
		ValueName   string      `json:"value_name"`
		ValueStruct interface{} `json:"value_struct"`
		Values      []struct {
			ID     string      `json:"id"`
			Name   string      `json:"name"`
			Struct interface{} `json:"struct"`
		} `json:"values"`
		AttributeGroupID   string `json:"attribute_group_id"`
		AttributeGroupName string `json:"attribute_group_name"`
	} `json:"attributes"`
	Warnings      []interface{} `json:"warnings"`
	ListingSource string        `json:"listing_source"`
	Variations    []struct {
		ID                    int64   `json:"id"`
		Price                 float64 `json:"price"`
		AttributeCombinations []struct {
			ID          string      `json:"id"`
			Name        string      `json:"name"`
			ValueID     string      `json:"value_id"`
			ValueName   string      `json:"value_name"`
			ValueStruct interface{} `json:"value_struct"`
			Values      []struct {
				ID     string      `json:"id"`
				Name   string      `json:"name"`
				Struct interface{} `json:"struct"`
			} `json:"values"`
		} `json:"attribute_combinations"`
		AvailableQuantity int           `json:"available_quantity"`
		SoldQuantity      int           `json:"sold_quantity"`
		SaleTerms         []interface{} `json:"sale_terms"`
		PictureIds        []string      `json:"picture_ids"`
		CatalogProductID  interface{}   `json:"catalog_product_id"`
	} `json:"variations"`
	Status              string        `json:"status"`
	SubStatus           []interface{} `json:"sub_status"`
	Tags                []string      `json:"tags"`
	Warranty            string        `json:"warranty"`
	CatalogProductID    interface{}   `json:"catalog_product_id"`
	DomainID            string        `json:"domain_id"`
	ParentItemID        interface{}   `json:"parent_item_id"`
	DifferentialPricing interface{}   `json:"differential_pricing"`
	DealIds             []interface{} `json:"deal_ids"`
	AutomaticRelist     bool          `json:"automatic_relist"`
	DateCreated         time.Time     `json:"date_created"`
	LastUpdated         time.Time     `json:"last_updated"`
	Health              int           `json:"health"`
	CatalogListing      bool          `json:"catalog_listing"`
	Channels            []string      `json:"channels"`
}

type MeliDescriptionResponse struct {
	Text          string `json:"text"`
	PlainText     string `json:"plain_text"`
	Description   string `json:"last_updated"`
	DecimalPlaces int    `json:"date_created"`
	Snapshot      int    `json:"snapshot"`
}

type MeliCurrencyResponse struct {
	Id            string `json:"id"`
	Symbol        string `json:"name"`
	Description   string `json:"description"`
	DecimalPlaces int    `json:"decimal_places"`
}

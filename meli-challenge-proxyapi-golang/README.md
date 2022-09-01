# MercadoLibre Challenge Proxy API - Golang

##### *Disclaimer: This project contained on this repository is for the only propose of education proposes. You shouldn't use this for any kind besides the reasons listed before. For a production environment you must use the [MercadoLibre API](https://api.mercadolibre.com/) and follow the [Mercado Libre API: Terms & Conditions](https://developers.mercadolibre.com.uy/es_ar/terminos-y-condiciones)*.

<div align="center">
    <table cellpadding="5">
        <tbody align="center">
            <tr>
                <td style="width:30%">
                    <a href="https://bit.ly/awesome-go-doppler">
                        <img src="https://i.imgur.com/ptbaGiH.png" width="200" alt="Doppler"><br/>
                    </a>
                </td>
                 <td style="width:30%">
                    <a href="https://bit.ly/awesome-go-doppler">
                        <img src="https://http2.mlstatic.com/frontend-assets/ml-web-navigation/ui-navigation/5.19.5/mercadolibre/logo__large_plus.png" width="200" alt="Doppler"><br/>
                    </a>
                </td>
            </tr>
        </tbody>
    </table>
</div>

##### *Gopher image by Nats Romanova, licensed under [Creative Commons 3.0 Attributions license][cc3-by].*

##### *Mercado Libre logo by Mercado Libre LLC under [Mercado Libre Terms & Conditions/9-Interactual property][meli-term-and-conditions].*

## Introduction
This was a project made on [Golang](https://go.dev/) for the propose of interact with [MercadoLibre API](https://api.mercadolibre.com/) and re-shape the data to create custom responses from it.

## Installation
* Clone repo
* `go run .`

## Endpoints

### GET /items/:id
Gets item data from a specific Item ID.

**Parameters**

| Name | Required | Type | Description                                                                                                                                                           |
| -------------:|:--------:|:-------:| --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `id` | required | string  | The item for which to perform the action. <br/><br/>

### GET items?q=:query
Gets items from a specific query parameter. 

**Parameters**

| Name | Required | Type | Description                                                                                                                                                           |
| -------------:|:--------:|:-------:| --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `query` | required | string  | Query term of the search. <br/><br/>

[cc3-by]: https://creativecommons.org/licenses/by/3.0/
[meli-term-and-conditions]: https://www.mercadolibre.com.ar/ayuda/terminos-y-condiciones-de-uso_991
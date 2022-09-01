# MercadoLibre Challenge Proxy API - NodeJS

##### *Disclaimer: This project contained on this repository is for the only propose of education proposes. You shouldn't use this for any kind besides the reasons listed before. For a production environment you must use the [MercadoLibre API](https://api.mercadolibre.com/) and follow the [Mercado Libre API: Terms & Conditions](https://developers.mercadolibre.com.uy/es_ar/terminos-y-condiciones)*.

<div align="center">
    <table cellpadding="5">
        <tbody align="center">
            <tr>
                <td style="width:30%">
                    <a href="https://bit.ly/awesome-go-doppler">
                        <img src="https://nodejs.org/static/images/logos/nodejs-new-pantone-white.svg" width="200" alt="Doppler"><br/>
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

##### *Mercado Libre logo by Mercado Libre LLC under [Mercado Libre Terms & Conditions/9-Interactual property][meli-term-and-conditions].*

## Introduction
This was a project made on [NodeJS](https://nodejs.org/es/) for the propose of interact with [MercadoLibre API](https://api.mercadolibre.com/) and re-shape the data to create custom responses from it.

## Installation
* Clone repo
* `npm install`
* `npm run dev`, to use [Nodemon](https://www.npmjs.com/package/nodemon) or `npm start`

## Endpoints

### GET /items/:id
Gets item data from a specific Item ID.

**Parameters**

| Name | Required | Type | Description                                                                                                                                                           |
| -------------:|:--------:|:-------:| --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `id` | required | string  | The item for which to perform the action. <br/><br/>

### GET /items?q=:query
Gets items from a specific query parameter. 

**Parameters**

| Name | Required | Type | Description                                                                                                                                                           |
| -------------:|:--------:|:-------:| --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `query` | required | string  | Query term of the search. <br/><br/>

## Dependencies
* [Express](http://expressjs.com/): Web framework to handle API routes. 
* [Cors](https://www.npmjs.com/package/cors): Middleware to handle CORS for Express.
* [Axios](https://axios-http.com/): To perform Http requests to the MercadoLibre API.
* [Nodemon](https://www.npmjs.com/package/nodemon): Tool to help developing, hot-reloading the project. 

[meli-term-and-conditions]: https://www.mercadolibre.com.ar/ayuda/terminos-y-condiciones-de-uso_991
const express = require('express');
const cors = require('cors');

const { response, request } = require("express");

class Server {
    constructor(){
        this.app = express();
        
        this.middleWares();

        this.routes();
    }

    middleWares(){
        this.app.use(cors());
        this.app.use(express.json());
        this.app.use(express.urlencoded({ 
            extended: true 
        }));

        this.app.use('/api', require('./routes'))
    }

    routes(){
        this.app.get("/", (pRequest, pResponse) => {
            pResponse.status(200).send(
                "Mercado Libre Challenge API by Igna"
            );
        })

        this.app.get("*", (pRequest, pResponse) => {
            pResponse.status(404).json({
                message: "ROUTE_NOT_FOUND"
            })
        })
    }

    listen(){
        const port = process.env.PORT || 3001;

        this.app.listen(port, () => {
            console.log(`API -> Running http://localhost:${port}`)
        })
    }
}

module.exports = Server;
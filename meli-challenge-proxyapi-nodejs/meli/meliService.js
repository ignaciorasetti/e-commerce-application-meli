const axios = require('axios').default;

class MeliService {
    constructor() {
        this.instance = axios.create({
            baseURL: 'https://api.mercadolibre.com/',
        });
    }

    async getItemDataById(pId) {
        try {
            const [
                { data: item }, 
                { data: description }
            ] = await Promise.all([
                this.instance.get(`items/${pId}`),
                this.instance.get(`items/${pId}/description`)
            ]);

            return { item, description };
            
        }catch (pError) {
            throw pError
        }
    }

    async getCurrencyById(pCurrencyId) {
        try {
            const { data } = await this.instance.get(`currencies/${pCurrencyId}`);

            return data;
        }catch (pError) {
            throw pError
        }
    }

    async getItemsDataByQuery(q , limit = 5) {
        try {              
            const { data: queryResult } = await this.instance.get(`/sites/MLU/search?q=${q}`, {
                params: {
                    limit: limit
                }
            });

            return queryResult
        }catch (pError) {
            throw pError
        }
    }
}

module.exports = new MeliService();

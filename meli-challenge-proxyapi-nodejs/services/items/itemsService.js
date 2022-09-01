const MeliService = require('../../meli/meliService')

class ItemsService {

    constructor(){}

    async getItemDataById(pId) {    
        try{
            const { item, description } = await MeliService.getItemDataById(pId);
            const { decimal_places } = await MeliService.getCurrencyById(item.currency_id)

            return { 
                item,
                description,
                decimal_places 
            }
        }catch(pError) {
            throw (pError)
        }
    }

    async getItemsDataByQuery(q) {    
        try{
            const queryResult = await MeliService.getItemsDataByQuery(q);
        
            const filters = queryResult.filters
            const categories = filters[0] && filters[0].values[0].path_from_root.map(category => category.name) || {};

            const items = await Promise.all(
                queryResult.results.map(async item => {
                    const currency = await MeliService.getCurrencyById(item.currency_id);

                    return {
                        id: item.id,
                        title: item.title,
                        price: {
                            amount: item.price,
                            currency: item.currency_id,
                            decimals: currency.decimal_places
                        },
                        picture: item.thumbnail,
                        condition: item.condition,
                        free_shipping: item.shipping.free_shipping
                    }
                })
            )

            return { categories, items }
        }catch(pError) {
            console.error(pError)
        }
    }
}

module.exports = new ItemsService();
import {
    ItemResponse,
    ItemsResponse,
} from "../../models/interfaces/ItemsResponse";
import api from "../baseApi";

/**
 * Get items from an specific query
 * @param id
 * @returns {Promise<ItemResponse>} A promise that contains the item data from specific id
 */
export const getItemById = async (id: string): Promise<ItemResponse> => {
    const { data } = await api.get("/items/" + id);

    return data;
};

/**
 * Get items from an specific query
 * @param query
 * @returns {Promise<ItemsResponse>} A promise that contains the items data from specific query
 */
export const getItemsById = async (query: string): Promise<ItemsResponse> => {
    const { data } = await api.get("/items?q=" + query);

    return data;
};

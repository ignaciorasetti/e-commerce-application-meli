import { useQuery, UseQueryResult } from "react-query";
import { useLocation } from "react-router-dom";

import { BreadCrumbs, ItemList, LoadingDots, NoItemsFound } from '../../components'
import useSEO from "../../hooks/useSEO";

import { getItemsById } from "../../api/ItemsService/itemsService";
import { ItemsResponse } from "../../models/interfaces/ItemsResponse";

import meli_mobile_logo from '../../assets/meli_mobile_logo.png';

import "./Items.scss";

export default function ItemsPage() {
    const location = useLocation();
    const query = new URLSearchParams(location.search);

    const searchValue = query.get("search") || "";

    const {
        data,
        isError,
        error,
        isLoading,
    }: UseQueryResult<ItemsResponse, Error> = useQuery<
        ItemsResponse,
        Error,
        ItemsResponse
    >(["item", searchValue], () => getItemsById(searchValue));

    const title = `${ searchValue } | MercadoLibre 📦`
    const description = `Encontrá ${searchValue} que buscas en Mercado Libre Uruguay. Tenemos miles de productos a precios increíbles. ¡Comprá ahora!`
    const ogImage = meli_mobile_logo

    useSEO({  
        title,
        description,
        ogTitle: title,
        ogDescription: description,
        ogImage: ogImage
    })

    if (isLoading) {
        return <LoadingDots />;
    }

    if (isError) {
        return <div>Error is: {error.message}</div>;
    }

    return (
        <main className="mt-5 px-2 mx-auto max-w-5xl">
            {data && data.categories?.length > 0 && (
                <BreadCrumbs categories={data?.categories} />
            )}

            {data && data.items?.length > 0 ? (
                <ItemList items={data?.items} />
            ) : (
                <NoItemsFound />
            )}
        </main>
    );
}

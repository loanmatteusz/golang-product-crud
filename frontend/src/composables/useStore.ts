import type { StorePreview, StoreResponse } from "~/types/store";


export const useStore = () => {
    const config = useRuntimeConfig();

    const create = async (data: Partial<StorePreview>) => {
        const { error } = await useFetch("/stores", {
            method: 'POST',
            body: data,
            baseURL: config.public.apiUrl,
        });

        if (error.value) {
            throw new Error(`Erro ao criar loja: ${error.value.message}`);
        }
    };

    const list = async () => {
        const { data, error } = await useFetch<StoreResponse[]>("/stores", {
            method: 'GET',
            baseURL: config.public.apiUrl,
        });

        if (error.value) {
            throw new Error(error.value.message)
        }

        const stroes = Array.from(data.value || []);

        return stroes.map((store): StorePreview => ({
            id: store.ID,
            name: store.Name,
            legal_name: store.LegalName,
            number: store.Number,
            address: {
                street: store.Address?.Street || '',
                number: store.Address?.Number || '',
                city: store.Address?.City || '',
                state: store.Address?.State || '',
                zip_code: store.Address?.ZipCode || '',
            },
            establishment_id: store.EstablishmentID,
        }));
    }

    const update = async (id: string, updateData: Partial<StorePreview>) => {
        const { error } = await useFetch(`/stores/${id}`, {
            method: 'PUT',
            body: updateData,
            baseURL: config.public.apiUrl,
        });

        if (error.value) {
            throw new Error(`Erro ao atualizar estabelecimento: ${error.value.message}`);
        }
    }

    const remove = async (id: string) => {
        const { error } = await useFetch(`/establishments/${id}`, {
            method: 'DELETE',
            baseURL: config.public.apiUrl,
        });

        if (error.value) {
            throw new Error(`Erro ao deletar loja: ${error.value.message}`);
        }
    }

    return {
        create,
        list,
        update,
        remove,
    }
}

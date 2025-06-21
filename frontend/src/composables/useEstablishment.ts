import type { EstablishmentPreview, EstablishmentResponse } from "~/types/establishment";

export const useEstablishment = () => {

    const API_URL = process.env.API_URL || 'http://localhost:3333/establishments';

    const create = async (data: Omit<EstablishmentPreview, 'id'>) => {
        const { error } = await useFetch(`${API_URL}`, {
            method: 'POST',
            body: data,
        });

        if (error.value) {
            throw new Error(`Erro ao criar estabelecimento: ${error.value.message}`);
        }
    }

    const list = async () => {
        const { data, error } = await useFetch<EstablishmentResponse[]>(`${API_URL}`, {
            method: 'GET',
        });

        if (error.value) {
            throw new Error(error.value.message)
        }

        const establishments = Array.from(data.value || []);

        return establishments.map((establishment): EstablishmentPreview => ({
            id: establishment.ID,
            name: establishment.Name,
            legal_name: establishment.LegalName,
            number: establishment.Number,
            address: {
                street: establishment.Address?.Street || '',
                number: establishment.Address?.Number || '',
                city: establishment.Address?.City || '',
                state: establishment.Address?.State || '',
                zip_code: establishment.Address?.ZipCode || '',
            },
        }));
    }

    const update = async (id: string, updateData: Partial<EstablishmentPreview>) => {
        const { error } = await useFetch(`${API_URL}/${id}`, {
            method: 'PUT',
            body: updateData,
        });

        if (error.value) {
            throw new Error(`Erro ao atualizar estabelecimento: ${error.value.message}`);
        }
    }

    const remove = async (id: string) => {
        const { error } = await useFetch(`${API_URL}/${id}`, {
            method: 'DELETE',
        });

        if (error.value) {
            throw new Error(`Erro ao deletar estabelecimento: ${error.value.message}`);
        }
    }

    return {
        create,
        list,
        update,
        remove,
    }
}

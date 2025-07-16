import type { Product } from "~/types/product";

export const useProduct = () => {
    const products: Product[] = [
        {
            ID: '1',
            Name: 'Product Test 1',
            Price: 1200,
            CategoryID: '1',
            CreatedAt: new Date(),
            UpdatedAt: new Date(),
        },
    ]; // Temporário / Estatico
    // const config = useRuntimeConfig();

    const create = async (data: Omit<Product, 'id'>) => {
        products.push(data);
    }

    const list = async () => {
        return products;
    }

    const update = async (id: string, updateData: Partial<Product>) => {
        console.log({ updateData });
    }

    const remove = async (id: string) => {
        const index = products.findIndex(p => p.ID === id);

        if (index !== -1) {
            products.splice(index, 1);
        }
    }

    return {
        create,
        list,
        update,
        remove,
    }
}

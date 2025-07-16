<script setup lang="ts">
    setPageLayout('private');
    useHead({
        title: 'Produtos',
    });

    import { useProduct } from '~/composables/useProducts';

    const { list } = useProduct();
    const { data: listProducts, error: listProductsError } = useAsyncData('products', async () => await list());
    if (listProductsError.value) {
        console.error("Erro to load product list.");
    }

    console.log({ listProducts: listProducts.value });
</script>

<template>
    <div class="w-full flex flex-col">
        <div class="w-full flex gap-4 items-center justify-center">
            <h1 class="font-bold">Produtos</h1>
        </div>

        <div class="flex flex-col gap-6">
            Lista de Produtos
        </div>
    </div>
</template>

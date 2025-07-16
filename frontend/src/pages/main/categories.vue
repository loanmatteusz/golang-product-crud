<script setup lang="ts">
    setPageLayout('private');
    useHead({
        title: 'Categorias',
    });

    import { onMounted, ref } from 'vue';
    import { columns } from '~/components/categories/columns';
    import DataTable from '~/components/categories/data-table.vue';
    import type { Category } from '~/types/category';

    const data = ref<Category[]>([]);

    async function getData(): Promise<Category[]> {
        return [
        {
                ID: '728ed52f',
                Name: 'Category Test',
                CreatedAt: new Date(),
                UpdatedAt: new Date(), 
            },
            {
                ID: '728ed52f',
                Name: 'Category Test',
                CreatedAt: new Date(),
                UpdatedAt: new Date(), 
            },
        ];
    }

    onMounted(async () => {
        data.value = await getData();
    });

</script>


<!-- TEMPLATE -->

<template>
    <div class="w-full flex flex-col">
        <div class="w-full flex gap-4 items-center justify-center">
            <h1 class="font-bold">Categorias</h1>
        </div>
        <div class="container py-10 mx-auto">
            <DataTable :columns="columns" :data="data" />
        </div>
    </div>
</template>

<script setup lang="ts">
    setPageLayout('private');
    useHead({
        title: 'Categorias',
    });

    import { onMounted } from 'vue';
    import { columns } from '~/components/categories/columns';
    import DataTable from '~/components/categories/data-table.vue';

    const page = ref(1);
    const limit = ref(10);

    const { categories, pagination, list, loading, error } = useCategories(page, limit);

    onMounted(() => {
        list();
    });

    const safeCategories = computed(() => categories.value ?? []);

    function goToPage(newPage: number) {
        page.value = newPage;
    }
</script>


<!-- TEMPLATE -->

<template>
    <div class="w-full flex flex-col items-center justify-center">
        <div class="w-full flex gap-4 items-center justify-center">
            <h1 class="font-bold">Categorias</h1>
        </div>
        <div class="w-full py-10">
            <DataTable class="w-full" v-if="!loading" :columns="columns" :data="safeCategories" />
            <div class="w-full flex items-center justify-center" v-if="loading">Carregando...</div>
            <div v-if="error" class="text-red-600">{{ error }}</div>
        </div>
        <div class="w-full">
            <AppTablePagination
                class="flex justify-end"
                :current-page="page"
                :total-items="pagination?.total || 0"
                :items-per-page="limit"
                @update:current-page="goToPage"
            />
        </div> 
    </div>
</template>

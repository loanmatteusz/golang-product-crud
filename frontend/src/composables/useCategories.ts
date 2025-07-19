import type { Category } from "~/types/category";


export const useCategories = (page: Ref<number>, limit: Ref<number>) => {
    const config = useRuntimeConfig();

    const {
        data: response,
        pending: loading,
        error,
        refresh: list,
    } = useAsyncData<CategoryListResponse>(
        () => `categories-page-${page.value}-limit-${limit.value}`,
        async () => {
            const { data } = await useFetch<CategoryListResponse>(`/categories?page=${page.value}&limit${limit.value}`, {
                method: 'GET',
                baseURL: config.public.apiUrl,
            });
            return data?.value!;
        },
        {
            watch: [page, limit],
        }
    );

    const categories = computed(() => response.value?.data || []);
    const pagination = computed(() => response.value?.pagination);

    return {
        categories,
        pagination,
        loading,
        error,
        list,
    }
}

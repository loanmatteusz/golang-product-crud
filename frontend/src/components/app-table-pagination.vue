<script setup lang="ts">
  import {
    Pagination,
    PaginationContent,
    PaginationItem,
    PaginationNext,
    PaginationPrevious,
    PaginationEllipsis,
  } from '@/components/ui/pagination';

  const props = defineProps<{
    currentPage: number;
    totalItems: number;
    itemsPerPage: number;
  }>();

  const emit = defineEmits<{
    (e: 'update:current-page', page: number): void;
  }>();

  const totalPages = computed(() => Math.ceil(props.totalItems / props.itemsPerPage));

  function changePage(newPage: number) {
    if (newPage >= 1 && newPage <= totalPages.value) {
      emit('update:current-page', newPage);
    }
  }
</script>

<template>
  <Pagination :page="props.currentPage" :items-per-page="props.itemsPerPage" :total="props.totalItems" v-slot="{ page, pageCount }">
    <PaginationContent>
      <PaginationPrevious @click="changePage(page - 1)" />

      <template v-for="item in Array.from({ length: pageCount }, (_, i) => i + 1)" :key="item">
        <PaginationItem
          :value="item"
          :is-active="item === page"
          @click="changePage(item)"
        >
          {{ item }}
        </PaginationItem>
      </template>

      <PaginationEllipsis v-if="pageCount > 5" :index="4" />

      <PaginationNext @click="changePage(page + 1)" />
    </PaginationContent>
  </Pagination>
</template>

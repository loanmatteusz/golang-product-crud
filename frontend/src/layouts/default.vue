<script setup lang="ts">
import { ref } from 'vue'
import type { NavigationMenuItem } from '@nuxt/ui'

const items = ref<NavigationMenuItem[][]>([
    [
        {
            label: 'GitHub',
            icon: 'i-simple-icons-github',
            to: 'https://github.com/loanmatteusz',
            target: '_blank'
        }
    ],
    [
        {
            label: 'Links',
            type: 'label'
        },
        {
            label: 'Establishment',
            icon: 'i-lucide-building',
            to: '/establishment'
        },
        {
            label: 'Store',
            icon: 'i-lucide-store',
            to: '/store'
        },
    ],
])
</script>

<template>
  <div class="flex min-h-screen">
    <aside class="w-52 border-r border-gray-200 py-6">
      <NuxtNavigationMenu orientation="vertical" :items="items" class="data-[orientation=vertical]:w-48" />
    </aside>
    <main class="flex-1 p-4">
      <NuxtPage />
    </main>
  </div>
</template>

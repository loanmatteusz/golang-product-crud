<script setup lang="ts">
  import { SidebarProvider, SidebarTrigger } from "@/components/ui/sidebar";
  import AppSidebar from "@/components/app-sidebar.vue";
</script>

<template>
  <div class="flex min-h-screen">
    <SidebarProvider>
      <AppSidebar />
      <main class="flex-1 p-4">
        <SidebarTrigger class="cursor-pointer" />
        <slot />
      </main>
    </SidebarProvider>
  </div>
</template>

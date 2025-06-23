<script setup lang="ts">
    import { Warehouse, Store, FileSearch } from "lucide-vue-next"
    import {
    Sidebar,
    SidebarContent,
    SidebarGroup,
    SidebarGroupContent,
    SidebarGroupLabel,
    SidebarMenu,
    SidebarMenuButton,
    SidebarMenuItem,
    } from "@/components/ui/sidebar"

    const items = [
      {
          title: "Establishments",
          url: "establishment",
          icon: Warehouse,
      },
      {
          title: "Store",
          url: "store",
          icon: Store,
      },
    ];

    const github = {
      title: "Github",
      url: "http://github.com/loanmatteusz/challenge-snet",
      icon: FileSearch,
    }
</script>

<template>
  <Sidebar>
    <SidebarContent>
      <SidebarGroup>
        <SidebarGroupLabel>Application</SidebarGroupLabel>
        <SidebarGroupContent>
          <SidebarMenu>
              <SidebarMenuItem v-for="item in items" :key="item.title">
                <SidebarMenuButton asChild>
                    <NuxtLink :to="{name: item.url}" class="cursor-pointer">
                      <component :is="item.icon" />
                      <span>{{item.title}}</span>
                    </NuxtLink>
                </SidebarMenuButton>
              </SidebarMenuItem>
          </SidebarMenu>
        </SidebarGroupContent>
      </SidebarGroup>
      <SidebarGroup>
        <SidebarGroupLabel>Github</SidebarGroupLabel>
        <SidebarGroupContent>
          <SidebarMenu>
              <SidebarMenuItem>
                <SidebarMenuButton asChild>
                    <a :href="github.url" target="_blank">
                      <component :is="github.icon" />
                      <span>Project Repository</span>
                    </a>
                </SidebarMenuButton>
              </SidebarMenuItem>
          </SidebarMenu>
        </SidebarGroupContent>
      </SidebarGroup>
    </SidebarContent>
  </Sidebar>
</template>

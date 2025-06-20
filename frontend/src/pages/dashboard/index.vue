<template>
    <div>
        <h1>Dashboard</h1>
        <NuxtLink to="/dashboard/stores">Lojas</NuxtLink>
        <NuxtLink to="/dashboard/establishments">Estabelecimentos</NuxtLink>
    </div>
</template>

<script setup>
    definePageMeta({
        middleware: 'auth',
    });
</script>

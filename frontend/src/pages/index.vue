<template>
  <div class="w-full items-center justify-center text-center">Loading...</div>
</template>

<script setup lang="ts">
    import { onMounted } from 'vue';
    import { useRouter } from 'vue-router';

    definePageMeta({
        layout: 'default'
    });

    const router = useRouter();

    onMounted(() => {
        const isAuth = localStorage.getItem('auth') === 'true'
        router.replace(isAuth ? '/establishment' : '/login')
    });
</script>

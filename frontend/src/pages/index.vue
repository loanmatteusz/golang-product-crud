<template>
  <div class="flex flex-col w-full h-[100vh] font-bold items-center justify-center text-center text-white">
    Loading...
  </div>
</template>

<script setup lang="ts">
    import { onMounted } from 'vue';
    import { useRouter } from 'vue-router';

    const router = useRouter();

    onMounted(() => {
        const isAuth = localStorage.getItem('auth') === 'true';
        router.replace(isAuth ? '/establishment' : '/login');
    });
</script>

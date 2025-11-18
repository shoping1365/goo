<template>
  <div class="flex flex-col items-center justify-center min-h-screen">
    <h1 class="text-3xl font-bold mb-4">با تشکر از خرید شما!</h1>
    <p class="mb-2">سفارش شما با موفقیت ثبت شد.</p>
    <p>
      شماره پیگیری سفارش:
      <span class="font-mono font-bold">{{ orderId }}</span>
    </p>
    <NuxtLink to="/" class="mt-6 btn btn-primary">بازگشت به صفحه اصلی</NuxtLink>
  </div>
</template>

<script lang="ts">
declare const useRoute: () => { query: Record<string, string | string[] | undefined>; params: Record<string, string> }
</script>

<script setup lang="ts">
import { computed } from 'vue';

const route = useRoute()
const orderId = computed(() => route.query.id || '---')
</script>

<style scoped>
/* تنوع استایل ساده */
</style> 
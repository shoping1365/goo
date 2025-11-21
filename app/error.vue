<script setup lang="ts">
import { onMounted } from 'vue';
import { useRouter } from 'vue-router';

// تعریف navigateTo برای Nuxt 3
declare const navigateTo: (to: string) => Promise<void>

const props = defineProps<{ error: { statusCode?: number; status?: number; [key: string]: unknown } }>()
const status = props.error?.statusCode || props.error?.status || 500
const router = useRouter()

// Redirect به صفحات خطای مجزا
if (process.server) {
  // در سرور مستقیم redirect کن
  if (status === 404) {
    await navigateTo('/404')
  } else if (status === 401) {
    await navigateTo('/401')
  } else if (status === 403) {
    await navigateTo('/403')
  } else if (status >= 500) {
    await navigateTo('/500')
  }
} else {
  // در client با onMounted
  onMounted(() => {
    if (status === 404) {
      router.push('/404')
    } else if (status === 401) {
      router.push('/401')
    } else if (status === 403) {
      router.push('/403')
    } else if (status >= 500) {
      router.push('/500')
    }
  })
}
</script>

<template>
  <div></div>
</template>
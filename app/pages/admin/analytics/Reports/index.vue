<template>
  <div class="p-6">
    <h1 class="text-2xl font-bold mb-4">Reports & Analytics</h1>
    <p>This section provides access to business reports and analytics dashboards.</p>
  </div>
</template>

<script setup lang="ts">
import { useAuth } from '~/composables/useAuth'

// تعریف definePageMeta برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void

// Reports & Analytics page
definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user, hasPermission } = useAuth()
</script> 


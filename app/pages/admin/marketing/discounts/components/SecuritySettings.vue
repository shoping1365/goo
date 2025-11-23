<template>
  <div v-if="hasAccess" class="bg-white rounded-lg shadow-sm border border-gray-200 w-full">
    <div class="p-6 border-b border-gray-200">
      <h2 class="text-lg font-semibold text-gray-900">تنظیمات امنیتی</h2>
      <p class="text-gray-600 mt-1">مدیریت محدودیت‌ها و سیاست‌های امنیتی کوپن و کمپین</p>
    </div>
    <form class="p-6 space-y-6" @submit.prevent="saveSettings">
      <div class="flex items-center gap-6">
        <label class="flex items-center cursor-pointer">
          <input v-model="settings.enableUsageLimit" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
          <span class="mr-2 text-sm text-gray-700">محدودیت تعداد استفاده فعال باشد</span>
        </label>
        <input v-if="settings.enableUsageLimit" v-model.number="settings.usageLimit" type="number" min="1" class="w-24 px-2 py-1 border border-gray-300 rounded-lg text-sm" placeholder="تعداد مجاز">
      </div>
      <div class="flex items-center gap-6">
        <label class="flex items-center cursor-pointer">
          <input v-model="settings.enableIpLimit" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
          <span class="mr-2 text-sm text-gray-700">محدودیت IP فعال باشد</span>
        </label>
        <input v-if="settings.enableIpLimit" v-model="settings.allowedIps" type="text" class="w-64 px-2 py-1 border border-gray-300 rounded-lg text-sm" placeholder="مثال: 192.168.1.1, 10.0.0.2">
      </div>
      <div class="flex items-center gap-6">
        <label class="flex items-center cursor-pointer">
          <input v-model="settings.enableTimeLimit" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
          <span class="mr-2 text-sm text-gray-700">محدودیت زمانی فعال باشد</span>
        </label>
        <input v-if="settings.enableTimeLimit" v-model.number="settings.timeLimitMinutes" type="number" min="1" class="w-24 px-2 py-1 border border-gray-300 rounded-lg text-sm" placeholder="دقیقه">
      </div>
      <button type="submit" class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors">ذخیره تنظیمات</button>
      <div v-if="success" class="text-green-600 mt-4">تنظیمات با موفقیت ذخیره شد.</div>
    </form>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useAuth } from '~/composables/useAuth'

const { user } = useAuth()
const router = useRouter()

const hasAccess = computed(() => {
  return ['admin', 'developer'].includes(user.value?.role || '')
})

watch(hasAccess, (newValue) => {
  if (!newValue) {
    router.push('/404')
  }
})

onMounted(() => {
  if (!hasAccess.value) {
    router.push('/404')
  }
})

const settings = ref({
  enableUsageLimit: false,
  usageLimit: 1,
  enableIpLimit: false,
  allowedIps: '',
  enableTimeLimit: false,
  timeLimitMinutes: 10
})
const success = ref(false)
function saveSettings() {
  success.value = true
  setTimeout(() => success.value = false, 2000)
}
</script> 

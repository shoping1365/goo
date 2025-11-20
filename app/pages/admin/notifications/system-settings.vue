<template>
  <div class="min-h-screen py-8">
    <!-- تایتل بالای صفحه -->
    <div class="bg-white w-full py-6 px-6 border-b border-gray-200 flex items-center justify-between">
      <h1 class="text-2xl font-bold text-gray-900 text-right">تنظیمات سامانه پیامک و اعلان‌ها</h1>
      <NuxtLink
to="/admin/notifications/system-settings-create"
        class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-purple-400 to-purple-600 hover:from-purple-500 hover:to-purple-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-purple-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105">
        <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
        </svg>
        ثبت درگاه
      </NuxtLink>
    </div>

    <!-- کارت‌های آماری وضعیت پیامک -->
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 mt-8">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
        <!-- کارت: پیامک ارسال شده امروز -->
        <TemplateCard 
          title="پیامک ارسال شده امروز" 
          :value="todaySMS" 
          variant="blue"
        >
          <template #icon>
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8h2a2 2 0 012 2v8a2 2 0 01-2 2H5a2 2 0 01-2-2v-8a2 2 0 012-2h2" />
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 3h-6a2 2 0 00-2 2v3h10V5a2 2 0 00-2-2z" />
            </svg>
          </template>
        </TemplateCard>

        <!-- کارت: نرخ موفقیت -->
        <TemplateCard 
          title="نرخ موفقیت" 
          :value="successRate + '%'" 
          variant="yellow"
        >
          <template #icon>
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </template>
        </TemplateCard>

        <!-- کارت: درگاه‌های فعال -->
        <TemplateCard 
          title="درگاه‌های فعال" 
          :value="activeGateways" 
          variant="green"
        >
          <template #icon>
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <circle cx="12" cy="12" r="10" stroke="currentColor" stroke-width="2" fill="none" />
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4" />
            </svg>
          </template>
        </TemplateCard>

        <!-- کارت: درگاه‌های غیرفعال -->
        <TemplateCard 
          title="درگاه‌های غیرفعال" 
          :value="inactiveGateways" 
          variant="red"
        >
          <template #icon>
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <circle cx="12" cy="12" r="10" stroke="currentColor" stroke-width="2" fill="none" />
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 9l-6 6m0-6l6 6" />
            </svg>
          </template>
        </TemplateCard>
      </div>
    </div>

    <!-- لیست داینامیک کامپوننت‌های SMS Gateway -->
    <div class="max-w-4xl mx-auto flex flex-col gap-6">
      <div v-if="pending" class="text-center py-8">
        <div class="inline-flex items-center px-4 py-2 font-semibold leading-6 text-sm shadow rounded-md text-white bg-purple-500 hover:bg-purple-400 transition ease-in-out duration-150 cursor-not-allowed">
          <svg class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
          در حال بارگذاری درگاه‌ها...
        </div>
      </div>
      
      <div v-else-if="error" class="text-center py-8">
        <div class="bg-red-50 border border-red-200 rounded-lg p-6">
          <div class="flex items-center justify-center mb-4">
            <svg class="w-8 h-8 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <h3 class="text-lg font-medium text-red-800 mb-2">خطا در بارگذاری درگاه‌ها</h3>
          <p class="text-red-600 mb-4">متأسفانه در بارگذاری درگاه‌های پیامک مشکلی پیش آمده است.</p>
          <button class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-red-600 hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500" @click="() => refresh()">
            تلاش مجدد
          </button>
        </div>
      </div>
      
      <div v-else-if="gateways?.data.length === 0" class="text-center py-8">
        <div class="bg-gray-50 border border-gray-200 rounded-lg p-6">
          <div class="flex items-center justify-center mb-4">
            <svg class="w-8 h-8 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4" />
            </svg>
          </div>
          <h3 class="text-lg font-medium text-gray-900 mb-2">هیچ درگاهی ثبت نشده</h3>
          <p class="text-gray-600 mb-4">هنوز هیچ درگاه پیامکی در سیستم ثبت نشده است.</p>
          <NuxtLink to="/admin/notifications/system-settings-create" class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-purple-600 hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-purple-500">
            ثبت اولین درگاه
          </NuxtLink>
        </div>
      </div>
      
      <DraggableGatewayList
        v-else
        :gateways="gatewaysArray"
        @gateway-deleted="handleGatewayDeleted"
        @priorities-updated="handlePrioritiesUpdated"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
// @ts-ignore: Nuxt macro
definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

import { computed, onMounted, ref } from 'vue'
import DraggableGatewayList from '~/components/admin/sms/DraggableGatewayList.vue'
import TemplateCard from '~/components/common/TemplateCard.vue'

// تعریف نوع SMSGateway (نمونه ساده)
type SMSGateway = {
  id: number
  name?: string
  type?: string
  is_active?: boolean
  priority?: number
  successCount?: number
  failureCount?: number
  status?: 'active' | 'inactive' | 'maintenance'
  api_url?: string
  sender_number?: string
  pattern_based?: boolean
  created_at?: string
  updated_at?: string
  [key: string]: any
}

// دریافت لیست درگاه‌ها از API
const gateways = ref<{ data: SMSGateway[] } | null>(null)
const pending = ref(true)
const error = ref<Error | null>(null)

// تابع دریافت درگاه‌ها
const fetchGateways = async () => {
  try {
    pending.value = true
    error.value = null
    const response = await $fetch<{ data: SMSGateway[] }>('/api/sms-gateways')
    gateways.value = response
  } catch (err) {
    error.value = err as Error
  } finally {
    pending.value = false
  }
}

// دریافت درگاه‌ها در ابتدا
onMounted(() => {
  fetchGateways()
})

// تابع refresh
const refresh = () => {
  fetchGateways()
}

// اطمینان از اینکه gateways همیشه آرایه باشد
const gatewaysArray = computed(() => {
  const data = Array.isArray(gateways.value?.data) ? gateways.value.data : []
  // تبدیل ساختار backend به frontend
  return data.map(gateway => ({
    ...gateway,
    status: (gateway.is_active ? 'active' : 'inactive') as 'active' | 'inactive' | 'maintenance',
    successCount: gateway.successCount || 0,
    failureCount: gateway.failureCount || 0,
    lastCheck: gateway.updated_at ? new Date(gateway.updated_at) : new Date()
  }))
})

const activeGateways = computed(() => gatewaysArray.value.filter(g => g.is_active).length)
const inactiveGateways = computed(() => gatewaysArray.value.filter(g => !g.is_active).length)
const todaySMS = computed(() => gatewaysArray.value.reduce((sum, g) => sum + (g.successCount || 0) + (g.failureCount || 0), 0))
const successRate = computed(() => {
  const total = gatewaysArray.value.reduce((sum, g) => sum + (g.successCount || 0) + (g.failureCount || 0), 0)
  const success = gatewaysArray.value.reduce((sum, g) => sum + (g.successCount || 0), 0)
  return total > 0 ? Math.round((success / total) * 100) : 0
})

const activeTab = ref('status')

const handleGatewayDeleted = () => {
  fetchGateways()
}

const handlePrioritiesUpdated = (priorities: Array<{id: number, priority: number}>) => {
  // به‌روزرسانی اولویت‌ها در داده‌های محلی بدون refresh
  if (gateways.value?.data) {
    const updatedGateways = gateways.value.data.map(gateway => {
      const priorityUpdate = priorities.find(p => p.id === gateway.id)
      if (priorityUpdate) {
        return { ...gateway, priority: priorityUpdate.priority }
      }
      return gateway
    })
    
    // مرتب‌سازی بر اساس اولویت جدید
    updatedGateways.sort((a, b) => (a.priority || 0) - (b.priority || 0))
    
    // به‌روزرسانی داده‌های محلی
    gateways.value.data = updatedGateways
  }
  
  // همچنین از سرور دریافت کن تا مطمئن شویم
  setTimeout(() => {
    fetchGateways()
  }, 100)
}
</script> 

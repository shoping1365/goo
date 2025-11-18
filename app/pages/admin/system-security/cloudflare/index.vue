<template>
  <div class="flex h-screen bg-gray-50" dir="rtl">
    <!-- سایدبار کلادفلر -->
    <div class="w-64 bg-white shadow-lg border-l border-gray-200 overflow-y-auto">
      <!-- لوگو و عنوان -->
      <div class="p-6 border-b border-gray-200">
        <div class="flex items-center space-x-3 space-x-reverse">
          <div class="w-8 h-8 bg-orange-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="currentColor" viewBox="0 0 24 24">
              <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/>
            </svg>
          </div>
          <h1 class="text-lg font-bold text-gray-900">کلادفلر</h1>
        </div>
      </div>

      <!-- منوی اصلی -->
      <nav class="p-6">
        <div class="space-y-1">
          <!-- Overview -->
          <NuxtLink 
            to="/admin/system-security/cloudflare" 
            class="flex items-center px-3 py-2 text-sm font-medium rounded-md transition-colors"
            :class="route.path === '/admin/system-security/cloudflare' ? 'bg-blue-50 text-blue-700' : 'text-gray-700 hover:bg-gray-50'"
          >
            <svg class="w-5 h-5 ml-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"/>
            </svg>
            نمای کلی
          </NuxtLink>

          <!-- Security -->
          <div class="mt-6">
            <h3 class="px-3 text-xs font-semibold text-gray-500 uppercase tracking-wider">امنیت</h3>
            <div class="mt-2 space-y-1">
              <NuxtLink 
                to="/admin/system-security/cloudflare/security/waf" 
                class="flex items-center px-3 py-2 text-sm font-medium rounded-md transition-colors"
                :class="route.path.includes('/waf') ? 'bg-blue-50 text-blue-700' : 'text-gray-700 hover:bg-gray-50'"
              >
                <svg class="w-5 h-5 ml-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"/>
                </svg>
                WAF
              </NuxtLink>
              <NuxtLink 
                to="/admin/system-security/cloudflare/security/ddos" 
                class="flex items-center px-3 py-2 text-sm font-medium rounded-md transition-colors"
                :class="route.path.includes('/ddos') ? 'bg-blue-50 text-blue-700' : 'text-gray-700 hover:bg-gray-50'"
              >
                <svg class="w-5 h-5 ml-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"/>
                </svg>
                DDoS Protection
              </NuxtLink>
              <NuxtLink 
                to="/admin/system-security/cloudflare/security/bot-management" 
                class="flex items-center px-3 py-2 text-sm font-medium rounded-md transition-colors"
                :class="route.path.includes('/bot-management') ? 'bg-blue-50 text-blue-700' : 'text-gray-700 hover:bg-gray-50'"
              >
                <svg class="w-5 h-5 ml-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.75 17L9 20l-1 1h8l-1-1-.75-3M3 13h18M5 17h14a2 2 0 002-2V5a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"/>
                </svg>
                مدیریت ربات‌ها
              </NuxtLink>
              <NuxtLink 
                to="/admin/system-security/cloudflare/security/rate-limiting" 
                class="flex items-center px-3 py-2 text-sm font-medium rounded-md transition-colors"
                :class="route.path.includes('/rate-limiting') ? 'bg-blue-50 text-blue-700' : 'text-gray-700 hover:bg-gray-50'"
              >
                <svg class="w-5 h-5 ml-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"/>
                </svg>
                محدودیت نرخ
              </NuxtLink>
            </div>
          </div>

          <!-- Performance -->
          <div class="mt-6">
            <h3 class="px-3 text-xs font-semibold text-gray-500 uppercase tracking-wider">عملکرد</h3>
            <div class="mt-2 space-y-1">
              <NuxtLink 
                to="/admin/system-security/cloudflare/performance/caching" 
                class="flex items-center px-3 py-2 text-sm font-medium rounded-md transition-colors"
                :class="route.path.includes('/caching') ? 'bg-blue-50 text-blue-700' : 'text-gray-700 hover:bg-gray-50'"
              >
                <svg class="w-5 h-5 ml-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"/>
                </svg>
                کش
              </NuxtLink>
              <NuxtLink 
                to="/admin/system-security/cloudflare/performance/page-rules" 
                class="flex items-center px-3 py-2 text-sm font-medium rounded-md transition-colors"
                :class="route.path.includes('/page-rules') ? 'bg-blue-50 text-blue-700' : 'text-gray-700 hover:bg-gray-50'"
              >
                <svg class="w-5 h-5 ml-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"/>
                </svg>
                قوانین صفحه
              </NuxtLink>
              <NuxtLink 
                to="/admin/system-security/cloudflare/performance/workers" 
                class="flex items-center px-3 py-2 text-sm font-medium rounded-md transition-colors"
                :class="route.path.includes('/workers') ? 'bg-blue-50 text-blue-700' : 'text-gray-700 hover:bg-gray-50'"
              >
                <svg class="w-5 h-5 ml-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"/>
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/>
                </svg>
                Workers
              </NuxtLink>
            </div>
          </div>

          <!-- DNS -->
          <div class="mt-6">
            <h3 class="px-3 text-xs font-semibold text-gray-500 uppercase tracking-wider">DNS</h3>
            <div class="mt-2 space-y-1">
              <NuxtLink 
                to="/admin/system-security/cloudflare/dns/records" 
                class="flex items-center px-3 py-2 text-sm font-medium rounded-md transition-colors"
                :class="route.path.includes('/records') ? 'bg-blue-50 text-blue-700' : 'text-gray-700 hover:bg-gray-50'"
              >
                <svg class="w-5 h-5 ml-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"/>
                </svg>
                رکوردهای DNS
              </NuxtLink>
              <NuxtLink 
                to="/admin/system-security/cloudflare/dns/analytics" 
                class="flex items-center px-3 py-2 text-sm font-medium rounded-md transition-colors"
                :class="route.path.includes('/analytics') ? 'bg-blue-50 text-blue-700' : 'text-gray-700 hover:bg-gray-50'"
              >
                <svg class="w-5 h-5 ml-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"/>
                </svg>
                تحلیل DNS
              </NuxtLink>
            </div>
          </div>

          <!-- Analytics -->
          <div class="mt-6">
            <h3 class="px-3 text-xs font-semibold text-gray-500 uppercase tracking-wider">تحلیل</h3>
            <div class="mt-2 space-y-1">
              <NuxtLink 
                to="/admin/system-security/cloudflare/analytics/traffic" 
                class="flex items-center px-3 py-2 text-sm font-medium rounded-md transition-colors"
                :class="route.path.includes('/traffic') ? 'bg-blue-50 text-blue-700' : 'text-gray-700 hover:bg-gray-50'"
              >
                <svg class="w-5 h-5 ml-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"/>
                </svg>
                ترافیک
              </NuxtLink>
              <NuxtLink 
                to="/admin/system-security/cloudflare/analytics/security" 
                class="flex items-center px-3 py-2 text-sm font-medium rounded-md transition-colors"
                :class="route.path.includes('/security') ? 'bg-blue-50 text-blue-700' : 'text-gray-700 hover:bg-gray-50'"
              >
                <svg class="w-5 h-5 ml-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"/>
                </svg>
                امنیت
              </NuxtLink>
              <NuxtLink 
                to="/admin/system-security/cloudflare/analytics/performance" 
                class="flex items-center px-3 py-2 text-sm font-medium rounded-md transition-colors"
                :class="route.path.includes('/performance') ? 'bg-blue-50 text-blue-700' : 'text-gray-700 hover:bg-gray-50'"
              >
                <svg class="w-5 h-5 ml-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"/>
                </svg>
                عملکرد
              </NuxtLink>
            </div>
          </div>

          <!-- Settings -->
          <div class="mt-6">
            <h3 class="px-3 text-xs font-semibold text-gray-500 uppercase tracking-wider">تنظیمات</h3>
            <div class="mt-2 space-y-1">
              <NuxtLink 
                to="/admin/system-security/cloudflare/settings/general" 
                class="flex items-center px-3 py-2 text-sm font-medium rounded-md transition-colors"
                :class="route.path.includes('/general') ? 'bg-blue-50 text-blue-700' : 'text-gray-700 hover:bg-gray-50'"
              >
                <svg class="w-5 h-5 ml-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"/>
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/>
                </svg>
                عمومی
              </NuxtLink>
              <NuxtLink 
                to="/admin/system-security/cloudflare/settings/api" 
                class="flex items-center px-3 py-2 text-sm font-medium rounded-md transition-colors"
                :class="route.path.includes('/api') ? 'bg-blue-50 text-blue-700' : 'text-gray-700 hover:bg-gray-50'"
              >
                <svg class="w-5 h-5 ml-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 9l3 3-3 3m5 0h3M5 20h14a2 2 0 002-2V6a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"/>
                </svg>
                API
              </NuxtLink>
            </div>
          </div>
        </div>
      </nav>
    </div>

    <!-- محتوای اصلی -->
    <div class="flex-1 overflow-auto">
      <div class="p-6">
        <!-- هدر صفحه -->
        <div class="mb-6">
          <h1 class="text-2xl font-bold text-gray-900">نمای کلی کلادفلر</h1>
          <p class="text-gray-600">مدیریت کامل سرویس‌های کلادفلر</p>
        </div>

        <!-- کارت‌های آمار -->
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
          <div class="bg-white rounded-lg shadow p-6 border border-gray-200">
            <div class="flex items-center">
              <div class="p-2 bg-blue-100 rounded-lg">
                <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"/>
                </svg>
              </div>
              <div class="mr-4">
                <p class="text-sm font-medium text-gray-600">درخواست‌های امروز</p>
                <p class="text-2xl font-bold text-gray-900">{{ requestsTodayFormatted }}</p>
              </div>
            </div>
          </div>

          <div class="bg-white rounded-lg shadow p-6 border border-gray-200">
            <div class="flex items-center">
              <div class="p-2 bg-green-100 rounded-lg">
                <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"/>
                </svg>
              </div>
              <div class="mr-4">
                <p class="text-sm font-medium text-gray-600">نرخ کش</p>
                <p class="text-2xl font-bold text-gray-900">{{ cacheHitRateFormatted }}</p>
              </div>
            </div>
          </div>

          <div class="bg-white rounded-lg shadow p-6 border border-gray-200">
            <div class="flex items-center">
              <div class="p-2 bg-red-100 rounded-lg">
                <svg class="w-6 h-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"/>
                </svg>
              </div>
              <div class="mr-4">
                <p class="text-sm font-medium text-gray-600">تهدیدات مسدود شده</p>
                <p class="text-2xl font-bold text-gray-900">{{ threatsBlockedFormatted }}</p>
              </div>
            </div>
          </div>

          <div class="bg-white rounded-lg shadow p-6 border border-gray-200">
            <div class="flex items-center">
              <div class="p-2 bg-purple-100 rounded-lg">
                <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"/>
                </svg>
              </div>
              <div class="mr-4">
                <p class="text-sm font-medium text-gray-600">زمان پاسخ</p>
                <p class="text-2xl font-bold text-gray-900">{{ responseTimeMsFormatted }}</p>
              </div>
            </div>
          </div>
        </div>

        <!-- نمودارها و اطلاعات بیشتر -->
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <div class="bg-white rounded-lg shadow p-6 border border-gray-200">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">ترافیک اخیر</h3>
            <div class="h-64 bg-gray-50 rounded flex items-center justify-center">
              <p class="text-gray-500">نمودار ترافیک</p>
            </div>
          </div>

          <div class="bg-white rounded-lg shadow p-6 border border-gray-200">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">تهدیدات امنیتی</h3>
            <div class="h-64 bg-gray-50 rounded flex items-center justify-center">
              <p class="text-gray-500">نمودار تهدیدات</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string }) => void
declare const useRoute: () => { params: Record<string, string>; query: Record<string, string>; path: string }
declare const useFetch: <T = unknown>(url: string, options?: { method?: string; key?: string; body?: unknown }) => Promise<{ data: { value: T }; error: { value: Error | null } }>

type CloudflareOverview = {
  requests: number
  cacheHitRate: number
  threats: number
  avgResponseMs: number
}

type CloudflareOverviewResponse = {
  success?: boolean
  data?: CloudflareOverview
}
</script>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';

definePageMeta({ layout: 'admin-main' })

const route = useRoute()

// وضعیت نمای کلی از API کلادفلر
const overview = ref<CloudflareOverview | null>(null)
const isOverviewLoading = ref(false)
const overviewError = ref('')

// قالب‌بندی اعداد برای نمایش
const requestsTodayFormatted = computed(() => overview.value ? new Intl.NumberFormat('fa-IR').format(overview.value.requests) : '—')
const cacheHitRateFormatted = computed(() => overview.value ? `${overview.value.cacheHitRate}%` : '—')
const threatsBlockedFormatted = computed(() => overview.value ? new Intl.NumberFormat('fa-IR').format(overview.value.threats) : '—')
const responseTimeMsFormatted = computed(() => overview.value ? `${overview.value.avgResponseMs}ms` : '—')

// فراخوانی داده‌ها هنگام ورود به صفحه
onMounted(async () => {
  try {
    isOverviewLoading.value = true
    const { data, error } = await useFetch<CloudflareOverviewResponse>('/api/admin/cloudflare/overview', {
      method: 'GET',
      key: 'cloudflare-overview',
    })
    if (error.value) {
      overviewError.value = error.value instanceof Error ? error.value.message : 'خطا در دریافت نمای کلی'
    } else if (data.value && data.value.success && data.value.data) {
      overview.value = data.value.data
    }
  } catch (e) {
    const error = e instanceof Error ? e : new Error('خطای نامشخص')
    overviewError.value = error.message
  } finally {
    isOverviewLoading.value = false
  }
})
</script> 

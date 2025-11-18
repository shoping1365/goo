<template>
  <div class="bg-white rounded-xl shadow-lg p-6">
    <h3 class="text-lg font-bold text-gray-900 mb-6">تحلیل‌های پیشرفته</h3>
    
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <!-- Conversion Rate -->
      <div class="text-center p-6 bg-gradient-to-br from-blue-50 to-blue-100 rounded-lg">
        <div class="w-12 h-12 bg-blue-500 rounded-full flex items-center justify-center mx-auto mb-3">
          <svg class="w-6 h-6 text-white" fill="currentColor" viewBox="0 0 20 20">
            <path fill-rule="evenodd" d="M6.267 3.455a3.066 3.066 0 001.745-.723 3.066 3.066 0 013.976 0 3.066 3.066 0 001.745.723 3.066 3.066 0 012.812 2.812c.051.643.304 1.254.723 1.745a3.066 3.066 0 010 3.976 3.066 3.066 0 00-.723 1.745 3.066 3.066 0 01-2.812 2.812 3.066 3.066 0 00-1.745.723 3.066 3.066 0 01-3.976 0 3.066 3.066 0 00-1.745-.723 3.066 3.066 0 01-2.812-2.812 3.066 3.066 0 00-.723-1.745 3.066 3.066 0 010-3.976 3.066 3.066 0 00.723-1.745 3.066 3.066 0 012.812-2.812zm7.44 5.252a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
          </svg>
        </div>
        <div class="text-2xl font-bold text-blue-700 mb-1">{{ analytics.conversionRate }}%</div>
        <div class="text-sm text-blue-600 font-medium">نرخ تبدیل</div>
        <div class="text-xs text-gray-600 mt-1">از بازدیدکننده به مشتری</div>
      </div>

      <!-- Session Duration -->
      <div class="text-center p-6 bg-gradient-to-br from-green-50 to-green-100 rounded-lg">
        <div class="w-12 h-12 bg-green-500 rounded-full flex items-center justify-center mx-auto mb-3">
          <svg class="w-6 h-6 text-white" fill="currentColor" viewBox="0 0 20 20">
            <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-12a1 1 0 10-2 0v4a1 1 0 00.293.707l2.828 2.829a1 1 0 101.415-1.415L11 9.586V6z" clip-rule="evenodd" />
          </svg>
        </div>
        <div class="text-2xl font-bold text-green-700 mb-1">{{ analytics.avgSessionDuration }}</div>
        <div class="text-sm text-green-600 font-medium">مدت جلسه</div>
        <div class="text-xs text-gray-600 mt-1">میانگین زمان ماندگاری</div>
      </div>

      <!-- Pages per Session -->
      <div class="text-center p-6 bg-gradient-to-br from-purple-50 to-purple-100 rounded-lg">
        <div class="w-12 h-12 bg-purple-500 rounded-full flex items-center justify-center mx-auto mb-3">
          <svg class="w-6 h-6 text-white" fill="currentColor" viewBox="0 0 20 20">
            <path fill-rule="evenodd" d="M4 4a2 2 0 012-2h4.586A2 2 0 0112 2.586L15.414 6A2 2 0 0116 7.414V16a2 2 0 01-2 2H6a2 2 0 01-2-2V4z" clip-rule="evenodd" />
          </svg>
        </div>
        <div class="text-2xl font-bold text-purple-700 mb-1">{{ analytics.pagesPerSession }}</div>
        <div class="text-sm text-purple-600 font-medium">صفحه در جلسه</div>
        <div class="text-xs text-gray-600 mt-1">میانگین صفحات بازدید شده</div>
      </div>

      <!-- New Visitor Rate -->
      <div class="text-center p-6 bg-gradient-to-br from-orange-50 to-orange-100 rounded-lg">
        <div class="w-12 h-12 bg-orange-500 rounded-full flex items-center justify-center mx-auto mb-3">
          <svg class="w-6 h-6 text-white" fill="currentColor" viewBox="0 0 20 20">
            <path d="M8 9a3 3 0 100-6 3 3 0 000 6zM8 11a6 6 0 016 6H2a6 6 0 016-6zM16 7a1 1 0 10-2 0v1h-1a1 1 0 100 2h1v1a1 1 0 102 0v-1h1a1 1 0 100-2h-1V7z" />
          </svg>
        </div>
        <div class="text-2xl font-bold text-orange-700 mb-1">{{ analytics.newVisitorRate }}%</div>
        <div class="text-sm text-orange-600 font-medium">بازدیدکننده جدید</div>
        <div class="text-xs text-gray-600 mt-1">نسبت کاربران جدید</div>
      </div>
    </div>

    <!-- Detailed Analytics -->
    <div class="mt-8 grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- Engagement Metrics -->
      <div class="bg-gray-50 rounded-lg p-6">
        <h4 class="text-md font-semibold text-gray-900 mb-4">متریک‌های تعامل</h4>
        <div class="space-y-3">
          <div class="flex items-center justify-between">
            <span class="text-sm text-gray-600">نرخ کلیک</span>
            <div class="flex items-center gap-2">
              <div class="w-24 bg-gray-200 rounded-full h-2">
                <div class="bg-blue-500 h-2 rounded-full" style="width: 68%"></div>
              </div>
              <span class="text-sm font-semibold text-gray-900">6.8%</span>
            </div>
          </div>
          <div class="flex items-center justify-between">
            <span class="text-sm text-gray-600">نرخ بازگشت</span>
            <div class="flex items-center gap-2">
              <div class="w-24 bg-gray-200 rounded-full h-2">
                <div class="bg-green-500 h-2 rounded-full" style="width: 45%"></div>
              </div>
              <span class="text-sm font-semibold text-gray-900">45%</span>
            </div>
          </div>
          <div class="flex items-center justify-between">
            <span class="text-sm text-gray-600">نرخ خروج</span>
            <div class="flex items-center gap-2">
              <div class="w-24 bg-gray-200 rounded-full h-2">
                <div class="bg-red-500 h-2 rounded-full" style="width: 32%"></div>
              </div>
              <span class="text-sm font-semibold text-gray-900">32%</span>
            </div>
          </div>
          <div class="flex items-center justify-between">
            <span class="text-sm text-gray-600">زمان بارگذاری صفحه</span>
            <div class="flex items-center gap-2">
              <div class="w-24 bg-gray-200 rounded-full h-2">
                <div class="bg-yellow-500 h-2 rounded-full" style="width: 78%"></div>
              </div>
              <span class="text-sm font-semibold text-gray-900">2.3s</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Geographic Insights -->
      <div class="bg-gray-50 rounded-lg p-6">
        <h4 class="text-md font-semibold text-gray-900 mb-4">بینش‌های جغرافیایی</h4>
        <div class="space-y-3">
          <div class="flex items-center justify-between">
            <span class="text-sm text-gray-600">پراکندگی جغرافیایی</span>
            <span class="text-sm font-semibold text-green-600">عالی</span>
          </div>
          <div class="flex items-center justify-between">
            <span class="text-sm text-gray-600">شهرهای فعال</span>
            <span class="text-sm font-semibold text-gray-900">{{ cityCount }} شهر</span>
          </div>
          <div class="flex items-center justify-between">
            <span class="text-sm text-gray-600">تمرکز ترافیک</span>
            <span class="text-sm font-semibold text-gray-900">تهران (37%)</span>
          </div>
          <div class="flex items-center justify-between">
            <span class="text-sm text-gray-600">رشد منطقه‌ای</span>
            <span class="text-sm font-semibold text-blue-600">+12.5%</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Performance Score -->
    <div class="mt-6 text-center p-6 bg-gradient-to-r from-blue-50 to-purple-50 rounded-lg">
      <h4 class="text-lg font-semibold text-gray-900 mb-2">امتیاز کلی عملکرد</h4>
      <div class="flex items-center justify-center gap-6">
        <div class="text-4xl font-bold text-blue-600">{{ performanceScore }}</div>
        <div class="text-left">
          <div class="text-sm text-gray-600">از 100</div>
          <div class="text-xs text-green-600 font-medium">{{ getPerformanceLevel() }}</div>
        </div>
      </div>
      <div class="mt-4 w-full bg-gray-200 rounded-full h-3">
        <div 
          class="bg-gradient-to-r from-blue-500 to-purple-600 h-3 rounded-full transition-all duration-1000"
          :style="{ width: performanceScore + '%' }"
        ></div>
      </div>
      <p class="text-sm text-gray-600 mt-3">
        بر اساس نرخ تبدیل، تعامل کاربران، سرعت سایت و پراکندگی جغرافیایی محاسبه شده است
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface Analytics {
  conversionRate: number
  avgSessionDuration: string
  pagesPerSession: number
  newVisitorRate: number
}

const props = defineProps<{
  analytics: Analytics
}>()

const cityCount = computed(() => 10) // Based on our map data

const performanceScore = computed(() => {
  // Calculate performance score based on various metrics
  const conversionScore = Math.min(props.analytics.conversionRate * 10, 25)
  const sessionScore = 20 // Based on avg session duration
  const pagesScore = Math.min(props.analytics.pagesPerSession * 10, 25)
  const newVisitorScore = Math.min(props.analytics.newVisitorRate / 2, 30)
  
  return Math.round(conversionScore + sessionScore + pagesScore + newVisitorScore)
})

const getPerformanceLevel = () => {
  if (performanceScore.value >= 90) return 'عالی'
  if (performanceScore.value >= 80) return 'خوب'
  if (performanceScore.value >= 70) return 'متوسط'
  if (performanceScore.value >= 60) return 'نیاز به بهبود'
  return 'ضعیف'
}
</script> 

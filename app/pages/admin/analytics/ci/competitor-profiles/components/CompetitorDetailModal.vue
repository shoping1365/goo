<template>
  <div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-6">
    <div class="bg-white rounded-xl shadow-xl w-full max-w-6xl max-h-[90vh] overflow-hidden">
      <!-- Header -->
      <div class="flex items-center justify-between p-6 border-b border-gray-200">
        <div class="flex items-center space-x-3 space-x-reverse">
          <div class="w-12 h-12 bg-gradient-to-br from-blue-500 to-blue-600 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"></path>
            </svg>
          </div>
          <div>
            <h2 class="text-2xl font-bold text-gray-900">{{ competitor.name }}</h2>
            <div class="flex items-center space-x-2 space-x-reverse mt-1">
              <span class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium" :class="getThreatLevelBadgeClass(competitor.threatLevel)">
                {{ getThreatLevelText(competitor.threatLevel) }}
              </span>
              <span class="text-sm text-gray-500">{{ competitor.companySize }}</span>
            </div>
          </div>
        </div>
        <div class="flex items-center space-x-2 space-x-reverse">
          <button 
            class="inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-lg text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
            @click="$emit('edit', competitor)"
          >
            <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
            </svg>
            ویرایش
          </button>
          <button 
            class="p-2 text-gray-400 hover:text-gray-600 hover:bg-gray-100 rounded-lg transition-colors"
            @click="$emit('close')"
          >
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>
      </div>

      <!-- Content -->
      <div class="overflow-y-auto max-h-[calc(90vh-120px)]">
        <div class="p-6">
          <!-- Basic Information -->
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-8">
            <div class="bg-gray-50 rounded-xl p-6">
              <h3 class="text-lg font-semibold text-gray-900 mb-4">اطلاعات پایه</h3>
              <div class="space-y-4">
                <div class="flex items-center justify-between">
                  <span class="text-sm font-medium text-gray-600">وب‌سایت:</span>
                  <a :href="competitor.website" target="_blank" class="text-blue-600 hover:text-blue-700 text-sm">{{ competitor.website }}</a>
                </div>
                <div class="flex items-center justify-between">
                  <span class="text-sm font-medium text-gray-600">سال تأسیس:</span>
                  <span class="text-sm text-gray-900">{{ competitor.foundedYear }}</span>
                </div>
                <div class="flex items-center justify-between">
                  <span class="text-sm font-medium text-gray-600">اندازه شرکت:</span>
                  <span class="text-sm text-gray-900">{{ competitor.companySize }}</span>
                </div>
                <div class="flex items-center justify-between">
                  <span class="text-sm font-medium text-gray-600">قیمت‌گذاری:</span>
                  <span class="text-sm text-gray-900">{{ competitor.pricing }}</span>
                </div>
              </div>
            </div>

            <div class="bg-gray-50 rounded-xl p-6">
              <h3 class="text-lg font-semibold text-gray-900 mb-4">بازارهای هدف</h3>
              <div class="flex flex-wrap gap-2">
                <span 
                  v-for="market in competitor.targetMarkets" 
                  :key="market"
                  class="inline-flex items-center px-3 py-1 rounded-full text-sm font-medium bg-blue-100 text-blue-800"
                >
                  {{ market }}
                </span>
              </div>
            </div>
          </div>

          <!-- Products and Services -->
          <div class="bg-white rounded-xl border border-gray-200 p-6 mb-8">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">محصولات و خدمات</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <h4 class="text-md font-medium text-gray-700 mb-3">محصولات کلیدی</h4>
                <div class="space-y-2">
                  <div 
                    v-for="product in competitor.keyProducts" 
                    :key="product"
                    class="flex items-center p-3 bg-green-50 rounded-lg"
                  >
                    <svg class="w-4 h-4 text-green-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                    </svg>
                    <span class="text-sm text-gray-900">{{ product }}</span>
                  </div>
                </div>
              </div>
              <div>
                <h4 class="text-md font-medium text-gray-700 mb-3">قیمت‌گذاری</h4>
                <div class="p-3 bg-blue-50 rounded-lg">
                  <span class="text-sm text-gray-900">{{ competitor.pricing }}</span>
                </div>
              </div>
            </div>
          </div>

          <!-- Strengths and Weaknesses -->
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-8">
            <div class="bg-white rounded-xl border border-gray-200 p-6">
              <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
                <svg class="w-5 h-5 text-green-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
                نقاط قوت
              </h3>
              <div class="space-y-2">
                <div 
                  v-for="strength in competitor.strengths" 
                  :key="strength"
                  class="flex items-start p-2 bg-green-50 rounded-lg"
                >
                  <span class="w-2 h-2 bg-green-500 rounded-full mt-2 ml-2 flex-shrink-0"></span>
                  <span class="text-sm text-gray-900">{{ strength }}</span>
                </div>
              </div>
            </div>

            <div class="bg-white rounded-xl border border-gray-200 p-6">
              <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
                <svg class="w-5 h-5 text-red-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
                نقاط ضعف
              </h3>
              <div class="space-y-2">
                <div 
                  v-for="weakness in competitor.weaknesses" 
                  :key="weakness"
                  class="flex items-start p-2 bg-red-50 rounded-lg"
                >
                  <span class="w-2 h-2 bg-red-500 rounded-full mt-2 ml-2 flex-shrink-0"></span>
                  <span class="text-sm text-gray-900">{{ weakness }}</span>
                </div>
              </div>
            </div>
          </div>

          <!-- Recent Changes -->
          <div class="bg-white rounded-xl border border-gray-200 p-6 mb-8">
            <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
              <svg class="w-5 h-5 text-blue-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
              </svg>
              تغییرات اخیر
            </h3>
            <div class="space-y-3">
              <div 
                v-for="change in competitor.recentChanges" 
                :key="change"
                class="flex items-start p-3 bg-blue-50 rounded-lg"
              >
                <svg class="w-4 h-4 text-blue-600 ml-2 mt-0.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
                </svg>
                <span class="text-sm text-gray-900">{{ change }}</span>
              </div>
            </div>
          </div>

          <!-- Financial Information (Placeholder) -->
          <div class="bg-white rounded-xl border border-gray-200 p-6 mb-8">
            <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
              <svg class="w-5 h-5 text-purple-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
              </svg>
              اطلاعات مالی
            </h3>
            <div class="text-center py-8">
              <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
              </svg>
              <h3 class="mt-2 text-sm font-medium text-gray-900">اطلاعات مالی در دسترس نیست</h3>
              <p class="mt-1 text-sm text-gray-500">این اطلاعات معمولاً از منابع عمومی قابل دسترسی نیست.</p>
            </div>
          </div>

          <!-- Organizational Structure (Placeholder) -->
          <div class="bg-white rounded-xl border border-gray-200 p-6 mb-8">
            <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
              <svg class="w-5 h-5 text-indigo-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"></path>
              </svg>
              ساختار سازمانی
            </h3>
            <div class="text-center py-8">
              <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
              </svg>
              <h3 class="mt-2 text-sm font-medium text-gray-900">اطلاعات ساختار سازمانی در دسترس نیست</h3>
              <p class="mt-1 text-sm text-gray-500">این اطلاعات نیاز به تحقیق و جمع‌آوری از منابع مختلف دارد.</p>
            </div>
          </div>

          <!-- Sales Channels (Placeholder) -->
          <div class="bg-white rounded-xl border border-gray-200 p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
              <svg class="w-5 h-5 text-orange-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z"></path>
              </svg>
              کانال‌های فروش
            </h3>
            <div class="text-center py-8">
              <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
              </svg>
              <h3 class="mt-2 text-sm font-medium text-gray-900">اطلاعات کانال‌های فروش در دسترس نیست</h3>
              <p class="mt-1 text-sm text-gray-500">این اطلاعات نیاز به تحقیق و جمع‌آوری از منابع مختلف دارد.</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { type CompetitorProfile } from '../../composables/useCompetitiveIntelligence'

// Props
interface Props {
  competitor: CompetitorProfile
}

defineProps<Props>()

// Emits
defineEmits<{
  close: []
  edit: [competitor: CompetitorProfile]
}>()

// Methods
const getThreatLevelBadgeClass = (level: string) => {
  switch (level) {
    case 'low': return 'bg-green-100 text-green-800'
    case 'medium': return 'bg-yellow-100 text-yellow-800'
    case 'high': return 'bg-red-100 text-red-800'
    default: return 'bg-gray-100 text-gray-800'
  }
}

const getThreatLevelText = (level: string) => {
  switch (level) {
    case 'low': return 'کم'
    case 'medium': return 'متوسط'
    case 'high': return 'زیاد'
    default: return 'نامشخص'
  }
}
</script>

<style scoped>
/* Custom styles */
</style>

<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200">
    <!-- هدر بخش -->
    <div class="px-6 py-4 border-b border-gray-200">
      <div class="flex items-center justify-between">
        <div>
          <h3 class="text-lg font-semibold text-gray-900">پشتیبانی و راهنما</h3>
          <p class="text-sm text-gray-600 mt-1">دسترسی سریع به راهنما، آموزش و پشتیبانی اتصال حسابداری</p>
        </div>
      </div>
    </div>

    <!-- تب‌های بخش‌های مختلف -->
    <div class="border-b border-gray-200">
      <nav class="-mb-px flex space-x-8 space-x-reverse px-6">
        <button
          v-for="tab in tabs"
          :key="tab.id"
          @click="activeTab = tab.id"
          :class="[
            activeTab === tab.id
              ? 'border-blue-500 text-blue-600'
              : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300',
            'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm'
          ]"
        >
          {{ tab.name }}
        </button>
      </nav>
    </div>

    <!-- محتوای تب‌ها -->
    <div class="p-6">
      <div v-if="activeTab === 'guide'">
        <SupportGuide />
      </div>
      <div v-if="activeTab === 'videos'">
        <SupportVideos />
      </div>
      <div v-if="activeTab === 'support'">
        <SupportTechnical />
      </div>
      <div v-if="activeTab === 'faq'">
        <SupportFAQ />
      </div>
      <div v-if="activeTab === 'api'">
        <SupportAPIDocs />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
// @ts-ignore
import SupportGuide from './support/SupportGuide.vue'
// @ts-ignore
import SupportVideos from './support/SupportVideos.vue'
// @ts-ignore
import SupportTechnical from './support/SupportTechnical.vue'
// @ts-ignore
import SupportFAQ from './support/SupportFAQ.vue'
// @ts-ignore
import SupportAPIDocs from './support/SupportAPIDocs.vue'

const tabs = ref([
  { id: 'guide', name: 'راهنمای اتصال' },
  { id: 'videos', name: 'ویدیوهای آموزشی' },
  { id: 'support', name: 'پشتیبانی فنی' },
  { id: 'faq', name: 'سوالات متداول' },
  { id: 'api', name: 'مستندات API' }
])
const activeTab = ref('guide')
</script>

<!--
  کامپوننت اصلی پشتیبانی و راهنما
  شامل:
  1. راهنمای اتصال
  2. ویدیوهای آموزشی
  3. پشتیبانی فنی
  4. سوالات متداول
  5. مستندات API
  طراحی مدرن و کاملاً ریسپانسیو
  توضیحات کامل به فارسی
--> 

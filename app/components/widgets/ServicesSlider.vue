<template>
  <div class="services-slider w-full">
    <div class="container mx-auto px-4 py-6">
      <!-- کارت‌های خدمات -->
      <div class="grid grid-cols-3 gap-6">
        <!-- کارت خرید حضوری -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 text-center">
          <div class="w-12 h-12 bg-blue-100 rounded-full flex items-center justify-center mx-auto mb-3">
            <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4m0 0L7 13m0 0l-2.5 5M7 13l2.5 5m6-5v6a2 2 0 01-2 2H9a2 2 0 01-2-2v-6m8 0V9a2 2 0 00-2-2H9a2 2 0 00-2 2v4.01"></path>
            </svg>
          </div>
          <h3 class="text-sm font-medium text-gray-900 mb-1">خرید حضوری</h3>
          <p class="text-xs text-gray-600">از فروشگاه</p>
        </div>

        <!-- کارت تحویل سریع -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 text-center">
          <div class="w-12 h-12 bg-green-100 rounded-full flex items-center justify-center mx-auto mb-3">
            <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
            </svg>
          </div>
          <h3 class="text-sm font-medium text-gray-900 mb-1">تحویل سریع</h3>
          <p class="text-xs text-gray-600">در سرتاسر ایران</p>
        </div>

        <!-- کارت تضمین اصالت -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 text-center">
          <div class="w-12 h-12 bg-purple-100 rounded-full flex items-center justify-center mx-auto mb-3">
            <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
          <h3 class="text-sm font-medium text-gray-900 mb-1">تضمین اصالت</h3>
          <p class="text-xs text-gray-600">تمام کالاها</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Widget } from '~/types/widget'

interface Props {
  widget: Widget
}

defineProps<Props>()
</script> 
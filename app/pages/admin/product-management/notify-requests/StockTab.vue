<script setup lang="ts">
import TemplateButton from '@/components/common/TemplateButton.vue';
import { useNotifyRequests } from '~/composables/useNotifyRequests';

interface Props {
  activeNavItem: string
}

const props = defineProps<Props>()

const { 
  stockSearchQuery,
  selectedStockItems,
  stockRequests,
  stockNotificationStates,
  applyFilters,
  toggleAllStock,
  bulkDeleteStock,
  bulkNotifyStock,
  sendStockNotification,
  clearStockError,
  getStatusClass
} = useNotifyRequests()
</script>

<template>

  <!-- تب موجودی -->
  <div v-if="props.activeNavItem === 'stock'" class="w-full">
    <!-- فیلترها و عملیات -->
    <div class="bg-white rounded-xl shadow-md border border-gray-100 p-6 mb-6">
      <div class="flex flex-row-reverse flex-wrap items-center justify-between gap-6 w-full">
        <!-- دکمه‌ها و سلکت‌ها -->
        <div class="flex flex-row-reverse items-center gap-3">
          <!-- دکمه‌های عملیات شرطی -->
          <div v-if="selectedStockItems.length > 0" class="flex items-center gap-2">
            <TemplateButton
              @click="bulkDeleteStock"
              bgGradient="bg-gradient-to-r from-red-400 to-red-600"
              textColor="text-white"
              hoverClass="hover:from-red-500 hover:to-red-700 hover:shadow-lg hover:scale-105"
              focusClass="focus:ring-2 focus:ring-offset-2 focus:ring-red-500"
              size="small"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
              </svg>
              <span>حذف ({{ selectedStockItems.length }})</span>
            </TemplateButton>
            <TemplateButton
              @click="bulkNotifyStock"
              bgGradient="bg-gradient-to-r from-green-400 to-green-600"
              textColor="text-white"
              hoverClass="hover:from-green-500 hover:to-green-700 hover:shadow-lg hover:scale-105"
              focusClass="focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
              size="small"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-5 5-5-5h5z"></path>
              </svg>
              <span>ارسال اطلاع رسانی ({{ selectedStockItems.length }})</span>
            </TemplateButton>
          </div>


        </div>

        <!-- جستجو -->
        <div class="relative">
          <input
              type="text"
              placeholder="جستجو در ایمیل و شماره تلفن..."
              class="w-64 pl-10 pr-4 py-2 border border-gray-300 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 text-right"
              v-model="stockSearchQuery"
              @input="applyFilters"
          >
          <svg class="absolute left-3 top-3 w-4 h-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
          </svg>
        </div>
      </div>
    </div>

    <!-- جدول موجودی -->
    <div class="bg-white rounded-xl shadow-md border border-gray-100 overflow-hidden">
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-blue-50">
          <tr class="text-right" dir="rtl">
            <th class="text-center py-3 px-4 text-sm font-medium text-blue-900">
              <input
                  type="checkbox"
                  @change="toggleAllStock"
                  :checked="selectedStockItems.length === stockRequests.length && stockRequests.length > 0"
                  class="w-4 h-4 text-blue-600 rounded border-gray-300 focus:ring-blue-500">
            </th>
            <th class="text-right py-3 px-4 text-sm font-medium text-blue-900" style="text-align: right;">محصول</th>
            <th class="text-right py-3 px-4 text-sm font-medium text-blue-900" style="text-align: right;">تاریخ درخواست</th>
            <th class="text-right py-3 px-4 text-sm font-medium text-blue-900" style="text-align: right;">قیمت</th>
            <th class="text-right py-3 px-4 text-sm font-medium text-blue-900" style="text-align: right;">روش اطلاع‌رسانی</th>
            <th class="text-right py-3 px-4 text-sm font-medium text-blue-900" style="text-align: right;">متقاضی</th>
            <th class="text-right py-3 px-4 text-sm font-medium text-blue-900" style="text-align: right;">وضعیت</th>
            <th class="text-right py-3 pl-4 text-sm font-medium text-blue-900" style="text-align: right; padding-right: 10rem;">عملیات</th>
          </tr>
          </thead>
          <tbody class="divide-y divide-gray-100">
          <template v-for="item in stockRequests" :key="item.id">
            <tr class="hover:bg-gray-50">
              <td class="py-3 px-4 text-center">
                <input
                    type="checkbox"
                    :value="item.id"
                    v-model="selectedStockItems"
                    class="w-4 h-4 text-blue-600 rounded border-gray-300 focus:ring-blue-500">
              </td>
              <td class="py-3 px-4 text-sm text-gray-900 text-right">
                {{ item.product }}
              </td>
              <td class="py-3 px-4 text-sm text-gray-600 text-right">
                <div>{{ item.date }}</div>
                <div class="text-xs text-gray-400">{{ item.time }}</div>
              </td>
              <td class="py-3 px-4 text-sm text-gray-900 font-medium text-right">
                {{ item.price.toLocaleString() }} تومان
              </td>
              <td class="py-3 px-4 text-right">
                      <span class="inline-flex px-3 py-1 text-xs font-semibold bg-green-100 text-green-800 rounded-full">
                        پیامک
                      </span>
              </td>
              <td class="py-3 px-4 text-sm text-blue-600 font-medium text-right">
                {{ item.phone }}
              </td>
              <td class="py-3 px-4 text-right">
                      <span :class="getStatusClass(item.status)">
                        {{ item.status }}
                      </span>
              </td>
              <td class="py-3 px-4">
                <div class="flex items-center gap-2 justify-end">
                  <TemplateButton
                    bgGradient="bg-gradient-to-r from-red-400 to-red-600"
                    textColor="text-white"
                    hoverClass="hover:from-red-500 hover:to-red-700 hover:shadow-lg hover:scale-105"
                    focusClass="focus:ring-2 focus:ring-offset-2 focus:ring-red-500"
                    size="small"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                    </svg>
                  </TemplateButton>

                  <!-- دکمه ارسال اطلاع‌رسانی -->
                  <TemplateButton
                    v-if="!stockNotificationStates[item.id]?.success"
                    @click="sendStockNotification(item.id)"
                    :disabled="stockNotificationStates[item.id]?.isLoading"
                    :bgGradient="stockNotificationStates[item.id]?.isLoading ? 'bg-gradient-to-r from-gray-400 to-gray-500' : 'bg-gradient-to-r from-green-400 to-green-600'"
                    textColor="text-white"
                    :hoverClass="stockNotificationStates[item.id]?.isLoading ? '' : 'hover:from-green-500 hover:to-green-700 hover:shadow-lg hover:scale-105'"
                    :focusClass="stockNotificationStates[item.id]?.isLoading ? '' : 'focus:ring-2 focus:ring-offset-2 focus:ring-green-500'"
                    size="small"
                  >
                    <svg v-if="stockNotificationStates[item.id]?.isLoading" class="w-3 h-3 animate-spin" fill="none" viewBox="0 0 24 24">
                      <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                      <path class="opacity-75" fill="currentColor" d="m4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                    </svg>
                    <span>{{ stockNotificationStates[item.id]?.isLoading ? 'در حال ارسال...' : 'ارسال اطلاع رسانی' }}</span>
                  </TemplateButton>

                  <!-- نمایش موفقیت -->
                  <span v-if="stockNotificationStates[item.id]?.success" class="inline-flex px-3 py-1 text-xs font-semibold bg-green-100 text-green-800 rounded-full">
                          ✓ ارسال شد
                        </span>
                </div>
              </td>
            </tr>

            <!-- ردیف خطا -->
            <tr v-if="stockNotificationStates[item.id]?.error" class="bg-red-50">
              <td colspan="8" class="py-3 px-4">
                <div class="flex items-center justify-between">
                  <div class="flex items-center gap-2 text-red-700">
                    <svg class="w-4 h-4 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                    </svg>
                    <span class="text-sm font-medium">{{ stockNotificationStates[item.id]?.error }}</span>
                  </div>
                  <button
                      @click="clearStockError(item.id)"
                      class="text-red-500 hover:text-red-700 transition-colors">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                    </svg>
                  </button>
                </div>
              </td>
            </tr>
          </template>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>

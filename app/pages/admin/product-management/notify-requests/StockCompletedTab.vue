<script setup lang="ts">
import TemplateButton from '@/components/common/TemplateButton.vue';
import { useNotifyRequests } from '~/composables/useNotifyRequests';

interface Props {
  activeNavItem: string
}

const props = defineProps<Props>()

const { 
  selectedStockCompletedItems,
  stockCompletedRequests,
  toggleAllStockCompleted,
  bulkDeleteStockCompleted,
  bulkResendStockNotification,
  getStatusClass
} = useNotifyRequests()
</script>

<template>

  <!-- تب اطلاع‌رسانی شده موجودی -->
  <div v-if="props.activeNavItem === 'stock-completed'" class="w-full">
    <!-- فیلترها و عملیات -->
    <div v-if="selectedStockCompletedItems.length > 0" class="bg-white rounded-xl shadow-md border border-gray-100 p-6 mb-6">
      <div class="flex flex-row-reverse items-center gap-3">
        <TemplateButton
          bg-gradient="bg-gradient-to-r from-red-400 to-red-600"
          text-color="text-white"
          hover-class="hover:from-red-500 hover:to-red-700 hover:shadow-lg hover:scale-105"
          focus-class="focus:ring-2 focus:ring-offset-2 focus:ring-red-500"
          size="small"
          @click="bulkDeleteStockCompleted"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
          </svg>
          <span>حذف ({{ selectedStockCompletedItems.length }})</span>
        </TemplateButton>
        <TemplateButton
          bg-gradient="bg-gradient-to-r from-blue-400 to-blue-600"
          text-color="text-white"
          hover-class="hover:from-blue-500 hover:to-blue-700 hover:shadow-lg hover:scale-105"
          focus-class="focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          size="small"
          @click="bulkResendStockNotification"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
          </svg>
          <span>اطلاع‌رسانی مجدد ({{ selectedStockCompletedItems.length }})</span>
        </TemplateButton>
      </div>
    </div>

    <div class="bg-white rounded-xl shadow-md border border-gray-100 overflow-hidden">
      <div class="bg-gray-50 px-6 py-4 border-b border-gray-100">
        <h3 class="text-lg font-semibold text-gray-900">اطلاع‌رسانی‌های موفق موجودی</h3>
        <p class="text-sm text-gray-600 mt-1">{{ stockCompletedRequests.length }} مورد</p>
      </div>
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-blue-50">
          <tr class="text-right" dir="rtl">
            <th class="text-center py-3 px-4 text-sm font-medium text-blue-900">
              <input
                  type="checkbox"
                  :checked="selectedStockCompletedItems.length === stockCompletedRequests.length && stockCompletedRequests.length > 0"
                  class="w-4 h-4 text-blue-600 rounded border-gray-300 focus:ring-blue-500"
                  @change="toggleAllStockCompleted">
            </th>
            <th class="text-right py-3 px-4 text-sm font-medium text-blue-900" style="text-align: right;">محصول</th>
            <th class="text-right py-3 px-4 text-sm font-medium text-blue-900" style="text-align: right;">تاریخ درخواست</th>
            <th class="text-right py-3 px-4 text-sm font-medium text-blue-900" style="text-align: right;">تاریخ اطلاع‌رسانی</th>
            <th class="text-right py-3 px-4 text-sm font-medium text-blue-900" style="text-align: right;">قیمت</th>
            <th class="text-right py-3 px-4 text-sm font-medium text-blue-900" style="text-align: right;">متقاضی</th>
            <th class="text-right py-3 px-4 text-sm font-medium text-blue-900" style="text-align: right;">وضعیت</th>
          </tr>
          </thead>
          <tbody class="divide-y divide-gray-100">
          <tr v-for="item in stockCompletedRequests" :key="item.id" class="hover:bg-gray-50">
            <td class="py-3 px-4 text-center">
              <input
                  v-model="selectedStockCompletedItems"
                  type="checkbox"
                  :value="item.id"
                  class="w-4 h-4 text-blue-600 rounded border-gray-300 focus:ring-blue-500">
            </td>
            <td class="py-3 px-4 text-sm text-gray-900 text-right">
              {{ item.product }}
            </td>
            <td class="py-3 px-4 text-sm text-gray-600 text-right">
              <div>{{ item.date }}</div>
              <div class="text-xs text-gray-400">{{ item.time }}</div>
            </td>
            <td class="py-3 px-4 text-sm text-green-600 font-medium text-right">
              <div>{{ item.notificationDate }}</div>
              <div class="text-xs text-green-400">{{ item.notificationTime }}</div>
            </td>
            <td class="py-3 px-4 text-sm text-gray-900 font-medium text-right">
              {{ item.price.toLocaleString() }} تومان
            </td>
            <td class="py-3 px-4 text-sm text-blue-600 font-medium text-right">
              {{ item.phone }}
            </td>
            <td class="py-3 px-4 text-right">
                     <span :class="getStatusClass(item.status)">
                       {{ item.status }}
                     </span>
            </td>
          </tr>
          <tr v-if="stockCompletedRequests.length === 0">
            <td colspan="7" class="py-8 px-4 text-center text-gray-500">
              <div class="flex flex-col items-center">
                <svg class="w-12 h-12 text-gray-300 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
                <p class="text-lg font-medium">هنوز اطلاع‌رسانی موفقی ثبت نشده</p>
                <p class="text-sm">اطلاع‌رسانی‌های موفق اینجا نمایش داده می‌شوند</p>
              </div>
            </td>
          </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>

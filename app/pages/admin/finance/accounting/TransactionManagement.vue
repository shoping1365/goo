<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200">
    <!-- هدر بخش -->
    <div class="px-6 py-4 border-b border-gray-200">
      <div class="flex items-center justify-between">
        <div>
          <h3 class="text-lg font-semibold text-gray-900">مدیریت تراکنش‌ها</h3>
          <p class="text-sm text-gray-600 mt-1">ثبت خودکار و مدیریت تراکنش‌های مالی</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button
            @click="refreshData"
            class="inline-flex items-center px-3 py-2 border border-gray-300 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          >
            <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
            </svg>
            بروزرسانی
          </button>
          <button
            @click="saveSettings"
            class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          >
            ذخیره تنظیمات
          </button>
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
      <!-- تب ثبت خودکار فاکتورها -->
      <div v-if="activeTab === 'invoices'" class="space-y-6">
        <InvoiceAutoRegistration />
      </div>

      <!-- تب ثبت خودکار پرداخت‌ها -->
      <div v-if="activeTab === 'payments'" class="space-y-6">
        <PaymentAutoRegistration />
      </div>

      <!-- تب ثبت خودکار برگشت‌ها -->
      <div v-if="activeTab === 'returns'" class="space-y-6">
        <ReturnAutoRegistration />
      </div>

      <!-- تب ثبت خودکار تخفیف‌ها -->
      <div v-if="activeTab === 'discounts'" class="space-y-6">
        <DiscountAutoRegistration />
      </div>

      <!-- تب مدیریت تراکنش‌های دوطرفه -->
      <div v-if="activeTab === 'dual'" class="space-y-6">
        <DualTransactionManagement />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// Import کامپوننت‌های زیربخش‌ها
// @ts-ignore
import InvoiceAutoRegistration from './transaction/InvoiceAutoRegistration.vue'
// @ts-ignore
import PaymentAutoRegistration from './transaction/PaymentAutoRegistration.vue'
// @ts-ignore
import ReturnAutoRegistration from './transaction/ReturnAutoRegistration.vue'
// @ts-ignore
import DiscountAutoRegistration from './transaction/DiscountAutoRegistration.vue'
// @ts-ignore
import DualTransactionManagement from './transaction/DualTransactionManagement.vue'

// تعریف تب‌های مختلف
const tabs = ref([
  { id: 'invoices', name: 'ثبت خودکار فاکتورها' },
  { id: 'payments', name: 'ثبت خودکار پرداخت‌ها' },
  { id: 'returns', name: 'ثبت خودکار برگشت‌ها' },
  { id: 'discounts', name: 'ثبت خودکار تخفیف‌ها' },
  { id: 'dual', name: 'تراکنش‌های دوطرفه' }
])

// تب فعال
const activeTab = ref('invoices')

// بروزرسانی داده‌ها
const refreshData = () => {
  // TODO: بروزرسانی داده‌ها از سرور
  console.log('بروزرسانی داده‌های مدیریت تراکنش‌ها')
}

// ذخیره تنظیمات
const saveSettings = () => {
  // TODO: ذخیره تنظیمات در سرور
  console.log('ذخیره تنظیمات مدیریت تراکنش‌ها')
}
</script>

<!--
  کامپوننت مدیریت تراکنش‌ها
  شامل ۵ بخش اصلی:
  1. ثبت خودکار فاکتورها
  2. ثبت خودکار پرداخت‌ها
  3. ثبت خودکار برگشت‌ها
  4. ثبت خودکار تخفیف‌ها
  5. مدیریت تراکنش‌های دوطرفه
  طراحی مدرن با تب‌بندی
  کاملاً ریسپانسیو
--> 

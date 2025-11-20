<template>
  <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
    <div class="mb-6">
      <h3 class="text-lg font-semibold text-gray-900">نگاشت حساب‌ها</h3>
      <p class="text-sm text-gray-600">مدیریت و نگاشت انواع حساب‌های مالی بین سیستم و نرم‌افزارهای حسابداری</p>
    </div>
    <!-- تب‌ها -->
    <div class="flex flex-wrap gap-2 border-b mb-6">
      <button
v-for="tab in tabs" :key="tab.value"
        :class="[ 'px-4 py-2 rounded-t', activeTab === tab.value ? 'bg-blue-600 text-white' : 'bg-gray-100 text-gray-700 hover:bg-gray-200']"
        @click="activeTab = tab.value"
      >
        {{ tab.label }}
      </button>
    </div>
    <!-- محتوای تب فعال -->
    <div>
      <SalesAccounts v-if="activeTab === 'sales'" />
      <PurchaseAccounts v-else-if="activeTab === 'purchase'" />
      <TaxAccounts v-else-if="activeTab === 'tax'" />
      <BankAccounts v-else-if="activeTab === 'bank'" />
      <AutoAccountMapping v-else-if="activeTab === 'auto'" />
      <AccountCodesManagement v-else-if="activeTab === 'codes'" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
// @ts-ignore
import SalesAccounts from './SalesAccounts.vue'
// @ts-ignore
import PurchaseAccounts from './PurchaseAccounts.vue'
// @ts-ignore
import TaxAccounts from './TaxAccounts.vue'
// @ts-ignore
import BankAccounts from './BankAccounts.vue'
// @ts-ignore
import AutoAccountMapping from './AutoAccountMapping.vue'
// @ts-ignore
import AccountCodesManagement from './AccountCodesManagement.vue'

const tabs = [
  { value: 'sales', label: 'حساب‌های فروش' },
  { value: 'purchase', label: 'حساب‌های خرید' },
  { value: 'tax', label: 'حساب‌های مالیاتی' },
  { value: 'bank', label: 'حساب‌های بانکی' },
  { value: 'auto', label: 'نگاشت خودکار' },
  { value: 'codes', label: 'مدیریت کدهای حساب' }
]
const activeTab = ref('sales')
</script>

<!--
  داشبورد نگاشت حساب‌ها با تب‌بندی
  هر تب یک زیرکامپوننت مستقل دارد
  توضیحات کامل به فارسی
--> 

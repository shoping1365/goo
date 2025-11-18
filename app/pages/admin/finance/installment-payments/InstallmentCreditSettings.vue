<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
    <div class="flex items-center justify-between mb-6">
      <h3 class="text-lg font-semibold text-gray-900">تنظیمات اعتباری</h3>
      <button @click="saveSettings" class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700">
        ذخیره تنظیمات
      </button>
    </div>

    <!-- Credit Score Settings -->
    <div class="mb-6">
      <h4 class="text-md font-semibold text-gray-900 mb-4">محاسبه امتیاز اعتباری</h4>
      <div class="space-y-4">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداقل امتیاز پذیرش</label>
            <input v-model="settings.minCreditScore" type="number" class="w-full border border-gray-300 rounded-md px-3 py-2">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر امتیاز</label>
            <input v-model="settings.maxCreditScore" type="number" class="w-full border border-gray-300 rounded-md px-3 py-2">
          </div>
        </div>
        
        <div>
          <h5 class="text-sm font-semibold text-gray-900 mb-3">وزن فاکتورها</h5>
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            <div>
              <label class="block text-sm text-gray-700 mb-1">سن ({{ settings.weights.age }}%)</label>
              <input v-model="settings.weights.age" type="range" min="0" max="100" class="w-full">
            </div>
            <div>
              <label class="block text-sm text-gray-700 mb-1">درآمد ({{ settings.weights.income }}%)</label>
              <input v-model="settings.weights.income" type="range" min="0" max="100" class="w-full">
            </div>
            <div>
              <label class="block text-sm text-gray-700 mb-1">سابقه ({{ settings.weights.history }}%)</label>
              <input v-model="settings.weights.history" type="range" min="0" max="100" class="w-full">
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Approval Rules -->
    <div class="mb-6">
      <h4 class="text-md font-semibold text-gray-900 mb-4">قوانین تایید</h4>
      <div class="space-y-4">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداقل سن</label>
            <input v-model="settings.rules.minAge" type="number" class="w-full border border-gray-300 rounded-md px-3 py-2">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر سن</label>
            <input v-model="settings.rules.maxAge" type="number" class="w-full border border-gray-300 rounded-md px-3 py-2">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداقل درآمد (ریال)</label>
            <input v-model="settings.rules.minIncome" type="number" class="w-full border border-gray-300 rounded-md px-3 py-2">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر مبلغ قابل تایید</label>
            <input v-model="settings.rules.maxAmount" type="number" class="w-full border border-gray-300 rounded-md px-3 py-2">
          </div>
        </div>
      </div>
    </div>

    <!-- Risk Assessment -->
    <div class="mb-6">
      <h4 class="text-md font-semibold text-gray-900 mb-4">ارزیابی ریسک</h4>
      <div class="space-y-4">
        <div>
          <label class="flex items-center">
            <input v-model="settings.riskAssessment.enabled" type="checkbox" class="rounded border-gray-300">
            <span class="mr-2">فعال‌سازی ارزیابی ریسک خودکار</span>
          </label>
        </div>
        
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">ریسک پایین (&lt; {{ settings.riskAssessment.lowThreshold }})</label>
            <input v-model="settings.riskAssessment.lowThreshold" type="number" class="w-full border border-gray-300 rounded-md px-3 py-2">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">ریسک متوسط (&lt; {{ settings.riskAssessment.mediumThreshold }})</label>
            <input v-model="settings.riskAssessment.mediumThreshold" type="number" class="w-full border border-gray-300 rounded-md px-3 py-2">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">ریسک بالا (&gt;= {{ settings.riskAssessment.mediumThreshold }})</label>
            <span class="text-sm text-gray-500">خودکار محاسبه می‌شود</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Current Status -->
    <div class="bg-blue-50 rounded-lg p-6">
      <h4 class="text-md font-semibold text-gray-900 mb-3">وضعیت فعلی سیستم</h4>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6 text-sm">
        <div>
          <span class="text-gray-600 block">درخواست‌های امروز</span>
          <span class="font-medium">{{ status.todayRequests }}</span>
        </div>
        <div>
          <span class="text-gray-600 block">نرخ تایید</span>
          <span class="font-medium">{{ status.approvalRate }}%</span>
        </div>
        <div>
          <span class="text-gray-600 block">میانگین امتیاز</span>
          <span class="font-medium">{{ status.avgScore }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'

const settings = ref({
  minCreditScore: 60,
  maxCreditScore: 100,
  weights: {
    age: 20,
    income: 40,
    history: 40
  },
  rules: {
    minAge: 18,
    maxAge: 65,
    minIncome: 10000000,
    maxAmount: 500000000
  },
  riskAssessment: {
    enabled: true,
    lowThreshold: 30,
    mediumThreshold: 70
  }
})

const status = ref({
  todayRequests: 156,
  approvalRate: 73,
  avgScore: 68
})

const saveSettings = async () => {
  try {
    await $fetch('/api/admin/installment-payments/credit-settings', {
      method: 'PUT',
      body: settings.value
    })
    alert('تنظیمات با موفقیت ذخیره شد')
  } catch (error) {
    console.error('خطا در ذخیره تنظیمات:', error)
    alert('خطا در ذخیره تنظیمات')
  }
}

onMounted(() => {
  // Load settings
})
</script>

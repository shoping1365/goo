<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200 w-full">
    <div class="p-6 border-b border-gray-200">
      <h2 class="text-lg font-semibold text-gray-900">تولید خودکار کوپن</h2>
      <p class="text-gray-600 mt-1">تولید انبوه کوپن با تنظیمات دلخواه</p>
    </div>
    <form class="p-6 space-y-6" @submit.prevent="generateCoupons">
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تعداد کوپن</label>
          <input v-model.number="form.count" type="number" min="1" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm" required>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نوع تخفیف</label>
          <select v-model="form.type" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
            <option value="percentage">درصدی</option>
            <option value="fixed">مبلغ ثابت</option>
            <option value="free_shipping">ارسال رایگان</option>
          </select>
        </div>
        <div v-if="form.type !== 'free_shipping'">
          <label class="block text-sm font-medium text-gray-700 mb-2">مقدار تخفیف</label>
          <input v-model.number="form.value" type="number" min="1" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm" required>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ انقضا</label>
          <input v-model="form.expiry" type="date" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm" required>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">پیشوند کد</label>
          <input v-model="form.prefix" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">پسوند کد</label>
          <input v-model="form.suffix" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
        </div>
      </div>
      <button type="submit" class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors">تولید کوپن‌ها</button>
      <div v-if="coupons.length" class="mt-8">
        <h3 class="text-md font-bold text-gray-900 mb-4">کوپن‌های تولیدشده:</h3>
        <ul class="space-y-2">
          <li v-for="code in coupons" :key="code" class="font-mono bg-gray-100 rounded px-3 py-2">{{ code }}</li>
        </ul>
      </div>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
const form = ref({
  count: 10,
  type: 'percentage',
  value: 10,
  expiry: '',
  prefix: '',
  suffix: ''
})
const coupons = ref<string[]>([])
function generateCoupons() {
  coupons.value = []
  for (let i = 0; i < form.value.count; i++) {
    const code = `${form.value.prefix}${Math.random().toString(36).substring(2, 8).toUpperCase()}${form.value.suffix}`
    coupons.value.push(code)
  }
}
</script> 

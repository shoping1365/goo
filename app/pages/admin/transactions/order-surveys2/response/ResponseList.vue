<template>
  <div class="space-y-6">
    <!-- Header & Filters -->
    <div class="flex flex-col md:flex-row md:items-center md:justify-between gap-6">
      <div>
        <h3 class="text-lg font-semibold text-gray-800">لیست پاسخ‌های دریافتی</h3>
        <p class="text-gray-600 text-sm">نمایش پاسخ‌های مشتریان به نظرسنجی</p>
      </div>
      <div class="flex items-center gap-2">
        <label class="text-sm text-gray-700">فیلتر بر اساس امتیاز:</label>
        <select v-model="filterScore" class="px-3 py-2 border border-gray-300 rounded-lg text-sm">
          <option value="">همه</option>
          <option v-for="score in [5,4,3,2,1]" :key="score" :value="score">{{ score }} ⭐</option>
        </select>
      </div>
    </div>
    <!-- Response Table -->
    <div class="bg-white rounded-lg border border-gray-200 overflow-x-auto">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-right text-xs font-medium text-gray-500">سفارش</th>
            <th class="px-4 py-3 text-right text-xs font-medium text-gray-500">مشتری</th>
            <th class="px-4 py-3 text-right text-xs font-medium text-gray-500">امتیاز</th>
            <th class="px-4 py-3 text-right text-xs font-medium text-gray-500">نظر متنی</th>
            <th class="px-4 py-3 text-right text-xs font-medium text-gray-500">تاریخ</th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-for="resp in filteredResponses" :key="resp.id" class="hover:bg-gray-50">
            <td class="px-4 py-3 whitespace-nowrap text-sm text-gray-900">#{{ resp.orderNumber }}</td>
            <td class="px-4 py-3 whitespace-nowrap text-sm text-gray-900">{{ resp.customerName }}</td>
            <td class="px-4 py-3 whitespace-nowrap">
              <span class="inline-flex items-center gap-1">
                <span v-for="i in resp.score" :key="i" class="text-yellow-400">★</span>
                <span v-for="i in 5-resp.score" :key="i" class="text-gray-300">★</span>
              </span>
            </td>
            <td class="px-4 py-3 whitespace-nowrap text-sm text-gray-700 max-w-xs truncate" :title="resp.comment">{{ resp.comment }}</td>
            <td class="px-4 py-3 whitespace-nowrap text-xs text-gray-500">{{ formatDate(resp.date) }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
<script setup lang="ts">
import { ref, computed } from 'vue'
const responses = ref([
  { id: 1, orderNumber: '1001', customerName: 'علی احمدی', score: 5, comment: 'عالی بود، ممنون!', date: '2024-06-01T10:30:00Z' },
  { id: 2, orderNumber: '1002', customerName: 'زهرا رضایی', score: 4, comment: 'خوب بود اما ارسال کمی دیر شد.', date: '2024-06-02T12:10:00Z' },
  { id: 3, orderNumber: '1003', customerName: 'محمد موسوی', score: 2, comment: 'محصول آسیب دیده بود.', date: '2024-06-03T09:20:00Z' },
  { id: 4, orderNumber: '1004', customerName: 'سارا کریمی', score: 5, comment: 'همه چیز عالی و سریع!', date: '2024-06-04T15:45:00Z' },
  { id: 5, orderNumber: '1005', customerName: 'رضا مرادی', score: 3, comment: 'معمولی بود.', date: '2024-06-05T11:00:00Z' }
])
const filterScore = ref('')
const filteredResponses = computed(() => {
  if (!filterScore.value) return responses.value
  return responses.value.filter(r => r.score === Number(filterScore.value))
})
const formatDate = (date: string) => new Date(date).toLocaleDateString('fa-IR')
</script> 

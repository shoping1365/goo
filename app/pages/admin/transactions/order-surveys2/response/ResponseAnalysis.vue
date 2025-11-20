<template>
  <div class="space-y-8">
    <!-- Header -->
    <div>
      <h3 class="text-lg font-semibold text-gray-800">تحلیل پاسخ‌های نظرسنجی</h3>
      <p class="text-gray-600 text-sm">آمار و تحلیل پیشرفته پاسخ‌های مشتریان</p>
    </div>
    <!-- Score Distribution Chart -->
    <div class="bg-white rounded-lg border border-gray-200 px-4 py-4">
      <h4 class="text-md font-semibold text-gray-800 mb-4">نمودار توزیع امتیازات</h4>
      <div class="w-full max-w-md mx-auto">
        <!-- جایگزین با چارت واقعی در آینده -->
        <div class="flex items-end gap-2 h-32">
          <div v-for="score in [5,4,3,2,1]" :key="score" class="flex flex-col items-center flex-1">
            <div :style="{height: scoreCounts[score]*15+'px'}" class="w-8 bg-blue-500 rounded-t-lg transition-all"></div>
            <span class="text-xs mt-1">{{ score }} ⭐</span>
            <span class="text-xs text-gray-500">{{ scoreCounts[score] }}</span>
          </div>
        </div>
      </div>
    </div>
    <!-- Keyword Analysis -->
    <div class="bg-white rounded-lg border border-gray-200 px-4 py-4">
      <h4 class="text-md font-semibold text-gray-800 mb-4">کلمات کلیدی پرتکرار</h4>
      <div class="flex flex-wrap gap-2">
        <span v-for="word in keywords" :key="word" class="px-3 py-1 bg-blue-100 text-blue-800 rounded-full text-xs font-medium">{{ word }}</span>
      </div>
    </div>
    <!-- Satisfaction Report -->
    <div class="bg-white rounded-lg border border-gray-200 px-4 py-4">
      <h4 class="text-md font-semibold text-gray-800 mb-4">گزارش رضایت مشتریان</h4>
      <div class="grid grid-cols-2 md:grid-cols-4 gapx-4 py-4">
        <div class="text-center">
          <div class="text-2xl font-bold text-green-600">{{ percent(5) }}%</div>
          <div class="text-xs text-gray-500">خیلی راضی</div>
        </div>
        <div class="text-center">
          <div class="text-2xl font-bold text-blue-600">{{ percent(4) }}%</div>
          <div class="text-xs text-gray-500">راضی</div>
        </div>
        <div class="text-center">
          <div class="text-2xl font-bold text-yellow-600">{{ percent(3) }}%</div>
          <div class="text-xs text-gray-500">معمولی</div>
        </div>
        <div class="text-center">
          <div class="text-2xl font-bold text-red-600">{{ percent(1,2) }}%</div>
          <div class="text-xs text-gray-500">ناراضی</div>
        </div>
      </div>
    </div>
    <!-- Period Comparison -->
    <div class="bg-white rounded-lg border border-gray-200 px-4 py-4">
      <h4 class="text-md font-semibold text-gray-800 mb-4">مقایسه دوره‌ای</h4>
      <div class="flex flex-col md:flex-row gapx-4 py-4 md:items-end">
        <div>
          <label class="block text-xs text-gray-700 mb-1">دوره اول</label>
          <input v-model="period1" type="month" class="px-3 py-2 border border-gray-300 rounded-lg">
        </div>
        <div>
          <label class="block text-xs text-gray-700 mb-1">دوره دوم</label>
          <input v-model="period2" type="month" class="px-3 py-2 border border-gray-300 rounded-lg">
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg text-sm" @click="comparePeriods">مقایسه</button>
      </div>
      <div v-if="comparisonResult" class="mt-4">
        <div class="text-sm text-gray-700">درصد رضایت دوره اول: <span class="font-bold">{{ comparisonResult.p1 }}%</span></div>
        <div class="text-sm text-gray-700">درصد رضایت دوره دوم: <span class="font-bold">{{ comparisonResult.p2 }}%</span></div>
        <div class="text-sm text-gray-700">تغییر: <span :class="comparisonResult.diff >= 0 ? 'text-green-600' : 'text-red-600'">{{ comparisonResult.diff }}%</span></div>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import { ref, computed } from 'vue'
const responses = ref([
  { id: 1, score: 5, comment: 'عالی بود، ممنون!', date: '2024-06-01T10:30:00Z' },
  { id: 2, score: 4, comment: 'خوب بود اما ارسال کمی دیر شد.', date: '2024-06-02T12:10:00Z' },
  { id: 3, score: 2, comment: 'محصول آسیب دیده بود.', date: '2024-06-03T09:20:00Z' },
  { id: 4, score: 5, comment: 'همه چیز عالی و سریع!', date: '2024-06-04T15:45:00Z' },
  { id: 5, score: 3, comment: 'معمولی بود.', date: '2024-06-05T11:00:00Z' },
  { id: 6, score: 1, comment: 'خیلی ناراضی بودم.', date: '2024-05-20T11:00:00Z' }
])
const scoreCounts = computed(() => {
  const counts: Record<number, number> = { 1: 0, 2: 0, 3: 0, 4: 0, 5: 0 }
  responses.value.forEach(r => counts[r.score]++)
  return counts
})
const keywords = computed(() => {
  // استخراج کلمات پرتکرار ساده (نمونه اولیه)
  const all = responses.value.map(r => r.comment).join(' ').replace(/[!،.]/g, '').split(' ')
  const freq: Record<string, number> = {}
  all.forEach(w => { if (w.length > 2) freq[w] = (freq[w] || 0) + 1 })
  return Object.entries(freq).sort((a,b) => b[1]-a[1]).slice(0, 8).map(([w]) => w)
})
const percent = (...scores: number[]) => {
  const total = responses.value.length
  const count = responses.value.filter(r => scores.includes(r.score)).length
  return total ? Math.round((count/total)*100) : 0
}
// مقایسه دوره‌ای (نمونه اولیه)
const period1 = ref('2024-06')
const period2 = ref('2024-05')
const comparisonResult = ref<null|{p1:number,p2:number,diff:number}>(null)
const comparePeriods = () => {
  const getPercent = (period: string) => {
    const filtered = responses.value.filter(r => r.date.startsWith(period))
    if (!filtered.length) return 0
    const satisfied = filtered.filter(r => r.score >= 4).length
    return Math.round((satisfied/filtered.length)*100)
  }
  const p1 = getPercent(period1.value)
  const p2 = getPercent(period2.value)
  comparisonResult.value = { p1, p2, diff: p1-p2 }
}
</script> 

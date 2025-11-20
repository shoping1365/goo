<template>
  <div class="bg-white rounded-xl shadow-lg p-6">
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 xl:grid-cols-6 gap-6">
      <!-- Date Range -->
      <div class="lg:col-span-2">
        <label class="block text-sm font-medium text-gray-700 mb-2">بازه زمانی</label>
        <div class="flex gap-2">
          <input
            v-model="filters.dateFrom"
            type="date"
            class="flex-1 border border-gray-300 rounded-lg px-3 py-2 text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            placeholder="از تاریخ"
          />
          <input
            v-model="filters.dateTo"
            type="date"
            class="flex-1 border border-gray-300 rounded-lg px-3 py-2 text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            placeholder="تا تاریخ"
          />
        </div>
      </div>

      <!-- City Filter -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">شهر</label>
        <select
          v-model="filters.city"
          class="w-full border border-gray-300 rounded-lg px-3 py-2 text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
        >
          <option value="">همه شهرها</option>
          <option value="تهران">تهران</option>
          <option value="اصفهان">اصفهان</option>
          <option value="مشهد">مشهد</option>
          <option value="شیراز">شیراز</option>
          <option value="تبریز">تبریز</option>
        </select>
      </div>

      <!-- Device Filter -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">دستگاه</label>
        <select
          v-model="filters.device"
          class="w-full border border-gray-300 rounded-lg px-3 py-2 text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
        >
          <option value="">همه دستگاه‌ها</option>
          <option value="موبایل">موبایل</option>
          <option value="دسکتاپ">دسکتاپ</option>
          <option value="تبلت">تبلت</option>
        </select>
      </div>

      <!-- Referrer Filter -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">منبع ترافیک</label>
        <select
          v-model="filters.referrer"
          class="w-full border border-gray-300 rounded-lg px-3 py-2 text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
        >
          <option value="">همه منابع</option>
          <option value="Google">Google</option>
          <option value="مستقیم">مستقیم</option>
          <option value="Instagram">Instagram</option>
          <option value="Telegram">Telegram</option>
          <option value="Facebook">Facebook</option>
        </select>
      </div>

      <!-- Action Buttons -->
      <div class="flex gap-2">
        <button
          class="flex-1 bg-blue-600 text-white px-4 py-2 rounded-lg text-sm font-medium hover:bg-blue-700 transition-colors duration-200"
          @click="applyFilters"
        >
          اعمال فیلتر
        </button>
        <button
          class="px-4 py-2 border border-gray-300 text-gray-700 rounded-lg text-sm font-medium hover:bg-gray-50 transition-colors duration-200"
          @click="resetFilters"
        >
          پاک کردن
        </button>
      </div>
    </div>

    <!-- Quick Filters -->
    <div class="mt-4 flex flex-wrap gap-2">
      <button
        v-for="quick in quickFilters"
        :key="quick.label"
        class="px-3 py-1 text-sm bg-gray-100 text-gray-700 rounded-full hover:bg-gray-200 transition-colors duration-200"
        :class="{ 'bg-blue-100 text-blue-700': isQuickFilterActive(quick) }"
        @click="applyQuickFilter(quick)"
      >
        {{ quick.label }}
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue';

const emit = defineEmits<{
  filter: [filters: Record<string, unknown>]
}>()

const filters = reactive({
  dateFrom: '',
  dateTo: '',
  city: '',
  device: '',
  referrer: ''
})

const quickFilters = [
  { label: 'امروز', days: 0 },
  { label: 'دیروز', days: 1 },
  { label: '7 روز گذشته', days: 7 },
  { label: '30 روز گذشته', days: 30 },
  { label: '3 ماه گذشته', days: 90 }
]

const activeQuickFilter = ref('')

const applyFilters = () => {
  emit('filter', { ...filters })
}

const resetFilters = () => {
  Object.assign(filters, {
    dateFrom: '',
    dateTo: '',
    city: '',
    device: '',
    referrer: ''
  })
  activeQuickFilter.value = ''
  emit('filter', { ...filters })
}

const applyQuickFilter = (quick: { label: string, days: number }) => {
  const today = new Date()
  const fromDate = new Date(today)
  fromDate.setDate(today.getDate() - quick.days)
  
  filters.dateFrom = fromDate.toISOString().split('T')[0]
  filters.dateTo = today.toISOString().split('T')[0]
  
  activeQuickFilter.value = quick.label
  emit('filter', { ...filters })
}

const isQuickFilterActive = (quick: { label: string, days: number }) => {
  return activeQuickFilter.value === quick.label
}
</script> 

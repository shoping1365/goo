<template>
  <div class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
    <div class="relative top-20 mx-auto p-5 border w-11/12 md:w-4/5 lg:w-3/4 shadow-lg rounded-md bg-white">
      <div class="mt-3">
        <!-- هدر مودال -->
        <div class="flex justify-between items-center mb-6">
          <h3 class="text-lg font-medium text-gray-900">گزارش‌های گرونه شانس</h3>
          <div class="flex items-center space-x-3 space-x-reverse">
            <button 
              class="bg-green-600 text-white px-4 py-2 rounded-lg hover:bg-green-700 transition-colors"
              @click="exportReport"
            >
              <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
              </svg>
              خروجی Excel
            </button>
            <button 
              class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors"
              @click="printReport"
            >
              <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 17h2a2 2 0 002-2v-4a2 2 0 00-2-2H5a2 2 0 00-2 2v4a2 2 0 002 2h2m2 4h6a2 2 0 002-2v-4a2 2 0 00-2-2H9a2 2 0 00-2 2v4a2 2 0 002 2zm8-12V5a2 2 0 00-2-2H9a2 2 0 00-2 2v4h10z"></path>
              </svg>
              چاپ
            </button>
            <button 
              class="text-gray-400 hover:text-gray-600"
              @click="$emit('close')"
            >
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
              </svg>
            </button>
          </div>
        </div>

        <!-- فیلترها -->
        <div class="bg-gray-50 p-6 rounded-lg mb-6">
          <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">بازه زمانی</label>
              <select 
                v-model="filters.timeRange"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              >
                <option value="7days">7 روز گذشته</option>
                <option value="30days">30 روز گذشته</option>
                <option value="90days">90 روز گذشته</option>
                <option value="1year">یک سال گذشته</option>
                <option value="custom">سفارشی</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نوع گرونه</label>
              <select 
                v-model="filters.drawType"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              >
                <option value="">همه انواع</option>
                <option value="daily">روزانه</option>
                <option value="weekly">هفتگی</option>
                <option value="monthly">ماهانه</option>
                <option value="purchase">خرید</option>
                <option value="registration">ثبت‌نام</option>
              </select>
            </div>
            <div class="flex items-end">
              <button 
                class="w-full bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition-colors"
                @click="applyFilters"
              >
                اعمال فیلتر
              </button>
            </div>
          </div>
        </div>

        <!-- آمار کلی -->
        <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
          <div class="bg-blue-50 p-6 rounded-lg border border-blue-200">
            <div class="text-center">
              <p class="text-sm text-blue-600">کل شرکت‌کنندگان</p>
              <p class="text-2xl font-bold text-blue-900">{{ reportStats.totalParticipants }}</p>
            </div>
          </div>
          <div class="bg-green-50 p-6 rounded-lg border border-green-200">
            <div class="text-center">
              <p class="text-sm text-green-600">جوایز اهدا شده</p>
              <p class="text-2xl font-bold text-green-900">{{ reportStats.totalPrizes }}</p>
            </div>
          </div>
          <div class="bg-purple-50 p-6 rounded-lg border border-purple-200">
            <div class="text-center">
              <p class="text-sm text-purple-600">نرخ تبدیل</p>
              <p class="text-2xl font-bold text-purple-900">{{ reportStats.conversionRate }}%</p>
            </div>
          </div>
          <div class="bg-yellow-50 p-6 rounded-lg border border-yellow-200">
            <div class="text-center">
              <p class="text-sm text-yellow-600">ارزش کل جوایز</p>
              <p class="text-2xl font-bold text-yellow-900">{{ formatCurrency(reportStats.totalPrizeValue) }}</p>
            </div>
          </div>
        </div>

        <!-- نمودار روند شرکت‌کنندگان -->
        <div class="bg-white rounded-lg border p-6 mb-6">
          <h4 class="text-lg font-medium text-gray-900 mb-4">روند شرکت‌کنندگان</h4>
          <div class="h-64 bg-gray-50 rounded-lg flex items-center justify-center">
            <div class="text-center text-gray-500">
              <svg class="w-16 h-16 mx-auto mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
              </svg>
              <p>نمودار روند شرکت‌کنندگان</p>
              <p class="text-sm">در اینجا نمودار نمایش داده می‌شود</p>
            </div>
          </div>
        </div>

        <!-- جدول بهترین گرونه‌ها -->
        <div class="bg-white rounded-lg border">
          <div class="px-6 py-4 border-b border-gray-200">
            <h4 class="text-lg font-medium text-gray-900">بهترین گرونه‌ها</h4>
          </div>
          <div class="overflow-x-auto">
            <table class="min-w-full divide-y divide-gray-200">
              <thead class="bg-gray-50">
                <tr>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نام گرونه</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">شرکت‌کنندگان</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">جوایز اهدا شده</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نرخ تبدیل</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">امتیاز</th>
                </tr>
              </thead>
              <tbody class="bg-white divide-y divide-gray-200">
                <tr v-for="draw in topDraws" :key="draw.id">
                  <td class="px-6 py-4 whitespace-nowrap">
                    <div class="text-sm font-medium text-gray-900">{{ draw.name }}</div>
                    <div class="text-sm text-gray-500">{{ draw.description }}</div>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap">
                    <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full" :class="getTypeBadgeClass(draw.type)">
                      {{ getTypeText(draw.type) }}
                    </span>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ draw.participants }}</td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ draw.awardedPrizes }}</td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ draw.conversionRate }}%</td>
                  <td class="px-6 py-4 whitespace-nowrap">
                    <div class="flex items-center">
                      <div class="flex items-center">
                        <svg 
                          v-for="star in 5" 
                          :key="star"
                          :class="star <= draw.score ? 'text-yellow-400' : 'text-gray-300'"
                          class="w-4 h-4" 
                          fill="currentColor" 
                          viewBox="0 0 20 20"
                        >
                          <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"></path>
                        </svg>
                      </div>
                      <span class="text-sm text-gray-900 mr-2">{{ draw.score }}/5</span>
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
defineEmits(['close'])

// متغیرهای reactive
const filters = ref({
  timeRange: '30days',
  drawType: ''
})

// آمار گزارش
const reportStats = ref({
  totalParticipants: 2847,
  totalPrizes: 156,
  conversionRate: 5.5,
  totalPrizeValue: 12500000
})

// بهترین گرونه‌ها
const topDraws = ref([
  {
    id: 1,
    name: 'گرونه هفتگی',
    description: 'قرعه‌کشی هفتگی برای مشتریان وفادار',
    type: 'weekly',
    participants: 1250,
    awardedPrizes: 25,
    conversionRate: 2.0,
    score: 5
  },
  {
    id: 2,
    name: 'گرونه خرید',
    description: 'قرعه‌کشی برای هر خرید بالای 500 هزار تومان',
    type: 'purchase',
    participants: 890,
    awardedPrizes: 45,
    conversionRate: 5.1,
    score: 4
  },
  {
    id: 3,
    name: 'گرونه ثبت‌نام',
    description: 'قرعه‌کشی برای کاربران جدید',
    type: 'registration',
    participants: 707,
    awardedPrizes: 86,
    conversionRate: 12.2,
    score: 5
  }
])

// توابع کمکی
const formatCurrency = (amount) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

const getTypeBadgeClass = (type) => {
  const classes = {
    daily: 'bg-blue-100 text-blue-800',
    weekly: 'bg-green-100 text-green-800',
    monthly: 'bg-purple-100 text-purple-800',
    purchase: 'bg-orange-100 text-orange-800',
    registration: 'bg-pink-100 text-pink-800'
  }
  return classes[type] || 'bg-gray-100 text-gray-800'
}

const getTypeText = (type) => {
  const texts = {
    daily: 'روزانه',
    weekly: 'هفتگی',
    monthly: 'ماهانه',
    purchase: 'خرید',
    registration: 'ثبت‌نام'
  }
  return texts[type] || type
}

// توابع عملیات
const applyFilters = () => {
  // در اینجا فیلترها اعمال می‌شوند
}

const exportReport = () => {
  // شبیه‌سازی خروجی Excel
  alert('گزارش در حال دانلود است...')
}

const printReport = () => {
  // شبیه‌سازی چاپ
  window.print()
}
</script>

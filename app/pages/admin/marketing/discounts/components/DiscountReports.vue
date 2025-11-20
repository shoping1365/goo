<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200">
    <!-- هدر بخش -->
    <div class="p-6 border-b border-gray-200">
      <div class="flex items-center justify-between">
        <div>
          <h2 class="text-lg font-semibold text-gray-900">گزارش‌ها و آمار</h2>
          <p class="text-gray-600 mt-1">آمار و گزارش‌های مربوط به کوپن‌ها و کمپین‌ها</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors" @click="exportReport">
            <span class="i-heroicons-arrow-down-tray ml-2"></span>
            خروجی اکسل
          </button>
          <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors" @click="refreshData">
            <span class="i-heroicons-arrow-path ml-2"></span>
            بروزرسانی
          </button>
        </div>
      </div>
    </div>

    <!-- فیلترهای گزارش -->
    <div class="p-6 border-b border-gray-200">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">بازه زمانی</label>
          <select v-model="filters.period" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500" @change="applyFilters">
            <option value="today">امروز</option>
            <option value="week">این هفته</option>
            <option value="month">این ماه</option>
            <option value="quarter">این فصل</option>
            <option value="year">امسال</option>
            <option value="custom">سفارشی</option>
          </select>
        </div>
        <div v-if="filters.period === 'custom'">
          <label class="block text-sm font-medium text-gray-700 mb-2">از تاریخ</label>
          <input v-model="filters.startDate" type="date" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500" @change="applyFilters">
        </div>
        <div v-if="filters.period === 'custom'">
          <label class="block text-sm font-medium text-gray-700 mb-2">تا تاریخ</label>
          <input v-model="filters.endDate" type="date" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500" @change="applyFilters">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نوع گزارش</label>
          <select v-model="filters.reportType" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500" @change="applyFilters">
            <option value="overview">نمای کلی</option>
            <option value="coupons">کوپن‌ها</option>
            <option value="campaigns">کمپین‌ها</option>
            <option value="performance">عملکرد</option>
          </select>
        </div>
      </div>
    </div>

    <!-- آمار کلی -->
    <div class="p-6">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
        <div class="bg-blue-50 p-6 rounded-lg">
          <div class="flex items-center">
            <div class="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center ml-3">
              <span class="i-heroicons-ticket text-blue-600 text-xl"></span>
            </div>
            <div>
              <p class="text-sm text-blue-600">کوپن‌های فعال</p>
              <p class="text-2xl font-bold text-blue-900">{{ stats.activeCoupons }}</p>
            </div>
          </div>
        </div>
        
        <div class="bg-green-50 p-6 rounded-lg">
          <div class="flex items-center">
            <div class="w-12 h-12 bg-green-100 rounded-lg flex items-center justify-center ml-3">
              <span class="i-heroicons-megaphone text-green-600 text-xl"></span>
            </div>
            <div>
              <p class="text-sm text-green-600">کمپین‌های فعال</p>
              <p class="text-2xl font-bold text-green-900">{{ stats.activeCampaigns }}</p>
            </div>
          </div>
        </div>
        
        <div class="bg-purple-50 p-6 rounded-lg">
          <div class="flex items-center">
            <div class="w-12 h-12 bg-purple-100 rounded-lg flex items-center justify-center ml-3">
              <span class="i-heroicons-currency-dollar text-purple-600 text-xl"></span>
            </div>
            <div>
              <p class="text-sm text-purple-600">تخفیف کل</p>
              <p class="text-2xl font-bold text-purple-900">{{ formatPrice(stats.totalDiscount) }}</p>
            </div>
          </div>
        </div>
        
        <div class="bg-orange-50 p-6 rounded-lg">
          <div class="flex items-center">
            <div class="w-12 h-12 bg-orange-100 rounded-lg flex items-center justify-center ml-3">
              <span class="i-heroicons-shopping-cart text-orange-600 text-xl"></span>
            </div>
            <div>
              <p class="text-sm text-orange-600">سفارشات</p>
              <p class="text-2xl font-bold text-orange-900">{{ formatNumber(stats.totalOrders) }}</p>
            </div>
          </div>
        </div>
      </div>

      <!-- نمودارها -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-8">
        <!-- نمودار استفاده از کوپن‌ها -->
        <div class="bg-white border border-gray-200 rounded-lg p-6">
          <h3 class="text-lg font-medium text-gray-900 mb-4">استفاده از کوپن‌ها</h3>
          <div class="space-y-3">
            <div v-for="coupon in topCoupons" :key="coupon.id" class="flex items-center justify-between">
              <div class="flex items-center">
                <div class="w-3 h-3 rounded-full ml-2" :style="{ backgroundColor: coupon.color }"></div>
                <span class="text-sm text-gray-700">{{ coupon.name }}</span>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <div class="w-32 bg-gray-200 rounded-full h-2">
                  <div class="bg-blue-600 h-2 rounded-full" :style="{ width: coupon.usageRate + '%' }"></div>
                </div>
                <span class="text-sm text-gray-600">{{ coupon.usageRate }}%</span>
              </div>
            </div>
          </div>
        </div>

        <!-- نمودار عملکرد کمپین‌ها -->
        <div class="bg-white border border-gray-200 rounded-lg p-6">
          <h3 class="text-lg font-medium text-gray-900 mb-4">عملکرد کمپین‌ها</h3>
          <div class="space-y-3">
            <div v-for="campaign in topCampaigns" :key="campaign.id" class="flex items-center justify-between">
              <div class="flex items-center">
                <div class="w-3 h-3 rounded-full ml-2" :style="{ backgroundColor: campaign.color }"></div>
                <span class="text-sm text-gray-700">{{ campaign.name }}</span>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <div class="w-32 bg-gray-200 rounded-full h-2">
                  <div class="bg-green-600 h-2 rounded-full" :style="{ width: campaign.performanceRate + '%' }"></div>
                </div>
                <span class="text-sm text-gray-600">{{ campaign.performanceRate }}%</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- جدول گزارش‌های تفصیلی -->
      <div class="bg-white border border-gray-200 rounded-lg">
        <div class="p-6 border-b border-gray-200">
          <h3 class="text-lg font-medium text-gray-900">گزارش تفصیلی</h3>
        </div>
        
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نام</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">استفاده</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تخفیف کل</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">فروش</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نرخ تبدیل</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">ROI</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="item in detailedReport" :key="item.id" class="hover:bg-gray-50">
                <td class="px-6 py-4 whitespace-nowrap">
                  <div>
                    <div class="text-sm font-medium text-gray-900">{{ item.name }}</div>
                    <div class="text-sm text-gray-500">{{ item.code }}</div>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span :class="`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${getTypeBadgeClass(item.type)}`">
                    {{ getTypeName(item.type) }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  {{ formatNumber(item.usageCount) }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  {{ formatPrice(item.totalDiscount) }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  {{ formatPrice(item.totalSales) }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  {{ item.conversionRate }}%
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  {{ item.roi }}%
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- آمار پیشرفته -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mt-8">
        <div class="bg-white border border-gray-200 rounded-lg p-6">
          <h4 class="font-medium text-gray-900 mb-4">بهترین ساعات استفاده</h4>
          <div class="space-y-2">
            <div v-for="hour in peakHours" :key="hour.hour" class="flex justify-between text-sm">
              <span class="text-gray-600">{{ hour.hour }}:۰۰</span>
              <span class="font-medium">{{ hour.usage }} استفاده</span>
            </div>
          </div>
        </div>
        
        <div class="bg-white border border-gray-200 rounded-lg p-6">
          <h4 class="font-medium text-gray-900 mb-4">محبوب‌ترین محصولات</h4>
          <div class="space-y-2">
            <div v-for="product in topProducts" :key="product.id" class="flex justify-between text-sm">
              <span class="text-gray-600">{{ product.name }}</span>
              <span class="font-medium">{{ product.discountCount }} تخفیف</span>
            </div>
          </div>
        </div>
        
        <div class="bg-white border border-gray-200 rounded-lg p-6">
          <h4 class="font-medium text-gray-900 mb-4">کاربران فعال</h4>
          <div class="space-y-2">
            <div v-for="user in topUsers" :key="user.id" class="flex justify-between text-sm">
              <span class="text-gray-600">{{ user.name }}</span>
              <span class="font-medium">{{ user.couponCount }} کوپن</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'

// فیلترها
const filters = ref({
  period: 'month',
  startDate: '',
  endDate: '',
  reportType: 'overview'
})

// آمار کلی
const stats = ref({
  activeCoupons: 15,
  activeCampaigns: 8,
  totalDiscount: 125000000,
  totalOrders: 2847
})

// کوپن‌های برتر
const topCoupons = ref([
  { id: 1, name: 'WELCOME20', usageRate: 85, color: '#3B82F6' },
  { id: 2, name: 'FREESHIP', usageRate: 72, color: '#10B981' },
  { id: 3, name: 'SAVE50K', usageRate: 68, color: '#8B5CF6' },
  { id: 4, name: 'BIRTHDAY15', usageRate: 45, color: '#F59E0B' },
  { id: 5, name: 'FLASH30', usageRate: 38, color: '#EF4444' }
])

// کمپین‌های برتر
const topCampaigns = ref([
  { id: 1, name: 'کمپین تابستانه', performanceRate: 92, color: '#3B82F6' },
  { id: 2, name: 'تخفیف نوروزی', performanceRate: 88, color: '#10B981' },
  { id: 3, name: 'فلش سیل', performanceRate: 76, color: '#8B5CF6' },
  { id: 4, name: 'کمپین زمستانه', performanceRate: 65, color: '#F59E0B' }
])

// گزارش تفصیلی
const detailedReport = ref([
  {
    id: 1,
    name: 'کوپن خوشامدگویی',
    code: 'WELCOME20',
    type: 'percentage',
    usageCount: 1250,
    totalDiscount: 25000000,
    totalSales: 125000000,
    conversionRate: 15.5,
    roi: 320
  },
  {
    id: 2,
    name: 'ارسال رایگان',
    code: 'FREESHIP',
    type: 'free_shipping',
    usageCount: 890,
    totalDiscount: 15000000,
    totalSales: 89000000,
    conversionRate: 12.8,
    roi: 280
  },
  {
    id: 3,
    name: 'تخفیف ۵۰ هزار تومانی',
    code: 'SAVE50K',
    type: 'fixed',
    usageCount: 720,
    totalDiscount: 36000000,
    totalSales: 72000000,
    conversionRate: 10.2,
    roi: 250
  }
])

// بهترین ساعات استفاده
const peakHours = ref([
  { hour: 10, usage: 245 },
  { hour: 14, usage: 312 },
  { hour: 18, usage: 289 },
  { hour: 20, usage: 267 },
  { hour: 22, usage: 198 }
])

// محبوب‌ترین محصولات
const topProducts = ref([
  { id: 1, name: 'لپ‌تاپ گیمینگ', discountCount: 156 },
  { id: 2, name: 'گوشی هوشمند', discountCount: 134 },
  { id: 3, name: 'هدفون بی‌سیم', discountCount: 98 },
  { id: 4, name: 'ساعت هوشمند', discountCount: 87 },
  { id: 5, name: 'تبلت', discountCount: 76 }
])

// کاربران فعال
const topUsers = ref([
  { id: 1, name: 'علی احمدی', couponCount: 12 },
  { id: 2, name: 'فاطمه محمدی', couponCount: 10 },
  { id: 3, name: 'محمد رضایی', couponCount: 8 },
  { id: 4, name: 'زهرا کریمی', couponCount: 7 },
  { id: 5, name: 'حسین نوری', couponCount: 6 }
])

// توابع کمکی
const formatPrice = (price: number): string => {
  if (price >= 1000000000) {
    return (price / 1000000000).toFixed(1) + ' میلیارد'
  } else if (price >= 1000000) {
    return (price / 1000000).toFixed(1) + ' میلیون'
  } else if (price >= 1000) {
    return (price / 1000).toFixed(0) + ' هزار'
  }
  return new Intl.NumberFormat('fa-IR').format(price) + ' تومان'
}

const formatNumber = (num: number): string => {
  if (num >= 1000000) {
    return (num / 1000000).toFixed(1) + 'M'
  } else if (num >= 1000) {
    return (num / 1000).toFixed(0) + 'K'
  }
  return new Intl.NumberFormat('fa-IR').format(num)
}

const getTypeName = (type: string): string => {
  switch (type) {
    case 'percentage': return 'درصدی'
    case 'fixed': return 'مبلغ ثابت'
    case 'free_shipping': return 'ارسال رایگان'
    default: return type
  }
}

const getTypeBadgeClass = (type: string): string => {
  switch (type) {
    case 'percentage': return 'bg-blue-100 text-blue-700'
    case 'fixed': return 'bg-green-100 text-green-700'
    case 'free_shipping': return 'bg-purple-100 text-purple-700'
    default: return 'bg-gray-100 text-gray-700'
  }
}

// توابع عملیات
const applyFilters = () => {
  // TODO: اعمال فیلترها و بارگذاری مجدد داده‌ها
  console.log('اعمال فیلترها:', filters.value)
}

const exportReport = () => {
  // TODO: خروجی اکسل از گزارش
  console.log('خروجی اکسل گزارش')
}

const refreshData = () => {
  // TODO: بروزرسانی داده‌ها
  console.log('بروزرسانی داده‌ها')
}

// بارگذاری داده‌ها
onMounted(async () => {
  try {
    // TODO: فراخوانی API برای دریافت گزارش‌ها
    // const response = await $fetch('/api/admin/discounts/reports', {
    //   params: filters.value
    // })
    // stats.value = response.stats
    // detailedReport.value = response.detailedReport
  } catch (error) {
    console.error('خطا در بارگذاری گزارش‌ها:', error)
  }
})
</script>

<!--
  کامپوننت گزارش‌گیری و آمار
  شامل:
  1. آمار کلی کوپن‌ها و کمپین‌ها
  2. نمودارهای عملکرد
  3. گزارش‌های تفصیلی
  4. فیلترهای پیشرفته
  5. خروجی اکسل
  توضیحات کامل به فارسی در کد
--> 

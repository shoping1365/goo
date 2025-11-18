<template>
  <div class="space-y-6">
    <!-- آمار کلی سفارش -->
    <div class="px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-blue-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
        </svg>
        آمار کلی سفارش
      </h3>

      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gapx-4 py-4">
        <!-- کارت آمار -->
        <div class="bg-blue-50 px-4 py-4 rounded-lg">
          <div class="flex items-center">
            <div class="p-2 bg-blue-500 rounded-lg">
              <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
              </svg>
            </div>
            <div class="mr-3">
              <p class="text-sm text-gray-600">مبلغ کل</p>
              <p class="text-lg font-bold text-gray-900">{{ formatPrice(statisticsReports.totalAmount) }} تومان</p>
            </div>
          </div>
        </div>

        <div class="bg-green-50 px-4 py-4 rounded-lg">
          <div class="flex items-center">
            <div class="p-2 bg-green-500 rounded-lg">
              <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
            <div class="mr-3">
              <p class="text-sm text-gray-600">تخفیف کل</p>
              <p class="text-lg font-bold text-gray-900">{{ formatPrice(statisticsReports.totalDiscount) }} تومان</p>
            </div>
          </div>
        </div>

        <div class="bg-purple-50 px-4 py-4 rounded-lg">
          <div class="flex items-center">
            <div class="p-2 bg-purple-500 rounded-lg">
              <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z"></path>
              </svg>
            </div>
            <div class="mr-3">
              <p class="text-sm text-gray-600">تعداد محصولات</p>
              <p class="text-lg font-bold text-gray-900">{{ statisticsReports.productCount }} عدد</p>
            </div>
          </div>
        </div>

        <div class="bg-orange-50 px-4 py-4 rounded-lg">
          <div class="flex items-center">
            <div class="p-2 bg-orange-500 rounded-lg">
              <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
            <div class="mr-3">
              <p class="text-sm text-gray-600">زمان پردازش</p>
              <p class="text-lg font-bold text-gray-900">{{ statisticsReports.processingTime }} دقیقه</p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- آمار مشتری -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-green-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
        </svg>
        آمار مشتری
      </h3>

      <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
        <!-- اطلاعات مشتری -->
        <div>
          <h4 class="text-md font-medium text-gray-700 mb-3">اطلاعات مشتری</h4>
          <div class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">تعداد سفارشات قبلی</label>
              <input 
                v-model="statisticsReports.customerStats.previousOrders"
                type="number" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="0"
                readonly
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">مجموع خریدهای قبلی</label>
              <input 
                v-model="statisticsReports.customerStats.totalPreviousSpent"
                type="number" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="0"
                readonly
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">متوسط خرید</label>
              <input 
                v-model="statisticsReports.customerStats.averageOrderValue"
                type="number" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="0"
                readonly
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">آخرین خرید</label>
              <input 
                v-model="statisticsReports.customerStats.lastOrderDate"
                type="date" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                readonly
              />
            </div>
          </div>
        </div>

        <!-- رفتار خرید -->
        <div>
          <h4 class="text-md font-medium text-gray-700 mb-3">رفتار خرید</h4>
          <div class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">دسته‌بندی محبوب</label>
              <input 
                v-model="statisticsReports.customerStats.favoriteCategory"
                type="text" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="دسته‌بندی محبوب"
                readonly
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">محصولات محبوب</label>
              <textarea 
                v-model="statisticsReports.customerStats.favoriteProducts"
                rows="3"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="محصولات محبوب مشتری"
                readonly
              ></textarea>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">نرخ بازگشت</label>
              <input 
                v-model="statisticsReports.customerStats.returnRate"
                type="number" 
                step="0.01"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="0"
                readonly
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">امتیاز رضایت</label>
              <input 
                v-model="statisticsReports.customerStats.satisfactionScore"
                type="number" 
                min="1"
                max="5"
                step="0.1"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="0"
                readonly
              />
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- آمار محصولات -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-purple-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"></path>
        </svg>
        آمار محصولات
      </h3>

      <div class="space-y-4">
        <div 
          v-for="(product, index) in statisticsReports.productStats" 
          :key="index"
          class="border border-gray-200 rounded-lg px-4 py-4"
        >
          <div class="grid grid-cols-1 md:grid-cols-4 gapx-4 py-4">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-1">نام محصول</label>
              <input 
                v-model="product.name"
                type="text" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="نام محصول"
                readonly
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-1">تعداد فروش</label>
              <input 
                v-model="product.salesCount"
                type="number" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="0"
                readonly
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-1">درآمد کل</label>
              <input 
                v-model="product.totalRevenue"
                type="number" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="0"
                readonly
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-1">امتیاز</label>
              <input 
                v-model="product.rating"
                type="number" 
                min="1"
                max="5"
                step="0.1"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="0"
                readonly
              />
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- گزارشات تحلیلی -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-red-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 17v-2m3 2v-4m3 4v-6m2 10H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
        </svg>
        گزارشات تحلیلی
      </h3>

      <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
        <!-- تحلیل زمانی -->
        <div>
          <h4 class="text-md font-medium text-gray-700 mb-3">تحلیل زمانی</h4>
          <div class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">ساعت اوج سفارش</label>
              <input 
                v-model="statisticsReports.timeAnalysis.peakHour"
                type="text" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="14:00"
                readonly
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">روز اوج سفارش</label>
              <input 
                v-model="statisticsReports.timeAnalysis.peakDay"
                type="text" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="شنبه"
                readonly
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">فصل اوج سفارش</label>
              <input 
                v-model="statisticsReports.timeAnalysis.peakSeason"
                type="text" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="تابستان"
                readonly
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">نرخ رشد ماهانه</label>
              <input 
                v-model="statisticsReports.timeAnalysis.monthlyGrowth"
                type="number" 
                step="0.01"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="0"
                readonly
              />
            </div>
          </div>
        </div>

        <!-- تحلیل جغرافیایی -->
        <div>
          <h4 class="text-md font-medium text-gray-700 mb-3">تحلیل جغرافیایی</h4>
          <div class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">منطقه محبوب</label>
              <input 
                v-model="statisticsReports.geographicAnalysis.popularRegion"
                type="text" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="تهران"
                readonly
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">شهر محبوب</label>
              <input 
                v-model="statisticsReports.geographicAnalysis.popularCity"
                type="text" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="تهران"
                readonly
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">نرخ تحویل موفق</label>
              <input 
                v-model="statisticsReports.geographicAnalysis.deliverySuccessRate"
                type="number" 
                step="0.01"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="0"
                readonly
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">متوسط زمان تحویل</label>
              <input 
                v-model="statisticsReports.geographicAnalysis.averageDeliveryTime"
                type="number" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="0"
                readonly
              />
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- گزارشات مالی -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-orange-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 7h6m0 10v-3m-3 3h.01M9 17h.01M9 14h.01M12 14h.01M15 11h.01M12 11h.01M9 11h.01M7 21h10a2 2 0 002-2V5a2 2 0 00-2-2H7a2 2 0 00-2 2v14a2 2 0 002 2z"></path>
        </svg>
        گزارشات مالی
      </h3>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
        <!-- خلاصه مالی -->
        <div class="space-y-3">
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">درآمد ناخالص:</span>
            <span class="font-medium">{{ formatPrice(statisticsReports.financialReport.grossRevenue) }} تومان</span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">هزینه‌های عملیاتی:</span>
            <span class="font-medium text-red-600">{{ formatPrice(statisticsReports.financialReport.operationalCosts) }} تومان</span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">مالیات:</span>
            <span class="font-medium text-red-600">{{ formatPrice(statisticsReports.financialReport.taxes) }} تومان</span>
          </div>
          <div class="border-t pt-2 flex justify-between text-lg font-bold">
            <span>سود خالص:</span>
            <span class="text-green-600">{{ formatPrice(statisticsReports.financialReport.netProfit) }} تومان</span>
          </div>
        </div>

        <!-- شاخص‌های مالی -->
        <div class="space-y-3">
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">نرخ سود:</span>
            <span class="font-medium">{{ statisticsReports.financialReport.profitMargin }}%</span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">نرخ بازگشت سرمایه:</span>
            <span class="font-medium">{{ statisticsReports.financialReport.roi }}%</span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">نرخ تبدیل:</span>
            <span class="font-medium">{{ statisticsReports.financialReport.conversionRate }}%</span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">ارزش متوسط سفارش:</span>
            <span class="font-medium">{{ formatPrice(statisticsReports.financialReport.averageOrderValue) }} تومان</span>
          </div>
        </div>
      </div>
    </div>

    <!-- دکمه‌های عملیات -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-indigo-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"></path>
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
        </svg>
        عملیات گزارشات
      </h3>
      
      <div class="flex flex-wrap gapx-4 py-4">
        <button 
          @click="generateReport"
          class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors flex items-center text-sm"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
          </svg>
          تولید گزارش
        </button>

        <button 
          @click="exportData"
          class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors flex items-center text-sm"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
          </svg>
          خروجی اکسل
        </button>

        <button 
          @click="printReport"
          class="px-4 py-2 bg-purple-600 text-white rounded-lg hover:bg-purple-700 transition-colors flex items-center text-sm"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 17h2a2 2 0 002-2v-4a2 2 0 00-2-2H5a2 2 0 00-2 2v4a2 2 0 002 2h2m2 4h6a2 2 0 002-2v-4a2 2 0 00-2-2H9a2 2 0 00-2 2v4a2 2 0 002 2zm8-12V5a2 2 0 00-2-2H9a2 2 0 00-2 2v4h10z"></path>
          </svg>
          چاپ گزارش
        </button>

        <button 
          @click="refreshStats"
          class="px-4 py-2 bg-orange-600 text-white rounded-lg hover:bg-orange-700 transition-colors flex items-center text-sm"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
          </svg>
          به‌روزرسانی آمار
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'

// تعریف props
const props = defineProps({
  modelValue: {
    type: Object,
    default: () => ({
      totalAmount: 0,
      totalDiscount: 0,
      productCount: 0,
      processingTime: 0,
      customerStats: {
        previousOrders: 0,
        totalPreviousSpent: 0,
        averageOrderValue: 0,
        lastOrderDate: '',
        favoriteCategory: '',
        favoriteProducts: '',
        returnRate: 0,
        satisfactionScore: 0
      },
      productStats: [],
      timeAnalysis: {
        peakHour: '',
        peakDay: '',
        peakSeason: '',
        monthlyGrowth: 0
      },
      geographicAnalysis: {
        popularRegion: '',
        popularCity: '',
        deliverySuccessRate: 0,
        averageDeliveryTime: 0
      },
      financialReport: {
        grossRevenue: 0,
        operationalCosts: 0,
        taxes: 0,
        netProfit: 0,
        profitMargin: 0,
        roi: 0,
        conversionRate: 0,
        averageOrderValue: 0
      }
    })
  }
})

// تعریف emit
const emit = defineEmits(['update:modelValue'])

// متغیر محلی
const statisticsReports = ref({ ...props.modelValue })

// توابع کمکی
const formatPrice = (price) => {
  if (!price) return '0'
  return new Intl.NumberFormat('fa-IR').format(price)
}

const generateReport = () => {
  // تولید گزارش
  console.log('تولید گزارش...')
}

const exportData = () => {
  // خروجی اکسل
  console.log('خروجی اکسل...')
}

const printReport = () => {
  // چاپ گزارش
  console.log('چاپ گزارش...')
}

const refreshStats = () => {
  // به‌روزرسانی آمار
  console.log('به‌روزرسانی آمار...')
}

// نظارت بر تغییرات
watch(statisticsReports, (newValue) => {
  emit('update:modelValue', newValue)
}, { deep: true })
</script> 

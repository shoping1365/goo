<template>
  <div class="space-y-8">
    <!-- بخش گزارش‌های فروش -->
    <section>
      <div class="bg-gradient-to-r from-blue-50 to-indigo-50 rounded-lg p-6 mb-6">
        <h2 class="text-2xl font-bold text-gray-900 mb-2">📈 گزارش‌های فروش</h2>
        <p class="text-gray-600">تحلیل جامع فروش و درآمد گیفت کارت‌ها</p>
      </div>
      
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- آمار کلی فروش -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">آمار کلی فروش</h3>
            <span class="text-xs bg-green-100 text-green-800 rounded-full px-3 py-1">مثبت</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">فروش امروز:</span>
              <span class="text-lg font-bold text-green-600">{{ formatCurrency(salesStats.todaySales) }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">فروش این ماه:</span>
              <span class="text-lg font-bold text-blue-600">{{ formatCurrency(salesStats.monthlySales) }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">فروش این سال:</span>
              <span class="text-lg font-bold text-purple-600">{{ formatCurrency(salesStats.yearlySales) }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">رشد نسبت به ماه قبل:</span>
              <span class="text-lg font-bold text-green-600">+{{ salesStats.growthRate }}%</span>
            </div>
          </div>
          
          <!-- نمودار روند فروش -->
          <div class="mt-4">
            <div class="flex justify-between text-sm text-gray-500 mb-1">
              <span>روند فروش هفتگی</span>
              <span>+{{ salesStats.weeklyGrowth }}%</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div
class="bg-gradient-to-r from-green-500 to-blue-500 h-2 rounded-full transition-all duration-300"
                   :style="{ width: salesStats.weeklyGrowth + '%' }"></div>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-green-600 text-white text-sm rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500">
              تحلیل فروش
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              گزارش
            </button>
          </div>
        </div>

        <!-- فروش بر اساس دسته‌بندی -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">فروش بر اساس دسته‌بندی</h3>
            <span class="text-xs bg-blue-100 text-blue-800 rounded-full px-3 py-1">دسته‌بندی</span>
          </div>
          
          <div class="space-y-4">
            <div v-for="category in salesByCategory" :key="category.name" class="flex items-center justify-between">
              <div class="flex items-center">
                <span class="text-2xl ml-2">{{ category.icon }}</span>
                <span class="text-sm text-gray-600">{{ category.name }}</span>
              </div>
              <div class="flex items-center">
                <div class="w-24 bg-gray-200 rounded-full h-2 ml-2">
                  <div
:class="category.color + ' h-2 rounded-full transition-all duration-300'"
                       :style="{ width: category.percentage + '%' }"></div>
                </div>
                <span class="text-sm font-bold text-gray-900">{{ formatCurrency(category.sales) }}</span>
              </div>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
              تحلیل دسته‌بندی
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              نمودار
            </button>
          </div>
        </div>
      </div>
    </section>

    <!-- بخش تحلیل رفتار -->
    <section>
      <div class="bg-gradient-to-r from-green-50 to-emerald-50 rounded-lg p-6 mb-6">
        <h2 class="text-2xl font-bold text-gray-900 mb-2">🧠 تحلیل رفتار</h2>
        <p class="text-gray-600">تحلیل رفتار مشتریان و الگوهای خرید</p>
      </div>
      
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- الگوهای خرید -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">الگوهای خرید</h3>
            <span class="text-xs bg-blue-100 text-blue-800 rounded-full px-3 py-1">تحلیل</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">میانگین مبلغ خرید:</span>
              <span class="text-lg font-bold text-blue-600">{{ formatCurrency(behaviorStats.averagePurchase) }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">تکرار خرید:</span>
              <span class="text-lg font-bold text-green-600">{{ behaviorStats.repeatPurchaseRate }}%</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">فاصله خرید:</span>
              <span class="text-lg font-bold text-purple-600">{{ behaviorStats.purchaseInterval }} روز</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">نرخ تبدیل:</span>
              <span class="text-lg font-bold text-orange-600">{{ behaviorStats.conversionRate }}%</span>
            </div>
          </div>
          
          <!-- نمودار الگوهای خرید -->
          <div class="mt-4">
            <div class="flex justify-between text-sm text-gray-500 mb-1">
              <span>الگوی خرید ماهانه</span>
              <span>{{ behaviorStats.monthlyPattern }}%</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div
class="bg-gradient-to-r from-blue-500 to-green-500 h-2 rounded-full transition-all duration-300"
                   :style="{ width: behaviorStats.monthlyPattern + '%' }"></div>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
              تحلیل رفتار
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              گزارش
            </button>
          </div>
        </div>

        <!-- ساعات اوج خرید -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">ساعات اوج خرید</h3>
            <span class="text-xs bg-orange-100 text-orange-800 rounded-full px-3 py-1">زمان</span>
          </div>
          
          <div class="space-y-4">
            <div v-for="hour in peakHours" :key="hour.time" class="flex items-center justify-between">
              <span class="text-sm text-gray-600">{{ hour.time }}</span>
              <div class="flex items-center">
                <div class="w-24 bg-gray-200 rounded-full h-2 ml-2">
                  <div
class="bg-gradient-to-r from-orange-500 to-red-500 h-2 rounded-full transition-all duration-300"
                       :style="{ width: hour.percentage + '%' }"></div>
                </div>
                <span class="text-sm font-bold text-gray-900">{{ hour.orderCount }}</span>
              </div>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-orange-600 text-white text-sm rounded-lg hover:bg-orange-700 focus:outline-none focus:ring-2 focus:ring-orange-500">
              تحلیل زمانی
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              نمودار
            </button>
          </div>
        </div>
      </div>
    </section>

    <!-- بخش پیش‌بینی‌ها -->
    <section>
      <div class="bg-gradient-to-r from-purple-50 to-pink-50 rounded-lg p-6 mb-6">
        <h2 class="text-2xl font-bold text-gray-900 mb-2">🔮 پیش‌بینی‌ها</h2>
        <p class="text-gray-600">پیش‌بینی فروش و روند آینده</p>
      </div>
      
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- پیش‌بینی فروش -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">پیش‌بینی فروش</h3>
            <span class="text-xs bg-purple-100 text-purple-800 rounded-full px-3 py-1">AI</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">پیش‌بینی هفته آینده:</span>
              <span class="text-lg font-bold text-purple-600">{{ formatCurrency(forecastStats.nextWeekSales) }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">پیش‌بینی ماه آینده:</span>
              <span class="text-lg font-bold text-blue-600">{{ formatCurrency(forecastStats.nextMonthSales) }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">پیش‌بینی فصل آینده:</span>
              <span class="text-lg font-bold text-green-600">{{ formatCurrency(forecastStats.nextQuarterSales) }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">دقت پیش‌بینی:</span>
              <span class="text-lg font-bold text-orange-600">{{ forecastStats.accuracy }}%</span>
            </div>
          </div>
          
          <!-- نمودار پیش‌بینی -->
          <div class="mt-4">
            <div class="flex justify-between text-sm text-gray-500 mb-1">
              <span>دقت پیش‌بینی</span>
              <span>{{ forecastStats.accuracy }}%</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div
class="bg-gradient-to-r from-purple-500 to-pink-500 h-2 rounded-full transition-all duration-300"
                   :style="{ width: forecastStats.accuracy + '%' }"></div>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-purple-600 text-white text-sm rounded-lg hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-purple-500">
              تحلیل پیش‌بینی
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              گزارش
            </button>
          </div>
        </div>

        <!-- روندهای آینده -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">روندهای آینده</h3>
            <span class="text-xs bg-pink-100 text-pink-800 rounded-full px-3 py-1">روند</span>
          </div>
          
          <div class="space-y-4">
            <div v-for="trend in futureTrends" :key="trend.name" class="flex items-center justify-between">
              <div class="flex items-center">
                <span :class="trend.color + ' w-3 h-3 rounded-full ml-2'"></span>
                <span class="text-sm text-gray-600">{{ trend.name }}</span>
              </div>
              <div class="flex items-center">
                <span class="text-sm font-bold text-gray-900">{{ trend.percentage }}%</span>
                <span class="text-xs text-gray-500 mr-2">{{ trend.timeframe }}</span>
                <span v-if="trend.direction === 'up'" class="text-green-500">↗</span>
                <span v-else class="text-red-500">↘</span>
              </div>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-pink-600 text-white text-sm rounded-lg hover:bg-pink-700 focus:outline-none focus:ring-2 focus:ring-pink-500">
              تحلیل روند
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              نمودار
            </button>
          </div>
        </div>
      </div>
    </section>

    <!-- بخش مقایسه‌ها -->
    <section>
      <div class="bg-gradient-to-r from-orange-50 to-red-50 rounded-lg p-6 mb-6">
        <h2 class="text-2xl font-bold text-gray-900 mb-2">📊 مقایسه‌ها</h2>
        <p class="text-gray-600">مقایسه عملکرد و روندها</p>
      </div>
      
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- مقایسه دوره‌ها -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">مقایسه دوره‌ها</h3>
            <span class="text-xs bg-orange-100 text-orange-800 rounded-full px-3 py-1">مقایسه</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">ماه جاری:</span>
              <span class="text-lg font-bold text-blue-600">{{ formatCurrency(comparisonStats.currentMonth) }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">ماه گذشته:</span>
              <span class="text-lg font-bold text-gray-600">{{ formatCurrency(comparisonStats.lastMonth) }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">رشد:</span>
              <span class="text-lg font-bold text-green-600">+{{ comparisonStats.monthlyGrowth }}%</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">سال گذشته:</span>
              <span class="text-lg font-bold text-purple-600">{{ formatCurrency(comparisonStats.lastYear) }}</span>
            </div>
          </div>
          
          <!-- نمودار مقایسه -->
          <div class="mt-4">
            <div class="flex justify-between text-sm text-gray-500 mb-1">
              <span>مقایسه ماهانه</span>
              <span>+{{ comparisonStats.monthlyGrowth }}%</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div
class="bg-gradient-to-r from-orange-500 to-red-500 h-2 rounded-full transition-all duration-300"
                   :style="{ width: Math.min(comparisonStats.monthlyGrowth, 100) + '%' }"></div>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-orange-600 text-white text-sm rounded-lg hover:bg-orange-700 focus:outline-none focus:ring-2 focus:ring-orange-500">
              مقایسه کامل
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              نمودار
            </button>
          </div>
        </div>

        <!-- مقایسه دسته‌بندی‌ها -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">مقایسه دسته‌بندی‌ها</h3>
            <span class="text-xs bg-red-100 text-red-800 rounded-full px-3 py-1">دسته‌بندی</span>
          </div>
          
          <div class="space-y-4">
            <div v-for="category in categoryComparison" :key="category.name" class="flex items-center justify-between">
              <div class="flex items-center">
                <span class="text-2xl ml-2">{{ category.icon }}</span>
                <span class="text-sm text-gray-600">{{ category.name }}</span>
              </div>
              <div class="flex items-center">
                <span class="text-sm font-bold text-gray-900">{{ formatCurrency(category.sales) }}</span>
                <span class="text-xs text-green-600 mr-2">+{{ category.growth }}%</span>
              </div>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-red-600 text-white text-sm rounded-lg hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500">
              تحلیل دسته‌بندی
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              نمودار
            </button>
          </div>
        </div>
      </div>
    </section>

    <!-- بخش داشبورد تعاملی -->
    <section>
      <div class="bg-gradient-to-r from-indigo-50 to-purple-50 rounded-lg p-6 mb-6">
        <h2 class="text-2xl font-bold text-gray-900 mb-2">🎛️ داشبورد تعاملی</h2>
        <p class="text-gray-600">داشبورد تعاملی با نمودارها و فیلترهای پیشرفته</p>
      </div>
      
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <!-- آمار کلی -->
        <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
          <div class="text-center p-6 bg-blue-50 rounded-lg">
            <div class="text-2xl font-bold text-blue-600">{{ interactiveStats.totalRevenue }}</div>
            <div class="text-sm text-gray-600">کل درآمد</div>
          </div>
          
          <div class="text-center p-6 bg-green-50 rounded-lg">
            <div class="text-2xl font-bold text-green-600">{{ interactiveStats.totalOrders }}</div>
            <div class="text-sm text-gray-600">کل سفارش‌ها</div>
          </div>
          
          <div class="text-center p-6 bg-purple-50 rounded-lg">
            <div class="text-2xl font-bold text-purple-600">{{ interactiveStats.averageOrderValue }}</div>
            <div class="text-sm text-gray-600">میانگین سفارش</div>
          </div>
          
          <div class="text-center p-6 bg-orange-50 rounded-lg">
            <div class="text-2xl font-bold text-orange-600">{{ interactiveStats.customerSatisfaction }}%</div>
            <div class="text-sm text-gray-600">رضایت مشتری</div>
          </div>
        </div>
        
        <!-- نمودار روند -->
        <div class="mt-6">
          <h4 class="text-sm font-medium text-gray-900 mb-3">روند فروش (12 ماه گذشته)</h4>
          <div class="overflow-x-auto">
            <div class="flex items-end space-x-2 space-x-reverse h-20 min-w-[900px]">
              <div v-for="(month, index) in interactiveTrend" :key="index" class="flex flex-col items-center min-w-[40px] max-w-[60px]">
                <div
class="w-full bg-gray-200 rounded-t"
                     :style="{ height: (month.sales / 1000000 * 80) + 'px' }">
                  <div
class="w-full bg-gradient-to-t from-blue-500 to-purple-500 rounded-t transition-all duration-300"
                       :style="{ height: (month.sales / 1000000 * 80) + 'px' }"></div>
                </div>
                <span class="text-xs text-gray-500 mt-1">{{ month.label }}</span>
              </div>
            </div>
          </div>
        </div>
        
        <!-- فیلترهای تعاملی -->
        <div class="mt-6 grid grid-cols-1 md:grid-cols-3 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">دوره زمانی</label>
            <select class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
              <option>7 روز گذشته</option>
              <option>30 روز گذشته</option>
              <option>3 ماه گذشته</option>
              <option>6 ماه گذشته</option>
              <option>1 سال گذشته</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">دسته‌بندی</label>
            <select class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
              <option>همه دسته‌بندی‌ها</option>
              <option>تولد</option>
              <option>عروسی</option>
              <option>سال نو</option>
              <option>ولنتاین</option>
              <option>مادر</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع گزارش</label>
            <select class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
              <option>فروش</option>
              <option>درآمد</option>
              <option>سفارش‌ها</option>
              <option>مشتریان</option>
            </select>
          </div>
        </div>
        
        <!-- دکمه‌های عملیات -->
        <div class="mt-6 flex space-x-2 space-x-reverse">
          <button class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
            اعمال فیلتر
          </button>
          <button class="px-4 py-2 bg-green-600 text-white text-sm rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500">
            خروجی Excel
          </button>
          <button class="px-4 py-2 bg-purple-600 text-white text-sm rounded-lg hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-purple-500">
            خروجی PDF
          </button>
          <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
            بازنشانی
          </button>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
// آمار فروش
const salesStats = {
  todaySales: 45000000,
  monthlySales: 1250000000,
  yearlySales: 15800000000,
  growthRate: 12.5,
  weeklyGrowth: 8.3
}

// فروش بر اساس دسته‌بندی
const salesByCategory = [
  { name: 'تولد', icon: '🎂', sales: 450000000, percentage: 28, color: 'bg-pink-500' },
  { name: 'عروسی', icon: '💒', sales: 350000000, percentage: 22, color: 'bg-red-500' },
  { name: 'سال نو', icon: '🎊', sales: 280000000, percentage: 18, color: 'bg-green-500' },
  { name: 'ولنتاین', icon: '💝', sales: 230000000, percentage: 15, color: 'bg-purple-500' },
  { name: 'مادر', icon: '🌹', sales: 180000000, percentage: 12, color: 'bg-orange-500' }
]

// آمار رفتار
const behaviorStats = {
  averagePurchase: 450000,
  repeatPurchaseRate: 65,
  purchaseInterval: 45,
  conversionRate: 78,
  monthlyPattern: 85
}

// ساعات اوج خرید
const peakHours = [
  { time: '18:00-20:00', percentage: 25, orderCount: 3850 },
  { time: '20:00-22:00', percentage: 22, orderCount: 3390 },
  { time: '14:00-16:00', percentage: 18, orderCount: 2780 },
  { time: '16:00-18:00', percentage: 15, orderCount: 2320 },
  { time: '10:00-12:00', percentage: 12, orderCount: 1850 }
]

// آمار پیش‌بینی
const forecastStats = {
  nextWeekSales: 350000000,
  nextMonthSales: 1400000000,
  nextQuarterSales: 4200000000,
  accuracy: 92
}

// روندهای آینده
const futureTrends = [
  { name: 'فروش آنلاین', direction: 'up', percentage: 15, timeframe: '3 ماه', color: 'bg-green-500' },
  { name: 'گیفت کارت دیجیتال', direction: 'up', percentage: 25, timeframe: '6 ماه', color: 'bg-blue-500' },
  { name: 'مناسبت‌های خاص', direction: 'up', percentage: 8, timeframe: '1 سال', color: 'bg-purple-500' },
  { name: 'فروش فیزیکی', direction: 'down', percentage: 5, timeframe: '6 ماه', color: 'bg-red-500' }
]

// آمار مقایسه
const comparisonStats = {
  currentMonth: 1250000000,
  lastMonth: 1100000000,
  monthlyGrowth: 13.6,
  lastYear: 980000000
}

// مقایسه دسته‌بندی‌ها
const categoryComparison = [
  { name: 'تولد', icon: '🎂', sales: 450000000, growth: 15, color: 'bg-pink-500' },
  { name: 'عروسی', icon: '💒', sales: 350000000, growth: 22, color: 'bg-red-500' },
  { name: 'سال نو', icon: '🎊', sales: 280000000, growth: 8, color: 'bg-green-500' },
  { name: 'ولنتاین', icon: '💝', sales: 230000000, growth: 18, color: 'bg-purple-500' },
  { name: 'مادر', icon: '🌹', sales: 180000000, growth: 12, color: 'bg-orange-500' }
]

// آمار تعاملی
const interactiveStats = {
  totalRevenue: '15.8 میلیارد',
  totalOrders: 15420,
  averageOrderValue: '1.02 میلیون',
  customerSatisfaction: 94
}

// روند تعاملی
const interactiveTrend = [
  { label: 'بهمن', sales: 1200000000 },
  { label: 'اسفند', sales: 1350000000 },
  { label: 'فروردین', sales: 1180000000 },
  { label: 'اردیبهشت', sales: 1420000000 },
  { label: 'خرداد', sales: 1580000000 },
  { label: 'تیر', sales: 1250000000 },
  { label: 'مرداد', sales: 1320000000 },
  { label: 'شهریور', sales: 1450000000 },
  { label: 'مهر', sales: 1380000000 },
  { label: 'آبان', sales: 1520000000 },
  { label: 'آذر', sales: 1680000000 },
  { label: 'دی', sales: 1580000000 }
]

// تابع فرمت کردن ارز
const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}
</script>

<!--
  مستندسازی:
  این کامپوننت شامل پنج بخش اصلی است:
  1. گزارش‌های فروش: آمار کلی فروش، فروش بر اساس دسته‌بندی
  2. تحلیل رفتار: الگوهای خرید، ساعات اوج خرید
  3. پیش‌بینی‌ها: پیش‌بینی فروش، روندهای آینده
  4. مقایسه‌ها: مقایسه دوره‌ها، مقایسه دسته‌بندی‌ها
  5. داشبورد تعاملی: آمار کلی، نمودار روند، فیلترهای تعاملی
  
  تمام توضیحات به فارسی و با طراحی ریسپانسیو ارائه شده است.
  نمودارها و آمار دقیق برای نمایش وضعیت فروش و تحلیل‌ها استفاده شده‌اند.
--> 

<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="bg-gradient-to-r from-indigo-50 to-purple-50 rounded-lg p-6">
      <h2 class="text-2xl font-bold text-gray-900 mb-2">📊 داشبورد تحلیلی کلی</h2>
      <p class="text-gray-600">نمایش جامع آمار و تحلیل‌های کلیدی کیف پول</p>
    </div>

    <!-- کارت‌های آمار کلیدی -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <!-- کاربران فعال -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">کاربران فعال</h3>
          <span class="text-xs bg-green-100 text-green-800 rounded-full px-3 py-1">+12%</span>
        </div>
        <div class="text-3xl font-bold text-green-600 mb-2">{{ analyticsStats.activeUsers }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-green-500">+{{ analyticsStats.userGrowth }}%</span>
          <span class="mx-2">از ماه قبل</span>
        </div>
        <!-- نمودار کوچک -->
        <div class="mt-3">
          <div class="flex items-end space-x-1 space-x-reverse h-8">
            <div v-for="(day, index) in userTrend" :key="index" 
                 class="flex-1 bg-green-200 rounded-t"
                 :style="{ height: (day.users / Math.max(...userTrend.map(d => d.users))) * 32 + 'px' }">
            </div>
          </div>
        </div>
      </div>

      <!-- حجم تراکنش‌ها -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">حجم تراکنش‌ها</h3>
          <span class="text-xs bg-blue-100 text-blue-800 rounded-full px-3 py-1">+8%</span>
        </div>
        <div class="text-3xl font-bold text-blue-600 mb-2">{{ formatCurrency(analyticsStats.transactionVolume) }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-blue-500">+{{ analyticsStats.volumeGrowth }}%</span>
          <span class="mx-2">از ماه قبل</span>
        </div>
        <!-- نمودار کوچک -->
        <div class="mt-3">
          <div class="flex items-end space-x-1 space-x-reverse h-8">
            <div v-for="(day, index) in volumeTrend" :key="index" 
                 class="flex-1 bg-blue-200 rounded-t"
                 :style="{ height: (day.volume / Math.max(...volumeTrend.map(d => d.volume))) * 32 + 'px' }">
            </div>
          </div>
        </div>
      </div>

      <!-- میانگین تراکنش -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">میانگین تراکنش</h3>
          <span class="text-xs bg-purple-100 text-purple-800 rounded-full px-3 py-1">+5%</span>
        </div>
        <div class="text-3xl font-bold text-purple-600 mb-2">{{ formatCurrency(analyticsStats.averageTransaction) }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-purple-500">+{{ analyticsStats.avgGrowth }}%</span>
          <span class="mx-2">از ماه قبل</span>
        </div>
        <!-- نمودار کوچک -->
        <div class="mt-3">
          <div class="flex items-end space-x-1 space-x-reverse h-8">
            <div v-for="(day, index) in avgTrend" :key="index" 
                 class="flex-1 bg-purple-200 rounded-t"
                 :style="{ height: (day.avg / Math.max(...avgTrend.map(d => d.avg))) * 32 + 'px' }">
            </div>
          </div>
        </div>
      </div>

      <!-- نرخ موفقیت -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">نرخ موفقیت</h3>
          <span class="text-xs bg-orange-100 text-orange-800 rounded-full px-3 py-1">+2%</span>
        </div>
        <div class="text-3xl font-bold text-orange-600 mb-2">{{ analyticsStats.successRate }}%</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-orange-500">+{{ analyticsStats.successGrowth }}%</span>
          <span class="mx-2">از ماه قبل</span>
        </div>
        <!-- نمودار دایره‌ای -->
        <div class="mt-3">
          <div class="w-12 h-12 rounded-full border-4 border-gray-200 flex items-center justify-center">
            <div class="w-8 h-8 rounded-full bg-gradient-to-r from-orange-400 to-orange-600 flex items-center justify-center">
              <span class="text-xs font-bold text-white">{{ analyticsStats.successRate }}%</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- نمودارهای تحلیلی -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- روند رشد کیف پول -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <h3 class="text-xl font-semibold text-gray-900">روند رشد کیف پول</h3>
          <div class="flex space-x-2 space-x-reverse">
            <button class="px-3 py-1 text-sm bg-blue-100 text-blue-800 rounded-lg hover:bg-blue-200">7 روز</button>
            <button class="px-3 py-1 text-sm bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200">30 روز</button>
            <button class="px-3 py-1 text-sm bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200">90 روز</button>
          </div>
        </div>
        
        <div class="h-64 flex items-end space-x-2 space-x-reverse overflow-x-auto">
          <div v-for="(day, index) in walletGrowthTrend" :key="index" class="flex-shrink-0 flex flex-col items-center min-w-12">
            <div class="w-full bg-gray-200 rounded-t relative"
                 :style="{ height: getChartHeight(day.balance) + 'px' }">
              <div class="w-full bg-gradient-to-t from-green-500 to-emerald-500 rounded-t transition-all duration-300 absolute bottom-0"
                   :style="{ height: getChartHeight(day.balance) + 'px' }"></div>
            </div>
            <span class="text-xs text-gray-500 mt-1 text-center">{{ day.date }}</span>
            <span class="text-xs text-gray-400 mt-1 text-center">{{ formatCurrency(day.balance) }}</span>
          </div>
        </div>
      </div>

      <!-- مقایسه با دوره‌های قبل -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-xl font-semibold text-gray-900 mb-6">مقایسه با دوره‌های قبل</h3>
        
        <div class="space-y-4">
          <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <div>
              <div class="text-sm font-medium text-gray-900">این ماه</div>
              <div class="text-xs text-gray-500">درآمد کل</div>
            </div>
            <div class="text-right">
              <div class="text-lg font-bold text-green-600">{{ formatCurrency(comparisonStats.thisMonth) }}</div>
              <div class="text-xs text-green-500">+{{ comparisonStats.monthlyGrowth }}%</div>
            </div>
          </div>
          
          <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <div>
              <div class="text-sm font-medium text-gray-900">ماه قبل</div>
              <div class="text-xs text-gray-500">درآمد کل</div>
            </div>
            <div class="text-right">
              <div class="text-lg font-bold text-blue-600">{{ formatCurrency(comparisonStats.lastMonth) }}</div>
              <div class="text-xs text-blue-500">+{{ comparisonStats.lastMonthGrowth }}%</div>
            </div>
          </div>
          
          <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <div>
              <div class="text-sm font-medium text-gray-900">این سال</div>
              <div class="text-xs text-gray-500">درآمد کل</div>
            </div>
            <div class="text-right">
              <div class="text-lg font-bold text-purple-600">{{ formatCurrency(comparisonStats.thisYear) }}</div>
              <div class="text-xs text-purple-500">+{{ comparisonStats.yearlyGrowth }}%</div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تحلیل‌های پیشرفته -->
    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <!-- تحلیل رفتار کاربران -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">تحلیل رفتار کاربران</h3>
        
        <div class="space-y-3">
          <div class="flex items-center justify-between">
            <span class="text-sm text-gray-600">کاربران جدید</span>
            <span class="text-sm font-medium text-green-600">{{ behaviorStats.newUsers }}</span>
          </div>
          <div class="w-full bg-gray-200 rounded-full h-2">
            <div class="bg-green-500 h-2 rounded-full" :style="{ width: behaviorStats.newUsersPercentage + '%' }"></div>
          </div>
          
          <div class="flex items-center justify-between">
            <span class="text-sm text-gray-600">کاربران فعال</span>
            <span class="text-sm font-medium text-blue-600">{{ behaviorStats.activeUsers }}</span>
          </div>
          <div class="w-full bg-gray-200 rounded-full h-2">
            <div class="bg-blue-500 h-2 rounded-full" :style="{ width: behaviorStats.activeUsersPercentage + '%' }"></div>
          </div>
          
          <div class="flex items-center justify-between">
            <span class="text-sm text-gray-600">کاربران VIP</span>
            <span class="text-sm font-medium text-purple-600">{{ behaviorStats.vipUsers }}</span>
          </div>
          <div class="w-full bg-gray-200 rounded-full h-2">
            <div class="bg-purple-500 h-2 rounded-full" :style="{ width: behaviorStats.vipUsersPercentage + '%' }"></div>
          </div>
        </div>
      </div>

      <!-- تحلیل تراکنش‌ها -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">تحلیل تراکنش‌ها</h3>
        
        <div class="space-y-3">
          <div class="flex items-center justify-between">
            <span class="text-sm text-gray-600">تراکنش‌های کوچک</span>
            <span class="text-sm font-medium text-green-600">{{ transactionAnalysis.smallTransactions }}%</span>
          </div>
          <div class="w-full bg-gray-200 rounded-full h-2">
            <div class="bg-green-500 h-2 rounded-full" :style="{ width: transactionAnalysis.smallTransactions + '%' }"></div>
          </div>
          
          <div class="flex items-center justify-between">
            <span class="text-sm text-gray-600">تراکنش‌های متوسط</span>
            <span class="text-sm font-medium text-blue-600">{{ transactionAnalysis.mediumTransactions }}%</span>
          </div>
          <div class="w-full bg-gray-200 rounded-full h-2">
            <div class="bg-blue-500 h-2 rounded-full" :style="{ width: transactionAnalysis.mediumTransactions + '%' }"></div>
          </div>
          
          <div class="flex items-center justify-between">
            <span class="text-sm text-gray-600">تراکنش‌های بزرگ</span>
            <span class="text-sm font-medium text-purple-600">{{ transactionAnalysis.largeTransactions }}%</span>
          </div>
          <div class="w-full bg-gray-200 rounded-full h-2">
            <div class="bg-purple-500 h-2 rounded-full" :style="{ width: transactionAnalysis.largeTransactions + '%' }"></div>
          </div>
        </div>
      </div>

      <!-- پیش‌بینی روندها -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">پیش‌بینی روندها</h3>
        
        <div class="space-y-4">
          <div class="p-3 bg-blue-50 rounded-lg">
            <div class="text-sm font-medium text-blue-900">پیش‌بینی ماه آینده</div>
            <div class="text-lg font-bold text-blue-600">{{ formatCurrency(predictionStats.nextMonthPrediction) }}</div>
            <div class="text-xs text-blue-500">+{{ predictionStats.nextMonthGrowth }}% رشد پیش‌بینی شده</div>
          </div>
          
          <div class="p-3 bg-green-50 rounded-lg">
            <div class="text-sm font-medium text-green-900">هدف ماهانه</div>
            <div class="text-lg font-bold text-green-600">{{ formatCurrency(predictionStats.monthlyTarget) }}</div>
            <div class="text-xs text-green-500">{{ predictionStats.targetProgress }}% پیشرفت</div>
          </div>
          
          <div class="p-3 bg-purple-50 rounded-lg">
            <div class="text-sm font-medium text-purple-900">پیش‌بینی سالانه</div>
            <div class="text-lg font-bold text-purple-600">{{ formatCurrency(predictionStats.yearlyPrediction) }}</div>
            <div class="text-xs text-purple-500">+{{ predictionStats.yearlyGrowth }}% رشد پیش‌بینی شده</div>
          </div>
        </div>
      </div>
    </div>

    <!-- شاخص‌های کلیدی عملکرد -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h3 class="text-xl font-semibold text-gray-900 mb-6">📈 شاخص‌های کلیدی عملکرد (KPI)</h3>
      
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <div class="text-center p-6 bg-gradient-to-br from-green-50 to-emerald-50 rounded-lg">
          <div class="text-2xl font-bold text-green-600">{{ kpiStats.customerSatisfaction }}%</div>
          <div class="text-sm text-gray-600">رضایت مشتری</div>
          <div class="text-xs text-green-500 mt-1">+2.5% از ماه قبل</div>
        </div>
        
        <div class="text-center p-6 bg-gradient-to-br from-blue-50 to-indigo-50 rounded-lg">
          <div class="text-2xl font-bold text-blue-600">{{ kpiStats.transactionSpeed }}s</div>
          <div class="text-sm text-gray-600">سرعت تراکنش</div>
          <div class="text-xs text-blue-500 mt-1">-0.3s از ماه قبل</div>
        </div>
        
        <div class="text-center p-6 bg-gradient-to-br from-purple-50 to-pink-50 rounded-lg">
          <div class="text-2xl font-bold text-purple-600">{{ kpiStats.uptime }}%</div>
          <div class="text-sm text-gray-600">در دسترس بودن</div>
          <div class="text-xs text-purple-500 mt-1">+0.1% از ماه قبل</div>
        </div>
        
        <div class="text-center p-6 bg-gradient-to-br from-orange-50 to-red-50 rounded-lg">
          <div class="text-2xl font-bold text-orange-600">{{ kpiStats.errorRate }}%</div>
          <div class="text-sm text-gray-600">نرخ خطا</div>
          <div class="text-xs text-orange-500 mt-1">-0.2% از ماه قبل</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// آمار تحلیلی
const analyticsStats = {
  activeUsers: 15420,
  userGrowth: 12.5,
  transactionVolume: 125000000000,
  volumeGrowth: 8.3,
  averageTransaction: 850000,
  avgGrowth: 5.2,
  successRate: 96.8,
  successGrowth: 2.1
}

// روند کاربران (7 روز)
const userTrend = [
  { users: 14500, date: 'شنبه' },
  { users: 14800, date: 'یکشنبه' },
  { users: 15200, date: 'دوشنبه' },
  { users: 14900, date: 'سه‌شنبه' },
  { users: 15600, date: 'چهارشنبه' },
  { users: 15800, date: 'پنج‌شنبه' },
  { users: 15420, date: 'جمعه' }
]

// روند حجم تراکنش‌ها (7 روز)
const volumeTrend = [
  { volume: 115000000000, date: 'شنبه' },
  { volume: 118000000000, date: 'یکشنبه' },
  { volume: 122000000000, date: 'دوشنبه' },
  { volume: 119000000000, date: 'سه‌شنبه' },
  { volume: 128000000000, date: 'چهارشنبه' },
  { volume: 131000000000, date: 'پنج‌شنبه' },
  { volume: 125000000000, date: 'جمعه' }
]

// روند میانگین تراکنش (7 روز)
const avgTrend = [
  { avg: 820000, date: 'شنبه' },
  { avg: 835000, date: 'یکشنبه' },
  { avg: 850000, date: 'دوشنبه' },
  { avg: 845000, date: 'سه‌شنبه' },
  { avg: 860000, date: 'چهارشنبه' },
  { avg: 870000, date: 'پنج‌شنبه' },
  { avg: 850000, date: 'جمعه' }
]

// روند رشد کیف پول (30 روز)
const walletGrowthTrend = [
  { balance: 115000000000, date: '1' },
  { balance: 118000000000, date: '5' },
  { balance: 122000000000, date: '10' },
  { balance: 125000000000, date: '15' },
  { balance: 128000000000, date: '20' },
  { balance: 131000000000, date: '25' },
  { balance: 125000000000, date: '30' }
]

// مقایسه با دوره‌های قبل
const comparisonStats = {
  thisMonth: 125000000000,
  monthlyGrowth: 8.3,
  lastMonth: 115000000000,
  lastMonthGrowth: 6.2,
  thisYear: 1580000000000,
  yearlyGrowth: 15.8
}

// تحلیل رفتار کاربران
const behaviorStats = {
  newUsers: 1250,
  newUsersPercentage: 8.1,
  activeUsers: 12450,
  activeUsersPercentage: 80.7,
  vipUsers: 1720,
  vipUsersPercentage: 11.2
}

// تحلیل تراکنش‌ها
const transactionAnalysis = {
  smallTransactions: 45,
  mediumTransactions: 35,
  largeTransactions: 20
}

// پیش‌بینی روندها
const predictionStats = {
  nextMonthPrediction: 135000000000,
  nextMonthGrowth: 8.0,
  monthlyTarget: 140000000000,
  targetProgress: 89.3,
  yearlyPrediction: 1850000000000,
  yearlyGrowth: 17.1
}

// شاخص‌های کلیدی عملکرد
const kpiStats = {
  customerSatisfaction: 94.5,
  transactionSpeed: 2.3,
  uptime: 99.8,
  errorRate: 0.15
}

// تابع فرمت کردن ارز
const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

// تابع محاسبه ارتفاع مناسب برای نمودار
const getChartHeight = (balance: number) => {
  const maxBalance = Math.max(...walletGrowthTrend.map(item => item.balance));
  const minBalance = Math.min(...walletGrowthTrend.map(item => item.balance));
  const range = maxBalance - minBalance;
  const height = 200; // حداکثر ارتفاع نمودار

  if (range === 0) return height;

  const percentage = ((balance - minBalance) / range) * 100;
  return (percentage / 100) * height;
}
</script>

<!--
  مستندسازی:
  این کامپوننت شامل داشبورد تحلیلی کلی کیف پول است که شامل:
  1. آمار کلیدی: کاربران فعال، حجم تراکنش‌ها، میانگین تراکنش، نرخ موفقیت
  2. نمودارهای روند: رشد کیف پول، مقایسه با دوره‌های قبل
  3. تحلیل‌های پیشرفته: رفتار کاربران، تحلیل تراکنش‌ها، پیش‌بینی روندها
  4. شاخص‌های کلیدی عملکرد (KPI)
  
  تمام آمار و نمودارها به صورت ریسپانسیو و با طراحی مدرن ارائه شده‌اند.
--> 

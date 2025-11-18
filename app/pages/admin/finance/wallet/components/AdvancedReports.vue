<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="bg-gradient-to-r from-teal-50 to-cyan-50 rounded-lg p-6">
      <h2 class="text-2xl font-bold text-gray-900 mb-2">📋 گزارش‌های پیشرفته</h2>
      <p class="text-gray-600">گزارش‌های تحلیلی پیشرفته بر اساس معیارهای مختلف</p>
    </div>

    <!-- فیلترهای پیشرفته -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">فیلترهای پیشرفته</h3>
      
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">بازه زمانی</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option value="today">امروز</option>
            <option value="week">هفته جاری</option>
            <option value="month" selected>ماه جاری</option>
            <option value="quarter">سه ماهه</option>
            <option value="year">سال جاری</option>
            <option value="custom">سفارشی</option>
          </select>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">گروه کاربری</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option value="all">همه گروه‌ها</option>
            <option value="regular">کاربران عادی</option>
            <option value="premium">کاربران پریمیوم</option>
            <option value="vip">کاربران VIP</option>
            <option value="enterprise">کاربران سازمانی</option>
          </select>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">منطقه جغرافیایی</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option value="all">همه مناطق</option>
            <option value="tehran">تهران</option>
            <option value="isfahan">اصفهان</option>
            <option value="shiraz">شیراز</option>
            <option value="mashhad">مشهد</option>
            <option value="other">سایر شهرها</option>
          </select>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نوع دستگاه</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option value="all">همه دستگاه‌ها</option>
            <option value="mobile">موبایل</option>
            <option value="desktop">دسکتاپ</option>
            <option value="tablet">تبلت</option>
          </select>
        </div>
      </div>
      
      <div class="mt-4 flex space-x-2 space-x-reverse">
        <button class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
          اعمال فیلتر
        </button>
        <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
          پاک کردن
        </button>
        <button class="px-4 py-2 bg-green-600 text-white text-sm rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500">
          خروجی اکسل
        </button>
      </div>
    </div>

    <!-- گزارش بر اساس گروه‌های کاربری -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">گزارش بر اساس گروه‌های کاربری</h3>
      
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- نمودار دایره‌ای -->
        <div>
          <div class="flex items-center justify-center h-64">
            <div class="relative w-48 h-48">
              <!-- کاربران عادی -->
              <div class="absolute inset-0 rounded-full border-8 border-green-500 transform rotate-0"></div>
              <!-- کاربران پریمیوم -->
              <div class="absolute inset-0 rounded-full border-8 border-blue-500 transform rotate-90"></div>
              <!-- کاربران VIP -->
              <div class="absolute inset-0 rounded-full border-8 border-purple-500 transform rotate-180"></div>
              <!-- کاربران سازمانی -->
              <div class="absolute inset-0 rounded-full border-8 border-orange-500 transform rotate-270"></div>
              
              <div class="absolute inset-0 flex items-center justify-center">
                <div class="text-center">
                  <div class="text-2xl font-bold text-gray-900">{{ userGroupStats.totalUsers }}</div>
                  <div class="text-sm text-gray-600">کل کاربران</div>
                </div>
              </div>
            </div>
          </div>
        </div>
        
        <!-- آمار تفصیلی -->
        <div class="space-y-4">
          <div class="flex items-center justify-between p-3 bg-green-50 rounded-lg">
            <div class="flex items-center">
              <div class="w-4 h-4 bg-green-500 rounded-full ml-3"></div>
              <div>
                <div class="text-sm font-medium text-gray-900">کاربران عادی</div>
                <div class="text-xs text-gray-500">{{ userGroupStats.regularUsers }} کاربر</div>
              </div>
            </div>
            <div class="text-right">
              <div class="text-lg font-bold text-green-600">{{ userGroupStats.regularPercentage }}%</div>
              <div class="text-xs text-green-500">{{ formatCurrency(userGroupStats.regularRevenue) }}</div>
            </div>
          </div>
          
          <div class="flex items-center justify-between p-3 bg-blue-50 rounded-lg">
            <div class="flex items-center">
              <div class="w-4 h-4 bg-blue-500 rounded-full ml-3"></div>
              <div>
                <div class="text-sm font-medium text-gray-900">کاربران پریمیوم</div>
                <div class="text-xs text-gray-500">{{ userGroupStats.premiumUsers }} کاربر</div>
              </div>
            </div>
            <div class="text-right">
              <div class="text-lg font-bold text-blue-600">{{ userGroupStats.premiumPercentage }}%</div>
              <div class="text-xs text-blue-500">{{ formatCurrency(userGroupStats.premiumRevenue) }}</div>
            </div>
          </div>
          
          <div class="flex items-center justify-between p-3 bg-purple-50 rounded-lg">
            <div class="flex items-center">
              <div class="w-4 h-4 bg-purple-500 rounded-full ml-3"></div>
              <div>
                <div class="text-sm font-medium text-gray-900">کاربران VIP</div>
                <div class="text-xs text-gray-500">{{ userGroupStats.vipUsers }} کاربر</div>
              </div>
            </div>
            <div class="text-right">
              <div class="text-lg font-bold text-purple-600">{{ userGroupStats.vipPercentage }}%</div>
              <div class="text-xs text-purple-500">{{ formatCurrency(userGroupStats.vipRevenue) }}</div>
            </div>
          </div>
          
          <div class="flex items-center justify-between p-3 bg-orange-50 rounded-lg">
            <div class="flex items-center">
              <div class="w-4 h-4 bg-orange-500 rounded-full ml-3"></div>
              <div>
                <div class="text-sm font-medium text-gray-900">کاربران سازمانی</div>
                <div class="text-xs text-gray-500">{{ userGroupStats.enterpriseUsers }} کاربر</div>
              </div>
            </div>
            <div class="text-right">
              <div class="text-lg font-bold text-orange-600">{{ userGroupStats.enterprisePercentage }}%</div>
              <div class="text-xs text-orange-500">{{ formatCurrency(userGroupStats.enterpriseRevenue) }}</div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- گزارش بر اساس مناطق جغرافیایی -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">گزارش بر اساس مناطق جغرافیایی</h3>
      
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- نمودار میله‌ای -->
        <div>
          <div class="h-64 flex items-end space-x-4 space-x-reverse">
            <div v-for="(region, index) in geographicStats" :key="index" class="flex flex-col items-center flex-1">
              <div class="w-full bg-gray-200 rounded-t relative"
                   :style="{ height: getChartHeight(region.revenue) + 'px' }">
                <div class="w-full bg-gradient-to-t from-blue-500 to-indigo-500 rounded-t transition-all duration-300 absolute bottom-0"
                     :style="{ height: getChartHeight(region.revenue) + 'px' }"></div>
              </div>
              <span class="text-xs text-gray-500 mt-2 text-center">{{ region.name }}</span>
              <span class="text-xs text-gray-400 mt-1 text-center">{{ formatCurrency(region.revenue) }}</span>
            </div>
          </div>
        </div>
        
        <!-- جدول آمار -->
        <div>
          <div class="overflow-x-auto">
            <table class="w-full text-sm text-right">
              <thead class="bg-gray-50">
                <tr>
                  <th class="px-4 py-3 text-gray-700 font-medium">منطقه</th>
                  <th class="px-4 py-3 text-gray-700 font-medium">کاربران</th>
                  <th class="px-4 py-3 text-gray-700 font-medium">درآمد</th>
                  <th class="px-4 py-3 text-gray-700 font-medium">میانگین</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-gray-200">
                <tr v-for="region in geographicStats" :key="region.name" class="hover:bg-gray-50">
                  <td class="px-4 py-3 text-gray-900">{{ region.name }}</td>
                  <td class="px-4 py-3 text-gray-600">{{ region.users }}</td>
                  <td class="px-4 py-3 font-medium text-blue-600">{{ formatCurrency(region.revenue) }}</td>
                  <td class="px-4 py-3 text-gray-600">{{ formatCurrency(region.averageRevenue) }}</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>

    <!-- گزارش بر اساس دستگاه‌ها -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">گزارش بر اساس دستگاه‌ها</h3>
      
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <!-- موبایل -->
        <div class="text-center p-6 bg-green-50 rounded-lg">
          <div class="text-4xl mb-4">📱</div>
          <div class="text-2xl font-bold text-green-600">{{ deviceStats.mobile.users }}</div>
          <div class="text-sm text-gray-600 mb-2">کاربران موبایل</div>
          <div class="text-lg font-bold text-green-600">{{ deviceStats.mobile.percentage }}%</div>
          <div class="text-xs text-gray-500">{{ formatCurrency(deviceStats.mobile.revenue) }} درآمد</div>
        </div>
        
        <!-- دسکتاپ -->
        <div class="text-center p-6 bg-blue-50 rounded-lg">
          <div class="text-4xl mb-4">💻</div>
          <div class="text-2xl font-bold text-blue-600">{{ deviceStats.desktop.users }}</div>
          <div class="text-sm text-gray-600 mb-2">کاربران دسکتاپ</div>
          <div class="text-lg font-bold text-blue-600">{{ deviceStats.desktop.percentage }}%</div>
          <div class="text-xs text-gray-500">{{ formatCurrency(deviceStats.desktop.revenue) }} درآمد</div>
        </div>
        
        <!-- تبلت -->
        <div class="text-center p-6 bg-purple-50 rounded-lg">
          <div class="text-4xl mb-4">📱</div>
          <div class="text-2xl font-bold text-purple-600">{{ deviceStats.tablet.users }}</div>
          <div class="text-sm text-gray-600 mb-2">کاربران تبلت</div>
          <div class="text-lg font-bold text-purple-600">{{ deviceStats.tablet.percentage }}%</div>
          <div class="text-xs text-gray-500">{{ formatCurrency(deviceStats.tablet.revenue) }} درآمد</div>
        </div>
      </div>
    </div>

    <!-- تحلیل رفتار کاربران -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">تحلیل رفتار کاربران</h3>
      
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- الگوهای استفاده -->
        <div>
          <h4 class="text-md font-medium text-gray-900 mb-3">الگوهای استفاده</h4>
          <div class="space-y-3">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">استفاده صبحگاهی (6-12)</span>
              <span class="text-sm font-medium text-green-600">{{ behaviorAnalysis.morningUsage }}%</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div class="bg-green-500 h-2 rounded-full" :style="{ width: behaviorAnalysis.morningUsage + '%' }"></div>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">استفاده عصرگاهی (12-18)</span>
              <span class="text-sm font-medium text-blue-600">{{ behaviorAnalysis.afternoonUsage }}%</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div class="bg-blue-500 h-2 rounded-full" :style="{ width: behaviorAnalysis.afternoonUsage + '%' }"></div>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">استفاده شبانه (18-24)</span>
              <span class="text-sm font-medium text-purple-600">{{ behaviorAnalysis.eveningUsage }}%</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div class="bg-purple-500 h-2 rounded-full" :style="{ width: behaviorAnalysis.eveningUsage + '%' }"></div>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">استفاده شبانه (0-6)</span>
              <span class="text-sm font-medium text-orange-600">{{ behaviorAnalysis.nightUsage }}%</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div class="bg-orange-500 h-2 rounded-full" :style="{ width: behaviorAnalysis.nightUsage + '%' }"></div>
            </div>
          </div>
        </div>
        
        <!-- الگوهای تراکنش -->
        <div>
          <h4 class="text-md font-medium text-gray-900 mb-3">الگوهای تراکنش</h4>
          <div class="space-y-3">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">تراکنش‌های روزانه</span>
              <span class="text-sm font-medium text-green-600">{{ behaviorAnalysis.dailyTransactions }}%</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div class="bg-green-500 h-2 rounded-full" :style="{ width: behaviorAnalysis.dailyTransactions + '%' }"></div>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">تراکنش‌های هفتگی</span>
              <span class="text-sm font-medium text-blue-600">{{ behaviorAnalysis.weeklyTransactions }}%</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div class="bg-blue-500 h-2 rounded-full" :style="{ width: behaviorAnalysis.weeklyTransactions + '%' }"></div>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">تراکنش‌های ماهانه</span>
              <span class="text-sm font-medium text-purple-600">{{ behaviorAnalysis.monthlyTransactions }}%</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div class="bg-purple-500 h-2 rounded-full" :style="{ width: behaviorAnalysis.monthlyTransactions + '%' }"></div>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">تراکنش‌های فصلی</span>
              <span class="text-sm font-medium text-orange-600">{{ behaviorAnalysis.seasonalTransactions }}%</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div class="bg-orange-500 h-2 rounded-full" :style="{ width: behaviorAnalysis.seasonalTransactions + '%' }"></div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- پیش‌بینی روندها -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">پیش‌بینی روندها</h3>
      
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <!-- پیش‌بینی کاربران -->
        <div class="p-6 bg-gradient-to-br from-green-50 to-emerald-50 rounded-lg">
          <h4 class="text-md font-medium text-green-900 mb-3">پیش‌بینی کاربران</h4>
          <div class="space-y-2">
            <div class="flex justify-between">
              <span class="text-sm text-green-700">ماه آینده:</span>
              <span class="text-sm font-bold text-green-900">{{ predictionStats.users.nextMonth }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-sm text-green-700">سه ماه آینده:</span>
              <span class="text-sm font-bold text-green-900">{{ predictionStats.users.nextQuarter }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-sm text-green-700">سال آینده:</span>
              <span class="text-sm font-bold text-green-900">{{ predictionStats.users.nextYear }}</span>
            </div>
          </div>
        </div>
        
        <!-- پیش‌بینی درآمد -->
        <div class="p-6 bg-gradient-to-br from-blue-50 to-indigo-50 rounded-lg">
          <h4 class="text-md font-medium text-blue-900 mb-3">پیش‌بینی درآمد</h4>
          <div class="space-y-2">
            <div class="flex justify-between">
              <span class="text-sm text-blue-700">ماه آینده:</span>
              <span class="text-sm font-bold text-blue-900">{{ formatCurrency(predictionStats.revenue.nextMonth) }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-sm text-blue-700">سه ماه آینده:</span>
              <span class="text-sm font-bold text-blue-900">{{ formatCurrency(predictionStats.revenue.nextQuarter) }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-sm text-blue-700">سال آینده:</span>
              <span class="text-sm font-bold text-blue-900">{{ formatCurrency(predictionStats.revenue.nextYear) }}</span>
            </div>
          </div>
        </div>
        
        <!-- پیش‌بینی تراکنش‌ها -->
        <div class="p-6 bg-gradient-to-br from-purple-50 to-pink-50 rounded-lg">
          <h4 class="text-md font-medium text-purple-900 mb-3">پیش‌بینی تراکنش‌ها</h4>
          <div class="space-y-2">
            <div class="flex justify-between">
              <span class="text-sm text-purple-700">ماه آینده:</span>
              <span class="text-sm font-bold text-purple-900">{{ predictionStats.transactions.nextMonth }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-sm text-purple-700">سه ماه آینده:</span>
              <span class="text-sm font-bold text-purple-900">{{ predictionStats.transactions.nextQuarter }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-sm text-purple-700">سال آینده:</span>
              <span class="text-sm font-bold text-purple-900">{{ predictionStats.transactions.nextYear }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// آمار گروه‌های کاربری
const userGroupStats = {
  totalUsers: 15420,
  regularUsers: 9250,
  regularPercentage: 60,
  regularRevenue: 45000000000,
  premiumUsers: 4620,
  premiumPercentage: 30,
  premiumRevenue: 55000000000,
  vipUsers: 1230,
  vipPercentage: 8,
  vipRevenue: 20000000000,
  enterpriseUsers: 320,
  enterprisePercentage: 2,
  enterpriseRevenue: 5000000000
}

// آمار جغرافیایی
const geographicStats = [
  { name: 'تهران', users: 6200, revenue: 65000000000, averageRevenue: 10483870 },
  { name: 'اصفهان', users: 3100, revenue: 28000000000, averageRevenue: 9032258 },
  { name: 'شیراز', users: 2100, revenue: 18000000000, averageRevenue: 8571428 },
  { name: 'مشهد', users: 2800, revenue: 25000000000, averageRevenue: 8928571 },
  { name: 'سایر شهرها', users: 1220, revenue: 9000000000, averageRevenue: 7377049 }
]

// آمار دستگاه‌ها
const deviceStats = {
  mobile: { users: 9250, percentage: 60, revenue: 75000000000 },
  desktop: { users: 4620, percentage: 30, revenue: 45000000000 },
  tablet: { users: 1550, percentage: 10, revenue: 5000000000 }
}

// تحلیل رفتار کاربران
const behaviorAnalysis = {
  morningUsage: 25,
  afternoonUsage: 35,
  eveningUsage: 30,
  nightUsage: 10,
  dailyTransactions: 45,
  weeklyTransactions: 30,
  monthlyTransactions: 20,
  seasonalTransactions: 5
}

// پیش‌بینی روندها
const predictionStats = {
  users: {
    nextMonth: 16800,
    nextQuarter: 18500,
    nextYear: 22000
  },
  revenue: {
    nextMonth: 135000000000,
    nextQuarter: 420000000000,
    nextYear: 1850000000000
  },
  transactions: {
    nextMonth: 16800,
    nextQuarter: 18500,
    nextYear: 22000
  }
}

// تابع فرمت کردن ارز
const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

// تابع محاسبه ارتفاع مناسب برای نمودار
const getChartHeight = (revenue: number) => {
  const maxRevenue = Math.max(...geographicStats.map(item => item.revenue));
  const minRevenue = Math.min(...geographicStats.map(item => item.revenue));
  const range = maxRevenue - minRevenue;
  const height = 200; // حداکثر ارتفاع نمودار

  if (range === 0) return height;

  const percentage = ((revenue - minRevenue) / range) * 100;
  return (percentage / 100) * height;
}
</script>

<!--
  مستندسازی:
  این کامپوننت شامل گزارش‌های پیشرفته کیف پول است که شامل:
  1. فیلترهای پیشرفته: بازه زمانی، گروه کاربری، منطقه جغرافیایی، نوع دستگاه
  2. گزارش گروه‌های کاربری: نمودار دایره‌ای و آمار تفصیلی
  3. گزارش مناطق جغرافیایی: نمودار میله‌ای و جدول آمار
  4. گزارش دستگاه‌ها: آمار موبایل، دسکتاپ و تبلت
  5. تحلیل رفتار کاربران: الگوهای استفاده و تراکنش
  6. پیش‌بینی روندها: کاربران، درآمد و تراکنش‌ها
  
  تمام گزارش‌ها به صورت ریسپانسیو و با طراحی مدرن ارائه شده‌اند.
--> 

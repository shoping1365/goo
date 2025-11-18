<template>
  <div class="space-y-8">
    <!-- بخش تاریخچه خرید -->
    <section>
      <div class="bg-gradient-to-r from-blue-50 to-indigo-50 rounded-lg p-6 mb-6">
        <h2 class="text-2xl font-bold text-gray-900 mb-2">🛒 تاریخچه خرید</h2>
        <p class="text-gray-600">مدیریت کامل تاریخچه خریدهای گیفت کارت مشتریان</p>
      </div>
      
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- آمار خریدهای مشتریان -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">آمار خریدهای مشتریان</h3>
            <span class="text-xs bg-green-100 text-green-800 rounded-full px-3 py-1">فعال</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">کل مشتریان:</span>
              <span class="text-lg font-bold text-blue-600">{{ customerStats.totalCustomers }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">مشتریان فعال:</span>
              <span class="text-lg font-bold text-green-600">{{ customerStats.activeCustomers }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">خریدهای امروز:</span>
              <span class="text-lg font-bold text-purple-600">{{ customerStats.todayPurchases }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">میانگین خرید:</span>
              <span class="text-lg font-bold text-orange-600">{{ formatCurrency(customerStats.averagePurchase) }}</span>
            </div>
          </div>
          
          <!-- نمودار روند خرید -->
          <div class="mt-4">
            <div class="flex justify-between text-sm text-gray-500 mb-1">
              <span>روند خرید ماهانه</span>
              <span>+{{ customerStats.purchaseGrowth }}%</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div class="bg-gradient-to-r from-blue-500 to-purple-500 h-2 rounded-full transition-all duration-300"
                   :style="{ width: customerStats.purchaseGrowth + '%' }"></div>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
              مشاهده آمار
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              گزارش
            </button>
          </div>
        </div>

        <!-- مشتریان برتر -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">مشتریان برتر</h3>
            <span class="text-xs bg-yellow-100 text-yellow-800 rounded-full px-3 py-1">ویژه</span>
          </div>
          
          <div class="space-y-3">
            <div v-for="(customer, index) in topCustomers" :key="customer.id" 
                 class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
              <div class="flex items-center space-x-3 space-x-reverse">
                <div class="w-8 h-8 bg-gradient-to-r from-yellow-400 to-orange-500 rounded-full flex items-center justify-center text-white text-sm font-bold">
                  {{ index + 1 }}
                </div>
                <div>
                  <div class="font-medium text-gray-900">{{ customer.name }}</div>
                  <div class="text-sm text-gray-500">{{ customer.email }}</div>
                </div>
              </div>
              <div class="text-right">
                <div class="font-bold text-green-600">{{ formatCurrency(customer.totalSpent) }}</div>
                <div class="text-xs text-gray-500">{{ customer.purchaseCount }} خرید</div>
              </div>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-yellow-600 text-white text-sm rounded-lg hover:bg-yellow-700 focus:outline-none focus:ring-2 focus:ring-yellow-500">
              مشاهده همه
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              مدیریت
            </button>
          </div>
        </div>
      </div>
    </section>

    <!-- بخش آدرس‌ها -->
    <section>
      <div class="bg-gradient-to-r from-green-50 to-emerald-50 rounded-lg p-6 mb-6">
        <h2 class="text-2xl font-bold text-gray-900 mb-2">📍 آدرس‌ها</h2>
        <p class="text-gray-600">مدیریت آدرس‌های ارسال گیفت کارت‌های فیزیکی</p>
      </div>
      
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- آمار آدرس‌ها -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">آمار آدرس‌ها</h3>
            <span class="text-xs bg-green-100 text-green-800 rounded-full px-3 py-1">فعال</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">کل آدرس‌ها:</span>
              <span class="text-lg font-bold text-blue-600">{{ addressStats.totalAddresses }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">آدرس‌های تأیید شده:</span>
              <span class="text-lg font-bold text-green-600">{{ addressStats.verifiedAddresses }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">آدرس‌های در انتظار:</span>
              <span class="text-lg font-bold text-orange-600">{{ addressStats.pendingAddresses }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">آدرس‌های نامعتبر:</span>
              <span class="text-lg font-bold text-red-600">{{ addressStats.invalidAddresses }}</span>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-green-600 text-white text-sm rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500">
              تأیید آدرس‌ها
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              مدیریت
            </button>
          </div>
        </div>

        <!-- مناطق پرترافیک -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">مناطق پرترافیک</h3>
            <span class="text-xs bg-purple-100 text-purple-800 rounded-full px-3 py-1">تحلیل</span>
          </div>
          
          <div class="space-y-3">
            <div v-for="region in topRegions" :key="region.name" 
                 class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
              <div class="flex items-center space-x-3 space-x-reverse">
                <div class="w-6 h-6 bg-purple-500 rounded-full flex items-center justify-center text-white text-xs">
                  {{ region.percentage }}%
                </div>
                <span class="font-medium text-gray-900">{{ region.name }}</span>
              </div>
              <span class="text-sm text-gray-600">{{ region.orderCount }} سفارش</span>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-purple-600 text-white text-sm rounded-lg hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-purple-500">
              تحلیل منطقه‌ای
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              نقشه
            </button>
          </div>
        </div>
      </div>
    </section>

    <!-- بخش ترجیحات -->
    <section>
      <div class="bg-gradient-to-r from-purple-50 to-pink-50 rounded-lg p-6 mb-6">
        <h2 class="text-2xl font-bold text-gray-900 mb-2">🎯 ترجیحات</h2>
        <p class="text-gray-600">مدیریت ترجیحات و علایق مشتریان</p>
      </div>
      
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- دسته‌بندی‌های محبوب -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">دسته‌بندی‌های محبوب</h3>
            <span class="text-xs bg-blue-100 text-blue-800 rounded-full px-3 py-1">محبوب</span>
          </div>
          
          <div class="space-y-3">
            <div v-for="category in popularCategories" :key="category.name" 
                 class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
              <div class="flex items-center space-x-3 space-x-reverse">
                <div class="w-8 h-8 rounded-lg flex items-center justify-center text-white text-sm"
                     :class="category.color">
                  {{ category.icon }}
                </div>
                <span class="font-medium text-gray-900">{{ category.name }}</span>
              </div>
              <div class="text-right">
                <div class="font-bold text-blue-600">{{ category.percentage }}%</div>
                <div class="text-xs text-gray-500">{{ category.purchaseCount }} خرید</div>
              </div>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
              تحلیل ترجیحات
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              گزارش
            </button>
          </div>
        </div>

        <!-- مناسبت‌های محبوب -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">مناسبت‌های محبوب</h3>
            <span class="text-xs bg-green-100 text-green-800 rounded-full px-3 py-1">فعال</span>
          </div>
          
          <div class="space-y-3">
            <div v-for="occasion in popularOccasions" :key="occasion.name" 
                 class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
              <div class="flex items-center space-x-3 space-x-reverse">
                <div class="w-8 h-8 rounded-lg flex items-center justify-center text-white text-sm"
                     :class="occasion.color">
                  {{ occasion.icon }}
                </div>
                <span class="font-medium text-gray-900">{{ occasion.name }}</span>
              </div>
              <div class="text-right">
                <div class="font-bold text-green-600">{{ occasion.percentage }}%</div>
                <div class="text-xs text-gray-500">{{ occasion.orderCount }} سفارش</div>
              </div>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-green-600 text-white text-sm rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500">
              مدیریت مناسبت‌ها
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              تقویم
            </button>
          </div>
        </div>
      </div>
    </section>

    <!-- بخش پروفایل‌ها -->
    <section>
      <div class="bg-gradient-to-r from-orange-50 to-red-50 rounded-lg p-6 mb-6">
        <h2 class="text-2xl font-bold text-gray-900 mb-2">👤 پروفایل‌ها</h2>
        <p class="text-gray-600">مدیریت پروفایل‌های مشتریان و اطلاعات شخصی</p>
      </div>
      
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- آمار پروفایل‌ها -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">آمار پروفایل‌ها</h3>
            <span class="text-xs bg-blue-100 text-blue-800 rounded-full px-3 py-1">فعال</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">پروفایل‌های کامل:</span>
              <span class="text-lg font-bold text-green-600">{{ profileStats.completeProfiles }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">پروفایل‌های ناقص:</span>
              <span class="text-lg font-bold text-orange-600">{{ profileStats.incompleteProfiles }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">پروفایل‌های تأیید شده:</span>
              <span class="text-lg font-bold text-blue-600">{{ profileStats.verifiedProfiles }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">پروفایل‌های جدید:</span>
              <span class="text-lg font-bold text-purple-600">{{ profileStats.newProfiles }}</span>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
              مدیریت پروفایل‌ها
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              تأیید
            </button>
          </div>
        </div>

        <!-- گروه‌های سنی -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">توزیع گروه‌های سنی</h3>
            <span class="text-xs bg-purple-100 text-purple-800 rounded-full px-3 py-1">تحلیل</span>
          </div>
          
          <div class="space-y-3">
            <div v-for="ageGroup in ageGroups" :key="ageGroup.range" 
                 class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
              <div class="flex items-center space-x-3 space-x-reverse">
                <div class="w-6 h-6 rounded-full flex items-center justify-center text-white text-xs"
                     :class="ageGroup.color">
                  {{ ageGroup.percentage }}%
                </div>
                <span class="font-medium text-gray-900">{{ ageGroup.range }}</span>
              </div>
              <span class="text-sm text-gray-600">{{ ageGroup.customerCount }} مشتری</span>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-purple-600 text-white text-sm rounded-lg hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-purple-500">
              تحلیل سنی
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              نمودار
            </button>
          </div>
        </div>
      </div>
    </section>

    <!-- داشبورد کلی مشتریان -->
    <section>
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-xl font-semibold text-gray-900 mb-4">📊 آمار کلی مشتریان</h3>
        
        <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
          <div class="text-center p-6 bg-blue-50 rounded-lg">
            <div class="text-2xl font-bold text-blue-600">{{ overallStats.totalCustomers }}</div>
            <div class="text-sm text-gray-600">کل مشتریان</div>
          </div>
          
          <div class="text-center p-6 bg-green-50 rounded-lg">
            <div class="text-2xl font-bold text-green-600">{{ overallStats.activeCustomers }}</div>
            <div class="text-sm text-gray-600">مشتریان فعال</div>
          </div>
          
          <div class="text-center p-6 bg-purple-50 rounded-lg">
            <div class="text-2xl font-bold text-purple-600">{{ overallStats.newCustomers }}</div>
            <div class="text-sm text-gray-600">مشتریان جدید</div>
          </div>
          
          <div class="text-center p-6 bg-orange-50 rounded-lg">
            <div class="text-2xl font-bold text-orange-600">{{ overallStats.retentionRate }}%</div>
            <div class="text-sm text-gray-600">نرخ بازگشت</div>
          </div>
        </div>
        
        <!-- نمودار روند مشتریان -->
        <div class="mt-6">
          <h4 class="text-sm font-medium text-gray-900 mb-3">روند مشتریان (6 ماه گذشته)</h4>
          <div class="flex items-end space-x-2 space-x-reverse h-32 overflow-x-auto">
            <div v-for="(month, index) in customerTrend" :key="index" class="flex-shrink-0 flex flex-col items-center min-w-16">
              <div class="w-full bg-gray-200 rounded-t relative"
                   :style="{ height: getChartHeight(month.customers) + 'px' }">
                <div class="w-full bg-gradient-to-t from-blue-500 to-purple-500 rounded-t transition-all duration-300 absolute bottom-0"
                     :style="{ height: getChartHeight(month.customers) + 'px' }"></div>
              </div>
              <span class="text-xs text-gray-500 mt-1 text-center">{{ month.label }}</span>
              <span class="text-xs text-gray-400 mt-1 text-center">{{ month.customers.toLocaleString() }}</span>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
// آمار مشتریان
const customerStats = {
  totalCustomers: 15420,
  activeCustomers: 12350,
  todayPurchases: 156,
  averagePurchase: 450000,
  purchaseGrowth: 15.5
}

// مشتریان برتر
const topCustomers = [
  { id: 1, name: 'علی احمدی', email: 'ali@example.com', totalSpent: 25000000, purchaseCount: 45 },
  { id: 2, name: 'فاطمه محمدی', email: 'fateme@example.com', totalSpent: 22000000, purchaseCount: 38 },
  { id: 3, name: 'حسن رضایی', email: 'hasan@example.com', totalSpent: 19800000, purchaseCount: 32 },
  { id: 4, name: 'مریم کریمی', email: 'maryam@example.com', totalSpent: 17500000, purchaseCount: 28 },
  { id: 5, name: 'محمد نوری', email: 'mohammad@example.com', totalSpent: 16200000, purchaseCount: 25 }
]

// آمار آدرس‌ها
const addressStats = {
  totalAddresses: 12500,
  verifiedAddresses: 11800,
  pendingAddresses: 450,
  invalidAddresses: 250
}

// مناطق پرترافیک
const topRegions = [
  { name: 'تهران', percentage: 35, orderCount: 5400 },
  { name: 'اصفهان', percentage: 18, orderCount: 2800 },
  { name: 'مشهد', percentage: 15, orderCount: 2300 },
  { name: 'شیراز', percentage: 12, orderCount: 1850 },
  { name: 'تبریز', percentage: 10, orderCount: 1550 }
]

// دسته‌بندی‌های محبوب
const popularCategories = [
  { name: 'تولد', icon: '🎂', percentage: 28, purchaseCount: 4320, color: 'bg-pink-500' },
  { name: 'عروسی', icon: '💒', percentage: 22, purchaseCount: 3400, color: 'bg-red-500' },
  { name: 'سال نو', icon: '🎊', percentage: 18, purchaseCount: 2780, color: 'bg-green-500' },
  { name: 'ولنتاین', icon: '💝', percentage: 15, purchaseCount: 2320, color: 'bg-purple-500' },
  { name: 'مادر', icon: '🌹', percentage: 12, purchaseCount: 1850, color: 'bg-orange-500' }
]

// مناسبت‌های محبوب
const popularOccasions = [
  { name: 'تولد', icon: '🎂', percentage: 32, orderCount: 4950, color: 'bg-pink-500' },
  { name: 'عروسی', icon: '💒', percentage: 25, orderCount: 3870, color: 'bg-red-500' },
  { name: 'سال نو', icon: '🎊', percentage: 20, orderCount: 3100, color: 'bg-green-500' },
  { name: 'ولنتاین', icon: '💝', percentage: 15, orderCount: 2320, color: 'bg-purple-500' },
  { name: 'مادر', icon: '🌹', percentage: 8, orderCount: 1240, color: 'bg-orange-500' }
]

// آمار پروفایل‌ها
const profileStats = {
  completeProfiles: 11800,
  incompleteProfiles: 3620,
  verifiedProfiles: 10500,
  newProfiles: 1250
}

// گروه‌های سنی
const ageGroups = [
  { range: '18-25', percentage: 25, customerCount: 3850, color: 'bg-blue-500' },
  { range: '26-35', percentage: 35, customerCount: 5390, color: 'bg-green-500' },
  { range: '36-45', percentage: 22, customerCount: 3390, color: 'bg-purple-500' },
  { range: '46-55', percentage: 12, customerCount: 1850, color: 'bg-orange-500' },
  { range: '55+', percentage: 6, customerCount: 940, color: 'bg-red-500' }
]

// آمار کلی
const overallStats = {
  totalCustomers: 15420,
  activeCustomers: 12350,
  newCustomers: 1250,
  retentionRate: 85
}

// روند مشتریان
const customerTrend = [
  { label: 'مرداد', customers: 14200 },
  { label: 'شهریور', customers: 14500 },
  { label: 'مهر', customers: 14800 },
  { label: 'آبان', customers: 15100 },
  { label: 'آذر', customers: 15300 },
  { label: 'دی', customers: 15420 }
]

// تابع فرمت کردن ارز
const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

// تابع محاسبه ارتفاع مناسب برای نمودار
const getChartHeight = (customers: number) => {
  const maxCustomers = Math.max(...customerTrend.map(item => item.customers));
  const minCustomers = Math.min(...customerTrend.map(item => item.customers));
  const range = maxCustomers - minCustomers;
  const height = (customers - minCustomers) / range * 80; // 80 is the max height of the bar
  return Math.max(10, height); // Ensure a minimum height
}
</script>

<!--
  مستندسازی:
  این کامپوننت شامل چهار بخش اصلی است:
  1. تاریخچه خرید: آمار خریدها، مشتریان برتر، روند خرید
  2. آدرس‌ها: آمار آدرس‌ها، مناطق پرترافیک، تأیید آدرس‌ها
  3. ترجیحات: دسته‌بندی‌های محبوب، مناسبت‌های محبوب
  4. پروفایل‌ها: آمار پروفایل‌ها، گروه‌های سنی
  
  تمام توضیحات به فارسی و با طراحی ریسپانسیو ارائه شده است.
  نمودارها و آمار دقیق برای نمایش وضعیت مشتریان استفاده شده‌اند.
-->
<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 

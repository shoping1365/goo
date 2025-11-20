<template>
  <div class="space-y-8">
    <!-- بخش تراکنش‌ها -->
    <section>
      <div class="bg-gradient-to-r from-green-50 to-emerald-50 rounded-lg p-6 mb-6">
        <h2 class="text-2xl font-bold text-gray-900 mb-2">💰 تراکنش‌ها</h2>
        <p class="text-gray-600">مدیریت کامل تراکنش‌های مالی گیفت کارت‌ها</p>
      </div>
      
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- ثبت خرید گیفت کارت -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">ثبت خرید گیفت کارت</h3>
            <span class="text-xs bg-green-100 text-green-800 rounded-full px-3 py-1">فعال</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">خریدهای امروز:</span>
              <span class="text-lg font-bold text-green-600">{{ transactionStats.todayPurchases }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">مبلغ کل امروز:</span>
              <span class="text-lg font-bold text-blue-600">{{ formatCurrency(transactionStats.todayAmount) }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">خریدهای موفق:</span>
              <span class="text-lg font-bold text-green-600">{{ transactionStats.successfulPurchases }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">خریدهای ناموفق:</span>
              <span class="text-lg font-bold text-red-600">{{ transactionStats.failedPurchases }}</span>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-green-600 text-white text-sm rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500">
              ثبت خرید جدید
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              تاریخچه
            </button>
          </div>
        </div>

        <!-- ردیابی استفاده از کارت‌ها -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">ردیابی استفاده از کارت‌ها</h3>
            <span class="text-xs bg-blue-100 text-blue-800 rounded-full px-3 py-1">فعال</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">کارت‌های استفاده شده:</span>
              <span class="text-lg font-bold text-blue-600">{{ transactionStats.usedCards }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">مبلغ مصرف شده:</span>
              <span class="text-lg font-bold text-purple-600">{{ formatCurrency(transactionStats.usedAmount) }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">کارت‌های فعال:</span>
              <span class="text-lg font-bold text-green-600">{{ transactionStats.activeCards }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">اعتبار باقی‌مانده:</span>
              <span class="text-lg font-bold text-orange-600">{{ formatCurrency(transactionStats.remainingBalance) }}</span>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
              مشاهده استفاده‌ها
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              گزارش
            </button>
          </div>
        </div>

        <!-- گزارش تراکنش‌های ناموفق -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">گزارش تراکنش‌های ناموفق</h3>
            <span class="text-xs bg-red-100 text-red-800 rounded-full px-3 py-1">هشدار</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">تراکنش‌های ناموفق:</span>
              <span class="text-lg font-bold text-red-600">{{ transactionStats.failedTransactions }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">در انتظار بررسی:</span>
              <span class="text-lg font-bold text-orange-600">{{ transactionStats.pendingReview }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">مبلغ از دست رفته:</span>
              <span class="text-lg font-bold text-red-600">{{ formatCurrency(transactionStats.lostAmount) }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">نرخ شکست:</span>
              <span class="text-lg font-bold text-yellow-600">{{ transactionStats.failureRate }}%</span>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-red-600 text-white text-sm rounded-lg hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500">
              بررسی خطاها
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              تحلیل
            </button>
          </div>
        </div>

        <!-- مدیریت بازپرداخت -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">مدیریت بازپرداخت</h3>
            <span class="text-xs bg-yellow-100 text-yellow-800 rounded-full px-3 py-1">مدیریت</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">درخواست‌های بازپرداخت:</span>
              <span class="text-lg font-bold text-blue-600">{{ refundStats.refundRequests }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">بازپرداخت‌های انجام شده:</span>
              <span class="text-lg font-bold text-green-600">{{ refundStats.completedRefunds }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">مبلغ کل بازپرداخت:</span>
              <span class="text-lg font-bold text-purple-600">{{ formatCurrency(refundStats.totalRefundAmount) }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">در انتظار تأیید:</span>
              <span class="text-lg font-bold text-orange-600">{{ refundStats.pendingApproval }}</span>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-yellow-600 text-white text-sm rounded-lg hover:bg-yellow-700 focus:outline-none focus:ring-2 focus:ring-yellow-500">
              بررسی درخواست‌ها
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              تاریخچه
            </button>
          </div>
        </div>
      </div>
    </section>

    <!-- بخش قیمت‌گذاری -->
    <section>
      <div class="bg-gradient-to-r from-blue-50 to-indigo-50 rounded-lg p-6 mb-6">
        <h2 class="text-2xl font-bold text-gray-900 mb-2">💎 قیمت‌گذاری</h2>
        <p class="text-gray-600">مدیریت قیمت‌ها، تخفیف‌ها و کارمزدها</p>
      </div>
      
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- تعیین قیمت کارت‌ها -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">تعیین قیمت کارت‌ها</h3>
            <span class="text-xs bg-green-100 text-green-800 rounded-full px-3 py-1">فعال</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">قیمت‌های تعریف شده:</span>
              <span class="text-lg font-bold text-blue-600">{{ pricingStats.definedPrices }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">میانگین قیمت:</span>
              <span class="text-lg font-bold text-green-600">{{ formatCurrency(pricingStats.averagePrice) }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">حداقل قیمت:</span>
              <span class="text-lg font-bold text-purple-600">{{ formatCurrency(pricingStats.minPrice) }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">حداکثر قیمت:</span>
              <span class="text-lg font-bold text-orange-600">{{ formatCurrency(pricingStats.maxPrice) }}</span>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
              تنظیم قیمت
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              مشاهده لیست
            </button>
          </div>
        </div>

        <!-- اعمال تخفیف‌های ویژه -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">اعمال تخفیف‌های ویژه</h3>
            <span class="text-xs bg-purple-100 text-purple-800 rounded-full px-3 py-1">فعال</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">تخفیف‌های فعال:</span>
              <span class="text-lg font-bold text-purple-600">{{ pricingStats.activeDiscounts }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">میانگین تخفیف:</span>
              <span class="text-lg font-bold text-green-600">{{ pricingStats.averageDiscount }}%</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">تخفیف‌های منقضی:</span>
              <span class="text-lg font-bold text-red-600">{{ pricingStats.expiredDiscounts }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">کل صرفه‌جویی:</span>
              <span class="text-lg font-bold text-blue-600">{{ formatCurrency(pricingStats.totalSavings) }}</span>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-purple-600 text-white text-sm rounded-lg hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-purple-500">
              ایجاد تخفیف
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              مدیریت
            </button>
          </div>
        </div>

        <!-- مدیریت کارمزد -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">مدیریت کارمزد</h3>
            <span class="text-xs bg-orange-100 text-orange-800 rounded-full px-3 py-1">فعال</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">کارمزد فعلی:</span>
              <span class="text-lg font-bold text-orange-600">{{ pricingStats.currentFee }}%</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">کارمزد جمع‌آوری شده:</span>
              <span class="text-lg font-bold text-green-600">{{ formatCurrency(pricingStats.collectedFees) }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">کارمزد ماهانه:</span>
              <span class="text-lg font-bold text-blue-600">{{ formatCurrency(pricingStats.monthlyFees) }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">تعداد تراکنش‌های کارمزد:</span>
              <span class="text-lg font-bold text-purple-600">{{ pricingStats.feeTransactions }}</span>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-orange-600 text-white text-sm rounded-lg hover:bg-orange-700 focus:outline-none focus:ring-2 focus:ring-orange-500">
              تنظیم کارمزد
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              گزارش
            </button>
          </div>
        </div>

        <!-- قیمت‌گذاری بر اساس مناسبت‌ها -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">قیمت‌گذاری بر اساس مناسبت‌ها</h3>
            <span class="text-xs bg-green-100 text-green-800 rounded-full px-3 py-1">فعال</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">مناسبت‌های فعال:</span>
              <span class="text-lg font-bold text-green-600">{{ pricingStats.activeOccasions }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">تخفیف مناسبت‌ها:</span>
              <span class="text-lg font-bold text-purple-600">{{ pricingStats.occasionDiscounts }}%</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">فروش مناسبت‌ها:</span>
              <span class="text-lg font-bold text-blue-600">{{ formatCurrency(pricingStats.occasionSales) }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">مناسبت بعدی:</span>
              <span class="text-sm font-medium text-gray-900">{{ pricingStats.nextOccasion }}</span>
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

    <!-- بخش گزارش‌های مالی -->
    <section>
      <div class="bg-gradient-to-r from-purple-50 to-pink-50 rounded-lg p-6 mb-6">
        <h2 class="text-2xl font-bold text-gray-900 mb-2">📊 گزارش‌های مالی</h2>
        <p class="text-gray-600">گزارش‌های جامع مالی و تحلیلی</p>
      </div>
      
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- درآمد کل از گیفت کارت‌ها -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">درآمد کل از گیفت کارت‌ها</h3>
            <span class="text-xs bg-green-100 text-green-800 rounded-full px-3 py-1">مثبت</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">درآمد امروز:</span>
              <span class="text-lg font-bold text-green-600">{{ financialStats.todayRevenue }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">درآمد این ماه:</span>
              <span class="text-lg font-bold text-blue-600">{{ financialStats.monthlyRevenue }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">درآمد این سال:</span>
              <span class="text-lg font-bold text-purple-600">{{ financialStats.yearlyRevenue }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">رشد نسبت به ماه قبل:</span>
              <span class="text-lg font-bold text-green-600">+{{ financialStats.growthRate }}%</span>
            </div>
          </div>
          
          <!-- نمودار درآمد -->
          <div class="mt-4">
            <div class="flex justify-between text-sm text-gray-500 mb-1">
              <span>روند درآمد</span>
              <span>{{ financialStats.revenueTrend }}%</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div
class="bg-gradient-to-r from-green-500 to-blue-500 h-2 rounded-full transition-all duration-300"
                   :style="{ width: financialStats.revenueTrend + '%' }"></div>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-green-600 text-white text-sm rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500">
              مشاهده جزئیات
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              خروجی
            </button>
          </div>
        </div>

        <!-- نرخ تبدیل (فروش به استفاده) -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">نرخ تبدیل (فروش به استفاده)</h3>
            <span class="text-xs bg-blue-100 text-blue-800 rounded-full px-3 py-1">تحلیل</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">نرخ تبدیل کلی:</span>
              <span class="text-lg font-bold text-blue-600">{{ financialStats.conversionRate }}%</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">کارت‌های فروخته شده:</span>
              <span class="text-lg font-bold text-green-600">{{ financialStats.soldCards }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">کارت‌های استفاده شده:</span>
              <span class="text-lg font-bold text-purple-600">{{ financialStats.usedCards }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">کارت‌های منقضی شده:</span>
              <span class="text-lg font-bold text-red-600">{{ financialStats.expiredCards }}</span>
            </div>
          </div>
          
          <!-- نمودار تبدیل -->
          <div class="mt-4">
            <div class="flex justify-between text-sm text-gray-500 mb-1">
              <span>نرخ تبدیل</span>
              <span>{{ financialStats.conversionRate }}%</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div
class="bg-gradient-to-r from-blue-500 to-purple-500 h-2 rounded-full transition-all duration-300"
                   :style="{ width: financialStats.conversionRate + '%' }"></div>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
              تحلیل تبدیل
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              گزارش
            </button>
          </div>
        </div>
      </div>
    </section>

    <!-- داشبورد کلی مالی -->
    <section>
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-xl font-semibold text-gray-900 mb-4">📈 آمار کلی مالی</h3>
        
        <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
          <div class="text-center p-6 bg-green-50 rounded-lg">
            <div class="text-2xl font-bold text-green-600">{{ overallStats.totalRevenue }}</div>
            <div class="text-sm text-gray-600">درآمد کل</div>
          </div>
          
          <div class="text-center p-6 bg-blue-50 rounded-lg">
            <div class="text-2xl font-bold text-blue-600">{{ overallStats.totalTransactions }}</div>
            <div class="text-sm text-gray-600">کل تراکنش‌ها</div>
          </div>
          
          <div class="text-center p-6 bg-purple-50 rounded-lg">
            <div class="text-2xl font-bold text-purple-600">{{ overallStats.profitMargin }}%</div>
            <div class="text-sm text-gray-600">حاشیه سود</div>
          </div>
          
          <div class="text-center p-6 bg-orange-50 rounded-lg">
            <div class="text-2xl font-bold text-orange-600">{{ overallStats.refundRate }}%</div>
            <div class="text-sm text-gray-600">نرخ بازپرداخت</div>
          </div>
        </div>
        
        <!-- نمودار روند مالی -->
        <div class="mt-6">
          <h4 class="text-sm font-medium text-gray-900 mb-3">روند مالی (6 ماه گذشته)</h4>
          <div class="flex items-end space-x-2 space-x-reverse h-32 overflow-x-auto">
            <div v-for="(month, index) in financialTrend" :key="index" class="flex-shrink-0 flex flex-col items-center min-w-16">
              <div
class="w-full bg-gray-200 rounded-t relative"
                   :style="{ height: getChartHeight(month.revenue) + 'px' }">
                <div
class="w-full bg-gradient-to-t from-green-500 to-blue-500 rounded-t transition-all duration-300 absolute bottom-0"
                     :style="{ height: getChartHeight(month.revenue) + 'px' }"></div>
              </div>
              <span class="text-xs text-gray-500 mt-1 text-center">{{ month.label }}</span>
              <span class="text-xs text-gray-400 mt-1 text-center">{{ formatCurrency(month.revenue) }}</span>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
// آمار تراکنش‌ها
const transactionStats = {
  todayPurchases: 156,
  todayAmount: 45000000,
  successfulPurchases: 145,
  failedPurchases: 11,
  usedCards: 8920,
  usedAmount: 125000000,
  activeCards: 6540,
  remainingBalance: 89000000,
  failedTransactions: 23,
  pendingReview: 8,
  lostAmount: 3500000,
  failureRate: 2.3
}

// آمار بازپرداخت
const refundStats = {
  refundRequests: 45,
  completedRefunds: 38,
  totalRefundAmount: 12500000,
  pendingApproval: 7
}

// آمار قیمت‌گذاری
const pricingStats = {
  definedPrices: 25,
  averagePrice: 500000,
  minPrice: 100000,
  maxPrice: 5000000,
  activeDiscounts: 12,
  averageDiscount: 15,
  expiredDiscounts: 3,
  totalSavings: 45000000,
  currentFee: 2.5,
  collectedFees: 12500000,
  monthlyFees: 2500000,
  feeTransactions: 15420,
  activeOccasions: 8,
  occasionDiscounts: 20,
  occasionSales: 89000000,
  nextOccasion: 'تولد (2 روز دیگر)'
}

// آمار مالی
const financialStats = {
  todayRevenue: '45,000,000 تومان',
  monthlyRevenue: '1,250,000,000 تومان',
  yearlyRevenue: '15,800,000,000 تومان',
  growthRate: 12.5,
  revenueTrend: 85,
  conversionRate: 78,
  soldCards: 15420,
  usedCards: 12030,
  expiredCards: 890
}

// آمار کلی
const overallStats = {
  totalRevenue: '15.8 میلیارد',
  totalTransactions: 15420,
  profitMargin: 35,
  refundRate: 2.1
}

// روند مالی
const financialTrend = [
  { label: 'مرداد', revenue: 1200000000 },
  { label: 'شهریور', revenue: 1350000000 },
  { label: 'مهر', revenue: 1180000000 },
  { label: 'آبان', revenue: 1420000000 },
  { label: 'آذر', revenue: 1580000000 },
  { label: 'دی', revenue: 1250000000 }
]

// تابع فرمت کردن ارز
const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

// تابع محاسبه ارتفاع مناسب برای نمودار
const getChartHeight = (revenue: number) => {
  const maxRevenue = Math.max(...financialTrend.map(item => item.revenue));
  const minRevenue = Math.min(...financialTrend.map(item => item.revenue));
  const range = maxRevenue - minRevenue;
  const height = 100; // حداکثر ارتفاع نمودار

  if (range === 0) return height; // اگر تمام مقادیر یکسان باشند

  const percentage = ((revenue - minRevenue) / range) * 100;
  return (percentage / 100) * height;
}
</script>

<!--
  مستندسازی:
  این کامپوننت شامل سه بخش اصلی است:
  1. تراکنش‌ها: ثبت خرید، ردیابی استفاده، گزارش خطاها، مدیریت بازپرداخت
  2. قیمت‌گذاری: تعیین قیمت، تخفیف‌ها، کارمزد، مناسبت‌ها
  3. گزارش‌های مالی: درآمد کل، نرخ تبدیل، آمار کلی
  
  تمام توضیحات به فارسی و با طراحی ریسپانسیو ارائه شده است.
  نمودارها و آمار دقیق برای نمایش وضعیت مالی استفاده شده‌اند.
--> 

<template>
  <div class="space-y-8">
    <!-- بخش چاپ و ارسال -->
    <section>
      <div class="bg-gradient-to-r from-green-50 to-emerald-50 rounded-lg p-6 mb-6">
        <h2 class="text-2xl font-bold text-gray-900 mb-2">🖨️ چاپ و ارسال</h2>
        <p class="text-gray-600">مدیریت کامل فرآیند چاپ، ارسال و ردیابی گیفت کارت‌های فیزیکی</p>
      </div>
      
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- چاپ خودکار کارت‌ها -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">چاپ خودکار کارت‌ها</h3>
            <span class="text-xs bg-green-100 text-green-800 rounded-full px-3 py-1">فعال</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">کارت‌های در انتظار چاپ:</span>
              <span class="text-lg font-bold text-blue-600">{{ printingStats.pendingCards }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">چاپ شده امروز:</span>
              <span class="text-lg font-bold text-green-600">{{ printingStats.printedToday }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">وضعیت چاپگر:</span>
              <span class="text-xs bg-green-100 text-green-800 rounded-full px-2 py-1">آنلاین</span>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
              شروع چاپ
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              تنظیمات
            </button>
          </div>
        </div>

        <!-- مدیریت آدرس‌های ارسال -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">مدیریت آدرس‌های ارسال</h3>
            <span class="text-xs bg-blue-100 text-blue-800 rounded-full px-3 py-1">مدیریت</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">آدرس‌های ثبت شده:</span>
              <span class="text-lg font-bold text-purple-600">{{ shippingStats.totalAddresses }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">آدرس‌های تأیید شده:</span>
              <span class="text-lg font-bold text-green-600">{{ shippingStats.verifiedAddresses }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">آدرس‌های مشکوک:</span>
              <span class="text-lg font-bold text-red-600">{{ shippingStats.suspiciousAddresses }}</span>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-purple-600 text-white text-sm rounded-lg hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-purple-500">
              افزودن آدرس
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              بررسی آدرس‌ها
            </button>
          </div>
        </div>

        <!-- ردیابی ارسال -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">ردیابی ارسال</h3>
            <span class="text-xs bg-yellow-100 text-yellow-800 rounded-full px-3 py-1">فعال</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">در حال ارسال:</span>
              <span class="text-lg font-bold text-blue-600">{{ trackingStats.inTransit }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">تحویل شده:</span>
              <span class="text-lg font-bold text-green-600">{{ trackingStats.delivered }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">در انتظار تأیید:</span>
              <span class="text-lg font-bold text-yellow-600">{{ trackingStats.pendingConfirmation }}</span>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-yellow-600 text-white text-sm rounded-lg hover:bg-yellow-700 focus:outline-none focus:ring-2 focus:ring-yellow-500">
              ردیابی
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              گزارش
            </button>
          </div>
        </div>

        <!-- تأیید دریافت -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">تأیید دریافت</h3>
            <span class="text-xs bg-green-100 text-green-800 rounded-full px-3 py-1">فعال</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">در انتظار تأیید:</span>
              <span class="text-lg font-bold text-orange-600">{{ confirmationStats.pendingConfirmation }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">تأیید شده:</span>
              <span class="text-lg font-bold text-green-600">{{ confirmationStats.confirmed }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">مشکل در تحویل:</span>
              <span class="text-lg font-bold text-red-600">{{ confirmationStats.deliveryIssues }}</span>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-green-600 text-white text-sm rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500">
              تأیید دسته‌ای
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              بررسی مشکلات
            </button>
          </div>
        </div>
      </div>
    </section>

    <!-- بخش انبارداری -->
    <section>
      <div class="bg-gradient-to-r from-orange-50 to-red-50 rounded-lg p-6 mb-6">
        <h2 class="text-2xl font-bold text-gray-900 mb-2">📦 انبارداری</h2>
        <p class="text-gray-600">مدیریت موجودی، هشدارها و سفارش خودکار کارت‌های فیزیکی</p>
      </div>
      
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- موجودی کارت‌های فیزیکی -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">موجودی کارت‌های فیزیکی</h3>
            <span class="text-xs bg-blue-100 text-blue-800 rounded-full px-3 py-1">مدیریت</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">موجودی کل:</span>
              <span class="text-lg font-bold text-blue-600">{{ inventoryStats.totalStock }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">موجودی قابل فروش:</span>
              <span class="text-lg font-bold text-green-600">{{ inventoryStats.availableStock }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">موجودی رزرو شده:</span>
              <span class="text-lg font-bold text-yellow-600">{{ inventoryStats.reservedStock }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">موجودی آسیب دیده:</span>
              <span class="text-lg font-bold text-red-600">{{ inventoryStats.damagedStock }}</span>
            </div>
          </div>
          
          <!-- نوار پیشرفت موجودی -->
          <div class="mt-4">
            <div class="flex justify-between text-sm text-gray-500 mb-1">
              <span>سطح موجودی</span>
              <span>{{ inventoryStats.stockLevel }}%</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div 
                class="h-2 rounded-full transition-all duration-300"
                :class="getStockLevelClass(inventoryStats.stockLevel)"
                :style="{ width: inventoryStats.stockLevel + '%' }"
              ></div>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
              بررسی موجودی
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              گزارش
            </button>
          </div>
        </div>

        <!-- هشدار موجودی کم -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">هشدار موجودی کم</h3>
            <span class="text-xs bg-red-100 text-red-800 rounded-full px-3 py-1">فعال</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">هشدارهای فعال:</span>
              <span class="text-lg font-bold text-red-600">{{ alertStats.activeAlerts }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">موجودی بحرانی:</span>
              <span class="text-lg font-bold text-orange-600">{{ alertStats.criticalStock }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">آخرین هشدار:</span>
              <span class="text-sm text-gray-900">{{ alertStats.lastAlert }}</span>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-red-600 text-white text-sm rounded-lg hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500">
              مشاهده هشدارها
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              تنظیمات
            </button>
          </div>
        </div>

        <!-- سفارش خودکار -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">سفارش خودکار</h3>
            <span class="text-xs bg-green-100 text-green-800 rounded-full px-3 py-1">فعال</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">سفارش‌های خودکار:</span>
              <span class="text-lg font-bold text-blue-600">{{ autoOrderStats.autoOrders }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">در انتظار تأیید:</span>
              <span class="text-lg font-bold text-yellow-600">{{ autoOrderStats.pendingApproval }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">آخرین سفارش:</span>
              <span class="text-sm text-gray-900">{{ autoOrderStats.lastOrder }}</span>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-green-600 text-white text-sm rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500">
              سفارش جدید
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              تاریخچه
            </button>
          </div>
        </div>

        <!-- مدیریت تاریخ انقضا -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">مدیریت تاریخ انقضا</h3>
            <span class="text-xs bg-orange-100 text-orange-800 rounded-full px-3 py-1">هشدار</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">منقضی شده:</span>
              <span class="text-lg font-bold text-red-600">{{ expiryStats.expired }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">در حال انقضا (7 روز):</span>
              <span class="text-lg font-bold text-orange-600">{{ expiryStats.expiringSoon }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">انقضا در ماه آینده:</span>
              <span class="text-lg font-bold text-yellow-600">{{ expiryStats.expiringNextMonth }}</span>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-orange-600 text-white text-sm rounded-lg hover:bg-orange-700 focus:outline-none focus:ring-2 focus:ring-orange-500">
              مدیریت انقضا
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              گزارش
            </button>
          </div>
        </div>
      </div>
    </section>

    <!-- داشبورد کلی -->
    <section>
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-xl font-semibold text-gray-900 mb-4">📊 آمار کلی مدیریت فیزیکی</h3>
        
        <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
          <div class="text-center p-6 bg-green-50 rounded-lg">
            <div class="text-2xl font-bold text-green-600">{{ overallStats.totalShipped }}</div>
            <div class="text-sm text-gray-600">کل ارسال شده</div>
          </div>
          
          <div class="text-center p-6 bg-blue-50 rounded-lg">
            <div class="text-2xl font-bold text-blue-600">{{ overallStats.inStock }}</div>
            <div class="text-sm text-gray-600">موجودی انبار</div>
          </div>
          
          <div class="text-center p-6 bg-yellow-50 rounded-lg">
            <div class="text-2xl font-bold text-yellow-600">{{ overallStats.pendingDelivery }}</div>
            <div class="text-sm text-gray-600">در انتظار تحویل</div>
          </div>
          
          <div class="text-center p-6 bg-red-50 rounded-lg">
            <div class="text-2xl font-bold text-red-600">{{ overallStats.deliveryIssues }}</div>
            <div class="text-sm text-gray-600">مشکلات تحویل</div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
// آمار چاپ و ارسال
const printingStats = {
  pendingCards: 156,
  printedToday: 89,
  printerStatus: 'online'
}

const shippingStats = {
  totalAddresses: 1247,
  verifiedAddresses: 1189,
  suspiciousAddresses: 58
}

const trackingStats = {
  inTransit: 234,
  delivered: 1891,
  pendingConfirmation: 45
}

const confirmationStats = {
  pendingConfirmation: 45,
  confirmed: 1891,
  deliveryIssues: 12
}

// آمار انبارداری
const inventoryStats = {
  totalStock: 15420,
  availableStock: 12890,
  reservedStock: 2340,
  damagedStock: 190,
  stockLevel: 83
}

const alertStats = {
  activeAlerts: 8,
  criticalStock: 156,
  lastAlert: '2 ساعت پیش'
}

const autoOrderStats = {
  autoOrders: 23,
  pendingApproval: 5,
  lastOrder: 'دیروز'
}

const expiryStats = {
  expired: 89,
  expiringSoon: 234,
  expiringNextMonth: 567
}

// آمار کلی
const overallStats = {
  totalShipped: 2156,
  inStock: 15420,
  pendingDelivery: 234,
  deliveryIssues: 12
}

// تابع تعیین کلاس سطح موجودی
const getStockLevelClass = (level: number) => {
  if (level >= 80) return 'bg-green-500'
  if (level >= 50) return 'bg-yellow-500'
  if (level >= 20) return 'bg-orange-500'
  return 'bg-red-500'
}
</script>

<!--
  مستندسازی:
  این کامپوننت شامل دو بخش اصلی است:
  1. چاپ و ارسال: چاپ خودکار، مدیریت آدرس‌ها، ردیابی، تأیید دریافت
  2. انبارداری: موجودی، هشدارها، سفارش خودکار، مدیریت انقضا
  
  تمام توضیحات به فارسی و با طراحی ریسپانسیو ارائه شده است.
  آمار و اعداد نمونه برای نمایش قابلیت‌ها استفاده شده‌اند.
--> 

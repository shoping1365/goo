<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="bg-gradient-to-r from-gray-50 to-slate-50 rounded-lg p-6">
      <h2 class="text-2xl font-bold text-gray-900 mb-2">⚙️ تنظیمات کیف پول</h2>
      <p class="text-gray-600">مدیریت تنظیمات کلی و امنیتی کیف پول</p>
    </div>

    <!-- تنظیمات عمومی -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات عمومی</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نام کیف پول</label>
            <input type="text" :value="generalSettings.walletName" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">واحد پول پیش‌فرض</label>
            <select class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
              <option value="IRR" selected>ریال ایران</option>
              <option value="USD">دلار آمریکا</option>
              <option value="EUR">یورو</option>
              <option value="GBP">پوند انگلیس</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">زبان پیش‌فرض</label>
            <select class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
              <option value="fa" selected>فارسی</option>
              <option value="en">انگلیسی</option>
              <option value="ar">عربی</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">منطقه زمانی</label>
            <select class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
              <option value="Asia/Tehran" selected>تهران (UTC+3:30)</option>
              <option value="UTC">UTC</option>
              <option value="America/New_York">نیویورک (UTC-5)</option>
            </select>
          </div>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر موجودی کیف پول</label>
            <input type="number" :value="generalSettings.maxWalletBalance" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداقل موجودی برای برداشت</label>
            <input type="number" :value="generalSettings.minWithdrawalBalance" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تعداد کارت‌های مجاز</label>
            <input type="number" :value="generalSettings.maxCardsPerUser" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">مدت اعتبار کد تأیید (دقیقه)</label>
            <input type="number" :value="generalSettings.verificationCodeExpiry" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
        </div>
      </div>
      
      <div class="mt-6">
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
          ذخیره تنظیمات عمومی
        </button>
      </div>
    </div>

    <!-- تنظیمات امنیتی -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات امنیتی</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="space-y-4">
          <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <div>
              <h4 class="font-medium text-gray-900">احراز هویت دو مرحله‌ای</h4>
              <p class="text-sm text-gray-600">استفاده از کد تأیید برای ورود</p>
            </div>
            <input type="checkbox" :checked="securitySettings.twoFactorAuth" 
                   class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
          </div>
          
          <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <div>
              <h4 class="font-medium text-gray-900">رمز عبور قوی</h4>
              <p class="text-sm text-gray-600">اجبار استفاده از رمز عبور پیچیده</p>
            </div>
            <input type="checkbox" :checked="securitySettings.strongPassword" 
                   class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
          </div>
          
          <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <div>
              <h4 class="font-medium text-gray-900">تأیید تراکنش‌های بزرگ</h4>
              <p class="text-sm text-gray-600">درخواست تأیید برای تراکنش‌های بالای حد مشخص</p>
            </div>
            <input type="checkbox" :checked="securitySettings.largeTransactionVerification" 
                   class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
          </div>
          
          <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <div>
              <h4 class="font-medium text-gray-900">ثبت فعالیت‌ها</h4>
              <p class="text-sm text-gray-600">ثبت تمام فعالیت‌های کاربران</p>
            </div>
            <input type="checkbox" :checked="securitySettings.activityLogging" 
                   class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
          </div>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حد آستانه تراکنش بزرگ (تومان)</label>
            <input type="number" :value="securitySettings.largeTransactionThreshold" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر تلاش ورود</label>
            <input type="number" :value="securitySettings.maxLoginAttempts" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">مدت قفل حساب (دقیقه)</label>
            <input type="number" :value="securitySettings.accountLockDuration" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">مدت اعتبار نشست (ساعت)</label>
            <input type="number" :value="securitySettings.sessionTimeout" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
        </div>
      </div>
      
      <div class="mt-6">
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
          ذخیره تنظیمات امنیتی
        </button>
      </div>
    </div>

    <!-- تنظیمات اعلان‌ها -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات اعلان‌ها</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="space-y-4">
          <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <div>
              <h4 class="font-medium text-gray-900">اعلان شارژ</h4>
              <p class="text-sm text-gray-600">ارسال اعلان هنگام شارژ کیف پول</p>
            </div>
            <input type="checkbox" :checked="notificationSettings.chargeNotification" 
                   class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
          </div>
          
          <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <div>
              <h4 class="font-medium text-gray-900">اعلان برداشت</h4>
              <p class="text-sm text-gray-600">ارسال اعلان هنگام برداشت از کیف پول</p>
            </div>
            <input type="checkbox" :checked="notificationSettings.withdrawalNotification" 
                   class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
          </div>
          
          <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <div>
              <h4 class="font-medium text-gray-900">اعلان تراکنش ناموفق</h4>
              <p class="text-sm text-gray-600">ارسال اعلان برای تراکنش‌های ناموفق</p>
            </div>
            <input type="checkbox" :checked="notificationSettings.failedTransactionNotification" 
                   class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
          </div>
          
          <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <div>
              <h4 class="font-medium text-gray-900">اعلان موجودی کم</h4>
              <p class="text-sm text-gray-600">اعلان هنگام کاهش موجودی</p>
            </div>
            <input type="checkbox" :checked="notificationSettings.lowBalanceNotification" 
                   class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
          </div>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حد آستانه موجودی کم (تومان)</label>
            <input type="number" :value="notificationSettings.lowBalanceThreshold" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">کانال‌های اعلان</label>
            <div class="space-y-2">
              <div class="flex items-center">
                <input type="checkbox" :checked="notificationSettings.emailNotifications" 
                       class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                <label class="mr-2 text-sm text-gray-700">ایمیل</label>
              </div>
              <div class="flex items-center">
                <input type="checkbox" :checked="notificationSettings.smsNotifications" 
                       class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                <label class="mr-2 text-sm text-gray-700">پیامک</label>
              </div>
              <div class="flex items-center">
                <input type="checkbox" :checked="notificationSettings.pushNotifications" 
                       class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                <label class="mr-2 text-sm text-gray-700">اعلان مرورگر</label>
              </div>
            </div>
          </div>
        </div>
      </div>
      
      <div class="mt-6">
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
          ذخیره تنظیمات اعلان‌ها
        </button>
      </div>
    </div>

    <!-- تنظیمات پیشرفته -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات پیشرفته</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="space-y-4">
          <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <div>
              <h4 class="font-medium text-gray-900">حالت نگهداری</h4>
              <p class="text-sm text-gray-600">فعال کردن حالت نگهداری سیستم</p>
            </div>
            <input type="checkbox" :checked="advancedSettings.maintenanceMode" 
                   class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
          </div>
          
          <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <div>
              <h4 class="font-medium text-gray-900">پشتیبان‌گیری خودکار</h4>
              <p class="text-sm text-gray-600">پشتیبان‌گیری روزانه از داده‌ها</p>
            </div>
            <input type="checkbox" :checked="advancedSettings.autoBackup" 
                   class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
          </div>
          
          <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <div>
              <h4 class="font-medium text-gray-900">لاگ‌گیری تفصیلی</h4>
              <p class="text-sm text-gray-600">ثبت جزئیات کامل فعالیت‌ها</p>
            </div>
            <input type="checkbox" :checked="advancedSettings.detailedLogging" 
                   class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
          </div>
          
          <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <div>
              <h4 class="font-medium text-gray-900">تست اتصال</h4>
              <p class="text-sm text-gray-600">تست خودکار اتصالات</p>
            </div>
            <input type="checkbox" :checked="advancedSettings.connectionTesting" 
                   class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
          </div>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">فاصله پشتیبان‌گیری (ساعت)</label>
            <input type="number" :value="advancedSettings.backupInterval" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">مدت نگهداری لاگ‌ها (روز)</label>
            <input type="number" :value="advancedSettings.logRetentionDays" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">فاصله تست اتصال (دقیقه)</label>
            <input type="number" :value="advancedSettings.connectionTestInterval" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">سطح لاگ</label>
            <select class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
              <option value="debug">Debug</option>
              <option value="info" selected>Info</option>
              <option value="warning">Warning</option>
              <option value="error">Error</option>
            </select>
          </div>
        </div>
      </div>
      
      <div class="mt-6">
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
          ذخیره تنظیمات پیشرفته
        </button>
      </div>
    </div>

    <!-- عملیات سیستمی -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">عملیات سیستمی</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <button class="px-4 py-3 bg-green-100 text-green-700 rounded-lg hover:bg-green-200 focus:outline-none focus:ring-2 focus:ring-green-500">
          <div class="text-center">
            <div class="text-lg mb-1">🔄</div>
            <div class="font-medium">پاک کردن کش</div>
            <div class="text-sm text-green-600">پاک کردن کش سیستم</div>
          </div>
        </button>
        
        <button class="px-4 py-3 bg-blue-100 text-blue-700 rounded-lg hover:bg-blue-200 focus:outline-none focus:ring-2 focus:ring-blue-500">
          <div class="text-center">
            <div class="text-lg mb-1">📊</div>
            <div class="font-medium">تست اتصال</div>
            <div class="text-sm text-blue-600">تست اتصالات خارجی</div>
          </div>
        </button>
        
        <button class="px-4 py-3 bg-purple-100 text-purple-700 rounded-lg hover:bg-purple-200 focus:outline-none focus:ring-2 focus:ring-purple-500">
          <div class="text-center">
            <div class="text-lg mb-1">💾</div>
            <div class="font-medium">پشتیبان‌گیری</div>
            <div class="text-sm text-purple-600">ایجاد پشتیبان دستی</div>
          </div>
        </button>
        
        <button class="px-4 py-3 bg-yellow-100 text-yellow-700 rounded-lg hover:bg-yellow-200 focus:outline-none focus:ring-2 focus:ring-yellow-500">
          <div class="text-center">
            <div class="text-lg mb-1">🔧</div>
            <div class="font-medium">بازسازی دیتابیس</div>
            <div class="text-sm text-yellow-600">بهینه‌سازی دیتابیس</div>
          </div>
        </button>
        
        <button class="px-4 py-3 bg-red-100 text-red-700 rounded-lg hover:bg-red-200 focus:outline-none focus:ring-2 focus:ring-red-500">
          <div class="text-center">
            <div class="text-lg mb-1">⚠️</div>
            <div class="font-medium">حالت نگهداری</div>
            <div class="text-sm text-red-600">فعال/غیرفعال کردن</div>
          </div>
        </button>
        
        <button class="px-4 py-3 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
          <div class="text-center">
            <div class="text-lg mb-1">📋</div>
            <div class="font-medium">گزارش سیستم</div>
            <div class="text-sm text-gray-600">نمایش وضعیت سیستم</div>
          </div>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// تنظیمات عمومی
const generalSettings = {
  walletName: 'کیف پول دیجیتال',
  maxWalletBalance: 1000000000,
  minWithdrawalBalance: 100000,
  maxCardsPerUser: 5,
  verificationCodeExpiry: 5
}

// تنظیمات امنیتی
const securitySettings = {
  twoFactorAuth: true,
  strongPassword: true,
  largeTransactionVerification: true,
  activityLogging: true,
  largeTransactionThreshold: 10000000,
  maxLoginAttempts: 5,
  accountLockDuration: 30,
  sessionTimeout: 24
}

// تنظیمات اعلان‌ها
const notificationSettings = {
  chargeNotification: true,
  withdrawalNotification: true,
  failedTransactionNotification: true,
  lowBalanceNotification: true,
  lowBalanceThreshold: 100000,
  emailNotifications: true,
  smsNotifications: true,
  pushNotifications: false
}

// تنظیمات پیشرفته
const advancedSettings = {
  maintenanceMode: false,
  autoBackup: true,
  detailedLogging: true,
  connectionTesting: true,
  backupInterval: 24,
  logRetentionDays: 30,
  connectionTestInterval: 15
}
</script>

<!--
  مستندسازی:
  این کامپوننت شامل بخش‌های زیر است:
  1. تنظیمات عمومی کیف پول
  2. تنظیمات امنیتی
  3. تنظیمات اعلان‌ها
  4. تنظیمات پیشرفته
  5. عملیات سیستمی
  
  تمام توضیحات به فارسی و با طراحی ریسپانسیو ارائه شده است.
--> 

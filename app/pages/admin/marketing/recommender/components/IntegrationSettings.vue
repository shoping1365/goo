<template>
  <div v-if="hasAccess">
  <div class="space-y-8">
    <!-- هدر بخش -->
    <div class="flex justify-between items-center">
      <div>
        <h2 class="text-2xl font-bold text-gray-900">یکپارچه‌سازی</h2>
        <p class="mt-1 text-sm text-gray-600">مدیریت اتصال سیستم پیشنهادات با سایر سرویس‌ها</p>
      </div>
      <button class="bg-indigo-600 text-white px-4 py-2 rounded-lg hover:bg-indigo-700 transition-colors">
        <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
        </svg>
        افزودن اتصال جدید
      </button>
    </div>

    <!-- آمار یکپارچه‌سازی -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-6">
      <div class="bg-gradient-to-r from-blue-500 to-blue-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">ایمیل مارکتینگ</p>
            <p class="text-2xl font-bold">24,567</p>
            <p class="text-xs opacity-75">پیشنهاد ارسال شده</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 4.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-green-500 to-green-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">پیامک</p>
            <p class="text-2xl font-bold">12,345</p>
            <p class="text-xs opacity-75">پیشنهاد ارسال شده</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-purple-500 to-purple-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">پوش نوتیفیکیشن</p>
            <p class="text-2xl font-bold">8,234</p>
            <p class="text-xs opacity-75">پیشنهاد ارسال شده</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-5 5v-5zM4.19 4.19A2 2 0 004 6v12a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-1.81 1.19z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-yellow-500 to-yellow-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">چت زنده</p>
            <p class="text-2xl font-bold">3,456</p>
            <p class="text-xs opacity-75">پیشنهاد ارسال شده</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-red-500 to-red-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">API خارجی</p>
            <p class="text-2xl font-bold">1,234</p>
            <p class="text-xs opacity-75">درخواست امروز</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4"></path>
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات یکپارچه‌سازی -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <!-- ایمیل مارکتینگ -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-blue-100 rounded-lg">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 4.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">ایمیل مارکتینگ</h3>
          </div>
          <label class="relative inline-flex items-center cursor-pointer">
            <input v-model="integrations.email.enabled" type="checkbox" class="sr-only peer">
            <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
          </label>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">فرکانس ارسال</label>
            <select v-model="integrations.email.frequency" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
              <option value="daily">روزانه</option>
              <option value="weekly">هفتگی</option>
              <option value="monthly">ماهانه</option>
              <option value="custom">سفارشی</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر پیشنهاد در هر ایمیل</label>
            <input v-model="integrations.email.maxRecommendations" type="number" min="1" max="10" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">قالب ایمیل</label>
            <select v-model="integrations.email.template" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
              <option value="simple">ساده</option>
              <option value="detailed">مفصل</option>
              <option value="promotional">تبلیغاتی</option>
            </select>
          </div>
        </div>
      </div>

      <!-- پیامک -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-green-100 rounded-lg">
              <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">پیامک</h3>
          </div>
          <label class="relative inline-flex items-center cursor-pointer">
            <input v-model="integrations.sms.enabled" type="checkbox" class="sr-only peer">
            <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-green-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-green-600"></div>
          </label>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">ساعت ارسال</label>
            <input v-model="integrations.sms.sendTime" type="time" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">متن پیشنهاد</label>
            <textarea v-model="integrations.sms.messageTemplate" rows="3" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500" placeholder="متن پیشنهاد برای پیامک..."></textarea>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر طول پیام</label>
            <input v-model="integrations.sms.maxLength" type="number" min="50" max="160" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500">
          </div>
        </div>
      </div>
    </div>

    <!-- پوش نوتیفیکیشن و چت زنده -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <!-- پوش نوتیفیکیشن -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-purple-100 rounded-lg">
              <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-5 5v-5zM4.19 4.19A2 2 0 004 6v12a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-1.81 1.19z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">پوش نوتیفیکیشن</h3>
          </div>
          <label class="relative inline-flex items-center cursor-pointer">
            <input v-model="integrations.push.enabled" type="checkbox" class="sr-only peer">
            <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-purple-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-purple-600"></div>
          </label>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع اعلان</label>
            <div class="space-y-2">
              <label class="flex items-center">
                <input v-model="integrations.push.types" type="checkbox" value="product" class="rounded border-gray-300 text-purple-600 focus:ring-purple-500">
                <span class="mr-2 text-sm text-gray-700">پیشنهاد محصول</span>
              </label>
              <label class="flex items-center">
                <input v-model="integrations.push.types" type="checkbox" value="discount" class="rounded border-gray-300 text-purple-600 focus:ring-purple-500">
                <span class="mr-2 text-sm text-gray-700">تخفیف ویژه</span>
              </label>
              <label class="flex items-center">
                <input v-model="integrations.push.types" type="checkbox" value="reminder" class="rounded border-gray-300 text-purple-600 focus:ring-purple-500">
                <span class="mr-2 text-sm text-gray-700">یادآوری</span>
              </label>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تصویر اعلان</label>
            <input type="file" accept="image/*" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">لینک عمیق</label>
            <input v-model="integrations.push.deepLink" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500" placeholder="https://example.com/product/123">
          </div>
        </div>
      </div>

      <!-- چت زنده -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-yellow-100 rounded-lg">
              <svg class="w-6 h-6 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">چت زنده</h3>
          </div>
          <label class="relative inline-flex items-center cursor-pointer">
            <input v-model="integrations.chat.enabled" type="checkbox" class="sr-only peer">
            <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-yellow-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-yellow-600"></div>
          </label>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">زمان تاخیر پیشنهاد</label>
            <input v-model="integrations.chat.delaySeconds" type="number" min="5" max="300" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500">
            <p class="text-xs text-gray-500 mt-1">ثانیه</p>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">شرایط پیشنهاد</label>
            <div class="space-y-2">
              <label class="flex items-center">
                <input v-model="integrations.chat.conditions" type="checkbox" value="idle" class="rounded border-gray-300 text-yellow-600 focus:ring-yellow-500">
                <span class="mr-2 text-sm text-gray-700">کاربر غیرفعال</span>
              </label>
              <label class="flex items-center">
                <input v-model="integrations.chat.conditions" type="checkbox" value="browsing" class="rounded border-gray-300 text-yellow-600 focus:ring-yellow-500">
                <span class="mr-2 text-sm text-gray-700">در حال مرور</span>
              </label>
              <label class="flex items-center">
                <input v-model="integrations.chat.conditions" type="checkbox" value="cart" class="rounded border-gray-300 text-yellow-600 focus:ring-yellow-500">
                <span class="mr-2 text-sm text-gray-700">سبد خرید خالی</span>
              </label>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">متن پیشنهاد چت</label>
            <textarea v-model="integrations.chat.messageTemplate" rows="3" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500" placeholder="متن پیشنهاد برای چت..."></textarea>
          </div>
        </div>
      </div>
    </div>

    <!-- API خارجی -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <div class="flex items-center">
          <div class="p-2 bg-red-100 rounded-lg">
            <svg class="w-6 h-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4"></path>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-900 mr-3">API خارجی</h3>
        </div>
        <label class="relative inline-flex items-center cursor-pointer">
          <input v-model="integrations.api.enabled" type="checkbox" class="sr-only peer">
          <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-red-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-red-600"></div>
        </label>
      </div>
      
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Endpoint URL</label>
            <input v-model="integrations.api.endpoint" type="url" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500" placeholder="https://api.example.com/recommendations">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">API Key</label>
            <input v-model="integrations.api.apiKey" type="password" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500" placeholder="API Key">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع احراز هویت</label>
            <select v-model="integrations.api.authType" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500">
              <option value="bearer">Bearer Token</option>
              <option value="basic">Basic Auth</option>
              <option value="api_key">API Key</option>
              <option value="oauth">OAuth 2.0</option>
            </select>
          </div>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">محدودیت درخواست</label>
            <input v-model="integrations.api.rateLimit" type="number" min="1" max="1000" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500">
            <p class="text-xs text-gray-500 mt-1">درخواست در دقیقه</p>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Timeout</label>
            <input v-model="integrations.api.timeout" type="number" min="1" max="60" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500">
            <p class="text-xs text-gray-500 mt-1">ثانیه</p>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Retry Count</label>
            <input v-model="integrations.api.retryCount" type="number" min="0" max="5" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500">
            <p class="text-xs text-gray-500 mt-1">تعداد تلاش مجدد</p>
          </div>
        </div>
      </div>
      
      <div class="mt-6 pt-6 border-t border-gray-200">
        <button class="bg-red-600 text-white px-4 py-2 rounded-lg hover:bg-red-700 transition-colors">
          تست اتصال
        </button>
        <button class="bg-gray-600 text-white px-4 py-2 rounded-lg hover:bg-gray-700 transition-colors mr-3">
          ذخیره تنظیمات
        </button>
      </div>
    </div>

    <!-- اتصالات فعال -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-semibold text-gray-900">اتصالات فعال</h3>
      </div>
      <div class="p-6">
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">آخرین ارسال</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">آمار امروز</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="connection in activeConnections" :key="connection.id">
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="flex items-center">
                    <div class="w-8 h-8 rounded-lg mr-3 flex items-center justify-center" :class="connection.bgColor">
                      <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="connection.icon"></path>
                      </svg>
                    </div>
                    <div>
                      <div class="text-sm font-medium text-gray-900">{{ connection.name }}</div>
                      <div class="text-sm text-gray-500">{{ connection.description }}</div>
                    </div>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full" :class="connection.statusClass">
                    {{ connection.status }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  {{ connection.lastSent }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  {{ connection.todayCount }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                  <button class="text-indigo-600 hover:text-indigo-900 mr-3">ویرایش</button>
                  <button class="text-red-600 hover:text-red-900">حذف</button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
  </div>
  <div v-else class="text-center p-8 text-gray-500">
    <p>شما مجوز دسترسی به این بخش را ندارید.</p>
  </div>
</template>

<script setup lang="ts">
import { useAuth } from '@/composables/useAuth';
import { computed, onMounted, ref, watch } from 'vue';
import { useRouter } from 'vue-router';

const { user } = useAuth();
const router = useRouter();

const hasAccess = computed(() => {
  return ['admin', 'developer'].includes(user.value?.role || '');
});

watch(hasAccess, (newValue) => {
  if (!newValue) {
    router.push('/404');
  }
});

onMounted(() => {
  if (!hasAccess.value) {
    router.push('/404');
  }
});

// تنظیمات یکپارچه‌سازی
const integrations = ref({
  email: {
    enabled: true,
    frequency: 'weekly',
    maxRecommendations: 5,
    template: 'simple'
  },
  sms: {
    enabled: false,
    sendTime: '10:00',
    messageTemplate: 'پیشنهاد ویژه برای شما: {product_name} با {discount}% تخفیف',
    maxLength: 160
  },
  push: {
    enabled: true,
    types: ['product', 'discount'],
    deepLink: ''
  },
  chat: {
    enabled: true,
    delaySeconds: 30,
    conditions: ['idle', 'browsing'],
    messageTemplate: 'آیا می‌خواهید محصولات مشابه را ببینید؟'
  },
  api: {
    enabled: false,
    endpoint: '',
    apiKey: '',
    authType: 'bearer',
    rateLimit: 100,
    timeout: 30,
    retryCount: 3
  }
})

// اتصالات فعال
const activeConnections = ref([
  {
    id: 1,
    name: 'ایمیل مارکتینگ',
    description: 'ارسال پیشنهادات از طریق ایمیل',
    status: 'فعال',
    statusClass: 'bg-green-100 text-green-800',
    lastSent: '2 ساعت پیش',
    todayCount: '1,234',
    bgColor: 'bg-blue-500',
    icon: 'M3 8l7.89 4.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z'
  },
  {
    id: 2,
    name: 'پیامک',
    description: 'ارسال پیشنهادات از طریق پیامک',
    status: 'غیرفعال',
    statusClass: 'bg-gray-100 text-gray-800',
    lastSent: '1 روز پیش',
    todayCount: '0',
    bgColor: 'bg-green-500',
    icon: 'M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z'
  },
  {
    id: 3,
    name: 'پوش نوتیفیکیشن',
    description: 'ارسال اعلان‌های پیشنهاد',
    status: 'فعال',
    statusClass: 'bg-green-100 text-green-800',
    lastSent: '30 دقیقه پیش',
    todayCount: '567',
    bgColor: 'bg-purple-500',
    icon: 'M15 17h5l-5 5v-5zM4.19 4.19A2 2 0 004 6v12a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-1.81 1.19z'
  },
  {
    id: 4,
    name: 'چت زنده',
    description: 'پیشنهادات در چت زنده',
    status: 'فعال',
    statusClass: 'bg-green-100 text-green-800',
    lastSent: '5 دقیقه پیش',
    todayCount: '89',
    bgColor: 'bg-yellow-500',
    icon: 'M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z'
  }
])
</script> 

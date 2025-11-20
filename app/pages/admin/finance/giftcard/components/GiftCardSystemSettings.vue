<template>
  <div class="space-y-8">
    <!-- بخش تنظیمات عمومی -->
    <section>
      <div class="bg-gradient-to-r from-blue-50 to-indigo-50 rounded-lg p-6 mb-6">
        <h2 class="text-2xl font-bold text-gray-900 mb-2">⚙️ تنظیمات عمومی</h2>
        <p class="text-gray-600">مدیریت تنظیمات کلی سیستم گیفت کارت</p>
      </div>
      
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- تنظیمات پایه -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">تنظیمات پایه</h3>
            <span class="text-xs bg-green-100 text-green-800 rounded-full px-3 py-1">فعال</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">واحد پول:</span>
              <select class="px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
                <option>تومان (IRR)</option>
                <option>ریال (IRR)</option>
                <option>دلار (USD)</option>
                <option>یورو (EUR)</option>
              </select>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">زبان سیستم:</span>
              <select class="px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
                <option>فارسی</option>
                <option>انگلیسی</option>
                <option>عربی</option>
              </select>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">منطقه زمانی:</span>
              <select class="px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
                <option>تهران (UTC+3:30)</option>
                <option>لندن (UTC+0)</option>
                <option>نیویورک (UTC-5)</option>
              </select>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">فرمت تاریخ:</span>
              <select class="px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
                <option>شمسی (1402/12/25)</option>
                <option>میلادی (2024/03/15)</option>
                <option>هجری قمری (1445/09/03)</option>
              </select>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
              ذخیره تنظیمات
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              بازنشانی
            </button>
          </div>
        </div>

        <!-- تنظیمات کارت -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">تنظیمات کارت</h3>
            <span class="text-xs bg-purple-100 text-purple-800 rounded-full px-3 py-1">پیشرفته</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">حداقل مبلغ کارت:</span>
              <input type="number" value="10000" class="px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 w-32" />
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">حداکثر مبلغ کارت:</span>
              <input type="number" value="10000000" class="px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 w-32" />
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">مدت اعتبار پیش‌فرض:</span>
              <select class="px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
                <option>1 سال</option>
                <option>6 ماه</option>
                <option>3 ماه</option>
                <option>1 ماه</option>
                <option>بدون محدودیت</option>
              </select>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">کد پیش‌فرض:</span>
              <input type="text" value="GC-" class="px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 w-32" />
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-purple-600 text-white text-sm rounded-lg hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-purple-500">
              اعمال تنظیمات
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              پیش‌نمایش
            </button>
          </div>
        </div>
      </div>
    </section>

    <!-- بخش تنظیمات امنیت -->
    <section>
      <div class="bg-gradient-to-r from-red-50 to-pink-50 rounded-lg p-6 mb-6">
        <h2 class="text-2xl font-bold text-gray-900 mb-2">🔒 تنظیمات امنیت</h2>
        <p class="text-gray-600">مدیریت امنیت و حریم خصوصی سیستم</p>
      </div>
      
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- امنیت احراز هویت -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">امنیت احراز هویت</h3>
            <span class="text-xs bg-red-100 text-red-800 rounded-full px-3 py-1">امنیت</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">احراز هویت دو مرحله‌ای:</span>
              <label class="relative inline-flex items-center cursor-pointer">
                <input type="checkbox" checked class="sr-only peer">
                <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
              </label>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">حداقل طول رمز عبور:</span>
              <input type="number" value="8" min="6" max="20" class="px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 w-20" />
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">انقضای رمز عبور:</span>
              <select class="px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
                <option>90 روز</option>
                <option>60 روز</option>
                <option>30 روز</option>
                <option>بدون انقضا</option>
              </select>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">محدودیت تلاش ورود:</span>
              <input type="number" value="5" min="3" max="10" class="px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 w-20" />
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-red-600 text-white text-sm rounded-lg hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500">
              ذخیره امنیت
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              تست امنیت
            </button>
          </div>
        </div>

        <!-- تنظیمات حریم خصوصی -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">تنظیمات حریم خصوصی</h3>
            <span class="text-xs bg-blue-100 text-blue-800 rounded-full px-3 py-1">خصوصی</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">ذخیره لاگ‌ها:</span>
              <label class="relative inline-flex items-center cursor-pointer">
                <input type="checkbox" checked class="sr-only peer">
                <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
              </label>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">مدت نگهداری لاگ:</span>
              <select class="px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
                <option>1 سال</option>
                <option>6 ماه</option>
                <option>3 ماه</option>
                <option>1 ماه</option>
              </select>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">رمزگذاری داده‌ها:</span>
              <label class="relative inline-flex items-center cursor-pointer">
                <input type="checkbox" checked class="sr-only peer">
                <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
              </label>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">GDPR Compliance:</span>
              <label class="relative inline-flex items-center cursor-pointer">
                <input type="checkbox" class="sr-only peer">
                <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
              </label>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
              ذخیره حریم خصوصی
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              گزارش
            </button>
          </div>
        </div>
      </div>
    </section>

    <!-- بخش تنظیمات اعلان‌ها -->
    <section>
      <div class="bg-gradient-to-r from-green-50 to-emerald-50 rounded-lg p-6 mb-6">
        <h2 class="text-2xl font-bold text-gray-900 mb-2">🔔 تنظیمات اعلان‌ها</h2>
        <p class="text-gray-600">مدیریت اعلان‌ها و هشدارهای سیستم</p>
      </div>
      
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- اعلان‌های ایمیل -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">اعلان‌های ایمیل</h3>
            <span class="text-xs bg-green-100 text-green-800 rounded-full px-3 py-1">فعال</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">کارت‌های منقضی:</span>
              <label class="relative inline-flex items-center cursor-pointer">
                <input type="checkbox" checked class="sr-only peer">
                <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
              </label>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">خریدهای جدید:</span>
              <label class="relative inline-flex items-center cursor-pointer">
                <input type="checkbox" checked class="sr-only peer">
                <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
              </label>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">خطاهای سیستم:</span>
              <label class="relative inline-flex items-center cursor-pointer">
                <input type="checkbox" checked class="sr-only peer">
                <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
              </label>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">گزارش‌های هفتگی:</span>
              <label class="relative inline-flex items-center cursor-pointer">
                <input type="checkbox" class="sr-only peer">
                <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
              </label>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-green-600 text-white text-sm rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500">
              ذخیره اعلان‌ها
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              تست ارسال
            </button>
          </div>
        </div>

        <!-- اعلان‌های درون برنامه‌ای -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">اعلان‌های درون برنامه‌ای</h3>
            <span class="text-xs bg-blue-100 text-blue-800 rounded-full px-3 py-1">سیستم</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">هشدارهای امنیتی:</span>
              <label class="relative inline-flex items-center cursor-pointer">
                <input type="checkbox" checked class="sr-only peer">
                <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
              </label>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">به‌روزرسانی‌های سیستم:</span>
              <label class="relative inline-flex items-center cursor-pointer">
                <input type="checkbox" checked class="sr-only peer">
                <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
              </label>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">گزارش‌های عملکرد:</span>
              <label class="relative inline-flex items-center cursor-pointer">
                <input type="checkbox" class="sr-only peer">
                <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
              </label>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">هشدارهای عملکرد:</span>
              <label class="relative inline-flex items-center cursor-pointer">
                <input type="checkbox" checked class="sr-only peer">
                <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
              </label>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
              ذخیره تنظیمات
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              مشاهده لاگ
            </button>
          </div>
        </div>
      </div>
    </section>

    <!-- بخش یکپارچه‌سازی -->
    <section>
      <div class="bg-gradient-to-r from-purple-50 to-pink-50 rounded-lg p-6 mb-6">
        <h2 class="text-2xl font-bold text-gray-900 mb-2">🔗 یکپارچه‌سازی</h2>
        <p class="text-gray-600">مدیریت اتصالات و API های خارجی</p>
      </div>
      
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- تنظیمات API -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">تنظیمات API</h3>
            <span class="text-xs bg-purple-100 text-purple-800 rounded-full px-3 py-1">API</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">وضعیت API:</span>
              <span class="text-xs bg-green-100 text-green-800 rounded-full px-3 py-1">فعال</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">محدودیت درخواست:</span>
              <input type="number" value="1000" class="px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 w-24" />
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">کلید API:</span>
              <button class="px-3 py-2 bg-purple-600 text-white text-sm rounded-lg hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-purple-500">
                تولید کلید جدید
              </button>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">Webhook URL:</span>
              <input type="url" placeholder="https://example.com/webhook" class="px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 w-48" />
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-purple-600 text-white text-sm rounded-lg hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-purple-500">
              ذخیره API
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              تست اتصال
            </button>
          </div>
        </div>

        <!-- اتصالات خارجی -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">اتصالات خارجی</h3>
            <span class="text-xs bg-orange-100 text-orange-800 rounded-full px-3 py-1">اتصال</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">پرداخت آنلاین:</span>
              <label class="relative inline-flex items-center cursor-pointer">
                <input type="checkbox" checked class="sr-only peer">
                <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
              </label>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">ارسال پیامک:</span>
              <label class="relative inline-flex items-center cursor-pointer">
                <input type="checkbox" checked class="sr-only peer">
                <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
              </label>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">ارسال ایمیل:</span>
              <label class="relative inline-flex items-center cursor-pointer">
                <input type="checkbox" checked class="sr-only peer">
                <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
              </label>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">CDN:</span>
              <label class="relative inline-flex items-center cursor-pointer">
                <input type="checkbox" class="sr-only peer">
                <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
              </label>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-orange-600 text-white text-sm rounded-lg hover:bg-orange-700 focus:outline-none focus:ring-2 focus:ring-orange-500">
              ذخیره اتصالات
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              تست اتصال
            </button>
          </div>
        </div>
      </div>
    </section>

    <!-- بخش پشتیبان‌گیری -->
    <section>
      <div class="bg-gradient-to-r from-yellow-50 to-orange-50 rounded-lg p-6 mb-6">
        <h2 class="text-2xl font-bold text-gray-900 mb-2">💾 پشتیبان‌گیری</h2>
        <p class="text-gray-600">مدیریت پشتیبان‌گیری و بازیابی داده‌ها</p>
      </div>
      
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- تنظیمات پشتیبان‌گیری -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">تنظیمات پشتیبان‌گیری</h3>
            <span class="text-xs bg-yellow-100 text-yellow-800 rounded-full px-3 py-1">پشتیبان</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">پشتیبان‌گیری خودکار:</span>
              <label class="relative inline-flex items-center cursor-pointer">
                <input type="checkbox" checked class="sr-only peer">
                <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
              </label>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">فرکانس پشتیبان‌گیری:</span>
              <select class="px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
                <option>روزانه</option>
                <option>هفتگی</option>
                <option>ماهانه</option>
              </select>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">نگهداری پشتیبان:</span>
              <select class="px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
                <option>30 روز</option>
                <option>60 روز</option>
                <option>90 روز</option>
                <option>1 سال</option>
              </select>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">آخرین پشتیبان:</span>
              <span class="text-sm text-gray-900">امروز 14:30</span>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-yellow-600 text-white text-sm rounded-lg hover:bg-yellow-700 focus:outline-none focus:ring-2 focus:ring-yellow-500">
              پشتیبان‌گیری دستی
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              تاریخچه
            </button>
          </div>
        </div>

        <!-- بازیابی و نگهداری -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">بازیابی و نگهداری</h3>
            <span class="text-xs bg-green-100 text-green-800 rounded-full px-3 py-1">نگهداری</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">وضعیت سیستم:</span>
              <span class="text-xs bg-green-100 text-green-800 rounded-full px-3 py-1">سالم</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">فضای دیسک:</span>
              <span class="text-sm text-gray-900">75% استفاده شده</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">حافظه سرور:</span>
              <span class="text-sm text-gray-900">60% استفاده شده</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">آخرین بهینه‌سازی:</span>
              <span class="text-sm text-gray-900">دیروز 23:00</span>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-green-600 text-white text-sm rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500">
              بهینه‌سازی سیستم
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              گزارش وضعیت
            </button>
          </div>
        </div>
      </div>
    </section>

    <!-- داشبورد تنظیمات -->
    <section>
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-xl font-semibold text-gray-900 mb-4">📊 وضعیت تنظیمات سیستم</h3>
        
        <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
          <div class="text-center p-6 bg-blue-50 rounded-lg">
            <div class="text-2xl font-bold text-blue-600">{{ systemStats.activeSettings }}</div>
            <div class="text-sm text-gray-600">تنظیمات فعال</div>
          </div>
          
          <div class="text-center p-6 bg-green-50 rounded-lg">
            <div class="text-2xl font-bold text-green-600">{{ systemStats.securityScore }}%</div>
            <div class="text-sm text-gray-600">امتیاز امنیت</div>
          </div>
          
          <div class="text-center p-6 bg-purple-50 rounded-lg">
            <div class="text-2xl font-bold text-purple-600">{{ systemStats.uptime }}%</div>
            <div class="text-sm text-gray-600">زمان فعالیت</div>
          </div>
          
          <div class="text-center p-6 bg-orange-50 rounded-lg">
            <div class="text-2xl font-bold text-orange-600">{{ systemStats.lastBackup }}</div>
            <div class="text-sm text-gray-600">آخرین پشتیبان</div>
          </div>
        </div>
        
        <!-- نمودار وضعیت سیستم -->
        <div class="mt-6">
          <h4 class="text-sm font-medium text-gray-900 mb-3">وضعیت سیستم (24 ساعت گذشته)</h4>
          <div class="flex items-end space-x-2 space-x-reverse h-32 overflow-x-auto">
            <div v-for="(hour, index) in systemStatus" :key="index" class="flex-shrink-0 flex flex-col items-center min-w-16">
              <div
class="w-full bg-gray-200 rounded-t relative"
                   :style="{ height: getChartHeight(hour.performance) + 'px' }">
                <div
class="w-full bg-gradient-to-t from-green-500 to-blue-500 rounded-t transition-all duration-300 absolute bottom-0"
                     :style="{ height: getChartHeight(hour.performance) + 'px' }"></div>
              </div>
              <span class="text-xs text-gray-500 mt-1 text-center">{{ hour.time }}</span>
              <span class="text-xs text-gray-400 mt-1 text-center">{{ Math.round(hour.performance * 100) }}%</span>
            </div>
          </div>
        </div>
        
        <!-- دکمه‌های عملیات -->
        <div class="mt-6 flex space-x-2 space-x-reverse">
          <button class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
            ذخیره همه تنظیمات
          </button>
          <button class="px-4 py-2 bg-green-600 text-white text-sm rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500">
            بازنشانی به پیش‌فرض
          </button>
          <button class="px-4 py-2 bg-purple-600 text-white text-sm rounded-lg hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-purple-500">
            خروجی تنظیمات
          </button>
          <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
            راهنما
          </button>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
// آمار سیستم
const systemStats = {
  activeSettings: 24,
  securityScore: 95,
  uptime: 99.9,
  lastBackup: '2 ساعت پیش'
}

// وضعیت سیستم
const systemStatus = [
  { time: '00:00', performance: 0.95 },
  { time: '04:00', performance: 0.92 },
  { time: '08:00', performance: 0.98 },
  { time: '12:00', performance: 0.99 },
  { time: '16:00', performance: 0.97 },
  { time: '20:00', performance: 0.96 },
  { time: '24:00', performance: 0.94 }
]

// تابع محاسبه ارتفاع مناسب برای نمودار
const getChartHeight = (performance: number) => {
  // حداکثر ارتفاع نمودار را تعیین کنید
  const maxHeight = 80; // به عنوان مثال، حداکثر 80 پیکسل
  // ارتفاع فعلی را بر اساس عملکرد محاسبه کنید
  return maxHeight * performance;
}
</script>

<!--
  مستندسازی:
  این کامپوننت شامل پنج بخش اصلی است:
  1. تنظیمات عمومی: تنظیمات پایه، تنظیمات کارت
  2. تنظیمات امنیت: امنیت احراز هویت، حریم خصوصی
  3. تنظیمات اعلان‌ها: اعلان‌های ایمیل، اعلان‌های درون برنامه‌ای
  4. یکپارچه‌سازی: تنظیمات API، اتصالات خارجی
  5. پشتیبان‌گیری: تنظیمات پشتیبان‌گیری، بازیابی و نگهداری
  
  تمام توضیحات به فارسی و با طراحی ریسپانسیو ارائه شده است.
  نمودارها و آمار دقیق برای نمایش وضعیت سیستم استفاده شده‌اند.
-->


<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 

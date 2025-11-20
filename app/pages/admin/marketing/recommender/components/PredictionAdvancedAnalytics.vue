<template>
  <div class="space-y-8">
    <!-- هدر بخش -->
    <div class="flex justify-between items-center">
      <div>
        <h2 class="text-2xl font-bold text-gray-900">پیش‌بینی و تحلیل پیشرفته</h2>
        <p class="mt-1 text-sm text-gray-600">پیش‌بینی تقاضا، قیمت، موجودی و تحلیل‌های پیشرفته</p>
      </div>
      <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors">
        <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"></path>
        </svg>
        اجرای پیش‌بینی
      </button>
    </div>

    <!-- آمار پیش‌بینی -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-6">
      <div class="bg-gradient-to-r from-blue-500 to-blue-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">دقت پیش‌بینی</p>
            <p class="text-2xl font-bold">94.2%</p>
            <p class="text-xs opacity-75">میانگین کلی</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-green-500 to-green-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">پیش‌بینی تقاضا</p>
            <p class="text-2xl font-bold">87.5%</p>
            <p class="text-xs opacity-75">دقت</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-purple-500 to-purple-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">پیش‌بینی قیمت</p>
            <p class="text-2xl font-bold">91.3%</p>
            <p class="text-xs opacity-75">دقت</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-yellow-500 to-yellow-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">پیش‌بینی موجودی</p>
            <p class="text-2xl font-bold">89.7%</p>
            <p class="text-xs opacity-75">دقت</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-red-500 to-red-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">پیش‌بینی ریزش</p>
            <p class="text-2xl font-bold">92.8%</p>
            <p class="text-xs opacity-75">دقت</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 17h8m0 0V9m0 8l-8-8-4 4-6-6"></path>
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- پیش‌بینی تقاضا -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-green-100 rounded-lg">
              <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">پیش‌بینی تقاضا</h3>
          </div>
          <button class="bg-green-600 text-white px-3 py-1 rounded-lg text-sm hover:bg-green-700 transition-colors">
            اجرا
          </button>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">بازه زمانی</label>
            <select v-model="demandPrediction.timeframe" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500">
              <option value="weekly">هفتگی</option>
              <option value="monthly">ماهانه</option>
              <option value="quarterly">فصلی</option>
              <option value="yearly">سالانه</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">دسته‌بندی محصولات</label>
            <select v-model="demandPrediction.category" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500">
              <option value="all">همه دسته‌بندی‌ها</option>
              <option value="electronics">الکترونیک</option>
              <option value="clothing">پوشاک</option>
              <option value="books">کتاب</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">فاکتورهای تاثیرگذار</label>
            <div class="space-y-2">
              <label class="flex items-center">
                <input v-model="demandPrediction.factors" type="checkbox" value="seasonal" class="rounded border-gray-300 text-green-600 focus:ring-green-500">
                <span class="mr-2 text-sm text-gray-700">فصلی</span>
              </label>
              <label class="flex items-center">
                <input v-model="demandPrediction.factors" type="checkbox" value="trending" class="rounded border-gray-300 text-green-600 focus:ring-green-500">
                <span class="mr-2 text-sm text-gray-700">روند</span>
              </label>
              <label class="flex items-center">
                <input v-model="demandPrediction.factors" type="checkbox" value="promotional" class="rounded border-gray-300 text-green-600 focus:ring-green-500">
                <span class="mr-2 text-sm text-gray-700">تبلیغاتی</span>
              </label>
            </div>
          </div>
          
          <div class="bg-gray-50 rounded-lg p-6">
            <h4 class="font-medium text-gray-900 mb-2">نتایج پیش‌بینی</h4>
            <div class="space-y-2 text-sm">
              <div class="flex justify-between">
                <span>محصولات پرفروش آینده:</span>
                <span class="font-medium">15 محصول</span>
              </div>
              <div class="flex justify-between">
                <span>افزایش تقاضا:</span>
                <span class="font-medium text-green-600">+23%</span>
              </div>
              <div class="flex justify-between">
                <span>دقت پیش‌بینی:</span>
                <span class="font-medium">87.5%</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- پیش‌بینی قیمت -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-purple-100 rounded-lg">
              <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">پیش‌بینی قیمت</h3>
          </div>
          <button class="bg-purple-600 text-white px-3 py-1 rounded-lg text-sm hover:bg-purple-700 transition-colors">
            اجرا
          </button>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع پیش‌بینی</label>
            <select v-model="pricePrediction.type" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500">
              <option value="increase">افزایش قیمت</option>
              <option value="decrease">کاهش قیمت</option>
              <option value="stable">ثبات قیمت</option>
              <option value="volatile">نوسان قیمت</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">بازه زمانی</label>
            <select v-model="pricePrediction.timeframe" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500">
              <option value="1_month">1 ماه</option>
              <option value="3_months">3 ماه</option>
              <option value="6_months">6 ماه</option>
              <option value="1_year">1 سال</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">فاکتورهای قیمت</label>
            <div class="space-y-2">
              <label class="flex items-center">
                <input v-model="pricePrediction.factors" type="checkbox" value="inflation" class="rounded border-gray-300 text-purple-600 focus:ring-purple-500">
                <span class="mr-2 text-sm text-gray-700">تورم</span>
              </label>
              <label class="flex items-center">
                <input v-model="pricePrediction.factors" type="checkbox" value="demand" class="rounded border-gray-300 text-purple-600 focus:ring-purple-500">
                <span class="mr-2 text-sm text-gray-700">تقاضا</span>
              </label>
              <label class="flex items-center">
                <input v-model="pricePrediction.factors" type="checkbox" value="competition" class="rounded border-gray-300 text-purple-600 focus:ring-purple-500">
                <span class="mr-2 text-sm text-gray-700">رقابت</span>
              </label>
            </div>
          </div>
          
          <div class="bg-gray-50 rounded-lg p-6">
            <h4 class="font-medium text-gray-900 mb-2">نتایج پیش‌بینی</h4>
            <div class="space-y-2 text-sm">
              <div class="flex justify-between">
                <span>تغییر قیمت پیش‌بینی:</span>
                <span class="font-medium text-red-600">+8.5%</span>
              </div>
              <div class="flex justify-between">
                <span>محصولات تحت تاثیر:</span>
                <span class="font-medium">234 محصول</span>
              </div>
              <div class="flex justify-between">
                <span>دقت پیش‌بینی:</span>
                <span class="font-medium">91.3%</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- پیش‌بینی موجودی و ریزش -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <!-- پیش‌بینی موجودی -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-yellow-100 rounded-lg">
              <svg class="w-6 h-6 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">پیش‌بینی موجودی</h3>
          </div>
          <button class="bg-yellow-600 text-white px-3 py-1 rounded-lg text-sm hover:bg-yellow-700 transition-colors">
            اجرا
          </button>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">آستانه هشدار</label>
            <input v-model="inventoryPrediction.threshold" type="number" min="1" max="100" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500">
            <p class="text-xs text-gray-500 mt-1">درصد موجودی</p>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">بازه زمانی</label>
            <select v-model="inventoryPrediction.timeframe" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500">
              <option value="1_week">1 هفته</option>
              <option value="2_weeks">2 هفته</option>
              <option value="1_month">1 ماه</option>
              <option value="3_months">3 ماه</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع محصولات</label>
            <div class="space-y-2">
              <label class="flex items-center">
                <input v-model="inventoryPrediction.productTypes" type="checkbox" value="fast_moving" class="rounded border-gray-300 text-yellow-600 focus:ring-yellow-500">
                <span class="mr-2 text-sm text-gray-700">پرفروش</span>
              </label>
              <label class="flex items-center">
                <input v-model="inventoryPrediction.productTypes" type="checkbox" value="seasonal" class="rounded border-gray-300 text-yellow-600 focus:ring-yellow-500">
                <span class="mr-2 text-sm text-gray-700">فصلی</span>
              </label>
              <label class="flex items-center">
                <input v-model="inventoryPrediction.productTypes" type="checkbox" value="all" class="rounded border-gray-300 text-yellow-600 focus:ring-yellow-500">
                <span class="mr-2 text-sm text-gray-700">همه</span>
              </label>
            </div>
          </div>
          
          <div class="bg-gray-50 rounded-lg p-6">
            <h4 class="font-medium text-gray-900 mb-2">نتایج پیش‌بینی</h4>
            <div class="space-y-2 text-sm">
              <div class="flex justify-between">
                <span>محصولات نیازمند موجودی:</span>
                <span class="font-medium text-red-600">45 محصول</span>
              </div>
              <div class="flex justify-between">
                <span>زمان تخلیه موجودی:</span>
                <span class="font-medium">2.3 هفته</span>
              </div>
              <div class="flex justify-between">
                <span>دقت پیش‌بینی:</span>
                <span class="font-medium">89.7%</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- پیش‌بینی ریزش -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-red-100 rounded-lg">
              <svg class="w-6 h-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 17h8m0 0V9m0 8l-8-8-4 4-6-6"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">پیش‌بینی ریزش</h3>
          </div>
          <button class="bg-red-600 text-white px-3 py-1 rounded-lg text-sm hover:bg-red-700 transition-colors">
            اجرا
          </button>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">آستانه ریزش</label>
            <input v-model="churnPrediction.threshold" type="number" min="1" max="100" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500">
            <p class="text-xs text-gray-500 mt-1">درصد احتمال ریزش</p>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">بازه زمانی</label>
            <select v-model="churnPrediction.timeframe" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500">
              <option value="1_month">1 ماه</option>
              <option value="3_months">3 ماه</option>
              <option value="6_months">6 ماه</option>
              <option value="1_year">1 سال</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">فاکتورهای ریزش</label>
            <div class="space-y-2">
              <label class="flex items-center">
                <input v-model="churnPrediction.factors" type="checkbox" value="inactivity" class="rounded border-gray-300 text-red-600 focus:ring-red-500">
                <span class="mr-2 text-sm text-gray-700">عدم فعالیت</span>
              </label>
              <label class="flex items-center">
                <input v-model="churnPrediction.factors" type="checkbox" value="complaints" class="rounded border-gray-300 text-red-600 focus:ring-red-500">
                <span class="mr-2 text-sm text-gray-700">شکایات</span>
              </label>
              <label class="flex items-center">
                <input v-model="churnPrediction.factors" type="checkbox" value="price_sensitivity" class="rounded border-gray-300 text-red-600 focus:ring-red-500">
                <span class="mr-2 text-sm text-gray-700">حساسیت قیمت</span>
              </label>
            </div>
          </div>
          
          <div class="bg-gray-50 rounded-lg p-6">
            <h4 class="font-medium text-gray-900 mb-2">نتایج پیش‌بینی</h4>
            <div class="space-y-2 text-sm">
              <div class="flex justify-between">
                <span>کاربران در معرض ریزش:</span>
                <span class="font-medium text-red-600">234 کاربر</span>
              </div>
              <div class="flex justify-between">
                <span>احتمال ریزش:</span>
                <span class="font-medium">12.3%</span>
              </div>
              <div class="flex justify-between">
                <span>دقت پیش‌بینی:</span>
                <span class="font-medium">92.8%</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- پیش‌بینی ارزش عمر مشتری -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <div class="flex items-center">
          <div class="p-2 bg-indigo-100 rounded-lg">
            <svg class="w-6 h-6 text-indigo-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"></path>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-900 mr-3">پیش‌بینی ارزش عمر مشتری</h3>
        </div>
        <button class="bg-indigo-600 text-white px-4 py-2 rounded-lg hover:bg-indigo-700 transition-colors">
          محاسبه CLV
        </button>
      </div>
      
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">بازه محاسبه</label>
            <select v-model="clvPrediction.timeframe" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500">
              <option value="1_year">1 سال</option>
              <option value="3_years">3 سال</option>
              <option value="5_years">5 سال</option>
              <option value="lifetime">کل عمر</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع مشتری</label>
            <select v-model="clvPrediction.customerType" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500">
              <option value="new">مشتریان جدید</option>
              <option value="regular">مشتریان عادی</option>
              <option value="premium">مشتریان ویژه</option>
              <option value="all">همه مشتریان</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">فاکتورهای محاسبه</label>
            <div class="space-y-2">
              <label class="flex items-center">
                <input v-model="clvPrediction.factors" type="checkbox" value="purchase_frequency" class="rounded border-gray-300 text-indigo-600 focus:ring-indigo-500">
                <span class="mr-2 text-sm text-gray-700">فرکانس خرید</span>
              </label>
              <label class="flex items-center">
                <input v-model="clvPrediction.factors" type="checkbox" value="average_order" class="rounded border-gray-300 text-indigo-600 focus:ring-indigo-500">
                <span class="mr-2 text-sm text-gray-700">میانگین سفارش</span>
              </label>
              <label class="flex items-center">
                <input v-model="clvPrediction.factors" type="checkbox" value="retention_rate" class="rounded border-gray-300 text-indigo-600 focus:ring-indigo-500">
                <span class="mr-2 text-sm text-gray-700">نرخ نگهداری</span>
              </label>
            </div>
          </div>
        </div>
        
        <div class="space-y-4">
          <div class="bg-gradient-to-r from-indigo-500 to-indigo-600 rounded-lg p-6 text-white">
            <h4 class="font-medium mb-4">نتایج CLV</h4>
            <div class="space-y-3">
              <div class="flex justify-between">
                <span>میانگین CLV:</span>
                <span class="font-bold">$1,234</span>
              </div>
              <div class="flex justify-between">
                <span>CLV بالا:</span>
                <span class="font-bold">$5,678</span>
              </div>
              <div class="flex justify-between">
                <span>CLV پایین:</span>
                <span class="font-bold">$234</span>
              </div>
              <div class="flex justify-between">
                <span>دقت پیش‌بینی:</span>
                <span class="font-bold">94.2%</span>
              </div>
            </div>
          </div>
          
          <div class="bg-gray-50 rounded-lg p-6">
            <h4 class="font-medium text-gray-900 mb-2">توصیه‌های بهینه‌سازی</h4>
            <div class="space-y-2 text-sm text-gray-600">
              <p>• افزایش نرخ نگهداری مشتریان جدید</p>
              <p>• بهبود تجربه خرید مشتریان ویژه</p>
              <p>• ارائه تخفیف‌های شخصی‌سازی شده</p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- گزارش‌های پیش‌بینی -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-semibold text-gray-900">گزارش‌های پیش‌بینی</h3>
      </div>
      <div class="p-6">
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع پیش‌بینی</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">آخرین اجرا</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">دقت</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="report in predictionReports" :key="report.id">
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="flex items-center">
                    <div class="w-8 h-8 rounded-lg mr-3 flex items-center justify-center" :class="report.bgColor">
                      <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="report.icon"></path>
                      </svg>
                    </div>
                    <div>
                      <div class="text-sm font-medium text-gray-900">{{ report.name }}</div>
                      <div class="text-sm text-gray-500">{{ report.description }}</div>
                    </div>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  {{ report.lastRun }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  {{ report.accuracy }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full" :class="report.statusClass">
                    {{ report.status }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                  <button class="text-indigo-600 hover:text-indigo-900 mr-3">مشاهده</button>
                  <button class="text-green-600 hover:text-green-900 mr-3">اجرا</button>
                  <button class="text-red-600 hover:text-red-900">حذف</button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
// پیش‌بینی تقاضا
const demandPrediction = ref({
  timeframe: 'monthly',
  category: 'all',
  factors: ['seasonal', 'trending']
})

// پیش‌بینی قیمت
const pricePrediction = ref({
  type: 'increase',
  timeframe: '3_months',
  factors: ['inflation', 'demand']
})

// پیش‌بینی موجودی
const inventoryPrediction = ref({
  threshold: 20,
  timeframe: '1_month',
  productTypes: ['fast_moving']
})

// پیش‌بینی ریزش
const churnPrediction = ref({
  threshold: 70,
  timeframe: '3_months',
  factors: ['inactivity', 'complaints']
})

// پیش‌بینی ارزش عمر مشتری
const clvPrediction = ref({
  timeframe: '3_years',
  customerType: 'all',
  factors: ['purchase_frequency', 'average_order', 'retention_rate']
})

// گزارش‌های پیش‌بینی
const predictionReports = ref([
  {
    id: 1,
    name: 'پیش‌بینی تقاضای ماهانه',
    description: 'پیش‌بینی تقاضای محصولات در ماه آینده',
    lastRun: '2 روز پیش',
    accuracy: '87.5%',
    status: 'تکمیل شده',
    statusClass: 'bg-green-100 text-green-800',
    bgColor: 'bg-green-500',
    icon: 'M13 7h8m0 0v8m0-8l-8 8-4-4-6 6'
  },
  {
    id: 2,
    name: 'پیش‌بینی قیمت فصلی',
    description: 'پیش‌بینی تغییرات قیمت در فصل آینده',
    lastRun: '1 هفته پیش',
    accuracy: '91.3%',
    status: 'تکمیل شده',
    statusClass: 'bg-green-100 text-green-800',
    bgColor: 'bg-purple-500',
    icon: 'M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1'
  },
  {
    id: 3,
    name: 'پیش‌بینی موجودی هفتگی',
    description: 'پیش‌بینی نیاز به موجودی در هفته آینده',
    lastRun: '1 روز پیش',
    accuracy: '89.7%',
    status: 'در حال اجرا',
    statusClass: 'bg-yellow-100 text-yellow-800',
    bgColor: 'bg-yellow-500',
    icon: 'M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4'
  },
  {
    id: 4,
    name: 'پیش‌بینی ریزش مشتریان',
    description: 'شناسایی مشتریان در معرض ریزش',
    lastRun: '3 روز پیش',
    accuracy: '92.8%',
    status: 'تکمیل شده',
    statusClass: 'bg-green-100 text-green-800',
    bgColor: 'bg-red-500',
    icon: 'M13 17h8m0 0V9m0 8l-8-8-4 4-6-6'
  },
  {
    id: 5,
    name: 'محاسبه CLV سالانه',
    description: 'محاسبه ارزش عمر مشتری برای سال آینده',
    lastRun: '1 هفته پیش',
    accuracy: '94.2%',
    status: 'تکمیل شده',
    statusClass: 'bg-green-100 text-green-800',
    bgColor: 'bg-indigo-500',
    icon: 'M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z'
  }
])
</script> 

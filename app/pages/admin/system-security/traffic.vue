<template>
  <div class="p-6" dir="rtl">
    <div class="mb-6 bg-white p-6 rounded-lg shadow-sm border border-gray-200">
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-4">
          <h1 class="text-2xl font-bold text-gray-900 mb-2">ترافیک آنلاین</h1>
          <p class="text-gray-600">نظارت بر ترافیک ورودی و خروجی سیستم</p>
        </div>
        <div class="flex items-center gap-4">
          <!-- Toggle: Traffic logging enabled -->
          <label class="flex items-center cursor-pointer select-none">
            <span class="ml-3 text-sm text-gray-600">ثبت لاگ ترافیک</span>
            <input type="checkbox" class="sr-only" :checked="trafficLoggingEnabled" @change="toggleTrafficLogging">
            <div :class="['w-12 h-7 rounded-full p-1 transition', trafficLoggingEnabled ? 'bg-green-500' : 'bg-gray-300']">
              <div :class="['w-5 h-5 bg-white rounded-full transition', trafficLoggingEnabled ? 'translate-x-5' : 'translate-x-0']"></div>
            </div>
          </label>
        
        <button 
          class="px-4 py-2 text-green-600 bg-green-50 border-2 border-green-200 rounded-lg hover:bg-green-100 hover:border-green-300 transition-all duration-200 flex items-center shadow-sm hover:shadow-md" 
          @click="exportFilteredData"
        >
          <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
          </svg>
          خروجی اکسل
        </button>
        </div>
      </div>
    </div>

    <!-- آمار کلی -->
    <div class="relative">
      <div class="relative grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-4 mb-6" style="z-index:1;">
        <TemplateCard
          title="کاربران آنلاین"
          :value="trafficStats.online_users || 0"
          variant="cyan"
        >
          <template #icon>
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197m13.5-9a2.5 2.5 0 11-5 0 2.5 2.5 0 015 0z"></path>
            </svg>
          </template>
        </TemplateCard>
        <TemplateCard
          title="درخواست‌های امروز"
          :value="trafficStats.today_requests || 0"
          variant="green"
        >
          <template #icon>
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
            </svg>
          </template>
        </TemplateCard>
        <TemplateCard
          title="درخواست‌های مشکوک"
          :value="trafficStats.suspicious_requests || 0"
          variant="yellow"
        >
          <template #icon>
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
            </svg>
          </template>
        </TemplateCard>
        <TemplateCard
          title="حمله‌های مسدود شده"
          :value="trafficStats.blocked_attacks || 0"
          variant="red"
        >
          <template #icon>
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"></path>
            </svg>
          </template>
        </TemplateCard>
        <TemplateCard
          title="IP های مسدود شده"
          :value="blockedIPsCount"
          variant="purple"
        >
          <template #icon>
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728L5.636 5.636m12.728 12.728L18.364 5.636M5.636 18.364l12.728-12.728"></path>
            </svg>
          </template>
        </TemplateCard>
      </div>
    </div>

    <!-- فیلترها و جستجوی پیشرفته -->
    <div class="bg-gradient-to-br from-white to-gray-50 p-8 rounded-2xl shadow-lg border border-gray-100 mb-8">
            <!-- هدر بخش فیلترها -->
      <div class="flex items-center justify-between mb-6 bg-gradient-to-l from-purple-100 via-white to-purple-50 rounded-xl p-4 shadow-sm">
        <div class="flex items-center gap-6">
          <div class="w-12 h-12 bg-gradient-to-br from-purple-500 to-purple-600 rounded-xl flex items-center justify-center mr-4 shadow-lg">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.207A1 1 0 013 6.5V4z"></path>
            </svg>
          </div>
          <div>
            <h3 class="text-xl font-bold text-gray-900">فیلترها و جستجوی پیشرفته</h3>
          </div>
        </div>
        <!-- حذف دایره سبز و کلمه فعال -->
      </div>
      
      <!-- فیلدهای فیلتر -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-6 gap-6">
        <!-- نوع آگهی -->
        <div class="group">
          <label class="block text-sm font-semibold text-gray-700 mb-3 group-hover:text-purple-600 transition-colors">
            <svg class="w-4 h-4 inline mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
            نوع آگهی
          </label>
          <select 
            v-model="filters.adType" 
            class="w-full px-4 py-3 border-2 border-gray-200 rounded-xl focus:outline-none focus:ring-4 focus:ring-purple-100 focus:border-purple-500 transition-all duration-200 bg-white shadow-sm hover:shadow-md"
            @change="applyFilters"
          >
            <option value="">همه انواع</option>
            <option value="ربات">🤖 ربات</option>
            <option value="انسان">👤 انسان</option>
          </select>
        </div>

        <!-- IP آدرس -->
        <div class="group">
          <label class="block text-sm font-semibold text-gray-700 mb-3 group-hover:text-purple-600 transition-colors">
            <svg class="w-4 h-4 inline mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 12a9 9 0 01-9 9m9-9a9 9 0 00-9-9m9 9H3m9 9v-9m0-9v9"></path>
            </svg>
            IP آدرس
          </label>
          <input 
            v-model="filters.ipAddress" 
            type="text"
            placeholder="جستجو IP..." 
            class="w-full px-4 py-3 border-2 border-gray-200 rounded-xl focus:outline-none focus:ring-4 focus:ring-purple-100 focus:border-purple-500 transition-all duration-200 bg-white shadow-sm hover:shadow-md"
            @input="applyFilters"
          />
        </div>

        <!-- موقعیت -->
        <div class="group">
          <label class="block text-sm font-semibold text-gray-700 mb-3 group-hover:text-purple-600 transition-colors">
            <svg class="w-4 h-4 inline mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"></path>
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"></path>
            </svg>
            موقعیت
          </label>
          <select 
            v-model="filters.location" 
            class="w-full px-4 py-3 border-2 border-gray-200 rounded-xl focus:outline-none focus:ring-4 focus:ring-purple-100 focus:border-purple-500 transition-all duration-200 bg-white shadow-sm hover:shadow-md"
            @change="applyFilters"
          >
            <option value="">همه شهرها</option>
            <option value="تهران">🏛️ تهران</option>
            <option value="اصفهان">🕌 اصفهان</option>
            <option value="مشهد">🕍 مشهد</option>
            <option value="شیراز">🌹 شیراز</option>
            <option value="تبریز">🏔️ تبریز</option>
            <option value="یزد">🏺 یزد</option>
            <option value="کرج">🏢 کرج</option>
            <option value="قم">🕌 قم</option>
          </select>
        </div>

        <!-- صفحه -->
        <div class="group">
          <label class="block text-sm font-semibold text-gray-700 mb-3 group-hover:text-purple-600 transition-colors">
            <svg class="w-4 h-4 inline mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
            </svg>
            صفحه
          </label>
          <input 
            v-model="filters.page" 
            type="text"
            placeholder="جستجو صفحه..." 
            class="w-full px-4 py-3 border-2 border-gray-200 rounded-xl focus:outline-none focus:ring-4 focus:ring-purple-100 focus:border-purple-500 transition-all duration-200 bg-white shadow-sm hover:shadow-md"
            @input="applyFilters"
          />
        </div>

        <!-- وضعیت -->
        <div class="group">
          <label class="block text-sm font-semibold text-gray-700 mb-3 group-hover:text-purple-600 transition-colors">
            <svg class="w-4 h-4 inline mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
            وضعیت
          </label>
          <select 
            v-model="filters.status" 
            class="w-full px-4 py-3 border-2 border-gray-200 rounded-xl focus:outline-none focus:ring-4 focus:ring-purple-100 focus:border-purple-500 transition-all duration-200 bg-white shadow-sm hover:shadow-md"
            @change="applyFilters"
          >
            <option value="">همه وضعیت‌ها</option>
            <option value="200">✅ موفق (200)</option>
            <option value="401">🔒 خطای احراز هویت (401)</option>
            <option value="404">❌ یافت نشد (404)</option>
            <option value="500">⚠️ خطای سرور (500)</option>
          </select>
        </div>

        <!-- بازه زمانی -->
        <div class="group">
          <label class="block text-sm font-semibold text-gray-700 mb-3 group-hover:text-purple-600 transition-colors">
            <svg class="w-4 h-4 inline mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
            بازه زمانی
          </label>
          <select 
            v-model="filters.timeRange" 
            class="w-full px-4 py-3 border-2 border-gray-200 rounded-xl focus:outline-none focus:ring-4 focus:ring-purple-100 focus:border-purple-500 transition-all duration-200 bg-white shadow-sm hover:shadow-md"
            @change="applyFilters"
          >
            <option value="">همه زمان‌ها</option>
            <option value="1h">⏰ آخرین 1 ساعت</option>
            <option value="6h">⏰ آخرین 6 ساعت</option>
            <option value="24h">📅 آخرین 24 ساعت</option>
            <option value="7d">📆 آخرین 7 روز</option>
            <option value="30d">🗓️ آخرین 30 روز</option>
          </select>
        </div>
      </div>

            <!-- دکمه‌های عملیات -->
      <div class="flex items-center justify-end mt-8 pt-6 border-t border-gray-200">
        <button 
          v-if="filters.adType || filters.ipAddress || filters.location || filters.page || filters.status || filters.timeRange"
          class="px-6 py-3 text-gray-600 bg-gray-100 border-2 border-gray-200 rounded-xl hover:bg-gray-200 hover:border-gray-300 transition-all duration-200 flex items-center shadow-sm hover:shadow-md" 
          @click="clearFilters"
        >
          <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
          پاک کردن فیلتر
        </button>
      </div>
    </div>

    <!-- جدول ترافیک -->
    <div class="bg-white p-6 rounded-lg shadow-md border border-gray-200 mb-8">
      <div class="flex items-center justify-between mb-4">
        <h3 class="text-lg font-semibold text-gray-900">جدول ترافیک</h3>
        <div class="text-sm text-gray-500">
          {{ generalTrafficStats.traffic_logs?.length || 0 }} رکورد
        </div>
      </div>
      
      <!-- جدول داده‌ها -->
      <div class="overflow-x-auto overflow-y-auto custom-scrollbar max-h-[50vh] md:max-h-[65vh]">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کاربر/IP</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع آگهی</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">موقعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">صفحه</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">زمان</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مرورگر</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع دستگاه</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">سیستم عامل</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نام هاست</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">پاسخ</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-if="isLoading">
              <td colspan="11" class="px-6 py-8 text-center text-gray-500">
                <div class="flex items-center justify-center">
                  <div class="animate-spin rounded-full h-6 w-6 border-b-2 border-blue-500"></div>
                  <span class="mr-2">در حال بارگذاری...</span>
                </div>
              </td>
            </tr>
            <tr v-else-if="!isLoading && getTableData().length === 0">
              <td colspan="11" class="px-6 py-8 text-center text-gray-500">
                هیچ داده‌ای یافت نشد
              </td>
            </tr>
            <tr v-for="item in getTableData()" v-else :key="item.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ item.ipAddress || '-' }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                <div class="flex items-center justify-center">
                  <div 
                    class="w-4 h-4 rounded-full"
                    :class="{
                      'bg-green-500': item.adType === 'بنیر',
                      'bg-gray-500': item.adType === 'اسلایدر',
                      'bg-yellow-500': item.adType === 'پاپ‌آپ',
                      'bg-red-500': item.adType === 'مشکوک'
                    }"
                    :title="item.adType || '-'"
                  ></div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ item.location || '-' }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ item.page || '-' }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ item.time || '-' }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ item.browser || '-' }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ item.deviceType || '-' }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ item.os || '-' }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ item.hostname || '-' }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ item.response || '-' }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                <div class="flex items-center space-x-2 space-x-reverse">
                  <button 
                    class="text-blue-600 hover:text-blue-900 transition-colors" 
                    title="مشاهده جزئیات"
                    @click="showTrafficDetails(item)"
                  >
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                    </svg>
                  </button>
                  <button 
                    class="text-red-600 hover:text-red-900 transition-colors" 
                    title="مسدود کردن IP"
                    @click="blockIP(item.ipAddress)"
                  >
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"></path>
                    </svg>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <!-- خط جداکننده و کنترل نمایش در پایین جدول ترافیک -->
      <hr class="my-2 border-t border-gray-200" />
      <div class="flex items-center justify-end mt-2">
        <div class="flex items-center space-x-2 space-x-reverse">
          <span class="text-sm text-gray-700">نمایش:</span>
          <select 
            v-model="trafficItemsPerPage" 
            class="border border-gray-300 rounded-md px-2 py-1 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
            @change="handleTrafficItemsPerPageChange"
          >
            <option value="10">10</option>
            <option value="25">25</option>
            <option value="50">50</option>
            <option value="100">100</option>
          </select>
        </div>
      </div>


    </div>

    <!-- جدول کاربران آنلاین -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200 flex items-center justify-between">
        <h3 class="text-lg font-semibold text-gray-900">کاربران آنلاین</h3>
        <div class="text-sm text-gray-500">
          {{ onlineUsers.length }} کاربر فعال
        </div>
      </div>
      <div class="overflow-x-auto overflow-y-auto custom-scrollbar max-h-[50vh] md:max-h-[65vh]">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کاربر/IP</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع آگهی</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">موقعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">صفحه</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">زمان</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مرورگر</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع دستگاه</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">سیستم عامل</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نام هاست</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">پاسخ</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-if="isOnlineUsersLoading">
              <td colspan="11" class="px-6 py-8 text-center text-gray-500">
                <div class="flex items-center justify-center">
                  <div class="animate-spin rounded-full h-6 w-6 border-b-2 border-blue-500"></div>
                  <span class="mr-2">در حال بارگذاری...</span>
                </div>
              </td>
            </tr>
            <tr v-else-if="onlineUsers.length === 0">
              <td colspan="11" class="px-6 py-8 text-center text-gray-500">
                هیچ کاربر آنلاینی یافت نشد
              </td>
            </tr>
            <tr v-for="user in onlineUsers" v-else :key="user.id || Math.random()">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="flex-shrink-0 h-8 w-8">
                    <div class="h-8 w-8 rounded-full bg-gray-300 flex items-center justify-center">
                      <span class="text-sm font-medium text-gray-700">{{ user.name?.charAt(0) || '?' }}</span>
                    </div>
                  </div>
                  <div class="mr-3">
                    <div class="text-sm font-medium text-gray-900">{{ user.name || 'نامشخص' }}</div>
                    <div class="text-sm text-gray-500">{{ user.ip || 'نامشخص' }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                <div class="flex items-center justify-center">
                  <template v-if="(user as any).adType">
                    <div 
                      class="w-4 h-4 rounded-full"
                      :class="{
                        'bg-green-500': (user as any).adType === 'بنیر',
                        'bg-gray-500': (user as any).adType === 'اسلایدر',
                        'bg-yellow-500': (user as any).adType === 'پاپ‌آپ',
                        'bg-red-500': (user as any).adType === 'مشکوک'
                      }"
                      :title="(user as any).adType"
                    ></div>
                  </template>
                  <template v-else>
                    <span>-</span>
                  </template>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ (user as any).location || '-' }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ (user as any).currentPage || '/' }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ (user as any).lastSeen }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ (user as any).browser || '-' }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ (user as any).deviceType || '-' }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ (user as any).os || '-' }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ (user as any).hostname || '-' }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ (user as any).response || '-' }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <div class="flex items-center space-x-2 space-x-reverse">
                  <button 
                    class="text-blue-600 hover:text-blue-900 transition-colors" 
                    title="مشاهده جزئیات کاربر"
                    @click="showUserDetails(user)"
                  >
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                    </svg>
                  </button>
                  <button 
                    class="text-red-600 hover:text-red-900 transition-colors" 
                    title="مسدود کردن IP"
                    @click="blockIP((user as any).ip)"
                  >
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"></path>
                    </svg>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <!-- خط جداکننده و کنترل نمایش در پایین جدول کاربران آنلاین -->
      <hr class="my-2 border-t border-gray-200" />
      <div class="flex items-center justify-end mt-2 mb-4 px-6">
        <div class="flex items-center space-x-2 space-x-reverse">
          <span class="text-sm text-gray-700">نمایش:</span>
          <select 
            v-model="onlineUsersItemsPerPage" 
            class="border border-gray-300 rounded-md px-2 py-1 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
            @change="handleOnlineUsersItemsPerPageChange"
          >
            <option value="10">10</option>
            <option value="25">25</option>
            <option value="50">50</option>
            <option value="100">100</option>
          </select>
        </div>
      </div>

    </div>

    <!-- مودال مسدود کردن IP -->
    <div v-if="showBlockModal" class="fixed inset-0 flex items-center justify-center z-50">
      <div 
        ref="blockModalRef"
        class="bg-white rounded-2xl p-8 w-[500px] max-w-full mx-4 shadow-2xl border-2 border-purple-500"
      >
        <div class="flex items-center mb-6">
          <div class="w-12 h-12 bg-gradient-to-br from-red-500 to-red-600 rounded-xl flex items-center justify-center ml-6 shadow-lg">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"></path>
            </svg>
          </div>
          <div>
            <h3 class="text-xl font-bold text-gray-900">مسدود کردن آدرس IP</h3>
          </div>
        </div>
        
        <div class="space-y-6">
          <div class="bg-gradient-to-r from-red-50 to-orange-50 p-4 rounded-xl border border-red-100">
            <label class="block text-sm font-semibold text-gray-700 mb-2 flex items-center">
              <svg class="w-4 h-4 mr-2 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 12a9 9 0 01-9 9m9-9a9 9 0 00-9-9m9 9H3m9 9v-9m0-9v9"></path>
              </svg>
              آدرس IP
            </label>
            <input 
              v-model="blockIPData.ip_address" 
              type="text" 
              class="w-full px-4 py-3 border-2 border-red-200 rounded-xl focus:outline-none focus:ring-4 focus:ring-red-100 focus:border-red-500 bg-white font-mono text-red-600"
              readonly
            />
          </div>
          
          <div>
            <label class="block text-sm font-semibold text-gray-700 mb-2 flex items-center">
              <svg class="w-4 h-4 mr-2 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
              </svg>
              دلیل مسدودیت
            </label>
            <textarea 
              v-model="blockIPData.reason" 
              rows="3"
              class="w-full px-4 py-3 border-2 border-gray-200 rounded-xl focus:outline-none focus:ring-4 focus:ring-blue-100 focus:border-blue-500 transition-all duration-200 resize-none"
              placeholder="دلیل مسدود کردن این آدرس IP را وارد کنید..."
            ></textarea>
          </div>
          
          <div>
            <label class="block text-sm font-semibold text-gray-700 mb-2 flex items-center">
              <svg class="w-4 h-4 mr-2 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
              مدت زمان مسدودیت (ساعت)
            </label>
            <input 
              v-model="blockIPData.duration" 
              type="number" 
              min="1"
              class="w-full px-4 py-3 border-2 border-gray-200 rounded-xl focus:outline-none focus:ring-4 focus:ring-blue-100 focus:border-blue-500 transition-all duration-200"
              placeholder="24"
            />
          </div>
        </div>
        
        <div class="flex justify-end gap-4 mt-8 pt-6 border-t border-gray-200">
          <button 
            class="px-6 py-3 text-gray-600 bg-gray-100 border-2 border-gray-200 rounded-xl hover:bg-gray-200 hover:border-gray-300 transition-all duration-200 flex items-center shadow-sm hover:shadow-md" 
            @click="showBlockModal = false"
          >
            <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
            انصراف
          </button>
          <button 
            class="px-6 py-3 bg-gradient-to-r from-red-500 to-red-600 text-white rounded-xl hover:from-red-600 hover:to-red-700 transition-all duration-200 flex items-center shadow-lg hover:shadow-xl transform hover:scale-105" 
            @click="confirmBlockIP"
          >
            <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"></path>
            </svg>
            مسدود کردن
          </button>
        </div>
      </div>
    </div>

    <!-- مودال جزئیات ترافیک -->
    <div v-if="showTrafficDetailsModal" class="fixed inset-0 flex items-center justify-center z-50">
      <div 
        ref="trafficModalRef"
        class="bg-white rounded-2xl p-8 w-11/12 max-w-7xl max-h-[85vh] overflow-hidden shadow-2xl border-2 border-purple-500"
      >
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="w-12 h-12 bg-gradient-to-br from-blue-500 to-blue-600 rounded-xl flex items-center justify-center ml-8 shadow-lg">
              <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
              </svg>
            </div>
            <div>
              <h3 class="text-xl font-bold text-gray-900">جزئیات ترافیک</h3>
              <p class="text-sm text-gray-600">
                IP: 
                <a 
                  :href="`https://check-host.net/ip-info?host=${(selectedTrafficItem as any)?.ipAddress || ''}`" 
                  target="_blank" 
                  class="text-blue-600 hover:text-blue-800 hover:underline transition-colors duration-200 font-mono"
                  title="مشاهده اطلاعات IP در check-host.net"
                >
                  {{ (selectedTrafficItem as any)?.ipAddress }}
                </a>
              </p>
            </div>
          </div>
          <div class="flex items-center gap-3">
            <button 
              class="px-4 py-2 text-green-600 bg-green-50 border-2 border-green-200 rounded-lg hover:bg-green-100 hover:border-green-300 transition-all duration-200 flex items-center shadow-sm hover:shadow-md" 
              @click="exportTrafficDetails"
            >
              <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
              </svg>
              خروجی اکسل
            </button>
            <button 
              class="w-10 h-10 bg-gray-100 hover:bg-gray-200 rounded-xl flex items-center justify-center transition-all duration-200 hover:scale-110" 
              @click="showTrafficDetailsModal = false"
            >
              <svg class="w-5 h-5 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
              </svg>
            </button>
          </div>
        </div>
        
        <div class="overflow-y-auto max-h-[65vh] custom-scrollbar bg-gray-50 rounded-xl p-4">
          <table class="min-w-full divide-y divide-gray-200 bg-white rounded-lg overflow-hidden shadow-sm">
            <thead class="bg-gradient-to-r from-blue-50 to-indigo-50">
              <tr>
                <th class="px-6 py-4 text-right text-xs font-semibold text-blue-700 uppercase tracking-wider">کاربر/IP</th>
                <th class="px-6 py-4 text-right text-xs font-semibold text-blue-700 uppercase tracking-wider">نوع آگهی</th>
                <th class="px-6 py-4 text-right text-xs font-semibold text-blue-700 uppercase tracking-wider">موقعیت</th>
                <th class="px-6 py-4 text-right text-xs font-semibold text-blue-700 uppercase tracking-wider">صفحه</th>
                <th class="px-6 py-4 text-right text-xs font-semibold text-blue-700 uppercase tracking-wider">زمان</th>
                <th class="px-6 py-4 text-right text-xs font-semibold text-blue-700 uppercase tracking-wider">مرورگر</th>
                <th class="px-6 py-4 text-right text-xs font-semibold text-blue-700 uppercase tracking-wider">نوع دستگاه</th>
                <th class="px-6 py-4 text-right text-xs font-semibold text-blue-700 uppercase tracking-wider">سیستم عامل</th>
                <th class="px-6 py-4 text-right text-xs font-semibold text-blue-700 uppercase tracking-wider">نام هاست</th>
                <th class="px-6 py-4 text-right text-xs font-semibold text-blue-700 uppercase tracking-wider">پاسخ</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-if="trafficDetails.length === 0">
                <td colspan="10" class="px-6 py-12 text-center text-gray-500">
                  <div class="flex flex-col items-center">
                    <svg class="w-12 h-12 text-gray-300 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
                    </svg>
                    <span class="text-lg font-medium">هیچ جزئیاتی یافت نشد</span>
                  </div>
                </td>
              </tr>
              <tr v-for="detail in trafficDetails" v-else :key="(detail as any).id" class="hover:bg-blue-50 transition-colors duration-200">
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900 font-medium">
                  <a 
                    :href="`https://check-host.net/ip-info?host=${(detail as any).ipAddress || (selectedTrafficItem as any)?.ipAddress || ''}`" 
                    target="_blank" 
                    class="text-blue-600 hover:text-blue-800 hover:underline transition-colors duration-200 font-mono"
                    title="مشاهده اطلاعات IP در check-host.net"
                  >
                    {{ (detail as any).ipAddress || (selectedTrafficItem as any)?.ipAddress || '-' }}
                  </a>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  <div class="flex items-center justify-center">
                    <div 
                      class="w-4 h-4 rounded-full"
                      :class="{
                        'bg-green-500': (detail as any).adType === 'بنیر',
                        'bg-gray-500': (detail as any).adType === 'اسلایدر',
                        'bg-yellow-500': (detail as any).adType === 'پاپ‌آپ',
                        'bg-red-500': (detail as any).adType === 'مشکوک'
                      }"
                      :title="(detail as any).adType || '-'"
                    ></div>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ (detail as any).location || '-' }}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900 font-medium">{{ (detail as any).page }}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ (detail as any).timestamp }}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ (detail as any).browser || '-' }}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ (detail as any).deviceType || '-' }}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ (detail as any).os || '-' }}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ (detail as any).hostname || '-' }}</td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span 
                    class="px-3 py-1 text-xs font-semibold rounded-full"
                    :class="{
                      'bg-green-100 text-green-800': (detail as any).status === 200,
                      'bg-red-100 text-red-800': (detail as any).status === 401,
                      'bg-yellow-100 text-yellow-800': (detail as any).status === 404,
                      'bg-gray-100 text-gray-800': (detail as any).status >= 500
                    }"
                  >
                    {{ (detail as any).status }}
                  </span>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>

    <!-- مودال جزئیات کاربر -->
    <div v-if="showUserDetailsModal" class="fixed inset-0 flex items-center justify-center z-50">
      <div 
        ref="userModalRef"
        class="bg-white rounded-2xl p-8 w-4/5 max-w-5xl max-h-[85vh] overflow-hidden shadow-2xl border-2 border-purple-500"
      >
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="w-12 h-12 bg-gradient-to-br from-green-500 to-green-600 rounded-xl flex items-center justify-center mr-4 shadow-lg">
              <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
              </svg>
            </div>
            <div>
              <h3 class="text-xl font-bold text-gray-900">جزئیات فعالیت کاربر</h3>
              <p class="text-sm text-gray-600">{{ (selectedUser as any)?.name }}</p>
            </div>
          </div>
          <button 
            class="w-10 h-10 bg-gray-100 hover:bg-gray-200 rounded-xl flex items-center justify-center transition-all duration-200 hover:scale-110" 
            @click="showUserDetailsModal = false"
          >
            <svg class="w-5 h-5 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>
        
        <div class="mb-6 p-6 bg-gradient-to-r from-green-50 to-emerald-50 rounded-xl border border-green-100">
          <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
            <div class="flex items-center">
              <div class="w-10 h-10 bg-green-100 rounded-lg flex items-center justify-center mr-3">
                <svg class="w-5 h-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
                </svg>
              </div>
              <div>
                <span class="text-xs font-medium text-gray-500">نام</span>
                <div class="text-sm font-semibold text-gray-900">{{ (selectedUser as any)?.name }}</div>
              </div>
            </div>
            <div class="flex items-center">
              <div class="w-10 h-10 bg-blue-100 rounded-lg flex items-center justify-center mr-3">
                <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 12a9 9 0 01-9 9m9-9a9 9 0 00-9-9m9 9H3m9 9v-9m0-9v9"></path>
                </svg>
              </div>
              <div>
                <span class="text-xs font-medium text-gray-500">IP</span>
                <div class="text-sm font-semibold text-gray-900 font-mono">{{ (selectedUser as any)?.ip }}</div>
              </div>
            </div>
            <div class="flex items-center">
              <div class="w-10 h-10 bg-purple-100 rounded-lg flex items-center justify-center mr-3">
                <svg class="w-5 h-5 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
                </svg>
              </div>
              <div>
                <span class="text-xs font-medium text-gray-500">صفحه فعلی</span>
                <div class="text-sm font-semibold text-gray-900">{{ (selectedUser as any)?.currentPage || '/' }}</div>
              </div>
            </div>
          </div>
        </div>
        
        <div class="overflow-y-auto max-h-[55vh] custom-scrollbar bg-gray-50 rounded-xl p-4">
          <table class="min-w-full divide-y divide-gray-200 bg-white rounded-lg overflow-hidden shadow-sm">
            <thead class="bg-gradient-to-r from-green-50 to-emerald-50">
              <tr>
                <th class="px-6 py-4 text-right text-xs font-semibold text-green-700 uppercase tracking-wider">صفحه</th>
                <th class="px-6 py-4 text-right text-xs font-semibold text-green-700 uppercase tracking-wider">عملیات</th>
                <th class="px-6 py-4 text-right text-xs font-semibold text-green-700 uppercase tracking-wider">زمان</th>
                <th class="px-6 py-4 text-right text-xs font-semibold text-green-700 uppercase tracking-wider">مدت زمان</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-if="userDetails.length === 0">
                <td colspan="4" class="px-6 py-12 text-center text-gray-500">
                  <div class="flex flex-col items-center">
                    <svg class="w-12 h-12 text-gray-300 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
                    </svg>
                    <span class="text-lg font-medium">هیچ فعالیتی یافت نشد</span>
                  </div>
                </td>
              </tr>
              <tr v-for="detail in userDetails" v-else :key="(detail as any).id" class="hover:bg-green-50 transition-colors duration-200">
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900 font-medium">{{ (detail as any).page }}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-700">{{ (detail as any).action }}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ (detail as any).timestamp }}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-600">{{ (detail as any).duration }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void

// تعریف interface ها
interface ApiResponse<T> {
  data?: T
  success?: boolean
  message?: string
}

interface OnlineUser {
  id: string
  ip: string
  userAgent: string
  lastActivity: string
  location?: string
}

interface TrafficDetail {
  id: string
  ipAddress: string
  timestamp: string
  request: string
  response: string
  userAgent: string
}

interface UserDetail {
  id: string
  ip: string
  userAgent: string
  lastActivity: string
  location?: string
}

export default {
  name: 'TrafficPage'
}
</script>

<script setup lang="ts">
import TemplateCard from '@/components/common/TemplateCard.vue';
import { onMounted, onUnmounted, ref, watchEffect } from 'vue';

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
// Auth disabled
// داده‌های واکنشی
const trafficStats = ref({
  online_users: 12,
  today_requests: 1250,
  suspicious_requests: 5,
  blocked_attacks: 2,
  hourly_traffic: []
})

type SettingResponse = {
  value?: string | number | boolean
  key?: string
  category?: string
  type?: string
}

// وضعیت ثبت لاگ ترافیک
const trafficLoggingEnabled = ref(true)
const loadTrafficLogging = async () => {
  try {
    const setting = await $fetch<SettingResponse>('/api/admin/settings/traffic.logging.enabled')
    const val = (setting?.value || '').toString().toLowerCase()
    trafficLoggingEnabled.value = val === 'true' || val === '1' || val === 'yes'
  } catch {
    trafficLoggingEnabled.value = true
  }
}
const toggleTrafficLogging = async () => {
  const newVal = trafficLoggingEnabled.value ? 'false' : 'true'
  await $fetch('/api/admin/settings/traffic.logging.enabled', { method: 'PUT', body: { key: 'traffic.logging.enabled', value: newVal, category: 'traffic', type: 'boolean' } })
  await loadTrafficLogging()
}

const generalTrafficStats = ref({
  traffic_logs: [],
  pagination: {
    current_page: 1,
    total_pages: 1,
    total_items: 0,
    items_per_page: 10
  }
})

// متغیرهای بارگذاری
const isLoading = ref(false) // فقط برای جدول ترافیک استفاده شود
const isOnlineUsersLoading = ref(false) // مخصوص جدول کاربران آنلاین

// متغیرهای تعداد آیتم در هر صفحه
const trafficItemsPerPage = ref(10)
const onlineUsersItemsPerPage = ref(10)

const onlineUsers = ref<OnlineUser[]>([])
const blockedIPsCount = ref(2)
const showBlockModal = ref(false)
const blockIPData = ref({
  ip_address: '',
  reason: '',
  duration: 24
})

// متغیرهای مودال جزئیات ترافیک
const showTrafficDetailsModal = ref(false)
const selectedTrafficItem = ref(null)
const trafficDetails = ref([])

// متغیرهای مودال جزئیات کاربر
const showUserDetailsModal = ref(false)
const selectedUser = ref(null)
const userDetails = ref([])

// متغیرهای فیلتر
const filters = ref({
  adType: '',
  ipAddress: '',
  location: '',
  page: '',
  status: '',
  timeRange: ''
})

const filteredTrafficCount = ref(0)
const totalTrafficCount = ref(0)

// refs برای مودال‌ها
const blockModalRef = ref(null)
const trafficModalRef = ref(null)
const userModalRef = ref(null)

// Debug: نظارت بر تغییرات داده‌ها
watchEffect(() => {
  // onlineUsers تغییر کرد
})

watchEffect(() => {
  // generalTrafficStats تغییر کرد
})

type TrafficStats = {
  online_users?: number
  today_requests?: number
  suspicious_requests?: number
  blocked_attacks?: number
  hourly_traffic?: unknown[]
}

type TrafficGeneralStats = {
  traffic_logs?: unknown[]
  pagination?: {
    current_page?: number
    total_pages?: number
    total_items?: number
    items_per_page?: number
  }
}

// دریافت آمار ترافیک
const fetchTrafficStats = async () => {
  try {
    const response = await $fetch<ApiResponse<TrafficStats>>('/api/admin/traffic/stats')
    trafficStats.value = (response.data ? {
      online_users: response.data.online_users ?? 12,
      today_requests: response.data.today_requests ?? 1250,
      suspicious_requests: response.data.suspicious_requests ?? 5,
      blocked_attacks: response.data.blocked_attacks ?? 2,
      hourly_traffic: response.data.hourly_traffic ?? []
    } : {
      online_users: 12,
      today_requests: 1250,
      suspicious_requests: 5,
      blocked_attacks: 2,
      hourly_traffic: []
    })
  } catch (error) {
    // خطا در دریافت آمار ترافیک
    // داده‌های نمونه برای تست
    trafficStats.value = {
      online_users: 12,
      today_requests: 1250,
      suspicious_requests: 5,
      blocked_attacks: 2,
      hourly_traffic: []
    }
  }
}

// دریافت آمار ترافیک کلی
const fetchGeneralTrafficStats = async () => {
  isLoading.value = true
  try {
    // درخواست آمار ترافیک کلی
    const response = await $fetch<ApiResponse<TrafficGeneralStats>>('/api/admin/traffic/general-stats')
    // پاسخ آمار ترافیک کلی
    generalTrafficStats.value = (response.data ? {
      traffic_logs: response.data.traffic_logs ?? [],
      pagination: {
        current_page: response.data.pagination?.current_page ?? 1,
        total_pages: response.data.pagination?.total_pages ?? 1,
        total_items: response.data.pagination?.total_items ?? 0,
        items_per_page: response.data.pagination?.items_per_page ?? 10
      }
    } : {
      traffic_logs: [],
      pagination: {
        current_page: 1,
        total_pages: 1,
        total_items: 0,
        items_per_page: 10
      }
    })
  } catch (error) {
    // خطا در دریافت آمار ترافیک کلی
    // استفاده از داده‌های تستی
    
    // داده‌های نمونه برای تست - همه داده‌ها در یک لیست
    const mockData = []
    
    for (let i = 1; i <= 100; i++) {
      const adTypes = ['بنیر', 'اسلایدر', 'پاپ‌آپ', 'مشکوک']
      const randomAdType = adTypes[i % adTypes.length]
      
      mockData.push({
        ID: i,
        AdType: randomAdType,
        Location: i % 4 === 0 ? 'تهران' : i % 4 === 1 ? 'اصفهان' : i % 4 === 2 ? 'مشهد' : 'شیراز',
        City: i % 4 === 0 ? 'تهران' : i % 4 === 1 ? 'اصفهان' : i % 4 === 2 ? 'مشهد' : 'شیراز',
        RequestPath: i % 5 === 0 ? '/admin/dashboard' : i % 5 === 1 ? '/products' : i % 5 === 2 ? '/cart' : i % 5 === 3 ? '/blog' : '/profile',
        CreatedAt: new Date(Date.now() - i * 60000).toISOString(),
        IPAddress: `192.168.1.${100 + i}`,
        Hostname: 'localhost',
        StatusCode: 200,
        ResponseTime: 150 + (i * 10),
        Display: 'نمایش داده شد',
        Browser: i % 3 === 0 ? 'Chrome' : i % 3 === 1 ? 'Firefox' : 'Edge',
        DeviceType: i % 2 === 0 ? 'Desktop' : 'Mobile',
        OS: i % 4 === 0 ? 'Windows' : i % 4 === 1 ? 'macOS' : i % 4 === 2 ? 'Linux' : 'Android'
      })
    }
    
    generalTrafficStats.value = {
      traffic_logs: mockData,
      pagination: {
        current_page: 1,
        total_pages: 1,
        total_items: mockData.length,
        items_per_page: mockData.length
      }
    }
    
    // داده‌های ترافیک بارگذاری شد
  } finally {
    isLoading.value = false
  }
}

// دریافت کاربران آنلاین
const fetchOnlineUsers = async () => {
  isOnlineUsersLoading.value = true
  try {
    // درخواست کاربران آنلاین
    const response: ApiResponse<OnlineUser[]> = await $fetch('/api/admin/traffic/online-users')
    // پاسخ کاربران آنلاین
    onlineUsers.value = response.data || []
  } catch (error) {
    // خطا در دریافت کاربران آنلاین
    // استفاده از داده‌های تستی برای کاربران آنلاین
    
    // داده‌های تستی - همه کاربران در یک لیست
    const allUsers = []
    const names = ['احمد محمدی', 'سارا احمدی', 'علی رضایی', 'فاطمه کریمی', 'محمد حسینی', 'زهرا نوری', 'حسین احمدی', 'مریم صادقی', 'رضا محمودی', 'نرگس کریمی', 'امیر رضایی', 'لیلا صادقی', 'علی کریمی', 'سارا نوری', 'محمد احمدی']
    const pages = ['/admin/dashboard', '/products', '/cart', '/blog', '/profile', '/orders', '/wishlist', '/search', '/categories', '/checkout', '/auth/login', '/register', '/about', '/contact', '/help']
    const adTypes = ['بنیر', 'اسلایدر', 'پاپ‌آپ', 'مشکوک']
    for (let i = 1; i <= 50; i++) {
      const randomName = names[i % names.length]
      const randomPage = pages[i % pages.length]
      const randomMinute = Math.floor(Math.random() * 60)
      const randomSecond = Math.floor(Math.random() * 60)
      const locations = ['تهران', 'اصفهان', 'مشهد', 'شیراز', 'تبریز', 'یزد', 'کرج', 'قم']
      const randomLocation = locations[i % locations.length]
      const randomAdType = adTypes[i % adTypes.length]
      allUsers.push({
        id: i,
        name: randomName,
        email: i % 5 === 0 ? '' : `${randomName.replace(/\s+/g, '').toLowerCase()}@example.com`,
        ip: `192.168.1.${100 + i}`,
        location: randomLocation,
        hostname: 'localhost',
        currentPage: randomPage,
        lastSeen: `14:${randomMinute.toString().padStart(2, '0')}:${randomSecond.toString().padStart(2, '0')}`,
        response: '200 (150ms)',
        adType: randomAdType,
        browser: i % 3 === 0 ? 'Chrome' : i % 3 === 1 ? 'Firefox' : 'Edge',
        deviceType: i % 2 === 0 ? 'Desktop' : 'Mobile',
        os: i % 4 === 0 ? 'Windows' : i % 4 === 1 ? 'macOS' : i % 4 === 2 ? 'Linux' : 'Android'
      })
    }
    
    onlineUsers.value = allUsers
    
    // داده‌های کاربران آنلاین بارگذاری شد
  } finally {
    isOnlineUsersLoading.value = false
  }
}

type BlockedIP = {
  id?: number
  ip_address?: string
  reason?: string
  created_at?: string
}

// دریافت تعداد IP های مسدود شده
const fetchBlockedIPsCount = async () => {
  try {
    const response = await $fetch<ApiResponse<BlockedIP[]>>('/api/admin/traffic/blocked-ips')
    blockedIPsCount.value = (response.data || []).length
  } catch (error) {
    // خطا در دریافت تعداد IP های مسدود شده
    // داده‌های نمونه برای تست
    blockedIPsCount.value = 2
  }
}





// تغییر تعداد آیتم در هر صفحه برای جدول ترافیک
const handleTrafficItemsPerPageChange = () => {
  // تعداد آیتم‌های ترافیک تغییر کرد
  // در اینجا می‌توانید منطق فیلتر کردن داده‌ها را اضافه کنید
}

// تغییر تعداد آیتم در هر صفحه برای جدول کاربران آنلاین
const handleOnlineUsersItemsPerPageChange = () => {
  // تعداد آیتم‌های کاربران آنلاین تغییر کرد
  // در اینجا می‌توانید منطق فیلتر کردن داده‌ها را اضافه کنید
}

// نمایش جزئیات ترافیک
const showTrafficDetails = async (item) => {
  selectedTrafficItem.value = item
  showTrafficDetailsModal.value = true
  
  try {
    // درخواست جزئیات ترافیک از API
    const response: ApiResponse<TrafficDetail[]> = await $fetch(`/api/admin/traffic/details/${item.ipAddress}`)
    trafficDetails.value = response.data || []
  } catch (error) {
    // خطا در دریافت جزئیات ترافیک
    // داده‌های تستی برای نمایش
    trafficDetails.value = [
      {
        id: 1,
        page: '/admin/dashboard',
        timestamp: new Date(Date.now() - 300000).toLocaleString('fa-IR'),
        action: 'بازدید صفحه',
        status: 200
      },
      {
        id: 2,
        page: '/products',
        timestamp: new Date(Date.now() - 600000).toLocaleString('fa-IR'),
        action: 'جستجوی محصول',
        status: 200
      },
      {
        id: 3,
        page: '/cart',
        timestamp: new Date(Date.now() - 900000).toLocaleString('fa-IR'),
        action: 'افزودن به سبد خرید',
        status: 200
      }
    ]
  }
}

// نمایش جزئیات کاربر
const showUserDetails = async (user) => {
  selectedUser.value = user
  showUserDetailsModal.value = true
  
  try {
    // درخواست جزئیات کاربر از API
    const response: ApiResponse<UserDetail[]> = await $fetch(`/api/admin/traffic/user-details/${user.ip}`)
    userDetails.value = response.data || []
  } catch (error) {
    // خطا در دریافت جزئیات کاربر
    // داده‌های تستی برای نمایش
    userDetails.value = [
      {
        id: 1,
        page: user.currentPage || '/',
        timestamp: user.lastSeen,
        action: 'صفحه فعلی',
        duration: '2 دقیقه'
      },
      {
        id: 2,
        page: '/products',
        timestamp: new Date(Date.now() - 300000).toLocaleString('fa-IR'),
        action: 'بازدید محصولات',
        duration: '5 دقیقه'
      },
      {
        id: 3,
        page: '/cart',
        timestamp: new Date(Date.now() - 600000).toLocaleString('fa-IR'),
        action: 'افزودن به سبد خرید',
        duration: '3 دقیقه'
      },
      {
        id: 4,
        page: '/profile',
        timestamp: new Date(Date.now() - 900000).toLocaleString('fa-IR'),
        action: 'مشاهده پروفایل',
        duration: '1 دقیقه'
      },
      {
        id: 5,
        page: '/search',
        timestamp: new Date(Date.now() - 1200000).toLocaleString('fa-IR'),
        action: 'جستجو',
        duration: '2 دقیقه'
      }
    ]
  }
}

// اعمال فیلترها
const applyFilters = () => {
  // فیلترها اعمال شدند
  // در اینجا منطق فیلتر کردن داده‌ها را اضافه کنید
  updateFilteredCount()
}

// پاک کردن فیلترها
const clearFilters = () => {
  filters.value = {
    adType: '',
    ipAddress: '',
    location: '',
    page: '',
    status: '',
    timeRange: ''
  }
  // فیلترها پاک شدند
  updateFilteredCount()
}

// به‌روزرسانی تعداد رکوردهای فیلتر شده
const updateFilteredCount = () => {
  // این تابع تعداد رکوردهای فیلتر شده را محاسبه می‌کند
  totalTrafficCount.value = generalTrafficStats.value.traffic_logs?.length || 0
  filteredTrafficCount.value = totalTrafficCount.value // فعلاً همه رکوردها
}

// خروجی اکسل
const exportFilteredData = () => {
  // خروجی اکسل درخواست شد
  // در اینجا منطق خروجی اکسل را اضافه کنید
  alert('خروجی اکسل در حال آماده‌سازی است...')
}

// خروجی اکسل جزئیات ترافیک
const exportTrafficDetails = () => {
  // خروجی اکسل جزئیات ترافیک درخواست شد
  // در اینجا منطق خروجی اکسل جزئیات ترافیک را اضافه کنید
  alert('خروجی اکسل جزئیات ترافیک در حال آماده‌سازی است...')
}

// مسدود کردن IP
const blockIP = (ip) => {
  blockIPData.value.ip_address = ip
  blockIPData.value.reason = ''
  blockIPData.value.duration = 24
  showBlockModal.value = true
}

// تایید مسدود کردن IP
const confirmBlockIP = async () => {
  try {
    await $fetch('/api/admin/traffic/block-ip', {
      method: 'POST',
      body: blockIPData.value
    })
    
    showBlockModal.value = false
    await refreshData()
    
    // نمایش پیام موفقیت
    alert('آدرس IP با موفقیت مسدود شد')
  } catch (error) {
    // خطا در مسدود کردن IP
    alert('خطا در مسدود کردن آدرس IP')
  }
}





// دریافت داده‌های جدول
const getTableData = () => {
  const trafficLogs = generalTrafficStats.value.traffic_logs || []
  
  return trafficLogs.map((log, index) => {
    return {
      id: log.ID,
      adType: log.AdType || '-',
      location: log.Location || log.City || '-',
      page: log.RequestPath || '-',
      time: new Date(log.CreatedAt).toLocaleString('fa-IR'),
      ipAddress: log.IPAddress || '-',
      hostname: log.Hostname || '-',
      response: log.StatusCode ? `${log.StatusCode} (${log.ResponseTime}ms)` : '-',
      display: log.Display || '-',
      browser: log.Browser || 'Chrome',
      deviceType: log.DeviceType || 'Desktop',
      os: log.OS || 'Windows'
    }
  })
}

// به‌روزرسانی داده‌ها
const refreshData = async () => {
  await Promise.all([
    fetchTrafficStats(),
    fetchGeneralTrafficStats(),
    fetchOnlineUsers(),
    fetchBlockedIPsCount()
  ])
}

// به‌روزرسانی خودکار هر 30 ثانیه
let refreshInterval

onMounted(async () => {
  // onMounted شروع شد
  
  // بارگذاری همه داده‌ها
  await Promise.all([
    fetchTrafficStats(),
    fetchGeneralTrafficStats(),
    fetchOnlineUsers(),
    fetchBlockedIPsCount(),
    loadTrafficLogging()
  ])
  
  // همه داده‌ها بارگذاری شدند
  
  // به‌روزرسانی تعداد رکوردها
  updateFilteredCount()
  
  // تنظیم به‌روزرسانی خودکار
  refreshInterval = setInterval(refreshData, 30000)
})

onUnmounted(() => {
  if (refreshInterval) {
    clearInterval(refreshInterval)
  }
})


</script>

<style scoped>
.custom-scrollbar {
  direction: ltr;
}
.custom-scrollbar table {
  direction: rtl;
}
.custom-scrollbar::-webkit-scrollbar {
  width: 8px;
}
.custom-scrollbar::-webkit-scrollbar-track {
  background: #f7fafc;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background: #cbd5e0;
  border-radius: 4px;
}
.custom-scrollbar::-webkit-scrollbar-thumb:hover {
  background: #a0aec0;
}
</style> 

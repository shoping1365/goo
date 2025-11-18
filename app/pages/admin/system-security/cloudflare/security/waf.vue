<template>
  <div class="flex h-screen bg-gray-50" dir="rtl">
    <!-- سایدبار کلادفلر -->
    <div class="w-64 bg-white shadow-lg border-l border-gray-200 overflow-y-auto">
      <!-- لوگو و عنوان -->
      <div class="p-6 border-b border-gray-200">
        <div class="flex items-center space-x-3 space-x-reverse">
          <div class="w-8 h-8 bg-orange-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="currentColor" viewBox="0 0 24 24">
              <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/>
            </svg>
          </div>
          <h1 class="text-lg font-bold text-gray-900">کلادفلر</h1>
        </div>
      </div>

      <!-- منوی اصلی -->
      <nav class="p-6">
        <div class="space-y-1">
          <!-- Overview -->
          <NuxtLink 
            to="/admin/system-security/cloudflare" 
            class="flex items-center px-3 py-2 text-sm font-medium rounded-md transition-colors"
            :class="route.path === '/admin/system-security/cloudflare' ? 'bg-blue-50 text-blue-700' : 'text-gray-700 hover:bg-gray-50'"
          >
            <svg class="w-5 h-5 ml-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"/>
            </svg>
            نمای کلی
          </NuxtLink>

          <!-- Security -->
          <div class="mt-6">
            <h3 class="px-3 text-xs font-semibold text-gray-500 uppercase tracking-wider">امنیت</h3>
            <div class="mt-2 space-y-1">
              <NuxtLink 
                to="/admin/system-security/cloudflare/security/waf" 
                class="flex items-center px-3 py-2 text-sm font-medium rounded-md transition-colors"
                :class="route.path.includes('/waf') ? 'bg-blue-50 text-blue-700' : 'text-gray-700 hover:bg-gray-50'"
              >
                <svg class="w-5 h-5 ml-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"/>
                </svg>
                WAF
              </NuxtLink>
              <NuxtLink 
                to="/admin/system-security/cloudflare/security/ddos" 
                class="flex items-center px-3 py-2 text-sm font-medium rounded-md transition-colors"
                :class="route.path.includes('/ddos') ? 'bg-blue-50 text-blue-700' : 'text-gray-700 hover:bg-gray-50'"
              >
                <svg class="w-5 h-5 ml-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"/>
                </svg>
                DDoS Protection
              </NuxtLink>
              <NuxtLink 
                to="/admin/system-security/cloudflare/security/bot-management" 
                class="flex items-center px-3 py-2 text-sm font-medium rounded-md transition-colors"
                :class="route.path.includes('/bot-management') ? 'bg-blue-50 text-blue-700' : 'text-gray-700 hover:bg-gray-50'"
              >
                <svg class="w-5 h-5 ml-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.75 17L9 20l-1 1h8l-1-1-.75-3M3 13h18M5 17h14a2 2 0 002-2V5a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"/>
                </svg>
                مدیریت ربات‌ها
              </NuxtLink>
              <NuxtLink 
                to="/admin/system-security/cloudflare/security/rate-limiting" 
                class="flex items-center px-3 py-2 text-sm font-medium rounded-md transition-colors"
                :class="route.path.includes('/rate-limiting') ? 'bg-blue-50 text-blue-700' : 'text-gray-700 hover:bg-gray-50'"
              >
                <svg class="w-5 h-5 ml-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"/>
                </svg>
                محدودیت نرخ
              </NuxtLink>
            </div>
          </div>

          <!-- Performance -->
          <div class="mt-6">
            <h3 class="px-3 text-xs font-semibold text-gray-500 uppercase tracking-wider">عملکرد</h3>
            <div class="mt-2 space-y-1">
              <NuxtLink 
                to="/admin/system-security/cloudflare/performance/caching" 
                class="flex items-center px-3 py-2 text-sm font-medium rounded-md transition-colors"
                :class="route.path.includes('/caching') ? 'bg-blue-50 text-blue-700' : 'text-gray-700 hover:bg-gray-50'"
              >
                <svg class="w-5 h-5 ml-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"/>
                </svg>
                کش
              </NuxtLink>
              <NuxtLink 
                to="/admin/system-security/cloudflare/performance/page-rules" 
                class="flex items-center px-3 py-2 text-sm font-medium rounded-md transition-colors"
                :class="route.path.includes('/page-rules') ? 'bg-blue-50 text-blue-700' : 'text-gray-700 hover:bg-gray-50'"
              >
                <svg class="w-5 h-5 ml-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"/>
                </svg>
                قوانین صفحه
              </NuxtLink>
              <NuxtLink 
                to="/admin/system-security/cloudflare/performance/workers" 
                class="flex items-center px-3 py-2 text-sm font-medium rounded-md transition-colors"
                :class="route.path.includes('/workers') ? 'bg-blue-50 text-blue-700' : 'text-gray-700 hover:bg-gray-50'"
              >
                <svg class="w-5 h-5 ml-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"/>
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/>
                </svg>
                Workers
              </NuxtLink>
            </div>
          </div>

          <!-- DNS -->
          <div class="mt-6">
            <h3 class="px-3 text-xs font-semibold text-gray-500 uppercase tracking-wider">DNS</h3>
            <div class="mt-2 space-y-1">
              <NuxtLink 
                to="/admin/system-security/cloudflare/dns/records" 
                class="flex items-center px-3 py-2 text-sm font-medium rounded-md transition-colors"
                :class="route.path.includes('/records') ? 'bg-blue-50 text-blue-700' : 'text-gray-700 hover:bg-gray-50'"
              >
                <svg class="w-5 h-5 ml-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"/>
                </svg>
                رکوردهای DNS
              </NuxtLink>
              <NuxtLink 
                to="/admin/system-security/cloudflare/dns/analytics" 
                class="flex items-center px-3 py-2 text-sm font-medium rounded-md transition-colors"
                :class="route.path.includes('/analytics') ? 'bg-blue-50 text-blue-700' : 'text-gray-700 hover:bg-gray-50'"
              >
                <svg class="w-5 h-5 ml-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"/>
                </svg>
                تحلیل DNS
              </NuxtLink>
            </div>
          </div>

          <!-- Analytics -->
          <div class="mt-6">
            <h3 class="px-3 text-xs font-semibold text-gray-500 uppercase tracking-wider">تحلیل</h3>
            <div class="mt-2 space-y-1">
              <NuxtLink 
                to="/admin/system-security/cloudflare/analytics/traffic" 
                class="flex items-center px-3 py-2 text-sm font-medium rounded-md transition-colors"
                :class="route.path.includes('/traffic') ? 'bg-blue-50 text-blue-700' : 'text-gray-700 hover:bg-gray-50'"
              >
                <svg class="w-5 h-5 ml-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"/>
                </svg>
                ترافیک
              </NuxtLink>
              <NuxtLink 
                to="/admin/system-security/cloudflare/analytics/security" 
                class="flex items-center px-3 py-2 text-sm font-medium rounded-md transition-colors"
                :class="route.path.includes('/security') ? 'bg-blue-50 text-blue-700' : 'text-gray-700 hover:bg-gray-50'"
              >
                <svg class="w-5 h-5 ml-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"/>
                </svg>
                امنیت
              </NuxtLink>
              <NuxtLink 
                to="/admin/system-security/cloudflare/analytics/performance" 
                class="flex items-center px-3 py-2 text-sm font-medium rounded-md transition-colors"
                :class="route.path.includes('/performance') ? 'bg-blue-50 text-blue-700' : 'text-gray-700 hover:bg-gray-50'"
              >
                <svg class="w-5 h-5 ml-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"/>
                </svg>
                عملکرد
              </NuxtLink>
            </div>
          </div>

          <!-- Settings -->
          <div class="mt-6">
            <h3 class="px-3 text-xs font-semibold text-gray-500 uppercase tracking-wider">تنظیمات</h3>
            <div class="mt-2 space-y-1">
              <NuxtLink 
                to="/admin/system-security/cloudflare/settings/general" 
                class="flex items-center px-3 py-2 text-sm font-medium rounded-md transition-colors"
                :class="route.path.includes('/general') ? 'bg-blue-50 text-blue-700' : 'text-gray-700 hover:bg-gray-50'"
              >
                <svg class="w-5 h-5 ml-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"/>
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/>
                </svg>
                عمومی
              </NuxtLink>
              <NuxtLink 
                to="/admin/system-security/cloudflare/settings/api" 
                class="flex items-center px-3 py-2 text-sm font-medium rounded-md transition-colors"
                :class="route.path.includes('/api') ? 'bg-blue-50 text-blue-700' : 'text-gray-700 hover:bg-gray-50'"
              >
                <svg class="w-5 h-5 ml-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 9l3 3-3 3m5 0h3M5 20h14a2 2 0 002-2V6a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"/>
                </svg>
                API
              </NuxtLink>
            </div>
          </div>
        </div>
      </nav>
    </div>

    <!-- محتوای اصلی -->
    <div class="flex-1 overflow-auto">
      <div class="p-6">
        <!-- هدر صفحه -->
        <div class="mb-6">
          <h1 class="text-2xl font-bold text-gray-900">Web Application Firewall (WAF)</h1>
          <p class="text-gray-600">مدیریت قوانین امنیتی و محافظت از برنامه وب</p>
        </div>

        <!-- کارت‌های آمار WAF -->
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
          <div class="bg-white rounded-lg shadow p-6 border border-gray-200">
            <div class="flex items-center">
              <div class="p-2 bg-red-100 rounded-lg">
                <svg class="w-6 h-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"/>
                </svg>
              </div>
              <div class="mr-4">
                <p class="text-sm font-medium text-gray-600">تهدیدات مسدود شده</p>
                <p class="text-2xl font-bold text-gray-900">1,847</p>
              </div>
            </div>
          </div>

          <div class="bg-white rounded-lg shadow p-6 border border-gray-200">
            <div class="flex items-center">
              <div class="p-2 bg-yellow-100 rounded-lg">
                <svg class="w-6 h-6 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"/>
                </svg>
              </div>
              <div class="mr-4">
                <p class="text-sm font-medium text-gray-600">چالش‌های ارسال شده</p>
                <p class="text-2xl font-bold text-gray-900">342</p>
              </div>
            </div>
          </div>

          <div class="bg-white rounded-lg shadow p-6 border border-gray-200">
            <div class="flex items-center">
              <div class="p-2 bg-green-100 rounded-lg">
                <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
                </svg>
              </div>
              <div class="mr-4">
                <p class="text-sm font-medium text-gray-600">قوانین فعال</p>
                <p class="text-2xl font-bold text-gray-900">24</p>
              </div>
            </div>
          </div>

          <div class="bg-white rounded-lg shadow p-6 border border-gray-200">
            <div class="flex items-center">
              <div class="p-2 bg-blue-100 rounded-lg">
                <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"/>
                </svg>
              </div>
              <div class="mr-4">
                <p class="text-sm font-medium text-gray-600">درخواست‌های بررسی شده</p>
                <p class="text-2xl font-bold text-gray-900">45.2K</p>
              </div>
            </div>
          </div>
        </div>

        <!-- دکمه افزودن قانون جدید -->
        <div class="mb-6">
          <button 
            @click="showAddRuleModal = true"
            class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors flex items-center"
          >
            <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"/>
            </svg>
            افزودن قانون جدید
          </button>
        </div>

        <!-- جدول قوانین WAF -->
        <div class="bg-white rounded-lg shadow border border-gray-200">
          <div class="px-6 py-4 border-b border-gray-200">
            <h3 class="text-lg font-semibold text-gray-900">قوانین WAF</h3>
          </div>
          <div class="overflow-x-auto">
            <table class="min-w-full divide-y divide-gray-200">
              <thead class="bg-gray-50">
                <tr>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نام قانون</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عبارت</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ ایجاد</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
                </tr>
              </thead>
              <tbody class="bg-white divide-y divide-gray-200">
                <tr v-for="rule in wafRules" :key="rule.id">
                  <td class="px-6 py-4 whitespace-nowrap">
                    <div class="text-sm font-medium text-gray-900">{{ rule.name }}</div>
                    <div class="text-sm text-gray-500">{{ rule.description }}</div>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap">
                    <div class="text-sm text-gray-900 font-mono bg-gray-100 px-2 py-1 rounded">{{ rule.filter?.expression || rule.expression }}</div>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap">
                    <span :class="getActionClass(rule.action)" class="px-2 py-1 text-xs font-medium rounded-full">
                      {{ getActionText(rule.action) }}
                    </span>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap">
                    <span :class="!rule.paused ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'" class="px-2 py-1 text-xs font-medium rounded-full">
                      {{ !rule.paused ? 'فعال' : 'غیرفعال' }}
                    </span>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ rule.createdAt }}</td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                    <div class="flex space-x-2 space-x-reverse">
                      <button 
                        @click="editRule(rule)"
                        class="text-blue-600 hover:text-blue-900 transition-colors"
                        title="ویرایش"
                      >
                        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"/>
                        </svg>
                      </button>
                      <button 
                        @click="toggleRule(rule)"
                        :class="!rule.paused ? 'text-red-600 hover:text-red-900' : 'text-green-600 hover:text-green-900'"
                        class="transition-colors"
                        :title="!rule.paused ? 'غیرفعال کردن' : 'فعال کردن'"
                      >
                        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728L5.636 5.636m12.728 12.728L18.364 5.636M5.636 18.364l12.728-12.728"/>
                        </svg>
                      </button>
                      <button 
                        @click="deleteRule(rule)"
                        class="text-red-600 hover:text-red-900 transition-colors"
                        title="حذف"
                      >
                        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"/>
                        </svg>
                      </button>
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>

    <!-- مودال افزودن قانون جدید -->
    <div v-if="showAddRuleModal" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
      <div class="relative top-20 mx-auto p-5 border w-full max-w-md sm:max-w-lg md:max-w-xl shadow-lg rounded-md bg-white">
        <div class="mt-3">
          <h3 class="text-lg font-medium text-gray-900 mb-4">افزودن قانون WAF جدید</h3>
          <form @submit.prevent="addNewRule">
            <div class="mb-4">
              <label class="block text-sm font-medium text-gray-700 mb-2">نام قانون</label>
              <input v-model="newRule.name" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-1 focus:ring-blue-500" required>
            </div>
            <div class="mb-4">
              <label class="block text-sm font-medium text-gray-700 mb-2">توضیحات</label>
              <textarea v-model="newRule.description" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-1 focus:ring-blue-500" rows="2"></textarea>
            </div>
            <div class="mb-4">
              <label class="block text-sm font-medium text-gray-700 mb-2">عبارت فیلتر</label>
                              <input v-model="newRule.expression" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-1 focus:ring-blue-500 font-mono" placeholder="(http.request.uri.query contains &quot;SELECT&quot;)" required>
            </div>
            <div class="mb-4">
              <label class="block text-sm font-medium text-gray-700 mb-2">عملیات</label>
              <select v-model="newRule.action" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-1 focus:ring-blue-500" required>
                <option value="block">مسدود کردن</option>
                <option value="challenge">چالش</option>
                <option value="js_challenge">چالش JavaScript</option>
                <option value="managed_challenge">چالش مدیریت شده</option>
              </select>
            </div>
            <div class="flex justify-end space-x-3 space-x-reverse">
              <button type="button" @click="showAddRuleModal = false" class="px-4 py-2 text-gray-600 border border-gray-300 rounded-md hover:bg-gray-50">
                انصراف
              </button>
              <button type="submit" class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700">
                افزودن
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string }) => void
declare const useRoute: () => { params: Record<string, string>; query: Record<string, string>; path: string }
declare const useFetch: <T = unknown>(url: string, options?: { method?: string; key?: string; body?: unknown }) => Promise<{ data: { value: T }; error: { value: Error | null } }>

type WAFRule = {
  id?: string
  name: string
  description?: string
  expression?: string
  filter?: {
    expression?: string
    description?: string
  }
  action: 'block' | 'challenge' | 'allow' | 'js_challenge' | 'managed_challenge'
  paused?: boolean
  createdAt?: string
}

type WAFRulesResponse = {
  success?: boolean
  data?: WAFRule[]
}
</script>

<script setup lang="ts">
import { onMounted, ref } from 'vue';

definePageMeta({ layout: 'admin-main' })

const route = useRoute()

// متغیرهای reactive
const showAddRuleModal = ref(false)
const newRule = ref({
  name: '',
  description: '',
  expression: '',
  action: 'block'
})

// لیست قوانین واقعی از API
const wafRules = ref<WAFRule[]>([])
const isLoading = ref(false)

async function fetchWafRules() {
  isLoading.value = true
  try {
    const { data, error } = await useFetch<WAFRulesResponse>('/api/admin/cloudflare/waf/rules', { method: 'GET', key: 'cf-waf-rules' })
    if (error.value) {
      console.error(error.value)
    } else if (data.value && data.value.success && data.value.data) {
      wafRules.value = data.value.data
    }
  } finally {
    isLoading.value = false
  }
}

onMounted(fetchWafRules)

// توابع کمکی
function getActionClass(action: string) {
  switch (action) {
    case 'block':
      return 'bg-red-100 text-red-800'
    case 'challenge':
      return 'bg-yellow-100 text-yellow-800'
    case 'js_challenge':
      return 'bg-orange-100 text-orange-800'
    case 'managed_challenge':
      return 'bg-purple-100 text-purple-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

function getActionText(action: string) {
  switch (action) {
    case 'block':
      return 'مسدود کردن'
    case 'challenge':
      return 'چالش'
    case 'js_challenge':
      return 'چالش JS'
    case 'managed_challenge':
      return 'چالش مدیریت شده'
    default:
      return 'نامشخص'
  }
}

// توابع عملیاتی
async function addNewRule() {
  try {
    isLoading.value = true
    const payload = {
      action: newRule.value.action,
      description: newRule.value.description,
      filter: { expression: newRule.value.expression, description: newRule.value.name },
      paused: false,
    }
    const { error } = await useFetch<WAFRulesResponse>('/api/admin/cloudflare/waf/rules', { method: 'POST', body: payload })
    if (error.value) {
      const errorMsg = error.value instanceof Error ? error.value.message : 'خطا در ایجاد قانون'
      alert(errorMsg)
      return
    }
    showAddRuleModal.value = false
    await fetchWafRules()
    newRule.value = { name: '', description: '', expression: '', action: 'block' }
  } finally {
    isLoading.value = false
  }
}

async function editRule(rule: WAFRule) {
  if (!rule.id) return
  const newDesc = prompt('توضیحات:', rule.description || '')
  if (newDesc === null) return
  try {
    isLoading.value = true
    const { error } = await useFetch<WAFRulesResponse>(`/api/admin/cloudflare/waf/${rule.id}`, { method: 'PUT', body: { description: newDesc } })
    if (error.value) {
      const errorMsg = error.value instanceof Error ? error.value.message : 'خطا در بروزرسانی قانون'
      alert(errorMsg)
      return
    }
    await fetchWafRules()
  } finally {
    isLoading.value = false
  }
}

async function toggleRule(rule: WAFRule) {
  if (!rule.id) return
  try {
    isLoading.value = true
    const { error } = await useFetch<WAFRulesResponse>(`/api/admin/cloudflare/waf/${rule.id}`, { method: 'PUT', body: { paused: !rule.paused } })
    if (error.value) {
      const errorMsg = error.value instanceof Error ? error.value.message : 'خطا در تغییر وضعیت قانون'
      alert(errorMsg)
      return
    }
    await fetchWafRules()
  } finally {
    isLoading.value = false
  }
}

async function deleteRule(rule: WAFRule) {
  if (!confirm(`آیا از حذف قانون "${rule.description || rule.id}" اطمینان دارید؟`)) return
  try {
    isLoading.value = true
    if (!rule.id) return
    const { error } = await useFetch<WAFRulesResponse>(`/api/admin/cloudflare/waf/${rule.id}`, { method: 'DELETE' })
    if (error.value) {
      const errorMsg = error.value instanceof Error ? error.value.message : 'خطا در حذف قانون'
      alert(errorMsg)
      return
    }
    await fetchWafRules()
  } finally {
    isLoading.value = false
  }
}
</script> 

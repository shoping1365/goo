<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between">
        <div>
          <h2 class="text-xl font-bold text-gray-900">امنیت و لاگ‌ها</h2>
          <p class="text-gray-600 mt-1">نظارت امنیتی، ثبت رویدادها و مدیریت تهدیدات کیف پول</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button
            @click="activeTab = 'overview'"
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'overview'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
          >
            نمای کلی
          </button>
          <button
            @click="activeTab = 'logs'"
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'logs'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
          >
            لاگ‌ها
          </button>
          <button
            @click="activeTab = 'threats'"
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'threats'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
          >
            تهدیدات
          </button>
          <button
            @click="activeTab = 'alerts'"
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'alerts'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
          >
            هشدارها
          </button>
        </div>
      </div>
    </div>

    <!-- محتوای تب‌ها -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200">
      <!-- نمای کلی امنیت -->
      <div v-if="activeTab === 'overview'" class="p-6">
        <div class="space-y-6">
          <!-- آمار امنیتی -->
          <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
            <div class="bg-white border border-gray-200 rounded-lg p-6">
              <div class="flex items-center">
                <div class="w-8 h-8 bg-green-100 rounded-full flex items-center justify-center">
                  <svg class="w-4 h-4 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                </div>
                <div class="mr-3">
                  <p class="text-sm font-medium text-gray-900">وضعیت امنیت</p>
                  <p class="text-lg font-bold text-green-600">عالی</p>
                </div>
              </div>
            </div>
            <div class="bg-white border border-gray-200 rounded-lg p-6">
              <div class="flex items-center">
                <div class="w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center">
                  <svg class="w-4 h-4 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                </div>
                <div class="mr-3">
                  <p class="text-sm font-medium text-gray-900">آخرین بررسی</p>
                  <p class="text-lg font-bold text-blue-600">2 ساعت پیش</p>
                </div>
              </div>
            </div>
            <div class="bg-white border border-gray-200 rounded-lg p-6">
              <div class="flex items-center">
                <div class="w-8 h-8 bg-yellow-100 rounded-full flex items-center justify-center">
                  <svg class="w-4 h-4 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z" />
                  </svg>
                </div>
                <div class="mr-3">
                  <p class="text-sm font-medium text-gray-900">هشدارهای فعال</p>
                  <p class="text-lg font-bold text-yellow-600">3</p>
                </div>
              </div>
            </div>
            <div class="bg-white border border-gray-200 rounded-lg p-6">
              <div class="flex items-center">
                <div class="w-8 h-8 bg-red-100 rounded-full flex items-center justify-center">
                  <svg class="w-4 h-4 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z" />
                  </svg>
                </div>
                <div class="mr-3">
                  <p class="text-sm font-medium text-gray-900">تهدیدات شناسایی شده</p>
                  <p class="text-lg font-bold text-red-600">1</p>
                </div>
              </div>
            </div>
          </div>

          <!-- نمودار فعالیت‌های امنیتی -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">فعالیت‌های امنیتی (24 ساعت گذشته)</h3>
            <div class="h-64 flex items-end space-x-2 space-x-reverse overflow-x-auto">
              <div v-for="(hour, index) in securityActivity" :key="index" class="flex-shrink-0 flex flex-col items-center min-w-12">
                <div class="w-full bg-gray-200 rounded-t relative"
                     :style="{ height: getChartHeight(hour.value) + 'px' }">
                  <div class="w-full bg-gradient-to-t from-blue-500 to-blue-600 rounded-t transition-all duration-300 absolute bottom-0"
                       :style="{ height: getChartHeight(hour.value) + 'px' }"></div>
                </div>
                <span class="text-xs text-gray-500 mt-1 text-center">{{ hour.label }}</span>
                <span class="text-xs text-gray-400 mt-1 text-center">{{ hour.value }}</span>
              </div>
            </div>
          </div>

          <!-- خلاصه رویدادهای اخیر -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">رویدادهای امنیتی اخیر</h3>
            <div class="space-y-3">
              <div v-for="event in recentSecurityEvents" :key="event.id" class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
                <div class="flex items-center">
                  <div :class="getEventIconClass(event.type)" class="w-8 h-8 rounded-full flex items-center justify-center ml-3">
                    <i :class="getEventIcon(event.type)" class="text-white text-sm"></i>
                  </div>
                  <div>
                    <p class="font-medium text-gray-900">{{ event.title }}</p>
                    <p class="text-sm text-gray-600">{{ event.description }}</p>
                  </div>
                </div>
                <div class="text-left">
                  <p class="text-sm text-gray-500">{{ event.time }}</p>
                  <span :class="getEventStatusClass(event.status)" class="text-xs px-2 py-1 rounded-full">
                    {{ event.status }}
                  </span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- لاگ‌های سیستم -->
      <div v-if="activeTab === 'logs'" class="p-6">
        <div class="space-y-6">
          <!-- فیلترهای لاگ -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">فیلتر لاگ‌ها</h3>
            <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">سطح لاگ</label>
                <select v-model="logFilters.level" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                  <option value="">همه سطوح</option>
                  <option value="info">اطلاعات</option>
                  <option value="warning">هشدار</option>
                  <option value="error">خطا</option>
                  <option value="critical">بحرانی</option>
                </select>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">نوع رویداد</label>
                <select v-model="logFilters.type" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                  <option value="">همه انواع</option>
                  <option value="login">ورود</option>
                  <option value="transaction">تراکنش</option>
                  <option value="withdrawal">برداشت</option>
                  <option value="security">امنیت</option>
                </select>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">از تاریخ</label>
                <input v-model="logFilters.startDate" type="datetime-local" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">تا تاریخ</label>
                <input v-model="logFilters.endDate" type="datetime-local" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
              </div>
            </div>
            <div class="mt-4 flex space-x-2 space-x-reverse">
              <button @click="applyLogFilters" class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors">
                اعمال فیلتر
              </button>
              <button @click="clearLogFilters" class="bg-gray-600 text-white px-4 py-2 rounded-lg hover:bg-gray-700 transition-colors">
                پاک کردن
              </button>
              <button @click="exportLogs" class="bg-green-600 text-white px-4 py-2 rounded-lg hover:bg-green-700 transition-colors">
                خروجی
              </button>
            </div>
          </div>

          <!-- جدول لاگ‌ها -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">لاگ‌های سیستم</h3>
            <div class="overflow-x-auto">
              <table class="min-w-full divide-y divide-gray-200">
                <thead class="bg-gray-50">
                  <tr>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">زمان</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">سطح</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کاربر</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">پیام</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">IP</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
                  </tr>
                </thead>
                <tbody class="bg-white divide-y divide-gray-200">
                  <tr v-for="log in filteredLogs" :key="log.id" class="hover:bg-gray-50">
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ log.timestamp }}</td>
                    <td class="px-6 py-4 whitespace-nowrap">
                      <span :class="getLogLevelClass(log.level)" class="px-2 py-1 text-xs rounded-full">
                        {{ log.level }}
                      </span>
                    </td>
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ log.type }}</td>
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ log.user }}</td>
                    <td class="px-6 py-4 text-sm text-gray-900">{{ log.message }}</td>
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ log.ip }}</td>
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                      <button class="text-blue-600 hover:text-blue-800">جزئیات</button>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>

      <!-- مدیریت تهدیدات -->
      <div v-if="activeTab === 'threats'" class="p-6">
        <div class="space-y-6">
          <!-- تهدیدات فعال -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">تهدیدات فعال</h3>
            <div class="space-y-4">
              <div v-for="threat in activeThreats" :key="threat.id" class="border border-red-200 bg-red-50 rounded-lg p-6">
                <div class="flex items-center justify-between mb-3">
                  <div class="flex items-center">
                    <div class="w-8 h-8 bg-red-100 rounded-full flex items-center justify-center ml-3">
                      <svg class="w-4 h-4 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z" />
                      </svg>
                    </div>
                    <div>
                      <h4 class="font-medium text-red-900">{{ threat.title }}</h4>
                      <p class="text-sm text-red-700">{{ threat.description }}</p>
                    </div>
                  </div>
                  <span :class="getThreatLevelClass(threat.level)" class="px-2 py-1 text-xs rounded-full">
                    {{ threat.level }}
                  </span>
                </div>
                <div class="grid grid-cols-2 md:grid-cols-4 gap-6 text-sm">
                  <div>
                    <span class="text-red-600">شناسایی شده:</span>
                    <span class="text-red-900">{{ threat.detectedAt }}</span>
                  </div>
                  <div>
                    <span class="text-red-600">منبع:</span>
                    <span class="text-red-900">{{ threat.source }}</span>
                  </div>
                  <div>
                    <span class="text-red-600">تعداد:</span>
                    <span class="text-red-900">{{ threat.count }}</span>
                  </div>
                  <div>
                    <span class="text-red-600">وضعیت:</span>
                    <span class="text-red-900">{{ threat.status }}</span>
                  </div>
                </div>
                <div class="mt-3 flex space-x-2 space-x-reverse">
                  <button class="bg-red-600 text-white px-3 py-2 rounded text-sm hover:bg-red-700 transition-colors">
                    مسدود کردن
                  </button>
                  <button class="bg-yellow-600 text-white px-3 py-2 rounded text-sm hover:bg-yellow-700 transition-colors">
                    بررسی
                  </button>
                  <button class="bg-gray-600 text-white px-3 py-2 rounded text-sm hover:bg-gray-700 transition-colors">
                    نادیده گرفتن
                  </button>
                </div>
              </div>
            </div>
          </div>

          <!-- آمار تهدیدات -->
          <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
            <div class="border border-gray-200 rounded-lg p-6">
              <h4 class="font-medium text-gray-900 mb-3">تهدیدات مسدود شده</h4>
              <div class="text-3xl font-bold text-green-600 mb-2">1,247</div>
              <p class="text-sm text-gray-600">در 24 ساعت گذشته</p>
            </div>
            <div class="border border-gray-200 rounded-lg p-6">
              <h4 class="font-medium text-gray-900 mb-3">تلاش‌های ورود ناموفق</h4>
              <div class="text-3xl font-bold text-yellow-600 mb-2">89</div>
              <p class="text-sm text-gray-600">در 24 ساعت گذشته</p>
            </div>
            <div class="border border-gray-200 rounded-lg p-6">
              <h4 class="font-medium text-gray-900 mb-3">IP های مسدود شده</h4>
              <div class="text-3xl font-bold text-red-600 mb-2">156</div>
              <p class="text-sm text-gray-600">کل مسدودیت‌ها</p>
            </div>
          </div>
        </div>
      </div>

      <!-- مدیریت هشدارها -->
      <div v-if="activeTab === 'alerts'" class="p-6">
        <div class="space-y-6">
          <!-- تنظیمات هشدار -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات هشدار</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <h4 class="font-medium text-gray-900 mb-3">هشدارهای امنیتی</h4>
                <div class="space-y-3">
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="text-sm font-medium text-gray-900">تلاش ورود ناموفق</p>
                      <p class="text-xs text-gray-600">هشدار پس از 5 تلاش ناموفق</p>
                    </div>
                    <label class="relative inline-flex items-center cursor-pointer">
                      <input type="checkbox" v-model="alertSettings.failedLogin" class="sr-only peer">
                      <div class="w-9 h-5 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-4 after:w-4 after:transition-all peer-checked:bg-blue-600"></div>
                    </label>
                  </div>
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="text-sm font-medium text-gray-900">تراکنش‌های مشکوک</p>
                      <p class="text-xs text-gray-600">تراکنش‌های بالای 10 میلیون تومان</p>
                    </div>
                    <label class="relative inline-flex items-center cursor-pointer">
                      <input type="checkbox" v-model="alertSettings.suspiciousTransaction" class="sr-only peer">
                      <div class="w-9 h-5 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-4 after:w-4 after:transition-all peer-checked:bg-blue-600"></div>
                    </label>
                  </div>
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="text-sm font-medium text-gray-900">ورود از IP جدید</p>
                      <p class="text-xs text-gray-600">ورود از آدرس IP جدید</p>
                    </div>
                    <label class="relative inline-flex items-center cursor-pointer">
                      <input type="checkbox" v-model="alertSettings.newIpLogin" class="sr-only peer">
                      <div class="w-9 h-5 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-4 after:w-4 after:transition-all peer-checked:bg-blue-600"></div>
                    </label>
                  </div>
                </div>
              </div>
              <div>
                <h4 class="font-medium text-gray-900 mb-3">کانال‌های هشدار</h4>
                <div class="space-y-3">
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="text-sm font-medium text-gray-900">ایمیل</p>
                      <p class="text-xs text-gray-600">ارسال هشدار به ایمیل</p>
                    </div>
                    <label class="relative inline-flex items-center cursor-pointer">
                      <input type="checkbox" v-model="alertSettings.email" class="sr-only peer">
                      <div class="w-9 h-5 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-4 after:w-4 after:transition-all peer-checked:bg-blue-600"></div>
                    </label>
                  </div>
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="text-sm font-medium text-gray-900">پیامک</p>
                      <p class="text-xs text-gray-600">ارسال هشدار به پیامک</p>
                    </div>
                    <label class="relative inline-flex items-center cursor-pointer">
                      <input type="checkbox" v-model="alertSettings.sms" class="sr-only peer">
                      <div class="w-9 h-5 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-4 after:w-4 after:transition-all peer-checked:bg-blue-600"></div>
                    </label>
                  </div>
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="text-sm font-medium text-gray-900">اعلان درون‌برنامه‌ای</p>
                      <p class="text-xs text-gray-600">نمایش اعلان در پنل</p>
                    </div>
                    <label class="relative inline-flex items-center cursor-pointer">
                      <input type="checkbox" v-model="alertSettings.inApp" class="sr-only peer">
                      <div class="w-9 h-5 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-4 after:w-4 after:transition-all peer-checked:bg-blue-600"></div>
                    </label>
                  </div>
                </div>
              </div>
            </div>
            <div class="mt-6">
              <button class="bg-blue-600 text-white px-6 py-2 rounded-lg hover:bg-blue-700 transition-colors">
                ذخیره تنظیمات
              </button>
            </div>
          </div>

          <!-- تاریخچه هشدارها -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">تاریخچه هشدارها</h3>
            <div class="space-y-3">
              <div v-for="alert in alertHistory" :key="alert.id" class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
                <div class="flex items-center">
                  <div :class="getAlertIconClass(alert.type)" class="w-8 h-8 rounded-full flex items-center justify-center ml-3">
                    <i :class="getAlertIcon(alert.type)" class="text-white text-sm"></i>
                  </div>
                  <div>
                    <p class="font-medium text-gray-900">{{ alert.title }}</p>
                    <p class="text-sm text-gray-600">{{ alert.description }}</p>
                  </div>
                </div>
                <div class="text-left">
                  <p class="text-sm text-gray-500">{{ alert.time }}</p>
                  <span :class="getAlertStatusClass(alert.status)" class="text-xs px-2 py-1 rounded-full">
                    {{ alert.status }}
                  </span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

// تعریف reactive data
const activeTab = ref('overview')

// آمار امنیتی
const securityStats = ref({
  status: 'عالی',
  lastCheck: '2 ساعت پیش',
  activeAlerts: 3,
  detectedThreats: 1
})

// فعالیت‌های امنیتی (24 ساعت)
const securityActivity = ref([
  { label: '00', value: 12 },
  { label: '02', value: 8 },
  { label: '04', value: 15 },
  { label: '06', value: 22 },
  { label: '08', value: 35 },
  { label: '10', value: 28 },
  { label: '12', value: 42 },
  { label: '14', value: 38 },
  { label: '16', value: 45 },
  { label: '18', value: 33 },
  { label: '20', value: 25 },
  { label: '22', value: 18 }
])

// رویدادهای امنیتی اخیر
const recentSecurityEvents = ref([
  {
    id: 1,
    type: 'login',
    title: 'ورود موفق کاربر',
    description: 'کاربر احمد محمدی با موفقیت وارد شد',
    time: '2 دقیقه پیش',
    status: 'موفق'
  },
  {
    id: 2,
    type: 'security',
    title: 'تلاش ورود ناموفق',
    description: '5 تلاش ناموفق از IP 192.168.1.100',
    time: '15 دقیقه پیش',
    status: 'هشدار'
  },
  {
    id: 3,
    type: 'transaction',
    title: 'تراکنش مشکوک',
    description: 'تراکنش 50 میلیون تومانی از حساب کاربری',
    time: '1 ساعت پیش',
    status: 'بررسی'
  }
])

// فیلترهای لاگ
const logFilters = ref({
  level: '',
  type: '',
  startDate: '',
  endDate: ''
})

// لاگ‌های سیستم
const systemLogs = ref([
  {
    id: 1,
    timestamp: '2024-01-31 14:30:25',
    level: 'info',
    type: 'login',
    user: 'احمد محمدی',
    message: 'ورود موفق به سیستم',
    ip: '192.168.1.100'
  },
  {
    id: 2,
    timestamp: '2024-01-31 14:25:10',
    level: 'warning',
    type: 'security',
    user: 'نامشخص',
    message: 'تلاش ورود ناموفق',
    ip: '192.168.1.101'
  },
  {
    id: 3,
    timestamp: '2024-01-31 14:20:15',
    level: 'error',
    type: 'transaction',
    user: 'فاطمه احمدی',
    message: 'خطا در پردازش تراکنش',
    ip: '192.168.1.102'
  }
])

// لاگ‌های فیلتر شده
const filteredLogs = computed(() => {
  let logs = systemLogs.value
  
  if (logFilters.value.level) {
    logs = logs.filter(log => log.level === logFilters.value.level)
  }
  
  if (logFilters.value.type) {
    logs = logs.filter(log => log.type === logFilters.value.type)
  }
  
  return logs
})

// تهدیدات فعال
const activeThreats = ref([
  {
    id: 1,
    title: 'حمله Brute Force',
    description: 'تلاش‌های مکرر ورود از IP 192.168.1.100',
    level: 'بالا',
    detectedAt: '2024-01-31 14:00',
    source: '192.168.1.100',
    count: 25,
    status: 'فعال'
  }
])

// تنظیمات هشدار
const alertSettings = ref({
  failedLogin: true,
  suspiciousTransaction: true,
  newIpLogin: false,
  email: true,
  sms: false,
  inApp: true
})

// تاریخچه هشدارها
const alertHistory = ref([
  {
    id: 1,
    type: 'security',
    title: 'هشدار امنیتی',
    description: 'تلاش ورود ناموفق از IP جدید',
    time: '10 دقیقه پیش',
    status: 'خوانده نشده'
  },
  {
    id: 2,
    type: 'transaction',
    title: 'تراکنش مشکوک',
    description: 'تراکنش بزرگ شناسایی شد',
    time: '1 ساعت پیش',
    status: 'بررسی شده'
  }
])

// توابع کمکی
const getChartHeight = (value: number) => {
  const maxValue = Math.max(...securityActivity.value.map(item => item.value))
  const height = 200
  return (value / maxValue) * height
}

const getEventIconClass = (type: string) => {
  const classes = {
    login: 'bg-green-100',
    security: 'bg-red-100',
    transaction: 'bg-yellow-100'
  }
  return classes[type] || 'bg-gray-100'
}

const getEventIcon = (type: string) => {
  const icons = {
    login: 'fas fa-sign-in-alt',
    security: 'fas fa-shield-alt',
    transaction: 'fas fa-exchange-alt'
  }
  return icons[type] || 'fas fa-info-circle'
}

const getEventStatusClass = (status: string) => {
  const classes = {
    'موفق': 'bg-green-100 text-green-800',
    'هشدار': 'bg-yellow-100 text-yellow-800',
    'بررسی': 'bg-blue-100 text-blue-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getLogLevelClass = (level: string) => {
  const classes = {
    info: 'bg-blue-100 text-blue-800',
    warning: 'bg-yellow-100 text-yellow-800',
    error: 'bg-red-100 text-red-800',
    critical: 'bg-purple-100 text-purple-800'
  }
  return classes[level] || 'bg-gray-100 text-gray-800'
}

const getThreatLevelClass = (level: string) => {
  const classes = {
    'کم': 'bg-green-100 text-green-800',
    'متوسط': 'bg-yellow-100 text-yellow-800',
    'بالا': 'bg-red-100 text-red-800'
  }
  return classes[level] || 'bg-gray-100 text-gray-800'
}

const getAlertIconClass = (type: string) => {
  const classes = {
    security: 'bg-red-100',
    transaction: 'bg-yellow-100'
  }
  return classes[type] || 'bg-gray-100'
}

const getAlertIcon = (type: string) => {
  const icons = {
    security: 'fas fa-shield-alt',
    transaction: 'fas fa-exchange-alt'
  }
  return icons[type] || 'fas fa-bell'
}

const getAlertStatusClass = (status: string) => {
  const classes = {
    'خوانده نشده': 'bg-red-100 text-red-800',
    'بررسی شده': 'bg-green-100 text-green-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

// توابع عملیاتی
const applyLogFilters = () => {
  // اعمال فیلترها
  console.log('فیلترها اعمال شدند')
}

const clearLogFilters = () => {
  logFilters.value = {
    level: '',
    type: '',
    startDate: '',
    endDate: ''
  }
}

const exportLogs = () => {
  // خروجی لاگ‌ها
  console.log('لاگ‌ها صادر شدند')
}
</script> 

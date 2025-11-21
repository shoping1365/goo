<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between">
        <div>
          <h2 class="text-xl font-bold text-gray-900">مدیریت اعلان‌ها</h2>
          <p class="text-gray-600 mt-1">تنظیمات و مدیریت اعلان‌های کیف پول</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'overview'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
            @click="activeTab = 'overview'"
          >
            نمای کلی
          </button>
          <button
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'settings'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
            @click="activeTab = 'settings'"
          >
            تنظیمات
          </button>
          <button
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'templates'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
            @click="activeTab = 'templates'"
          >
            قالب‌ها
          </button>
          <button
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'history'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
            @click="activeTab = 'history'"
          >
            تاریخچه
          </button>
        </div>
      </div>
    </div>

    <!-- محتوای تب‌ها -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200">
      <!-- نمای کلی -->
      <div v-if="activeTab === 'overview'" class="p-6">
        <div class="space-y-6">
          <!-- آمار اعلان‌ها -->
          <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
            <div class="bg-white border border-gray-200 rounded-lg p-6">
              <div class="flex items-center">
                <div class="w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center">
                  <svg class="w-4 h-4 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-5 5v-5zM4.19 4.19A4 4 0 004 6v6a4 4 0 004 4h6a4 4 0 004-4V6a4 4 0 00-4-4H8a4 4 0 00-2.83 1.17z" />
                  </svg>
                </div>
                <div class="mr-3">
                  <p class="text-sm font-medium text-gray-900">اعلان‌های امروز</p>
                  <p class="text-lg font-bold text-blue-600">{{ notificationStats.todayNotifications }}</p>
                </div>
              </div>
            </div>
            <div class="bg-white border border-gray-200 rounded-lg p-6">
              <div class="flex items-center">
                <div class="w-8 h-8 bg-green-100 rounded-full flex items-center justify-center">
                  <svg class="w-4 h-4 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                  </svg>
                </div>
                <div class="mr-3">
                  <p class="text-sm font-medium text-gray-900">ارسال موفق</p>
                  <p class="text-lg font-bold text-green-600">{{ notificationStats.successRate }}%</p>
                </div>
              </div>
            </div>
            <div class="bg-white border border-gray-200 rounded-lg p-6">
              <div class="flex items-center">
                <div class="w-8 h-8 bg-purple-100 rounded-full flex items-center justify-center">
                  <svg class="w-4 h-4 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                </div>
                <div class="mr-3">
                  <p class="text-sm font-medium text-gray-900">در صف ارسال</p>
                  <p class="text-lg font-bold text-purple-600">{{ notificationStats.queuedNotifications }}</p>
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
                  <p class="text-sm font-medium text-gray-900">خطاهای ارسال</p>
                  <p class="text-lg font-bold text-yellow-600">{{ notificationStats.failedNotifications }}</p>
                </div>
              </div>
            </div>
          </div>

          <!-- نمودار اعلان‌ها -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">روند اعلان‌ها (7 روز گذشته)</h3>
            <div class="h-64 flex items-end space-x-2 space-x-reverse overflow-x-auto">
              <div v-for="(day, index) in notificationTrend" :key="index" class="flex-shrink-0 flex flex-col items-center min-w-16">
                <div
class="w-full bg-gray-200 rounded-t relative"
                     :style="{ height: getChartHeight(day.count) + 'px' }">
                  <div
class="w-full bg-gradient-to-t from-blue-500 to-blue-600 rounded-t transition-all duration-300 absolute bottom-0"
                       :style="{ height: getChartHeight(day.count) + 'px' }"></div>
                </div>
                <span class="text-xs text-gray-500 mt-1 text-center">{{ day.date }}</span>
                <span class="text-xs text-gray-400 mt-1 text-center">{{ day.count }}</span>
              </div>
            </div>
          </div>

          <!-- اعلان‌های اخیر -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">اعلان‌های اخیر</h3>
            <div class="space-y-3">
              <div v-for="notification in recentNotifications" :key="notification.id" class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
                <div class="flex items-center">
                  <div :class="getNotificationIconClass(notification.type)" class="w-8 h-8 rounded-full flex items-center justify-center ml-3">
                    <i :class="getNotificationIcon(notification.type)" class="text-white text-sm"></i>
                  </div>
                  <div>
                    <p class="font-medium text-gray-900">{{ notification.title }}</p>
                    <p class="text-sm text-gray-600">{{ notification.message }}</p>
                  </div>
                </div>
                <div class="text-left">
                  <p class="text-sm text-gray-500">{{ notification.time }}</p>
                  <span :class="getNotificationStatusClass(notification.status)" class="text-xs px-2 py-1 rounded-full">
                    {{ notification.status }}
                  </span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- تنظیمات -->
      <div v-if="activeTab === 'settings'" class="p-6">
        <div class="space-y-6">
          <!-- تنظیمات کانال‌های ارسال -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات کانال‌های ارسال</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <h4 class="font-medium text-gray-900 mb-3">ایمیل</h4>
                <div class="space-y-3">
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="text-sm font-medium text-gray-900">فعال کردن ایمیل</p>
                      <p class="text-xs text-gray-600">ارسال اعلان‌ها از طریق ایمیل</p>
                    </div>
                    <label class="relative inline-flex items-center cursor-pointer">
                      <input v-model="notificationSettings.email.enabled" type="checkbox" class="sr-only peer">
                      <div class="w-9 h-5 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-4 after:w-4 after:transition-all peer-checked:bg-blue-600"></div>
                    </label>
                  </div>
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">SMTP سرور</label>
                    <input v-model="notificationSettings.email.smtpServer" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="smtp.gmail.com">
                  </div>
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">پورت</label>
                    <input v-model="notificationSettings.email.port" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="587">
                  </div>
                </div>
              </div>
              <div>
                <h4 class="font-medium text-gray-900 mb-3">پیامک</h4>
                <div class="space-y-3">
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="text-sm font-medium text-gray-900">فعال کردن پیامک</p>
                      <p class="text-xs text-gray-600">ارسال اعلان‌ها از طریق پیامک</p>
                    </div>
                    <label class="relative inline-flex items-center cursor-pointer">
                      <input v-model="notificationSettings.sms.enabled" type="checkbox" class="sr-only peer">
                      <div class="w-9 h-5 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-4 after:w-4 after:transition-all peer-checked:bg-blue-600"></div>
                    </label>
                  </div>
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">API کلید</label>
                    <input v-model="notificationSettings.sms.apiKey" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="API کلید پیامک">
                  </div>
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">شماره فرستنده</label>
                    <input v-model="notificationSettings.sms.senderNumber" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="09123456789">
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- تنظیمات اعلان‌ها -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات اعلان‌ها</h3>
            <div class="space-y-4">
              <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
                <div>
                  <h4 class="font-medium text-gray-900">اعلان تراکنش</h4>
                  <p class="text-sm text-gray-600">اعلان برای تراکنش‌های جدید</p>
                </div>
                <label class="relative inline-flex items-center cursor-pointer">
                  <input v-model="notificationSettings.transaction.enabled" type="checkbox" class="sr-only peer">
                  <div class="w-9 h-5 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-4 after:w-4 after:transition-all peer-checked:bg-blue-600"></div>
                </label>
              </div>
              <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
                <div>
                  <h4 class="font-medium text-gray-900">اعلان برداشت</h4>
                  <p class="text-sm text-gray-600">اعلان برای درخواست‌های برداشت</p>
                </div>
                <label class="relative inline-flex items-center cursor-pointer">
                  <input v-model="notificationSettings.withdrawal.enabled" type="checkbox" class="sr-only peer">
                  <div class="w-9 h-5 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-4 after:w-4 after:transition-all peer-checked:bg-blue-600"></div>
                </label>
              </div>
              <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
                <div>
                  <h4 class="font-medium text-gray-900">اعلان امنیت</h4>
                  <p class="text-sm text-gray-600">اعلان برای رویدادهای امنیتی</p>
                </div>
                <label class="relative inline-flex items-center cursor-pointer">
                  <input v-model="notificationSettings.security.enabled" type="checkbox" class="sr-only peer">
                  <div class="w-9 h-5 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-4 after:w-4 after:transition-all peer-checked:bg-blue-600"></div>
                </label>
              </div>
            </div>
            <div class="mt-6">
              <button class="bg-blue-600 text-white px-6 py-2 rounded-lg hover:bg-blue-700 transition-colors" @click="saveNotificationSettings">
                ذخیره تنظیمات
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- قالب‌ها -->
      <div v-if="activeTab === 'templates'" class="p-6">
        <div class="space-y-6">
          <!-- ایجاد قالب جدید -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">ایجاد قالب جدید</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">نام قالب</label>
                <input v-model="newTemplate.name" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="نام قالب">
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">نوع قالب</label>
                <select v-model="newTemplate.type" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                  <option value="">انتخاب کنید</option>
                  <option value="email">ایمیل</option>
                  <option value="sms">پیامک</option>
                  <option value="push">اعلان درون‌برنامه‌ای</option>
                </select>
              </div>
              <div class="md:col-span-2">
                <label class="block text-sm font-medium text-gray-700 mb-1">موضوع</label>
                <input v-model="newTemplate.subject" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="موضوع اعلان">
              </div>
              <div class="md:col-span-2">
                <label class="block text-sm font-medium text-gray-700 mb-1">محتوا</label>
                <textarea v-model="newTemplate.content" rows="6" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="محتوی قالب اعلان"></textarea>
                <p class="text-xs text-gray-500 mt-1">متغیرهای قابل استفاده: &#123;&#123;user_name&#125;&#125;, &#123;&#123;amount&#125;&#125;, &#123;&#123;transaction_id&#125;&#125;, &#123;&#123;date&#125;&#125;</p>
              </div>
              <div class="md:col-span-2">
                <button class="bg-blue-600 text-white px-6 py-2 rounded-lg hover:bg-blue-700 transition-colors" @click="createTemplate">
                  ایجاد قالب
                </button>
              </div>
            </div>
          </div>

          <!-- جدول قالب‌ها -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">قالب‌های اعلان</h3>
            <div class="overflow-x-auto">
              <table class="min-w-full divide-y divide-gray-200">
                <thead class="bg-gray-50">
                  <tr>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نام</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">موضوع</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
                  </tr>
                </thead>
                <tbody class="bg-white divide-y divide-gray-200">
                  <tr v-for="template in notificationTemplates" :key="template.id" class="hover:bg-gray-50">
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ template.name }}</td>
                    <td class="px-6 py-4 whitespace-nowrap">
                      <span :class="getTemplateTypeClass(template.type)" class="px-2 py-1 text-xs rounded-full">
                        {{ template.type }}
                      </span>
                    </td>
                    <td class="px-6 py-4 text-sm text-gray-900">{{ template.subject }}</td>
                    <td class="px-6 py-4 whitespace-nowrap">
                      <span :class="getTemplateStatusClass(template.status)" class="px-2 py-1 text-xs rounded-full">
                        {{ template.status }}
                      </span>
                    </td>
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                      <div class="flex space-x-2 space-x-reverse">
                        <button class="text-blue-600 hover:text-blue-800">ویرایش</button>
                        <button class="text-green-600 hover:text-green-800">تست</button>
                        <button class="text-red-600 hover:text-red-800">حذف</button>
                      </div>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>

      <!-- تاریخچه -->
      <div v-if="activeTab === 'history'" class="p-6">
        <div class="space-y-6">
          <!-- فیلترهای تاریخچه -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">فیلتر تاریخچه</h3>
            <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">نوع اعلان</label>
                <select v-model="historyFilters.type" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                  <option value="">همه انواع</option>
                  <option value="email">ایمیل</option>
                  <option value="sms">پیامک</option>
                  <option value="push">اعلان درون‌برنامه‌ای</option>
                </select>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">وضعیت</label>
                <select v-model="historyFilters.status" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                  <option value="">همه وضعیت‌ها</option>
                  <option value="sent">ارسال شده</option>
                  <option value="failed">ناموفق</option>
                  <option value="pending">در انتظار</option>
                </select>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">از تاریخ</label>
                <input v-model="historyFilters.startDate" type="datetime-local" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">تا تاریخ</label>
                <input v-model="historyFilters.endDate" type="datetime-local" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
              </div>
            </div>
            <div class="mt-4 flex space-x-2 space-x-reverse">
              <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors" @click="applyHistoryFilters">
                اعمال فیلتر
              </button>
              <button class="bg-gray-600 text-white px-4 py-2 rounded-lg hover:bg-gray-700 transition-colors" @click="clearHistoryFilters">
                پاک کردن
              </button>
            </div>
          </div>

          <!-- جدول تاریخچه -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">تاریخچه اعلان‌ها</h3>
            <div class="overflow-x-auto">
              <table class="min-w-full divide-y divide-gray-200">
                <thead class="bg-gray-50">
                  <tr>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">زمان</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">گیرنده</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">موضوع</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
                  </tr>
                </thead>
                <tbody class="bg-white divide-y divide-gray-200">
                  <tr v-for="notification in filteredHistory" :key="notification.id" class="hover:bg-gray-50">
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ notification.timestamp }}</td>
                    <td class="px-6 py-4 whitespace-nowrap">
                      <span :class="getNotificationTypeClass(notification.type)" class="px-2 py-1 text-xs rounded-full">
                        {{ notification.type }}
                      </span>
                    </td>
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ notification.recipient }}</td>
                    <td class="px-6 py-4 text-sm text-gray-900">{{ notification.subject }}</td>
                    <td class="px-6 py-4 whitespace-nowrap">
                      <span :class="getNotificationStatusClass(notification.status)" class="px-2 py-1 text-xs rounded-full">
                        {{ notification.status }}
                      </span>
                    </td>
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
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'

// تعریف reactive data
const activeTab = ref('overview')

// آمار اعلان‌ها
const notificationStats = ref({
  todayNotifications: 1250,
  successRate: 98.5,
  queuedNotifications: 45,
  failedNotifications: 12
})

// روند اعلان‌ها (7 روز)
const notificationTrend = ref([
  { date: 'شنبه', count: 180 },
  { date: 'یکشنبه', count: 220 },
  { date: 'دوشنبه', count: 195 },
  { date: 'سه‌شنبه', count: 250 },
  { date: 'چهارشنبه', count: 280 },
  { date: 'پنج‌شنبه', count: 320 },
  { date: 'جمعه', count: 290 }
])

// اعلان‌های اخیر
const recentNotifications = ref([
  {
    id: 1,
    type: 'transaction',
    title: 'تراکنش جدید',
    message: 'تراکنش 500,000 تومانی با موفقیت انجام شد',
    time: '2 دقیقه پیش',
    status: 'ارسال شده'
  },
  {
    id: 2,
    type: 'withdrawal',
    title: 'درخواست برداشت',
    message: 'درخواست برداشت 2,000,000 تومانی ثبت شد',
    time: '5 دقیقه پیش',
    status: 'در انتظار'
  },
  {
    id: 3,
    type: 'security',
    title: 'هشدار امنیت',
    message: 'ورود از IP جدید شناسایی شد',
    time: '10 دقیقه پیش',
    status: 'ارسال شده'
  }
])

// تنظیمات اعلان‌ها
const notificationSettings = ref({
  email: {
    enabled: true,
    smtpServer: 'smtp.gmail.com',
    port: 587
  },
  sms: {
    enabled: true,
    apiKey: '',
    senderNumber: ''
  },
  transaction: {
    enabled: true
  },
  withdrawal: {
    enabled: true
  },
  security: {
    enabled: true
  }
})

// قالب جدید
const newTemplate = ref({
  name: '',
  type: '',
  subject: '',
  content: ''
})

// قالب‌های اعلان
const notificationTemplates = ref([
  {
    id: 1,
    name: 'قالب تراکنش',
    type: 'ایمیل',
    subject: 'تراکنش جدید انجام شد',
    status: 'فعال'
  },
  {
    id: 2,
    name: 'قالب برداشت',
    type: 'پیامک',
    subject: 'درخواست برداشت شما ثبت شد',
    status: 'فعال'
  },
  {
    id: 3,
    name: 'قالب امنیت',
    type: 'اعلان درون‌برنامه‌ای',
    subject: 'هشدار امنیتی',
    status: 'غیرفعال'
  }
])

// فیلترهای تاریخچه
const historyFilters = ref({
  type: '',
  status: '',
  startDate: '',
  endDate: ''
})

// تاریخچه اعلان‌ها
const notificationHistory = ref([
  {
    id: 1,
    timestamp: '2024-01-31 14:30:25',
    type: 'ایمیل',
    recipient: 'user@example.com',
    subject: 'تراکنش جدید انجام شد',
    status: 'ارسال شده'
  },
  {
    id: 2,
    timestamp: '2024-01-31 14:25:10',
    type: 'پیامک',
    recipient: '09123456789',
    subject: 'درخواست برداشت شما ثبت شد',
    status: 'ارسال شده'
  },
  {
    id: 3,
    timestamp: '2024-01-31 14:20:15',
    type: 'اعلان درون‌برنامه‌ای',
    recipient: 'کاربر احمد',
    subject: 'هشدار امنیتی',
    status: 'ناموفق'
  }
])

// تاریخچه فیلتر شده
const filteredHistory = computed(() => {
  let history = notificationHistory.value
  
  if (historyFilters.value.type) {
    history = history.filter(item => item.type === historyFilters.value.type)
  }
  
  if (historyFilters.value.status) {
    history = history.filter(item => item.status === historyFilters.value.status)
  }
  
  return history
})

// توابع کمکی
const getChartHeight = (value: number) => {
  const maxValue = Math.max(...notificationTrend.value.map(item => item.count))
  const height = 200
  return (value / maxValue) * height
}

const getNotificationIconClass = (type: string) => {
  const classes = {
    transaction: 'bg-green-100',
    withdrawal: 'bg-blue-100',
    security: 'bg-red-100'
  }
  return classes[type] || 'bg-gray-100'
}

const getNotificationIcon = (type: string) => {
  const icons = {
    transaction: 'fas fa-exchange-alt',
    withdrawal: 'fas fa-money-bill-wave',
    security: 'fas fa-shield-alt'
  }
  return icons[type] || 'fas fa-bell'
}

const getNotificationStatusClass = (status: string) => {
  const classes = {
    'ارسال شده': 'bg-green-100 text-green-800',
    'در انتظار': 'bg-yellow-100 text-yellow-800',
    'ناموفق': 'bg-red-100 text-red-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getTemplateTypeClass = (type: string) => {
  const classes = {
    'ایمیل': 'bg-blue-100 text-blue-800',
    'پیامک': 'bg-green-100 text-green-800',
    'اعلان درون‌برنامه‌ای': 'bg-purple-100 text-purple-800'
  }
  return classes[type] || 'bg-gray-100 text-gray-800'
}

const getTemplateStatusClass = (status: string) => {
  const classes = {
    'فعال': 'bg-green-100 text-green-800',
    'غیرفعال': 'bg-red-100 text-red-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getNotificationTypeClass = (type: string) => {
  const classes = {
    'ایمیل': 'bg-blue-100 text-blue-800',
    'پیامک': 'bg-green-100 text-green-800',
    'اعلان درون‌برنامه‌ای': 'bg-purple-100 text-purple-800'
  }
  return classes[type] || 'bg-gray-100 text-gray-800'
}

// توابع عملیاتی
const saveNotificationSettings = () => {

}

const createTemplate = () => {

}

const applyHistoryFilters = () => {

}

const clearHistoryFilters = () => {
  historyFilters.value = {
    type: '',
    status: '',
    startDate: '',
    endDate: ''
  }
}
</script> 

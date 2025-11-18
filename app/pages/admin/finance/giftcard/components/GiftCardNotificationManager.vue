<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex justify-between items-center">
        <div>
          <h2 class="text-xl font-bold text-gray-900">مدیریت اطلاع‌رسانی گیفت کارت</h2>
          <p class="text-gray-600 mt-1">مدیریت ایمیل‌ها و پیامک‌های خودکار سیستم گیفت کارت</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button 
            @click="testNotification"
            class="px-4 py-2 bg-yellow-600 text-white text-sm font-medium rounded-lg hover:bg-yellow-700 focus:outline-none focus:ring-2 focus:ring-yellow-500 focus:ring-offset-2"
          >
            تست اطلاع‌رسانی
          </button>
          <button 
            @click="saveSettings"
            :disabled="isSaving"
            class="px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed"
          >
            <span v-if="isSaving">در حال ذخیره...</span>
            <span v-else>ذخیره تنظیمات</span>
          </button>
        </div>
      </div>
    </div>

    <!-- آمار اطلاع‌رسانی -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-blue-100 rounded-lg flex items-center justify-center">
              <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 4.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">ایمیل‌های ارسال شده</p>
            <p class="text-2xl font-bold text-gray-900">{{ emailStats.sent }}</p>
            <p class="text-xs text-green-600">{{ emailStats.successRate }}% نرخ موفقیت</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-green-100 rounded-lg flex items-center justify-center">
              <svg class="w-5 h-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 18h.01M8 21h8a2 2 0 002-2V5a2 2 0 00-2-2H8a2 2 0 00-2 2v14a2 2 0 002 2z"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">پیامک‌های ارسال شده</p>
            <p class="text-2xl font-bold text-gray-900">{{ smsStats.sent }}</p>
            <p class="text-xs text-green-600">{{ smsStats.successRate }}% نرخ موفقیت</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-yellow-100 rounded-lg flex items-center justify-center">
              <svg class="w-5 h-5 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">در انتظار ارسال</p>
            <p class="text-2xl font-bold text-gray-900">{{ pendingNotifications }}</p>
            <p class="text-xs text-yellow-600">در صف ارسال</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-red-100 rounded-lg flex items-center justify-center">
              <svg class="w-5 h-5 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">خطاهای ارسال</p>
            <p class="text-2xl font-bold text-gray-900">{{ failedNotifications }}</p>
            <p class="text-xs text-red-600">نیاز به بررسی</p>
          </div>
        </div>
      </div>
    </div>

    <!-- تب‌های مدیریت -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200">
      <div class="border-b border-gray-200">
        <nav class="-mb-px flex space-x-8 space-x-reverse px-6">
          <button
            @click="activeTab = 'email'"
            :class="{
              'border-blue-500 text-blue-600': activeTab === 'email',
              'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300': activeTab !== 'email'
            }"
            class="whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm"
          >
            مدیریت ایمیل‌ها
          </button>
          <button
            @click="activeTab = 'sms'"
            :class="{
              'border-blue-500 text-blue-600': activeTab === 'sms',
              'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300': activeTab !== 'sms'
            }"
            class="whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm"
          >
            مدیریت پیامک‌ها
          </button>
          <button
            @click="activeTab = 'templates'"
            :class="{
              'border-blue-500 text-blue-600': activeTab === 'templates',
              'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300': activeTab !== 'templates'
            }"
            class="whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm"
          >
            قالب‌های پیام
          </button>
          <button
            @click="activeTab = 'settings'"
            :class="{
              'border-blue-500 text-blue-600': activeTab === 'settings',
              'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300': activeTab !== 'settings'
            }"
            class="whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm"
          >
            تنظیمات ارسال
          </button>
        </nav>
      </div>

      <!-- محتوای تب‌ها -->
      <div class="p-6">
        <!-- تب مدیریت ایمیل‌ها -->
        <div v-if="activeTab === 'email'" class="space-y-6">
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <!-- تنظیمات ایمیل -->
            <div class="space-y-6">
              <div class="bg-gray-50 rounded-lg p-6">
                <h3 class="text-lg font-medium text-gray-900 mb-4">تنظیمات ارسال ایمیل</h3>
                <div class="space-y-4">
                  <div class="flex items-center justify-between">
                    <div>
                      <label class="text-sm font-medium text-gray-700">فعال کردن ارسال ایمیل</label>
                      <p class="text-xs text-gray-500">ارسال خودکار ایمیل‌های اطلاع‌رسانی</p>
                    </div>
                    <Switch
                      v-model="emailSettings.enabled"
                      :class="emailSettings.enabled ? 'bg-blue-600' : 'bg-gray-200'"
                      class="relative inline-flex h-6 w-11 items-center rounded-full transition-colors focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
                    >
                      <span
                        :class="emailSettings.enabled ? 'translate-x-6' : 'translate-x-1'"
                        class="inline-block h-4 w-4 transform rounded-full bg-white transition-transform"
                      />
                    </Switch>
                  </div>
                  
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">آدرس ایمیل فرستنده</label>
                    <input
                      v-model="emailSettings.senderEmail"
                      type="email"
                      class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                      placeholder="noreply@example.com"
                    />
                  </div>
                  
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">نام فرستنده</label>
                    <input
                      v-model="emailSettings.senderName"
                      type="text"
                      class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                      placeholder="سیستم گیفت کارت"
                    />
                  </div>
                </div>
              </div>

              <!-- انواع ایمیل‌ها -->
              <div class="bg-gray-50 rounded-lg p-6">
                <h3 class="text-lg font-medium text-gray-900 mb-4">انواع ایمیل‌های ارسالی</h3>
                <div class="space-y-4">
                  <div class="flex items-center justify-between">
                    <div>
                      <label class="text-sm font-medium text-gray-700">تأیید خرید</label>
                      <p class="text-xs text-gray-500">ارسال ایمیل تأیید پس از خرید کارت</p>
                    </div>
                    <Switch
                      v-model="emailSettings.purchaseConfirmation"
                      :class="emailSettings.purchaseConfirmation ? 'bg-blue-600' : 'bg-gray-200'"
                      class="relative inline-flex h-6 w-11 items-center rounded-full transition-colors focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
                    >
                      <span
                        :class="emailSettings.purchaseConfirmation ? 'translate-x-6' : 'translate-x-1'"
                        class="inline-block h-4 w-4 transform rounded-full bg-white transition-transform"
                      />
                    </Switch>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <div>
                      <label class="text-sm font-medium text-gray-700">ارسال خودکار کارت</label>
                      <p class="text-xs text-gray-500">ارسال کارت به گیرنده در زمان مقرر</p>
                    </div>
                    <Switch
                      v-model="emailSettings.autoDelivery"
                      :class="emailSettings.autoDelivery ? 'bg-blue-600' : 'bg-gray-200'"
                      class="relative inline-flex h-6 w-11 items-center rounded-full transition-colors focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
                    >
                      <span
                        :class="emailSettings.autoDelivery ? 'translate-x-6' : 'translate-x-1'"
                        class="inline-block h-4 w-4 transform rounded-full bg-white transition-transform"
                      />
                    </Switch>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <div>
                      <label class="text-sm font-medium text-gray-700">یادآوری انقضا</label>
                      <p class="text-xs text-gray-500">ارسال یادآوری قبل از انقضای کارت</p>
                    </div>
                    <Switch
                      v-model="emailSettings.expiryReminder"
                      :class="emailSettings.expiryReminder ? 'bg-blue-600' : 'bg-gray-200'"
                      class="relative inline-flex h-6 w-11 items-center rounded-full transition-colors focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
                    >
                      <span
                        :class="emailSettings.expiryReminder ? 'translate-x-6' : 'translate-x-1'"
                        class="inline-block h-4 w-4 transform rounded-full bg-white transition-transform"
                      />
                    </Switch>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <div>
                      <label class="text-sm font-medium text-gray-700">اطلاع‌رسانی استفاده</label>
                      <p class="text-xs text-gray-500">ارسال اطلاع‌رسانی هنگام استفاده از کارت</p>
                    </div>
                    <Switch
                      v-model="emailSettings.usageNotification"
                      :class="emailSettings.usageNotification ? 'bg-blue-600' : 'bg-gray-200'"
                      class="relative inline-flex h-6 w-11 items-center rounded-full transition-colors focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
                    >
                      <span
                        :class="emailSettings.usageNotification ? 'translate-x-6' : 'translate-x-1'"
                        class="inline-block h-4 w-4 transform rounded-full bg-white transition-transform"
                      />
                    </Switch>
                  </div>
                </div>
              </div>
            </div>

            <!-- تاریخچه ایمیل‌ها -->
            <div class="bg-gray-50 rounded-lg p-6">
              <h3 class="text-lg font-medium text-gray-900 mb-4">تاریخچه ارسال ایمیل‌ها</h3>
              <div class="space-y-3">
                <div v-for="email in recentEmails" :key="email.id" class="bg-white rounded-lg p-3 border border-gray-200">
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="text-sm font-medium text-gray-900">{{ email.recipient }}</p>
                      <p class="text-xs text-gray-600">{{ email.subject }}</p>
                      <p class="text-xs text-gray-500">{{ formatDate(email.sentAt) }}</p>
                    </div>
                    <span :class="getEmailStatusClass(email.status)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                      {{ getEmailStatusLabel(email.status) }}
                    </span>
                  </div>
                </div>
              </div>
              <button class="w-full mt-4 px-4 py-2 text-sm text-blue-600 hover:text-blue-800">
                مشاهده همه ایمیل‌ها
              </button>
            </div>
          </div>
        </div>

        <!-- تب مدیریت پیامک‌ها -->
        <div v-if="activeTab === 'sms'" class="space-y-6">
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <!-- تنظیمات پیامک -->
            <div class="space-y-6">
              <div class="bg-gray-50 rounded-lg p-6">
                <h3 class="text-lg font-medium text-gray-900 mb-4">تنظیمات ارسال پیامک</h3>
                <div class="space-y-4">
                  <div class="flex items-center justify-between">
                    <div>
                      <label class="text-sm font-medium text-gray-700">فعال کردن ارسال پیامک</label>
                      <p class="text-xs text-gray-500">ارسال خودکار پیامک‌های اطلاع‌رسانی</p>
                    </div>
                    <Switch
                      v-model="smsSettings.enabled"
                      :class="smsSettings.enabled ? 'bg-blue-600' : 'bg-gray-200'"
                      class="relative inline-flex h-6 w-11 items-center rounded-full transition-colors focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
                    >
                      <span
                        :class="smsSettings.enabled ? 'translate-x-6' : 'translate-x-1'"
                        class="inline-block h-4 w-4 transform rounded-full bg-white transition-transform"
                      />
                    </Switch>
                  </div>
                  
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">شماره فرستنده</label>
                    <input
                      v-model="smsSettings.senderNumber"
                      type="text"
                      class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                      placeholder="100000000"
                    />
                  </div>
                  
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">API Key</label>
                    <input
                      v-model="smsSettings.apiKey"
                      type="password"
                      class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                      placeholder="کلید API سرویس پیامک"
                    />
                  </div>
                </div>
              </div>

              <!-- انواع پیامک‌ها -->
              <div class="bg-gray-50 rounded-lg p-6">
                <h3 class="text-lg font-medium text-gray-900 mb-4">انواع پیامک‌های ارسالی</h3>
                <div class="space-y-4">
                  <div class="flex items-center justify-between">
                    <div>
                      <label class="text-sm font-medium text-gray-700">تأیید خرید</label>
                      <p class="text-xs text-gray-500">ارسال پیامک تأیید پس از خرید کارت</p>
                    </div>
                    <Switch
                      v-model="smsSettings.purchaseConfirmation"
                      :class="smsSettings.purchaseConfirmation ? 'bg-blue-600' : 'bg-gray-200'"
                      class="relative inline-flex h-6 w-11 items-center rounded-full transition-colors focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
                    >
                      <span
                        :class="smsSettings.purchaseConfirmation ? 'translate-x-6' : 'translate-x-1'"
                        class="inline-block h-4 w-4 transform rounded-full bg-white transition-transform"
                      />
                    </Switch>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <div>
                      <label class="text-sm font-medium text-gray-700">یادآوری انقضا</label>
                      <p class="text-xs text-gray-500">ارسال یادآوری قبل از انقضای کارت</p>
                    </div>
                    <Switch
                      v-model="smsSettings.expiryReminder"
                      :class="smsSettings.expiryReminder ? 'bg-blue-600' : 'bg-gray-200'"
                      class="relative inline-flex h-6 w-11 items-center rounded-full transition-colors focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
                    >
                      <span
                        :class="smsSettings.expiryReminder ? 'translate-x-6' : 'translate-x-1'"
                        class="inline-block h-4 w-4 transform rounded-full bg-white transition-transform"
                      />
                    </Switch>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <div>
                      <label class="text-sm font-medium text-gray-700">اطلاع‌رسانی استفاده</label>
                      <p class="text-xs text-gray-500">ارسال اطلاع‌رسانی هنگام استفاده از کارت</p>
                    </div>
                    <Switch
                      v-model="smsSettings.usageNotification"
                      :class="smsSettings.usageNotification ? 'bg-blue-600' : 'bg-gray-200'"
                      class="relative inline-flex h-6 w-11 items-center rounded-full transition-colors focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
                    >
                      <span
                        :class="smsSettings.usageNotification ? 'translate-x-6' : 'translate-x-1'"
                        class="inline-block h-4 w-4 transform rounded-full bg-white transition-transform"
                      />
                    </Switch>
                  </div>
                </div>
              </div>
            </div>

            <!-- تاریخچه پیامک‌ها -->
            <div class="bg-gray-50 rounded-lg p-6">
              <h3 class="text-lg font-medium text-gray-900 mb-4">تاریخچه ارسال پیامک‌ها</h3>
              <div class="space-y-3">
                <div v-for="sms in recentSms" :key="sms.id" class="bg-white rounded-lg p-3 border border-gray-200">
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="text-sm font-medium text-gray-900">{{ sms.recipient }}</p>
                      <p class="text-xs text-gray-600">{{ sms.message.substring(0, 50) }}...</p>
                      <p class="text-xs text-gray-500">{{ formatDate(sms.sentAt) }}</p>
                    </div>
                    <span :class="getSmsStatusClass(sms.status)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                      {{ getSmsStatusLabel(sms.status) }}
                    </span>
                  </div>
                </div>
              </div>
              <button class="w-full mt-4 px-4 py-2 text-sm text-blue-600 hover:text-blue-800">
                مشاهده همه پیامک‌ها
              </button>
            </div>
          </div>
        </div>

        <!-- تب قالب‌های پیام -->
        <div v-if="activeTab === 'templates'" class="space-y-6">
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <!-- قالب‌های ایمیل -->
            <div class="space-y-4">
              <h3 class="text-lg font-medium text-gray-900">قالب‌های ایمیل</h3>
              <div class="space-y-3">
                <div v-for="template in emailTemplates" :key="template.id" class="bg-gray-50 rounded-lg p-6">
                  <div class="flex items-center justify-between mb-2">
                    <h4 class="text-sm font-medium text-gray-900">{{ template.name }}</h4>
                    <button 
                      @click="editTemplate(template)"
                      class="text-blue-600 hover:text-blue-800 text-sm"
                    >
                      ویرایش
                    </button>
                  </div>
                  <p class="text-xs text-gray-600 mb-2">{{ template.description }}</p>
                  <div class="text-xs text-gray-500">
                    <span class="font-medium">متغیرها:</span> {{ template.variables.join(', ') }}
                  </div>
                </div>
              </div>
            </div>

            <!-- قالب‌های پیامک -->
            <div class="space-y-4">
              <h3 class="text-lg font-medium text-gray-900">قالب‌های پیامک</h3>
              <div class="space-y-3">
                <div v-for="template in smsTemplates" :key="template.id" class="bg-gray-50 rounded-lg p-6">
                  <div class="flex items-center justify-between mb-2">
                    <h4 class="text-sm font-medium text-gray-900">{{ template.name }}</h4>
                    <button 
                      @click="editTemplate(template)"
                      class="text-blue-600 hover:text-blue-800 text-sm"
                    >
                      ویرایش
                    </button>
                  </div>
                  <p class="text-xs text-gray-600 mb-2">{{ template.description }}</p>
                  <div class="text-xs text-gray-500">
                    <span class="font-medium">متغیرها:</span> {{ template.variables.join(', ') }}
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- تب تنظیمات ارسال -->
        <div v-if="activeTab === 'settings'" class="space-y-6">
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <!-- تنظیمات عمومی -->
            <div class="bg-gray-50 rounded-lg p-6">
              <h3 class="text-lg font-medium text-gray-900 mb-4">تنظیمات عمومی</h3>
              <div class="space-y-4">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-1">تأخیر ارسال (دقیقه)</label>
                  <input
                    v-model.number="generalSettings.delayMinutes"
                    type="number"
                    min="0"
                    max="60"
                    class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                  />
                </div>
                
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-1">حداکثر تلاش ارسال</label>
                  <input
                    v-model.number="generalSettings.maxRetries"
                    type="number"
                    min="1"
                    max="5"
                    class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                  />
                </div>
                
                <div class="flex items-center justify-between">
                  <div>
                    <label class="text-sm font-medium text-gray-700">لاگ کامل ارسال</label>
                    <p class="text-xs text-gray-500">ثبت تمام جزئیات ارسال</p>
                  </div>
                  <Switch
                    v-model="generalSettings.fullLogging"
                    :class="generalSettings.fullLogging ? 'bg-blue-600' : 'bg-gray-200'"
                    class="relative inline-flex h-6 w-11 items-center rounded-full transition-colors focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
                  >
                    <span
                      :class="generalSettings.fullLogging ? 'translate-x-6' : 'translate-x-1'"
                      class="inline-block h-4 w-4 transform rounded-full bg-white transition-transform"
                    />
                  </Switch>
                </div>
              </div>
            </div>

            <!-- تنظیمات زمان‌بندی -->
            <div class="bg-gray-50 rounded-lg p-6">
              <h3 class="text-lg font-medium text-gray-900 mb-4">تنظیمات زمان‌بندی</h3>
              <div class="space-y-4">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-1">روزهای قبل از انقضا</label>
                  <input
                    v-model.number="generalSettings.expiryReminderDays"
                    type="number"
                    min="1"
                    max="30"
                    class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                  />
                </div>
                
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-1">ساعت ارسال یادآوری</label>
                  <input
                    v-model="generalSettings.reminderHour"
                    type="time"
                    class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                  />
                </div>
                
                <div class="flex items-center justify-between">
                  <div>
                    <label class="text-sm font-medium text-gray-700">ارسال در تعطیلات</label>
                    <p class="text-xs text-gray-500">ارسال اطلاع‌رسانی در روزهای تعطیل</p>
                  </div>
                  <Switch
                    v-model="generalSettings.sendOnHolidays"
                    :class="generalSettings.sendOnHolidays ? 'bg-blue-600' : 'bg-gray-200'"
                    class="relative inline-flex h-6 w-11 items-center rounded-full transition-colors focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
                  >
                    <span
                      :class="generalSettings.sendOnHolidays ? 'translate-x-6' : 'translate-x-1'"
                      class="inline-block h-4 w-4 transform rounded-full bg-white transition-transform"
                    />
                  </Switch>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- مودال ویرایش قالب -->
    <GiftCardTemplateEditModal 
      v-if="showTemplateModal"
      :template="selectedTemplate"
      @close="showTemplateModal = false"
      @saved="handleTemplateSaved"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted } from 'vue'
import { Switch } from '@headlessui/vue'

// کامپوننت‌های مورد نیاز
import GiftCardTemplateEditModal from './GiftCardTemplateEditModal.vue'

// Reactive data
const activeTab = ref('email')
const isSaving = ref(false)
const showTemplateModal = ref(false)
const selectedTemplate = ref(null)

// آمار اطلاع‌رسانی
const emailStats = reactive({
  sent: 1250,
  successRate: 98.5
})

const smsStats = reactive({
  sent: 890,
  successRate: 95.2
})

const pendingNotifications = ref(15)
const failedNotifications = ref(8)

// تنظیمات ایمیل
const emailSettings = reactive({
  enabled: true,
  senderEmail: 'noreply@giftcard.com',
  senderName: 'سیستم گیفت کارت',
  purchaseConfirmation: true,
  autoDelivery: true,
  expiryReminder: true,
  usageNotification: false
})

// تنظیمات پیامک
const smsSettings = reactive({
  enabled: true,
  senderNumber: '100000000',
  apiKey: '',
  purchaseConfirmation: true,
  expiryReminder: true,
  usageNotification: false
})

// تنظیمات عمومی
const generalSettings = reactive({
  delayMinutes: 5,
  maxRetries: 3,
  fullLogging: true,
  expiryReminderDays: 7,
  reminderHour: '09:00',
  sendOnHolidays: false
})

// تاریخچه ایمیل‌ها
const recentEmails = ref([
  {
    id: 1,
    recipient: 'ali@example.com',
    subject: 'تأیید خرید گیفت کارت',
    status: 'sent',
    sentAt: new Date(Date.now() - 3600000)
  },
  {
    id: 2,
    recipient: 'fateme@example.com',
    subject: 'یادآوری انقضای کارت',
    status: 'sent',
    sentAt: new Date(Date.now() - 7200000)
  },
  {
    id: 3,
    recipient: 'hasan@example.com',
    subject: 'اطلاع‌رسانی استفاده',
    status: 'failed',
    sentAt: new Date(Date.now() - 10800000)
  }
])

// تاریخچه پیامک‌ها
const recentSms = ref([
  {
    id: 1,
    recipient: '09123456789',
    message: 'گیفت کارت شما با موفقیت خریداری شد. کد: GC-001',
    status: 'sent',
    sentAt: new Date(Date.now() - 3600000)
  },
  {
    id: 2,
    recipient: '09187654321',
    message: 'یادآوری: کارت شما تا 7 روز دیگر منقضی می‌شود',
    status: 'sent',
    sentAt: new Date(Date.now() - 7200000)
  },
  {
    id: 3,
    recipient: '09111111111',
    message: 'کارت شما استفاده شد. مبلغ باقی‌مانده: 250,000 تومان',
    status: 'failed',
    sentAt: new Date(Date.now() - 10800000)
  }
])

// قالب‌های ایمیل
const emailTemplates = ref([
  {
    id: 1,
    name: 'تأیید خرید',
    description: 'ایمیل تأیید پس از خرید موفق گیفت کارت',
    variables: ['نام خریدار', 'مبلغ', 'کد کارت', 'تاریخ انقضا'],
    type: 'email'
  },
  {
    id: 2,
    name: 'ارسال کارت',
    description: 'ایمیل ارسال کارت به گیرنده',
    variables: ['نام گیرنده', 'نام فرستنده', 'پیام شخصی', 'کد کارت'],
    type: 'email'
  },
  {
    id: 3,
    name: 'یادآوری انقضا',
    description: 'یادآوری قبل از انقضای کارت',
    variables: ['نام گیرنده', 'کد کارت', 'تاریخ انقضا', 'مبلغ باقی‌مانده'],
    type: 'email'
  }
])

// قالب‌های پیامک
const smsTemplates = ref([
  {
    id: 1,
    name: 'تأیید خرید',
    description: 'پیامک تأیید پس از خرید موفق',
    variables: ['مبلغ', 'کد کارت'],
    type: 'sms'
  },
  {
    id: 2,
    name: 'یادآوری انقضا',
    description: 'یادآوری قبل از انقضای کارت',
    variables: ['کد کارت', 'تاریخ انقضا'],
    type: 'sms'
  },
  {
    id: 3,
    name: 'اطلاع‌رسانی استفاده',
    description: 'اطلاع‌رسانی هنگام استفاده از کارت',
    variables: ['مبلغ استفاده شده', 'مبلغ باقی‌مانده'],
    type: 'sms'
  }
])

// Methods
const saveSettings = async () => {
  isSaving.value = true
  try {
    // شبیه‌سازی API call
    await new Promise(resolve => setTimeout(resolve, 1000))
    
    console.log('تنظیمات اطلاع‌رسانی ذخیره شد:', {
      emailSettings,
      smsSettings,
      generalSettings
    })
    
    alert('تنظیمات با موفقیت ذخیره شد')
  } catch (error) {
    console.error('خطا در ذخیره تنظیمات:', error)
    alert('خطا در ذخیره تنظیمات')
  } finally {
    isSaving.value = false
  }
}

const testNotification = () => {
  if (confirm('آیا می‌خواهید یک اطلاع‌رسانی تست ارسال شود؟')) {
    console.log('ارسال اطلاع‌رسانی تست...')
    alert('اطلاع‌رسانی تست ارسال شد')
  }
}

const editTemplate = (template: any) => {
  selectedTemplate.value = template
  showTemplateModal.value = true
}

const handleTemplateSaved = (template: any) => {
  // به‌روزرسانی قالب در لیست
  if (template.type === 'email') {
    const index = emailTemplates.value.findIndex(t => t.id === template.id)
    if (index > -1) {
      emailTemplates.value[index] = template
    }
  } else {
    const index = smsTemplates.value.findIndex(t => t.id === template.id)
    if (index > -1) {
      smsTemplates.value[index] = template
    }
  }
  
  showTemplateModal.value = false
  alert('قالب با موفقیت ذخیره شد')
}

const formatDate = (date: Date) => {
  return new Intl.DateTimeFormat('fa-IR').format(date)
}

const getEmailStatusClass = (status: string) => {
  const classes = {
    sent: 'bg-green-100 text-green-800',
    failed: 'bg-red-100 text-red-800',
    pending: 'bg-yellow-100 text-yellow-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getEmailStatusLabel = (status: string) => {
  const labels = {
    sent: 'ارسال شد',
    failed: 'ناموفق',
    pending: 'در انتظار'
  }
  return labels[status] || status
}

const getSmsStatusClass = (status: string) => {
  const classes = {
    sent: 'bg-green-100 text-green-800',
    failed: 'bg-red-100 text-red-800',
    pending: 'bg-yellow-100 text-yellow-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getSmsStatusLabel = (status: string) => {
  const labels = {
    sent: 'ارسال شد',
    failed: 'ناموفق',
    pending: 'در انتظار'
  }
  return labels[status] || status
}

// Lifecycle
onMounted(() => {
  console.log('Gift card notification manager component mounted')
})
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 

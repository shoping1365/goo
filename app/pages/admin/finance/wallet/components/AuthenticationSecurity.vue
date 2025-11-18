<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between">
        <div>
          <h2 class="text-xl font-bold text-gray-900">احراز هویت و امنیت تراکنش‌ها</h2>
          <p class="text-gray-600 mt-1">مدیریت امنیت، احراز هویت و کنترل دسترسی‌ها</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button
            @click="activeTab = 'authentication'"
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'authentication'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
          >
            احراز هویت
          </button>
          <button
            @click="activeTab = 'security'"
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'security'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
          >
            امنیت
          </button>
          <button
            @click="activeTab = 'permissions'"
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'permissions'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
          >
            دسترسی‌ها
          </button>
          <button
            @click="activeTab = 'monitoring'"
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'monitoring'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
          >
            نظارت
          </button>
        </div>
      </div>
    </div>

    <!-- محتوای تب‌ها -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200">
      <!-- احراز هویت -->
      <div v-if="activeTab === 'authentication'" class="p-6">
        <div class="space-y-6">
          <!-- تنظیمات احراز هویت -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات احراز هویت</h3>
            <div class="space-y-4">
              <div class="flex items-center justify-between">
                <div>
                  <h4 class="font-medium text-gray-900">احراز هویت دو مرحله‌ای (2FA)</h4>
                  <p class="text-gray-600 text-sm">فعال‌سازی احراز هویت دو مرحله‌ای برای امنیت بیشتر</p>
                </div>
                <label class="relative inline-flex items-center cursor-pointer">
                  <input type="checkbox" v-model="authSettings.twoFactor" class="sr-only peer">
                  <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
                </label>
              </div>
              <div class="flex items-center justify-between">
                <div>
                  <h4 class="font-medium text-gray-900">رمز عبور قوی</h4>
                  <p class="text-gray-600 text-sm">اجبار استفاده از رمز عبور پیچیده</p>
                </div>
                <label class="relative inline-flex items-center cursor-pointer">
                  <input type="checkbox" v-model="authSettings.strongPassword" class="sr-only peer">
                  <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
                </label>
              </div>
              <div class="flex items-center justify-between">
                <div>
                  <h4 class="font-medium text-gray-900">تغییر خودکار رمز عبور</h4>
                  <p class="text-gray-600 text-sm">اجبار تغییر رمز عبور هر 90 روز</p>
                </div>
                <label class="relative inline-flex items-center cursor-pointer">
                  <input type="checkbox" v-model="authSettings.autoPasswordChange" class="sr-only peer">
                  <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
                </label>
              </div>
            </div>
          </div>

          <!-- روش‌های احراز هویت -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">روش‌های احراز هویت</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div class="border border-gray-200 rounded-lg p-6">
                <div class="flex items-center space-x-3 space-x-reverse mb-3">
                  <div class="w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center">
                    <svg class="w-4 h-4 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
                    </svg>
                  </div>
                  <h4 class="font-medium text-gray-900">رمز عبور</h4>
                </div>
                <p class="text-gray-600 text-sm mb-3">احراز هویت با رمز عبور و نام کاربری</p>
                <label class="relative inline-flex items-center cursor-pointer">
                  <input type="checkbox" v-model="authMethods.password" class="sr-only peer">
                  <div class="w-9 h-5 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-4 after:w-4 after:transition-all peer-checked:bg-blue-600"></div>
                </label>
              </div>
              <div class="border border-gray-200 rounded-lg p-6">
                <div class="flex items-center space-x-3 space-x-reverse mb-3">
                  <div class="w-8 h-8 bg-green-100 rounded-full flex items-center justify-center">
                    <svg class="w-4 h-4 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                  </div>
                  <h4 class="font-medium text-gray-900">کد تایید</h4>
                </div>
                <p class="text-gray-600 text-sm mb-3">ارسال کد تایید به ایمیل یا پیامک</p>
                <label class="relative inline-flex items-center cursor-pointer">
                  <input type="checkbox" v-model="authMethods.otp" class="sr-only peer">
                  <div class="w-9 h-5 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-4 after:w-4 after:transition-all peer-checked:bg-blue-600"></div>
                </label>
              </div>
              <div class="border border-gray-200 rounded-lg p-6">
                <div class="flex items-center space-x-3 space-x-reverse mb-3">
                  <div class="w-8 h-8 bg-purple-100 rounded-full flex items-center justify-center">
                    <svg class="w-4 h-4 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 7a2 2 0 012 2m4 0a6 6 0 01-7.743 5.743L11 17H9v2H7v2H4a1 1 0 01-1-1v-2.586a1 1 0 01.293-.707l5.964-5.964A6 6 0 1121 9z" />
                    </svg>
                  </div>
                  <h4 class="font-medium text-gray-900">کلید امنیتی</h4>
                </div>
                <p class="text-gray-600 text-sm mb-3">استفاده از کلیدهای امنیتی فیزیکی</p>
                <label class="relative inline-flex items-center cursor-pointer">
                  <input type="checkbox" v-model="authMethods.securityKey" class="sr-only peer">
                  <div class="w-9 h-5 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-4 after:w-4 after:transition-all peer-checked:bg-blue-600"></div>
                </label>
              </div>
              <div class="border border-gray-200 rounded-lg p-6">
                <div class="flex items-center space-x-3 space-x-reverse mb-3">
                  <div class="w-8 h-8 bg-orange-100 rounded-full flex items-center justify-center">
                    <svg class="w-4 h-4 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.75 17L9 20l-1 1h8l-1-1-.75-3M3 13h18M5 17h14a2 2 0 002-2V5a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
                    </svg>
                  </div>
                  <h4 class="font-medium text-gray-900">احراز هویت زیستی</h4>
                </div>
                <p class="text-gray-600 text-sm mb-3">استفاده از اثر انگشت یا تشخیص چهره</p>
                <label class="relative inline-flex items-center cursor-pointer">
                  <input type="checkbox" v-model="authMethods.biometric" class="sr-only peer">
                  <div class="w-9 h-5 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-4 after:w-4 after:transition-all peer-checked:bg-blue-600"></div>
                </label>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- امنیت -->
      <div v-if="activeTab === 'security'" class="p-6">
        <div class="space-y-6">
          <!-- تنظیمات امنیتی -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات امنیتی</h3>
            <div class="space-y-4">
              <div class="flex items-center justify-between">
                <div>
                  <h4 class="font-medium text-gray-900">محدودیت تلاش ورود</h4>
                  <p class="text-gray-600 text-sm">قفل شدن حساب پس از تلاش‌های ناموفق</p>
                </div>
                <select v-model="securitySettings.maxLoginAttempts" class="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                  <option value="3">3 تلاش</option>
                  <option value="5">5 تلاش</option>
                  <option value="10">10 تلاش</option>
                </select>
              </div>
              <div class="flex items-center justify-between">
                <div>
                  <h4 class="font-medium text-gray-900">مدت زمان قفل</h4>
                  <p class="text-gray-600 text-sm">مدت زمان قفل شدن حساب (دقیقه)</p>
                </div>
                <select v-model="securitySettings.lockoutDuration" class="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                  <option value="15">15 دقیقه</option>
                  <option value="30">30 دقیقه</option>
                  <option value="60">1 ساعت</option>
                  <option value="1440">24 ساعت</option>
                </select>
              </div>
              <div class="flex items-center justify-between">
                <div>
                  <h4 class="font-medium text-gray-900">اعلان ورود مشکوک</h4>
                  <p class="text-gray-600 text-sm">ارسال اعلان برای ورود از مکان‌های جدید</p>
                </div>
                <label class="relative inline-flex items-center cursor-pointer">
                  <input type="checkbox" v-model="securitySettings.suspiciousLoginAlert" class="sr-only peer">
                  <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
                </label>
              </div>
            </div>
          </div>

          <!-- رمزنگاری -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات رمزنگاری</h3>
            <div class="space-y-4">
              <div class="flex items-center justify-between">
                <div>
                  <h4 class="font-medium text-gray-900">الگوریتم رمزنگاری</h4>
                  <p class="text-gray-600 text-sm">انتخاب الگوریتم رمزنگاری داده‌ها</p>
                </div>
                <select v-model="securitySettings.encryptionAlgorithm" class="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                  <option value="AES-256">AES-256</option>
                  <option value="AES-192">AES-192</option>
                  <option value="AES-128">AES-128</option>
                </select>
              </div>
              <div class="flex items-center justify-between">
                <div>
                  <h4 class="font-medium text-gray-900">کلیدهای رمزنگاری</h4>
                  <p class="text-gray-600 text-sm">مدیریت کلیدهای رمزنگاری</p>
                </div>
                <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors">
                  مدیریت کلیدها
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- دسترسی‌ها -->
      <div v-if="activeTab === 'permissions'" class="p-6">
        <div class="space-y-6">
          <!-- نقش‌های کاربری -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">نقش‌های کاربری</h3>
            <div class="space-y-4">
              <div v-for="role in userRoles" :key="role.id" class="border border-gray-200 rounded-lg p-6">
                <div class="flex items-center justify-between mb-3">
                  <h4 class="font-medium text-gray-900">{{ role.name }}</h4>
                  <button class="text-blue-600 hover:text-blue-700 text-sm">ویرایش</button>
                </div>
                <p class="text-gray-600 text-sm mb-3">{{ role.description }}</p>
                <div class="flex flex-wrap gap-2">
                  <span v-for="permission in role.permissions" :key="permission" class="px-2 py-1 bg-blue-100 text-blue-800 text-xs rounded">
                    {{ permission }}
                  </span>
                </div>
              </div>
            </div>
          </div>

          <!-- مجوزهای دسترسی -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">مجوزهای دسترسی</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div v-for="permission in permissions" :key="permission.id" class="border border-gray-200 rounded-lg p-6">
                <div class="flex items-center justify-between mb-2">
                  <h4 class="font-medium text-gray-900">{{ permission.name }}</h4>
                  <label class="relative inline-flex items-center cursor-pointer">
                    <input type="checkbox" v-model="permission.enabled" class="sr-only peer">
                    <div class="w-9 h-5 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-4 after:w-4 after:transition-all peer-checked:bg-blue-600"></div>
                  </label>
                </div>
                <p class="text-gray-600 text-sm">{{ permission.description }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- نظارت -->
      <div v-if="activeTab === 'monitoring'" class="p-6">
        <div class="space-y-6">
          <!-- فعالیت‌های اخیر -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">فعالیت‌های اخیر</h3>
            <div class="space-y-3">
              <div v-for="activity in recentActivities" :key="activity.id" class="flex items-center space-x-3 space-x-reverse p-3 bg-gray-50 rounded-lg">
                <div :class="`w-8 h-8 rounded-full flex items-center justify-center ${activity.type === 'login' ? 'bg-green-100' : activity.type === 'transaction' ? 'bg-blue-100' : 'bg-red-100'}`">
                  <svg class="w-4 h-4" :class="activity.type === 'login' ? 'text-green-600' : activity.type === 'transaction' ? 'text-blue-600' : 'text-red-600'" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path v-if="activity.type === 'login'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 16l-4-4m0 0l4-4m-4 4h14m-5 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h7a3 3 0 013 3v1" />
                    <path v-else-if="activity.type === 'transaction'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1" />
                    <path v-else stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z" />
                  </svg>
                </div>
                <div class="flex-1">
                  <p class="text-sm font-medium text-gray-900">{{ activity.description }}</p>
                  <p class="text-xs text-gray-500">{{ activity.time }}</p>
                </div>
                <div class="text-xs text-gray-500">{{ activity.ip }}</div>
              </div>
            </div>
          </div>

          <!-- آمار امنیتی -->
          <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
            <div class="bg-white border border-gray-200 rounded-lg p-6">
              <div class="flex items-center">
                <div class="w-8 h-8 bg-green-100 rounded-full flex items-center justify-center">
                  <svg class="w-4 h-4 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                </div>
                <div class="mr-3">
                  <p class="text-sm font-medium text-gray-900">ورودهای موفق</p>
                  <p class="text-lg font-bold text-green-600">{{ securityStats.successfulLogins }}</p>
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
                  <p class="text-sm font-medium text-gray-900">تلاش‌های ناموفق</p>
                  <p class="text-lg font-bold text-red-600">{{ securityStats.failedLogins }}</p>
                </div>
              </div>
            </div>
            <div class="bg-white border border-gray-200 rounded-lg p-6">
              <div class="flex items-center">
                <div class="w-8 h-8 bg-yellow-100 rounded-full flex items-center justify-center">
                  <svg class="w-4 h-4 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                </div>
                <div class="mr-3">
                  <p class="text-sm font-medium text-gray-900">هشدارهای امنیتی</p>
                  <p class="text-lg font-bold text-yellow-600">{{ securityStats.securityAlerts }}</p>
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
import { ref } from 'vue'
// تعریف reactive data
const activeTab = ref('authentication')

// تنظیمات احراز هویت
const authSettings = ref({
  twoFactor: true,
  strongPassword: true,
  autoPasswordChange: true
})

// روش‌های احراز هویت
const authMethods = ref({
  password: true,
  otp: true,
  securityKey: false,
  biometric: false
})

// تنظیمات امنیتی
const securitySettings = ref({
  maxLoginAttempts: 5,
  lockoutDuration: 30,
  suspiciousLoginAlert: true,
  encryptionAlgorithm: 'AES-256'
})

// نقش‌های کاربری
const userRoles = ref([
  {
    id: 1,
    name: 'مدیر سیستم',
    description: 'دسترسی کامل به تمام بخش‌های کیف پول',
    permissions: ['مدیریت کاربران', 'تنظیمات امنیتی', 'گزارش‌گیری', 'مدیریت تراکنش‌ها']
  },
  {
    id: 2,
    name: 'کارشناس مالی',
    description: 'دسترسی به بخش‌های مالی و گزارش‌گیری',
    permissions: ['مشاهده تراکنش‌ها', 'گزارش‌گیری', 'مدیریت موجودی']
  },
  {
    id: 3,
    name: 'اپراتور',
    description: 'دسترسی محدود به عملیات روزانه',
    permissions: ['مشاهده تراکنش‌ها', 'تایید تراکنش‌های کوچک']
  }
])

// مجوزهای دسترسی
const permissions = ref([
  { id: 1, name: 'مشاهده موجودی', description: 'مشاهده موجودی کیف پول‌ها', enabled: true },
  { id: 2, name: 'مدیریت تراکنش‌ها', description: 'ایجاد و ویرایش تراکنش‌ها', enabled: true },
  { id: 3, name: 'گزارش‌گیری', description: 'دسترسی به گزارش‌های مالی', enabled: true },
  { id: 4, name: 'تنظیمات امنیتی', description: 'تغییر تنظیمات امنیتی', enabled: false },
  { id: 5, name: 'مدیریت کاربران', description: 'افزودن و حذف کاربران', enabled: false },
  { id: 6, name: 'پشتیبان‌گیری', description: 'ایجاد و بازیابی پشتیبان', enabled: false }
])

// فعالیت‌های اخیر
const recentActivities = ref([
  {
    id: 1,
    type: 'login',
    description: 'ورود موفق کاربر admin',
    time: '2 دقیقه پیش',
    ip: '192.168.1.100'
  },
  {
    id: 2,
    type: 'transaction',
    description: 'تراکنش جدید: شارژ کیف پول',
    time: '5 دقیقه پیش',
    ip: '192.168.1.100'
  },
  {
    id: 3,
    type: 'security',
    description: 'تلاش ورود ناموفق از IP جدید',
    time: '10 دقیقه پیش',
    ip: '203.0.113.45'
  }
])

// آمار امنیتی
const securityStats = ref({
  successfulLogins: 1247,
  failedLogins: 23,
  securityAlerts: 5
})
</script> 

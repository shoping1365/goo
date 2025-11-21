<template>
  <div class="p-6">
    <!-- عنوان صفحه -->
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">مدیریت احراز هویت دو مرحله‌ای</h1>
      <p class="text-gray-600 mt-2">مدیریت و نظارت بر احراز هویت دو مرحله‌ای کاربران</p>
    </div>

    <!-- کارت‌های آمار -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
      <!-- کل کاربران -->
      <div class="bg-gradient-to-r from-blue-500 to-blue-600 rounded-lg p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-blue-100 text-sm">کل کاربران</p>
            <p class="text-3xl font-bold">1,234</p>
          </div>
          <div class="text-blue-200">
            <svg class="w-8 h-8" fill="currentColor" viewBox="0 0 20 20">
              <path d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
        </div>
      </div>

      <!-- کاربران فعال -->
      <div class="bg-gradient-to-r from-green-500 to-green-600 rounded-lg p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-green-100 text-sm">فعال</p>
            <p class="text-3xl font-bold">856</p>
          </div>
          <div class="text-green-200">
            <svg class="w-8 h-8" fill="currentColor" viewBox="0 0 20 20">
              <path d="M10 12a2 2 0 100-4 2 2 0 000 4z" />
              <path fill-rule="evenodd" d="M.458 10C1.732 5.943 5.522 3 10 3s8.268 2.943 9.542 7c-1.274 4.057-5.064 7-9.542 7S1.732 14.057.458 10zM14 10a4 4 0 11-8 0 4 4 0 018 0z" clip-rule="evenodd" />
            </svg>
          </div>
        </div>
      </div>

      <!-- کاربران غیرفعال -->
      <div class="bg-gradient-to-r from-red-500 to-red-600 rounded-lg p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-red-100 text-sm">غیرفعال</p>
            <p class="text-3xl font-bold">378</p>
          </div>
          <div class="text-red-200">
            <svg class="w-8 h-8" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-8-3a1 1 0 00-.867.5 1 1 0 11-1.731-1A3 3 0 0113 8a3.001 3.001 0 01-2 2.83V11a1 1 0 11-2 0v-1a1 1 0 011-1 1 1 0 100-2zm0 8a1 1 0 100-2 1 1 0 000 2z" clip-rule="evenodd" />
            </svg>
          </div>
        </div>
      </div>

      <!-- درصد فعال -->
      <div class="bg-gradient-to-r from-purple-500 to-purple-600 rounded-lg p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-purple-100 text-sm">درصد فعال</p>
            <p class="text-3xl font-bold">69%</p>
          </div>
          <div class="text-purple-200">
            <svg class="w-8 h-8" fill="currentColor" viewBox="0 0 20 20">
              <path d="M2 10a8 8 0 018-8v8h8a8 8 0 11-16 0z" />
              <path d="M12 2.252A8.014 8.014 0 0117.748 8H12V2.252z" />
            </svg>
          </div>
        </div>
      </div>
    </div>



    <!-- بخش ایجاد تایید دو مرحله‌ای -->
    <div v-if="showCreateForm" class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-6">
      <div class="mb-4">
        <h3 class="text-lg font-semibold text-gray-900">ایجاد تایید دو مرحله‌ای</h3>
      </div>

      <!-- فرم ایجاد -->
      <div class="border-t border-gray-200 pt-4">
        <form class="space-y-4" @submit.prevent="createTwoFactor">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <!-- نام کاربر -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نام کاربر</label>
              <input
                v-model="newTwoFactor.name"
                type="text"
                required
                placeholder="نام کامل کاربر"
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>

            <!-- ایمیل -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">ایمیل</label>
              <input
                v-model="newTwoFactor.email"
                type="email"
                required
                placeholder="example@domain.com"
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>

            <!-- شماره تلفن -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">شماره تلفن</label>
              <input
                v-model="newTwoFactor.phone"
                type="tel"
                required
                placeholder="09123456789"
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>

            <!-- نوع احراز هویت -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نوع احراز هویت</label>
              <select
                v-model="newTwoFactor.type"
                required
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="">انتخاب کنید</option>
                <option value="sms">پیامک (SMS)</option>
                <option value="email">ایمیل</option>
                <option value="app">اپلیکیشن (Google Authenticator)</option>
                <option value="backup">کدهای پشتیبان</option>
              </select>
            </div>

            <!-- مدت اعتبار -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">مدت اعتبار (دقیقه)</label>
              <input
                v-model="newTwoFactor.expiry"
                type="number"
                min="1"
                max="60"
                required
                placeholder="5"
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>

            <!-- تعداد تلاش مجاز -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">تعداد تلاش مجاز</label>
              <input
                v-model="newTwoFactor.maxAttempts"
                type="number"
                min="1"
                max="10"
                required
                placeholder="3"
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
          </div>

          <!-- توضیحات -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">توضیحات</label>
            <textarea
              v-model="newTwoFactor.description"
              rows="3"
              placeholder="توضیحات اختیاری..."
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            ></textarea>
          </div>

          <!-- تنظیمات پیشرفته -->
          <div class="border-t border-gray-200 pt-4">
            <h4 class="text-md font-medium text-gray-900 mb-3">تنظیمات پیشرفته</h4>
            <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
              <div class="flex items-center">
                <input
                  id="forceEnable"
                  v-model="newTwoFactor.forceEnable"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label for="forceEnable" class="mr-2 block text-sm text-gray-900">
                  اجبار فعال‌سازی
                </label>
              </div>
              <div class="flex items-center">
                <input
                  id="requireOnLogin"
                  v-model="newTwoFactor.requireOnLogin"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label for="requireOnLogin" class="mr-2 block text-sm text-gray-900">
                  الزامی در ورود
                </label>
              </div>
              <div class="flex items-center">
                <input
                  id="sendNotification"
                  v-model="newTwoFactor.sendNotification"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label for="sendNotification" class="mr-2 block text-sm text-gray-900">
                  ارسال اعلان
                </label>
              </div>
            </div>
          </div>

          <!-- دکمه‌های عملیات -->
          <div class="flex gap-3 pt-4">
            <button
              type="submit"
              class="px-6 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors flex items-center gap-2"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
              </svg>
              ایجاد تایید دو مرحله‌ای
            </button>
            <button
              type="button"
              class="px-6 py-2 bg-gray-500 text-white rounded-lg hover:bg-gray-600 transition-colors"
              @click="resetForm"
            >
              پاک کردن فرم
            </button>
            <button
              type="button"
              class="px-6 py-2 bg-red-500 text-white rounded-lg hover:bg-red-600 transition-colors"
              @click="showCreateForm = false"
            >
              انصراف
            </button>
          </div>
        </form>
      </div>
    </div>

    <!-- بخش فیلترها و عملیات -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-6">
      <div class="flex flex-col lg:flex-row gap-6">
        <!-- جستجو -->
        <div class="flex-1">
          <label class="block text-sm font-medium text-gray-700 mb-2">جستجو</label>
          <input
            v-model="searchQuery"
            type="text"
            placeholder="جستجو بر اساس نام، ایمیل یا شماره تلفن..."
            class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
        </div>

        <!-- وضعیت -->
        <div class="lg:w-48">
          <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
          <select
            v-model="statusFilter"
            class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه</option>
            <option value="active">فعال</option>
            <option value="inactive">غیرفعال</option>
          </select>
        </div>

        <!-- نوع دستگاه -->
        <div class="lg:w-48">
          <label class="block text-sm font-medium text-gray-700 mb-2">نوع دستگاه</label>
          <select
            v-model="deviceTypeFilter"
            class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه</option>
            <option value="mobile">موبایل</option>
            <option value="desktop">دسکتاپ</option>
            <option value="tablet">تبلت</option>
          </select>
        </div>

        <!-- تاریخ -->
        <div class="lg:w-48">
          <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ</label>
          <input
            v-model="dateFilter"
            type="date"
            class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
        </div>

        <!-- دکمه‌های عملیات -->
        <div class="flex gap-2 items-end">
          <button
            class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
            @click="applyFilters"
          >
            اعمال فیلتر
          </button>
          <button
            class="px-6 py-2 bg-gray-500 text-white rounded-lg hover:bg-gray-600 transition-colors"
            @click="clearFilters"
          >
            پاک کردن
          </button>
        </div>
      </div>

      <!-- دکمه‌های عملیات اضافی -->
      <div class="flex items-center justify-between mt-4 pt-4 border-t border-gray-200">
        <div class="flex items-center gap-3">
          <!-- دکمه فعال‌سازی تایید دو مرحله‌ای -->
          <button
            class="px-4 py-2 bg-purple-600 text-white rounded-lg hover:bg-purple-700 transition-colors flex items-center gap-2"
            @click="showSetupForm = !showSetupForm"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
            </svg>
            {{ showSetupForm ? 'بستن فعال‌سازی' : 'فعال‌سازی برای خودم' }}
          </button>

          <!-- دکمه خروجی اکسل -->
          <button
            class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors flex items-center gap-2"
            @click="exportToExcel"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
            </svg>
            خروجی اکسل
          </button>
        </div>

        <!-- آمار سریع -->
        <div class="flex items-center gap-6 text-sm text-gray-600">
          <span>کل: {{ filteredUsers.length }}</span>
          <span>فعال: {{ activeUsersCount }}</span>
          <span>غیرفعال: {{ inactiveUsersCount }}</span>
        </div>
      </div>

      <!-- فرم فعال‌سازی (در کنار دکمه‌ها) -->
      <div v-if="showSetupForm" class="mt-4 p-6 bg-gray-50 rounded-lg border">
        <div class="space-y-4">
          <!-- مرحله 1: انتخاب روش -->
          <div v-if="setupStep === 1">
            <h4 class="text-md font-medium text-gray-900 mb-3">مرحله 1: انتخاب روش احراز هویت</h4>
            <div class="grid grid-cols-2 md:grid-cols-4 gap-3">
              <div
                :class="[
                  'p-3 border-2 rounded-lg cursor-pointer transition-all text-center',
                  selectedMethod === 'sms' ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-gray-300'
                ]"
                @click="selectMethod('sms')"
              >
                <div class="w-8 h-8 bg-blue-100 rounded-lg flex items-center justify-center mx-auto mb-2">
                  <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" />
                  </svg>
                </div>
                <h5 class="font-medium text-gray-900 text-sm">پیامک</h5>
              </div>

              <div
                :class="[
                  'p-3 border-2 rounded-lg cursor-pointer transition-all text-center',
                  selectedMethod === 'email' ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-gray-300'
                ]"
                @click="selectMethod('email')"
              >
                <div class="w-8 h-8 bg-green-100 rounded-lg flex items-center justify-center mx-auto mb-2">
                  <svg class="w-5 h-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 4.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
                  </svg>
                </div>
                <h5 class="font-medium text-gray-900 text-sm">ایمیل</h5>
              </div>

              <div
                :class="[
                  'p-3 border-2 rounded-lg cursor-pointer transition-all text-center',
                  selectedMethod === 'app' ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-gray-300'
                ]"
                @click="selectMethod('app')"
              >
                <div class="w-8 h-8 bg-purple-100 rounded-lg flex items-center justify-center mx-auto mb-2">
                  <svg class="w-5 h-5 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 18h.01M8 21h8a2 2 0 002-2V5a2 2 0 00-2-2H8a2 2 0 00-2 2v14a2 2 0 002 2z" />
                  </svg>
                </div>
                <h5 class="font-medium text-gray-900 text-sm">اپلیکیشن</h5>
              </div>

              <div
                :class="[
                  'p-3 border-2 rounded-lg cursor-pointer transition-all text-center',
                  selectedMethod === 'backup' ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-gray-300'
                ]"
                @click="selectMethod('backup')"
              >
                <div class="w-8 h-8 bg-orange-100 rounded-lg flex items-center justify-center mx-auto mb-2">
                  <svg class="w-5 h-5 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
                  </svg>
                </div>
                <h5 class="font-medium text-gray-900 text-sm">کدهای پشتیبان</h5>
              </div>
            </div>

            <div class="flex justify-end mt-4">
              <button
                :disabled="!selectedMethod"
                :class="[
                  'px-4 py-2 rounded-lg transition-colors text-sm',
                  selectedMethod ? 'bg-blue-600 text-white hover:bg-blue-700' : 'bg-gray-300 text-gray-500 cursor-not-allowed'
                ]"
                @click="nextStep"
              >
                مرحله بعد
              </button>
            </div>
          </div>

          <!-- مرحله 2: تنظیمات -->
          <div v-if="setupStep === 2">
            <h4 class="text-md font-medium text-gray-900 mb-3">مرحله 2: تنظیمات {{ getMethodName() }}</h4>
            
            <div class="space-y-3">
              <!-- تنظیمات SMS -->
              <div v-if="selectedMethod === 'sms'">
                <label class="block text-sm font-medium text-gray-700 mb-1">شماره تلفن</label>
                <input
                  v-model="setupData.phone"
                  type="tel"
                  placeholder="09123456789"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-sm"
                />
              </div>

              <!-- تنظیمات Email -->
              <div v-if="selectedMethod === 'email'">
                <label class="block text-sm font-medium text-gray-700 mb-1">ایمیل</label>
                <input
                  v-model="setupData.email"
                  type="email"
                  placeholder="example@domain.com"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-sm"
                />
              </div>

              <!-- تنظیمات App -->
              <div v-if="selectedMethod === 'app'">
                <div class="bg-blue-50 p-3 rounded-lg">
                  <p class="text-sm text-blue-800 mb-2">QR کد یا کد دستی برای Google Authenticator</p>
                  <div class="text-center">
                    <div class="w-32 h-32 bg-gray-200 rounded-lg flex items-center justify-center mx-auto mb-2">
                      <span class="text-gray-500 text-xs">QR Code</span>
                    </div>
                    <code class="bg-gray-100 px-2 py-1 rounded text-xs font-mono">JBSWY3DPEHPK3PXP</code>
                  </div>
                </div>
              </div>

              <!-- تنظیمات Backup -->
              <div v-if="selectedMethod === 'backup'">
                <div class="bg-yellow-50 border border-yellow-200 rounded-lg p-3">
                  <p class="text-sm text-yellow-800">کدهای پشتیبان شما:</p>
                  <div class="grid grid-cols-4 gap-1 mt-2">
                    <div v-for="(code, index) in backupCodes.slice(0, 4)" :key="index" class="bg-gray-100 p-1 rounded text-center font-mono text-xs">
                      {{ code }}
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <div class="flex justify-between mt-4">
              <button
                class="px-4 py-2 bg-gray-500 text-white rounded-lg hover:bg-gray-600 transition-colors text-sm"
                @click="previousStep"
              >
                مرحله قبل
              </button>
              <button
                :disabled="!canProceedToNext"
                :class="[
                  'px-4 py-2 rounded-lg transition-colors text-sm',
                  canProceedToNext ? 'bg-blue-600 text-white hover:bg-blue-700' : 'bg-gray-300 text-gray-500 cursor-not-allowed'
                ]"
                @click="nextStep"
              >
                مرحله بعد
              </button>
            </div>
          </div>

          <!-- مرحله 3: تایید -->
          <div v-if="setupStep === 3">
            <h4 class="text-md font-medium text-gray-900 mb-3">مرحله 3: تایید و فعال‌سازی</h4>
            
            <div class="space-y-3">
              <div v-if="selectedMethod !== 'backup'">
                <label class="block text-sm font-medium text-gray-700 mb-1">کد تایید</label>
                <input
                  v-model="setupData.verificationCode"
                  type="text"
                  maxlength="6"
                  placeholder="123456"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-center text-lg font-mono"
                />
              </div>

              <div class="bg-blue-50 border border-blue-200 rounded-lg p-3">
                <p class="text-sm text-blue-800">پس از فعال‌سازی، در هر ورود نیاز به کد تایید خواهید داشت.</p>
              </div>
            </div>

            <div class="flex justify-between mt-4">
              <button
                class="px-4 py-2 bg-gray-500 text-white rounded-lg hover:bg-gray-600 transition-colors text-sm"
                @click="previousStep"
              >
                مرحله قبل
              </button>
              <button
                :disabled="!canActivate"
                :class="[
                  'px-4 py-2 rounded-lg transition-colors text-sm flex items-center gap-1',
                  canActivate ? 'bg-green-600 text-white hover:bg-green-700' : 'bg-gray-300 text-gray-500 cursor-not-allowed'
                ]"
                @click="activateTwoFactor"
              >
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                </svg>
                فعال‌سازی
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- جدول کاربران -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                کاربر
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                وضعیت
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                نوع دستگاه
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                آخرین ورود
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                عملیات
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="user in paginatedUsers" :key="user.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="">
                    <img class="h-10 w-10 rounded-full" :src="user.avatar" :alt="user.name" />
                  </div>
                  <div class="mr-4">
                    <div class="text-sm font-medium text-gray-900">{{ user.name }}</div>
                    <div class="text-sm text-gray-500">{{ user.email }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span
                  :class="[
                    'inline-flex px-2 py-1 text-xs font-semibold rounded-full',
                    user.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'
                  ]"
                >
                  {{ user.status === 'active' ? 'فعال' : 'غیرفعال' }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ user.deviceType === 'mobile' ? 'موبایل' : user.deviceType === 'desktop' ? 'دسکتاپ' : 'تبلت' }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ user.lastLogin }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <div class="flex items-center gap-2">
                  <button
                    class="text-blue-600 hover:text-blue-900 transition-colors"
                    title="مشاهده جزئیات"
                    @click="viewUser(user.id)"
                  >
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                    </svg>
                  </button>
                  <button
                    :class="[
                      'transition-colors',
                      user.status === 'active' ? 'text-red-600 hover:text-red-900' : 'text-green-600 hover:text-green-900'
                    ]"
                    :title="user.status === 'active' ? 'غیرفعال کردن' : 'فعال کردن'"
                    @click="toggleTwoFactor(user.id)"
                  >
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
                    </svg>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- صفحه‌بندی -->
      <div class="bg-white px-4 py-3 border-t border-gray-200 sm:px-6">
        <div class="flex items-center justify-between">
          <div class="flex-1 flex justify-between sm:hidden">
            <button
              :disabled="currentPage === 1"
              :class="[
                'relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md',
                currentPage === 1 ? 'bg-gray-100 text-gray-400 cursor-not-allowed' : 'bg-white text-gray-700 hover:bg-gray-50'
              ]"
              @click="previousPage"
            >
              قبلی
            </button>
            <button
              :disabled="currentPage === totalPages"
              :class="[
                'relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md',
                currentPage === totalPages ? 'bg-gray-100 text-gray-400 cursor-not-allowed' : 'bg-white text-gray-700 hover:bg-gray-50'
              ]"
              @click="nextPage"
            >
              بعدی
            </button>
          </div>
          <div class="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
            <div>
              <p class="text-sm text-gray-700">
                نمایش
                <span class="font-medium">{{ startIndex + 1 }}</span>
                تا
                <span class="font-medium">{{ endIndex }}</span>
                از
                <span class="font-medium">{{ filteredUsers.length }}</span>
                نتیجه
              </p>
            </div>
            <div>
              <nav class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px">
                <button
                  :disabled="currentPage === 1"
                  :class="[
                    'relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50',
                    currentPage === 1 ? 'cursor-not-allowed opacity-50' : ''
                  ]"
                  @click="previousPage"
                >
                  <span class="sr-only">قبلی</span>
                  <svg class="h-5 w-5" fill="currentColor" viewBox="0 0 20 20">
                    <path fill-rule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd" />
                  </svg>
                </button>
                
                <button
                  v-for="page in visiblePages"
                  :key="page"
                  :class="[
                    'relative inline-flex items-center px-4 py-2 border text-sm font-medium',
                    page === currentPage
                      ? 'z-10 bg-blue-50 border-blue-500 text-blue-600'
                      : 'bg-white border-gray-300 text-gray-500 hover:bg-gray-50'
                  ]"
                  @click="goToPage(page)"
                >
                  {{ page }}
                </button>
                
                <button
                  :disabled="currentPage === totalPages"
                  :class="[
                    'relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50',
                    currentPage === totalPages ? 'cursor-not-allowed opacity-50' : ''
                  ]"
                  @click="nextPage"
                >
                  <span class="sr-only">بعدی</span>
                  <svg class="h-5 w-5" fill="currentColor" viewBox="0 0 20 20">
                    <path fill-rule="evenodd" d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z" clip-rule="evenodd" />
                  </svg>
                </button>
              </nav>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// متغیرهای واکنشی
const searchQuery = ref('')
const statusFilter = ref('')
const deviceTypeFilter = ref('')
const dateFilter = ref('')
const currentPage = ref(1)
const itemsPerPage = ref(10)
const showCreateForm = ref(false)
const showSetupForm = ref(false)

// متغیرهای فرم ایجاد تایید دو مرحله‌ای
const newTwoFactor = ref({
  name: '',
  email: '',
  phone: '',
  type: '',
  expiry: 5,
  maxAttempts: 3,
  description: '',
  forceEnable: false,
  requireOnLogin: false,
  sendNotification: true
})

// متغیرهای فعال‌سازی تایید دو مرحله‌ای
const setupStep = ref(1)
const selectedMethod = ref('')
const setupData = ref({
  phone: '',
  email: '',
  verificationCode: ''
})

// کدهای پشتیبان نمونه
const backupCodes = ref([
  'ABC123',
  'DEF456',
  'GHI789',
  'JKL012',
  'MNO345',
  'PQR678',
  'STU901',
  'VWX234'
])

// داده‌های نمونه
const users = ref([
  {
    id: 1,
    name: 'علی احمدی',
    email: 'ali@example.com',
    avatar: '/default-avatar.png',
    status: 'active',
    deviceType: 'mobile',
    lastLogin: '2024-01-15 14:30'
  },
  {
    id: 2,
    name: 'فاطمه محمدی',
    email: 'fateme@example.com',
    avatar: '/default-avatar.png',
    status: 'inactive',
    deviceType: 'desktop',
    lastLogin: '2024-01-14 09:15'
  },
  {
    id: 3,
    name: 'محمد رضایی',
    email: 'mohammad@example.com',
    avatar: '/default-avatar.png',
    status: 'active',
    deviceType: 'tablet',
    lastLogin: '2024-01-15 16:45'
  },
  {
    id: 4,
    name: 'زهرا کریمی',
    email: 'zahra@example.com',
    avatar: '/default-avatar.png',
    status: 'active',
    deviceType: 'mobile',
    lastLogin: '2024-01-15 11:20'
  },
  {
    id: 5,
    name: 'حسین نوری',
    email: 'hossein@example.com',
    avatar: '/default-avatar.png',
    status: 'inactive',
    deviceType: 'desktop',
    lastLogin: '2024-01-13 08:30'
  }
])

// محاسبه کاربران فیلتر شده
const filteredUsers = computed(() => {
  return users.value.filter(user => {
    const matchesSearch = !searchQuery.value || 
      user.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      user.email.toLowerCase().includes(searchQuery.value.toLowerCase())
    
    const matchesStatus = !statusFilter.value || user.status === statusFilter.value
    const matchesDevice = !deviceTypeFilter.value || user.deviceType === deviceTypeFilter.value
    
    return matchesSearch && matchesStatus && matchesDevice
  })
})

// محاسبه کاربران صفحه‌بندی شده
const paginatedUsers = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage.value
  const end = start + itemsPerPage.value
  return filteredUsers.value.slice(start, end)
})

// محاسبه تعداد کل صفحات
const totalPages = computed(() => {
  return Math.ceil(filteredUsers.value.length / itemsPerPage.value)
})

// محاسبه شاخص‌های شروع و پایان
const startIndex = computed(() => {
  return (currentPage.value - 1) * itemsPerPage.value
})

const endIndex = computed(() => {
  return Math.min(startIndex.value + itemsPerPage.value, filteredUsers.value.length)
})

// محاسبه صفحات قابل مشاهده
const visiblePages = computed(() => {
  const pages = []
  const maxVisible = 5
  let start = Math.max(1, currentPage.value - Math.floor(maxVisible / 2))
  const end = Math.min(totalPages.value, start + maxVisible - 1)
  
  if (end - start + 1 < maxVisible) {
    start = Math.max(1, end - maxVisible + 1)
  }
  
  for (let i = start; i <= end; i++) {
    pages.push(i)
  }
  
  return pages
})

// آمار کاربران
const activeUsersCount = computed(() => {
  return filteredUsers.value.filter(user => user.status === 'active').length
})

const inactiveUsersCount = computed(() => {
  return filteredUsers.value.filter(user => user.status === 'inactive').length
})

// توابع صفحه‌بندی
const nextPage = () => {
  if (currentPage.value < totalPages.value) {
    currentPage.value++
  }
}

const previousPage = () => {
  if (currentPage.value > 1) {
    currentPage.value--
  }
}

const goToPage = (page) => {
  currentPage.value = page
}

// توابع فیلتر
const applyFilters = () => {
  currentPage.value = 1 // بازگشت به صفحه اول
}

const clearFilters = () => {
  searchQuery.value = ''
  statusFilter.value = ''
  deviceTypeFilter.value = ''
  dateFilter.value = ''
  currentPage.value = 1
}

// توابع عملیات
const viewUser = (_userId) => {
  // مشاهده جزئیات کاربر
}

const toggleTwoFactor = (userId) => {
  // تغییر وضعیت احراز هویت دو مرحله‌ای
  const user = users.value.find(u => u.id === userId)
  if (user) {
    user.status = user.status === 'active' ? 'inactive' : 'active'
  }
}

// توابع فرم ایجاد تایید دو مرحله‌ای
const createTwoFactor = async () => {
  try {
    // اعتبارسنجی فرم
    if (!newTwoFactor.value.name || !newTwoFactor.value.email || !newTwoFactor.value.phone || !newTwoFactor.value.type) {
      alert('لطفاً تمام فیلدهای الزامی را پر کنید')
      return
    }

    // اعتبارسنجی ایمیل
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
    if (!emailRegex.test(newTwoFactor.value.email)) {
      alert('لطفاً یک ایمیل معتبر وارد کنید')
      return
    }

    // اعتبارسنجی شماره تلفن
    const phoneRegex = /^09\d{9}$/
    if (!phoneRegex.test(newTwoFactor.value.phone)) {
      alert('لطفاً یک شماره تلفن معتبر وارد کنید (مثال: 09123456789)')
      return
    }

    // شبیه‌سازی ارسال درخواست به سرور
    
    // نمایش پیام موفقیت
    alert('تایید دو مرحله‌ای با موفقیت ایجاد شد!')
    
    // پاک کردن فرم
    resetForm()
    
    // بستن فرم
    showCreateForm.value = false
    
    // در اینجا می‌توانید درخواست API ارسال کنید
    // const response = await $fetch('/api/admin/two-factor', {
    //   method: 'POST',
    //   body: newTwoFactor.value
    // })
    
  } catch (error) {
    console.error('خطا در ایجاد تایید دو مرحله‌ای:', error)
    alert('خطا در ایجاد تایید دو مرحله‌ای. لطفاً دوباره تلاش کنید.')
  }
}

const resetForm = () => {
  newTwoFactor.value = {
    name: '',
    email: '',
    phone: '',
    type: '',
    expiry: 5,
    maxAttempts: 3,
    description: '',
    forceEnable: false,
    requireOnLogin: false,
    sendNotification: true
  }
}

// Computed properties برای فعال‌سازی
const canProceedToNext = computed(() => {
  if (setupStep.value === 2) {
    if (selectedMethod.value === 'sms') {
      return setupData.value.phone && /^09\d{9}$/.test(setupData.value.phone)
    } else if (selectedMethod.value === 'email') {
      return setupData.value.email && /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(setupData.value.email)
    } else if (selectedMethod.value === 'app' || selectedMethod.value === 'backup') {
      return true
    }
  }
  return false
})

const canActivate = computed(() => {
  if (setupStep.value === 3) {
    if (selectedMethod.value === 'backup') {
      return true
    } else {
      return setupData.value.verificationCode && setupData.value.verificationCode.length === 6
    }
  }
  return false
})

// توابع فعال‌سازی تایید دو مرحله‌ای
const selectMethod = (method) => {
  selectedMethod.value = method
}

const nextStep = () => {
  if (setupStep.value < 3) {
    setupStep.value++
  }
}

const previousStep = () => {
  if (setupStep.value > 1) {
    setupStep.value--
  }
}

const getMethodName = () => {
  const methodNames = {
    sms: 'پیامک (SMS)',
    email: 'ایمیل',
    app: 'اپلیکیشن',
    backup: 'کدهای پشتیبان'
  }
  return methodNames[selectedMethod.value] || ''
}

const activateTwoFactor = async () => {
  try {
    // اعتبارسنجی نهایی
    if (selectedMethod.value !== 'backup' && (!setupData.value.verificationCode || setupData.value.verificationCode.length !== 6)) {
      alert('لطفاً کد تایید 6 رقمی را وارد کنید')
      return
    }

    // شبیه‌سازی فعال‌سازی

    // نمایش پیام موفقیت
    alert('تایید دو مرحله‌ای با موفقیت فعال شد!')
    
    // ریست کردن فرم
    resetSetupForm()
    
    // بستن فرم
    showSetupForm.value = false
    
    // در اینجا می‌توانید درخواست API ارسال کنید
    // const response = await $fetch('/api/admin/two-factor/setup', {
    //   method: 'POST',
    //   body: {
    //     method: selectedMethod.value,
    //     data: setupData.value
    //   }
    // })
    
  } catch (error) {
    console.error('خطا در فعال‌سازی تایید دو مرحله‌ای:', error)
    alert('خطا در فعال‌سازی تایید دو مرحله‌ای. لطفاً دوباره تلاش کنید.')
  }
}

const resetSetupForm = () => {
  setupStep.value = 1
  selectedMethod.value = ''
  setupData.value = {
    phone: '',
    email: '',
    verificationCode: ''
  }
}

// تابع خروجی اکسل
const exportToExcel = () => {
  try {
    // آماده‌سازی داده‌ها برای اکسل
    const excelData = filteredUsers.value.map(user => ({
      'نام کاربر': user.name,
      'ایمیل': user.email,
      'وضعیت': user.status === 'active' ? 'فعال' : 'غیرفعال',
      'نوع دستگاه': user.deviceType === 'mobile' ? 'موبایل' : user.deviceType === 'desktop' ? 'دسکتاپ' : 'تبلت',
      'آخرین ورود': user.lastLogin
    }))

    // ایجاد CSV
    const headers = Object.keys(excelData[0])
    const csvContent = [
      headers.join(','),
      ...excelData.map(row => headers.map(header => `"${row[header]}"`).join(','))
    ].join('\n')

    // اضافه کردن BOM برای پشتیبانی از فارسی
    const BOM = '\uFEFF'
    const blob = new Blob([BOM + csvContent], { type: 'text/csv;charset=utf-8;' })
    
    // دانلود فایل
    const link = document.createElement('a')
    const url = URL.createObjectURL(blob)
    link.setAttribute('href', url)
    link.setAttribute('download', `two-factor-auth-report-${new Date().toISOString().split('T')[0]}.csv`)
    link.style.visibility = 'hidden'
    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link)

    alert('فایل اکسل با موفقیت دانلود شد!')
    
  } catch (error) {
    console.error('خطا در ایجاد خروجی اکسل:', error)
    alert('خطا در ایجاد خروجی اکسل. لطفاً دوباره تلاش کنید.')
  }
}

// نظارت بر تغییرات فیلترها
watch([searchQuery, statusFilter, deviceTypeFilter, dateFilter], () => {
  currentPage.value = 1
})
</script>

<style scoped>
/* استایل‌های سفارشی برای صفحه */
</style> 

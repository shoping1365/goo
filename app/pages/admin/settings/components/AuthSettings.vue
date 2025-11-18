<template>
  <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
    <h2 class="text-2xl font-bold text-gray-900 mb-6 flex items-center">
      <i class="i-heroicons-shield-check text-blue-600 mr-3"></i>
      تنظیمات احراز هویت
    </h2>

    <!-- تب‌های احراز هویت -->
    <div class="border-b border-gray-200 mb-6">
      <nav class="-mb-px flex space-x-8" aria-label="Tabs">
        <button
          @click="$emit('update:auth-active-tab', 'login-register')"
          :class="[
            'whitespace-nowrap py-2 px-1 border-b-2 font-medium text-sm',
            authActiveTab === 'login-register'
              ? 'border-blue-500 text-blue-600'
              : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
          ]"
        >
          <i class="i-heroicons-user-plus ml-2"></i>
          ورود و ثبت‌نام
        </button>
        <button
          @click="$emit('update:auth-active-tab', 'security')"
          :class="[
            'whitespace-nowrap py-2 px-1 border-b-2 font-medium text-sm',
            authActiveTab === 'security'
              ? 'border-blue-500 text-blue-600'
              : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
          ]"
        >
          <i class="i-heroicons-shield-exclamation ml-2"></i>
          امنیت و محدودیت‌ها
        </button>
        <button
          @click="$emit('update:auth-active-tab', 'jwt')"
          :class="[
            'whitespace-nowrap py-2 px-1 border-b-2 font-medium text-sm',
            authActiveTab === 'jwt'
              ? 'border-blue-500 text-blue-600'
              : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
          ]"
        >
          <i class="i-heroicons-key ml-2"></i>
          تنظیمات JWT
        </button>
      </nav>
    </div>

    <div v-if="authActiveTab === 'login-register'" class="space-y-8">
      <!-- محتوای تب ورود و ثبت‌نام -->
      <!-- تنظیمات سیستم یکپارچه ورود/ثبت‌نام -->
      <div class="mb-8">
        <h3 class="text-lg font-semibold text-gray-800 mb-4 flex items-center">
          <i class="i-heroicons-user-plus text-blue-600 mr-2"></i>
          سیستم یکپارچه ورود و ثبت‌نام
        </h3>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <!-- فعال‌سازی سیستم یکپارچه -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700">
              فعال‌سازی سیستم یکپارچه
            </label>
            <div class="flex items-center">
              <input
                v-model="authSettings.unifiedAuthEnabled"
                type="checkbox"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <span class="ml-2 text-sm text-gray-600">فعال</span>
            </div>
            <p class="text-xs text-gray-500">کاربران با یک فیلد موبایل وارد می‌شوند (خودکار ثبت‌نام در صورت عدم وجود)</p>
          </div>

          <!-- ثبت‌نام خودکار -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700">
              ثبت‌نام خودکار
            </label>
            <div class="flex items-center">
              <input
                v-model="authSettings.autoRegisterEnabled"
                type="checkbox"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <span class="ml-2 text-sm text-gray-600">فعال</span>
            </div>
            <p class="text-xs text-gray-500">کاربران جدید به صورت خودکار ثبت‌نام می‌شوند</p>
          </div>

          <!-- انتخاب روش پیش‌فرض -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700">
              روش احراز هویت پیش‌فرض
            </label>
            <select
              v-model="authSettings.defaultAuthMethod"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="otp">کد تایید (OTP)</option>
              <option value="password">پسورد</option>
            </select>
            <p class="text-xs text-gray-500">روش پیش‌فرض نمایش داده شده به کاربر</p>
          </div>

          <!-- نقش پیش‌فرض کاربران جدید -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700">
              نقش پیش‌فرض کاربران جدید
            </label>
            <select
              v-model="authSettings.defaultUserRole"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="user">کاربر عادی</option>
              <option value="customer">مشتری</option>
              <option value="guest">مهمان</option>
            </select>
            <p class="text-xs text-gray-500">نقش اختصاص داده شده به کاربران تازه ثبت‌نام شده</p>
          </div>
        </div>
      </div>

      <!-- تنظیمات OTP -->
      <div class="mb-8">
        <h3 class="text-lg font-semibold text-gray-800 mb-4 flex items-center">
          <i class="i-heroicons-device-phone-mobile text-green-600 mr-2"></i>
          تنظیمات احراز هویت با موبایل (OTP)
        </h3>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <!-- فعال‌سازی OTP -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700">
              فعال‌سازی احراز هویت با موبایل
            </label>
            <div class="flex items-center">
              <input
                v-model="authSettings.mobileAuthEnabled"
                type="checkbox"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <span class="ml-2 text-sm text-gray-600">فعال</span>
            </div>
            <p class="text-xs text-gray-500">کاربران می‌توانند با شماره موبایل و کد تایید وارد شوند</p>
          </div>

          <!-- طول کد OTP -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700">
              طول کد تایید
            </label>
            <select
              v-model="authSettings.otpLength"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="4">4 رقم</option>
              <option value="5">5 رقم</option>
              <option value="6">6 رقم</option>
            </select>
            <p class="text-xs text-gray-500">تعداد ارقام کد تایید ارسالی</p>
          </div>

          <!-- مدت اعتبار OTP -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700">
              مدت اعتبار کد تایید (دقیقه)
            </label>
            <input
              v-model="authSettings.otpExpiryMinutes"
              type="number"
              min="1"
              max="60"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
            <p class="text-xs text-gray-500">زمان اعتبار کد تایید به دقیقه</p>
          </div>

          <!-- حداکثر تلاش OTP -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700">
              حداکثر تلاش برای کد تایید
            </label>
            <input
              v-model="authSettings.maxOtpAttempts"
              type="number"
              min="1"
              max="10"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
            <p class="text-xs text-gray-500">تعداد تلاش مجاز برای وارد کردن کد تایید</p>
          </div>

          <!-- فاصله ارسال مجدد -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700">
              فاصله ارسال مجدد (ثانیه)
            </label>
            <input
              v-model="authSettings.otpResendCooldown"
              type="number"
              min="30"
              max="300"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
            <p class="text-xs text-gray-500">حداقل فاصله زمانی برای ارسال مجدد کد</p>
          </div>

          <!-- کد پترن SMS -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700">
              کد پترن SMS
            </label>
            <input
              v-model="authSettings.otpPatternCode"
              type="text"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="مثال: t2cfmnyo0c"
            />
            <p class="text-xs text-gray-500">کد پترن SMS برای ارسال کد تایید (از جدول sms_patterns)</p>
          </div>
        </div>
      </div>

      <!-- تنظیمات ورود با یوزرنیم -->
      <div class="mb-8">
        <h3 class="text-lg font-semibold text-gray-800 mb-4 flex items-center">
          <i class="i-heroicons-user-circle text-indigo-600 mr-2"></i>
          تنظیمات ورود با یوزرنیم و پسورد
        </h3>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <!-- فعال‌سازی ورود با یوزرنیم -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700">
              فعال‌سازی ورود با یوزرنیم
            </label>
            <div class="flex items-center">
              <input
                v-model="authSettings.usernameAuthEnabled"
                type="checkbox"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <span class="ml-2 text-sm text-gray-600">فعال</span>
            </div>
            <p class="text-xs text-gray-500">کاربران می‌توانند با یوزرنیم و پسورد وارد شوند</p>
          </div>

          <!-- حداقل طول پسورد -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700">
              حداقل طول پسورد
            </label>
            <input
              v-model="authSettings.minPasswordLength"
              type="number"
              min="6"
              max="20"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
            <p class="text-xs text-gray-500">حداقل تعداد کاراکترهای پسورد</p>
          </div>

          <!-- حداکثر تلاش ورود -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700">
              حداکثر تلاش ورود
            </label>
            <input
              v-model="authSettings.maxLoginAttempts"
              type="number"
              min="3"
              max="10"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
            <p class="text-xs text-gray-500">تعداد تلاش مجاز برای ورود</p>
          </div>

          <!-- مدت قفل حساب -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700">
              مدت قفل حساب (دقیقه)
            </label>
            <input
              v-model="authSettings.accountLockoutMinutes"
              type="number"
              min="5"
              max="1440"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
            <p class="text-xs text-gray-500">مدت زمان قفل شدن حساب پس از تلاش‌های ناموفق</p>
          </div>

          <!-- مدت اعتبار پسورد -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700">
              مدت اعتبار پسورد (روز)
            </label>
            <input
              v-model="authSettings.passwordExpiryDays"
              type="number"
              min="30"
              max="365"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
            <p class="text-xs text-gray-500">پس از این مدت کاربر باید پسورد خود را تغییر دهد</p>
          </div>
        </div>
      </div>
    </div>

    <div v-else-if="authActiveTab === 'security'" class="space-y-8">
      <!-- محتوای تب امنیت و محدودیت‌ها -->
      <!-- تنظیمات امنیتی -->
      <div class="mb-8">
        <h3 class="text-lg font-semibold text-gray-800 mb-4 flex items-center">
          <i class="i-heroicons-shield-exclamation text-red-600 mr-2"></i>
          تنظیمات امنیتی پیشرفته
        </h3>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <!-- احراز هویت دو مرحله‌ای -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700">
              احراز هویت دو مرحله‌ای
            </label>
            <div class="flex items-center">
              <input
                v-model="authSettings.twoFactorEnabled"
                type="checkbox"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <span class="ml-2 text-sm text-gray-600">فعال</span>
            </div>
            <p class="text-xs text-gray-500">نیاز به تایید اضافی پس از ورود</p>
          </div>

          <!-- تشخیص فعالیت مشکوک -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700">
              تشخیص فعالیت مشکوک
            </label>
            <div class="flex items-center">
              <input
                v-model="authSettings.suspiciousActivityDetection"
                type="checkbox"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <span class="ml-2 text-sm text-gray-600">فعال</span>
            </div>
            <p class="text-xs text-gray-500">نظارت خودکار بر ورودهای مشکوک</p>
          </div>

          <!-- تایم‌اوت نشست -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700">
              تایم‌اوت نشست (دقیقه)
            </label>
            <input
              v-model="authSettings.sessionTimeoutMinutes"
              type="number"
              min="5"
              max="1440"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
            <p class="text-xs text-gray-500">زمان خودکار خروج کاربر از سیستم</p>
          </div>

          <!-- نگهداری تاریخچه ورود -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700">
              نگهداری تاریخچه ورود (روز)
            </label>
            <input
              v-model="authSettings.loginHistoryRetentionDays"
              type="number"
              min="30"
              max="730"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
            <p class="text-xs text-gray-500">مدت نگهداری تاریخچه ورود کاربران</p>
          </div>

          <!-- مسدود کردن خودکار -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700">
              مسدود کردن خودکار پس از ورود ناموفق
            </label>
            <input
              v-model="authSettings.autoBlockFailedLogins"
              type="number"
              min="5"
              max="50"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
            <p class="text-xs text-gray-500">تعداد تلاش ناموفق قبل از مسدود شدن IP</p>
          </div>

          <!-- مدت مسدود بودن -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700">
              مدت مسدود بودن (ساعت)
            </label>
            <input
              v-model="authSettings.autoBlockDurationHours"
              type="number"
              min="1"
              max="168"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
            <p class="text-xs text-gray-500">مدت زمان مسدود بودن IP مشکوک</p>
          </div>
        </div>
      </div>

      <!-- تنظیمات تایید هویت -->
      <div class="mb-8">
        <h3 class="text-lg font-semibold text-gray-800 mb-4 flex items-center">
          <i class="i-heroicons-identification text-purple-600 mr-2"></i>
          تنظیمات تایید هویت
        </h3>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <!-- تایید ایمیل -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700">
              تایید ایمیل
            </label>
            <div class="flex items-center">
              <input
                v-model="authSettings.emailVerificationEnabled"
                type="checkbox"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <span class="ml-2 text-sm text-gray-600">فعال</span>
            </div>
            <p class="text-xs text-gray-500">الزام تایید ایمیل برای ثبت‌نام</p>
          </div>

          <!-- تایید موبایل -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700">
              تایید موبایل
            </label>
            <div class="flex items-center">
              <input
                v-model="authSettings.phoneVerificationEnabled"
                type="checkbox"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <span class="ml-2 text-sm text-gray-600">فعال</span>
            </div>
            <p class="text-xs text-gray-500">الزام تایید شماره موبایل برای ثبت‌نام</p>
          </div>
        </div>
      </div>
    </div>

    <div v-else-if="authActiveTab === 'jwt'" class="space-y-8">
      <!-- محتوای تب تنظیمات JWT -->
      <div class="bg-yellow-50 border border-yellow-200 rounded-xl p-6">
        <div class="flex items-center mb-4">
          <i class="i-heroicons-exclamation-triangle text-yellow-600 mr-2"></i>
          <h4 class="text-lg font-semibold text-yellow-800">هشدار امنیتی</h4>
        </div>
        <p class="text-yellow-700">
          تغییر تنظیمات JWT باعث خروج تمام کاربران از سیستم می‌شود. این تغییرات را با احتیاط انجام دهید.
        </p>
      </div>

      <!-- کلید مخفی JWT -->
      <div class="mb-8">
        <h3 class="text-lg font-semibold text-gray-800 mb-4 flex items-center">
          <i class="i-heroicons-key text-purple-600 mr-2"></i>
          مدیریت کلید مخفی
        </h3>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <!-- کلید مخفی فعلی -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700">
              کلید مخفی فعلی
            </label>
            <div class="relative">
              <input
                v-model="authSettings.jwtSecret"
                :type="showJwtSecret ? 'text' : 'password'"
                class="w-full px-3 py-2 pr-10 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="کلید مخفی JWT"
                readonly
              />
              <button
                @click="$emit('toggle-jwt-secret')"
                type="button"
                class="absolute inset-y-0 right-0 pr-3 flex items-center"
              >
                <i :class="showJwtSecret ? 'i-heroicons-eye-slash' : 'i-heroicons-eye'" class="text-gray-400 hover:text-gray-600"></i>
              </button>
            </div>
            <p class="text-xs text-gray-500">کلید مخفی برای امضای توکن‌های JWT</p>
          </div>

          <!-- تولید کلید جدید -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700">
              تولید کلید جدید
            </label>
            <button
              @click="$emit('generate-new-jwt-secret')"
              type="button"
              class="w-full px-4 py-2 bg-orange-500 text-white rounded-lg hover:bg-orange-600 transition-colors"
            >
              <i class="i-heroicons-arrow-path mr-2"></i>
              تولید کلید جدید
            </button>
            <p class="text-xs text-gray-500">تولید کلید مخفی جدید (تمام توکن‌های فعلی باطل می‌شوند)</p>
          </div>
        </div>
      </div>

      <!-- تنظیمات زمان‌بندی -->
      <div class="mb-8">
        <h3 class="text-lg font-semibold text-gray-800 mb-4 flex items-center">
          <i class="i-heroicons-clock text-green-600 mr-2"></i>
          تنظیمات زمان‌بندی
        </h3>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <!-- مدت اعتبار توکن -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700">
              مدت اعتبار توکن (ساعت)
            </label>
            <input
              v-model="authSettings.jwtExpiryHours"
              type="number"
              min="1"
              max="720"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
            <p class="text-xs text-gray-500">مدت زمان اعتبار توکن JWT به ساعت</p>
          </div>

          <!-- مدت اعتبار توکن تازه‌سازی -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700">
              مدت اعتبار توکن تازه‌سازی (روز)
            </label>
            <input
              v-model="authSettings.refreshTokenExpiryDays"
              type="number"
              min="1"
              max="365"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
            <p class="text-xs text-gray-500">مدت زمان اعتبار توکن تازه‌سازی به روز</p>
          </div>
        </div>
      </div>

      <!-- تنظیمات نشست -->
      <div class="mb-8">
        <h3 class="text-lg font-semibold text-gray-800 mb-4 flex items-center">
          <i class="i-heroicons-device-tablet text-blue-600 mr-2"></i>
          مدیریت نشست‌ها
        </h3>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <!-- حداکثر نشست‌های همزمان -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700">
              حداکثر نشست‌های همزمان
            </label>
            <input
              v-model="authSettings.maxConcurrentSessions"
              type="number"
              min="1"
              max="20"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
            <p class="text-xs text-gray-500">حداکثر تعداد نشست‌های همزمان برای هر کاربر</p>
          </div>

          <!-- تایم‌اوت نشست -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700">
              تایم‌اوت نشست (دقیقه)
            </label>
            <input
              v-model="authSettings.sessionTimeoutMinutes"
              type="number"
              min="5"
              max="1440"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
            <p class="text-xs text-gray-500">زمان خودکار خروج کاربر از سیستم</p>
          </div>
        </div>
      </div>
    </div>

    <!-- دکمه ذخیره -->
    <div class="mt-8 flex justify-end">
      <button
        @click="$emit('save')"
        :disabled="savingAuth"
        class="px-6 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:opacity-50 transition-colors flex items-center"
      >
        <i v-if="savingAuth" class="i-heroicons-arrow-path animate-spin mr-2"></i>
        <i v-else class="i-heroicons-check mr-2"></i>
        {{ savingAuth ? 'در حال ذخیره...' : 'ذخیره تنظیمات' }}
      </button>
    </div>
  </div>
</template>

<script setup>
// تعریف Props
defineProps({
  authSettings: {
    type: Object,
    required: true
  },
  authActiveTab: {
    type: String,
    default: 'login-register'
  },
  showJwtSecret: {
    type: Boolean,
    default: false
  },
  savingAuth: {
    type: Boolean,
    default: false
  }
})

// تعریف Emits
defineEmits([
  'save',
  'generate-new-jwt-secret',
  'update:auth-active-tab',
  'toggle-jwt-secret'
])
</script>

<template>
  <div>
    <div class="mb-6">
      <h2 class="text-xl font-semibold text-gray-900 mb-2">تنظیمات احراز هویت</h2>
      <p class="text-gray-600">تنظیم پارامترهای سیستم احراز هویت</p>
    </div>

    <form class="space-y-8" @submit.prevent="saveSettings">
      <!-- OTP Settings -->
      <div class="bg-gray-50 px-4 py-4 rounded-lg">
        <h3 class="text-lg font-medium text-gray-900 mb-4">تنظیمات OTP</h3>
        <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              احراز هویت با موبایل
            </label>
            <div class="flex items-center">
              <input
                v-model="authSettings.mobileAuthEnabled"
                type="checkbox"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <span class="mr-2 text-sm text-gray-600">فعال</span>
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              طول کد OTP
            </label>
            <select
              v-model="authSettings.otpLength"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            >
              <option value="4">4 رقم</option>
              <option value="5">5 رقم</option>
              <option value="6">6 رقم</option>
            </select>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              مدت اعتبار OTP (دقیقه)
            </label>
            <input
              v-model.number="authSettings.otpExpiryMinutes"
              type="number"
              min="1"
              max="60"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              حداکثر تلاش OTP
            </label>
            <input
              v-model.number="authSettings.maxOtpAttempts"
              type="number"
              min="1"
              max="10"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              فاصله ارسال مجدد (ثانیه)
            </label>
            <input
              v-model.number="authSettings.otpResendCooldown"
              type="number"
              min="30"
              max="300"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>
        </div>
      </div>

      <!-- JWT Settings -->
      <div class="bg-gray-50 px-4 py-4 rounded-lg">
        <h3 class="text-lg font-medium text-gray-900 mb-4">تنظیمات JWT</h3>
        <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              مدت اعتبار توکن (ساعت)
            </label>
            <input
              v-model.number="authSettings.jwtExpiryHours"
              type="number"
              min="1"
              max="168"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              مدت اعتبار توکن تازه‌سازی (روز)
            </label>
            <input
              v-model.number="authSettings.refreshTokenExpiryDays"
              type="number"
              min="1"
              max="365"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              حداکثر نشست‌های همزمان
            </label>
            <input
              v-model.number="authSettings.maxConcurrentSessions"
              type="number"
              min="1"
              max="20"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              کلید JWT
            </label>
            <div class="flex space-x-2">
              <input
                v-model="authSettings.jwtSecret"
                :type="showJwtSecret ? 'text' : 'password'"
                class="flex-1 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                placeholder="کلید JWT"
              />
              <button
                type="button"
                class="px-3 py-2 border border-gray-300 rounded-md hover:bg-gray-50"
                @click="showJwtSecret = !showJwtSecret"
              >
                <span v-if="showJwtSecret">مخفی</span>
                <span v-else>نمایش</span>
              </button>
              <button
                type="button"
                class="px-3 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700"
                @click="generateNewJwtSecret"
              >
                تولید جدید
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- Username/Password Settings -->
      <div class="bg-gray-50 px-4 py-4 rounded-lg">
        <h3 class="text-lg font-medium text-gray-900 mb-4">تنظیمات یوزرنیم و پسورد</h3>
        <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              احراز هویت با یوزرنیم
            </label>
            <div class="flex items-center">
              <input
                v-model="authSettings.usernameAuthEnabled"
                type="checkbox"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <span class="mr-2 text-sm text-gray-600">فعال</span>
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              حداقل طول پسورد
            </label>
            <input
              v-model.number="authSettings.minPasswordLength"
              type="number"
              min="6"
              max="50"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              حداکثر تلاش ورود
            </label>
            <input
              v-model.number="authSettings.maxLoginAttempts"
              type="number"
              min="3"
              max="20"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              مدت قفل حساب (دقیقه)
            </label>
            <input
              v-model.number="authSettings.accountLockoutMinutes"
              type="number"
              min="5"
              max="1440"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              مدت اعتبار پسورد (روز)
            </label>
            <input
              v-model.number="authSettings.passwordExpiryDays"
              type="number"
              min="30"
              max="365"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>
        </div>
      </div>

      <!-- Security Settings -->
      <div class="bg-gray-50 px-4 py-4 rounded-lg">
        <h3 class="text-lg font-medium text-gray-900 mb-4">تنظیمات امنیتی</h3>
        <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              احراز هویت دو مرحله‌ای
            </label>
            <div class="flex items-center">
              <input
                v-model="authSettings.twoFactorEnabled"
                type="checkbox"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <span class="mr-2 text-sm text-gray-600">فعال</span>
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              تشخیص فعالیت مشکوک
            </label>
            <div class="flex items-center">
              <input
                v-model="authSettings.suspiciousActivityDetection"
                type="checkbox"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <span class="mr-2 text-sm text-gray-600">فعال</span>
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              تایم‌اوت نشست (دقیقه)
            </label>
            <input
              v-model.number="authSettings.sessionTimeoutMinutes"
              type="number"
              min="5"
              max="480"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              نگهداری تاریخچه ورود (روز)
            </label>
            <input
              v-model.number="authSettings.loginHistoryRetentionDays"
              type="number"
              min="30"
              max="1095"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              بلاک خودکار پس از تلاش ناموفق
            </label>
            <input
              v-model.number="authSettings.autoBlockFailedLogins"
              type="number"
              min="5"
              max="50"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              مدت بلاک خودکار (ساعت)
            </label>
            <input
              v-model.number="authSettings.autoBlockDurationHours"
              type="number"
              min="1"
              max="168"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>
        </div>
      </div>

      <!-- Roles Settings -->
      <div class="bg-gray-50 px-4 py-4 rounded-lg">
        <h3 class="text-lg font-medium text-gray-900 mb-4">تنظیمات نقش‌ها</h3>
        <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              نقش پیش‌فرض کاربران
            </label>
            <select
              v-model="authSettings.defaultUserRole"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            >
              <option value="user">کاربر</option>
              <option value="staff">کارمند</option>
            </select>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              تایید ایمیل
            </label>
            <div class="flex items-center">
              <input
                v-model="authSettings.emailVerificationEnabled"
                type="checkbox"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <span class="mr-2 text-sm text-gray-600">فعال</span>
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              تایید موبایل
            </label>
            <div class="flex items-center">
              <input
                v-model="authSettings.phoneVerificationEnabled"
                type="checkbox"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <span class="mr-2 text-sm text-gray-600">فعال</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Save Button -->
      <div class="flex justify-end">
        <button
          type="submit"
          :disabled="savingAuth"
          class="px-6 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          <span v-if="savingAuth">در حال ذخیره...</span>
          <span v-else>ذخیره تنظیمات</span>
        </button>
      </div>
    </form>
  </div>
</template>

<script setup>
import { reactive, ref } from 'vue'

// Reactive data
const savingAuth = ref(false)
const showJwtSecret = ref(false)

const authSettings = reactive({
  // OTP Settings
  mobileAuthEnabled: true,
  otpLength: 5,
  otpExpiryMinutes: 5,
  maxOtpAttempts: 3,
  otpResendCooldown: 60,
  
  // JWT Settings
  jwtExpiryHours: 24,
  refreshTokenExpiryDays: 30,
  jwtSecret: '',
  maxConcurrentSessions: 5,
  
  // Username/Password Settings
  usernameAuthEnabled: true,
  minPasswordLength: 8,
  maxLoginAttempts: 5,
  accountLockoutMinutes: 15,
  passwordExpiryDays: 90,
  
  // Security Settings
  twoFactorEnabled: false,
  suspiciousActivityDetection: true,
  sessionTimeoutMinutes: 30,
  loginHistoryRetentionDays: 365,
  autoBlockFailedLogins: 10,
  autoBlockDurationHours: 24,
  
  // Roles Settings
  defaultUserRole: 'user',
  emailVerificationEnabled: false,
  phoneVerificationEnabled: true
})

// Methods
const generateNewJwtSecret = () => {
  const chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789'
  let result = ''
  for (let i = 0; i < 64; i++) {
    result += chars.charAt(Math.floor(Math.random() * chars.length))
  }
  authSettings.jwtSecret = result
}

const loadAuthSettings = async () => {
  try {
    const { data } = await $fetch('/api/admin/settings/auth')
    if (data) {
      Object.assign(authSettings, data)
    }
  } catch (error) {
    console.error('خطا در بارگذاری تنظیمات:', error)
  }
}

const saveSettings = async () => {
  try {
    savingAuth.value = true
    await $fetch('/api/admin/settings/auth', {
      method: 'POST',
      body: authSettings
    })
    alert('تنظیمات با موفقیت ذخیره شد')
  } catch (error) {
    console.error('خطا در ذخیره تنظیمات:', error)
    alert('خطا در ذخیره تنظیمات')
  } finally {
    savingAuth.value = false
  }
}

// Lifecycle
onMounted(() => {
  loadAuthSettings()
})
</script> 

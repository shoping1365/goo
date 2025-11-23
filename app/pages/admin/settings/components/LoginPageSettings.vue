<template>
  <div v-if="hasAccess" class="login-page-settings">
    <div class="settings-header">
      <h2 class="text-2xl font-bold text-gray-900 mb-4">تنظیمات صفحه ورود و ثبت‌نام</h2>
      <p class="text-gray-600 mb-6">مدیریت تنظیمات مربوط به صفحه ورود و ثبت‌نام کاربران</p>
    </div>

    <div class="settings-content space-y-8">
      <!-- تنظیمات OTP -->
      <div class="settings-section">
        <h3 class="text-lg font-semibold text-gray-800 mb-4">تنظیمات کد تایید (OTP)</h3>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div class="form-group">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              طول کد تایید
            </label>
            <input
              v-model="loginSettings.otpLength"
              type="number"
              min="4"
              max="8"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="5"
            />
            <p class="text-xs text-gray-500 mt-1">تعداد ارقام کد تایید (پیشنهاد: 5)</p>
          </div>

          <div class="form-group">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              مدت اعتبار کد (دقیقه)
            </label>
            <input
              v-model="loginSettings.otpExpiryMinutes"
              type="number"
              min="1"
              max="60"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="5"
            />
            <p class="text-xs text-gray-500 mt-1">مدت زمان اعتبار کد تایید</p>
          </div>

          <div class="form-group">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              حداکثر تلاش برای کد تایید
            </label>
            <input
              v-model="loginSettings.maxOtpAttempts"
              type="number"
              min="1"
              max="10"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="3"
            />
            <p class="text-xs text-gray-500 mt-1">تعداد تلاش‌های مجاز برای وارد کردن کد</p>
          </div>

          <div class="form-group">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              تاخیر ارسال مجدد (ثانیه)
            </label>
            <input
              v-model="loginSettings.otpResendDelaySeconds"
              type="number"
              min="30"
              max="300"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="60"
            />
            <p class="text-xs text-gray-500 mt-1">تاخیر بین ارسال‌های مجدد کد</p>
          </div>
        </div>
      </div>

      <!-- تنظیمات محدودیت ارسال مجدد -->
      <div class="settings-section">
        <h3 class="text-lg font-semibold text-gray-800 mb-4">محدودیت‌های ارسال مجدد</h3>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div class="form-group">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              زمان انتظار تلاش 1 و 2 (دقیقه)
            </label>
            <input
              v-model="loginSettings.firstAttemptsWaitTime"
              type="number"
              min="1"
              max="10"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="2"
            />
            <p class="text-xs text-gray-500 mt-1">زمان انتظار برای تلاش‌های اول و دوم</p>
          </div>

          <div class="form-group">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              زمان انتظار تلاش 3 و 4 (دقیقه)
            </label>
            <input
              v-model="loginSettings.secondAttemptsWaitTime"
              type="number"
              min="1"
              max="15"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="3"
            />
            <p class="text-xs text-gray-500 mt-1">زمان انتظار برای تلاش‌های سوم و چهارم</p>
          </div>

          <div class="form-group">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              زمان انتظار تلاش 5 (دقیقه)
            </label>
            <input
              v-model="loginSettings.fifthAttemptWaitTime"
              type="number"
              min="1"
              max="30"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="5"
            />
            <p class="text-xs text-gray-500 mt-1">زمان انتظار برای تلاش پنجم</p>
          </div>

          <div class="form-group">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              زمان بلاک شدن (دقیقه)
            </label>
            <input
              v-model="loginSettings.blockTimeMinutes"
              type="number"
              min="30"
              max="1440"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="60"
            />
            <p class="text-xs text-gray-500 mt-1">زمان بلاک شدن بعد از تلاش‌های متعدد</p>
          </div>
        </div>
      </div>

      <!-- تنظیمات امنیتی -->
      <div class="settings-section">
        <h3 class="text-lg font-semibold text-gray-800 mb-4">تنظیمات امنیتی</h3>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div class="form-group">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              حداکثر تلاش ورود
            </label>
            <input
              v-model="loginSettings.maxLoginAttempts"
              type="number"
              min="3"
              max="10"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="5"
            />
            <p class="text-xs text-gray-500 mt-1">تعداد تلاش‌های مجاز برای ورود</p>
          </div>

          <div class="form-group">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              مدت قفل حساب (دقیقه)
            </label>
            <input
              v-model="loginSettings.lockoutDurationMinutes"
              type="number"
              min="5"
              max="1440"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="15"
            />
            <p class="text-xs text-gray-500 mt-1">مدت زمان قفل شدن حساب بعد از تلاش‌های ناموفق</p>
          </div>

          <div class="form-group">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              مدت نشست (دقیقه)
            </label>
            <input
              v-model="loginSettings.sessionTimeoutMinutes"
              type="number"
              min="15"
              max="1440"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="1440"
            />
            <p class="text-xs text-gray-500 mt-1">مدت زمان اعتبار نشست کاربر (24 ساعت = 1440 دقیقه)</p>
          </div>

          <div class="form-group">
            <label class="flex items-center">
              <input
                v-model="loginSettings.mobileAuthEnabled"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
              />
              <span class="ml-2 text-sm font-medium text-gray-700">فعال‌سازی احراز هویت با موبایل</span>
            </label>
            <p class="text-xs text-gray-500 mt-1">اجازه ورود با کد تایید موبایل</p>
          </div>
        </div>
      </div>

      <!-- تنظیمات UI -->
      <div class="settings-section">
        <h3 class="text-lg font-semibold text-gray-800 mb-4">تنظیمات رابط کاربری</h3>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div class="form-group">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              عنوان صفحه ورود
            </label>
            <input
              v-model="loginSettings.loginPageTitle"
              type="text"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="ورود و ثبت‌نام"
            />
            <p class="text-xs text-gray-500 mt-1">عنوان نمایشی در صفحه ورود</p>
          </div>

          <div class="form-group">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              متن راهنما
            </label>
            <input
              v-model="loginSettings.helperText"
              type="text"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="لطفا شماره موبایل خود را وارد کنید"
            />
            <p class="text-xs text-gray-500 mt-1">متن راهنمای زیر فیلد شماره موبایل</p>
          </div>

          <div class="form-group">
            <label class="flex items-center">
              <input
                v-model="loginSettings.showBackButton"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
              />
              <span class="ml-2 text-sm font-medium text-gray-700">نمایش دکمه بازگشت</span>
            </label>
            <p class="text-xs text-gray-500 mt-1">نمایش دکمه بازگشت در صفحه ورود</p>
          </div>

          <div class="form-group">
            <label class="flex items-center">
              <input
                v-model="loginSettings.showPasswordLogin"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
              />
              <span class="ml-2 text-sm font-medium text-gray-700">نمایش گزینه ورود با رمز عبور</span>
            </label>
            <p class="text-xs text-gray-500 mt-1">نمایش لینک ورود با رمز عبور در صفحه OTP</p>
          </div>
        </div>
      </div>
    </div>

    <!-- دکمه‌های عملیات -->
    <div class="settings-actions mt-8 flex justify-end space-x-4 space-x-reverse">
      <button
        type="button"
        class="px-4 py-2 border border-gray-300 rounded-lg text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
        @click="resetSettings"
      >
        بازنشانی
      </button>
      <button
        :disabled="saving"
        class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50 disabled:cursor-not-allowed"
        @click="saveSettings"
      >
        <span v-if="saving">در حال ذخیره...</span>
        <span v-else>ذخیره تنظیمات</span>
      </button>
    </div>
  </div>
</template>

<script lang="ts">
declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>
</script>

<script setup lang="ts">
import { computed, onMounted, reactive, watch } from 'vue';
import { useAuth } from '~/composables/useAuth';

// احراز هویت
const { user, isAuthenticated } = useAuth();

// بررسی دسترسی admin
const hasAccess = computed(() => {
  if (!isAuthenticated.value) {
    return false;
  }

  const userRole = user.value?.role?.toLowerCase() || '';
  const adminRoles = ['admin', 'developer'];
  return adminRoles.includes(userRole);
});

// بررسی احراز هویت و دسترسی admin - نمایش 404 در صورت عدم دسترسی
const checkAuth = async (): Promise<void> => {
  if (!hasAccess.value) {
    await navigateTo('/404', { external: false });
  }
};

// بررسی احراز هویت در هنگام mount
onMounted(async () => {
  await checkAuth();
});

// بررسی احراز هویت هنگام تغییر وضعیت احراز هویت
watch([isAuthenticated, hasAccess], async () => {
  if (!hasAccess.value) {
    await checkAuth();
  }
});

interface LoginSettings {
  otpLength: number
  otpExpiryMinutes: number
  maxOtpAttempts: number
  otpResendDelaySeconds: number
  firstAttemptsWaitTime: number
  secondAttemptsWaitTime: number
  fifthAttemptWaitTime: number
  blockTimeMinutes: number
  maxLoginAttempts: number
  lockoutDurationMinutes: number
  sessionTimeoutMinutes: number
  mobileAuthEnabled: boolean
  loginPageTitle: string
  helperText: string
  showBackButton: boolean
  showPasswordLogin: boolean
}

const props = defineProps<{
  settings: LoginSettings
  saving: boolean
}>()

const emit = defineEmits(['save', 'reset'])

// تنظیمات پیش‌فرض
const defaultSettings = {
  // تنظیمات OTP
  otpLength: 5,
  otpExpiryMinutes: 5,
  maxOtpAttempts: 3,
  otpResendDelaySeconds: 60,
  
  // محدودیت‌های ارسال مجدد
  firstAttemptsWaitTime: 2,
  secondAttemptsWaitTime: 3,
  fifthAttemptWaitTime: 5,
  blockTimeMinutes: 60,
  
  // تنظیمات امنیتی
  maxLoginAttempts: 5,
  lockoutDurationMinutes: 15,
  sessionTimeoutMinutes: 1440,
  mobileAuthEnabled: true,
  
  // تنظیمات UI
  loginPageTitle: 'ورود و ثبت‌نام',
  helperText: 'لطفا شماره موبایل خود را وارد کنید',
  showBackButton: true,
  showPasswordLogin: true
}

// تنظیمات محلی
const loginSettings = reactive({
  ...defaultSettings,
  ...props.settings
})

// تماشای تغییرات props
watch(() => props.settings, (newSettings) => {
  Object.assign(loginSettings, defaultSettings, newSettings)
}, { deep: true })

// تابع ذخیره تنظیمات
const saveSettings = () => {
  emit('save', { ...loginSettings })
}

// تابع بازنشانی تنظیمات
const resetSettings = () => {
  Object.assign(loginSettings, defaultSettings)
  emit('reset')
}
</script>

<style scoped>
.login-page-settings {
  padding: 1.5rem;
  background-color: white;
  border-radius: 0.5rem;
  box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.1), 0 1px 2px 0 rgba(0, 0, 0, 0.06);
}

.settings-header {
  border-bottom: 1px solid #e5e7eb;
  padding-bottom: 1.5rem;
  margin-bottom: 1.5rem;
}

.settings-section {
  border-bottom: 1px solid #f3f4f6;
  padding-bottom: 1.5rem;
}

.settings-section:last-child {
  border-bottom: none;
  padding-bottom: 0;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
}

.settings-actions {
  border-top: 1px solid #e5e7eb;
  padding-top: 1.5rem;
}
</style>

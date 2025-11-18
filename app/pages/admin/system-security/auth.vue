<template>
  <div class="p-6 space-y-6" dir="rtl">
    <h1 class="text-2xl font-bold">تنظیمات احراز هویت</h1>

    <!-- حالت بارگذاری -->
    <div v-if="pending" class="text-center py-10">در حال بارگذاری...</div>

    <div v-else class="max-w-2xl space-y-8">
      <!-- تنظیمات عمومی JWT -->
      <section>
        <h2 class="text-xl font-semibold mb-4">JWT</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <FormInputNumber v-model="model.jwt_expiry_hours" label="مدت اعتبار توکن (ساعت)" />
          <FormInputNumber v-model="model.refresh_token_expiry_days" label="مدت اعتبار Refresh Token (روز)" />
          <FormInputText v-model="model.jwt_secret" label="کلید امضای JWT" type="password" />
        </div>
      </section>

      <!-- OTP -->
      <section>
        <h2 class="text-xl font-semibold mb-4">ورود با پیامک (OTP)</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <FormToggle v-model="model.mobile_auth_enabled" label="فعال‌سازی ورود با موبایل" />
          <FormInputNumber v-model="model.otp_length" label="تعداد ارقام OTP" />
          <FormInputNumber v-model="model.otp_expiry_minutes" label="زمان انقضا (دقیقه)" />
          <FormInputNumber v-model="model.max_otp_attempts" label="حداکثر تلاش مجاز" />
          <FormInputNumber v-model="model.otp_resend_cooldown" label="فاصله ارسال مجدد (ثانیه)" />
        </div>
      </section>

      <!-- محدودیت‌ها و امنیت -->
      <section>
        <h2 class="text-xl font-semibold mb-4">محدودیت‌ها و امنیت</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <FormInputNumber v-model="model.max_login_attempts" label="تلاش لاگین مجاز" />
          <FormInputNumber v-model="model.account_lockout_minutes" label="مدت قفل حساب (دقیقه)" />
          <FormInputNumber v-model="model.auto_block_failed_logins" label="تعداد ناموفق برای بلاک خودکار" />
          <FormInputNumber v-model="model.auto_block_duration_hours" label="مدت بلاک خودکار (ساعت)" />
          <FormToggle v-model="model.two_factor_enabled" label="فعال‌سازی ۲مرحله‌ای" />
          <FormToggle v-model="model.suspicious_activity_detection" label="تشخیص فعالیت مشکوک" />
        </div>
      </section>

      <!-- ذخیره -->
      <div class="pt-4">
        <button @click="save" class="px-6 py-2 bg-green-600 text-white rounded hover:bg-green-700">ذخیره</button>
        <span v-if="saving" class="ml-3">در حال ذخیره...</span>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string }) => void
declare const useFetch: <T = unknown>(url: string, options?: { key?: string; server?: boolean; default?: () => T; swr?: boolean; query?: Record<string, string | number>; method?: string; body?: unknown; credentials?: string }) => { data: { value: T }; pending: { value: boolean }; error: { value: Error | null }; refresh: () => Promise<void> }
</script>

<script setup lang="ts">
// کامپوننت‌های فرم کوچک (Toggle, InputNumber, InputText) فرض می‌شود قبلاً در پروژه وجود دارند
import { reactive, ref } from 'vue';

// meta
definePageMeta({ layout: 'admin-main' })

// استفاده از useAuth برای چک کردن پرمیژن‌ها
// Auth disabled
const pending = ref(true)
const saving = ref(false)

// مدل تنظیمات
const model = reactive({
  mobile_auth_enabled: true,
  otp_length: 5,
  otp_expiry_minutes: 5,
  max_otp_attempts: 3,
  otp_resend_cooldown: 60,
  jwt_expiry_hours: 24,
  refresh_token_expiry_days: 30,
  jwt_secret: '',
  max_concurrent_sessions: 5,
  username_auth_enabled: true,
  min_password_length: 8,
  max_login_attempts: 5,
  account_lockout_minutes: 15,
  password_expiry_days: 90,
  two_factor_enabled: false,
  suspicious_activity_detection: true,
  session_timeout_minutes: 30,
  login_history_retention_days: 365,
  auto_block_failed_logins: 10,
  auto_block_duration_hours: 24,
  default_user_role: 'user',
  email_verification_enabled: false,
  phone_verification_enabled: true,
})

// دریافت تنظیمات از API
const fetchSettings = async () => {
  const { data, pending: p } = await useFetch('/api/admin/settings/category/auth', { key: 'auth-settings', credentials: 'include' })
  pending.value = p.value
  if (data.value) Object.assign(model, data.value)
}

// ذخیره تنظیمات
const save = async () => {
  saving.value = true
  const { error } = await useFetch('/api/admin/settings/category/auth', {
    method: 'PUT',
    body: model,
    credentials: 'include',
  })
  saving.value = false
  if (error.value) alert('خطا در ذخیره تنظیمات')
}

fetchSettings()
</script>

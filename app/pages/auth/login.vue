<template>
  <div class="min-h-screen relative overflow-hidden font-iransans">
    <!-- پس‌زمینه گرادیان زیبا -->
    <div class="absolute inset-0 bg-gradient-to-br from-blue-50 via-indigo-50 to-purple-50"></div>
    
    <!-- دایره‌های تزئینی -->
    <div class="absolute top-20 left-10 w-72 h-72 bg-gradient-to-r from-blue-400/20 to-purple-400/20 rounded-full blur-3xl animate-pulse"></div>
    <div class="absolute bottom-20 right-10 w-96 h-96 bg-gradient-to-r from-indigo-400/20 to-pink-400/20 rounded-full blur-3xl animate-pulse delay-1000"></div>
    
    <!-- محتوای اصلی -->
    <div class="relative z-10 flex flex-col justify-start min-h-screen pt-20 px-4 sm:px-6 lg:px-8">
      <div class="sm:mx-auto sm:w-full sm:max-w-md">
        
        <!-- لوگو و عنوان -->
        <div class="text-center mb-8">
          <h2 class="text-2xl font-bold bg-gradient-to-r from-gray-900 to-gray-600 bg-clip-text text-transparent font-iransans">
            ورود و ثبت نام
          </h2>
        </div>
      </div>

      <div class="sm:mx-auto sm:w-full sm:max-w-md">
        <div class="bg-white/80 backdrop-blur-sm rounded-3xl shadow-2xl border border-white/20 p-8">
          <!-- فرم ورود ساده -->
          <div class="space-y-6">
            <!-- مرحله اول: وارد کردن شماره موبایل -->
            <div v-if="isClient && !otpSent">
              <div class="space-y-2">
                <div class="flex justify-end items-center mb-16">
                  <button
                    class="inline-flex items-center space-x-2 rtl:space-x-reverse text-pink-400 hover:text-pink-500 transition-colors duration-200 bg-pink-50 hover:bg-pink-100 px-4 py-2 rounded-lg"
                    @click="goBack"
                  >
                    <span class="text-sm">بازگشت</span>
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
                    </svg>
                  </button>
                </div>
                
                <p class="text-sm text-gray-600 text-center mb-4 font-medium">
                  {{ isDevUser ? 'لطفا نام کاربری خود را وارد کنید' : 'لطفا شماره موبایل خود را وارد کنید' }}
                </p>
                <div class="relative">
                  <div class="absolute inset-y-0 right-0 pr-3 flex items-center pointer-events-none">
                    <svg class="h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z"></path>
                    </svg>
                  </div>
                  <input
                    v-model="emailOrPhone"
                    type="text"
                    :placeholder="isDevUser ? 'نام کاربری خود را وارد کنید' : 'شماره موبایل خود را وارد کنید'"
                    required
                    :class="[
                      'w-full pr-10 pl-4 py-4 rounded-2xl text-sm focus:ring-2 focus:border-transparent transition-all duration-200 bg-white/50 backdrop-blur-sm text-right border-2',
                      error ? 'border-red-300 focus:ring-red-500' : 'border-blue-300 focus:ring-blue-500'
                    ]"
                  />
                </div>
                <div v-if="error" class="text-red-500 text-sm mt-1 text-right">{{ error }}</div>
              </div>

              <button
                :disabled="loading || !isPhoneValid"
                :class="[
                  isPhoneValid ? 'bg-green-400 hover:bg-green-500' : 'bg-blue-400 hover:bg-blue-500',
                  'w-full text-white py-4 rounded-2xl transition-all duration-300 font-semibold shadow-lg hover:shadow-xl transform hover:scale-[1.02] mt-12',
                  loading ? 'bg-gray-400 transform-none' : ''
                ]"
                @click="sendOTP"
              >
                <div class="flex items-center justify-center space-x-2 rtl:space-x-reverse">
                  <svg v-if="loading" class="animate-spin h-5 w-5" fill="none" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                  </svg>
                  <span v-if="loading">در حال ارسال...</span>
                  <span v-else>ورود</span>
                </div>
              </button>
            </div>

            <!-- مرحله دوم: وارد کردن کد تایید یا پسورد -->
            <div v-if="isClient && otpSent" class="space-y-6">
              <div class="text-center">
                <h3 class="text-2xl font-bold text-gray-900 mb-6">
                  {{ showPasswordField ? 'رمز عبور را وارد کنید' : 'کد تایید را وارد کنید' }}
                </h3>
                <p class="text-sm text-gray-600">
                  <span v-if="loading">در حال ارسال کد تایید...</span>
                  <span v-else-if="showPasswordField">برای کاربر {{ emailOrPhone }} رمز عبور وارد کنید</span>
                  <span v-else>کد تایید برای شماره {{ emailOrPhone }} پیامک شد</span>
                </p>
              </div>
              
              <div class="space-y-2">
                <!-- باکس پسورد برای کاربر dev -->
                <input
                  v-if="showPasswordField"
                  v-model="password"
                  type="password"
                  placeholder="رمز عبور را وارد کنید"
                  class="w-full px-4 py-4 border border-blue-300 rounded-2xl text-sm focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-200 bg-white/50 backdrop-blur-sm text-right"
                />
                <!-- باکس کد تایید برای کاربران عادی -->
                <input
                  v-else
                  v-model="otpCode"
                  type="text"
                  placeholder="کد تایید را وارد کنید"
                  maxlength="5"
                  class="w-full px-4 py-4 border border-blue-300 rounded-2xl text-sm focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-200 bg-white/50 backdrop-blur-sm text-center text-2xl font-mono tracking-widest"
                />
              </div>

              <div class="text-center mt-6">
                <a href="#" class="inline-flex items-center space-x-2 rtl:space-x-reverse text-blue-600 hover:text-blue-700 text-sm">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"></path>
                  </svg>
                  <span>ورود با رمز عبور</span>
                </a>
              </div>

              <div class="text-center text-sm text-gray-600">
                <span v-if="isBlocked" class="text-red-600 font-medium">
                  {{ resendStatusText }}
                </span>
                <span v-else-if="resendTimer > 0" class="text-gray-600">
                  {{ resendStatusText }}
                </span>
                <span v-else>
                  <button 
                    :disabled="loading || !isPhoneValid" 
                    class="text-blue-600 hover:text-blue-700 font-medium"
                    @click="sendOTP"
                  >
                    {{ resendStatusText }}
                  </button>
                </span>
              </div>

              <!-- دکمه ورود با پسورد برای کاربر dev -->
              <button
                v-if="showPasswordField"
                :disabled="loading || !isPasswordValid"
                :class="[
                  isPasswordValid ? 'bg-green-400 hover:bg-green-500' : 'bg-blue-400 hover:bg-blue-500',
                  'w-full text-white py-4 rounded-2xl transition-all duration-300 font-semibold shadow-lg hover:shadow-xl transform hover:scale-[1.02] disabled:transform-none'
                ]"
                @click="loginWithPasswordHandler"
              >
                <div class="flex items-center justify-center space-x-2 rtl:space-x-reverse">
                  <svg v-if="loading" class="animate-spin h-5 w-5" fill="none" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                  </svg>
                  <span v-if="loading">در حال ورود...</span>
                  <span v-else>ورود</span>
                </div>
              </button>
              
              <!-- دکمه تایید کد OTP برای کاربران عادی -->
              <button
                v-else
                :disabled="otpLoading || !isOtpValid"
                :class="[
                  isOtpValid ? 'bg-green-400 hover:bg-green-500' : 'bg-blue-400 hover:bg-blue-500',
                  'w-full text-white py-4 rounded-2xl transition-all duration-300 font-semibold shadow-lg hover:shadow-xl transform hover:scale-[1.02] disabled:transform-none'
                ]"
                @click="verifyOTPHandler"
              >
                <div class="flex items-center justify-center space-x-2 rtl:space-x-reverse">
                  <svg v-if="otpLoading" class="animate-spin h-5 w-5" fill="none" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                  </svg>
                  <span v-if="otpLoading">در حال تایید...</span>
                  <span v-else>تایید</span>
                </div>
              </button>
            </div>
          </div>

          <!-- Footer -->
          <div class="mt-8 text-center">
            <p class="text-xs text-gray-500">
              ورود شما به معنای پذیرش شرایط استفاده و قوانین حریم خصوصی است
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string }) => void
declare const navigateTo: (to: string) => Promise<void>
</script>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue';
import { useOTP } from '~/composables/useOTP';
import type { User } from '~/types/user';
import { validationUtils } from '~/utils/validation';

// استفاده از layout ساده برای صفحه احراز هویت (بدون Header/Footer)
definePageMeta({
  layout: 'auth'
});

const { sendOTP: sendOTPRequest, loading: otpLoading, error: otpError } = useOTP();

// متغیرهای فرم
const emailOrPhone = ref('');
const password = ref('');

// متغیرهای state
const loading = ref(false);
const error = ref('');

// متغیرهای OTP
const otpSent = ref(false);
const otpCode = ref('');
const resendTimer = ref(0);
const resendAttempts = ref(0);
const isBlocked = ref(false);
let otpTimerInterval: NodeJS.Timeout | null = null;

// متغیرهای مربوط به کاربر dev
const isDevUser = ref(false);
const showPasswordField = ref(false);

// متغیر برای تشخیص کلاینت
const isClient = ref(false);

onMounted(() => {
  isClient.value = true;
});

// computed برای بررسی وضعیت دکمه
const isPhoneValid = computed(() => {
  if (emailOrPhone.value === 'dev') {
    return true;
  }
  return emailOrPhone.value && validationUtils.isValidPhoneNumber(emailOrPhone.value);
});

// computed برای بررسی وضعیت دکمه ورود با پسورد
const isPasswordValid = computed(() => {
  return password.value && password.value.length >= 4;
});

// computed برای بررسی وضعیت کد OTP
const isOtpValid = computed(() => {
  return otpCode.value && otpCode.value.length === 5;
});

// computed برای نمایش وضعیت تلاش‌ها
const resendStatusText = computed(() => {
  if (isBlocked.value) {
    return 'حساب شما به مدت 1 ساعت بلاک شده است';
  }
  if (resendTimer.value > 0) {
    const minutes = Math.floor(resendTimer.value / 60);
    const seconds = resendTimer.value % 60;
    return `${minutes}:${String(seconds).padStart(2, '0')} مانده تا دریافت مجدد کد`;
  }
  return 'ارسال مجدد کد';
});


const sendOTP = async () => {
  error.value = '';
  
  if (!emailOrPhone.value) {
    error.value = 'لطفاً شماره موبایل را وارد کنید';
    return;
  }

  // بررسی کاربر dev
  if (emailOrPhone.value === 'dev') {
    isDevUser.value = true;
    showPasswordField.value = true;
    otpSent.value = true;
    return;
  }

  if (!isPhoneValid.value) {
    error.value = 'شماره موبایل باید 11 رقم و با 09 شروع شود (مثال: 09123456789)';
    return;
  }

  // اطمینان از اینکه شماره دقیقاً 11 رقم است
  if (emailOrPhone.value.length !== 11) {
    error.value = `شماره موبایل باید دقیقاً 11 رقم باشد (${emailOrPhone.value.length} رقم وارد شده)`;
    return;
  }

  // بررسی محدودیت‌ها
  if (resendTimer.value > 0) {
    const minutes = Math.floor(resendTimer.value / 60);
    const seconds = resendTimer.value % 60;
    error.value = `لطفاً ${minutes}:${String(seconds).padStart(2, '0')} دقیقه دیگر تلاش کنید`;
    return;
  }

  if (isBlocked.value) {
    error.value = 'به دلیل تلاش‌های متعدد، حساب شما به مدت 1 ساعت مسدود شده است';
    return;
  }

  loading.value = true;
  otpSent.value = true;
  startResendTimer();

  try {
    const success = await sendOTPRequest(emailOrPhone.value);
    
    if (!success) {
      // بازگشت به مرحله قبل در صورت خطا
      otpSent.value = false;
      resendAttempts.value--;
      error.value = otpError.value || 'خطا در ارسال کد تایید';
    }
  } catch (err: unknown) {
    const e = err as { data?: { message?: string }, message?: string };
    otpSent.value = false;
    resendAttempts.value--;
    error.value = e?.data?.message || e?.message || 'خطای نامشخصی رخ داده است';
  } finally {
    loading.value = false;
  }
}

const loginWithPasswordHandler = async () => {
  if (!emailOrPhone.value || !password.value) {
    error.value = 'لطفاً نام کاربری و رمز عبور را وارد کنید';
    return;
  }

  loading.value = true;
  error.value = '';

  try {
    const response = await $fetch<{ success: boolean; message?: string; user?: User }>('/api/auth/login-password', {
      method: 'POST',
      body: { 
        identifier: emailOrPhone.value,
        password: password.value 
      },
    });

    // Go backend token ها را در cookie ذخیره می‌کند و از response حذف می‌کند
    // پس بجای چک کردن token، success را چک می‌کنیم
    if (response?.success) {
      // Redirect بر اساس role کاربر
      const user = response?.user;
      const isAdmin = user?.role_id === 100 || 
                     user?.username === 'dev' || 
                     ['admin', 'developer', 'super_admin', 'manager', 'operator'].includes(user?.user_role?.name?.toLowerCase());
      
      if (isAdmin) {
        // Admin
        navigateTo('/admin');
      } else {
        // User عادی - redirect به پنل کاربری
        navigateTo('/account');
      }
    } else {
      error.value = response?.message || 'خطا در ورود';
    }
  } catch (err: unknown) {
    const e = err as { data?: { error?: string; message?: string }, message?: string };
    error.value = e?.data?.error || e?.data?.message || e?.message || 'نام کاربری یا رمز عبور اشتباه است';
  } finally {
    loading.value = false;
  }
}


const verifyOTPHandler = async () => {
  if (!emailOrPhone.value || !otpCode.value) {
    error.value = 'لطفاً شماره موبایل و کد تایید را وارد کنید';
    return;
  }

  error.value = '';

  try {
    const response = await $fetch<{ success: boolean; message?: string; user?: User }>('/api/auth/verify-otp', {
      method: 'POST',
      body: { 
        mobile: emailOrPhone.value, 
        code: otpCode.value 
      },
    });

    // Go backend token ها را در cookie ذخیره می‌کند
    if (response?.success) {
      // پاکسازی محدودیت‌های OTP
      resendTimer.value = 0;
      resendAttempts.value = 0;
      isBlocked.value = false;
      if (otpTimerInterval) {
        clearInterval(otpTimerInterval);
        otpTimerInterval = null;
      }
      
      // Redirect بر اساس role کاربر
      const user = response?.user;
      const isAdmin = user?.role_id === 100 || 
                     user?.username === 'dev' || 
                     ['admin', 'developer', 'super_admin', 'manager', 'operator'].includes(user?.user_role?.name?.toLowerCase());
      
      if (isAdmin) {
        // Admin
        navigateTo('/admin');
      } else {
        // User عادی - redirect به پنل کاربری
        navigateTo('/account');
      }
    } else {
      error.value = response?.message || 'کد تایید اشتباه است';
    }
  } catch (err: unknown) {
    const e = err as { data?: { message?: string }, message?: string };
    error.value = e?.data?.message || e?.message || 'خطای نامشخصی رخ داده است';
  }
}

/**
 * ریست کردن وضعیت OTP
 */
const resetOTPState = () => {
  otpSent.value = false;
  otpCode.value = '';
  error.value = '';
  isDevUser.value = false;
  showPasswordField.value = false;
  password.value = '';
}

/**
 * بازگشت به صفحه قبلی
 */
const goBack = () => {
  resetOTPState();
  
  const previousPath = sessionStorage.getItem('auth_redirect_path');
  if (previousPath && previousPath !== '/auth/login') {
    navigateTo(previousPath);
  } else {
    navigateTo('/');
  }
}

/**
 * شروع تایمر ارسال مجدد با الگوریتم Exponential Backoff
 */
const startResendTimer = () => {
  resendAttempts.value++;
  
  let waitTime = 0;
  if (resendAttempts.value <= 2) {
    waitTime = 120; // 2 دقیقه
  } else if (resendAttempts.value <= 4) {
    waitTime = 180; // 3 دقیقه
  } else if (resendAttempts.value === 5) {
    waitTime = 300; // 5 دقیقه
  } else {
    // بلاک کردن برای 1 ساعت
    isBlocked.value = true;
    waitTime = 3600;
  }
  
  resendTimer.value = waitTime;
  
  if (otpTimerInterval) {
    clearInterval(otpTimerInterval);
  }
  
  otpTimerInterval = setInterval(() => {
    resendTimer.value--;
    
    if (resendTimer.value <= 0) {
      clearInterval(otpTimerInterval!);
      otpTimerInterval = null;
      
      if (isBlocked.value && resendAttempts.value >= 6) {
        isBlocked.value = false;
        resendAttempts.value = 0;
      }
    }
  }, 1000);
}

// پاک کردن خطا هنگام تایپ
watch(emailOrPhone, () => {
  if (error.value) error.value = '';
})

watch(otpCode, () => {
  if (error.value) error.value = '';
})

watch(password, () => {
  if (error.value) error.value = '';
})

// ریست کردن state هنگام تغییر otpSent
watch(otpSent, (newValue) => {
  if (!newValue) {
    resetOTPState();
  }
})
</script>
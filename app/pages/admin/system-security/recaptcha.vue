<template>
  <div class="p-6" dir="rtl">
    <!-- عنوان صفحه -->
    <div class="mb-6 bg-white rounded-lg shadow-md border border-gray-200 p-6">
      <h1 class="text-2xl font-bold text-gray-900 mb-2">ریکپچا و Turnstile</h1>
      <p class="text-gray-600">مدیریت و تنظیمات reCAPTCHA و Cloudflare Turnstile برای محافظت از فرم‌ها</p>
    </div>

    <!-- آمار کلی -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
      <TemplateCard
        title="وضعیت reCAPTCHA"
        :value="recaptchaStatus"
        variant="blue"
      />
      <TemplateCard
        title="درخواست‌های موفق"
        :value="successfulRequests"
        variant="green"
      />
      <TemplateCard
        title="درخواست‌های مسدود شده"
        :value="blockedRequests"
        variant="red"
      />
      <TemplateCard
        title="نرخ موفقیت"
        :value="`${successRate}%`"
        variant="yellow"
      />
    </div>

    <!-- کنترل‌های reCAPTCHA -->
    <div class="bg-white rounded-2xl shadow-lg border border-gray-200 mb-8 p-6">
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6 items-center">
        <!-- وضعیت reCAPTCHA -->
        <div>
          <div class="text-xs text-gray-500 mb-2 font-medium">وضعیت reCAPTCHA</div>
          <div class="flex gap-2">
            <button :class="recaptchaEnabled ? 'bg-green-500 text-white' : 'bg-gray-200 text-gray-700'" class="px-4 py-1 rounded-full text-sm font-bold transition-colors" @click="toggleRecaptcha(true)">فعال</button>
            <button :class="!recaptchaEnabled ? 'bg-gray-400 text-white' : 'bg-gray-200 text-gray-700'" class="px-4 py-1 rounded-full text-sm font-bold transition-colors" @click="toggleRecaptcha(false)">غیرفعال</button>
          </div>
        </div>
        <!-- نوع reCAPTCHA -->
        <div>
          <div class="text-xs text-gray-500 mb-2 font-medium">نوع reCAPTCHA</div>
          <select v-model="recaptchaType" class="w-full px-3 py-2 border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-400 text-sm bg-gray-50" @change="changeRecaptchaType">
            <option value="v2">reCAPTCHA v2</option>
            <option value="v3">reCAPTCHA v3</option>
            <option value="invisible">reCAPTCHA Invisible</option>
            <option value="cloudflare">Cloudflare Turnstile</option>
          </select>
        </div>
        <!-- عملیات سریع -->
        <div>
          <div class="flex gap-6 justify-end">
            <TemplateButton 
              bg-gradient="bg-gradient-to-r from-cyan-400 to-blue-600" 
              text-color="text-white" 
              border-color="border border-blue-500" 
              hover-class="hover:from-cyan-500 hover:to-blue-700" 
              size="medium" 
              @click="testRecaptcha"
            >
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
              تست reCAPTCHA
            </TemplateButton>
            <TemplateButton 
              bg-gradient="bg-gradient-to-r from-emerald-400 to-green-600" 
              text-color="text-white" 
              border-color="border border-green-500" 
              hover-class="hover:from-emerald-500 hover:to-green-700" 
              size="medium" 
              @click="showSettingsModal = true"
            >
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"></path>
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
              </svg>
              تنظیمات
            </TemplateButton>
          </div>
        </div>
      </div>
    </div>

    <!-- تست Turnstile -->
    <div class="bg-white p-6 rounded-lg shadow-md border border-gray-200 mb-8">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">تست Cloudflare Turnstile</h3>
      <div class="bg-gradient-to-r from-blue-50 to-indigo-50 p-6 rounded-xl border border-blue-100">
        <div class="flex items-center mb-4">
          <svg class="w-6 h-6 text-blue-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"></path>
          </svg>
          <span class="text-blue-900 font-medium">نمونه فرم محافظت شده با Turnstile</span>
        </div>
        
        <form class="space-y-4" @submit.prevent="testTurnstileForm">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نام</label>
            <input v-model="testForm.name" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="نام خود را وارد کنید">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">ایمیل</label>
            <input v-model="testForm.email" type="email" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="ایمیل خود را وارد کنید">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">پیام</label>
            <textarea v-model="testForm.message" rows="3" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="پیام خود را وارد کنید"></textarea>
          </div>
          
          <!-- Turnstile Component -->
          <div class="flex justify-center">
            <TurnstileWidget 
              :site-key="turnstileSiteKey"
              :theme="turnstileTheme"
              :size="turnstileSize"
              @token="onTurnstileToken"
              @expired="onTurnstileExpired"
              @error="onTurnstileError"
            />
          </div>
          
          <div class="flex justify-end">
            <TemplateButton 
              type="submit"
              :disabled="!turnstileToken"
              bg-gradient="bg-gradient-to-r from-blue-500 to-indigo-600" 
              text-color="text-white" 
              border-color="border border-blue-600" 
              hover-class="hover:from-blue-600 hover:to-indigo-700" 
              size="medium"
            >
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8"></path>
              </svg>
              ارسال فرم
            </TemplateButton>
          </div>
        </form>
        
        <div v-if="turnstileResult" class="mt-4 p-3 rounded-lg" :class="turnstileResult.success ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'">
          <div class="flex items-center">
            <svg v-if="turnstileResult.success" class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
            <svg v-else class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
            <span class="font-medium">{{ turnstileResult.message }}</span>
          </div>
        </div>
      </div>
    </div>

    <!-- نمودار عملکرد -->
    <div class="bg-white p-6 rounded-lg shadow-md border border-gray-200 mb-8">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">نمودار عملکرد reCAPTCHA</h3>
      <div class="h-64 bg-gray-50 rounded-lg flex items-center justify-center">
        <p class="text-gray-500">نمودار عملکرد reCAPTCHA اینجا نمایش داده می‌شود</p>
      </div>
    </div>

    <!-- تنظیمات فرم‌ها -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 mb-8">
      <div class="px-6 py-4 border-b border-gray-200 flex justify-between items-center">
        <h3 class="text-lg font-semibold text-gray-900">تنظیمات فرم‌ها</h3>
        <TemplateButton 
          bg-gradient="bg-gradient-to-r from-emerald-400 to-green-600" 
          text-color="text-white" 
          border-color="border border-green-500" 
          hover-class="hover:from-emerald-500 hover:to-green-700" 
          size="medium" 
          @click="showAddFormModal = true"
        >
          <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
          </svg>
          افزودن فرم جدید
        </TemplateButton>
      </div>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نام فرم</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع reCAPTCHA</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">آستانه امتیاز</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">آخرین تست</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="form in recaptchaForms" :key="form.id">
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{{ form.name }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ form.type }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ form.threshold }}</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="form.enabled ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ form.enabled ? 'فعال' : 'غیرفعال' }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ form.lastTest }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-2" @click="toggleForm(form.id)">
                  {{ form.enabled ? 'غیرفعال' : 'فعال' }}
                </button>
                <button class="text-green-600 hover:text-green-900 ml-2" @click="editForm(form.id)">ویرایش</button>
                <button class="text-purple-600 hover:text-purple-900 ml-2" @click="testFormById(form.id)">تست</button>
                <button class="text-red-600 hover:text-red-900" @click="deleteForm(form.id)">حذف</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- لاگ‌های reCAPTCHA -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-semibold text-gray-900">لاگ‌های reCAPTCHA</h3>
      </div>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">زمان</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">IP کاربر</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">فرم</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">امتیاز</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نتیجه</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">جزئیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="log in recaptchaLogs" :key="log.id">
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ log.timestamp }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ log.userIP }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ log.formName }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ log.score }}</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="log.success ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ log.success ? 'موفق' : 'ناموفق' }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ log.details }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- مودال تنظیمات -->
    <div v-if="showSettingsModal" class="fixed top-0 left-0 right-0 bottom-0 flex items-center justify-center z-50">
      <!-- Backdrop نامرئی برای بستن مودال -->
      <div class="absolute inset-0" @click="showSettingsModal = false"></div>
      <!-- مودال -->
      <div class="relative bg-gradient-to-br from-white to-gray-50 rounded-3xl shadow-2xl max-w-4xl w-full mx-4 p-6 transform transition-all duration-300 scale-100 border-2 border-purple-500 max-h-[90vh] overflow-y-auto" @click.stop>
        <!-- دکمه بستن -->
        <button class="absolute top-6 left-6 text-gray-400 hover:text-gray-600 transition-all duration-200 hover:scale-110" @click="showSettingsModal = false">
          <svg class="w-7 h-7" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
        </button>
        <div class="text-center mb-6">
          <div class="w-12 h-12 bg-gradient-to-br from-blue-500 to-blue-600 rounded-2xl flex items-center justify-center mx-auto mb-3 shadow-lg">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"></path>
            </svg>
          </div>
          <h3 class="text-xl font-bold text-gray-900 mb-1">تنظیمات reCAPTCHA</h3>
          <p class="text-gray-600 text-sm">مدیریت کلیدها و تنظیمات امنیتی</p>
        </div>
        
        <div class="space-y-4">
          <!-- کلیدهای reCAPTCHA -->
          <div class="bg-gradient-to-r from-blue-50 to-indigo-50 p-6 rounded-2xl border border-blue-100">
            <h4 class="text-base font-semibold text-blue-900 mb-3 flex items-center">
              <svg class="w-4 h-4 ml-2 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 7a2 2 0 012 2m4 0a6 6 0 01-7.743 5.743L11 17H9v2H7v2H4a1 1 0 01-1-1v-2.586a1 1 0 01.293-.707l5.964-5.964A6 6 0 1121 9z"></path>
              </svg>
              کلیدهای reCAPTCHA
            </h4>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-blue-800 mb-1">کلید سایت reCAPTCHA</label>
                <input v-model="settings.recaptchaSiteKey" type="text" class="w-full px-3 py-2 border border-blue-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 bg-white shadow-sm">
              </div>
              
              <div>
                <label class="block text-sm font-medium text-blue-800 mb-1">کلید مخفی reCAPTCHA</label>
                <input v-model="settings.recaptchaSecretKey" type="password" class="w-full px-3 py-2 border border-blue-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 bg-white shadow-sm">
              </div>
            </div>
          </div>

          <!-- کلیدهای Turnstile -->
          <div class="bg-gradient-to-r from-green-50 to-emerald-50 p-6 rounded-2xl border border-green-100">
            <h4 class="text-base font-semibold text-green-900 mb-3 flex items-center">
              <svg class="w-4 h-4 ml-2 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"></path>
              </svg>
              کلیدهای Cloudflare Turnstile
            </h4>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-green-800 mb-1">کلید سایت Turnstile</label>
                <input v-model="settings.turnstileSiteKey" type="text" class="w-full px-3 py-2 border border-green-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-green-500 bg-white shadow-sm">
              </div>
              
              <div>
                <label class="block text-sm font-medium text-green-800 mb-1">کلید مخفی Turnstile</label>
                <input v-model="settings.turnstileSecretKey" type="password" class="w-full px-3 py-2 border border-green-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-green-500 bg-white shadow-sm">
              </div>
            </div>
          </div>

          <!-- تنظیمات عمومی -->
          <div class="bg-gradient-to-r from-green-50 to-emerald-50 p-6 rounded-2xl border border-green-100">
            <h4 class="text-base font-semibold text-green-900 mb-3 flex items-center">
              <svg class="w-4 h-4 ml-2 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"></path>
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
              </svg>
              تنظیمات عمومی
            </h4>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-green-800 mb-1">آستانه پیش‌فرض امتیاز</label>
                <input v-model="settings.defaultThreshold" type="number" min="0" max="1" step="0.1" class="w-full px-3 py-2 border border-green-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-green-500 bg-white shadow-sm">
              </div>
              
              <div>
                <label class="block text-sm font-medium text-green-800 mb-1">زمان انقضای توکن (ثانیه)</label>
                <input v-model="settings.tokenExpiry" type="number" min="60" max="3600" class="w-full px-3 py-2 border border-green-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-green-500 bg-white shadow-sm">
              </div>
            </div>
          </div>

          <!-- تنظیمات پیشرفته -->
          <div class="bg-gradient-to-r from-purple-50 to-pink-50 p-6 rounded-2xl border border-purple-100">
            <h4 class="text-base font-semibold text-purple-900 mb-3 flex items-center">
              <svg class="w-4 h-4 ml-2 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"></path>
              </svg>
              تنظیمات پیشرفته
            </h4>
            <div class="grid grid-cols-1 md:grid-cols-3 gap-3">
              <label class="flex items-center p-2 bg-white rounded-xl border border-purple-200 hover:bg-purple-50 transition-colors duration-200 cursor-pointer">
                <input v-model="settings.enableLogging" type="checkbox" class="ml-2 rounded border-purple-300 text-purple-600 focus:ring-purple-500 focus:ring-2">
                <span class="text-sm font-medium text-purple-800">فعال‌سازی لاگ‌گیری</span>
              </label>
              
              <label class="flex items-center p-2 bg-white rounded-xl border border-purple-200 hover:bg-purple-50 transition-colors duration-200 cursor-pointer">
                <input v-model="settings.enableAnalytics" type="checkbox" class="ml-2 rounded border-purple-300 text-purple-600 focus:ring-purple-500 focus:ring-2">
                <span class="text-sm font-medium text-purple-800">فعال‌سازی تحلیل‌گر</span>
              </label>
              
              <label class="flex items-center p-2 bg-white rounded-xl border border-purple-200 hover:bg-purple-50 transition-colors duration-200 cursor-pointer">
                <input v-model="settings.enableFallback" type="checkbox" class="ml-2 rounded border-purple-300 text-purple-600 focus:ring-purple-500 focus:ring-2">
                <span class="text-sm font-medium text-purple-800">فعال‌سازی حالت پشتیبان</span>
              </label>
            </div>
          </div>
        </div>

        <div class="flex justify-end space-x-3 space-x-reverse mt-6">
          <TemplateButton 
            bg-gradient="bg-gradient-to-r from-gray-400 to-gray-600" 
            text-color="text-white" 
            border-color="border border-gray-500" 
            hover-class="hover:from-gray-500 hover:to-gray-700" 
            size="medium" 
            @click="showSettingsModal = false"
          >
            انصراف
          </TemplateButton>
          <TemplateButton 
            bg-gradient="bg-gradient-to-r from-indigo-400 to-blue-600" 
            text-color="text-white" 
            border-color="border border-blue-500" 
            hover-class="hover:from-indigo-500 hover:to-blue-700" 
            size="medium" 
            @click="saveSettings"
          >
            ذخیره تنظیمات
          </TemplateButton>
        </div>
      </div>
    </div>

    <!-- مودال افزودن فرم -->
    <div v-if="showAddFormModal" class="fixed top-0 left-0 right-0 bottom-0 flex items-center justify-center z-50">
      <!-- Backdrop نامرئی برای بستن مودال -->
      <div class="absolute inset-0" @click="showAddFormModal = false"></div>
      <!-- مودال -->
      <div class="relative bg-gradient-to-br from-white to-gray-50 rounded-3xl shadow-2xl max-w-lg w-full mx-4 p-8 transform transition-all duration-300 scale-100 border-2 border-purple-500" @click.stop>
        <!-- دکمه بستن -->
        <button class="absolute top-6 left-6 text-gray-400 hover:text-gray-600 transition-all duration-200 hover:scale-110" @click="showAddFormModal = false">
          <svg class="w-7 h-7" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
        </button>
        <div class="text-center mb-8">
          <div class="w-16 h-16 bg-gradient-to-br from-green-500 to-green-600 rounded-2xl flex items-center justify-center mx-auto mb-4 shadow-lg">
            <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
            </svg>
          </div>
          <h3 class="text-2xl font-bold text-gray-900 mb-2">افزودن فرم جدید</h3>
          <p class="text-gray-600 text-sm">ایجاد فرم جدید برای محافظت با reCAPTCHA</p>
        </div>
        
        <div class="space-y-6">
          <div class="bg-gradient-to-r from-blue-50 to-indigo-50 p-6 rounded-2xl border border-blue-100">
            <label class="block text-sm font-medium text-blue-800 mb-2">نام فرم</label>
            <input v-model="newForm.name" type="text" class="w-full px-4 py-3 border border-blue-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 bg-white shadow-sm" placeholder="مثال: فرم تماس">
          </div>
          
          <div class="bg-gradient-to-r from-green-50 to-emerald-50 p-6 rounded-2xl border border-green-100">
            <label class="block text-sm font-medium text-green-800 mb-2">نوع reCAPTCHA</label>
            <select v-model="newForm.type" class="w-full px-4 py-3 border border-green-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-green-500 bg-white shadow-sm">
              <option value="v2">reCAPTCHA v2</option>
              <option value="v3">reCAPTCHA v3</option>
              <option value="invisible">reCAPTCHA Invisible</option>
              <option value="cloudflare">Cloudflare Turnstile</option>
            </select>
          </div>
          
          <div class="bg-gradient-to-r from-purple-50 to-pink-50 p-6 rounded-2xl border border-purple-100">
            <label class="block text-sm font-medium text-purple-800 mb-2">آستانه امتیاز</label>
            <input v-model="newForm.threshold" type="number" min="0" max="1" step="0.1" class="w-full px-4 py-3 border border-purple-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-purple-500 bg-white shadow-sm" placeholder="0.5">
          </div>
        </div>

        <div class="flex justify-end space-x-3 space-x-reverse mt-8">
          <TemplateButton 
            bg-gradient="bg-gradient-to-r from-gray-400 to-gray-600" 
            text-color="text-white" 
            border-color="border border-gray-500" 
            hover-class="hover:from-gray-500 hover:to-gray-700" 
            size="medium" 
            @click="showAddFormModal = false"
          >
            انصراف
          </TemplateButton>
          <TemplateButton 
            bg-gradient="bg-gradient-to-r from-emerald-400 to-green-600" 
            text-color="text-white" 
            border-color="border border-green-500" 
            hover-class="hover:from-emerald-500 hover:to-green-700" 
            size="medium" 
            @click="addForm"
          >
            افزودن فرم
          </TemplateButton>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
declare const $fetch: <T = unknown>(url: string, options?: { method?: string; body?: unknown }) => Promise<T>
</script>

<script setup lang="ts">
import TemplateButton from '@/components/common/TemplateButton.vue'
import TemplateCard from '@/components/common/TemplateCard.vue'
import TurnstileWidget from '@/components/common/TurnstileWidget.vue'
import { ref } from 'vue'

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
// Auth disabled
// تعریف interface ها
interface RecaptchaSettings {
  recaptchaSiteKey: string
  recaptchaSecretKey: string
  turnstileSiteKey: string
  turnstileSecretKey: string
  defaultThreshold: number
  tokenExpiry: number
  enableLogging: boolean
  enableAnalytics: boolean
  enableFallback: boolean
}

interface TestForm {
  name: string
  email?: string
  message?: string
  type: string
  threshold: number
}

// interface TestResult {
//   success: boolean
//   score?: number
//   message: string
// }

// interface RequestLog {
//   id: number
//   timestamp: string
//   userIP: string
//   formName: string
//   score: number
//   success: boolean
//   details: string
// }

interface Rule {
  id: number
  name: string
  type: string
  threshold: number
  enabled: boolean
  lastTest: string
}

// متغیرهای reactive
const recaptchaStatus = ref('فعال')
const successfulRequests = ref(15420)
const blockedRequests = ref(1234)
const successRate = ref(92.6)
const recaptchaEnabled = ref(true)
const recaptchaType = ref('v3')
const showSettingsModal = ref(false)
const showAddFormModal = ref(false)

// تنظیمات reCAPTCHA و Turnstile
const settings = ref<RecaptchaSettings>({
  recaptchaSiteKey: '6LcXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX',
  recaptchaSecretKey: '6LcXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX',
  turnstileSiteKey: '1x00000000000000000000AA',
  turnstileSecretKey: '1x0000000000000000000000000000000AA',
  defaultThreshold: 0.5,
  tokenExpiry: 300,
  enableLogging: true,
  enableAnalytics: true,
  enableFallback: false
})

// متغیرهای تست Turnstile
const turnstileToken = ref('')
const turnstileResult = ref(null)
const testForm = ref<TestForm>({
  name: '',
  email: '',
  message: '',
  type: 'v3',
  threshold: 0.5
})

// تنظیمات Turnstile
const turnstileSiteKey = ref('1x00000000000000000000AA')
const turnstileTheme = ref('light')
const turnstileSize = ref('normal')

// توابع Turnstile
const onTurnstileToken = (token: string) => {
  turnstileToken.value = token
  // console.log('Turnstile token received:', token)
}

const onTurnstileExpired = () => {
  turnstileToken.value = ''
  // console.log('Turnstile token expired')
}

const onTurnstileError = (error: string) => {
  turnstileToken.value = ''
  // console.log('Turnstile error:', error)
}

// فرم جدید
const newForm = ref<TestForm>({
  name: '',
  type: 'v3',
  threshold: 0.5
})

// داده‌های نمونه
const recaptchaForms = ref<Rule[]>([
  {
    id: 1,
    name: 'فرم تماس',
    type: 'reCAPTCHA v3',
    threshold: 0.5,
    enabled: true,
    lastTest: '2 ساعت پیش'
  },
  {
    id: 2,
    name: 'فرم ثبت‌نام',
    type: 'reCAPTCHA v2',
    threshold: 0.7,
    enabled: true,
    lastTest: '1 روز پیش'
  },
  {
    id: 3,
    name: 'فرم ورود',
    type: 'reCAPTCHA Invisible',
    threshold: 0.6,
    enabled: false,
    lastTest: '3 روز پیش'
  }
])

const recaptchaLogs = ref([
  {
    id: 1,
    timestamp: '2024-01-15 14:30:25',
    userIP: '192.168.1.100',
    formName: 'فرم تماس',
    score: 0.8,
    success: true,
    details: 'درخواست موفق'
  },
  {
    id: 2,
    timestamp: '2024-01-15 14:28:15',
    userIP: '192.168.1.101',
    formName: 'فرم ثبت‌نام',
    score: 0.2,
    success: false,
    details: 'امتیاز پایین'
  },
  {
    id: 3,
    timestamp: '2024-01-15 14:25:45',
    userIP: '192.168.1.102',
    formName: 'فرم ورود',
    score: 0.9,
    success: true,
    details: 'درخواست موفق'
  }
])

// توابع
const toggleRecaptcha = (enabled: boolean) => {
  recaptchaEnabled.value = enabled
  recaptchaStatus.value = enabled ? 'فعال' : 'غیرفعال'
}

const changeRecaptchaType = () => {
  // console.log('نوع reCAPTCHA تغییر کرد به:', recaptchaType.value)
}

const testRecaptcha = () => {
  // console.log('تست reCAPTCHA انجام شد')
}

const testTurnstileForm = async () => {
  if (!turnstileToken.value) {
    turnstileResult.value = {
      success: false,
      message: 'لطفاً ابتدا Turnstile را تکمیل کنید'
    }
    return
  }

  try {
    // ارسال درخواست به endpoint اعتبارسنجی Turnstile
    const response = await $fetch<{ success?: boolean }>('/api/turnstile-test', {
      method: 'POST',
      body: {
        token: turnstileToken.value
      }
    })

    if (response.success) {
      turnstileResult.value = {
        success: true,
        message: 'اعتبارسنجی Turnstile موفق بود! فرم ارسال شد.'
      }
      
      // اضافه کردن به لاگ‌ها
      recaptchaLogs.value.unshift({
        id: Date.now(),
        timestamp: new Date().toLocaleString('fa-IR'),
        userIP: '127.0.0.1',
        formName: 'تست Turnstile',
        score: 1.0,
        success: true,
        details: 'اعتبارسنجی Turnstile موفق'
      })
      
      // پاک کردن فرم
      testForm.value = { name: '', email: '', message: '', type: 'v3', threshold: 0.5 }
      turnstileToken.value = ''
    } else {
      turnstileResult.value = {
        success: false,
        message: 'اعتبارسنجی Turnstile ناموفق بود'
      }
    }
  } catch (error) {
    console.error('خطا در اعتبارسنجی Turnstile:', error)
    turnstileResult.value = {
      success: false,
      message: 'خطا در ارتباط با سرور'
    }
  }
}

const saveSettings = () => {
  // console.log('تنظیمات ذخیره شد:', settings.value)
  showSettingsModal.value = false
}

const addForm = () => {
  if (newForm.value.name.trim()) {
    recaptchaForms.value.push({
      id: Date.now(),
      name: newForm.value.name,
      type: newForm.value.type,
      threshold: newForm.value.threshold,
      enabled: true,
      lastTest: 'همین الان'
    })
    newForm.value = { name: '', type: 'v3', threshold: 0.5 }
    showAddFormModal.value = false
  }
}

const toggleForm = (id: number) => {
  const form = recaptchaForms.value.find(f => f.id === id)
  if (form) {
    form.enabled = !form.enabled
  }
}

const editForm = (id: number) => {
  // console.log('ویرایش فرم با ID:', id)
}

const testFormById = (id: number) => {
  // console.log('تست فرم با ID:', id)
}

const deleteForm = (id: number) => {
  recaptchaForms.value = recaptchaForms.value.filter(f => f.id !== id)
}
</script> 

<template>
  <div class="settings-layout" dir="rtl">
    <SettingsMenu 
      :menu="menu" 
      :active-key="activeKey" 
      :expanded="expanded"
      @update:active-key="activeKey = $event"
      @menu-click="handleMenuClick"
    />
    
    <div class="settings-content">
      <HomeSettings 
        v-if="activeKey === 'home'"
        :settings="settings"
        :saving="saving"
        @save="saveSettings"
        @reset="resetSettings"
        @select-image="selectImage"
        @add-phone="addPhone"
        @remove-phone="removePhone"
        @add-admin-phone="addAdminPhone"
        @remove-admin-phone="removeAdminPhone"
        @add-working-hour="addWorkingHour"
        @remove-working-hour="removeWorkingHour"
        @add-location="addLocation"
        @remove-location="removeLocation"
        @add-location-phone="addLocationPhone"
        @remove-location-phone="removeLocationPhone"
      />
      <GeneralSettings v-else-if="activeKey === 'general'" />
      <ProductSettings v-else-if="activeKey === 'product'" />
      <ShopSettings v-else-if="activeKey === 'shop'" />
      <InvoicePrintSettings 
        v-else-if="activeKey === 'invoice-print'"
        :invoice-settings="invoiceSettings"
        :print-settings="printSettings"
        :template-settings="templateSettings"
        :saving-invoice-print="savingInvoicePrint"
        @save="saveInvoicePrintSettings"
        @reset="resetInvoicePrintSettings"
      />
      <ApiSettings 
        v-else-if="activeKey === 'api-settings'"
        :open-a-i-settings="openAISettings || defaultOpenAI"
        :available-models="availableModels"
        :section-configs="sectionConfigs"
        :show-api-key="showApiKey"
        :fetching-usage="fetchingUsage"
        :testing-connection="testingConnection"
        :saving-api="savingApi"
        @save="saveApiSettings"
        @reset="resetApiSettings"
        @test-connection="testOpenAIConnection"
        @fetch-usage="fetchUsageData"
        @toggle-api-key="showApiKey = !showApiKey"
        @add-consuming-page="addConsumingPage"
        @remove-consuming-page="removeConsumingPage"
      />
      <AutomationSettings v-else-if="activeKey === 'automation'" />
      <IntegrationSettings v-else-if="activeKey === 'integration'" />
      <SocialMediaSettings 
        v-else-if="activeKey === 'social-media'"
        :settings="socialMediaSettings"
        :saving="savingSocialMedia"
        @save="saveSocialMediaSettings"
        @reset="resetSocialMediaSettings"
      />
      <AuthSettings 
        v-else-if="activeKey === 'auth'" 
        :auth-settings="authSettings" 
        :auth-active-tab="authActiveTab" 
        :show-jwt-secret="showJwtSecret" 
        :saving-auth="savingAuth" 
        @save="saveAuthSettings" 
        @generate-new-jwt-secret="generateNewJwtSecret" 
        @update:auth-active-tab="authActiveTab = $event" 
        @toggle-jwt-secret="showJwtSecret = !showJwtSecret" 
      />
      <UserSettings v-else-if="activeKey === 'user'" />
      <ReviewSettings v-else-if="activeKey === 'reviews'" :config="reviewSettings" @save="saveReviewSettings" />
      <LoginPageSettings 
        v-else-if="activeKey === 'login-page'" 
        :settings="loginPageSettings" 
        :saving="savingLoginPage" 
        @save="saveLoginPageSettings" 
        @reset="resetLoginPageSettings" 
      />
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, watch, onMounted } from 'vue'

// Import settings components
import SettingsMenu from './components/SettingsMenu.vue'
import HomeSettings from './components/HomeSettings.vue'
import GeneralSettings from './components/GeneralSettings.vue'
import ProductSettings from './components/ProductSettings.vue'
import ShopSettings from './components/ShopSettings.vue'
import InvoicePrintSettings from './components/InvoicePrintSettings.vue'
import ApiSettings from './components/ApiSettings.vue'
import AutomationSettings from './components/AutomationSettings.vue'
import IntegrationSettings from './components/IntegrationSettings.vue'
import SocialMediaSettings from './components/SocialMediaSettings.vue'
import AuthSettings from './components/AuthSettings.vue'
import UserSettings from './components/UserSettings.vue'
import ReviewSettings from './components/ReviewSettings.vue'
import LoginPageSettings from './components/LoginPageSettings.vue'

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

const menu = [
  { key: 'home', title: 'تنظیمات خانه', icon: 'i-heroicons-home', children: [] },
  { key: 'general', title: 'تنظیمات عمومی', icon: 'i-heroicons-cog-6-tooth', children: [] },
  { key: 'product', title: 'تنظیمات محصول', icon: 'i-heroicons-shopping-cart', children: [] },
  { key: 'shop', title: 'تنظیمات فروشگاه', icon: 'i-heroicons-store', children: [] },
  { key: 'invoice-print', title: 'پیکربندی فاکتور و چاپ', icon: 'i-heroicons-printer', children: [] },
  { key: 'api-settings', title: 'تنظیمات API', icon: 'i-heroicons-code-bracket', children: [] },
  { key: 'automation', title: 'تنظیمات اتوماسیون', icon: 'i-heroicons-bolt', children: [] },
  { key: 'integration', title: 'تنظیمات یکپارچه‌سازی', icon: 'i-heroicons-link', children: [] },
  { key: 'social-media', title: 'شبکه‌های اجتماعی', icon: 'i-heroicons-share', children: [] },
  { key: 'auth', title: 'تنظیمات احراز هویت', icon: 'i-heroicons-shield-check', children: [] },
  { key: 'login-page', title: 'تنظیمات صفحه ورود و ثبت‌نام', icon: 'i-heroicons-key', children: [] },
  { key: 'user', title: 'تنظیمات کاربری', icon: 'i-heroicons-user', children: [] },
  { key: 'reviews', title: 'تنظیمات نظرات', icon: 'i-heroicons-chat-bubble-bottom-center-text', children: [] },
]

const activeKey = ref('home')
const expanded = ref({})
const saving = ref(false)
const savingInvoicePrint = ref(false)
const savingApi = ref(false)
const savingAuth = ref(false)
const savingSocialMedia = ref(false)
const testingConnection = ref(false)
const showApiKey = ref(false)
const showJwtSecret = ref(false)
const fetchingUsage = ref(false)
const savingLoginPage = ref(false)

// تنظیمات فاکتور و چاپ
const invoiceSettings = reactive({
  companyName: 'شرکت من',
  taxNumber: '',
  companyAddress: '',
  companyPhone: ''
})

const printSettings = reactive({
  defaultPaperSize: 'a4',
  orientation: 'portrait',
  defaultFont: 'iransans',
  fontSize: '12',
  showLogo: true,
  showQRCode: true,
  showBarcode: true
})

const templateSettings = reactive({
  primaryColor: '#3B82F6',
  secondaryColor: '#1E40AF',
  showHeader: true,
  showFooter: true,
  showPageNumbers: true
})

// تنظیمات نظرات
const reviewSettings = reactive({
  allowedImageTypes: 'image/jpeg,image/png,image/webp',
  allowedVideoTypes: 'video/mp4,video/webm',
  maxFilesPerReview: 7,
  maxFileSizeMb: 50,
})

// شیء پیش‌فرض برای جلوگیری از خطای undefined هنگام رندر اولیه
const defaultOpenAI = {
  enabled: false,
  apiKey: '',
  apiUrl: 'https://api.openai.com/v1',
  rateLimit: 60,
  timeout: 30,
  maxDailyCost: 0,
  accountBalance: '0.00',
  monthlyUsage: '0.00',
  totalRequests: '0',
  todayRequests: '0',
  lastBalanceUpdate: '',
  consumingPages: [''],
  sectionModels: {}
}

const saveReviewSettings = async (payload) => {
  Object.assign(reviewSettings, payload)
  // TODO: اتصال به API تنظیمات بک‌اند (SettingService) در صورتی که کلیدها مشخص شد
}

// تنظیمات صفحه ورود و ثبت‌نام
const loginPageSettings = reactive({
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
})

// تنظیمات شبکه‌های اجتماعی
const socialMediaSettings = reactive({
  instagram: '',
  instagram_enabled: true,
  telegram: '',
  telegram_enabled: true,
  whatsapp: '',
  whatsapp_enabled: true,
  twitter: '',
  twitter_enabled: true,
  linkedin: '',
  linkedin_enabled: true,
  facebook: '',
  facebook_enabled: true,
  youtube: '',
  youtube_enabled: true,
  aparat: '',
  aparat_enabled: true,
  rubika: '',
  rubika_enabled: true,
  eitaa: '',
  eitaa_enabled: true,
  bale: '',
  bale_enabled: true,
  tiktok: '',
  tiktok_enabled: true,
  pinterest: '',
  pinterest_enabled: true,
  discord: '',
  discord_enabled: true,
  github: '',
  github_enabled: true,
  customLinks: []
})

// تب فعال احراز هویت
const authActiveTab = ref('login-register')

// تنظیمات احراز هویت
const authSettings = reactive({
  // تنظیمات سیستم یکپارچه
  unifiedAuthEnabled: true,
  autoRegisterEnabled: true,
  defaultAuthMethod: 'otp',

  // تنظیمات OTP
  mobileAuthEnabled: true,
  otpLength: 5,
  otpExpiryMinutes: 5,
  maxOtpAttempts: 3,
  otpResendCooldown: 60,
  otpPatternCode: '',

  // تنظیمات JWT
  jwtExpiryHours: 24,
  refreshTokenExpiryDays: 30,
  jwtSecret: '',
  maxConcurrentSessions: 5,

  // تنظیمات ورود با یوزرنیم
  usernameAuthEnabled: true,
  minPasswordLength: 8,
  maxLoginAttempts: 5,
  accountLockoutMinutes: 15,
  passwordExpiryDays: 90,

  // تنظیمات امنیتی
  twoFactorEnabled: false,
  suspiciousActivityDetection: true,
  sessionTimeoutMinutes: 30,
  loginHistoryRetentionDays: 365,
  autoBlockFailedLogins: 10,
  autoBlockDurationHours: 24,

  // تنظیمات نقش‌ها
  defaultUserRole: 'user',
  emailVerificationEnabled: false,
  phoneVerificationEnabled: true
})

// تنظیمات OpenAI
const openAISettings = reactive({
  enabled: false,
  apiKey: '',
  apiUrl: 'https://api.openai.com/v1',
  defaultModel: '',
  temperature: 0.7,
  rateLimit: 60,
  timeout: 30,
  maxDailyCost: 10.00,
  accountBalance: '0.00',
  monthlyUsage: '0.00',
  totalRequests: '0',
  todayRequests: '0',
  lastBalanceUpdate: '',
  consumingPages: [''],
  sectionModels: {},
  availableModels: []
})

// لیست مدل‌های موجود (جایگزین شده بر اساس لیست ارائه‌شده)
const availableModels = ref([
  { id: 'gpt-5', name: 'gpt-5', description: '', category: 'chat', max_tokens: 128000, is_active: true, cost_per_1k: 'input —' },
  { id: 'gpt-5-high', name: 'gpt-5-high', description: '', category: 'chat', max_tokens: 128000, is_active: true, cost_per_1k: 'input —' },
  { id: 'o3', name: 'o3', description: '', category: 'chat', max_tokens: 128000, is_active: true, cost_per_1k: 'input —' },
  { id: 'gpt-4.1', name: 'gpt-4.1', description: '', category: 'chat', max_tokens: 128000, is_active: true, cost_per_1k: 'input —' },
  { id: 'gpt-5-mini', name: 'gpt-5-mini', description: '', category: 'chat', max_tokens: 128000, is_active: true, cost_per_1k: 'input —' },
  { id: 'gpt-4o-mini', name: 'gpt-4o-mini (Vision)', description: 'مدل سبک با توانایی دیدن تصویر', category: 'vision', max_tokens: 128000, is_active: true, cost_per_1k: 'input —' },
  { id: 'gpt-5-nano', name: 'gpt-5-nano', description: '', category: 'chat', max_tokens: 128000, is_active: true, cost_per_1k: 'input —' },
  { id: 'gpt-4o-mini-search-preview', name: 'gpt-4o-mini-search-preview', description: '', category: 'chat', max_tokens: 128000, is_active: true, cost_per_1k: 'input —' },
  { id: 'gpt-4o-mini', name: 'gpt-4o-mini', description: '', category: 'chat', max_tokens: 128000, is_active: true, cost_per_1k: 'input —' },
  { id: 'dall-e-3', name: 'DALL·E 3', description: 'مدل تولید تصویر OpenAI', category: 'image', max_tokens: 0, is_active: true, cost_per_1k: 'input —' },
])

// تنظیمات بخش‌ها
const sectionConfigs = reactive({
  chatbot: {
    title: 'چت بات',
    description: 'مدل مورد استفاده برای چت بات‌های پشتیبانی',
    model: 'gpt-3.5-turbo',
    temperature: 0.7,
    maxTokens: 1000,
    isEnabled: true
  },
  content_generation: {
    title: 'تولید محتوا',
    description: 'مدل برای تولید محتوای متنی و مقالات',
    model: 'gpt-4',
    temperature: 0.8,
    maxTokens: 2000,
    isEnabled: true
  },
  seo_optimization: {
    title: 'بهینه‌سازی SEO',
    description: 'مدل برای بهینه‌سازی محتوا برای موتورهای جستجو',
    model: 'gpt-4-turbo',
    temperature: 0.6,
    maxTokens: 1500,
    isEnabled: true
  },
  product_description: {
    title: 'توضیحات محصول',
    description: 'مدل برای تولید توضیحات محصولات',
    model: 'gpt-3.5-turbo',
    temperature: 0.7,
    maxTokens: 800,
    isEnabled: true
  },
  customer_support: {
    title: 'پشتیبانی مشتری',
    description: 'مدل برای پاسخ‌دهی به سوالات مشتریان',
    model: 'gpt-4',
    temperature: 0.5,
    maxTokens: 1200,
    isEnabled: true
  },
  translation: {
    title: 'ترجمه',
    description: 'مدل برای ترجمه متون',
    model: 'gpt-3.5-turbo',
    temperature: 0.3,
    maxTokens: 1000,
    isEnabled: true
  },
  code_generation: {
    title: 'تولید کد',
    description: 'مدل برای تولید و دیباگ کد',
    model: 'gpt-4',
    temperature: 0.2,
    maxTokens: 2000,
    isEnabled: true
  },
  data_analysis: {
    title: 'تحلیل داده',
    description: 'مدل برای تحلیل و تفسیر داده‌ها',
    model: 'gpt-4-turbo',
    temperature: 0.4,
    maxTokens: 1500,
    isEnabled: true
  },
  image_seo: {
    title: 'سئو تصاویر',
    description: 'بهینه‌سازی alt، کپشن و پیشنهاد کلیدواژه برای تصاویر رسانه',
    model: 'gpt-4o-mini',
    temperature: 0.4,
    maxTokens: 600,
    isEnabled: true,
    link: '/admin/media-management/image-compression',
    linkText: 'بخش هوش مصنوعی تصاویر'
  }
})

function handleMenuClick(item) {
  if (item.children && item.children.length > 0) {
    // اگر آیتم فرزند دارد، فقط expand/collapse کن
    toggleGroup(item.key)
  } else {
    // اگر آیتم فرزند ندارد، activeKey را تغییر بده
    activeKey.value = item.key
  }
}

function toggleGroup(k) {
  expanded.value[k] = !expanded.value[k]
}

// توابع مدیریت تنظیمات
function selectImage(type) {
  // اینجا می‌توانید منطق انتخاب تصویر را اضافه کنید
}

function addAdminPhone() {
  settings.adminPhones.push('')
}

function removeAdminPhone(index) {
  if (settings.adminPhones.length > 1) {
    settings.adminPhones.splice(index, 1)
  }
}

function addPhone() {
  settings.phones.push('')
}

function removePhone(index) {
  if (settings.phones.length > 1) {
    settings.phones.splice(index, 1)
  }
}

function addWorkingHour() {
  settings.workingHours.push('')
}

function removeWorkingHour(index) {
  if (settings.workingHours.length > 1) {
    settings.workingHours.splice(index, 1)
  }
}

function addLocation() {
  settings.locations.push({
    id: Date.now(),
    title: '',
    address: '',
    phones: ['']
  })
}

function removeLocation(index) {
  if (settings.locations.length > 1) {
    settings.locations.splice(index, 1)
  }
}

function addLocationPhone(locationIndex) {
  settings.locations[locationIndex].phones.push('')
}

function removeLocationPhone(locationIndex, phoneIndex) {
  if (settings.locations[locationIndex].phones.length > 1) {
    settings.locations[locationIndex].phones.splice(phoneIndex, 1)
  }
}

// توابع مدیریت تنظیمات API
function addConsumingPage() {
  openAISettings.consumingPages.push('')
}

function removeConsumingPage(index) {
  if (openAISettings.consumingPages.length > 1) {
    openAISettings.consumingPages.splice(index, 1)
  }
}

const saveApiSettings = async () => {
  try {
    savingApi.value = true
    
    // شروع ذخیره تنظیمات API
    
    // بروزرسانی تنظیمات بخش‌ها
    updateSectionModels()
    
    // آماده‌سازی داده‌ها برای ارسال
    const apiSettingsData = {
      openai: {
        enabled: openAISettings.enabled,
        api_key: openAISettings.apiKey,
        api_url: openAISettings.apiUrl,
        rate_limit: parseInt(openAISettings.rateLimit),
        timeout: parseInt(openAISettings.timeout),
        max_daily_cost: parseFloat(openAISettings.maxDailyCost),
        consuming_pages: openAISettings.consumingPages.filter(page => page.trim() !== ''),
        section_models: openAISettings.sectionModels,
        available_models: availableModels.value
      }
    }
    
    // داده‌های ارسالی
    
    // ارسال درخواست به API
    const response = await $fetch('/api/admin/api-settings', {
      method: 'PUT',
      body: apiSettingsData,
      credentials: 'include'
    })
    
    // پاسخ دریافتی
    
    // نمایش پیام موفقیت
    showSuccessMessage('تنظیمات API با موفقیت ذخیره شد')
    
  } catch (error) {
    // خطا در ذخیره تنظیمات API
    
    // نمایش خطای دقیق به کاربر
    let errorMessage = 'خطا در ذخیره تنظیمات API'
    
    if (error.data) {
      if (error.data.message) {
        errorMessage = error.data.message
      } else if (error.data.error) {
        errorMessage = error.data.error
      } else if (typeof error.data === 'string') {
        errorMessage = error.data
      }
    } else if (error.statusMessage) {
      errorMessage = error.statusMessage
    }
    
    showErrorMessage(errorMessage)
  } finally {
    savingApi.value = false
  }
}

const testOpenAIConnection = async () => {
  try {
    testingConnection.value = true
    
    // شروع تست اتصال OpenAI
    
    const response = await $fetch('/api/admin/api-settings/test-openai', {
      method: 'POST',
      body: {
        api_key: openAISettings.apiKey,
        api_url: openAISettings.apiUrl
      },
      credentials: 'include'
    })
    
    // پاسخ تست اتصال
    
    if (response.status === 'success') {
      showSuccessMessage('اتصال به OpenAI با موفقیت برقرار شد')
      // بروزرسانی اطلاعات حساب
      if (response.data) {
        openAISettings.accountBalance = response.data.balance || '0.00'
        openAISettings.lastBalanceUpdate = new Date().toLocaleString('fa-IR')
      }
    } else {
      showErrorMessage(response.message || 'خطا در اتصال به OpenAI')
    }
    
  } catch (error) {
    // خطا در تست اتصال OpenAI
    
    // بررسی نوع خطا و نمایش پیام مناسب
    let errorMessage = 'خطا در اتصال به OpenAI'
    
    if (error.data?.message) {
      errorMessage = error.data.message
    } else if (error.message) {
      errorMessage = error.message
    } else if (error.status === 401) {
      errorMessage = 'API Key نامعتبر است. لطفاً کلید API صحیح را وارد کنید.'
    } else if (error.status === 403) {
      errorMessage = 'دسترسی به OpenAI محدود شده است. لطفاً تنظیمات حساب خود را بررسی کنید.'
    } else if (error.status === 404) {
      errorMessage = 'آدرس API نامعتبر است. لطفاً URL صحیح را وارد کنید.'
    } else if (error.status >= 500) {
      errorMessage = 'خطای سرور. لطفاً بعداً تلاش کنید.'
    }
    
    showErrorMessage(errorMessage)
  } finally {
    testingConnection.value = false
  }
}

const fetchUsageData = async () => {
  try {
    fetchingUsage.value = true
    
    // شروع دریافت آمار استفاده OpenAI
    
    const response = await $fetch('/api/admin/api-settings/fetch-usage', {
      method: 'POST',
      credentials: 'include'
    })
    
    // پاسخ دریافت آمار
    
    if (response.status === 'success' && response.data) {
      // بروزرسانی آمار استفاده
      openAISettings.accountBalance = response.data.account_balance || '0.00'
      openAISettings.monthlyUsage = response.data.monthly_usage || '0.00'
      openAISettings.totalRequests = response.data.total_requests || '0'
      openAISettings.todayRequests = response.data.today_requests || '0'
      openAISettings.lastBalanceUpdate = response.data.last_balance_update || new Date().toLocaleString('fa-IR')
      
      showSuccessMessage('آمار استفاده با موفقیت بروزرسانی شد')
    } else {
      showErrorMessage(response.message || 'خطا در دریافت آمار استفاده')
    }
    
  } catch (error) {
    // خطا در دریافت آمار استفاده
    
    let errorMessage = 'خطا در دریافت آمار استفاده'
    
    if (error.data?.message) {
      errorMessage = error.data.message
    } else if (error.message) {
      errorMessage = error.message
    } else if (error.status === 401) {
      errorMessage = 'API Key نامعتبر است یا دسترسی به اطلاعات حساب ندارد. لطفاً کلید API صحیح را وارد کنید.'
    } else if (error.status === 403) {
      errorMessage = 'دسترسی به اطلاعات حساب محدود شده است. ممکن است API Key شما دسترسی به بخش billing نداشته باشد.'
    } else if (error.status === 404) {
      errorMessage = 'آدرس API نامعتبر است. لطفاً URL صحیح را وارد کنید.'
    } else if (error.status >= 500) {
      errorMessage = 'خطای سرور. لطفاً بعداً تلاش کنید.'
    }
    
    // بررسی خطاهای خاص مربوط به دیتابیس
    if (errorMessage.includes('جدول تنظیمات API وجود ندارد') || 
        errorMessage.includes('مایگریشن‌ها را اجرا کنید')) {
      errorMessage = 'خطا در تنظیمات دیتابیس: جدول تنظیمات API وجود ندارد. لطفاً با مدیر سیستم تماس بگیرید.'
    } else if (errorMessage.includes('تنظیمات OpenAI یافت نشد') || 
               errorMessage.includes('تنظیمات API را ذخیره کنید')) {
      errorMessage = 'تنظیمات OpenAI یافت نشد. لطفاً ابتدا تنظیمات API را ذخیره کنید.'
    } else if (errorMessage.includes('OpenAI فعال نیست')) {
      errorMessage = 'OpenAI فعال نیست. لطفاً ابتدا OpenAI را فعال کنید.'
    } else if (errorMessage.includes('API Key تنظیم نشده است')) {
      errorMessage = 'API Key تنظیم نشده است. لطفاً ابتدا کلید API را وارد کنید.'
    }
    
    // نمایش پیام خطای دقیق‌تر
    if (errorMessage.includes('دسترسی به اطلاعات حساب محدود شده است') || 
        errorMessage.includes('دسترسی به آمار استفاده محدود شده است')) {
      errorMessage += '\n\nنکته: برخی از API Key ها ممکن است دسترسی به اطلاعات billing نداشته باشند. این امر طبیعی است و آمار استفاده با مقادیر پیش‌فرض نمایش داده خواهد شد.'
    }
    
    showErrorMessage(errorMessage)
  } finally {
    fetchingUsage.value = false
  }
}

const maskAPIKey = (apiKey) => {
  if (!apiKey || apiKey.length <= 8) {
    return '***'
  }
  return apiKey.substring(0, 4) + '...' + apiKey.substring(apiKey.length - 4)
}

const resetApiSettings = () => {
  // بازنشانی تنظیمات OpenAI
  Object.assign(openAISettings, {
    enabled: false,
    apiKey: '',
    apiUrl: 'https://api.openai.com/v1',
    defaultModel: '',
    temperature: 0.7,
    rateLimit: 60,
    timeout: 30,
    maxDailyCost: 10.00,
    accountBalance: '0.00',
    monthlyUsage: '0.00',
    totalRequests: '0',
    todayRequests: '0',
    lastBalanceUpdate: '',
    consumingPages: [''],
    sectionModels: {},
    availableModels: []
  })
  
  // بازنشانی تنظیمات بخش‌ها
  Object.keys(sectionConfigs).forEach(key => {
    sectionConfigs[key].isEnabled = true
    sectionConfigs[key].temperature = 0.7
    sectionConfigs[key].maxTokens = 1000
  })
  
  // تنظیم مدل‌های پیش‌فرض برای بخش‌ها
  sectionConfigs.chatbot.model = 'gpt-3.5-turbo'
  sectionConfigs.content_generation.model = 'gpt-4'
  sectionConfigs.seo_optimization.model = 'gpt-4-turbo'
  sectionConfigs.product_description.model = 'gpt-3.5-turbo'
  sectionConfigs.customer_support.model = 'gpt-4'
  sectionConfigs.translation.model = 'gpt-3.5-turbo'
  sectionConfigs.code_generation.model = 'gpt-4'
  sectionConfigs.data_analysis.model = 'gpt-4-turbo'
  
  showSuccessMessage('تنظیمات API بازنشانی شد')
}

const saveSettings = async () => {
  try {
    saving.value = true
    
    // شروع ذخیره تنظیمات
    
    // آماده‌سازی داده‌ها برای ارسال
    const normalizePhoneArray = (phones) => {
      if (!Array.isArray(phones)) return []
      return phones
        .map(phone => (phone || '').trim())
        .filter(Boolean)
    }

    const settingsData = {
      shopNameFa: settings.shopNameFa,
      shopNameEn: settings.shopNameEn,
      logo: settings.logo,
      logoRetina: settings.logoRetina,
      favicon: settings.favicon,
      defaultLanguage: settings.defaultLanguage,
      timezone: settings.timezone,
      defaultCurrency: settings.defaultCurrency,
      maintenanceMode: settings.maintenanceMode,
      maintenanceMessage: settings.maintenanceMessage,
      phones: normalizePhoneArray(settings.phones),
      email: settings.email,
      adminPhones: normalizePhoneArray(settings.adminPhones),
      address: settings.address,
      secondaryAddress: settings.secondaryAddress,
  address_secondary: settings.secondaryAddress,
      latitude: settings.coordinates.split(',')[0]?.trim() || '',
      longitude: settings.coordinates.split(',')[1]?.trim() || '',
      workingHours: Array.isArray(settings.workingHours)
        ? settings.workingHours.map(hour => (hour || '').trim()).filter(Boolean)
        : [],
      shortDescription: settings.shortDescription,
      locations: Array.isArray(settings.locations)
        ? settings.locations.map(location => ({
          id: location.id,
          title: (location.title || '').trim(),
          address: (location.address || '').trim(),
          phones: normalizePhoneArray(location.phones)
        }))
        : []
    }
    
    // داده‌های ارسالی
    
    // ارسال یک درخواست تجمیعی
    const response = await $fetch('/api/admin/shop-settings', {
      method: 'PUT',
      body: settingsData
    })
    
    // پاسخ دریافتی
    
    // نمایش پیام موفقیت
    showSuccessMessage('تنظیمات با موفقیت ذخیره شد')
    
  } catch (error) {
    // خطا در ذخیره تنظیمات
    
    // نمایش خطای دقیق به کاربر
    let errorMessage = 'خطا در ذخیره تنظیمات'
    
    if (error.data) {
      // اگر خطا از سرور آمده باشد
      if (error.data.message) {
        errorMessage = error.data.message
      } else if (error.data.error) {
        errorMessage = error.data.error
      } else if (typeof error.data === 'string') {
        errorMessage = error.data
      }
    } else if (error.statusMessage) {
      errorMessage = error.statusMessage
    }
    
    showErrorMessage(errorMessage)
  } finally {
    saving.value = false
  }
}

function resetSettings() {
  // بازنشانی تنظیمات به مقادیر پیش‌فرض
  Object.assign(settings, {
    shopNameFa: 'فروشگاه من',
    shopNameEn: 'My Shop',
    logo: '',
    logoRetina: '',
    favicon: '',
    defaultLanguage: 'fa',
    timezone: 'Asia/Tehran',
    defaultCurrency: 'IRR',
    maintenanceMode: false,
    maintenanceMessage: 'فروشگاه در حال تعمیر و نگهداری است. لطفاً بعداً مراجعه کنید.',
    phones: [''],
    email: '',
    adminPhones: [''],
    address: '',
    secondaryAddress: '',
    latitude: '',
    longitude: '',
    workingHours: [''], // Reset working hours
    shortDescription: '',
    coordinates: '' // Reset combined coordinates
  })
}

// ذخیره تنظیمات فاکتور و چاپ
const saveInvoicePrintSettings = async () => {
  try {
    savingInvoicePrint.value = true
    
    // شروع ذخیره تنظیمات فاکتور و چاپ
    
    // آماده‌سازی داده‌ها برای ارسال
    const invoicePrintData = {
      invoice: invoiceSettings,
      print: printSettings,
      template: templateSettings
    }
    
    console.log('داده‌های ارسالی:', invoicePrintData)
    
    // ارسال درخواست به API
    const response = await $fetch('/api/admin/invoice-print-settings', {
      method: 'PUT',
      body: invoicePrintData
    })
    
    // پاسخ دریافتی
    
    // نمایش پیام موفقیت
    showSuccessMessage('تنظیمات فاکتور و چاپ با موفقیت ذخیره شد')
    
  } catch (error) {
    // خطا در ذخیره تنظیمات فاکتور و چاپ
    
    // نمایش خطای دقیق به کاربر
    let errorMessage = 'خطا در ذخیره تنظیمات فاکتور و چاپ'
    
    if (error.data) {
      // اگر خطا از سرور آمده باشد
      if (error.data.message) {
        errorMessage = error.data.message
      } else if (error.data.error) {
        errorMessage = error.data.error
      } else if (typeof error.data === 'string') {
        errorMessage = error.data
      }
    } else if (error.statusMessage) {
      errorMessage = error.statusMessage
    }
    
    showErrorMessage(errorMessage)
  } finally {
    savingInvoicePrint.value = false
  }
}

// بازنشانی تنظیمات فاکتور و چاپ
const resetInvoicePrintSettings = () => {
  // بازنشانی تنظیمات فاکتور
  Object.assign(invoiceSettings, {
    companyName: 'شرکت من',
    taxNumber: '',
    companyAddress: '',
    companyPhone: ''
  })
  
  // بازنشانی تنظیمات چاپ
  Object.assign(printSettings, {
    defaultPaperSize: 'a4',
    orientation: 'portrait',
    defaultFont: 'iransans',
    fontSize: '12',
    showLogo: true,
    showQRCode: true,
    showBarcode: true
  })
  
  // بازنشانی تنظیمات قالب
  Object.assign(templateSettings, {
    primaryColor: '#3B82F6',
    secondaryColor: '#1E40AF',
    showHeader: true,
    showFooter: true,
    showPageNumbers: true
  })
  
  showSuccessMessage('تنظیمات فاکتور و چاپ بازنشانی شد')
}

// بارگذاری تنظیمات اولیه - فقط یک بار اجرا شود
let settingsLoaded = false

onMounted(async () => {
  if (settingsLoaded) return // جلوگیری از اجرای مجدد
  
  try {
    // بارگذاری تنظیمات
    const response = await $fetch('/api/admin/shop-settings')
    
    if (response.status === 'success' && response.data) {
      Object.assign(settings, response.data)
      
      // ترکیب عرض و طول جغرافیایی
      if (response.data.latitude && response.data.longitude) {
        settings.coordinates = `${response.data.latitude}, ${response.data.longitude}`
      }

      // نرمال‌سازی آدرس‌ها و شماره‌های تماس هر آدرس
      const parseLocationPhones = (rawPhones) => {
        if (Array.isArray(rawPhones)) {
          return rawPhones.length ? rawPhones : ['']
        }
        if (typeof rawPhones === 'string') {
          const parts = rawPhones
            .split(/[\n,|;]/)
            .map(part => part.trim())
            .filter(Boolean)
          return parts.length ? parts : ['']
        }
        return ['']
      }

      let rawLocations = response.data.locations ?? response.data.addresses ?? []

      if (typeof rawLocations === 'string') {
        try {
          rawLocations = JSON.parse(rawLocations)
        } catch (parseError) {
          console.warn('⚠️ Failed to parse locations JSON:', parseError)
          rawLocations = []
        }
      }

      if (!Array.isArray(rawLocations)) {
        rawLocations = []
      }

  const normalizedLocations = rawLocations.map((location, index) => {
        const phones = parseLocationPhones(location?.phones ?? location?.Phones ?? location?.phone ?? location?.Phone)
        return {
          id: location?.id ?? location?.Id ?? (Date.now() + index),
          title: location?.title ?? location?.Title ?? '',
          address: location?.address ?? location?.Address ?? '',
          phones
        }
      })

      if (!normalizedLocations.length) {
        normalizedLocations.push({
          id: Date.now(),
          title: '',
          address: '',
          phones: ['']
        })
      }

      settings.locations = normalizedLocations

      // تبدیل ساعات کاری از string به array
      if (response.data.workingHours) {
        if (Array.isArray(response.data.workingHours)) {
          settings.workingHours = response.data.workingHours
        } else {
          settings.workingHours = response.data.workingHours.split(' | ').filter(h => h.trim())
        }
      }
      
      // اگر خالی بود، یک فیلد خالی اضافه کن
      if (settings.workingHours.length === 0) {
        settings.workingHours = ['']
      }
      
      // اطمینان از وجود آرایه‌های شماره تلفن
      if (!settings.phones || settings.phones.length === 0) {
        settings.phones = ['']
      }
      if (!settings.adminPhones || settings.adminPhones.length === 0) {
        settings.adminPhones = ['']
      }
      
      // تنظیمات بارگذاری شد
    }
    
    // بارگذاری تنظیمات فاکتور و چاپ
    try {
      // بارگذاری تنظیمات فاکتور و چاپ
      const invoicePrintResponse = await $fetch('/api/admin/invoice-print-settings')
      
      if (invoicePrintResponse.status === 'success' && invoicePrintResponse.data) {
        // بارگذاری تنظیمات فاکتور
        if (invoicePrintResponse.data.invoice) {
          Object.assign(invoiceSettings, invoicePrintResponse.data.invoice)
        }
        
        // بارگذاری تنظیمات چاپ
        if (invoicePrintResponse.data.print) {
          Object.assign(printSettings, invoicePrintResponse.data.print)
        }
        
        // بارگذاری تنظیمات قالب
        if (invoicePrintResponse.data.template) {
          Object.assign(templateSettings, invoicePrintResponse.data.template)
        }
        
        // تنظیمات فاکتور و چاپ بارگذاری شد
      }
    } catch (invoicePrintError) {
      // خطا در بارگذاری تنظیمات فاکتور و چاپ
      // این خطا نباید مانع بارگذاری تنظیمات اصلی شود
    }

    // بارگذاری تنظیمات API
    try {
      // بارگذاری تنظیمات API
      const apiSettingsResponse = await $fetch('/api/admin/api-settings')
      
      // پاسخ تنظیمات API
      
      if ((apiSettingsResponse.success || apiSettingsResponse.status === 'success') && apiSettingsResponse.data) {
        // بارگذاری تنظیمات OpenAI
        if (apiSettingsResponse.data.openai) {
          const openaiData = apiSettingsResponse.data.openai
          // داده‌های OpenAI
          
          // بروزرسانی تنظیمات OpenAI با مقادیر دقیق
          openAISettings.enabled = openaiData.enabled || false
          openAISettings.apiKey = openaiData.api_key || ''
          openAISettings.apiUrl = openaiData.api_url || 'https://api.openai.com/v1'
          openAISettings.defaultModel = openaiData.default_model || ''
          openAISettings.temperature = parseFloat(openaiData.temperature) || 0.7
          openAISettings.rateLimit = parseInt(openaiData.rate_limit) || 60
          openAISettings.timeout = parseInt(openaiData.timeout) || 30
          openAISettings.maxDailyCost = parseFloat(openaiData.max_daily_cost) || 10.0
          openAISettings.accountBalance = openaiData.account_balance || '0.00'
          openAISettings.monthlyUsage = openaiData.monthly_usage || '0.00'
          openAISettings.totalRequests = openaiData.total_requests || '0'
          openAISettings.todayRequests = openaiData.today_requests || '0'
          openAISettings.lastBalanceUpdate = openaiData.last_balance_update || ''
          
          // اطمینان از وجود آرایه صفحات مصرف کننده
          if (openaiData.consuming_pages && openaiData.consuming_pages.length > 0) {
            openAISettings.consumingPages = openaiData.consuming_pages
          } else {
            openAISettings.consumingPages = ['']
          }
          
          // بارگذاری مدل‌های موجود
          if (openaiData.available_models && openaiData.available_models.length > 0) {
            availableModels.value = openaiData.available_models
          }
          
          // بارگذاری تنظیمات بخش‌ها
          if (openaiData.section_models) {
            openAISettings.sectionModels = openaiData.section_models
            Object.keys(openaiData.section_models).forEach(sectionKey => {
              if (sectionConfigs[sectionKey]) {
                sectionConfigs[sectionKey].model = openaiData.section_models[sectionKey]
                sectionConfigs[sectionKey].isEnabled = true
              }
            })
          }
          
          // تنظیمات OpenAI بارگذاری شد
        }
        
        // تنظیمات API بارگذاری شد
        
        // دریافت آمار استفاده اگر OpenAI فعال است
        if (openAISettings.enabled && openAISettings.apiKey) {
          try {
            // دریافت آمار استفاده OpenAI
            await fetchUsageData()
          } catch (usageError) {
            // خطا در دریافت آمار استفاده
            // این خطا نباید مانع بارگذاری تنظیمات اصلی شود
          }
        }
      }
    } catch (apiError) {
      // خطا در بارگذاری تنظیمات API
      // این خطا نباید مانع بارگذاری تنظیمات اصلی شود
    }

    // مدل سکشن «سئو تصاویر» را از تنظیمات بخوان و روی کارت اعمال کن
    try {
      const imageSeoRes = await $fetch('/api/admin/settings')
      const arr = Array.isArray(imageSeoRes?.data) ? imageSeoRes.data : (Array.isArray(imageSeoRes) ? imageSeoRes : [])
      const map = {}
      for (const s of arr) { map[s.key || s.Key] = s.value ?? s.Value }
      if (map['image_seo.model'] && sectionConfigs.image_seo) {
        sectionConfigs.image_seo.model = map['image_seo.model']
        sectionConfigs.image_seo.isEnabled = true
      }
    } catch (e) {
      // load image_seo model override failed
    }

    // بارگذاری تنظیمات احراز هویت
    try {
      // بارگذاری تنظیمات احراز هویت
      await loadAuthSettings()
    } catch (authError) {
      // خطا در بارگذاری تنظیمات احراز هویت
      // این خطا نباید مانع بارگذاری تنظیمات اصلی شود
    }

    // بارگذاری تنظیمات صفحه ورود و ثبت‌نام
    try {
      // بارگذاری تنظیمات صفحه ورود
      await loadLoginPageSettings()
    } catch (loginPageError) {
      // خطا در بارگذاری تنظیمات صفحه ورود
      // این خطا نباید مانع بارگذاری تنظیمات اصلی شود
    }

    // بارگذاری تنظیمات شبکه‌های اجتماعی
    try {
      await loadSocialMediaSettings()
    } catch (socialMediaError) {
      // خطا در بارگذاری تنظیمات شبکه‌های اجتماعی
      // این خطا نباید مانع بارگذاری تنظیمات اصلی شود
    }
    
    settingsLoaded = true
  } catch (error) {
    // خطا در بارگذاری تنظیمات
    settingsLoaded = true // حتی در صورت خطا، علامت‌گذاری کن
  }
})

const showErrorMessage = (message) => {
  // اگر از toast استفاده می‌کنید
  if (typeof $toast !== 'undefined') {
    $toast.error(message)
  } else {
    // یا از alert استفاده کنید
    alert(message)
  }
}

const showSuccessMessage = (message) => {
  // اگر از toast استفاده می‌کنید
  if (typeof $toast !== 'undefined') {
    $toast.success(message)
  } else {
    // یا از alert استفاده کنید
    alert(message)
  }
}

// توابع تنظیمات صفحه ورود و ثبت‌نام
const saveLoginPageSettings = async (payload) => {
  try {
    savingLoginPage.value = true
    
    // شروع ذخیره تنظیمات صفحه ورود
    
    // بروزرسانی تنظیمات محلی
    Object.assign(loginPageSettings, payload)
    
    // آماده‌سازی داده‌ها برای ارسال
    const loginPageData = {
      // تنظیمات OTP
      'auth.otp.length': loginPageSettings.otpLength,
      'auth.otp.expiry_minutes': loginPageSettings.otpExpiryMinutes,
      'auth.max_otp_attempts': loginPageSettings.maxOtpAttempts,
      'auth.otp.resend_delay_seconds': loginPageSettings.otpResendDelaySeconds,
      
      // محدودیت‌های ارسال مجدد
      'auth.first_attempts_wait_time': loginPageSettings.firstAttemptsWaitTime,
      'auth.second_attempts_wait_time': loginPageSettings.secondAttemptsWaitTime,
      'auth.fifth_attempt_wait_time': loginPageSettings.fifthAttemptWaitTime,
      'auth.block_time_minutes': loginPageSettings.blockTimeMinutes,
      
      // تنظیمات امنیتی
      'auth.max_login_attempts': loginPageSettings.maxLoginAttempts,
      'auth.lockout_duration_minutes': loginPageSettings.lockoutDurationMinutes,
      'auth.session_timeout_minutes': loginPageSettings.sessionTimeoutMinutes,
      'auth.mobile_auth_enabled': loginPageSettings.mobileAuthEnabled,
      
      // تنظیمات UI
      'auth.login_page_title': loginPageSettings.loginPageTitle,
      'auth.helper_text': loginPageSettings.helperText,
      'auth.show_back_button': loginPageSettings.showBackButton,
      'auth.show_password_login': loginPageSettings.showPasswordLogin
    }
    
    console.log('داده‌های ارسالی:', loginPageData)
    
    // ارسال درخواست به API
    const response = await $fetch('/api/admin/settings', {
      method: 'PUT',
      body: loginPageData
    })
    
    // پاسخ دریافتی
    
    // نمایش پیام موفقیت
    showSuccessMessage('تنظیمات صفحه ورود و ثبت‌نام با موفقیت ذخیره شد')
    
  } catch (error) {
    // خطا در ذخیره تنظیمات صفحه ورود
    
    // نمایش خطای دقیق به کاربر
    let errorMessage = 'خطا در ذخیره تنظیمات صفحه ورود'
    
    if (error.data) {
      if (error.data.message) {
        errorMessage = error.data.message
      } else if (error.data.error) {
        errorMessage = error.data.error
      } else if (typeof error.data === 'string') {
        errorMessage = error.data
      }
    } else if (error.statusMessage) {
      errorMessage = error.statusMessage
    }
    
    showErrorMessage(errorMessage)
  } finally {
    savingLoginPage.value = false
  }
}

const resetLoginPageSettings = () => {
  // بازنشانی تنظیمات صفحه ورود
  Object.assign(loginPageSettings, {
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
  })
  
  showSuccessMessage('تنظیمات صفحه ورود و ثبت‌نام بازنشانی شد')
}

// ذخیره تنظیمات شبکه‌های اجتماعی
const socialMediaEndpoint = '/nuxt/admin/settings/social-media'

const saveSocialMediaSettings = async (payload) => {
  try {
    savingSocialMedia.value = true
    
    console.log('💾 Saving social media settings:', payload)
    
    const sanitizedPayload = {
      ...payload,
      customLinks: Array.isArray(payload.customLinks)
        ? payload.customLinks
            .map(link => ({
              title: (link?.title || '').trim(),
              url: (link?.url || '').trim()
            }))
            .filter(link => link.title.length || link.url.length)
        : []
    }

    const payloadToSave = JSON.parse(JSON.stringify(sanitizedPayload))

    // ارسال داده‌ها به endpoint اختصاصی شبکه‌های اجتماعی
  const response = await $fetch(socialMediaEndpoint, {
      method: 'PUT',
      body: payloadToSave
    })
    
    console.log('✅ Save response:', response)
    
    // بروزرسانی state محلی
  Object.assign(socialMediaSettings, payloadToSave)
    
    // نمایش پیام موفقیت
    showSuccessMessage('تنظیمات شبکه‌های اجتماعی با موفقیت ذخیره شد')
    
  } catch (error) {
    console.error('❌ خطا در ذخیره تنظیمات شبکه‌های اجتماعی:', error)
    showErrorMessage('خطا در ذخیره تنظیمات شبکه‌های اجتماعی')
  } finally {
    savingSocialMedia.value = false
  }
}

// بازنشانی تنظیمات شبکه‌های اجتماعی
const resetSocialMediaSettings = () => {
  Object.assign(socialMediaSettings, {
    instagram: '',
    instagram_enabled: true,
    telegram: '',
    telegram_enabled: true,
    whatsapp: '',
    whatsapp_enabled: true,
    twitter: '',
    twitter_enabled: true,
    linkedin: '',
    linkedin_enabled: true,
    facebook: '',
    facebook_enabled: true,
    youtube: '',
    youtube_enabled: true,
    aparat: '',
    aparat_enabled: true,
    rubika: '',
    rubika_enabled: true,
    eitaa: '',
    eitaa_enabled: true,
    bale: '',
    bale_enabled: true,
    tiktok: '',
    tiktok_enabled: true,
    pinterest: '',
    pinterest_enabled: true,
    discord: '',
    discord_enabled: true,
    github: '',
    github_enabled: true,
    customLinks: []
  })
  
  showSuccessMessage('تنظیمات شبکه‌های اجتماعی بازنشانی شد')
}

const mapSocialMediaRecordsToObject = (records) => {
  const settingsMap = {}
  const prefixedKeys = new Set()

  records.forEach((item) => {
    if (!item) return
    const key = item.key ?? item.Key
    if (!key) return

    const value = item.value ?? item.Value ?? ''
    const cleanKey = key.replace(/^social-media\./, '')
    const isPrefixed = key.startsWith('social-media.')

    if (isPrefixed) {
      prefixedKeys.add(cleanKey)
      settingsMap[cleanKey] = value
    } else if (!prefixedKeys.has(cleanKey) && settingsMap[cleanKey] === undefined) {
      settingsMap[cleanKey] = value
    }
  })

  const rawCustomLinks = settingsMap.custom_links ?? settingsMap.customLinks
  if (typeof rawCustomLinks === 'string') {
    try {
      settingsMap.customLinks = JSON.parse(rawCustomLinks)
    } catch (error) {
      console.warn('⚠️ Failed to parse custom_links JSON:', error)
      settingsMap.customLinks = []
    }
  } else if (Array.isArray(rawCustomLinks)) {
    settingsMap.customLinks = rawCustomLinks
  } else {
    settingsMap.customLinks = []
  }

  return settingsMap
}

const normalizeSocialMediaResponse = (raw) => {
  if (!raw) {
    return { success: false, data: {} }
  }

  if (raw.success && raw.data) {
    if (Array.isArray(raw.data)) {
      return { success: true, data: mapSocialMediaRecordsToObject(raw.data) }
    }
    if (raw.data && typeof raw.data === 'object') {
      return { success: true, data: raw.data }
    }
  }

  if (Array.isArray(raw)) {
    return { success: true, data: mapSocialMediaRecordsToObject(raw) }
  }

  if (raw.data && Array.isArray(raw.data)) {
    return { success: true, data: mapSocialMediaRecordsToObject(raw.data) }
  }

  const recordKey = raw.key ?? raw.Key
  const recordValue = raw.value ?? raw.Value

  if (recordKey === 'social-media' && typeof recordValue === 'string') {
    try {
      const parsed = JSON.parse(recordValue)
      return { success: true, data: parsed }
    } catch (error) {
      console.warn('⚠️ Failed to parse social-media JSON value:', error)
    }
  }

  if (recordKey) {
    return { success: true, data: mapSocialMediaRecordsToObject([raw]) }
  }

  if (raw.data && typeof raw.data === 'object') {
    return { success: true, data: raw.data }
  }

  return { success: false, data: {} }
}

// بارگذاری تنظیمات شبکه‌های اجتماعی
const loadSocialMediaSettings = async () => {
  try {
    console.log('🔄 Loading social media settings...')
  const response = await $fetch(socialMediaEndpoint)
    
    console.log('📥 Response:', response)

    const normalized = normalizeSocialMediaResponse(response)

    if (normalized.success && normalized.data) {
      const settings = normalized.data
      console.log('📊 Settings data:', settings)
      
      // بارگذاری تنظیمات به صورت مستقیم
      Object.keys(socialMediaSettings).forEach(key => {
        if (key === 'customLinks') {
          let parsedLinks = []
          const rawLinks = settings.custom_links ?? settings.customLinks
          if (typeof rawLinks === 'string') {
            try {
              parsedLinks = JSON.parse(rawLinks)
            } catch (parseError) {
              console.warn('⚠️ Unable to parse custom links JSON:', parseError)
            }
          } else if (Array.isArray(rawLinks)) {
            parsedLinks = rawLinks
          }
          socialMediaSettings.customLinks = Array.isArray(parsedLinks) ? parsedLinks : []
          console.log('✅ Loaded customLinks:', socialMediaSettings.customLinks)
          return
        }

        if (settings[key] !== undefined) {
          // تبدیل رشته boolean به boolean واقعی
          if (key.endsWith('_enabled')) {
            socialMediaSettings[key] = settings[key] === 'true' || settings[key] === true
            console.log(`✅ Loaded ${key}:`, socialMediaSettings[key])
          } else {
            socialMediaSettings[key] = settings[key]
            console.log(`✅ Loaded ${key}:`, socialMediaSettings[key])
          }
        } else {
          console.log(`⚠️ Key not found in response: ${key}`)
        }
      })
      
      console.log('🎯 Final socialMediaSettings:', socialMediaSettings)
    } else {
      console.log('❌ Invalid response format')
    }
  } catch (error) {
    console.error('❌ خطا در بارگذاری تنظیمات شبکه‌های اجتماعی:', error)
  }
}

// بارگذاری تنظیمات صفحه ورود و ثبت‌نام
const loadLoginPageSettings = async () => {
  try {
    const response = await $fetch('/api/admin/settings')
    
    if (response && response.data) {
      const settings = response.data
      const settingsMap = {}
      
      // تبدیل آرایه تنظیمات به map برای دسترسی آسان‌تر
      if (Array.isArray(settings)) {
        settings.forEach(setting => {
          settingsMap[setting.key || setting.Key] = setting.value ?? setting.Value
        })
      }
      
      // بارگذاری تنظیمات OTP
      if (settingsMap['auth.otp.length']) {
        loginPageSettings.otpLength = parseInt(settingsMap['auth.otp.length']) || 5
      }
      if (settingsMap['auth.otp.expiry_minutes']) {
        loginPageSettings.otpExpiryMinutes = parseInt(settingsMap['auth.otp.expiry_minutes']) || 5
      }
      if (settingsMap['auth.max_otp_attempts']) {
        loginPageSettings.maxOtpAttempts = parseInt(settingsMap['auth.max_otp_attempts']) || 3
      }
      if (settingsMap['auth.otp.resend_delay_seconds']) {
        loginPageSettings.otpResendDelaySeconds = parseInt(settingsMap['auth.otp.resend_delay_seconds']) || 60
      }
      
      // بارگذاری محدودیت‌های ارسال مجدد
      if (settingsMap['auth.first_attempts_wait_time']) {
        loginPageSettings.firstAttemptsWaitTime = parseInt(settingsMap['auth.first_attempts_wait_time']) || 2
      }
      if (settingsMap['auth.second_attempts_wait_time']) {
        loginPageSettings.secondAttemptsWaitTime = parseInt(settingsMap['auth.second_attempts_wait_time']) || 3
      }
      if (settingsMap['auth.fifth_attempt_wait_time']) {
        loginPageSettings.fifthAttemptWaitTime = parseInt(settingsMap['auth.fifth_attempt_wait_time']) || 5
      }
      if (settingsMap['auth.block_time_minutes']) {
        loginPageSettings.blockTimeMinutes = parseInt(settingsMap['auth.block_time_minutes']) || 60
      }
      
      // بارگذاری تنظیمات امنیتی
      if (settingsMap['auth.max_login_attempts']) {
        loginPageSettings.maxLoginAttempts = parseInt(settingsMap['auth.max_login_attempts']) || 5
      }
      if (settingsMap['auth.lockout_duration_minutes']) {
        loginPageSettings.lockoutDurationMinutes = parseInt(settingsMap['auth.lockout_duration_minutes']) || 15
      }
      if (settingsMap['auth.session_timeout_minutes']) {
        loginPageSettings.sessionTimeoutMinutes = parseInt(settingsMap['auth.session_timeout_minutes']) || 1440
      }
      if (settingsMap['auth.mobile_auth_enabled'] !== undefined) {
        loginPageSettings.mobileAuthEnabled = settingsMap['auth.mobile_auth_enabled'] === 'true' || settingsMap['auth.mobile_auth_enabled'] === true
      }
      
      // بارگذاری تنظیمات UI
      if (settingsMap['auth.login_page_title']) {
        loginPageSettings.loginPageTitle = settingsMap['auth.login_page_title']
      }
      if (settingsMap['auth.helper_text']) {
        loginPageSettings.helperText = settingsMap['auth.helper_text']
      }
      if (settingsMap['auth.show_back_button'] !== undefined) {
        loginPageSettings.showBackButton = settingsMap['auth.show_back_button'] === 'true' || settingsMap['auth.show_back_button'] === true
      }
      if (settingsMap['auth.show_password_login'] !== undefined) {
        loginPageSettings.showPasswordLogin = settingsMap['auth.show_password_login'] === 'true' || settingsMap['auth.show_password_login'] === true
      }
      
      // تنظیمات صفحه ورود بارگذاری شد
    }
  } catch (error) {
    // خطا در بارگذاری تنظیمات صفحه ورود
    // در صورت خطا، تنظیمات پیش‌فرض حفظ می‌شوند
  }
}

// تولید کلید مخفی جدید JWT
const generateNewJwtSecret = () => {
  const chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*'
  let result = ''
  for (let i = 0; i < 64; i++) {
    result += chars.charAt(Math.floor(Math.random() * chars.length))
  }
  authSettings.jwtSecret = result
  showSuccessMessage('کلید مخفی جدید تولید شد')
}

// ذخیره تنظیمات احراز هویت
const saveAuthSettings = async () => {
  try {
    savingAuth.value = true
    
    // شروع ذخیره تنظیمات احراز هویت
    
    // آماده‌سازی داده‌ها برای ارسال
    const settingsData = {
                      // تنظیمات سیستم یکپارچه
                unified_auth_enabled: authSettings.unifiedAuthEnabled,
                auto_register_enabled: authSettings.autoRegisterEnabled,
                default_auth_method: authSettings.defaultAuthMethod,

                      // تنظیمات OTP
                mobile_auth_enabled: authSettings.mobileAuthEnabled,
                otp_length: authSettings.otpLength,
                otp_expiry_minutes: authSettings.otpExpiryMinutes,
                max_otp_attempts: authSettings.maxOtpAttempts,
                otp_resend_cooldown: authSettings.otpResendCooldown,
                otp_pattern_code: authSettings.otpPatternCode,

                // تنظیمات JWT
                jwt_expiry_hours: authSettings.jwtExpiryHours,
                refresh_token_expiry_days: authSettings.refreshTokenExpiryDays,
                jwt_secret: authSettings.jwtSecret,
                max_concurrent_sessions: authSettings.maxConcurrentSessions,

                // تنظیمات ورود با یوزرنیم
                username_auth_enabled: authSettings.usernameAuthEnabled,
                min_password_length: authSettings.minPasswordLength,
                max_login_attempts: authSettings.maxLoginAttempts,
                account_lockout_minutes: authSettings.accountLockoutMinutes,
                password_expiry_days: authSettings.passwordExpiryDays,

                // تنظیمات امنیتی
                two_factor_enabled: authSettings.twoFactorEnabled,
                suspicious_activity_detection: authSettings.suspiciousActivityDetection,
                session_timeout_minutes: authSettings.sessionTimeoutMinutes,
                login_history_retention_days: authSettings.loginHistoryRetentionDays,
                auto_block_failed_logins: authSettings.autoBlockFailedLogins,
                auto_block_duration_hours: authSettings.autoBlockDurationHours,

                // تنظیمات نقش‌ها
                default_user_role: authSettings.defaultUserRole,
                email_verification_enabled: authSettings.emailVerificationEnabled,
                phone_verification_enabled: authSettings.phoneVerificationEnabled
    }
    
    // داده‌های ارسالی
    
    // ارسال تنظیمات احراز هویت
    const response = await $fetch('/api/admin/settings/auth', {
      method: 'PUT',
      body: settingsData
    })
    
    // پاسخ دریافتی
    
    // نمایش پیام موفقیت
    showSuccessMessage('تنظیمات احراز هویت با موفقیت ذخیره شد')
    
  } catch (error) {
    // خطا در ذخیره تنظیمات احراز هویت
    showErrorMessage('خطا در ذخیره تنظیمات احراز هویت')
  } finally {
    savingAuth.value = false
  }
}

// بارگذاری تنظیمات احراز هویت
const loadAuthSettings = async () => {
  try {
    const response = await $fetch('/api/admin/settings/auth')
    
    if (response && response.data) {
      const settings = response.data
      
      // تنظیمات سیستم یکپارچه
      authSettings.unifiedAuthEnabled = settings.unified_auth_enabled !== undefined ? settings.unified_auth_enabled : true
      authSettings.autoRegisterEnabled = settings.auto_register_enabled !== undefined ? settings.auto_register_enabled : true
      authSettings.defaultAuthMethod = settings.default_auth_method || 'otp'

      // تنظیمات OTP
      authSettings.mobileAuthEnabled = settings.mobile_auth_enabled || true
      authSettings.otpLength = settings.otp_length || 5
      authSettings.otpExpiryMinutes = settings.otp_expiry_minutes || 5
      authSettings.maxOtpAttempts = settings.max_otp_attempts || 3
      authSettings.otpResendCooldown = settings.otp_resend_cooldown || 60
      authSettings.otpPatternCode = settings.otp_pattern_code || ''

      // تنظیمات JWT
      authSettings.jwtExpiryHours = settings.jwt_expiry_hours || 24
      authSettings.refreshTokenExpiryDays = settings.refresh_token_expiry_days || 30
      authSettings.jwtSecret = settings.jwt_secret || ''
      authSettings.maxConcurrentSessions = settings.max_concurrent_sessions || 5

      // تنظیمات ورود با یوزرنیم
      authSettings.usernameAuthEnabled = settings.username_auth_enabled || true
      authSettings.minPasswordLength = settings.min_password_length || 8
      authSettings.maxLoginAttempts = settings.max_login_attempts || 5
      authSettings.accountLockoutMinutes = settings.account_lockout_minutes || 15
      authSettings.passwordExpiryDays = settings.password_expiry_days || 90

      // تنظیمات امنیتی
      authSettings.twoFactorEnabled = settings.two_factor_enabled || false
      authSettings.suspiciousActivityDetection = settings.suspicious_activity_detection || true
      authSettings.sessionTimeoutMinutes = settings.session_timeout_minutes || 30
      authSettings.loginHistoryRetentionDays = settings.login_history_retention_days || 365
      authSettings.autoBlockFailedLogins = settings.auto_block_failed_logins || 10
      authSettings.autoBlockDurationHours = settings.auto_block_duration_hours || 24

      // تنظیمات نقش‌ها
      authSettings.defaultUserRole = settings.default_user_role || 'user'
      authSettings.emailVerificationEnabled = settings.email_verification_enabled || false
      authSettings.phoneVerificationEnabled = settings.phone_verification_enabled || true
    }
  } catch (error) {
    // خطا در بارگذاری تنظیمات احراز هویت
  }
}

// تنظیمات فروشگاه
const settings = reactive({
  // اطلاعات پایه
  shopNameFa: 'فروشگاه من',
  shopNameEn: 'My Shop',
  logo: '',
  logoRetina: '',
  favicon: '',
  
  // تنظیمات منطقه‌ای
  defaultLanguage: 'fa',
  timezone: 'Asia/Tehran',
  defaultCurrency: 'IRR',
  
  // وضعیت فروشگاه
  maintenanceMode: false,
  maintenanceMessage: 'فروشگاه در حال تعمیر و نگهداری است. لطفاً بعداً مراجعه کنید.',
  
  // اطلاعات تماس - ساختار جدید
  locations: [
    {
      id: Date.now(),
      title: '',
      address: '',
      phones: ['']
    }
  ],
  phones: [''],
  email: '',
  adminPhones: [''],
  address: '',
  secondaryAddress: '',
  
  // اطلاعات اضافی
  workingHours: [''], // Changed to array for multiple working hours
  shortDescription: '',
  coordinates: '', // Added for combined coordinates
})

// متدهای کمکی
const getModelStatusClass = (isActive) => {
  return isActive ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800'
}

// بروزرسانی تنظیمات بخش‌ها
const updateSectionModels = async () => {
  const sectionModels = {}
  Object.keys(sectionConfigs).forEach(key => {
    if (sectionConfigs[key].isEnabled) {
      sectionModels[key] = sectionConfigs[key].model
    }
  })
  openAISettings.sectionModels = sectionModels
  // همگام‌سازی مدل «سئو تصاویر» با صفحهٔ هوش مصنوعی تصاویر
  try {
    if (sectionConfigs.image_seo && sectionConfigs.image_seo.model) {
      $fetch('/api/admin/settings', {
        method: 'PUT',
        body: [
          { key: 'image_seo.model', value: sectionConfigs.image_seo.model, category: 'image_seo', type: 'string' }
        ]
      }).catch(()=>{})
    }
  } catch (e) {
    // sync image_seo model failed
  }
}

// تماشای تغییرات در تنظیمات بخش‌ها
watch(sectionConfigs, () => {
  updateSectionModels()
}, { deep: true })
</script>

<style scoped>
.settings-layout {
  display: flex;
  flex-direction: row-reverse;
  gap: 0;
  min-height: 70vh;
  direction: rtl;
  position: relative;
}
.settings-menu {
  background: #222d32;
  padding: 16px 0 16px 16px;
  border-radius: 0 8px 8px 0;
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  min-width: 180px;
  max-width: 200px;
  width: 200px;
  margin-left: 0;
  margin-right: 0;
  height: fit-content;
  text-align: right;
  box-shadow: 0 2px 8px #0001;
  position: absolute;
  right: 0;
  top: 0;
  z-index: 2;
}
.settings-menu-item {
  color: #fff;
  padding: 8px 0;
  border-radius: 6px;
  font-size: 1em;
  display: flex;
  flex-direction: row;
  align-items: center;
  cursor: pointer;
  transition: background 0.2s;
  width: 100%;
  justify-content: flex-end;
  margin-bottom: 2px;
  text-align: right;
}
.settings-menu-item.active {
  background: #e74c3c;
  color: #fff;
}
.settings-menu-item:hover {
  background: #2d3842;
}
.settings-menu-item .ml-2 {
  margin-left: 8px;
  margin-right: 0;
  font-size: 1.1em;
}
.settings-content {
  flex: 1;
  background: #fff;
  border-radius: 8px 0 0 8px;
  padding: 32px 24px;
  min-height: 400px;
  box-shadow: 0 2px 8px #0001;
  margin-left: 0;
  text-align: right;
  margin-right: 210px;
}
.settings-content h2, .settings-content p, .settings-content > div {
  text-align: right;
  direction: rtl;
}
.settings-menu-header {
  color:#aaa;
  font-weight:bold;
  padding:6px 0;
  text-align:right;
}
.settings-menu-item.child {
  padding-right: 20px;
  font-size: 0.9em;
}

/* Custom slider styles */
.slider {
  -webkit-appearance: none;
  appearance: none;
  background: transparent;
  cursor: pointer;
}

.slider::-webkit-slider-track {
  background: #e5e7eb;
  height: 8px;
  border-radius: 4px;
}

.slider::-webkit-slider-thumb {
  -webkit-appearance: none;
  appearance: none;
  background: #10b981;
  height: 20px;
  width: 20px;
  border-radius: 50%;
  cursor: pointer;
  transition: all 0.2s ease;
}

.slider::-webkit-slider-thumb:hover {
  background: #059669;
  transform: scale(1.1);
}

.slider::-moz-range-track {
  background: #e5e7eb;
  height: 8px;
  border-radius: 4px;
  border: none;
}

.slider::-moz-range-thumb {
  background: #10b981;
  height: 20px;
  width: 20px;
  border-radius: 50%;
  cursor: pointer;
  border: none;
  transition: all 0.2s ease;
}

.slider::-moz-range-thumb:hover {
  background: #059669;
  transform: scale(1.1);
}

.slider:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.slider:disabled::-webkit-slider-thumb {
  background: #9ca3af;
  cursor: not-allowed;
}

.slider:disabled::-moz-range-thumb {
  background: #9ca3af;
  cursor: not-allowed;
}
</style> 

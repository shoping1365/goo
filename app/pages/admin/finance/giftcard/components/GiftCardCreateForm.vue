<template>
  <div v-if="hasAccess" class="space-y-6">
    <!-- Header -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex justify-between items-center">
        <div>
          <h2 class="text-xl font-bold text-gray-900">ایجاد گیفت کارت جدید</h2>
          <p class="text-gray-600 mt-1">ایجاد کارت هدیه جدید با تنظیمات کامل</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button 
            class="px-4 py-2 bg-gray-600 text-white text-sm font-medium rounded-lg hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
            @click="resetForm"
          >
            پاک کردن فرم
          </button>
          <button 
            class="px-4 py-2 bg-yellow-600 text-white text-sm font-medium rounded-lg hover:bg-yellow-700 focus:outline-none focus:ring-2 focus:ring-yellow-500 focus:ring-offset-2"
            @click="saveDraft"
          >
            ذخیره پیش‌نویس
          </button>
        </div>
      </div>
    </div>

    <!-- تب‌های فرم -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200">
      <div class="border-b border-gray-200">
        <nav class="-mb-px flex space-x-8 space-x-reverse" aria-label="Tabs">
          <button
            v-for="tab in tabs"
            :key="tab.id"
            :class="[
              activeTab === tab.id
                ? 'border-blue-500 text-blue-600'
                : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300',
              'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm'
            ]"
            @click="activeTab = tab.id"
          >
            {{ tab.label }}
            <span v-if="tab.hasError" class="ml-2 inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium bg-red-100 text-red-800">
              خطا
            </span>
          </button>
        </nav>
      </div>

      <div class="p-6">
        <!-- تب اطلاعات پایه -->
        <div v-if="activeTab === 'basic'" class="space-y-6">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <!-- نوع کارت -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نوع کارت *</label>
              <select
                v-model="form.type"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                :class="{ 'border-red-500': errors.type }"
              >
                <option value="">انتخاب کنید</option>
                <option value="digital">دیجیتال</option>
                <option value="physical">فیزیکی</option>
                <option value="virtual">مجازی</option>
              </select>
              <p v-if="errors.type" class="mt-1 text-sm text-red-600">{{ errors.type }}</p>
            </div>

            <!-- دسته‌بندی -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">دسته‌بندی</label>
              <select
                v-model="form.category"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="">انتخاب کنید</option>
                <option value="birthday">تولد</option>
                <option value="wedding">عروسی</option>
                <option value="newyear">سال نو</option>
                <option value="christmas">کریسمس</option>
                <option value="graduation">فارغ‌التحصیلی</option>
                <option value="anniversary">سالگرد</option>
                <option value="general">عمومی</option>
              </select>
            </div>

            <!-- مبلغ -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">مبلغ (تومان) *</label>
              <input
                v-model.number="form.amount"
                type="number"
                min="1000"
                step="1000"
                placeholder="مثال: 100000"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                :class="{ 'border-red-500': errors.amount }"
              />
              <p v-if="errors.amount" class="mt-1 text-sm text-red-600">{{ errors.amount }}</p>
            </div>

            <!-- تعداد -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">تعداد *</label>
              <input
                v-model.number="form.quantity"
                type="number"
                min="1"
                max="1000"
                placeholder="مثال: 1"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                :class="{ 'border-red-500': errors.quantity }"
              />
              <p v-if="errors.quantity" class="mt-1 text-sm text-red-600">{{ errors.quantity }}</p>
            </div>

            <!-- تاریخ انقضا -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ انقضا *</label>
              <input
                v-model="form.expiryDate"
                type="date"
                :min="minExpiryDate"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                :class="{ 'border-red-500': errors.expiryDate }"
              />
              <p v-if="errors.expiryDate" class="mt-1 text-sm text-red-600">{{ errors.expiryDate }}</p>
            </div>

            <!-- مدت اعتبار -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">مدت اعتبار (روز)</label>
              <input
                v-model.number="form.validityDays"
                type="number"
                min="1"
                max="365"
                placeholder="مثال: 30"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              />
              <p class="mt-1 text-sm text-gray-500">در صورت خالی بودن، از تاریخ انقضا استفاده می‌شود</p>
            </div>
          </div>

          <!-- کد کارت -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">کد کارت</label>
            <div class="flex space-x-2 space-x-reverse">
              <input
                v-model="form.codePrefix"
                type="text"
                placeholder="پیشوند (مثال: GC)"
                class="flex-1 px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              />
              <span class="px-3 py-2 text-gray-500">-</span>
              <input
                v-model="form.codeSuffix"
                type="text"
                placeholder="پسوند (مثال: 2024)"
                class="flex-1 px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              />
              <button
                type="button"
                class="px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
                @click="generateCode"
              >
                تولید کد
              </button>
            </div>
            <p class="mt-1 text-sm text-gray-500">نمونه: {{ generatedCode }}</p>
          </div>
        </div>

        <!-- تب طراحی و ظاهر -->
        <div v-if="activeTab === 'design'" class="space-y-6">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <!-- قالب -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">قالب کارت</label>
              <select
                v-model="form.template"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="">قالب پیش‌فرض</option>
                <option value="birthday">تولد</option>
                <option value="wedding">عروسی</option>
                <option value="christmas">کریسمس</option>
                <option value="newyear">سال نو</option>
                <option value="custom">سفارشی</option>
              </select>
            </div>

            <!-- رنگ اصلی -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">رنگ اصلی</label>
              <input
                v-model="form.primaryColor"
                type="color"
                class="w-full h-10 border border-gray-300 rounded-md"
              />
            </div>

            <!-- رنگ ثانویه -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">رنگ ثانویه</label>
              <input
                v-model="form.secondaryColor"
                type="color"
                class="w-full h-10 border border-gray-300 rounded-md"
              />
            </div>

            <!-- فونت -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">فونت</label>
              <select
                v-model="form.font"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="IRANSans">ایران سنس</option>
                <option value="Vazir">وزیر</option>
                <option value="Shabnam">شبنم</option>
                <option value="Yekan">یکان</option>
              </select>
            </div>
          </div>

          <!-- پیش‌نمایش -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">پیش‌نمایش کارت</label>
            <div class="border border-gray-300 rounded-lg p-6 bg-gray-50">
              <div 
                class="w-full h-32 rounded-lg flex items-center justify-center text-white font-bold text-lg"
                :style="{
                  background: `linear-gradient(135deg, ${form.primaryColor || '#3B82F6'} 0%, ${form.secondaryColor || '#8B5CF6'} 100%)`,
                  fontFamily: getFontFamily(form.font)
                }"
              >
                {{ form.amount ? formatCurrency(form.amount) : 'مبلغ کارت' }}
              </div>
            </div>
          </div>
        </div>

        <!-- تب گیرنده و ارسال -->
        <div v-if="activeTab === 'recipient'" class="space-y-6">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <!-- نام گیرنده -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نام گیرنده</label>
              <input
                v-model="form.recipientName"
                type="text"
                placeholder="نام و نام خانوادگی"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              />
            </div>

            <!-- ایمیل گیرنده -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">ایمیل گیرنده</label>
              <input
                v-model="form.recipientEmail"
                type="email"
                placeholder="example@email.com"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                :class="{ 'border-red-500': errors.recipientEmail }"
              />
              <p v-if="errors.recipientEmail" class="mt-1 text-sm text-red-600">{{ errors.recipientEmail }}</p>
            </div>

            <!-- شماره موبایل گیرنده -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">شماره موبایل گیرنده</label>
              <input
                v-model="form.recipientPhone"
                type="tel"
                placeholder="09123456789"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              />
            </div>

            <!-- روش تحویل -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">روش تحویل</label>
              <select
                v-model="form.deliveryMethod"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="email">ایمیل</option>
                <option value="sms">پیامک</option>
                <option value="both">هر دو</option>
                <option value="manual">دستی</option>
              </select>
            </div>

            <!-- تاریخ تحویل -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ تحویل</label>
              <input
                v-model="form.deliveryDate"
                type="datetime-local"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              />
            </div>

            <!-- پیام شخصی -->
            <div class="md:col-span-2">
              <label class="block text-sm font-medium text-gray-700 mb-2">پیام شخصی</label>
              <textarea
                v-model="form.personalMessage"
                rows="4"
                placeholder="پیام شخصی برای گیرنده..."
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              ></textarea>
            </div>
          </div>
        </div>

        <!-- تب تنظیمات پیشرفته -->
        <div v-if="activeTab === 'advanced'" class="space-y-6">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <!-- محدودیت استفاده -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">محدودیت استفاده</label>
              <select
                v-model="form.usageLimit"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="unlimited">نامحدود</option>
                <option value="once">یک بار</option>
                <option value="multiple">چند بار</option>
              </select>
            </div>

            <!-- تعداد استفاده مجاز -->
            <div v-if="form.usageLimit === 'multiple'">
              <label class="block text-sm font-medium text-gray-700 mb-2">تعداد استفاده مجاز</label>
              <input
                v-model.number="form.maxUsageCount"
                type="number"
                min="2"
                max="10"
                placeholder="مثال: 3"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              />
            </div>

            <!-- حداقل مبلغ سفارش -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">حداقل مبلغ سفارش</label>
              <input
                v-model.number="form.minOrderAmount"
                type="number"
                min="0"
                placeholder="مثال: 50000"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              />
            </div>

            <!-- حداکثر مبلغ سفارش -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر مبلغ سفارش</label>
              <input
                v-model.number="form.maxOrderAmount"
                type="number"
                min="0"
                placeholder="مثال: 1000000"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              />
            </div>

            <!-- دسته‌بندی‌های مجاز -->
            <div class="md:col-span-2">
              <label class="block text-sm font-medium text-gray-700 mb-2">دسته‌بندی‌های مجاز</label>
              <div class="grid grid-cols-2 md:grid-cols-4 gap-2">
                <label v-for="category in availableCategories" :key="category.value" class="flex items-center">
                  <input
                    v-model="form.allowedCategories"
                    :value="category.value"
                    type="checkbox"
                    class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                  />
                  <span class="mr-2 text-sm text-gray-700">{{ category.label }}</span>
                </label>
              </div>
            </div>

            <!-- نیاز به تأیید -->
            <div class="md:col-span-2">
              <div class="flex items-center">
                <input
                  v-model="form.requiresApproval"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label class="mr-2 text-sm text-gray-700">نیاز به تأیید قبل از استفاده</label>
              </div>
            </div>

            <!-- فعال بودن -->
            <div class="md:col-span-2">
              <div class="flex items-center">
                <input
                  v-model="form.isActive"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label class="mr-2 text-sm text-gray-700">کارت فعال باشد</label>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- دکمه‌های عملیات -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex justify-between items-center">
        <div class="flex items-center space-x-3 space-x-reverse">
          <button
            :disabled="activeTab === 'basic'"
            class="px-4 py-2 bg-gray-600 text-white text-sm font-medium rounded-lg hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed"
            @click="previousTab"
          >
            قبلی
          </button>
          <button
            :disabled="activeTab === 'advanced'"
            class="px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed"
            @click="nextTab"
          >
            بعدی
          </button>
        </div>
        
        <div class="flex items-center space-x-3 space-x-reverse">
          <button
            class="px-6 py-2 bg-green-600 text-white text-sm font-medium rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2"
            @click="validateAndCreate"
          >
            ایجاد کارت
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>
</script>

<script setup lang="ts">
import { computed, onMounted, reactive, ref, watch } from 'vue';
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


// Props
interface InitialData {
  [key: string]: unknown
}

const props = defineProps<{
  initialData?: InitialData
}>()

interface GiftCardData {
  id?: number | string
  code?: string
  amount?: number
  [key: string]: unknown
}

// Emits
const emit = defineEmits<{
  'create-card': [cardData: GiftCardData]
  'save-draft': [cardData: GiftCardData]
}>()

// Reactive data
const activeTab = ref('basic')

interface Tab {
  id: string
  label: string
  hasError?: boolean
}

const tabs = ref<Tab[]>([
  { id: 'basic', label: 'اطلاعات پایه' },
  { id: 'design', label: 'طراحی و ظاهر' },
  { id: 'recipient', label: 'گیرنده و ارسال' },
  { id: 'advanced', label: 'تنظیمات پیشرفته' }
])

const form = reactive({
  // اطلاعات پایه
  type: '',
  category: '',
  amount: null,
  quantity: 1,
  expiryDate: '',
  validityDays: null,
  codePrefix: 'GC',
  codeSuffix: '',
  
  // طراحی
  template: '',
  primaryColor: '#3B82F6',
  secondaryColor: '#8B5CF6',
  font: 'IRANSans',
  
  // گیرنده
  recipientName: '',
  recipientEmail: '',
  recipientPhone: '',
  deliveryMethod: 'email',
  deliveryDate: '',
  personalMessage: '',
  
  // پیشرفته
  usageLimit: 'unlimited',
  maxUsageCount: 3,
  minOrderAmount: null,
  maxOrderAmount: null,
  allowedCategories: [],
  requiresApproval: false,
  isActive: true
})

const errors = reactive({
  type: '',
  amount: '',
  quantity: '',
  expiryDate: '',
  recipientEmail: ''
})

const availableCategories = [
  { value: 'electronics', label: 'الکترونیک' },
  { value: 'clothing', label: 'پوشاک' },
  { value: 'books', label: 'کتاب' },
  { value: 'home', label: 'خانه و آشپزخانه' },
  { value: 'sports', label: 'ورزشی' },
  { value: 'beauty', label: 'زیبایی' },
  { value: 'toys', label: 'اسباب بازی' },
  { value: 'food', label: 'مواد غذایی' }
]

// Computed properties
const minExpiryDate = computed(() => {
  const today = new Date()
  today.setDate(today.getDate() + 1)
  return today.toISOString().split('T')[0]
})

const generatedCode = computed(() => {
  const prefix = form.codePrefix || 'GC'
  const suffix = form.codeSuffix || new Date().getFullYear()
  const random = Math.random().toString(36).substring(2, 8).toUpperCase()
  return `${prefix}-${suffix}-${random}`
})

// Methods
const nextTab = () => {
  const currentIndex = tabs.value.findIndex(tab => tab.id === activeTab.value)
  if (currentIndex < tabs.value.length - 1) {
    activeTab.value = tabs.value[currentIndex + 1].id
  }
}

const previousTab = () => {
  const currentIndex = tabs.value.findIndex(tab => tab.id === activeTab.value)
  if (currentIndex > 0) {
    activeTab.value = tabs.value[currentIndex - 1].id
  }
}

const generateCode = () => {
  form.codeSuffix = new Date().getFullYear().toString()
}

const validateForm = () => {
  let isValid = true
  
  // پاک کردن خطاهای قبلی
  Object.keys(errors).forEach(key => {
    errors[key] = ''
  })
  
  // اعتبارسنجی نوع کارت
  if (!form.type) {
    errors.type = 'نوع کارت الزامی است'
    isValid = false
  }
  
  // اعتبارسنجی مبلغ
  if (!form.amount || form.amount < 1000) {
    errors.amount = 'مبلغ باید حداقل 1,000 تومان باشد'
    isValid = false
  }
  
  // اعتبارسنجی تعداد
  if (!form.quantity || form.quantity < 1) {
    errors.quantity = 'تعداد باید حداقل 1 باشد'
    isValid = false
  }
  
  // اعتبارسنجی تاریخ انقضا
  if (!form.expiryDate) {
    errors.expiryDate = 'تاریخ انقضا الزامی است'
    isValid = false
  } else {
    const expiryDate = new Date(form.expiryDate)
    const today = new Date()
    if (expiryDate <= today) {
      errors.expiryDate = 'تاریخ انقضا باید بعد از امروز باشد'
      isValid = false
    }
  }
  
  // اعتبارسنجی ایمیل گیرنده
  if (form.recipientEmail && !isValidEmail(form.recipientEmail)) {
    errors.recipientEmail = 'ایمیل معتبر نیست'
    isValid = false
  }
  
  return isValid
}

const isValidEmail = (email: string) => {
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  return emailRegex.test(email)
}

const validateAndCreate = () => {
  if (validateForm()) {
    const cardData = {
      ...form,
      code: generatedCode.value,
      createdAt: new Date().toISOString()
    }
    emit('create-card', cardData)
  } else {
    // رفتن به اولین تب با خطا
    if (errors.type || errors.amount || errors.quantity || errors.expiryDate) {
      activeTab.value = 'basic'
    } else if (errors.recipientEmail) {
      activeTab.value = 'recipient'
    }
  }
}

const saveDraft = () => {
  const cardData = {
    ...form,
    code: generatedCode.value,
    status: 'draft',
    createdAt: new Date().toISOString()
  }
  emit('save-draft', cardData)
}

const resetForm = () => {
  if (confirm('آیا از پاک کردن فرم اطمینان دارید؟')) {
    Object.keys(form).forEach(key => {
      if (Array.isArray(form[key])) {
        form[key] = []
      } else if (typeof form[key] === 'number') {
        form[key] = null
      } else if (typeof form[key] === 'boolean') {
        form[key] = key === 'isActive'
      } else {
        form[key] = key === 'codePrefix' ? 'GC' : ''
      }
    })
    activeTab.value = 'basic'
  }
}

const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

const getFontFamily = (font: string) => {
  const fontMap = {
    'IRANSans': 'IRANSans, sans-serif',
    'Vazir': 'Vazir, sans-serif',
    'Shabnam': 'Shabnam, sans-serif',
    'Yekan': 'Yekan, sans-serif'
  }
  return fontMap[font] || 'IRANSans, sans-serif'
}

// Lifecycle
onMounted(() => {
  if (props.initialData) {
    Object.assign(form, props.initialData)
  }
  
  // تنظیم تاریخ انقضا پیش‌فرض (30 روز آینده)
  if (!form.expiryDate) {
    const defaultExpiry = new Date()
    defaultExpiry.setDate(defaultExpiry.getDate() + 30)
    form.expiryDate = defaultExpiry.toISOString().split('T')[0]
  }
})
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 

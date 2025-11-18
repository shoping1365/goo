<template>
  <div class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
    <div class="relative top-20 mx-auto p-5 border w-11/12 max-w-2xl shadow-lg rounded-md bg-white">
      <!-- هدر مودال -->
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-lg font-semibold text-gray-900">تعریف مناسبت جدید</h3>
        <button
          @click="$emit('close')"
          class="text-gray-400 hover:text-gray-600"
        >
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
        </button>
      </div>

      <!-- فرم تعریف مناسبت -->
      <form @submit.prevent="createOccasion" class="space-y-6">
        <!-- نام مناسبت -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            نام مناسبت <span class="text-red-500">*</span>
          </label>
          <input
            v-model="form.name"
            type="text"
            required
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="مثال: تخفیف تولد، تخفیف عروسی، تخفیف سال نو"
          />
        </div>

        <!-- نوع مناسبت -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            نوع مناسبت <span class="text-red-500">*</span>
          </label>
          <select
            v-model="form.type"
            required
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          >
            <option value="">انتخاب کنید</option>
            <option value="birthday">تولد</option>
            <option value="wedding">عروسی</option>
            <option value="new_year">سال نو</option>
            <option value="religious">مذهبی</option>
            <option value="national">ملی</option>
            <option value="seasonal">فصلی</option>
            <option value="custom">سفارشی</option>
          </select>
        </div>

        <!-- تاریخ‌های مناسبت -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              تاریخ شروع <span class="text-red-500">*</span>
            </label>
            <input
              v-model="form.startDate"
              type="date"
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              تاریخ پایان <span class="text-red-500">*</span>
            </label>
            <input
              v-model="form.endDate"
              type="date"
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            />
          </div>
        </div>

        <!-- تخفیف مناسبت -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            تخفیف مناسبت (%) <span class="text-red-500">*</span>
          </label>
          <input
            v-model="form.discount"
            type="number"
            required
            min="0"
            max="100"
            step="0.1"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="مثال: 15"
          />
          <p class="text-xs text-gray-500 mt-1">
            درصد تخفیف اعمال شده در این مناسبت
          </p>
        </div>

        <!-- کارمزد ویژه -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            کارمزد ویژه (%)
          </label>
          <input
            v-model="form.specialCommission"
            type="number"
            min="0"
            max="100"
            step="0.1"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="مثال: 2"
          />
          <p class="text-xs text-gray-500 mt-1">
            کارمزد ویژه اعمال شده در این مناسبت (اختیاری)
          </p>
        </div>

        <!-- محدودیت‌های مبلغ -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              حداقل مبلغ خرید (تومان)
            </label>
            <input
              v-model="form.minAmount"
              type="number"
              min="0"
              step="1000"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="مثال: 500000"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              حداکثر مبلغ تخفیف (تومان)
            </label>
            <input
              v-model="form.maxDiscountAmount"
              type="number"
              min="0"
              step="1000"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="مثال: 200000"
            />
          </div>
        </div>

        <!-- محدودیت‌های کاربر -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            محدودیت‌های کاربر
          </label>
          <div class="space-y-2">
            <label class="flex items-center">
              <input
                v-model="form.isFirstTimeOnly"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
              />
              <span class="mr-2 text-sm text-gray-700">فقط برای اولین خرید</span>
            </label>
            <label class="flex items-center">
              <input
                v-model="form.isNewUserOnly"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
              />
              <span class="mr-2 text-sm text-gray-700">فقط برای کاربران جدید</span>
            </label>
            <label class="flex items-center">
              <input
                v-model="form.isVipOnly"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
              />
              <span class="mr-2 text-sm text-gray-700">فقط برای کاربران VIP</span>
            </label>
            <label class="flex items-center">
              <input
                v-model="form.isAgeSpecific"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
              />
              <span class="mr-2 text-sm text-gray-700">محدودیت سنی</span>
            </label>
          </div>
        </div>

        <!-- محدودیت سنی -->
        <div v-if="form.isAgeSpecific" class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              حداقل سن
            </label>
            <input
              v-model="form.minAge"
              type="number"
              min="0"
              max="120"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="مثال: 18"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              حداکثر سن
            </label>
            <input
              v-model="form.maxAge"
              type="number"
              min="0"
              max="120"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="مثال: 65"
            />
          </div>
        </div>

        <!-- تعداد استفاده -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              تعداد استفاده
            </label>
            <input
              v-model="form.usageLimit"
              type="number"
              min="0"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="نامحدود = 0"
            />
            <p class="text-xs text-gray-500 mt-1">
              تعداد دفعاتی که این مناسبت قابل استفاده است
            </p>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              استفاده شده
            </label>
            <input
              v-model="form.usedCount"
              type="number"
              min="0"
              readonly
              class="w-full px-3 py-2 border border-gray-300 rounded-lg bg-gray-50"
              placeholder="0"
            />
          </div>
        </div>

        <!-- کد مناسبت -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            کد مناسبت
          </label>
          <div class="flex">
            <input
              v-model="form.code"
              type="text"
              class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="مثال: BIRTHDAY2024"
            />
            <button
              type="button"
              @click="generateCode"
              class="px-4 py-2 bg-gray-600 text-white text-sm font-medium rounded-l-lg hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
            >
              تولید
            </button>
          </div>
          <p class="text-xs text-gray-500 mt-1">
            کد اختیاری برای استفاده از مناسبت
          </p>
        </div>

        <!-- توضیحات -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            توضیحات
          </label>
          <textarea
            v-model="form.description"
            rows="3"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="توضیحات مربوط به این مناسبت..."
          ></textarea>
        </div>

        <!-- وضعیت -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            وضعیت <span class="text-red-500">*</span>
          </label>
          <select
            v-model="form.status"
            required
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          >
            <option value="">انتخاب کنید</option>
            <option value="active">فعال</option>
            <option value="inactive">غیرفعال</option>
            <option value="upcoming">آینده</option>
            <option value="expired">منقضی شده</option>
          </select>
        </div>

        <!-- نمایش محاسبه مناسبت -->
        <div v-if="form.discount && form.minAmount" class="bg-green-50 p-6 rounded-lg">
          <h4 class="text-sm font-medium text-green-900 mb-2">نمونه محاسبه مناسبت</h4>
          <div class="space-y-2 text-sm">
            <div class="flex justify-between">
              <span class="text-green-600">مبلغ خرید:</span>
              <span class="font-medium text-green-900">{{ formatCurrency(sampleAmount) }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-green-600">تخفیف مناسبت:</span>
              <span class="font-medium text-green-900">{{ formatCurrency(sampleDiscount) }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-green-600">کارمزد ویژه:</span>
              <span class="font-medium text-green-900">{{ formatCurrency(sampleCommission) }}</span>
            </div>
            <div class="flex justify-between border-t border-green-200 pt-2">
              <span class="text-green-600 font-medium">مبلغ نهایی:</span>
              <span class="font-bold text-green-900">{{ formatCurrency(sampleFinalAmount) }}</span>
            </div>
          </div>
        </div>

        <!-- دکمه‌های عملیات -->
        <div class="flex justify-end space-x-3 space-x-reverse">
          <button
            type="button"
            @click="$emit('close')"
            class="px-4 py-2 bg-gray-300 text-gray-700 text-sm font-medium rounded-lg hover:bg-gray-400 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
          >
            انصراف
          </button>
          <button
            type="submit"
            class="px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
          >
            تعریف مناسبت
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
// تعریف props و emits
const emit = defineEmits(['close', 'created'])

// تعریف متغیرهای reactive
const form = ref({
  name: '',
  type: '',
  startDate: '',
  endDate: '',
  discount: '',
  specialCommission: '',
  minAmount: '',
  maxDiscountAmount: '',
  isFirstTimeOnly: false,
  isNewUserOnly: false,
  isVipOnly: false,
  isAgeSpecific: false,
  minAge: '',
  maxAge: '',
  usageLimit: '',
  usedCount: 0,
  code: '',
  description: '',
  status: 'active'
})

// محاسبات نمونه
const sampleAmount = computed(() => {
  return form.value.minAmount ? parseFloat(form.value.minAmount) : 1000000
})

const sampleDiscount = computed(() => {
  if (!form.value.discount) return 0
  const amount = sampleAmount.value
  const discount = (amount * parseFloat(form.value.discount)) / 100
  const maxDiscount = form.value.maxDiscountAmount ? parseFloat(form.value.maxDiscountAmount) : Infinity
  return Math.min(discount, maxDiscount)
})

const sampleCommission = computed(() => {
  if (!form.value.specialCommission) return 0
  const amount = sampleAmount.value
  return (amount * parseFloat(form.value.specialCommission)) / 100
})

const sampleFinalAmount = computed(() => {
  return sampleAmount.value - sampleDiscount.value + sampleCommission.value
})

// توابع کمکی
const generateCode = () => {
  const chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789'
  let result = ''
  for (let i = 0; i < 8; i++) {
    result += chars.charAt(Math.floor(Math.random() * chars.length))
  }
  form.value.code = result
}

// تابع ایجاد مناسبت
const createOccasion = () => {
  // اعتبارسنجی فرم
  if (!form.value.name || !form.value.type || !form.value.startDate || !form.value.endDate || !form.value.discount || !form.value.status) {
    alert('لطفاً تمام فیلدهای اجباری را پر کنید')
    return
  }

  // اعتبارسنجی تخفیف
  if (form.value.discount < 0 || form.value.discount > 100) {
    alert('درصد تخفیف باید بین 0 تا 100 باشد')
    return
  }

  // اعتبارسنجی کارمزد ویژه
  if (form.value.specialCommission && (form.value.specialCommission < 0 || form.value.specialCommission > 100)) {
    alert('کارمزد ویژه باید بین 0 تا 100 درصد باشد')
    return
  }

  // اعتبارسنجی تاریخ‌ها
  if (new Date(form.value.startDate) >= new Date(form.value.endDate)) {
    alert('تاریخ شروع باید قبل از تاریخ پایان باشد')
    return
  }

  // اعتبارسنجی محدودیت سنی
  if (form.value.isAgeSpecific) {
    if (form.value.minAge && form.value.maxAge && parseFloat(form.value.minAge) >= parseFloat(form.value.maxAge)) {
      alert('حداقل سن باید کمتر از حداکثر سن باشد')
      return
    }
  }

  // اعتبارسنجی حداقل مبلغ
  if (form.value.minAmount && form.value.minAmount <= 0) {
    alert('حداقل مبلغ خرید باید بیشتر از صفر باشد')
    return
  }

  // اعتبارسنجی حداکثر تخفیف
  if (form.value.maxDiscountAmount && form.value.maxDiscountAmount <= 0) {
    alert('حداکثر مبلغ تخفیف باید بیشتر از صفر باشد')
    return
  }

  // ایجاد مناسبت جدید
  const newOccasion = {
    name: form.value.name,
    type: form.value.type,
    startDate: form.value.startDate,
    endDate: form.value.endDate,
    discount: parseFloat(form.value.discount),
    specialCommission: form.value.specialCommission ? parseFloat(form.value.specialCommission) : 0,
    minAmount: form.value.minAmount ? parseFloat(form.value.minAmount) : 0,
    maxDiscountAmount: form.value.maxDiscountAmount ? parseFloat(form.value.maxDiscountAmount) : null,
    isFirstTimeOnly: form.value.isFirstTimeOnly,
    isNewUserOnly: form.value.isNewUserOnly,
    isVipOnly: form.value.isVipOnly,
    isAgeSpecific: form.value.isAgeSpecific,
    minAge: form.value.minAge ? parseInt(form.value.minAge) : null,
    maxAge: form.value.maxAge ? parseInt(form.value.maxAge) : null,
    usageLimit: form.value.usageLimit ? parseInt(form.value.usageLimit) : 0,
    usedCount: 0,
    code: form.value.code || null,
    description: form.value.description || '',
    status: form.value.status
  }

  // ارسال به کامپوننت والد
  emit('created', newOccasion)
}

// تابع فرمت کردن مبلغ
const formatCurrency = (amount) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 

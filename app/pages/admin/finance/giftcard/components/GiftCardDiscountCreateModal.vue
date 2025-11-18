<template>
  <div class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
    <div class="relative top-20 mx-auto p-5 border w-11/12 max-w-2xl shadow-lg rounded-md bg-white">
      <!-- هدر مودال -->
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-lg font-semibold text-gray-900">تعریف تخفیف جدید</h3>
        <button
          @click="$emit('close')"
          class="text-gray-400 hover:text-gray-600"
        >
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
        </button>
      </div>

      <!-- فرم تعریف تخفیف -->
      <form @submit.prevent="createDiscount" class="space-y-6">
        <!-- نام تخفیف -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            نام تخفیف <span class="text-red-500">*</span>
          </label>
          <input
            v-model="form.name"
            type="text"
            required
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="مثال: تخفیف ویژه تولد، تخفیف عروسی"
          />
        </div>

        <!-- نوع تخفیف -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            نوع تخفیف <span class="text-red-500">*</span>
          </label>
          <select
            v-model="form.type"
            required
            @change="onDiscountTypeChange"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          >
            <option value="">انتخاب کنید</option>
            <option value="percentage">درصدی</option>
            <option value="fixed">مبلغ ثابت</option>
            <option value="buy_one_get_one">یک به یک</option>
            <option value="free_shipping">ارسال رایگان</option>
            <option value="cashback">کش‌بک</option>
          </select>
        </div>

        <!-- مقدار تخفیف -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            مقدار تخفیف <span class="text-red-500">*</span>
          </label>
          <div class="flex">
            <input
              v-model="form.value"
              :type="form.type === 'percentage' ? 'number' : 'number'"
              required
              :min="form.type === 'percentage' ? 0 : 0"
              :max="form.type === 'percentage' ? 100 : null"
              :step="form.type === 'percentage' ? 0.1 : 1000"
              class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              :placeholder="getDiscountPlaceholder()"
            />
            <span class="inline-flex items-center px-3 py-2 border border-r-0 border-gray-300 bg-gray-50 text-gray-500 text-sm rounded-l-lg">
              {{ getDiscountUnit() }}
            </span>
          </div>
          <p class="text-xs text-gray-500 mt-1">
            {{ getDiscountDescription() }}
          </p>
        </div>

        <!-- حداقل مبلغ خرید -->
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
          <p class="text-xs text-gray-500 mt-1">
            تخفیف فقط برای خریدهای بالاتر از این مبلغ اعمال می‌شود
          </p>
        </div>

        <!-- حداکثر مبلغ تخفیف -->
        <div v-if="form.type === 'percentage'">
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
          <p class="text-xs text-gray-500 mt-1">
            حداکثر مبلغی که می‌توان از تخفیف استفاده کرد
          </p>
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
              تعداد دفعاتی که این تخفیف قابل استفاده است
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
          </div>
        </div>

        <!-- تاریخ‌های اعتبار -->
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
              تاریخ انقضا <span class="text-red-500">*</span>
            </label>
            <input
              v-model="form.expiresAt"
              type="date"
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            />
          </div>
        </div>

        <!-- کد تخفیف -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            کد تخفیف
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
            کد اختیاری برای استفاده از تخفیف
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
            placeholder="توضیحات مربوط به این تخفیف..."
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
            <option value="scheduled">برنامه‌ریزی شده</option>
          </select>
        </div>

        <!-- نمایش محاسبه تخفیف -->
        <div v-if="form.value && form.minAmount" class="bg-green-50 p-6 rounded-lg">
          <h4 class="text-sm font-medium text-green-900 mb-2">نمونه محاسبه تخفیف</h4>
          <div class="space-y-2 text-sm">
            <div class="flex justify-between">
              <span class="text-green-600">مبلغ خرید:</span>
              <span class="font-medium text-green-900">{{ formatCurrency(sampleAmount) }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-green-600">تخفیف:</span>
              <span class="font-medium text-green-900">{{ formatCurrency(sampleDiscount) }}</span>
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
            تعریف تخفیف
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
  value: '',
  minAmount: '',
  maxDiscountAmount: '',
  usageLimit: '',
  usedCount: 0,
  isFirstTimeOnly: false,
  isNewUserOnly: false,
  isVipOnly: false,
  startDate: '',
  expiresAt: '',
  code: '',
  description: '',
  status: 'active'
})

// محاسبات نمونه
const sampleAmount = computed(() => {
  return form.value.minAmount ? parseFloat(form.value.minAmount) : 1000000
})

const sampleDiscount = computed(() => {
  if (!form.value.value || !form.value.type) return 0
  
  const amount = sampleAmount.value
  
  if (form.value.type === 'percentage') {
    const discount = (amount * parseFloat(form.value.value)) / 100
    const maxDiscount = form.value.maxDiscountAmount ? parseFloat(form.value.maxDiscountAmount) : Infinity
    return Math.min(discount, maxDiscount)
  } else if (form.value.type === 'fixed') {
    return parseFloat(form.value.value)
  } else if (form.value.type === 'buy_one_get_one') {
    return amount / 2
  } else if (form.value.type === 'free_shipping') {
    return 50000 // فرض بر ارسال رایگان
  } else if (form.value.type === 'cashback') {
    return (amount * parseFloat(form.value.value)) / 100
  }
  
  return 0
})

const sampleFinalAmount = computed(() => {
  return sampleAmount.value - sampleDiscount.value
})

// توابع کمکی
const getDiscountPlaceholder = () => {
  const placeholders = {
    percentage: 'مثال: 15',
    fixed: 'مثال: 200000',
    buy_one_get_one: 'مثال: 1',
    free_shipping: 'مثال: 1',
    cashback: 'مثال: 10'
  }
  return placeholders[form.value.type] || ''
}

const getDiscountUnit = () => {
  const units = {
    percentage: '%',
    fixed: 'تومان',
    buy_one_get_one: 'عدد',
    free_shipping: 'عدد',
    cashback: '%'
  }
  return units[form.value.type] || ''
}

const getDiscountDescription = () => {
  const descriptions = {
    percentage: 'درصد تخفیف از کل مبلغ خرید',
    fixed: 'مبلغ ثابت تخفیف',
    buy_one_get_one: 'خرید یک عدد، یک عدد هدیه',
    free_shipping: 'ارسال رایگان',
    cashback: 'درصد کش‌بک به کیف پول'
  }
  return descriptions[form.value.type] || ''
}

const onDiscountTypeChange = () => {
  // پاک کردن مقدار قبلی هنگام تغییر نوع تخفیف
  form.value.value = ''
  form.value.maxDiscountAmount = ''
}

const generateCode = () => {
  const chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789'
  let result = ''
  for (let i = 0; i < 8; i++) {
    result += chars.charAt(Math.floor(Math.random() * chars.length))
  }
  form.value.code = result
}

// تابع ایجاد تخفیف
const createDiscount = () => {
  // اعتبارسنجی فرم
  if (!form.value.name || !form.value.type || !form.value.value || !form.value.startDate || !form.value.expiresAt || !form.value.status) {
    alert('لطفاً تمام فیلدهای اجباری را پر کنید')
    return
  }

  // اعتبارسنجی مقدار تخفیف
  if (form.value.type === 'percentage' && (form.value.value < 0 || form.value.value > 100)) {
    alert('درصد تخفیف باید بین 0 تا 100 باشد')
    return
  }

  if (form.value.type === 'fixed' && form.value.value <= 0) {
    alert('مبلغ تخفیف باید بیشتر از صفر باشد')
    return
  }

  // اعتبارسنجی تاریخ‌ها
  if (new Date(form.value.startDate) >= new Date(form.value.expiresAt)) {
    alert('تاریخ شروع باید قبل از تاریخ انقضا باشد')
    return
  }

  // اعتبارسنجی حداقل مبلغ
  if (form.value.minAmount && form.value.minAmount <= 0) {
    alert('حداقل مبلغ خرید باید بیشتر از صفر باشد')
    return
  }

  // اعتبارسنجی حداکثر تخفیف
  if (form.value.type === 'percentage' && form.value.maxDiscountAmount && form.value.maxDiscountAmount <= 0) {
    alert('حداکثر مبلغ تخفیف باید بیشتر از صفر باشد')
    return
  }

  // ایجاد تخفیف جدید
  const newDiscount = {
    name: form.value.name,
    type: form.value.type,
    value: parseFloat(form.value.value),
    minAmount: form.value.minAmount ? parseFloat(form.value.minAmount) : 0,
    maxDiscountAmount: form.value.maxDiscountAmount ? parseFloat(form.value.maxDiscountAmount) : null,
    usageLimit: form.value.usageLimit ? parseInt(form.value.usageLimit) : 0,
    usedCount: 0,
    isFirstTimeOnly: form.value.isFirstTimeOnly,
    isNewUserOnly: form.value.isNewUserOnly,
    isVipOnly: form.value.isVipOnly,
    startDate: form.value.startDate,
    expiresAt: form.value.expiresAt,
    code: form.value.code || null,
    description: form.value.description || '',
    status: form.value.status
  }

  // ارسال به کامپوننت والد
  emit('created', newDiscount)
}

// تابع فرمت کردن مبلغ
const formatCurrency = (amount) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 

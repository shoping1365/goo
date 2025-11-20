<template>
  <div class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
    <div class="relative top-20 mx-auto p-5 border w-11/12 max-w-2xl shadow-lg rounded-md bg-white">
      <!-- هدر مودال -->
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-lg font-semibold text-gray-900">تعریف قیمت جدید</h3>
        <button
          class="text-gray-400 hover:text-gray-600"
          @click="$emit('close')"
        >
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
        </button>
      </div>

      <!-- فرم تعریف قیمت -->
      <form class="space-y-6" @submit.prevent="createPricing">
        <!-- نام قیمت -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            نام قیمت <span class="text-red-500">*</span>
          </label>
          <input
            v-model="form.name"
            type="text"
            required
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="مثال: کارت تولد، کارت عروسی"
          />
        </div>

        <!-- نوع کارت -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            نوع کارت <span class="text-red-500">*</span>
          </label>
          <select
            v-model="form.type"
            required
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          >
            <option value="">انتخاب کنید</option>
            <option value="fixed">مبلغ ثابت</option>
            <option value="variable">مبلغ متغیر</option>
            <option value="themed">موضوعی</option>
            <option value="corporate">شرکتی</option>
          </select>
        </div>

        <!-- مبلغ پایه -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            مبلغ پایه (تومان) <span class="text-red-500">*</span>
          </label>
          <input
            v-model="form.baseAmount"
            type="number"
            required
            min="0"
            step="1000"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="مثال: 500000"
          />
        </div>

        <!-- کارمزد -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            کارمزد (%) <span class="text-red-500">*</span>
          </label>
          <input
            v-model="form.commission"
            type="number"
            required
            min="0"
            max="100"
            step="0.1"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="مثال: 5"
          />
          <p class="text-xs text-gray-500 mt-1">
            درصد کارمزد از مبلغ پایه
          </p>
        </div>

        <!-- نمایش قیمت نهایی -->
        <div v-if="form.baseAmount && form.commission" class="bg-blue-50 p-6 rounded-lg">
          <h4 class="text-sm font-medium text-blue-900 mb-2">محاسبه قیمت نهایی</h4>
          <div class="grid grid-cols-2 gap-6 text-sm">
            <div>
              <span class="text-blue-600">مبلغ پایه:</span>
              <span class="font-medium text-blue-900 mr-2">{{ formatCurrency(form.baseAmount) }}</span>
            </div>
            <div>
              <span class="text-blue-600">کارمزد:</span>
              <span class="font-medium text-blue-900 mr-2">{{ formatCurrency(commissionAmount) }}</span>
            </div>
            <div class="col-span-2 border-t border-blue-200 pt-2">
              <span class="text-blue-600 font-medium">قیمت نهایی:</span>
              <span class="font-bold text-blue-900 text-lg mr-2">{{ formatCurrency(finalPrice) }}</span>
            </div>
          </div>
        </div>

        <!-- محدوده مبلغ (برای کارت‌های متغیر) -->
        <div v-if="form.type === 'variable'" class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              حداقل مبلغ (تومان)
            </label>
            <input
              v-model="form.minAmount"
              type="number"
              min="0"
              step="1000"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="مثال: 100000"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              حداکثر مبلغ (تومان)
            </label>
            <input
              v-model="form.maxAmount"
              type="number"
              min="0"
              step="1000"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="مثال: 5000000"
            />
          </div>
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
            placeholder="توضیحات مربوط به این قیمت..."
          ></textarea>
        </div>

        <!-- شرایط خاص -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            شرایط خاص
          </label>
          <div class="space-y-2">
            <label class="flex items-center">
              <input
                v-model="form.hasMinimumOrder"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
              />
              <span class="mr-2 text-sm text-gray-700">حداقل مبلغ سفارش</span>
            </label>
            <label class="flex items-center">
              <input
                v-model="form.hasMaximumOrder"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
              />
              <span class="mr-2 text-sm text-gray-700">حداکثر مبلغ سفارش</span>
            </label>
            <label class="flex items-center">
              <input
                v-model="form.isLimitedTime"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
              />
              <span class="mr-2 text-sm text-gray-700">محدودیت زمانی</span>
            </label>
          </div>
        </div>

        <!-- محدودیت‌های زمانی -->
        <div v-if="form.isLimitedTime" class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              تاریخ شروع
            </label>
            <input
              v-model="form.startDate"
              type="date"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              تاریخ پایان
            </label>
            <input
              v-model="form.endDate"
              type="date"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            />
          </div>
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
            <option value="pending">در انتظار</option>
          </select>
        </div>

        <!-- دکمه‌های عملیات -->
        <div class="flex justify-end space-x-3 space-x-reverse">
          <button
            type="button"
            class="px-4 py-2 bg-gray-300 text-gray-700 text-sm font-medium rounded-lg hover:bg-gray-400 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
            @click="$emit('close')"
          >
            انصراف
          </button>
          <button
            type="submit"
            class="px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
          >
            تعریف قیمت
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
  baseAmount: '',
  commission: '',
  minAmount: '',
  maxAmount: '',
  description: '',
  hasMinimumOrder: false,
  hasMaximumOrder: false,
  isLimitedTime: false,
  startDate: '',
  endDate: '',
  status: 'active'
})

// محاسبه کارمزد
const commissionAmount = computed(() => {
  if (!form.value.baseAmount || !form.value.commission) return 0
  return (parseFloat(form.value.baseAmount) * parseFloat(form.value.commission)) / 100
})

// محاسبه قیمت نهایی
const finalPrice = computed(() => {
  if (!form.value.baseAmount || !form.value.commission) return 0
  return parseFloat(form.value.baseAmount) + commissionAmount.value
})

// تابع ایجاد قیمت
const createPricing = () => {
  // اعتبارسنجی فرم
  if (!form.value.name || !form.value.type || !form.value.baseAmount || !form.value.commission || !form.value.status) {
    alert('لطفاً تمام فیلدهای اجباری را پر کنید')
    return
  }

  // اعتبارسنجی مبلغ
  if (form.value.baseAmount <= 0) {
    alert('مبلغ پایه باید بیشتر از صفر باشد')
    return
  }

  // اعتبارسنجی کارمزد
  if (form.value.commission < 0 || form.value.commission > 100) {
    alert('کارمزد باید بین 0 تا 100 درصد باشد')
    return
  }

  // اعتبارسنجی محدوده مبلغ برای کارت‌های متغیر
  if (form.value.type === 'variable') {
    if (form.value.minAmount && form.value.maxAmount && parseFloat(form.value.minAmount) >= parseFloat(form.value.maxAmount)) {
      alert('حداقل مبلغ باید کمتر از حداکثر مبلغ باشد')
      return
    }
  }

  // اعتبارسنجی تاریخ‌ها
  if (form.value.isLimitedTime) {
    if (!form.value.startDate || !form.value.endDate) {
      alert('برای محدودیت زمانی، تاریخ شروع و پایان الزامی است')
      return
    }
    if (new Date(form.value.startDate) >= new Date(form.value.endDate)) {
      alert('تاریخ شروع باید قبل از تاریخ پایان باشد')
      return
    }
  }

  // ایجاد قیمت جدید
  const newPricing = {
    name: form.value.name,
    type: form.value.type,
    baseAmount: parseFloat(form.value.baseAmount),
    commission: parseFloat(form.value.commission),
    minAmount: form.value.minAmount ? parseFloat(form.value.minAmount) : null,
    maxAmount: form.value.maxAmount ? parseFloat(form.value.maxAmount) : null,
    description: form.value.description || '',
    hasMinimumOrder: form.value.hasMinimumOrder,
    hasMaximumOrder: form.value.hasMaximumOrder,
    isLimitedTime: form.value.isLimitedTime,
    startDate: form.value.startDate || null,
    endDate: form.value.endDate || null,
    status: form.value.status
  }

  // ارسال به کامپوننت والد
  emit('created', newPricing)
}

// تابع فرمت کردن مبلغ
const formatCurrency = (amount) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 

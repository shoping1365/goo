<template>
  <div class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
    <div class="relative top-20 mx-auto p-5 border w-11/12 max-w-2xl shadow-lg rounded-md bg-white">
      <!-- هدر مودال -->
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-lg font-semibold text-gray-900">تعریف کارمزد جدید</h3>
        <button
          class="text-gray-400 hover:text-gray-600"
          @click="$emit('close')"
        >
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
        </button>
      </div>

      <!-- فرم تعریف کارمزد -->
      <form class="space-y-6" @submit.prevent="createCommission">
        <!-- نام کارمزد -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            نام کارمزد <span class="text-red-500">*</span>
          </label>
          <input
            v-model="form.name"
            type="text"
            required
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="مثال: کارمزد استاندارد، کارمزد ویژه"
          />
        </div>

        <!-- نوع کارمزد -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            نوع کارمزد <span class="text-red-500">*</span>
          </label>
          <select
            v-model="form.type"
            required
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            @change="onCommissionTypeChange"
          >
            <option value="">انتخاب کنید</option>
            <option value="percentage">درصدی</option>
            <option value="fixed">مبلغ ثابت</option>
            <option value="tiered">پلکانی</option>
            <option value="progressive">تصاعدی</option>
          </select>
        </div>

        <!-- نرخ کارمزد -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            نرخ کارمزد <span class="text-red-500">*</span>
          </label>
          <div class="flex">
            <input
              v-model="form.rate"
              type="number"
              required
              :min="form.type === 'percentage' ? 0 : 0"
              :max="form.type === 'percentage' ? 100 : null"
              :step="form.type === 'percentage' ? 0.1 : 1000"
              class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              :placeholder="getCommissionPlaceholder()"
            />
            <span class="inline-flex items-center px-3 py-2 border border-r-0 border-gray-300 bg-gray-50 text-gray-500 text-sm rounded-l-lg">
              {{ getCommissionUnit() }}
            </span>
          </div>
          <p class="text-xs text-gray-500 mt-1">
            {{ getCommissionDescription() }}
          </p>
        </div>

        <!-- محدوده مبلغ -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
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
              placeholder="مثال: 0"
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
              placeholder="مثال: 10000000"
            />
          </div>
        </div>

        <!-- کارمزد پلکانی -->
        <div v-if="form.type === 'tiered'" class="space-y-4">
          <div class="flex items-center justify-between">
            <h4 class="text-sm font-medium text-gray-700">سطوح کارمزد پلکانی</h4>
            <button
              type="button"
              class="px-3 py-1 bg-blue-600 text-white text-xs font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
              @click="addTier"
            >
              افزودن سطح
            </button>
          </div>
          
          <div v-for="(tier, index) in form.tiers" :key="index" class="grid grid-cols-3 gap-2 p-3 bg-gray-50 rounded-lg">
            <div>
              <label class="block text-xs font-medium text-gray-600 mb-1">از مبلغ</label>
              <input
                v-model="tier.from"
                type="number"
                min="0"
                step="1000"
                class="w-full px-2 py-1 text-xs border border-gray-300 rounded focus:outline-none focus:ring-1 focus:ring-blue-500"
                placeholder="0"
              />
            </div>
            <div>
              <label class="block text-xs font-medium text-gray-600 mb-1">تا مبلغ</label>
              <input
                v-model="tier.to"
                type="number"
                min="0"
                step="1000"
                class="w-full px-2 py-1 text-xs border border-gray-300 rounded focus:outline-none focus:ring-1 focus:ring-blue-500"
                placeholder="1000000"
              />
            </div>
            <div class="flex items-end gap-1">
              <div class="flex-1">
                <label class="block text-xs font-medium text-gray-600 mb-1">نرخ (%)</label>
                <input
                  v-model="tier.rate"
                  type="number"
                  min="0"
                  max="100"
                  step="0.1"
                  class="w-full px-2 py-1 text-xs border border-gray-300 rounded focus:outline-none focus:ring-1 focus:ring-blue-500"
                  placeholder="5"
                />
              </div>
              <button
                type="button"
                class="px-2 py-1 bg-red-600 text-white text-xs rounded hover:bg-red-700"
                @click="removeTier(index)"
              >
                حذف
              </button>
            </div>
          </div>
        </div>

        <!-- کارمزد تصاعدی -->
        <div v-if="form.type === 'progressive'" class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              نرخ پایه (%)
            </label>
            <input
              v-model="form.baseRate"
              type="number"
              min="0"
              max="100"
              step="0.1"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="مثال: 3"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              ضریب تصاعدی
            </label>
            <input
              v-model="form.progressiveFactor"
              type="number"
              min="0"
              step="0.1"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="مثال: 0.5"
            />
            <p class="text-xs text-gray-500 mt-1">
              برای هر 1 میلیون تومان افزایش، این مقدار به نرخ پایه اضافه می‌شود
            </p>
          </div>
        </div>

        <!-- شرایط خاص -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            شرایط خاص
          </label>
          <div class="space-y-2">
            <label class="flex items-center">
              <input
                v-model="form.hasMinimumCommission"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
              />
              <span class="mr-2 text-sm text-gray-700">حداقل کارمزد</span>
            </label>
            <label class="flex items-center">
              <input
                v-model="form.hasMaximumCommission"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
              />
              <span class="mr-2 text-sm text-gray-700">حداکثر کارمزد</span>
            </label>
            <label class="flex items-center">
              <input
                v-model="form.isHolidayRate"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
              />
              <span class="mr-2 text-sm text-gray-700">نرخ تعطیلات</span>
            </label>
          </div>
        </div>

        <!-- محدودیت‌های کارمزد -->
        <div v-if="form.hasMinimumCommission || form.hasMaximumCommission" class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div v-if="form.hasMinimumCommission">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              حداقل کارمزد (تومان)
            </label>
            <input
              v-model="form.minCommission"
              type="number"
              min="0"
              step="1000"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="مثال: 10000"
            />
          </div>
          <div v-if="form.hasMaximumCommission">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              حداکثر کارمزد (تومان)
            </label>
            <input
              v-model="form.maxCommission"
              type="number"
              min="0"
              step="1000"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="مثال: 500000"
            />
          </div>
        </div>

        <!-- تاریخ‌های اعتبار -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
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
              تاریخ انقضا
            </label>
            <input
              v-model="form.endDate"
              type="date"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
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
            placeholder="توضیحات مربوط به این کارمزد..."
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

        <!-- نمایش محاسبه کارمزد -->
        <div v-if="form.rate && form.minAmount" class="bg-blue-50 p-6 rounded-lg">
          <h4 class="text-sm font-medium text-blue-900 mb-2">نمونه محاسبه کارمزد</h4>
          <div class="space-y-2 text-sm">
            <div class="flex justify-between">
              <span class="text-blue-600">مبلغ تراکنش:</span>
              <span class="font-medium text-blue-900">{{ formatCurrency(sampleAmount) }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-blue-600">کارمزد محاسبه شده:</span>
              <span class="font-medium text-blue-900">{{ formatCurrency(sampleCommission) }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-blue-600">نرخ کارمزد:</span>
              <span class="font-medium text-blue-900">{{ form.type === 'percentage' ? form.rate + '%' : formatCurrency(form.rate) }}</span>
            </div>
          </div>
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
            تعریف کارمزد
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
  rate: '',
  minAmount: '',
  maxAmount: '',
  baseRate: '',
  progressiveFactor: '',
  tiers: [],
  hasMinimumCommission: false,
  hasMaximumCommission: false,
  isHolidayRate: false,
  minCommission: '',
  maxCommission: '',
  startDate: '',
  endDate: '',
  description: '',
  status: 'active'
})

// محاسبات نمونه
const sampleAmount = computed(() => {
  return form.value.minAmount ? parseFloat(form.value.minAmount) : 1000000
})

const sampleCommission = computed(() => {
  if (!form.value.rate || !form.value.type) return 0
  
  const amount = sampleAmount.value
  
  if (form.value.type === 'percentage') {
    return (amount * parseFloat(form.value.rate)) / 100
  } else if (form.value.type === 'fixed') {
    return parseFloat(form.value.rate)
  } else if (form.value.type === 'tiered') {
    // محاسبه کارمزد پلکانی
    for (const tier of form.value.tiers) {
      if (amount >= parseFloat(tier.from) && amount <= parseFloat(tier.to)) {
        return (amount * parseFloat(tier.rate)) / 100
      }
    }
    return 0
  } else if (form.value.type === 'progressive') {
    // محاسبه کارمزد تصاعدی
    const baseRate = parseFloat(form.value.baseRate) || 0
    const factor = parseFloat(form.value.progressiveFactor) || 0
    const additionalRate = Math.floor(amount / 1000000) * factor
    return (amount * (baseRate + additionalRate)) / 100
  }
  
  return 0
})

// توابع کمکی
const getCommissionPlaceholder = () => {
  const placeholders = {
    percentage: 'مثال: 5',
    fixed: 'مثال: 50000',
    tiered: 'مثال: 5',
    progressive: 'مثال: 3'
  }
  return placeholders[form.value.type] || ''
}

const getCommissionUnit = () => {
  const units = {
    percentage: '%',
    fixed: 'تومان',
    tiered: '%',
    progressive: '%'
  }
  return units[form.value.type] || ''
}

const getCommissionDescription = () => {
  const descriptions = {
    percentage: 'درصد کارمزد از کل مبلغ تراکنش',
    fixed: 'مبلغ ثابت کارمزد',
    tiered: 'کارمزد پلکانی بر اساس محدوده مبلغ',
    progressive: 'کارمزد تصاعدی بر اساس مبلغ تراکنش'
  }
  return descriptions[form.value.type] || ''
}

const onCommissionTypeChange = () => {
  // پاک کردن مقادیر قبلی هنگام تغییر نوع کارمزد
  form.value.rate = ''
  form.value.baseRate = ''
  form.value.progressiveFactor = ''
  form.value.tiers = []
}

const addTier = () => {
  form.value.tiers.push({
    from: '',
    to: '',
    rate: ''
  })
}

const removeTier = (index) => {
  form.value.tiers.splice(index, 1)
}

// تابع ایجاد کارمزد
const createCommission = () => {
  // اعتبارسنجی فرم
  if (!form.value.name || !form.value.type || !form.value.status) {
    alert('لطفاً تمام فیلدهای اجباری را پر کنید')
    return
  }

  // اعتبارسنجی نرخ کارمزد
  if (form.value.type === 'percentage' && (!form.value.rate || form.value.rate < 0 || form.value.rate > 100)) {
    alert('نرخ کارمزد درصدی باید بین 0 تا 100 باشد')
    return
  }

  if (form.value.type === 'fixed' && (!form.value.rate || form.value.rate <= 0)) {
    alert('نرخ کارمزد ثابت باید بیشتر از صفر باشد')
    return
  }

  // اعتبارسنجی کارمزد پلکانی
  if (form.value.type === 'tiered') {
    if (form.value.tiers.length === 0) {
      alert('برای کارمزد پلکانی حداقل یک سطح باید تعریف شود')
      return
    }
    
    for (const tier of form.value.tiers) {
      if (!tier.from || !tier.to || !tier.rate) {
        alert('تمام فیلدهای سطوح کارمزد پلکانی باید پر شوند')
        return
      }
      if (parseFloat(tier.from) >= parseFloat(tier.to)) {
        alert('مبلغ شروع باید کمتر از مبلغ پایان باشد')
        return
      }
      if (tier.rate < 0 || tier.rate > 100) {
        alert('نرخ کارمزد باید بین 0 تا 100 درصد باشد')
        return
      }
    }
  }

  // اعتبارسنجی کارمزد تصاعدی
  if (form.value.type === 'progressive') {
    if (!form.value.baseRate || form.value.baseRate < 0) {
      alert('نرخ پایه باید بیشتر از صفر باشد')
      return
    }
    if (!form.value.progressiveFactor || form.value.progressiveFactor < 0) {
      alert('ضریب تصاعدی باید بیشتر از صفر باشد')
      return
    }
  }

  // اعتبارسنجی تاریخ‌ها
  if (form.value.startDate && form.value.endDate) {
    if (new Date(form.value.startDate) >= new Date(form.value.endDate)) {
      alert('تاریخ شروع باید قبل از تاریخ انقضا باشد')
      return
    }
  }

  // ایجاد کارمزد جدید
  const newCommission = {
    name: form.value.name,
    type: form.value.type,
    rate: form.value.type === 'percentage' || form.value.type === 'fixed' ? parseFloat(form.value.rate) : null,
    minAmount: form.value.minAmount ? parseFloat(form.value.minAmount) : 0,
    maxAmount: form.value.maxAmount ? parseFloat(form.value.maxAmount) : null,
    baseRate: form.value.baseRate ? parseFloat(form.value.baseRate) : null,
    progressiveFactor: form.value.progressiveFactor ? parseFloat(form.value.progressiveFactor) : null,
    tiers: form.value.tiers.map(tier => ({
      from: parseFloat(tier.from),
      to: parseFloat(tier.to),
      rate: parseFloat(tier.rate)
    })),
    hasMinimumCommission: form.value.hasMinimumCommission,
    hasMaximumCommission: form.value.hasMaximumCommission,
    isHolidayRate: form.value.isHolidayRate,
    minCommission: form.value.minCommission ? parseFloat(form.value.minCommission) : null,
    maxCommission: form.value.maxCommission ? parseFloat(form.value.maxCommission) : null,
    startDate: form.value.startDate || null,
    endDate: form.value.endDate || null,
    description: form.value.description || '',
    status: form.value.status
  }

  // ارسال به کامپوننت والد
  emit('created', newCommission)
}

// تابع فرمت کردن مبلغ
const formatCurrency = (amount) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 

<template>
  <div class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
    <div class="relative top-20 mx-auto p-5 border w-11/12 md:w-3/4 lg:w-1/2 shadow-lg rounded-md bg-white">
      <!-- هدر مودال -->
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-lg font-semibold text-gray-900">تعریف محدودیت جدید</h3>
        <button
          @click="$emit('close')"
          class="text-gray-400 hover:text-gray-600"
        >
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
        </button>
      </div>

      <!-- فرم تعریف محدودیت -->
      <form @submit.prevent="handleSubmit" class="space-y-6">
        <!-- اطلاعات اصلی -->
        <div class="bg-gray-50 p-6 rounded-lg">
          <h4 class="text-md font-semibold text-gray-900 mb-4">اطلاعات اصلی</h4>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                نام محدودیت <span class="text-red-500">*</span>
              </label>
              <input
                v-model="form.name"
                type="text"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="مثال: محدودیت روزانه"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                نوع محدودیت <span class="text-red-500">*</span>
              </label>
              <select
                v-model="form.type"
                required
                @change="handleTypeChange"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              >
                <option value="">انتخاب کنید</option>
                <option value="daily">روزانه</option>
                <option value="weekly">هفتگی</option>
                <option value="monthly">ماهانه</option>
                <option value="yearly">سالانه</option>
                <option value="transaction_count">تعداد تراکنش</option>
                <option value="amount_limit">محدودیت مبلغ</option>
                <option value="time_limit">محدودیت زمانی</option>
                <option value="category_limit">محدودیت دسته‌بندی</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                مقدار محدودیت <span class="text-red-500">*</span>
              </label>
              <input
                v-model="form.limit"
                type="number"
                required
                min="1"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                :placeholder="getLimitPlaceholder()"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                واحد <span class="text-red-500">*</span>
              </label>
              <select
                v-model="form.unit"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              >
                <option value="">انتخاب کنید</option>
                <option value="toman">تومان</option>
                <option value="rial">ریال</option>
                <option value="dollar">دلار</option>
                <option value="euro">یورو</option>
                <option value="count">تعداد</option>
                <option value="hour">ساعت</option>
                <option value="day">روز</option>
                <option value="week">هفته</option>
                <option value="month">ماه</option>
              </select>
            </div>
          </div>
        </div>

        <!-- دوره زمانی -->
        <div class="bg-gray-50 p-6 rounded-lg">
          <h4 class="text-md font-semibold text-gray-900 mb-4">دوره زمانی</h4>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                دوره <span class="text-red-500">*</span>
              </label>
              <select
                v-model="form.period"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              >
                <option value="">انتخاب کنید</option>
                <option value="day">روزانه</option>
                <option value="week">هفتگی</option>
                <option value="month">ماهانه</option>
                <option value="year">سالانه</option>
                <option value="custom">سفارشی</option>
              </select>
            </div>
            <div v-if="form.period === 'custom'">
              <label class="block text-sm font-medium text-gray-700 mb-2">
                تعداد روزهای سفارشی
              </label>
              <input
                v-model="form.customDays"
                type="number"
                min="1"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="تعداد روزها"
              />
            </div>
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
        </div>

        <!-- محدوده اعمال -->
        <div class="bg-gray-50 p-6 rounded-lg">
          <h4 class="text-md font-semibold text-gray-900 mb-4">محدوده اعمال</h4>
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                نوع کاربران هدف
              </label>
              <div class="space-y-2">
                <div class="flex items-center">
                  <input
                    v-model="form.targetUsers.all"
                    type="checkbox"
                    id="all-users"
                    class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                  />
                  <label for="all-users" class="mr-2 block text-sm text-gray-900">
                    همه کاربران
                  </label>
                </div>
                <div class="flex items-center">
                  <input
                    v-model="form.targetUsers.new"
                    type="checkbox"
                    id="new-users"
                    class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                  />
                  <label for="new-users" class="mr-2 block text-sm text-gray-900">
                    کاربران جدید
                  </label>
                </div>
                <div class="flex items-center">
                  <input
                    v-model="form.targetUsers.premium"
                    type="checkbox"
                    id="premium-users"
                    class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                  />
                  <label for="premium-users" class="mr-2 block text-sm text-gray-900">
                    کاربران پریمیوم
                  </label>
                </div>
                <div class="flex items-center">
                  <input
                    v-model="form.targetUsers.vip"
                    type="checkbox"
                    id="vip-users"
                    class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                  />
                  <label for="vip-users" class="mr-2 block text-sm text-gray-900">
                    کاربران VIP
                  </label>
                </div>
              </div>
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                دسته‌بندی‌های محصولات
              </label>
              <div class="grid grid-cols-2 md:grid-cols-3 gap-2">
                <div class="flex items-center">
                  <input
                    v-model="form.targetCategories.electronics"
                    type="checkbox"
                    id="electronics"
                    class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                  />
                  <label for="electronics" class="mr-2 block text-sm text-gray-900">
                    الکترونیک
                  </label>
                </div>
                <div class="flex items-center">
                  <input
                    v-model="form.targetCategories.clothing"
                    type="checkbox"
                    id="clothing"
                    class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                  />
                  <label for="clothing" class="mr-2 block text-sm text-gray-900">
                    پوشاک
                  </label>
                </div>
                <div class="flex items-center">
                  <input
                    v-model="form.targetCategories.home"
                    type="checkbox"
                    id="home"
                    class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                  />
                  <label for="home" class="mr-2 block text-sm text-gray-900">
                    خانه و آشپزخانه
                  </label>
                </div>
                <div class="flex items-center">
                  <input
                    v-model="form.targetCategories.books"
                    type="checkbox"
                    id="books"
                    class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                  />
                  <label for="books" class="mr-2 block text-sm text-gray-900">
                    کتاب
                  </label>
                </div>
                <div class="flex items-center">
                  <input
                    v-model="form.targetCategories.sports"
                    type="checkbox"
                    id="sports"
                    class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                  />
                  <label for="sports" class="mr-2 block text-sm text-gray-900">
                    ورزشی
                  </label>
                </div>
                <div class="flex items-center">
                  <input
                    v-model="form.targetCategories.beauty"
                    type="checkbox"
                    id="beauty"
                    class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                  />
                  <label for="beauty" class="mr-2 block text-sm text-gray-900">
                    زیبایی
                  </label>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- شرایط خاص -->
        <div class="bg-gray-50 p-6 rounded-lg">
          <h4 class="text-md font-semibold text-gray-900 mb-4">شرایط خاص</h4>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                حداقل مبلغ سفارش
              </label>
              <input
                v-model="form.minOrderAmount"
                type="number"
                min="0"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="حداقل مبلغ سفارش"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                حداکثر مبلغ سفارش
              </label>
              <input
                v-model="form.maxOrderAmount"
                type="number"
                min="0"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="حداکثر مبلغ سفارش"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                حداقل تعداد محصول
              </label>
              <input
                v-model="form.minProductCount"
                type="number"
                min="1"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="حداقل تعداد محصول"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                حداکثر تعداد محصول
              </label>
              <input
                v-model="form.maxProductCount"
                type="number"
                min="1"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="حداکثر تعداد محصول"
              />
            </div>
          </div>
        </div>

        <!-- تنظیمات پیشرفته -->
        <div class="bg-gray-50 p-6 rounded-lg">
          <h4 class="text-md font-semibold text-gray-900 mb-4">تنظیمات پیشرفته</h4>
          <div class="space-y-4">
            <div class="flex items-center">
              <input
                v-model="form.settings.cumulative"
                type="checkbox"
                id="cumulative"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <label for="cumulative" class="mr-2 block text-sm text-gray-900">
                محدودیت تجمعی (در طول دوره)
              </label>
            </div>
            <div class="flex items-center">
              <input
                v-model="form.settings.rollover"
                type="checkbox"
                id="rollover"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <label for="rollover" class="mr-2 block text-sm text-gray-900">
                انتقال باقی‌مانده به دوره بعد
              </label>
            </div>
            <div class="flex items-center">
              <input
                v-model="form.settings.autoReset"
                type="checkbox"
                id="auto-reset"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <label for="auto-reset" class="mr-2 block text-sm text-gray-900">
                بازنشانی خودکار در پایان دوره
              </label>
            </div>
            <div class="flex items-center">
              <input
                v-model="form.settings.notifications"
                type="checkbox"
                id="notifications"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <label for="notifications" class="mr-2 block text-sm text-gray-900">
                ارسال اعلان هنگام نزدیک شدن به محدودیت
              </label>
            </div>
          </div>
        </div>

        <!-- پیام‌های خطا -->
        <div class="bg-gray-50 p-6 rounded-lg">
          <h4 class="text-md font-semibold text-gray-900 mb-4">پیام‌های خطا</h4>
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                پیام خطای محدودیت
              </label>
              <textarea
                v-model="form.errorMessages.limitReached"
                rows="2"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="پیام نمایش داده شده هنگام رسیدن به محدودیت"
              ></textarea>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                پیام هشدار نزدیک شدن به محدودیت
              </label>
              <textarea
                v-model="form.errorMessages.warning"
                rows="2"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="پیام هشدار هنگام نزدیک شدن به محدودیت"
              ></textarea>
            </div>
          </div>
        </div>

        <!-- وضعیت و توضیحات -->
        <div class="bg-gray-50 p-6 rounded-lg">
          <h4 class="text-md font-semibold text-gray-900 mb-4">وضعیت و توضیحات</h4>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                وضعیت <span class="text-red-500">*</span>
              </label>
              <select
                v-model="form.status"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              >
                <option value="active">فعال</option>
                <option value="inactive">غیرفعال</option>
                <option value="draft">پیش‌نویس</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                اولویت
              </label>
              <select
                v-model="form.priority"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              >
                <option value="low">کم</option>
                <option value="medium">متوسط</option>
                <option value="high">زیاد</option>
                <option value="critical">بحرانی</option>
              </select>
            </div>
            <div class="md:col-span-2">
              <label class="block text-sm font-medium text-gray-700 mb-2">
                توضیحات
              </label>
              <textarea
                v-model="form.description"
                rows="3"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="توضیحات اضافی درباره این محدودیت"
              ></textarea>
            </div>
          </div>
        </div>

        <!-- دکمه‌های عملیات -->
        <div class="flex justify-end gap-3 pt-6 border-t border-gray-200">
          <button
            type="button"
            @click="$emit('close')"
            class="px-4 py-2 bg-gray-300 text-gray-700 text-sm font-medium rounded-lg hover:bg-gray-400 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
          >
            انصراف
          </button>
          <button
            type="submit"
            :disabled="isSubmitting"
            class="px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed"
          >
            <span v-if="isSubmitting">در حال ثبت...</span>
            <span v-else>ثبت محدودیت</span>
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
// تعریف props و emits
const emit = defineEmits(['close', 'created'])

// متغیرهای reactive
const isSubmitting = ref(false)

// فرم داده‌ها
const form = ref({
  // اطلاعات اصلی
  name: '',
  type: '',
  limit: '',
  unit: '',
  
  // دوره زمانی
  period: '',
  customDays: '',
  startDate: '',
  endDate: '',
  
  // محدوده اعمال
  targetUsers: {
    all: false,
    new: false,
    premium: false,
    vip: false
  },
  targetCategories: {
    electronics: false,
    clothing: false,
    home: false,
    books: false,
    sports: false,
    beauty: false
  },
  
  // شرایط خاص
  minOrderAmount: '',
  maxOrderAmount: '',
  minProductCount: '',
  maxProductCount: '',
  
  // تنظیمات پیشرفته
  settings: {
    cumulative: false,
    rollover: false,
    autoReset: true,
    notifications: true
  },
  
  // پیام‌های خطا
  errorMessages: {
    limitReached: '',
    warning: ''
  },
  
  // وضعیت و توضیحات
  status: 'active',
  priority: 'medium',
  description: ''
})

// توابع
const getLimitPlaceholder = () => {
  const placeholders = {
    daily: 'مثال: 500000 (تومان)',
    weekly: 'مثال: 2000000 (تومان)',
    monthly: 'مثال: 5000000 (تومان)',
    yearly: 'مثال: 50000000 (تومان)',
    transaction_count: 'مثال: 10 (تعداد)',
    amount_limit: 'مثال: 1000000 (تومان)',
    time_limit: 'مثال: 24 (ساعت)',
    category_limit: 'مثال: 5 (تعداد)'
  }
  return placeholders[form.value.type] || 'مقدار محدودیت'
}

const handleTypeChange = () => {
  // تنظیم واحد پیش‌فرض بر اساس نوع
  const defaultUnits = {
    daily: 'toman',
    weekly: 'toman',
    monthly: 'toman',
    yearly: 'toman',
    transaction_count: 'count',
    amount_limit: 'toman',
    time_limit: 'hour',
    category_limit: 'count'
  }
  
  if (defaultUnits[form.value.type]) {
    form.value.unit = defaultUnits[form.value.type]
  }
  
  // تنظیم پیام‌های پیش‌فرض
  if (form.value.type === 'transaction_count') {
    form.value.errorMessages.limitReached = 'تعداد تراکنش‌های مجاز شما در این دوره به پایان رسیده است.'
    form.value.errorMessages.warning = 'شما نزدیک به محدودیت تعداد تراکنش هستید.'
  } else if (form.value.type.includes('amount') || form.value.type.includes('daily') || form.value.type.includes('weekly') || form.value.type.includes('monthly') || form.value.type.includes('yearly')) {
    form.value.errorMessages.limitReached = 'مبلغ مجاز شما در این دوره به پایان رسیده است.'
    form.value.errorMessages.warning = 'شما نزدیک به محدودیت مبلغ مجاز هستید.'
  }
}

const validateForm = () => {
  const errors = []
  
  // اعتبارسنجی نام
  if (!form.value.name.trim()) {
    errors.push('نام محدودیت الزامی است')
  }
  
  // اعتبارسنجی نوع
  if (!form.value.type) {
    errors.push('نوع محدودیت الزامی است')
  }
  
  // اعتبارسنجی مقدار
  if (!form.value.limit || form.value.limit <= 0) {
    errors.push('مقدار محدودیت باید بزرگتر از صفر باشد')
  }
  
  // اعتبارسنجی واحد
  if (!form.value.unit) {
    errors.push('واحد الزامی است')
  }
  
  // اعتبارسنجی دوره
  if (!form.value.period) {
    errors.push('دوره زمانی الزامی است')
  }
  
  // اعتبارسنجی دوره سفارشی
  if (form.value.period === 'custom' && (!form.value.customDays || form.value.customDays <= 0)) {
    errors.push('تعداد روزهای سفارشی الزامی است')
  }
  
  // اعتبارسنجی تاریخ‌ها
  if (form.value.startDate && form.value.endDate && form.value.startDate > form.value.endDate) {
    errors.push('تاریخ شروع نمی‌تواند بعد از تاریخ پایان باشد')
  }
  
  return errors
}

const handleSubmit = async () => {
  const errors = validateForm()
  
  if (errors.length > 0) {
    alert('خطاهای زیر را برطرف کنید:\n' + errors.join('\n'))
    return
  }
  
  isSubmitting.value = true
  
  try {
    // شبیه‌سازی ارسال به سرور
    await new Promise(resolve => setTimeout(resolve, 1000))
    
    // ایجاد محدودیت جدید
    const newRestriction = {
      name: form.value.name,
      type: form.value.type,
      limit: parseInt(form.value.limit),
      unit: form.value.unit,
      period: form.value.period,
      status: form.value.status,
      priority: form.value.priority,
      // سایر اطلاعات...
    }
    
    emit('created', newRestriction)
    
    // پاک کردن فرم
    form.value = {
      name: '',
      type: '',
      limit: '',
      unit: '',
      period: '',
      customDays: '',
      startDate: '',
      endDate: '',
      targetUsers: {
        all: false,
        new: false,
        premium: false,
        vip: false
      },
      targetCategories: {
        electronics: false,
        clothing: false,
        home: false,
        books: false,
        sports: false,
        beauty: false
      },
      minOrderAmount: '',
      maxOrderAmount: '',
      minProductCount: '',
      maxProductCount: '',
      settings: {
        cumulative: false,
        rollover: false,
        autoReset: true,
        notifications: true
      },
      errorMessages: {
        limitReached: '',
        warning: ''
      },
      status: 'active',
      priority: 'medium',
      description: ''
    }
    
  } catch (error) {
    console.error('خطا در ثبت محدودیت:', error)
    alert('خطا در ثبت محدودیت. لطفاً دوباره تلاش کنید.')
  } finally {
    isSubmitting.value = false
  }
}
</script>

<style scoped>
/* استایل‌های خاص مودال */
</style> 

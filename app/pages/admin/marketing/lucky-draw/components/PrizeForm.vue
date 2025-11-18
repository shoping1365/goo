<template>
  <div class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
    <div class="relative top-20 mx-auto p-5 border w-11/12 md:w-3/4 lg:w-1/2 shadow-lg rounded-md bg-white">
      <div class="mt-3">
        <!-- هدر مودال -->
        <div class="flex justify-between items-center mb-6">
          <h3 class="text-lg font-medium text-gray-900">
            {{ isEditing ? 'ویرایش جایزه' : 'ایجاد جایزه جدید' }}
          </h3>
          <button 
            @click="$emit('close')"
            class="text-gray-400 hover:text-gray-600"
          >
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>

        <!-- فرم -->
        <form @submit.prevent="savePrize" class="space-y-6">
          <!-- نام و نوع جایزه -->
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نام جایزه *</label>
              <input 
                v-model="form.name"
                type="text"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                placeholder="نام جایزه را وارد کنید"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نوع جایزه *</label>
              <select 
                v-model="form.type"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              >
                <option value="">انتخاب کنید</option>
                <option value="gift_card">کارت هدیه</option>
                <option value="free_product">محصول رایگان</option>
                <option value="discount">تخفیف</option>
                <option value="free_shipping">ارسال رایگان</option>
                <option value="cash">نقدی</option>
              </select>
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">توضیحات</label>
            <textarea 
              v-model="form.description"
              rows="3"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              placeholder="توضیحات جایزه را وارد کنید"
            ></textarea>
          </div>

          <!-- ارزش و تعداد -->
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">ارزش جایزه (تومان)</label>
              <input 
                v-model.number="form.value"
                type="number"
                min="0"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                placeholder="0"
              />
              <p class="text-xs text-gray-500 mt-1">برای جوایز رایگان صفر وارد کنید</p>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">تعداد *</label>
              <input 
                v-model.number="form.quantity"
                type="number"
                min="1"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                placeholder="تعداد جایزه"
              />
            </div>
          </div>

          <!-- تنظیمات خاص نوع جایزه -->
          <div v-if="form.type === 'gift_card'" class="space-y-4">
            <h4 class="text-md font-medium text-gray-900">تنظیمات کارت هدیه</h4>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">حداقل مبلغ استفاده</label>
                <input 
                  v-model.number="form.minAmount"
                  type="number"
                  min="0"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                  placeholder="0"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ انقضا</label>
                <input 
                  v-model="form.expiryDate"
                  type="date"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                />
              </div>
            </div>
          </div>

          <div v-if="form.type === 'discount'" class="space-y-4">
            <h4 class="text-md font-medium text-gray-900">تنظیمات تخفیف</h4>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">درصد تخفیف</label>
                <input 
                  v-model.number="form.discountPercent"
                  type="number"
                  min="1"
                  max="100"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                  placeholder="30"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">حداقل مبلغ خرید</label>
                <input 
                  v-model.number="form.minPurchaseAmount"
                  type="number"
                  min="0"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                  placeholder="0"
                />
              </div>
            </div>
          </div>

          <div v-if="form.type === 'free_shipping'" class="space-y-4">
            <h4 class="text-md font-medium text-gray-900">تنظیمات ارسال رایگان</h4>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">مدت اعتبار (روز)</label>
                <input 
                  v-model.number="form.validityDays"
                  type="number"
                  min="1"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                  placeholder="30"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">حداقل مبلغ خرید</label>
                <input 
                  v-model.number="form.minPurchaseAmount"
                  type="number"
                  min="0"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                  placeholder="0"
                />
              </div>
            </div>
          </div>

          <!-- شرایط استفاده -->
          <div class="space-y-4">
            <h4 class="text-md font-medium text-gray-900">شرایط استفاده</h4>
            
            <div class="flex items-center">
              <input 
                v-model="form.requireRegistration"
                type="checkbox"
                id="requireRegistration"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <label for="requireRegistration" class="mr-2 text-sm text-gray-700">نیاز به ثبت‌نام</label>
            </div>

            <div class="flex items-center">
              <input 
                v-model="form.oneTimeUse"
                type="checkbox"
                id="oneTimeUse"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <label for="oneTimeUse" class="mr-2 text-sm text-gray-700">استفاده یکبار</label>
            </div>

            <div class="flex items-center">
              <input 
                v-model="form.combinable"
                type="checkbox"
                id="combinable"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <label for="combinable" class="mr-2 text-sm text-gray-700">قابل ترکیب با سایر تخفیف‌ها</label>
            </div>
          </div>

          <!-- وضعیت -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
            <select 
              v-model="form.status"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            >
              <option value="active">فعال</option>
              <option value="inactive">غیرفعال</option>
              <option value="draft">پیش‌نویس</option>
            </select>
          </div>

          <!-- دکمه‌های عملیات -->
          <div class="flex justify-end space-x-3 space-x-reverse">
            <button 
              type="button"
              @click="$emit('close')"
              class="px-4 py-2 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400 transition-colors"
            >
              انصراف
            </button>
            <button 
              type="submit"
              :disabled="isSaving"
              class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors disabled:opacity-50"
            >
              {{ isSaving ? 'در حال ذخیره...' : (isEditing ? 'ویرایش' : 'ایجاد') }}
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  prize: {
    type: Object,
    default: null
  }
})

const emit = defineEmits(['close', 'saved'])

// متغیرهای reactive
const isSaving = ref(false)
const form = ref({
  name: '',
  description: '',
  type: '',
  value: 0,
  quantity: 1,
  status: 'active',
  // تنظیمات خاص کارت هدیه
  minAmount: 0,
  expiryDate: '',
  // تنظیمات خاص تخفیف
  discountPercent: 0,
  minPurchaseAmount: 0,
  // تنظیمات خاص ارسال رایگان
  validityDays: 30,
  // شرایط استفاده
  requireRegistration: false,
  oneTimeUse: true,
  combinable: false
})

// محاسبه‌گرها
const isEditing = computed(() => !!props.prize)

// تنظیم مقادیر اولیه
onMounted(() => {
  if (props.prize) {
    form.value = { ...props.prize }
  } else {
    // تنظیم تاریخ انقضای پیش‌فرض (یک ماه بعد)
    const nextMonth = new Date()
    nextMonth.setMonth(nextMonth.getMonth() + 1)
    form.value.expiryDate = nextMonth.toISOString().split('T')[0]
  }
})

// توابع
const savePrize = async () => {
  try {
    isSaving.value = true
    
    // شبیه‌سازی API call
    await new Promise(resolve => setTimeout(resolve, 1000))
    
    // ارسال داده به کامپوننت والد
    const savedData = {
      ...form.value,
      id: props.prize?.id || Date.now()
    }
    
    emit('saved', savedData)
  } catch (error) {
    console.error('خطا در ذخیره جایزه:', error)
    alert('خطا در ذخیره جایزه')
  } finally {
    isSaving.value = false
  }
}
</script>

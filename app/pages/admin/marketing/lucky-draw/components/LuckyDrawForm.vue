<template>
  <div class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
    <div class="relative top-20 mx-auto p-5 border w-11/12 md:w-3/4 lg:w-1/2 shadow-lg rounded-md bg-white">
      <div class="mt-3">
        <!-- هدر مودال -->
        <div class="flex justify-between items-center mb-6">
          <h3 class="text-lg font-medium text-gray-900">
            {{ isEditing ? 'ویرایش گرونه' : 'ایجاد گرونه جدید' }}
          </h3>
          <button 
            class="text-gray-400 hover:text-gray-600"
            @click="$emit('close')"
          >
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>

        <!-- فرم -->
        <form class="space-y-6" @submit.prevent="saveLuckyDraw">
          <!-- نام و توضیحات -->
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نام گرونه *</label>
              <input 
                v-model="form.name"
                type="text"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                placeholder="نام گرونه را وارد کنید"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نوع گرونه *</label>
              <select 
                v-model="form.type"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              >
                <option value="">انتخاب کنید</option>
                <option value="daily">روزانه</option>
                <option value="weekly">هفتگی</option>
                <option value="monthly">ماهانه</option>
                <option value="purchase">خرید</option>
                <option value="registration">ثبت‌نام</option>
              </select>
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">توضیحات</label>
            <textarea 
              v-model="form.description"
              rows="3"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              placeholder="توضیحات گرونه را وارد کنید"
            ></textarea>
          </div>

          <!-- تاریخ‌ها -->
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ شروع *</label>
              <input 
                v-model="form.startDate"
                type="date"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ پایان *</label>
              <input 
                v-model="form.endDate"
                type="date"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
            </div>
          </div>

          <!-- تعداد جوایز و شرکت‌کنندگان -->
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">تعداد جوایز *</label>
              <input 
                v-model.number="form.prizeCount"
                type="number"
                min="1"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                placeholder="تعداد جوایز"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر شرکت‌کنندگان</label>
              <input 
                v-model.number="form.maxParticipants"
                type="number"
                min="1"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                placeholder="نامحدود"
              />
            </div>
          </div>

          <!-- شرایط شرکت -->
          <div class="space-y-4">
            <h4 class="text-md font-medium text-gray-900">شرایط شرکت</h4>
            
            <div class="flex items-center">
              <input 
                id="requirePurchase"
                v-model="form.requirePurchase"
                type="checkbox"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <label for="requirePurchase" class="mr-2 text-sm text-gray-700">نیاز به خرید</label>
            </div>

            <div v-if="form.requirePurchase" class="mr-6">
              <label class="block text-sm font-medium text-gray-700 mb-2">حداقل مبلغ خرید (تومان)</label>
              <input 
                v-model.number="form.minPurchaseAmount"
                type="number"
                min="0"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                placeholder="0"
              />
            </div>

            <div class="flex items-center">
              <input 
                id="requireRegistration"
                v-model="form.requireRegistration"
                type="checkbox"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <label for="requireRegistration" class="mr-2 text-sm text-gray-700">نیاز به ثبت‌نام</label>
            </div>

            <div class="flex items-center">
              <input 
                id="oneEntryPerUser"
                v-model="form.oneEntryPerUser"
                type="checkbox"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <label for="oneEntryPerUser" class="mr-2 text-sm text-gray-700">یک شرکت در هر کاربر</label>
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
              class="px-4 py-2 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400 transition-colors"
              @click="$emit('close')"
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
  luckyDraw: {
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
  startDate: '',
  endDate: '',
  prizeCount: 1,
  maxParticipants: null,
  requirePurchase: false,
  minPurchaseAmount: 0,
  requireRegistration: false,
  oneEntryPerUser: false,
  status: 'active'
})

// محاسبه‌گرها
const isEditing = computed(() => !!props.luckyDraw)

// تنظیم مقادیر اولیه
onMounted(() => {
  if (props.luckyDraw) {
    form.value = { ...props.luckyDraw }
  } else {
    // تنظیم تاریخ‌های پیش‌فرض
    const today = new Date()
    const nextMonth = new Date(today.getFullYear(), today.getMonth() + 1, today.getDate())
    
    form.value.startDate = today.toISOString().split('T')[0]
    form.value.endDate = nextMonth.toISOString().split('T')[0]
  }
})

// توابع
const saveLuckyDraw = async () => {
  try {
    isSaving.value = true
    
    // شبیه‌سازی API call
    await new Promise(resolve => setTimeout(resolve, 1000))
    
    // ارسال داده به کامپوننت والد
    const savedData = {
      ...form.value,
      id: props.luckyDraw?.id || Date.now(),
      participants: props.luckyDraw?.participants || 0
    }
    
    emit('saved', savedData)
  } catch (error) {
    console.error('خطا در ذخیره گرونه:', error)
    alert('خطا در ذخیره گرونه')
  } finally {
    isSaving.value = false
  }
}
</script>

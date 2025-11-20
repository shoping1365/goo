<template>
  <div v-if="isOpen" class="fixed inset-0 flex items-center justify-center z-50 p-6">
    <div class="bg-white rounded-lg max-w-2xl w-full max-h-[90vh] overflow-y-auto border-2 border-blue-200">
      <!-- هدر مودال -->
      <div class="flex items-center justify-between p-6 border-b border-blue-200 text-right">
        <button class="text-gray-400 hover:text-gray-600" @click="closeModal">
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
        </button>
        <h2 class="text-xl font-bold text-gray-900">
          {{ isEditing ? 'ویرایش آدرس' : 'افزودن آدرس جدید' }}
        </h2>
      </div>

      <!-- فرم ویرایش آدرس -->
      <form class="p-6 space-y-4 text-right" @submit.prevent="handleSubmit">
        <!-- ردیف اول: نام دریافت‌کننده -->
        <div>
          <label for="recipient_name" class="block text-sm font-medium text-gray-700 mb-2">
            نام دریافت‌کننده *
          </label>
            <input
              id="recipient_name"
              v-model="form.recipient_name"
            type="text"
            required
            class="w-full px-3 py-2 text-base border border-gray-300 rounded focus:outline-none focus:ring-1 focus:ring-teal-500 focus:border-transparent text-right"
            placeholder="نام و نام خانوادگی"
          />
        </div>

        <!-- ردیف دوم: شماره تلفن و شماره موبایل -->
        <div class="grid grid-cols-2 gap-6">
          <div>
            <label for="recipient_mobile" class="block text-xs font-medium text-gray-700 mb-1">
              شماره تلفن *
            </label>
            <input
              id="recipient_mobile"
              v-model="form.recipient_mobile"
              type="tel"
              required
              class="w-full px-2 py-1 text-sm border border-gray-300 rounded focus:outline-none focus:ring-1 focus:ring-teal-500 focus:border-transparent text-right"
              placeholder="09123456789"
            />
          </div>
          <div>
            <label for="phone" class="block text-xs font-medium text-gray-700 mb-1">
              شماره موبایل
            </label>
            <input
              id="phone"
              v-model="form.phone"
              type="tel"
              class="w-full px-2 py-1 text-sm border border-gray-300 rounded focus:outline-none focus:ring-1 focus:ring-teal-500 focus:border-transparent text-right"
              placeholder="09123456789"
            />
          </div>
        </div>

        <!-- ردیف سوم: شهر و استان -->
        <div class="grid grid-cols-2 gap-6">
          <div>
            <label for="city" class="block text-xs font-medium text-gray-700 mb-1">
              شهر *
            </label>
            <input
              id="city"
              v-model="form.city"
              type="text"
              required
              class="w-full px-2 py-1 text-sm border border-gray-300 rounded focus:outline-none focus:ring-1 focus:ring-teal-500 focus:border-transparent text-right"
              placeholder="نام شهر"
            />
          </div>
          <div>
            <label for="province" class="block text-xs font-medium text-gray-700 mb-1">
              استان *
            </label>
            <input
              id="province"
              v-model="form.province"
              type="text"
              required
              class="w-full px-2 py-1 text-sm border border-gray-300 rounded focus:outline-none focus:ring-1 focus:ring-teal-500 focus:border-transparent text-right"
              placeholder="نام استان"
            />
          </div>
        </div>

        <!-- آدرس کامل -->
        <div>
          <label for="street" class="block text-xs font-medium text-gray-700 mb-1">
            آدرس کامل *
          </label>
          <textarea
            id="street"
            v-model="form.street"
            required
            rows="2"
            class="w-full px-2 py-1 text-sm border border-blue-300 rounded focus:outline-none focus:ring-1 focus:ring-teal-500 focus:border-transparent text-right"
            placeholder="آدرس کامل شامل خیابان، کوچه، پلاک و... "
          ></textarea>
        </div>

        <!-- ردیف چهارم: کد پستی و برچسب آدرس -->
        <div class="grid grid-cols-2 gap-6">
          <div>
            <label for="label" class="block text-xs font-medium text-gray-700 mb-1">
              برچسب آدرس
            </label>
            <input
              id="label"
              v-model="form.label"
              type="text"
              class="w-full px-2 py-1 text-sm border border-gray-300 rounded focus:outline-none focus:ring-1 focus:ring-teal-500 focus:border-transparent text-right"
              placeholder="مثل: خانه، محل کار، و..."
            />
          </div>
          <div>
            <label for="postal_code" class="block text-xs font-medium text-gray-700 mb-1">
              کد پستی
            </label>
            <input
              id="postal_code"
              v-model="form.postal_code"
              type="text"
              class="w-full px-2 py-1 text-sm border border-gray-300 rounded focus:outline-none focus:ring-1 focus:ring-teal-500 focus:border-transparent text-right"
              placeholder="1234567890"
            />
          </div>
        </div>

        <!-- دکمه‌های عملیات -->
        <div class="pt-4 border-t border-blue-200">
          <div class="flex items-center justify-between">
            <!-- سمت چپ: دکمه‌ها -->
            <div class="flex items-center space-x-3 space-x-reverse">
              <button
                type="button"
                class="px-3 py-1 text-xs font-medium text-gray-700 bg-white border border-gray-300 rounded hover:bg-gray-50 focus:outline-none focus:ring-1 focus:ring-offset-1 focus:ring-teal-500"
                @click="closeModal"
              >
                انصراف
              </button>
              <div class="w-4"></div>
              <button
                type="submit"
                :disabled="loading"
                class="px-3 py-1 text-xs font-medium text-white bg-teal-600 border border-transparent rounded hover:bg-teal-700 focus:outline-none focus:ring-1 focus:ring-offset-1 focus:ring-teal-500 disabled:opacity-50 disabled:cursor-not-allowed"
              >
                <span v-if="loading" class="flex items-center">
                  <svg class="animate-spin -ml-1 mr-1 h-3 w-3 text-white" fill="none" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                  </svg>
                  در حال ذخیره...
                </span>
                <span v-else>
                  {{ isEditing ? 'ذخیره تغییرات' : 'افزودن آدرس' }}
                </span>
              </button>
            </div>
            
            <!-- سمت راست: آدرس پیش‌فرض -->
            <div class="flex items-center">
              <label for="is_default" class="ml-2 block text-xs text-gray-900">
                تنظیم به عنوان آدرس پیش‌فرض
              </label>
              <input
                id="is_default"
                v-model="form.is_default"
                type="checkbox"
                class="h-3 w-3 text-teal-600 focus:ring-teal-500 border-gray-300 rounded"
              />
            </div>
          </div>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useAddresses } from '~/composables/useAddresses'
import { ref, computed, watch } from 'vue'

// تعریف نوع‌های TypeScript
type Address = {
  id?: number
  recipient_name?: string
  recipient_mobile?: string
  phone?: string
  province?: string
  city?: string
  street?: string
  full_address?: string
  postal_code?: string
  label?: string
  is_default?: boolean
}

// تعریف props
const props = defineProps<{
  isOpen: boolean
  address?: Address | null
}>()

// تعریف emits
const emit = defineEmits<{
  close: []
  saved: [address: Address]
}>()

// استفاده از composable آدرس‌ها
const { addAddress, updateAddress } = useAddresses()

// حالت‌های فرم
const loading = ref(false)
const isEditing = computed(() => !!props.address?.id)

// فرم داده‌ها
const form = ref({
  recipient_name: '',
  recipient_mobile: '',
  phone: '',
  province: '',
  city: '',
  street: '',
  postal_code: '',
  label: '',
  is_default: false
})

// مقداردهی اولیه فرم
const initializeForm = () => {
  if (props.address) {
    form.value = {
      recipient_name: props.address.recipient_name || '',
      recipient_mobile: props.address.recipient_mobile || props.address.phone || '',
      phone: props.address.phone || '',
      province: props.address.province || '',
      city: props.address.city || '',
      street: props.address.street || props.address.full_address || '',
      postal_code: props.address.postal_code || '',
      label: props.address.label || '',
      is_default: props.address.is_default || false
    }
  } else {
    form.value = {
      recipient_name: '',
      recipient_mobile: '',
      phone: '',
      province: '',
      city: '',
      street: '',
      postal_code: '',
      label: '',
      is_default: false
    }
  }
}

// بستن مودال
const closeModal = () => {
  emit('close')
}

// ارسال فرم
const handleSubmit = async () => {
  try {
    loading.value = true
    
    const addressData = {
      ...form.value,
      full_address: form.value.street, // برای سازگاری با API
      phone: form.value.recipient_mobile // استفاده از شماره تلفن اصلی
    }

    let result
    if (isEditing.value) {
      result = await updateAddress(props.address.id, addressData)
    } else {
      result = await addAddress(addressData)
    }

    emit('saved', result)
    closeModal()
  } catch (error) {
    console.error('خطا در ذخیره آدرس:', error)
    // TODO: نمایش پیام خطا به کاربر
  } finally {
    loading.value = false
  }
}

// مقداردهی اولیه هنگام تغییر props
watch(() => props.isOpen, (newValue) => {
  if (newValue) {
    initializeForm()
  }
}, { immediate: true })
</script>

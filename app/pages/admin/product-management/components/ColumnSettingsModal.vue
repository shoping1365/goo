<template>
  <div v-if="isOpen" class="fixed inset-0 z-50 flex items-center justify-center pointer-events-none">
    <!-- Backdrop -->
    
    <!-- Modal -->
    <div class="relative bg-white rounded-lg shadow-xl w-full max-w-md sm:max-w-lg md:max-w-xl mx-4 pointer-events-auto">
      <!-- Header -->
      <div class="flex items-center justify-between p-6 border-b bg-blue-50">
        <h3 class="text-lg font-semibold text-gray-900">
          تنظیمات ستون‌ها
        </h3>
        <button
          class="text-gray-400 hover:text-gray-600 transition-colors"
          @click="closeModal"
        >
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
        </button>
      </div>

      <!-- Content -->
      <div class="p-6">
        <p class="text-sm text-gray-600 mb-4">
          ستون‌هایی که می‌خواهید در جدول نمایش داده شوند را انتخاب کنید:
        </p>

        <div class="space-y-3 max-h-64 overflow-y-auto">
          <div
            v-for="column in availableColumns"
            :key="column.key"
            class="flex items-center"
          >
            <input
              :id="column.key"
              v-model="selectedColumns"
              :value="column.key"
              type="checkbox"
              class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
            />
            <label
              :for="column.key"
              class="mr-3 text-sm font-medium text-gray-700 cursor-pointer"
            >
              {{ column.label }}
            </label>
          </div>
        </div>
      </div>

      <!-- Footer -->
      <div class="flex items-center justify-end space-x-3 space-x-reverse p-6 border-t bg-blue-50">
        <button
          class="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          @click="closeModal"
        >
          انصراف
        </button>
        <button
          :disabled="isLoading"
          class="px-4 py-2 text-sm font-medium text-white bg-blue-600 border border-transparent rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50 disabled:cursor-not-allowed"
          @click="saveSettings"
        >
          <span v-if="isLoading">در حال ذخیره...</span>
          <span v-else>ذخیره</span>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
// useToast is now available globally through the composable

const props = defineProps({
  isOpen: {
    type: Boolean,
    default: false
  },
  visibleColumns: {
    type: Array,
    default: () => []
  }
})

const emit = defineEmits(['close', 'update:visibleColumns'])

const isLoading = ref(false)

// ستون‌های موجود برای انتخاب
const availableColumns = [
  { key: 'select', label: 'انتخاب' },
  { key: 'index', label: '#' },
  { key: 'image', label: 'تصویر' },
  { key: 'name', label: 'نام' },
  { key: 'sku', label: 'شناسه کالا در انبار' },
  { key: 'price', label: 'قیمت فروش' },
  { key: 'old_price', label: 'قیمت خط خورده' },
  { key: 'cost', label: 'قیمت خرید' },
  { key: 'track_inventory', label: 'مدیریت موجودی' },
  { key: 'stock_quantity', label: 'تعداد انبار' },
  { key: 'isNew', label: 'علامت گذاری به عنوان جدید' },
  { key: 'isPublished', label: 'منتشر شده' },
  { key: 'showOnHomepage', label: 'نمایش در صفحه اصلی' },
  { key: 'disableBuyButton', label: 'دکمه خرید غیرفعال باشد - ناموجود' },
  { key: 'callForPrice', label: 'تماس برای قیمت' },
]

// ستون‌های انتخاب شده
const selectedColumns = ref([])

// وقتی modal باز می‌شود، ستون‌های فعلی را تنظیم کن
watch(() => props.isOpen, (newValue) => {
  if (newValue) {
    selectedColumns.value = [...props.visibleColumns]
  }
})

// بستن modal
const closeModal = () => {
  emit('close')
}

// ذخیره تنظیمات
const saveSettings = async () => {
  try {
    isLoading.value = true
    
    // ارسال تنظیمات به سرور
    await $fetch('/api/admin/admin-settings/bulk-edit-columns', {
      method: 'PUT',
      body: {
        value: JSON.stringify(selectedColumns.value)
      }
    })

    // ارسال ستون‌های جدید به parent
    emit('update:visibleColumns', selectedColumns.value)
    
    // بستن modal
    closeModal()
    
    // نمایش پیام موفقیت
    useToast().success('تنظیمات با موفقیت ذخیره شد')
    
  } catch (error) {
    console.error('خطا در ذخیره تنظیمات:', error)
    useToast().error('خطا در ذخیره تنظیمات')
  } finally {
    isLoading.value = false
  }
}
</script> 

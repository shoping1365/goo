<template>
  <div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-6">
    <div class="bg-white rounded-lg shadow-xl w-full max-w-2xl max-h-[90vh] overflow-y-auto">
      <!-- هدر مودال -->
      <div class="p-6 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <h2 class="text-xl font-semibold text-gray-900">
            {{ coupon ? 'ویرایش کوپن' : 'افزودن کوپن جدید' }}
          </h2>
          <button @click="$emit('cancel')" class="text-gray-400 hover:text-gray-600">
            <span class="i-heroicons-x-mark text-xl"></span>
          </button>
        </div>
      </div>

      <!-- فرم -->
      <form @submit.prevent="handleSubmit" class="p-6 space-y-6">
        <!-- اطلاعات اصلی -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نام کوپن</label>
            <input v-model="form.name" type="text" required class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">کد کوپن</label>
            <input v-model="form.code" type="text" required class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">توضیحات</label>
          <textarea v-model="form.description" rows="3" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"></textarea>
        </div>

        <!-- تنظیمات تخفیف -->
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع تخفیف</label>
            <select v-model="form.type" @change="updateDiscountValue" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
              <option value="percentage">درصدی</option>
              <option value="fixed">مبلغ ثابت</option>
              <option value="free_shipping">ارسال رایگان</option>
            </select>
          </div>
          <div v-if="form.type !== 'free_shipping'">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              {{ form.type === 'percentage' ? 'درصد تخفیف' : 'مبلغ تخفیف' }}
            </label>
            <input v-model="form.discountValue" type="number" min="0" :max="form.type === 'percentage' ? 100 : undefined" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداقل مبلغ سفارش</label>
            <input v-model="form.minOrderAmount" type="number" min="0" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
          </div>
        </div>

        <!-- محدودیت‌ها -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر تعداد استفاده</label>
            <input v-model="form.maxUses" type="number" min="0" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تعداد استفاده برای هر کاربر</label>
            <input v-model="form.maxUsesPerUser" type="number" min="0" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
          </div>
        </div>

        <!-- تاریخ‌ها -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ شروع</label>
            <input v-model="form.startsAt" type="datetime-local" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ انقضا</label>
            <input v-model="form.expiresAt" type="datetime-local" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
          </div>
        </div>

        <!-- وضعیت -->
        <div>
          <label class="flex items-center">
            <input v-model="form.isActive" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
            <span class="mr-2 text-sm text-gray-700">کوپن فعال باشد</span>
          </label>
        </div>

        <!-- دکمه‌ها -->
        <div class="flex items-center justify-end gap-3 pt-6 border-t border-gray-200">
          <button type="button" @click="$emit('cancel')" class="px-4 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200 transition-colors">
            انصراف
          </button>
          <button type="submit" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors">
            {{ coupon ? 'ویرایش' : 'افزودن' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'

// Props
interface Props {
  coupon?: any
}

const props = withDefaults(defineProps<Props>(), {
  coupon: undefined
})

// Emits
const emit = defineEmits<{
  save: [coupon: any]
  cancel: []
}>()

// فرم
const form = ref({
  name: '',
  code: '',
  description: '',
  type: 'percentage',
  discountValue: 0,
  minOrderAmount: 0,
  maxUses: null,
  maxUsesPerUser: 1,
  startsAt: '',
  expiresAt: '',
  isActive: true
})

// بروزرسانی مقدار تخفیف
const updateDiscountValue = () => {
  if (form.value.type === 'free_shipping') {
    form.value.discountValue = 0
  }
}

// ارسال فرم
const handleSubmit = () => {
  emit('save', { ...form.value })
}

// مقداردهی اولیه
onMounted(() => {
  if (props.coupon) {
    form.value = { ...props.coupon }
  }
})
</script> 

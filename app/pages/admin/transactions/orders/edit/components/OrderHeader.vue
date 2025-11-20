<template>
  <div class="mb-4">
    <div class="w-full px-4 sm:px-6 lg:px-8">
      <div class="flex justify-between items-center py-3">
        <div class="flex items-center">
          <NuxtLink 
            to="/admin/transactions/orders" 
            class="text-gray-500 hover:text-gray-700 mr-4 transition-colors"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path>
            </svg>
          </NuxtLink>
          <div>
            <h1 class="text-lg font-bold text-gray-900">ویرایش سفارش #{{ orderNumber }}</h1>
            <p class="text-xs text-gray-600 mt-1">مشاهده و ویرایش جزئیات سفارش</p>
          </div>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button 
            class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors flex items-center"
            @click="saveOrder"
          >
            <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
            </svg>
            ذخیره تغییرات
          </button>
          <button 
            class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors flex items-center"
            @click="printOrder"
          >
            <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 17h2a2 2 0 002-2v-4a2 2 0 00-2-2H5a2 2 0 00-2 2v4a2 2 0 002 2h2m2 4h6a2 2 0 002-2v-4a2 2 0 00-2-2H9a2 2 0 00-2 2v4a2 2 0 002 2zm8-12V5a2 2 0 00-2-2H9a2 2 0 00-2 2v4h10z"></path>
            </svg>
            چاپ سفارش
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
// تعریف props
defineProps({
  orderNumber: {
    type: String,
    required: true
  }
})

// تعریف emit
const emit = defineEmits(['save-order', 'print-order'])

// توابع
const saveOrder = () => {
  emit('save-order')
}

const printOrder = () => {
  emit('print-order')
}
</script> 

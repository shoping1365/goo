<template>
  <div class="px-4 py-4 mb-6">
    <h3 class="text-lg font-semibold text-gray-900 mb-4">فیلتر سفارشات</h3>
    <div class="grid grid-cols-1 md:grid-cols-4 gapx-4 py-4">
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">جستجو</label>
        <input 
          v-model="searchTerm"
          type="text" 
          class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          placeholder="شماره سفارش یا نام مشتری"
        />
      </div>
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
        <select 
          v-model="statusFilter"
          class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
        >
          <option value="">همه وضعیت‌ها</option>
          <option value="pending">در انتظار پرداخت</option>
          <option value="paid">پرداخت شده</option>
          <option value="processing">در حال پردازش</option>
          <option value="shipped">ارسال شده</option>
          <option value="delivered">تحویل داده شده</option>
          <option value="cancelled">لغو شده</option>
          <option value="refunded">بازگشت وجه</option>
        </select>
      </div>
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">از تاریخ</label>
        <input 
          v-model="dateFrom"
          type="date" 
          class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
        />
      </div>
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">تا تاریخ</label>
        <input 
          v-model="dateTo"
          type="date" 
          class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
        />
      </div>
    </div>
    <div class="flex justify-end mt-4">
      <button 
        @click="applyFilters"
        class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
      >
        اعمال فیلتر
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'

// تعریف props برای دریافت مقادیر از کامپوننت والد
const props = defineProps({
  modelValue: {
    type: Object,
    default: () => ({
      searchTerm: '',
      statusFilter: '',
      dateFrom: '',
      dateTo: ''
    })
  }
})

// تعریف emit برای ارسال تغییرات به کامپوننت والد
const emit = defineEmits(['update:modelValue', 'apply-filters'])

// متغیرهای محلی
const searchTerm = ref(props.modelValue.searchTerm)
const statusFilter = ref(props.modelValue.statusFilter)
const dateFrom = ref(props.modelValue.dateFrom)
const dateTo = ref(props.modelValue.dateTo)

// نظارت بر تغییرات و ارسال به والد
watch([searchTerm, statusFilter, dateFrom, dateTo], () => {
  emit('update:modelValue', {
    searchTerm: searchTerm.value,
    statusFilter: statusFilter.value,
    dateFrom: dateFrom.value,
    dateTo: dateTo.value
  })
}, { deep: true })

// اعمال فیلترها
const applyFilters = () => {
  emit('apply-filters')
}
</script> 

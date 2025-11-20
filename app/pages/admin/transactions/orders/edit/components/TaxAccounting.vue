<template>
  <div class="space-y-6">
    <!-- اطلاعات مالیاتی -->
    <div class="px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-red-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 7h6m0 10v-3m-3 3h.01M9 17h.01M9 14h.01M12 14h.01M15 11h.01M12 11h.01M9 11h.01M7 21h10a2 2 0 002-2V5a2 2 0 00-2-2H7a2 2 0 00-2 2v14a2 2 0 002 2z"></path>
        </svg>
        اطلاعات مالیاتی
      </h3>

      <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
        <!-- محاسبات مالیاتی -->
        <div>
          <h4 class="text-md font-medium text-gray-700 mb-3">محاسبات مالیاتی</h4>
          <div class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">نرخ مالیات منطقه‌ای (%)</label>
              <input 
                v-model="taxAccounting.regionalTaxRate"
                type="number" 
                step="0.01"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="9.00"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">مبلغ مشمول مالیات</label>
              <input 
                v-model="taxAccounting.taxableAmount"
                type="number" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="0"
                readonly
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">مبلغ مالیات محاسبه شده</label>
              <input 
                v-model="taxAccounting.calculatedTax"
                type="number" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="0"
                readonly
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">معافیت مالیاتی</label>
              <select 
                v-model="taxAccounting.taxExemption"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="">بدون معافیت</option>
                <option value="export">صادرات</option>
                <option value="diplomatic">سفارت‌خانه</option>
                <option value="charity">خیریه</option>
                <option value="government">دولتی</option>
                <option value="special">ویژه</option>
              </select>
            </div>
          </div>
        </div>

        <!-- فاکتور مالیاتی -->
        <div>
          <h4 class="text-md font-medium text-gray-700 mb-3">فاکتور مالیاتی</h4>
          <div class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">شماره فاکتور مالیاتی</label>
              <input 
                v-model="taxAccounting.taxInvoiceNumber"
                type="text" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="TAX-2024-001"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">تاریخ صدور فاکتور مالیاتی</label>
              <input 
                v-model="taxAccounting.taxInvoiceDate"
                type="date" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">وضعیت فاکتور مالیاتی</label>
              <select 
                v-model="taxAccounting.taxInvoiceStatus"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="pending">در انتظار صدور</option>
                <option value="issued">صادر شده</option>
                <option value="cancelled">لغو شده</option>
                <option value="amended">اصلاح شده</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">کد اقتصادی</label>
              <input 
                v-model="taxAccounting.economicCode"
                type="text" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="کد اقتصادی مشتری"
              />
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- اطلاعات حسابداری -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-green-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 7h6m0 10v-3m-3 3h.01M9 17h.01M9 14h.01M12 14h.01M15 11h.01M12 11h.01M9 11h.01M7 21h10a2 2 0 002-2V5a2 2 0 00-2-2H7a2 2 0 00-2 2v14a2 2 0 002 2z"></path>
        </svg>
        اطلاعات حسابداری
      </h3>

      <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
        <!-- کدهای حساب -->
        <div>
          <h4 class="text-md font-medium text-gray-700 mb-3">کدهای حساب</h4>
          <div class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">حساب فروش</label>
              <select 
                v-model="taxAccounting.salesAccount"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="">انتخاب حساب</option>
                <option value="4001">4001 - فروش کالا</option>
                <option value="4002">4002 - فروش خدمات</option>
                <option value="4003">4003 - فروش دیجیتال</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">حساب مالیات</label>
              <select 
                v-model="taxAccounting.taxAccount"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="">انتخاب حساب</option>
                <option value="2201">2201 - مالیات بر ارزش افزوده</option>
                <option value="2202">2202 - مالیات تکلیفی</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">حساب بانکی</label>
              <select 
                v-model="taxAccounting.bankAccount"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="">انتخاب حساب</option>
                <option value="1011">1011 - بانک ملی</option>
                <option value="1012">1012 - بانک ملت</option>
                <option value="1013">1013 - بانک صادرات</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">مرکز هزینه</label>
              <select 
                v-model="taxAccounting.costCenter"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="">انتخاب مرکز هزینه</option>
                <option value="CC001">CC001 - فروشگاه آنلاین</option>
                <option value="CC002">CC002 - انبار</option>
                <option value="CC003">CC003 - پشتیبانی</option>
              </select>
            </div>
          </div>
        </div>

        <!-- تراکنشات حسابداری -->
        <div>
          <h4 class="text-md font-medium text-gray-700 mb-3">تراکنشات حسابداری</h4>
          <div class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">شماره سند حسابداری</label>
              <input 
                v-model="taxAccounting.accountingDocumentNumber"
                type="text" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="DOC-2024-001"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">تاریخ ثبت سند</label>
              <input 
                v-model="taxAccounting.accountingDate"
                type="date" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">وضعیت سند</label>
              <select 
                v-model="taxAccounting.accountingStatus"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="draft">پیش‌نویس</option>
                <option value="posted">ثبت شده</option>
                <option value="cancelled">لغو شده</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">یادداشت حسابداری</label>
              <textarea 
                v-model="taxAccounting.accountingNotes"
                rows="2"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="یادداشت‌های حسابداری..."
              ></textarea>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- خلاصه مالیاتی -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-purple-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
        </svg>
        خلاصه مالیاتی
      </h3>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
        <!-- محاسبات -->
        <div class="space-y-3">
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">جمع کل محصولات:</span>
            <span class="font-medium">{{ formatPrice(taxSummary.subtotal) }} تومان</span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">تخفیف‌ها:</span>
            <span class="font-medium text-green-600">-{{ formatPrice(taxSummary.discounts) }} تومان</span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">مشمول مالیات:</span>
            <span class="font-medium">{{ formatPrice(taxSummary.taxableAmount) }} تومان</span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">مالیات بر ارزش افزوده:</span>
            <span class="font-medium text-red-600">{{ formatPrice(taxSummary.vatAmount) }} تومان</span>
          </div>
          <div class="border-t pt-2 flex justify-between text-lg font-bold">
            <span>مبلغ نهایی:</span>
            <span class="text-blue-600">{{ formatPrice(taxSummary.totalAmount) }} تومان</span>
          </div>
        </div>

        <!-- وضعیت مالیاتی -->
        <div class="space-y-3">
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">وضعیت مالیاتی:</span>
            <span class="font-medium" :class="getTaxStatusColor(taxAccounting.taxInvoiceStatus)">
              {{ getTaxStatusText(taxAccounting.taxInvoiceStatus) }}
            </span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">وضعیت حسابداری:</span>
            <span class="font-medium" :class="getAccountingStatusColor(taxAccounting.accountingStatus)">
              {{ getAccountingStatusText(taxAccounting.accountingStatus) }}
            </span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">معافیت مالیاتی:</span>
            <span class="font-medium">{{ getExemptionText(taxAccounting.taxExemption) }}</span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">نرخ مالیات:</span>
            <span class="font-medium">{{ taxAccounting.regionalTaxRate }}%</span>
          </div>
        </div>
      </div>
    </div>

    <!-- تاریخچه مالیاتی -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-orange-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
        </svg>
        تاریخچه مالیاتی
      </h3>
      
      <div class="space-y-3">
        <div 
          v-for="(history, index) in taxAccounting.history" 
          :key="index"
          class="flex items-center p-3 bg-gray-50 rounded-lg"
        >
          <div class="w-3 h-3 rounded-full mr-3" :class="getHistoryStatusColor(history.type)"></div>
          <div class="flex-1">
            <div class="text-sm font-medium text-gray-900">{{ getHistoryTypeText(history.type) }}</div>
            <div class="text-xs text-gray-500">{{ formatDate(history.date) }} - {{ history.user }}</div>
            <div v-if="history.note" class="text-xs text-gray-600 mt-1">{{ history.note }}</div>
          </div>
          <div class="text-xs text-gray-500">{{ formatTime(history.date) }}</div>
        </div>
      </div>

      <!-- دکمه افزودن تاریخچه -->
      <div class="mt-4">
        <button 
          class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors flex items-center text-sm"
          @click="addTaxHistory"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
          </svg>
          افزودن تاریخچه مالیاتی
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'

// تعریف props
const props = defineProps({
  modelValue: {
    type: Object,
    default: () => ({
      regionalTaxRate: 9,
      taxableAmount: 0,
      calculatedTax: 0,
      taxExemption: '',
      taxInvoiceNumber: '',
      taxInvoiceDate: '',
      taxInvoiceStatus: 'pending',
      economicCode: '',
      salesAccount: '',
      taxAccount: '',
      bankAccount: '',
      costCenter: '',
      accountingDocumentNumber: '',
      accountingDate: '',
      accountingStatus: 'draft',
      accountingNotes: '',
      history: []
    })
  }
})

// تعریف emit
const emit = defineEmits(['update:modelValue'])

// متغیر محلی
const taxAccounting = ref({ ...props.modelValue })

// محاسبات خلاصه مالیاتی
const taxSummary = computed(() => {
  const subtotal = 4200000 // از سفارش
  const discounts = 100000 // از سفارش
  const taxableAmount = subtotal - discounts
  const vatAmount = (taxableAmount * taxAccounting.value.regionalTaxRate) / 100
  const totalAmount = taxableAmount + vatAmount

  return {
    subtotal,
    discounts,
    taxableAmount,
    vatAmount,
    totalAmount
  }
})

// توابع کمکی
const getTaxStatusText = (status) => {
  const statusMap = {
    'pending': 'در انتظار صدور',
    'issued': 'صادر شده',
    'cancelled': 'لغو شده',
    'amended': 'اصلاح شده'
  }
  return statusMap[status] || status
}

const getTaxStatusColor = (status) => {
  const colors = {
    'pending': 'text-yellow-600',
    'issued': 'text-green-600',
    'cancelled': 'text-red-600',
    'amended': 'text-blue-600'
  }
  return colors[status] || 'text-gray-600'
}

const getAccountingStatusText = (status) => {
  const statusMap = {
    'draft': 'پیش‌نویس',
    'posted': 'ثبت شده',
    'cancelled': 'لغو شده'
  }
  return statusMap[status] || status
}

const getAccountingStatusColor = (status) => {
  const colors = {
    'draft': 'text-yellow-600',
    'posted': 'text-green-600',
    'cancelled': 'text-red-600'
  }
  return colors[status] || 'text-gray-600'
}

const getExemptionText = (exemption) => {
  const exemptionMap = {
    '': 'بدون معافیت',
    'export': 'صادرات',
    'diplomatic': 'سفارت‌خانه',
    'charity': 'خیریه',
    'government': 'دولتی',
    'special': 'ویژه'
  }
  return exemptionMap[exemption] || 'نامشخص'
}

const getHistoryTypeText = (type) => {
  const typeMap = {
    'tax_calculation': 'محاسبه مالیات',
    'tax_invoice_issued': 'صدور فاکتور مالیاتی',
    'accounting_entry': 'ثبت حسابداری',
    'tax_exemption': 'اعمال معافیت',
    'tax_amendment': 'اصلاح مالیاتی'
  }
  return typeMap[type] || type
}

const getHistoryStatusColor = (type) => {
  const colors = {
    'tax_calculation': 'bg-blue-400',
    'tax_invoice_issued': 'bg-green-400',
    'accounting_entry': 'bg-purple-400',
    'tax_exemption': 'bg-yellow-400',
    'tax_amendment': 'bg-orange-400'
  }
  return colors[type] || 'bg-gray-400'
}

const formatPrice = (price) => {
  if (!price) return '0'
  return new Intl.NumberFormat('fa-IR').format(price)
}

const formatDate = (date) => {
  if (!date) return 'نامشخص'
  return new Date(date).toLocaleDateString('fa-IR')
}

const formatTime = (date) => {
  if (!date) return ''
  return new Date(date).toLocaleTimeString('fa-IR', { 
    hour: '2-digit', 
    minute: '2-digit' 
  })
}

const addTaxHistory = () => {
  const newEntry = {
    type: 'tax_calculation',
    date: new Date().toISOString(),
    user: 'مدیر سیستم',
    note: 'محاسبه مالیات جدید'
  }
  taxAccounting.value.history.unshift(newEntry)
}

// نظارت بر تغییرات
watch(taxAccounting, (newValue) => {
  emit('update:modelValue', newValue)
}, { deep: true })

// به‌روزرسانی محاسبات مالیاتی
watch(() => taxSummary.value.taxableAmount, (newValue) => {
  taxAccounting.value.taxableAmount = newValue
  taxAccounting.value.calculatedTax = (newValue * taxAccounting.value.regionalTaxRate) / 100
})
</script> 

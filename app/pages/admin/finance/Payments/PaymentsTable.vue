<template>
  <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
    <div class="flex items-center justify-between mb-6">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">جدول پرداخت‌ها</h3>
        <p class="text-sm text-gray-500 mt-1">Payments Table</p>
      </div>
      <div class="flex items-center space-x-2 space-x-reverse">
        <button class="px-3 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200 transition-colors" @click="refreshTable">
          بروزرسانی
        </button>
        <button class="px-3 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 transition-colors" @click="search = ''">
          پاک‌کردن جستجو
        </button>
      </div>
    </div>

    <!-- جستجو -->
    <div class="mb-4">
      <input v-model="search" type="text" placeholder="جستجوی کد پیگیری یا شماره سفارش..." class="w-full px-4 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500" />
    </div>

    <!-- جدول -->
    <div class="overflow-x-auto">
      <table class="min-w-full text-sm text-right">
        <thead>
          <tr class="bg-gray-50">
            <th class="py-3 px-4 font-semibold text-gray-700">کد پیگیری</th>
            <th class="py-3 px-4 font-semibold text-gray-700">تاریخ و ساعت</th>
            <th class="py-3 px-4 font-semibold text-gray-700">نام پرداخت‌کننده</th>
            <th class="py-3 px-4 font-semibold text-gray-700">مبلغ</th>
            <th class="py-3 px-4 font-semibold text-gray-700">روش پرداخت</th>
            <th class="py-3 px-4 font-semibold text-gray-700">وضعیت</th>
            <th class="py-3 px-4 font-semibold text-gray-700">شماره سفارش</th>
            <th class="py-3 px-4 font-semibold text-gray-700">عملیات</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="payment in filteredPayments" :key="payment.id" class="border-b hover:bg-gray-50 transition-colors">
            <td class="py-2 px-4">{{ payment.transactionId }}</td>
            <td class="py-2 px-4">{{ formatDate(String(payment.date || '')) }}</td>
            <td class="py-2 px-4">{{ payment.customerName }}</td>
            <td class="py-2 px-4">{{ formatPrice(payment.amount || 0) }}</td>
            <td class="py-2 px-4">{{ getMethodName(String(payment.paymentMethod || '')) }}</td>
            <td class="py-2 px-4">
              <span :class="statusClass(String(payment.status || ''))">{{ getStatusName(String(payment.status || '')) }}</span>
            </td>
            <td class="py-2 px-4">{{ payment.orderNumber }}</td>
            <td class="py-2 px-4">
              <!-- منوی عملیات -->
              <div class="relative inline-block text-left">
                <button class="p-2 rounded-full hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-blue-500" @click="openActionMenu(typeof payment.id === 'number' ? payment.id : Number(payment.id) || 0)">
                  <svg class="w-5 h-5 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <circle cx="12" cy="6" r="1.5" />
                    <circle cx="12" cy="12" r="1.5" />
                    <circle cx="12" cy="18" r="1.5" />
                  </svg>
                </button>
                <div v-if="actionMenuOpen === payment.id" class="origin-top-left absolute right-0 mt-2 w-44 rounded-md shadow-lg bg-white ring-1 ring-black ring-opacity-5 z-20">
                  <div class="py-1">
                    <button class="w-full text-right px-4 py-2 text-sm text-gray-700 hover:bg-gray-100" @click="showManualConfirm(payment)">تایید پرداخت دستی</button>
                    <button class="w-full text-right px-4 py-2 text-sm text-gray-700 hover:bg-gray-100" @click="showRefund(payment)">مرجوع کردن وجه</button>
                    <button class="w-full text-right px-4 py-2 text-sm text-gray-700 hover:bg-gray-100" @click="showSendReceipt(payment)">ارسال رسید</button>
                    <button class="w-full text-right px-4 py-2 text-sm text-gray-700 hover:bg-gray-100" @click="showNote(payment)">یادداشت</button>
                  </div>
                </div>
              </div>
            </td>
          </tr>
          <tr v-if="filteredPayments.length === 0">
            <td colspan="8" class="text-center py-8 text-gray-400">هیچ پرداختی یافت نشد.</td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- مودال تایید پرداخت دستی -->
    <div v-if="modal.manualConfirm" class="fixed inset-0 z-30 flex items-center justify-center bg-black bg-opacity-40">
      <div class="bg-white rounded-xl p-6 w-full max-w-md shadow-lg">
        <h4 class="text-lg font-bold mb-4">تایید پرداخت دستی</h4>
        <p class="mb-4 text-gray-700">آیا از تایید پرداخت کارت به کارت برای <span class="font-semibold">{{ selectedPayment?.customerName }}</span> با مبلغ <span class="font-semibold">{{ formatPrice(selectedPayment?.amount) }}</span> مطمئن هستید؟</p>
        <div class="flex justify-end space-x-2 space-x-reverse">
          <button class="px-4 py-2 bg-green-500 text-white rounded-lg hover:bg-green-600" @click="confirmManualPayment">تایید</button>
          <button class="px-4 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200" @click="closeModal">انصراف</button>
        </div>
      </div>
    </div>

    <!-- مودال مرجوع کردن وجه -->
    <div v-if="modal.refund" class="fixed inset-0 z-30 flex items-center justify-center bg-black bg-opacity-40">
      <div class="bg-white rounded-xl p-6 w-full max-w-md shadow-lg">
        <h4 class="text-lg font-bold mb-4">مرجوع کردن وجه</h4>
        <p class="mb-4 text-gray-700">آیا از مرجوع کردن وجه برای <span class="font-semibold">{{ selectedPayment?.customerName }}</span> با مبلغ <span class="font-semibold">{{ formatPrice(selectedPayment?.amount) }}</span> مطمئن هستید؟</p>
        <div class="flex justify-end space-x-2 space-x-reverse">
          <button class="px-4 py-2 bg-red-500 text-white rounded-lg hover:bg-red-600" @click="confirmRefund">مرجوع</button>
          <button class="px-4 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200" @click="closeModal">انصراف</button>
        </div>
      </div>
    </div>

    <!-- مودال ارسال رسید -->
    <div v-if="modal.sendReceipt" class="fixed inset-0 z-30 flex items-center justify-center bg-black bg-opacity-40">
      <div class="bg-white rounded-xl p-6 w-full max-w-md shadow-lg">
        <h4 class="text-lg font-bold mb-4">ارسال رسید به مشتری</h4>
        <p class="mb-4 text-gray-700">آیا رسید پرداخت برای <span class="font-semibold">{{ selectedPayment?.customerName }}</span> ارسال شود؟</p>
        <div class="flex justify-end space-x-2 space-x-reverse">
          <button class="px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600" @click="confirmSendReceipt">ارسال</button>
          <button class="px-4 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200" @click="closeModal">انصراف</button>
        </div>
      </div>
    </div>

    <!-- مودال یادداشت -->
    <div v-if="modal.note" class="fixed inset-0 z-30 flex items-center justify-center bg-black bg-opacity-40">
      <div class="bg-white rounded-xl p-6 w-full max-w-md shadow-lg">
        <h4 class="text-lg font-bold mb-4">یادداشت تراکنش</h4>
        <textarea v-model="noteText" rows="4" class="w-full border border-gray-300 rounded-lg p-2 mb-4 text-sm" placeholder="یادداشت خود را وارد کنید..."></textarea>
        <div class="flex justify-end space-x-2 space-x-reverse">
          <button class="px-4 py-2 bg-purple-500 text-white rounded-lg hover:bg-purple-600" @click="saveNote">ذخیره</button>
          <button class="px-4 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200" @click="closeModal">انصراف</button>
        </div>
      </div>
    </div>

    <!-- پیام موفقیت -->
    <div v-if="successMessage" class="fixed bottom-6 right-6 z-40 bg-green-500 text-white px-6 py-3 rounded-lg shadow-lg">
      {{ successMessage }}
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'

interface Payment {
  id?: number | string
  amount?: number
  status?: string
  transactionId?: string
  orderNumber?: string | number
  date?: string
  customerName?: string
  paymentMethod?: string
  note?: string
  [key: string]: unknown
}
const props = defineProps<{ payments: Payment[] }>()

const search = ref('')
const actionMenuOpen = ref<number|null>(null)
const modal = ref({ manualConfirm: false, refund: false, sendReceipt: false, note: false })

const selectedPayment = ref<Payment | null>(null)
const noteText = ref('')
const successMessage = ref('')

const filteredPayments = computed(() => {
  if (!search.value) return props.payments
  return props.payments.filter(p => {
    const transactionId = String(p.transactionId || '')
    const orderNumber = String(p.orderNumber || '')
    return (transactionId && transactionId.includes(search.value)) ||
           (orderNumber && orderNumber.includes(search.value))
  })
})

const formatDate = (date: string) => {
  const d = new Date(date)
  return d.toLocaleDateString('fa-IR') + ' ' + d.toLocaleTimeString('fa-IR', { hour: '2-digit', minute: '2-digit' })
}
const formatPrice = (price: number) => {
  if (price >= 1000000) return (price / 1000000).toFixed(1) + 'M'
  if (price >= 1000) return (price / 1000).toFixed(0) + 'K'
  return new Intl.NumberFormat('fa-IR').format(price)
}
const getMethodName = (method: string) => {
  switch (method) {
    case 'zarinpal': return 'زرین‌پال'
    case 'nextpay': return 'نکست‌پی'
    case 'card_to_card': return 'کارت به کارت'
    case 'wallet': return 'کیف پول'
    case 'cash': return 'نقدی'
    default: return method
  }
}
const getStatusName = (status: string) => {
  switch (status) {
    case 'success': return 'موفق'
    case 'failed': return 'ناموفق'
    case 'pending': return 'در انتظار'
    case 'refunded': return 'مرجوع'
    default: return status
  }
}
const statusClass = (status: string) => {
  switch (status) {
    case 'success': return 'inline-block px-2 py-1 rounded-full bg-green-100 text-green-700 text-xs'
    case 'failed': return 'inline-block px-2 py-1 rounded-full bg-red-100 text-red-700 text-xs'
    case 'pending': return 'inline-block px-2 py-1 rounded-full bg-yellow-100 text-yellow-700 text-xs'
    case 'refunded': return 'inline-block px-2 py-1 rounded-full bg-purple-100 text-purple-700 text-xs'
    default: return 'inline-block px-2 py-1 rounded-full bg-gray-100 text-gray-700 text-xs'
  }
}

function openActionMenu(id: number) {
  actionMenuOpen.value = actionMenuOpen.value === id ? null : id
}
function closeModal() {
  modal.value = { manualConfirm: false, refund: false, sendReceipt: false, note: false }
  selectedPayment.value = null
  noteText.value = ''
  actionMenuOpen.value = null
}
function showManualConfirm(payment: Payment) {
  selectedPayment.value = payment
  modal.value = { manualConfirm: true, refund: false, sendReceipt: false, note: false }
  actionMenuOpen.value = null
}
function showRefund(payment: Payment) {
  selectedPayment.value = payment
  modal.value = { manualConfirm: false, refund: true, sendReceipt: false, note: false }
  actionMenuOpen.value = null
}
function showSendReceipt(payment: Payment) {
  selectedPayment.value = payment
  modal.value = { manualConfirm: false, refund: false, sendReceipt: true, note: false }
  actionMenuOpen.value = null
}
function showNote(payment: Payment) {
  selectedPayment.value = payment
  noteText.value = String(payment.note || '')
  modal.value = { manualConfirm: false, refund: false, sendReceipt: false, note: true }
  actionMenuOpen.value = null
}
function confirmManualPayment() {
  // TODO: فراخوانی API تایید پرداخت دستی
  successMessage.value = 'پرداخت دستی با موفقیت تایید شد.'
  closeModal()
}
function confirmRefund() {
  // TODO: فراخوانی API مرجوع کردن وجه
  successMessage.value = 'وجه با موفقیت مرجوع شد.'
  closeModal()
}
function confirmSendReceipt() {
  // TODO: فراخوانی API ارسال رسید
  successMessage.value = 'رسید پرداخت برای مشتری ارسال شد.'
  closeModal()
}
function saveNote() {
  // TODO: ذخیره یادداشت برای تراکنش
  successMessage.value = 'یادداشت تراکنش ذخیره شد.'
  closeModal()
}
function refreshTable() {
  // TODO: بروزرسانی داده‌ها از سرور
  successMessage.value = 'جدول بروزرسانی شد.'
}

watch(successMessage, (val) => {
  if (val) setTimeout(() => (successMessage.value = ''), 2500)
})
</script>

<!--
  جدول پرداخت‌ها با ستون عملیات و مودال‌های عملیاتی
  - تایید پرداخت دستی
  - مرجوع کردن وجه
  - ارسال رسید
  - یادداشت‌گذاری
  - پیام موفقیت
  توضیحات کامل به فارسی در کد
--> 

<template>
  <div class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
    <div class="relative top-20 mx-auto p-5 border w-11/12 max-w-2xl shadow-lg rounded-md bg-white">
      <!-- هدر مودال -->
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-lg font-semibold text-gray-900">ثبت تراکنش جدید</h3>
        <button
          @click="$emit('close')"
          class="text-gray-400 hover:text-gray-600"
        >
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
        </button>
      </div>

      <!-- فرم ثبت تراکنش -->
      <form @submit.prevent="createTransaction" class="space-y-6">
        <!-- نوع تراکنش -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            نوع تراکنش <span class="text-red-500">*</span>
          </label>
          <select
            v-model="form.type"
            required
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          >
            <option value="">انتخاب کنید</option>
            <option value="purchase">خرید کارت</option>
            <option value="usage">استفاده کارت</option>
            <option value="refund">بازپرداخت</option>
            <option value="expiry">انقضا</option>
          </select>
        </div>

        <!-- مبلغ -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            مبلغ (تومان) <span class="text-red-500">*</span>
          </label>
          <input
            v-model="form.amount"
            type="number"
            required
            min="0"
            step="1000"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="مثال: 500000"
          />
        </div>

        <!-- کد گیفت کارت -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            کد گیفت کارت
          </label>
          <input
            v-model="form.giftCardCode"
            type="text"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="مثال: GC-2024-001"
          />
        </div>

        <!-- اطلاعات کاربر -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              نام کاربر <span class="text-red-500">*</span>
            </label>
            <input
              v-model="form.user"
              type="text"
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="مثال: احمد محمدی"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              ایمیل کاربر
            </label>
            <input
              v-model="form.userEmail"
              type="email"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="مثال: ahmad@example.com"
            />
          </div>
        </div>

        <!-- روش پرداخت -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            روش پرداخت <span class="text-red-500">*</span>
          </label>
          <select
            v-model="form.paymentMethod"
            required
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          >
            <option value="">انتخاب کنید</option>
            <option value="online">پرداخت آنلاین</option>
            <option value="gift_card">گیفت کارت</option>
            <option value="refund">بازپرداخت</option>
            <option value="system">سیستم</option>
          </select>
        </div>

        <!-- توضیحات -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            توضیحات
          </label>
          <textarea
            v-model="form.description"
            rows="3"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="توضیحات مربوط به تراکنش..."
          ></textarea>
        </div>

        <!-- شماره سفارش -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            شماره سفارش
          </label>
          <input
            v-model="form.orderId"
            type="text"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="مثال: ORD-2024-001"
          />
        </div>

        <!-- وضعیت -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            وضعیت <span class="text-red-500">*</span>
          </label>
          <select
            v-model="form.status"
            required
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          >
            <option value="">انتخاب کنید</option>
            <option value="successful">موفق</option>
            <option value="failed">ناموفق</option>
            <option value="pending">در انتظار</option>
          </select>
        </div>

        <!-- پاسخ درگاه پرداخت -->
        <div v-if="form.status === 'failed'">
          <label class="block text-sm font-medium text-gray-700 mb-2">
            دلیل ناموفق بودن
          </label>
          <select
            v-model="form.gatewayResponse"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          >
            <option value="">انتخاب کنید</option>
            <option value="insufficient_funds">موجودی ناکافی</option>
            <option value="card_declined">کارت رد شد</option>
            <option value="network_error">خطای شبکه</option>
            <option value="timeout">زمان انتظار تمام شد</option>
            <option value="invalid_card">کارت نامعتبر</option>
          </select>
        </div>

        <!-- شناسه تراکنش درگاه -->
        <div v-if="form.status === 'successful'">
          <label class="block text-sm font-medium text-gray-700 mb-2">
            شناسه تراکنش درگاه
          </label>
          <input
            v-model="form.gatewayTransactionId"
            type="text"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="مثال: GT-001"
          />
        </div>

        <!-- دکمه‌های عملیات -->
        <div class="flex justify-end space-x-3 space-x-reverse">
          <button
            type="button"
            @click="$emit('close')"
            class="px-4 py-2 bg-gray-300 text-gray-700 text-sm font-medium rounded-lg hover:bg-gray-400 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
          >
            انصراف
          </button>
          <button
            type="submit"
            class="px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
          >
            ثبت تراکنش
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
// تعریف props و emits
const emit = defineEmits(['close', 'created'])

// تعریف متغیرهای reactive
const form = ref({
  type: '',
  amount: '',
  giftCardCode: '',
  user: '',
  userEmail: '',
  paymentMethod: '',
  description: '',
  orderId: '',
  status: 'successful',
  gatewayResponse: 'success',
  gatewayTransactionId: ''
})

// تابع ایجاد تراکنش
const createTransaction = () => {
  // اعتبارسنجی فرم
  if (!form.value.type || !form.value.amount || !form.value.user || !form.value.paymentMethod || !form.value.status) {
    alert('لطفاً تمام فیلدهای اجباری را پر کنید')
    return
  }

  // اعتبارسنجی مبلغ
  if (form.value.amount <= 0) {
    alert('مبلغ باید بیشتر از صفر باشد')
    return
  }

  // اعتبارسنجی ایمیل
  if (form.value.userEmail && !isValidEmail(form.value.userEmail)) {
    alert('لطفاً ایمیل معتبر وارد کنید')
    return
  }

  // ایجاد تراکنش جدید
  const newTransaction = {
    type: form.value.type,
    amount: parseInt(form.value.amount),
    giftCardCode: form.value.giftCardCode || null,
    user: form.value.user,
    userEmail: form.value.userEmail || null,
    paymentMethod: form.value.paymentMethod,
    description: form.value.description || '',
    orderId: form.value.orderId || null,
    status: form.value.status,
    gatewayResponse: form.value.gatewayResponse,
    gatewayTransactionId: form.value.gatewayTransactionId || null
  }

  // ارسال به کامپوننت والد
  emit('created', newTransaction)
}

// تابع اعتبارسنجی ایمیل
const isValidEmail = (email) => {
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  return emailRegex.test(email)
}
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 

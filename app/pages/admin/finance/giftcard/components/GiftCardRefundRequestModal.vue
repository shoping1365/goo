<template>
  <div class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
    <div class="relative top-20 mx-auto p-5 border w-11/12 max-w-2xl shadow-lg rounded-md bg-white">
      <!-- هدر مودال -->
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-lg font-semibold text-gray-900">درخواست بازپرداخت جدید</h3>
        <button
          class="text-gray-400 hover:text-gray-600"
          @click="$emit('close')"
        >
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
        </button>
      </div>

      <!-- فرم درخواست بازپرداخت -->
      <form class="space-y-6" @submit.prevent="createRefundRequest">
        <!-- کد گیفت کارت -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            کد گیفت کارت <span class="text-red-500">*</span>
          </label>
          <div class="flex gap-2">
            <input
              v-model="form.giftCardCode"
              type="text"
              required
              class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="مثال: GC-2024-001"
            />
            <button
              type="button"
              class="px-4 py-2 bg-gray-600 text-white text-sm font-medium rounded-lg hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
              @click="searchGiftCard"
            >
              جستجو
            </button>
          </div>
        </div>

        <!-- اطلاعات گیفت کارت -->
        <div v-if="giftCardInfo" class="bg-gray-50 p-6 rounded-lg">
          <h4 class="text-sm font-medium text-gray-900 mb-2">اطلاعات گیفت کارت</h4>
          <div class="grid grid-cols-2 gap-6 text-sm">
            <div>
              <span class="text-gray-600">مبلغ اصلی:</span>
              <span class="font-medium text-gray-900 mr-2">{{ formatCurrency(giftCardInfo.amount) }}</span>
            </div>
            <div>
              <span class="text-gray-600">مبلغ باقی‌مانده:</span>
              <span class="font-medium text-gray-900 mr-2">{{ formatCurrency(giftCardInfo.remainingAmount) }}</span>
            </div>
            <div>
              <span class="text-gray-600">وضعیت:</span>
              <span class="font-medium text-gray-900 mr-2">{{ getStatusText(giftCardInfo.status) }}</span>
            </div>
            <div>
              <span class="text-gray-600">تاریخ انقضا:</span>
              <span class="font-medium text-gray-900 mr-2">{{ formatDate(giftCardInfo.expiresAt) }}</span>
            </div>
          </div>
        </div>

        <!-- مبلغ بازپرداخت -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            مبلغ بازپرداخت (تومان) <span class="text-red-500">*</span>
          </label>
          <input
            v-model="form.amount"
            type="number"
            required
            min="0"
            step="1000"
            :max="giftCardInfo ? giftCardInfo.remainingAmount : undefined"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="مبلغ بازپرداخت را وارد کنید"
          />
          <p v-if="giftCardInfo" class="text-xs text-gray-500 mt-1">
            حداکثر مبلغ قابل بازپرداخت: {{ formatCurrency(giftCardInfo.remainingAmount) }}
          </p>
        </div>

        <!-- اطلاعات درخواست کننده -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              نام درخواست کننده <span class="text-red-500">*</span>
            </label>
            <input
              v-model="form.requester"
              type="text"
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="مثال: احمد محمدی"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              ایمیل درخواست کننده
            </label>
            <input
              v-model="form.requesterEmail"
              type="email"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="مثال: ahmad@example.com"
            />
          </div>
        </div>

        <!-- دلیل بازپرداخت -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            دلیل بازپرداخت <span class="text-red-500">*</span>
          </label>
          <select
            v-model="form.reason"
            required
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          >
            <option value="">انتخاب کنید</option>
            <option value="duplicate_purchase">خرید تکراری</option>
            <option value="wrong_amount">مبلغ اشتباه</option>
            <option value="technical_issue">مشکل فنی</option>
            <option value="customer_request">درخواست مشتری</option>
            <option value="expired_card">کارت منقضی شده</option>
            <option value="other">سایر</option>
          </select>
        </div>

        <!-- توضیحات -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            توضیحات <span class="text-red-500">*</span>
          </label>
          <textarea
            v-model="form.description"
            rows="4"
            required
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="توضیحات کامل دلیل بازپرداخت..."
          ></textarea>
        </div>

        <!-- شناسه تراکنش -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            شناسه تراکنش اصلی
          </label>
          <input
            v-model="form.transactionId"
            type="text"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="مثال: TXN-001"
          />
        </div>

        <!-- پیوست فایل -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            پیوست فایل (اختیاری)
          </label>
          <input
            type="file"
            accept=".pdf,.jpg,.jpeg,.png,.doc,.docx"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            @change="handleFileUpload"
          />
          <p class="text-xs text-gray-500 mt-1">
            فرمت‌های مجاز: PDF, JPG, PNG, DOC, DOCX (حداکثر 5MB)
          </p>
        </div>

        <!-- دکمه‌های عملیات -->
        <div class="flex justify-end space-x-3 space-x-reverse">
          <button
            type="button"
            class="px-4 py-2 bg-gray-300 text-gray-700 text-sm font-medium rounded-lg hover:bg-gray-400 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
            @click="$emit('close')"
          >
            انصراف
          </button>
          <button
            type="submit"
            class="px-4 py-2 bg-red-600 text-white text-sm font-medium rounded-lg hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2"
          >
            ثبت درخواست بازپرداخت
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
  giftCardCode: '',
  amount: '',
  requester: '',
  requesterEmail: '',
  reason: '',
  description: '',
  transactionId: '',
  attachment: null
})

const giftCardInfo = ref(null)

// تابع جستجوی گیفت کارت
const searchGiftCard = () => {
  if (!form.value.giftCardCode) {
    alert('لطفاً کد گیفت کارت را وارد کنید')
    return
  }

  // اینجا کد جستجوی گیفت کارت از API را اضافه کنید
  // برای نمونه، داده‌های نمونه استفاده می‌کنیم
  giftCardInfo.value = {
    code: form.value.giftCardCode,
    amount: 500000,
    remainingAmount: 500000,
    status: 'active',
    expiresAt: new Date(Date.now() + 2592000000),
    recipient: 'فاطمه احمدی',
    recipientEmail: 'fateme@example.com'
  }

  // تنظیم مبلغ پیش‌فرض
  form.value.amount = giftCardInfo.value.remainingAmount
}

// تابع ایجاد درخواست بازپرداخت
const createRefundRequest = () => {
  // اعتبارسنجی فرم
  if (!form.value.giftCardCode || !form.value.amount || !form.value.requester || !form.value.reason || !form.value.description) {
    alert('لطفاً تمام فیلدهای اجباری را پر کنید')
    return
  }

  // اعتبارسنجی مبلغ
  if (form.value.amount <= 0) {
    alert('مبلغ باید بیشتر از صفر باشد')
    return
  }

  if (giftCardInfo.value && form.value.amount > giftCardInfo.value.remainingAmount) {
    alert('مبلغ بازپرداخت نمی‌تواند بیشتر از مبلغ باقی‌مانده کارت باشد')
    return
  }

  // اعتبارسنجی ایمیل
  if (form.value.requesterEmail && !isValidEmail(form.value.requesterEmail)) {
    alert('لطفاً ایمیل معتبر وارد کنید')
    return
  }

  // ایجاد درخواست بازپرداخت جدید
  const newRefundRequest = {
    giftCardCode: form.value.giftCardCode,
    amount: parseInt(form.value.amount),
    requester: form.value.requester,
    requesterEmail: form.value.requesterEmail || null,
    reason: form.value.reason,
    description: form.value.description,
    transactionId: form.value.transactionId || null,
    attachment: form.value.attachment
  }

  // ارسال به کامپوننت والد
  emit('created', newRefundRequest)
}

// تابع مدیریت آپلود فایل
const handleFileUpload = (event) => {
  const file = event.target.files[0]
  if (file) {
    // اعتبارسنجی اندازه فایل (5MB)
    if (file.size > 5 * 1024 * 1024) {
      alert('حجم فایل نمی‌تواند بیشتر از 5MB باشد')
      event.target.value = ''
      return
    }

    // اعتبارسنجی نوع فایل
    const allowedTypes = ['application/pdf', 'image/jpeg', 'image/jpg', 'image/png', 'application/msword', 'application/vnd.openxmlformats-officedocument.wordprocessingml.document']
    if (!allowedTypes.includes(file.type)) {
      alert('نوع فایل مجاز نیست')
      event.target.value = ''
      return
    }

    form.value.attachment = file
  }
}

// تابع اعتبارسنجی ایمیل
const isValidEmail = (email) => {
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  return emailRegex.test(email)
}

// تابع فرمت کردن مبلغ
const formatCurrency = (amount) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

// تابع فرمت کردن تاریخ
const formatDate = (date) => {
  return new Intl.DateTimeFormat('fa-IR', {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  }).format(date)
}

// تابع دریافت متن وضعیت
const getStatusText = (status) => {
  const statusMap = {
    active: 'فعال',
    used: 'استفاده شده',
    expired: 'منقضی شده',
    blocked: 'مسدود شده'
  }
  return statusMap[status] || status
}
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 

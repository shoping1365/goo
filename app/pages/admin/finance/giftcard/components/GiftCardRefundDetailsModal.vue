<template>
  <div class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
    <div class="relative top-10 mx-auto p-5 border w-11/12 max-w-4xl shadow-lg rounded-md bg-white">
      <!-- هدر مودال -->
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-lg font-semibold text-gray-900">جزئیات بازپرداخت #{{ refund.id }}</h3>
        <button
          class="text-gray-400 hover:text-gray-600"
          @click="$emit('close')"
        >
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
        </button>
      </div>

      <!-- محتوای مودال -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
        <!-- اطلاعات اصلی بازپرداخت -->
        <div class="space-y-6">
          <!-- کارت اطلاعات اصلی -->
          <div class="bg-white border border-gray-200 rounded-lg p-6">
            <h4 class="text-lg font-semibold text-gray-900 mb-4">اطلاعات اصلی</h4>
            <div class="space-y-4">
              <div class="flex justify-between items-center">
                <span class="text-sm font-medium text-gray-600">شناسه بازپرداخت:</span>
                <span class="text-sm text-gray-900 font-mono">#{{ refund.id }}</span>
              </div>
              <div class="flex justify-between items-center">
                <span class="text-sm font-medium text-gray-600">کد گیفت کارت:</span>
                <span class="text-sm font-mono text-gray-900">{{ refund.giftCardCode }}</span>
              </div>
              <div class="flex justify-between items-center">
                <span class="text-sm font-medium text-gray-600">مبلغ بازپرداخت:</span>
                <span class="text-lg font-bold text-gray-900">{{ formatCurrency(refund.amount) }}</span>
              </div>
              <div class="flex justify-between items-center">
                <span class="text-sm font-medium text-gray-600">وضعیت:</span>
                <span
                  :class="{
                    'bg-yellow-100 text-yellow-800': refund.status === 'pending',
                    'bg-green-100 text-green-800': refund.status === 'approved',
                    'bg-red-100 text-red-800': refund.status === 'rejected',
                    'bg-blue-100 text-blue-800': refund.status === 'completed'
                  }"
                  class="px-2 py-1 text-xs font-medium rounded-full"
                >
                  {{ getStatusText(refund.status) }}
                </span>
              </div>
              <div class="flex justify-between items-center">
                <span class="text-sm font-medium text-gray-600">تاریخ درخواست:</span>
                <span class="text-sm text-gray-900">{{ formatDateTime(refund.createdAt) }}</span>
              </div>
            </div>
          </div>

          <!-- اطلاعات درخواست کننده -->
          <div class="bg-white border border-gray-200 rounded-lg p-6">
            <h4 class="text-lg font-semibold text-gray-900 mb-4">اطلاعات درخواست کننده</h4>
            <div class="space-y-4">
              <div class="flex items-center">
                <div class="flex-shrink-0 h-12 w-12">
                  <div class="h-12 w-12 rounded-full bg-gray-300 flex items-center justify-center">
                    <span class="text-lg font-medium text-gray-700">
                      {{ refund.requester.charAt(0) }}
                    </span>
                  </div>
                </div>
                <div class="mr-4">
                  <div class="text-lg font-medium text-gray-900">{{ refund.requester }}</div>
                  <div class="text-sm text-gray-500">{{ refund.requesterEmail || 'ایمیل ثبت نشده' }}</div>
                </div>
              </div>
            </div>
          </div>

          <!-- اطلاعات دلیل بازپرداخت -->
          <div class="bg-white border border-gray-200 rounded-lg p-6">
            <h4 class="text-lg font-semibold text-gray-900 mb-4">دلیل بازپرداخت</h4>
            <div class="space-y-4">
              <div class="flex justify-between items-center">
                <span class="text-sm font-medium text-gray-600">دلیل:</span>
                <span class="text-sm text-gray-900">{{ getReasonText(refund.reason) }}</span>
              </div>
              <div>
                <span class="text-sm font-medium text-gray-600">توضیحات:</span>
                <p class="text-sm text-gray-900 mt-2">{{ refund.description }}</p>
              </div>
            </div>
          </div>
        </div>

        <!-- اطلاعات فنی و پرداخت -->
        <div class="space-y-6">
          <!-- اطلاعات تراکنش -->
          <div class="bg-white border border-gray-200 rounded-lg p-6">
            <h4 class="text-lg font-semibold text-gray-900 mb-4">اطلاعات تراکنش</h4>
            <div class="space-y-4">
              <div v-if="refund.transactionId" class="flex justify-between items-center">
                <span class="text-sm font-medium text-gray-600">شناسه تراکنش:</span>
                <span class="text-sm font-mono text-gray-900">{{ refund.transactionId }}</span>
              </div>
              <div v-if="refund.approvedBy" class="flex justify-between items-center">
                <span class="text-sm font-medium text-gray-600">تایید کننده:</span>
                <span class="text-sm text-gray-900">{{ refund.approvedBy }}</span>
              </div>
              <div v-if="refund.approvedAt" class="flex justify-between items-center">
                <span class="text-sm font-medium text-gray-600">تاریخ تایید:</span>
                <span class="text-sm text-gray-900">{{ formatDateTime(refund.approvedAt) }}</span>
              </div>
              <div v-if="refund.rejectionReason" class="flex justify-between items-center">
                <span class="text-sm font-medium text-gray-600">دلیل رد:</span>
                <span class="text-sm text-red-600">{{ refund.rejectionReason }}</span>
              </div>
            </div>
          </div>

          <!-- آمار و تحلیل -->
          <div class="bg-white border border-gray-200 rounded-lg p-6">
            <h4 class="text-lg font-semibold text-gray-900 mb-4">آمار و تحلیل</h4>
            <div class="grid grid-cols-2 gap-6">
              <div class="text-center p-6 bg-gray-50 rounded-lg">
                <div class="text-2xl font-bold text-red-600">{{ refund.amount }}</div>
                <div class="text-xs text-gray-500">مبلغ بازپرداخت (تومان)</div>
              </div>
              <div class="text-center p-6 bg-gray-50 rounded-lg">
                <div class="text-2xl font-bold text-blue-600">{{ getRefundAge(refund.createdAt) }}</div>
                <div class="text-xs text-gray-500">روز از درخواست</div>
              </div>
            </div>
          </div>

          <!-- تاریخچه تغییرات -->
          <div class="bg-white border border-gray-200 rounded-lg p-6">
            <h4 class="text-lg font-semibold text-gray-900 mb-4">تاریخچه تغییرات</h4>
            <div class="space-y-3">
              <div class="flex items-center">
                <div class="w-2 h-2 bg-blue-400 rounded-full"></div>
                <div class="mr-3">
                  <div class="text-sm font-medium text-gray-900">درخواست بازپرداخت ثبت شد</div>
                  <div class="text-xs text-gray-500">{{ formatDateTime(refund.createdAt) }}</div>
                </div>
              </div>
              <div v-if="refund.status === 'approved'" class="flex items-center">
                <div class="w-2 h-2 bg-green-400 rounded-full"></div>
                <div class="mr-3">
                  <div class="text-sm font-medium text-gray-900">بازپرداخت تایید شد</div>
                  <div class="text-xs text-gray-500">{{ formatDateTime(refund.approvedAt) }}</div>
                </div>
              </div>
              <div v-if="refund.status === 'rejected'" class="flex items-center">
                <div class="w-2 h-2 bg-red-400 rounded-full"></div>
                <div class="mr-3">
                  <div class="text-sm font-medium text-gray-900">بازپرداخت رد شد</div>
                  <div class="text-xs text-gray-500">{{ formatDateTime(refund.createdAt) }}</div>
                </div>
              </div>
              <div v-if="refund.status === 'completed'" class="flex items-center">
                <div class="w-2 h-2 bg-purple-400 rounded-full"></div>
                <div class="mr-3">
                  <div class="text-sm font-medium text-gray-900">بازپرداخت انجام شد</div>
                  <div class="text-xs text-gray-500">{{ formatDateTime(refund.createdAt) }}</div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- دکمه‌های عملیات -->
      <div class="flex justify-end space-x-3 space-x-reverse mt-8 pt-6 border-t border-gray-200">
        <button
          class="px-4 py-2 bg-gray-600 text-white text-sm font-medium rounded-lg hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
          @click="printRefund"
        >
          <svg class="w-4 h-4 mr-2 inline" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 17h2a2 2 0 002-2v-4a2 2 0 00-2-2H5a2 2 0 00-2 2v4a2 2 0 002 2h2m2 4h6a2 2 0 002-2v-4a2 2 0 00-2-2H9a2 2 0 00-2 2v4a2 2 0 002 2zm8-12V5a2 2 0 00-2-2H9a2 2 0 00-2 2v4h10z"></path>
          </svg>
          چاپ
        </button>
        <button
          class="px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
          @click="exportRefund"
        >
          <svg class="w-4 h-4 mr-2 inline" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
          </svg>
          دانلود
        </button>
        <button
          v-if="refund.status === 'pending'"
          class="px-4 py-2 bg-green-600 text-white text-sm font-medium rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2"
          @click="approveRefund"
        >
          <svg class="w-4 h-4 mr-2 inline" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
          </svg>
          تایید
        </button>
        <button
          v-if="refund.status === 'pending'"
          class="px-4 py-2 bg-red-600 text-white text-sm font-medium rounded-lg hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2"
          @click="rejectRefund"
        >
          <svg class="w-4 h-4 mr-2 inline" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
          رد
        </button>
        <button
          class="px-4 py-2 bg-gray-300 text-gray-700 text-sm font-medium rounded-lg hover:bg-gray-400 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
          @click="$emit('close')"
        >
          بستن
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
// تعریف props و emits
const props = defineProps({
  refund: {
    type: Object,
    required: true
  }
})

const _emit = defineEmits(['close'])

// تابع دریافت متن وضعیت
const getStatusText = (status) => {
  const statusMap = {
    pending: 'در انتظار',
    approved: 'تایید شده',
    rejected: 'رد شده',
    completed: 'انجام شده'
  }
  return statusMap[status] || status
}

// تابع دریافت متن دلیل
const getReasonText = (reason) => {
  const reasonMap = {
    duplicate_purchase: 'خرید تکراری',
    wrong_amount: 'مبلغ اشتباه',
    technical_issue: 'مشکل فنی',
    customer_request: 'درخواست مشتری',
    expired_card: 'کارت منقضی شده',
    other: 'سایر'
  }
  return reasonMap[reason] || reason
}

// تابع فرمت کردن مبلغ
const formatCurrency = (amount) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

// تابع فرمت کردن تاریخ و زمان
const formatDateTime = (date) => {
  return new Intl.DateTimeFormat('fa-IR', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  }).format(date)
}

// تابع محاسبه سن درخواست
const getRefundAge = (timestamp) => {
  const now = new Date()
  const refundDate = new Date(timestamp)
  const diffTime = Math.abs(now - refundDate)
  const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24))
  return diffDays
}

// تابع چاپ بازپرداخت
const printRefund = () => {
  const printWindow = window.open('', '_blank')
  const printContent = `
    <!DOCTYPE html>
    <html dir="rtl" lang="fa">
    <head>
      <meta charset="UTF-8">
      <title>بازپرداخت #${props.refund.id}</title>
      <style>
        body { 
          font-family: 'IRANSans', Arial, sans-serif; 
          margin: 0; 
          padding: 20px; 
          background-color: #f9fafb;
        }
        .container {
          max-width: 800px;
          margin: 0 auto;
          background: white;
          padding: 30px;
          border-radius: 8px;
          box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
        }
        .header {
          text-align: center;
          border-bottom: 2px solid #e5e7eb;
          padding-bottom: 20px;
          margin-bottom: 30px;
        }
        .refund-id {
          font-size: 24px;
          font-weight: bold;
          color: #374151;
        }
        .info-grid {
          display: grid;
          grid-template-columns: 1fr 1fr;
          gap: 20px;
          margin-bottom: 30px;
        }
        .info-item {
          padding: 15px;
          background-color: #f9fafb;
          border-radius: 6px;
        }
        .info-label {
          font-size: 12px;
          color: #6b7280;
          margin-bottom: 5px;
        }
        .info-value {
          font-size: 16px;
          font-weight: 600;
          color: #111827;
        }
        .amount {
          font-size: 28px;
          font-weight: bold;
          color: #dc2626;
          text-align: center;
          padding: 20px;
          background-color: #fef2f2;
          border-radius: 8px;
          margin: 20px 0;
        }
        @media print {
          body { background-color: white; }
          .container { box-shadow: none; }
        }
      </style>
    </head>
    <body>
      <div class="container">
        <div class="header">
          <div class="refund-id">بازپرداخت #${props.refund.id}</div>
          <div style="margin-top: 10px; color: #6b7280;">
            ${formatDateTime(props.refund.createdAt)}
          </div>
        </div>
        
        <div class="amount">
          ${formatCurrency(props.refund.amount)}
        </div>
        
        <div class="info-grid">
          <div class="info-item">
            <div class="info-label">کد گیفت کارت</div>
            <div class="info-value">${props.refund.giftCardCode}</div>
          </div>
          <div class="info-item">
            <div class="info-label">وضعیت</div>
            <div class="info-value">${getStatusText(props.refund.status)}</div>
          </div>
          <div class="info-item">
            <div class="info-label">درخواست کننده</div>
            <div class="info-value">${props.refund.requester}</div>
          </div>
          <div class="info-item">
            <div class="info-label">دلیل</div>
            <div class="info-value">${getReasonText(props.refund.reason)}</div>
          </div>
        </div>
        
        <div style="margin-top: 20px; padding: 15px; background-color: #f9fafb; border-radius: 6px;">
          <div style="font-size: 12px; color: #6b7280; margin-bottom: 5px;">توضیحات</div>
          <div style="font-size: 14px; color: #111827;">${props.refund.description}</div>
        </div>
      </div>
    </body>
    </html>
  `
  
  printWindow.document.write(printContent)
  printWindow.document.close()
  printWindow.print()
}

// تابع دانلود بازپرداخت
const exportRefund = () => {
  const refundData = {
    id: props.refund.id,
    giftCardCode: props.refund.giftCardCode,
    amount: props.refund.amount,
    status: getStatusText(props.refund.status),
    requester: props.refund.requester,
    requesterEmail: props.refund.requesterEmail,
    reason: getReasonText(props.refund.reason),
    description: props.refund.description,
    transactionId: props.refund.transactionId,
    createdAt: formatDateTime(props.refund.createdAt),
    approvedBy: props.refund.approvedBy,
    approvedAt: props.refund.approvedAt ? formatDateTime(props.refund.approvedAt) : null,
    rejectionReason: props.refund.rejectionReason
  }

  const jsonData = JSON.stringify(refundData, null, 2)
  const blob = new Blob([jsonData], { type: 'application/json' })
  const url = URL.createObjectURL(blob)
  const a = document.createElement('a')
  a.href = url
  a.download = `refund-${props.refund.id}.json`
  document.body.appendChild(a)
  a.click()
  document.body.removeChild(a)
  URL.revokeObjectURL(url)

  alert('فایل بازپرداخت با موفقیت دانلود شد')
}

// تابع تایید بازپرداخت
const approveRefund = () => {
  if (confirm(`آیا از تایید بازپرداخت #${props.refund.id} به مبلغ ${formatCurrency(props.refund.amount)} اطمینان دارید؟`)) {
    // اینجا کد تایید بازپرداخت را اضافه کنید

    alert('بازپرداخت با موفقیت تایید شد')
  }
}

// تابع رد بازپرداخت
const rejectRefund = () => {
  const rejectionReason = prompt('دلیل رد بازپرداخت را وارد کنید:')
  if (rejectionReason) {
    // اینجا کد رد بازپرداخت را اضافه کنید

    alert('بازپرداخت رد شد')
  }
}
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 

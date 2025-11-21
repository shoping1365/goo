<template>
  <div class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
    <div class="relative top-10 mx-auto p-5 border w-11/12 max-w-4xl shadow-lg rounded-md bg-white">
      <!-- هدر مودال -->
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-lg font-semibold text-gray-900">جزئیات تراکنش #{{ transaction.id }}</h3>
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
        <!-- اطلاعات اصلی تراکنش -->
        <div class="space-y-6">
          <!-- کارت اطلاعات اصلی -->
          <div class="bg-white border border-gray-200 rounded-lg p-6">
            <h4 class="text-lg font-semibold text-gray-900 mb-4">اطلاعات اصلی</h4>
            <div class="space-y-4">
              <div class="flex justify-between items-center">
                <span class="text-sm font-medium text-gray-600">شناسه تراکنش:</span>
                <span class="text-sm text-gray-900 font-mono">#{{ transaction.id }}</span>
              </div>
              <div class="flex justify-between items-center">
                <span class="text-sm font-medium text-gray-600">نوع تراکنش:</span>
                <span
                  :class="{
                    'bg-blue-100 text-blue-800': transaction.type === 'purchase',
                    'bg-green-100 text-green-800': transaction.type === 'usage',
                    'bg-red-100 text-red-800': transaction.type === 'refund',
                    'bg-yellow-100 text-yellow-800': transaction.type === 'expiry'
                  }"
                  class="px-2 py-1 text-xs font-medium rounded-full"
                >
                  {{ getTransactionTypeText(transaction.type) }}
                </span>
              </div>
              <div class="flex justify-between items-center">
                <span class="text-sm font-medium text-gray-600">مبلغ:</span>
                <span class="text-lg font-bold text-gray-900">{{ formatCurrency(transaction.amount) }}</span>
              </div>
              <div class="flex justify-between items-center">
                <span class="text-sm font-medium text-gray-600">وضعیت:</span>
                <span
                  :class="{
                    'bg-green-100 text-green-800': transaction.status === 'successful',
                    'bg-red-100 text-red-800': transaction.status === 'failed',
                    'bg-yellow-100 text-yellow-800': transaction.status === 'pending',
                    'bg-blue-100 text-blue-800': transaction.status === 'refunded'
                  }"
                  class="px-2 py-1 text-xs font-medium rounded-full"
                >
                  {{ getStatusText(transaction.status) }}
                </span>
              </div>
              <div class="flex justify-between items-center">
                <span class="text-sm font-medium text-gray-600">تاریخ و زمان:</span>
                <span class="text-sm text-gray-900">{{ formatDateTime(transaction.timestamp) }}</span>
              </div>
            </div>
          </div>

          <!-- اطلاعات کاربر -->
          <div class="bg-white border border-gray-200 rounded-lg p-6">
            <h4 class="text-lg font-semibold text-gray-900 mb-4">اطلاعات کاربر</h4>
            <div class="space-y-4">
              <div class="flex items-center">
                <div class="flex-shrink-0 h-12 w-12">
                  <div class="h-12 w-12 rounded-full bg-gray-300 flex items-center justify-center">
                    <span class="text-lg font-medium text-gray-700">
                      {{ transaction.user.charAt(0) }}
                    </span>
                  </div>
                </div>
                <div class="mr-4">
                  <div class="text-lg font-medium text-gray-900">{{ transaction.user }}</div>
                  <div class="text-sm text-gray-500">{{ transaction.userEmail || 'ایمیل ثبت نشده' }}</div>
                </div>
              </div>
            </div>
          </div>

          <!-- اطلاعات گیفت کارت -->
          <div v-if="transaction.giftCardCode" class="bg-white border border-gray-200 rounded-lg p-6">
            <h4 class="text-lg font-semibold text-gray-900 mb-4">اطلاعات گیفت کارت</h4>
            <div class="space-y-4">
              <div class="flex justify-between items-center">
                <span class="text-sm font-medium text-gray-600">کد کارت:</span>
                <span class="text-sm font-mono text-gray-900">{{ transaction.giftCardCode }}</span>
              </div>
              <div class="flex justify-between items-center">
                <span class="text-sm font-medium text-gray-600">توضیحات:</span>
                <span class="text-sm text-gray-900">{{ transaction.description || 'توضیحی ثبت نشده' }}</span>
              </div>
            </div>
          </div>
        </div>

        <!-- اطلاعات فنی و پرداخت -->
        <div class="space-y-6">
          <!-- اطلاعات پرداخت -->
          <div class="bg-white border border-gray-200 rounded-lg p-6">
            <h4 class="text-lg font-semibold text-gray-900 mb-4">اطلاعات پرداخت</h4>
            <div class="space-y-4">
              <div class="flex justify-between items-center">
                <span class="text-sm font-medium text-gray-600">روش پرداخت:</span>
                <span class="text-sm text-gray-900">{{ getPaymentMethodText(transaction.paymentMethod) }}</span>
              </div>
              <div v-if="transaction.orderId" class="flex justify-between items-center">
                <span class="text-sm font-medium text-gray-600">شماره سفارش:</span>
                <span class="text-sm font-mono text-gray-900">{{ transaction.orderId }}</span>
              </div>
              <div v-if="transaction.gatewayTransactionId" class="flex justify-between items-center">
                <span class="text-sm font-medium text-gray-600">شناسه درگاه:</span>
                <span class="text-sm font-mono text-gray-900">{{ transaction.gatewayTransactionId }}</span>
              </div>
              <div class="flex justify-between items-center">
                <span class="text-sm font-medium text-gray-600">پاسخ درگاه:</span>
                <span
                  :class="{
                    'bg-green-100 text-green-800': transaction.gatewayResponse === 'success',
                    'bg-red-100 text-red-800': transaction.gatewayResponse !== 'success',
                    'bg-yellow-100 text-yellow-800': transaction.gatewayResponse === 'pending'
                  }"
                  class="px-2 py-1 text-xs font-medium rounded-full"
                >
                  {{ getGatewayResponseText(transaction.gatewayResponse) }}
                </span>
              </div>
            </div>
          </div>

          <!-- آمار و تحلیل -->
          <div class="bg-white border border-gray-200 rounded-lg p-6">
            <h4 class="text-lg font-semibold text-gray-900 mb-4">آمار و تحلیل</h4>
            <div class="grid grid-cols-2 gap-6">
              <div class="text-center p-6 bg-gray-50 rounded-lg">
                <div class="text-2xl font-bold text-blue-600">{{ transaction.amount }}</div>
                <div class="text-xs text-gray-500">مبلغ (تومان)</div>
              </div>
              <div class="text-center p-6 bg-gray-50 rounded-lg">
                <div class="text-2xl font-bold text-green-600">{{ getTransactionAge(transaction.timestamp) }}</div>
                <div class="text-xs text-gray-500">روز از ایجاد</div>
              </div>
            </div>
          </div>

          <!-- تاریخچه تغییرات -->
          <div class="bg-white border border-gray-200 rounded-lg p-6">
            <h4 class="text-lg font-semibold text-gray-900 mb-4">تاریخچه تغییرات</h4>
            <div class="space-y-3">
              <div class="flex items-center">
                <div class="w-2 h-2 bg-green-400 rounded-full"></div>
                <div class="mr-3">
                  <div class="text-sm font-medium text-gray-900">تراکنش ایجاد شد</div>
                  <div class="text-xs text-gray-500">{{ formatDateTime(transaction.timestamp) }}</div>
                </div>
              </div>
              <div v-if="transaction.status === 'successful'" class="flex items-center">
                <div class="w-2 h-2 bg-green-400 rounded-full"></div>
                <div class="mr-3">
                  <div class="text-sm font-medium text-gray-900">تراکنش با موفقیت انجام شد</div>
                  <div class="text-xs text-gray-500">{{ formatDateTime(transaction.timestamp) }}</div>
                </div>
              </div>
              <div v-if="transaction.status === 'failed'" class="flex items-center">
                <div class="w-2 h-2 bg-red-400 rounded-full"></div>
                <div class="mr-3">
                  <div class="text-sm font-medium text-gray-900">تراکنش ناموفق بود</div>
                  <div class="text-xs text-gray-500">{{ formatDateTime(transaction.timestamp) }}</div>
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
          @click="printTransaction"
        >
          <svg class="w-4 h-4 mr-2 inline" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 17h2a2 2 0 002-2v-4a2 2 0 00-2-2H5a2 2 0 00-2 2v4a2 2 0 002 2h2m2 4h6a2 2 0 002-2v-4a2 2 0 00-2-2H9a2 2 0 00-2 2v4a2 2 0 002 2zm8-12V5a2 2 0 00-2-2H9a2 2 0 00-2 2v4h10z"></path>
          </svg>
          چاپ
        </button>
        <button
          class="px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
          @click="exportTransaction"
        >
          <svg class="w-4 h-4 mr-2 inline" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
          </svg>
          دانلود
        </button>
        <button
          v-if="transaction.status === 'successful'"
          class="px-4 py-2 bg-red-600 text-white text-sm font-medium rounded-lg hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2"
          @click="refundTransaction"
        >
          <svg class="w-4 h-4 mr-2 inline" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h10a8 8 0 018 8v2M3 10l6 6m-6-6l6-6"></path>
          </svg>
          بازپرداخت
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
  transaction: {
    type: Object,
    required: true
  }
})

const _emit = defineEmits(['close'])

// تابع دریافت متن نوع تراکنش
const getTransactionTypeText = (type) => {
  const typeMap = {
    purchase: 'خرید کارت',
    usage: 'استفاده کارت',
    refund: 'بازپرداخت',
    expiry: 'انقضا'
  }
  return typeMap[type] || type
}

// تابع دریافت متن وضعیت
const getStatusText = (status) => {
  const statusMap = {
    successful: 'موفق',
    failed: 'ناموفق',
    pending: 'در انتظار',
    refunded: 'بازپرداخت شده'
  }
  return statusMap[status] || status
}

// تابع دریافت متن روش پرداخت
const getPaymentMethodText = (method) => {
  const methodMap = {
    online: 'پرداخت آنلاین',
    gift_card: 'گیفت کارت',
    refund: 'بازپرداخت',
    system: 'سیستم'
  }
  return methodMap[method] || method
}

// تابع دریافت متن پاسخ درگاه
const getGatewayResponseText = (response) => {
  const responseMap = {
    success: 'موفق',
    pending: 'در انتظار',
    insufficient_funds: 'موجودی ناکافی',
    card_declined: 'کارت رد شد',
    network_error: 'خطای شبکه',
    timeout: 'زمان انتظار تمام شد',
    invalid_card: 'کارت نامعتبر'
  }
  return responseMap[response] || response
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

// تابع محاسبه سن تراکنش
const getTransactionAge = (timestamp) => {
  const now = new Date()
  const transactionDate = new Date(timestamp)
  const diffTime = Math.abs(now - transactionDate)
  const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24))
  return diffDays
}

// تابع چاپ تراکنش
const printTransaction = () => {
  const printWindow = window.open('', '_blank')
  const printContent = `
    <!DOCTYPE html>
    <html dir="rtl" lang="fa">
    <head>
      <meta charset="UTF-8">
      <title>تراکنش #${props.transaction.id}</title>
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
        .transaction-id {
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
          color: #059669;
          text-align: center;
          padding: 20px;
          background-color: #ecfdf5;
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
          <div class="transaction-id">تراکنش #${props.transaction.id}</div>
          <div style="margin-top: 10px; color: #6b7280;">
            ${formatDateTime(props.transaction.timestamp)}
          </div>
        </div>
        
        <div class="amount">
          ${formatCurrency(props.transaction.amount)}
        </div>
        
        <div class="info-grid">
          <div class="info-item">
            <div class="info-label">نوع تراکنش</div>
            <div class="info-value">${getTransactionTypeText(props.transaction.type)}</div>
          </div>
          <div class="info-item">
            <div class="info-label">وضعیت</div>
            <div class="info-value">${getStatusText(props.transaction.status)}</div>
          </div>
          <div class="info-item">
            <div class="info-label">کاربر</div>
            <div class="info-value">${props.transaction.user}</div>
          </div>
          <div class="info-item">
            <div class="info-label">روش پرداخت</div>
            <div class="info-value">${getPaymentMethodText(props.transaction.paymentMethod)}</div>
          </div>
          ${props.transaction.giftCardCode ? `
          <div class="info-item">
            <div class="info-label">کد گیفت کارت</div>
            <div class="info-value">${props.transaction.giftCardCode}</div>
          </div>
          ` : ''}
          ${props.transaction.orderId ? `
          <div class="info-item">
            <div class="info-label">شماره سفارش</div>
            <div class="info-value">${props.transaction.orderId}</div>
          </div>
          ` : ''}
        </div>
        
        ${props.transaction.description ? `
        <div style="margin-top: 20px; padding: 15px; background-color: #f9fafb; border-radius: 6px;">
          <div style="font-size: 12px; color: #6b7280; margin-bottom: 5px;">توضیحات</div>
          <div style="font-size: 14px; color: #111827;">${props.transaction.description}</div>
        </div>
        ` : ''}
      </div>
    </body>
    </html>
  `
  
  printWindow.document.write(printContent)
  printWindow.document.close()
  printWindow.print()
}

// تابع دانلود تراکنش
const exportTransaction = () => {
  const transactionData = {
    id: props.transaction.id,
    type: getTransactionTypeText(props.transaction.type),
    amount: props.transaction.amount,
    status: getStatusText(props.transaction.status),
    user: props.transaction.user,
    userEmail: props.transaction.userEmail,
    paymentMethod: getPaymentMethodText(props.transaction.paymentMethod),
    giftCardCode: props.transaction.giftCardCode,
    orderId: props.transaction.orderId,
    description: props.transaction.description,
    timestamp: formatDateTime(props.transaction.timestamp),
    gatewayResponse: getGatewayResponseText(props.transaction.gatewayResponse),
    gatewayTransactionId: props.transaction.gatewayTransactionId
  }

  const jsonData = JSON.stringify(transactionData, null, 2)
  const blob = new Blob([jsonData], { type: 'application/json' })
  const url = URL.createObjectURL(blob)
  const a = document.createElement('a')
  a.href = url
  a.download = `transaction-${props.transaction.id}.json`
  document.body.appendChild(a)
  a.click()
  document.body.removeChild(a)
  URL.revokeObjectURL(url)

  alert('فایل تراکنش با موفقیت دانلود شد')
}

// تابع بازپرداخت تراکنش
const refundTransaction = () => {
  if (confirm(`آیا از بازپرداخت تراکنش #${props.transaction.id} به مبلغ ${formatCurrency(props.transaction.amount)} اطمینان دارید؟`)) {
    // اینجا کد بازپرداخت را اضافه کنید

    alert('درخواست بازپرداخت ثبت شد و در حال پردازش است')
  }
}
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 

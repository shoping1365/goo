<template>
  <div class="fixed inset-0 bg-gray-600 bg-opacity-50 flex items-center justify-center z-50">
    <div class="bg-white rounded-lg p-6 w-full max-w-4xl max-h-[90vh] overflow-y-auto">
      <div class="flex justify-between items-center mb-6">
        <h3 class="text-lg font-bold text-gray-900">جزئیات گیفت کارت</h3>
        <button class="text-gray-400 hover:text-gray-600" @click="$emit('close')">
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>

      <div v-if="giftCard" class="space-y-6">
        <!-- اطلاعات اصلی -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div class="bg-gray-50 rounded-lg p-6">
            <h4 class="font-medium text-gray-900 mb-3">اطلاعات اصلی</h4>
            <div class="space-y-2 text-sm">
              <div class="flex justify-between">
                <span class="text-gray-600">کد کارت:</span>
                <span class="font-mono font-medium">{{ giftCard.code }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">مبلغ:</span>
                <span class="font-medium">{{ formatAmount(giftCard.amount) }} تومان</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">باقی‌مانده:</span>
                <span class="font-medium">{{ formatAmount(giftCard.remainingAmount) }} تومان</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">وضعیت:</span>
                <span :class="getStatusClasses(giftCard.status)" class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium">
                  {{ getStatusText(giftCard.status) }}
                </span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">نوع:</span>
                <span class="font-medium">{{ getTypeText(giftCard.type) }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">دسته‌بندی:</span>
                <span class="font-medium">{{ getCategoryText(giftCard.category) }}</span>
              </div>
            </div>
          </div>

          <div class="bg-gray-50 rounded-lg p-6">
            <h4 class="font-medium text-gray-900 mb-3">اطلاعات گیرنده</h4>
            <div class="space-y-2 text-sm">
              <div class="flex justify-between">
                <span class="text-gray-600">نام:</span>
                <span class="font-medium">{{ giftCard.recipientName || giftCard.recipient }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">ایمیل:</span>
                <span class="font-medium">{{ giftCard.recipientEmail }}</span>
              </div>
              <div v-if="giftCard.recipientPhone" class="flex justify-between">
                <span class="text-gray-600">تلفن:</span>
                <span class="font-medium">{{ giftCard.recipientPhone }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">روش ارسال:</span>
                <span class="font-medium">{{ getDeliveryMethodText(giftCard.deliveryMethod) }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">تاریخ ارسال:</span>
                <span class="font-medium">{{ formatDate(giftCard.deliveryDate) }}</span>
              </div>
            </div>
          </div>
        </div>

        <!-- پیام شخصی -->
        <div v-if="giftCard.personalMessage" class="bg-blue-50 rounded-lg p-6">
          <h4 class="font-medium text-blue-900 mb-2">پیام شخصی</h4>
          <p class="text-blue-800">{{ giftCard.personalMessage }}</p>
        </div>

        <!-- تنظیمات استفاده -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div class="bg-gray-50 rounded-lg p-6">
            <h4 class="font-medium text-gray-900 mb-3">تنظیمات استفاده</h4>
            <div class="space-y-2 text-sm">
              <div class="flex justify-between">
                <span class="text-gray-600">محدودیت استفاده:</span>
                <span class="font-medium">{{ getUsageLimitText(giftCard.usageLimit) }}</span>
              </div>
              <div v-if="giftCard.usageLimit === 'multiple'" class="flex justify-between">
                <span class="text-gray-600">تعداد استفاده مجاز:</span>
                <span class="font-medium">{{ giftCard.maxUsageCount }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">حداقل مبلغ سفارش:</span>
                <span class="font-medium">{{ giftCard.minOrderAmount ? formatAmount(giftCard.minOrderAmount) + ' تومان' : 'بدون محدودیت' }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">حداکثر مبلغ سفارش:</span>
                <span class="font-medium">{{ giftCard.maxOrderAmount ? formatAmount(giftCard.maxOrderAmount) + ' تومان' : 'بدون محدودیت' }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">استفاده جزئی:</span>
                <span class="font-medium">{{ giftCard.allowPartialUsage ? 'مجاز' : 'غیرمجاز' }}</span>
              </div>
            </div>
          </div>

          <div class="bg-gray-50 rounded-lg p-6">
            <h4 class="font-medium text-gray-900 mb-3">تنظیمات امنیتی</h4>
            <div class="space-y-2 text-sm">
              <div class="flex justify-between">
                <span class="text-gray-600">تأیید هویت:</span>
                <span class="font-medium">{{ giftCard.requireVerification ? 'ضروری' : 'اختیاری' }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">تمدید خودکار:</span>
                <span class="font-medium">{{ giftCard.autoRenew ? 'فعال' : 'غیرفعال' }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">تاریخ انقضا:</span>
                <span class="font-medium">{{ formatDate(giftCard.expiryDate) }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">تاریخ ایجاد:</span>
                <span class="font-medium">{{ formatDate(giftCard.createdAt) }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">ایجاد شده توسط:</span>
                <span class="font-medium">{{ giftCard.createdBy }}</span>
              </div>
            </div>
          </div>
        </div>

        <!-- دسته‌بندی‌های مجاز -->
        <div v-if="giftCard.allowedCategories && giftCard.allowedCategories.length > 0" class="bg-gray-50 rounded-lg p-6">
          <h4 class="font-medium text-gray-900 mb-3">دسته‌بندی‌های مجاز</h4>
          <div class="flex flex-wrap gap-2">
            <span 
              v-for="category in giftCard.allowedCategories" 
              :key="category"
              class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800"
            >
              {{ getCategoryText(category) }}
            </span>
          </div>
        </div>

        <!-- تاریخچه استفاده -->
        <div v-if="giftCard.usageHistory && giftCard.usageHistory.length > 0" class="bg-gray-50 rounded-lg p-6">
          <h4 class="font-medium text-gray-900 mb-3">تاریخچه استفاده</h4>
          <div class="space-y-3">
            <div 
              v-for="usage in giftCard.usageHistory" 
              :key="usage.id"
              class="flex justify-between items-center p-3 bg-white rounded-lg border border-gray-200"
            >
              <div>
                <div class="font-medium text-gray-900">{{ formatAmount(usage.amount) }} تومان</div>
                <div class="text-sm text-gray-600">{{ usage.description || 'استفاده از گیفت کارت' }}</div>
                <div class="text-xs text-gray-500">{{ formatDate(usage.date) }}</div>
              </div>
              <div class="text-right">
                <div class="text-sm font-medium text-gray-900">سفارش: {{ usage.orderId }}</div>
              </div>
            </div>
          </div>
        </div>

        <!-- دکمه‌های عملیات -->
        <div class="flex justify-end space-x-3 space-x-reverse pt-6 border-t border-gray-200">
          <button 
            class="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-lg hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
            @click="$emit('close')"
          >
            بستن
          </button>
          <button 
            class="px-4 py-2 text-sm font-medium text-blue-600 bg-blue-50 border border-blue-200 rounded-lg hover:bg-blue-100 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
            @click="copyCode"
          >
            کپی کد
          </button>
          <button 
            v-if="giftCard.deliveryMethod !== 'manual'"
            class="px-4 py-2 text-sm font-medium text-green-600 bg-green-50 border border-green-200 rounded-lg hover:bg-green-100 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2"
            @click="resendGiftCard"
          >
            ارسال مجدد
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

// Props
const props = defineProps<{
  giftCard: any
}>()

// Emits
const emit = defineEmits<{
  close: []
}>()

// Computed properties
const giftCard = computed(() => props.giftCard)

// Methods
const formatAmount = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount)
}

const formatDate = (date: string | Date) => {
  return new Date(date).toLocaleDateString('fa-IR', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}

const getStatusText = (status: string) => {
  const statusMap = {
    active: 'فعال',
    used: 'استفاده شده',
    expired: 'منقضی شده',
    inactive: 'غیرفعال'
  }
  return statusMap[status] || status
}

const getStatusClasses = (status: string) => {
  const classesMap = {
    active: 'bg-green-100 text-green-800',
    used: 'bg-blue-100 text-blue-800',
    expired: 'bg-red-100 text-red-800',
    inactive: 'bg-gray-100 text-gray-800'
  }
  return classesMap[status] || 'bg-gray-100 text-gray-800'
}

const getTypeText = (type: string) => {
  const typeMap = {
    digital: 'دیجیتال',
    physical: 'فیزیکی',
    hybrid: 'ترکیبی'
  }
  return typeMap[type] || type
}

const getCategoryText = (category: string) => {
  const categoryMap = {
    birthday: 'تولد',
    wedding: 'عروسی',
    newyear: 'سال نو',
    corporate: 'شرکتی',
    discount: 'تخفیف ویژه',
    general: 'عمومی',
    electronics: 'الکترونیک',
    clothing: 'پوشاک',
    home: 'خانه و آشپزخانه',
    beauty: 'زیبایی'
  }
  return categoryMap[category] || category
}

const getDeliveryMethodText = (method: string) => {
  const methodMap = {
    email: 'ایمیل',
    sms: 'پیامک',
    both: 'ایمیل و پیامک',
    manual: 'ارسال دستی'
  }
  return methodMap[method] || method
}

const getUsageLimitText = (limit: string) => {
  const limitMap = {
    unlimited: 'نامحدود',
    once: 'یک بار',
    multiple: 'چند بار'
  }
  return limitMap[limit] || limit
}

const copyCode = async () => {
  try {
    await navigator.clipboard.writeText(giftCard.value.code)
    alert('کد کپی شد')
  } catch (err) {
    console.error('خطا در کپی کردن کد:', err)
  }
}

const resendGiftCard = async () => {
  if (confirm(`آیا می‌خواهید گیفت کارت ${giftCard.value.code} مجدداً ارسال شود؟`)) {
    try {
      // در نسخه واقعی: ارسال به API
      console.log('Resending gift card:', giftCard.value.id)
      alert('گیفت کارت با موفقیت ارسال شد')
    } catch (error) {
      console.error('خطا در ارسال مجدد:', error)
      alert('خطا در ارسال مجدد گیفت کارت')
    }
  }
}
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 

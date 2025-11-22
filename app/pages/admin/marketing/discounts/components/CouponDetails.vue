<template>
  <div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-6">
    <div class="bg-white rounded-lg shadow-xl w-full max-w-4xl max-h-[90vh] overflow-y-auto">
      <!-- هدر مودال -->
      <div class="p-6 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <h2 class="text-xl font-semibold text-gray-900">جزئیات کوپن</h2>
          <button class="text-gray-400 hover:text-gray-600" @click="$emit('close')">
            <span class="i-heroicons-x-mark text-xl"></span>
          </button>
        </div>
      </div>

      <!-- محتوای جزئیات -->
      <div class="p-6">
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
          <!-- اطلاعات اصلی -->
          <div class="space-y-6">
            <div>
              <h3 class="text-lg font-medium text-gray-900 mb-4">اطلاعات اصلی</h3>
              <div class="space-y-4">
                <div>
                  <label class="block text-sm font-medium text-gray-700">نام کوپن</label>
                  <p class="mt-1 text-sm text-gray-900">{{ coupon.name }}</p>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700">کد کوپن</label>
                  <div class="mt-1 flex items-center gap-2">
                    <span class="font-mono text-sm bg-gray-100 px-3 py-2 rounded-lg">{{ coupon.code }}</span>
                    <button class="text-blue-600 hover:text-blue-800" @click="copyCode">
                      <span class="i-heroicons-clipboard-document"></span>
                    </button>
                  </div>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700">توضیحات</label>
                  <p class="mt-1 text-sm text-gray-900">{{ coupon.description || 'توضیحی ثبت نشده' }}</p>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700">وضعیت</label>
                  <span :class="`mt-1 inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${getStatusBadgeClass(String(coupon.status || ''))}`">
                    {{ getStatusName(String(coupon.status || '')) }}
                  </span>
                </div>
              </div>
            </div>

            <!-- تنظیمات تخفیف -->
            <div>
              <h3 class="text-lg font-medium text-gray-900 mb-4">تنظیمات تخفیف</h3>
              <div class="space-y-4">
                <div>
                  <label class="block text-sm font-medium text-gray-700">نوع تخفیف</label>
                  <div class="mt-1 flex items-center gap-2">
                    <span :class="`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${getTypeBadgeClass(String(coupon.type || ''))}`">
                      {{ getTypeName(String(coupon.type || '')) }}
                    </span>
                    <span class="text-sm text-gray-900">{{ formatDiscount(coupon) }}</span>
                  </div>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700">حداقل مبلغ سفارش</label>
                  <p class="mt-1 text-sm text-gray-900">{{ formatCurrency(Number(coupon.minOrderAmount || 0)) }}</p>
                </div>
                <div v-if="coupon.maxDiscount">
                  <label class="block text-sm font-medium text-gray-700">حداکثر تخفیف</label>
                  <p class="mt-1 text-sm text-gray-900">{{ formatCurrency(Number(coupon.maxDiscount || 0)) }}</p>
                </div>
              </div>
            </div>
          </div>

          <!-- محدودیت‌ها و آمار -->
          <div class="space-y-6">
            <!-- محدودیت‌ها -->
            <div>
              <h3 class="text-lg font-medium text-gray-900 mb-4">محدودیت‌ها</h3>
              <div class="space-y-4">
                <div>
                  <label class="block text-sm font-medium text-gray-700">حداکثر تعداد استفاده</label>
                  <p class="mt-1 text-sm text-gray-900">{{ coupon.maxUses || 'نامحدود' }}</p>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700">تعداد استفاده برای هر کاربر</label>
                  <p class="mt-1 text-sm text-gray-900">{{ coupon.maxUsesPerUser || 'نامحدود' }}</p>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700">تاریخ شروع</label>
                  <p class="mt-1 text-sm text-gray-900">{{ formatDate(String(coupon.startsAt || '')) }}</p>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700">تاریخ انقضا</label>
                  <p class="mt-1 text-sm text-gray-900">{{ formatDate(String(coupon.expiresAt || '')) }}</p>
                </div>
              </div>
            </div>

            <!-- آمار استفاده -->
            <div>
              <h3 class="text-lg font-medium text-gray-900 mb-4">آمار استفاده</h3>
              <div class="space-y-4">
                <div>
                  <label class="block text-sm font-medium text-gray-700">تعداد استفاده شده</label>
                  <p class="mt-1 text-sm text-gray-900">{{ coupon.usedCount || 0 }}</p>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700">تعداد باقی‌مانده</label>
                  <p class="mt-1 text-sm text-gray-900">{{ getRemainingUses() }}</p>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700">درصد استفاده</label>
                  <div class="mt-1">
                    <div class="w-full bg-gray-200 rounded-full h-2">
                      <div :class="`h-2 rounded-full ${getUsageProgressClass()}`" :style="{ width: getUsagePercentage() + '%' }"></div>
                    </div>
                    <p class="mt-1 text-xs text-gray-600">{{ getUsagePercentage() }}% استفاده شده</p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- تاریخچه استفاده -->
        <div class="mt-8">
          <h3 class="text-lg font-medium text-gray-900 mb-4">تاریخچه استفاده</h3>
          <div class="overflow-x-auto">
            <table class="w-full min-w-max divide-y divide-gray-200">
              <thead class="bg-gray-50">
                <tr>
                  <th class="px-3 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کاربر</th>
                  <th class="px-3 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">سفارش</th>
                  <th class="px-3 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مبلغ تخفیف</th>
                  <th class="px-3 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ استفاده</th>
                </tr>
              </thead>
              <tbody class="bg-white divide-y divide-gray-200">
                <tr v-for="usage in (Array.isArray(coupon.usageHistory) ? coupon.usageHistory : [])" :key="(usage as { id?: number | string }).id" class="hover:bg-gray-50">
                  <td class="px-3 py-4 whitespace-nowrap text-sm text-gray-900">{{ String((usage as { userName?: string }).userName || '') }}</td>
                  <td class="px-3 py-4 whitespace-nowrap text-sm text-gray-900">#{{ String((usage as { orderId?: string }).orderId || '') }}</td>
                  <td class="px-3 py-4 whitespace-nowrap text-sm text-gray-900">{{ formatCurrency(Number((usage as { discountAmount?: number }).discountAmount || 0)) }}</td>
                  <td class="px-3 py-4 whitespace-nowrap text-sm text-gray-900">{{ formatDate(String((usage as { usedAt?: string }).usedAt || '')) }}</td>
                </tr>
              </tbody>
            </table>
          </div>
          <div v-if="!Array.isArray(coupon.usageHistory) || coupon.usageHistory.length === 0" class="text-center py-8">
            <p class="text-gray-500">هنوز از این کوپن استفاده نشده است.</p>
          </div>
        </div>

        <!-- دکمه‌های عملیات -->
        <div class="flex items-center justify-end gap-3 mt-8 pt-6 border-t border-gray-200">
          <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors" @click="editCoupon">
            ویرایش کوپن
          </button>
          <button class="px-4 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200 transition-colors" @click="$emit('close')">
            بستن
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// Props
interface Coupon {
  id?: number | string
  code?: string
  discount?: number
  status?: string
  minOrderAmount?: number
  maxDiscount?: number
  startsAt?: string
  expiresAt?: string
  maxUses?: number
  usedCount?: number
  usageHistory?: Array<{ id?: number | string; userName?: string; orderId?: string; discountAmount?: number; usedAt?: string; [key: string]: unknown }>
  [key: string]: unknown
}

interface Props {
  coupon: Coupon
}

const props = defineProps<Props>()
const emit = defineEmits<{
  close: []
  edit: [coupon: Coupon]
}>()

// توابع کمکی
const getStatusBadgeClass = (status: string) => {
  switch (status) {
    case 'active':
      return 'bg-green-100 text-green-800'
    case 'inactive':
      return 'bg-gray-100 text-gray-800'
    case 'expired':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getStatusName = (status: string) => {
  switch (status) {
    case 'active':
      return 'فعال'
    case 'inactive':
      return 'غیرفعال'
    case 'expired':
      return 'منقضی شده'
    default:
      return 'نامشخص'
  }
}

const getTypeBadgeClass = (type: string) => {
  switch (type) {
    case 'percentage':
      return 'bg-blue-100 text-blue-800'
    case 'fixed':
      return 'bg-purple-100 text-purple-800'
    case 'free_shipping':
      return 'bg-green-100 text-green-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getTypeName = (type: string) => {
  switch (type) {
    case 'percentage':
      return 'درصدی'
    case 'fixed':
      return 'مبلغ ثابت'
    case 'free_shipping':
      return 'ارسال رایگان'
    default:
      return 'نامشخص'
  }
}

const formatDiscount = (coupon: Coupon) => {
  if (coupon.type === 'percentage') {
    return `${coupon.discountValue}%`
  } else if (coupon.type === 'fixed') {
    return formatCurrency(Number(coupon.discountValue || 0))
  } else if (coupon.type === 'free_shipping') {
    return 'ارسال رایگان'
  }
  return ''
}

const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR', {
    style: 'currency',
    currency: 'IRR'
  }).format(amount)
}

const formatDate = (date: string) => {
  return new Date(date).toLocaleDateString('fa-IR', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}

const getRemainingUses = () => {
  const maxUses = Number(props.coupon.maxUses || 0)
  if (!maxUses) return 'نامحدود'
  return Math.max(0, maxUses - Number(props.coupon.usedCount || 0))
}

const getUsagePercentage = () => {
  const maxUses = Number(props.coupon.maxUses || 0)
  if (!maxUses) return 0
  return Math.round((Number(props.coupon.usedCount || 0) / maxUses) * 100)
}

const getUsageProgressClass = () => {
  const percentage = getUsagePercentage()
  if (percentage >= 80) return 'bg-red-500'
  if (percentage >= 60) return 'bg-yellow-500'
  return 'bg-green-500'
}

const copyCode = () => {
  navigator.clipboard.writeText(props.coupon.code)
  // می‌توانید یک پیام موفقیت نمایش دهید
}

const editCoupon = () => {
  emit('edit', props.coupon)
}
</script> 

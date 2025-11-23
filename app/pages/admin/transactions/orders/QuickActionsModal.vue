<template>
  <div v-if="isOpen && hasAccess" class="fixed inset-0 z-50 overflow-y-auto">
    <!-- Backdrop -->
    <div class="fixed inset-0 bg-black bg-opacity-50 transition-opacity" @click="closeModal"></div>
    
    <!-- Modal -->
    <div class="flex min-h-full items-center justify-center px-4 py-4">
      <div class="relative w-full max-w-md bg-white rounded-lg shadow-xl">
        <!-- Header -->
        <div class="flex items-center justify-between px-4 py-4 border-b border-gray-200">
          <div class="flex items-center space-x-3">
            <div class="w-10 h-10 bg-purple-100 rounded-full flex items-center justify-center">
              <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z"></path>
              </svg>
            </div>
            <div>
              <h3 class="text-lg font-semibold text-gray-900">عملیات سریع</h3>
              <p class="text-sm text-gray-500">سفارش: {{ order?.orderNumber }}</p>
            </div>
          </div>
          <button class="text-gray-400 hover:text-gray-600 transition-colors" @click="closeModal">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>

        <!-- Content -->
        <div class="px-4 py-4">
          <div class="space-y-3">
            <!-- Status Change -->
            <div class="bg-gray-50 px-4 py-4 rounded-lg">
              <h4 class="text-sm font-medium text-gray-900 mb-3">تغییر وضعیت</h4>
              <div class="grid grid-cols-2 gap-2">
                <button
                  v-for="status in availableStatuses"
                  :key="status.value"
                  :class="[
                    'px-3 py-2 text-sm rounded-lg transition-colors',
                    order?.status === status.value
                      ? 'bg-blue-100 text-blue-800 border border-blue-300'
                      : 'bg-white text-gray-700 border border-gray-300 hover:bg-gray-50'
                  ]"
                  @click="changeStatus(status.value)"
                >
                  {{ status.label }}
                </button>
              </div>
            </div>

            <!-- Quick Actions -->
            <div class="space-y-2">
              <button
                class="w-full flex items-center justify-between p-3 text-right bg-blue-50 hover:bg-blue-100 rounded-lg transition-colors"
                @click="sendNotification"
              >
                <div class="flex items-center">
                  <svg class="w-5 h-5 text-blue-600 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-5 5v-5z"></path>
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 7h6m0 10v-3m-3 3h.01M9 17h.01M9 14h.01M12 14h.01M15 11h.01M12 11h.01M9 11h.01M7 21h10a2 2 0 002-2V5a2 2 0 00-2-2H7a2 2 0 00-2 2v14a2 2 0 002 2z"></path>
                  </svg>
                  <span class="text-sm font-medium text-blue-900">ارسال پیامک</span>
                </div>
                <svg class="w-4 h-4 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
                </svg>
              </button>

              <button
                class="w-full flex items-center justify-between p-3 text-right bg-green-50 hover:bg-green-100 rounded-lg transition-colors"
                @click="printInvoice"
              >
                <div class="flex items-center">
                  <svg class="w-5 h-5 text-green-600 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 17h2a2 2 0 002-2v-4a2 2 0 00-2-2H5a2 2 0 00-2 2v4a2 2 0 002 2h2m2 4h6a2 2 0 002-2v-4a2 2 0 00-2-2H9a2 2 0 00-2 2v4a2 2 0 002 2zm8-12V5a2 2 0 00-2-2H9a2 2 0 00-2 2v4h10z"></path>
                  </svg>
                  <span class="text-sm font-medium text-green-900">چاپ فاکتور</span>
                </div>
                <svg class="w-4 h-4 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
                </svg>
              </button>

              <button
                class="w-full flex items-center justify-between p-3 text-right bg-purple-50 hover:bg-purple-100 rounded-lg transition-colors"
                @click="duplicateOrder"
              >
                <div class="flex items-center">
                  <svg class="w-5 h-5 text-purple-600 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"></path>
                  </svg>
                  <span class="text-sm font-medium text-purple-900">تکرار سفارش</span>
                </div>
                <svg class="w-4 h-4 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
                </svg>
              </button>

              <button
                class="w-full flex items-center justify-between p-3 text-right bg-orange-50 hover:bg-orange-100 rounded-lg transition-colors"
                @click="exportOrder"
              >
                <div class="flex items-center">
                  <svg class="w-5 h-5 text-orange-600 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
                  </svg>
                  <span class="text-sm font-medium text-orange-900">خروجی Excel</span>
                </div>
                <svg class="w-4 h-4 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
                </svg>
              </button>
            </div>

            <!-- Danger Zone -->
            <div class="border-t border-gray-200 pt-4">
              <h4 class="text-sm font-medium text-red-900 mb-3">منطقه خطر</h4>
              <div class="space-y-2">
                <button
                  class="w-full flex items-center justify-between p-3 text-right bg-red-50 hover:bg-red-100 rounded-lg transition-colors"
                  @click="cancelOrder"
                >
                  <div class="flex items-center">
                    <svg class="w-5 h-5 text-red-600 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                    </svg>
                    <span class="text-sm font-medium text-red-900">لغو سفارش</span>
                  </div>
                  <svg class="w-4 h-4 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
                  </svg>
                </button>

                <button
                  v-if="canDeleteOrder"
                  class="w-full flex items-center justify-between p-3 text-right bg-red-50 hover:bg-red-100 rounded-lg transition-colors"
                  @click="deleteOrder"
                >
                  <div class="flex items-center">
                    <svg class="w-5 h-5 text-red-600 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                    </svg>
                    <span class="text-sm font-medium text-red-900">حذف سفارش</span>
                  </div>
                  <svg class="w-4 h-4 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>

        <!-- Footer -->
        <div class="flex items-center justify-end px-4 py-4 border-t border-gray-200">
          <button class="px-4 py-2 text-gray-700 bg-gray-100 hover:bg-gray-200 rounded-lg transition-colors" @click="closeModal">
            بستن
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>
declare const hasPermission: (perm: string) => boolean
</script>

<script setup lang="ts">
import { computed, watch } from 'vue';
import { useAuth } from '~/composables/useAuth';

// احراز هویت
const { user, isAuthenticated } = useAuth()

// بررسی دسترسی admin
const hasAccess = computed(() => {
  if (!isAuthenticated.value) {
    return false
  }

  const userRole = user.value?.role?.toLowerCase() || ''
  const adminRoles = ['admin', 'developer']
  return adminRoles.includes(userRole)
})

// بررسی احراز هویت و دسترسی admin - نمایش 404 در صورت عدم دسترسی
const checkAuth = async () => {
  if (!hasAccess.value) {
    await navigateTo('/404', { external: false })
  }
}

// بررسی احراز هویت هنگام تغییر وضعیت احراز هویت
watch([isAuthenticated, hasAccess], async () => {
  if (!hasAccess.value) {
    await checkAuth()
  }
})

// Auth disabled
const canDeleteOrder = computed(() => hasPermission('order.delete'))

// بررسی احراز هویت هنگام باز شدن modal
const props = defineProps({
  isOpen: {
    type: Boolean,
    default: false
  },
  order: {
    type: Object,
    default: null
  }
})

watch(() => props.isOpen, async (newValue) => {
  if (newValue) {
    await checkAuth()
  }
})

const emit = defineEmits(['close', 'status-change', 'notification', 'print', 'duplicate', 'export', 'cancel', 'delete'])

const availableStatuses = computed(() => [
  { value: 'pending', label: 'در انتظار پرداخت' },
  { value: 'paid', label: 'پرداخت شده' },
  { value: 'processing', label: 'در حال پردازش' },
  { value: 'shipped', label: 'ارسال شده' },
  { value: 'delivered', label: 'تحویل داده شده' },
  { value: 'cancelled', label: 'لغو شده' }
])

const closeModal = () => {
  emit('close')
}

const changeStatus = (status) => {
  emit('status-change', { orderId: props.order?.id, status })
  closeModal()
}

const sendNotification = () => {
  emit('notification', props.order)
  closeModal()
}

const printInvoice = () => {
  emit('print', props.order)
  closeModal()
}

const duplicateOrder = () => {
  emit('duplicate', props.order)
  closeModal()
}

const exportOrder = () => {
  emit('export', props.order)
  closeModal()
}

const cancelOrder = () => {
  if (confirm('آیا از لغو این سفارش اطمینان دارید؟')) {
    emit('cancel', props.order)
    closeModal()
  }
}

const deleteOrder = () => {
  if (confirm('آیا از حذف این سفارش اطمینان دارید؟ این عملیات غیرقابل بازگشت است.')) {
    emit('delete', props.order)
    closeModal()
  }
}
</script> 

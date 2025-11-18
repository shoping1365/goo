<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="flex items-center justify-between">
      <div>
        <h3 class="text-lg font-semibold text-gray-800">ارسال دستی SMS</h3>
        <p class="text-gray-600 text-sm">انتخاب و ارسال پیامک برای سفارشات خاص</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button 
          @click="refreshOrders"
          class="px-4 py-2 bg-blue-600 text-white rounded-lg text-sm hover:bg-blue-700 transition-colors"
        >
          بروزرسانی لیست
        </button>
      </div>
    </div>

    <!-- Quick Actions -->
    <div class="bg-blue-50 rounded-lg p-6">
      <h4 class="text-md font-semibold text-blue-800 mb-3">عملیات سریع</h4>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <button 
          @click="selectAllNotSent"
          class="flex items-center justify-center space-x-2 space-x-reverse p-3 bg-white rounded-lg border border-blue-200 hover:border-blue-300 transition-colors"
        >
          <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
          </svg>
          <span class="text-blue-800 font-medium">انتخاب همه ارسال نشده</span>
        </button>
        
        <button 
          @click="selectTodayOrders"
          class="flex items-center justify-center space-x-2 space-x-reverse p-3 bg-white rounded-lg border border-blue-200 hover:border-blue-300 transition-colors"
        >
          <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
          </svg>
          <span class="text-blue-800 font-medium">سفارشات امروز</span>
        </button>
        
        <button 
          @click="selectHighValueOrders"
          class="flex items-center justify-center space-x-2 space-x-reverse p-3 bg-white rounded-lg border border-blue-200 hover:border-blue-300 transition-colors"
        >
          <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
          </svg>
          <span class="text-blue-800 font-medium">سفارشات با ارزش بالا</span>
        </button>
      </div>
    </div>

    <!-- Selection Summary -->
    <div v-if="selectedOrders.length > 0" class="bg-green-50 rounded-lg p-6">
      <div class="flex items-center justify-between">
        <div class="flex items-center space-x-4 space-x-reverse">
          <div class="flex items-center space-x-2 space-x-reverse">
            <svg class="w-5 h-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
            <span class="text-green-800 font-medium">{{ selectedOrders.length }} سفارش انتخاب شده</span>
          </div>
          <button 
            @click="clearSelection"
            class="text-green-600 hover:text-green-800 text-sm"
          >
            پاک کردن انتخاب
          </button>
        </div>
        
        <div class="flex items-center space-x-3 space-x-reverse">
          <button 
            @click="sendBulkSMS"
            :disabled="sending"
            class="px-4 py-2 bg-green-600 hover:bg-green-700 disabled:bg-gray-400 text-white rounded-lg font-medium transition-colors flex items-center space-x-2 space-x-reverse"
          >
            <svg v-if="sending" class="animate-spin w-5 h-5" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8"></path>
            </svg>
            <span>{{ sending ? 'در حال ارسال...' : 'ارسال SMS گروهی' }}</span>
          </button>
          
          <button 
            @click="previewBulkSMS"
            class="px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg font-medium transition-colors"
          >
            پیش‌نمایش
          </button>
        </div>
      </div>
    </div>

    <!-- Orders Table -->
    <div class="bg-white rounded-lg border border-gray-200 overflow-hidden">
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                <input 
                  v-model="selectAll"
                  @change="toggleSelectAll"
                  type="checkbox" 
                  class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                >
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                شماره سفارش
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                مشتری
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                تاریخ سفارش
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                وضعیت SMS
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                عملیات
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr 
              v-for="order in orders" 
              :key="order.id"
              class="hover:bg-gray-50 transition-colors"
            >
              <td class="px-6 py-4 whitespace-nowrap">
                <input 
                  :checked="selectedOrders.includes(order.id)"
                  @change="toggleOrderSelection(order.id)"
                  type="checkbox" 
                  class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                >
              </td>
              
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm font-medium text-gray-900">#{{ order.orderNumber }}</div>
              </td>
              
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="">
                    <div class="h-10 w-10 rounded-full bg-gradient-to-r from-blue-400 to-purple-500 flex items-center justify-center">
                      <span class="text-white font-medium text-sm">{{ getInitials(order.customerName) }}</span>
                    </div>
                  </div>
                  <div class="mr-4">
                    <div class="text-sm font-medium text-gray-900">{{ order.customerName }}</div>
                    <div class="text-sm text-gray-500">{{ order.phoneNumber }}</div>
                  </div>
                </div>
              </td>
              
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ formatDate(order.orderDate) }}
              </td>
              
              <td class="px-6 py-4 whitespace-nowrap">
                <span 
                  :class="[
                    'inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium',
                    getSMSStatusClass(order.smsStatus)
                  ]"
                >
                  {{ getSMSStatusText(order.smsStatus) }}
                </span>
              </td>
              
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <div class="flex items-center space-x-2 space-x-reverse">
                  <button 
                    @click="sendSingleSMS(order.id)"
                    :disabled="order.smsStatus === 'sent' || order.smsStatus === 'delivered'"
                    class="text-blue-600 hover:text-blue-900 disabled:text-gray-400 transition-colors"
                    title="ارسال SMS"
                  >
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8"></path>
                    </svg>
                  </button>
                  
                  <button 
                    @click="previewSMS(order.id)"
                    class="text-green-600 hover:text-green-900 transition-colors"
                    title="پیش‌نمایش SMS"
                  >
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                    </svg>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Pagination -->
    <div class="flex items-center justify-between">
      <div class="text-sm text-gray-700">
        نمایش {{ startIndex + 1 }} تا {{ endIndex }} از {{ totalOrders }} سفارش
      </div>
      <div class="flex space-x-2 space-x-reverse">
        <button 
          @click="previousPage"
          :disabled="currentPage === 1"
          class="px-3 py-1 border border-gray-300 rounded text-sm disabled:opacity-50 disabled:cursor-not-allowed hover:bg-gray-50"
        >
          قبلی
        </button>
        <span class="px-3 py-1 text-sm text-gray-700">
          صفحه {{ currentPage }} از {{ totalPages }}
        </span>
        <button 
          @click="nextPage"
          :disabled="currentPage >= totalPages"
          class="px-3 py-1 border border-gray-300 rounded text-sm disabled:opacity-50 disabled:cursor-not-allowed hover:bg-gray-50"
        >
          بعدی
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'

interface Order {
  id: number
  orderNumber: string
  customerName: string
  phoneNumber: string
  orderDate: string
  smsStatus: string
}

const props = defineProps<{
  orders: Order[]
  selectedOrders: number[]
}>()

const emit = defineEmits<{
  'send-single': [orderId: number]
  'send-bulk': [orderIds: number[]]
  'select-orders': [orders: number[]]
}>()

// Reactive data
const sending = ref(false)
const currentPage = ref(1)
const itemsPerPage = ref(20)

// Computed properties
const selectAll = computed({
  get: () => props.selectedOrders.length === props.orders.length && props.orders.length > 0,
  set: (value: boolean) => {
    if (value) {
      emit('select-orders', props.orders.map(order => order.id))
    } else {
      emit('select-orders', [])
    }
  }
})

const totalOrders = computed(() => props.orders.length)
const totalPages = computed(() => Math.ceil(totalOrders.value / itemsPerPage.value))
const startIndex = computed(() => (currentPage.value - 1) * itemsPerPage.value)
const endIndex = computed(() => Math.min(startIndex.value + itemsPerPage.value, totalOrders.value))

// Methods
const sendSingleSMS = async (orderId: number) => {
  emit('send-single', orderId)
}

const sendBulkSMS = async () => {
  if (props.selectedOrders.length === 0) return
  
  sending.value = true
  try {
    emit('send-bulk', props.selectedOrders)
    // Wait for the operation to complete
    await new Promise(resolve => setTimeout(resolve, 2000))
  } finally {
    sending.value = false
  }
}

const toggleSelectAll = () => {
  selectAll.value = !selectAll.value
}

const selectAllNotSent = () => {
  const notSentOrders = props.orders
    .filter(order => order.smsStatus === 'not_sent')
    .map(order => order.id)
  emit('select-orders', notSentOrders)
}

const selectTodayOrders = () => {
  const today = new Date().toISOString().split('T')[0]
  const todayOrders = props.orders
    .filter(order => order.orderDate === today)
    .map(order => order.id)
  emit('select-orders', todayOrders)
}

const selectHighValueOrders = () => {
  // This would filter orders with high value (implementation depends on your data structure)
  const highValueOrders = props.orders
    .filter(order => order.orderNumber.includes('VIP')) // Example filter
    .map(order => order.id)
  emit('select-orders', highValueOrders)
}

const clearSelection = () => {
  emit('select-orders', [])
}

const toggleOrderSelection = (orderId: number) => {
  const newSelection = [...props.selectedOrders]
  const index = newSelection.indexOf(orderId)
  
  if (index > -1) {
    newSelection.splice(index, 1)
  } else {
    newSelection.push(orderId)
  }
  
  emit('select-orders', newSelection)
}

const refreshOrders = () => {
  // This would trigger a refresh of the orders list
  console.log('Refreshing orders...')
}

const previewSMS = (orderId: number) => {
  console.log(`Previewing SMS for order ${orderId}`)
}

const previewBulkSMS = () => {
  console.log(`Previewing bulk SMS for ${props.selectedOrders.length} orders`)
}

const previousPage = () => {
  if (currentPage.value > 1) {
    currentPage.value--
  }
}

const nextPage = () => {
  if (currentPage.value < totalPages.value) {
    currentPage.value++
  }
}

// Utility functions
const getInitials = (name: string) => {
  return name.split(' ').map(n => n[0]).join('').toUpperCase()
}

const formatDate = (date: string) => {
  return new Date(date).toLocaleDateString('fa-IR')
}

const getSMSStatusClass = (status: string) => {
  const classes = {
    sent: 'bg-blue-100 text-blue-800',
    delivered: 'bg-green-100 text-green-800',
    failed: 'bg-red-100 text-red-800',
    pending: 'bg-yellow-100 text-yellow-800',
    not_sent: 'bg-gray-100 text-gray-800'
  }
  return classes[status as keyof typeof classes] || 'bg-gray-100 text-gray-800'
}

const getSMSStatusText = (status: string) => {
  const texts = {
    sent: 'ارسال شده',
    delivered: 'تحویل شده',
    failed: 'ناموفق',
    pending: 'در انتظار',
    not_sent: 'ارسال نشده'
  }
  return texts[status as keyof typeof texts] || status
}

// Watch for changes in selected orders
watch(() => props.selectedOrders, (newSelection) => {
  // Update local state if needed
}, { deep: true })
</script>

<style scoped>
.space-x-reverse > :not([hidden]) ~ :not([hidden]) {
  --tw-space-x-reverse: 1;
}
</style> 

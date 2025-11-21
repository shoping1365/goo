<template>
  <div class="min-h-screen bg-gray-50 py-8" dir="rtl">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <!-- لایوت اصلی با سایدبار -->
      <div class="flex flex-col lg:flex-row gap-8">
        <!-- سایدبار -->
        <div class="lg:w-80 flex-shrink-8">
          <AccountSidebar />
        </div>

        <!-- محتوای اصلی -->
        <div class="flex-1">
          <!-- عنوان و تب‌های فیلتر وضعیت -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-6">
            <h1 class="text-sm font-bold text-gray-900 text-right mb-4">تاریخچه سفارشات</h1>
            <div class="flex border-b border-gray-200">
              <button 
                v-for="(filter, index) in orderFilters" 
                :key="filter.key"
                class="pb-4 px-2 text-gray-500 hover:text-gray-700 transition-colors relative"
                :class="[
                  { 'text-red-600': activeFilter === filter.key },
                  index < orderFilters.length - 1 ? 'ml-8' : ''
                ]"
                @click="activeFilter = filter.key"
              >
                {{ filter.label }}
                <span v-if="filter.count > 0" class="mr-2 text-sm font-medium">{{ filter.count }}</span>
                <div v-if="activeFilter === filter.key" class="absolute bottom-0 left-0 right-0 h-0.5 bg-red-600"></div>
              </button>
            </div>
          </div>

          <div v-if="loading" class="text-center text-gray-500">در حال بارگذاری ...</div>
          <div v-else-if="filteredOrders.length === 0" class="text-center text-gray-500">هیچ سفارشی یافت نشد.</div>

          <div v-else class="space-y-6">
            <!-- کارت سفارش -->
            <div v-for="order in filteredOrders" :key="order.id" class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 cursor-pointer" @click="openOrderDetail(order)">
              <!-- وضعیت کلی سفارش -->
              <div class="mb-6">
                <div class="text-right mb-4">
                  <div class="flex items-center justify-between mb-2">
                    <div class="text-base">کد سفارش {{ order.order_number || order.id }}</div>
                    <div class="text-sm text-gray-600">{{ formatDate(order.created_at) }}</div>
                  </div>
                  <div class="flex items-center justify-between">
                    <div class="text-sm">مبلغ {{ formatPrice(order.final_amount) }}</div>
                    <div class="flex items-center space-x-3 space-x-reverse">
                      <div class="w-3 h-3 rounded-full" :class="getStatusColor(order.status)"></div>
                      <span class="font-medium" :class="getStatusTextColor(order.status)">{{ getStatusText(order.status) }}</span>
                    </div>
                  </div>
                </div>
              </div>

              <!-- مرسوله‌ها -->
              <div class="space-y-4">
                <div v-for="(shipment, index) in getShipments(order)" :key="index" class="border border-gray-200 rounded-lg p-6">
                  <!-- هدر مرسوله -->
                  <div class="flex items-center justify-between mb-4">
                    <div class="flex items-center space-x-2 space-x-reverse">
                      <svg class="w-4 h-4 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"></path>
                      </svg>
                      <span class="text-sm font-medium">مرسوله {{ index + 1 }} از {{ getShipments(order).length }}</span>
                    </div>
                  </div>

                  <!-- تصویر محصول -->
                  <div class="mb-4">
                    <img :src="shipment.product_image || '/default-product.svg'" :alt="shipment.product_name" class="w-16 h-16 object-cover rounded-lg cursor-pointer" @click.stop="goToProduct(shipment.product_id)">
                  </div>

                  <!-- نوار پیشرفت -->
                  <div class="space-y-2">
                    <div class="flex items-center justify-between text-sm">
                      <span class="text-gray-500">مرحله بعد: تحویل به مامور ارسال</span>
                      <span class="text-green-600 font-medium">ورود به مرکز توزیع کرج</span>
                    </div>
                    <div class="w-full bg-gray-200 rounded-full h-2">
                      <div class="bg-green-500 h-2 rounded-full" style="width: 60%"></div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import AccountSidebar from '@/components/account/AccountSidebar.vue'
import { computed, onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'

// تعریف definePageMeta برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string }) => void

// navigateTo is available globally in Nuxt 3

// تعریف interface برای Order Item
interface OrderItem {
  id: number
  product_id: number
  product_name: string
  quantity: number
  price: number
  product_image?: string
}

// تعریف interface برای Order
interface Order {
  id: number
  order_number?: string
  status: 'processing' | 'shipped' | 'pending' | 'delivered' | 'refunded' | 'cancelled'
  payment_status?: string
  final_amount?: number
  created_at: string
  items: OrderItem[]
}

// تعریف interface برای API Response
interface OrdersApiResponse {
  success: boolean
  data: Order[]
  message?: string
}

const router = useRouter()

// middleware حذف شده
definePageMeta({ layout: 'default' })

// const { isAuthenticated } = useAuth() // احراز هویت غیرفعال شده است
// const router = useRouter()

const orders = ref<Order[]>([])
const loading = ref<boolean>(false)
const activeFilter = ref<string>('current')

// فیلترهای سفارش - شمارش واقعی
const orderFilters = computed(() => {
  if (!orders.value || !Array.isArray(orders.value)) {
    return [
      { key: 'current', label: 'جاری', count: 0 },
      { key: 'delivered', label: 'تحویل شده', count: 0 },
      { key: 'returned', label: 'مرجوع شده', count: 0 },
      { key: 'cancelled', label: 'لغو شده', count: 0 }
    ]
  }

  const currentCount = orders.value.filter(order => 
    order.status === 'processing' || order.status === 'shipped' || order.status === 'pending'
  ).length

  const deliveredCount = orders.value.filter(order => 
    order.status === 'delivered'
  ).length

  const returnedCount = orders.value.filter(order => 
    order.status === 'refunded'
  ).length

  const cancelledCount = orders.value.filter(order => 
    order.status === 'cancelled'
  ).length

  return [
    { key: 'current', label: 'جاری', count: currentCount },
    { key: 'delivered', label: 'تحویل شده', count: deliveredCount },
    { key: 'returned', label: 'مرجوع شده', count: returnedCount },
    { key: 'cancelled', label: 'لغو شده', count: cancelledCount }
  ]
})

// سفارشات فیلتر شده
const filteredOrders = computed(() => {
  if (!orders.value || !Array.isArray(orders.value)) {
    return []
  }
  
  if (activeFilter.value === 'current') {
    return orders.value.filter(order => 
      order.status === 'processing' || order.status === 'shipped' || order.status === 'pending'
    )
  } else if (activeFilter.value === 'delivered') {
    return orders.value.filter(order => order.status === 'delivered')
  } else if (activeFilter.value === 'returned') {
    return orders.value.filter(order => order.status === 'refunded')
  } else if (activeFilter.value === 'cancelled') {
    return orders.value.filter(order => order.status === 'cancelled')
  }
  return orders.value
})

onMounted(async () => {
  try {
    loading.value = true

    const response = await $fetch<OrdersApiResponse>('/api/orders', { credentials: 'include' })

    if (response.success) {
      orders.value = response.data

    } else {
      console.error('خطا در دریافت سفارش‌ها:', response.message)
      // در صورت خطا، داده‌های نمونه اضافه می‌کنیم
      orders.value = [
        {
          id: 1,
          order_number: 'ORD-123456',
          status: 'processing',
          payment_status: 'paid',
          final_amount: 3715500,
          created_at: '2024-10-01T10:00:00Z',
          items: [
            {
              id: 1,
              product_id: 1,
              product_name: 'کابل HDMI',
              quantity: 2,
              price: 50000
            },
            {
              id: 2,
              product_id: 2,
              product_name: 'بازوی مانیتور',
              quantity: 1,
              price: 200000
            }
          ]
        }
      ]
    }
  } catch (e) {
    console.error('خطا در دریافت سفارش‌ها', e)
    console.error('جزئیات خطا:', {
      message: e.message,
      status: e.status,
      statusText: e.statusText,
      data: e.data
    })
    // در صورت خطا، داده‌های نمونه اضافه می‌کنیم
    orders.value = [
      {
        id: 1,
        order_number: 'ORD-123456',
        status: 'processing',
        payment_status: 'paid',
        final_amount: 3715500,
        created_at: '2024-10-01T10:00:00Z',
        items: [
          {
            id: 1,
            product_id: 1,
            product_name: 'کابل HDMI',
            quantity: 2,
            price: 50000
          },
          {
            id: 2,
            product_id: 2,
            product_name: 'بازوی مانیتور',
            quantity: 1,
            price: 200000
          }
        ]
      }
    ]
  } finally {
    loading.value = false
  }
})

function formatPrice(val: number | null | undefined) {
  if (!val || isNaN(val)) return 'نامشخص'
  return new Intl.NumberFormat('fa-IR').format(val) + ' تومان'
}

function formatDate(d: string | null | undefined) {
  if (!d) return 'نامشخص'
  try {
    const date = new Date(d)
    if (isNaN(date.getTime())) return 'نامشخص'
    return date.toLocaleDateString('fa-IR')
  } catch {
    return 'نامشخص'
  }
}

// تابع برای دریافت مرسوله‌های سفارش
function getShipments(order: { items?: Array<Record<string, unknown>>; [key: string]: unknown }) {
  if (!order.items || order.items.length === 0) {
    return [{
      product_name: 'محصول سفارش',
      quantity: 1,
      product_image: '/default-product.svg',
      product_id: null
    }]
  }
  
  return order.items.map((item: OrderItem) => ({
    product_name: item.product_name || `محصول ${item.product_id}`,
    quantity: item.quantity || 1,
    product_image: item.product_image || '/default-product.svg',
    product_id: item.product_id
  }))
}

function openOrderDetail(order: Order) {
  // هدایت به صفحه جزئیات سفارش
  router.push(`/account/orders/${order.id}`)
}

function goToProduct(productId: number) {
  if (productId) {
    // هدایت به صفحه محصول
    router.push(`/product/${productId}`)
  }
}

// توابع وضعیت سفارش
function getStatusColor(status: string) {
  switch (status) {
    case 'processing':
    case 'pending':
      return 'bg-yellow-500'
    case 'shipped':
      return 'bg-blue-500'
    case 'delivered':
      return 'bg-green-500'
    case 'cancelled':
      return 'bg-red-500'
    case 'refunded':
      return 'bg-gray-500'
    default:
      return 'bg-gray-400'
  }
}

function getStatusTextColor(status: string) {
  switch (status) {
    case 'processing':
    case 'pending':
      return 'text-yellow-600'
    case 'shipped':
      return 'text-blue-600'
    case 'delivered':
      return 'text-green-600'
    case 'cancelled':
      return 'text-red-600'
    case 'refunded':
      return 'text-gray-600'
    default:
      return 'text-gray-500'
  }
}

function getStatusText(status: string) {
  switch (status) {
    case 'processing':
      return 'در حال پردازش'
    case 'pending':
      return 'در انتظار'
    case 'shipped':
      return 'ارسال شده'
    case 'delivered':
      return 'تحویل شده'
    case 'cancelled':
      return 'لغو شده'
    case 'refunded':
      return 'مرجوع شده'
    default:
      return 'نامشخص'
  }
}
</script> 
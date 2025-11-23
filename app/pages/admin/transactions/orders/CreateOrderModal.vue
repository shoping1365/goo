<template>
  <div v-if="isOpen && hasAccess" class="fixed inset-0 z-50 overflow-y-auto">
    <!-- Backdrop -->
    <div class="fixed inset-0 bg-black bg-opacity-50 transition-opacity" @click="closeModal"></div>
    
    <!-- Modal -->
    <div class="flex min-h-full items-center justify-center px-4 py-4">
      <div class="relative w-full max-w-4xl bg-white rounded-lg shadow-xl">
        <!-- Header -->
        <div class="flex items-center justify-between px-4 py-4 border-b border-gray-200">
          <div class="flex items-center space-x-3">
            <div class="w-10 h-10 bg-green-100 rounded-full flex items-center justify-center">
              <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
              </svg>
            </div>
            <div>
              <h3 class="text-lg font-semibold text-gray-900">ایجاد سفارش جدید</h3>
              <p class="text-sm text-gray-500">اطلاعات سفارش را وارد کنید</p>
            </div>
          </div>
          <button class="text-gray-400 hover:text-gray-600 transition-colors" @click="closeModal">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>

        <!-- Content -->
        <div class="px-4 py-4 max-h-[70vh] overflow-y-auto">
          <form class="space-y-6" @submit.prevent="createOrder">
            <!-- Customer Information -->
            <div class="bg-gray-50 px-4 py-4 rounded-lg">
              <h4 class="text-md font-semibold text-gray-900 mb-4">اطلاعات مشتری</h4>
              <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">نام و نام خانوادگی *</label>
                  <input
                    v-model="form.customer.name"
                    type="text"
                    required
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors"
                    placeholder="نام مشتری"
                  />
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">شماره تلفن *</label>
                  <input
                    v-model="form.customer.phone"
                    type="tel"
                    required
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors"
                    placeholder="09123456789"
                  />
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">ایمیل</label>
                  <input
                    v-model="form.customer.email"
                    type="email"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors"
                    placeholder="example@email.com"
                  />
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">کد ملی</label>
                  <input
                    v-model="form.customer.nationalCode"
                    type="text"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors"
                    placeholder="1234567890"
                  />
                </div>
              </div>
            </div>

            <!-- Shipping Information -->
            <div class="bg-gray-50 px-4 py-4 rounded-lg">
              <h4 class="text-md font-semibold text-gray-900 mb-4">اطلاعات ارسال</h4>
              <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">استان *</label>
                  <select
                    v-model="form.shipping.province"
                    required
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors"
                  >
                    <option value="">انتخاب استان</option>
                    <option value="تهران">تهران</option>
                    <option value="اصفهان">اصفهان</option>
                    <option value="خراسان رضوی">خراسان رضوی</option>
                    <option value="فارس">فارس</option>
                    <option value="آذربایجان شرقی">آذربایجان شرقی</option>
                  </select>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">شهر *</label>
                  <input
                    v-model="form.shipping.city"
                    type="text"
                    required
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors"
                    placeholder="نام شهر"
                  />
                </div>
                <div class="md:col-span-2">
                  <label class="block text-sm font-medium text-gray-700 mb-2">آدرس کامل *</label>
                  <textarea
                    v-model="form.shipping.address"
                    required
                    rows="3"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors"
                    placeholder="آدرس کامل پستی"
                  ></textarea>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">کد پستی *</label>
                  <input
                    v-model="form.shipping.postalCode"
                    type="text"
                    required
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors"
                    placeholder="1234567890"
                  />
                </div>
              </div>
            </div>

            <!-- Order Items -->
            <div class="bg-gray-50 px-4 py-4 rounded-lg">
              <div class="flex items-center justify-between mb-4">
                <h4 class="text-md font-semibold text-gray-900">محصولات سفارش</h4>
                <button
                  type="button"
                  class="px-3 py-1 text-sm text-blue-600 hover:text-blue-800 border border-blue-300 hover:border-blue-500 rounded-lg transition-colors"
                  @click="addItem"
                >
                  + افزودن محصول
                </button>
              </div>
              
              <div class="space-y-4">
                <div v-for="(item, index) in form.items" :key="index" class="flex items-center space-x-4 px-4 py-4 bg-white rounded-lg border border-gray-200">
                  <div class="flex-1">
                    <input
                      v-model="item.productName"
                      type="text"
                      required
                      class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors"
                      placeholder="نام محصول"
                    />
                  </div>
                  <div class="w-24">
                    <input
                      v-model="item.sku"
                      type="text"
                      class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors"
                      placeholder="SKU"
                    />
                  </div>
                  <div class="w-24">
                    <input
                      v-model="item.quantity"
                      type="number"
                      required
                      min="1"
                      class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors"
                      placeholder="تعداد"
                    />
                  </div>
                  <div class="w-32">
                    <input
                      v-model="item.price"
                      type="number"
                      required
                      min="0"
                      class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors"
                      placeholder="قیمت (تومان)"
                    />
                  </div>
                  <button
                    type="button"
                    class="text-red-600 hover:text-red-800 p-1 rounded transition-colors"
                    @click="removeItem(index)"
                  >
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                    </svg>
                  </button>
                </div>
              </div>
            </div>

            <!-- Order Details -->
            <div class="bg-gray-50 px-4 py-4 rounded-lg">
              <h4 class="text-md font-semibold text-gray-900 mb-4">جزئیات سفارش</h4>
              <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت سفارش</label>
                  <select
                    v-model="form.status"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors"
                  >
                    <option value="pending">در انتظار پرداخت</option>
                    <option value="paid">پرداخت شده</option>
                    <option value="processing">در حال پردازش</option>
                    <option value="shipped">ارسال شده</option>
                    <option value="delivered">تحویل داده شده</option>
                  </select>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">روش پرداخت</label>
                  <select
                    v-model="form.payment.method"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors"
                  >
                    <option value="online">پرداخت آنلاین</option>
                    <option value="cash">پرداخت نقدی</option>
                    <option value="bank_transfer">انتقال بانکی</option>
                  </select>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">توضیحات</label>
                  <textarea
                    v-model="form.notes"
                    rows="3"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors"
                    placeholder="توضیحات اضافی..."
                  ></textarea>
                </div>
              </div>
            </div>

            <!-- Order Summary -->
            <div class="bg-blue-50 px-4 py-4 rounded-lg">
              <h4 class="text-md font-semibold text-blue-900 mb-4">خلاصه سفارش</h4>
              <div class="grid grid-cols-1 md:grid-cols-3 gapx-4 py-4">
                <div>
                  <p class="text-sm text-blue-600">تعداد محصولات</p>
                  <p class="text-lg font-semibold text-blue-900">{{ form.items.length }}</p>
                </div>
                <div>
                  <p class="text-sm text-blue-600">مجموع قیمت</p>
                  <p class="text-lg font-semibold text-blue-900">{{ formatPrice(totalAmount) }} تومان</p>
                </div>
                <div>
                  <p class="text-sm text-blue-600">وضعیت</p>
                  <p class="text-lg font-semibold text-blue-900">{{ getStatusText(form.status) }}</p>
                </div>
              </div>
            </div>
          </form>
        </div>

        <!-- Footer -->
        <div class="flex items-center justify-end space-x-3 px-4 py-4 border-t border-gray-200">
          <button class="px-4 py-2 text-gray-700 bg-gray-100 hover:bg-gray-200 rounded-lg transition-colors" @click="closeModal">
            انصراف
          </button>
          <button
            :disabled="isSubmitting"
            class="px-6 py-2 text-white bg-blue-600 hover:bg-blue-700 disabled:bg-gray-400 rounded-lg transition-colors flex items-center"
            @click="createOrder"
          >
            <svg v-if="isSubmitting" class="animate-spin -ml-1 mr-3 h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            {{ isSubmitting ? 'در حال ایجاد...' : 'ایجاد سفارش' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>
</script>

<script setup lang="ts">
import { computed, ref, watch } from 'vue';
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

// بررسی احراز هویت هنگام باز شدن modal
const _props = defineProps({
  isOpen: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['close', 'created'])

const isSubmitting = ref(false)

const form = ref({
  customer: {
    name: '',
    phone: '',
    email: '',
    nationalCode: ''
  },
  shipping: {
    province: '',
    city: '',
    address: '',
    postalCode: ''
  },
  items: [
    {
      productName: '',
      sku: '',
      quantity: 1,
      price: 0
    }
  ],
  status: 'pending',
  payment: {
    method: 'online'
  },
  notes: ''
})

const totalAmount = computed(() => {
  return form.value.items.reduce((total, item) => {
    const price = typeof item.price === 'number' ? item.price : parseInt(String(item.price)) || 0
    const quantity = typeof item.quantity === 'number' ? item.quantity : parseInt(String(item.quantity)) || 0
    return total + price * quantity
  }, 0)
})

const closeModal = () => {
  emit('close')
}

const addItem = () => {
  form.value.items.push({
    productName: '',
    sku: '',
    quantity: 1,
    price: 0
  })
}

const removeItem = (index) => {
  if (form.value.items.length > 1) {
    form.value.items.splice(index, 1)
  }
}

const getStatusText = (status) => {
  const statusMap = {
    'pending': 'در انتظار پرداخت',
    'paid': 'پرداخت شده',
    'processing': 'در حال پردازش',
    'shipped': 'ارسال شده',
    'delivered': 'تحویل داده شده'
  }
  return statusMap[status] || status
}

const formatPrice = (price) => {
  if (!price) return '0'
  return new Intl.NumberFormat('fa-IR').format(price)
}

const createOrder = async () => {
  if (isSubmitting.value) return

  // Validate required fields
  if (!form.value.customer.name || !form.value.customer.phone) {
    alert('لطفاً اطلاعات مشتری را تکمیل کنید')
    return
  }

  if (!form.value.shipping.province || !form.value.shipping.city || !form.value.shipping.address) {
    alert('لطفاً اطلاعات ارسال را تکمیل کنید')
    return
  }

  if (form.value.items.some(item => !item.productName || !item.price)) {
    alert('لطفاً اطلاعات محصولات را تکمیل کنید')
    return
  }

  isSubmitting.value = true

  try {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 1000))
    
    // Create order object
    const order = {
      id: Date.now(),
      orderNumber: `ORD-${Date.now()}`,
      customer: { ...form.value.customer },
      shipping: { ...form.value.shipping },
      items: form.value.items.map(item => ({ ...item })),
      status: form.value.status,
      payment: { ...form.value.payment },
      notes: form.value.notes,
      totalAmount: totalAmount.value,
      createdAt: new Date().toISOString()
    }

    emit('created', order)
    closeModal()
    
    // Reset form
    form.value = {
      customer: { name: '', phone: '', email: '', nationalCode: '' },
      shipping: { province: '', city: '', address: '', postalCode: '' },
      items: [{ productName: '', sku: '', quantity: 1, price: 0 }],
      status: 'pending',
      payment: { method: 'online' },
      notes: ''
    }
  } catch (error) {
    console.error('Error creating order:', error)
    alert('خطا در ایجاد سفارش')
  } finally {
    isSubmitting.value = false
  }
}
</script> 

<template>
  <div v-if="show" class="fixed inset-0 bg-black bg-opacity-50 z-50 overflow-y-auto" @click="closeModal">
    <div class="min-h-screen px-4 text-center">
      <div class="inline-block align-bottom bg-white rounded-xl text-right overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-5xl sm:w-full" @click.stop>
        
        <!-- Header -->
        <div class="bg-gradient-to-r from-purple-600 to-purple-700 px-6 py-4">
          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-10 h-10 bg-white bg-opacity-20 rounded-lg flex items-center justify-center ml-3">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
                </svg>
              </div>
              <div>
                <h3 class="text-xl font-bold text-white">ویرایش سفارش #{{ editedOrder.order_number }}</h3>
                <p class="text-purple-100 text-sm">تغییر اطلاعات سفارش و وضعیت</p>
              </div>
            </div>
            <button class="text-white hover:text-gray-200 transition-colors" @click="closeModal">
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
              </svg>
            </button>
          </div>
        </div>

        <!-- Content -->
        <div class="max-h-[calc(100vh-200px)] overflow-y-auto bg-gray-50">
          <div class="px-4 py-4">
            <form class="space-y-6" @submit.prevent="saveOrder">
              
              <!-- تب‌ها -->
              <div class="border-b border-gray-200">
                <nav class="-mb-px flex space-x-8 space-x-reverse">
                  <button 
                    v-for="tab in tabs" 
                    :key="tab.id"
                    type="button"
                    :class="[
                      'py-2 px-1 border-b-2 font-medium text-sm',
                      activeTab === tab.id 
                        ? 'border-purple-500 text-purple-600' 
                        : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
                    ]"
                    @click="activeTab = tab.id"
                  >
                    {{ tab.name }}
                  </button>
                </nav>
              </div>

              <!-- محتوای تب‌ها -->
              <div class="bg-white rounded-lg border border-gray-200 px-4 py-4">
                
                <!-- تب اطلاعات کلی -->
                <div v-if="activeTab === 'general'" class="space-y-6">
                  <div class="grid grid-cols-1 md:grid-cols-3 gapx-4 py-4">
                    <div>
                      <label class="block text-sm font-medium text-gray-700 mb-2">شماره سفارش</label>
                      <input 
                        v-model="editedOrder.order_number" 
                        type="text"
                        class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500 text-left"
                        readonly
                      >
                    </div>
                    
                    <div>
                      <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت سفارش</label>
                      <select 
                        v-model="editedOrder.status" 
                        class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
                      >
                        <option value="pending">در انتظار تایید</option>
                        <option value="confirmed">تایید شده</option>
                        <option value="processing">در حال آماده‌سازی</option>
                        <option value="shipped">ارسال شده</option>
                        <option value="delivered">تحویل شده</option>
                        <option value="cancelled">لغو شده</option>
                        <option value="returned">مرجوع شده</option>
                      </select>
                    </div>

                    <div>
                      <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت پرداخت</label>
                      <select 
                        v-model="editedOrder.payment_status" 
                        class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
                      >
                        <option value="pending">در انتظار پرداخت</option>
                        <option value="paid">پرداخت شده</option>
                        <option value="failed">پرداخت ناموفق</option>
                        <option value="refunded">بازپرداخت شده</option>
                      </select>
                    </div>
                  </div>

                  <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
                    <div>
                      <label class="block text-sm font-medium text-gray-700 mb-2">اولویت سفارش</label>
                      <select 
                        v-model="editedOrder.priority" 
                        class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
                      >
                        <option value="low">کم</option>
                        <option value="normal">عادی</option>
                        <option value="high">بالا</option>
                        <option value="urgent">فوری</option>
                      </select>
                    </div>

                    <div>
                      <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ تحویل مورد انتظار</label>
                      <input 
                        v-model="editedOrder.expected_delivery" 
                        type="date"
                        class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
                      >
                    </div>
                  </div>

                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-2">یادداشت ادمین</label>
                    <textarea 
                      v-model="editedOrder.admin_notes" 
                      rows="4"
                      class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
                      placeholder="یادداشت‌های مربوط به سفارش..."
                    ></textarea>
                  </div>
                </div>

                <!-- تب اطلاعات مشتری -->
                <div v-if="activeTab === 'customer'" class="space-y-6">
                  <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
                    <div>
                      <label class="block text-sm font-medium text-gray-700 mb-2">نام و نام خانوادگی *</label>
                      <input 
                        v-model="editedOrder.customer.name" 
                        type="text"
                        required
                        class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
                      >
                    </div>

                    <div>
                      <label class="block text-sm font-medium text-gray-700 mb-2">شماره تماس *</label>
                      <input 
                        v-model="editedOrder.customer.phone" 
                        type="tel"
                        required
                        class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
                        dir="ltr"
                      >
                    </div>
                  </div>

                  <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
                    <div>
                      <label class="block text-sm font-medium text-gray-700 mb-2">ایمیل</label>
                      <input 
                        v-model="editedOrder.customer.email" 
                        type="email"
                        class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
                        dir="ltr"
                      >
                    </div>

                    <div>
                      <label class="block text-sm font-medium text-gray-700 mb-2">نوع مشتری</label>
                      <select 
                        v-model="editedOrder.customer.type" 
                        class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
                      >
                        <option value="regular">عادی</option>
                        <option value="vip">ویژه</option>
                        <option value="wholesale">عمده</option>
                      </select>
                    </div>
                  </div>
                </div>

                <!-- تب محصولات -->
                <div v-if="activeTab === 'products'" class="space-y-6">
                  <div class="flex justify-between items-center">
                    <h5 class="text-lg font-semibold text-gray-900">محصولات سفارش</h5>
                    <button 
                      type="button"
                      class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors flex items-center"
                      @click="addProduct"
                    >
                      <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
                      </svg>
                      افزودن محصول
                    </button>
                  </div>

                  <div class="space-y-4">
                    <div 
                      v-for="(item, index) in editedOrder.items" 
                      :key="index"
                      class="bg-gray-50 rounded-lg border px-4 py-4"
                    >
                      <div class="grid grid-cols-1 md:grid-cols-5 gapx-4 py-4 items-end">
                        <div>
                          <label class="block text-sm font-medium text-gray-700 mb-2">نام محصول</label>
                          <input 
                            v-model="item.name" 
                            type="text"
                            required
                            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
                          >
                        </div>

                        <div>
                          <label class="block text-sm font-medium text-gray-700 mb-2">کد محصول</label>
                          <input 
                            v-model="item.sku" 
                            type="text"
                            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
                          >
                        </div>

                        <div>
                          <label class="block text-sm font-medium text-gray-700 mb-2">قیمت (تومان)</label>
                          <input 
                            v-model.number="item.price" 
                            type="number"
                            min="0"
                            required
                            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
                            @input="calculateTotals"
                          >
                        </div>

                        <div>
                          <label class="block text-sm font-medium text-gray-700 mb-2">تعداد</label>
                          <input 
                            v-model.number="item.quantity" 
                            type="number"
                            min="1"
                            required
                            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
                            @input="calculateTotals"
                          >
                        </div>

                        <div class="flex space-x-2 space-x-reverse">
                          <button 
                            type="button"
                            class="px-3 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700 transition-colors"
                            @click="removeProduct(index)"
                          >
                            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                            </svg>
                          </button>
                        </div>
                      </div>
                    </div>
                  </div>

                  <!-- محاسبات مالی -->
                  <div class="bg-blue-50 rounded-lg px-4 py-4 border border-blue-200">
                    <div class="space-y-3">
                      <div class="flex justify-between">
                        <span class="text-gray-700">جمع محصولات:</span>
                        <span class="font-semibold">{{ formatPrice(calculatedSubtotal) }} تومان</span>
                      </div>
                      
                      <div class="flex justify-between items-center">
                        <span class="text-gray-700">تخفیف:</span>
                        <div class="flex items-center space-x-2 space-x-reverse">
                          <input 
                            v-model.number="editedOrder.discount" 
                            type="number"
                            min="0"
                            class="w-20 px-2 py-1 text-sm border border-gray-300 rounded focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
                            @input="calculateTotals"
                          >
                          <select 
                            v-model="discountType" 
                            class="px-2 py-1 text-sm border border-gray-300 rounded focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
                            @change="calculateTotals"
                          >
                            <option value="percent">درصد</option>
                            <option value="fixed">تومان</option>
                          </select>
                        </div>
                      </div>
                      
                      <div class="flex justify-between">
                        <span class="text-gray-700">هزینه ارسال:</span>
                        <span class="font-semibold">{{ formatPrice(editedOrder.shipping_cost) }} تومان</span>
                      </div>
                      
                      <div class="flex justify-between border-t pt-3 text-lg font-bold">
                        <span>مبلغ نهایی:</span>
                        <span class="text-green-600">{{ formatPrice(calculatedTotal) }} تومان</span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </form>
          </div>
        </div>

        <!-- Footer -->
        <div class="bg-gray-100 px-6 py-4 flex justify-between">
          <div class="flex space-x-2 space-x-reverse">
            <button 
              type="button"
              class="px-6 py-2 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors"
              @click="resetForm"
            >
              بازنشانی
            </button>
          </div>
          
          <div class="flex space-x-2 space-x-reverse">
            <button 
              type="button"
              class="px-6 py-2 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors" 
              @click="closeModal"
            >
              لغو
            </button>
            <button 
              type="button"
              :disabled="isSaving"
              :class="[
                'px-6 py-2 rounded-lg font-medium transition-colors',
                isSaving
                  ? 'bg-gray-400 text-gray-200 cursor-not-allowed'
                  : 'bg-purple-600 text-white hover:bg-purple-700'
              ]"
              @click="saveOrder"
            >
              <div v-if="isSaving" class="flex items-center">
                <svg class="animate-spin -ml-1 mr-3 h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                  <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                </svg>
                در حال ذخیره...
              </div>
              <span v-else>ذخیره تغییرات</span>
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, watch } from 'vue';

interface Props {
  show: boolean
  order: Record<string, unknown>
}

const props = defineProps<Props>()
const emit = defineEmits(['close', 'save'])

// State
const activeTab = ref('general')
const isSaving = ref(false)
const discountType = ref('percent')

// Tabs
const tabs = [
  { id: 'general', name: 'اطلاعات کلی' },
  { id: 'customer', name: 'مشتری' },
  { id: 'products', name: 'محصولات' }
]

// Form data
const editedOrder = ref({
  order_number: '',
  status: 'pending',
  payment_status: 'pending',
  priority: 'normal',
  expected_delivery: '',
  admin_notes: '',
  shipping_cost: 0,
  discount: 0,
  total_amount: 0,
  subtotal: 0,
  customer: {
    name: '',
    phone: '',
    email: '',
    type: 'regular'
  },
  items: []
})

// Watch for prop changes
watch(() => props.order, (newOrder) => {
  if (newOrder && props.show) {
    resetForm()
  }
}, { immediate: true })

// Watch for show prop changes
watch(() => props.show, (newShow) => {
  if (newShow && props.order) {
    resetForm()
  }
  if (!newShow) {
    // Reset active tab when modal closes
    activeTab.value = 'general'
  }
})

// Computed
const calculatedSubtotal = computed(() => {
  return editedOrder.value.items.reduce((sum, item: { price: number; quantity: number }) => sum + (item.price * item.quantity), 0)
})

const calculatedTotal = computed(() => {
  const subtotal = calculatedSubtotal.value
  let discount = 0
  
  if (discountType.value === 'percent') {
    discount = (subtotal * editedOrder.value.discount) / 100
  } else {
    discount = editedOrder.value.discount
  }
  
  return subtotal - discount + editedOrder.value.shipping_cost
})

// Methods
const formatPrice = (price: number): string => {
  if (!price) return '0'
  return new Intl.NumberFormat('fa-IR').format(price)
}

const resetForm = () => {
  if (props.order) {
    editedOrder.value = JSON.parse(JSON.stringify(props.order))
  } else {
    editedOrder.value = {
      order_number: '12345',
      status: 'processing',
      payment_status: 'paid',
      priority: 'normal',
      expected_delivery: '',
      admin_notes: '',
      shipping_cost: 200000,
      discount: 0,
      total_amount: 12200000,
      subtotal: 12000000,
      customer: {
        name: 'علی احمدی',
        phone: '09121234567',
        email: 'ali@example.com',
        type: 'regular'
      },
      items: [
        {
          id: 1,
          name: 'گوشی موبایل سامسونگ Galaxy S23',
          sku: 'SAM-S23-128',
          price: 12000000,
          quantity: 1
        }
      ]
    }
  }
}

const calculateTotals = () => {
  // Force reactivity update
}

const addProduct = () => {
  editedOrder.value.items.push({
    id: Date.now(),
    name: '',
    sku: '',
    price: 0,
    quantity: 1
  })
}

const removeProduct = (index: number) => {
  editedOrder.value.items.splice(index, 1)
}

const saveOrder = async () => {
  isSaving.value = true
  
  try {
    // Validate required fields
    if (!editedOrder.value.customer.name.trim()) {
      alert('نام مشتری الزامی است')
      return
    }
    
    if (!editedOrder.value.customer.phone.trim()) {
      alert('شماره تماس الزامی است')
      return
    }
    
    if (editedOrder.value.items.length === 0) {
      alert('حداقل یک محصول باید در سفارش وجود داشته باشد')
      return
    }
    
    // Update total amount
    editedOrder.value.total_amount = calculatedTotal.value
    editedOrder.value.subtotal = calculatedSubtotal.value
    
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 1500))
    
    emit('save', editedOrder.value)
    // console.log('سفارش ذخیره شد:', editedOrder.value)
    
  } catch {
    // console.error('خطا در ذخیره سفارش:', error)
    alert('خطا در ذخیره سفارش. لطفاً دوباره تلاش کنید.')
  } finally {
    isSaving.value = false
  }
}

const closeModal = () => {
  if (!isSaving.value) {
    emit('close')
  }
}
</script>

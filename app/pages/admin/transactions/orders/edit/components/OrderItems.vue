<template>
  <div class="px-4 py-4">
    <div class="flex justify-between items-center mb-4">
      <h3 class="text-lg font-semibold text-gray-900 flex items-center">
        <svg class="w-5 h-5 text-green-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"></path>
        </svg>
        محصولات سفارش
      </h3>
      <button 
        @click="addNewItem"
        class="px-3 py-1.5 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors flex items-center text-sm"
      >
        <svg class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
        </svg>
        افزودن محصول
      </button>
    </div>

    <div class="overflow-x-auto">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              محصول
            </th>
            <th class="px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              SKU
            </th>
            <th class="px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              قیمت واحد
            </th>
            <th class="px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              تعداد
            </th>
            <th class="px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              تخفیف
            </th>
            <th class="px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              قیمت کل
            </th>
            <th class="px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              عملیات
            </th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-for="(item, index) in orderItems" :key="index" class="hover:bg-gray-50 transition-colors">
            <!-- محصول -->
            <td class="px-4 py-4">
              <div class="flex items-center">
                <img 
                  :src="item.image || '/default-product.svg'" 
                  :alt="item.name"
                  class="w-12 h-12 rounded-lg object-cover mr-3"
                />
                <div>
                  <div class="text-sm font-medium text-gray-900">{{ item.name }}</div>
                  <div class="text-sm text-gray-500">{{ item.category }}</div>
                </div>
              </div>
            </td>

            <!-- SKU -->
            <td class="px-4 py-4">
              <input 
                v-model="item.sku"
                type="text" 
                class="w-full px-2 py-1 border border-gray-300 rounded text-sm focus:ring-1 focus:ring-blue-500 focus:border-blue-500"
                placeholder="SKU محصول"
              />
            </td>

            <!-- قیمت واحد -->
            <td class="px-4 py-4">
              <input 
                v-model="item.unitPrice"
                type="number" 
                class="w-full px-2 py-1 border border-gray-300 rounded text-sm focus:ring-1 focus:ring-blue-500 focus:border-blue-500"
                placeholder="0"
                @input="calculateItemTotal(index)"
              />
            </td>

            <!-- تعداد -->
            <td class="px-4 py-4">
              <div class="flex items-center">
                <button 
                  @click="decreaseQuantity(index)"
                  class="w-8 h-8 border border-gray-300 rounded-r flex items-center justify-center hover:bg-gray-50"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 12H4"></path>
                  </svg>
                </button>
                <input 
                  v-model="item.quantity"
                  type="number" 
                  class="w-16 px-2 py-1 border-t border-b border-gray-300 text-center text-sm focus:ring-1 focus:ring-blue-500 focus:border-blue-500"
                  @input="calculateItemTotal(index)"
                />
                <button 
                  @click="increaseQuantity(index)"
                  class="w-8 h-8 border border-gray-300 rounded-l flex items-center justify-center hover:bg-gray-50"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
                  </svg>
                </button>
              </div>
            </td>

            <!-- تخفیف -->
            <td class="px-4 py-4">
              <input 
                v-model="item.discount"
                type="number" 
                class="w-full px-2 py-1 border border-gray-300 rounded text-sm focus:ring-1 focus:ring-blue-500 focus:border-blue-500"
                placeholder="0"
                @input="calculateItemTotal(index)"
              />
            </td>

            <!-- قیمت کل -->
            <td class="px-4 py-4">
              <div class="text-sm font-medium text-gray-900">
                {{ formatPrice(item.totalPrice) }} تومان
              </div>
            </td>

            <!-- عملیات -->
            <td class="px-4 py-4">
              <button 
                @click="removeItem(index)"
                class="text-red-600 hover:text-red-800 transition-colors"
              >
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                </svg>
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- خلاصه سفارش -->
    <div class="mt-6 pt-6 border-t border-gray-200">
      <div class="flex justify-end">
        <div class="w-64 space-y-2">
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">جمع کل:</span>
            <span class="font-medium">{{ formatPrice(subtotal) }} تومان</span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">تخفیف کل:</span>
            <span class="font-medium text-green-600">-{{ formatPrice(totalDiscount) }} تومان</span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">هزینه ارسال:</span>
            <span class="font-medium">{{ formatPrice(shippingCost) }} تومان</span>
          </div>
          <div class="border-t pt-2 flex justify-between text-lg font-bold">
            <span>مبلغ نهایی:</span>
            <span class="text-blue-600">{{ formatPrice(totalAmount) }} تومان</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'

// تعریف props
const props = defineProps({
  modelValue: {
    type: Array,
    default: () => []
  },
  shippingCost: {
    type: Number,
    default: 0
  }
})

// تعریف emit
const emit = defineEmits(['update:modelValue', 'update:shippingCost'])

// متغیر محلی
const orderItems = ref([...props.modelValue])

// محاسبات
const subtotal = computed(() => {
  return orderItems.value.reduce((sum, item) => sum + (item.unitPrice * item.quantity), 0)
})

const totalDiscount = computed(() => {
  return orderItems.value.reduce((sum, item) => sum + (item.discount || 0), 0)
})

const totalAmount = computed(() => {
  return subtotal.value - totalDiscount.value + props.shippingCost
})

// توابع
const calculateItemTotal = (index) => {
  const item = orderItems.value[index]
  const total = (item.unitPrice * item.quantity) - (item.discount || 0)
  item.totalPrice = Math.max(0, total)
}

const increaseQuantity = (index) => {
  orderItems.value[index].quantity++
  calculateItemTotal(index)
}

const decreaseQuantity = (index) => {
  if (orderItems.value[index].quantity > 1) {
    orderItems.value[index].quantity--
    calculateItemTotal(index)
  }
}

const removeItem = (index) => {
  orderItems.value.splice(index, 1)
}

const addNewItem = () => {
  orderItems.value.push({
    name: '',
    sku: '',
    category: '',
    image: '',
    unitPrice: 0,
    quantity: 1,
    discount: 0,
    totalPrice: 0
  })
}

// فرمت قیمت
const formatPrice = (price) => {
  if (!price) return '0'
  return new Intl.NumberFormat('fa-IR').format(price)
}

// نظارت بر تغییرات
watch(orderItems, (newValue) => {
  emit('update:modelValue', newValue)
}, { deep: true })
</script> 

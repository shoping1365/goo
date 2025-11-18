<template>
  <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
    <div class="flex items-center justify-between mb-6">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">آمار مشتریان</h3>
        <p class="text-sm text-gray-500 mt-1">Customer Payment Statistics</p>
      </div>
      <div class="flex items-center space-x-2 space-x-reverse">
        <button class="text-sm text-blue-600 hover:text-blue-800">مشاهده جزئیات</button>
      </div>
    </div>

    <!-- آمار کلی مشتریان -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
      <div class="bg-blue-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-blue-600">{{ totalCustomers }}</div>
        <div class="text-sm text-blue-500">کل مشتریان</div>
      </div>
      <div class="bg-green-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-green-600">{{ activeCustomers }}</div>
        <div class="text-sm text-green-500">مشتریان فعال</div>
      </div>
      <div class="bg-purple-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-purple-600">{{ formatPrice(averageCustomerSpending) }}</div>
        <div class="text-sm text-purple-500">میانگین خرید</div>
      </div>
      <div class="bg-orange-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-orange-600">{{ repeatCustomers }}</div>
        <div class="text-sm text-orange-500">مشتریان تکرار</div>
      </div>
    </div>

    <!-- برترین مشتریان -->
    <div class="mb-6">
      <h4 class="font-medium text-gray-900 mb-3">برترین مشتریان</h4>
      <div class="space-y-3">
        <div 
          v-for="(customer, index) in topCustomers" 
          :key="customer.id"
          class="flex items-center justify-between p-3 bg-gray-50 rounded-lg hover:bg-gray-100 transition-colors cursor-pointer"
        >
          <div class="flex items-center">
            <div class="w-8 h-8 rounded-full bg-blue-100 flex items-center justify-center ml-3 text-sm font-medium text-blue-600">
              {{ index + 1 }}
            </div>
            <div>
              <div class="font-medium text-gray-900">{{ customer.name }}</div>
              <div class="text-sm text-gray-500">{{ customer.email }}</div>
            </div>
          </div>
          <div class="text-right">
            <div class="font-semibold text-gray-900">{{ formatPrice(customer.totalSpent) }}</div>
            <div class="text-xs text-gray-500">{{ customer.orderCount }} سفارش</div>
          </div>
        </div>
      </div>
    </div>

    <!-- توزیع مبالغ خرید مشتریان -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
      <div>
        <h4 class="font-medium text-gray-900 mb-3">توزیع مبالغ خرید</h4>
        <div class="space-y-3">
          <div 
            v-for="range in spendingRanges" 
            :key="range.range"
            class="flex items-center justify-between"
          >
            <span class="text-sm text-gray-600">{{ range.range }}</span>
            <div class="flex items-center">
              <div class="w-20 bg-gray-200 rounded-full h-2 ml-2">
                <div 
                  class="bg-blue-600 h-2 rounded-full"
                  :style="{ width: range.percentage + '%' }"
                ></div>
              </div>
              <span class="text-sm font-medium text-gray-900">{{ range.count }}</span>
            </div>
          </div>
        </div>
      </div>

      <div>
        <h4 class="font-medium text-gray-900 mb-3">نوع مشتریان</h4>
        <div class="space-y-3">
          <div 
            v-for="type in customerTypes" 
            :key="type.type"
            class="flex items-center justify-between p-2 bg-gray-50 rounded"
          >
            <div class="flex items-center">
              <div 
                class="w-3 h-3 rounded-full ml-2"
                :class="getTypeColor(type.type)"
              ></div>
              <span class="text-sm font-medium">{{ type.name }}</span>
            </div>
            <div class="text-right">
              <div class="text-sm font-semibold">{{ type.count }}</div>
              <div class="text-xs text-gray-500">{{ type.percentage }}%</div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- آمار تکرار خرید -->
    <div class="mt-6 pt-6 border-t border-gray-200">
      <h4 class="font-medium text-gray-900 mb-3">آمار تکرار خرید</h4>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div class="text-center">
          <div class="text-lg font-bold text-blue-600">{{ firstTimeCustomers }}</div>
          <div class="text-sm text-gray-500">مشتریان جدید</div>
        </div>
        <div class="text-center">
          <div class="text-lg font-bold text-green-600">{{ returningCustomers }}</div>
          <div class="text-sm text-gray-500">مشتریان بازگشت</div>
        </div>
        <div class="text-center">
          <div class="text-lg font-bold text-purple-600">{{ loyalCustomers }}</div>
          <div class="text-sm text-gray-500">مشتریان وفادار</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

// داده‌های برترین مشتریان
const topCustomers = ref([
  {
    id: 1,
    name: 'علی احمدی',
    email: 'ali.ahmadi@example.com',
    totalSpent: 8500000,
    orderCount: 12
  },
  {
    id: 2,
    name: 'فاطمه محمدی',
    email: 'fateme.mohammadi@example.com',
    totalSpent: 7200000,
    orderCount: 8
  },
  {
    id: 3,
    name: 'محمد رضایی',
    email: 'mohammad.rezaei@example.com',
    totalSpent: 6500000,
    orderCount: 15
  },
  {
    id: 4,
    name: 'زهرا کریمی',
    email: 'zahra.karimi@example.com',
    totalSpent: 5800000,
    orderCount: 6
  },
  {
    id: 5,
    name: 'حسن نوری',
    email: 'hasan.nouri@example.com',
    totalSpent: 5200000,
    orderCount: 9
  }
])

// توزیع مبالغ خرید
const spendingRanges = ref([
  { range: 'کمتر از ۵۰۰ هزار تومان', count: 45, percentage: 25 },
  { range: '۵۰۰ هزار تا ۱ میلیون تومان', count: 78, percentage: 35 },
  { range: '۱ تا ۲ میلیون تومان', count: 56, percentage: 20 },
  { range: '۲ تا ۵ میلیون تومان', count: 23, percentage: 15 },
  { range: 'بیش از ۵ میلیون تومان', count: 12, percentage: 5 }
])

// نوع مشتریان
const customerTypes = ref([
  { type: 'new', name: 'مشتریان جدید', count: 89, percentage: 40 },
  { type: 'returning', name: 'مشتریان بازگشت', count: 67, percentage: 30 },
  { type: 'loyal', name: 'مشتریان وفادار', count: 45, percentage: 20 },
  { type: 'vip', name: 'مشتریان VIP', count: 23, percentage: 10 }
])

// آمار کلی
const totalCustomers = ref(224)
const activeCustomers = ref(189)
const repeatCustomers = ref(135)
const firstTimeCustomers = ref(89)
const returningCustomers = ref(67)
const loyalCustomers = ref(45)

// محاسبات
const averageCustomerSpending = computed(() => {
  const total = topCustomers.value.reduce((sum, customer) => sum + customer.totalSpent, 0)
  return total / topCustomers.value.length
})

// رنگ‌های نوع مشتری
const getTypeColor = (type: string) => {
  switch (type) {
    case 'new': return 'bg-blue-500'
    case 'returning': return 'bg-green-500'
    case 'loyal': return 'bg-purple-500'
    case 'vip': return 'bg-orange-500'
    default: return 'bg-gray-500'
  }
}

// فرمت کردن قیمت
const formatPrice = (price: number) => {
  if (price >= 1000000) {
    return (price / 1000000).toFixed(1) + 'M'
  } else if (price >= 1000) {
    return (price / 1000).toFixed(0) + 'K'
  }
  return price.toString()
}
</script> 

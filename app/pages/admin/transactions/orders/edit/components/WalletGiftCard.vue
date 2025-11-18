<template>
  <div class="space-y-6">
    <!-- اطلاعات کیف پول -->
    <div class="px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-green-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"></path>
        </svg>
        اطلاعات کیف پول
      </h3>

      <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
        <!-- موجودی کیف پول -->
        <div>
          <h4 class="text-md font-medium text-gray-700 mb-3">موجودی کیف پول</h4>
          <div class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">موجودی فعلی</label>
              <input 
                v-model="walletGiftCard.currentBalance"
                type="number" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="0"
                readonly
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">مبلغ استفاده شده در این سفارش</label>
              <input 
                v-model="walletGiftCard.usedAmount"
                type="number" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="0"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">موجودی پس از سفارش</label>
              <input 
                v-model="remainingBalance"
                type="number" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="0"
                readonly
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">وضعیت کیف پول</label>
              <select 
                v-model="walletGiftCard.walletStatus"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="active">فعال</option>
                <option value="suspended">معلق</option>
                <option value="blocked">مسدود</option>
                <option value="expired">منقضی شده</option>
              </select>
            </div>
          </div>
        </div>

        <!-- تاریخچه کیف پول -->
        <div>
          <h4 class="text-md font-medium text-gray-700 mb-3">آخرین تراکنشات کیف پول</h4>
          <div class="space-y-2 max-h-48 overflow-y-auto">
            <div 
              v-for="(transaction, index) in walletGiftCard.recentTransactions" 
              :key="index"
              class="flex items-center justify-between p-2 bg-gray-50 rounded text-sm"
            >
              <div>
                <div class="font-medium text-gray-900">{{ transaction.type }}</div>
                <div class="text-xs text-gray-500">{{ formatDate(transaction.date) }}</div>
              </div>
              <div class="text-right">
                <div class="font-medium" :class="transaction.amount > 0 ? 'text-green-600' : 'text-red-600'">
                  {{ transaction.amount > 0 ? '+' : '' }}{{ formatPrice(transaction.amount) }} تومان
                </div>
                <div class="text-xs text-gray-500">{{ transaction.description }}</div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- اطلاعات کارت‌های هدیه -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-purple-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v13m0-13V6a2 2 0 112 2h-2zm0 0V5.5A2.5 2.5 0 109.5 8H12zm-7 4h14M5 12a2 2 0 110-4h14a2 2 0 110 4M5 12v7a2 2 0 002 2h10a2 2 0 002-2v-7"></path>
        </svg>
        کارت‌های هدیه استفاده شده
      </h3>

      <div class="space-y-4">
        <!-- کارت‌های هدیه فعال -->
        <div>
          <h4 class="text-md font-medium text-gray-700 mb-3">کارت‌های هدیه در این سفارش</h4>
          <div class="space-y-3">
            <div 
              v-for="(giftCard, index) in walletGiftCard.usedGiftCards" 
              :key="index"
              class="border border-gray-200 rounded-lg px-4 py-4"
            >
              <div class="grid grid-cols-1 md:grid-cols-3 gapx-4 py-4">
                <div>
                  <label class="block text-sm font-medium text-gray-600 mb-1">کد کارت هدیه</label>
                  <input 
                    v-model="giftCard.code"
                    type="text" 
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                    placeholder="GIFT-XXXX-XXXX"
                  />
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-600 mb-1">مبلغ استفاده شده</label>
                  <input 
                    v-model="giftCard.usedAmount"
                    type="number" 
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                    placeholder="0"
                  />
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-600 mb-1">وضعیت</label>
                  <select 
                    v-model="giftCard.status"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                  >
                    <option value="active">فعال</option>
                    <option value="used">استفاده شده</option>
                    <option value="expired">منقضی شده</option>
                    <option value="invalid">نامعتبر</option>
                  </select>
                </div>
              </div>
              <div class="mt-3 grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
                <div>
                  <label class="block text-sm font-medium text-gray-600 mb-1">تاریخ انقضا</label>
                  <input 
                    v-model="giftCard.expiryDate"
                    type="date" 
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                  />
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-600 mb-1">موجودی باقی‌مانده</label>
                  <input 
                    v-model="giftCard.remainingBalance"
                    type="number" 
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                    placeholder="0"
                    readonly
                  />
                </div>
              </div>
            </div>
          </div>

          <!-- دکمه افزودن کارت هدیه -->
          <div class="mt-4">
            <button 
              @click="addGiftCard"
              class="px-4 py-2 bg-purple-600 text-white rounded-lg hover:bg-purple-700 transition-colors flex items-center text-sm"
            >
              <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
              </svg>
              افزودن کارت هدیه
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- خلاصه مالی -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-blue-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 7h6m0 10v-3m-3 3h.01M9 17h.01M9 14h.01M12 14h.01M15 11h.01M12 11h.01M9 11h.01M7 21h10a2 2 0 002-2V5a2 2 0 00-2-2H7a2 2 0 00-2 2v14a2 2 0 002 2z"></path>
        </svg>
        خلاصه مالی
      </h3>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
        <!-- محاسبات تخفیف -->
        <div class="space-y-3">
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">مبلغ کل سفارش:</span>
            <span class="font-medium">{{ formatPrice(totalOrderAmount) }} تومان</span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">تخفیف کیف پول:</span>
            <span class="font-medium text-green-600">-{{ formatPrice(walletGiftCard.usedAmount) }} تومان</span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">تخفیف کارت هدیه:</span>
            <span class="font-medium text-green-600">-{{ formatPrice(totalGiftCardDiscount) }} تومان</span>
          </div>
          <div class="border-t pt-2 flex justify-between text-lg font-bold">
            <span>مبلغ قابل پرداخت:</span>
            <span class="text-blue-600">{{ formatPrice(finalAmount) }} تومان</span>
          </div>
        </div>

        <!-- وضعیت پرداخت -->
        <div class="space-y-3">
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">وضعیت کیف پول:</span>
            <span class="font-medium" :class="getWalletStatusColor(walletGiftCard.walletStatus)">
              {{ getWalletStatusText(walletGiftCard.walletStatus) }}
            </span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">تعداد کارت‌های هدیه:</span>
            <span class="font-medium">{{ walletGiftCard.usedGiftCards.length }} کارت</span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">کل تخفیف اعمال شده:</span>
            <span class="font-medium text-green-600">{{ formatPrice(totalDiscount) }} تومان</span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">درصد تخفیف:</span>
            <span class="font-medium">{{ discountPercentage }}%</span>
          </div>
        </div>
      </div>
    </div>

    <!-- تاریخچه تراکنشات -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-orange-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
        </svg>
        تاریخچه تراکنشات
      </h3>
      
      <div class="space-y-3">
        <div 
          v-for="(history, index) in walletGiftCard.transactionHistory" 
          :key="index"
          class="flex items-center p-3 bg-gray-50 rounded-lg"
        >
          <div class="w-3 h-3 rounded-full mr-3" :class="getTransactionTypeColor(history.type)"></div>
          <div class="flex-1">
            <div class="text-sm font-medium text-gray-900">{{ getTransactionTypeText(history.type) }}</div>
            <div class="text-xs text-gray-500">{{ formatDate(history.date) }} - {{ history.user }}</div>
            <div v-if="history.note" class="text-xs text-gray-600 mt-1">{{ history.note }}</div>
          </div>
          <div class="text-right">
            <div class="text-sm font-medium" :class="history.amount > 0 ? 'text-green-600' : 'text-red-600'">
              {{ history.amount > 0 ? '+' : '' }}{{ formatPrice(history.amount) }} تومان
            </div>
            <div class="text-xs text-gray-500">{{ formatTime(history.date) }}</div>
          </div>
        </div>
      </div>

      <!-- دکمه افزودن تاریخچه -->
      <div class="mt-4">
        <button 
          @click="addTransactionHistory"
          class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors flex items-center text-sm"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
          </svg>
          افزودن تراکنش
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'

// تعریف props
const props = defineProps({
  modelValue: {
    type: Object,
    default: () => ({
      currentBalance: 500000,
      usedAmount: 0,
      walletStatus: 'active',
      recentTransactions: [],
      usedGiftCards: [],
      transactionHistory: []
    })
  }
})

// تعریف emit
const emit = defineEmits(['update:modelValue'])

// متغیر محلی
const walletGiftCard = ref({ ...props.modelValue })

// محاسبات
const remainingBalance = computed(() => {
  return walletGiftCard.value.currentBalance - walletGiftCard.value.usedAmount
})

const totalOrderAmount = computed(() => {
  return 4250000 // از سفارش
})

const totalGiftCardDiscount = computed(() => {
  return walletGiftCard.value.usedGiftCards.reduce((sum, card) => sum + (card.usedAmount || 0), 0)
})

const totalDiscount = computed(() => {
  return walletGiftCard.value.usedAmount + totalGiftCardDiscount.value
})

const finalAmount = computed(() => {
  return totalOrderAmount.value - totalDiscount.value
})

const discountPercentage = computed(() => {
  if (totalOrderAmount.value === 0) return 0
  return Math.round((totalDiscount.value / totalOrderAmount.value) * 100)
})

// توابع کمکی
const getWalletStatusText = (status) => {
  const statusMap = {
    'active': 'فعال',
    'suspended': 'معلق',
    'blocked': 'مسدود',
    'expired': 'منقضی شده'
  }
  return statusMap[status] || status
}

const getWalletStatusColor = (status) => {
  const colors = {
    'active': 'text-green-600',
    'suspended': 'text-yellow-600',
    'blocked': 'text-red-600',
    'expired': 'text-gray-600'
  }
  return colors[status] || 'text-gray-600'
}

const getTransactionTypeText = (type) => {
  const typeMap = {
    'wallet_use': 'استفاده از کیف پول',
    'gift_card_use': 'استفاده از کارت هدیه',
    'wallet_recharge': 'شارژ کیف پول',
    'gift_card_purchase': 'خرید کارت هدیه',
    'refund': 'بازپرداخت'
  }
  return typeMap[type] || type
}

const getTransactionTypeColor = (type) => {
  const colors = {
    'wallet_use': 'bg-red-400',
    'gift_card_use': 'bg-purple-400',
    'wallet_recharge': 'bg-green-400',
    'gift_card_purchase': 'bg-blue-400',
    'refund': 'bg-orange-400'
  }
  return colors[type] || 'bg-gray-400'
}

const formatPrice = (price) => {
  if (!price) return '0'
  return new Intl.NumberFormat('fa-IR').format(price)
}

const formatDate = (date) => {
  if (!date) return 'نامشخص'
  return new Date(date).toLocaleDateString('fa-IR')
}

const formatTime = (date) => {
  if (!date) return ''
  return new Date(date).toLocaleTimeString('fa-IR', { 
    hour: '2-digit', 
    minute: '2-digit' 
  })
}

const addGiftCard = () => {
  walletGiftCard.value.usedGiftCards.push({
    code: '',
    usedAmount: 0,
    status: 'active',
    expiryDate: '',
    remainingBalance: 0
  })
}

const addTransactionHistory = () => {
  const newEntry = {
    type: 'wallet_use',
    date: new Date().toISOString(),
    user: 'مدیر سیستم',
    amount: -walletGiftCard.value.usedAmount,
    note: 'استفاده از کیف پول در سفارش'
  }
  walletGiftCard.value.transactionHistory.unshift(newEntry)
}

// نظارت بر تغییرات
watch(walletGiftCard, (newValue) => {
  emit('update:modelValue', newValue)
}, { deep: true })
</script> 

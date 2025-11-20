<template>
  <div class="bg-white rounded-lg shadow p-6">
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-xl font-semibold text-gray-900">مدیریت پاداش‌ها</h2>
      <button 
        class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors"
        @click="showRewardForm = true; editingReward = null"
      >
        افزودن پاداش جدید
      </button>
    </div>

    <!-- لیست پاداش‌ها -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div 
        v-for="reward in rewards" 
        :key="reward.id"
        class="border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow"
      >
        <div class="flex justify-between items-start mb-3">
          <div class="flex items-center">
            <div :class="getRewardIconColor(reward.type)" class="p-2 rounded-lg">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="getRewardIcon(reward.type)"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">{{ reward.name }}</h3>
          </div>
          <span :class="reward.active ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'" class="px-2 py-1 text-xs font-semibold rounded-full">
            {{ reward.active ? 'فعال' : 'غیرفعال' }}
          </span>
        </div>
        
        <p class="text-gray-600 text-sm mb-3">{{ reward.description }}</p>
        
        <div class="space-y-2 text-sm">
          <div class="flex justify-between">
            <span class="text-gray-500">نوع:</span>
            <span class="font-semibold text-gray-900">{{ getRewardTypeName(reward.type) }}</span>
          </div>
          <div class="flex justify-between">
            <span class="text-gray-500">امتیاز مورد نیاز:</span>
            <span class="font-semibold text-gray-900">{{ reward.pointsRequired }} امتیاز</span>
          </div>
          <div class="flex justify-between">
            <span class="text-gray-500">ارزش:</span>
            <span class="font-semibold text-gray-900">{{ formatRewardValue(reward) }}</span>
          </div>
          <div class="flex justify-between">
            <span class="text-gray-500">تعداد استفاده:</span>
            <span class="font-semibold text-gray-900">{{ reward.usageCount }}</span>
          </div>
        </div>
        
        <div class="flex justify-end space-x-2 space-x-reverse mt-4">
          <button 
            class="text-blue-600 hover:text-blue-800 text-sm"
            @click="editingReward = reward; showRewardForm = true"
          >
            ویرایش
          </button>
          <button 
            :class="reward.active ? 'text-red-600 hover:text-red-800' : 'text-green-600 hover:text-green-800'"
            class="text-sm"
            @click="toggleRewardStatus(reward)"
          >
            {{ reward.active ? 'غیرفعال' : 'فعال' }}
          </button>
        </div>
      </div>
    </div>

    <!-- مودال فرم پاداش -->
    <div v-if="showRewardForm" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
      <div class="relative top-20 mx-auto p-5 border w-full max-w-md sm:max-w-lg md:max-w-xl shadow-lg rounded-md bg-white">
        <div class="mt-3">
          <h3 class="text-lg font-medium text-gray-900 mb-4">
            {{ editingReward ? 'ویرایش پاداش' : 'افزودن پاداش جدید' }}
          </h3>
          
          <form @submit.prevent="saveReward">
            <div class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-gray-700">نوع پاداش</label>
                <select 
                  v-model="rewardForm.type"
                  class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                  required
                >
                  <option value="discount_percent">تخفیف درصدی</option>
                  <option value="discount_amount">تخفیف مبلغی</option>
                  <option value="free_shipping">ارسال رایگان</option>
                  <option value="gift">هدیه</option>
                  <option value="cashback">بازگشت وجه</option>
                </select>
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700">نام پاداش</label>
                <input 
                  v-model="rewardForm.name" 
                  type="text" 
                  class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                  required
                >
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700">توضیحات</label>
                <textarea 
                  v-model="rewardForm.description" 
                  rows="2"
                  class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                ></textarea>
              </div>
              
              <div class="grid grid-cols-2 gap-6">
                <div>
                  <label class="block text-sm font-medium text-gray-700">امتیاز مورد نیاز</label>
                  <input 
                    v-model="rewardForm.pointsRequired" 
                    type="number" 
                    class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                    required
                  >
                </div>
                
                <div>
                  <label class="block text-sm font-medium text-gray-700">ارزش</label>
                  <input 
                    v-model="rewardForm.value" 
                    type="number" 
                    class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                    required
                  >
                </div>
              </div>
              
              <div class="grid grid-cols-2 gap-6">
                <div>
                  <label class="block text-sm font-medium text-gray-700">حداقل خرید</label>
                  <input 
                    v-model="rewardForm.minPurchase" 
                    type="number" 
                    class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                  >
                </div>
                
                <div>
                  <label class="block text-sm font-medium text-gray-700">حداکثر استفاده</label>
                  <input 
                    v-model="rewardForm.maxUsage" 
                    type="number" 
                    class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                  >
                </div>
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700">تاریخ انقضا (روز)</label>
                <input 
                  v-model="rewardForm.expiryDays" 
                  type="number" 
                  class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                >
              </div>
            </div>
            
            <div class="flex justify-end space-x-3 space-x-reverse mt-6">
              <button 
                type="button"
                class="bg-gray-300 text-gray-700 px-4 py-2 rounded-md hover:bg-gray-400 transition-colors"
                @click="showRewardForm = false"
              >
                انصراف
              </button>
              <button 
                type="submit"
                class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition-colors"
              >
                {{ editingReward ? 'ویرایش' : 'افزودن' }}
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'

// پاداش‌ها
const rewards = ref([
  {
    id: 1,
    type: 'discount_percent',
    name: 'تخفیف 10 درصدی',
    description: 'تخفیف 10 درصدی روی کل خرید',
    pointsRequired: 1000,
    value: 10,
    minPurchase: 100000,
    maxUsage: 1,
    expiryDays: 30,
    usageCount: 45,
    active: true
  },
  {
    id: 2,
    type: 'discount_amount',
    name: 'تخفیف 50 هزار تومانی',
    description: 'تخفیف 50 هزار تومانی روی خرید',
    pointsRequired: 500,
    value: 50000,
    minPurchase: 200000,
    maxUsage: 2,
    expiryDays: 60,
    usageCount: 23,
    active: true
  },
  {
    id: 3,
    type: 'free_shipping',
    name: 'ارسال رایگان',
    description: 'ارسال رایگان برای سفارش‌های بالای 200 هزار تومان',
    pointsRequired: 300,
    value: 0,
    minPurchase: 200000,
    maxUsage: 3,
    expiryDays: 90,
    usageCount: 67,
    active: true
  },
  {
    id: 4,
    type: 'gift',
    name: 'هدیه ویژه',
    description: 'دریافت یک هدیه رایگان',
    pointsRequired: 800,
    value: 0,
    minPurchase: 0,
    maxUsage: 1,
    expiryDays: 45,
    usageCount: 12,
    active: true
  },
  {
    id: 5,
    type: 'cashback',
    name: 'بازگشت وجه',
    description: 'بازگشت 5 درصد وجه خرید',
    pointsRequired: 1500,
    value: 5,
    minPurchase: 500000,
    maxUsage: 1,
    expiryDays: 120,
    usageCount: 8,
    active: true
  }
])

// State برای مودال
const showRewardForm = ref(false)
const editingReward = ref(null)

// فرم پاداش
const rewardForm = reactive({
  type: 'discount_percent',
  name: '',
  description: '',
  pointsRequired: 0,
  value: 0,
  minPurchase: 0,
  maxUsage: 1,
  expiryDays: 30
})

// رنگ آیکون‌ها
function getRewardIconColor(type: string) {
  const colors = {
    discount_percent: 'bg-green-500',
    discount_amount: 'bg-blue-500',
    free_shipping: 'bg-purple-500',
    gift: 'bg-pink-500',
    cashback: 'bg-yellow-500'
  }
  return colors[type] || 'bg-gray-500'
}

// آیکون‌ها
function getRewardIcon(type: string) {
  const icons = {
    discount_percent: 'M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1',
    discount_amount: 'M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1',
    free_shipping: 'M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4',
    gift: 'M12 8v13m0-13V6a2 2 0 112 2h-2zm0 0V5.5A2.5 2.5 0 109.5 8H12zm-7 4h14M5 12a2 2 0 110-4h14a2 2 0 110 4M5 12v7a2 2 0 002 2h10a2 2 0 002-2v-7',
    cashback: 'M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z'
  }
  return icons[type] || 'M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z'
}

// نام نوع پاداش
function getRewardTypeName(type: string) {
  const names = {
    discount_percent: 'تخفیف درصدی',
    discount_amount: 'تخفیف مبلغی',
    free_shipping: 'ارسال رایگان',
    gift: 'هدیه',
    cashback: 'بازگشت وجه'
  }
  return names[type] || 'نامشخص'
}

// فرمت ارزش پاداش
function formatRewardValue(reward: any) {
  if (reward.type === 'discount_percent') {
    return `${reward.value}%`
  } else if (reward.type === 'discount_amount') {
    return `${reward.value.toLocaleString()} تومان`
  } else if (reward.type === 'free_shipping') {
    return 'رایگان'
  } else if (reward.type === 'gift') {
    return 'هدیه'
  } else if (reward.type === 'cashback') {
    return `${reward.value}%`
  }
  return 'نامشخص'
}

// تغییر وضعیت پاداش
function toggleRewardStatus(reward: any) {
  reward.active = !reward.active
  // TODO: فراخوانی API برای تغییر وضعیت
}

// ذخیره پاداش
function saveReward() {
  if (editingReward.value) {
    // ویرایش پاداش موجود
    const index = rewards.value.findIndex(r => r.id === editingReward.value.id)
    if (index !== -1) {
      rewards.value[index] = { ...editingReward.value, ...rewardForm }
    }
  } else {
    // افزودن پاداش جدید
    const newReward = {
      id: Date.now(),
      ...rewardForm,
      usageCount: 0,
      active: true
    }
    rewards.value.push(newReward)
  }
  
  // پاک کردن فرم
  Object.assign(rewardForm, {
    type: 'discount_percent',
    name: '',
    description: '',
    pointsRequired: 0,
    value: 0,
    minPurchase: 0,
    maxUsage: 1,
    expiryDays: 30
  })
  
  showRewardForm.value = false
  editingReward.value = null
}
</script> 

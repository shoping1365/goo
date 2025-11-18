<template>
  <div class="bg-white rounded-lg shadow p-6">
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-xl font-semibold text-gray-900">قوانین امتیازدهی</h2>
      <button 
        @click="showRuleForm = true; editingRule = null"
        class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors"
      >
        افزودن قانون جدید
      </button>
    </div>

    <!-- لیست قوانین -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div 
        v-for="rule in pointRules" 
        :key="rule.id"
        class="border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow"
      >
        <div class="flex justify-between items-start mb-3">
          <div class="flex items-center">
            <div :class="getRuleIconColor(rule.type)" class="p-2 rounded-lg">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="getRuleIcon(rule.type)"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">{{ rule.name }}</h3>
          </div>
          <span :class="rule.active ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'" class="px-2 py-1 text-xs font-semibold rounded-full">
            {{ rule.active ? 'فعال' : 'غیرفعال' }}
          </span>
        </div>
        
        <p class="text-gray-600 text-sm mb-3">{{ rule.description }}</p>
        
        <div class="space-y-2 text-sm">
          <div class="flex justify-between">
            <span class="text-gray-500">امتیاز:</span>
            <span class="font-semibold text-gray-900">{{ rule.points }} امتیاز</span>
          </div>
          <div class="flex justify-between">
            <span class="text-gray-500">حداکثر در روز:</span>
            <span class="font-semibold text-gray-900">{{ rule.maxPerDay }}</span>
          </div>
          <div class="flex justify-between">
            <span class="text-gray-500">حداکثر در ماه:</span>
            <span class="font-semibold text-gray-900">{{ rule.maxPerMonth }}</span>
          </div>
        </div>
        
        <div class="flex justify-end space-x-2 space-x-reverse mt-4">
          <button 
            @click="editingRule = rule; showRuleForm = true"
            class="text-blue-600 hover:text-blue-800 text-sm"
          >
            ویرایش
          </button>
          <button 
            @click="toggleRuleStatus(rule)"
            :class="rule.active ? 'text-red-600 hover:text-red-800' : 'text-green-600 hover:text-green-800'"
            class="text-sm"
          >
            {{ rule.active ? 'غیرفعال' : 'فعال' }}
          </button>
        </div>
      </div>
    </div>

    <!-- مودال فرم قانون -->
    <div v-if="showRuleForm" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
      <div class="relative top-20 mx-auto p-5 border w-full max-w-md sm:max-w-lg md:max-w-xl shadow-lg rounded-md bg-white">
        <div class="mt-3">
          <h3 class="text-lg font-medium text-gray-900 mb-4">
            {{ editingRule ? 'ویرایش قانون' : 'افزودن قانون جدید' }}
          </h3>
          
          <form @submit.prevent="saveRule">
            <div class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-gray-700">نوع قانون</label>
                <select 
                  v-model="ruleForm.type"
                  class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                  required
                >
                  <option value="purchase">خرید</option>
                  <option value="registration">ثبت‌نام</option>
                  <option value="referral">معرفی دوستان</option>
                  <option value="review">بررسی محصول</option>
                  <option value="profile">تکمیل پروفایل</option>
                  <option value="birthday">تولد</option>
                  <option value="special_day">روزهای خاص</option>
                </select>
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700">نام قانون</label>
                <input 
                  v-model="ruleForm.name" 
                  type="text" 
                  class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                  required
                >
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700">توضیحات</label>
                <textarea 
                  v-model="ruleForm.description" 
                  rows="2"
                  class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                ></textarea>
              </div>
              
              <div class="grid grid-cols-2 gap-6">
                <div>
                  <label class="block text-sm font-medium text-gray-700">امتیاز</label>
                  <input 
                    v-model="ruleForm.points" 
                    type="number" 
                    class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                    required
                  >
                </div>
                
                <div>
                  <label class="block text-sm font-medium text-gray-700">ضریب</label>
                  <input 
                    v-model="ruleForm.multiplier" 
                    type="number" 
                    step="0.1"
                    class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                  >
                </div>
              </div>
              
              <div class="grid grid-cols-2 gap-6">
                <div>
                  <label class="block text-sm font-medium text-gray-700">حداکثر در روز</label>
                  <input 
                    v-model="ruleForm.maxPerDay" 
                    type="number" 
                    class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                  >
                </div>
                
                <div>
                  <label class="block text-sm font-medium text-gray-700">حداکثر در ماه</label>
                  <input 
                    v-model="ruleForm.maxPerMonth" 
                    type="number" 
                    class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                  >
                </div>
              </div>
            </div>
            
            <div class="flex justify-end space-x-3 space-x-reverse mt-6">
              <button 
                type="button"
                @click="showRuleForm = false"
                class="bg-gray-300 text-gray-700 px-4 py-2 rounded-md hover:bg-gray-400 transition-colors"
              >
                انصراف
              </button>
              <button 
                type="submit"
                class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition-colors"
              >
                {{ editingRule ? 'ویرایش' : 'افزودن' }}
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

// قوانین امتیازدهی
const pointRules = ref([
  {
    id: 1,
    type: 'purchase',
    name: 'امتیاز خرید',
    description: 'امتیاز برای هر تومان خرید',
    points: 1,
    multiplier: 1,
    maxPerDay: 1000,
    maxPerMonth: 10000,
    active: true
  },
  {
    id: 2,
    type: 'registration',
    name: 'امتیاز ثبت‌نام',
    description: 'امتیاز برای ثبت‌نام اولیه',
    points: 100,
    multiplier: 1,
    maxPerDay: 1,
    maxPerMonth: 1,
    active: true
  },
  {
    id: 3,
    type: 'referral',
    name: 'امتیاز معرفی',
    description: 'امتیاز برای معرفی دوستان',
    points: 50,
    multiplier: 1,
    maxPerDay: 10,
    maxPerMonth: 100,
    active: true
  },
  {
    id: 4,
    type: 'review',
    name: 'امتیاز بررسی',
    description: 'امتیاز برای بررسی محصولات',
    points: 20,
    multiplier: 1,
    maxPerDay: 5,
    maxPerMonth: 50,
    active: true
  },
  {
    id: 5,
    type: 'profile',
    name: 'امتیاز پروفایل',
    description: 'امتیاز برای تکمیل پروفایل',
    points: 30,
    multiplier: 1,
    maxPerDay: 1,
    maxPerMonth: 1,
    active: true
  },
  {
    id: 6,
    type: 'birthday',
    name: 'امتیاز تولد',
    description: 'امتیاز برای روز تولد',
    points: 200,
    multiplier: 1,
    maxPerDay: 1,
    maxPerMonth: 1,
    active: true
  }
])

// State برای مودال
const showRuleForm = ref(false)
const editingRule = ref(null)

// فرم قانون
const ruleForm = reactive({
  type: 'purchase',
  name: '',
  description: '',
  points: 0,
  multiplier: 1,
  maxPerDay: 0,
  maxPerMonth: 0
})

// رنگ آیکون‌ها
function getRuleIconColor(type: string) {
  const colors = {
    purchase: 'bg-green-500',
    registration: 'bg-blue-500',
    referral: 'bg-purple-500',
    review: 'bg-yellow-500',
    profile: 'bg-indigo-500',
    birthday: 'bg-pink-500',
    special_day: 'bg-red-500'
  }
  return colors[type] || 'bg-gray-500'
}

// آیکون‌ها
function getRuleIcon(type: string) {
  const icons = {
    purchase: 'M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1',
    registration: 'M18 9v3m0 0v3m0-3h3m-3 0h-3m-2-5a4 4 0 11-8 0 4 4 0 018 0zM3 20a6 6 0 0112 0v1H3v-1z',
    referral: 'M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z',
    review: 'M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z',
    profile: 'M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z',
    birthday: 'M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z',
    special_day: 'M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z'
  }
  return icons[type] || 'M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z'
}

// تغییر وضعیت قانون
function toggleRuleStatus(rule: any) {
  rule.active = !rule.active
  // TODO: فراخوانی API برای تغییر وضعیت
}

// ذخیره قانون
function saveRule() {
  if (editingRule.value) {
    // ویرایش قانون موجود
    const index = pointRules.value.findIndex(r => r.id === editingRule.value.id)
    if (index !== -1) {
      pointRules.value[index] = { ...editingRule.value, ...ruleForm }
    }
  } else {
    // افزودن قانون جدید
    const newRule = {
      id: Date.now(),
      ...ruleForm,
      active: true
    }
    pointRules.value.push(newRule)
  }
  
  // پاک کردن فرم
  Object.assign(ruleForm, {
    type: 'purchase',
    name: '',
    description: '',
    points: 0,
    multiplier: 1,
    maxPerDay: 0,
    maxPerMonth: 0
  })
  
  showRuleForm.value = false
  editingRule.value = null
}
</script> 

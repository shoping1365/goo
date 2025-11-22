<template>
  <div v-if="hasAccess" class="bg-white rounded-lg shadow p-6">
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-xl font-semibold text-gray-900">قوانین پاداش‌دهی</h2>
      <button 
        class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors"
        @click="showRuleForm = true; editingRule = null"
      >
        افزودن قانون جدید
      </button>
    </div>

    <!-- لیست قوانین -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <div v-for="rule in rewardRules" :key="rule.id" class="border border-gray-200 rounded-lg p-6">
        <div class="flex justify-between items-start mb-4">
          <div>
            <h3 class="text-lg font-semibold text-gray-900">{{ rule.name }}</h3>
            <p class="text-sm text-gray-600 mt-1">{{ rule.description }}</p>
          </div>
          <span :class="rule.active ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'" class="px-2 py-1 text-xs font-semibold rounded-full">
            {{ rule.active ? 'فعال' : 'غیرفعال' }}
          </span>
        </div>
        
        <div class="space-y-3">
          <div class="flex justify-between">
            <span class="text-sm text-gray-600">نوع پاداش:</span>
            <span class="text-sm font-medium">{{ getRewardTypeLabel(rule.rewardType) }}</span>
          </div>
          
          <div class="flex justify-between">
            <span class="text-sm text-gray-600">مقدار پاداش:</span>
            <span class="text-sm font-medium">{{ formatReward(rule.rewardValue, rule.rewardType) }}</span>
          </div>
          
          <div class="flex justify-between">
            <span class="text-sm text-gray-600">حداقل خرید:</span>
            <span class="text-sm font-medium">{{ rule.minPurchaseAmount.toLocaleString() }} تومان</span>
          </div>
          
          <div class="flex justify-between">
            <span class="text-sm text-gray-600">حداکثر پاداش:</span>
            <span class="text-sm font-medium">{{ rule.maxRewardAmount.toLocaleString() }} تومان</span>
          </div>
          
          <div class="flex justify-between">
            <span class="text-sm text-gray-600">تاریخ اعتبار:</span>
            <span class="text-sm font-medium">{{ formatDate(rule.validUntil) }}</span>
          </div>
        </div>
        
        <div class="flex justify-end space-x-2 space-x-reverse mt-4">
          <button 
            class="text-blue-600 hover:text-blue-900 text-sm"
            @click="editingRule = rule; showRuleForm = true"
          >
            ویرایش
          </button>
          <button 
            :class="rule.active ? 'text-red-600 hover:text-red-900' : 'text-green-600 hover:text-green-900'"
            class="text-sm"
            @click="toggleRuleStatus(rule)"
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
            {{ editingRule ? 'ویرایش قانون پاداش' : 'افزودن قانون جدید' }}
          </h3>
          
          <form @submit.prevent="saveRule">
            <div class="space-y-4">
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
              
              <div>
                <label class="block text-sm font-medium text-gray-700">نوع پاداش</label>
                <select 
                  v-model="ruleForm.rewardType"
                  class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                  required
                >
                  <option value="percentage">درصدی</option>
                  <option value="fixed">مبلغ ثابت</option>
                  <option value="points">امتیاز</option>
                </select>
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700">مقدار پاداش</label>
                <input 
                  v-model="ruleForm.rewardValue" 
                  type="number" 
                  class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                  required
                >
              </div>
              
              <div class="grid grid-cols-2 gap-6">
                <div>
                  <label class="block text-sm font-medium text-gray-700">حداقل خرید (تومان)</label>
                  <input 
                    v-model="ruleForm.minPurchaseAmount" 
                    type="number" 
                    class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                  >
                </div>
                
                <div>
                  <label class="block text-sm font-medium text-gray-700">حداکثر پاداش (تومان)</label>
                  <input 
                    v-model="ruleForm.maxRewardAmount" 
                    type="number" 
                    class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                  >
                </div>
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700">تاریخ انقضا</label>
                <input 
                  v-model="ruleForm.validUntil" 
                  type="date" 
                  class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                >
              </div>
            </div>
            
            <div class="flex justify-end space-x-3 space-x-reverse mt-6">
              <button 
                type="button"
                class="bg-gray-300 text-gray-700 px-4 py-2 rounded-md hover:bg-gray-400 transition-colors"
                @click="showRuleForm = false"
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
import { computed, onMounted, reactive, ref, watch } from 'vue';
import { useAuth } from '~/composables/useAuth';

declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>;

// احراز هویت
const { user, isAuthenticated } = useAuth();

// بررسی دسترسی admin
const hasAccess = computed(() => {
  if (!isAuthenticated.value) {
    return false;
  }

  const userRole = user.value?.role?.toLowerCase() || '';
  const adminRoles = ['admin', 'developer', 'super_admin', 'manager', 'operator'];
  return adminRoles.includes(userRole);
});

// بررسی احراز هویت و دسترسی admin - نمایش 404 در صورت عدم دسترسی
const checkAuth = async (): Promise<void> => {
  if (!hasAccess.value) {
    await navigateTo('/404', { external: false });
  }
};

// بررسی احراز هویت در هنگام mount
onMounted(async () => {
  await checkAuth();
});

// بررسی احراز هویت هنگام تغییر وضعیت احراز هویت
watch([isAuthenticated, hasAccess], async () => {
  if (!hasAccess.value) {
    await checkAuth();
  }
});

// قوانین پاداش
const rewardRules = ref([
  {
    id: 1,
    name: 'پاداش ارجاع اول',
    description: 'پاداش برای اولین ارجاع موفق هر کاربر',
    rewardType: 'percentage',
    rewardValue: 10,
    minPurchaseAmount: 100000,
    maxRewardAmount: 50000,
    validUntil: '2024-12-31',
    active: true
  },
  {
    id: 2,
    name: 'پاداش ارجاع ویژه',
    description: 'پاداش ویژه برای ارجاعات با خرید بالا',
    rewardType: 'fixed',
    rewardValue: 25000,
    minPurchaseAmount: 500000,
    maxRewardAmount: 100000,
    validUntil: '2024-12-31',
    active: true
  },
  {
    id: 3,
    name: 'پاداش امتیازی',
    description: 'پاداش امتیازی برای ارجاعات موفق',
    rewardType: 'points',
    rewardValue: 1000,
    minPurchaseAmount: 200000,
    maxRewardAmount: 5000,
    validUntil: '2024-12-31',
    active: true
  }
])

// State برای مودال
const showRuleForm = ref(false)
const editingRule = ref(null)

// فرم قانون
const ruleForm = reactive({
  name: '',
  description: '',
  rewardType: 'percentage',
  rewardValue: 0,
  minPurchaseAmount: 0,
  maxRewardAmount: 0,
  validUntil: ''
})

// برچسب نوع پاداش
function getRewardTypeLabel(type: string) {
  const labels = {
    percentage: 'درصدی',
    fixed: 'مبلغ ثابت',
    points: 'امتیاز'
  }
  return labels[type] || type
}

// فرمت پاداش
function formatReward(value: number, type: string) {
  if (type === 'percentage') {
    return `${value}%`
  } else if (type === 'fixed') {
    return `${value.toLocaleString()} تومان`
  } else {
    return `${value} امتیاز`
  }
}

// فرمت تاریخ
function formatDate(dateString: string) {
  const date = new Date(dateString)
  return date.toLocaleDateString('fa-IR')
}

interface RewardRule {
  id: number
  name: string
  description: string
  rewardType: string
  rewardValue: number
  minPurchaseAmount: number
  maxRewardAmount: number
  validUntil: string
  active: boolean
}

// تغییر وضعیت قانون
function toggleRuleStatus(rule: RewardRule) {
  rule.active = !rule.active
  // TODO: فراخوانی API برای تغییر وضعیت
}

// ذخیره قانون
function saveRule() {
  if (editingRule.value) {
    // ویرایش قانون موجود
    const index = rewardRules.value.findIndex(r => r.id === editingRule.value.id)
    if (index !== -1) {
      rewardRules.value[index] = { ...editingRule.value, ...ruleForm }
    }
  } else {
    // افزودن قانون جدید
    const newRule = {
      id: Date.now(),
      ...ruleForm,
      active: true
    }
    rewardRules.value.push(newRule)
  }
  
  // پاک کردن فرم
  Object.assign(ruleForm, {
    name: '',
    description: '',
    rewardType: 'percentage',
    rewardValue: 0,
    minPurchaseAmount: 0,
    maxRewardAmount: 0,
    validUntil: ''
  })
  
  showRuleForm.value = false
  editingRule.value = null
}
</script> 

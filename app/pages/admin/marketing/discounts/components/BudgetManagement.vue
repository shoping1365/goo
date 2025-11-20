<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200">
    <!-- هدر بخش -->
    <div class="p-6 border-b border-gray-200">
      <div class="flex items-center justify-between">
        <div>
          <h2 class="text-lg font-semibold text-gray-900">مدیریت بودجه و محدودیت‌های مالی</h2>
          <p class="text-gray-600 mt-1">تعیین و کنترل بودجه کمپین‌ها و کوپن‌ها</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors" @click="showBudgetForm = true">
            <span class="i-heroicons-plus ml-2"></span>
            افزودن بودجه جدید
          </button>
        </div>
      </div>
    </div>

    <!-- آمار بودجه -->
    <div class="p-6 border-b border-gray-200">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div class="bg-blue-50 p-6 rounded-lg">
          <div class="flex items-center">
            <div class="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center ml-3">
              <span class="i-heroicons-banknotes text-blue-600 text-xl"></span>
            </div>
            <div>
              <p class="text-sm text-blue-600">بودجه کل</p>
              <p class="text-2xl font-bold text-blue-900">{{ formatPrice(stats.totalBudget) }}</p>
            </div>
          </div>
        </div>
        <div class="bg-green-50 p-6 rounded-lg">
          <div class="flex items-center">
            <div class="w-12 h-12 bg-green-100 rounded-lg flex items-center justify-center ml-3">
              <span class="i-heroicons-currency-dollar text-green-600 text-xl"></span>
            </div>
            <div>
              <p class="text-sm text-green-600">مصرف شده</p>
              <p class="text-2xl font-bold text-green-900">{{ formatPrice(stats.spentBudget) }}</p>
            </div>
          </div>
        </div>
        <div class="bg-orange-50 p-6 rounded-lg">
          <div class="flex items-center">
            <div class="w-12 h-12 bg-orange-100 rounded-lg flex items-center justify-center ml-3">
              <span class="i-heroicons-exclamation-triangle text-orange-600 text-xl"></span>
            </div>
            <div>
              <p class="text-sm text-orange-600">نزدیک به سقف</p>
              <p class="text-2xl font-bold text-orange-900">{{ stats.nearLimitCount }}</p>
            </div>
          </div>
        </div>
        <div class="bg-purple-50 p-6 rounded-lg">
          <div class="flex items-center">
            <div class="w-12 h-12 bg-purple-100 rounded-lg flex items-center justify-center ml-3">
              <span class="i-heroicons-clipboard-document-list text-purple-600 text-xl"></span>
            </div>
            <div>
              <p class="text-sm text-purple-600">کمپین‌های فعال</p>
              <p class="text-2xl font-bold text-purple-900">{{ stats.activeCampaigns }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- لیست بودجه کمپین‌ها -->
    <div class="overflow-x-auto">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نام کمپین</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">بودجه کل</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مصرف شده</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">سقف روزانه</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">سقف ماهانه</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">درصد مصرف</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-for="budget in budgets" :key="budget.id" :class="budget.percentUsed >= 90 ? 'bg-orange-50' : ''">
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="text-sm font-medium text-gray-900">{{ budget.campaignName }}</div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ formatPrice(budget.totalBudget) }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ formatPrice(budget.spentBudget) }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ formatPrice(budget.dailyLimit) }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ formatPrice(budget.monthlyLimit) }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
              <div class="flex items-center">
                <div class="w-24 bg-gray-200 rounded-full h-2 ml-2">
                  <div class="h-2 rounded-full" :class="budget.percentUsed >= 90 ? 'bg-orange-500' : 'bg-blue-600'" :style="{ width: budget.percentUsed + '%' }"></div>
                </div>
                <span>{{ budget.percentUsed }}%</span>
              </div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <span :class="budget.percentUsed >= 90 ? 'inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-orange-100 text-orange-700' : 'inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-700'">
                {{ budget.percentUsed >= 90 ? 'نزدیک به سقف' : 'عادی' }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
              <div class="flex items-center space-x-2 space-x-reverse">
                <button class="text-blue-600 hover:text-blue-900" @click="editBudget(budget)">
                  <span class="i-heroicons-pencil-square"></span>
                </button>
                <button class="text-red-600 hover:text-red-900" @click="deleteBudget(budget)">
                  <span class="i-heroicons-trash"></span>
                </button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- پیام عدم وجود بودجه -->
    <div v-if="budgets.length === 0" class="text-center py-12">
      <div class="w-16 h-16 bg-gray-100 rounded-full flex items-center justify-center mx-auto mb-4">
        <span class="i-heroicons-banknotes text-gray-400 text-xl"></span>
      </div>
      <h3 class="text-lg font-medium text-gray-900 mb-2">هیچ بودجه‌ای ثبت نشده</h3>
      <p class="text-gray-600">برای کمپین‌های خود بودجه تعیین کنید.</p>
    </div>

    <!-- مودال فرم بودجه -->
    <div v-if="showBudgetForm" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg shadow-xl max-w-md w-full mx-4">
        <div class="p-6 border-b border-gray-200">
          <div class="flex items-center justify-between">
            <h3 class="text-lg font-semibold text-gray-900">{{ editingBudget ? 'ویرایش بودجه' : 'افزودن بودجه جدید' }}</h3>
            <button class="text-gray-400 hover:text-gray-600" @click="closeForm">
              <span class="i-heroicons-x-mark text-xl"></span>
            </button>
          </div>
        </div>
        <form class="p-6 space-y-6" @submit.prevent="handleSubmit">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نام کمپین *</label>
            <input v-model="form.campaignName" type="text" required class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="نام کمپین">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">بودجه کل *</label>
            <input v-model.number="form.totalBudget" type="number" min="10000" required class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="مثال: 1000000">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">سقف روزانه</label>
            <input v-model.number="form.dailyLimit" type="number" min="0" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="مثال: 50000">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">سقف ماهانه</label>
            <input v-model.number="form.monthlyLimit" type="number" min="0" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="مثال: 200000">
          </div>
        </form>
        <div class="p-6 border-t border-gray-200 flex justify-end space-x-3 space-x-reverse">
          <button class="px-6 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200 transition-colors" @click="closeForm">
            انصراف
          </button>
          <button class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors" @click="handleSubmit">
            ذخیره
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed } from 'vue'

const showBudgetForm = ref(false)
const editingBudget = ref<any>(null)

const stats = ref({
  totalBudget: 50000000,
  spentBudget: 31200000,
  nearLimitCount: 2,
  activeCampaigns: 5
})

const budgets = ref([
  {
    id: 1,
    campaignName: 'کمپین نوروزی',
    totalBudget: 10000000,
    spentBudget: 9500000,
    dailyLimit: 500000,
    monthlyLimit: 3000000,
    percentUsed: 95
  },
  {
    id: 2,
    campaignName: 'کمپین تابستانه',
    totalBudget: 20000000,
    spentBudget: 12000000,
    dailyLimit: 800000,
    monthlyLimit: 5000000,
    percentUsed: 60
  },
  {
    id: 3,
    campaignName: 'کوپن ویژه مشتریان وفادار',
    totalBudget: 5000000,
    spentBudget: 4000000,
    dailyLimit: 200000,
    monthlyLimit: 1000000,
    percentUsed: 80
  }
])

const form = reactive({
  campaignName: '',
  totalBudget: 0,
  dailyLimit: 0,
  monthlyLimit: 0
})

const formatPrice = (price: number): string => {
  return new Intl.NumberFormat('fa-IR').format(price) + ' تومان'
}

const editBudget = (budget: any) => {
  editingBudget.value = budget
  Object.assign(form, budget)
  showBudgetForm.value = true
}

const deleteBudget = async (budget: any) => {
  if (confirm(`آیا از حذف بودجه کمپین "${budget.campaignName}" اطمینان دارید؟`)) {
    try {
      const index = budgets.value.findIndex(b => b.id === budget.id)
      if (index !== -1) {
        budgets.value.splice(index, 1)
      }
    } catch (error) {
      console.error('خطا در حذف بودجه:', error)
    }
  }
}

const handleSubmit = async () => {
  if (!form.campaignName || !form.totalBudget) {
    alert('لطفاً فیلدهای اجباری را پر کنید')
    return
  }
  if (editingBudget.value) {
    Object.assign(editingBudget.value, form)
  } else {
    budgets.value.unshift({
      id: Date.now(),
      ...form,
      spentBudget: 0,
      percentUsed: 0
    })
  }
  closeForm()
}

const closeForm = () => {
  showBudgetForm.value = false
  editingBudget.value = null
  Object.assign(form, { campaignName: '', totalBudget: 0, dailyLimit: 0, monthlyLimit: 0 })
}
</script>

<!--
  کامپوننت مدیریت بودجه و محدودیت‌های مالی کمپین‌ها
  شامل:
  1. تعیین بودجه کل، سقف روزانه و ماهانه
  2. نمایش آمار مصرف بودجه و هشدار نزدیک به سقف
  3. فرم افزودن/ویرایش بودجه برای هر کمپین
  توضیحات کامل به فارسی در کد
--> 

<template>
  <div class="bg-white rounded-lg shadow p-6">
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-xl font-semibold text-gray-900">مدیریت کدهای ارجاع</h2>
      <button 
        class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors"
        @click="showCodeForm = true; editingCode = null"
      >
        افزودن کد جدید
      </button>
    </div>

    <!-- لیست کدهای ارجاع -->
    <div class="overflow-x-auto">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              کد ارجاع
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              کاربر
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              تعداد استفاده
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              ارجاعات موفق
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              تاریخ ایجاد
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              وضعیت
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              عملیات
            </th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-for="code in referralCodes" :key="code.id">
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="flex items-center">
                <span class="text-sm font-mono text-gray-900 bg-gray-100 px-2 py-1 rounded">{{ code.code }}</span>
                <button 
                  class="mr-2 text-blue-600 hover:text-blue-800"
                  title="کپی کردن"
                  @click="copyToClipboard(code.code)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"></path>
                  </svg>
                </button>
              </div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="flex items-center">
                <img class="h-8 w-8 rounded-full ml-2" :src="code.userAvatar" :alt="code.userName">
                <div>
                  <div class="text-sm font-medium text-gray-900">{{ code.userName }}</div>
                  <div class="text-sm text-gray-500">{{ code.userEmail }}</div>
                </div>
              </div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="text-sm text-gray-900">{{ code.usageCount }}</div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="text-sm text-gray-900">{{ code.successfulReferrals }}</div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="text-sm text-gray-900">{{ formatDate(code.createdAt) }}</div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <span :class="code.active ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'" class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full">
                {{ code.active ? 'فعال' : 'غیرفعال' }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
              <button 
                class="text-blue-600 hover:text-blue-900 ml-4"
                @click="editingCode = code; showCodeForm = true"
              >
                ویرایش
              </button>
              <button 
                :class="code.active ? 'text-red-600 hover:text-red-900' : 'text-green-600 hover:text-green-900'"
                @click="toggleCodeStatus(code)"
              >
                {{ code.active ? 'غیرفعال' : 'فعال' }}
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- مودال فرم کد -->
    <div v-if="showCodeForm" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
      <div class="relative top-20 mx-auto p-5 border w-full max-w-md sm:max-w-lg md:max-w-xl shadow-lg rounded-md bg-white">
        <div class="mt-3">
          <h3 class="text-lg font-medium text-gray-900 mb-4">
            {{ editingCode ? 'ویرایش کد ارجاع' : 'افزودن کد جدید' }}
          </h3>
          
          <form @submit.prevent="saveCode">
            <div class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-gray-700">کد ارجاع</label>
                <input 
                  v-model="codeForm.code" 
                  type="text" 
                  class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                  required
                >
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700">کاربر</label>
                <select 
                  v-model="codeForm.userId"
                  class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                  required
                >
                  <option value="">انتخاب کاربر</option>
                  <option v-for="user in users" :key="user.id" :value="user.id">
                    {{ user.name }} ({{ user.email }})
                  </option>
                </select>
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700">توضیحات</label>
                <textarea 
                  v-model="codeForm.description" 
                  rows="2"
                  class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                ></textarea>
              </div>
              
              <div class="grid grid-cols-2 gap-6">
                <div>
                  <label class="block text-sm font-medium text-gray-700">حداکثر استفاده</label>
                  <input 
                    v-model="codeForm.maxUsage" 
                    type="number" 
                    class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                  >
                </div>
                
                <div>
                  <label class="block text-sm font-medium text-gray-700">تاریخ انقضا</label>
                  <input 
                    v-model="codeForm.expiryDate" 
                    type="date" 
                    class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                  >
                </div>
              </div>
            </div>
            
            <div class="flex justify-end space-x-3 space-x-reverse mt-6">
              <button 
                type="button"
                class="bg-gray-300 text-gray-700 px-4 py-2 rounded-md hover:bg-gray-400 transition-colors"
                @click="showCodeForm = false"
              >
                انصراف
              </button>
              <button 
                type="submit"
                class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition-colors"
              >
                {{ editingCode ? 'ویرایش' : 'افزودن' }}
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

// کدهای ارجاع
const referralCodes = ref([
  {
    id: 1,
    code: 'ALI123',
    userName: 'علی احمدی',
    userEmail: 'ali@example.com',
    userAvatar: '/default-avatar.png',
    usageCount: 15,
    successfulReferrals: 8,
    createdAt: '2024-01-01',
    active: true
  },
  {
    id: 2,
    code: 'FATEME456',
    userName: 'فاطمه محمدی',
    userEmail: 'fateme@example.com',
    userAvatar: '/default-avatar.png',
    usageCount: 12,
    successfulReferrals: 6,
    createdAt: '2024-01-05',
    active: true
  },
  {
    id: 3,
    code: 'MOHAMMAD789',
    userName: 'محمد رضایی',
    userEmail: 'mohammad@example.com',
    userAvatar: '/default-avatar.png',
    usageCount: 8,
    successfulReferrals: 4,
    createdAt: '2024-01-10',
    active: true
  }
])

// کاربران
const users = ref([
  { id: 1, name: 'علی احمدی', email: 'ali@example.com' },
  { id: 2, name: 'فاطمه محمدی', email: 'fateme@example.com' },
  { id: 3, name: 'محمد رضایی', email: 'mohammad@example.com' }
])

// State برای مودال
const showCodeForm = ref(false)
const editingCode = ref(null)

// فرم کد
const codeForm = reactive({
  code: '',
  userId: '',
  description: '',
  maxUsage: 0,
  expiryDate: ''
})

// کپی کردن کد
function copyToClipboard(code: string) {
  navigator.clipboard.writeText(code).then(() => {
    alert('کد کپی شد!')
  })
}

// تغییر وضعیت کد
function toggleCodeStatus(code: any) {
  code.active = !code.active
  // TODO: فراخوانی API برای تغییر وضعیت
}

// فرمت تاریخ
function formatDate(dateString: string) {
  const date = new Date(dateString)
  return date.toLocaleDateString('fa-IR')
}

// ذخیره کد
function saveCode() {
  if (editingCode.value) {
    // ویرایش کد موجود
    const index = referralCodes.value.findIndex(c => c.id === editingCode.value.id)
    if (index !== -1) {
      referralCodes.value[index] = { ...editingCode.value, ...codeForm }
    }
  } else {
    // افزودن کد جدید
    const selectedUser = users.value.find(u => u.id === parseInt(codeForm.userId))
    const newCode = {
      id: Date.now(),
      code: codeForm.code,
      userName: selectedUser?.name || '',
      userEmail: selectedUser?.email || '',
      userAvatar: '/default-avatar.png',
      usageCount: 0,
      successfulReferrals: 0,
      createdAt: new Date().toISOString().slice(0, 10),
      active: true
    }
    referralCodes.value.push(newCode)
  }
  
  // پاک کردن فرم
  Object.assign(codeForm, {
    code: '',
    userId: '',
    description: '',
    maxUsage: 0,
    expiryDate: ''
  })
  
  showCodeForm.value = false
  editingCode.value = null
}
</script> 

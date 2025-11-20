<template>
  <div class="bg-white rounded-lg shadow p-6">
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-xl font-semibold text-gray-900">مدیریت درخواست‌های پاداش</h2>
      <div class="flex space-x-2 space-x-reverse">
        <select v-model="statusFilter" class="border border-gray-300 rounded-md px-3 py-2">
          <option value="">همه وضعیت‌ها</option>
          <option value="pending">در انتظار</option>
          <option value="approved">تایید شده</option>
          <option value="rejected">رد شده</option>
        </select>
        <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors">
          فیلتر
        </button>
      </div>
    </div>

    <!-- لیست درخواست‌ها -->
    <div class="overflow-x-auto">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              کاربر
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              نوع پاداش
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              مبلغ درخواستی
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              تاریخ درخواست
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
          <tr v-for="request in filteredRequests" :key="request.id">
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="flex items-center">
                <img class="h-8 w-8 rounded-full ml-2" :src="request.userAvatar" :alt="request.userName">
                <div>
                  <div class="text-sm font-medium text-gray-900">{{ request.userName }}</div>
                  <div class="text-sm text-gray-500">{{ request.userEmail }}</div>
                </div>
              </div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <span class="text-sm text-gray-900">{{ getRewardTypeLabel(request.rewardType) }}</span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="text-sm font-medium text-green-600">{{ request.amount.toLocaleString() }} تومان</div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="text-sm text-gray-900">{{ formatDate(request.requestDate) }}</div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <span :class="getStatusClass(request.status)" class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full">
                {{ getStatusLabel(request.status) }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
              <button 
                v-if="request.status === 'pending'"
                class="text-green-600 hover:text-green-900 ml-4"
                @click="approveRequest(request)"
              >
                تایید
              </button>
              <button 
                v-if="request.status === 'pending'"
                class="text-red-600 hover:text-red-900"
                @click="rejectRequest(request)"
              >
                رد
              </button>
              <button 
                class="text-blue-600 hover:text-blue-900"
                @click="showRequestDetails(request)"
              >
                جزئیات
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- مودال جزئیات درخواست -->
    <div v-if="showDetails" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
      <div class="relative top-20 mx-auto p-5 border w-full max-w-md sm:max-w-lg md:max-w-xl shadow-lg rounded-md bg-white">
        <div class="mt-3">
          <h3 class="text-lg font-medium text-gray-900 mb-4">جزئیات درخواست پاداش</h3>
          
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700">کاربر:</label>
              <p class="text-sm text-gray-900">{{ selectedRequest?.userName }}</p>
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700">نوع پاداش:</label>
              <p class="text-sm text-gray-900">{{ getRewardTypeLabel(selectedRequest?.rewardType) }}</p>
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700">مبلغ درخواستی:</label>
              <p class="text-sm text-gray-900">{{ selectedRequest?.amount.toLocaleString() }} تومان</p>
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700">تاریخ درخواست:</label>
              <p class="text-sm text-gray-900">{{ formatDate(selectedRequest?.requestDate) }}</p>
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700">توضیحات:</label>
              <p class="text-sm text-gray-900">{{ selectedRequest?.description || 'توضیحی ارائه نشده' }}</p>
            </div>
            
            <div v-if="selectedRequest?.status === 'rejected'">
              <label class="block text-sm font-medium text-gray-700">دلیل رد:</label>
              <p class="text-sm text-gray-900">{{ selectedRequest?.rejectionReason || 'دلیلی ارائه نشده' }}</p>
            </div>
          </div>
          
          <div class="flex justify-end mt-6">
            <button 
              class="bg-gray-300 text-gray-700 px-4 py-2 rounded-md hover:bg-gray-400 transition-colors"
              @click="showDetails = false"
            >
              بستن
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

// درخواست‌های پاداش
const requests = ref([
  {
    id: 1,
    userName: 'علی احمدی',
    userEmail: 'ali@example.com',
    userAvatar: '/default-avatar.png',
    rewardType: 'cash',
    amount: 50000,
    requestDate: '2024-01-15',
    status: 'pending',
    description: 'درخواست پاداش برای ارجاعات موفق ماه گذشته'
  },
  {
    id: 2,
    userName: 'فاطمه محمدی',
    userEmail: 'fateme@example.com',
    userAvatar: '/default-avatar.png',
    rewardType: 'credit',
    amount: 75000,
    requestDate: '2024-01-14',
    status: 'approved',
    description: 'پاداش ارجاعات هفته گذشته'
  },
  {
    id: 3,
    userName: 'محمد رضایی',
    userEmail: 'mohammad@example.com',
    userAvatar: '/default-avatar.png',
    rewardType: 'cash',
    amount: 30000,
    requestDate: '2024-01-13',
    status: 'rejected',
    description: 'درخواست پاداش ارجاعات',
    rejectionReason: 'مبلغ درخواستی بیش از حد مجاز است'
  }
])

// فیلترها
const statusFilter = ref('')
const showDetails = ref(false)
const selectedRequest = ref(null)

// محاسبه درخواست‌های فیلتر شده
const filteredRequests = computed(() => {
  let filtered = requests.value
  
  if (statusFilter.value) {
    filtered = filtered.filter(r => r.status === statusFilter.value)
  }
  
  return filtered
})

// برچسب نوع پاداش
function getRewardTypeLabel(type: string) {
  const labels = {
    cash: 'نقدی',
    credit: 'اعتبار',
    points: 'امتیاز'
  }
  return labels[type] || type
}

// برچسب وضعیت
function getStatusLabel(status: string) {
  const labels = {
    pending: 'در انتظار',
    approved: 'تایید شده',
    rejected: 'رد شده'
  }
  return labels[status] || status
}

// کلاس وضعیت
function getStatusClass(status: string) {
  const classes = {
    pending: 'bg-yellow-100 text-yellow-800',
    approved: 'bg-green-100 text-green-800',
    rejected: 'bg-red-100 text-red-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

// فرمت تاریخ
function formatDate(dateString: string) {
  const date = new Date(dateString)
  return date.toLocaleDateString('fa-IR')
}

// تایید درخواست
function approveRequest(request: any) {
  request.status = 'approved'
  // TODO: فراخوانی API برای تایید درخواست
  console.log('درخواست تایید شد:', request.id)
}

// رد درخواست
function rejectRequest(request: any) {
  const reason = prompt('دلیل رد درخواست:')
  if (reason) {
    request.status = 'rejected'
    request.rejectionReason = reason
    // TODO: فراخوانی API برای رد درخواست
    console.log('درخواست رد شد:', request.id)
  }
}

// نمایش جزئیات درخواست
function showRequestDetails(request: any) {
  selectedRequest.value = request
  showDetails.value = true
}
</script> 

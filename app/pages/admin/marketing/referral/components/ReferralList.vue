<template>
  <div class="bg-white rounded-lg shadow p-6">
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-xl font-semibold text-gray-900">لیست ارجاعات</h2>
      <div class="flex space-x-2 space-x-reverse">
        <select v-model="statusFilter" class="border border-gray-300 rounded-md px-3 py-2">
          <option value="">همه وضعیت‌ها</option>
          <option value="pending">در انتظار</option>
          <option value="successful">موفق</option>
          <option value="failed">ناموفق</option>
        </select>
        <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors">
          فیلتر
        </button>
      </div>
    </div>

    <!-- لیست ارجاعات -->
    <div class="overflow-x-auto">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              ارجاع‌دهنده
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              ارجاع‌شونده
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              کد ارجاع
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              مبلغ خرید
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              پاداش
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              تاریخ ارجاع
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
          <tr v-for="referral in filteredReferrals" :key="referral.id">
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="flex items-center">
                <img class="h-8 w-8 rounded-full ml-2" :src="referral.referrerAvatar" :alt="referral.referrerName">
                <div>
                  <div class="text-sm font-medium text-gray-900">{{ referral.referrerName }}</div>
                  <div class="text-sm text-gray-500">{{ referral.referrerEmail }}</div>
                </div>
              </div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="flex items-center">
                <img class="h-8 w-8 rounded-full ml-2" :src="referral.referredAvatar" :alt="referral.referredName">
                <div>
                  <div class="text-sm font-medium text-gray-900">{{ referral.referredName }}</div>
                  <div class="text-sm text-gray-500">{{ referral.referredEmail }}</div>
                </div>
              </div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <span class="text-sm font-mono text-gray-900 bg-gray-100 px-2 py-1 rounded">{{ referral.referralCode }}</span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="text-sm text-gray-900">{{ referral.purchaseAmount.toLocaleString() }} تومان</div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="text-sm font-medium text-green-600">{{ referral.rewardAmount.toLocaleString() }} تومان</div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="text-sm text-gray-900">{{ formatDate(referral.referralDate) }}</div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <span :class="getStatusClass(referral.status)" class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full">
                {{ getStatusLabel(referral.status) }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
              <button 
                class="text-blue-600 hover:text-blue-900 ml-4"
                @click="$emit('show-referral-details', referral)"
              >
                جزئیات
              </button>
              <button 
                class="text-green-600 hover:text-green-900"
                @click="$emit('edit-referral', referral)"
              >
                ویرایش
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- صفحه‌بندی -->
    <div class="flex justify-between items-center mt-6">
      <div class="text-sm text-gray-700">
        نمایش {{ (currentPage - 1) * pageSize + 1 }} تا {{ Math.min(currentPage * pageSize, totalItems) }} از {{ totalItems }} مورد
      </div>
      <div class="flex space-x-2 space-x-reverse">
        <button 
          :disabled="currentPage === 1"
          class="px-3 py-2 border border-gray-300 rounded-md text-sm disabled:opacity-50 disabled:cursor-not-allowed"
          @click="currentPage--"
        >
          قبلی
        </button>
        <span class="px-3 py-2 text-sm text-gray-700">{{ currentPage }} از {{ totalPages }}</span>
        <button 
          :disabled="currentPage === totalPages"
          class="px-3 py-2 border border-gray-300 rounded-md text-sm disabled:opacity-50 disabled:cursor-not-allowed"
          @click="currentPage++"
        >
          بعدی
        </button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>
</script>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue'
import { useAuth } from '~/composables/useAuth'

// احراز هویت
const { user, isAuthenticated } = useAuth();

// بررسی دسترسی admin
const hasAccess = computed(() => {
  if (!isAuthenticated.value) {
    return false;
  }

  const userRole = user.value?.role?.toLowerCase() || ''
  const adminRoles = ['admin', 'developer']
  return adminRoles.includes(userRole)
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

// تعریف emit events
defineEmits<{
  (e: 'show-referral-details', referral: unknown): void
  (e: 'edit-referral', referral: unknown): void
}>()

// ارجاعات
const referrals = ref([
  {
    id: 1,
    referrerName: 'علی احمدی',
    referrerEmail: 'ali@example.com',
    referrerAvatar: '/default-avatar.png',
    referredName: 'فاطمه محمدی',
    referredEmail: 'fateme@example.com',
    referredAvatar: '/default-avatar.png',
    referralCode: 'ALI123',
    purchaseAmount: 250000,
    rewardAmount: 25000,
    referralDate: '2024-01-15',
    status: 'successful'
  },
  {
    id: 2,
    referrerName: 'محمد رضایی',
    referrerEmail: 'mohammad@example.com',
    referrerAvatar: '/default-avatar.png',
    referredName: 'زهرا کریمی',
    referredEmail: 'zahra@example.com',
    referredAvatar: '/default-avatar.png',
    referralCode: 'MOHAMMAD789',
    purchaseAmount: 180000,
    rewardAmount: 18000,
    referralDate: '2024-01-14',
    status: 'pending'
  },
  {
    id: 3,
    referrerName: 'فاطمه محمدی',
    referrerEmail: 'fateme@example.com',
    referrerAvatar: '/default-avatar.png',
    referredName: 'احمد نوری',
    referredEmail: 'ahmad@example.com',
    referredAvatar: '/default-avatar.png',
    referralCode: 'FATEME456',
    purchaseAmount: 320000,
    rewardAmount: 32000,
    referralDate: '2024-01-13',
    status: 'successful'
  }
])

// فیلترها و صفحه‌بندی
const statusFilter = ref('')
const currentPage = ref(1)
const pageSize = 10

// محاسبه ارجاعات فیلتر شده
const filteredReferrals = computed(() => {
  let filtered = referrals.value
  
  if (statusFilter.value) {
    filtered = filtered.filter(r => r.status === statusFilter.value)
  }
  
  return filtered
})

// محاسبه آمار صفحه‌بندی
const totalItems = computed(() => filteredReferrals.value.length)
const totalPages = computed(() => Math.ceil(totalItems.value / pageSize))

// برچسب وضعیت
function getStatusLabel(status: string) {
  const labels = {
    pending: 'در انتظار',
    successful: 'موفق',
    failed: 'ناموفق'
  }
  return labels[status] || status
}

// کلاس وضعیت
function getStatusClass(status: string) {
  const classes = {
    pending: 'bg-yellow-100 text-yellow-800',
    successful: 'bg-green-100 text-green-800',
    failed: 'bg-red-100 text-red-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

// فرمت تاریخ
function formatDate(dateString: string) {
  const date = new Date(dateString)
  return date.toLocaleDateString('fa-IR')
}
</script> 

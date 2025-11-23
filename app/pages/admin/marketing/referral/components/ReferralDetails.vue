<template>
  <div v-if="hasAccess && referral" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
    <div class="relative top-20 mx-auto p-5 border w-full max-w-md sm:max-w-lg md:max-w-xl shadow-lg rounded-md bg-white">
      <div class="mt-3">
        <h3 class="text-lg font-medium text-gray-900 mb-4">جزئیات ارجاع</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700">ارجاع‌دهنده:</label>
            <p class="text-sm text-gray-900">{{ referral.referrerName }}</p>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">ارجاع‌شونده:</label>
            <p class="text-sm text-gray-900">{{ referral.referredName }}</p>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">کد ارجاع:</label>
            <p class="text-sm text-gray-900">{{ referral.referralCode }}</p>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">مبلغ خرید:</label>
            <p class="text-sm text-gray-900">{{ referral.purchaseAmount?.toLocaleString() }} تومان</p>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">پاداش:</label>
            <p class="text-sm text-gray-900">{{ referral.rewardAmount?.toLocaleString() }} تومان</p>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">تاریخ ارجاع:</label>
            <p class="text-sm text-gray-900">{{ formatDate(referral.referralDate) }}</p>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">وضعیت:</label>
            <p class="text-sm text-gray-900">{{ getStatusLabel(referral.status) }}</p>
          </div>
        </div>
        <div class="flex justify-end mt-6 space-x-2 space-x-reverse">
          <button class="bg-green-600 text-white px-4 py-2 rounded-md hover:bg-green-700 transition-colors" @click="$emit('edit', referral)">ویرایش</button>
          <button class="bg-gray-300 text-gray-700 px-4 py-2 rounded-md hover:bg-gray-400 transition-colors" @click="$emit('close')">بستن</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>
</script>

<script setup lang="ts">
import { computed, onMounted, watch } from 'vue';
import { useAuth } from '~/composables/useAuth';

// احراز هویت
const { user, isAuthenticated } = useAuth()

// بررسی دسترسی admin
const hasAccess = computed(() => {
  if (!isAuthenticated.value) {
    return false
  }

  const userRole = user.value?.role?.toLowerCase() || ''
  const adminRoles = ['admin', 'developer']
  return adminRoles.includes(userRole)
})

// بررسی احراز هویت و دسترسی admin - نمایش 404 در صورت عدم دسترسی
const checkAuth = async (): Promise<void> => {
  if (!hasAccess.value) {
    await navigateTo('/404', { external: false })
  }
}

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

const _props = defineProps({ referral: Object });
const _emit = defineEmits(['close', 'edit']);

function formatDate(dateString: string) {
  if (!dateString) return ''
  const date = new Date(dateString)
  return date.toLocaleDateString('fa-IR')
}
function getStatusLabel(status: string) {
  const labels = {
    pending: 'در انتظار',
    successful: 'موفق',
    failed: 'ناموفق'
  }
  return labels[status] || status
}
</script> 

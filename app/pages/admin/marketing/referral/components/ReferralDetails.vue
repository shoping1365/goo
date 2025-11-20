<template>
  <div v-if="referral" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
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

<script setup lang="ts">
const props = defineProps({ referral: Object })
const emit = defineEmits(['close', 'edit'])

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

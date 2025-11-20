<template>
  <div class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
    <div class="relative top-20 mx-auto p-5 border w-full max-w-md sm:max-w-lg md:max-w-xl shadow-lg rounded-md bg-white">
      <div class="mt-3">
        <h3 class="text-lg font-medium text-gray-900 mb-4">
          {{ referral ? 'ویرایش ارجاع' : 'افزودن ارجاع جدید' }}
        </h3>
        <form @submit.prevent="save">
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700">ارجاع‌دهنده</label>
              <input v-model="form.referrerName" type="text" class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">ارجاع‌شونده</label>
              <input v-model="form.referredName" type="text" class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">کد ارجاع</label>
              <input v-model="form.referralCode" type="text" class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">مبلغ خرید</label>
              <input v-model="form.purchaseAmount" type="number" class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">پاداش</label>
              <input v-model="form.rewardAmount" type="number" class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">تاریخ ارجاع</label>
              <input v-model="form.referralDate" type="date" class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">وضعیت</label>
              <select v-model="form.status" class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2">
                <option value="pending">در انتظار</option>
                <option value="successful">موفق</option>
                <option value="failed">ناموفق</option>
              </select>
            </div>
          </div>
          <div class="flex justify-end space-x-3 space-x-reverse mt-6">
            <button type="button" class="bg-gray-300 text-gray-700 px-4 py-2 rounded-md hover:bg-gray-400 transition-colors" @click="$emit('cancel')">انصراف</button>
            <button type="submit" class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition-colors">
              {{ referral ? 'ویرایش' : 'افزودن' }}
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive, watch, toRefs } from 'vue'
const props = defineProps({ referral: Object })
const emit = defineEmits(['save', 'cancel'])

const form = reactive({
  referrerName: '',
  referredName: '',
  referralCode: '',
  purchaseAmount: 0,
  rewardAmount: 0,
  referralDate: '',
  status: 'pending'
})

watch(() => props.referral, (val) => {
  if (val) Object.assign(form, val)
  else {
    form.referrerName = ''
    form.referredName = ''
    form.referralCode = ''
    form.purchaseAmount = 0
    form.rewardAmount = 0
    form.referralDate = ''
    form.status = 'pending'
  }
}, { immediate: true })

function save() {
  emit('save', { ...form })
}
</script> 

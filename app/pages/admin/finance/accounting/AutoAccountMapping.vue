<template>
  <div>
    <div class="flex items-center justify-between mb-4">
      <h4 class="text-base font-semibold text-gray-900">نگاشت خودکار حساب‌ها</h4>
      <button class="px-4 py-2 bg-green-600 hover:bg-green-700 text-white rounded-lg" @click="runAutoMapping">اجرای نگاشت خودکار</button>
    </div>
    <div class="overflow-x-auto">
      <table class="w-full text-sm">
        <thead>
          <tr class="border-b border-gray-200 bg-gray-50">
            <th class="text-right py-3 px-4 font-medium text-gray-600">حساب سیستم</th>
            <th class="text-right py-3 px-4 font-medium text-gray-600">کد سیستم</th>
            <th class="text-right py-3 px-4 font-medium text-gray-600">حساب نرم‌افزار</th>
            <th class="text-right py-3 px-4 font-medium text-gray-600">کد نرم‌افزار</th>
            <th class="text-right py-3 px-4 font-medium text-gray-600">وضعیت</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="mapping in autoMappings" :key="mapping.id" class="border-b border-gray-100 hover:bg-gray-50">
            <td class="py-3 px-4">{{ mapping.systemAccount }}</td>
            <td class="py-3 px-4 font-mono">{{ mapping.systemCode }}</td>
            <td class="py-3 px-4">{{ mapping.softwareAccount }}</td>
            <td class="py-3 px-4 font-mono">{{ mapping.softwareCode }}</td>
            <td class="py-3 px-4">
              <span :class="mapping.status === 'matched' ? 'text-green-600' : 'text-yellow-600'">
                {{ mapping.status === 'matched' ? 'نگاشت شد' : 'نیاز به بررسی' }}
              </span>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
<script setup lang="ts">
import { ref } from 'vue'
const autoMappings = ref([
  { id: 1, systemAccount: 'فروش کالا', systemCode: '4001', softwareAccount: 'فروش کالا', softwareCode: '5001', status: 'matched' },
  { id: 2, systemAccount: 'خرید کالا', systemCode: '5001', softwareAccount: 'خرید کالا', softwareCode: '6001', status: 'matched' },
  { id: 3, systemAccount: 'مالیات', systemCode: '2201', softwareAccount: 'مالیات بر ارزش افزوده', softwareCode: '2201', status: 'matched' },
  { id: 4, systemAccount: 'بانک', systemCode: '1101', softwareAccount: 'بانک ملت', softwareCode: '1101', status: 'pending' }
])
const runAutoMapping = () => {
  // TODO: اجرای نگاشت خودکار
  alert('نگاشت خودکار اجرا شد!')
}
</script>
<!-- نگاشت خودکار حساب‌ها --> 

<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200 w-full">
    <div class="p-6 border-b border-gray-200 flex items-center justify-between">
      <div>
        <h2 class="text-lg font-semibold text-gray-900">مدیریت کمپین‌ها</h2>
        <p class="text-gray-600 mt-1">ایجاد و مدیریت کمپین‌های تخفیف</p>
      </div>
      <button @click="$emit('show-form', null)" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors">
        <span class="i-heroicons-plus ml-2"></span>
        افزودن کمپین جدید
      </button>
    </div>
    <div class="p-6 border-b border-gray-200">
      <input v-model="search" type="text" placeholder="جستجو بر اساس نام یا توضیح..." class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
    </div>
    <div class="p-6">
      <table class="w-full text-sm">
        <thead>
          <tr class="border-b border-gray-200">
            <th class="py-2 px-3 text-right">نام کمپین</th>
            <th class="py-2 px-3 text-right">نوع</th>
            <th class="py-2 px-3 text-right">توضیحات</th>
            <th class="py-2 px-3 text-right">تاریخ شروع</th>
            <th class="py-2 px-3 text-right">تاریخ پایان</th>
            <th class="py-2 px-3 text-right">وضعیت</th>
            <th class="py-2 px-3 text-right">عملیات</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="campaign in filteredCampaigns" :key="campaign.id" class="border-b border-gray-100 hover:bg-gray-50">
            <td class="py-2 px-3">{{ campaign.name }}</td>
            <td class="py-2 px-3">{{ campaign.type }}</td>
            <td class="py-2 px-3">{{ campaign.description }}</td>
            <td class="py-2 px-3">{{ campaign.startsAt }}</td>
            <td class="py-2 px-3">{{ campaign.endsAt }}</td>
            <td class="py-2 px-3">
              <span :class="['px-2 py-1 rounded-full text-xs', getStatusClass(campaign.status)]">{{ getStatusText(campaign.status) }}</span>
            </td>
            <td class="py-2 px-3 flex gap-2">
              <button @click="$emit('show-form', campaign)" class="text-blue-600 hover:text-blue-900">ویرایش</button>
              <button @click="deleteCampaign(campaign.id)" class="text-red-600 hover:text-red-900">حذف</button>
            </td>
          </tr>
          <tr v-if="filteredCampaigns.length === 0">
            <td colspan="7" class="text-center text-gray-400 py-8">کمپینی یافت نشد.</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
const search = ref('')
const campaigns = ref([
  { id: 1, name: 'کمپین تابستان', type: 'فصلی', description: 'تخفیف ویژه تابستان', startsAt: '1403/04/01', endsAt: '1403/04/31', status: 'active' },
  { id: 2, name: 'کمپین جمعه سیاه', type: 'فلش', description: 'تخفیف ویژه جمعه سیاه', startsAt: '1403/08/01', endsAt: '1403/08/02', status: 'draft' }
])
const filteredCampaigns = computed(() => {
  if (!search.value) return campaigns.value
  return campaigns.value.filter(c =>
    c.name.includes(search.value) ||
    c.description.includes(search.value)
  )
})
function deleteCampaign(id: number) {
  campaigns.value = campaigns.value.filter(c => c.id !== id)
}
function getStatusClass(status: string) {
  switch (status) {
    case 'active': return 'bg-green-100 text-green-800'
    case 'draft': return 'bg-gray-100 text-gray-800'
    case 'ended': return 'bg-red-100 text-red-800'
    default: return 'bg-gray-100 text-gray-800'
  }
}
function getStatusText(status: string) {
  switch (status) {
    case 'active': return 'فعال'
    case 'draft': return 'پیش‌نویس'
    case 'ended': return 'پایان یافته'
    default: return 'نامشخص'
  }
}
</script> 

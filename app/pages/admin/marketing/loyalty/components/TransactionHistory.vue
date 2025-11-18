<template>
  <div class="bg-white rounded-lg shadow p-6">
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-xl font-semibold text-gray-900">تاریخچه تراکنش‌ها</h2>
      <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors">
        بروزرسانی
      </button>
    </div>
    <div class="overflow-x-auto">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کاربر</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع تراکنش</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">امتیاز</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">توضیحات</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ</th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-for="tx in transactions" :key="tx.id">
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="flex items-center">
                <img class="h-8 w-8 rounded-full ml-2" :src="tx.avatar" :alt="tx.user">
                <span class="text-sm font-medium text-gray-900">{{ tx.user }}</span>
              </div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <span :class="getTypeColor(tx.type)" class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full">
                {{ getTypeName(tx.type) }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ tx.points }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ tx.description }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ formatDate(tx.date) }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const transactions = ref([
  {
    id: 1,
    user: 'علی احمدی',
    avatar: '/default-avatar.png',
    type: 'earn',
    points: 100,
    description: 'دریافت امتیاز بابت خرید',
    date: '2024-01-10'
  },
  {
    id: 2,
    user: 'فاطمه محمدی',
    avatar: '/default-avatar.png',
    type: 'spend',
    points: -50,
    description: 'استفاده از امتیاز برای تخفیف',
    date: '2024-01-09'
  },
  {
    id: 3,
    user: 'محمد رضایی',
    avatar: '/default-avatar.png',
    type: 'earn',
    points: 200,
    description: 'دریافت امتیاز بابت معرفی دوست',
    date: '2024-01-08'
  },
  {
    id: 4,
    user: 'زهرا کریمی',
    avatar: '/default-avatar.png',
    type: 'spend',
    points: -30,
    description: 'استفاده از امتیاز برای ارسال رایگان',
    date: '2024-01-07'
  },
  {
    id: 5,
    user: 'حسین نوری',
    avatar: '/default-avatar.png',
    type: 'earn',
    points: 150,
    description: 'دریافت امتیاز بابت تکمیل پروفایل',
    date: '2024-01-06'
  }
])

function getTypeColor(type) {
  return type === 'earn' ? 'bg-green-100 text-green-800' : 'bg-blue-100 text-blue-800'
}

function getTypeName(type) {
  return type === 'earn' ? 'دریافت' : 'استفاده'
}

function formatDate(dateString) {
  const date = new Date(dateString)
  return date.toLocaleDateString('fa-IR')
}
</script> 

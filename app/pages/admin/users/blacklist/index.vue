<template>
  <div class="space-y-6">
    <!-- هدر صفحه -->
    <div class="flex justify-between items-center">
      <h1 class="text-2xl font-bold text-gray-900">لیست سیاه کاربران</h1>
      <button class="bg-red-600 text-white px-4 py-2 rounded-lg hover:bg-red-700 transition-colors">
        افزودن کاربر به لیست سیاه
      </button>
    </div>

    <!-- فیلترها -->
    <div class="bg-white p-6 rounded-lg shadow">
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">جستجو</label>
          <input 
            type="text" 
            placeholder="جستجو بر اساس نام، ایمیل یا شماره تلفن..."
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">دلیل مسدودیت</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option value="">همه دلایل</option>
            <option value="spam">اسپم</option>
            <option value="fraud">کلاهبرداری</option>
            <option value="abuse">سوء استفاده</option>
            <option value="other">سایر</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">وضعیت</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option value="">همه وضعیت‌ها</option>
            <option value="active">فعال</option>
            <option value="expired">منقضی شده</option>
          </select>
        </div>
      </div>
    </div>

    <!-- جدول کاربران -->
    <div class="bg-white rounded-lg shadow overflow-hidden">
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                کاربر
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                دلیل مسدودیت
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                تاریخ مسدودیت
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                مدت مسدودیت
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
            <tr v-for="user in blacklistedUsers" :key="user.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="">
                    <img class="h-10 w-10 rounded-full" :src="user.avatar" :alt="user.name" />
                  </div>
                  <div class="mr-4">
                    <div class="text-sm font-medium text-gray-900">{{ user.name }}</div>
                    <div class="text-sm text-gray-500">{{ user.email }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full" 
                      :class="getReasonClass(user.reason)">
                  {{ user.reason }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ formatDate(user.blockedAt) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ user.duration }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full" 
                      :class="getStatusClass(user.status)">
                  {{ user.status }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">ویرایش</button>
                <button class="text-green-600 hover:text-green-900 ml-3">حذف از لیست</button>
                <button v-if="canDeleteBlacklistedUser" class="text-red-600 hover:text-red-900">حذف</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- پیجینیشن -->
    <div class="bg-white px-4 py-3 flex items-center justify-between border-t border-gray-200 sm:px-6">
      <div class="flex-1 flex justify-between sm:hidden">
        <button class="relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50">
          قبلی
        </button>
        <button class="mr-3 relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50">
          بعدی
        </button>
      </div>
      <div class="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
        <div>
          <p class="text-sm text-gray-700">
            نمایش
            <span class="font-medium">1</span>
            تا
            <span class="font-medium">10</span>
            از
            <span class="font-medium">97</span>
            نتیجه
          </p>
        </div>
        <div>
          <nav class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px">
            <button class="relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50">
              <span class="sr-only">قبلی</span>
              <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd" />
              </svg>
            </button>
            <button class="bg-blue-50 border-blue-500 text-blue-600 relative inline-flex items-center px-4 py-2 border text-sm font-medium">
              1
            </button>
            <button class="bg-white border-gray-300 text-gray-500 hover:bg-gray-50 relative inline-flex items-center px-4 py-2 border text-sm font-medium">
              2
            </button>
            <button class="bg-white border-gray-300 text-gray-500 hover:bg-gray-50 relative inline-flex items-center px-4 py-2 border text-sm font-medium">
              3
            </button>
            <button class="relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50">
              <span class="sr-only">بعدی</span>
              <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z" clip-rule="evenodd" />
              </svg>
            </button>
          </nav>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
</script>

<script setup lang="ts">
import { computed, ref } from 'vue';

// تعریف layout صفحه
definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// تعریف interface برای کاربر
interface BlacklistedUser {
  id: number
  name: string
  email: string
  avatar: string
  reason: string
  blockedAt: string
  duration: string
  status: string
}

// داده‌های نمونه
const blacklistedUsers = ref<BlacklistedUser[]>([
  {
    id: 1,
    name: 'احمد محمدی',
    email: 'ahmad@example.com',
    avatar: '/default-avatar.svg',
    reason: 'spam',
    blockedAt: '2024-01-15',
    duration: '30 روز',
    status: 'active'
  },
  {
    id: 2,
    name: 'فاطمه احمدی',
    email: 'fateme@example.com',
    avatar: '/default-avatar.svg',
    reason: 'fraud',
    blockedAt: '2024-01-10',
    duration: '90 روز',
    status: 'active'
  },
  {
    id: 3,
    name: 'علی رضایی',
    email: 'ali@example.com',
    avatar: '/default-avatar.svg',
    reason: 'abuse',
    blockedAt: '2024-01-05',
    duration: '7 روز',
    status: 'expired'
  }
])

// تابع فرمت کردن تاریخ
const formatDate = (dateString) => {
  return new Date(dateString).toLocaleDateString('fa-IR')
}

// تابع کلاس‌های دلیل مسدودیت
const getReasonClass = (reason: string) => {
  const classes: Record<string, string> = {
    spam: 'bg-yellow-100 text-yellow-800',
    fraud: 'bg-red-100 text-red-800',
    abuse: 'bg-orange-100 text-orange-800',
    other: 'bg-gray-100 text-gray-800'
  }
  return classes[reason] || classes.other
}

// تابع کلاس‌های وضعیت
const getStatusClass = (status: string) => {
  const classes: Record<string, string> = {
    active: 'bg-red-100 text-red-800',
    expired: 'bg-gray-100 text-gray-800'
  }
  return classes[status] || classes.expired
}

// Permission check (simple implementation)
const hasPermission = (permission: string) => true

// Computed برای چک کردن پرمیژن حذف
const canDeleteBlacklistedUser = computed(() => hasPermission('blacklist.delete'))
</script> 

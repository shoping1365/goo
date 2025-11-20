<template>
  <div class="space-y-8">
    <!-- Header -->
    <div>
      <h3 class="text-lg font-semibold text-gray-800">مدیریت دسترسی و امنیت</h3>
      <p class="text-gray-600 text-sm">تعریف نقش‌ها، مجوزها و نظارت بر فعالیت‌ها</p>
    </div>
    <!-- Roles Management -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h4 class="text-md font-semibold text-gray-800 mb-4">مدیریت نقش‌ها</h4>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div v-for="role in roles" :key="role.id" class="border border-gray-200 rounded-lg p-6">
          <div class="flex items-center justify-between mb-2">
            <h5 class="font-medium text-gray-900">{{ role.name }}</h5>
            <span :class="getRoleBadgeClass(role.level)">{{ getRoleLevelText(role.level) }}</span>
          </div>
          <p class="text-sm text-gray-600 mb-3">{{ role.description }}</p>
          <div class="space-y-2">
            <div v-for="permission in role.permissions" :key="permission" class="flex items-center">
              <svg class="w-4 h-4 text-green-500 mr-2" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd"></path>
              </svg>
              <span class="text-xs text-gray-700">{{ permission }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- User Management -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h4 class="text-md font-semibold text-gray-800 mb-4">مدیریت کاربران</h4>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-4 py-3 text-right text-xs font-medium text-gray-500">کاربر</th>
              <th class="px-4 py-3 text-right text-xs font-medium text-gray-500">نقش</th>
              <th class="px-4 py-3 text-right text-xs font-medium text-gray-500">وضعیت</th>
              <th class="px-4 py-3 text-right text-xs font-medium text-gray-500">آخرین ورود</th>
              <th class="px-4 py-3 text-right text-xs font-medium text-gray-500">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="userItem in users" :key="userItem.id" class="hover:bg-gray-50">
              <td class="px-4 py-3 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="w-8 h-8 bg-gradient-to-r from-blue-400 to-purple-500 rounded-full flex items-center justify-center">
                    <span class="text-white text-xs font-medium">{{ getInitials(userItem.name) }}</span>
                  </div>
                  <div class="mr-3">
                    <div class="text-sm font-medium text-gray-900">{{ userItem.name }}</div>
                    <div class="text-xs text-gray-500">{{ userItem.email }}</div>
                  </div>
                </div>
              </td>
              <td class="px-4 py-3 whitespace-nowrap">
                <select v-model="userItem.roleId" class="text-sm border border-gray-300 rounded px-2 py-1">
                  <option v-for="role in roles" :key="role.id" :value="role.id">{{ role.name }}</option>
                </select>
              </td>
              <td class="px-4 py-3 whitespace-nowrap">
                <span :class="userItem.active ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'" class="inline-flex px-2 py-1 text-xs font-medium rounded-full">
                  {{ userItem.active ? 'فعال' : 'غیرفعال' }}
                </span>
              </td>
              <td class="px-4 py-3 whitespace-nowrap text-xs text-gray-500">{{ formatDate(userItem.lastLogin) }}</td>
              <td class="px-4 py-3 whitespace-nowrap text-sm">
                <button class="text-blue-600 hover:text-blue-800 mr-2" @click="toggleUserStatus(userItem.id)">
                  {{ userItem.active ? 'غیرفعال' : 'فعال' }}
                </button>
                <button v-if="hasPermission('user.delete')" class="text-red-600 hover:text-red-800" @click="deleteUser(userItem.id)">حذف</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
    <!-- Activity Log -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h4 class="text-md font-semibold text-gray-800 mb-4">لاگ فعالیت‌ها</h4>
      <div class="space-y-3">
        <div v-for="activity in activities" :key="activity.id" class="flex items-start space-x-3 space-x-reverse p-3 bg-gray-50 rounded-lg">
          <div class="flex-shrink-0">
            <div :class="getActivityIconClass(activity.type)" class="w-8 h-8 rounded-full flex items-center justify-center">
              <svg class="w-4 h-4 text-white" fill="currentColor" viewBox="0 0 20 20">
                <path v-if="activity.type === 'login'" d="M3 4a1 1 0 011-1h12a1 1 0 011 1v2a1 1 0 01-1 1H4a1 1 0 01-1-1V4zM3 10a1 1 0 011-1h6a1 1 0 011 1v6a1 1 0 01-1 1H4a1 1 0 01-1-1v-6zM14 9a1 1 0 00-1 1v6a1 1 0 001 1h2a1 1 0 001-1v-6a1 1 0 00-1-1h-2z"></path>
                <path v-else-if="activity.type === 'action'" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                <path v-else d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z"></path>
              </svg>
            </div>
          </div>
          <div class="flex-1">
            <div class="text-sm text-gray-900">{{ activity.description }}</div>
            <div class="text-xs text-gray-500 mt-1">{{ formatDate(activity.timestamp) }} - {{ activity.user }}</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
declare const useAuth: () => { user: { id?: number; name?: string; email?: string } | null; hasPermission: (perm: string) => boolean }
</script>

<script setup lang="ts">
import { ref } from 'vue';

// استفاده از کامپوزابل احراز هویت
const { hasPermission } = useAuth()

const roles = ref([
  {
    id: 1,
    name: 'مدیر کل',
    level: 'admin',
    description: 'دسترسی کامل به تمام بخش‌ها',
    permissions: ['مدیریت کاربران', 'تنظیمات سیستم', 'گزارشات کامل', 'پشتیبان‌گیری']
  },
  {
    id: 2,
    name: 'اپراتور',
    level: 'operator',
    description: 'دسترسی به عملیات روزانه',
    permissions: ['ارسال پیامک', 'مشاهده پاسخ‌ها', 'گزارشات محدود']
  },
  {
    id: 3,
    name: 'مشاهده‌کننده',
    level: 'viewer',
    description: 'فقط مشاهده گزارشات',
    permissions: ['مشاهده پاسخ‌ها', 'گزارشات پایه']
  }
])
const users = ref([
  { id: 1, name: 'علی احمدی', email: 'ali@example.com', roleId: 1, active: true, lastLogin: '2024-06-07T10:30:00Z' },
  { id: 2, name: 'زهرا رضایی', email: 'zahra@example.com', roleId: 2, active: true, lastLogin: '2024-06-07T09:15:00Z' },
  { id: 3, name: 'محمد موسوی', email: 'mohammad@example.com', roleId: 3, active: false, lastLogin: '2024-06-06T16:45:00Z' }
])
const activities = ref([
  { id: 1, type: 'login', description: 'ورود به سیستم', user: 'علی احمدی', timestamp: '2024-06-07T10:30:00Z' },
  { id: 2, type: 'action', description: 'ارسال پیامک گروهی به 50 سفارش', user: 'زهرا رضایی', timestamp: '2024-06-07T09:20:00Z' },
  { id: 3, type: 'action', description: 'ایجاد پشتیبان جدید', user: 'علی احمدی', timestamp: '2024-06-07T08:45:00Z' },
  { id: 4, type: 'warning', description: 'تلاش ورود ناموفق', user: 'کاربر ناشناس', timestamp: '2024-06-07T08:30:00Z' }
])
const getRoleBadgeClass = (level: string) => {
  return level === 'admin' ? 'bg-red-100 text-red-800 px-2 py-1 rounded-full text-xs' :
         level === 'operator' ? 'bg-blue-100 text-blue-800 px-2 py-1 rounded-full text-xs' :
         'bg-gray-100 text-gray-800 px-2 py-1 rounded-full text-xs'
}
const getRoleLevelText = (level: string) => {
  return level === 'admin' ? 'مدیر' : level === 'operator' ? 'اپراتور' : 'مشاهده‌کننده'
}
const getInitials = (name: string) => {
  return name.split(' ').map(n => n[0]).join('').toUpperCase()
}
const getActivityIconClass = (type: string) => {
  return type === 'login' ? 'bg-green-500' : type === 'action' ? 'bg-blue-500' : 'bg-yellow-500'
}
const formatDate = (date: string) => new Date(date).toLocaleString('fa-IR')
const toggleUserStatus = (userId: number) => {
  const user = users.value.find(u => u.id === userId)
  if (user) user.active = !user.active
}
const deleteUser = (userId: number) => {
  if (confirm('آیا مطمئن هستید که می‌خواهید این کاربر را حذف کنید؟')) {
    users.value = users.value.filter(u => u.id !== userId)
  }
}
</script> 

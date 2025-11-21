<template>
  <div class="bg-white rounded-lg shadow p-6">
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-xl font-semibold text-gray-900">لیست اعضای وفاداری</h2>
      <div class="flex space-x-2 space-x-reverse">
        <input 
          v-model="searchQuery"
          type="text" 
          placeholder="جستجو در اعضا..."
          class="border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
        >
        <select 
          v-model="levelFilter"
          class="border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
        >
          <option value="">همه سطوح</option>
          <option value="برنزی">برنزی</option>
          <option value="نقره‌ای">نقره‌ای</option>
          <option value="طلایی">طلایی</option>
          <option value="الماس">الماس</option>
        </select>
      </div>
    </div>

    <!-- جدول اعضا -->
    <div class="overflow-x-auto">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              کاربر
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              سطح
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              امتیاز کل
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              امتیاز استفاده شده
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              تاریخ عضویت
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              آخرین فعالیت
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
          <tr v-for="member in filteredMembers" :key="member.id">
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="flex items-center">
                <div class="">
                  <img class="h-10 w-10 rounded-full" :src="member.avatar" :alt="member.name">
                </div>
                <div class="mr-4">
                  <div class="text-sm font-medium text-gray-900">{{ member.name }}</div>
                  <div class="text-sm text-gray-500">{{ member.email }}</div>
                </div>
              </div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="flex items-center">
                <div :class="getLevelColor(member.level)" class="w-8 h-8 rounded-full flex items-center justify-center">
                  <span class="text-white font-semibold text-sm">{{ getLevelIcon(member.level) }}</span>
                </div>
                <span class="mr-2 text-sm text-gray-900">{{ member.level }}</span>
              </div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="text-sm text-gray-900">{{ member.totalPoints.toLocaleString() }}</div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="text-sm text-gray-900">{{ member.usedPoints.toLocaleString() }}</div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="text-sm text-gray-900">{{ formatDate(member.joinDate) }}</div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="text-sm text-gray-900">{{ formatDate(member.lastActivity) }}</div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <span :class="member.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'" class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full">
                {{ member.status === 'active' ? 'فعال' : 'غیرفعال' }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
              <button 
                class="text-blue-600 hover:text-blue-900 ml-4"
                @click="$emit('show-member-details', member)"
              >
                جزئیات
              </button>
              <button 
                class="text-green-600 hover:text-green-900"
                @click="$emit('edit-member', member)"
              >
                ویرایش
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- صفحه‌بندی -->
    <div class="flex items-center justify-between mt-6">
      <div class="text-sm text-gray-700">
        نمایش {{ (currentPage - 1) * pageSize + 1 }} تا {{ Math.min(currentPage * pageSize, filteredMembers.length) }} از {{ filteredMembers.length }} عضو
      </div>
      <div class="flex space-x-2 space-x-reverse">
        <button 
          :disabled="currentPage === 1"
          class="px-3 py-2 border border-gray-300 rounded-md text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
          @click="currentPage--"
        >
          قبلی
        </button>
        <button 
          :disabled="currentPage >= totalPages"
          class="px-3 py-2 border border-gray-300 rounded-md text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
          @click="currentPage++"
        >
          بعدی
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'

// تعریف emit events
defineEmits(['show-member-details', 'edit-member'])

// اعضای وفاداری
const members = ref([
  {
    id: 1,
    name: 'علی احمدی',
    email: 'ali@example.com',
    avatar: '/default-avatar.png',
    level: 'الماس',
    totalPoints: 15000,
    usedPoints: 8500,
    joinDate: '2023-01-15',
    lastActivity: '2024-01-10',
    status: 'active'
  },
  {
    id: 2,
    name: 'فاطمه محمدی',
    email: 'fateme@example.com',
    avatar: '/default-avatar.png',
    level: 'طلایی',
    totalPoints: 8500,
    usedPoints: 4200,
    joinDate: '2023-03-20',
    lastActivity: '2024-01-08',
    status: 'active'
  },
  {
    id: 3,
    name: 'محمد رضایی',
    email: 'mohammad@example.com',
    avatar: '/default-avatar.png',
    level: 'نقره‌ای',
    totalPoints: 3200,
    usedPoints: 1800,
    joinDate: '2023-06-10',
    lastActivity: '2024-01-05',
    status: 'active'
  },
  {
    id: 4,
    name: 'زهرا کریمی',
    email: 'zahra@example.com',
    avatar: '/default-avatar.png',
    level: 'برنزی',
    totalPoints: 800,
    usedPoints: 300,
    joinDate: '2023-09-05',
    lastActivity: '2024-01-02',
    status: 'active'
  },
  {
    id: 5,
    name: 'حسین نوری',
    email: 'hossein@example.com',
    avatar: '/default-avatar.png',
    level: 'برنزی',
    totalPoints: 1200,
    usedPoints: 0,
    joinDate: '2023-11-15',
    lastActivity: '2023-12-20',
    status: 'inactive'
  }
])

// فیلترها و جستجو
const searchQuery = ref('')
const levelFilter = ref('')
const currentPage = ref(1)
const pageSize = 10

// فیلتر کردن اعضا
const filteredMembers = computed(() => {
  let filtered = members.value

  // فیلتر بر اساس جستجو
  if (searchQuery.value) {
    filtered = filtered.filter(member => 
      member.name.includes(searchQuery.value) || 
      member.email.includes(searchQuery.value)
    )
  }

  // فیلتر بر اساس سطح
  if (levelFilter.value) {
    filtered = filtered.filter(member => member.level === levelFilter.value)
  }

  return filtered
})

// محاسبه تعداد صفحات
const totalPages = computed(() => Math.ceil(filteredMembers.value.length / pageSize))

// رنگ‌های سطوح
function getLevelColor(level: string) {
  const colors = {
    'برنزی': 'bg-yellow-600',
    'نقره‌ای': 'bg-gray-500',
    'طلایی': 'bg-yellow-500',
    'الماس': 'bg-blue-600'
  }
  return colors[level] || 'bg-gray-400'
}

// آیکون‌های سطوح
function getLevelIcon(level: string) {
  const icons = {
    'برنزی': '🥉',
    'نقره‌ای': '🥈',
    'طلایی': '🥇',
    'الماس': '💎'
  }
  return icons[level] || '⭐'
}

// فرمت تاریخ
function formatDate(dateString: string) {
  const date = new Date(dateString)
  return date.toLocaleDateString('fa-IR')
}
</script> 

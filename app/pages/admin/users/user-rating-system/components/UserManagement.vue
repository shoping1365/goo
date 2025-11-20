<template>
  <div class="bg-white rounded-lg shadow">
    <div class="px-6 py-4 border-b border-gray-200">
      <div class="flex items-center justify-between">
        <h3 class="text-lg font-medium text-gray-900">مدیریت کاربران</h3>
        <div class="flex space-x-3 space-x-reverse">
          <input 
            v-model="searchQuery" 
            type="text" 
            placeholder="جستجو در کاربران..." 
            class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
          <select 
            v-model="filterLevel" 
            class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            <option value="">همه سطوح</option>
            <option value="bronze">برنزی</option>
            <option value="silver">نقره‌ای</option>
            <option value="gold">طلایی</option>
            <option value="platinum">پلاتینیوم</option>
            <option value="diamond">الماس</option>
          </select>
          <select 
            v-model="filterStatus" 
            class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            <option value="">همه وضعیت‌ها</option>
            <option value="active">فعال</option>
            <option value="blocked">مسدود</option>
            <option value="pending">در انتظار تایید</option>
          </select>
        </div>
      </div>
    </div>
    
    <div class="overflow-x-auto">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کاربر</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">امتیاز</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">سطح</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">آخرین فعالیت</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-for="user in filteredUsers" :key="user.id" class="hover:bg-gray-50">
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="flex items-center">
                <div class="">
                  <img class="h-10 w-10 rounded-full" :src="user.avatar || '/default-avatar.png'" :alt="user.name">
                </div>
                <div class="mr-4">
                  <div class="text-sm font-medium text-gray-900">{{ user.name }}</div>
                  <div class="text-sm text-gray-500">{{ user.email }}</div>
                </div>
              </div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="flex items-center">
                <span class="text-sm font-medium text-gray-900">{{ user.score }}</span>
                <div class="mr-2 flex items-center">
                  <svg
v-for="star in 5" :key="star" 
                       :class="star <= Math.floor(user.score / 200) ? 'text-yellow-400' : 'text-gray-300'" 
                       class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                    <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"></path>
                  </svg>
                </div>
              </div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <span :class="getLevelClass(user.level)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                {{ getLevelText(user.level) }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <span :class="getStatusClass(user.status)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                {{ getStatusText(user.status) }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
              {{ formatDate(user.lastActivity) }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
              <button class="text-blue-600 hover:text-blue-900 mr-3" @click="viewUserDetails(user)">جزئیات</button>
              <button class="text-green-600 hover:text-green-900 mr-3" @click="editUser(user)">ویرایش</button>
              <button :class="user.status === 'blocked' ? 'text-green-600 hover:text-green-900' : 'text-red-600 hover:text-red-900'" @click="toggleUserStatus(user)">
                {{ user.status === 'blocked' ? 'فعال‌سازی' : 'مسدودسازی' }}
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Pagination -->
    <div class="bg-white px-4 py-3 flex items-center justify-between border-t border-gray-200 sm:px-6">
      <div class="flex-1 flex justify-between sm:hidden">
        <button :disabled="currentPage === 1" class="relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50" @click="previousPage">
          قبلی
        </button>
        <button :disabled="currentPage === totalPages" class="mr-3 relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50" @click="nextPage">
          بعدی
        </button>
      </div>
      <div class="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
        <div>
          <p class="text-sm text-gray-700">
            نمایش <span class="font-medium">{{ paginationStart }}</span> تا <span class="font-medium">{{ paginationEnd }}</span> از <span class="font-medium">{{ totalUsers }}</span> نتیجه
          </p>
        </div>
        <div>
          <nav class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px">
            <button :disabled="currentPage === 1" class="relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50" @click="previousPage">
              <span class="sr-only">قبلی</span>
              <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd" />
              </svg>
            </button>
            <button
v-for="page in visiblePages" :key="page" 
                    :class="page === currentPage ? 'bg-blue-50 border-blue-500 text-blue-600' : 'bg-white border-gray-300 text-gray-500 hover:bg-gray-50'"
                    class="relative inline-flex items-center px-4 py-2 border text-sm font-medium"
                    @click="goToPage(page)">
              {{ page }}
            </button>
            <button :disabled="currentPage === totalPages" class="relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50" @click="nextPage">
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

<script setup lang="ts">
import { computed, ref } from 'vue';

// تعریف interface برای کاربر
interface User {
  id: number
  name: string
  email: string
  avatar?: string
  score: number
  level: string
  status: string
  lastActivity: string
}

// Props
const props = defineProps<{
  users: User[]
}>()

// Emits
const emit = defineEmits<{
  viewDetails: [user: User]
  editUser: [user: User]
  toggleStatus: [user: User]
}>()

// Reactive data
const searchQuery = ref('')
const filterLevel = ref('')
const filterStatus = ref('')
const currentPage = ref(1)
const perPage = ref(10)

// Computed properties
const filteredUsers = computed(() => {
  let filtered = props.users

  if (searchQuery.value) {
    filtered = filtered.filter(user => 
      user.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      user.email.toLowerCase().includes(searchQuery.value.toLowerCase())
    )
  }

  if (filterLevel.value) {
    filtered = filtered.filter(user => user.level === filterLevel.value)
  }

  if (filterStatus.value) {
    filtered = filtered.filter(user => user.status === filterStatus.value)
  }

  return filtered
})

const totalUsers = computed(() => filteredUsers.value.length)
const totalPages = computed(() => Math.ceil(totalUsers.value / perPage.value))
const paginationStart = computed(() => (currentPage.value - 1) * perPage.value + 1)
const paginationEnd = computed(() => Math.min(currentPage.value * perPage.value, totalUsers.value))

const visiblePages = computed(() => {
  const pages = []
  const start = Math.max(1, currentPage.value - 2)
  const end = Math.min(totalPages.value, currentPage.value + 2)
  
  for (let i = start; i <= end; i++) {
    pages.push(i)
  }
  
  return pages
})

// Methods
const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleDateString('fa-IR')
}

const getLevelClass = (level: string) => {
  switch (level) {
    case 'bronze':
      return 'bg-orange-100 text-orange-800'
    case 'silver':
      return 'bg-gray-100 text-gray-800'
    case 'gold':
      return 'bg-yellow-100 text-yellow-800'
    case 'platinum':
      return 'bg-blue-100 text-blue-800'
    case 'diamond':
      return 'bg-purple-100 text-purple-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getLevelText = (level: string) => {
  switch (level) {
    case 'bronze':
      return 'برنزی'
    case 'silver':
      return 'نقره‌ای'
    case 'gold':
      return 'طلایی'
    case 'platinum':
      return 'پلاتینیوم'
    case 'diamond':
      return 'الماس'
    default:
      return 'نامشخص'
  }
}

const getStatusClass = (status: string) => {
  switch (status) {
    case 'active':
      return 'bg-green-100 text-green-800'
    case 'blocked':
      return 'bg-red-100 text-red-800'
    case 'pending':
      return 'bg-yellow-100 text-yellow-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getStatusText = (status: string) => {
  switch (status) {
    case 'active':
      return 'فعال'
    case 'blocked':
      return 'مسدود'
    case 'pending':
      return 'در انتظار تایید'
    default:
      return 'نامشخص'
  }
}

const viewUserDetails = (user: User) => {
  emit('viewDetails', user)
}

const editUser = (user: User) => {
  emit('editUser', user)
}

const toggleUserStatus = (user: User) => {
  emit('toggleStatus', user)
}

const previousPage = () => {
  if (currentPage.value > 1) {
    currentPage.value--
  }
}

const nextPage = () => {
  if (currentPage.value < totalPages.value) {
    currentPage.value++
  }
}

const goToPage = (page: number) => {
  currentPage.value = page
}
</script> 

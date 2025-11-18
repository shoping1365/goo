<template>
  <div>
    <div class="mb-6">
      <h2 class="text-xl font-semibold text-gray-900 mb-2">مدیریت کاربران</h2>
      <p class="text-gray-600">مشاهده، مدیریت و کنترل کاربران سیستم</p>
    </div>

    <!-- Search and Filters -->
    <div class="bg-white p-6 rounded-lg shadow mb-6">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div>
          <input
            v-model="filters.search"
            type="text"
            placeholder="جستجو بر اساس نام، موبایل یا ایمیل..."
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        </div>
        <div>
          <select
            v-model="filters.role"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            <option value="">همه نقش‌ها</option>
            <option value="user">کاربر</option>
            <option value="staff">کارمند</option>
            <option value="admin">ادمین</option>
          </select>
        </div>
        <div>
          <select
            v-model="filters.status"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            <option value="">همه وضعیت‌ها</option>
            <option value="active">فعال</option>
            <option value="inactive">غیرفعال</option>
            <option value="blocked">بلاک شده</option>
          </select>
        </div>
        <div>
          <button
            @click="loadUsers"
            class="w-full px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700"
          >
            جستجو
          </button>
        </div>
      </div>
    </div>

    <!-- Users Table -->
    <div class="bg-white rounded-lg shadow overflow-hidden">
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                کاربر
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                نقش
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                وضعیت
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                آخرین ورود
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                تاریخ ثبت‌نام
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                عملیات
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="user in users" :key="user.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="">
                    <div class="h-10 w-10 rounded-full bg-gray-300 flex items-center justify-center">
                      <span class="text-sm font-medium text-gray-700">
                        {{ user.name ? user.name.charAt(0) : user.mobile.charAt(0) }}
                      </span>
                    </div>
                  </div>
                  <div class="mr-4">
                    <div class="text-sm font-medium text-gray-900">
                      {{ user.name || 'بدون نام' }}
                    </div>
                    <div class="text-sm text-gray-500">{{ user.mobile }}</div>
                    <div v-if="user.email" class="text-sm text-gray-500">{{ user.email }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span
                  :class="[
                    'px-2 inline-flex text-xs leading-5 font-semibold rounded-full',
                    user.role === 'admin' ? 'bg-red-100 text-red-800' :
                    user.role === 'staff' ? 'bg-yellow-100 text-yellow-800' :
                    'bg-green-100 text-green-800'
                  ]"
                >
                  {{ getRoleLabel(user.role) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex flex-col space-y-1">
                  <span
                    :class="[
                      'px-2 inline-flex text-xs leading-5 font-semibold rounded-full',
                      user.is_blocked ? 'bg-red-100 text-red-800' :
                      user.status === 'active' ? 'bg-green-100 text-green-800' :
                      'bg-gray-100 text-gray-800'
                    ]"
                  >
                    {{ user.is_blocked ? 'بلاک شده' : getStatusLabel(user.status) }}
                  </span>
                  <span v-if="user.is_blocked" class="text-xs text-gray-500">
                    {{ formatDate(user.blocked_at) }}
                  </span>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ user.last_login_at ? formatDate(user.last_login_at) : 'هرگز' }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ formatDate(user.created_at) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <div class="flex space-x-2">
                  <button
                    @click="viewUserDetails(user)"
                    class="text-blue-600 hover:text-blue-900"
                  >
                    جزئیات
                  </button>
                  <button
                    v-if="!user.is_blocked"
                    @click="blockUser(user)"
                    class="text-red-600 hover:text-red-900"
                  >
                    بلاک
                  </button>
                  <button
                    v-else
                    @click="unblockUser(user)"
                    class="text-green-600 hover:text-green-900"
                  >
                    آنبلاک
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Pagination -->
      <div class="bg-white px-4 py-3 flex items-center justify-between border-t border-gray-200 sm:px-6">
        <div class="flex-1 flex justify-between sm:hidden">
          <button
            @click="previousPage"
            :disabled="currentPage === 1"
            class="relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50"
          >
            قبلی
          </button>
          <button
            @click="nextPage"
            :disabled="currentPage >= totalPages"
            class="mr-3 relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50"
          >
            بعدی
          </button>
        </div>
        <div class="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
          <div>
            <p class="text-sm text-gray-700">
              نمایش
              <span class="font-medium">{{ (currentPage - 1) * pageSize + 1 }}</span>
              تا
              <span class="font-medium">{{ Math.min(currentPage * pageSize, total) }}</span>
              از
              <span class="font-medium">{{ total }}</span>
              نتیجه
            </p>
          </div>
          <div>
            <nav class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px">
              <button
                @click="previousPage"
                :disabled="currentPage === 1"
                class="relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50"
              >
                قبلی
              </button>
              <button
                v-for="page in visiblePages"
                :key="page"
                @click="goToPage(page)"
                :class="[
                  'relative inline-flex items-center px-4 py-2 border text-sm font-medium',
                  page === currentPage
                    ? 'z-10 bg-blue-50 border-blue-500 text-blue-600'
                    : 'bg-white border-gray-300 text-gray-500 hover:bg-gray-50'
                ]"
              >
                {{ page }}
              </button>
              <button
                @click="nextPage"
                :disabled="currentPage >= totalPages"
                class="relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50"
              >
                بعدی
              </button>
            </nav>
          </div>
        </div>
      </div>
    </div>

    <!-- User Details Modal -->
    <div v-if="selectedUser" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
      <div class="relative top-20 mx-auto p-5 border w-11/12 md:w-3/4 lg:w-1/2 shadow-lg rounded-md bg-white">
        <div class="mt-3">
          <div class="flex justify-between items-center mb-4">
            <h3 class="text-lg font-medium text-gray-900">جزئیات کاربر</h3>
            <button
              @click="selectedUser = null"
              class="text-gray-400 hover:text-gray-600"
            >
              ✕
            </button>
          </div>
          
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-6">
            <div>
              <label class="block text-sm font-medium text-gray-700">نام</label>
              <p class="mt-1 text-sm text-gray-900">{{ selectedUser.name || 'بدون نام' }}</p>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">موبایل</label>
              <p class="mt-1 text-sm text-gray-900">{{ selectedUser.mobile }}</p>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">ایمیل</label>
              <p class="mt-1 text-sm text-gray-900">{{ selectedUser.email || 'تنظیم نشده' }}</p>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">نقش</label>
              <p class="mt-1 text-sm text-gray-900">{{ getRoleLabel(selectedUser.role) }}</p>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">وضعیت</label>
              <p class="mt-1 text-sm text-gray-900">{{ getStatusLabel(selectedUser.status) }}</p>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">تاریخ ثبت‌نام</label>
              <p class="mt-1 text-sm text-gray-900">{{ formatDate(selectedUser.created_at) }}</p>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">آخرین ورود</label>
              <p class="mt-1 text-sm text-gray-900">{{ selectedUser.last_login_at ? formatDate(selectedUser.last_login_at) : 'هرگز' }}</p>
            </div>
            <div v-if="selectedUser.is_blocked">
              <label class="block text-sm font-medium text-gray-700">دلیل بلاک</label>
              <p class="mt-1 text-sm text-gray-900">{{ selectedUser.block_reason || 'تنظیم نشده' }}</p>
            </div>
          </div>

          <div class="flex justify-end space-x-3">
            <button
              @click="selectedUser = null"
              class="px-4 py-2 border border-gray-300 rounded-md text-gray-700 hover:bg-gray-50"
            >
              بستن
            </button>
            <button
              v-if="!selectedUser.is_blocked"
              @click="blockUser(selectedUser)"
              class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-700"
            >
              بلاک کاربر
            </button>
            <button
              v-else
              @click="unblockUser(selectedUser)"
              class="px-4 py-2 bg-green-600 text-white rounded-md hover:bg-green-700"
            >
              آنبلاک کاربر
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Block User Modal -->
    <div v-if="showBlockModal" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
      <div class="relative top-20 mx-auto p-5 border w-full max-w-md sm:max-w-lg md:max-w-xl shadow-lg rounded-md bg-white">
        <div class="mt-3">
          <h3 class="text-lg font-medium text-gray-900 mb-4">بلاک کردن کاربر</h3>
          <div class="mb-4">
            <label class="block text-sm font-medium text-gray-700 mb-2">دلیل بلاک کردن</label>
            <textarea
              v-model="blockReason"
              rows="3"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              placeholder="دلیل بلاک کردن کاربر را وارد کنید..."
            ></textarea>
          </div>
          <div class="flex justify-end space-x-3">
            <button
              @click="showBlockModal = false"
              class="px-4 py-2 border border-gray-300 rounded-md text-gray-700 hover:bg-gray-50"
            >
              انصراف
            </button>
            <button
              @click="confirmBlockUser"
              :disabled="!blockReason.trim()"
              class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-700 disabled:opacity-50"
            >
              بلاک کردن
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'

// Reactive data
const users = ref([])
const total = ref(0)
const currentPage = ref(1)
const pageSize = ref(10)
const selectedUser = ref(null)
const showBlockModal = ref(false)
const userToBlock = ref(null)
const blockReason = ref('')

const filters = reactive({
  search: '',
  role: '',
  status: ''
})

// Computed
const totalPages = computed(() => Math.ceil(total.value / pageSize.value))

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
const getRoleLabel = (role) => {
  const labels = {
    'admin': 'ادمین',
    'staff': 'کارمند',
    'user': 'کاربر'
  }
  return labels[role] || role
}

const getStatusLabel = (status) => {
  const labels = {
    'active': 'فعال',
    'inactive': 'غیرفعال',
    'blocked': 'بلاک شده'
  }
  return labels[status] || status
}

const formatDate = (date) => {
  if (!date) return ''
  return new Date(date).toLocaleDateString('fa-IR')
}

const loadUsers = async () => {
  try {
    const params = new URLSearchParams({
      page: currentPage.value,
      limit: pageSize.value,
      ...filters
    })

    const { data } = await $fetch(`/api/admin/users?${params}`)
    users.value = data.users
    total.value = data.total
  } catch (error) {
    console.error('خطا در بارگذاری کاربران:', error)
  }
}

const viewUserDetails = (user) => {
  selectedUser.value = user
}

const blockUser = (user) => {
  userToBlock.value = user
  showBlockModal.value = true
}

const confirmBlockUser = async () => {
  if (!blockReason.value.trim()) return

  try {
    await $fetch(`/api/admin/users/${userToBlock.value.id}/block`, {
      method: 'POST',
      body: { reason: blockReason.value }
    })
    
    showBlockModal.value = false
    blockReason.value = ''
    userToBlock.value = null
    
    // Reload users
    await loadUsers()
    
    alert('کاربر با موفقیت بلاک شد')
  } catch (error) {
    console.error('خطا در بلاک کردن کاربر:', error)
    alert('خطا در بلاک کردن کاربر')
  }
}

const unblockUser = async (user) => {
  if (!confirm('آیا از آنبلاک کردن این کاربر اطمینان دارید؟')) return

  try {
    await $fetch(`/api/admin/users/${user.id}/unblock`, {
      method: 'POST'
    })
    
    // Reload users
    await loadUsers()
    
    alert('کاربر با موفقیت آنبلاک شد')
  } catch (error) {
    console.error('خطا در آنبلاک کردن کاربر:', error)
    alert('خطا در آنبلاک کردن کاربر')
  }
}

const previousPage = () => {
  if (currentPage.value > 1) {
    currentPage.value--
    loadUsers()
  }
}

const nextPage = () => {
  if (currentPage.value < totalPages.value) {
    currentPage.value++
    loadUsers()
  }
}

const goToPage = (page) => {
  currentPage.value = page
  loadUsers()
}

// Lifecycle
onMounted(() => {
  loadUsers()
})
</script> 

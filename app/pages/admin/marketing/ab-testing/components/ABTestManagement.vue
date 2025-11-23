<template>
  <div v-if="hasAccess" class="mb-10">
    <!-- هدر و دکمه ایجاد تست جدید -->
    <div class="flex items-center justify-between mb-6">
      <h2 class="text-xl font-bold text-gray-900">مدیریت تست‌ها</h2>
      <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 flex items-center" @click="$emit('create-test')">
        <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
        </svg>
        ایجاد تست جدید
      </button>
    </div>

    <!-- فیلترها و جستجو -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6 mb-6">
      <div class="flex flex-wrap items-center gap-6">
        <!-- جستجو -->
        <div class="flex-1 min-w-64">
          <input
            v-model="searchQuery"
            type="text"
            placeholder="جستجو در نام تست‌ها، توضیحات یا نتایج..."
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
        </div>

        <!-- فیلتر نوع تست -->
        <select v-model="filterType" class="px-3 py-2 border border-gray-300 rounded-lg text-sm">
          <option value="">همه انواع</option>
          <option value="page">صفحه</option>
          <option value="button">دکمه</option>
          <option value="price">قیمت</option>
          <option value="image">تصویر</option>
          <option value="product">محصول</option>
        </select>

        <!-- فیلتر وضعیت -->
        <select v-model="filterStatus" class="px-3 py-2 border border-gray-300 rounded-lg text-sm">
          <option value="">همه وضعیت‌ها</option>
          <option value="active">فعال</option>
          <option value="inactive">غیرفعال</option>
          <option value="pending">در انتظار</option>
          <option value="completed">تکمیل شده</option>
        </select>

        <!-- فیلتر تاریخ -->
        <select v-model="filterDate" class="px-3 py-2 border border-gray-300 rounded-lg text-sm">
          <option value="">همه تاریخ‌ها</option>
          <option value="today">امروز</option>
          <option value="week">هفته اخیر</option>
          <option value="month">ماه اخیر</option>
        </select>

        <!-- فیلتر نتیجه -->
        <select v-model="filterResult" class="px-3 py-2 border border-gray-300 rounded-lg text-sm">
          <option value="">همه نتایج</option>
          <option value="winner_a">برنده A</option>
          <option value="winner_b">برنده B</option>
          <option value="no_winner">بدون برنده</option>
        </select>

        <!-- پاک کردن فیلترها -->
        <button class="px-3 py-2 text-sm text-gray-600 hover:text-gray-800 hover:bg-gray-100 rounded-lg" @click="clearFilters">
          پاک کردن فیلترها
        </button>
      </div>
    </div>

    <!-- جدول لیست تست‌ها -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 overflow-hidden">
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نام تست</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ شروع/پایان</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">شرکت‌کنندگان</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نرخ تبدیل</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="test in filteredTests" :key="test.id" class="hover:bg-gray-50">
              <!-- نام تست -->
              <td class="px-6 py-4 whitespace-nowrap">
                <div>
                  <div class="text-sm font-medium text-gray-900">{{ test.name }}</div>
                  <div class="text-sm text-gray-500">{{ test.description }}</div>
                </div>
              </td>

              <!-- نوع تست -->
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium" :class="getTypeClass(test.type)">
                  {{ getTypeLabel(test.type) }}
                </span>
              </td>

              <!-- وضعیت -->
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium" :class="getStatusClass(test.status)">
                  {{ getStatusLabel(test.status) }}
                </span>
              </td>

              <!-- تاریخ -->
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                <div>{{ formatDate(test.startDate) }}</div>
                <div class="text-gray-500">{{ formatDate(test.endDate) }}</div>
              </td>

              <!-- شرکت‌کنندگان -->
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ formatNumber(test.participants) }}
              </td>

              <!-- نرخ تبدیل -->
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm text-gray-900">
                  A: {{ test.conversionRateA }}% / B: {{ test.conversionRateB }}%
                </div>
                <div class="text-xs" :class="test.winner === 'A' ? 'text-blue-600' : test.winner === 'B' ? 'text-green-600' : 'text-gray-500'">
                  برنده: {{ test.winner || 'نامشخص' }}
                </div>
              </td>

              <!-- عملیات -->
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <div class="flex items-center space-x-3 space-x-reverse">
                  <!-- مشاهده نتایج -->
                  <button class="text-blue-600 hover:text-blue-900" title="مشاهده نتایج" @click="$emit('view-results', test.id)">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                    </svg>
                  </button>

                  <!-- ویرایش -->
                  <button class="text-green-600 hover:text-green-900" title="ویرایش" @click="$emit('edit-test', test.id)">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                    </svg>
                  </button>

                  <!-- فعال/غیرفعال -->
                  <button :class="test.status === 'active' ? 'text-yellow-600 hover:text-yellow-900' : 'text-green-600 hover:text-green-900'" :title="test.status === 'active' ? 'غیرفعال کردن' : 'فعال کردن'" @click="toggleTestStatus(test)">
                    <svg v-if="test.status === 'active'" class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 9v6m4-6v6m7-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                    <svg v-else class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.828 14.828a4 4 0 01-5.656 0M9 10h1m4 0h1m-6 4h1m4 0h1m2-7a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                  </button>

                  <!-- کپی -->
                  <button class="text-purple-600 hover:text-purple-900" title="کپی کردن" @click="$emit('duplicate-test', test.id)">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z" />
                    </svg>
                  </button>

                  <!-- حذف -->
                  <button 
                    v-if="canDeleteTest"
                    class="text-red-600 hover:text-red-900" 
                    title="حذف" 
                    @click="$emit('delete-test', test.id)"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                    </svg>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- پیام خالی -->
      <div v-if="filteredTests.length === 0" class="px-6 py-12 text-center">
        <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
        </svg>
        <h3 class="mt-2 text-sm font-medium text-gray-900">هیچ تستی یافت نشد</h3>
        <p class="mt-1 text-sm text-gray-500">تست جدیدی ایجاد کنید یا فیلترها را تغییر دهید.</p>
        <div class="mt-6">
          <button class="inline-flex items-center px-4 py-2 border border-transparent shadow-sm text-sm font-medium rounded-lg text-white bg-blue-600 hover:bg-blue-700" @click="$emit('create-test')">
            <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
            </svg>
            ایجاد تست جدید
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useAuth } from '@/composables/useAuth'
import { computed, onMounted, ref, watch } from 'vue'
import { useRouter } from 'vue-router'

const { user } = useAuth()
const router = useRouter()

const hasAccess = computed(() => {
  return user.value?.role === 'admin' || user.value?.role === 'developer'
})

watch(() => user.value, (newUser) => {
  if (!newUser || (newUser.role !== 'admin' && newUser.role !== 'developer')) {
    router.push('/404')
  }
})

onMounted(() => {
  if (!user.value || (user.value.role !== 'admin' && user.value.role !== 'developer')) {
    router.push('/404')
  }
})


// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { hasPermission } = useAuth()

// Computed برای چک کردن پرمیژن حذف
const canDeleteTest = computed(() => hasPermission('ab-test.delete'))

// Events
const _emit = defineEmits(['create-test', 'edit-test', 'view-results', 'duplicate-test', 'delete-test'])

// فیلترها
const searchQuery = ref('')
const filterType = ref('')
const filterStatus = ref('')
const filterDate = ref('')
const filterResult = ref('')

// داده‌های تست (mock data)
const tests = ref([
  {
    id: 1,
    name: 'تست دکمه خرید صفحه محصول',
    description: 'مقایسه متن و رنگ دکمه خرید',
    type: 'button',
    status: 'active',
    startDate: '2024-01-15',
    endDate: '2024-02-15',
    participants: 2450,
    conversionRateA: 12.3,
    conversionRateB: 15.8,
    winner: 'B'
  },
  {
    id: 2,
    name: 'تست قیمت محصولات',
    description: 'آزمایش قیمت‌گذاری مختلف',
    type: 'price',
    status: 'completed',
    startDate: '2024-01-01',
    endDate: '2024-01-31',
    participants: 1890,
    conversionRateA: 18.5,
    conversionRateB: 16.2,
    winner: 'A'
  },
  {
    id: 3,
    name: 'تست تصویر صفحه اصلی',
    description: 'مقایسه تصاویر مختلف هرو',
    type: 'image',
    status: 'pending',
    startDate: '2024-02-01',
    endDate: '2024-03-01',
    participants: 0,
    conversionRateA: 0,
    conversionRateB: 0,
    winner: null
  }
])

// فیلتر کردن تست‌ها
const filteredTests = computed(() => {
  let filtered = tests.value

  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    filtered = filtered.filter(test => 
      test.name.toLowerCase().includes(query) ||
      test.description.toLowerCase().includes(query)
    )
  }

  if (filterType.value) {
    filtered = filtered.filter(test => test.type === filterType.value)
  }

  if (filterStatus.value) {
    filtered = filtered.filter(test => test.status === filterStatus.value)
  }

  if (filterResult.value) {
    if (filterResult.value === 'winner_a') {
      filtered = filtered.filter(test => test.winner === 'A')
    } else if (filterResult.value === 'winner_b') {
      filtered = filtered.filter(test => test.winner === 'B')
    } else if (filterResult.value === 'no_winner') {
      filtered = filtered.filter(test => !test.winner)
    }
  }

  return filtered
})

// پاک کردن فیلترها
const clearFilters = () => {
  searchQuery.value = ''
  filterType.value = ''
  filterStatus.value = ''
  filterDate.value = ''
  filterResult.value = ''
}

interface Test {
  id?: number | string
  status?: string
  [key: string]: unknown
}

// فعال/غیرفعال کردن تست
const toggleTestStatus = (test: Test) => {
  test.status = test.status === 'active' ? 'inactive' : 'active'
}

// کلاس‌های نوع تست
const getTypeClass = (type: string) => {
  const classes = {
    page: 'bg-blue-100 text-blue-800',
    button: 'bg-green-100 text-green-800',
    price: 'bg-yellow-100 text-yellow-800',
    image: 'bg-purple-100 text-purple-800',
    product: 'bg-indigo-100 text-indigo-800'
  }
  return classes[type as keyof typeof classes] || 'bg-gray-100 text-gray-800'
}

// برچسب نوع تست
const getTypeLabel = (type: string) => {
  const labels = {
    page: 'صفحه',
    button: 'دکمه',
    price: 'قیمت',
    image: 'تصویر',
    product: 'محصول'
  }
  return labels[type as keyof typeof labels] || type
}

// کلاس‌های وضعیت
const getStatusClass = (status: string) => {
  const classes = {
    active: 'bg-green-100 text-green-800',
    inactive: 'bg-red-100 text-red-800',
    pending: 'bg-yellow-100 text-yellow-800',
    completed: 'bg-blue-100 text-blue-800'
  }
  return classes[status as keyof typeof classes] || 'bg-gray-100 text-gray-800'
}

// برچسب وضعیت
const getStatusLabel = (status: string) => {
  const labels = {
    active: 'فعال',
    inactive: 'غیرفعال',
    pending: 'در انتظار',
    completed: 'تکمیل شده'
  }
  return labels[status as keyof typeof labels] || status
}

// فرمت کردن تاریخ
const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleDateString('fa-IR')
}

// فرمت کردن اعداد
const formatNumber = (num: number) => {
  return new Intl.NumberFormat('fa-IR').format(num)
}
</script>

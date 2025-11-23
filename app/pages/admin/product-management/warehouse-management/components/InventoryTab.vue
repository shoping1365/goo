<template>
  <div v-if="hasAccess" class="space-y-6">
    <!-- هدر و ابزارها -->
    <div class="flex flex-col lg:flex-row lg:items-center lg:justify-between gap-6">
      <div>
        <h2 class="text-2xl font-bold text-gray-900">موجودی کالاها</h2>
        <p class="text-gray-600">مدیریت موجودی کالاها در این انبار</p>
      </div>
      
      <div class="flex flex-col sm:flex-row gap-3">
        <!-- جستجو -->
        <div class="relative">
          <input 
            v-model="searchQuery"
            type="text" 
            placeholder="جستجو در کالاها..." 
            class="w-64 bg-white border border-gray-300 rounded-xl pl-10 pr-4 py-2 text-gray-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          >
          <svg class="absolute left-3 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
        </div>
        
        <!-- فیلترها -->
        <div class="flex gap-2">
          <select v-model="categoryFilter" class="bg-white border border-gray-300 rounded-xl px-3 py-2 text-gray-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent">
            <option value="">همه دسته‌ها</option>
            <option value="الکترونیک">الکترونیک</option>
            <option value="پوشاک">پوشاک</option>
            <option value="کتاب">کتاب</option>
          </select>
          
          <select v-model="stockFilter" class="bg-white border border-gray-300 rounded-xl px-3 py-2 text-gray-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent">
            <option value="">همه موجودی‌ها</option>
            <option value="موجود">موجود</option>
            <option value="کم موجود">کم موجود</option>
            <option value="ناموجود">ناموجود</option>
          </select>
        </div>
        
        <!-- افزودن کالا -->
        <button class="bg-gradient-to-r from-blue-600 to-blue-700 text-white px-6 py-2 rounded-xl hover:from-blue-700 hover:to-blue-800 transition-all duration-200 shadow-lg hover:shadow-xl flex items-center gap-2">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
          </svg>
          افزودن کالا به این انبار
        </button>
      </div>
    </div>

    <!-- جدول موجودی -->
    <div class="bg-white rounded-2xl shadow-lg overflow-hidden">
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-gradient-to-r from-blue-50 to-indigo-50">
            <tr>
              <th class="px-6 py-4 text-right text-sm font-bold text-gray-800">نام کالا</th>
              <th class="px-6 py-4 text-right text-sm font-bold text-gray-800">کد کالا</th>
              <th class="px-6 py-4 text-right text-sm font-bold text-gray-800">موجودی</th>
              <th class="px-6 py-4 text-right text-sm font-bold text-gray-800">حداقل موجودی</th>
              <th class="px-6 py-4 text-right text-sm font-bold text-gray-800">تاریخ انقضا</th>
              <th class="px-6 py-4 text-right text-sm font-bold text-gray-800">محل قفسه</th>
              <th class="px-6 py-4 text-right text-sm font-bold text-gray-800">عملیات</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="item in filteredItems" :key="item.id" class="hover:bg-gray-50 transition-colors">
              <td class="px-6 py-4">
                <div class="flex items-center gap-3">
                  <div class="w-10 h-10 bg-gradient-to-r from-blue-500 to-indigo-600 rounded-lg flex items-center justify-center">
                    <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
                    </svg>
                  </div>
                  <div>
                    <div class="font-medium text-gray-900">{{ item.name }}</div>
                    <div class="text-sm text-gray-500">{{ item.category }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4">
                <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800">
                  {{ item.sku }}
                </span>
              </td>
              <td class="px-6 py-4">
                <div class="text-center">
                  <div class="text-lg font-bold" :class="getStockColor(item.stock)">{{ item.stock }}</div>
                  <div class="text-xs text-gray-500">عدد</div>
                </div>
              </td>
              <td class="px-6 py-4">
                <div class="text-center">
                  <div class="text-sm font-medium text-gray-900">{{ item.minStock }}</div>
                  <div class="text-xs text-gray-500">حداقل</div>
                </div>
              </td>
              <td class="px-6 py-4">
                <div class="text-center">
                  <div class="text-sm font-medium" :class="getExpiryColor(item.expiryDate)">{{ formatDate(item.expiryDate) }}</div>
                  <div class="text-xs text-gray-500">تاریخ انقضا</div>
                </div>
              </td>
              <td class="px-6 py-4">
                <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-gray-100 text-gray-800">
                  {{ item.location }}
                </span>
              </td>
              <td class="px-6 py-4">
                <div class="flex items-center gap-2">
                  <button class="p-1.5 text-blue-600 hover:bg-blue-50 rounded-lg transition-colors" title="مشاهده">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                    </svg>
                  </button>
                  <button class="p-1.5 text-green-600 hover:bg-green-50 rounded-lg transition-colors" title="ویرایش">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                    </svg>
                  </button>
                  <button class="p-1.5 text-orange-600 hover:bg-orange-50 rounded-lg transition-colors" title="تنظیم موجودی">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6V4m0 2a2 2 0 100 4m0-4a2 2 0 100 4m-6 8a2 2 0 100-4m0 4a2 2 0 100 4m0-4v2m0-6V4m6 6v10m6-2a2 2 0 100-4m0 4a2 2 0 100 4m0-4v2m0-6V4" />
                    </svg>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      
      <!-- آمار پایین جدول -->
      <div class="bg-gray-50 px-6 py-4 border-t border-gray-200">
        <div class="flex items-center justify-between text-sm text-gray-600">
          <span>نمایش {{ filteredItems.length }} از {{ items.length }} کالا</span>
          <div class="flex items-center gap-6">
            <span>کل موجودی: {{ totalStock }}</span>
            <span>کالاهای کم موجود: {{ lowStockCount }}</span>
            <span>کالاهای منقضی: {{ expiredCount }}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue';
import { useAuth } from '~/composables/useAuth';

declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>;

// احراز هویت
const { user, isAuthenticated } = useAuth()

// بررسی دسترسی admin
const hasAccess = computed(() => {
  if (!isAuthenticated.value) {
    return false
  }

  const userRole = user.value?.role?.toLowerCase() || ''
  const adminRoles = ['admin', 'developer']
  return adminRoles.includes(userRole)
})

// بررسی احراز هویت و دسترسی admin - نمایش 404 در صورت عدم دسترسی
const checkAuth = async () => {
  if (!hasAccess.value) {
    await navigateTo('/404', { external: false })
  }
}

// بررسی احراز هویت در هنگام mount
onMounted(async () => {
  await checkAuth()
})

// بررسی احراز هویت هنگام تغییر وضعیت احراز هویت
watch([isAuthenticated, hasAccess], async () => {
  if (!hasAccess.value) {
    await checkAuth()
  }
})

const searchQuery = ref('')

// Props
interface Props {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  warehouse: any
}
defineProps<Props>()

// داده‌های نمونه موجودی
const items = ref([
  {
    id: 1,
    name: 'لپ‌تاپ Dell XPS 13',
    sku: 'LAP-DELL-001',
    category: 'الکترونیک',
    stock: 45,
    minStock: 10,
    expiryDate: '2025-12-31',
    location: 'A-01-01'
  },
  {
    id: 2,
    name: 'کتاب React.js در عمل',
    sku: 'BOOK-REACT-001',
    category: 'کتاب',
    stock: 120,
    minStock: 20,
    expiryDate: '2026-06-30',
    location: 'B-02-03'
  },
  {
    id: 3,
    name: 'تی‌شرت نخی مردانه',
    sku: 'TSHIRT-M-001',
    category: 'پوشاک',
    stock: 8,
    minStock: 15,
    expiryDate: '2024-12-31',
    location: 'C-03-02'
  },
  {
    id: 4,
    name: 'ماوس بی‌سیم Logitech',
    sku: 'MOUSE-LOG-001',
    category: 'الکترونیک',
    stock: 67,
    minStock: 25,
    expiryDate: '2025-08-31',
    location: 'A-01-05'
  }
])

// فیلترها
const categoryFilter = ref('')
const stockFilter = ref('')

// محاسبه کالاهای فیلتر شده
const filteredItems = computed(() => {
  return items.value.filter(item => {
    const matchesSearch = item.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
                         item.sku.toLowerCase().includes(searchQuery.value.toLowerCase())
    
    const matchesCategory = !categoryFilter.value || item.category === categoryFilter.value
    const matchesStock = !stockFilter.value || getStockStatus(item.stock, item.minStock) === stockFilter.value
    
    return matchesSearch && matchesCategory && matchesStock
  })
})

// آمار
const totalStock = computed(() => 
  items.value.reduce((sum, item) => sum + item.stock, 0)
)

const lowStockCount = computed(() => 
  items.value.filter(item => item.stock <= item.minStock).length
)

const expiredCount = computed(() => 
  items.value.filter(item => new Date(item.expiryDate) < new Date()).length
)

// توابع کمکی
const getStockColor = (stock: number) => {
  if (stock === 0) return 'text-red-600'
  if (stock < 10) return 'text-orange-600'
  return 'text-green-600'
}

const getStockStatus = (stock: number, min: number) => {
  if (stock === 0) return 'ناموجود'
  if (stock <= min) return 'کم موجود'
  return 'موجود'
}

const getExpiryColor = (date: string) => {
  if (!date) return ''
  const d = new Date(date)
  const now = new Date()
  if (d < now) return 'text-red-600'
  const diffTime = d.getTime() - now.getTime();
  const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24)); 
  if (diffDays < 30) return 'text-orange-600'
  return 'text-gray-600'
}

const formatDate = (date: string) => {
  if (!date) return '-'
  return new Date(date).toLocaleDateString('fa-IR')
}
</script>

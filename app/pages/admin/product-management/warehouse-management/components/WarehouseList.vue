<template>
  <div class="space-y-6">
    <!-- هدر و ابزارها -->
    <div class="bg-gradient-to-br from-blue-50 via-indigo-50 to-purple-50 rounded-3xl shadow-xl p-8 border border-blue-100">
      <div class="flex flex-col lg:flex-row lg:items-center lg:justify-between gap-6">
        <div class="space-y-2">
          <div class="flex items-center gap-3">
            <div class="w-12 h-12 bg-gradient-to-br from-blue-600 to-indigo-700 rounded-2xl flex items-center justify-center shadow-lg">
              <svg class="w-7 h-7 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
              </svg>
            </div>
            <div>
              <h2 class="text-3xl font-bold bg-gradient-to-r from-blue-600 to-indigo-700 bg-clip-text text-transparent">لیست انبارها</h2>
              <p class="text-gray-600 text-lg">مدیریت و مشاهده تمام انبارهای ثبت‌شده</p>
            </div>
          </div>
        </div>
        
        <div class="flex flex-col sm:flex-row gap-6">
          <!-- جستجو -->
          <div class="relative group">
            <div class="absolute inset-0 bg-gradient-to-r from-blue-500 to-indigo-600 rounded-2xl blur opacity-20 group-hover:opacity-30 transition-opacity"></div>
            <div class="relative">
              <input 
                v-model="searchQuery"
                type="text" 
                placeholder="جستجو در انبارها..." 
                class="w-72 bg-white/80 backdrop-blur-sm border-0 rounded-2xl pl-12 pr-6 py-3 text-gray-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:bg-white shadow-lg"
              >
              <svg class="absolute left-4 top-1/2 transform -translate-y-1/2 w-6 h-6 text-blue-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
            </div>
          </div>
          
          <!-- فیلترها -->
          <div class="flex gap-3">
            <div class="relative group">
              <div class="absolute inset-0 bg-gradient-to-r from-green-500 to-emerald-600 rounded-xl blur opacity-20 group-hover:opacity-30 transition-opacity"></div>
              <select v-model="cityFilter" class="relative bg-white/80 backdrop-blur-sm border-0 rounded-xl px-4 py-3 text-gray-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:bg-white shadow-lg min-w-[140px]">
                <option value="">🌍 همه شهرها</option>
                <option value="تهران">🏛️ تهران</option>
                <option value="اصفهان">🕌 اصفهان</option>
                <option value="شیراز">🌹 شیراز</option>
                <option value="مشهد">🕋 مشهد</option>
              </select>
            </div>
            
            <div class="relative group">
              <div class="absolute inset-0 bg-gradient-to-r from-purple-500 to-pink-600 rounded-xl blur opacity-20 group-hover:opacity-30 transition-opacity"></div>
              <select v-model="statusFilter" class="relative bg-white/80 backdrop-blur-sm border-0 rounded-xl px-4 py-3 text-gray-700 focus:outline-none focus:ring-2 focus:ring-purple-500 focus:bg-white shadow-lg min-w-[160px]">
                <option value="">📊 همه وضعیت‌ها</option>
                <option value="فعال">✅ فعال</option>
                <option value="غیرفعال">❌ غیرفعال</option>
              </select>
            </div>
          </div>
          
          <!-- افزودن انبار جدید -->
          <div class="relative group">
            <div class="absolute inset-0 bg-gradient-to-r from-blue-600 to-indigo-700 rounded-2xl blur opacity-30 group-hover:opacity-50 transition-opacity"></div>
            <button :disabled="creating" class="relative bg-gradient-to-r from-blue-600 to-indigo-700 text-white px-8 py-3 rounded-2xl hover:from-blue-700 hover:to-indigo-800 transition-all duration-300 shadow-xl hover:shadow-2xl hover:scale-105 flex items-center gap-3 font-semibold" @click="onAddClick">
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
              </svg>
              {{ showCreate ? (creating ? 'در حال ثبت...' : 'ثبت انبار') : 'افزودن انبار جدید' }}
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- فرم ایجاد انبار جدید -->
    <div v-if="showCreate">
      <WarehouseForm v-model="newWarehouse" />
    </div>

    <!-- جدول انبارها -->
    <div class="bg-white/80 backdrop-blur-sm rounded-3xl shadow-2xl overflow-hidden border border-white/20">
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-gradient-to-r from-blue-50 to-indigo-50">
            <tr>
              <th class="px-8 py-6 text-right text-sm font-bold text-gray-800 border-b border-blue-100">🏢 نام انبار</th>
              <th class="px-8 py-6 text-right text-sm font-bold text-gray-800 border-b border-blue-100">🏷️ کد انبار</th>
              <th class="px-8 py-6 text-right text-sm font-bold text-gray-800 border-b border-blue-100">📍 موقعیت</th>
              <th class="px-8 py-6 text-right text-sm font-bold text-gray-800 border-b border-blue-100">👤 مدیر انبار</th>
              <th class="px-8 py-6 text-right text-sm font-bold text-gray-800 border-b border-blue-100">📦 تعداد SKU</th>
              <th class="px-8 py-6 text-right text-sm font-bold text-gray-800 border-b border-blue-100">📊 ظرفیت</th>
              <th class="px-8 py-6 text-right text-sm font-bold text-gray-800 border-b border-blue-100">🔄 وضعیت</th>
              <th class="px-8 py-6 text-right text-sm font-bold text-gray-800 border-b border-blue-100">⚙️ عملیات</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-blue-50">
            <tr v-for="warehouse in filteredWarehouses" :key="warehouse.id" class="hover:bg-gradient-to-r hover:from-blue-50/50 hover:to-indigo-50/50 transition-all duration-300 group">
              <td class="px-8 py-6">
                <div class="flex items-center gap-6">
                  <div class="w-12 h-12 bg-gradient-to-br from-blue-500 to-indigo-600 rounded-2xl flex items-center justify-center shadow-lg group-hover:scale-110 transition-transform duration-300">
                    <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
                    </svg>
                  </div>
                  <div>
                    <div class="font-bold text-lg text-gray-900 group-hover:text-blue-600 transition-colors">{{ warehouse.name }}</div>
                    <div class="text-sm text-gray-500 bg-gray-100 px-3 py-1 rounded-full inline-block">{{ warehouse.type }}</div>
                  </div>
                </div>
              </td>
              <td class="px-8 py-6">
                <span class="inline-flex items-center px-4 py-2 rounded-2xl text-sm font-bold bg-gradient-to-r from-blue-100 to-indigo-100 text-blue-800 border border-blue-200 shadow-sm group-hover:shadow-md transition-all duration-300">
                  <svg class="w-4 h-4 mr-2 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z" />
                  </svg>
                  {{ warehouse.code }}
                </span>
              </td>
              <td class="px-8 py-6">
                <div class="space-y-2">
                  <div class="flex items-center gap-2">
                    <svg class="w-4 h-4 text-blue-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
                    </svg>
                    <span class="text-sm font-bold text-gray-900">{{ warehouse.city }}</span>
                  </div>
                  <div class="text-xs text-gray-500 bg-gray-50 px-3 py-1 rounded-lg max-w-[200px] truncate">{{ warehouse.address }}</div>
                </div>
              </td>
              <td class="px-8 py-6">
                <div class="flex items-center gap-3">
                  <div class="w-10 h-10 bg-gradient-to-br from-green-500 to-blue-600 rounded-2xl flex items-center justify-center shadow-lg group-hover:scale-110 transition-transform duration-300">
                    <span class="text-white text-sm font-bold">{{ warehouse.manager.initials }}</span>
                  </div>
                  <div class="space-y-1">
                    <div class="text-sm font-bold text-gray-900 group-hover:text-green-600 transition-colors">{{ warehouse.manager.name }}</div>
                    <div class="text-xs text-gray-500 bg-gray-50 px-2 py-1 rounded-lg">{{ warehouse.manager.phone }}</div>
                  </div>
                </div>
              </td>
              <td class="px-8 py-6">
                <div class="text-center">
                  <div class="text-2xl font-bold text-gray-900 group-hover:text-blue-600 transition-colors">{{ warehouse.skuCount }}</div>
                  <div class="text-xs text-gray-500 bg-blue-50 px-3 py-1 rounded-full inline-block">SKU فعال</div>
                </div>
              </td>
              <td class="px-8 py-6">
                <div class="w-36">
                  <div class="flex items-center justify-between text-sm mb-2">
                    <span class="text-gray-600 font-medium">{{ warehouse.capacity.used }}</span>
                    <span class="text-gray-900 font-bold">{{ warehouse.capacity.total }}</span>
                  </div>
                  <div class="w-full bg-gray-200 rounded-full h-3 shadow-inner">
                    <div 
                      class="bg-gradient-to-r from-blue-500 to-indigo-600 h-3 rounded-full transition-all duration-500 shadow-lg" 
                      :style="{ width: `${warehouse.capacity.total > 0 ? (warehouse.capacity.used / warehouse.capacity.total) * 100 : 0}%` }"
                    ></div>
                  </div>
                  <div class="text-xs text-gray-500 mt-2 bg-gray-50 px-2 py-1 rounded-lg text-center">
                    {{ warehouse.capacity.total > 0 ? Math.round((warehouse.capacity.used / warehouse.capacity.total) * 100) : 0 }}% پر شده
                  </div>
                </div>
              </td>
              <td class="px-6 py-4">
                <span 
                  :class="[
                    'inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium',
                    warehouse.status === 'فعال' 
                      ? 'bg-green-100 text-green-800' 
                      : 'bg-red-100 text-red-800'
                  ]"
                >
                  <span 
                    :class="[
                      'w-2 h-2 rounded-full mr-1.5',
                      warehouse.status === 'فعال' ? 'bg-green-400' : 'bg-red-400'
                    ]"
                  ></span>
                  {{ warehouse.status }}
                </span>
              </td>
              <td class="px-6 py-4">
                <div class="flex items-center gap-2">
                  <NuxtLink 
                    :to="`/admin/product-management/warehouse-management/${warehouse.id}`"
                    class="p-1.5 text-blue-600 hover:bg-blue-50 rounded-lg transition-colors" 
                    title="مشاهده جزئیات"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                    </svg>
                  </NuxtLink>
                  <button class="p-1.5 text-green-600 hover:bg-green-50 rounded-lg transition-colors" title="ویرایش">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                    </svg>
                  </button>
                  <button class="p-1.5 text-red-600 hover:bg-red-50 rounded-lg transition-colors" title="حذف" @click="onRemove(warehouse)">
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
      
             <!-- آمار پایین جدول -->
       <div class="bg-gradient-to-r from-blue-50 to-indigo-50 px-8 py-6 border-t border-blue-100">
         <div class="flex flex-col lg:flex-row lg:items-center lg:justify-between gap-6">
           <div class="flex items-center gap-2 text-gray-700">
              <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
              </svg>
              <span class="font-medium">نمایش {{ filteredWarehouses.length }} از {{ warehouses.length }} انبار</span>
            </div>
            <div class="flex flex-wrap items-center gap-6">
              <div class="flex items-center gap-2 bg-white/60 px-4 py-2 rounded-xl shadow-sm">
                <div class="w-3 h-3 bg-green-500 rounded-full"></div>
                <span class="text-sm font-medium text-gray-700">انبارهای فعال: {{ activeWarehousesCount }}</span>
              </div>
              <div class="flex items-center gap-2 bg-white/60 px-4 py-2 rounded-xl shadow-sm">
                <div class="w-3 h-3 bg-blue-500 rounded-full"></div>
                <span class="text-sm font-medium text-gray-700">کل ظرفیت: {{ totalCapacity }}</span>
              </div>
              <div class="flex items-center gap-2 bg-white/60 px-4 py-2 rounded-xl shadow-sm">
                <div class="w-3 h-3 bg-purple-500 rounded-full"></div>
                <span class="text-sm font-medium text-gray-700">میانگین پر شدن: {{ averageCapacityUsage }}%</span>
              </div>
            </div>
          </div>
        </div>
    </div>

    <!-- نقشه تعاملی -->
    <div class="bg-white/80 backdrop-blur-sm rounded-3xl shadow-2xl p-8 border border-white/20">
      <h3 class="text-2xl font-bold text-gray-900 mb-6 flex items-center gap-3">
        <div class="w-10 h-10 bg-gradient-to-br from-indigo-500 to-purple-600 rounded-2xl flex items-center justify-center shadow-lg">
          <svg class="h-6 w-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 20l-5.447-2.724A1 1 0 013 16.382V5.618a1 1 0 011.447-.894L9 7m0 13l6-3m-6 3V7m6 10l4.553 2.276A1 1 0 0021 18.382V7.618a1 1 0 00-1.447-.894L15 4m0 13V4m0 0L9 7" />
          </svg>
        </div>
        نقشه تعاملی موقعیت انبارها
      </h3>
      
      <div class="bg-gradient-to-br from-indigo-50 via-purple-50 to-pink-50 rounded-3xl p-8 h-96 flex items-center justify-center border border-indigo-100">
        <div class="text-center">
          <div class="w-24 h-24 bg-gradient-to-br from-indigo-400 to-purple-500 rounded-3xl flex items-center justify-center mx-auto mb-6 shadow-2xl">
            <svg class="h-12 w-12 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 20l-5.447-2.724A1 1 0 013 16.382V5.618a1 1 0 011.447-.894L9 7m0 13l6-3m-6 3V7m6 10l4.553 2.276A1 1 0 0021 18.382V7.618a1 1 0 00-1.447-.894L15 4m0 13V4m0 0L9 7" />
            </svg>
          </div>
          <h4 class="text-xl font-bold text-gray-900 mb-3">نقشه تعاملی انبارها</h4>
          <p class="text-gray-600 mb-6">کلیک روی هر نقطه برای مشاهده جزئیات انبار</p>
          
          <!-- نقاط انبارها روی نقشه -->
          <div class="flex items-center justify-center gap-12 mb-8">
            <div v-for="warehouse in warehouses" :key="warehouse.id" class="text-center group cursor-pointer">
              <div 
                class="w-6 h-6 bg-gradient-to-br from-blue-500 to-indigo-600 rounded-full mx-auto mb-3 shadow-lg group-hover:scale-125 group-hover:shadow-2xl transition-all duration-300"
                :title="warehouse.name"
              ></div>
              <div class="text-sm font-medium text-gray-700 group-hover:text-blue-600 transition-colors">{{ warehouse.city }}</div>
              <div class="text-xs text-gray-500">{{ warehouse.skuCount }} SKU</div>
            </div>
          </div>
          
          <div class="bg-white/60 backdrop-blur-sm rounded-2xl p-6 shadow-lg">
            <p class="text-sm font-medium text-gray-700">
              <!-- آمار شهرها به‌صورت استاتیک نمایش داده می‌شود -->
              🏛️ تهران | 🕌 اصفهان | 🌹 شیراز | 🕋 مشهد
            </p>
          </div>
        </div>
      </div>
    </div>

    <!-- Dialog و Notification سفارشی -->
    <ConfirmDialog ref="confirmDialogRef" />
    <ChatNotification :notification="notification" />
  </div>
</template>

<script lang="ts">
declare const useFetch: <T>(url: string, options?: unknown) => Promise<{ data: { value: T }; refresh: () => Promise<void> }>
</script>

<script setup lang="ts">
import ConfirmDialog from '@/components/common/ConfirmDialog.vue'
import ChatNotification from '@/pages/admin/support-management/live-chat/components/ChatNotification.vue'
import { computed, ref } from 'vue'
import WarehouseForm from './WarehouseForm.vue'

type ApiWarehouse = {
  id: number
  code: string
  name: string
  city?: string
  address?: string
  is_active: boolean
  is_default: boolean
  priority: number
}

// بارگذاری لیست انبارها از API
const { data: apiWarehouses, refresh } = await useFetch<ApiWarehouse[]>(
  '/api/admin/warehouses',
  { key: 'admin-warehouses-list', defaultCache: true, swr: true }
)

// نرمال‌سازی داده‌ها برای نمایش در UI فعلی بدون تغییر استایل‌ها
const warehouses = computed(() => {
  const list = apiWarehouses.value || []
  return list.map(w => ({
    id: w.id,
    name: w.name,
    code: w.code,
    type: 'انبار',
    city: w.city || '-',
    address: w.address || '-',
    manager: { name: '-', phone: '-', initials: '-' },
    skuCount: 0,
    capacity: { used: 0, total: 0 },
    status: w.is_active ? 'فعال' : 'غیرفعال'
  }))
})

// فیلترها
const searchQuery = ref('')
const cityFilter = ref('')
const statusFilter = ref('')

// محاسبه انبارهای فیلتر شده
const filteredWarehouses = computed(() => {
  return warehouses.value.filter(warehouse => {
    const q = searchQuery.value.toLowerCase()
    const name = String(warehouse.name || '').toLowerCase()
    const code = String(warehouse.code || '').toLowerCase()
    const city = String(warehouse.city || '').toLowerCase()

    const matchesSearch = name.includes(q) || code.includes(q) || city.includes(q)
    const matchesCity = !cityFilter.value || warehouse.city === cityFilter.value
    const matchesStatus = !statusFilter.value || warehouse.status === statusFilter.value
    return matchesSearch && matchesCity && matchesStatus
  })
})

// آمار
const activeWarehousesCount = computed(() => 
  warehouses.value.filter(w => w.status === 'فعال').length
)

const totalCapacity = computed(() => 
  warehouses.value.reduce((sum, w) => sum + (w.capacity?.total || 0), 0)
)

const averageCapacityUsage = computed(() => {
  const totals = warehouses.value
  const totalUsed = totals.reduce((sum, w) => sum + (w.capacity?.used || 0), 0)
  const totalCap = totals.reduce((sum, w) => sum + (w.capacity?.total || 0), 0)
  return totalCap > 0 ? Math.round((totalUsed / totalCap) * 100) : 0
})

// ایجاد انبار جدید
const showCreate = ref(false)
const creating = ref(false)
const notification = ref<{ message: string; type: 'success' | 'info' | 'warning' | 'error' } | null>(null)
const newWarehouse = ref({ code: '', name: '', city: '', address: '', is_default: false, is_active: true, priority: 100 })
const confirmDialogRef = ref<InstanceType<typeof ConfirmDialog> | null>(null)

async function onAddClick() {
  if (!showCreate.value) {
    showCreate.value = true
    return
  }
  if (!newWarehouse.value.code || !newWarehouse.value.name) return
  try {
    creating.value = true
    const created = await $fetch<ApiWarehouse>('/api/admin/warehouses', { method: 'POST', body: newWarehouse.value })
    // به‌روزرسانی فوری لیست در حافظه برای نمایش بلافاصله
    apiWarehouses.value = [ ...(apiWarehouses.value || []), created ]
    showCreate.value = false
    newWarehouse.value = { code: '', name: '', city: '', address: '', is_default: false, is_active: true, priority: 100 }
    await refresh()
    notification.value = { message: 'انبار با موفقیت ایجاد شد', type: 'success' }
    setTimeout(() => (notification.value = null), 2500)
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
  } catch (e: any) {
    notification.value = { message: e?.data?.message || e?.statusMessage || 'خطا در ثبت انبار', type: 'error' }
    setTimeout(() => (notification.value = null), 3000)
  } finally {
    creating.value = false
  }
}

// حذف انبار
async function onRemove(w: { id: number }) {
  if (!w?.id) return
  const ok = await confirmDialogRef.value?.show({
    title: 'تایید حذف',
    message: 'آیا از حذف این انبار مطمئن هستید؟ این عملیات غیرقابل بازگشت است.',
    confirmText: 'حذف',
    cancelText: 'انصراف',
    type: 'danger'
  })
  if (!ok) return
  try {
    await $fetch(`/api/admin/warehouses/${w.id}`, { method: 'DELETE' })
    // حذف خوش‌بینانه از لیست محلی
    apiWarehouses.value = (apiWarehouses.value || []).filter(x => x.id !== w.id)
    await refresh()
    notification.value = { message: 'انبار با موفقیت حذف شد', type: 'success' }
    setTimeout(() => (notification.value = null), 2500)
  } catch (e) {
    notification.value = { message: 'خطا در حذف انبار', type: 'error' }
    setTimeout(() => (notification.value = null), 3000)
  }
}
</script>

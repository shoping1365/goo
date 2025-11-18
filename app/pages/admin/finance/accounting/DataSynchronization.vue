<template>
  <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
    <!-- هدر -->
    <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-6 mb-6">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">همگام‌سازی داده‌ها</h3>
        <p class="text-sm text-gray-600">مدیریت همگام‌سازی خودکار و دستی داده‌ها با نرم‌افزارهای حسابداری</p>
      </div>
      
      <!-- دکمه‌های عملیاتی -->
      <div class="flex flex-wrap gap-3">
        <button 
          @click="startManualSync"
          :disabled="isSyncing"
          class="inline-flex items-center gap-2 px-4 py-2 bg-blue-600 hover:bg-blue-700 disabled:bg-gray-400 text-white rounded-lg transition-colors duration-200"
        >
          <svg v-if="isSyncing" class="w-4 h-4 animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
          </svg>
          <svg v-else class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
          </svg>
          {{ isSyncing ? 'در حال همگام‌سازی...' : 'شروع همگام‌سازی' }}
        </button>
        
        <button 
          @click="showSyncSettings = true"
          class="inline-flex items-center gap-2 px-4 py-2 bg-gray-600 hover:bg-gray-700 text-white rounded-lg transition-colors duration-200"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"></path>
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
          </svg>
          تنظیمات همگام‌سازی
        </button>
      </div>
    </div>

    <!-- آمار همگام‌سازی -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-6">
      <!-- آخرین همگام‌سازی -->
      <div class="bg-blue-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-blue-600">{{ syncStats.lastSyncTime }}</div>
            <div class="text-sm text-blue-700">آخرین همگام‌سازی</div>
          </div>
          <div class="w-10 h-10 bg-blue-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
        <div class="mt-2 text-xs text-blue-600">
          {{ syncStats.lastSyncDuration }} مدت زمان
        </div>
      </div>

      <!-- داده‌های همگام‌سازی شده -->
      <div class="bg-green-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-green-600">{{ formatNumber(syncStats.syncedRecords) }}</div>
            <div class="text-sm text-green-700">رکورد همگام شده</div>
          </div>
          <div class="w-10 h-10 bg-green-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
        <div class="mt-2 text-xs text-green-600">
          امروز: {{ formatNumber(syncStats.todayRecords) }}
        </div>
      </div>

      <!-- نرخ موفقیت -->
      <div class="bg-purple-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-purple-600">{{ syncStats.successRate }}%</div>
            <div class="text-sm text-purple-700">نرخ موفقیت</div>
          </div>
          <div class="w-10 h-10 bg-purple-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
          </div>
        </div>
        <div class="mt-2 text-xs text-purple-600">
          {{ syncStats.totalSyncs }} همگام‌سازی کل
        </div>
      </div>

      <!-- خطاهای همگام‌سازی -->
      <div class="bg-red-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-red-600">{{ syncStats.syncErrors }}</div>
            <div class="text-sm text-red-700">خطای همگام‌سازی</div>
          </div>
          <div class="w-10 h-10 bg-red-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
        <div class="mt-2 text-xs text-red-600">
          آخرین: {{ syncStats.lastErrorTime }}
        </div>
      </div>
    </div>

    <!-- تنظیمات همگام‌سازی -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- زمان‌بندی همگام‌سازی -->
      <div class="bg-gray-50 rounded-lg p-6">
        <h4 class="font-medium text-gray-900 mb-4">زمان‌بندی همگام‌سازی</h4>
        <div class="space-y-4">
          <div>
            <label class="flex items-center">
              <input 
                v-model="syncSettings.autoSync"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">همگام‌سازی خودکار</span>
            </label>
          </div>
          
          <div v-if="syncSettings.autoSync" class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">فرکانس همگام‌سازی</label>
              <select 
                v-model="syncSettings.frequency"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="hourly">هر ساعت</option>
                <option value="daily">روزانه</option>
                <option value="weekly">هفتگی</option>
                <option value="custom">سفارشی</option>
              </select>
            </div>
            
            <div v-if="syncSettings.frequency === 'custom'">
              <label class="block text-sm font-medium text-gray-700 mb-2">فاصله زمانی (دقیقه)</label>
              <input 
                v-model.number="syncSettings.customInterval"
                type="number"
                min="5"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">ساعت شروع</label>
              <input 
                v-model="syncSettings.startTime"
                type="time"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
          </div>
        </div>
      </div>

      <!-- نوع داده‌های همگام‌سازی -->
      <div class="bg-gray-50 rounded-lg p-6">
        <h4 class="font-medium text-gray-900 mb-4">نوع داده‌های همگام‌سازی</h4>
        <div class="space-y-3">
          <label class="flex items-center">
            <input 
              v-model="syncSettings.syncInvoices"
              type="checkbox"
              class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
            />
            <span class="mr-2 text-sm text-gray-700">فاکتورها و رسیدها</span>
          </label>
          
          <label class="flex items-center">
            <input 
              v-model="syncSettings.syncCustomers"
              type="checkbox"
              class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
            />
            <span class="mr-2 text-sm text-gray-700">مشتریان</span>
          </label>
          
          <label class="flex items-center">
            <input 
              v-model="syncSettings.syncProducts"
              type="checkbox"
              class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
            />
            <span class="mr-2 text-sm text-gray-700">محصولات</span>
          </label>
          
          <label class="flex items-center">
            <input 
              v-model="syncSettings.syncPayments"
              type="checkbox"
              class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
            />
            <span class="mr-2 text-sm text-gray-700">پرداخت‌ها</span>
          </label>
          
          <label class="flex items-center">
            <input 
              v-model="syncSettings.syncTaxes"
              type="checkbox"
              class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
            />
            <span class="mr-2 text-sm text-gray-700">مالیات</span>
          </label>
          
          <label class="flex items-center">
            <input 
              v-model="syncSettings.syncInventory"
              type="checkbox"
              class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
            />
            <span class="mr-2 text-sm text-gray-700">موجودی</span>
          </label>
        </div>
      </div>
    </div>

    <!-- تاریخچه همگام‌سازی -->
    <div class="mt-6">
      <h4 class="font-medium text-gray-900 mb-4">تاریخچه همگام‌سازی</h4>
      <div class="overflow-x-auto">
        <table class="w-full text-sm">
          <thead>
            <tr class="border-b border-gray-200 bg-gray-50">
              <th class="text-right py-3 px-4 font-medium text-gray-600">تاریخ</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">نوع</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">وضعیت</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">رکوردها</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">مدت زمان</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">جزئیات</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="sync in syncHistory" :key="sync.id" class="border-b border-gray-100 hover:bg-gray-50">
              <td class="py-3 px-4 text-gray-900">{{ sync.date }}</td>
              <td class="py-3 px-4">
                <span 
                  class="px-2 py-1 rounded-full text-xs font-medium"
                  :class="sync.type === 'auto' ? 'bg-blue-100 text-blue-700' : 'bg-green-100 text-green-700'"
                >
                  {{ sync.type === 'auto' ? 'خودکار' : 'دستی' }}
                </span>
              </td>
              <td class="py-3 px-4">
                <span 
                  class="px-2 py-1 rounded-full text-xs font-medium"
                  :class="sync.status === 'success' ? 'bg-green-100 text-green-700' : 'bg-red-100 text-red-700'"
                >
                  {{ sync.status === 'success' ? 'موفق' : 'ناموفق' }}
                </span>
              </td>
              <td class="py-3 px-4 text-gray-900">{{ formatNumber(sync.records) }}</td>
              <td class="py-3 px-4 text-gray-900">{{ sync.duration }}</td>
              <td class="py-3 px-4">
                <button 
                  @click="showSyncDetails(sync)"
                  class="text-blue-600 hover:text-blue-800 text-xs"
                >
                  مشاهده جزئیات
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- مودال تنظیمات همگام‌سازی -->
    <div v-if="showSyncSettings" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-xl p-6 w-full max-w-2xl mx-4 max-h-[90vh] overflow-y-auto">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">تنظیمات همگام‌سازی</h3>
          <button @click="showSyncSettings = false" class="text-gray-400 hover:text-gray-600">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>

        <div class="space-y-6">
          <!-- تنظیمات پیشرفته -->
          <div>
            <h4 class="font-medium text-gray-900 mb-3">تنظیمات پیشرفته</h4>
            <div class="space-y-3">
              <label class="flex items-center">
                <input 
                  v-model="syncSettings.retryOnError"
                  type="checkbox"
                  class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                />
                <span class="mr-2 text-sm text-gray-700">تلاش مجدد در صورت خطا</span>
              </label>
              
              <label class="flex items-center">
                <input 
                  v-model="syncSettings.notifyOnError"
                  type="checkbox"
                  class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                />
                <span class="mr-2 text-sm text-gray-700">اعلان در صورت خطا</span>
              </label>
              
              <label class="flex items-center">
                <input 
                  v-model="syncSettings.backupBeforeSync"
                  type="checkbox"
                  class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                />
                <span class="mr-2 text-sm text-gray-700">پشتیبان‌گیری قبل از همگام‌سازی</span>
              </label>
            </div>
          </div>

          <!-- تنظیمات اتصال -->
          <div>
            <h4 class="font-medium text-gray-900 mb-3">تنظیمات اتصال</h4>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">Timeout (ثانیه)</label>
                <input 
                  v-model.number="syncSettings.timeout"
                  type="number"
                  min="10"
                  max="300"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر تلاش</label>
                <input 
                  v-model.number="syncSettings.maxRetries"
                  type="number"
                  min="1"
                  max="10"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                />
              </div>
            </div>
          </div>

          <!-- دکمه‌ها -->
          <div class="flex gap-3 pt-4">
            <button 
              @click="saveSyncSettings"
              class="flex-1 px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg transition-colors duration-200"
            >
              ذخیره تنظیمات
            </button>
            <button 
              @click="showSyncSettings = false"
              class="flex-1 px-4 py-2 bg-gray-200 hover:bg-gray-300 text-gray-700 rounded-lg transition-colors duration-200"
            >
              انصراف
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'

// وضعیت همگام‌سازی
const isSyncing = ref(false)
const showSyncSettings = ref(false)

// آمار همگام‌سازی
const syncStats = ref({
  lastSyncTime: '2 ساعت پیش',
  lastSyncDuration: '3 دقیقه',
  syncedRecords: 15420,
  todayRecords: 1250,
  successRate: 98.5,
  totalSyncs: 156,
  syncErrors: 2,
  lastErrorTime: '1 روز پیش'
})

// تنظیمات همگام‌سازی
const syncSettings = ref({
  autoSync: true,
  frequency: 'daily',
  customInterval: 60,
  startTime: '02:00',
  syncInvoices: true,
  syncCustomers: true,
  syncProducts: true,
  syncPayments: true,
  syncTaxes: true,
  syncInventory: true,
  retryOnError: true,
  notifyOnError: true,
  backupBeforeSync: false,
  timeout: 30,
  maxRetries: 3
})

// تاریخچه همگام‌سازی
const syncHistory = ref([
  {
    id: 1,
    date: '2024-01-15 14:30',
    type: 'auto',
    status: 'success',
    records: 1250,
    duration: '2 دقیقه'
  },
  {
    id: 2,
    date: '2024-01-15 02:00',
    type: 'auto',
    status: 'success',
    records: 980,
    duration: '1.5 دقیقه'
  },
  {
    id: 3,
    date: '2024-01-14 14:30',
    type: 'manual',
    status: 'success',
    records: 2100,
    duration: '4 دقیقه'
  },
  {
    id: 4,
    date: '2024-01-14 02:00',
    type: 'auto',
    status: 'error',
    records: 0,
    duration: '30 ثانیه'
  }
])

// فرمت اعداد
const formatNumber = (num: number): string => {
  return new Intl.NumberFormat('fa-IR').format(num)
}

// شروع همگام‌سازی دستی
const startManualSync = async () => {
  try {
    isSyncing.value = true
    // TODO: شروع همگام‌سازی دستی
    console.log('همگام‌سازی دستی شروع شد')
    
    // شبیه‌سازی همگام‌سازی
    await new Promise(resolve => setTimeout(resolve, 3000))
    
    console.log('همگام‌سازی دستی تکمیل شد')
  } catch (error) {
    console.error('خطا در همگام‌سازی:', error)
  } finally {
    isSyncing.value = false
  }
}

// ذخیره تنظیمات همگام‌سازی
const saveSyncSettings = async () => {
  try {
    // TODO: ذخیره تنظیمات
    console.log('تنظیمات همگام‌سازی ذخیره شد:', syncSettings.value)
    showSyncSettings.value = false
  } catch (error) {
    console.error('خطا در ذخیره تنظیمات:', error)
  }
}

// نمایش جزئیات همگام‌سازی
const showSyncDetails = (sync: any) => {
  // TODO: نمایش جزئیات همگام‌سازی
  console.log('جزئیات همگام‌سازی:', sync)
}

// بارگذاری اولیه
onMounted(() => {
  // TODO: بارگذاری تنظیمات و آمار همگام‌سازی از API
})
</script>

<!--
  کامپوننت همگام‌سازی داده‌ها
  شامل:
  1. آمار همگام‌سازی
  2. تنظیمات زمان‌بندی
  3. انتخاب نوع داده‌ها
  4. تاریخچه همگام‌سازی
  5. تنظیمات پیشرفته
  6. طراحی مدرن و کاملاً ریسپانسیو
--> 

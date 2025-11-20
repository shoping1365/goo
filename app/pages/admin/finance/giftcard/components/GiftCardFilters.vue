<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex justify-between items-center">
        <div>
          <h2 class="text-xl font-bold text-gray-900">فیلترهای گیفت کارت</h2>
          <p class="text-gray-600 mt-1">فیلتر کردن گیفت کارت‌ها بر اساس معیارهای مختلف</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button 
            class="px-4 py-2 bg-gray-600 text-white text-sm font-medium rounded-lg hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
            @click="resetAllFilters"
          >
            پاک کردن همه فیلترها
          </button>
          <button 
            class="px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
            @click="applyFilters"
          >
            اعمال فیلترها
          </button>
        </div>
      </div>
    </div>

    <!-- فیلترهای اصلی -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- فیلتر وضعیت کارت -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-medium text-gray-900 mb-4">وضعیت کارت</h3>
        
        <div class="space-y-3">
          <div v-for="status in cardStatuses" :key="status.value" class="flex items-center">
            <input
              v-model="filters.statuses"
              :value="status.value"
              type="checkbox"
              class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
            />
            <label class="mr-3 text-sm text-gray-700 flex items-center">
              <div :class="status.colorClass" class="w-3 h-3 rounded-full ml-2"></div>
              {{ status.label }}
              <span class="text-xs text-gray-500 mr-1">({{ status.count }})</span>
            </label>
          </div>
        </div>

        <!-- فیلترهای اضافی وضعیت -->
        <div class="mt-4 pt-4 border-t border-gray-200">
          <h4 class="text-sm font-medium text-gray-700 mb-3">فیلترهای پیشرفته وضعیت</h4>
          
          <div class="space-y-3">
            <div class="flex items-center">
              <input
                v-model="filters.isUsed"
                type="checkbox"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <label class="mr-2 text-sm text-gray-700">فقط کارت‌های استفاده شده</label>
            </div>
            
            <div class="flex items-center">
              <input
                v-model="filters.isExpired"
                type="checkbox"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <label class="mr-2 text-sm text-gray-700">فقط کارت‌های منقضی شده</label>
            </div>
            
            <div class="flex items-center">
              <input
                v-model="filters.isExpiringSoon"
                type="checkbox"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <label class="mr-2 text-sm text-gray-700">کارت‌های در حال انقضا (7 روز آینده)</label>
            </div>
            
            <div class="flex items-center">
              <input
                v-model="filters.requireVerification"
                type="checkbox"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <label class="mr-2 text-sm text-gray-700">کارت‌های نیازمند تأیید</label>
            </div>
          </div>
        </div>
      </div>

      <!-- فیلتر نوع کارت -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-medium text-gray-900 mb-4">نوع کارت</h3>
        
        <div class="space-y-3">
          <div v-for="type in cardTypes" :key="type.value" class="flex items-center">
            <input
              v-model="filters.types"
              :value="type.value"
              type="checkbox"
              class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
            />
            <label class="mr-3 text-sm text-gray-700 flex items-center">
              <component :is="type.icon" class="w-4 h-4 ml-2 text-gray-500" />
              {{ type.label }}
              <span class="text-xs text-gray-500 mr-1">({{ type.count }})</span>
            </label>
          </div>
        </div>

        <!-- فیلترهای اضافی نوع -->
        <div class="mt-4 pt-4 border-t border-gray-200">
          <h4 class="text-sm font-medium text-gray-700 mb-3">فیلترهای پیشرفته نوع</h4>
          
          <div class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">دسته‌بندی</label>
              <select
                v-model="filters.category"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="">همه دسته‌بندی‌ها</option>
                <option v-for="category in cardCategories" :key="category.value" :value="category.value">
                  {{ category.label }} ({{ category.count }})
                </option>
              </select>
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">روش تحویل</label>
              <div class="space-y-2">
                <div v-for="method in deliveryMethods" :key="method.value" class="flex items-center">
                  <input
                    v-model="filters.deliveryMethods"
                    :value="method.value"
                    type="checkbox"
                    class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                  />
                  <label class="mr-2 text-sm text-gray-700">{{ method.label }}</label>
                </div>
              </div>
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">محدودیت استفاده</label>
              <select
                v-model="filters.usageLimit"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="">همه</option>
                <option value="unlimited">نامحدود</option>
                <option value="once">یک بار</option>
                <option value="multiple">چند بار</option>
              </select>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- فیلتر محدوده مبلغ -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-medium text-gray-900 mb-4">محدوده مبلغ</h3>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <!-- مبلغ کارت -->
        <div>
          <h4 class="text-sm font-medium text-gray-700 mb-3">مبلغ کارت</h4>
          
          <div class="space-y-4">
            <div>
              <label class="block text-sm text-gray-600 mb-1">حداقل مبلغ</label>
              <div class="relative">
                <input
                  v-model.number="filters.amountMin"
                  type="number"
                  min="0"
                  placeholder="0"
                  class="w-full px-3 py-2 pr-12 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                />
                <div class="absolute inset-y-0 right-0 pr-3 flex items-center">
                  <span class="text-gray-500 text-sm">تومان</span>
                </div>
              </div>
            </div>
            
            <div>
              <label class="block text-sm text-gray-600 mb-1">حداکثر مبلغ</label>
              <div class="relative">
                <input
                  v-model.number="filters.amountMax"
                  type="number"
                  min="0"
                  placeholder="نامحدود"
                  class="w-full px-3 py-2 pr-12 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                />
                <div class="absolute inset-y-0 right-0 pr-3 flex items-center">
                  <span class="text-gray-500 text-sm">تومان</span>
                </div>
              </div>
            </div>
          </div>

          <!-- پیش‌تعریف‌های مبلغ -->
          <div class="mt-4">
            <label class="block text-sm font-medium text-gray-700 mb-2">محدوده‌های پیش‌تعریف</label>
            <div class="grid grid-cols-2 gap-2">
              <button 
                v-for="range in amountRanges" 
                :key="range.id"
                class="px-3 py-2 text-xs border border-gray-300 rounded-md hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
                @click="setAmountRange(range)"
              >
                {{ range.label }}
              </button>
            </div>
          </div>
        </div>

        <!-- مبلغ باقی‌مانده -->
        <div>
          <h4 class="text-sm font-medium text-gray-700 mb-3">مبلغ باقی‌مانده</h4>
          
          <div class="space-y-4">
            <div>
              <label class="block text-sm text-gray-600 mb-1">حداقل مبلغ باقی‌مانده</label>
              <div class="relative">
                <input
                  v-model.number="filters.remainingAmountMin"
                  type="number"
                  min="0"
                  placeholder="0"
                  class="w-full px-3 py-2 pr-12 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                />
                <div class="absolute inset-y-0 right-0 pr-3 flex items-center">
                  <span class="text-gray-500 text-sm">تومان</span>
                </div>
              </div>
            </div>
            
            <div>
              <label class="block text-sm text-gray-600 mb-1">حداکثر مبلغ باقی‌مانده</label>
              <div class="relative">
                <input
                  v-model.number="filters.remainingAmountMax"
                  type="number"
                  min="0"
                  placeholder="نامحدود"
                  class="w-full px-3 py-2 pr-12 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                />
                <div class="absolute inset-y-0 right-0 pr-3 flex items-center">
                  <span class="text-gray-500 text-sm">تومان</span>
                </div>
              </div>
            </div>
          </div>

          <!-- درصد استفاده -->
          <div class="mt-4">
            <label class="block text-sm font-medium text-gray-700 mb-2">درصد استفاده</label>
            <div class="space-y-2">
              <div>
                <label class="block text-xs text-gray-600 mb-1">حداقل درصد استفاده</label>
                <input
                  v-model.number="filters.usagePercentageMin"
                  type="range"
                  min="0"
                  max="100"
                  class="w-full"
                />
                <div class="flex justify-between text-xs text-gray-500">
                  <span>0%</span>
                  <span>{{ filters.usagePercentageMin || 0 }}%</span>
                  <span>100%</span>
                </div>
              </div>
              
              <div>
                <label class="block text-xs text-gray-600 mb-1">حداکثر درصد استفاده</label>
                <input
                  v-model.number="filters.usagePercentageMax"
                  type="range"
                  min="0"
                  max="100"
                  class="w-full"
                />
                <div class="flex justify-between text-xs text-gray-500">
                  <span>0%</span>
                  <span>{{ filters.usagePercentageMax || 100 }}%</span>
                  <span>100%</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- فیلتر محدوده تاریخ -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-medium text-gray-900 mb-4">محدوده تاریخ</h3>
      
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <!-- تاریخ ایجاد -->
        <div>
          <h4 class="text-sm font-medium text-gray-700 mb-3">تاریخ ایجاد</h4>
          
          <div class="space-y-3">
            <div>
              <label class="block text-sm text-gray-600 mb-1">از تاریخ</label>
              <input
                v-model="filters.createdDateFrom"
                type="date"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
            
            <div>
              <label class="block text-sm text-gray-600 mb-1">تا تاریخ</label>
              <input
                v-model="filters.createdDateTo"
                type="date"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
          </div>

          <!-- پیش‌تعریف‌های تاریخ ایجاد -->
          <div class="mt-3">
            <div class="space-y-1">
              <button 
                v-for="period in createdDatePeriods" 
                :key="period.id"
                class="w-full text-right px-2 py-1 text-xs border border-gray-300 rounded hover:bg-gray-50"
                @click="setCreatedDatePeriod(period)"
              >
                {{ period.label }}
              </button>
            </div>
          </div>
        </div>

        <!-- تاریخ انقضا -->
        <div>
          <h4 class="text-sm font-medium text-gray-700 mb-3">تاریخ انقضا</h4>
          
          <div class="space-y-3">
            <div>
              <label class="block text-sm text-gray-600 mb-1">از تاریخ</label>
              <input
                v-model="filters.expiryDateFrom"
                type="date"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
            
            <div>
              <label class="block text-sm text-gray-600 mb-1">تا تاریخ</label>
              <input
                v-model="filters.expiryDateTo"
                type="date"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
          </div>

          <!-- پیش‌تعریف‌های تاریخ انقضا -->
          <div class="mt-3">
            <div class="space-y-1">
              <button 
                v-for="period in expiryDatePeriods" 
                :key="period.id"
                class="w-full text-right px-2 py-1 text-xs border border-gray-300 rounded hover:bg-gray-50"
                @click="setExpiryDatePeriod(period)"
              >
                {{ period.label }}
              </button>
            </div>
          </div>
        </div>

        <!-- تاریخ تحویل -->
        <div>
          <h4 class="text-sm font-medium text-gray-700 mb-3">تاریخ تحویل</h4>
          
          <div class="space-y-3">
            <div>
              <label class="block text-sm text-gray-600 mb-1">از تاریخ</label>
              <input
                v-model="filters.deliveryDateFrom"
                type="date"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
            
            <div>
              <label class="block text-sm text-gray-600 mb-1">تا تاریخ</label>
              <input
                v-model="filters.deliveryDateTo"
                type="date"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
          </div>

          <!-- پیش‌تعریف‌های تاریخ تحویل -->
          <div class="mt-3">
            <div class="space-y-1">
              <button 
                v-for="period in deliveryDatePeriods" 
                :key="period.id"
                class="w-full text-right px-2 py-1 text-xs border border-gray-300 rounded hover:bg-gray-50"
                @click="setDeliveryDatePeriod(period)"
              >
                {{ period.label }}
              </button>
            </div>
          </div>
        </div>

        <!-- آخرین استفاده -->
        <div>
          <h4 class="text-sm font-medium text-gray-700 mb-3">آخرین استفاده</h4>
          
          <div class="space-y-3">
            <div>
              <label class="block text-sm text-gray-600 mb-1">از تاریخ</label>
              <input
                v-model="filters.lastUsedDateFrom"
                type="date"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
            
            <div>
              <label class="block text-sm text-gray-600 mb-1">تا تاریخ</label>
              <input
                v-model="filters.lastUsedDateTo"
                type="date"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
          </div>

          <!-- پیش‌تعریف‌های آخرین استفاده -->
          <div class="mt-3">
            <div class="space-y-1">
              <button 
                v-for="period in lastUsedDatePeriods" 
                :key="period.id"
                class="w-full text-right px-2 py-1 text-xs border border-gray-300 rounded hover:bg-gray-50"
                @click="setLastUsedDatePeriod(period)"
              >
                {{ period.label }}
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- خلاصه فیلترها -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-medium text-gray-900 mb-4">خلاصه فیلترهای اعمال شده</h3>
      
      <div v-if="activeFiltersCount === 0" class="text-center py-4">
        <p class="text-gray-500">هیچ فیلتری اعمال نشده است</p>
      </div>
      
      <div v-else class="space-y-3">
        <div v-if="filters.statuses.length > 0" class="flex items-center space-x-2 space-x-reverse">
          <span class="text-sm font-medium text-gray-700">وضعیت:</span>
          <div class="flex flex-wrap gap-1">
            <span 
              v-for="status in filters.statuses" 
              :key="status"
              class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium bg-blue-100 text-blue-800"
            >
              {{ getStatusLabel(status) }}
            </span>
          </div>
        </div>
        
        <div v-if="filters.types.length > 0" class="flex items-center space-x-2 space-x-reverse">
          <span class="text-sm font-medium text-gray-700">نوع:</span>
          <div class="flex flex-wrap gap-1">
            <span 
              v-for="type in filters.types" 
              :key="type"
              class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium bg-green-100 text-green-800"
            >
              {{ getTypeLabel(type) }}
            </span>
          </div>
        </div>
        
        <div v-if="filters.amountMin || filters.amountMax" class="flex items-center space-x-2 space-x-reverse">
          <span class="text-sm font-medium text-gray-700">مبلغ کارت:</span>
          <span class="text-sm text-gray-600">
            {{ filters.amountMin ? formatCurrency(filters.amountMin) : '0' }} - 
            {{ filters.amountMax ? formatCurrency(filters.amountMax) : 'نامحدود' }}
          </span>
        </div>
        
        <div v-if="filters.createdDateFrom || filters.createdDateTo" class="flex items-center space-x-2 space-x-reverse">
          <span class="text-sm font-medium text-gray-700">تاریخ ایجاد:</span>
          <span class="text-sm text-gray-600">
            {{ filters.createdDateFrom || 'نامحدود' }} تا {{ filters.createdDateTo || 'نامحدود' }}
          </span>
        </div>
        
        <div class="flex items-center space-x-2 space-x-reverse">
          <span class="text-sm font-medium text-gray-700">تعداد فیلترهای فعال:</span>
          <span class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium bg-purple-100 text-purple-800">
            {{ activeFiltersCount }}
          </span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted } from 'vue'

// Props
const props = defineProps<{
  giftCards: any[]
}>()

// Emits
const emit = defineEmits<{
  'apply-filters': [filters: any]
}>()

// Reactive data
const filters = reactive({
  // وضعیت کارت
  statuses: [],
  isUsed: false,
  isExpired: false,
  isExpiringSoon: false,
  requireVerification: false,
  
  // نوع کارت
  types: [],
  category: '',
  deliveryMethods: [],
  usageLimit: '',
  
  // محدوده مبلغ
  amountMin: null,
  amountMax: null,
  remainingAmountMin: null,
  remainingAmountMax: null,
  usagePercentageMin: null,
  usagePercentageMax: null,
  
  // محدوده تاریخ
  createdDateFrom: '',
  createdDateTo: '',
  expiryDateFrom: '',
  expiryDateTo: '',
  deliveryDateFrom: '',
  deliveryDateTo: '',
  lastUsedDateFrom: '',
  lastUsedDateTo: ''
})

// داده‌های مرجع
const cardStatuses = [
  { value: 'active', label: 'فعال', colorClass: 'bg-green-500', count: 0 },
  { value: 'used', label: 'استفاده شده', colorClass: 'bg-blue-500', count: 0 },
  { value: 'expired', label: 'منقضی شده', colorClass: 'bg-red-500', count: 0 },
  { value: 'suspended', label: 'معلق', colorClass: 'bg-yellow-500', count: 0 },
  { value: 'cancelled', label: 'لغو شده', colorClass: 'bg-gray-500', count: 0 }
]

const cardTypes = [
  { value: 'digital', label: 'دیجیتال', icon: 'DeviceMobileIcon', count: 0 },
  { value: 'physical', label: 'فیزیکی', icon: 'CreditCardIcon', count: 0 },
  { value: 'virtual', label: 'مجازی', icon: 'GlobeAltIcon', count: 0 }
]

const cardCategories = [
  { value: 'birthday', label: 'تولد', count: 0 },
  { value: 'wedding', label: 'عروسی', count: 0 },
  { value: 'newyear', label: 'سال نو', count: 0 },
  { value: 'christmas', label: 'کریسمس', count: 0 },
  { value: 'graduation', label: 'فارغ‌التحصیلی', count: 0 },
  { value: 'anniversary', label: 'سالگرد', count: 0 },
  { value: 'general', label: 'عمومی', count: 0 }
]

const deliveryMethods = [
  { value: 'email', label: 'ایمیل' },
  { value: 'sms', label: 'پیامک' },
  { value: 'both', label: 'هر دو' },
  { value: 'physical', label: 'فیزیکی' }
]

const amountRanges = [
  { id: 1, label: 'کمتر از 100,000 تومان', min: 0, max: 100000 },
  { id: 2, label: '100,000 - 500,000 تومان', min: 100000, max: 500000 },
  { id: 3, label: '500,000 - 1,000,000 تومان', min: 500000, max: 1000000 },
  { id: 4, label: 'بیش از 1,000,000 تومان', min: 1000000, max: null }
]

const createdDatePeriods = [
  { id: 1, label: 'امروز', from: 'today', to: 'today' },
  { id: 2, label: 'هفته جاری', from: 'week', to: 'today' },
  { id: 3, label: 'ماه جاری', from: 'month', to: 'today' },
  { id: 4, label: '30 روز گذشته', from: '30days', to: 'today' }
]

const expiryDatePeriods = [
  { id: 1, label: 'امروز منقضی می‌شوند', from: 'today', to: 'today' },
  { id: 2, label: '7 روز آینده', from: 'today', to: '7days' },
  { id: 3, label: '30 روز آینده', from: 'today', to: '30daysLater' },
  { id: 4, label: 'قبلاً منقضی شده', from: 'past', to: 'today' }
]

const deliveryDatePeriods = [
  { id: 1, label: 'امروز تحویل داده شد', from: 'today', to: 'today' },
  { id: 2, label: 'هفته جاری', from: 'week', to: 'today' },
  { id: 3, label: 'ماه جاری', from: 'month', to: 'today' },
  { id: 4, label: 'آینده', from: 'today', to: 'future' }
]

const lastUsedDatePeriods = [
  { id: 1, label: 'امروز', from: 'today', to: 'today' },
  { id: 2, label: 'هفته گذشته', from: 'week', to: 'today' },
  { id: 3, label: 'ماه گذشته', from: 'month', to: 'today' },
  { id: 4, label: 'هرگز استفاده نشده', from: 'never', to: 'never' }
]

// Computed properties
const activeFiltersCount = computed(() => {
  let count = 0
  
  if (filters.statuses.length > 0) count++
  if (filters.types.length > 0) count++
  if (filters.category) count++
  if (filters.deliveryMethods.length > 0) count++
  if (filters.usageLimit) count++
  if (filters.isUsed) count++
  if (filters.isExpired) count++
  if (filters.isExpiringSoon) count++
  if (filters.requireVerification) count++
  if (filters.amountMin !== null) count++
  if (filters.amountMax !== null) count++
  if (filters.remainingAmountMin !== null) count++
  if (filters.remainingAmountMax !== null) count++
  if (filters.usagePercentageMin !== null) count++
  if (filters.usagePercentageMax !== null) count++
  if (filters.createdDateFrom) count++
  if (filters.createdDateTo) count++
  if (filters.expiryDateFrom) count++
  if (filters.expiryDateTo) count++
  if (filters.deliveryDateFrom) count++
  if (filters.deliveryDateTo) count++
  if (filters.lastUsedDateFrom) count++
  if (filters.lastUsedDateTo) count++
  
  return count
})

// Methods
const applyFilters = () => {
  emit('apply-filters', { ...filters })
  console.log('فیلترها اعمال شد:', filters)
}

const resetAllFilters = () => {
  Object.keys(filters).forEach(key => {
    if (Array.isArray(filters[key])) {
      filters[key] = []
    } else if (typeof filters[key] === 'number') {
      filters[key] = null
    } else if (typeof filters[key] === 'boolean') {
      filters[key] = false
    } else {
      filters[key] = ''
    }
  })
  console.log('همه فیلترها پاک شد')
}

const setAmountRange = (range: any) => {
  filters.amountMin = range.min
  filters.amountMax = range.max
}

const setCreatedDatePeriod = (period: any) => {
  const { from, to } = period
  filters.createdDateFrom = getDateFromPeriod(from)
  filters.createdDateTo = getDateFromPeriod(to)
}

const setExpiryDatePeriod = (period: any) => {
  const { from, to } = period
  filters.expiryDateFrom = getDateFromPeriod(from)
  filters.expiryDateTo = getDateFromPeriod(to)
}

const setDeliveryDatePeriod = (period: any) => {
  const { from, to } = period
  filters.deliveryDateFrom = getDateFromPeriod(from)
  filters.deliveryDateTo = getDateFromPeriod(to)
}

const setLastUsedDatePeriod = (period: any) => {
  const { from, to } = period
  filters.lastUsedDateFrom = getDateFromPeriod(from)
  filters.lastUsedDateTo = getDateFromPeriod(to)
}

const getDateFromPeriod = (period: string) => {
  const today = new Date()
  
  switch (period) {
    case 'today':
      return today.toISOString().split('T')[0]
    case 'week':
      const weekAgo = new Date(today.getTime() - 7 * 24 * 60 * 60 * 1000)
      return weekAgo.toISOString().split('T')[0]
    case 'month':
      const monthAgo = new Date(today.getTime() - 30 * 24 * 60 * 60 * 1000)
      return monthAgo.toISOString().split('T')[0]
    case '30days':
      const thirtyDaysAgo = new Date(today.getTime() - 30 * 24 * 60 * 60 * 1000)
      return thirtyDaysAgo.toISOString().split('T')[0]
    case '7days':
      const sevenDaysLater = new Date(today.getTime() + 7 * 24 * 60 * 60 * 1000)
      return sevenDaysLater.toISOString().split('T')[0]
    case '30daysLater':
      const thirtyDaysLater = new Date(today.getTime() + 30 * 24 * 60 * 60 * 1000)
      return thirtyDaysLater.toISOString().split('T')[0]
    case 'past':
      return '1900-01-01'
    case 'future':
      return '2100-12-31'
    case 'never':
      return ''
    default:
      return ''
  }
}

const updateCounts = () => {
  // به‌روزرسانی تعداد کارت‌ها در هر وضعیت
  cardStatuses.forEach(status => {
    status.count = props.giftCards.filter(card => card.status === status.value).length
  })
  
  // به‌روزرسانی تعداد کارت‌ها در هر نوع
  cardTypes.forEach(type => {
    type.count = props.giftCards.filter(card => card.type === type.value).length
  })
  
  // به‌روزرسانی تعداد کارت‌ها در هر دسته‌بندی
  cardCategories.forEach(category => {
    category.count = props.giftCards.filter(card => card.category === category.value).length
  })
}

const getStatusLabel = (status: string) => {
  const statusObj = cardStatuses.find(s => s.value === status)
  return statusObj ? statusObj.label : status
}

const getTypeLabel = (type: string) => {
  const typeObj = cardTypes.find(t => t.value === type)
  return typeObj ? typeObj.label : type
}

const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

// Lifecycle
onMounted(() => {
  updateCounts()
  console.log('Gift card filters component mounted')
})
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 

<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex justify-between items-center">
        <div>
          <h2 class="text-xl font-bold text-gray-900">جستجو و فیلتر پیشرفته</h2>
          <p class="text-gray-600 mt-1">جستجو و فیلتر پیشرفته گیفت کارت‌ها بر اساس معیارهای مختلف</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button 
            class="px-4 py-2 bg-gray-600 text-white text-sm font-medium rounded-lg hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
            @click="clearAllFilters"
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
      <!-- جستجوی متنی -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-medium text-gray-900 mb-4">جستجوی متنی</h3>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">کد کارت</label>
            <div class="relative">
              <input
                v-model="filters.cardCode"
                type="text"
                placeholder="مثال: GC-BIRTHDAY-2024-001"
                class="w-full px-3 py-2 pr-10 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              />
              <div class="absolute inset-y-0 right-0 pr-3 flex items-center">
                <svg class="h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
                </svg>
              </div>
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نام خریدار</label>
            <input
              v-model="filters.buyerName"
              type="text"
              placeholder="نام خریدار را وارد کنید"
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نام گیرنده</label>
            <input
              v-model="filters.recipientName"
              type="text"
              placeholder="نام گیرنده را وارد کنید"
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">ایمیل گیرنده</label>
            <input
              v-model="filters.recipientEmail"
              type="email"
              placeholder="ایمیل گیرنده را وارد کنید"
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">شماره تماس گیرنده</label>
            <input
              v-model="filters.recipientPhone"
              type="tel"
              placeholder="شماره تماس گیرنده را وارد کنید"
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            />
          </div>
        </div>
      </div>

      <!-- فیلترهای تاریخ -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-medium text-gray-900 mb-4">فیلترهای تاریخ</h3>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ ایجاد</label>
            <div class="grid grid-cols-2 gap-2">
              <div>
                <label class="block text-xs text-gray-500 mb-1">از تاریخ</label>
                <input
                  v-model="filters.createdDateFrom"
                  type="date"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                />
              </div>
              <div>
                <label class="block text-xs text-gray-500 mb-1">تا تاریخ</label>
                <input
                  v-model="filters.createdDateTo"
                  type="date"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                />
              </div>
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ انقضا</label>
            <div class="grid grid-cols-2 gap-2">
              <div>
                <label class="block text-xs text-gray-500 mb-1">از تاریخ</label>
                <input
                  v-model="filters.expiryDateFrom"
                  type="date"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                />
              </div>
              <div>
                <label class="block text-xs text-gray-500 mb-1">تا تاریخ</label>
                <input
                  v-model="filters.expiryDateTo"
                  type="date"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                />
              </div>
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ تحویل</label>
            <div class="grid grid-cols-2 gap-2">
              <div>
                <label class="block text-xs text-gray-500 mb-1">از تاریخ</label>
                <input
                  v-model="filters.deliveryDateFrom"
                  type="date"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                />
              </div>
              <div>
                <label class="block text-xs text-gray-500 mb-1">تا تاریخ</label>
                <input
                  v-model="filters.deliveryDateTo"
                  type="date"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                />
              </div>
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">آخرین استفاده</label>
            <div class="grid grid-cols-2 gap-2">
              <div>
                <label class="block text-xs text-gray-500 mb-1">از تاریخ</label>
                <input
                  v-model="filters.lastUsedDateFrom"
                  type="date"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                />
              </div>
              <div>
                <label class="block text-xs text-gray-500 mb-1">تا تاریخ</label>
                <input
                  v-model="filters.lastUsedDateTo"
                  type="date"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- فیلترهای مالی و وضعیت -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- فیلترهای مالی -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-medium text-gray-900 mb-4">فیلترهای مالی</h3>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">مبلغ کارت</label>
            <div class="grid grid-cols-2 gap-2">
              <div>
                <label class="block text-xs text-gray-500 mb-1">حداقل</label>
                <input
                  v-model.number="filters.amountMin"
                  type="number"
                  min="0"
                  placeholder="0"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                />
              </div>
              <div>
                <label class="block text-xs text-gray-500 mb-1">حداکثر</label>
                <input
                  v-model.number="filters.amountMax"
                  type="number"
                  min="0"
                  placeholder="نامحدود"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                />
              </div>
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">مبلغ باقی‌مانده</label>
            <div class="grid grid-cols-2 gap-2">
              <div>
                <label class="block text-xs text-gray-500 mb-1">حداقل</label>
                <input
                  v-model.number="filters.remainingAmountMin"
                  type="number"
                  min="0"
                  placeholder="0"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                />
              </div>
              <div>
                <label class="block text-xs text-gray-500 mb-1">حداکثر</label>
                <input
                  v-model.number="filters.remainingAmountMax"
                  type="number"
                  min="0"
                  placeholder="نامحدود"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                />
              </div>
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">درصد استفاده</label>
            <div class="grid grid-cols-2 gap-2">
              <div>
                <label class="block text-xs text-gray-500 mb-1">حداقل</label>
                <input
                  v-model.number="filters.usagePercentageMin"
                  type="number"
                  min="0"
                  max="100"
                  placeholder="0%"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                />
              </div>
              <div>
                <label class="block text-xs text-gray-500 mb-1">حداکثر</label>
                <input
                  v-model.number="filters.usagePercentageMax"
                  type="number"
                  min="0"
                  max="100"
                  placeholder="100%"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                />
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- فیلترهای وضعیت و نوع -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-medium text-gray-900 mb-4">فیلترهای وضعیت و نوع</h3>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت کارت</label>
            <div class="space-y-2">
              <div v-for="status in cardStatuses" :key="status.value" class="flex items-center">
                <input
                  v-model="filters.statuses"
                  :value="status.value"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label class="mr-2 text-sm text-gray-700">{{ status.label }}</label>
              </div>
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع کارت</label>
            <div class="space-y-2">
              <div v-for="type in cardTypes" :key="type.value" class="flex items-center">
                <input
                  v-model="filters.types"
                  :value="type.value"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label class="mr-2 text-sm text-gray-700">{{ type.label }}</label>
              </div>
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">دسته‌بندی</label>
            <div class="space-y-2">
              <div v-for="category in cardCategories" :key="category.value" class="flex items-center">
                <input
                  v-model="filters.categories"
                  :value="category.value"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label class="mr-2 text-sm text-gray-700">{{ category.label }}</label>
              </div>
            </div>
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
        </div>
      </div>
    </div>

    <!-- فیلترهای پیشرفته -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-medium text-gray-900 mb-4">فیلترهای پیشرفته</h3>
      
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
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

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نیاز به تأیید</label>
          <select
            v-model="filters.requireVerification"
            class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه</option>
            <option value="true">بله</option>
            <option value="false">خیر</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">استفاده جزئی</label>
          <select
            v-model="filters.allowPartialUsage"
            class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه</option>
            <option value="true">مجاز</option>
            <option value="false">غیرمجاز</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تمدید خودکار</label>
          <select
            v-model="filters.autoRenew"
            class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه</option>
            <option value="true">فعال</option>
            <option value="false">غیرفعال</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">کارت‌های استفاده شده</label>
          <select
            v-model="filters.isUsed"
            class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه</option>
            <option value="true">استفاده شده</option>
            <option value="false">استفاده نشده</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">کارت‌های منقضی شده</label>
          <select
            v-model="filters.isExpired"
            class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه</option>
            <option value="true">منقضی شده</option>
            <option value="false">معتبر</option>
          </select>
        </div>
      </div>
    </div>

    <!-- نتایج جستجو -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <div class="flex justify-between items-center">
          <h3 class="text-lg font-medium text-gray-900">نتایج جستجو</h3>
          <div class="flex items-center space-x-3 space-x-reverse">
            <span class="text-sm text-gray-500">{{ filteredResults.length }} نتیجه یافت شد</span>
            <button 
              class="px-3 py-1 bg-green-600 text-white text-sm font-medium rounded-md hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2"
              @click="exportResults"
            >
              خروجی Excel
            </button>
          </div>
        </div>
      </div>
      
      <div class="p-6">
        <div v-if="filteredResults.length === 0" class="text-center py-8">
          <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
          </svg>
          <h3 class="mt-2 text-sm font-medium text-gray-900">نتیجه‌ای یافت نشد</h3>
          <p class="mt-1 text-sm text-gray-500">با فیلترهای انتخاب شده هیچ گیفت کارتی یافت نشد.</p>
        </div>
        
        <div v-else class="space-y-4">
          <div 
            v-for="card in paginatedResults" 
            :key="card.id"
            class="border border-gray-200 rounded-lg p-6 hover:bg-gray-50 transition-colors duration-200"
          >
            <div class="flex items-center justify-between">
              <div class="flex items-center space-x-4 space-x-reverse">
                <div class="flex-shrink-0">
                  <div :class="getStatusColorClass(card.status)" class="w-3 h-3 rounded-full"></div>
                </div>
                <div>
                  <h4 class="text-sm font-medium text-gray-900">{{ card.code }}</h4>
                  <p class="text-xs text-gray-500">{{ card.recipientName }} - {{ card.recipientEmail }}</p>
                </div>
              </div>
              <div class="flex items-center space-x-4 space-x-reverse">
                <div class="text-right">
                  <p class="text-sm font-medium text-gray-900">{{ formatCurrency(card.amount) }}</p>
                  <p class="text-xs text-gray-500">باقی‌مانده: {{ formatCurrency(card.remainingAmount) }}</p>
                </div>
                <div class="flex items-center space-x-2 space-x-reverse">
                  <button 
                    class="text-blue-600 hover:text-blue-800 text-sm"
                    @click="viewCardDetails(card)"
                  >
                    جزئیات
                  </button>
                  <button 
                    class="text-green-600 hover:text-green-800 text-sm"
                    @click="editCard(card)"
                  >
                    ویرایش
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
        
        <!-- صفحه‌بندی -->
        <div v-if="totalPages > 1" class="mt-6 flex items-center justify-between">
          <div class="flex items-center space-x-2 space-x-reverse">
            <button 
              :disabled="currentPage === 1"
              class="px-3 py-1 border border-gray-300 rounded-md text-sm disabled:opacity-50 disabled:cursor-not-allowed"
              @click="previousPage"
            >
              قبلی
            </button>
            <span class="text-sm text-gray-700">
              صفحه {{ currentPage }} از {{ totalPages }}
            </span>
            <button 
              :disabled="currentPage === totalPages"
              class="px-3 py-1 border border-gray-300 rounded-md text-sm disabled:opacity-50 disabled:cursor-not-allowed"
              @click="nextPage"
            >
              بعدی
            </button>
          </div>
          <div class="flex items-center space-x-2 space-x-reverse">
            <span class="text-sm text-gray-500">نمایش:</span>
            <select 
              v-model="pageSize"
              class="px-2 py-1 border border-gray-300 rounded-md text-sm"
              @change="currentPage = 1"
            >
              <option value="10">10</option>
              <option value="25">25</option>
              <option value="50">50</option>
              <option value="100">100</option>
            </select>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted, watch } from 'vue'

// Props
interface GiftCard {
  id?: number | string
  [key: string]: unknown
}

const props = defineProps<{
  giftCards: GiftCard[]
}>()

// Emits
const emit = defineEmits<{
  'view-details': [card: GiftCard]
  'edit-card': [card: GiftCard]
}>()

// Reactive data
const currentPage = ref(1)
const pageSize = ref(25)

// فیلترها
const filters = reactive({
  // جستجوی متنی
  cardCode: '',
  buyerName: '',
  recipientName: '',
  recipientEmail: '',
  recipientPhone: '',
  
  // فیلترهای تاریخ
  createdDateFrom: '',
  createdDateTo: '',
  expiryDateFrom: '',
  expiryDateTo: '',
  deliveryDateFrom: '',
  deliveryDateTo: '',
  lastUsedDateFrom: '',
  lastUsedDateTo: '',
  
  // فیلترهای مالی
  amountMin: null,
  amountMax: null,
  remainingAmountMin: null,
  remainingAmountMax: null,
  usagePercentageMin: null,
  usagePercentageMax: null,
  
  // فیلترهای وضعیت و نوع
  statuses: [],
  types: [],
  categories: [],
  deliveryMethods: [],
  
  // فیلترهای پیشرفته
  usageLimit: '',
  requireVerification: '',
  allowPartialUsage: '',
  autoRenew: '',
  isUsed: '',
  isExpired: ''
})

// داده‌های مرجع
const cardStatuses = [
  { value: 'active', label: 'فعال' },
  { value: 'used', label: 'استفاده شده' },
  { value: 'expired', label: 'منقضی شده' },
  { value: 'suspended', label: 'معلق' },
  { value: 'cancelled', label: 'لغو شده' }
]

const cardTypes = [
  { value: 'digital', label: 'دیجیتال' },
  { value: 'physical', label: 'فیزیکی' },
  { value: 'virtual', label: 'مجازی' }
]

const cardCategories = [
  { value: 'birthday', label: 'تولد' },
  { value: 'wedding', label: 'عروسی' },
  { value: 'newyear', label: 'سال نو' },
  { value: 'christmas', label: 'کریسمس' },
  { value: 'graduation', label: 'فارغ‌التحصیلی' },
  { value: 'anniversary', label: 'سالگرد' },
  { value: 'general', label: 'عمومی' }
]

const deliveryMethods = [
  { value: 'email', label: 'ایمیل' },
  { value: 'sms', label: 'پیامک' },
  { value: 'both', label: 'هر دو' },
  { value: 'physical', label: 'فیزیکی' }
]

// Computed properties
const filteredResults = computed(() => {
  let results = [...props.giftCards]

  // فیلتر کد کارت
  if (filters.cardCode) {
    results = results.filter(card => 
      card.code.toLowerCase().includes(filters.cardCode.toLowerCase())
    )
  }

  // فیلتر نام خریدار
  if (filters.buyerName) {
    results = results.filter(card => 
      card.buyerName?.toLowerCase().includes(filters.buyerName.toLowerCase())
    )
  }

  // فیلتر نام گیرنده
  if (filters.recipientName) {
    results = results.filter(card => 
      card.recipientName?.toLowerCase().includes(filters.recipientName.toLowerCase())
    )
  }

  // فیلتر ایمیل گیرنده
  if (filters.recipientEmail) {
    results = results.filter(card => 
      card.recipientEmail?.toLowerCase().includes(filters.recipientEmail.toLowerCase())
    )
  }

  // فیلتر شماره تماس گیرنده
  if (filters.recipientPhone) {
    results = results.filter(card => 
      card.recipientPhone?.includes(filters.recipientPhone)
    )
  }

  // فیلترهای تاریخ
  if (filters.createdDateFrom) {
    results = results.filter(card => 
      new Date(card.createdAt) >= new Date(filters.createdDateFrom)
    )
  }

  if (filters.createdDateTo) {
    results = results.filter(card => 
      new Date(card.createdAt) <= new Date(filters.createdDateTo)
    )
  }

  if (filters.expiryDateFrom) {
    results = results.filter(card => 
      new Date(card.expiryDate) >= new Date(filters.expiryDateFrom)
    )
  }

  if (filters.expiryDateTo) {
    results = results.filter(card => 
      new Date(card.expiryDate) <= new Date(filters.expiryDateTo)
    )
  }

  if (filters.deliveryDateFrom) {
    results = results.filter(card => 
      new Date(card.deliveryDate) >= new Date(filters.deliveryDateFrom)
    )
  }

  if (filters.deliveryDateTo) {
    results = results.filter(card => 
      new Date(card.deliveryDate) <= new Date(filters.deliveryDateTo)
    )
  }

  // فیلترهای مالی
  if (filters.amountMin !== null) {
    results = results.filter(card => card.amount >= filters.amountMin)
  }

  if (filters.amountMax !== null) {
    results = results.filter(card => card.amount <= filters.amountMax)
  }

  if (filters.remainingAmountMin !== null) {
    results = results.filter(card => card.remainingAmount >= filters.remainingAmountMin)
  }

  if (filters.remainingAmountMax !== null) {
    results = results.filter(card => card.remainingAmount <= filters.remainingAmountMax)
  }

  if (filters.usagePercentageMin !== null) {
    results = results.filter(card => {
      const usagePercentage = ((card.amount - card.remainingAmount) / card.amount) * 100
      return usagePercentage >= filters.usagePercentageMin
    })
  }

  if (filters.usagePercentageMax !== null) {
    results = results.filter(card => {
      const usagePercentage = ((card.amount - card.remainingAmount) / card.amount) * 100
      return usagePercentage <= filters.usagePercentageMax
    })
  }

  // فیلترهای وضعیت
  if (filters.statuses.length > 0) {
    results = results.filter(card => filters.statuses.includes(card.status))
  }

  // فیلترهای نوع
  if (filters.types.length > 0) {
    results = results.filter(card => filters.types.includes(card.type))
  }

  // فیلترهای دسته‌بندی
  if (filters.categories.length > 0) {
    results = results.filter(card => filters.categories.includes(card.category))
  }

  // فیلترهای روش تحویل
  if (filters.deliveryMethods.length > 0) {
    results = results.filter(card => filters.deliveryMethods.includes(card.deliveryMethod))
  }

  // فیلترهای پیشرفته
  if (filters.usageLimit) {
    results = results.filter(card => card.usageLimit === filters.usageLimit)
  }

  if (filters.requireVerification !== '') {
    results = results.filter(card => card.requireVerification === (filters.requireVerification === 'true'))
  }

  if (filters.allowPartialUsage !== '') {
    results = results.filter(card => card.allowPartialUsage === (filters.allowPartialUsage === 'true'))
  }

  if (filters.autoRenew !== '') {
    results = results.filter(card => card.autoRenew === (filters.autoRenew === 'true'))
  }

  if (filters.isUsed !== '') {
    results = results.filter(card => card.isUsed === (filters.isUsed === 'true'))
  }

  if (filters.isExpired !== '') {
    const now = new Date()
    if (filters.isExpired === 'true') {
      results = results.filter(card => new Date(card.expiryDate) < now)
    } else {
      results = results.filter(card => new Date(card.expiryDate) >= now)
    }
  }

  return results
})

const totalPages = computed(() => Math.ceil(filteredResults.value.length / pageSize.value))

const paginatedResults = computed(() => {
  const start = (currentPage.value - 1) * pageSize.value
  const end = start + pageSize.value
  return filteredResults.value.slice(start, end)
})

// Methods
const applyFilters = () => {
  currentPage.value = 1

}

const clearAllFilters = () => {
  Object.keys(filters).forEach(key => {
    if (Array.isArray(filters[key])) {
      filters[key] = []
    } else if (typeof filters[key] === 'number') {
      filters[key] = null
    } else {
      filters[key] = ''
    }
  })
  currentPage.value = 1

}

const exportResults = () => {

  // اینجا منطق خروجی Excel اضافه می‌شود
}

const viewCardDetails = (card: GiftCard) => {
  emit('view-details', card)
}

const editCard = (card: GiftCard) => {
  emit('edit-card', card)
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

const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

const getStatusColorClass = (status: string) => {
  const classes = {
    active: 'bg-green-500',
    used: 'bg-blue-500',
    expired: 'bg-red-500',
    suspended: 'bg-yellow-500',
    cancelled: 'bg-gray-500'
  }
  return classes[status] || 'bg-gray-500'
}

// Watchers
watch(filters, () => {
  currentPage.value = 1
}, { deep: true })

// Lifecycle
onMounted(() => {

})
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 

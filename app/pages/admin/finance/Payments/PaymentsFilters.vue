<template>
  <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
    <div class="flex items-center justify-between mb-6">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">فیلترها و تنظیمات</h3>
        <p class="text-sm text-gray-500 mt-1">جستجو و فیلتر کردن پرداخت‌ها بر اساس معیارهای مختلف</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button 
          @click="clearFilters"
          class="text-sm text-gray-500 hover:text-gray-700 transition-colors"
        >
          پاک کردن فیلترها
        </button>
        <button 
          @click="toggleAdvanced"
          class="inline-flex items-center text-sm text-blue-600 hover:text-blue-700 transition-colors"
        >
          <svg class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
          </svg>
          {{ showAdvanced ? 'فیلترهای ساده' : 'فیلترهای پیشرفته' }}
        </button>
      </div>
    </div>

    <!-- فیلترهای اصلی -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-6">
      <!-- بازه زمانی -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">بازه زمانی</label>
        <select 
          v-model="filters.dateRange"
          @change="applyFilters"
          class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
        >
          <option value="today">امروز</option>
          <option value="yesterday">دیروز</option>
          <option value="week">این هفته</option>
          <option value="lastWeek">هفته گذشته</option>
          <option value="month">این ماه</option>
          <option value="lastMonth">ماه گذشته</option>
          <option value="quarter">سه ماه گذشته</option>
          <option value="year">سال جاری</option>
          <option value="custom">انتخاب دستی</option>
        </select>
      </div>

      <!-- وضعیت پرداخت -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت پرداخت</label>
        <select 
          v-model="filters.status"
          @change="applyFilters"
          class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
        >
          <option value="">همه وضعیت‌ها</option>
          <option value="success">موفق</option>
          <option value="failed">ناموفق</option>
          <option value="pending">در انتظار</option>
          <option value="cancelled">لغو شده</option>
          <option value="refunded">مرجوع شده</option>
        </select>
      </div>

      <!-- روش پرداخت -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">روش پرداخت</label>
        <select 
          v-model="filters.paymentMethod"
          @change="applyFilters"
          class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
        >
          <option value="">همه روش‌ها</option>
          <option value="card">کارت به کارت</option>
          <option value="gateway">درگاه بانکی</option>
          <option value="wallet">کیف پول</option>
          <option value="cash">پرداخت نقدی</option>
          <option value="check">چک</option>
        </select>
      </div>

      <!-- مبلغ -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">مبلغ (تومان)</label>
        <select 
          v-model="filters.amountRange"
          @change="applyFilters"
          class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
        >
          <option value="">همه مبالغ</option>
          <option value="0-100000">کمتر از ۱۰۰ هزار</option>
          <option value="100000-500000">۱۰۰ تا ۵۰۰ هزار</option>
          <option value="500000-1000000">۵۰۰ هزار تا ۱ میلیون</option>
          <option value="1000000-5000000">۱ تا ۵ میلیون</option>
          <option value="5000000+">بیش از ۵ میلیون</option>
        </select>
      </div>
    </div>

    <!-- فیلترهای پیشرفته -->
    <div v-if="showAdvanced" class="border-t border-gray-200 pt-6">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-6">
        <!-- شماره سفارش -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">شماره سفارش</label>
          <input 
            v-model="filters.orderId"
            type="text"
            placeholder="مثال: ORD-1234"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            @input="applyFilters"
          />
        </div>

        <!-- شماره تراکنش -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">شماره تراکنش</label>
          <input 
            v-model="filters.transactionId"
            type="text"
            placeholder="مثال: TRX-5678"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            @input="applyFilters"
          />
        </div>

        <!-- نام مشتری -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نام مشتری</label>
          <input 
            v-model="filters.customerName"
            type="text"
            placeholder="جستجو بر اساس نام"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            @input="applyFilters"
          />
        </div>

        <!-- شماره موبایل -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">شماره موبایل</label>
          <input 
            v-model="filters.phoneNumber"
            type="text"
            placeholder="مثال: 09123456789"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            @input="applyFilters"
          />
        </div>

        <!-- بانک -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">بانک</label>
          <select 
            v-model="filters.bank"
            @change="applyFilters"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه بانک‌ها</option>
            <option value="melli">ملی</option>
            <option value="mellat">ملت</option>
            <option value="parsian">پارسیان</option>
            <option value="saman">سامان</option>
            <option value="pasargad">پاسارگاد</option>
            <option value="saderat">صادرات</option>
            <option value="keshavarzi">کشاورزی</option>
          </select>
        </div>

        <!-- درگاه پرداخت -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">درگاه پرداخت</label>
          <select 
            v-model="filters.gateway"
            @change="applyFilters"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه درگاه‌ها</option>
            <option value="zarinpal">زرین‌پال</option>
            <option value="nextpay">نکست‌پی</option>
            <option value="idpay">آیدی‌پی</option>
            <option value="payir">پی‌ایر</option>
            <option value="payping">پی‌پینگ</option>
          </select>
        </div>
      </div>

      <!-- بازه مبلغ دقیق -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حداقل مبلغ (تومان)</label>
          <input 
            v-model="filters.minAmount"
            type="number"
            placeholder="0"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            @input="applyFilters"
          />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر مبلغ (تومان)</label>
          <input 
            v-model="filters.maxAmount"
            type="number"
            placeholder="نامحدود"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            @input="applyFilters"
          />
        </div>
      </div>

      <!-- بازه تاریخ دقیق -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">از تاریخ</label>
          <input 
            v-model="filters.startDate"
            type="date"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            @change="applyFilters"
          />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تا تاریخ</label>
          <input 
            v-model="filters.endDate"
            type="date"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            @change="applyFilters"
          />
        </div>
      </div>
    </div>

    <!-- دکمه‌های عملیات -->
    <div class="flex items-center justify-between pt-4 border-t border-gray-200">
      <div class="flex items-center space-x-4 space-x-reverse">
        <span class="text-sm text-gray-500">
          {{ activeFiltersCount }} فیلتر فعال
        </span>
        <span class="text-sm text-gray-500">
          {{ filteredResultsCount }} نتیجه یافت شد
        </span>
      </div>
      
      <div class="flex items-center space-x-3 space-x-reverse">
        <button 
          @click="saveFilterPreset"
          class="inline-flex items-center px-3 py-2 text-sm text-gray-600 bg-gray-100 rounded-lg hover:bg-gray-200 transition-colors"
        >
          <svg class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z"></path>
          </svg>
          ذخیره فیلتر
        </button>
        
        <button 
          @click="applyFilters"
          class="inline-flex items-center px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 transition-colors"
        >
          <svg class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
          </svg>
          اعمال فیلتر
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch, onMounted } from 'vue'

// Props
interface Props {
  totalResults?: number
}

const props = withDefaults(defineProps<Props>(), {
  totalResults: 0
})

// Emits
const emit = defineEmits<{
  filter: [filters: any]
}>()

// Reactive data
const showAdvanced = ref(false)
const filteredResultsCount = ref(props.totalResults)

// فیلترها
const filters = ref({
  dateRange: 'week',
  status: '',
  paymentMethod: '',
  amountRange: '',
  orderId: '',
  transactionId: '',
  customerName: '',
  phoneNumber: '',
  bank: '',
  gateway: '',
  minAmount: '',
  maxAmount: '',
  startDate: '',
  endDate: ''
})

// Computed
const activeFiltersCount = computed(() => {
  let count = 0
  Object.values(filters.value).forEach(value => {
    if (value && value !== 'week') count++
  })
  return count
})

// Methods
const applyFilters = () => {
  emit('filter', { ...filters.value })
  // شبیه‌سازی تعداد نتایج فیلتر شده
  filteredResultsCount.value = Math.floor(Math.random() * 500) + 50
}

const clearFilters = () => {
  filters.value = {
    dateRange: 'week',
    status: '',
    paymentMethod: '',
    amountRange: '',
    orderId: '',
    transactionId: '',
    customerName: '',
    phoneNumber: '',
    bank: '',
    gateway: '',
    minAmount: '',
    maxAmount: '',
    startDate: '',
    endDate: ''
  }
  applyFilters()
}

const toggleAdvanced = () => {
  showAdvanced.value = !showAdvanced.value
}

const saveFilterPreset = () => {
  // TODO: ذخیره فیلتر به عنوان پیش‌فرض
  console.log('Saving filter preset:', filters.value)
}

// Watch for prop changes
watch(() => props.totalResults, (newValue) => {
  filteredResultsCount.value = newValue
})

// Apply filters on mount
onMounted(() => {
  applyFilters()
})
</script>

<!--
  کامپوننت فیلترهای پیشرفته صفحه پرداخت‌ها
  - فیلترهای اصلی (زمان، وضعیت، روش، مبلغ)
  - فیلترهای پیشرفته (شماره سفارش، مشتری، بانک و...)
  - نمایش تعداد فیلترهای فعال
  - ذخیره فیلتر به عنوان پیش‌فرض
  - کاملاً ریسپانسیو
  توضیحات کامل به فارسی در کد
--> 

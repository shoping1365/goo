<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="flex justify-between items-center">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">مدیریت قیمت‌گذاری</h3>
        <p class="text-sm text-gray-600 mt-1">تنظیم قیمت‌ها و تخفیف‌های گیفت کارت</p>
      </div>
      <button 
        @click="showCreateModal = true"
        class="px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
      >
        افزودن قیمت جدید
      </button>
    </div>

    <!-- کارت‌های آمار -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-blue-100 rounded-lg flex items-center justify-center">
              <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1" />
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">میانگین قیمت</p>
            <p class="text-2xl font-semibold text-gray-900">{{ formatAmount(averagePrice) }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-green-100 rounded-lg flex items-center justify-center">
              <svg class="w-5 h-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" />
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">بیشترین فروش</p>
            <p class="text-2xl font-semibold text-gray-900">{{ formatAmount(bestSellingPrice) }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-yellow-100 rounded-lg flex items-center justify-center">
              <svg class="w-5 h-5 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z" />
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">تخفیف فعال</p>
            <p class="text-2xl font-semibold text-gray-900">{{ activeDiscounts }}%</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-purple-100 rounded-lg flex items-center justify-center">
              <svg class="w-5 h-5 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">درآمد کل</p>
            <p class="text-2xl font-semibold text-gray-900">{{ formatAmount(totalRevenue) }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- جدول قیمت‌ها -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <div class="flex justify-between items-center">
          <h4 class="text-lg font-medium text-gray-900">لیست قیمت‌ها</h4>
          <div class="flex items-center space-x-3 space-x-reverse">
            <select v-model="priceFilter" class="px-3 py-2 text-sm border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
              <option value="">همه دسته‌ها</option>
              <option value="birthday">تولد</option>
              <option value="wedding">عروسی</option>
              <option value="newyear">سال نو</option>
              <option value="corporate">شرکتی</option>
              <option value="general">عمومی</option>
            </select>
            <select v-model="statusFilter" class="px-3 py-2 text-sm border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
              <option value="">همه وضعیت‌ها</option>
              <option value="active">فعال</option>
              <option value="inactive">غیرفعال</option>
            </select>
          </div>
        </div>
      </div>

      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">دسته‌بندی</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مبلغ پایه</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تخفیف</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">قیمت نهایی</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="price in filteredPrices" :key="price.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <span :class="getCategoryClasses(price.category)" class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium">
                    {{ getCategoryText(price.category) }}
                  </span>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                {{ formatAmount(price.basePrice) }} تومان
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                <span v-if="price.discount > 0" class="text-green-600 font-medium">{{ price.discount }}%</span>
                <span v-else class="text-gray-500">بدون تخفیف</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                {{ formatAmount(price.finalPrice) }} تومان
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getStatusClasses(price.status)" class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium">
                  {{ getStatusText(price.status) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <div class="flex items-center space-x-2 space-x-reverse">
                  <button 
                    @click="editPrice(price)"
                    class="text-blue-600 hover:text-blue-900 p-1 rounded hover:bg-blue-50"
                    title="ویرایش"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                    </svg>
                  </button>
                  <button 
                    @click="togglePriceStatus(price)"
                    :class="price.status === 'active' ? 'text-yellow-600 hover:text-yellow-900' : 'text-green-600 hover:text-green-900'"
                    class="p-1 rounded hover:bg-gray-50"
                    :title="price.status === 'active' ? 'غیرفعال کردن' : 'فعال کردن'"
                  >
                    <svg v-if="price.status === 'active'" class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728L5.636 5.636m12.728 12.728L18.364 5.636M5.636 18.364l12.728-12.728" />
                    </svg>
                    <svg v-else class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                    </svg>
                  </button>
                  <button 
                    @click="deletePrice(price)"
                    class="text-red-600 hover:text-red-900 p-1 rounded hover:bg-red-50"
                    title="حذف"
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
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

// Reactive data
const showCreateModal = ref(false)
const priceFilter = ref('')
const statusFilter = ref('')

// داده‌های نمونه
const prices = ref([
  {
    id: 1,
    category: 'birthday',
    basePrice: 500000,
    discount: 10,
    finalPrice: 450000,
    status: 'active',
    createdAt: new Date(Date.now() - 86400000)
  },
  {
    id: 2,
    category: 'wedding',
    basePrice: 1000000,
    discount: 15,
    finalPrice: 850000,
    status: 'active',
    createdAt: new Date(Date.now() - 172800000)
  },
  {
    id: 3,
    category: 'newyear',
    basePrice: 200000,
    discount: 0,
    finalPrice: 200000,
    status: 'inactive',
    createdAt: new Date(Date.now() - 259200000)
  },
  {
    id: 4,
    category: 'corporate',
    basePrice: 2000000,
    discount: 20,
    finalPrice: 1600000,
    status: 'active',
    createdAt: new Date(Date.now() - 345600000)
  }
])

// Computed properties
const filteredPrices = computed(() => {
  let filtered = prices.value

  if (priceFilter.value) {
    filtered = filtered.filter(price => price.category === priceFilter.value)
  }

  if (statusFilter.value) {
    filtered = filtered.filter(price => price.status === statusFilter.value)
  }

  return filtered
})

const averagePrice = computed(() => {
  const activePrices = prices.value.filter(price => price.status === 'active')
  if (activePrices.length === 0) return 0
  const total = activePrices.reduce((sum, price) => sum + price.finalPrice, 0)
  return Math.round(total / activePrices.length)
})

const bestSellingPrice = computed(() => {
  // در نسخه واقعی: محاسبه بر اساس آمار فروش
  return 500000
})

const activeDiscounts = computed(() => {
  const activePrices = prices.value.filter(price => price.status === 'active' && price.discount > 0)
  if (activePrices.length === 0) return 0
  const totalDiscount = activePrices.reduce((sum, price) => sum + price.discount, 0)
  return Math.round(totalDiscount / activePrices.length)
})

const totalRevenue = computed(() => {
  // در نسخه واقعی: محاسبه بر اساس آمار فروش
  return 45678000
})

// Methods
const formatAmount = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount)
}

const getCategoryText = (category: string) => {
  const categoryMap = {
    birthday: 'تولد',
    wedding: 'عروسی',
    newyear: 'سال نو',
    corporate: 'شرکتی',
    general: 'عمومی'
  }
  return categoryMap[category] || category
}

const getCategoryClasses = (category: string) => {
  const classesMap = {
    birthday: 'bg-pink-100 text-pink-800',
    wedding: 'bg-purple-100 text-purple-800',
    newyear: 'bg-green-100 text-green-800',
    corporate: 'bg-blue-100 text-blue-800',
    general: 'bg-gray-100 text-gray-800'
  }
  return classesMap[category] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status: string) => {
  const statusMap = {
    active: 'فعال',
    inactive: 'غیرفعال'
  }
  return statusMap[status] || status
}

const getStatusClasses = (status: string) => {
  const classesMap = {
    active: 'bg-green-100 text-green-800',
    inactive: 'bg-gray-100 text-gray-800'
  }
  return classesMap[status] || 'bg-gray-100 text-gray-800'
}

const editPrice = (price: any) => {
  // در نسخه واقعی: باز کردن مودال ویرایش
  console.log('Edit price:', price)
}

const togglePriceStatus = async (price: any) => {
  try {
    price.status = price.status === 'active' ? 'inactive' : 'active'
    // در نسخه واقعی: ارسال به API
    console.log('Toggling price status:', price.id, price.status)
    alert(`قیمت ${price.status === 'active' ? 'فعال' : 'غیرفعال'} شد`)
  } catch (error) {
    console.error('خطا در تغییر وضعیت قیمت:', error)
    alert('خطا در تغییر وضعیت قیمت')
  }
}

const deletePrice = async (price: any) => {
  if (confirm(`آیا از حذف قیمت ${getCategoryText(price.category)} مطمئن هستید؟`)) {
    try {
      const index = prices.value.findIndex(p => p.id === price.id)
      if (index > -1) {
        prices.value.splice(index, 1)
      }
      // در نسخه واقعی: ارسال به API
      console.log('Deleting price:', price.id)
      alert('قیمت حذف شد')
    } catch (error) {
      console.error('خطا در حذف قیمت:', error)
      alert('خطا در حذف قیمت')
    }
  }
}
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 

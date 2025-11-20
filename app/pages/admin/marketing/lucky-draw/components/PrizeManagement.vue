<template>
  <div class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
    <div class="relative top-20 mx-auto p-5 border w-11/12 md:w-4/5 lg:w-3/4 shadow-lg rounded-md bg-white">
      <div class="mt-3">
        <!-- هدر مودال -->
        <div class="flex justify-between items-center mb-6">
          <h3 class="text-lg font-medium text-gray-900">مدیریت جوایز</h3>
          <div class="flex items-center space-x-3 space-x-reverse">
            <button 
              class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors"
              @click="showPrizeForm = true"
            >
              افزودن جایزه جدید
            </button>
            <button 
              class="text-gray-400 hover:text-gray-600"
              @click="$emit('close')"
            >
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
              </svg>
            </button>
          </div>
        </div>

        <!-- آمار جوایز -->
        <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
          <div class="bg-blue-50 p-6 rounded-lg border border-blue-200">
            <div class="text-center">
              <p class="text-sm text-blue-600">کل جوایز</p>
              <p class="text-2xl font-bold text-blue-900">{{ prizeStats.total }}</p>
            </div>
          </div>
          <div class="bg-green-50 p-6 rounded-lg border border-green-200">
            <div class="text-center">
              <p class="text-sm text-green-600">جوایز فعال</p>
              <p class="text-2xl font-bold text-green-900">{{ prizeStats.active }}</p>
            </div>
          </div>
          <div class="bg-purple-50 p-6 rounded-lg border border-purple-200">
            <div class="text-center">
              <p class="text-sm text-purple-600">اهدا شده</p>
              <p class="text-2xl font-bold text-purple-900">{{ prizeStats.awarded }}</p>
            </div>
          </div>
          <div class="bg-yellow-50 p-6 rounded-lg border border-yellow-200">
            <div class="text-center">
              <p class="text-sm text-yellow-600">ارزش کل</p>
              <p class="text-2xl font-bold text-yellow-900">{{ formatCurrency(prizeStats.totalValue) }}</p>
            </div>
          </div>
        </div>

        <!-- جدول جوایز -->
        <div class="bg-white rounded-lg border">
          <div class="overflow-x-auto">
            <table class="min-w-full divide-y divide-gray-200">
              <thead class="bg-gray-50">
                <tr>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نام جایزه</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">ارزش</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تعداد</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
                </tr>
              </thead>
              <tbody class="bg-white divide-y divide-gray-200">
                <tr v-for="prize in prizes" :key="prize.id">
                  <td class="px-6 py-4 whitespace-nowrap">
                    <div class="flex items-center">
                      <div class="w-10 h-10 bg-gray-200 rounded-lg flex items-center justify-center mr-3">
                        <svg class="w-6 h-6 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v13m0-13V6a2 2 0 112 2h-2zm0 0V5.5A2.5 2.5 0 109.5 8H12z"></path>
                        </svg>
                      </div>
                      <div>
                        <div class="text-sm font-medium text-gray-900">{{ prize.name }}</div>
                        <div class="text-sm text-gray-500">{{ prize.description }}</div>
                      </div>
                    </div>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap">
                    <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full" :class="getPrizeTypeBadgeClass(prize.type)">
                      {{ getPrizeTypeText(prize.type) }}
                    </span>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ formatCurrency(prize.value) }}</td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ prize.quantity }}</td>
                  <td class="px-6 py-4 whitespace-nowrap">
                    <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full" :class="getStatusBadgeClass(prize.status)">
                      {{ getStatusText(prize.status) }}
                    </span>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                    <button class="text-blue-600 hover:text-blue-900 ml-3" @click="editPrize(prize)">ویرایش</button>
                    <button class="text-red-600 hover:text-red-900" @click="deletePrize(prize.id)">حذف</button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <!-- مودال فرم جایزه -->
        <PrizeForm 
          v-if="showPrizeForm" 
          :prize="editingPrize"
          @close="closePrizeForm"
          @saved="onPrizeSaved"
        />
      </div>
    </div>
  </div>
</template>

<script setup>
import PrizeForm from './PrizeForm.vue'

const emit = defineEmits(['close'])

// متغیرهای reactive
const showPrizeForm = ref(false)
const editingPrize = ref(null)

// آمار جوایز
const prizeStats = ref({
  total: 25,
  active: 20,
  awarded: 156,
  totalValue: 12500000
})

// لیست جوایز
const prizes = ref([
  {
    id: 1,
    name: 'کارت هدیه 500 هزار تومانی',
    description: 'کارت هدیه برای خرید از فروشگاه',
    type: 'gift_card',
    value: 500000,
    quantity: 50,
    status: 'active'
  },
  {
    id: 2,
    name: 'محصول رایگان',
    description: 'انتخاب یک محصول به دلخواه',
    type: 'free_product',
    value: 0,
    quantity: 100,
    status: 'active'
  },
  {
    id: 3,
    name: 'تخفیف 30 درصدی',
    description: 'تخفیف روی کل خرید',
    type: 'discount',
    value: 0,
    quantity: 200,
    status: 'active'
  },
  {
    id: 4,
    name: 'ارسال رایگان',
    description: 'ارسال رایگان برای یک ماه',
    type: 'free_shipping',
    value: 0,
    quantity: 150,
    status: 'active'
  }
])

// توابع کمکی
const formatCurrency = (amount) => {
  if (amount === 0) return 'رایگان'
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

const getPrizeTypeBadgeClass = (type) => {
  const classes = {
    gift_card: 'bg-blue-100 text-blue-800',
    free_product: 'bg-green-100 text-green-800',
    discount: 'bg-purple-100 text-purple-800',
    free_shipping: 'bg-orange-100 text-orange-800',
    cash: 'bg-yellow-100 text-yellow-800'
  }
  return classes[type] || 'bg-gray-100 text-gray-800'
}

const getPrizeTypeText = (type) => {
  const texts = {
    gift_card: 'کارت هدیه',
    free_product: 'محصول رایگان',
    discount: 'تخفیف',
    free_shipping: 'ارسال رایگان',
    cash: 'نقدی'
  }
  return texts[type] || type
}

const getStatusBadgeClass = (status) => {
  const classes = {
    active: 'bg-green-100 text-green-800',
    inactive: 'bg-red-100 text-red-800',
    draft: 'bg-gray-100 text-gray-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status) => {
  const texts = {
    active: 'فعال',
    inactive: 'غیرفعال',
    draft: 'پیش‌نویس'
  }
  return texts[status] || status
}

// توابع مدیریت
const editPrize = (prize) => {
  editingPrize.value = { ...prize }
  showPrizeForm.value = true
}

const deletePrize = (id) => {
  if (confirm('آیا از حذف این جایزه اطمینان دارید؟')) {
    prizes.value = prizes.value.filter(prize => prize.id !== id)
    prizeStats.value.total--
    
    // به‌روزرسانی آمار
    updatePrizeStats()
  }
}

const closePrizeForm = () => {
  showPrizeForm.value = false
  editingPrize.value = null
}

const onPrizeSaved = (savedPrize) => {
  if (editingPrize.value) {
    // ویرایش
    const index = prizes.value.findIndex(prize => prize.id === savedPrize.id)
    if (index !== -1) {
      prizes.value[index] = savedPrize
    }
  } else {
    // ایجاد جدید
    savedPrize.id = Date.now()
    prizes.value.unshift(savedPrize)
    prizeStats.value.total++
  }
  
  // به‌روزرسانی آمار
  updatePrizeStats()
  closePrizeForm()
}

const updatePrizeStats = () => {
  prizeStats.value.active = prizes.value.filter(p => p.status === 'active').length
  prizeStats.value.totalValue = prizes.value.reduce((sum, p) => sum + (p.value * p.quantity), 0)
}
</script>

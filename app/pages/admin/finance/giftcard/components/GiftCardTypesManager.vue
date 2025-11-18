<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">مدیریت انواع گیفت کارت</h3>
        <p class="text-sm text-gray-600">تنظیم و مدیریت انواع مختلف گیفت کارت</p>
      </div>
      <button
        @click="showCreateModal = true"
        class="inline-flex items-center px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
      >
        <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path>
        </svg>
        افزودن نوع جدید
      </button>
    </div>

    <!-- آمار کلی انواع -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div class="bg-white p-6 rounded-lg border border-gray-200">
        <div class="flex items-center">
          <div class="p-2 bg-blue-100 rounded-lg">
            <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21a4 4 0 01-4-4V5a2 2 0 012-2h4a2 2 0 012 2v12a4 4 0 01-4 4zM21 5a2 2 0 00-2-2h-4a2 2 0 00-2 2v12a4 4 0 004 4h4a2 2 0 002-2V5z"></path>
            </svg>
          </div>
          <div class="ml-3">
            <p class="text-sm font-medium text-gray-600">کل انواع</p>
            <p class="text-2xl font-bold text-gray-900">{{ types.length }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white p-6 rounded-lg border border-gray-200">
        <div class="flex items-center">
          <div class="p-2 bg-green-100 rounded-lg">
            <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
          <div class="ml-3">
            <p class="text-sm font-medium text-gray-600">فعال</p>
            <p class="text-2xl font-bold text-gray-900">{{ activeTypes.length }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white p-6 rounded-lg border border-gray-200">
        <div class="flex items-center">
          <div class="p-2 bg-purple-100 rounded-lg">
            <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
            </svg>
          </div>
          <div class="ml-3">
            <p class="text-sm font-medium text-gray-600">متغیر</p>
            <p class="text-2xl font-bold text-gray-900">{{ variableTypes.length }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white p-6 rounded-lg border border-gray-200">
        <div class="flex items-center">
          <div class="p-2 bg-orange-100 rounded-lg">
            <svg class="w-6 h-6 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"></path>
            </svg>
          </div>
          <div class="ml-3">
            <p class="text-sm font-medium text-gray-600">موضوعی</p>
            <p class="text-2xl font-bold text-gray-900">{{ themedTypes.length }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- فیلترها و جستجو -->
    <div class="bg-white p-6 rounded-lg border border-gray-200">
      <div class="flex flex-col sm:flex-row gap-6">
        <div class="flex-1">
          <input
            v-model="searchQuery"
            type="text"
            placeholder="جستجو در انواع..."
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          />
        </div>
        <div class="flex gap-2">
          <select
            v-model="statusFilter"
            class="px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          >
            <option value="">همه وضعیت‌ها</option>
            <option value="active">فعال</option>
            <option value="inactive">غیرفعال</option>
          </select>
          <select
            v-model="categoryFilter"
            class="px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          >
            <option value="">همه دسته‌ها</option>
            <option value="fixed">مبلغ ثابت</option>
            <option value="variable">مبلغ متغیر</option>
            <option value="themed">موضوعی</option>
            <option value="corporate">شرکتی</option>
            <option value="discount">تخفیف ویژه</option>
          </select>
        </div>
      </div>
    </div>

    <!-- لیست انواع -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div
        v-for="type in filteredTypes"
        :key="type.id"
        class="bg-white rounded-lg border border-gray-200 overflow-hidden hover:shadow-lg transition-shadow"
      >
        <!-- هدر کارت -->
        <div class="relative h-32 bg-gradient-to-br from-blue-400 to-purple-600">
          <div class="absolute inset-0 flex items-center justify-center">
            <div class="text-white text-center">
              <div class="text-3xl mb-2">{{ type.icon }}</div>
              <h3 class="text-lg font-semibold">{{ type.name }}</h3>
            </div>
          </div>
          <div class="absolute top-2 right-2">
            <span
              :class="{
                'bg-green-100 text-green-800': type.status === 'active',
                'bg-red-100 text-red-800': type.status === 'inactive'
              }"
              class="px-2 py-1 text-xs font-medium rounded-full"
            >
              {{ getStatusText(type.status) }}
            </span>
          </div>
        </div>

        <!-- اطلاعات نوع -->
        <div class="p-6">
          <div class="flex items-center justify-between mb-3">
            <h4 class="text-lg font-semibold text-gray-900">{{ type.name }}</h4>
            <span class="text-sm text-gray-500">{{ type.category }}</span>
          </div>
          
          <p class="text-sm text-gray-600 mb-4">{{ type.description }}</p>

          <!-- تنظیمات مبلغ -->
          <div class="space-y-2 mb-4">
            <div class="flex items-center justify-between text-sm">
              <span class="text-gray-600">نوع مبلغ:</span>
              <span class="text-gray-900">{{ getAmountTypeText(type.amountType) }}</span>
            </div>
            <div v-if="type.amountType === 'fixed'" class="flex items-center justify-between text-sm">
              <span class="text-gray-600">مبلغ ثابت:</span>
              <span class="text-gray-900 font-medium">{{ formatCurrency(type.fixedAmount) }}</span>
            </div>
            <div v-if="type.amountType === 'variable'" class="space-y-1">
              <div class="flex items-center justify-between text-sm">
                <span class="text-gray-600">حداقل:</span>
                <span class="text-gray-900 font-medium">{{ formatCurrency(type.minAmount) }}</span>
              </div>
              <div class="flex items-center justify-between text-sm">
                <span class="text-gray-600">حداکثر:</span>
                <span class="text-gray-900 font-medium">{{ formatCurrency(type.maxAmount) }}</span>
              </div>
            </div>
            <div class="flex items-center justify-between text-sm">
              <span class="text-gray-600">تاریخ انقضا:</span>
              <span class="text-gray-900">{{ type.expiryDays }} روز</span>
            </div>
          </div>

          <!-- آمار استفاده -->
          <div class="grid grid-cols-2 gap-6 mb-4">
            <div class="text-center">
              <p class="text-2xl font-bold text-blue-600">{{ type.usageCount }}</p>
              <p class="text-xs text-gray-500">استفاده شده</p>
            </div>
            <div class="text-center">
              <p class="text-2xl font-bold text-green-600">{{ type.totalValue }}</p>
              <p class="text-xs text-gray-500">کل ارزش (هزار)</p>
            </div>
          </div>

          <!-- دکمه‌های عملیات -->
          <div class="flex gap-2">
            <button
              @click="editType(type)"
              class="flex-1 px-3 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
            >
              ویرایش
            </button>
            <button
              @click="viewTypeDetails(type)"
              class="flex-1 px-3 py-2 bg-gray-600 text-white text-sm font-medium rounded-lg hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
            >
              جزئیات
            </button>
            <button
              @click="deleteType(type)"
              class="px-3 py-2 bg-red-600 text-white text-sm font-medium rounded-lg hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
              </svg>
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- مودال ایجاد نوع جدید -->
    <GiftCardTypeCreateModal
      v-if="showCreateModal"
      @close="showCreateModal = false"
      @created="handleTypeCreated"
    />

    <!-- مودال ویرایش نوع -->
    <GiftCardTypeEditModal
      v-if="showEditModal"
      :type="selectedType"
      @close="showEditModal = false"
      @updated="handleTypeUpdated"
    />
  </div>
</template>

<script setup>
// تعریف متغیرهای reactive
const searchQuery = ref('')
const statusFilter = ref('')
const categoryFilter = ref('')
const showCreateModal = ref(false)
const showEditModal = ref(false)
const selectedType = ref(null)

// داده‌های نمونه انواع گیفت کارت
const types = ref([
  {
    id: 1,
    name: 'گیفت کارت تولد',
    description: 'کارت‌های ویژه برای جشن تولد با مبلغ ثابت',
    category: 'موضوعی',
    status: 'active',
    icon: '🎂',
    amountType: 'fixed',
    fixedAmount: 500000,
    minAmount: null,
    maxAmount: null,
    expiryDays: 365,
    usageCount: 1250,
    totalValue: 625000,
    themes: ['تولد', 'جشن'],
    isDefault: true,
    allowCustomMessage: true,
    allowRecipientEmail: true,
    autoSendEmail: true,
    emailTemplate: 'birthday_template'
  },
  {
    id: 2,
    name: 'گیفت کارت عروسی',
    description: 'کارت‌های مجلل برای مراسم عروسی',
    category: 'موضوعی',
    status: 'active',
    icon: '💒',
    amountType: 'variable',
    fixedAmount: null,
    minAmount: 1000000,
    maxAmount: 10000000,
    expiryDays: 730,
    usageCount: 890,
    totalValue: 4450000,
    themes: ['عروسی', 'مجلل'],
    isDefault: false,
    allowCustomMessage: true,
    allowRecipientEmail: true,
    autoSendEmail: true,
    emailTemplate: 'wedding_template'
  },
  {
    id: 3,
    name: 'گیفت کارت سال نو',
    description: 'کارت‌های شاد برای تبریک سال نو',
    category: 'موضوعی',
    status: 'active',
    icon: '🎊',
    amountType: 'fixed',
    fixedAmount: 200000,
    minAmount: null,
    maxAmount: null,
    expiryDays: 180,
    usageCount: 2100,
    totalValue: 420000,
    themes: ['سال نو', 'جشن'],
    isDefault: true,
    allowCustomMessage: true,
    allowRecipientEmail: true,
    autoSendEmail: true,
    emailTemplate: 'newyear_template'
  },
  {
    id: 4,
    name: 'گیفت کارت شرکتی',
    description: 'کارت‌های رسمی برای استفاده شرکتی',
    category: 'شرکتی',
    status: 'active',
    icon: '🏢',
    amountType: 'variable',
    fixedAmount: null,
    minAmount: 500000,
    maxAmount: 5000000,
    expiryDays: 365,
    usageCount: 650,
    totalValue: 1950000,
    themes: ['شرکتی', 'رسمی'],
    isDefault: false,
    allowCustomMessage: false,
    allowRecipientEmail: true,
    autoSendEmail: false,
    emailTemplate: 'corporate_template'
  },
  {
    id: 5,
    name: 'گیفت کارت تخفیف ویژه',
    description: 'کارت‌های تخفیف با درصد مشخص',
    category: 'تخفیف',
    status: 'active',
    icon: '🎯',
    amountType: 'fixed',
    fixedAmount: 100000,
    minAmount: null,
    maxAmount: null,
    expiryDays: 90,
    usageCount: 320,
    totalValue: 32000,
    themes: ['تخفیف', 'ویژه'],
    isDefault: false,
    allowCustomMessage: false,
    allowRecipientEmail: true,
    autoSendEmail: true,
    emailTemplate: 'discount_template'
  },
  {
    id: 6,
    name: 'گیفت کارت عمومی',
    description: 'کارت‌های عمومی برای استفاده آزاد',
    category: 'مبلغ ثابت',
    status: 'inactive',
    icon: '🎁',
    amountType: 'variable',
    fixedAmount: null,
    minAmount: 50000,
    maxAmount: 1000000,
    expiryDays: 365,
    usageCount: 180,
    totalValue: 90000,
    themes: ['عمومی'],
    isDefault: false,
    allowCustomMessage: true,
    allowRecipientEmail: true,
    autoSendEmail: true,
    emailTemplate: 'general_template'
  }
])

// محاسبه انواع فعال
const activeTypes = computed(() => types.value.filter(type => type.status === 'active'))

// محاسبه انواع متغیر
const variableTypes = computed(() => types.value.filter(type => type.amountType === 'variable'))

// محاسبه انواع موضوعی
const themedTypes = computed(() => types.value.filter(type => type.category === 'موضوعی'))

// فیلتر کردن انواع بر اساس جستجو و فیلترها
const filteredTypes = computed(() => {
  return types.value.filter(type => {
    const matchesSearch = type.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
                         type.description.toLowerCase().includes(searchQuery.value.toLowerCase())
    
    const matchesStatus = !statusFilter.value || type.status === statusFilter.value
    
    const matchesCategory = !categoryFilter.value || type.category === categoryFilter.value
    
    return matchesSearch && matchesStatus && matchesCategory
  })
})

// تابع دریافت متن وضعیت
const getStatusText = (status) => {
  const statusMap = {
    active: 'فعال',
    inactive: 'غیرفعال'
  }
  return statusMap[status] || status
}

// تابع دریافت متن نوع مبلغ
const getAmountTypeText = (amountType) => {
  const amountTypeMap = {
    fixed: 'مبلغ ثابت',
    variable: 'مبلغ متغیر'
  }
  return amountTypeMap[amountType] || amountType
}

// تابع فرمت کردن مبلغ
const formatCurrency = (amount) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

// تابع ویرایش نوع
const editType = (type) => {
  selectedType.value = type
  showEditModal.value = true
}

// تابع مشاهده جزئیات نوع
const viewTypeDetails = (type) => {
  // اینجا می‌توانید مودال جزئیات را باز کنید
  console.log('View type details:', type)
}

// تابع حذف نوع
const deleteType = (type) => {
  if (confirm(`آیا از حذف نوع "${type.name}" اطمینان دارید؟`)) {
    const index = types.value.findIndex(t => t.id === type.id)
    if (index > -1) {
      types.value.splice(index, 1)
      // نمایش پیام موفقیت
      alert('نوع با موفقیت حذف شد')
    }
  }
}

// تابع مدیریت ایجاد نوع جدید
const handleTypeCreated = (newType) => {
  types.value.push({
    ...newType,
    id: Date.now(),
    usageCount: 0,
    totalValue: 0
  })
  showCreateModal.value = false
}

// تابع مدیریت بروزرسانی نوع
const handleTypeUpdated = (updatedType) => {
  const index = types.value.findIndex(t => t.id === updatedType.id)
  if (index > -1) {
    types.value[index] = { ...types.value[index], ...updatedType }
  }
  showEditModal.value = false
}
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 

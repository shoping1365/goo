<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex justify-between items-center">
        <div>
          <h2 class="text-xl font-bold text-gray-900">فیلترهای ذخیره شده</h2>
          <p class="text-gray-600 mt-1">مدیریت فیلترهای ذخیره شده برای جستجوی سریع</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button 
            class="px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
            @click="showSaveFilterModal = true"
          >
            ذخیره فیلتر جدید
          </button>
        </div>
      </div>
    </div>

    <!-- لیست فیلترهای ذخیره شده -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div 
        v-for="filter in savedFilters" 
        :key="filter.id"
        class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 hover:shadow-md transition-shadow duration-200"
      >
        <div class="flex items-start justify-between mb-4">
          <div class="flex items-center space-x-2 space-x-reverse">
            <div :class="getFilterIconClass(filter.type)" class="w-8 h-8 rounded-lg flex items-center justify-center">
              <component :is="getFilterIcon(filter.type)" class="w-4 h-4" />
            </div>
            <div>
              <h3 class="text-sm font-medium text-gray-900">{{ filter.name }}</h3>
              <p class="text-xs text-gray-500">{{ filter.description }}</p>
            </div>
          </div>
          <div class="flex items-center space-x-1 space-x-reverse">
            <button 
              class="text-gray-400 hover:text-gray-600"
              title="ویرایش"
              @click="editFilter(filter)"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
              </svg>
            </button>
            <button 
              class="text-gray-400 hover:text-red-600"
              title="حذف"
              @click="deleteFilter(filter)"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
              </svg>
            </button>
          </div>
        </div>

        <!-- نمایش فیلترها -->
        <div class="space-y-2 mb-4">
          <div v-if="filter.criteria.cardCode" class="flex items-center text-xs text-gray-600">
            <span class="font-medium ml-1">کد کارت:</span>
            <span>{{ filter.criteria.cardCode }}</span>
          </div>
          <div v-if="filter.criteria.recipientName" class="flex items-center text-xs text-gray-600">
            <span class="font-medium ml-1">گیرنده:</span>
            <span>{{ filter.criteria.recipientName }}</span>
          </div>
          <div v-if="filter.criteria.statuses && filter.criteria.statuses.length > 0" class="flex items-center text-xs text-gray-600">
            <span class="font-medium ml-1">وضعیت:</span>
            <span>{{ filter.criteria.statuses.join(', ') }}</span>
          </div>
          <div v-if="filter.criteria.amountMin || filter.criteria.amountMax" class="flex items-center text-xs text-gray-600">
            <span class="font-medium ml-1">مبلغ:</span>
            <span>
              {{ filter.criteria.amountMin ? formatCurrency(filter.criteria.amountMin as number) : '0' }} - 
              {{ filter.criteria.amountMax ? formatCurrency(filter.criteria.amountMax as number) : 'نامحدود' }}
            </span>
          </div>
          <div v-if="filter.criteria.createdDateFrom || filter.criteria.createdDateTo" class="flex items-center text-xs text-gray-600">
            <span class="font-medium ml-1">تاریخ ایجاد:</span>
            <span>
              {{ filter.criteria.createdDateFrom || 'نامحدود' }} تا 
              {{ filter.criteria.createdDateTo || 'نامحدود' }}
            </span>
          </div>
        </div>

        <!-- آمار -->
        <div class="flex items-center justify-between text-xs text-gray-500 mb-4">
          <span>{{ filter.resultCount }} نتیجه</span>
          <span>{{ formatDate(filter.lastUsed) }}</span>
        </div>

        <!-- عملیات -->
        <div class="flex items-center space-x-2 space-x-reverse">
          <button 
            class="flex-1 px-3 py-2 bg-blue-600 text-white text-xs font-medium rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
            @click="applyFilter(filter)"
          >
            اعمال فیلتر
          </button>
          <button 
            class="px-3 py-2 bg-gray-100 text-gray-700 text-xs font-medium rounded-md hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
            title="اشتراک‌گذاری"
            @click="shareFilter(filter)"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.367 2.684 3 3 0 00-5.367-2.684z"></path>
            </svg>
          </button>
        </div>
      </div>
    </div>

    <!-- پیام خالی -->
    <div v-if="savedFilters.length === 0" class="text-center py-12">
      <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.207A1 1 0 013 6.5V4z"></path>
      </svg>
      <h3 class="mt-2 text-sm font-medium text-gray-900">هیچ فیلتر ذخیره‌ای وجود ندارد</h3>
      <p class="mt-1 text-sm text-gray-500">برای شروع، یک فیلتر جدید ذخیره کنید.</p>
      <button 
        class="mt-4 px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
        @click="showSaveFilterModal = true"
      >
        ذخیره فیلتر جدید
      </button>
    </div>

    <!-- مودال ذخیره فیلتر -->
    <div v-if="showSaveFilterModal" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
      <div class="relative top-20 mx-auto p-5 border w-full max-w-md sm:max-w-lg md:max-w-xl shadow-lg rounded-md bg-white">
        <div class="mt-3">
          <h3 class="text-lg font-medium text-gray-900 mb-4">ذخیره فیلتر جدید</h3>
          
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">نام فیلتر</label>
              <input
                v-model="newFilter.name"
                type="text"
                placeholder="مثال: کارت‌های فعال با مبلغ بالا"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">توضیحات</label>
              <textarea
                v-model="newFilter.description"
                rows="3"
                placeholder="توضیحات مختصر درباره این فیلتر"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              ></textarea>
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">نوع فیلتر</label>
              <select
                v-model="newFilter.type"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="personal">شخصی</option>
                <option value="shared">اشتراکی</option>
                <option value="system">سیستمی</option>
              </select>
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">دسترسی</label>
              <div class="space-y-2">
                <div class="flex items-center">
                  <input
                    v-model="newFilter.isPublic"
                    type="checkbox"
                    class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                  />
                  <label class="mr-2 text-sm text-gray-700">عمومی (قابل مشاهده برای همه)</label>
                </div>
                <div class="flex items-center">
                  <input
                    v-model="newFilter.isDefault"
                    type="checkbox"
                    class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                  />
                  <label class="mr-2 text-sm text-gray-700">پیش‌فرض (اعمال خودکار)</label>
                </div>
              </div>
            </div>
          </div>
          
          <div class="flex items-center justify-end space-x-3 space-x-reverse mt-6">
            <button 
              class="px-4 py-2 bg-gray-600 text-white text-sm font-medium rounded-md hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
              @click="showSaveFilterModal = false"
            >
              انصراف
            </button>
            <button 
              class="px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
              @click="saveFilter"
            >
              ذخیره
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- مودال اشتراک‌گذاری -->
    <div v-if="showShareModal" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
      <div class="relative top-20 mx-auto p-5 border w-full max-w-md sm:max-w-lg md:max-w-xl shadow-lg rounded-md bg-white">
        <div class="mt-3">
          <h3 class="text-lg font-medium text-gray-900 mb-4">اشتراک‌گذاری فیلتر</h3>
          
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">لینک اشتراک‌گذاری</label>
              <div class="flex">
                <input
                  :value="shareLink"
                  type="text"
                  readonly
                  class="flex-1 px-3 py-2 border border-gray-300 rounded-r-md shadow-sm bg-gray-50"
                />
                <button 
                  class="px-3 py-2 bg-blue-600 text-white text-sm font-medium rounded-l-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
                  @click="copyShareLink"
                >
                  کپی
                </button>
              </div>
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">اشتراک‌گذاری با کاربران</label>
              <select
                v-model="selectedUsers"
                multiple
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              >
                <option v-for="user in availableUsers" :key="user.id" :value="user.id">
                  {{ user.name }} ({{ user.role }})
                </option>
              </select>
            </div>
          </div>
          
          <div class="flex items-center justify-end space-x-3 space-x-reverse mt-6">
            <button 
              class="px-4 py-2 bg-gray-600 text-white text-sm font-medium rounded-md hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
              @click="showShareModal = false"
            >
              بستن
            </button>
            <button 
              class="px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
              @click="shareWithUsers"
            >
              اشتراک‌گذاری
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, reactive, ref } from 'vue';

// Props
interface CurrentFilters {
  [key: string]: unknown
}

const props = defineProps<{
  currentFilters: CurrentFilters
}>()

// Emits
interface FilterCriteria {
  [key: string]: unknown
}

const emit = defineEmits<{
  'apply-filter': [filters: FilterCriteria]
}>()

// Reactive data
const showSaveFilterModal = ref(false)
const showShareModal = ref(false)
const selectedFilter = ref(null)
const selectedUsers = ref([])

// فیلتر جدید
const newFilter = reactive({
  name: '',
  description: '',
  type: 'personal',
  isPublic: false,
  isDefault: false,
  criteria: {}
})

interface FilterCriteria {
  statuses?: string[]
  amountMin?: number
  amountMax?: number
  cardCode?: string
  recipientName?: string
  expiryDateFrom?: string
  expiryDateTo?: string
  createdDateFrom?: string
  createdDateTo?: string
  isUsed?: string
}

interface SavedFilter {
  id: number
  name: string
  description: string
  type: string
  isPublic: boolean
  isDefault: boolean
  criteria: FilterCriteria
  resultCount: number
  lastUsed: Date
  createdBy: string
  createdAt: Date
}

// فیلترهای ذخیره شده نمونه
const savedFilters = ref<SavedFilter[]>([
  {
    id: 1,
    name: 'کارت‌های فعال با مبلغ بالا',
    description: 'کارت‌های فعال با مبلغ بیشتر از 500,000 تومان',
    type: 'personal',
    isPublic: false,
    isDefault: false,
    criteria: {
      statuses: ['active'],
      amountMin: 500000,
      cardCode: '',
      recipientName: ''
    },
    resultCount: 15,
    lastUsed: new Date(Date.now() - 86400000),
    createdBy: 'مدیر سیستم',
    createdAt: new Date(Date.now() - 172800000)
  },
  {
    id: 2,
    name: 'کارت‌های منقضی شده',
    description: 'کارت‌هایی که در 30 روز آینده منقضی می‌شوند',
    type: 'shared',
    isPublic: true,
    isDefault: false,
    criteria: {
      statuses: ['active'],
      expiryDateFrom: new Date().toISOString().split('T')[0],
      expiryDateTo: new Date(Date.now() + 2592000000).toISOString().split('T')[0]
    },
    resultCount: 8,
    lastUsed: new Date(Date.now() - 3600000),
    createdBy: 'مدیر مالی',
    createdAt: new Date(Date.now() - 259200000)
  },
  {
    id: 3,
    name: 'کارت‌های استفاده نشده',
    description: 'کارت‌هایی که هنوز استفاده نشده‌اند',
    type: 'system',
    isPublic: true,
    isDefault: true,
    criteria: {
      isUsed: 'false',
      statuses: ['active']
    },
    resultCount: 42,
    lastUsed: new Date(Date.now() - 1800000),
    createdBy: 'سیستم',
    createdAt: new Date(Date.now() - 518400000)
  }
])

// کاربران موجود
const availableUsers = ref([
  { id: 1, name: 'علی احمدی', role: 'مدیر سیستم' },
  { id: 2, name: 'فاطمه محمدی', role: 'مدیر مالی' },
  { id: 3, name: 'حسن رضایی', role: 'پشتیبانی' },
  { id: 4, name: 'مریم کریمی', role: 'فروش' }
])

// Computed properties
const shareLink = computed(() => {
  if (!selectedFilter.value) return ''
  return `${window.location.origin}/admin/finance/giftcard?filter=${selectedFilter.value.id}`
})

// Methods
const applyFilter = (filter: SavedFilter) => {
  emit('apply-filter', filter.criteria)

}

const editFilter = (filter: SavedFilter) => {
  selectedFilter.value = filter
  Object.assign(newFilter, {
    name: filter.name,
    description: filter.description,
    type: filter.type,
    isPublic: filter.isPublic,
    isDefault: filter.isDefault,
    criteria: filter.criteria
  })
  showSaveFilterModal.value = true
}

const deleteFilter = (filter: SavedFilter) => {
  if (confirm(`آیا مطمئن هستید که می‌خواهید فیلتر "${filter.name}" را حذف کنید؟`)) {
    const index = savedFilters.value.findIndex(f => f.id === filter.id)
    if (index > -1) {
      savedFilters.value.splice(index, 1)
    }

  }
}

const shareFilter = (filter: SavedFilter) => {
  selectedFilter.value = filter
  showShareModal.value = true
}

const saveFilter = () => {
  if (!newFilter.name.trim()) {
    alert('لطفاً نام فیلتر را وارد کنید')
    return
  }

  const filter = {
    id: Date.now(),
    ...newFilter,
    criteria: { ...props.currentFilters },
    resultCount: 0,
    lastUsed: new Date(),
    createdBy: 'کاربر فعلی',
    createdAt: new Date()
  }

  savedFilters.value.unshift(filter)
  showSaveFilterModal.value = false
  
  // پاک کردن فرم
  Object.assign(newFilter, {
    name: '',
    description: '',
    type: 'personal',
    isPublic: false,
    isDefault: false,
    criteria: {}
  })

}

const copyShareLink = () => {
  navigator.clipboard.writeText(shareLink.value).then(() => {
    alert('لینک کپی شد')
  })
}

const shareWithUsers = () => {
  if (selectedUsers.value.length === 0) {
    alert('لطفاً حداقل یک کاربر انتخاب کنید')
    return
  }

  showShareModal.value = false
  selectedUsers.value = []
}

const getFilterIcon = (type: string) => {
  const icons = {
    personal: 'UserIcon',
    shared: 'UsersIcon',
    system: 'CogIcon'
  }
  return icons[type] || 'FilterIcon'
}

const getFilterIconClass = (type: string) => {
  const classes = {
    personal: 'bg-blue-100 text-blue-600',
    shared: 'bg-green-100 text-green-600',
    system: 'bg-purple-100 text-purple-600'
  }
  return classes[type] || 'bg-gray-100 text-gray-600'
}

const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

const formatDate = (date: Date) => {
  return new Intl.DateTimeFormat('fa-IR').format(date)
}

// Lifecycle
onMounted(() => {

})
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 

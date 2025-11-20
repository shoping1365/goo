<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">مدیریت تم‌های گیفت کارت</h3>
        <p class="text-sm text-gray-600">تنظیم و مدیریت تم‌های مختلف گیفت کارت</p>
      </div>
      <button
        class="inline-flex items-center px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
        @click="showCreateModal = true"
      >
        <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path>
        </svg>
        افزودن تم جدید
      </button>
    </div>

    <!-- آمار کلی تم‌ها -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div class="bg-white p-6 rounded-lg border border-gray-200">
        <div class="flex items-center">
          <div class="p-2 bg-blue-100 rounded-lg">
            <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21a4 4 0 01-4-4V5a2 2 0 012-2h4a2 2 0 012 2v12a4 4 0 01-4 4zM21 5a2 2 0 00-2-2h-4a2 2 0 00-2 2v12a4 4 0 004 4h4a2 2 0 002-2V5z"></path>
            </svg>
          </div>
          <div class="ml-3">
            <p class="text-sm font-medium text-gray-600">کل تم‌ها</p>
            <p class="text-2xl font-bold text-gray-900">{{ themes.length }}</p>
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
            <p class="text-2xl font-bold text-gray-900">{{ activeThemes.length }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white p-6 rounded-lg border border-gray-200">
        <div class="flex items-center">
          <div class="p-2 bg-yellow-100 rounded-lg">
            <svg class="w-6 h-6 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
          <div class="ml-3">
            <p class="text-sm font-medium text-gray-600">در انتظار</p>
            <p class="text-2xl font-bold text-gray-900">{{ pendingThemes.length }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white p-6 rounded-lg border border-gray-200">
        <div class="flex items-center">
          <div class="p-2 bg-purple-100 rounded-lg">
            <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"></path>
            </svg>
          </div>
          <div class="ml-3">
            <p class="text-sm font-medium text-gray-600">محبوب</p>
            <p class="text-2xl font-bold text-gray-900">{{ popularThemes.length }}</p>
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
            placeholder="جستجو در تم‌ها..."
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
            <option value="pending">در انتظار</option>
          </select>
          <select
            v-model="categoryFilter"
            class="px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          >
            <option value="">همه دسته‌ها</option>
            <option value="birthday">تولد</option>
            <option value="wedding">عروسی</option>
            <option value="newyear">سال نو</option>
            <option value="corporate">شرکتی</option>
            <option value="general">عمومی</option>
          </select>
        </div>
      </div>
    </div>

    <!-- لیست تم‌ها -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div
        v-for="theme in filteredThemes"
        :key="theme.id"
        class="bg-white rounded-lg border border-gray-200 overflow-hidden hover:shadow-lg transition-shadow"
      >
        <!-- تصویر تم -->
        <div class="relative h-48 bg-gradient-to-br from-blue-400 to-purple-600">
          <div class="absolute inset-0 flex items-center justify-center">
            <div class="text-white text-center">
              <div class="text-4xl mb-2">{{ theme.icon }}</div>
              <h3 class="text-lg font-semibold">{{ theme.name }}</h3>
            </div>
          </div>
          <div class="absolute top-2 right-2">
            <span
              :class="{
                'bg-green-100 text-green-800': theme.status === 'active',
                'bg-red-100 text-red-800': theme.status === 'inactive',
                'bg-yellow-100 text-yellow-800': theme.status === 'pending'
              }"
              class="px-2 py-1 text-xs font-medium rounded-full"
            >
              {{ getStatusText(theme.status) }}
            </span>
          </div>
        </div>

        <!-- اطلاعات تم -->
        <div class="p-6">
          <div class="flex items-center justify-between mb-3">
            <h4 class="text-lg font-semibold text-gray-900">{{ theme.name }}</h4>
            <span class="text-sm text-gray-500">{{ theme.category }}</span>
          </div>
          
          <p class="text-sm text-gray-600 mb-4">{{ theme.description }}</p>

          <!-- آمار تم -->
          <div class="grid grid-cols-2 gap-6 mb-4">
            <div class="text-center">
              <p class="text-2xl font-bold text-blue-600">{{ theme.usageCount }}</p>
              <p class="text-xs text-gray-500">استفاده شده</p>
            </div>
            <div class="text-center">
              <p class="text-2xl font-bold text-green-600">{{ theme.rating }}</p>
              <p class="text-xs text-gray-500">امتیاز</p>
            </div>
          </div>

          <!-- تنظیمات تم -->
          <div class="space-y-2 mb-4">
            <div class="flex items-center justify-between text-sm">
              <span class="text-gray-600">رنگ اصلی:</span>
              <div class="flex items-center">
                <div
                  class="w-4 h-4 rounded-full mr-2"
                  :style="{ backgroundColor: theme.primaryColor }"
                ></div>
                <span class="text-gray-900">{{ theme.primaryColor }}</span>
              </div>
            </div>
            <div class="flex items-center justify-between text-sm">
              <span class="text-gray-600">فونت:</span>
              <span class="text-gray-900">{{ theme.fontFamily }}</span>
            </div>
            <div class="flex items-center justify-between text-sm">
              <span class="text-gray-600">سایز:</span>
              <span class="text-gray-900">{{ theme.fontSize }}</span>
            </div>
          </div>

          <!-- دکمه‌های عملیات -->
          <div class="flex gap-2">
            <button
              class="flex-1 px-3 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
              @click="editTheme(theme)"
            >
              ویرایش
            </button>
            <button
              class="flex-1 px-3 py-2 bg-gray-600 text-white text-sm font-medium rounded-lg hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
              @click="previewTheme(theme)"
            >
              پیش‌نمایش
            </button>
            <button
              class="px-3 py-2 bg-red-600 text-white text-sm font-medium rounded-lg hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2"
              @click="deleteTheme(theme)"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
              </svg>
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- مودال ایجاد تم جدید -->
    <GiftCardThemeCreateModal
      v-if="showCreateModal"
      @close="showCreateModal = false"
      @created="handleThemeCreated"
    />

    <!-- مودال ویرایش تم -->
    <GiftCardThemeEditModal
      v-if="showEditModal"
      :theme="selectedTheme"
      @close="showEditModal = false"
      @updated="handleThemeUpdated"
    />

    <!-- مودال پیش‌نمایش تم -->
    <GiftCardThemePreviewModal
      v-if="showPreviewModal"
      :theme="selectedTheme"
      @close="showPreviewModal = false"
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
const showPreviewModal = ref(false)
const selectedTheme = ref(null)

// داده‌های نمونه تم‌ها
const themes = ref([
  {
    id: 1,
    name: 'تم تولد شاد',
    description: 'تم زیبا و شاد برای کارت‌های تولد با رنگ‌های گرم و تصاویر جذاب',
    category: 'تولد',
    status: 'active',
    icon: '🎂',
    primaryColor: '#FF6B6B',
    fontFamily: 'IRANSans',
    fontSize: '16px',
    usageCount: 1250,
    rating: 4.8,
    backgroundColor: '#FFF5F5',
    textColor: '#2D3748',
    borderColor: '#FFE2E2',
    borderRadius: '12px',
    shadow: '0 4px 6px rgba(0, 0, 0, 0.1)',
    animation: 'fadeIn',
    customCSS: '',
    isDefault: true
  },
  {
    id: 2,
    name: 'تم عروسی مجلل',
    description: 'تم لوکس و مجلل برای کارت‌های عروسی با طراحی کلاسیک',
    category: 'عروسی',
    status: 'active',
    icon: '💒',
    primaryColor: '#9F7AEA',
    fontFamily: 'IRANSans',
    fontSize: '18px',
    usageCount: 890,
    rating: 4.9,
    backgroundColor: '#FAF5FF',
    textColor: '#2D3748',
    borderColor: '#E9D8FD',
    borderRadius: '16px',
    shadow: '0 10px 15px rgba(0, 0, 0, 0.1)',
    animation: 'slideIn',
    customCSS: '',
    isDefault: false
  },
  {
    id: 3,
    name: 'تم سال نو',
    description: 'تم شاد و رنگارنگ برای تبریک سال نو با المان‌های جشن',
    category: 'سال نو',
    status: 'active',
    icon: '🎊',
    primaryColor: '#38B2AC',
    fontFamily: 'IRANSans',
    fontSize: '16px',
    usageCount: 2100,
    rating: 4.7,
    backgroundColor: '#F0FFF4',
    textColor: '#2D3748',
    borderColor: '#C6F6D5',
    borderRadius: '20px',
    shadow: '0 8px 25px rgba(0, 0, 0, 0.15)',
    animation: 'bounce',
    customCSS: '',
    isDefault: true
  },
  {
    id: 4,
    name: 'تم شرکتی حرفه‌ای',
    description: 'تم رسمی و حرفه‌ای برای کارت‌های شرکتی و تجاری',
    category: 'شرکتی',
    status: 'active',
    icon: '🏢',
    primaryColor: '#2D3748',
    fontFamily: 'IRANSans',
    fontSize: '14px',
    usageCount: 650,
    rating: 4.6,
    backgroundColor: '#F7FAFC',
    textColor: '#2D3748',
    borderColor: '#E2E8F0',
    borderRadius: '8px',
    shadow: '0 2px 4px rgba(0, 0, 0, 0.1)',
    animation: 'none',
    customCSS: '',
    isDefault: false
  },
  {
    id: 5,
    name: 'تم تخفیف ویژه',
    description: 'تم جذاب برای کارت‌های تخفیف با طراحی چشم‌نواز',
    category: 'تخفیف',
    status: 'pending',
    icon: '🎯',
    primaryColor: '#F56565',
    fontFamily: 'IRANSans',
    fontSize: '16px',
    usageCount: 320,
    rating: 4.5,
    backgroundColor: '#FFF5F5',
    textColor: '#2D3748',
    borderColor: '#FEB2B2',
    borderRadius: '14px',
    shadow: '0 6px 20px rgba(0, 0, 0, 0.12)',
    animation: 'pulse',
    customCSS: '',
    isDefault: false
  },
  {
    id: 6,
    name: 'تم عمومی ساده',
    description: 'تم ساده و کاربردی برای استفاده عمومی',
    category: 'عمومی',
    status: 'inactive',
    icon: '🎁',
    primaryColor: '#4299E1',
    fontFamily: 'IRANSans',
    fontSize: '15px',
    usageCount: 180,
    rating: 4.3,
    backgroundColor: '#EBF8FF',
    textColor: '#2D3748',
    borderColor: '#BEE3F8',
    borderRadius: '10px',
    shadow: '0 4px 6px rgba(0, 0, 0, 0.07)',
    animation: 'none',
    customCSS: '',
    isDefault: false
  }
])

// محاسبه تم‌های فعال
const activeThemes = computed(() => themes.value.filter(theme => theme.status === 'active'))

// محاسبه تم‌های در انتظار
const pendingThemes = computed(() => themes.value.filter(theme => theme.status === 'pending'))

// محاسبه تم‌های محبوب
const popularThemes = computed(() => themes.value.filter(theme => theme.usageCount > 1000))

// فیلتر کردن تم‌ها بر اساس جستجو و فیلترها
const filteredThemes = computed(() => {
  return themes.value.filter(theme => {
    const matchesSearch = theme.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
                         theme.description.toLowerCase().includes(searchQuery.value.toLowerCase())
    
    const matchesStatus = !statusFilter.value || theme.status === statusFilter.value
    
    const matchesCategory = !categoryFilter.value || theme.category === categoryFilter.value
    
    return matchesSearch && matchesStatus && matchesCategory
  })
})

// تابع دریافت متن وضعیت
const getStatusText = (status) => {
  const statusMap = {
    active: 'فعال',
    inactive: 'غیرفعال',
    pending: 'در انتظار'
  }
  return statusMap[status] || status
}

// تابع ویرایش تم
const editTheme = (theme) => {
  selectedTheme.value = theme
  showEditModal.value = true
}

// تابع پیش‌نمایش تم
const previewTheme = (theme) => {
  selectedTheme.value = theme
  showPreviewModal.value = true
}

// تابع حذف تم
const deleteTheme = (theme) => {
  if (confirm(`آیا از حذف تم "${theme.name}" اطمینان دارید؟`)) {
    const index = themes.value.findIndex(t => t.id === theme.id)
    if (index > -1) {
      themes.value.splice(index, 1)
      // نمایش پیام موفقیت
      alert('تم با موفقیت حذف شد')
    }
  }
}

// تابع مدیریت ایجاد تم جدید
const handleThemeCreated = (newTheme) => {
  themes.value.push({
    ...newTheme,
    id: Date.now(),
    usageCount: 0,
    rating: 0
  })
  showCreateModal.value = false
}

// تابع مدیریت بروزرسانی تم
const handleThemeUpdated = (updatedTheme) => {
  const index = themes.value.findIndex(t => t.id === updatedTheme.id)
  if (index > -1) {
    themes.value[index] = { ...themes.value[index], ...updatedTheme }
  }
  showEditModal.value = false
}
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 

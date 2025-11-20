<template>
  <div class="space-y-6">
    <!-- هدر صفحه -->
    <div class="flex justify-between items-center">
      <div>
        <h1 class="text-2xl font-bold text-gray-900">گرونه شانس</h1>
        <p class="text-gray-600">مدیریت قرعه‌کشی‌ها و جوایز</p>
      </div>
      <button 
        class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors"
        @click="showLuckyDrawForm = true"
      >
        ایجاد گرونه جدید
      </button>
    </div>

    <!-- آمار کلی -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div class="bg-white p-6 rounded-lg shadow border">
        <div class="flex items-center">
          <div class="p-2 bg-blue-100 rounded-lg">
            <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
          </div>
          <div class="mr-3">
            <p class="text-sm text-gray-600">کل گرونه‌ها</p>
            <p class="text-2xl font-bold text-gray-900">{{ luckyDrawStats.totalDraws }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white p-6 rounded-lg shadow border">
        <div class="flex items-center">
          <div class="p-2 bg-green-100 rounded-lg">
            <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
          <div class="mr-3">
            <p class="text-sm text-gray-600">گرونه‌های فعال</p>
            <p class="text-2xl font-bold text-gray-900">{{ luckyDrawStats.activeDraws }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white p-6 rounded-lg shadow border">
        <div class="flex items-center">
          <div class="p-2 bg-purple-100 rounded-lg">
            <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"></path>
            </svg>
          </div>
          <div class="mr-3">
            <p class="text-sm text-gray-600">شرکت‌کنندگان</p>
            <p class="text-2xl font-bold text-gray-900">{{ luckyDrawStats.totalParticipants }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white p-6 rounded-lg shadow border">
        <div class="flex items-center">
          <div class="p-2 bg-yellow-100 rounded-lg">
            <svg class="w-6 h-6 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v13m0-13V6a2 2 0 112 2h-2zm0 0V5.5A2.5 2.5 0 109.5 8H12z"></path>
            </svg>
          </div>
          <div class="mr-3">
            <p class="text-sm text-gray-600">جوایز اهدا شده</p>
            <p class="text-2xl font-bold text-gray-900">{{ luckyDrawStats.totalPrizes }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- کارت‌های ابزار -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <div class="bg-white p-6 rounded-lg shadow border hover:shadow-lg transition-shadow cursor-pointer" @click="showPrizeManagement = true">
        <div class="text-center">
          <div class="mx-auto w-12 h-12 bg-blue-100 rounded-full flex items-center justify-center mb-4">
            <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v13m0-13V6a2 2 0 112 2h-2zm0 0V5.5A2.5 2.5 0 109.5 8H12z"></path>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-900 mb-2">مدیریت جوایز</h3>
          <p class="text-gray-600 text-sm">تعریف و مدیریت جوایز گرونه</p>
        </div>
      </div>

      <div class="bg-white p-6 rounded-lg shadow border hover:shadow-lg transition-shadow cursor-pointer" @click="showLuckyDrawReports = true">
        <div class="text-center">
          <div class="mx-auto w-12 h-12 bg-green-100 rounded-full flex items-center justify-center mb-4">
            <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-900 mb-2">گزارش‌ها</h3>
          <p class="text-gray-600 text-sm">گزارش‌های تحلیلی و آماری</p>
        </div>
      </div>

      <div class="bg-white p-6 rounded-lg shadow border hover:shadow-lg transition-shadow cursor-pointer" @click="showLuckyDrawSettings = true">
        <div class="text-center">
          <div class="mx-auto w-12 h-12 bg-purple-100 rounded-full flex items-center justify-center mb-4">
            <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"></path>
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-900 mb-2">تنظیمات</h3>
          <p class="text-gray-600 text-sm">تنظیمات سیستم گرونه شانس</p>
        </div>
      </div>

      <div class="bg-white p-6 rounded-lg shadow border hover:shadow-lg transition-shadow cursor-pointer" @click="showLuckyDrawForm = true">
        <div class="text-center">
          <div class="mx-auto w-12 h-12 bg-orange-100 rounded-full flex items-center justify-center mb-4">
            <svg class="w-6 h-6 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-900 mb-2">ایجاد جدید</h3>
          <p class="text-gray-600 text-sm">ایجاد گرونه شانس جدید</p>
        </div>
      </div>
    </div>

    <!-- جدول گرونه‌های فعال -->
    <div class="bg-white rounded-lg shadow border">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-semibold text-gray-900">گرونه‌های فعال</h3>
      </div>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نام گرونه</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ شروع</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ پایان</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">شرکت‌کنندگان</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="draw in activeLuckyDraws" :key="draw.id">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm font-medium text-gray-900">{{ draw.name }}</div>
                <div class="text-sm text-gray-500">{{ draw.description }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full" :class="getTypeBadgeClass(draw.type)">
                  {{ getTypeText(draw.type) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ formatDate(draw.startDate) }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ formatDate(draw.endDate) }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ draw.participants }}</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-green-100 text-green-800">
                  فعال
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3" @click="editLuckyDraw(draw)">ویرایش</button>
                <button class="text-red-600 hover:text-red-900" @click="deleteLuckyDraw(draw.id)">حذف</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- مودال‌ها -->
    <LuckyDrawForm 
      v-if="showLuckyDrawForm" 
      :lucky-draw="editingLuckyDraw"
      @close="closeLuckyDrawForm"
      @saved="onLuckyDrawSaved"
    />

    <PrizeManagement 
      v-if="showPrizeManagement" 
      @close="showPrizeManagement = false"
    />

    <LuckyDrawReports 
      v-if="showLuckyDrawReports" 
      @close="showLuckyDrawReports = false"
    />

    <LuckyDrawSettings 
      v-if="showLuckyDrawSettings" 
      @close="showLuckyDrawSettings = false"
    />
  </div>
</template>

<script setup>

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user, hasPermission } = useAuth()

// متغیرهای reactive
const showLuckyDrawForm = ref(false)
const showPrizeManagement = ref(false)
const showLuckyDrawReports = ref(false)
const showLuckyDrawSettings = ref(false)
const editingLuckyDraw = ref(null)

// آمار گرونه شانس
const luckyDrawStats = ref({
  totalDraws: 12,
  activeDraws: 5,
  totalParticipants: 2847,
  totalPrizes: 156
})

// لیست گرونه‌های فعال
const activeLuckyDraws = ref([
  {
    id: 1,
    name: 'گرونه هفتگی',
    description: 'قرعه‌کشی هفتگی برای مشتریان وفادار',
    type: 'weekly',
    startDate: '2024-01-01',
    endDate: '2024-12-31',
    participants: 1250,
    status: 'active'
  },
  {
    id: 2,
    name: 'گرونه خرید',
    description: 'قرعه‌کشی برای هر خرید بالای 500 هزار تومان',
    type: 'purchase',
    startDate: '2024-01-01',
    endDate: '2024-12-31',
    participants: 890,
    status: 'active'
  },
  {
    id: 3,
    name: 'گرونه ثبت‌نام',
    description: 'قرعه‌کشی برای کاربران جدید',
    type: 'registration',
    startDate: '2024-01-01',
    endDate: '2024-12-31',
    participants: 707,
    status: 'active'
  }
])

// توابع کمکی
const getTypeBadgeClass = (type) => {
  const classes = {
    daily: 'bg-blue-100 text-blue-800',
    weekly: 'bg-green-100 text-green-800',
    monthly: 'bg-purple-100 text-purple-800',
    purchase: 'bg-orange-100 text-orange-800',
    registration: 'bg-pink-100 text-pink-800'
  }
  return classes[type] || 'bg-gray-100 text-gray-800'
}

const getTypeText = (type) => {
  const texts = {
    daily: 'روزانه',
    weekly: 'هفتگی',
    monthly: 'ماهانه',
    purchase: 'خرید',
    registration: 'ثبت‌نام'
  }
  return texts[type] || type
}

const formatDate = (dateString) => {
  return new Date(dateString).toLocaleDateString('fa-IR')
}

// توابع مدیریت
const editLuckyDraw = (draw) => {
  editingLuckyDraw.value = { ...draw }
  showLuckyDrawForm.value = true
}

const deleteLuckyDraw = (id) => {
  if (confirm('آیا از حذف این گرونه اطمینان دارید؟')) {
    activeLuckyDraws.value = activeLuckyDraws.value.filter(draw => draw.id !== id)
    luckyDrawStats.value.totalDraws--
    luckyDrawStats.value.activeDraws--
  }
}

const closeLuckyDrawForm = () => {
  showLuckyDrawForm.value = false
  editingLuckyDraw.value = null
}

const onLuckyDrawSaved = (savedDraw) => {
  if (editingLuckyDraw.value) {
    // ویرایش
    const index = activeLuckyDraws.value.findIndex(draw => draw.id === savedDraw.id)
    if (index !== -1) {
      activeLuckyDraws.value[index] = savedDraw
    }
  } else {
    // ایجاد جدید
    savedDraw.id = Date.now()
    activeLuckyDraws.value.unshift(savedDraw)
    luckyDrawStats.value.totalDraws++
    luckyDrawStats.value.activeDraws++
  }
  closeLuckyDrawForm()
}
</script>


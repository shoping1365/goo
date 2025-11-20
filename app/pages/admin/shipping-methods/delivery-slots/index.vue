<template>
  <div class="space-y-6">
    <!-- هدر صفحه -->
    <div class="flex justify-between items-center">
      <h1 class="text-2xl font-bold text-gray-900">زمان‌بندی تحویل</h1>
      <div class="flex space-x-3 space-x-reverse">
        <button class="bg-green-600 text-white px-4 py-2 rounded-lg hover:bg-green-700 transition-colors">
          افزودن بازه زمانی
        </button>
        <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors">
          تنظیمات پیش‌فرض
        </button>
      </div>
    </div>

    <!-- آمار کلی -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div class="bg-white p-6 rounded-lg shadow">
        <div class="flex items-center">
          <div class="p-2 bg-blue-100 rounded-lg">
            <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">کل بازه‌های زمانی</p>
            <p class="text-2xl font-semibold text-gray-900">{{ slotStats.total }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white p-6 rounded-lg shadow">
        <div class="flex items-center">
          <div class="p-2 bg-green-100 rounded-lg">
            <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">فعال</p>
            <p class="text-2xl font-semibold text-gray-900">{{ slotStats.active }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white p-6 rounded-lg shadow">
        <div class="flex items-center">
          <div class="p-2 bg-yellow-100 rounded-lg">
            <svg class="w-6 h-6 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">رزرو شده امروز</p>
            <p class="text-2xl font-semibold text-gray-900">{{ slotStats.reserved }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white p-6 rounded-lg shadow">
        <div class="flex items-center">
          <div class="p-2 bg-purple-100 rounded-lg">
            <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"></path>
            </svg>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">متوسط ظرفیت</p>
            <p class="text-2xl font-semibold text-gray-900">{{ slotStats.avgCapacity }}%</p>
          </div>
        </div>
      </div>
    </div>

    <!-- تقویم هفتگی -->
    <div class="bg-white rounded-lg shadow overflow-hidden">
      <div class="px-6 py-4 border-b border-gray-200">
        <div class="flex justify-between items-center">
          <h3 class="text-lg font-semibold text-gray-900">تقویم هفتگی</h3>
          <div class="flex space-x-2 space-x-reverse">
            <button class="px-3 py-1 text-sm bg-blue-100 text-blue-800 rounded-md">هفته قبل</button>
            <button class="px-3 py-1 text-sm bg-blue-600 text-white rounded-md">هفته جاری</button>
            <button class="px-3 py-1 text-sm bg-blue-100 text-blue-800 rounded-md">هفته بعد</button>
          </div>
        </div>
      </div>
      <div class="overflow-x-auto">
        <table class="min-w-full">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">ساعت</th>
              <th v-for="day in weekDays" :key="day.name" class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                {{ day.name }}
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="timeSlot in timeSlots" :key="timeSlot.time">
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                {{ timeSlot.time }}
              </td>
              <td v-for="day in weekDays" :key="day.name" class="px-6 py-4 whitespace-nowrap">
                <div class="text-center">
                  <div
v-if="getSlotStatus(day.name, timeSlot.time)" 
                       class="inline-flex items-center px-2 py-1 text-xs font-semibold rounded-full"
                       :class="getSlotStatusClass(getSlotStatus(day.name, timeSlot.time))">
                    {{ getSlotStatusText(getSlotStatus(day.name, timeSlot.time)) }}
                  </div>
                  <div v-else class="text-gray-400 text-xs">خالی</div>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- فیلترها -->
    <div class="bg-white p-6 rounded-lg shadow">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">منطقه</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option value="">همه مناطق</option>
            <option value="tehran">تهران</option>
            <option value="isfahan">اصفهان</option>
            <option value="shiraz">شیراز</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">روش ارسال</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option value="">همه روش‌ها</option>
            <option value="express">ارسال سریع</option>
            <option value="standard">ارسال عادی</option>
            <option value="scheduled">ارسال زمان‌دار</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">وضعیت</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option value="">همه وضعیت‌ها</option>
            <option value="available">موجود</option>
            <option value="reserved">رزرو شده</option>
            <option value="full">پر</option>
            <option value="disabled">غیرفعال</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">تاریخ</label>
          <input 
            type="date" 
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        </div>
      </div>
    </div>

    <!-- جدول بازه‌های زمانی -->
    <div class="bg-white rounded-lg shadow overflow-hidden">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-semibold text-gray-900">بازه‌های زمانی</h3>
      </div>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                بازه زمانی
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                روزهای هفته
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                ظرفیت
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                رزرو شده
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                وضعیت
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                عملیات
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="slot in deliverySlots" :key="slot.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm font-medium text-gray-900">{{ slot.timeRange }}</div>
                <div class="text-sm text-gray-500">{{ slot.duration }} دقیقه</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex flex-wrap gap-1">
                  <span
v-for="day in slot.days" :key="day" 
                        class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-blue-100 text-blue-800">
                    {{ day }}
                  </span>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ slot.capacity }} سفارش
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ slot.reserved }} سفارش
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span
class="inline-flex px-2 py-1 text-xs font-semibold rounded-full" 
                      :class="getSlotStatusClass(slot.status)">
                  {{ getSlotStatusText(slot.status) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">ویرایش</button>
                <button class="text-green-600 hover:text-green-900 ml-3">کپی</button>
                <button class="text-red-600 hover:text-red-900">حذف</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup>

// استفاده از useAuth برای چک کردن پرمیژن‌ها
// const { user, hasPermission } = useAuth()

// آمار بازه‌های زمانی
const slotStats = ref({
  total: 24,
  active: 20,
  reserved: 156,
  avgCapacity: 75
})

// روزهای هفته
const weekDays = ref([
  { name: 'شنبه', key: 'saturday' },
  { name: 'یکشنبه', key: 'sunday' },
  { name: 'دوشنبه', key: 'monday' },
  { name: 'سه‌شنبه', key: 'tuesday' },
  { name: 'چهارشنبه', key: 'wednesday' },
  { name: 'پنج‌شنبه', key: 'thursday' },
  { name: 'جمعه', key: 'friday' }
])

// بازه‌های زمانی
const timeSlots = ref([
  { time: '08:00 - 10:00' },
  { time: '10:00 - 12:00' },
  { time: '12:00 - 14:00' },
  { time: '14:00 - 16:00' },
  { time: '16:00 - 18:00' },
  { time: '18:00 - 20:00' }
])

// داده‌های نمونه بازه‌های زمانی
const deliverySlots = ref([
  {
    id: 1,
    timeRange: '08:00 - 10:00',
    duration: 120,
    days: ['شنبه', 'یکشنبه', 'دوشنبه', 'سه‌شنبه', 'چهارشنبه'],
    capacity: 50,
    reserved: 35,
    status: 'available'
  },
  {
    id: 2,
    timeRange: '10:00 - 12:00',
    duration: 120,
    days: ['شنبه', 'یکشنبه', 'دوشنبه', 'سه‌شنبه', 'چهارشنبه'],
    capacity: 50,
    reserved: 50,
    status: 'full'
  },
  {
    id: 3,
    timeRange: '14:00 - 16:00',
    duration: 120,
    days: ['شنبه', 'یکشنبه', 'دوشنبه', 'سه‌شنبه', 'چهارشنبه'],
    capacity: 50,
    reserved: 25,
    status: 'available'
  },
  {
    id: 4,
    timeRange: '16:00 - 18:00',
    duration: 120,
    days: ['شنبه', 'یکشنبه', 'دوشنبه', 'سه‌شنبه', 'چهارشنبه'],
    capacity: 50,
    reserved: 46,
    status: 'reserved'
  }
])

// تابع دریافت وضعیت بازه زمانی
const getSlotStatus = (day, time) => {
  // این تابع وضعیت بازه زمانی را برای روز و زمان مشخص برمی‌گرداند
  const statuses = {
    'شنبه': {
      '08:00 - 10:00': 'available',
      '10:00 - 12:00': 'full',
      '14:00 - 16:00': 'available',
      '16:00 - 18:00': 'reserved'
    }
  }
  return statuses[day]?.[time] || null
}

// تابع کلاس‌های وضعیت بازه زمانی
const getSlotStatusClass = (status) => {
  const classes = {
    available: 'bg-green-100 text-green-800',
    reserved: 'bg-yellow-100 text-yellow-800',
    full: 'bg-red-100 text-red-800',
    disabled: 'bg-gray-100 text-gray-800'
  }
  return classes[status] || classes.disabled
}

// تابع متن وضعیت بازه زمانی
const getSlotStatusText = (status) => {
  const texts = {
    available: 'موجود',
    reserved: 'رزرو شده',
    full: 'پر',
    disabled: 'غیرفعال'
  }
  return texts[status] || 'نامشخص'
}

// تنظیمات صفحه
definePageMeta({
  title: 'زمان‌بندی تحویل',
  layout: 'admin-main',
  middleware: 'admin'
})
</script> 


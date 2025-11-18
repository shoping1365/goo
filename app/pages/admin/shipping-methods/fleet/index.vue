<template>
  <div class="space-y-6">
    <!-- هدر صفحه -->
    <div class="flex justify-between items-center">
      <h1 class="text-2xl font-bold text-gray-900">مدیریت ناوگان</h1>
      <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors">
        افزودن وسیله نقلیه جدید
      </button>
    </div>

    <!-- آمار کلی -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div class="bg-white p-6 rounded-lg shadow">
        <div class="flex items-center">
          <div class="p-2 bg-blue-100 rounded-lg">
            <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4"></path>
            </svg>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">کل وسایل نقلیه</p>
            <p class="text-2xl font-semibold text-gray-900">{{ fleetStats.total }}</p>
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
            <p class="text-2xl font-semibold text-gray-900">{{ fleetStats.active }}</p>
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
            <p class="text-sm font-medium text-gray-600">در حال سرویس</p>
            <p class="text-2xl font-semibold text-gray-900">{{ fleetStats.maintenance }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white p-6 rounded-lg shadow">
        <div class="flex items-center">
          <div class="p-2 bg-red-100 rounded-lg">
            <svg class="w-6 h-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">غیرفعال</p>
            <p class="text-2xl font-semibold text-gray-900">{{ fleetStats.inactive }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- فیلترها -->
    <div class="bg-white p-6 rounded-lg shadow">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">جستجو</label>
          <input 
            type="text" 
            placeholder="جستجو بر اساس پلاک یا مدل..."
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">نوع وسیله نقلیه</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option value="">همه انواع</option>
            <option value="car">خودرو</option>
            <option value="van">ون</option>
            <option value="truck">کامیون</option>
            <option value="motorcycle">موتورسیکلت</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">وضعیت</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option value="">همه وضعیت‌ها</option>
            <option value="active">فعال</option>
            <option value="maintenance">در حال سرویس</option>
            <option value="inactive">غیرفعال</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">راننده</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option value="">همه رانندگان</option>
            <option value="assigned">دارای راننده</option>
            <option value="unassigned">بدون راننده</option>
          </select>
        </div>
      </div>
    </div>

    <!-- جدول ناوگان -->
    <div class="bg-white rounded-lg shadow overflow-hidden">
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                وسیله نقلیه
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                راننده
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                وضعیت
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                آخرین سرویس
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                مسافت طی شده
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                عملیات
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="vehicle in fleetVehicles" :key="vehicle.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="flex-shrink-0 h-12 w-12">
                    <img class="h-12 w-12 rounded-lg object-cover" :src="vehicle.image" :alt="vehicle.model" />
                  </div>
                  <div class="mr-4">
                    <div class="text-sm font-medium text-gray-900">{{ vehicle.model }}</div>
                    <div class="text-sm text-gray-500">{{ vehicle.plate }}</div>
                    <div class="text-xs text-gray-400">{{ vehicle.year }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div v-if="vehicle.driver" class="flex items-center">
                  <div class="flex-shrink-0 h-8 w-8">
                    <img class="h-8 w-8 rounded-full" :src="vehicle.driver.avatar" :alt="vehicle.driver.name" />
                  </div>
                  <div class="mr-3">
                    <div class="text-sm font-medium text-gray-900">{{ vehicle.driver.name }}</div>
                    <div class="text-sm text-gray-500">{{ vehicle.driver.phone }}</div>
                  </div>
                </div>
                <span v-else class="text-sm text-gray-500">بدون راننده</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full" 
                      :class="getStatusClass(vehicle.status)">
                  {{ getStatusText(vehicle.status) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ formatDate(vehicle.lastService) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ vehicle.mileage }} کیلومتر
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">ویرایش</button>
                <button class="text-green-600 hover:text-green-900 ml-3">سرویس</button>
                <button class="text-red-600 hover:text-red-900">حذف</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- پیجینیشن -->
    <div class="bg-white px-4 py-3 flex items-center justify-between border-t border-gray-200 sm:px-6">
      <div class="flex-1 flex justify-between sm:hidden">
        <button class="relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50">
          قبلی
        </button>
        <button class="mr-3 relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50">
          بعدی
        </button>
      </div>
      <div class="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
        <div>
          <p class="text-sm text-gray-700">
            نمایش
            <span class="font-medium">1</span>
            تا
            <span class="font-medium">10</span>
            از
            <span class="font-medium">25</span>
            نتیجه
          </p>
        </div>
        <div>
          <nav class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px">
            <button class="relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50">
              <span class="sr-only">قبلی</span>
              <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd" />
              </svg>
            </button>
            <button class="bg-blue-50 border-blue-500 text-blue-600 relative inline-flex items-center px-4 py-2 border text-sm font-medium">
              1
            </button>
            <button class="bg-white border-gray-300 text-gray-500 hover:bg-gray-50 relative inline-flex items-center px-4 py-2 border text-sm font-medium">
              2
            </button>
            <button class="bg-white border-gray-300 text-gray-500 hover:bg-gray-50 relative inline-flex items-center px-4 py-2 border text-sm font-medium">
              3
            </button>
            <button class="relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50">
              <span class="sr-only">بعدی</span>
              <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z" clip-rule="evenodd" />
              </svg>
            </button>
          </nav>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user, hasPermission } = useAuth()

// آمار ناوگان
const fleetStats = ref({
  total: 25,
  active: 18,
  maintenance: 4,
  inactive: 3
})

// داده‌های نمونه ناوگان
const fleetVehicles = ref([
  {
    id: 1,
    model: 'پژو 206',
    plate: '12-345-67',
    year: '2020',
    image: '/vehicle-1.jpg',
    status: 'active',
    lastService: '2024-01-10',
    mileage: 45000,
    driver: {
      name: 'احمد محمدی',
      phone: '09123456789',
      avatar: '/driver-1.jpg'
    }
  },
  {
    id: 2,
    model: 'سایپا 111',
    plate: '98-765-43',
    year: '2019',
    image: '/vehicle-2.jpg',
    status: 'maintenance',
    lastService: '2024-01-15',
    mileage: 32000,
    driver: null
  },
  {
    id: 3,
    model: 'پراید',
    plate: '11-222-33',
    year: '2021',
    image: '/vehicle-3.jpg',
    status: 'active',
    lastService: '2024-01-08',
    mileage: 28000,
    driver: {
      name: 'علی رضایی',
      phone: '09987654321',
      avatar: '/driver-2.jpg'
    }
  }
])

// تابع فرمت کردن تاریخ
const formatDate = (dateString) => {
  return new Date(dateString).toLocaleDateString('fa-IR')
}

// تابع کلاس‌های وضعیت
const getStatusClass = (status) => {
  const classes = {
    active: 'bg-green-100 text-green-800',
    maintenance: 'bg-yellow-100 text-yellow-800',
    inactive: 'bg-red-100 text-red-800'
  }
  return classes[status] || classes.inactive
}

// تابع متن وضعیت
const getStatusText = (status) => {
  const texts = {
    active: 'فعال',
    maintenance: 'در حال سرویس',
    inactive: 'غیرفعال'
  }
  return texts[status] || 'نامشخص'
}

// تنظیمات صفحه
definePageMeta({
  title: 'مدیریت ناوگان',
  layout: 'admin-main',
  middleware: 'admin'
})
</script> 


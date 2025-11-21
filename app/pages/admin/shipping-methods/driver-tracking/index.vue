<template>
  <div class="driver-tracking-page min-h-screen p-6" dir="rtl">
    <!-- هدر صفحه -->
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900 mb-2">رهگیری رانندگان</h1>
      <p class="text-gray-600">مدیریت و رهگیری موقعیت رانندگان در زمان واقعی</p>
    </div>

    <!-- فیلترها و جستجو -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-6">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">جستجو راننده</label>
          <input 
            v-model="searchQuery" 
            type="text" 
            placeholder="نام یا شماره راننده..."
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">وضعیت</label>
          <select 
            v-model="statusFilter" 
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            <option value="">همه</option>
            <option value="active">فعال</option>
            <option value="inactive">غیرفعال</option>
            <option value="busy">مشغول</option>
            <option value="offline">آفلاین</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">منطقه</label>
          <select 
            v-model="areaFilter" 
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            <option value="">همه مناطق</option>
            <option value="tehran">تهران</option>
            <option value="isfahan">اصفهان</option>
            <option value="shiraz">شیراز</option>
            <option value="mashhad">مشهد</option>
          </select>
        </div>
        <div class="flex items-end">
          <button 
            class="w-full bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500"
            @click="refreshData"
          >
            <span class="i-heroicons-arrow-path ml-2"></span>
            به‌روزرسانی
          </button>
        </div>
      </div>
    </div>

    <!-- آمار کلی -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="p-2 bg-green-100 rounded-lg">
            <span class="i-heroicons-user-group text-green-600 text-xl"></span>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">کل رانندگان</p>
            <p class="text-2xl font-bold text-gray-900">{{ stats.totalDrivers }}</p>
          </div>
        </div>
      </div>
      
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="p-2 bg-blue-100 rounded-lg">
            <span class="i-heroicons-truck text-blue-600 text-xl"></span>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">رانندگان فعال</p>
            <p class="text-2xl font-bold text-gray-900">{{ stats.activeDrivers }}</p>
          </div>
        </div>
      </div>
      
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="p-2 bg-yellow-100 rounded-lg">
            <span class="i-heroicons-clock text-yellow-600 text-xl"></span>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">در حال تحویل</p>
            <p class="text-2xl font-bold text-gray-900">{{ stats.deliveringDrivers }}</p>
          </div>
        </div>
      </div>
      
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="p-2 bg-red-100 rounded-lg">
            <span class="i-heroicons-exclamation-triangle text-red-600 text-xl"></span>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">آفلاین</p>
            <p class="text-2xl font-bold text-gray-900">{{ stats.offlineDrivers }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- جدول رانندگان -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-medium text-gray-900">لیست رانندگان</h3>
      </div>
      
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">راننده</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">موقعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">آخرین به‌روزرسانی</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">سرعت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="driver in filteredDrivers" :key="driver.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="">
                    <img class="h-10 w-10 rounded-full" :src="driver.avatar" :alt="driver.name" />
                  </div>
                  <div class="mr-4">
                    <div class="text-sm font-medium text-gray-900">{{ driver.name }}</div>
                    <div class="text-sm text-gray-500">{{ driver.phone }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getStatusClass(driver.status)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                  {{ getStatusText(driver.status) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ driver.location }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ formatTime(driver.lastUpdate) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ driver.speed }} کیلومتر/ساعت
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button 
                  class="text-blue-600 hover:text-blue-900 ml-4"
                  @click="viewDriverDetails(driver)"
                >
                  جزئیات
                </button>
                <button 
                  class="text-green-600 hover:text-green-900"
                  @click="sendMessage(driver)"
                >
                  پیام
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      
      <!-- پیجینیشن -->
      <div class="bg-white px-4 py-3 flex items-center justify-between border-t border-gray-200 sm:px-6">
        <div class="flex-1 flex justify-between sm:hidden">
          <button 
            :disabled="currentPage === 1"
            class="relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50"
            @click="previousPage"
          >
            قبلی
          </button>
          <button 
            :disabled="currentPage === totalPages"
            class="mr-3 relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50"
            @click="nextPage"
          >
            بعدی
          </button>
        </div>
        <div class="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
          <div>
            <p class="text-sm text-gray-700">
              نمایش 
              <span class="font-medium">{{ (currentPage - 1) * pageSize + 1 }}</span>
              تا 
              <span class="font-medium">{{ Math.min(currentPage * pageSize, totalItems) }}</span>
              از 
              <span class="font-medium">{{ totalItems }}</span>
              نتیجه
            </p>
          </div>
          <div>
            <nav class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px">
              <button 
                :disabled="currentPage === 1"
                class="relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50"
                @click="previousPage"
              >
                <span class="i-heroicons-chevron-left"></span>
              </button>
              <button 
                v-for="page in visiblePages" 
                :key="page"
                :class="[
                  'relative inline-flex items-center px-4 py-2 border text-sm font-medium',
                  page === currentPage
                    ? 'z-10 bg-blue-50 border-blue-500 text-blue-600'
                    : 'bg-white border-gray-300 text-gray-500 hover:bg-gray-50'
                ]"
                @click="goToPage(page)"
              >
                {{ page }}
              </button>
              <button 
                :disabled="currentPage === totalPages"
                class="relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50"
                @click="nextPage"
              >
                <span class="i-heroicons-chevron-right"></span>
              </button>
            </nav>
          </div>
        </div>
      </div>
    </div>

    <!-- مودال جزئیات راننده -->
    <div v-if="showDriverModal" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
      <div class="relative top-20 mx-auto p-5 border w-full max-w-md sm:max-w-lg md:max-w-xl shadow-lg rounded-md bg-white">
        <div class="mt-3">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-medium text-gray-900">جزئیات راننده</h3>
            <button class="text-gray-400 hover:text-gray-600" @click="showDriverModal = false">
              <span class="i-heroicons-x-mark text-xl"></span>
            </button>
          </div>
          
          <div v-if="selectedDriver" class="space-y-4">
            <div class="flex items-center">
              <img :src="selectedDriver.avatar" :alt="selectedDriver.name" class="h-16 w-16 rounded-full" />
              <div class="mr-4">
                <h4 class="text-lg font-medium text-gray-900">{{ selectedDriver.name }}</h4>
                <p class="text-gray-500">{{ selectedDriver.phone }}</p>
              </div>
            </div>
            
            <div class="grid grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700">وضعیت</label>
                <p class="text-sm text-gray-900">{{ getStatusText(selectedDriver.status) }}</p>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700">سرعت</label>
                <p class="text-sm text-gray-900">{{ selectedDriver.speed }} کیلومتر/ساعت</p>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700">موقعیت</label>
                <p class="text-sm text-gray-900">{{ selectedDriver.location }}</p>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700">آخرین به‌روزرسانی</label>
                <p class="text-sm text-gray-900">{{ formatTime(selectedDriver.lastUpdate) }}</p>
              </div>
            </div>
            
            <div class="flex justify-end space-x-2 space-x-reverse">
              <button class="px-4 py-2 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400" @click="showDriverModal = false">
                بستن
              </button>
              <button class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700" @click="sendMessage(selectedDriver)">
                ارسال پیام
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; title?: string }) => void
</script>

<script setup lang="ts">
import { computed, ref } from 'vue';
// تعریف interface برای راننده
interface Driver {
  id: number;
  name: string;
  phone: string;
  avatar: string;
  status: 'active' | 'inactive' | 'busy' | 'offline';
  location: string;
  lastUpdate: string;
  speed: number;
}

// تعریف interface برای آمار
interface Stats {
  totalDrivers: number;
  activeDrivers: number;
  deliveringDrivers: number;
  offlineDrivers: number;
}

// متغیرهای reactive
const searchQuery = ref('');
const statusFilter = ref('');
const areaFilter = ref('');
const currentPage = ref(1);
const pageSize = ref(10);
const showDriverModal = ref(false);
const selectedDriver = ref<Driver | null>(null);

// داده‌های نمونه
const drivers = ref<Driver[]>([
  {
    id: 1,
    name: 'علی احمدی',
    phone: '09123456789',
    avatar: '/default-avatar.png',
    status: 'active',
    location: 'تهران، خیابان ولیعصر',
    lastUpdate: '2024-01-15T10:30:00Z',
    speed: 45
  },
  {
    id: 2,
    name: 'محمد رضایی',
    phone: '09187654321',
    avatar: '/default-avatar.png',
    status: 'busy',
    location: 'تهران، خیابان انقلاب',
    lastUpdate: '2024-01-15T10:25:00Z',
    speed: 0
  },
  {
    id: 3,
    name: 'حسن کریمی',
    phone: '09111111111',
    avatar: '/default-avatar.png',
    status: 'offline',
    location: 'تهران، خیابان آزادی',
    lastUpdate: '2024-01-15T09:15:00Z',
    speed: 0
  },
  {
    id: 4,
    name: 'رضا نوری',
    phone: '09222222222',
    avatar: '/default-avatar.png',
    status: 'active',
    location: 'تهران، خیابان شریعتی',
    lastUpdate: '2024-01-15T10:28:00Z',
    speed: 35
  }
]);

const stats = ref<Stats>({
  totalDrivers: 4,
  activeDrivers: 2,
  deliveringDrivers: 1,
  offlineDrivers: 1
});

// محاسبه‌های computed
const filteredDrivers = computed(() => {
  let filtered = drivers.value;
  
  if (searchQuery.value) {
    filtered = filtered.filter(driver => 
      driver.name.includes(searchQuery.value) || 
      driver.phone.includes(searchQuery.value)
    );
  }
  
  if (statusFilter.value) {
    filtered = filtered.filter(driver => driver.status === statusFilter.value);
  }
  
  return filtered;
});

const totalItems = computed(() => filteredDrivers.value.length);
const totalPages = computed(() => Math.ceil(totalItems.value / pageSize.value));

const visiblePages = computed(() => {
  const pages = [];
  const start = Math.max(1, currentPage.value - 2);
  const end = Math.min(totalPages.value, currentPage.value + 2);
  
  for (let i = start; i <= end; i++) {
    pages.push(i);
  }
  
  return pages;
});

// متدها
const getStatusClass = (status: string) => {
  switch (status) {
    case 'active':
      return 'bg-green-100 text-green-800';
    case 'busy':
      return 'bg-yellow-100 text-yellow-800';
    case 'offline':
      return 'bg-red-100 text-red-800';
    default:
      return 'bg-gray-100 text-gray-800';
  }
};

const getStatusText = (status: string) => {
  switch (status) {
    case 'active':
      return 'فعال';
    case 'busy':
      return 'مشغول';
    case 'offline':
      return 'آفلاین';
    case 'inactive':
      return 'غیرفعال';
    default:
      return 'نامشخص';
  }
};

const formatTime = (timeString: string) => {
  const date = new Date(timeString);
  return date.toLocaleString('fa-IR');
};

const refreshData = () => {
  // اینجا می‌توانید API call برای به‌روزرسانی داده‌ها اضافه کنید
};

const viewDriverDetails = (driver: Driver) => {
  selectedDriver.value = driver;
  showDriverModal.value = true;
};

const sendMessage = (driver: Driver) => {
  // اینجا می‌توانید مودال ارسال پیام را باز کنید
  alert(`ارسال پیام به ${driver.name}`);
};

const previousPage = () => {
  if (currentPage.value > 1) {
    currentPage.value--;
  }
};

const nextPage = () => {
  if (currentPage.value < totalPages.value) {
    currentPage.value++;
  }
};

const goToPage = (page: number) => {
  currentPage.value = page;
};

// تنظیمات صفحه
definePageMeta({
  title: 'رهگیری رانندگان',
  layout: 'admin-main'
});
</script>

<style scoped>
.driver-tracking-page {
  min-height: 100vh;
  background-color: #f9fafb;
}

/* انیمیشن‌ها */
.fade-enter-active, .fade-leave-active {
  transition: opacity 0.3s;
}

.fade-enter-from, .fade-leave-to {
  opacity: 0;
}
</style> 

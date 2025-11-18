<template>
  <div class="notifications-page" dir="rtl">
    <!-- هدر صفحه -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-6">
      <h1 class="text-2xl font-bold text-gray-900 mb-2">پیامک و اعلان‌ها</h1>
      <p class="text-gray-600">مدیریت پیامک‌ها، اعلان‌ها و تنظیمات ارتباطی</p>
    </div>

    <!-- آمار کلی -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-6">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div class="rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="p-2 bg-blue-100 rounded-lg">
            <span class="i-heroicons-envelope text-blue-600 text-xl"></span>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">کل پیامک‌ها</p>
            <p class="text-2xl font-bold text-gray-900">{{ stats.totalSms }}</p>
          </div>
        </div>
      </div>
      
      <div class="rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="p-2 bg-green-100 rounded-lg">
            <span class="i-heroicons-bell text-green-600 text-xl"></span>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">اعلان‌های ارسال شده</p>
            <p class="text-2xl font-bold text-gray-900">{{ stats.sentNotifications }}</p>
          </div>
        </div>
      </div>
      
      <div class="rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="p-2 bg-yellow-100 rounded-lg">
            <span class="i-heroicons-clock text-yellow-600 text-xl"></span>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">در انتظار ارسال</p>
            <p class="text-2xl font-bold text-gray-900">{{ stats.pendingNotifications }}</p>
          </div>
        </div>
      </div>
      
      <div class="rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="p-2 bg-red-100 rounded-lg">
            <span class="i-heroicons-exclamation-triangle text-red-600 text-xl"></span>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">خطاهای ارسال</p>
            <p class="text-2xl font-bold text-gray-900">{{ stats.failedNotifications }}</p>
          </div>
        </div>
      </div>
    </div>
    </div>

    <!-- تب‌های اصلی -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 mb-6">
      <div class="border-b border-gray-200">
        <nav class="flex space-x-8 space-x-reverse px-6">
          <button 
            v-for="tab in tabs" 
            :key="tab.key"
            @click="activeTab = tab.key"
            :class="[
              'py-4 px-1 border-b-2 font-medium text-sm',
              activeTab === tab.key
                ? 'border-blue-500 text-blue-600'
                : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
            ]"
          >
            {{ tab.title }}
          </button>
        </nav>
      </div>

      <!-- محتوای تب‌ها -->
      <div class="p-6">
        <!-- تب پیامک‌ها -->
        <div v-if="activeTab === 'sms'">
          <div class="flex justify-between items-center mb-4">
            <h3 class="text-lg font-medium text-gray-900">مدیریت پیامک‌ها</h3>
            <button 
              v-if="hasPermission('notification.create')"
              class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700"
            >
              ارسال پیامک جدید
            </button>
          </div>
          
          <div class="overflow-x-auto">
            <table class="min-w-full divide-y divide-gray-200">
              <thead>
                <tr>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase">گیرنده</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase">متن پیام</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase">وضعیت</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase">تاریخ ارسال</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase">عملیات</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-gray-200">
                <tr v-for="sms in smsList" :key="sms.id">
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ sms.recipient }}</td>
                  <td class="px-6 py-4 text-sm text-gray-900 max-w-xs truncate">{{ sms.message }}</td>
                  <td class="px-6 py-4 whitespace-nowrap">
                    <span :class="getSmsStatusClass(sms.status)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                      {{ getSmsStatusText(sms.status) }}
                    </span>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ formatDate(sms.sentAt) }}</td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                    <button 
                      v-if="hasPermission('notification.read')"
                      class="text-blue-600 hover:text-blue-900 ml-3"
                    >
                      مشاهده
                    </button>
                    <button 
                      v-if="hasPermission('notification.delete')"
                      class="text-red-600 hover:text-red-900"
                    >
                      حذف
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <!-- تب اعلان‌ها -->
        <div v-else-if="activeTab === 'notifications'">
          <div class="flex justify-between items-center mb-4">
            <h3 class="text-lg font-medium text-gray-900">مدیریت اعلان‌ها</h3>
            <button 
              v-if="hasPermission('notification.create')"
              class="bg-green-600 text-white px-4 py-2 rounded-md hover:bg-green-700"
            >
              ایجاد اعلان جدید
            </button>
          </div>
          
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            <div v-for="notification in notificationsList" :key="notification.id" 
                 class="border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow">
              <div class="flex items-start justify-between mb-2">
                <h4 class="font-medium text-gray-900">{{ notification.title }}</h4>
                <span :class="getNotificationTypeClass(notification.type)" 
                      class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                  {{ getNotificationTypeText(notification.type) }}
                </span>
              </div>
              <p class="text-sm text-gray-600 mb-3">{{ notification.message }}</p>
              <div class="flex items-center justify-between text-xs text-gray-500">
                <span>{{ formatDate(notification.createdAt) }}</span>
                <div class="flex space-x-2 space-x-reverse">
                  <button 
                    v-if="hasPermission('notification.update')"
                    class="text-blue-600 hover:text-blue-900"
                  >
                    ویرایش
                  </button>
                  <button 
                    v-if="hasPermission('notification.delete')"
                    class="text-red-600 hover:text-red-900"
                  >
                    حذف
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- تب قالب‌ها -->
        <div v-else-if="activeTab === 'templates'">
          <div class="flex justify-between items-center mb-4">
            <h3 class="text-lg font-medium text-gray-900">قالب‌های پیام</h3>
            <button 
              v-if="hasPermission('notification.create')"
              class="bg-purple-600 text-white px-4 py-2 rounded-md hover:bg-purple-700"
            >
              افزودن قالب جدید
            </button>
          </div>
          
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div v-for="template in templatesList" :key="template.id" 
                 class="border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow">
              <div class="flex items-start justify-between mb-2">
                <h4 class="font-medium text-gray-900">{{ template.name }}</h4>
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full text-gray-800">
                  {{ template.category }}
                </span>
              </div>
              <p class="text-sm text-gray-600 mb-3">{{ template.content }}</p>
              <div class="flex items-center justify-between">
                <span class="text-xs text-gray-500">{{ template.usageCount }} بار استفاده شده</span>
                <div class="flex space-x-2 space-x-reverse">
                  <button 
                    v-if="hasPermission('notification.update')"
                    class="text-blue-600 hover:text-blue-900 text-sm"
                  >
                    ویرایش
                  </button>
                  <button 
                    v-if="hasPermission('notification.use')"
                    class="text-green-600 hover:text-green-900 text-sm"
                  >
                    استفاده
                  </button>
                  <button 
                    v-if="hasPermission('notification.delete')"
                    class="text-red-600 hover:text-red-900 text-sm"
                  >
                    حذف
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- تب تنظیمات -->
        <div v-else-if="activeTab === 'settings'">
          <h3 class="text-lg font-medium text-gray-900 mb-4">تنظیمات پیامک و اعلان</h3>
          
          <div class="space-y-6">
            <!-- تنظیمات SMS -->
            <div class="bg-white p-6 rounded-lg">
              <h4 class="font-medium text-gray-900 mb-3">تنظیمات SMS</h4>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-1">API Key</label>
                  <input type="password" class="w-full px-3 py-2 border border-gray-300 rounded-md" 
                         placeholder="کلید API پیامک" />
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-1">شماره فرستنده</label>
                  <input type="text" class="w-full px-3 py-2 border border-gray-300 rounded-md" 
                         placeholder="شماره فرستنده" />
                </div>
              </div>
            </div>

            <!-- تنظیمات اعلان‌ها -->
            <div class="bg-gray-50 p-6 rounded-lg">
              <h4 class="font-medium text-gray-900 mb-3">تنظیمات اعلان‌ها</h4>
              <div class="space-y-3">
                <div class="flex items-center justify-between">
                  <span class="text-sm text-gray-700">اعلان سفارش جدید</span>
                  <button class="relative inline-flex flex-shrink-0 h-6 w-11 border-2 border-transparent rounded-full cursor-pointer transition-colors ease-in-out duration-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500" 
                          :class="settings.newOrderNotification ? 'bg-blue-600' : 'bg-gray-200'">
                    <span class="pointer-events-none inline-block h-5 w-5 rounded-full bg-white shadow transform ring-0 transition ease-in-out duration-200" 
                          :class="settings.newOrderNotification ? 'translate-x-5' : 'translate-x-0'"></span>
                  </button>
                </div>
                <div class="flex items-center justify-between">
                  <span class="text-sm text-gray-700">اعلان تغییر وضعیت سفارش</span>
                  <button class="relative inline-flex flex-shrink-0 h-6 w-11 border-2 border-transparent rounded-full cursor-pointer transition-colors ease-in-out duration-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500" 
                          :class="settings.orderStatusNotification ? 'bg-blue-600' : 'bg-gray-200'">
                    <span class="pointer-events-none inline-block h-5 w-5 rounded-full bg-white shadow transform ring-0 transition ease-in-out duration-200" 
                          :class="settings.orderStatusNotification ? 'translate-x-5' : 'translate-x-0'"></span>
                  </button>
                </div>
                <div class="flex items-center justify-between">
                  <span class="text-sm text-gray-700">اعلان موجودی محصول</span>
                  <button class="relative inline-flex flex-shrink-0 h-6 w-11 border-2 border-transparent rounded-full cursor-pointer transition-colors ease-in-out duration-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500" 
                          :class="settings.inventoryNotification ? 'bg-blue-600' : 'bg-gray-200'">
                    <span class="pointer-events-none inline-block h-5 w-5 rounded-full bg-white shadow transform ring-0 transition ease-in-out duration-200" 
                          :class="settings.inventoryNotification ? 'translate-x-5' : 'translate-x-0'"></span>
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { title?: string; layout?: string; middleware?: string | string[] }) => void
</script>

<script setup lang="ts">
import { ref } from 'vue';
import { useAuth } from '~/composables/useAuth';
// Import useAuth for permission checking
const { user, hasPermission } = useAuth();

// تعریف interface ها
interface Sms {
  id: number;
  recipient: string;
  message: string;
  status: 'sent' | 'pending' | 'failed';
  sentAt: string;
}

interface Notification {
  id: number;
  title: string;
  message: string;
  type: 'info' | 'warning' | 'success' | 'error';
  createdAt: string;
}

interface Template {
  id: number;
  name: string;
  content: string;
  category: string;
  usageCount: number;
}

interface Settings {
  newOrderNotification: boolean;
  orderStatusNotification: boolean;
  inventoryNotification: boolean;
}

// متغیرهای reactive
const activeTab = ref('sms');

const tabs = ref([
  { key: 'sms', title: 'پیامک‌ها' },
  { key: 'notifications', title: 'اعلان‌ها' },
  { key: 'templates', title: 'قالب‌ها' },
  { key: 'settings', title: 'تنظیمات' }
]);

const stats = ref({
  totalSms: 1250,
  sentNotifications: 890,
  pendingNotifications: 45,
  failedNotifications: 12
});

const smsList = ref<Sms[]>([
  {
    id: 1,
    recipient: '09123456789',
    message: 'سفارش شما با موفقیت ثبت شد. کد پیگیری: 12345',
    status: 'sent',
    sentAt: '2024-01-15T10:30:00Z'
  },
  {
    id: 2,
    recipient: '09987654321',
    message: 'وضعیت سفارش شما به "در حال ارسال" تغییر یافت',
    status: 'pending',
    sentAt: '2024-01-15T11:15:00Z'
  },
  {
    id: 3,
    recipient: '09111111111',
    message: 'محصول مورد نظر شما موجود شد',
    status: 'failed',
    sentAt: '2024-01-15T09:45:00Z'
  }
]);

const notificationsList = ref<Notification[]>([
  {
    id: 1,
    title: 'سفارش جدید',
    message: 'سفارش جدیدی با شماره #12345 ثبت شده است',
    type: 'info',
    createdAt: '2024-01-15T10:30:00Z'
  },
  {
    id: 2,
    title: 'موجودی کم',
    message: 'محصول "لپ‌تاپ اپل" موجودی کمی دارد',
    type: 'warning',
    createdAt: '2024-01-15T09:15:00Z'
  },
  {
    id: 3,
    title: 'ارسال موفق',
    message: 'سفارش #12340 با موفقیت ارسال شد',
    type: 'success',
    createdAt: '2024-01-15T08:45:00Z'
  }
]);

const templatesList = ref<Template[]>([
  {
    id: 1,
    name: 'تایید سفارش',
    content: 'سفارش شما با موفقیت ثبت شد. کد پیگیری: {order_id}',
    category: 'سفارش',
    usageCount: 156
  },
  {
    id: 2,
    name: 'تغییر وضعیت',
    content: 'وضعیت سفارش شما به "{status}" تغییر یافت',
    category: 'وضعیت',
    usageCount: 89
  },
  {
    id: 3,
    name: 'موجودی',
    content: 'محصول "{product_name}" موجود شد',
    category: 'موجودی',
    usageCount: 34
  }
]);

const settings = ref<Settings>({
  newOrderNotification: true,
  orderStatusNotification: true,
  inventoryNotification: false
});

// متدها
const getSmsStatusClass = (status: string) => {
  switch (status) {
    case 'sent':
      return 'bg-green-100 text-green-800';
    case 'pending':
      return 'bg-yellow-100 text-yellow-800';
    case 'failed':
      return 'bg-red-100 text-red-800';
    default:
      return 'bg-gray-100 text-gray-800';
  }
};

const getSmsStatusText = (status: string) => {
  switch (status) {
    case 'sent':
      return 'ارسال شده';
    case 'pending':
      return 'در انتظار';
    case 'failed':
      return 'ناموفق';
    default:
      return 'نامشخص';
  }
};

const getNotificationTypeClass = (type: string) => {
  switch (type) {
    case 'info':
      return 'bg-blue-100 text-blue-800';
    case 'warning':
      return 'bg-yellow-100 text-yellow-800';
    case 'success':
      return 'bg-green-100 text-green-800';
    case 'error':
      return 'bg-red-100 text-red-800';
    default:
      return 'bg-gray-100 text-gray-800';
  }
};

const getNotificationTypeText = (type: string) => {
  switch (type) {
    case 'info':
      return 'اطلاعات';
    case 'warning':
      return 'هشدار';
    case 'success':
      return 'موفق';
    case 'error':
      return 'خطا';
    default:
      return 'نامشخص';
  }
};

const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleString('fa-IR');
};

// تنظیمات صفحه
definePageMeta({
  title: 'پیامک و اعلان‌ها',
  layout: 'admin-main',
  middleware: 'admin'
});
</script>

<style scoped>
</style> 


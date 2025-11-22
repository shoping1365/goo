<template>
  <div class="support-ticket-page" dir="rtl">
    <!-- هدر صفحه -->
    <div class="bg-white w-full rounded-lg shadow-sm border border-gray-200 mb-6">
      <div class="mb-6 p-6">
        <h1 class="text-2xl font-bold text-gray-900 mb-2">پشتیبانی و تیکت‌ها</h1>
        <p class="text-gray-600">مدیریت درخواست‌های پشتیبانی و تیکت‌های کاربران</p>
      </div>
    </div>

    <!-- آمار کلی -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="p-2 bg-blue-100 rounded-lg">
            <span class="i-heroicons-ticket text-blue-600 text-xl"></span>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">کل تیکت‌ها</p>
            <p class="text-2xl font-bold text-gray-900">{{ stats.totalTickets }}</p>
          </div>
        </div>
      </div>
      
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="p-2 bg-yellow-100 rounded-lg">
            <span class="i-heroicons-clock text-yellow-600 text-xl"></span>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">در انتظار پاسخ</p>
            <p class="text-2xl font-bold text-gray-900">{{ stats.pendingTickets }}</p>
          </div>
        </div>
      </div>
      
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="p-2 bg-green-100 rounded-lg">
            <span class="i-heroicons-check-circle text-green-600 text-xl"></span>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">حل شده</p>
            <p class="text-2xl font-bold text-gray-900">{{ stats.resolvedTickets }}</p>
          </div>
        </div>
      </div>
      
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="p-2 bg-red-100 rounded-lg">
            <span class="i-heroicons-exclamation-triangle text-red-600 text-xl"></span>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">بحرانی</p>
            <p class="text-2xl font-bold text-gray-900">{{ stats.criticalTickets }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- فیلترها و جستجو -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-6">
      <div class="grid grid-cols-1 md:grid-cols-5 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">جستجو</label>
          <input 
            v-model="searchQuery" 
            type="text" 
            placeholder="جستجو در تیکت‌ها..."
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">وضعیت</label>
          <select 
            v-model="statusFilter" 
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            <option value="">همه وضعیت‌ها</option>
            <option value="open">باز</option>
            <option value="pending">در انتظار</option>
            <option value="in_progress">در حال بررسی</option>
            <option value="resolved">حل شده</option>
            <option value="closed">بسته</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">اولویت</label>
          <select 
            v-model="priorityFilter" 
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            <option value="">همه اولویت‌ها</option>
            <option value="low">کم</option>
            <option value="medium">متوسط</option>
            <option value="high">زیاد</option>
            <option value="critical">بحرانی</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">دپارتمان</label>
          <select 
            v-model="departmentFilter" 
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            <option value="">همه دپارتمان‌ها</option>
            <option value="technical">فنی</option>
            <option value="billing">مالی</option>
            <option value="sales">فروش</option>
            <option value="general">عمومی</option>
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

    <!-- جدول تیکت‌ها -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <div class="flex justify-between items-center">
          <h3 class="text-lg font-medium text-gray-900">لیست تیکت‌ها</h3>
          <button class="bg-green-600 text-white px-4 py-2 rounded-md hover:bg-green-700">
            تیکت جدید
          </button>
        </div>
      </div>
      
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead>
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تیکت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کاربر</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">موضوع</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">اولویت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">آخرین به‌روزرسانی</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="ticket in filteredTickets" :key="ticket.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm font-medium text-gray-900">#{{ ticket.id }}</div>
                <div class="text-sm text-gray-500">{{ ticket.department }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="flex-shrink-0 h-8 w-8">
                    <img class="h-8 w-8 rounded-full" :src="ticket.user.avatar" :alt="ticket.user.name" />
                  </div>
                  <div class="mr-3">
                    <div class="text-sm font-medium text-gray-900">{{ ticket.user.name }}</div>
                    <div class="text-sm text-gray-500">{{ ticket.user.email }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4">
                <div class="text-sm font-medium text-gray-900">{{ ticket.subject }}</div>
                <div class="text-sm text-gray-500 max-w-xs truncate">{{ ticket.message }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getStatusClass(ticket.status)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                  {{ getStatusText(ticket.status) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getPriorityClass(ticket.priority)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                  {{ getPriorityText(ticket.priority) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ formatDate(ticket.updatedAt) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button 
                  class="text-blue-600 hover:text-blue-900 ml-3"
                  @click="viewTicket(ticket)"
                >
                  مشاهده
                </button>
                <button 
                  class="text-green-600 hover:text-green-900 ml-3"
                  @click="assignTicket(ticket)"
                >
                  واگذار
                </button>
                <button 
                  class="text-red-600 hover:text-red-900"
                  @click="closeTicket(ticket)"
                >
                  بستن
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      
      <!-- پیجینیشن -->
      <div class="px-4 py-3 flex items-center justify-between border-t border-gray-200 sm:px-6">
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

    <!-- مودال جزئیات تیکت -->
          <div v-if="showTicketModal" class="fixed inset-0 bg-opacity-50 overflow-y-auto h-full w-full z-50">
      <div class="relative top-20 mx-auto p-5 border w-full max-w-4xl shadow-lg rounded-md bg-white">
        <div class="mt-3">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-medium text-gray-900">جزئیات تیکت #{{ selectedTicket?.id }}</h3>
            <button class="text-gray-400 hover:text-gray-600" @click="showTicketModal = false">
              <span class="i-heroicons-x-mark text-xl"></span>
            </button>
          </div>
          
          <div v-if="selectedTicket" class="space-y-6">
            <!-- اطلاعات تیکت -->
            <div class="p-6 rounded-lg">
              <div class="grid grid-cols-2 md:grid-cols-4 gap-6">
                <div>
                  <label class="block text-sm font-medium text-gray-700">وضعیت</label>
                  <span :class="getStatusClass(selectedTicket.status)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                    {{ getStatusText(selectedTicket.status) }}
                  </span>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700">اولویت</label>
                  <span :class="getPriorityClass(selectedTicket.priority)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                    {{ getPriorityText(selectedTicket.priority) }}
                  </span>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700">دپارتمان</label>
                  <p class="text-sm text-gray-900">{{ selectedTicket.department }}</p>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700">تاریخ ایجاد</label>
                  <p class="text-sm text-gray-900">{{ formatDate(selectedTicket.createdAt) }}</p>
                </div>
              </div>
            </div>

            <!-- پیام اصلی -->
            <div>
              <h4 class="font-medium text-gray-900 mb-2">پیام اصلی</h4>
              <div class="border border-gray-200 rounded-lg p-6">
                <div class="flex items-start mb-3">
                  <img :src="selectedTicket.user.avatar" :alt="selectedTicket.user.name" class="h-8 w-8 rounded-full" />
                  <div class="mr-3">
                    <p class="text-sm font-medium text-gray-900">{{ selectedTicket.user.name }}</p>
                    <p class="text-xs text-gray-500">{{ formatDate(selectedTicket.createdAt) }}</p>
                  </div>
                </div>
                <p class="text-sm text-gray-700">{{ selectedTicket.message }}</p>
              </div>
            </div>

            <!-- پاسخ‌ها -->
            <div>
              <h4 class="font-medium text-gray-900 mb-2">پاسخ‌ها</h4>
              <div class="space-y-3">
                <div
v-for="reply in selectedTicket.replies" :key="reply.id" 
                     class="bg-white border border-gray-200 rounded-lg p-6">
                  <div class="flex items-start mb-3">
                    <img :src="reply.user.avatar" :alt="reply.user.name" class="h-8 w-8 rounded-full" />
                    <div class="mr-3">
                      <p class="text-sm font-medium text-gray-900">{{ reply.user.name }}</p>
                      <p class="text-xs text-gray-500">{{ formatDate(reply.createdAt) }}</p>
                    </div>
                  </div>
                  <p class="text-sm text-gray-700">{{ reply.message }}</p>
                </div>
              </div>
            </div>

            <!-- فرم پاسخ -->
            <div>
              <h4 class="font-medium text-gray-900 mb-2">ارسال پاسخ</h4>
              <div class="space-y-3">
                <textarea 
                  v-model="replyMessage"
                  rows="4"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                  placeholder="پاسخ خود را بنویسید..."
                ></textarea>
                <div class="flex justify-end space-x-2 space-x-reverse">
                  <button class="px-4 py-2 text-gray-700 rounded-md hover:bg-gray-400" @click="showTicketModal = false">
                    انصراف
                  </button>
                  <button class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700" @click="sendReply">
                    ارسال پاسخ
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
declare const definePageMeta: (meta: { layout?: string; title?: string; middleware?: string | string[] }) => void
</script>

<script setup lang="ts">
import { computed, ref } from 'vue';

// تعریف interface ها
interface User {
  id: number;
  name: string;
  email: string;
  avatar: string;
}

interface Reply {
  id: number;
  message: string;
  user: User;
  createdAt: string;
}

interface Ticket {
  id: number;
  subject: string;
  message: string;
  status: 'open' | 'pending' | 'in_progress' | 'resolved' | 'closed';
  priority: 'low' | 'medium' | 'high' | 'critical';
  department: string;
  user: User;
  createdAt: string;
  updatedAt: string;
  replies: Reply[];
}

// متغیرهای reactive
const searchQuery = ref('');
const statusFilter = ref('');
const priorityFilter = ref('');
const departmentFilter = ref('');
const currentPage = ref(1);
const pageSize = ref(10);
const showTicketModal = ref(false);
const selectedTicket = ref<Ticket | null>(null);
const replyMessage = ref('');

const stats = ref({
  totalTickets: 156,
  pendingTickets: 23,
  resolvedTickets: 89,
  criticalTickets: 5
});

const ticketsList = ref<Ticket[]>([
  {
    id: 1001,
    subject: 'مشکل در ورود به حساب کاربری',
    message: 'من نمی‌توانم وارد حساب کاربری خود شوم. پیام خطا دریافت می‌کنم.',
    status: 'open',
    priority: 'high',
    department: 'فنی',
    user: {
      id: 1,
      name: 'علی احمدی',
      email: 'ali@example.com',
      avatar: '/user-1.jpg'
    },
    createdAt: '2024-01-15T10:30:00Z',
    updatedAt: '2024-01-15T10:30:00Z',
    replies: []
  },
  {
    id: 1002,
    subject: 'درخواست بازگشت وجه',
    message: 'می‌خواهم وجه سفارش شماره 12345 را بازگردانید.',
    status: 'pending',
    priority: 'medium',
    department: 'مالی',
    user: {
      id: 2,
      name: 'فاطمه رضایی',
      email: 'fateme@example.com',
      avatar: '/user-2.jpg'
    },
    createdAt: '2024-01-15T09:15:00Z',
    updatedAt: '2024-01-15T11:20:00Z',
    replies: [
      {
        id: 1,
        message: 'درخواست شما دریافت شد. در حال بررسی هستیم.',
        user: {
          id: 101,
          name: 'پشتیبانی مالی',
          email: 'finance@company.com',
          avatar: '/support-1.jpg'
        },
        createdAt: '2024-01-15T11:20:00Z'
      }
    ]
  },
  {
    id: 1003,
    subject: 'سوال درباره محصول',
    message: 'آیا این محصول گارانتی دارد؟',
    status: 'resolved',
    priority: 'low',
    department: 'عمومی',
    user: {
      id: 3,
      name: 'محمد کریمی',
      email: 'mohammad@example.com',
      avatar: '/user-3.jpg'
    },
    createdAt: '2024-01-14T16:45:00Z',
    updatedAt: '2024-01-15T08:30:00Z',
    replies: [
      {
        id: 2,
        message: 'بله، این محصول 18 ماه گارانتی دارد.',
        user: {
          id: 102,
          name: 'پشتیبانی عمومی',
          email: 'support@company.com',
          avatar: '/support-2.jpg'
        },
        createdAt: '2024-01-15T08:30:00Z'
      }
    ]
  }
]);

// محاسبه‌های computed
const filteredTickets = computed(() => {
  let filtered = ticketsList.value;
  
  if (searchQuery.value) {
    filtered = filtered.filter(ticket => 
      ticket.subject.includes(searchQuery.value) || 
      ticket.message.includes(searchQuery.value) ||
      ticket.user.name.includes(searchQuery.value)
    );
  }
  
  if (statusFilter.value) {
    filtered = filtered.filter(ticket => ticket.status === statusFilter.value);
  }
  
  if (priorityFilter.value) {
    filtered = filtered.filter(ticket => ticket.priority === priorityFilter.value);
  }
  
  if (departmentFilter.value) {
    filtered = filtered.filter(ticket => ticket.department === departmentFilter.value);
  }
  
  return filtered;
});

const totalItems = computed(() => filteredTickets.value.length);
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
    case 'open':
      return 'bg-blue-100 text-blue-800';
    case 'pending':
      return 'bg-yellow-100 text-yellow-800';
    case 'in_progress':
      return 'bg-purple-100 text-purple-800';
    case 'resolved':
      return 'bg-green-100 text-green-800';
    case 'closed':
      return '';
    default:
      return '';
  }
};

const getStatusText = (status: string) => {
  switch (status) {
    case 'open':
      return 'باز';
    case 'pending':
      return 'در انتظار';
    case 'in_progress':
      return 'در حال بررسی';
    case 'resolved':
      return 'حل شده';
    case 'closed':
      return 'بسته';
    default:
      return 'نامشخص';
  }
};

const getPriorityClass = (priority: string) => {
  switch (priority) {
    case 'low':
      return 'bg-green-100 text-green-800';
    case 'medium':
      return 'bg-yellow-100 text-yellow-800';
    case 'high':
      return 'bg-orange-100 text-orange-800';
    case 'critical':
      return 'bg-red-100 text-red-800';
    default:
      return '';
  }
};

const getPriorityText = (priority: string) => {
  switch (priority) {
    case 'low':
      return 'کم';
    case 'medium':
      return 'متوسط';
    case 'high':
      return 'زیاد';
    case 'critical':
      return 'بحرانی';
    default:
      return 'نامشخص';
  }
};

const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleString('fa-IR');
};

const refreshData = () => {
  // اینجا می‌توانید API call برای به‌روزرسانی داده‌ها اضافه کنید
};

const viewTicket = (ticket: Ticket) => {
  selectedTicket.value = ticket;
  showTicketModal.value = true;
};

const assignTicket = (ticket: Ticket) => {
  // اینجا می‌توانید مودال واگذاری تیکت را باز کنید
  alert(`واگذاری تیکت #${ticket.id}`);
};

const closeTicket = (ticket: Ticket) => {
  // اینجا می‌توانید تیکت را ببندید
  if (confirm(`آیا از بستن تیکت #${ticket.id} اطمینان دارید؟`)) {
  }
};

const sendReply = () => {
  if (replyMessage.value.trim()) {
    // اینجا می‌توانید پاسخ را ارسال کنید
    replyMessage.value = '';
    showTicketModal.value = false;
  }
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
  title: 'پشتیبانی و تیکت‌ها',
  layout: 'admin-main',
  middleware: 'admin'
});
</script>

<style scoped>
.support-tickets-page {
  min-height: 100vh;
}
</style> 

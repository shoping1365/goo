<template>
  <div class="bg-white/80 backdrop-blur-md rounded-2xl shadow-lg border border-gray-200/50 p-6">
    <div class="flex items-center justify-between mb-6">
      <h3 class="text-lg font-semibold text-gray-900 flex items-center">
        <svg class="w-5 h-5 ml-2 text-orange-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
        </svg>
        مدیریت تیکت‌ها
      </h3>
      <button 
        class="px-4 py-2 text-sm font-medium text-white bg-orange-500 rounded-lg hover:bg-orange-600 transition-colors"
        @click="showNewTicketModal = true"
      >
        تیکت جدید
      </button>
    </div>

    <!-- Ticket Stats -->
    <div class="grid grid-cols-2 lg:grid-cols-4 gap-6 mb-6">
      <div class="bg-red-50 rounded-lg p-3 text-center">
        <div class="text-xl font-bold text-red-600">{{ ticketStats.urgent }}</div>
        <div class="text-sm text-red-600">فوری</div>
      </div>
      <div class="bg-yellow-50 rounded-lg p-3 text-center">
        <div class="text-xl font-bold text-yellow-600">{{ ticketStats.pending }}</div>
        <div class="text-sm text-yellow-600">در انتظار</div>
      </div>
      <div class="bg-blue-50 rounded-lg p-3 text-center">
        <div class="text-xl font-bold text-blue-600">{{ ticketStats.inProgress }}</div>
        <div class="text-sm text-blue-600">در حال بررسی</div>
      </div>
      <div class="bg-green-50 rounded-lg p-3 text-center">
        <div class="text-xl font-bold text-green-600">{{ ticketStats.resolved }}</div>
        <div class="text-sm text-green-600">حل شده</div>
      </div>
    </div>

    <!-- Ticket Filters -->
    <div class="flex flex-wrap gap-2 mb-6">
      <button 
        v-for="filter in ticketFilters"
        :key="filter.id"
        :class="[
          activeTicketFilter === filter.id
            ? 'bg-orange-500 text-white'
            : 'bg-gray-100 text-gray-700 hover:bg-gray-200',
          'px-3 py-1 text-sm rounded-lg transition-colors'
        ]"
        @click="activeTicketFilter = filter.id"
      >
        {{ filter.name }}
      </button>
    </div>

    <!-- Tickets List -->
    <div class="space-y-3 max-h-96 overflow-y-auto">
      <div 
        v-for="ticket in filteredTickets" 
        :key="ticket.id"
        :class="[
          selectedTicket?.id === ticket.id ? 'ring-2 ring-orange-500' : '',
          'p-6 bg-gray-50 rounded-lg cursor-pointer hover:bg-gray-100 transition-colors'
        ]"
        @click="selectTicket(ticket)"
      >
        <div class="flex items-start justify-between">
          <div class="flex-1">
            <div class="flex items-center space-x-2 space-x-reverse mb-2">
              <span class="text-sm font-medium text-gray-900">#{{ ticket.id }}</span>
              <span 
                :class="[
                  getPriorityColor(ticket.priority),
                  'px-2 py-1 text-xs font-medium rounded-full'
                ]"
              >
                {{ getPriorityText(ticket.priority) }}
              </span>
              <span 
                :class="[
                  getStatusColor(ticket.status),
                  'px-2 py-1 text-xs font-medium rounded-full'
                ]"
              >
                {{ getStatusText(ticket.status) }}
              </span>
            </div>
            <h4 class="text-sm font-medium text-gray-900 mb-1">{{ ticket.title }}</h4>
            <p class="text-xs text-gray-600 mb-2">{{ ticket.description }}</p>
            <div class="flex items-center justify-between text-xs text-gray-500">
              <span>{{ ticket.customerName }}</span>
              <span>{{ formatTime(ticket.createdAt) }}</span>
            </div>
          </div>
          <div class="flex items-center space-x-2 space-x-reverse">
            <div 
              v-if="ticket.unreadMessages > 0"
              class="w-5 h-5 bg-red-500 text-white text-xs rounded-full flex items-center justify-center"
            >
              {{ ticket.unreadMessages }}
            </div>
            <button 
              class="text-xs text-blue-600 hover:text-blue-800"
              @click.stop="assignTicket(ticket)"
            >
              {{ ticket.assignedTo ? 'تغییر مسئول' : 'اختصاص' }}
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Selected Ticket Details -->
    <div v-if="selectedTicket" class="mt-6 p-6 bg-gray-50 rounded-lg">
      <div class="flex items-center justify-between mb-4">
        <h4 class="text-lg font-medium text-gray-900">جزئیات تیکت #{{ selectedTicket.id }}</h4>
        <div class="flex space-x-2 space-x-reverse">
          <button 
            class="px-3 py-1 text-xs bg-blue-500 text-white rounded-lg hover:bg-blue-600"
            @click="updateTicketStatus(selectedTicket, 'in_progress')"
          >
            شروع بررسی
          </button>
          <button 
            class="px-3 py-1 text-xs bg-green-500 text-white rounded-lg hover:bg-green-600"
            @click="updateTicketStatus(selectedTicket, 'resolved')"
          >
            حل شده
          </button>
        </div>
      </div>
      
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <div>
          <h5 class="text-sm font-medium text-gray-900 mb-2">اطلاعات مشتری</h5>
          <div class="text-sm text-gray-600 space-y-1">
            <div>نام: {{ selectedTicket.customerName }}</div>
            <div>ایمیل: {{ selectedTicket.customerEmail }}</div>
            <div>تلفن: {{ selectedTicket.customerPhone }}</div>
          </div>
        </div>
        <div>
          <h5 class="text-sm font-medium text-gray-900 mb-2">اطلاعات تیکت</h5>
          <div class="text-sm text-gray-600 space-y-1">
            <div>دسته‌بندی: {{ selectedTicket.category }}</div>
            <div>اولویت: {{ getPriorityText(selectedTicket.priority) }}</div>
            <div>وضعیت: {{ getStatusText(selectedTicket.status) }}</div>
            <div>تاریخ ایجاد: {{ formatDate(selectedTicket.createdAt) }}</div>
          </div>
        </div>
      </div>

      <!-- Ticket Messages -->
      <div class="mt-4">
        <h5 class="text-sm font-medium text-gray-900 mb-2">پیام‌های تیکت</h5>
        <div class="space-y-2 max-h-32 overflow-y-auto">
          <div 
            v-for="message in selectedTicket.messages" 
            :key="message.id"
            class="p-2 bg-white rounded text-xs"
          >
            <div class="flex items-center justify-between mb-1">
              <span class="font-medium">{{ message.sender }}</span>
              <span class="text-gray-500">{{ formatTime(message.timestamp) }}</span>
            </div>
            <div class="text-gray-700">{{ message.content }}</div>
          </div>
        </div>
      </div>
    </div>

    <!-- New Ticket Modal -->
    <div v-if="showNewTicketModal" class="fixed inset-0 z-50 overflow-y-auto" dir="rtl">
      <div class="flex items-center justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
        <div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" @click="showNewTicketModal = false"></div>
        
        <div class="inline-block align-bottom bg-white rounded-2xl text-right overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-lg sm:w-full">
          <div class="bg-white px-6 py-6">
            <div class="flex items-center justify-between mb-6">
              <h3 class="text-xl font-bold text-gray-900">تیکت جدید</h3>
              <button class="text-gray-400 hover:text-gray-600" @click="showNewTicketModal = false">
                <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                </svg>
              </button>
            </div>
            
            <form class="space-y-4" @submit.prevent="createTicket">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">عنوان</label>
                <input 
                  v-model="newTicket.title"
                  type="text"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent"
                  required
                >
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">دسته‌بندی</label>
                <select 
                  v-model="newTicket.category"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent"
                  required
                >
                  <option value="">انتخاب کنید</option>
                  <option value="technical">مشکلات فنی</option>
                  <option value="billing">مشکلات پرداخت</option>
                  <option value="product">سوالات محصول</option>
                  <option value="shipping">مشکلات ارسال</option>
                  <option value="other">سایر</option>
                </select>
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">اولویت</label>
                <select 
                  v-model="newTicket.priority"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent"
                  required
                >
                  <option value="low">کم</option>
                  <option value="medium">متوسط</option>
                  <option value="high">زیاد</option>
                  <option value="urgent">فوری</option>
                </select>
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">توضیحات</label>
                <textarea 
                  v-model="newTicket.description"
                  rows="3"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent"
                  required
                ></textarea>
              </div>
              
              <div class="flex justify-end space-x-3 space-x-reverse">
                <button 
                  type="button"
                  class="px-4 py-2 text-sm font-medium text-gray-700 bg-gray-100 border border-gray-300 rounded-lg hover:bg-gray-200"
                  @click="showNewTicketModal = false"
                >
                  انصراف
                </button>
                <button 
                  type="submit"
                  class="px-4 py-2 text-sm font-medium text-white bg-orange-500 border border-transparent rounded-lg hover:bg-orange-600"
                >
                  ایجاد تیکت
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'

interface TicketMessage {
  id: number
  sender: string
  content: string
  timestamp: Date
}

interface Ticket {
  id: number
  title: string
  description: string
  customerName: string
  customerEmail: string
  customerPhone: string
  category: string
  priority: string
  status: string
  assignedTo: string | null
  unreadMessages: number
  createdAt: Date
  messages: TicketMessage[]
}

// Reactive data
const showNewTicketModal = ref(false)
const activeTicketFilter = ref('all')
const selectedTicket = ref<Ticket | null>(null)

// Ticket stats
const ticketStats = ref({
  urgent: 5,
  pending: 12,
  inProgress: 8,
  resolved: 45
})

// Ticket filters
const ticketFilters = ref([
  { id: 'all', name: 'همه' },
  { id: 'urgent', name: 'فوری' },
  { id: 'pending', name: 'در انتظار' },
  { id: 'in_progress', name: 'در حال بررسی' },
  { id: 'resolved', name: 'حل شده' }
])

// Tickets data
const tickets = ref<Ticket[]>([
  {
    id: 1001,
    title: 'مشکل در پرداخت آنلاین',
    description: 'کاربر نمی‌تواند پرداخت خود را تکمیل کند',
    customerName: 'احمد محمدی',
    customerEmail: 'ahmad@example.com',
    customerPhone: '09123456789',
    category: 'billing',
    priority: 'high',
    status: 'pending',
    assignedTo: null,
    unreadMessages: 2,
    createdAt: new Date(Date.now() - 3600000),
    messages: [
      {
        id: 1,
        sender: 'احمد محمدی',
        content: 'سلام، نمی‌تونم پرداخت کنم',
        timestamp: new Date(Date.now() - 3600000)
      },
      {
        id: 2,
        sender: 'پشتیبانی',
        content: 'سلام، لطفاً جزئیات بیشتری ارائه دهید',
        timestamp: new Date(Date.now() - 1800000)
      }
    ]
  },
  {
    id: 1002,
    title: 'سوال در مورد ویژگی‌های محصول',
    description: 'کاربر می‌خواهد اطلاعات بیشتری در مورد محصول بداند',
    customerName: 'فاطمه احمدی',
    customerEmail: 'fateme@example.com',
    customerPhone: '09187654321',
    category: 'product',
    priority: 'medium',
    status: 'in_progress',
    assignedTo: 'علی رضایی',
    unreadMessages: 0,
    createdAt: new Date(Date.now() - 7200000),
    messages: [
      {
        id: 1,
        sender: 'فاطمه احمدی',
        content: 'سلام، می‌خوام بدونم این محصول چه ویژگی‌هایی داره',
        timestamp: new Date(Date.now() - 7200000)
      }
    ]
  },
  {
    id: 1003,
    title: 'مشکل در ارسال سفارش',
    description: 'سفارش کاربر ارسال نشده است',
    customerName: 'علی رضایی',
    customerEmail: 'ali@example.com',
    customerPhone: '09111222333',
    category: 'shipping',
    priority: 'urgent',
    status: 'pending',
    assignedTo: null,
    unreadMessages: 1,
    createdAt: new Date(Date.now() - 1800000),
    messages: [
      {
        id: 1,
        sender: 'علی رضایی',
        content: 'سفارشم هنوز ارسال نشده، لطفاً بررسی کنید',
        timestamp: new Date(Date.now() - 1800000)
      }
    ]
  }
])

// New ticket form
const newTicket = ref({
  title: '',
  category: '',
  priority: 'medium',
  description: ''
})

// Computed properties
const filteredTickets = computed(() => {
  if (activeTicketFilter.value === 'all') {
    return tickets.value
  }
  return tickets.value.filter(ticket => ticket.status === activeTicketFilter.value)
})

// Methods
const selectTicket = (ticket: Ticket) => {
  selectedTicket.value = ticket
}

const assignTicket = (ticket: Ticket) => {
  ticket.assignedTo = 'شما'
}

const updateTicketStatus = (ticket: Ticket, status: string) => {
  ticket.status = status
}

const createTicket = () => {
  const ticket: Ticket = {
    id: Date.now(),
    title: newTicket.value.title,
    description: newTicket.value.description,
    customerName: 'کاربر جدید',
    customerEmail: 'user@example.com',
    customerPhone: '09100000000',
    category: newTicket.value.category,
    priority: newTicket.value.priority,
    status: 'pending',
    assignedTo: null,
    unreadMessages: 0,
    createdAt: new Date(),
    messages: []
  }
  
  tickets.value.unshift(ticket)
  showNewTicketModal.value = false
  newTicket.value = { title: '', category: '', priority: 'medium', description: '' }
}

const getPriorityColor = (priority: string): string => {
  const colors = {
    low: 'bg-gray-100 text-gray-800',
    medium: 'bg-yellow-100 text-yellow-800',
    high: 'bg-orange-100 text-orange-800',
    urgent: 'bg-red-100 text-red-800'
  }
  return colors[priority as keyof typeof colors] || colors.low
}

const getPriorityText = (priority: string): string => {
  const texts = {
    low: 'کم',
    medium: 'متوسط',
    high: 'زیاد',
    urgent: 'فوری'
  }
  return texts[priority as keyof typeof texts] || 'نامشخص'
}

const getStatusColor = (status: string): string => {
  const colors = {
    pending: 'bg-yellow-100 text-yellow-800',
    in_progress: 'bg-blue-100 text-blue-800',
    resolved: 'bg-green-100 text-green-800'
  }
  return colors[status as keyof typeof colors] || colors.pending
}

const getStatusText = (status: string): string => {
  const texts = {
    pending: 'در انتظار',
    in_progress: 'در حال بررسی',
    resolved: 'حل شده'
  }
  return texts[status as keyof typeof texts] || 'نامشخص'
}

const formatTime = (timestamp: Date): string => {
  return timestamp.toLocaleTimeString('fa-IR', { 
    hour: '2-digit', 
    minute: '2-digit' 
  })
}

const formatDate = (timestamp: Date): string => {
  return timestamp.toLocaleDateString('fa-IR')
}
</script>

<style scoped>
/* Custom styles for ticket manager */
</style>

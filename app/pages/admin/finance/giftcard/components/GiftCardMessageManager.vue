<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200">
    <!-- Header -->
    <div class="px-6 py-4 border-b border-gray-200">
      <div class="flex justify-between items-center">
        <div>
          <h3 class="text-lg font-semibold text-gray-900">مدیریت پیام‌های گیفت کارت</h3>
          <p class="text-sm text-gray-600 mt-1">مدیریت پیام‌های پیش‌فرض و سفارشی برای گیفت کارت‌ها</p>
        </div>
        <button 
          @click="showCreateModal = true"
          class="px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
        >
          افزودن پیام جدید
        </button>
      </div>
    </div>

    <!-- تب‌های مدیریت -->
    <div class="border-b border-gray-200">
      <nav class="-mb-px flex space-x-8 space-x-reverse px-6">
        <button
          @click="activeTab = 'default'"
          :class="{
            'border-blue-500 text-blue-600': activeTab === 'default',
            'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300': activeTab !== 'default'
          }"
          class="whitespace-nowrap py-2 px-1 border-b-2 font-medium text-sm"
        >
          پیام‌های پیش‌فرض
        </button>
        <button
          @click="activeTab = 'custom'"
          :class="{
            'border-blue-500 text-blue-600': activeTab === 'custom',
            'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300': activeTab !== 'custom'
          }"
          class="whitespace-nowrap py-2 px-1 border-b-2 font-medium text-sm"
        >
          پیام‌های سفارشی
        </button>
        <button
          @click="activeTab = 'templates'"
          :class="{
            'border-blue-500 text-blue-600': activeTab === 'templates',
            'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300': activeTab !== 'templates'
          }"
          class="whitespace-nowrap py-2 px-1 border-b-2 font-medium text-sm"
        >
          قالب‌های پیام
        </button>
        <button
          @click="activeTab = 'settings'"
          :class="{
            'border-blue-500 text-blue-600': activeTab === 'settings',
            'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300': activeTab !== 'settings'
          }"
          class="whitespace-nowrap py-2 px-1 border-b-2 font-medium text-sm"
        >
          تنظیمات
        </button>
      </nav>
    </div>

    <!-- محتوای تب‌ها -->
    <div class="p-6">
      <!-- تب پیام‌های پیش‌فرض -->
      <div v-if="activeTab === 'default'" class="space-y-6">
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <div 
            v-for="message in defaultMessages" 
            :key="message.id"
            class="border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow"
          >
            <div class="flex justify-between items-start mb-3">
              <div class="flex items-center space-x-2 space-x-reverse">
                <span :class="getCategoryClasses(message.category)" class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium">
                  {{ getCategoryText(message.category) }}
                </span>
                <span v-if="message.isActive" class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800">
                  فعال
                </span>
              </div>
              <div class="flex items-center space-x-1 space-x-reverse">
                <button 
                  @click="editMessage(message)"
                  class="text-blue-600 hover:text-blue-900 p-1 rounded hover:bg-blue-50"
                  title="ویرایش"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                  </svg>
                </button>
                <button 
                  @click="toggleMessageStatus(message)"
                  :class="message.isActive ? 'text-yellow-600 hover:text-yellow-900' : 'text-green-600 hover:text-green-900'"
                  class="p-1 rounded hover:bg-gray-50"
                  :title="message.isActive ? 'غیرفعال کردن' : 'فعال کردن'"
                >
                  <svg v-if="message.isActive" class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728L5.636 5.636m12.728 12.728L18.364 5.636M5.636 18.364l12.728-12.728" />
                  </svg>
                  <svg v-else class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                  </svg>
                </button>
              </div>
            </div>
            
            <h4 class="font-medium text-gray-900 mb-2">{{ message.title }}</h4>
            <p class="text-sm text-gray-600 mb-3 line-clamp-3">{{ message.content }}</p>
            
            <div class="flex justify-between items-center text-xs text-gray-500">
              <span>{{ message.usageCount }} بار استفاده شده</span>
              <span>{{ formatDate(message.updatedAt) }}</span>
            </div>
          </div>
        </div>
      </div>

      <!-- تب پیام‌های سفارشی -->
      <div v-if="activeTab === 'custom'" class="space-y-6">
        <div class="flex justify-between items-center">
          <div class="flex items-center space-x-4 space-x-reverse">
            <select v-model="customMessageFilter" class="px-3 py-2 text-sm border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
              <option value="">همه دسته‌ها</option>
              <option value="birthday">تولد</option>
              <option value="wedding">عروسی</option>
              <option value="newyear">سال نو</option>
              <option value="corporate">شرکتی</option>
              <option value="general">عمومی</option>
            </select>
            <select v-model="customMessageStatus" class="px-3 py-2 text-sm border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
              <option value="">همه وضعیت‌ها</option>
              <option value="active">فعال</option>
              <option value="inactive">غیرفعال</option>
              <option value="pending">در انتظار تأیید</option>
            </select>
          </div>
          <div class="relative">
            <input 
              v-model="customMessageSearch" 
              type="text" 
              placeholder="جستجو در پیام‌ها..."
              class="pl-10 pr-4 py-2 text-sm border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 w-64"
            />
            <svg class="absolute left-3 top-2.5 h-4 w-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
          </div>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <div 
            v-for="message in filteredCustomMessages" 
            :key="message.id"
            class="border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow"
          >
            <div class="flex justify-between items-start mb-3">
              <div class="flex items-center space-x-2 space-x-reverse">
                <span :class="getCategoryClasses(message.category)" class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium">
                  {{ getCategoryText(message.category) }}
                </span>
                <span :class="getStatusClasses(message.status)" class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium">
                  {{ getStatusText(message.status) }}
                </span>
              </div>
              <div class="flex items-center space-x-1 space-x-reverse">
                <button 
                  @click="editMessage(message)"
                  class="text-blue-600 hover:text-blue-900 p-1 rounded hover:bg-blue-50"
                  title="ویرایش"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                  </svg>
                </button>
                <button 
                  @click="approveMessage(message)"
                  v-if="message.status === 'pending'"
                  class="text-green-600 hover:text-green-900 p-1 rounded hover:bg-green-50"
                  title="تأیید"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                  </svg>
                </button>
                <button 
                  @click="deleteMessage(message)"
                  class="text-red-600 hover:text-red-900 p-1 rounded hover:bg-red-50"
                  title="حذف"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                  </svg>
                </button>
              </div>
            </div>
            
            <h4 class="font-medium text-gray-900 mb-2">{{ message.title }}</h4>
            <p class="text-sm text-gray-600 mb-3 line-clamp-3">{{ message.content }}</p>
            
            <div class="flex justify-between items-center text-xs text-gray-500">
              <span>{{ message.usageCount }} بار استفاده شده</span>
              <span>{{ formatDate(message.createdAt) }}</span>
            </div>
            
            <div v-if="message.createdBy" class="mt-2 text-xs text-gray-400">
              ایجاد شده توسط: {{ message.createdBy }}
            </div>
          </div>
        </div>
      </div>

      <!-- تب قالب‌های پیام -->
      <div v-if="activeTab === 'templates'" class="space-y-6">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div 
            v-for="template in messageTemplates" 
            :key="template.id"
            class="border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow"
          >
            <div class="flex justify-between items-start mb-3">
              <h4 class="font-medium text-gray-900">{{ template.name }}</h4>
              <button 
                @click="useTemplate(template)"
                class="text-blue-600 hover:text-blue-900 text-sm font-medium"
              >
                استفاده
              </button>
            </div>
            
            <p class="text-sm text-gray-600 mb-3">{{ template.description }}</p>
            
            <div class="bg-gray-50 rounded p-3 text-sm text-gray-700 font-mono">
              {{ template.template }}
            </div>
            
            <div class="mt-3 text-xs text-gray-500">
              متغیرهای موجود: {{ template.variables.join(', ') }}
            </div>
          </div>
        </div>
      </div>

      <!-- تب تنظیمات -->
      <div v-if="activeTab === 'settings'" class="space-y-6">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <!-- محدودیت تعداد کاراکتر -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h4 class="font-medium text-gray-900 mb-3">محدودیت تعداد کاراکتر</h4>
            <div class="space-y-3">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">حداکثر کاراکتر برای پیام شخصی</label>
                <input 
                  v-model.number="settings.maxPersonalMessageLength" 
                  type="number" 
                  min="50" 
                  max="1000"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">حداکثر کاراکتر برای پیام پیش‌فرض</label>
                <input 
                  v-model.number="settings.maxDefaultMessageLength" 
                  type="number" 
                  min="50" 
                  max="500"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                />
              </div>
            </div>
          </div>

          <!-- فیلتر محتوا -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h4 class="font-medium text-gray-900 mb-3">فیلتر محتوا</h4>
            <div class="space-y-3">
              <div class="flex items-center">
                <input 
                  type="checkbox" 
                  v-model="settings.enableContentFilter" 
                  id="enableContentFilter" 
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label for="enableContentFilter" class="mr-2 block text-sm text-gray-900">فعال‌سازی فیلتر محتوای نامناسب</label>
              </div>
              <div class="flex items-center">
                <input 
                  type="checkbox" 
                  v-model="settings.requireApproval" 
                  id="requireApproval" 
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label for="requireApproval" class="mr-2 block text-sm text-gray-900">نیاز به تأیید پیام‌های سفارشی</label>
              </div>
              <div class="flex items-center">
                <input 
                  type="checkbox" 
                  v-model="settings.autoTranslate" 
                  id="autoTranslate" 
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label for="autoTranslate" class="mr-2 block text-sm text-gray-900">ترجمه خودکار پیام‌ها</label>
              </div>
            </div>
          </div>

          <!-- کلمات ممنوعه -->
          <div class="border border-gray-200 rounded-lg p-6 md:col-span-2">
            <h4 class="font-medium text-gray-900 mb-3">کلمات ممنوعه</h4>
            <div class="space-y-3">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">کلمات ممنوعه (هر کلمه در یک خط)</label>
                <textarea 
                  v-model="settings.bannedWords" 
                  rows="4"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                  placeholder="کلمه1&#10;کلمه2&#10;کلمه3"
                ></textarea>
              </div>
              <div class="flex items-center space-x-4 space-x-reverse">
                <button 
                  @click="addBannedWord"
                  class="px-3 py-1 bg-red-600 text-white text-sm rounded hover:bg-red-700"
                >
                  افزودن کلمه
                </button>
                <button 
                  @click="clearBannedWords"
                  class="px-3 py-1 bg-gray-600 text-white text-sm rounded hover:bg-gray-700"
                >
                  پاک کردن همه
                </button>
              </div>
            </div>
          </div>
        </div>

        <!-- دکمه ذخیره تنظیمات -->
        <div class="flex justify-end">
          <button 
            @click="saveSettings"
            class="px-6 py-2 bg-green-600 text-white text-sm font-medium rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2"
          >
            ذخیره تنظیمات
          </button>
        </div>
      </div>
    </div>

    <!-- مودال ایجاد/ویرایش پیام -->
    <GiftCardMessageCreateModal 
      v-if="showCreateModal"
      :message="editingMessage"
      @close="closeModal"
      @created="handleMessageCreated"
      @updated="handleMessageUpdated"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'

// کامپوننت‌های مورد نیاز
import GiftCardMessageCreateModal from './GiftCardMessageCreateModal.vue'

// Reactive data
const activeTab = ref('default')
const showCreateModal = ref(false)
const editingMessage = ref(null)
const customMessageFilter = ref('')
const customMessageStatus = ref('')
const customMessageSearch = ref('')

// تنظیمات
const settings = ref({
  maxPersonalMessageLength: 500,
  maxDefaultMessageLength: 200,
  enableContentFilter: true,
  requireApproval: true,
  autoTranslate: false,
  bannedWords: ''
})

// پیام‌های پیش‌فرض
const defaultMessages = ref([
  {
    id: 1,
    title: 'پیام تولد',
    content: 'تولدت مبارک! آرزوی شادی و موفقیت برای شما دارم.',
    category: 'birthday',
    isActive: true,
    usageCount: 156,
    updatedAt: new Date(Date.now() - 86400000)
  },
  {
    id: 2,
    title: 'پیام عروسی',
    content: 'عروسی مبارک! با آرزوی زندگی سرشار از عشق و شادی.',
    category: 'wedding',
    isActive: true,
    usageCount: 89,
    updatedAt: new Date(Date.now() - 172800000)
  },
  {
    id: 3,
    title: 'پیام سال نو',
    content: 'سال نو مبارک! با آرزوی سالی پر از موفقیت و شادی.',
    category: 'newyear',
    isActive: true,
    usageCount: 234,
    updatedAt: new Date(Date.now() - 259200000)
  },
  {
    id: 4,
    title: 'پیام شرکتی',
    content: 'با تشکر از همکاری شما. موفقیت‌های بیشتری در پیش رو دارید.',
    category: 'corporate',
    isActive: false,
    usageCount: 45,
    updatedAt: new Date(Date.now() - 345600000)
  }
])

// پیام‌های سفارشی
const customMessages = ref([
  {
    id: 1,
    title: 'پیام شخصی تولد',
    content: 'در روز تولدت، آرزوی شادی و موفقیت دارم. امیدوارم سالی پر از لحظات زیبا داشته باشی.',
    category: 'birthday',
    status: 'active',
    usageCount: 23,
    createdAt: new Date(Date.now() - 86400000),
    createdBy: 'احمد محمدی'
  },
  {
    id: 2,
    title: 'پیام تشکر',
    content: 'از لطف و محبت شما متشکرم. امیدوارم این هدیه کوچک شادی‌بخش باشد.',
    category: 'general',
    status: 'pending',
    usageCount: 0,
    createdAt: new Date(Date.now() - 172800000),
    createdBy: 'فاطمه احمدی'
  }
])

// قالب‌های پیام
const messageTemplates = ref([
  {
    id: 1,
    name: 'قالب تولد',
    description: 'قالب مناسب برای مناسبت‌های تولد',
    template: 'تولدت مبارک {نام}! آرزوی {سال} سال زندگی سرشار از {موفقیت} برای شما دارم.',
    variables: ['نام', 'سال', 'موفقیت']
  },
  {
    id: 2,
    name: 'قالب تشکر',
    description: 'قالب مناسب برای ابراز تشکر',
    template: 'از {لطف} شما متشکرم. این {هدیه} کوچک نشان‌دهنده {احترام} من به شماست.',
    variables: ['لطف', 'هدیه', 'احترام']
  },
  {
    id: 3,
    name: 'قالب عروسی',
    description: 'قالب مناسب برای مراسم عروسی',
    template: 'عروسی مبارک {نام‌عروس} و {نام‌داماد}! با آرزوی زندگی سرشار از {عشق} و {شادی}.',
    variables: ['نام‌عروس', 'نام‌داماد', 'عشق', 'شادی']
  }
])

// Computed properties
const filteredCustomMessages = computed(() => {
  let filtered = customMessages.value

  if (customMessageFilter.value) {
    filtered = filtered.filter(message => message.category === customMessageFilter.value)
  }

  if (customMessageStatus.value) {
    filtered = filtered.filter(message => message.status === customMessageStatus.value)
  }

  if (customMessageSearch.value) {
    const query = customMessageSearch.value.toLowerCase()
    filtered = filtered.filter(message => 
      message.title.toLowerCase().includes(query) ||
      message.content.toLowerCase().includes(query) ||
      (message.createdBy || '').toLowerCase().includes(query)
    )
  }

  return filtered
})

// Methods
const formatDate = (date: Date) => {
  return date.toLocaleDateString('fa-IR', {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  })
}

const getCategoryText = (category: string) => {
  const categoryMap = {
    birthday: 'تولد',
    wedding: 'عروسی',
    newyear: 'سال نو',
    corporate: 'شرکتی',
    general: 'عمومی'
  }
  return categoryMap[category] || category
}

const getCategoryClasses = (category: string) => {
  const classesMap = {
    birthday: 'bg-pink-100 text-pink-800',
    wedding: 'bg-purple-100 text-purple-800',
    newyear: 'bg-green-100 text-green-800',
    corporate: 'bg-blue-100 text-blue-800',
    general: 'bg-gray-100 text-gray-800'
  }
  return classesMap[category] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status: string) => {
  const statusMap = {
    active: 'فعال',
    inactive: 'غیرفعال',
    pending: 'در انتظار تأیید'
  }
  return statusMap[status] || status
}

const getStatusClasses = (status: string) => {
  const classesMap = {
    active: 'bg-green-100 text-green-800',
    inactive: 'bg-gray-100 text-gray-800',
    pending: 'bg-yellow-100 text-yellow-800'
  }
  return classesMap[status] || 'bg-gray-100 text-gray-800'
}

const editMessage = (message: any) => {
  editingMessage.value = message
  showCreateModal.value = true
}

const toggleMessageStatus = async (message: any) => {
  try {
    message.isActive = !message.isActive
    // در نسخه واقعی: ارسال به API
    console.log('Toggling message status:', message.id, message.isActive)
    alert(`پیام ${message.isActive ? 'فعال' : 'غیرفعال'} شد`)
  } catch (error) {
    console.error('خطا در تغییر وضعیت پیام:', error)
    alert('خطا در تغییر وضعیت پیام')
  }
}

const approveMessage = async (message: any) => {
  if (confirm(`آیا می‌خواهید پیام "${message.title}" تأیید شود؟`)) {
    try {
      message.status = 'active'
      // در نسخه واقعی: ارسال به API
      console.log('Approving message:', message.id)
      alert('پیام تأیید شد')
    } catch (error) {
      console.error('خطا در تأیید پیام:', error)
      alert('خطا در تأیید پیام')
    }
  }
}

const deleteMessage = async (message: any) => {
  if (confirm(`آیا از حذف پیام "${message.title}" مطمئن هستید؟`)) {
    try {
      const index = customMessages.value.findIndex(m => m.id === message.id)
      if (index > -1) {
        customMessages.value.splice(index, 1)
      }
      // در نسخه واقعی: ارسال به API
      console.log('Deleting message:', message.id)
      alert('پیام حذف شد')
    } catch (error) {
      console.error('خطا در حذف پیام:', error)
      alert('خطا در حذف پیام')
    }
  }
}

const useTemplate = (template: any) => {
  editingMessage.value = {
    title: '',
    content: template.template,
    category: 'general',
    isTemplate: true,
    templateId: template.id
  }
  showCreateModal.value = true
}

const addBannedWord = () => {
  const word = prompt('کلمه ممنوعه جدید را وارد کنید:')
  if (word && word.trim()) {
    const words = settings.value.bannedWords.split('\n').filter(w => w.trim())
    if (!words.includes(word.trim())) {
      words.push(word.trim())
      settings.value.bannedWords = words.join('\n')
    }
  }
}

const clearBannedWords = () => {
  if (confirm('آیا از پاک کردن همه کلمات ممنوعه مطمئن هستید؟')) {
    settings.value.bannedWords = ''
  }
}

const saveSettings = async () => {
  try {
    // در نسخه واقعی: ارسال به API
    console.log('Saving settings:', settings.value)
    alert('تنظیمات با موفقیت ذخیره شد')
  } catch (error) {
    console.error('خطا در ذخیره تنظیمات:', error)
    alert('خطا در ذخیره تنظیمات')
  }
}

const closeModal = () => {
  showCreateModal.value = false
  editingMessage.value = null
}

const handleMessageCreated = (message: any) => {
  customMessages.value.unshift(message)
  closeModal()
}

const handleMessageUpdated = (message: any) => {
  const index = customMessages.value.findIndex(m => m.id === message.id)
  if (index > -1) {
    customMessages.value[index] = message
  }
  closeModal()
}

// Lifecycle
onMounted(() => {
  // بارگذاری تنظیمات از API
  console.log('Message manager mounted')
})
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
.line-clamp-3 {
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style> 

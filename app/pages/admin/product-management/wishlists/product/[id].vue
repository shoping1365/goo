<template>
 
    <!-- Page Header -->
    <div class="bg-white shadow-sm border-b border-gray-200 mb-4">
      <div class="w-full px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center py-3">
          <div class="flex items-center">
            <button 
              class="ml-4 p-2 text-gray-600 hover:text-gray-800 hover:bg-gray-100 rounded-lg transition-colors"
              title="بازگشت"
              @click="$router.back()"
            >
              ←
            </button>
            <div>
              <h1 class="text-lg font-bold text-gray-900">مدیریت کاربران علاقمند</h1>
              <p class="text-xs text-gray-600 mt-1">{{ product?.name }}</p>
            </div>
          </div>
          <div class="flex space-x-2 space-x-reverse">
            <button 
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-blue-400 to-blue-600 hover:from-blue-500 hover:to-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105"
              @click="showPriceModal = true"
            >
              تغییر قیمت
            </button>
            <button 
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-green-400 to-green-600 hover:from-green-500 hover:to-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105"
              @click="showSmsModal = true"
            >
              ارسال پیام
            </button>
          </div>
        </div>
      </div>
    </div>

    <div class="w-full px-4 py-4">

      <!-- Product Summary Card -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6 mb-6">
        <div class="flex items-center">
          <img 
            :src="product?.image" 
            :alt="product?.name"
            class="w-24 h-24 rounded-lg object-cover border border-gray-200"
          />
          <div class="mr-6 flex-1">
            <h2 class="text-lg font-semibold text-gray-900">{{ product?.name }}</h2>
            <div class="grid grid-cols-2 md:grid-cols-4 gap-6 mt-4">
              <div class="bg-blue-50 rounded-lg p-3">
                <div class="text-xs text-blue-600 font-medium">قیمت فعلی</div>
                <div class="text-lg font-bold text-blue-900">{{ formatPrice(product?.price) }}</div>
                <div v-if="product?.discount" class="text-xs text-red-600">{{ product.discount }}% تخفیف</div>
              </div>
              <div class="bg-green-50 rounded-lg p-3">
                <div class="text-xs text-green-600 font-medium">علاقمندان</div>
                <div class="text-lg font-bold text-green-900">{{ wishlistUsers.length }} کاربر</div>
              </div>
              <div class="bg-purple-50 rounded-lg p-3">
                <div class="text-xs text-purple-600 font-medium">وضعیت موجودی</div>
                <div class="text-sm font-bold text-purple-900">
                  {{ product?.in_stock ? 'موجود' : 'ناموجود' }}
                </div>
              </div>
              <div class="bg-orange-50 rounded-lg p-3">
                <div class="text-xs text-orange-600 font-medium">دسته‌بندی</div>
                <div class="text-sm font-bold text-orange-900">{{ product?.category }}</div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Users Management -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200">
        <!-- Filters -->
        <div class="p-6 border-b border-gray-200">
          <div class="flex items-center justify-between">
            <h3 class="text-lg font-semibold text-gray-900">لیست کاربران علاقمند</h3>
            <div class="flex items-center space-x-2 space-x-reverse">
              <div class="relative">
                <input 
                  v-model="searchQuery"
                  type="text"
                  placeholder="جستجو کاربر..."
                  class="w-64 border border-gray-300 rounded-lg px-3 py-2 text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                />
              </div>
              <select 
                v-model="sortBy"
                class="border border-gray-300 rounded-lg px-3 py-2 text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="date_desc">جدیدترین</option>
                <option value="date_asc">قدیمی‌ترین</option>
                <option value="name_asc">نام (الف-ی)</option>
                <option value="name_desc">نام (ی-الف)</option>
              </select>
            </div>
          </div>
        </div>

        <!-- Users Table -->
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-4 py-3 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">
                  <input 
                    type="checkbox" 
                    :checked="isAllSelected"
                    class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                    @change="toggleSelectAll"
                  />
                </th>
                <th class="px-4 py-3 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">کاربر</th>
                <th class="px-4 py-3 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">تاریخ علاقمندی</th>
                <th class="px-4 py-3 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">مدت زمان</th>
                <th class="px-4 py-3 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">عملیات</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="userWishlist in paginatedUsers" :key="userWishlist.id" class="hover:bg-gray-50">
                <td class="px-4 py-4 whitespace-nowrap text-center">
                  <input 
                    v-model="selectedUsers" 
                    type="checkbox"
                    :value="userWishlist.id"
                    class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                  />
                </td>
                <td class="px-4 py-4 whitespace-nowrap">
                  <div class="flex items-center">
                    <div class="">
                      <div class="h-10 w-10 rounded-full bg-blue-100 flex items-center justify-center">
                        <span class="text-sm font-medium text-blue-700">
                          {{ userWishlist.user.name.charAt(0) }}
                        </span>
                      </div>
                    </div>
                    <div class="mr-4">
                      <div class="text-sm font-medium text-gray-900">{{ userWishlist.user.name }}</div>
                      <div class="text-sm text-gray-500">{{ userWishlist.user.email }}</div>
                      <div v-if="userWishlist.user.phone" class="text-xs text-gray-400">{{ userWishlist.user.phone }}</div>
                    </div>
                  </div>
                </td>
                <td class="px-4 py-4 whitespace-nowrap text-center">
                  <div class="text-sm text-gray-900">{{ formatDate(userWishlist.created_at) }}</div>
                  <div class="text-xs text-gray-500">{{ formatTime(userWishlist.created_at) }}</div>
                </td>
                <td class="px-4 py-4 whitespace-nowrap text-center">
                  <span class="text-sm text-gray-600">{{ getRelativeTime(userWishlist.created_at) }}</span>
                </td>
                <td class="px-4 py-4 whitespace-nowrap text-center">
                  <div class="flex items-center justify-center space-x-1 space-x-reverse">
                    <button 
                      class="p-2 text-green-600 hover:bg-green-100 rounded-lg transition-colors"
                      title="ارسال پیام شخصی"
                      @click="sendPersonalSms(userWishlist.user)"
                    >
                      💬
                    </button>
                    <button 
                      class="p-2 text-blue-600 hover:bg-blue-100 rounded-lg transition-colors"
                      title="مشاهده پروفایل"
                      @click="viewUserProfile(userWishlist.user.id)"
                    >
                      👤
                    </button>
                    <button 
                      class="p-2 text-red-600 hover:bg-red-100 rounded-lg transition-colors"
                      title="حذف از لیست"
                      @click="removeFromWishlist(userWishlist.id)"
                    >
                      🗑️
                    </button>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <!-- Pagination -->
        <div class="bg-gray-50 px-4 py-3 flex items-center justify-between border-t border-gray-200">
          <div class="flex items-center">
            <p class="text-sm text-gray-700">
              نمایش {{ paginationInfo.start }} تا {{ paginationInfo.end }} از {{ paginationInfo.total }} کاربر
            </p>
          </div>
          <div class="flex items-center space-x-1 space-x-reverse">
            <button 
              :disabled="currentPage === 1"
              class="px-3 py-1 border border-gray-300 rounded-md text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
              @click="currentPage = Math.max(1, currentPage - 1)"
            >
              قبلی
            </button>
            <template v-for="page in visiblePages" :key="page">
              <button 
                v-if="page !== '...'"
                :class="[
                  'px-3 py-1 border text-sm font-medium rounded-md',
                  currentPage === page 
                    ? 'border-blue-500 bg-blue-50 text-blue-600' 
                    : 'border-gray-300 bg-white text-gray-700 hover:bg-gray-50'
                ]"
                @click="currentPage = Number(page)"
              >
                {{ page }}
              </button>
              <span v-else class="px-3 py-1 text-gray-500">...</span>
            </template>
            <button 
              :disabled="currentPage === totalPages"
              class="px-3 py-1 border border-gray-300 rounded-md text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
              @click="currentPage = Math.min(totalPages, currentPage + 1)"
            >
              بعدی
            </button>
          </div>
        </div>
      </div>

      <!-- Bulk Actions -->
      <div v-if="selectedUsers.length > 0" class="fixed bottom-6 right-6 bg-white rounded-lg shadow-lg border border-gray-200 p-6">
        <div class="flex items-center space-x-3 space-x-reverse">
          <span class="text-sm font-medium text-gray-700">{{ selectedUsers.length }} کاربر انتخاب شده</span>
          <button 
            class="px-3 py-1.5 bg-green-600 text-white text-sm font-medium rounded-lg hover:bg-green-700 transition-colors"
            @click="sendBulkSms"
          >
            ارسال پیام گروهی
          </button>
          <button 
            class="px-3 py-1.5 bg-red-600 text-white text-sm font-medium rounded-lg hover:bg-red-700 transition-colors"
            @click="removeBulkFromWishlist"
          >
            حذف گروهی
          </button>
          <button 
            class="px-3 py-1.5 bg-gray-300 text-gray-700 text-sm font-medium rounded-lg hover:bg-gray-400 transition-colors"
            @click="selectedUsers = []"
          >
            لغو انتخاب
          </button>
        </div>
      </div>
    </div>

    <!-- Price Change Modal -->
    <div v-if="showPriceModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg shadow-xl p-6 w-full max-w-md mx-4">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">تغییر قیمت محصول</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">قیمت جدید (تومان)</label>
            <input 
              v-model="newPrice"
              type="number"
              class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              :placeholder="product?.price?.toString()"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">درصد تخفیف (اختیاری)</label>
            <input 
              v-model="newDiscount"
              type="number"
              min="0"
              max="100"
              class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="0"
            />
          </div>
          <div class="flex items-center">
            <input 
              v-model="notifyUsersOfPriceChange"
              type="checkbox"
              class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
            />
            <label class="mr-2 text-sm text-gray-700">اطلاع‌رسانی تغییر قیمت به کاربران</label>
          </div>
        </div>
        <div class="flex justify-end space-x-2 space-x-reverse mt-6">
          <button 
            class="px-4 py-2 bg-gray-300 text-gray-700 text-sm font-medium rounded-lg hover:bg-gray-400 transition-colors"
            @click="showPriceModal = false"
          >
            انصراف
          </button>
          <button 
            class="px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 transition-colors"
            @click="updatePrice"
          >
            ذخیره تغییرات
          </button>
        </div>
      </div>
    </div>

    <!-- SMS Modal -->
    <div v-if="showSmsModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg shadow-xl p-6 w-full max-w-lg mx-4">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">ارسال پیام به کاربران علاقمند</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع پیام</label>
            <select 
              v-model="smsType"
              class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="discount">اطلاع‌رسانی تخفیف</option>
              <option value="availability">اطلاع‌رسانی موجودی</option>
              <option value="price_drop">کاهش قیمت</option>
              <option value="custom">پیام سفارشی</option>
            </select>
          </div>
          <div v-if="smsType === 'discount'">
            <label class="block text-sm font-medium text-gray-700 mb-2">درصد تخفیف</label>
            <input 
              v-model="discountPercent"
              type="number"
              min="1"
              max="99"
              class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="20"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">متن پیام</label>
            <textarea 
              v-model="smsMessage"
              rows="4"
              class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              :placeholder="getSmsTemplate()"
            ></textarea>
            <div class="text-xs text-gray-500 mt-1">تعداد کاراکتر: {{ smsMessage.length }}/160</div>
          </div>
          <div class="bg-blue-50 rounded-lg p-3">
            <div class="text-sm text-blue-800">
              <strong>گیرندگان:</strong> {{ selectedUsers.length > 0 ? selectedUsers.length : wishlistUsers.length }} کاربر
            </div>
            <div class="text-xs text-blue-600 mt-1">
              هزینه تقریبی: {{ calculateSmsPrice() }} تومان
            </div>
          </div>
        </div>
        <div class="flex justify-end space-x-2 space-x-reverse mt-6">
          <button 
            class="px-4 py-2 bg-gray-300 text-gray-700 text-sm font-medium rounded-lg hover:bg-gray-400 transition-colors"
            @click="showSmsModal = false"
          >
            انصراف
          </button>
          <button 
            class="px-4 py-2 bg-green-600 text-white text-sm font-medium rounded-lg hover:bg-green-700 transition-colors"
            @click="sendSms"
          >
            ارسال پیام
          </button>
        </div>
      </div>
    </div>
  
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string }) => void
declare const useHead: (head: { title?: string }) => void
declare const navigateTo: (to: string) => Promise<void>
</script>

<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import { useRoute } from 'vue-router'

definePageMeta({
  layout: 'admin-main'
})

useHead({
  title: 'مدیریت کاربران علاقمند - پنل مدیریت'
})

// Types
interface Product {
  id: number
  name: string
  sku: string
  price: number
  discount?: number
  category: string
  image: string
  in_stock: boolean
  wishlist_count: number
}

interface User {
  id: number
  name: string
  email: string
  phone?: string
}

interface Wishlist {
  id: number
  product: Product
  user: User
  created_at: string
}

// Route params
const route = useRoute()
const productId = Number(route.params.id) || 1

// Reactive data
const searchQuery = ref('')
const sortBy = ref('date_desc')
const selectedUsers = ref<number[]>([])
const currentPage = ref(1)
const itemsPerPage = ref(15)

// Modals
const showPriceModal = ref(false)
const showSmsModal = ref(false)

// Price modal data
const newPrice = ref<number>()
const newDiscount = ref<number>(0)
const notifyUsersOfPriceChange = ref(true)

// SMS modal data
const smsType = ref('discount')
const discountPercent = ref<number>(20)
const smsMessage = ref('')

// Mock data (in real app, fetch from API)
const product = ref<Product>({
  id: productId,
  name: 'گوشی سامسونگ گلکسی A54',
  sku: 'SAM-A54-128',
  price: 8500000,
  discount: 15,
  category: 'الکترونیکی',
  image: 'https://via.placeholder.com/200x200?text=Phone',
  in_stock: true,
  wishlist_count: 245
})

const wishlistUsers = ref<Wishlist[]>([
  {
    id: 1,
    product: product.value,
    user: { id: 1001, name: 'علی احمدی', email: 'ali@example.com', phone: '09123456789' },
    created_at: '2024-01-15T10:30:00Z'
  },
  {
    id: 2,
    product: product.value,
    user: { id: 1006, name: 'فاطمه رضایی', email: 'fatemeh@example.com', phone: '09187654321' },
    created_at: '2024-01-16T14:20:00Z'
  },
  {
    id: 3,
    product: product.value,
    user: { id: 1007, name: 'حمید کریمی', email: 'hamid@example.com' },
    created_at: '2024-01-14T09:15:00Z'
  },
  {
    id: 4,
    product: product.value,
    user: { id: 1008, name: 'سارا احمدی', email: 'sara@example.com', phone: '09112345678' },
    created_at: '2024-01-13T16:45:00Z'
  },
  {
    id: 5,
    product: product.value,
    user: { id: 1009, name: 'محمد حسینی', email: 'mohammad@example.com', phone: '09198765432' },
    created_at: '2024-01-12T11:30:00Z'
  }
])

// Computed properties
const filteredUsers = computed(() => {
  let filtered = wishlistUsers.value

  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    filtered = filtered.filter(w => 
      w.user.name.toLowerCase().includes(query) ||
      w.user.email.toLowerCase().includes(query)
    )
  }

  // Sort
  filtered.sort((a, b) => {
    switch (sortBy.value) {
      case 'date_desc':
        return new Date(b.created_at).getTime() - new Date(a.created_at).getTime()
      case 'date_asc':
        return new Date(a.created_at).getTime() - new Date(b.created_at).getTime()
      case 'name_asc':
        return a.user.name.localeCompare(b.user.name, 'fa')
      case 'name_desc':
        return b.user.name.localeCompare(a.user.name, 'fa')
      default:
        return 0
    }
  })

  return filtered
})

const totalPages = computed(() => Math.ceil(filteredUsers.value.length / itemsPerPage.value))

const paginatedUsers = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage.value
  const end = start + itemsPerPage.value
  return filteredUsers.value.slice(start, end)
})

const paginationInfo = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage.value + 1
  const end = Math.min(start + itemsPerPage.value - 1, filteredUsers.value.length)
  return {
    start,
    end,
    total: filteredUsers.value.length
  }
})

const visiblePages = computed(() => {
  const pages = []
  const total = totalPages.value
  const current = currentPage.value
  
  if (total <= 7) {
    for (let i = 1; i <= total; i++) {
      pages.push(i)
    }
  } else {
    if (current <= 4) {
      for (let i = 1; i <= 5; i++) pages.push(i)
      pages.push('...')
      pages.push(total)
    } else if (current >= total - 3) {
      pages.push(1)
      pages.push('...')
      for (let i = total - 4; i <= total; i++) pages.push(i)
    } else {
      pages.push(1)
      pages.push('...')
      for (let i = current - 1; i <= current + 1; i++) pages.push(i)
      pages.push('...')
      pages.push(total)
    }
  }
  
  return pages
})

const isAllSelected = computed(() => {
  return paginatedUsers.value.length > 0 && 
         selectedUsers.value.length === paginatedUsers.value.length
})

// Methods
const toggleSelectAll = () => {
  if (isAllSelected.value) {
    selectedUsers.value = []
  } else {
    selectedUsers.value = paginatedUsers.value.map(w => w.id)
  }
}

const formatPrice = (price: number) => {
  return new Intl.NumberFormat('fa-IR').format(price) + ' تومان'
}

const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleDateString('fa-IR')
}

const formatTime = (dateString: string) => {
  return new Date(dateString).toLocaleTimeString('fa-IR', { 
    hour: '2-digit', 
    minute: '2-digit' 
  })
}

const getRelativeTime = (dateString: string) => {
  const date = new Date(dateString)
  const now = new Date()
  const diffInMs = now.getTime() - date.getTime()
  const diffInHours = Math.floor(diffInMs / (1000 * 60 * 60))
  const diffInDays = Math.floor(diffInHours / 24)
  
  if (diffInDays > 0) {
    return `${diffInDays} روز پیش`
  } else if (diffInHours > 0) {
    return `${diffInHours} ساعت پیش`
  } else {
    return 'همین الان'
  }
}

const updatePrice = () => {
  if (newPrice.value && newPrice.value > 0) {
    product.value.price = newPrice.value
    if (newDiscount.value) {
      product.value.discount = newDiscount.value
    }
    
    if (notifyUsersOfPriceChange.value) {
      // Send price change notification
      // console.log('اطلاع‌رسانی تغییر قیمت به کاربران...')
    }
    
    showPriceModal.value = false
    alert('قیمت محصول با موفقیت تغییر یافت!')
  }
}

const getSmsTemplate = () => {
  switch (smsType.value) {
    case 'discount':
      return `سلام، محصول مورد علاقه شما ${product.value.name} با ${discountPercent.value}% تخفیف در دسترس است!`
    case 'availability':
      return `سلام، محصول مورد علاقه شما ${product.value.name} دوباره موجود شد!`
    case 'price_drop':
      return `سلام، قیمت محصول مورد علاقه شما ${product.value.name} کاهش یافت!`
    default:
      return ''
  }
}

const calculateSmsPrice = () => {
  const recipientCount = selectedUsers.value.length > 0 ? selectedUsers.value.length : wishlistUsers.value.length
  const pricePerSms = 150 // تومان
  return recipientCount * pricePerSms
}

const sendSms = () => {
  const recipients = selectedUsers.value.length > 0 
    ? wishlistUsers.value.filter(w => selectedUsers.value.includes(w.id))
    : wishlistUsers.value
  
  // console.log('ارسال پیام به', recipients.length, 'کاربر:', smsMessage.value)
  
  showSmsModal.value = false
  selectedUsers.value = []
  alert(`پیام با موفقیت به ${recipients.length} کاربر ارسال شد!`)
}

const sendBulkSms = () => {
  smsMessage.value = getSmsTemplate()
  showSmsModal.value = true
}

const sendPersonalSms = (user: User) => {
  // console.log('ارسال پیام شخصی به', user.name)
  alert(`پیام شخصی به ${user.name} ارسال شد!`)
}

const viewUserProfile = (userId: number) => {
  navigateTo(`/admin/users/${userId}`)
}

const removeFromWishlist = (wishlistId: number) => {
  wishlistUsers.value = wishlistUsers.value.filter(w => w.id !== wishlistId)
  selectedUsers.value = selectedUsers.value.filter(id => id !== wishlistId)
}

const removeBulkFromWishlist = () => {
  wishlistUsers.value = wishlistUsers.value.filter(w => !selectedUsers.value.includes(w.id))
  selectedUsers.value = []
}

// Initialize SMS message when type changes
watch(smsType, () => {
  smsMessage.value = getSmsTemplate()
})

watch(discountPercent, () => {
  if (smsType.value === 'discount') {
    smsMessage.value = getSmsTemplate()
  }
})
</script> 
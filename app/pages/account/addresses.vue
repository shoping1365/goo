<template>
  <div class="min-h-screen bg-gray-50 py-8">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <!-- لایوت اصلی با سایدبار -->
      <div class="flex flex-col lg:flex-row gap-8 justify-center">
        <!-- محتوای اصلی -->
        <div class="flex-1 max-w-4xl">
          <!-- عنوان و دکمه افزودن آدرس جدید -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-6">
            <div class="flex items-center justify-between">
              <button 
                class="flex items-center space-x-2 text-gray-600 hover:text-gray-800 transition-colors"
                @click="openAddModal"
              >
                <span class="text-sm font-medium">افزودن آدرس جدید</span>
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
                </svg>
              </button>
              <h1 class="text-sm font-bold text-gray-900">آدرس‌ها</h1>
            </div>
          </div>

          <!-- حالت بارگذاری -->
          <div v-if="loading" class="space-y-4">
            <div v-for="i in 3" :key="i" class="bg-white border border-gray-200 rounded-lg p-6 animate-pulse">
              <div class="flex items-center justify-between mb-4">
                <div class="w-6 h-6 bg-gray-200 rounded"></div>
                <div class="w-5 h-5 bg-gray-200 rounded"></div>
              </div>
              <div class="space-y-2">
                <div class="h-4 bg-gray-200 rounded w-3/4"></div>
                <div class="h-4 bg-gray-200 rounded w-1/2"></div>
                <div class="h-4 bg-gray-200 rounded w-2/3"></div>
              </div>
            </div>
          </div>

          <!-- لیست آدرس‌های واقعی کاربر -->
          <div v-else-if="addresses.length > 0" class="space-y-4">
            <div 
              v-for="address in addresses" 
              :key="address.id"
              :class="[
                'bg-white rounded-lg p-6 relative transition-colors',
                address.is_default ? 'border-2 border-teal-500' : 'border border-gray-200'
              ]"
            >
              <!-- آیکون‌های بالای کارت -->
              <div class="flex items-center justify-between mb-4">
                <div class="flex items-center space-x-2 space-x-reverse">
                  <!-- دکمه تنظیم به عنوان پیش‌فرض -->
                  <button 
                    v-if="!address.is_default"
                    class="text-gray-400 hover:text-teal-600 transition-colors"
                    title="تنظیم به عنوان پیش‌فرض"
                    @click="setAsDefault(address.id)"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
                    </svg>
                  </button>
                  <!-- دکمه حذف -->
                  <button 
                    class="text-gray-400 hover:text-red-600 transition-colors"
                    title="حذف آدرس"
                    @click="confirmDelete(address)"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                    </svg>
                  </button>
                  <!-- دکمه ویرایش -->
                  <button 
                    class="text-gray-400 hover:text-blue-600 transition-colors"
                    title="ویرایش آدرس"
                    @click="openEditModal(address)"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
                    </svg>
                  </button>
                </div>
                <div class="flex items-center space-x-2 space-x-reverse">
                  <!-- برچسب پیش‌فرض -->
                  <span v-if="address.is_default" class="bg-teal-100 text-teal-800 text-xs px-2 py-1 rounded-full">
                    پیش‌فرض
                  </span>
                  <svg 
                    :class="[
                      'w-6 h-6',
                      address.is_default ? 'text-teal-600' : 'text-gray-600'
                    ]" 
                    fill="none" 
                    stroke="currentColor" 
                    viewBox="0 0 24 24"
                  >
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"></path>
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"></path>
                  </svg>
                </div>
              </div>
              
              <!-- جزئیات آدرس -->
              <div class="space-y-2 text-right">
                <p class="text-gray-900 font-medium">
                  {{ address.full_address || address.street || 'آدرس نامشخص' }}
                </p>
                <p v-if="address.postal_code" class="text-gray-600">
                  کد پستی: {{ address.postal_code }}
                </p>
                <p class="text-gray-600">
                  {{ address.recipient_name || 'نام دریافت‌کننده نامشخص' }}
                  <span v-if="address.recipient_mobile || address.phone">
                    | {{ address.recipient_mobile || address.phone }}
                  </span>
                </p>
                <p v-if="address.province || address.city" class="text-gray-500 text-sm">
                  {{ [address.province, address.city].filter(Boolean).join('، ') }}
                </p>
              </div>
            </div>
          </div>

          <!-- حالت خالی - عدم وجود آدرس -->
          <div v-else class="text-center py-12">
            <svg class="w-16 h-16 text-gray-300 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"></path>
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"></path>
            </svg>
            <h3 class="text-lg font-medium text-gray-900 mb-2">هنوز آدرسی ثبت نشده</h3>
            <p class="text-gray-500 mb-6">برای شروع خرید، ابتدا آدرس خود را اضافه کنید</p>
            <button 
              class="bg-teal-600 text-white px-6 py-3 rounded-lg hover:bg-teal-700 transition-colors"
              @click="openAddModal"
            >
              افزودن اولین آدرس
            </button>
          </div>
        </div>

        <!-- سایدبار -->
        <div class="lg:w-80 flex-shrink-0 lg:ml-8">
          <AccountSidebar />
        </div>
      </div>
    </div>

    <!-- مودال ویرایش/افزودن آدرس -->
    <AddressEditModal
      :is-open="isModalOpen"
      :address="selectedAddress"
      @close="closeModal"
      @saved="handleAddressSaved"
    />

    <!-- مودال تایید حذف -->
    <div v-if="showDeleteConfirm" class="fixed inset-0 flex items-center justify-center z-50 p-6">
      <div class="bg-white rounded-lg max-w-md w-full p-6 shadow-xl border-2 border-blue-200">
        <div class="flex items-center justify-between mb-4">
          <div class="flex-shrink-0">
            <svg class="w-6 h-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
            </svg>
          </div>
          <div class="text-right">
            <h3 class="text-lg font-medium text-gray-900">تایید حذف آدرس</h3>
          </div>
        </div>
        <div class="mb-6 text-right">
          <p class="text-sm text-gray-500">
            آیا مطمئن هستید که می‌خواهید این آدرس را حذف کنید؟ این عمل قابل بازگشت نیست.
          </p>
        </div>
        <div class="flex items-center justify-start">
          <button
            :disabled="deleting"
            class="px-4 py-2 text-sm font-medium text-white bg-red-600 border border-transparent rounded-md hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 disabled:opacity-50 disabled:cursor-not-allowed"
            @click="deleteAddress"
          >
            <span v-if="deleting" class="flex items-center">
              <svg class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              در حال حذف...
            </span>
            <span v-else>حذف آدرس</span>
          </button>
          <div class="w-4"></div>
          <button
            class="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-teal-500"
            @click="showDeleteConfirm = false"
          >
            انصراف
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import AccountSidebar from '@/components/account/AccountSidebar.vue'
import AddressEditModal from '@/components/account/AddressEditModal.vue'

// تنظیم عنوان صفحه
definePageMeta({
  layout: 'default',
  middleware: ['auth']
})

// تنظیم عنوان صفحه
useHead({
  title: 'آدرس‌ها - حساب کاربری'
})

// استفاده از composable آدرس‌ها
const { addresses, loading, fetchAddresses, deleteAddress: deleteAddressAPI, setDefaultAddress } = useAddresses()

// حالت‌های مودال
const isModalOpen = ref(false)
const selectedAddress = ref(null)
const showDeleteConfirm = ref(false)
const addressToDelete = ref(null)
const deleting = ref(false)

// باز کردن مودال افزودن آدرس جدید
const openAddModal = () => {
  selectedAddress.value = null
  isModalOpen.value = true
}

// باز کردن مودال ویرایش آدرس
const openEditModal = (address) => {
  selectedAddress.value = address
  isModalOpen.value = true
}

// بستن مودال
const closeModal = () => {
  isModalOpen.value = false
  selectedAddress.value = null
}

// تایید حذف آدرس
const confirmDelete = (address) => {
  addressToDelete.value = address
  showDeleteConfirm.value = true
}

// حذف آدرس
const deleteAddress = async () => {
  if (!addressToDelete.value) return
  
  try {
    deleting.value = true
    await deleteAddressAPI(addressToDelete.value.id)
    showDeleteConfirm.value = false
    addressToDelete.value = null
    // بارگذاری مجدد آدرس‌ها
    await fetchAddresses(true)
  } catch (error) {
    console.error('خطا در حذف آدرس:', error)
  } finally {
    deleting.value = false
  }
}

// تنظیم آدرس به عنوان پیش‌فرض
const setAsDefault = async (addressId) => {
  try {
    await setDefaultAddress(addressId)
    // بارگذاری مجدد آدرس‌ها
    await fetchAddresses(true)
  } catch (error) {
    console.error('خطا در تنظیم آدرس پیش‌فرض:', error)
  }
}

// مدیریت ذخیره آدرس
const handleAddressSaved = async () => {
  // بارگذاری مجدد آدرس‌ها
  await fetchAddresses(true)
}

// بارگذاری آدرس‌ها هنگام mount شدن کامپوننت
onMounted(async () => {
  try {
    await fetchAddresses()
  } catch (error) {
    console.error('خطا در بارگذاری آدرس‌ها:', error)
  }
})
</script> 
<template>
  <div class="checkout-root bg-[#eef2ff] min-h-[100vh] py-10" dir="rtl" style="font-family: 'IranYekan', 'Vazir', 'IRANSansX', 'Tahoma', sans-serif;">
    <div class="max-w-6xl mx-auto">
      <!-- آدرس و زمان ارسال -->
      <div class="bg-white rounded-2xl shadow px-4 py-4 mb-6">
        <h2 class="text-xl font-bold text-[#1a2341] mb-4">آدرس و زمان ارسال</h2>
      </div>

      <div class="flex flex-col lg:flex-row gap-8">
        <!-- فرم اطلاعات -->
        <div class="flex-1">
          <!-- اگر آدرس انتخاب شده وجود داشت، خلاصه آدرس را نمایش بده -->
          <AddressSummary v-if="summaryAddress" :address="summaryAddress" @edit="showAddressModal=true" class="mb-6" />

          <!-- فرم اطلاعات در صورتی که آدرسی انتخاب نشده باشد -->
          <div v-if="!summaryAddress" class="bg-white rounded-2xl shadow px-4 py-4 mb-6">
            <h2 class="text-xl font-bold text-[#1a2341] mb-4">اطلاعات شخصی</h2>
            <form @submit.prevent class="space-y-4">
              <!-- اطلاعات شخصی -->
              <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">نام و نام خانوادگی *</label>
                  <input 
                    v-model="formData.fullName" 
                    type="text" 
                    required
                    class="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-[#e60023] focus:border-transparent"
                    placeholder="نام و نام خانوادگی خود را وارد کنید"
                  >
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">شماره موبایل *</label>
                  <input 
                    v-model="formData.phone" 
                    type="tel" 
                    required
                    class="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-[#e60023] focus:border-transparent"
                    placeholder="09xxxxxxxxx"
                  >
                </div>
              </div>

              <!-- آدرس -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">آدرس کامل *</label>
                <textarea 
                  v-model="formData.address" 
                  required
                  rows="3"
                  class="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-[#e60023] focus:border-transparent"
                  placeholder="آدرس کامل خود را وارد کنید"
                ></textarea>
              </div>

              <!-- استان و شهر -->
              <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">استان *</label>
                  <select v-model="selectedProvinceId" required class="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-[#e60023] focus:border-transparent">
                    <option value="" disabled selected>انتخاب استان</option>
                    <option v-for="p in provinces" :key="p.id" :value="p.id">{{ p.name }}</option>
                  </select>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">شهر *</label>
                  <select v-model="formData.cityId" required class="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-[#e60023] focus:border-transparent">
                    <option value="" disabled selected>انتخاب شهر</option>
                    <option v-for="c in cities" :key="c.id" :value="c.id">{{ c.name }}</option>
                  </select>
                </div>
              </div>

              <!-- کد پستی -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">کد پستی</label>
                <input 
                  v-model="formData.postalCode" 
                  type="text" 
                  class="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-[#e60023] focus:border-transparent"
                  placeholder="کد پستی ۱۰ رقمی"
                >
              </div>

              <!-- توضیحات -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">توضیحات سفارش</label>
                <textarea 
                  v-model="formData.notes" 
                  rows="2"
                  class="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-[#e60023] focus:border-transparent"
                  placeholder="توضیحات اضافی (اختیاری)"
                ></textarea>
              </div>
            </form>
          </div>
          <!-- دکمه تغییر آدرس زمانی‌که آدرسی انتخاب نشده ولی لیست آدرس‌ها وجود دارد -->
          <button v-if="!summaryAddress && addresses.length>0" @click="showAddressModal=true" class="mb-6 text-sm text-[#e60023] font-bold hover:underline">
            انتخاب از آدرس‌های ذخیره‌شده
          </button>
          
          <!-- مودال آدرس‌ها -->
          <AddressModal :open="showAddressModal" @close="showAddressModal=false" @select="onAddressSelected" @addresses-changed="fetchAddresses" />

          <!-- لیست محصولات -->
          <div class="bg-white rounded-2xl shadow px-4 py-4 mb-6">
            <h2 class="text-lg font-bold text-[#1a2341] mb-3">محصولات سفارش</h2>
            <div class="flex gap-3 overflow-x-auto hide-scrollbar pb-2">
              <div v-for="item in cartItems" :key="item.id" class="min-w-[120px] bg-white border border-gray-200 rounded-lg p-2 flex flex-col items-center">
                <!-- تصویر محصول -->
                <div class="mb-2">
                  <img :src="item.image || '/default-product.svg'" :alt="item.name" class="w-16 h-16 object-cover rounded">
                </div>
                
                <!-- کنترل تعداد -->
                <div class="bg-gray-50 border border-gray-200 rounded p-1 flex items-center justify-between w-full">
                  <button 
                    v-if="item.quantity > 1"
                    @click="updateQuantity(item, item.quantity - 1)"
                    class="w-5 h-5 flex items-center justify-center rounded-full bg-red-500 text-white text-xs font-bold hover:bg-red-600"
                  >
                    -
                  </button>
                  <button 
                    v-else
                    @click="removeItem(item)"
                    class="w-5 h-5 flex items-center justify-center rounded-full bg-red-500 text-white text-xs font-bold hover:bg-red-600"
                  >
                    🗑️
                  </button>
                  
                  <div class="text-center">
                    <div class="text-red-500 font-bold text-sm">{{ item.quantity }}</div>
                  </div>
                  
                  <button 
                    @click="updateQuantity(item, item.quantity + 1)"
                    :disabled="item.quantity >= (item.stock_quantity || 999)"
                    class="w-5 h-5 flex items-center justify-center rounded-full bg-red-500 text-white text-xs font-bold hover:bg-red-600 disabled:opacity-50 disabled:cursor-not-allowed"
                  >
                    +
                  </button>
                </div>
              </div>
            </div>
          </div>
 
          <!-- روش پرداخت -->
          <div class="bg-white rounded-2xl shadow px-4 py-4">
            <h2 class="text-xl font-bold text-[#1a2341] mb-4">روش پرداخت</h2>
            <div class="space-y-3">
              <label class="flex items-center px-4 py-4 border border-gray-200 rounded-xl cursor-pointer hover:bg-gray-50">
                <input 
                  v-model="formData.paymentMethod" 
                  type="radio" 
                  value="online" 
                  class="mr-3 text-[#e60023] focus:ring-[#e60023]"
                >
                <div>
                  <div class="font-medium">پرداخت آنلاین</div>
                  <div class="text-sm text-gray-600">پرداخت امن از طریق درگاه‌های بانکی</div>
                </div>
              </label>
              
              <label class="flex items-center px-4 py-4 border border-gray-200 rounded-xl cursor-pointer hover:bg-gray-50">
                <input 
                  v-model="formData.paymentMethod" 
                  type="radio" 
                  value="cod" 
                  class="mr-3 text-[#e60023] focus:ring-[#e60023]"
                >
                <div>
                  <div class="font-medium">پرداخت در محل</div>
                  <div class="text-sm text-gray-600">پرداخت نقدی هنگام تحویل کالا</div>
                </div>
              </label>
            </div>
          </div>
        </div>

        <!-- خلاصه سفارش -->
        <div class="w-full lg:w-[400px]">
          <div class="bg-white rounded-2xl shadow px-4 py-4 sticky topx-4 py-4">
            <h2 class="text-xl font-bold text-[#1a2341] mb-4">خلاصه سفارش</h2>
            

            <!-- محاسبات -->
            <div class="space-y-2 border-t pt-4">
              <div class="flex justify-between text-gray-600">
                <span>قیمت کالاها</span>
                <span>{{ formatPrice(cartTotal) }}</span>
              </div>
              <div class="flex justify-between text-gray-600">
                <span>هزینه ارسال</span>
                <span>{{ formatPrice(shippingCost) }}</span>
              </div>
              <div class="flex justify-between font-bold text-lg text-[#1a2341] border-t pt-2">
                <span>مبلغ قابل پرداخت</span>
                <span>{{ formatPrice(totalAmount) }}</span>
              </div>
            </div>

            <!-- دکمه ثبت سفارش -->
            <button 
              type="button"
              @click="submitOrder"
              :disabled="orderLoading"
              class="w-full py-4 rounded-xl bg-[#e60023] text-white font-bold text-lg shadow-lg hover:bg-[#c9001b] transition disabled:opacity-50 disabled:cursor-not-allowed mt-6"
            >
              <span v-if="orderLoading">در حال پردازش...</span>
              <span v-else>ثبت و پرداخت سفارش</span>
            </button>

            <div class="text-xs text-gray-500 mt-3 text-center">
              با کلیک روی دکمه بالا، قوانین و مقررات فروشگاه را پذیرفته‌اید
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- ConfirmDialog -->
    <ConfirmDialog ref="confirmDialog" />
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; title?: string; middleware?: string[] }) => void
declare const navigateTo: (to: { path: string; query?: Record<string, string | number> } | string) => Promise<void>
declare const useAuth: () => { user: { value: { name?: string; id?: number } | null } }
</script>

<script setup lang="ts">
import ConfirmDialog from '@/components/common/ConfirmDialog.vue'
import { computed, onMounted, onUnmounted, ref, watch } from 'vue'
import { useAddresses } from '~/composables/useAddresses'
import { useCart } from '~/composables/useCart'
import AddressModal from '~/pages/checkout/components/AddressModal.vue'
import AddressSummary from '~/pages/checkout/components/AddressSummary.vue'

// تنظیم متا صفحه
definePageMeta({
  layout: 'default',
  title: 'تکمیل سفارش',
  middleware: ['auth']
})

// استفاده از composable احراز هویت
// Auth disabled
const { user } = useAuth()

// استفاده از composable سبد خرید
const { 
  cartItems, 
  cartTotal, 
  loading: cartLoading, 
  fetchCart, 
  clearCart,
  updateCartItem
} = useCart()

// وضعیت پردازش ثبت سفارش (مستقل از cartLoading که فقط مربوط به عملیات سبد است)
const orderLoading = ref(false)

// استفاده از composable آدرس‌ها

const { addresses, fetchAddresses, addAddress } = useAddresses()

const selectedAddressId = ref<number | null>(null)

// وضعیت مودال
const showAddressModal = ref(false)

// آدرس انتخاب‌شده بر اساس id
const summaryAddress = computed(()=>{
  return addresses.value.find(a=>a.id===selectedAddressId.value) ?? null
})

function onAddressSelected(addr:any){
  handleAddressSelection(addr.id)
}

// تابع انتخاب آدرس باید async باشد تا بتوان از await استفاده کرد
async function handleAddressSelection(id: number) {
  selectedAddressId.value = id
  const addr = addresses.value.find(a => a.id === id)
  if (addr) {
    formData.value.fullName = addr.recipient_name || user.value?.name || ''
    formData.value.phone = addr.recipient_mobile || ''
    formData.value.address = addr.full_address || ''
    formData.value.postalCode = addr.postal_code || ''
    selectedProvinceId.value = addr.province_id || null
    
    // اگر شهرها در کش موجود نیستند، دریافت کن
    if (addr.province_id && !citiesCache.has(addr.province_id)) {
      await fetchCities(addr.province_id)
    } else if (addr.province_id) {
      cities.value = citiesCache.get(addr.province_id)!
    }
    
    formData.value.cityId = addr.city_id || null
  }
}

// داده‌های فرم
const formData = ref({
  fullName: '',
  phone: '',
  address: '',
  postalCode: '',
  notes: '',
  paymentMethod: 'online',
  cityId: null as number | null
})

// هزینه ارسال
const shippingCost = ref<number>(0)

// استان / شهر - با کش
const provinces = ref<any[]>([])
const cities = ref<any[]>([])
const selectedProvinceId = ref<number|null>(null)
const provincesLoaded = ref(false)
const citiesCache = new Map<number, any[]>()

async function fetchProvinces(){
  if (provincesLoaded.value) return
  try {
    provinces.value = await $fetch('/api/geo/provinces')
    provincesLoaded.value = true
  } catch (error) {
    console.error('خطا در دریافت استان‌ها:', error)
  }
}

async function fetchCities(pid: number){
  if (citiesCache.has(pid)) {
    cities.value = citiesCache.get(pid)!
    return
  }
  
  try {
    const citiesData = await $fetch<any[]>(`/api/geo/provinces/${pid}/cities`)
    cities.value = citiesData
    citiesCache.set(pid, citiesData)
  } catch (error) {
    console.error('خطا در دریافت شهرها:', error)
    cities.value = []
  }
}

// بهینه‌سازی watch با debounce
let citiesFetchTimeout: NodeJS.Timeout | null = null
watch(selectedProvinceId, (val) => {
  if (citiesFetchTimeout) {
    clearTimeout(citiesFetchTimeout)
  }
  
  if (val) {
    citiesFetchTimeout = setTimeout(() => {
      fetchCities(val)
    }, 300) // تاخیر 300ms برای جلوگیری از فراخوانی‌های مکرر
  } else {
    cities.value = []
  }
})

// محاسبه مبلغ کل
const totalAmount = computed(() => {
  return cartTotal.value + shippingCost.value
})

// فرمت کردن قیمت
function formatPrice(val) {
  return val.toLocaleString('fa-IR') + ' تومان'
}

// به‌روزرسانی تعداد محصول
async function updateQuantity(item, newQuantity) {
  if (newQuantity < 1) {
    // حذف محصول از سبد خرید
    try {
      await $fetch(`/api/cart/remove`, {
        method: 'DELETE',
        body: { cart_item_id: item.id }
      })
      await fetchCart()
    } catch (error) {
      console.error('خطا در حذف محصول:', error)
    }
  } else {
    // به‌روزرسانی تعداد
    try {
      await updateCartItem(item.id, newQuantity)
    } catch (error) {
      console.error('خطا در به‌روزرسانی تعداد:', error)
    }
  }
}

// حذف مستقیم محصول
async function removeItem(item) {
  try {
    await $fetch(`/api/cart/remove`, {
      method: 'DELETE',
      body: { cart_item_id: item.id }
    })
    await fetchCart()
  } catch (error) {
    console.error('خطا در حذف محصول:', error)
  }
}

// دریافت سبد خرید در ابتدای بارگذاری
onMounted(async () => {
  await fetchCart()
  await fetchAddresses()
  await fetchProvinces()
  
  // اگر کاربر آدرس‌هایی دارد، آدرس پیش‌فرض را انتخاب کن
  if (addresses.value.length > 0) {
    // پیدا کردن آدرس پیش‌فرض
    const defaultAddress = addresses.value.find(addr => addr.is_default)
    if (defaultAddress) {
      selectedAddressId.value = defaultAddress.id
      // پر کردن فرم با اطلاعات آدرس پیش‌فرض
      await handleAddressSelection(defaultAddress.id)
    } else {
      // اگر آدرس پیش‌فرضی وجود ندارد، اولین آدرس را انتخاب کن
      selectedAddressId.value = addresses.value[0].id
      await handleAddressSelection(addresses.value[0].id)
    }
  }
  
  // پرکردن فیلدها با اطلاعات کاربر در صورت خالی بودن (فقط اگر آدرسی انتخاب نشده)
  if (user.value && !selectedAddressId.value) {
    if (!formData.value.fullName) formData.value.fullName = user.value.name || ''
    if (!formData.value.phone) formData.value.phone = ''
  }
})

// پاک‌سازی timeout ها
onUnmounted(() => {
  if (citiesFetchTimeout) {
    clearTimeout(citiesFetchTimeout)
  }
})

// ثبت سفارش
async function submitOrder() {
  // جلوگیری از ارسال تکراری - باید در ابتدا چک شود
  if (orderLoading.value) {
    console.log('در حال پردازش سفارش قبلی...')
    return
  }

  // تنظیم فوری orderLoading برای جلوگیری از double-click
  orderLoading.value = true

  console.log('شروع ثبت سفارش - cartItems:', cartItems.value)
  console.log('تعداد آیتم‌های سبد خرید:', cartItems.value.length)

  try {
    // اگر کاربر آدرس انتخاب نکرده و addresses خالی است، ابتدا ذخیره آدرس
    if (!selectedAddressId.value) {
      try {
        const addrRes = await addAddress({
          full_address: formData.value.address,
          postal_code: formData.value.postalCode,
          recipient_name: formData.value.fullName,
          recipient_mobile: formData.value.phone,
          is_default: true
        })
        selectedAddressId.value = addrRes.id
      } catch(e){
        console.error('خطا در ذخیره آدرس',e)
        return
      }
    }

    if (!formData.value.fullName || !formData.value.phone || !formData.value.address) {
      const confirmDialog = ref()
      await confirmDialog.value?.show({
        title: 'خطا',
        message: 'لطفاً اطلاعات ضروری را تکمیل کنید',
        confirmText: 'تأیید',
        cancelText: '',
        type: 'danger'
      })
      return
    }

    if (!formData.value.paymentMethod) {
      const confirmDialog = ref()
      await confirmDialog.value?.show({
        title: 'خطا',
        message: 'لطفاً روش پرداخت را انتخاب کنید',
        confirmText: 'تأیید',
        cancelText: '',
        type: 'danger'
      })
      return
    }

    if (cartItems.value.length === 0) {
      const confirmDialog = ref()
      await confirmDialog.value?.show({
        title: 'خطا',
        message: 'سبد خرید شما خالی است',
        confirmText: 'تأیید',
        cancelText: '',
        type: 'danger'
      })
      return
    }
    
    // فراخوانی API ثبت سفارش
    const orderData = {
      ...formData.value,
      shipping_address_id: selectedAddressId.value,
      payment_method: formData.value.paymentMethod,
      items: cartItems.value.map(i => ({ 
        product_id: i.product_id, 
        quantity: i.quantity || 1,
        product: {
          name: i.name,
          price: i.price
        }
      }))
    }
    
    console.log('ارسال داده‌های سفارش:', orderData)
    
    const res = await $fetch<{success?: boolean, data?: {orderId?: number, orderNumber?: string}, id?: number}>('/api/orders/create', {
      method: 'POST',
      body: orderData
    })
    
    console.log('پاسخ API:', res)
    
    if (res && res.success && res.data) {
      // پاک‌کردن سبد خرید
      await clearCart()

      // نوتیفیکیشن موفقیت ساده
      const oid = res.data.orderId
      const orderNumber = res.data.orderNumber
      
      // استفاده از ConfirmDialog به جای alert
      const confirmDialog = ref()
      await confirmDialog.value?.show({
        title: 'موفقیت',
        message: `سفارش با موفقیت ثبت شد. شماره سفارش: ${orderNumber}`,
        confirmText: 'تأیید',
        cancelText: '',
        type: 'info'
      })

      // هدایت کاربر به صفحه جزئیات سفارش
      navigateTo({ path: '/orders/completed', query: { id: oid } })
      return
    } else {
      console.error('فرمت پاسخ API نامعتبر:', res)
      throw new Error('فرمت پاسخ API نامعتبر است')
    }
    
  } catch (error) {
    console.error('خطا در ثبت سفارش:', error)
    
    // نمایش پیام خطا
    let errorMessage = 'خطا در ثبت سفارش. لطفاً دوباره تلاش کنید'
    
    if (error.data?.statusMessage) {
      errorMessage = error.data.statusMessage
    } else if (error.message) {
      errorMessage = error.message
    }
    
    // استفاده از ConfirmDialog به جای alert
    const confirmDialog = ref()
    await confirmDialog.value?.show({
      title: 'خطا',
      message: errorMessage,
      confirmText: 'تأیید',
      cancelText: '',
      type: 'danger'
    })
  } finally {
    orderLoading.value = false
  }
}
</script>

<style scoped>
.checkout-root { 
  font-family: 'IranYekan', 'Vazir', 'IRANSansX', 'Tahoma', sans-serif; 
}

.hide-scrollbar {
  -ms-overflow-style: none;
  scrollbar-width: none;
}

.hide-scrollbar::-webkit-scrollbar {
  display: none;
}
</style> 
<template>
  <div class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
    <div class="relative top-20 mx-auto p-5 border w-11/12 max-w-2xl shadow-lg rounded-md bg-white">
      <!-- هدر مودال -->
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-lg font-semibold text-gray-900">ثبت مشتری جدید</h3>
        <button
          class="text-gray-400 hover:text-gray-600"
          @click="$emit('close')"
        >
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
        </button>
      </div>

      <!-- فرم ثبت مشتری -->
      <form class="space-y-6" @submit.prevent="createCustomer">
        <!-- اطلاعات شخصی -->
        <div class="bg-gray-50 p-6 rounded-lg">
          <h4 class="text-md font-medium text-gray-900 mb-4">اطلاعات شخصی</h4>
          
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                نام و نام خانوادگی <span class="text-red-500">*</span>
              </label>
              <input
                v-model="form.fullName"
                type="text"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="مثال: احمد محمدی"
              />
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                کد مشتری
              </label>
              <div class="flex">
                <input
                  v-model="form.customerCode"
                  type="text"
                  class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                  placeholder="CUST-001"
                />
                <button
                  type="button"
                  class="px-4 py-2 bg-gray-600 text-white text-sm font-medium rounded-l-lg hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
                  @click="generateCustomerCode"
                >
                  تولید
                </button>
              </div>
            </div>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mt-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                تاریخ تولد
              </label>
              <input
                v-model="form.birthDate"
                type="date"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              />
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                جنسیت
              </label>
              <select
                v-model="form.gender"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              >
                <option value="">انتخاب کنید</option>
                <option value="male">مرد</option>
                <option value="female">زن</option>
                <option value="other">سایر</option>
              </select>
            </div>
          </div>
        </div>

        <!-- اطلاعات تماس -->
        <div class="bg-gray-50 p-6 rounded-lg">
          <h4 class="text-md font-medium text-gray-900 mb-4">اطلاعات تماس</h4>
          
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                ایمیل <span class="text-red-500">*</span>
              </label>
              <input
                v-model="form.email"
                type="email"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="example@email.com"
              />
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                شماره موبایل <span class="text-red-500">*</span>
              </label>
              <input
                v-model="form.phone"
                type="tel"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="09123456789"
              />
            </div>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mt-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                شماره ثابت
              </label>
              <input
                v-model="form.landline"
                type="tel"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="02112345678"
              />
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                کد ملی
              </label>
              <input
                v-model="form.nationalCode"
                type="text"
                maxlength="10"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="1234567890"
              />
            </div>
          </div>
        </div>

        <!-- اطلاعات آدرس -->
        <div class="bg-gray-50 p-6 rounded-lg">
          <h4 class="text-md font-medium text-gray-900 mb-4">اطلاعات آدرس</h4>
          
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                استان
              </label>
              <select
                v-model="form.province"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                @change="onProvinceChange"
              >
                <option value="">انتخاب کنید</option>
                <option value="tehran">تهران</option>
                <option value="isfahan">اصفهان</option>
                <option value="shiraz">شیراز</option>
                <option value="mashhad">مشهد</option>
                <option value="tabriz">تبریز</option>
                <option value="kerman">کرمان</option>
                <option value="yazd">یزد</option>
                <option value="kermanshah">کرمانشاه</option>
              </select>
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                شهر
              </label>
              <select
                v-model="form.city"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              >
                <option value="">ابتدا استان را انتخاب کنید</option>
                <option v-for="city in availableCities" :key="city" :value="city">
                  {{ city }}
                </option>
              </select>
            </div>
          </div>

          <div class="mt-4">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              آدرس کامل
            </label>
            <textarea
              v-model="form.address"
              rows="3"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="خیابان، کوچه، پلاک، طبقه..."
            ></textarea>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mt-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                کد پستی
              </label>
              <input
                v-model="form.postalCode"
                type="text"
                maxlength="10"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="1234567890"
              />
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                شماره واحد
              </label>
              <input
                v-model="form.unitNumber"
                type="text"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="مثال: 5"
              />
            </div>
          </div>
        </div>

        <!-- اطلاعات تکمیلی -->
        <div class="bg-gray-50 p-6 rounded-lg">
          <h4 class="text-md font-medium text-gray-900 mb-4">اطلاعات تکمیلی</h4>
          
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                منبع معرفی
              </label>
              <select
                v-model="form.referralSource"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              >
                <option value="">انتخاب کنید</option>
                <option value="website">وب‌سایت</option>
                <option value="social_media">شبکه‌های اجتماعی</option>
                <option value="friend">معرفی دوست</option>
                <option value="advertisement">تبلیغات</option>
                <option value="search">جستجو</option>
                <option value="other">سایر</option>
              </select>
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                نوع مشتری
              </label>
              <select
                v-model="form.customerType"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              >
                <option value="regular">عادی</option>
                <option value="vip">VIP</option>
                <option value="premium">پریمیوم</option>
                <option value="corporate">شرکتی</option>
              </select>
            </div>
          </div>

          <div class="mt-4">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              توضیحات
            </label>
            <textarea
              v-model="form.notes"
              rows="3"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="توضیحات اضافی درباره مشتری..."
            ></textarea>
          </div>
        </div>

        <!-- تنظیمات اعلان‌ها -->
        <div class="bg-gray-50 p-6 rounded-lg">
          <h4 class="text-md font-medium text-gray-900 mb-4">تنظیمات اعلان‌ها</h4>
          
          <div class="space-y-3">
            <label class="flex items-center">
              <input
                v-model="form.notifications.email"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
              />
              <span class="mr-2 text-sm text-gray-700">اعلان‌های ایمیلی</span>
            </label>
            
            <label class="flex items-center">
              <input
                v-model="form.notifications.sms"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
              />
              <span class="mr-2 text-sm text-gray-700">اعلان‌های پیامکی</span>
            </label>
            
            <label class="flex items-center">
              <input
                v-model="form.notifications.push"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
              />
              <span class="mr-2 text-sm text-gray-700">اعلان‌های push</span>
            </label>
            
            <label class="flex items-center">
              <input
                v-model="form.notifications.newsletter"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
              />
              <span class="mr-2 text-sm text-gray-700">خبرنامه و تخفیف‌ها</span>
            </label>
          </div>
        </div>

        <!-- وضعیت -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            وضعیت <span class="text-red-500">*</span>
          </label>
          <select
            v-model="form.status"
            required
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          >
            <option value="">انتخاب کنید</option>
            <option value="active">فعال</option>
            <option value="inactive">غیرفعال</option>
            <option value="pending">در انتظار تایید</option>
          </select>
        </div>

        <!-- دکمه‌های عملیات -->
        <div class="flex justify-end space-x-3 space-x-reverse">
          <button
            type="button"
            class="px-4 py-2 bg-gray-300 text-gray-700 text-sm font-medium rounded-lg hover:bg-gray-400 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
            @click="$emit('close')"
          >
            انصراف
          </button>
          <button
            type="submit"
            class="px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
          >
            ثبت مشتری
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
// تعریف props و emits
const emit = defineEmits(['close', 'created'])

// تعریف متغیرهای reactive
const form = ref({
  fullName: '',
  customerCode: '',
  birthDate: '',
  gender: '',
  email: '',
  phone: '',
  landline: '',
  nationalCode: '',
  province: '',
  city: '',
  address: '',
  postalCode: '',
  unitNumber: '',
  referralSource: '',
  customerType: 'regular',
  notes: '',
  notifications: {
    email: true,
    sms: true,
    push: false,
    newsletter: true
  },
  status: 'active'
})

// شهرهای موجود بر اساس استان
const citiesByProvince = {
  tehran: ['تهران', 'شهریار', 'ورامین', 'فیروزکوه', 'دماوند'],
  isfahan: ['اصفهان', 'کاشان', 'نجف‌آباد', 'خمینی‌شهر', 'شاهین‌شهر'],
  shiraz: ['شیراز', 'مرودشت', 'جهرم', 'فسا', 'کازرون'],
  mashhad: ['مشهد', 'نیشابور', 'سبزوار', 'تربت حیدریه', 'کاشمر'],
  tabriz: ['تبریز', 'مراغه', 'میانه', 'اهر', 'مرند'],
  kerman: ['کرمان', 'رفسنجان', 'سیرجان', 'بم', 'جیرفت'],
  yazd: ['یزد', 'میبد', 'اردکان', 'تفت', 'ابرکوه'],
  kermanshah: ['کرمانشاه', 'اسلام‌آباد غرب', 'کنگاور', 'سنقر', 'پاوه']
}

const availableCities = computed(() => {
  return form.value.province ? citiesByProvince[form.value.province] || [] : []
})

// توابع کمکی
const generateCustomerCode = () => {
  const chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789'
  let result = 'CUST-'
  for (let i = 0; i < 6; i++) {
    result += chars.charAt(Math.floor(Math.random() * chars.length))
  }
  form.value.customerCode = result
}

const onProvinceChange = () => {
  form.value.city = ''
}

// تابع ایجاد مشتری
const createCustomer = () => {
  // اعتبارسنجی فرم
  if (!form.value.fullName || !form.value.email || !form.value.phone || !form.value.status) {
    alert('لطفاً تمام فیلدهای اجباری را پر کنید')
    return
  }

  // اعتبارسنجی ایمیل
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  if (!emailRegex.test(form.value.email)) {
    alert('لطفاً ایمیل معتبر وارد کنید')
    return
  }

  // اعتبارسنجی شماره موبایل
  const phoneRegex = /^09\d{9}$/
  if (!phoneRegex.test(form.value.phone)) {
    alert('لطفاً شماره موبایل معتبر وارد کنید')
    return
  }

  // اعتبارسنجی کد ملی
  if (form.value.nationalCode && form.value.nationalCode.length !== 10) {
    alert('کد ملی باید 10 رقم باشد')
    return
  }

  // اعتبارسنجی کد پستی
  if (form.value.postalCode && form.value.postalCode.length !== 10) {
    alert('کد پستی باید 10 رقم باشد')
    return
  }

  // ایجاد مشتری جدید
  const newCustomer = {
    fullName: form.value.fullName,
    customerCode: form.value.customerCode || generateCustomerCode(),
    birthDate: form.value.birthDate || null,
    gender: form.value.gender || null,
    email: form.value.email,
    phone: form.value.phone,
    landline: form.value.landline || null,
    nationalCode: form.value.nationalCode || null,
    province: form.value.province || null,
    city: form.value.city || null,
    address: form.value.address || null,
    postalCode: form.value.postalCode || null,
    unitNumber: form.value.unitNumber || null,
    referralSource: form.value.referralSource || null,
    customerType: form.value.customerType,
    notes: form.value.notes || '',
    notifications: form.value.notifications,
    status: form.value.status,
    createdAt: new Date()
  }

  // ارسال به کامپوننت والد
  emit('created', newCustomer)
}
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 

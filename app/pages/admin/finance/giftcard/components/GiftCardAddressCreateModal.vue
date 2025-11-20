<template>
  <div class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
    <div class="relative top-20 mx-auto p-5 border w-11/12 max-w-2xl shadow-lg rounded-md bg-white">
      <!-- هدر مودال -->
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-lg font-semibold text-gray-900">افزودن آدرس جدید</h3>
        <button
          class="text-gray-400 hover:text-gray-600"
          @click="$emit('close')"
        >
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
        </button>
      </div>

      <!-- فرم آدرس -->
      <form class="space-y-6" @submit.prevent="createAddress">
        <!-- عنوان آدرس -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            عنوان آدرس <span class="text-red-500">*</span>
          </label>
          <input
            v-model="form.title"
            type="text"
            required
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="مثال: آدرس خانه، آدرس محل کار"
          />
        </div>

        <!-- اطلاعات گیرنده -->
        <div class="bg-gray-50 p-6 rounded-lg">
          <h4 class="text-md font-medium text-gray-900 mb-4">اطلاعات گیرنده</h4>
          
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
                placeholder="نام گیرنده"
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
                استان <span class="text-red-500">*</span>
              </label>
              <select
                v-model="form.province"
                required
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
                <option value="kerman">کرمانشاه</option>
                <option value="alborz">البرز</option>
                <option value="qom">قم</option>
                <option value="markazi">مرکزی</option>
                <option value="hamadan">همدان</option>
                <option value="zanjan">زنجان</option>
                <option value="qazvin">قزوین</option>
                <option value="gilan">گیلان</option>
                <option value="mazandaran">مازندران</option>
                <option value="golestan">گلستان</option>
                <option value="semnan">سمنان</option>
                <option value="ardabil">اردبیل</option>
                <option value="west_azerbaijan">آذربایجان غربی</option>
                <option value="east_azerbaijan">آذربایجان شرقی</option>
                <option value="kurdistan">کردستان</option>
                <option value="kermanshah">کرمانشاه</option>
                <option value="ilam">ایلام</option>
                <option value="lorestan">لرستان</option>
                <option value="chaharmahal">چهارمحال و بختیاری</option>
                <option value="kohgiluyeh">کهگیلویه و بویراحمد</option>
                <option value="bushehr">بوشهر</option>
                <option value="fars">فارس</option>
                <option value="hormozgan">هرمزگان</option>
                <option value="sistan">سیستان و بلوچستان</option>
                <option value="kerman">کرمان</option>
                <option value="yazd">یزد</option>
                <option value="south_khorasan">خراسان جنوبی</option>
                <option value="razavi_khorasan">خراسان رضوی</option>
                <option value="north_khorasan">خراسان شمالی</option>
              </select>
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                شهر <span class="text-red-500">*</span>
              </label>
              <select
                v-model="form.city"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              >
                <option value="">ابتدا استان را انتخاب کنید</option>
                <option v-for="city in availableCities" :key="city" :value="city">
                  {{ city }}
                </option>
              </select>
            </div>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mt-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                بخش/شهرستان
              </label>
              <input
                v-model="form.district"
                type="text"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="مثال: بخش مرکزی"
              />
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                روستا/دهستان
              </label>
              <input
                v-model="form.village"
                type="text"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="مثال: روستای نمونه"
              />
            </div>
          </div>

          <div class="mt-4">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              آدرس کامل <span class="text-red-500">*</span>
            </label>
            <textarea
              v-model="form.address"
              rows="3"
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="خیابان، کوچه، پلاک، طبقه، واحد..."
            ></textarea>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mt-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                کد پستی <span class="text-red-500">*</span>
              </label>
              <input
                v-model="form.postalCode"
                type="text"
                maxlength="10"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="1234567890"
              />
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                شماره پلاک
              </label>
              <input
                v-model="form.plaqueNumber"
                type="text"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="مثال: 123"
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
                نوع آدرس
              </label>
              <select
                v-model="form.addressType"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              >
                <option value="home">خانه</option>
                <option value="work">محل کار</option>
                <option value="relative">خانواده</option>
                <option value="friend">دوست</option>
                <option value="other">سایر</option>
              </select>
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                اولویت ارسال
              </label>
              <select
                v-model="form.priority"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              >
                <option value="normal">عادی</option>
                <option value="high">بالا</option>
                <option value="urgent">فوری</option>
              </select>
            </div>
          </div>

          <div class="mt-4">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              راهنمای دسترسی
            </label>
            <textarea
              v-model="form.accessGuide"
              rows="2"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="راهنمای دسترسی به آدرس (اختیاری)..."
            ></textarea>
          </div>

          <div class="mt-4">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              توضیحات اضافی
            </label>
            <textarea
              v-model="form.notes"
              rows="2"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="توضیحات اضافی درباره آدرس..."
            ></textarea>
          </div>
        </div>

        <!-- تنظیمات آدرس -->
        <div class="bg-gray-50 p-6 rounded-lg">
          <h4 class="text-md font-medium text-gray-900 mb-4">تنظیمات آدرس</h4>
          
          <div class="space-y-3">
            <label class="flex items-center">
              <input
                v-model="form.isDefault"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
              />
              <span class="mr-2 text-sm text-gray-700">آدرس پیش‌فرض</span>
            </label>
            
            <label class="flex items-center">
              <input
                v-model="form.isActive"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
              />
              <span class="mr-2 text-sm text-gray-700">آدرس فعال</span>
            </label>
            
            <label class="flex items-center">
              <input
                v-model="form.requiresSignature"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
              />
              <span class="mr-2 text-sm text-gray-700">نیاز به امضا</span>
            </label>
            
            <label class="flex items-center">
              <input
                v-model="form.isBusinessAddress"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
              />
              <span class="mr-2 text-sm text-gray-700">آدرس تجاری</span>
            </label>
          </div>
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
            افزودن آدرس
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
  title: '',
  fullName: '',
  phone: '',
  landline: '',
  nationalCode: '',
  province: '',
  city: '',
  district: '',
  village: '',
  address: '',
  postalCode: '',
  plaqueNumber: '',
  unitNumber: '',
  addressType: 'home',
  priority: 'normal',
  accessGuide: '',
  notes: '',
  isDefault: false,
  isActive: true,
  requiresSignature: false,
  isBusinessAddress: false
})

// شهرهای موجود بر اساس استان
const citiesByProvince = {
  tehran: ['تهران', 'شهریار', 'ورامین', 'فیروزکوه', 'دماوند', 'پردیس', 'پردیس', 'رباط کریم', 'پیشوا', 'ری'],
  isfahan: ['اصفهان', 'کاشان', 'نجف‌آباد', 'خمینی‌شهر', 'شاهین‌شهر', 'فولادشهر', 'مبارکه', 'گلپایگان', 'نطنز', 'خوانسار'],
  shiraz: ['شیراز', 'مرودشت', 'جهرم', 'فسا', 'کازرون', 'لار', 'داراب', 'نی ریز', 'آباده', 'اقلید'],
  mashhad: ['مشهد', 'نیشابور', 'سبزوار', 'تربت حیدریه', 'کاشمر', 'گناباد', 'قوچان', 'تربت جام', 'چناران', 'سرخس'],
  tabriz: ['تبریز', 'مراغه', 'میانه', 'اهر', 'مرند', 'بناب', 'سراب', 'شبستر', 'هشترود', 'کلیبر'],
  kerman: ['کرمان', 'رفسنجان', 'سیرجان', 'بم', 'جیرفت', 'کهنوج', 'زرند', 'بافت', 'راور', 'کوهبنان'],
  yazd: ['یزد', 'میبد', 'اردکان', 'تفت', 'ابرکوه', 'بافق', 'مهریز', 'اشکذر', 'خاتم', 'بهاباد'],
  alborz: ['کرج', 'فردیس', 'محمدشهر', 'ماهدشت', 'اشتهارد', 'چهارباغ', 'کمال شهر', 'مشکین دشت', 'ماهدشت', 'اشتهارد'],
  qom: ['قم', 'پردیسان', 'قنوات', 'دستجرد', 'سلفچگان', 'جعفریه', 'کهک', 'سلفچگان', 'دستجرد', 'قنوات'],
  markazi: ['اراک', 'خمین', 'محلات', 'دلیجان', 'ساوه', 'تفرش', 'آشتیان', 'کمیجان', 'شازند', 'زرندیه']
}

const availableCities = computed(() => {
  return form.value.province ? citiesByProvince[form.value.province] || [] : []
})

// توابع کمکی
const onProvinceChange = () => {
  form.value.city = ''
}

// تابع ایجاد آدرس
const createAddress = () => {
  // اعتبارسنجی فرم
  if (!form.value.title || !form.value.fullName || !form.value.phone || !form.value.province || !form.value.city || !form.value.address || !form.value.postalCode) {
    alert('لطفاً تمام فیلدهای اجباری را پر کنید')
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
  if (form.value.postalCode.length !== 10) {
    alert('کد پستی باید 10 رقم باشد')
    return
  }

  // اعتبارسنجی شماره ثابت
  if (form.value.landline) {
    const landlineRegex = /^0\d{10}$/
    if (!landlineRegex.test(form.value.landline)) {
      alert('لطفاً شماره ثابت معتبر وارد کنید')
      return
    }
  }

  // ایجاد آدرس جدید
  const newAddress = {
    title: form.value.title,
    fullName: form.value.fullName,
    phone: form.value.phone,
    landline: form.value.landline || null,
    nationalCode: form.value.nationalCode || null,
    province: form.value.province,
    city: form.value.city,
    district: form.value.district || null,
    village: form.value.village || null,
    address: form.value.address,
    postalCode: form.value.postalCode,
    plaqueNumber: form.value.plaqueNumber || null,
    unitNumber: form.value.unitNumber || null,
    addressType: form.value.addressType,
    priority: form.value.priority,
    accessGuide: form.value.accessGuide || null,
    notes: form.value.notes || null,
    isDefault: form.value.isDefault,
    isActive: form.value.isActive,
    requiresSignature: form.value.requiresSignature,
    isBusinessAddress: form.value.isBusinessAddress,
    createdAt: new Date()
  }

  // ارسال به کامپوننت والد
  emit('created', newAddress)
}
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 

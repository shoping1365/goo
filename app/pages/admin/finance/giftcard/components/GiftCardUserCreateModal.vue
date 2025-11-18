<template>
  <div class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
    <div class="relative top-20 mx-auto p-5 border w-11/12 md:w-3/4 lg:w-1/2 shadow-lg rounded-md bg-white">
      <!-- هدر مودال -->
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-lg font-semibold text-gray-900">ثبت استفاده‌کننده جدید</h3>
        <button
          @click="$emit('close')"
          class="text-gray-400 hover:text-gray-600"
        >
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
        </button>
      </div>

      <!-- فرم ثبت -->
      <form @submit.prevent="handleSubmit" class="space-y-6">
        <!-- اطلاعات شخصی -->
        <div class="bg-gray-50 p-6 rounded-lg">
          <h4 class="text-md font-semibold text-gray-900 mb-4">اطلاعات شخصی</h4>
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
                placeholder="نام و نام خانوادگی"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                کد ملی <span class="text-red-500">*</span>
              </label>
              <input
                v-model="form.nationalCode"
                type="text"
                required
                maxlength="10"
                pattern="[0-9]{10}"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="کد ملی ۱۰ رقمی"
              />
            </div>
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
              </select>
            </div>
          </div>
        </div>

        <!-- اطلاعات تماس -->
        <div class="bg-gray-50 p-6 rounded-lg">
          <h4 class="text-md font-semibold text-gray-900 mb-4">اطلاعات تماس</h4>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                شماره موبایل <span class="text-red-500">*</span>
              </label>
              <input
                v-model="form.mobile"
                type="tel"
                required
                pattern="09[0-9]{9}"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="09xxxxxxxxx"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                ایمیل
              </label>
              <input
                v-model="form.email"
                type="email"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="example@email.com"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                شماره تلفن ثابت
              </label>
              <input
                v-model="form.phone"
                type="tel"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="021xxxxxxxx"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                آدرس وب‌سایت
              </label>
              <input
                v-model="form.website"
                type="url"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="https://example.com"
              />
            </div>
          </div>
        </div>

        <!-- آدرس -->
        <div class="bg-gray-50 p-6 rounded-lg">
          <h4 class="text-md font-semibold text-gray-900 mb-4">آدرس</h4>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                استان
              </label>
              <select
                v-model="form.province"
                @change="handleProvinceChange"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              >
                <option value="">انتخاب استان</option>
                <option value="tehran">تهران</option>
                <option value="isfahan">اصفهان</option>
                <option value="shiraz">شیراز</option>
                <option value="mashhad">مشهد</option>
                <option value="tabriz">تبریز</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                شهر
              </label>
              <select
                v-model="form.city"
                :disabled="!form.province"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              >
                <option value="">انتخاب شهر</option>
                <option v-if="form.province === 'tehran'" value="tehran">تهران</option>
                <option v-if="form.province === 'tehran'" value="karaj">کرج</option>
                <option v-if="form.province === 'isfahan'" value="isfahan">اصفهان</option>
                <option v-if="form.province === 'isfahan'" value="kashan">کاشان</option>
                <option v-if="form.province === 'shiraz'" value="shiraz">شیراز</option>
                <option v-if="form.province === 'mashhad'" value="mashhad">مشهد</option>
                <option v-if="form.province === 'tabriz'" value="tabriz">تبریز</option>
              </select>
            </div>
            <div class="md:col-span-2">
              <label class="block text-sm font-medium text-gray-700 mb-2">
                آدرس کامل
              </label>
              <textarea
                v-model="form.address"
                rows="3"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="آدرس کامل پستی"
              ></textarea>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                کد پستی
              </label>
              <input
                v-model="form.postalCode"
                type="text"
                maxlength="10"
                pattern="[0-9]{10}"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="کد پستی ۱۰ رقمی"
              />
            </div>
          </div>
        </div>

        <!-- اطلاعات تکمیلی -->
        <div class="bg-gray-50 p-6 rounded-lg">
          <h4 class="text-md font-semibold text-gray-900 mb-4">اطلاعات تکمیلی</h4>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                شغل
              </label>
              <input
                v-model="form.occupation"
                type="text"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="شغل"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                شرکت
              </label>
              <input
                v-model="form.company"
                type="text"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="نام شرکت"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                سطح تحصیلات
              </label>
              <select
                v-model="form.education"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              >
                <option value="">انتخاب کنید</option>
                <option value="diploma">دیپلم</option>
                <option value="associate">کاردانی</option>
                <option value="bachelor">کارشناسی</option>
                <option value="master">کارشناسی ارشد</option>
                <option value="phd">دکترا</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                درآمد ماهانه
              </label>
              <select
                v-model="form.monthlyIncome"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              >
                <option value="">انتخاب کنید</option>
                <option value="under_5m">کمتر از ۵ میلیون</option>
                <option value="5m_10m">۵ تا ۱۰ میلیون</option>
                <option value="10m_20m">۱۰ تا ۲۰ میلیون</option>
                <option value="20m_50m">۲۰ تا ۵۰ میلیون</option>
                <option value="over_50m">بیش از ۵۰ میلیون</option>
              </select>
            </div>
          </div>
        </div>

        <!-- تنظیمات اعلان‌ها -->
        <div class="bg-gray-50 p-6 rounded-lg">
          <h4 class="text-md font-semibold text-gray-900 mb-4">تنظیمات اعلان‌ها</h4>
          <div class="space-y-3">
            <div class="flex items-center">
              <input
                v-model="form.notifications.sms"
                type="checkbox"
                id="sms-notifications"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <label for="sms-notifications" class="mr-2 block text-sm text-gray-900">
                اعلان‌های پیامکی
              </label>
            </div>
            <div class="flex items-center">
              <input
                v-model="form.notifications.email"
                type="checkbox"
                id="email-notifications"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <label for="email-notifications" class="mr-2 block text-sm text-gray-900">
                اعلان‌های ایمیل
              </label>
            </div>
            <div class="flex items-center">
              <input
                v-model="form.notifications.push"
                type="checkbox"
                id="push-notifications"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <label for="push-notifications" class="mr-2 block text-sm text-gray-900">
                اعلان‌های push
              </label>
            </div>
            <div class="flex items-center">
              <input
                v-model="form.notifications.newsletter"
                type="checkbox"
                id="newsletter"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <label for="newsletter" class="mr-2 block text-sm text-gray-900">
                خبرنامه و تخفیف‌ها
              </label>
            </div>
          </div>
        </div>

        <!-- تنظیمات امنیتی -->
        <div class="bg-gray-50 p-6 rounded-lg">
          <h4 class="text-md font-semibold text-gray-900 mb-4">تنظیمات امنیتی</h4>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                رمز عبور <span class="text-red-500">*</span>
              </label>
              <input
                v-model="form.password"
                type="password"
                required
                minlength="8"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="حداقل ۸ کاراکتر"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                تکرار رمز عبور <span class="text-red-500">*</span>
              </label>
              <input
                v-model="form.confirmPassword"
                type="password"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="تکرار رمز عبور"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                سوال امنیتی
              </label>
              <select
                v-model="form.securityQuestion"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              >
                <option value="">انتخاب کنید</option>
                <option value="birth_place">محل تولد شما کجاست؟</option>
                <option value="first_pet">نام اولین حیوان خانگی شما چیست؟</option>
                <option value="mother_maiden">نام خانوادگی مادر شما چیست؟</option>
                <option value="favorite_color">رنگ مورد علاقه شما چیست؟</option>
                <option value="first_school">نام اولین مدرسه شما چیست؟</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                پاسخ سوال امنیتی
              </label>
              <input
                v-model="form.securityAnswer"
                type="text"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="پاسخ سوال امنیتی"
              />
            </div>
          </div>
        </div>

        <!-- وضعیت حساب -->
        <div class="bg-gray-50 p-6 rounded-lg">
          <h4 class="text-md font-semibold text-gray-900 mb-4">وضعیت حساب</h4>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                وضعیت حساب <span class="text-red-500">*</span>
              </label>
              <select
                v-model="form.status"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              >
                <option value="active">فعال</option>
                <option value="inactive">غیرفعال</option>
                <option value="pending">در انتظار تایید</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                سطح دسترسی
              </label>
              <select
                v-model="form.accessLevel"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              >
                <option value="basic">پایه</option>
                <option value="premium">پریمیوم</option>
                <option value="vip">VIP</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                تاریخ انقضا حساب
              </label>
              <input
                v-model="form.expiryDate"
                type="date"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                توضیحات
              </label>
              <textarea
                v-model="form.notes"
                rows="2"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="توضیحات اضافی"
              ></textarea>
            </div>
          </div>
        </div>

        <!-- دکمه‌های عملیات -->
        <div class="flex justify-end gap-3 pt-6 border-t border-gray-200">
          <button
            type="button"
            @click="$emit('close')"
            class="px-4 py-2 bg-gray-300 text-gray-700 text-sm font-medium rounded-lg hover:bg-gray-400 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
          >
            انصراف
          </button>
          <button
            type="submit"
            :disabled="isSubmitting"
            class="px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed"
          >
            <span v-if="isSubmitting">در حال ثبت...</span>
            <span v-else>ثبت استفاده‌کننده</span>
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
// تعریف props و emits
const emit = defineEmits(['close', 'created'])

// متغیرهای reactive
const isSubmitting = ref(false)

// فرم داده‌ها
const form = ref({
  // اطلاعات شخصی
  fullName: '',
  nationalCode: '',
  birthDate: '',
  gender: '',
  
  // اطلاعات تماس
  mobile: '',
  email: '',
  phone: '',
  website: '',
  
  // آدرس
  province: '',
  city: '',
  address: '',
  postalCode: '',
  
  // اطلاعات تکمیلی
  occupation: '',
  company: '',
  education: '',
  monthlyIncome: '',
  
  // تنظیمات اعلان‌ها
  notifications: {
    sms: true,
    email: false,
    push: true,
    newsletter: false
  },
  
  // تنظیمات امنیتی
  password: '',
  confirmPassword: '',
  securityQuestion: '',
  securityAnswer: '',
  
  // وضعیت حساب
  status: 'active',
  accessLevel: 'basic',
  expiryDate: '',
  notes: ''
})

// توابع
const handleProvinceChange = () => {
  form.value.city = ''
}

const validateForm = () => {
  const errors = []
  
  // اعتبارسنجی نام
  if (!form.value.fullName.trim()) {
    errors.push('نام و نام خانوادگی الزامی است')
  }
  
  // اعتبارسنجی کد ملی
  if (!form.value.nationalCode || form.value.nationalCode.length !== 10) {
    errors.push('کد ملی باید ۱۰ رقم باشد')
  }
  
  // اعتبارسنجی موبایل
  if (!form.value.mobile || !/^09[0-9]{9}$/.test(form.value.mobile)) {
    errors.push('شماره موبایل معتبر نیست')
  }
  
  // اعتبارسنجی ایمیل
  if (form.value.email && !/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(form.value.email)) {
    errors.push('ایمیل معتبر نیست')
  }
  
  // اعتبارسنجی رمز عبور
  if (form.value.password.length < 8) {
    errors.push('رمز عبور باید حداقل ۸ کاراکتر باشد')
  }
  
  if (form.value.password !== form.value.confirmPassword) {
    errors.push('رمز عبور و تکرار آن یکسان نیستند')
  }
  
  return errors
}

const handleSubmit = async () => {
  const errors = validateForm()
  
  if (errors.length > 0) {
    alert('خطاهای زیر را برطرف کنید:\n' + errors.join('\n'))
    return
  }
  
  isSubmitting.value = true
  
  try {
    // شبیه‌سازی ارسال به سرور
    await new Promise(resolve => setTimeout(resolve, 1000))
    
    // ایجاد کاربر جدید
    const newUser = {
      name: form.value.fullName,
      email: form.value.email,
      mobile: form.value.mobile,
      status: form.value.status,
      accessLevel: form.value.accessLevel,
      // سایر اطلاعات...
    }
    
    emit('created', newUser)
    
    // پاک کردن فرم
    form.value = {
      fullName: '',
      nationalCode: '',
      birthDate: '',
      gender: '',
      mobile: '',
      email: '',
      phone: '',
      website: '',
      province: '',
      city: '',
      address: '',
      postalCode: '',
      occupation: '',
      company: '',
      education: '',
      monthlyIncome: '',
      notifications: {
        sms: true,
        email: false,
        push: true,
        newsletter: false
      },
      password: '',
      confirmPassword: '',
      securityQuestion: '',
      securityAnswer: '',
      status: 'active',
      accessLevel: 'basic',
      expiryDate: '',
      notes: ''
    }
    
  } catch (error) {
    console.error('خطا در ثبت استفاده‌کننده:', error)
    alert('خطا در ثبت استفاده‌کننده. لطفاً دوباره تلاش کنید.')
  } finally {
    isSubmitting.value = false
  }
}
</script>

<style scoped>
/* استایل‌های خاص مودال */
</style> 

<template>
  <AdminPanelTopBar v-if="isAdmin" />
  <div class="min-h-screen bg-gray-50 py-8">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <!-- لایوت اصلی با سایدبار -->
      <div class="flex flex-col lg:flex-row gap-8 justify-center">
        <!-- محتوای اصلی -->
        <div class="flex-1 max-w-4xl">
          <!-- عنوان -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-6">
            <h1 class="text-sm font-bold text-gray-900 text-right">اطلاعات حساب کاربری</h1>
          </div>
          <!-- فرم اطلاعات شخصی -->
          <div class="bg-white rounded-lg border border-gray-200 p-6">
            <div class="mb-6">
              <h2 class="text-xl font-bold text-gray-900 text-right">اطلاعات شخصی</h2>
            </div>

            <!-- فرم ویرایش -->
            <form @submit.prevent="saveProfile">
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6 rtl">
            <!-- کد ملی -->
                <div>
                  <label class="block text-xs font-medium text-gray-700 mb-1 text-right">
                    کد ملی
                    <span v-if="verifiedFields.national_code" class="mr-1 text-xs bg-green-100 text-green-700 px-2 py-0.5 rounded">✓</span>
                    <span v-else-if="hasVerificationRequest('national_code')" class="mr-1 text-xs bg-yellow-100 text-yellow-700 px-2 py-0.5 rounded">⏳</span>
                  </label>
                  <input
                    v-model="profileForm.nationalCode"
                    :disabled="!isEditing"
                    type="text"
                    maxlength="10"
                    class="w-full px-3 py-2 border border-gray-300 rounded text-sm text-right focus:ring-2 focus:ring-blue-500 focus:border-blue-500 disabled:bg-gray-50 disabled:text-gray-600"
                    placeholder="کد ملی (10 رقم)"
                  />
              </div>

                <!-- نام و نام خانوادگی -->
                <div>
                  <label class="block text-xs font-medium text-gray-700 mb-1 text-right">
                    نام و نام خانوادگی
                  </label>
                  <input
                    v-model="profileForm.fullName"
                    :disabled="!isEditing"
                    type="text"
                    class="w-full px-3 py-2 border border-gray-300 rounded text-sm text-right focus:ring-2 focus:ring-blue-500 focus:border-blue-500 disabled:bg-gray-50 disabled:text-gray-600"
                    placeholder="نام کامل خود را وارد کنید"
                  />
              </div>

                <!-- ایمیل -->
                <div>
                  <label class="block text-xs font-medium text-gray-700 mb-1 text-right">
                    شماره موبایل
                    <span class="mr-1 text-xs bg-green-100 text-green-700 px-2 py-0.5 rounded">✓</span>
                  </label>
                  <input
                    :value="user?.mobile"
                    disabled
                    type="text"
                    class="w-full px-3 py-2 border border-gray-300 rounded text-sm text-right bg-gray-50 text-gray-600"
                  />
            </div>

                <!-- شماره شبا -->
                <div>
                  <label class="block text-xs font-medium text-gray-700 mb-1 text-right">
                    شماره شبا
                    <span v-if="verifiedFields.sheba_number" class="mr-1 text-xs bg-green-100 text-green-700 px-2 py-0.5 rounded">✓</span>
                    <span v-else-if="hasVerificationRequest('sheba_number')" class="mr-1 text-xs bg-yellow-100 text-yellow-700 px-2 py-0.5 rounded">⏳</span>
                  </label>
                  <div class="flex gap-1">
                    <span class="px-2 py-2 bg-gray-100 border border-gray-300 rounded text-sm">IR</span>
                    <input
                      v-model="profileForm.shebaNumber"
                      :disabled="!isEditing"
                      type="text"
                      maxlength="24"
                      class="flex-1 px-3 py-2 border border-gray-300 rounded text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500 disabled:bg-gray-50 disabled:text-gray-600 ltr text-right"
                      placeholder="24 رقم"
                    />
              </div>
            </div>

            <!-- ایمیل -->
                <div>
                  <label class="block text-xs font-medium text-gray-700 mb-1 text-right">
                    ایمیل
                  </label>
                  <input
                    v-model="profileForm.email"
                    :disabled="!isEditing"
                    type="email"
                    class="w-full px-3 py-2 border border-gray-300 rounded text-sm text-right focus:ring-2 focus:ring-blue-500 focus:border-blue-500 disabled:bg-gray-50 disabled:text-gray-600"
                    placeholder="ایمیل"
                  />
                </div>

                <!-- جنسیت -->
                <div>
                  <label class="block text-xs font-medium text-gray-700 mb-1 text-right">
                    تاریخ تولد
                    <span v-if="verifiedFields.birth_date" class="mr-1 text-xs bg-green-100 text-green-700 px-2 py-0.5 rounded">✓</span>
                  </label>
                  <input
                    v-model="profileForm.birthDate"
                    :disabled="!isEditing"
                    type="text"
                    placeholder="1365/06/29"
                    class="w-full px-3 py-2 border border-gray-300 rounded text-sm text-right focus:ring-2 focus:ring-blue-500 focus:border-blue-500 disabled:bg-gray-50 disabled:text-gray-600"
                  />
                </div>

                <!-- کد اقتصادی -->
                <div>
                  <label class="block text-xs font-medium text-gray-700 mb-1 text-right">
                    شغل
                  </label>
                  <input
                    v-model="profileForm.job"
                    :disabled="!isEditing"
                    type="text"
                    class="w-full px-3 py-2 border border-gray-300 rounded text-sm text-right focus:ring-2 focus:ring-blue-500 focus:border-blue-500 disabled:bg-gray-50 disabled:text-gray-600"
                    placeholder="شغل"
                  />
                </div>

                <!-- کد اقتصادی -->
                <div>
                  <label class="block text-xs font-medium text-gray-700 mb-1 text-right">
                    کد اقتصادی <span class="text-xs text-gray-500">(اختیاری)</span>
                  </label>
                  <input
                    v-model="profileForm.economicCode"
                    :disabled="!isEditing"
                    type="text"
                    class="w-full px-3 py-2 border border-gray-300 rounded text-sm text-right focus:ring-2 focus:ring-blue-500 focus:border-blue-500 disabled:bg-gray-50 disabled:text-gray-600"
                    placeholder="کد اقتصادی"
                  />
                </div>

                <!-- دکمه ویرایش اطلاعات -->
                <div v-if="!isEditing" class="flex items-end">
                  <button
                    class="w-full px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700 transition-colors text-sm font-medium"
                    @click="startEditing"
                  >
                    ویرایش اطلاعات
                </button>
              </div>

                <!-- جنسیت -->
                <div>
                  <label class="block text-xs font-medium text-gray-700 mb-1 text-right">
                    جنسیت <span class="text-xs text-gray-500">(اختیاری)</span>
                  </label>
                  <select
                    v-model="profileForm.gender"
                    :disabled="!isEditing"
                    class="w-full px-3 py-2 border border-gray-300 rounded text-sm text-right focus:ring-2 focus:ring-blue-500 focus:border-blue-500 disabled:bg-gray-50 disabled:text-gray-600"
                  >
                    <option value="">انتخاب کنید</option>
                    <option value="male">مرد</option>
                    <option value="female">زن</option>
                  </select>
              </div>
            </div>

              <!-- دکمه‌های عملیات و پیام‌ها -->
              <div v-if="isEditing" class="mt-4 space-y-3">
                <div class="flex gap-2 justify-end">
                  <button
                    type="button"
                    class="px-4 py-2 bg-gray-200 text-gray-700 rounded hover:bg-gray-300 transition-colors text-sm font-medium"
                    @click="cancelEditing"
                  >
                    انصراف
                  </button>
                  <button
                    type="submit"
                    :disabled="loading"
                    class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700 disabled:bg-gray-400 transition-colors text-sm font-medium"
                  >
                    {{ loading ? 'در حال ذخیره...' : 'ذخیره' }}
                  </button>
                </div>

                <!-- پیام خطا -->
                <div v-if="error" class="p-2 bg-red-50 border border-red-200 rounded text-red-600 text-xs">
                  {{ error }}
                </div>
              </div>
              
              <!-- پیام موفقیت -->
              <div v-if="successMessage" class="mt-4 p-2 bg-green-50 border border-green-200 rounded text-green-600 text-xs">
                {{ successMessage }}
              </div>
            </form>
          </div>

          <!-- بخش آپلود مدارک احراز هویت -->
          <div class="mt-6 bg-white rounded-lg border border-gray-200 p-6">
            <h2 class="text-lg font-bold text-gray-900 mb-4 text-right">مدارک احراز هویت</h2>
            
            <div class="space-y-4">
              <!-- آپلود کارت ملی -->
              <div class="border-2 border-dashed border-gray-300 rounded-lg p-6">
                <div class="flex items-start justify-between mb-2">
                  <div class="text-right flex-1">
                    <h3 class="text-sm font-medium text-gray-900">تصویر کارت ملی</h3>
                    <p class="text-xs text-gray-500 mt-1">تصویر واضح و خوانا از کارت ملی خود آپلود کنید</p>
                  </div>
                  <span v-if="uploadedDocs.national_card" class="text-xs bg-green-100 text-green-700 px-2 py-1 rounded">✓ آپلود شده</span>
                  <span v-else class="text-xs bg-gray-100 text-gray-600 px-2 py-1 rounded">آپلود نشده</span>
                </div>
                
                <div class="flex items-center gap-2 justify-end">
                  <input 
                    ref="nationalCardInput"
                    type="file" 
                    accept="image/*"
                    class="hidden"
                    @change="handleFileUpload($event, 'national_card')"
                  />
                  <button
                    class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700 transition-colors text-sm"
                    @click="$refs.nationalCardInput.click()"
                  >
                    {{ uploadedDocs.national_card ? 'تغییر فایل' : 'انتخاب فایل' }}
                  </button>
                  <span v-if="uploadingDoc === 'national_card'" class="text-xs text-blue-600">در حال آپلود...</span>
                </div>

                <!-- نمایش تصویر آپلود شده -->
                <div v-if="uploadedDocs.national_card" class="mt-3">
                  <img :src="uploadedDocs.national_card" class="w-32 h-32 object-cover rounded border" />
                </div>
              </div>

              <!-- آپلود سلفی با کارت ملی -->
              <div class="border-2 border-dashed border-gray-300 rounded-lg p-6">
                <div class="flex items-start justify-between mb-2">
                  <div class="text-right flex-1">
                    <h3 class="text-sm font-medium text-gray-900">سلفی با کارت ملی</h3>
                    <p class="text-xs text-gray-500 mt-1">تصویری از خود به همراه کارت ملی (برای احراز هویت)</p>
                  </div>
                  <span v-if="uploadedDocs.selfie_with_id" class="text-xs bg-green-100 text-green-700 px-2 py-1 rounded">✓ آپلود شده</span>
                  <span v-else class="text-xs bg-gray-100 text-gray-600 px-2 py-1 rounded">آپلود نشده</span>
                </div>
                
                <div class="flex items-center gap-2 justify-end">
                  <input 
                    ref="selfieInput"
                    type="file" 
                    accept="image/*"
                    class="hidden"
                    @change="handleFileUpload($event, 'selfie_with_id')"
                  />
                  <button
                    class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700 transition-colors text-sm"
                    @click="$refs.selfieInput.click()"
                  >
                    {{ uploadedDocs.selfie_with_id ? 'تغییر فایل' : 'انتخاب فایل' }}
                  </button>
                  <span v-if="uploadingDoc === 'selfie_with_id'" class="text-xs text-blue-600">در حال آپلود...</span>
                </div>

                <!-- نمایش تصویر آپلود شده -->
                <div v-if="uploadedDocs.selfie_with_id" class="mt-3">
                  <img :src="uploadedDocs.selfie_with_id" class="w-32 h-32 object-cover rounded border" />
                </div>
              </div>

              <!-- دکمه ارسال مدارک برای تایید -->
              <div v-if="canSubmitDocuments" class="pt-4">
                <button
                  :disabled="submittingDocs"
                  class="w-full px-4 py-3 bg-green-600 text-white rounded-lg hover:bg-green-700 disabled:bg-gray-400 transition-colors font-medium"
                  @click="submitDocumentsForVerification"
                >
                  {{ submittingDocs ? 'در حال ارسال...' : 'ارسال مدارک برای تایید ادمین' }}
                </button>
              </div>

              <!-- پیام موفقیت آپلود مدارک -->
              <div v-if="docsSuccessMessage" class="p-3 bg-green-50 border border-green-200 rounded text-green-600 text-sm text-right">
                {{ docsSuccessMessage }}
              </div>
            </div>
          </div>

          <!-- بخش رمز عبور -->
          <div class="bg-white rounded-lg border border-gray-200 p-6 mt-6">
            <div class="flex items-center justify-between mb-3">
              <h2 class="text-lg font-bold text-gray-900 text-right">رمز عبور</h2>
              <div class="flex items-center gap-2">
                <span v-if="hasPassword" class="bg-green-100 text-green-700 text-xs px-2 py-0.5 rounded">✓ تنظیم شده</span>
                <span v-else class="bg-yellow-100 text-yellow-700 text-xs px-2 py-0.5 rounded">تنظیم نشده</span>
                <button
                  class="text-blue-600 hover:text-blue-700 transition-colors text-xs font-medium"
                  @click="showPasswordForm = !showPasswordForm"
                >
                  {{ hasPassword ? 'تغییر' : 'تنظیم' }}
                </button>
              </div>
            </div>

            <!-- فرم رمز عبور -->
            <div v-if="showPasswordForm" class="p-3 bg-gray-50 rounded space-y-3">
                  <div v-if="hasPassword">
                <label class="block text-xs font-medium text-gray-700 mb-1">رمز عبور فعلی</label>
                    <input 
                      v-model="currentPassword" 
                      type="password" 
                  placeholder="رمز فعلی"
                  class="w-full px-3 py-2 border border-gray-300 rounded text-sm text-right focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                    />
                  </div>
                  
                  <div>
                <label class="block text-xs font-medium text-gray-700 mb-1">
                      {{ hasPassword ? 'رمز عبور جدید' : 'رمز عبور' }}
                    </label>
                    <input 
                      v-model="newPassword" 
                      type="password" 
                  placeholder="رمز جدید"
                  class="w-full px-3 py-2 border border-gray-300 rounded text-sm text-right focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                    />
                  </div>
                  
                  <div>
                <label class="block text-xs font-medium text-gray-700 mb-1">تکرار رمز عبور</label>
                    <input 
                      v-model="confirmPassword" 
                      type="password" 
                  placeholder="تکرار رمز جدید"
                  class="w-full px-3 py-2 border border-gray-300 rounded text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                    />
                  </div>
                  
              <div v-if="passwordError" class="text-red-600 text-xs">{{ passwordError }}</div>
                  
              <div class="flex gap-2">
                    <button 
                      :disabled="passwordLoading || !isPasswordFormValid"
                      class="px-3 py-2 bg-blue-600 text-white rounded hover:bg-blue-700 disabled:bg-gray-400 transition-colors text-xs font-medium"
                  @click="handleSetPassword"
                    >
                  {{ passwordLoading ? 'در حال ذخیره...' : 'ذخیره' }}
                    </button>
                    <button 
                      class="px-3 py-2 bg-gray-200 text-gray-700 rounded hover:bg-gray-300 transition-colors text-xs font-medium"
                  @click="cancelPasswordForm"
                    >
                      انصراف
                </button>
              </div>
            </div>
          </div>

          <!-- بخش اطلاعات حقوقی -->
          <div class="bg-white rounded-lg border border-gray-200 p-6 mt-6">
            <div class="flex items-center justify-between mb-3">
              <h2 class="text-lg font-bold text-red-600 text-right">اطلاعات حقوقی</h2>
              <button
                class="text-blue-600 hover:text-blue-700 font-medium text-xs transition-colors"
                @click="showLegalInfoForm = !showLegalInfoForm"
              >
                {{ showLegalInfoForm ? 'بستن' : 'ثبت' }}
              </button>
            </div>
            <p class="text-gray-600 text-xs mb-3 text-right">
              برای خرید سازمانی با فاکتور رسمی و گواهی ارزش افزوده
            </p>

            <!-- فرم اطلاعات حقوقی -->
            <div v-if="showLegalInfoForm" class="space-y-3">
              <div class="grid grid-cols-1 md:grid-cols-2 gap-3">
                <div>
                  <label class="block text-xs font-medium text-gray-700 mb-1 text-right">نام شرکت</label>
                  <input
                    v-model="legalForm.companyName"
                    type="text"
                    class="w-full px-3 py-2 border border-gray-300 rounded text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                    placeholder="نام شرکت"
                  />
                </div>

                <div>
                  <label class="block text-xs font-medium text-gray-700 mb-1 text-right">شماره ثبت</label>
                  <input
                    v-model="legalForm.registrationNumber"
                    type="text"
                    class="w-full px-3 py-2 border border-gray-300 rounded text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                    placeholder="شماره ثبت"
                  />
                </div>

                <div>
                  <label class="block text-xs font-medium text-gray-700 mb-1 text-right">شناسه ملی</label>
                  <input
                    v-model="legalForm.nationalId"
                    type="text"
                    class="w-full px-3 py-2 border border-gray-300 rounded text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                    placeholder="شناسه ملی"
                  />
                </div>

                <div>
                  <label class="block text-xs font-medium text-gray-700 mb-1 text-right">کد اقتصادی</label>
                  <input
                    v-model="legalForm.economicCode"
                    type="text"
                    class="w-full px-3 py-2 border border-gray-300 rounded text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                    placeholder="کد اقتصادی"
                  />
                </div>

                <div class="md:col-span-2">
                  <label class="block text-xs font-medium text-gray-700 mb-1 text-right">آدرس</label>
                  <textarea
                    v-model="legalForm.address"
                    rows="2"
                    class="w-full px-3 py-2 border border-gray-300 rounded text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                    placeholder="آدرس کامل"
                  ></textarea>
                </div>

                <div>
                  <label class="block text-xs font-medium text-gray-700 mb-1 text-right">شماره تماس</label>
                  <input
                    v-model="legalForm.phone"
                    type="text"
                    class="w-full px-3 py-2 border border-gray-300 rounded text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                    placeholder="شماره تماس"
                  />
                </div>
              </div>

              <button
                class="px-4 py-2 bg-red-600 text-white rounded hover:bg-red-700 transition-colors text-xs font-medium"
                @click="saveLegalInfo"
              >
                ذخیره
            </button>
            </div>
          </div>
        </div>

        <!-- سایدبار -->
        <div class="lg:w-80 flex-shrink-0 lg:ml-8">
          <AccountSidebar />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import AdminPanelTopBar from '~/components/admin/ui/AdminPanelTopBar.vue'
import AccountSidebar from '@/components/account/AccountSidebar.vue'
import { useAdmin } from '~/composables/useAdmin'

// تنظیم عنوان صفحه
definePageMeta({
  layout: 'default',
  middleware: ['auth']
})

useHead({
  title: 'اطلاعات حساب کاربری'
})

// استفاده از useAuth و useAdmin
// Auth disabled
const { isAdmin } = useAdmin()

// وضعیت ویرایش
const isEditing = ref(false)
const loading = ref(false)
const error = ref('')
const successMessage = ref('')

// مدارک آپلود شده
const uploadedDocs = ref({
  national_card: '',
  selfie_with_id: ''
})
const uploadingDoc = ref('')
const submittingDocs = ref(false)
const docsSuccessMessage = ref('')

// بررسی آماده بودن برای ارسال مدارک
const canSubmitDocuments = computed(() => {
  return uploadedDocs.value.national_card && uploadedDocs.value.selfie_with_id
})

// فرم اطلاعات شخصی
const profileForm = ref({
  fullName: '',
  nationalCode: '',
  email: '',
  shebaNumber: '',
  birthDate: '',
  gender: '',
  job: '',
  economicCode: ''
})

// فیلدهای تایید شده
const verifiedFields = ref({})
const verificationRequests = ref([])

// فرم رمز عبور
const showPasswordForm = ref(false)
const hasPassword = ref(false)
const currentPassword = ref('')
const newPassword = ref('')
const confirmPassword = ref('')
const passwordLoading = ref(false)
const passwordError = ref('')

// فرم اطلاعات حقوقی
const showLegalInfoForm = ref(false)
const legalForm = ref({
  companyName: '',
  registrationNumber: '',
  nationalId: '',
  economicCode: '',
  address: '',
  phone: ''
})

// بررسی معتبر بودن فرم رمز عبور
const isPasswordFormValid = computed(() => {
  if (!hasPassword.value) {
    return newPassword.value && 
           confirmPassword.value && 
           newPassword.value === confirmPassword.value && 
           newPassword.value.length >= 8
  }
  
  return currentPassword.value && 
         newPassword.value && 
         confirmPassword.value && 
         newPassword.value === confirmPassword.value && 
         newPassword.value.length >= 8
})

// شروع ویرایش
const startEditing = () => {
  isEditing.value = true
  error.value = ''
  successMessage.value = ''
}

// لغو ویرایش
const cancelEditing = () => {
  isEditing.value = false
  loadProfileData()
}

// بارگیری اطلاعات پروفایل
const loadProfileData = async () => {
  if (!user.value) return

  try {
    const response = await $fetch(`/api/users/profile`)
    const profileData = response.profile_data || {}

    // ترکیب نام و نام خانوادگی
    const fullName = [profileData.first_name, profileData.last_name]
      .filter(Boolean)
      .join(' ') || user.value.name || ''

    profileForm.value = {
      fullName: fullName,
      nationalCode: profileData.national_code || '',
      email: user.value.email || '',
      shebaNumber: profileData.sheba_number || '',
      birthDate: profileData.birth_date || '',
      gender: profileData.gender || '',
      job: profileData.job || '',
      economicCode: profileData.economic_code || ''
    }

    // بارگیری اطلاعات حقوقی
    legalForm.value = {
      companyName: profileData.legal_company_name || '',
      registrationNumber: profileData.legal_registration_number || '',
      nationalId: profileData.legal_national_id || '',
      economicCode: profileData.legal_economic_code || '',
      address: profileData.legal_address || '',
      phone: profileData.legal_phone || ''
    }
  } catch (err) {
    console.error('خطا در بارگیری اطلاعات:', err)
  }
}

// بارگیری فیلدهای تایید شده
const loadVerifiedFields = async () => {
  if (!user.value?.id) return

  try {
    const response = await $fetch(`/api/users/${user.value.id}/verified-fields`, {
      credentials: 'include'
    })
    verifiedFields.value = response.data || {}
  } catch (err) {
    console.error('خطا در بارگیری فیلدهای تایید شده:', err)
  }
}

// بارگیری درخواست‌های احراز هویت
const loadVerificationRequests = async () => {
  if (!user.value?.id) return

  try {
    const response = await $fetch(`/api/users/${user.value.id}/verifications`, {
      credentials: 'include'
    })
    verificationRequests.value = response.data || []
  } catch (err) {
    console.error('خطا در بارگیری درخواست‌های احراز هویت:', err)
  }
}

// بررسی وجود درخواست برای یک فیلد
const hasVerificationRequest = (fieldName) => {
  return verificationRequests.value.some(
    req => req.field_name === fieldName && req.status === 'pending'
  )
}

// ذخیره پروفایل
const saveProfile = async () => {
  loading.value = true
  error.value = ''
  successMessage.value = ''

  try {
    // تقسیم نام کامل به نام و نام خانوادگی
    const nameParts = profileForm.value.fullName.trim().split(' ')
    const firstName = nameParts[0] || ''
    const lastName = nameParts.slice(1).join(' ') || ''

    // ذخیره در profile_data کاربر
    const profileData = {
      first_name: firstName,
      last_name: lastName,
      national_code: profileForm.value.nationalCode,
      sheba_number: profileForm.value.shebaNumber,
      birth_date: profileForm.value.birthDate,
      gender: profileForm.value.gender,
      job: profileForm.value.job,
      economic_code: profileForm.value.economicCode
    }

    await $fetch(`/api/users/${user.value.id}`, {
      method: 'PUT',
      body: {
        name: profileForm.value.fullName,
        email: profileForm.value.email,
        profile_data: profileData
      },
      credentials: 'include'
    })

    // ارسال درخواست احراز هویت برای کد ملی (اگر پر شده باشد)
    if (profileForm.value.nationalCode && !verifiedFields.value.national_code) {
      await $fetch('/api/verifications', {
        method: 'POST',
        body: {
          user_id: user.value.id,
          field_name: 'national_code',
          field_value: profileForm.value.nationalCode,
          documents: [],
          metadata: {}
        },
        credentials: 'include'
      })
    }

    // ارسال درخواست احراز هویت برای شماره شبا (اگر پر شده باشد)
    if (profileForm.value.shebaNumber && !verifiedFields.value.sheba_number) {
      await $fetch('/api/verifications', {
        method: 'POST',
        body: {
          user_id: user.value.id,
          field_name: 'sheba_number',
          field_value: 'IR' + profileForm.value.shebaNumber,
          documents: [],
          metadata: {}
        },
        credentials: 'include'
      })
    }

    // ارسال درخواست احراز هویت برای تاریخ تولد (اگر پر شده باشد)
    if (profileForm.value.birthDate && !verifiedFields.value.birth_date) {
      await $fetch('/api/verifications', {
        method: 'POST',
        body: {
          user_id: user.value.id,
          field_name: 'birth_date',
          field_value: profileForm.value.birthDate,
          documents: [],
          metadata: {}
        },
        credentials: 'include'
      })
    }

    successMessage.value = 'اطلاعات با موفقیت ذخیره شد و برای تایید ادمین ارسال شد.'
    isEditing.value = false

    // بارگیری مجدد داده‌ها
    await loadVerifiedFields()
    await loadVerificationRequests()
    await fetchUser()

    setTimeout(() => {
      successMessage.value = ''
    }, 5000)
  } catch (err) {
    error.value = err.data?.message || 'خطا در ذخیره اطلاعات'
  } finally {
    loading.value = false
  }
}

// ذخیره اطلاعات حقوقی
const saveLegalInfo = async () => {
  try {
    const legalData = {
      legal_company_name: legalForm.value.companyName,
      legal_registration_number: legalForm.value.registrationNumber,
      legal_national_id: legalForm.value.nationalId,
      legal_economic_code: legalForm.value.economicCode,
      legal_address: legalForm.value.address,
      legal_phone: legalForm.value.phone
    }

    await $fetch(`/api/users/${user.value.id}`, {
      method: 'PUT',
      body: {
        profile_data: legalData
      },
      credentials: 'include'
    })

    // ارسال درخواست احراز هویت برای اطلاعات حقوقی
    await $fetch('/api/verifications', {
      method: 'POST',
      body: {
        user_id: user.value.id,
        field_name: 'legal_info',
        field_value: JSON.stringify(legalData),
        documents: [],
        metadata: legalData
      },
      credentials: 'include'
    })

    alert('اطلاعات حقوقی با موفقیت ذخیره شد و برای تایید ادمین ارسال شد.')
    showLegalInfoForm.value = false
  } catch (err) {
    alert('خطا در ذخیره اطلاعات حقوقی')
  }
}

// بررسی وضعیت رمز عبور
const checkUserPassword = async () => {
  if (user.value?.mobile) {
    try {
      const response = await $fetch('/api/auth/check-user', {
        method: 'POST',
        body: { mobile: user.value.mobile }
      })
      hasPassword.value = response.hasPassword || false
    } catch (error) {
      console.error('خطا در بررسی وضعیت رمز عبور:', error)
    }
  }
}

// تنظیم/تغییر رمز عبور
const handleSetPassword = async () => {
  passwordError.value = ''
  
  if (newPassword.value !== confirmPassword.value) {
    passwordError.value = 'رمز عبور و تکرار آن باید یکسان باشد'
    return
  }
  
  if (newPassword.value.length < 8) {
    passwordError.value = 'رمز عبور باید حداقل ۸ کاراکتر باشد'
    return
  }
  
  passwordLoading.value = true
  
  try {
    const requestBody = { 
      newPassword: newPassword.value 
    }
    
    if (hasPassword.value) {
      requestBody.currentPassword = currentPassword.value
    }
    
    await $fetch('/api/auth/set-password', {
      method: 'POST',
      body: requestBody,
      credentials: 'include'
    })
    
    hasPassword.value = true
    cancelPasswordForm()
    alert('رمز عبور با موفقیت تنظیم شد')
  } catch (error) {
    passwordError.value = error?.data?.message || 'خطا در تنظیم رمز عبور'
  } finally {
    passwordLoading.value = false
  }
}

// لغو فرم رمز عبور
const cancelPasswordForm = () => {
  showPasswordForm.value = false
  currentPassword.value = ''
  newPassword.value = ''
  confirmPassword.value = ''
  passwordError.value = ''
}

// آپلود فایل
const handleFileUpload = async (event, docType) => {
  const file = event.target.files?.[0]
  if (!file) return

  // بررسی سایز فایل (حداکثر 5MB)
  if (file.size > 5 * 1024 * 1024) {
    alert('حجم فایل نباید بیشتر از 5 مگابایت باشد')
    return
  }

  uploadingDoc.value = docType

  try {
    const formData = new FormData()
    formData.append('file', file)
    formData.append('category', 'verification')
    formData.append('purpose', 'user_verification')

    const response = await $fetch('/api/media/upload', {
      method: 'POST',
      body: formData,
      credentials: 'include'
    })

    if (response?.data?.path) {
      uploadedDocs.value[docType] = response.data.path
      docsSuccessMessage.value = 'فایل با موفقیت آپلود شد'
      setTimeout(() => {
        docsSuccessMessage.value = ''
      }, 3000)
    }
  } catch (error) {
    console.error('خطا در آپلود فایل:', error)
    alert('خطا در آپلود فایل. لطفا دوباره تلاش کنید.')
  } finally {
    uploadingDoc.value = ''
  }
}

// ارسال مدارک برای تایید
const submitDocumentsForVerification = async () => {
  if (!canSubmitDocuments.value) return

  submittingDocs.value = true
  
  try {
    // ارسال درخواست احراز هویت برای مدارک
    await $fetch('/api/verifications', {
      method: 'POST',
      body: {
        user_id: user.value.id,
        field_name: 'identity_documents',
        field_value: 'documents_uploaded',
        documents: [
          uploadedDocs.value.national_card,
          uploadedDocs.value.selfie_with_id
        ],
        metadata: {
          national_card_url: uploadedDocs.value.national_card,
          selfie_url: uploadedDocs.value.selfie_with_id
        }
      },
      credentials: 'include'
    })

    docsSuccessMessage.value = 'مدارک شما با موفقیت برای ادمین ارسال شد و در حال بررسی است.'
    
    // بارگیری مجدد درخواست‌ها
    await loadVerificationRequests()
    
  } catch (error) {
    console.error('خطا در ارسال مدارک:', error)
    alert('خطا در ارسال مدارک. لطفا دوباره تلاش کنید.')
  } finally {
    submittingDocs.value = false
  }
}

// بارگیری اطلاعات در ابتدا
onMounted(async () => {
  await fetchUser()
  await checkUserPassword()
  await loadProfileData()
  await loadVerifiedFields()
  await loadVerificationRequests()
})
</script>

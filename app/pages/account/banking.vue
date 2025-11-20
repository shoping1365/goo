<template>
  <div class="min-h-screen bg-gray-50 py-8">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <!-- لایوت اصلی با سایدبار -->
      <div class="flex flex-col lg:flex-row gap-8">
        <!-- محتوای اصلی -->
        <div class="flex-1 max-w-4xl">
          <!-- عنوان -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-6">
            <div class="flex items-center justify-between">
              <button
                class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors text-sm font-medium"
                @click="showAddForm = true"
              >
                افزودن حساب جدید
              </button>
              <h1 class="text-sm font-bold text-gray-900 text-right">اطلاعات بانکی</h1>
            </div>
          </div>

          <!-- لیست حساب‌های بانکی -->
          <div v-if="bankingAccounts.length > 0" class="space-y-4 mb-6">
            <div
              v-for="account in bankingAccounts"
              :key="account.id"
              class="bg-white rounded-lg border border-gray-200 p-6"
            >
              <div class="flex items-center justify-between mb-4">
                <div class="flex items-center gap-3">
                  <div v-if="account.is_default" class="bg-green-100 text-green-800 px-2 py-1 rounded-full text-xs font-medium">
                    پیشفرض
                  </div>
                  <div v-if="account.is_verified" class="bg-blue-100 text-blue-800 px-2 py-1 rounded-full text-xs font-medium">
                    تایید شده
                  </div>
                </div>
                <div class="flex items-center gap-2">
                  <button
                    v-if="!account.is_default"
                    class="px-3 py-1 bg-gray-100 text-gray-700 rounded text-sm hover:bg-gray-200 transition-colors"
                    @click="setAsDefault(account.id)"
                  >
                    تنظیم به عنوان پیشفرض
                  </button>
                  <button
                    class="px-3 py-1 bg-blue-100 text-blue-700 rounded text-sm hover:bg-blue-200 transition-colors"
                    @click="editAccount(account)"
                  >
                    ویرایش
                  </button>
                  <button
                    class="px-3 py-1 bg-red-100 text-red-700 rounded text-sm hover:bg-red-200 transition-colors"
                    @click="deleteAccount(account.id)"
                  >
                    حذف
                  </button>
                </div>
              </div>
              
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6 text-sm">
                <div>
                  <span class="text-gray-500">نام بانک:</span>
                  <span class="text-gray-900 font-medium">{{ account.bank_name }}</span>
                </div>
                <div>
                  <span class="text-gray-500">شماره کارت:</span>
                  <span class="text-gray-900 font-medium">{{ formatCardNumber(account.card_number) }}</span>
                </div>
                <div>
                  <span class="text-gray-500">شماره حساب:</span>
                  <span class="text-gray-900 font-medium">{{ account.account_number }}</span>
                </div>
                <div v-if="account.sheba_number">
                  <span class="text-gray-500">شماره شبا:</span>
                  <span class="text-gray-900 font-medium">{{ formatShebaNumber(account.sheba_number) }}</span>
                </div>
                <div v-if="account.account_holder_name">
                  <span class="text-gray-500">نام صاحب حساب:</span>
                  <span class="text-gray-900 font-medium">{{ account.account_holder_name }}</span>
                </div>
                <div v-if="account.account_type">
                  <span class="text-gray-500">نوع حساب:</span>
                  <span class="text-gray-900 font-medium">{{ account.account_type }}</span>
                </div>
              </div>
            </div>
          </div>

          <!-- پیام خالی -->
          <div v-else-if="!showAddForm" class="text-center py-12">
            <div class="text-gray-400 mb-4">
              <svg class="mx-auto h-12 w-12" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-medium text-gray-900 mb-2">هنوز حسابی اضافه نکرده‌اید</h3>
            <p class="text-gray-500 mb-6">برای شروع، اولین حساب بانکی خود را اضافه کنید</p>
          </div>

          <!-- فرم افزودن/ویرایش -->
          <div v-if="showAddForm" class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
            <h2 class="text-lg font-semibold text-gray-900 mb-6 text-right">
              {{ editingAccount ? 'ویرایش حساب' : 'افزودن حساب جدید' }}
            </h2>

            <!-- پیام موفقیت/خطا -->
            <div v-if="message" :class="messageType === 'success' ? 'bg-green-50 text-green-800' : 'bg-red-50 text-red-800'" class="p-3 rounded-lg mb-6 text-sm text-right">
              {{ message }}
            </div>

            <form class="space-y-6" @submit.prevent="submitForm">
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <!-- نام بانک -->
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2 text-right">نام بانک</label>
                  <input
                    v-model="bankingForm.bankName"
                    type="text"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-right"
                    placeholder="نام بانک را وارد کنید"
                    required
                  />
                </div>

                <!-- شماره کارت -->
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2 text-right">شماره کارت</label>
                  <input
                    v-model="bankingForm.cardNumber"
                    type="text"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-right"
                    placeholder="شماره کارت را وارد کنید"
                    maxlength="19"
                    required
                    @input="formatCardNumberInput"
                    @blur="detectBankFromCard"
                  />
                </div>

                <!-- شماره حساب -->
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2 text-right">شماره حساب</label>
                  <input
                    v-model="bankingForm.accountNumber"
                    type="text"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-right"
                    placeholder="شماره حساب را وارد کنید"
                    required
                  />
                </div>

                <!-- شماره شبا -->
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2 text-right">
                    شماره شبا
                    <span class="text-gray-400 text-xs">(اختیاری)</span>
                  </label>
                  <input
                    v-model="bankingForm.shebaNumber"
                    type="text"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-right"
                    placeholder="شماره شبا را وارد کنید (اختیاری)"
                    maxlength="26"
                    @input="formatShebaNumberInput"
                    @blur="detectBankFromSheba"
                  />
                </div>

                <!-- نام صاحب حساب -->
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2 text-right">نام صاحب حساب</label>
                  <input
                    v-model="bankingForm.accountHolderName"
                    type="text"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-right"
                    placeholder="نام صاحب حساب را وارد کنید"
                    required
                  />
                </div>

                <!-- نوع حساب -->
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2 text-right">نوع حساب</label>
                  <select
                    v-model="bankingForm.accountType"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-right"
                  >
                    <option value="">انتخاب کنید</option>
                    <option value="جاری">جاری</option>
                    <option value="قرض الحسنه">قرض الحسنه</option>
                    <option value="پس‌انداز">پس‌انداز</option>
                    <option value="سپرده کوتاه مدت">سپرده کوتاه مدت</option>
                    <option value="سپرده بلند مدت">سپرده بلند مدت</option>
                  </select>
                </div>
              </div>

              <!-- تنظیم به عنوان حساب پیشفرض -->
              <div class="md:col-span-2">
                <label class="flex items-center justify-end gap-2 text-sm text-gray-700">
                  <span>تنظیم به عنوان حساب پیشفرض</span>
                  <input
                    v-model="bankingForm.isDefault"
                    type="checkbox"
                    class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                  />
                </label>
              </div>

              <!-- دکمه‌های عملیات -->
              <div class="flex justify-between gap-3 mt-6">
                <div class="flex gap-3">
                  <button
                    type="button"
                    class="px-4 py-2 text-gray-700 bg-gray-100 hover:bg-gray-200 rounded-lg transition-colors text-sm font-medium"
                    @click="cancelForm"
                  >
                    انصراف
                  </button>
                  <button
                    type="submit"
                    :disabled="loading"
                    class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors text-sm font-medium disabled:opacity-50 disabled:cursor-not-allowed"
                  >
                    <span v-if="loading">در حال ذخیره...</span>
                    <span v-else>{{ editingAccount ? 'به‌روزرسانی' : 'افزودن حساب' }}</span>
                  </button>
                </div>
                <div class="flex gap-3">
                  <button
                    v-if="editingAccount && !editingAccount.is_default"
                    type="button"
                    class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors text-sm font-medium"
                    @click="setAsDefault(editingAccount.id)"
                  >
                    تنظیم به عنوان حساب پیشفرض
                  </button>
                </div>
              </div>
            </form>
          </div>
        </div>
        
        <!-- سایدبار -->
        <div class="lg:w-64 flex-shrink-0">
          <AccountSidebar />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import AccountSidebar from '@/components/account/AccountSidebar.vue'

definePageMeta({
  middleware: ['auth']
})

// متغیرهای reactive
const bankingAccounts = ref([])
const showAddForm = ref(false)
const editingAccount = ref(null)
const loading = ref(false)
const message = ref('')
const messageType = ref('')

// فرم اطلاعات بانکی
const bankingForm = ref({
  bankName: '',
  cardNumber: '',
  accountNumber: '',
  shebaNumber: '',
  accountHolderName: '',
  accountType: '',
  isDefault: false
})

// بارگذاری اطلاعات بانکی
const loadBankingAccounts = async () => {
  try {
    const response = await $fetch('/api/user/banking-info')
    if (response.success) {
      bankingAccounts.value = response.data || []
    }
  } catch (error) {
    console.error('خطا در بارگذاری اطلاعات بانکی:', error)
  }
}

// تشخیص بانک از شماره کارت
const detectBankFromCard = async () => {
  const cardNumber = bankingForm.value.cardNumber.replace(/\D/g, '')
  if (cardNumber.length >= 16) {
    try {
      const response = await $fetch('/api/user/banking-info/auto-complete/card', {
        method: 'POST',
        body: {
          card_number: cardNumber
        }
      })
      
      if (response.success && response.data) {
        if (response.data.bank_name) {
          bankingForm.value.bankName = response.data.bank_name
        }
        if (response.data.sheba_number && !bankingForm.value.shebaNumber) {
          bankingForm.value.shebaNumber = response.data.sheba_number
        }
        
        message.value = 'نام بانک تشخیص داده شد'
        messageType.value = 'success'
        
        setTimeout(() => {
          message.value = ''
        }, 3000)
      }
    } catch (error) {
      console.error('خطا در تکمیل خودکار:', error)
    }
  }
}

// تشخیص بانک از شماره شبا
const detectBankFromSheba = async () => {
  const shebaNumber = bankingForm.value.shebaNumber.replace(/\s/g, '').toUpperCase()
  if (shebaNumber.length >= 26) {
    try {
      const response = await $fetch('/api/user/banking-info/auto-complete/sheba', {
        method: 'POST',
        body: {
          sheba_number: shebaNumber
        }
      })
      
      if (response.success && response.data) {
        if (response.data.bank_name) {
          bankingForm.value.bankName = response.data.bank_name
        }
        if (response.data.card_number && !bankingForm.value.cardNumber) {
          bankingForm.value.cardNumber = response.data.card_number
        }
        
        message.value = 'نام بانک تشخیص داده شد'
        messageType.value = 'success'
        
        setTimeout(() => {
          message.value = ''
        }, 3000)
      }
    } catch (error) {
      console.error('خطا در تکمیل خودکار:', error)
    }
  }
}

// فرمت کردن شماره کارت
const formatCardNumberInput = (event) => {
  let value = event.target.value.replace(/\D/g, '')
  value = value.replace(/(\d{4})(?=\d)/g, '$1-')
  bankingForm.value.cardNumber = value
}

// فرمت کردن شماره شبا
const formatShebaNumberInput = (event) => {
  let value = event.target.value.replace(/\s/g, '').toUpperCase()
  if (value.length > 2 && !value.startsWith('IR')) {
    value = 'IR' + value.substring(2)
  }
  bankingForm.value.shebaNumber = value
}

// فرمت کردن شماره کارت برای نمایش
const formatCardNumber = (cardNumber) => {
  if (!cardNumber) return ''
  return cardNumber.replace(/(\d{4})(?=\d)/g, '$1-')
}

// فرمت کردن شماره شبا برای نمایش
const formatShebaNumber = (shebaNumber) => {
  if (!shebaNumber) return ''
  return shebaNumber.replace(/(.{4})/g, '$1 ').trim()
}

// ارسال فرم
const submitForm = async () => {
  loading.value = true
  message.value = ''

  try {
    const url = editingAccount.value 
      ? `/api/user/banking-info/${editingAccount.value.id}` 
      : '/api/user/banking-info'
    
    const method = editingAccount.value ? 'PUT' : 'POST'
    
    const response = await $fetch(url, {
      method,
      body: bankingForm.value
    })

    if (response.success) {
      message.value = editingAccount.value ? 'حساب با موفقیت به‌روزرسانی شد' : 'حساب جدید با موفقیت اضافه شد'
      messageType.value = 'success'
      
      await loadBankingAccounts()
      resetForm()
      
      setTimeout(() => {
        message.value = ''
      }, 3000)
    }
  } catch (error) {
    message.value = error.data?.message || 'خطا در ذخیره اطلاعات'
    messageType.value = 'error'
  } finally {
    loading.value = false
  }
}

// ویرایش حساب
const editAccount = (account) => {
  editingAccount.value = account
  bankingForm.value = {
    bankName: account.bank_name,
    cardNumber: account.card_number,
    accountNumber: account.account_number,
    shebaNumber: account.sheba_number,
    accountHolderName: account.account_holder_name || '',
    accountType: account.account_type || '',
    isDefault: account.is_default
  }
  showAddForm.value = true
}

// حذف حساب
const deleteAccount = async (accountId) => {
  if (confirm('آیا از حذف این حساب مطمئن هستید؟')) {
    try {
      const response = await $fetch(`/api/user/banking-info/${accountId}`, {
        method: 'DELETE'
      })

      if (response.success) {
        message.value = 'حساب با موفقیت حذف شد'
        messageType.value = 'success'
        await loadBankingAccounts()
        
        setTimeout(() => {
          message.value = ''
        }, 3000)
      }
    } catch (error) {
      message.value = error.data?.message || 'خطا در حذف حساب'
      messageType.value = 'error'
    }
  }
}

// تنظیم به عنوان پیشفرض
const setAsDefault = async (accountId) => {
  try {
    const response = await $fetch(`/api/user/banking-info/${accountId}/set-default`, {
      method: 'POST'
    })

    if (response.success) {
      message.value = 'حساب پیشفرض با موفقیت تغییر کرد'
      messageType.value = 'success'
      await loadBankingAccounts()
      
      setTimeout(() => {
        message.value = ''
      }, 3000)
    }
  } catch (error) {
    message.value = error.data?.message || 'خطا در تغییر حساب پیشفرض'
    messageType.value = 'error'
  }
}

// لغو فرم
const cancelForm = () => {
  resetForm()
  showAddForm.value = false
}

// ریست کردن فرم
const resetForm = () => {
  bankingForm.value = {
    bankName: '',
    cardNumber: '',
    accountNumber: '',
    shebaNumber: '',
    accountHolderName: '',
    accountType: '',
    isDefault: false
  }
  editingAccount.value = null
  message.value = ''
}

// بارگذاری اولیه
onMounted(() => {
  loadBankingAccounts()
})
</script>
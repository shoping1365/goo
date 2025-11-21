<template>
  <div v-if="show" class="fixed inset-0 bg-black bg-opacity-50 overflow-y-auto h-full w-full z-50 flex items-center justify-center p-6">
    <div class="relative w-full max-w-4xl bg-white rounded-2xl shadow-2xl overflow-hidden max-h-[90vh] overflow-y-auto">
      <!-- Header -->
      <div class="bg-gradient-to-r from-purple-600 to-blue-600 px-8 py-6 text-white">
        <div class="flex items-center justify-between">
          <div class="flex items-center space-x-3 space-x-reverse">
            <div class="w-10 h-10 bg-white bg-opacity-20 rounded-xl flex items-center justify-center">
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
              </svg>
            </div>
            <div>
              <h3 class="text-2xl font-bold">{{ props.editingPattern ? 'ویرایش پترن' : 'ایجاد پترن جدید' }}</h3>
              <p class="text-purple-100 text-sm mt-1">تعریف الگوی جدید برای ارسال پیامک</p>
            </div>
          </div>
          <button 
            class="w-8 h-8 bg-white bg-opacity-20 rounded-lg flex items-center justify-center hover:bg-opacity-30 transition-all duration-200" 
            @click="$emit('close')"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>
      </div>

      <!-- Form Content -->
      <div class="p-6">
        <form class="space-y-6" @submit.prevent="savePattern">
          <!-- Basic Information Section -->
          <div class="bg-gradient-to-br from-blue-50 to-indigo-50 rounded-xl p-6 border border-blue-100">
            <div class="flex items-center space-x-2 space-x-reverse mb-3">
              <div class="w-5 h-5 bg-blue-500 rounded-lg flex items-center justify-center">
                <svg class="w-3 h-3 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
              </div>
              <h4 class="text-base font-semibold text-gray-800">اطلاعات پایه</h4>
            </div>
            
            <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
              <div>
                <label class="block text-sm font-semibold text-gray-700 mb-2">شناسه پترن</label>
                <input 
                  v-model="form.fixedId" 
                  type="number" 
                  :disabled="!!props.editingPattern"
                  class="block w-full px-3 py-2 rounded-lg border border-gray-200 shadow-sm focus:border-blue-400 focus:ring-2 focus:ring-blue-200 focus:ring-opacity-50 transition-all duration-200"
                  placeholder="شناسه خودکار"
                >
                <p class="mt-1 text-xs text-gray-500">شناسه یکتا پترن (در ویرایش غیرقابل تغییر)</p>
              </div>
              <div>
                <label class="block text-sm font-semibold text-gray-700 mb-2">نام پترن</label>
              <input 
                v-model="form.name" 
                type="text" 
                required 
                  class="block w-full px-3 py-2 rounded-lg border border-gray-200 shadow-sm focus:border-blue-400 focus:ring-2 focus:ring-blue-200 focus:ring-opacity-50 transition-all duration-200"
                  placeholder="مثال: تایید سفارش"
              >
            </div>
            <div>
                <label class="block text-sm font-semibold text-gray-700 mb-2">نوع</label>
              <select 
                v-model="form.type" 
                required 
                  class="block w-full px-3 py-2 rounded-lg border border-gray-200 shadow-sm focus:border-blue-400 focus:ring-2 focus:ring-blue-200 focus:ring-opacity-50 transition-all duration-200"
              >
                <option value="sms">پیامک</option>
                <option value="email">ایمیل</option>
                <option value="push">اعلان مرورگر</option>
              </select>
            </div>
            <div>
                <label class="block text-sm font-semibold text-gray-700 mb-2">درگاه پیامکی</label>
              <select 
                v-model="form.gatewayId" 
                required 
                  class="block w-full px-3 py-2 rounded-lg border border-gray-200 shadow-sm focus:border-blue-400 focus:ring-2 focus:ring-blue-200 focus:ring-opacity-50 transition-all duration-200"
              >
                <option value="1">IPPanel</option>
                <option value="2">کاوه‌نگار</option>
                <option value="3">ملی پیامک</option>
                <option value="4">نوین</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-2">دریافت‌کننده</label>
              <select v-model="form.scope" class="block w-full px-3 py-2 rounded-lg border border-gray-200 shadow-sm focus:border-blue-400 focus:ring-2 focus:ring-blue-200">
                <option value="customer">مشتری</option>
                <option value="manager">مدیر</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-2">هدف/آیتم</label>
              <select v-model="form.feature" class="block w-full px-3 py-2 rounded-lg border border-gray-200 shadow-sm focus:border-blue-400 focus:ring-2 focus:ring-blue-200" @change="applyDefaults">
                <optgroup v-if="form.scope === 'customer'" label="مشتری">
                  <option value="auth_otp">کد تأیید ورود</option>
                  <option value="order_confirmation">تایید سفارش</option>
                  <option value="order_shipped">ارسال سفارش</option>
                </optgroup>
                <optgroup v-if="form.scope === 'manager'" label="مدیر">
                  <option value="admin_failover">اعلان خطای سیستم</option>
                  <option value="low_stock">کمبود موجودی</option>
                  <option value="admin_order">اعلان سفارشات به مدیر</option>
                  <option value="security_issue">مشکلات امنیتی</option>
                  <option value="gateway_test">تست ارسال درگاه</option>
                </optgroup>
              </select>
              <p v-if="helperText" class="text-xs text-gray-500 mt-1">{{ helperText }}</p>
            </div>
          </div>
          
            <div class="mt-4">
              <label class="block text-sm font-semibold text-gray-700 mb-2">توضیحات</label>
            <textarea 
              v-model="form.description" 
              rows="2" 
                class="block w-full px-3 py-2 rounded-lg border border-gray-200 shadow-sm focus:border-blue-400 focus:ring-2 focus:ring-blue-200 focus:ring-opacity-50 transition-all duration-200"
                placeholder="توضیحات مختصری درباره این پترن..."
            ></textarea>
            </div>
          </div>
          
          <!-- Pattern Configuration Section -->
          <div class="bg-gradient-to-br from-green-50 to-emerald-50 rounded-xl p-6 border border-green-100">
            <div class="flex items-center space-x-2 space-x-reverse mb-3">
              <div class="w-5 h-5 bg-green-500 rounded-lg flex items-center justify-center">
                <svg class="w-3 h-3 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"></path>
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                </svg>
              </div>
              <h4 class="text-base font-semibold text-gray-800">تنظیمات الگو</h4>
            </div>
            
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
                <label class="block text-sm font-semibold text-gray-700 mb-2">کد الگو</label>
            <input 
              v-model="form.patternCode" 
              type="text" 
              required 
                  class="block w-full px-3 py-2 rounded-lg border border-gray-200 shadow-sm focus:border-green-400 focus:ring-2 focus:ring-green-200 focus:ring-opacity-50 transition-all duration-200"
              placeholder="مثال: order_confirmation"
            >
                <p class="mt-1 text-sm text-gray-600 flex items-center">
                  <svg class="w-3 h-3 ml-1 text-green-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                  </svg>
                  کد یکتا برای شناسایی این الگو
                </p>
          </div>
          <div>
                <label class="block text-sm font-semibold text-gray-700 mb-2">وضعیت</label>
              <select 
                v-model="form.status" 
                required 
                  class="block w-full px-3 py-2 rounded-lg border border-gray-200 shadow-sm focus:border-green-400 focus:ring-2 focus:ring-green-200 focus:ring-opacity-50 transition-all duration-200"
              >
                <option value="active">فعال</option>
                <option value="inactive">غیرفعال</option>
                <option value="draft">پیش‌نویس</option>
              </select>
            </div>
            </div>
            
            <div class="mt-4">
              <label class="block text-sm font-semibold text-gray-700 mb-2">متغیرهای الگو</label>
              <textarea 
                v-model="form.variables" 
                rows="2" 
                class="block w-full px-3 py-2 rounded-lg border border-gray-200 shadow-sm focus:border-green-400 focus:ring-2 focus:ring-green-200 focus:ring-opacity-50 transition-all duration-200"
                :placeholder="variablesPlaceholder"
              ></textarea>
              <p class="mt-1 text-sm text-gray-600 flex items-center">
                <svg class="w-3 h-3 ml-1 text-green-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
                متغیرهای قابل استفاده در الگو (با کاما جدا کنید)
              </p>
            </div>
          </div>

          <!-- Message Template Section -->
          <div class="bg-gradient-to-br from-purple-50 to-pink-50 rounded-xl p-6 border border-purple-100">
            <div class="flex items-center space-x-2 space-x-reverse mb-3">
              <div class="w-5 h-5 bg-purple-500 rounded-lg flex items-center justify-center">
                <svg class="w-3 h-3 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"></path>
                </svg>
              </div>
              <h4 class="text-base font-semibold text-gray-800">الگوی پیام</h4>
            </div>
            
            <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
              <div class="md:col-span-2">
                <label class="block text-sm font-semibold text-gray-700 mb-2">متن پیام</label>
                <textarea 
                  v-model="form.content" 
                  rows="4" 
                  required 
                  class="block w-full px-3 py-2 rounded-lg border border-gray-200 shadow-sm focus:border-purple-400 focus:ring-2 focus:ring-purple-200 focus:ring-opacity-50 transition-all duration-200" 
                  placeholder="مثال: سلام {name}، سفارش شما با شماره {order_id} ثبت شد."
                ></textarea>
                <p class="mt-1 text-sm text-gray-600 flex items-center">
                  <svg class="w-3 h-3 ml-1 text-purple-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                  </svg>
                  از متغیرهای تعریف‌شده در بالا استفاده کنید
                </p>
              </div>
            <div>
                <label class="block text-sm font-semibold text-gray-700 mb-2">حداکثر طول</label>
              <input 
                v-model="form.maxLength" 
                type="number" 
                  class="block w-full px-3 py-2 rounded-lg border border-gray-200 shadow-sm focus:border-purple-400 focus:ring-2 focus:ring-purple-200 focus:ring-opacity-50 transition-all duration-200"
                placeholder="160"
              >
                <div class="mt-3 p-2 bg-white rounded-lg border border-gray-200">
                  <div class="text-xs font-medium text-gray-700 mb-1">پیش‌نمایش:</div>
                  <div class="text-xs text-gray-600 bg-gray-50 p-2 rounded border">
                    {{ form.content || 'متن پیام اینجا نمایش داده می‌شود...' }}
                  </div>
                  <div class="text-xs text-gray-500 mt-1">
                    طول: {{ form.content.length }} / {{ form.maxLength }}
                  </div>
                </div>
              </div>
            </div>
          </div>
          
          <!-- Action Buttons -->
          <div class="flex items-center justify-end space-x-4 space-x-reverse pt-4 border-t border-gray-200">
            <button 
              type="button" 
              class="px-4 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200 transition-all duration-200 font-medium"
              @click="$emit('close')"
            >
              انصراف
            </button>
            <button 
              type="submit"
              class="px-6 py-2 bg-gradient-to-r from-purple-600 to-blue-600 text-white rounded-lg hover:from-purple-700 hover:to-blue-700 transition-all duration-200 font-medium shadow-lg hover:shadow-xl transform hover:scale-105"
            >
              <span class="flex items-center">
                <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
                </svg>
                {{ props.editingPattern ? 'ویرایش پترن' : 'ایجاد پترن' }}
              </span>
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch, computed } from 'vue'

// تعریف interface ها
interface Pattern {
  id: number
  name: string
  description: string
  type: 'sms' | 'email' | 'push'
  category: 'order' | 'marketing' | 'notification' | 'verification'
  scope: 'customer' | 'manager'
  feature: string
  content: string
  status: 'active' | 'inactive' | 'draft'
  usageCount: number
  maxLength?: number
  gatewayId: number
  gatewayName: string
  patternCode: string
  variables: string
}

interface PatternForm {
  fixedId?: number
  name: string
  description: string
  type: 'sms' | 'email' | 'push'
  scope?: 'manager' | 'customer'
  feature?: string
  content: string
  status: 'active' | 'inactive' | 'draft'
  maxLength?: number
  gatewayId: number
  gatewayName: string
  patternCode: string
  variables: string
  message_template?: string // اضافه کردن فیلد message_template
}

// Props
interface Props {
  show: boolean
  editingPattern?: Pattern | null
}

const props = defineProps<Props>()

// Events
const emit = defineEmits<{
  close: []
  save: [pattern: PatternForm]
}>()

// Reactive form data
const form = ref<PatternForm>({
  fixedId: undefined,
  name: '',
  description: '',
  type: 'sms',
  scope: 'customer',
  feature: '',
  content: '',
  status: 'active',
  maxLength: 160,
  gatewayId: 1,
  gatewayName: 'IPPanel',
  patternCode: '',
  variables: '',
  message_template: ''
})

// Methods
const savePattern = () => {

  // تعیین نام درگاه بر اساس ID
  const gatewayNames: { [key: number]: string } = {
    1: 'IPPanel',
    2: 'کاوه‌نگار',
    3: 'ملی پیامک',
    4: 'نوین'
  }
  
  const formData = {
    ...form.value,
    gatewayName: gatewayNames[form.value.gatewayId] || 'IPPanel',
    message_template: form.value.content // تبدیل content به message_template
  }

  emit('save', formData)
}

const resetForm = () => {
  form.value = {
    fixedId: undefined,
    name: '',
    description: '',
    type: 'sms',
    scope: 'customer',
    feature: 'auth_otp',
    content: '',
    status: 'active',
    maxLength: 160,
    gatewayId: 1,
    gatewayName: 'IPPanel',
    patternCode: '',
    variables: '',
    message_template: ''
  }
}

// Watch for editing pattern changes
watch(() => props.editingPattern, (newPattern) => {
  if (newPattern) {
    form.value = {
      fixedId: newPattern.id,
      name: newPattern.name,
      description: newPattern.description,
      type: newPattern.type,
      scope: newPattern.scope || 'customer',
      feature: newPattern.feature || '',
      content: newPattern.content,
      status: newPattern.status,
      maxLength: newPattern.maxLength,
      gatewayId: newPattern.gatewayId,
      gatewayName: newPattern.gatewayName,
      patternCode: newPattern.patternCode,
      variables: newPattern.variables,
      message_template: newPattern.content // استفاده از content برای message_template
    }
  } else {
    resetForm()
  }
}, { immediate: true })

// -- اضافه کردن واچر برای تغییر دریافت‌کننده --
watch(() => form.value.scope, (newScope, oldScope) => {
  if (newScope !== oldScope && !props.editingPattern) {
    // فقط در حالت ایجاد پترن جدید، وقتی دریافت‌کننده تغییر می‌کند، آیتم هدف را به یک گزینه پیش‌فرض تغییر می‌دهیم
    if (newScope === 'customer') {
      form.value.feature = 'auth_otp';
    } else if (newScope === 'manager') {
      form.value.feature = 'admin_failover';
    } else {
      form.value.feature = '';
    }
  }
});

// Defaults per feature
const variablesPlaceholder = computed(() => 'name, order_id, ...')
const helperText = computed(() => {
  switch (form.value.feature) {
    case 'gateway_test':
      return 'پترن تست برای بررسی عملکرد درگاه‌های پیامکی - متغیرهای gateway و date به صورت خودکار تنظیم می‌شوند'
    case 'auth_otp':
      return 'پترن ارسال کد تایید - متغیر code به صورت خودکار تنظیم می‌شود'
    case 'order_confirmation':
      return 'پترن تایید سفارش - متغیرهای name و order_id به صورت خودکار تنظیم می‌شوند'
    case 'order_shipped':
      return 'پترن اطلاع‌رسانی ارسال سفارش - متغیرهای name و order_id به صورت خودکار تنظیم می‌شوند'
    case 'admin_failover':
      return 'پترن اعلان خطای سیستم - متغیر error_message به صورت خودکار تنظیم می‌شود'
    case 'low_stock':
      return 'پترن کمبود موجودی - متغیر product_name به صورت خودکار تنظیم می‌شود'
    case 'admin_order':
      return 'پترن اعلان سفارشات به مدیر - متغیرهای order_id و amount به صورت خودکار تنظیم می‌شوند'
    case 'security_issue':
      return 'پترن مشکلات امنیتی - متغیر issue_description به صورت خودکار تنظیم می‌شود'
    default:
      return ''
  }
})

const applyDefaults = () => {
  // فقط در حالت ایجاد پترن جدید، تنظیم مقادیر پیش‌فرض بر اساس feature انتخاب شده
  if (props.editingPattern) {
    return // در حالت ویرایش، مقادیر پیش‌فرض اعمال نکن
  }
  
  // تنظیم مقادیر پیش‌فرض بر اساس feature انتخاب شده
  switch (form.value.feature) {
    case 'gateway_test':
      // تنظیمات پیش‌فرض برای تست ارسال درگاه
      form.value.content = 'مدیر محترم\nاین یک تست ارسال {gateway} در تاریخ {date} میباشد'
      form.value.variables = 'gateway,date'
      form.value.name = 'تست ارسال درگاه'
      form.value.description = 'پترن تست برای بررسی عملکرد درگاه‌های پیامکی'
      form.value.patternCode = 'gateway_test'
      break
    case 'auth_otp':
      form.value.content = 'کد تایید شما: {code}'
      form.value.variables = 'code'
      break
    case 'order_confirmation':
      form.value.content = 'سلام {name}، سفارش شما با شماره {order_id} ثبت شد.'
      form.value.variables = 'name,order_id'
      break
    case 'order_shipped':
      form.value.content = 'سلام {name}، سفارش شما با شماره {order_id} ارسال شد.'
      form.value.variables = 'name,order_id'
      break
    case 'admin_failover':
      form.value.content = 'خطای سیستم: {error_message}'
      form.value.variables = 'error_message'
      break
    case 'low_stock':
      form.value.content = 'کمبود موجودی: محصول {product_name}'
      form.value.variables = 'product_name'
      break
    case 'admin_order':
      form.value.content = 'سفارش جدید: {order_id} - مبلغ: {amount}'
      form.value.variables = 'order_id,amount'
      break
    case 'security_issue':
      form.value.content = 'مشکل امنیتی: {issue_description}'
      form.value.variables = 'issue_description'
      break
    default:
      // برای سایر موارد، مقادیر را پاک نکنیم
      break
  }
}
</script> 
<template>
  <div class="footer-custom-widget" :style="widgetStyle">
    <div class="text-center">
      <!-- عنوان بخش -->
      <h4 v-if="title" class="text-white font-semibold mb-4 text-lg">{{ title }}</h4>
      
      <!-- محتوای سفارشی -->
      <div v-if="content" class="custom-content mb-4">
        <!-- HTML content -->
        <!-- 
          ⚠️ امنیت XSS: استفاده از v-html خطرناک است!
          
          این کد محتوای HTML را بدون sanitization نمایش می‌دهد که می‌تواند منجر به حملات XSS شود.
          
          ✅ راه حل صحیح:
          1. قبل از استفاده از v-html، محتوا را با کتابخانه sanitization (مثل DOMPurify) پاکسازی کنید
          2. یا از {{ }} به جای v-html استفاده کنید اگر HTML نیاز نیست
          3. محتوای کاربر را هرگز بدون sanitization در v-html قرار ندهید
          
          مثال صحیح:
          import DOMPurify from 'dompurify'
          const sanitizedContent = computed(() => DOMPurify.sanitize(props.content))
          <div v-html="sanitizedContent"></div>
        -->
        <div v-if="isHtml" class="text-white text-opacity-90" v-html="content"></div>
        
        <!-- متن ساده -->
        <div v-else class="text-white text-opacity-90 text-sm leading-relaxed whitespace-pre-line">
          {{ content }}
        </div>
      </div>
      
      <!-- تصویر -->
      <div v-if="imageUrl" class="custom-image mb-4">
        <img
          :src="imageUrl"
          :alt="imageAlt || 'تصویر سفارشی'"
          class="w-full h-auto rounded-lg shadow-lg"
          :class="imageClasses"
        />
      </div>
      
      <!-- دکمه‌ها -->
      <div v-if="buttons && buttons.length > 0" class="custom-buttons mb-4">
        <div class="flex flex-wrap justify-center gap-3">
          <button
            v-for="button in buttons"
            :key="button.id"
            :class="[
              'px-4 py-2 rounded-lg font-medium text-sm transition-all duration-200',
              button.variant === 'primary' ? 'bg-blue-600 hover:bg-blue-700 text-white' : '',
              button.variant === 'secondary' ? 'bg-white bg-opacity-20 hover:bg-opacity-30 text-white border border-white border-opacity-30' : '',
              button.variant === 'outline' ? 'bg-transparent hover:bg-white hover:text-gray-900 text-white border border-white border-opacity-50' : '',
              button.variant === 'ghost' ? 'bg-transparent hover:bg-white hover:text-gray-900 text-white' : '',
              button.size === 'small' ? 'px-3 py-1.5 text-xs' : '',
              button.size === 'large' ? 'px-6 py-3 text-base' : ''
            ]"
            :disabled="button.disabled"
            @click="handleButtonClick(button)"
          >
            <span v-if="button.icon" class="mr-2">{{ button.icon }}</span>
            <span>{{ button.text }}</span>
          </button>
        </div>
      </div>
      
      <!-- لیست -->
      <div v-if="list && list.length > 0" class="custom-list mb-4">
        <ul class="text-white text-opacity-80 text-sm space-y-2">
          <li
            v-for="item in list"
            :key="item.id"
            class="flex items-center justify-center space-x-2 space-x-reverse"
          >
            <span v-if="item.icon" class="text-lg">{{ item.icon }}</span>
            <span v-if="item.bullet" class="w-2 h-2 bg-white rounded-full opacity-60"></span>
            <span>{{ item.text }}</span>
          </li>
        </ul>
      </div>
      
      <!-- جدول -->
      <div v-if="table && table.length > 0" class="custom-table mb-4">
        <div class="overflow-x-auto">
          <table class="w-full text-white text-opacity-80 text-sm">
            <thead>
              <tr class="border-b border-white border-opacity-20">
                <th
                  v-for="header in tableHeaders"
                  :key="header"
                  class="px-3 py-2 text-center font-medium"
                >
                  {{ header }}
                </th>
              </tr>
            </thead>
            <tbody>
              <tr
                v-for="(row, rowIndex) in table"
                :key="rowIndex"
                class="border-b border-white border-opacity-10"
              >
                <td
                  v-for="(cell, cellIndex) in row"
                  :key="cellIndex"
                  class="px-3 py-2 text-center"
                >
                  {{ cell }}
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
      
      <!-- فرم -->
      <form v-if="form" class="custom-form mb-4" @submit.prevent="handleFormSubmit">
        <div class="space-y-3">
          <div
            v-for="field in form.fields"
            :key="field.id"
            class="text-right"
          >
            <label v-if="field.label" :for="field.id" class="block text-white text-opacity-80 text-sm mb-1">
              {{ field.label }}
              <span v-if="field.required" class="text-red-400">*</span>
            </label>
            
            <!-- فیلد متن -->
            <input
              v-if="field.type === 'text' || field.type === 'email' || field.type === 'tel'"
              :id="field.id"
              v-model="formData[field.id]"
              :type="field.type"
              :placeholder="field.placeholder"
              :required="field.required"
              class="w-full px-3 py-2 bg-white bg-opacity-20 border border-white border-opacity-30 rounded-lg text-white placeholder-white placeholder-opacity-60 focus:outline-none focus:ring-2 focus:ring-white focus:ring-opacity-50 focus:border-white transition-all duration-200"
            />
            
            <!-- فیلد textarea -->
            <textarea
              v-else-if="field.type === 'textarea'"
              :id="field.id"
              v-model="formData[field.id]"
              :placeholder="field.placeholder"
              :required="field.required"
              :rows="field.rows || 3"
              class="w-full px-3 py-2 bg-white bg-opacity-20 border border-white border-opacity-30 rounded-lg text-white placeholder-white placeholder-opacity-60 focus:outline-none focus:ring-2 focus:ring-white focus:ring-opacity-50 focus:border-white transition-all duration-200 resize-none"
            ></textarea>
            
            <!-- فیلد select -->
            <select
              v-else-if="field.type === 'select'"
              :id="field.id"
              v-model="formData[field.id]"
              :required="field.required"
              class="w-full px-3 py-2 bg-white bg-opacity-20 border border-white border-opacity-30 rounded-lg text-white focus:outline-none focus:ring-2 focus:ring-white focus:ring-opacity-50 focus:border-white transition-all duration-200"
            >
              <option value="" disabled selected>{{ field.placeholder || 'انتخاب کنید' }}</option>
              <option
                v-for="option in field.options"
                :key="option.value"
                :value="option.value"
              >
                {{ option.label }}
              </option>
            </select>
            
            <!-- فیلد checkbox -->
            <div v-else-if="field.type === 'checkbox'" class="flex items-center space-x-2 space-x-reverse">
              <input
                :id="field.id"
                v-model="formData[field.id]"
                type="checkbox"
                :required="field.required"
                class="w-4 h-4 text-blue-600 bg-white bg-opacity-20 border-white border-opacity-30 rounded focus:ring-white focus:ring-opacity-50"
              />
              <label :for="field.id" class="text-white text-opacity-80 text-sm">
                {{ field.label }}
              </label>
            </div>
            
            <!-- فیلد radio -->
            <div v-else-if="field.type === 'radio'" class="space-y-2">
              <div
                v-for="option in field.options"
                :key="option.value"
                class="flex items-center space-x-2 space-x-reverse"
              >
                <input
                  :id="`${field.id}-${option.value}`"
                  v-model="formData[field.id]"
                  type="radio"
                  :value="option.value"
                  :required="field.required"
                  class="w-4 h-4 text-blue-600 bg-white bg-opacity-20 border-white border-opacity-30 focus:ring-white focus:ring-opacity-50"
                />
                <label :for="`${field.id}-${option.value}`" class="text-white text-opacity-80 text-sm">
                  {{ option.label }}
                </label>
              </div>
            </div>
          </div>
          
          <!-- دکمه ارسال -->
          <button
            type="submit"
            :disabled="isSubmitting"
            class="w-full px-4 py-2 bg-blue-600 hover:bg-blue-700 disabled:bg-gray-500 text-white font-medium rounded-lg transition-all duration-200"
          >
            <span v-if="!isSubmitting">{{ form.submitText || 'ارسال' }}</span>
            <span v-else class="flex items-center justify-center space-x-2 space-x-reverse">
              <div class="w-4 h-4 border-2 border-white border-t-transparent rounded-full animate-spin"></div>
              <span>در حال ارسال...</span>
            </span>
          </button>
        </div>
      </form>
      
      <!-- پیام‌ها -->
      <div v-if="successMessage" class="bg-green-500 bg-opacity-20 border border-green-400 rounded-lg p-3 text-green-300 text-sm mb-4">
        <div class="flex items-center justify-center space-x-2 space-x-reverse">
          <span>✅</span>
          <span>{{ successMessage }}</span>
        </div>
      </div>
      
      <div v-if="errorMessage" class="bg-red-500 bg-opacity-20 border border-red-400 rounded-lg p-3 text-red-300 text-sm mb-4">
        <div class="flex items-center justify-center space-x-2 space-x-reverse">
          <span>❌</span>
          <span>{{ errorMessage }}</span>
        </div>
      </div>
      
      <!-- محتوای خالی -->
      <div v-if="!hasContent" class="text-white text-opacity-60 text-sm">
        <div class="bg-white bg-opacity-10 rounded-lg p-6">
          <div class="text-4xl mb-3">🔧</div>
          <p>محتوای سفارشی تعریف نشده است</p>
          <p class="text-xs mt-2">لطفاً محتوا، تصویر یا فرم مورد نظر را تنظیم کنید</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'

interface Button {
  id: string
  text: string
  icon?: string
  variant?: 'primary' | 'secondary' | 'outline' | 'ghost'
  size?: 'small' | 'medium' | 'large'
  disabled?: boolean
  action?: string
}

interface ListItem {
  id: string
  text: string
  icon?: string
  bullet?: boolean
}

interface FormField {
  id: string
  type: 'text' | 'email' | 'tel' | 'textarea' | 'select' | 'checkbox' | 'radio'
  label?: string
  placeholder?: string
  required?: boolean
  options?: Array<{ value: string; label: string }>
  rows?: number
}

interface Form {
  fields: FormField[]
  submitText?: string
  action?: string
}

interface Props {
  title?: string
  content?: string
  isHtml?: boolean
  imageUrl?: string
  imageAlt?: string
  imageClasses?: string
  buttons?: Button[]
  list?: ListItem[]
  table?: string[][]
  tableHeaders?: string[]
  form?: Form
  paddingRight?: number
  paddingLeft?: number
  align?: 'left' | 'center' | 'right'
}

const props = withDefaults(defineProps<Props>(), {
  title: 'محتوای سفارشی',
  content: '',
  isHtml: false,
  imageUrl: '',
  imageAlt: '',
  imageClasses: '',
  buttons: undefined,
  list: undefined,
  table: undefined,
  tableHeaders: undefined,
  form: undefined,
  paddingRight: 0,
  paddingLeft: 0,
  align: 'center'
})

// متغیرهای محلی
const formData = ref<Record<string, any>>({})
const isSubmitting = ref(false)
const successMessage = ref('')
const errorMessage = ref('')

// بررسی وجود محتوا
const hasContent = computed(() => {
  return props.content || props.imageUrl || props.buttons?.length || props.list?.length || props.table?.length || props.form
})

// مدیریت کلیک دکمه
const handleButtonClick = (button: Button) => {
  if (button.action) {
    // اینجا می‌توانید عملیات مورد نظر را انجام دهید
    console.log('کلیک روی دکمه:', button)
  }
}

// مدیریت ارسال فرم
const handleFormSubmit = async () => {
  if (!props.form) return
  
  isSubmitting.value = true
  successMessage.value = ''
  errorMessage.value = ''
  
  try {
    // شبیه‌سازی API call
    await new Promise(resolve => setTimeout(resolve, 1500))
    
    // موفقیت
    successMessage.value = 'فرم با موفقیت ارسال شد!'
    
    // پاک کردن فرم
    formData.value = {}
    
    // پاک کردن پیام موفقیت بعد از 5 ثانیه
    setTimeout(() => {
      successMessage.value = ''
    }, 5000)
    
  } catch (error) {
    // خطا
    errorMessage.value = 'خطا در ارسال فرم. لطفاً دوباره تلاش کنید.'
    
    // پاک کردن پیام خطا بعد از 5 ثانیه
    setTimeout(() => {
      errorMessage.value = ''
    }, 5000)
  } finally {
    isSubmitting.value = false
  }
}

// استایل کامپوننت بر اساس چینش
const widgetStyle = computed(() => ({
  paddingRight: `${props.paddingRight}px`,
  paddingLeft: `${props.paddingLeft}px`,
  display: 'flex',
  alignItems: 'center',
  justifyContent: getJustifyContent(props.align),
  width: '100%',
  height: '100%'
}))

// تابع تعیین justify-content بر اساس چینش
function getJustifyContent(align: string): string {
  switch (align) {
    case 'left':
      return 'flex-start'  // در RTL: چپ = flex-start
    case 'center':
      return 'center'
    case 'right':
      return 'flex-end'  // در RTL: راست = flex-end
    default:
      return 'center'
  }
}
</script>

<style scoped>
.footer-custom-widget {
  transition: all 0.2s ease;
}

.footer-custom-widget button:hover:not(:disabled) {
  transform: translateY(-1px);
}

.footer-custom-widget input:focus,
.footer-custom-widget textarea:focus,
.footer-custom-widget select:focus {
  background-color: rgba(255, 255, 255, 0.3);
}

.footer-custom-widget .custom-image img:hover {
  transform: scale(1.02);
  transition: transform 0.2s ease;
}

.footer-custom-widget .custom-table {
  border-radius: 0.5rem;
  overflow: hidden;
}

.footer-custom-widget .custom-table th {
  background-color: rgba(255, 255, 255, 0.1);
}

.footer-custom-widget .custom-table tr:hover {
  background-color: rgba(255, 255, 255, 0.05);
}

/* انیمیشن برای دکمه‌ها */
@keyframes bounce {
  0%, 100% { transform: translateY(0); }
  50% { transform: translateY(-2px); }
}

.footer-custom-widget button:active:not(:disabled) {
  animation: bounce 0.1s ease-in-out;
}
</style>

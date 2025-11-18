<template>
  <div class="fixed inset-0 bg-gray-600 bg-opacity-50 flex items-center justify-center z-50">
    <div class="bg-white rounded-lg p-6 w-full max-w-2xl max-h-[90vh] overflow-y-auto">
      <div class="flex justify-between items-center mb-6">
        <h3 class="text-lg font-bold text-gray-900">
          {{ editingMessage ? 'ویرایش پیام' : 'ایجاد پیام جدید' }}
        </h3>
        <button @click="$emit('close')" class="text-gray-400 hover:text-gray-600">
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>

      <form @submit.prevent="handleSubmit" class="space-y-6">
        <!-- عنوان پیام -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">عنوان پیام *</label>
          <input 
            v-model="formData.title" 
            type="text" 
            required
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="مثال: پیام تولد"
          />
        </div>

        <!-- دسته‌بندی -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">دسته‌بندی *</label>
          <select v-model="formData.category" required class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent">
            <option value="">انتخاب دسته‌بندی</option>
            <option value="birthday">تولد</option>
            <option value="wedding">عروسی</option>
            <option value="newyear">سال نو</option>
            <option value="corporate">شرکتی</option>
            <option value="general">عمومی</option>
          </select>
        </div>

        <!-- محتوای پیام -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">محتوای پیام *</label>
          <textarea 
            v-model="formData.content" 
            rows="6" 
            required
            :maxlength="maxContentLength"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent resize-none"
            placeholder="محتوای پیام خود را اینجا بنویسید..."
          ></textarea>
          <div class="flex justify-between items-center mt-1">
            <div class="text-xs text-gray-500">
              {{ formData.content.length }}/{{ maxContentLength }} کاراکتر
            </div>
            <div v-if="formData.content.length > maxContentLength * 0.9" class="text-xs text-orange-600">
              نزدیک به حد مجاز
            </div>
          </div>
        </div>

        <!-- متغیرهای قالب -->
        <div v-if="formData.isTemplate" class="border border-blue-200 rounded-lg p-6 bg-blue-50">
          <h4 class="text-sm font-medium text-blue-900 mb-3">متغیرهای قالب</h4>
          <div class="grid grid-cols-2 gap-3">
            <div v-for="variable in templateVariables" :key="variable" class="flex items-center space-x-2 space-x-reverse">
              <span class="text-xs bg-blue-100 text-blue-800 px-2 py-1 rounded">{{ variable }}</span>
              <span class="text-xs text-blue-700">قابل جایگزینی</span>
            </div>
          </div>
          <p class="text-xs text-blue-600 mt-2">
            از متغیرها با فرمت {نام_متغیر} استفاده کنید
          </p>
        </div>

        <!-- تنظیمات پیشرفته -->
        <div class="border border-gray-200 rounded-lg p-6">
          <h4 class="text-sm font-medium text-gray-900 mb-3">تنظیمات پیشرفته</h4>
          <div class="space-y-3">
            <!-- نوع پیام -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نوع پیام</label>
              <select v-model="formData.type" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent">
                <option value="default">پیش‌فرض</option>
                <option value="custom">سفارشی</option>
                <option value="template">قالب</option>
              </select>
            </div>

            <!-- زبان -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">زبان</label>
              <select v-model="formData.language" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent">
                <option value="fa">فارسی</option>
                <option value="en">انگلیسی</option>
                <option value="ar">عربی</option>
              </select>
            </div>

            <!-- اولویت -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">اولویت نمایش</label>
              <select v-model="formData.priority" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent">
                <option value="low">کم</option>
                <option value="medium">متوسط</option>
                <option value="high">زیاد</option>
              </select>
            </div>

            <!-- تنظیمات اضافی -->
            <div class="space-y-2">
              <div class="flex items-center">
                <input 
                  type="checkbox" 
                  v-model="formData.isActive" 
                  id="isActive" 
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label for="isActive" class="mr-2 block text-sm text-gray-900">فعال</label>
              </div>
              <div class="flex items-center">
                <input 
                  type="checkbox" 
                  v-model="formData.allowPersonalization" 
                  id="allowPersonalization" 
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label for="allowPersonalization" class="mr-2 block text-sm text-gray-900">اجازه شخصی‌سازی</label>
              </div>
              <div class="flex items-center">
                <input 
                  type="checkbox" 
                  v-model="formData.requireApproval" 
                  id="requireApproval" 
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label for="requireApproval" class="mr-2 block text-sm text-gray-900">نیاز به تأیید</label>
              </div>
            </div>
          </div>
        </div>

        <!-- پیش‌نمایش -->
        <div v-if="formData.content" class="border border-gray-200 rounded-lg p-6">
          <h4 class="text-sm font-medium text-gray-900 mb-3">پیش‌نمایش پیام</h4>
          <div class="bg-gray-50 rounded p-3 text-sm text-gray-700">
            <div class="font-medium mb-1">{{ formData.title }}</div>
            <div>{{ formData.content }}</div>
            <div class="text-xs text-gray-500 mt-2">
              دسته‌بندی: {{ getCategoryText(formData.category) }} | 
              نوع: {{ getTypeText(formData.type) }} | 
              زبان: {{ getLanguageText(formData.language) }}
            </div>
          </div>
        </div>

        <!-- دکمه‌های عملیات -->
        <div class="flex justify-end space-x-3 space-x-reverse pt-6 border-t border-gray-200">
          <button 
            type="button" 
            @click="$emit('close')"
            class="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-lg hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
          >
            انصراف
          </button>
          <button 
            type="submit" 
            :disabled="isSubmitting || !isFormValid"
            class="px-6 py-2 text-sm font-medium text-white bg-blue-600 border border-transparent rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 disabled:opacity-50"
          >
            {{ isSubmitting ? 'در حال ذخیره...' : (editingMessage ? 'به‌روزرسانی' : 'ایجاد') }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'

// Props
const props = defineProps<{
  message?: any
}>()

// Emits
const emit = defineEmits<{
  close: []
  created: [message: any]
  updated: [message: any]
}>()

// Reactive data
const isSubmitting = ref(false)

// فرم داده‌ها
const formData = ref({
  title: '',
  content: '',
  category: '',
  type: 'custom',
  language: 'fa',
  priority: 'medium',
  isActive: true,
  allowPersonalization: true,
  requireApproval: false,
  isTemplate: false,
  templateId: null
})

// متغیرهای قالب
const templateVariables = ref([])

// Computed properties
const editingMessage = computed(() => props.message)

const maxContentLength = computed(() => {
  return formData.value.type === 'default' ? 200 : 500
})

const isFormValid = computed(() => {
  return formData.value.title.trim() && 
         formData.value.content.trim() && 
         formData.value.category &&
         formData.value.content.length <= maxContentLength.value
})

// Methods
const getCategoryText = (category: string) => {
  const categoryMap = {
    birthday: 'تولد',
    wedding: 'عروسی',
    newyear: 'سال نو',
    corporate: 'شرکتی',
    general: 'عمومی'
  }
  return categoryMap[category] || category
}

const getTypeText = (type: string) => {
  const typeMap = {
    default: 'پیش‌فرض',
    custom: 'سفارشی',
    template: 'قالب'
  }
  return typeMap[type] || type
}

const getLanguageText = (language: string) => {
  const languageMap = {
    fa: 'فارسی',
    en: 'انگلیسی',
    ar: 'عربی'
  }
  return languageMap[language] || language
}

const extractTemplateVariables = (content: string) => {
  const regex = /\{([^}]+)\}/g
  const variables = []
  let match
  
  while ((match = regex.exec(content)) !== null) {
    if (!variables.includes(match[1])) {
      variables.push(match[1])
    }
  }
  
  return variables
}

const validateContent = () => {
  const errors = []
  
  // بررسی طول محتوا
  if (formData.value.content.length > maxContentLength.value) {
    errors.push(`طول محتوا نباید بیشتر از ${maxContentLength.value} کاراکتر باشد`)
  }
  
  // بررسی کلمات ممنوعه (در نسخه واقعی از API دریافت می‌شود)
  const bannedWords = ['کلمه1', 'کلمه2', 'کلمه3']
  const content = formData.value.content.toLowerCase()
  
  for (const word of bannedWords) {
    if (content.includes(word.toLowerCase())) {
      errors.push(`کلمه "${word}" در محتوا مجاز نیست`)
    }
  }
  
  return errors
}

const handleSubmit = async () => {
  const errors = validateContent()
  if (errors.length > 0) {
    alert('خطاهای زیر را برطرف کنید:\n' + errors.join('\n'))
    return
  }

  isSubmitting.value = true

  try {
    const messageData = {
      id: editingMessage.value?.id || Date.now(),
      title: formData.value.title,
      content: formData.value.content,
      category: formData.value.category,
      type: formData.value.type,
      language: formData.value.language,
      priority: formData.value.priority,
      isActive: formData.value.isActive,
      allowPersonalization: formData.value.allowPersonalization,
      requireApproval: formData.value.requireApproval,
      isTemplate: formData.value.isTemplate,
      templateId: formData.value.templateId,
      usageCount: editingMessage.value?.usageCount || 0,
      createdAt: editingMessage.value?.createdAt || new Date(),
      updatedAt: new Date(),
      createdBy: editingMessage.value?.createdBy || 'کاربر فعلی',
      status: formData.value.requireApproval ? 'pending' : 'active'
    }

    // در نسخه واقعی: ارسال به API
    console.log('Saving message:', messageData)

    // شبیه‌سازی تأخیر
    await new Promise(resolve => setTimeout(resolve, 1000))

    if (editingMessage.value) {
      emit('updated', messageData)
    } else {
      emit('created', messageData)
    }
  } catch (error) {
    console.error('Error saving message:', error)
    alert('خطا در ذخیره پیام')
  } finally {
    isSubmitting.value = false
  }
}

// Watchers
watch(() => formData.value.content, (newContent) => {
  if (formData.value.isTemplate) {
    templateVariables.value = extractTemplateVariables(newContent)
  }
})

// Lifecycle
onMounted(() => {
  if (editingMessage.value) {
    // ویرایش پیام موجود
    formData.value = { ...editingMessage.value }
  }
})
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 

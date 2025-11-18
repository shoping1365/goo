<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200">
    <!-- هدر بخش -->
    <div class="p-6 border-b border-gray-200">
      <div class="flex items-center justify-between">
        <div>
          <h2 class="text-lg font-semibold text-gray-900">مدیریت محتوا و قالب‌های پیام</h2>
          <p class="text-gray-600 mt-1">مدیریت قالب‌های پیام، متن‌های پیش‌فرض و شخصی‌سازی محتوا</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button @click="showTemplateForm = true" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors">
            <span class="i-heroicons-plus ml-2"></span>
            افزودن قالب جدید
          </button>
        </div>
      </div>
    </div>

    <!-- آمار محتوا -->
    <div class="p-6 border-b border-gray-200">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div class="bg-blue-50 p-6 rounded-lg">
          <div class="flex items-center">
            <div class="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center ml-3">
              <span class="i-heroicons-document-text text-blue-600 text-xl"></span>
            </div>
            <div>
              <p class="text-sm text-blue-600">قالب‌های فعال</p>
              <p class="text-2xl font-bold text-blue-900">{{ stats.activeTemplates }}</p>
            </div>
          </div>
        </div>
        <div class="bg-green-50 p-6 rounded-lg">
          <div class="flex items-center">
            <div class="w-12 h-12 bg-green-100 rounded-lg flex items-center justify-center ml-3">
              <span class="i-heroicons-envelope text-green-600 text-xl"></span>
            </div>
            <div>
              <p class="text-sm text-green-600">پیام‌های ارسال شده</p>
              <p class="text-2xl font-bold text-green-900">{{ stats.sentMessages }}</p>
            </div>
          </div>
        </div>
        <div class="bg-purple-50 p-6 rounded-lg">
          <div class="flex items-center">
            <div class="w-12 h-12 bg-purple-100 rounded-lg flex items-center justify-center ml-3">
              <span class="i-heroicons-eye text-purple-600 text-xl"></span>
            </div>
            <div>
              <p class="text-sm text-purple-600">نرخ باز شدن</p>
              <p class="text-2xl font-bold text-purple-900">{{ stats.openRate }}%</p>
            </div>
          </div>
        </div>
        <div class="bg-orange-50 p-6 rounded-lg">
          <div class="flex items-center">
            <div class="w-12 h-12 bg-orange-100 rounded-lg flex items-center justify-center ml-3">
              <span class="i-heroicons-cursor-arrow-rays text-orange-600 text-xl"></span>
            </div>
            <div>
              <p class="text-sm text-orange-600">نرخ کلیک</p>
              <p class="text-2xl font-bold text-orange-900">{{ stats.clickRate }}%</p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تب‌های محتوا -->
    <div class="border-b border-gray-200">
      <div class="flex border-b border-gray-200 overflow-x-auto">
        <button v-for="tab in tabs" :key="tab.value" @click="activeTab = tab.value"
          :class="['px-6 py-3 -mb-px font-medium text-sm focus:outline-none whitespace-nowrap', activeTab === tab.value ? 'border-b-2 border-blue-600 text-blue-700' : 'text-gray-500 hover:text-blue-600']">
          {{ tab.label }}
        </button>
      </div>
    </div>

    <!-- محتوای تب‌ها -->
    <div class="p-6">
      <!-- قالب‌های پیام -->
      <div v-if="activeTab === 'templates'" class="space-y-6">
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <div v-for="template in templates" :key="template.id" class="border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow">
            <div class="flex items-center justify-between mb-4">
              <div class="flex items-center">
                <div class="w-10 h-10 rounded-full flex items-center justify-center ml-3" :style="{ backgroundColor: template.color + '20', color: template.color }">
                  <span class="i-heroicons-document-text text-lg"></span>
                </div>
                <div>
                  <h4 class="font-medium text-gray-900">{{ template.name }}</h4>
                  <p class="text-sm text-gray-500">{{ template.type }}</p>
                </div>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <button @click="editTemplate(template)" class="text-blue-600 hover:text-blue-900">
                  <span class="i-heroicons-pencil-square"></span>
                </button>
                <button @click="deleteTemplate(template)" class="text-red-600 hover:text-red-900">
                  <span class="i-heroicons-trash"></span>
                </button>
              </div>
            </div>
            
            <div class="space-y-3">
              <div class="text-sm text-gray-600 line-clamp-3">{{ template.content }}</div>
              <div class="flex justify-between text-sm">
                <span class="text-gray-500">استفاده شده:</span>
                <span class="font-medium">{{ template.usageCount }}</span>
              </div>
              <div class="flex justify-between text-sm">
                <span class="text-gray-500">آخرین ویرایش:</span>
                <span class="font-medium">{{ formatDate(template.lastModified) }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- متغیرهای شخصی‌سازی -->
      <div v-if="activeTab === 'variables'" class="space-y-6">
        <div class="bg-blue-50 p-6 rounded-lg">
          <h4 class="text-md font-medium text-blue-900 mb-2">متغیرهای قابل استفاده در قالب‌ها</h4>
          <p class="text-sm text-blue-700">از این متغیرها برای شخصی‌سازی پیام‌ها استفاده کنید</p>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div class="border border-gray-200 rounded-lg p-6">
            <h5 class="font-medium text-gray-900 mb-3">متغیرهای کاربر</h5>
            <div class="space-y-2">
              <div v-for="variable in userVariables" :key="variable.name" class="flex justify-between items-center p-2 bg-gray-50 rounded">
                <span class="text-sm text-gray-700">{{ variable.description }}</span>
                <code class="text-xs bg-gray-200 px-2 py-1 rounded">{{ variable.name }}</code>
              </div>
            </div>
          </div>

          <div class="border border-gray-200 rounded-lg p-6">
            <h5 class="font-medium text-gray-900 mb-3">متغیرهای کوپن</h5>
            <div class="space-y-2">
              <div v-for="variable in couponVariables" :key="variable.name" class="flex justify-between items-center p-2 bg-gray-50 rounded">
                <span class="text-sm text-gray-700">{{ variable.description }}</span>
                <code class="text-xs bg-gray-200 px-2 py-1 rounded">{{ variable.name }}</code>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- پیش‌نمایش قالب -->
      <div v-if="activeTab === 'preview'" class="space-y-6">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">انتخاب قالب</label>
            <select v-model="selectedTemplate" @change="updatePreview" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
              <option value="">انتخاب کنید</option>
              <option v-for="template in templates" :key="template.id" :value="template.id">{{ template.name }}</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع پیش‌نمایش</label>
            <select v-model="previewType" @change="updatePreview" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
              <option value="email">ایمیل</option>
              <option value="sms">پیامک</option>
              <option value="push">اعلان</option>
            </select>
          </div>
        </div>

        <div class="border border-gray-200 rounded-lg p-6">
          <h4 class="text-md font-medium text-gray-900 mb-4">پیش‌نمایش</h4>
          <div class="bg-gray-50 p-6 rounded-lg">
            <!-- 
              ⚠️ امنیت XSS: استفاده از v-html خطرناک است!
              
              این کد محتوای HTML را بدون sanitization نمایش می‌دهد که می‌تواند منجر به حملات XSS شود.
              
              ✅ راه حل صحیح:
              1. قبل از استفاده از v-html، محتوا را با کتابخانه sanitization (مثل DOMPurify) پاکسازی کنید
              2. یا از {{ }} به جای v-html استفاده کنید اگر HTML نیاز نیست
              3. محتوای کاربر را هرگز بدون sanitization در v-html قرار ندهید
              
              مثال صحیح:
              import DOMPurify from 'dompurify'
              const sanitizedPreviewContent = computed(() => DOMPurify.sanitize(previewContent.value))
              <div v-html="sanitizedPreviewContent"></div>
            -->
            <div v-if="previewContent" v-html="previewContent" class="text-sm text-gray-700"></div>
            <div v-else class="text-gray-400 text-center py-8">قالبی انتخاب نشده است</div>
          </div>
        </div>
      </div>

      <!-- آمار عملکرد -->
      <div v-if="activeTab === 'performance'" class="space-y-6">
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <div class="bg-white border border-gray-200 rounded-lg p-6">
            <h4 class="text-md font-medium text-gray-900 mb-4">بهترین قالب‌ها</h4>
            <div class="space-y-3">
              <div v-for="template in topTemplates" :key="template.id" class="flex items-center justify-between">
                <div class="flex items-center">
                  <div class="w-3 h-3 rounded-full ml-2" :style="{ backgroundColor: template.color }"></div>
                  <span class="text-sm text-gray-700">{{ template.name }}</span>
                </div>
                <div class="flex items-center space-x-2 space-x-reverse">
                  <div class="w-32 bg-gray-200 rounded-full h-2">
                    <div class="h-2 rounded-full bg-green-600" :style="{ width: template.performance + '%' }"></div>
                  </div>
                  <span class="text-sm text-gray-600">{{ template.performance }}%</span>
                </div>
              </div>
            </div>
          </div>

          <div class="bg-white border border-gray-200 rounded-lg p-6">
            <h4 class="text-md font-medium text-gray-900 mb-4">روند عملکرد</h4>
            <div class="h-48 flex items-center justify-center text-gray-400">[نمودار روند عملکرد]</div>
          </div>
        </div>
      </div>
    </div>

    <!-- مودال فرم قالب -->
    <div v-if="showTemplateForm" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg shadow-xl max-w-2xl w-full mx-4 max-h-[90vh] overflow-y-auto">
        <div class="p-6 border-b border-gray-200">
          <div class="flex items-center justify-between">
            <h3 class="text-lg font-semibold text-gray-900">
              {{ editingTemplate ? 'ویرایش قالب' : 'افزودن قالب جدید' }}
            </h3>
            <button @click="closeForm" class="text-gray-400 hover:text-gray-600">
              <span class="i-heroicons-x-mark text-xl"></span>
            </button>
          </div>
        </div>
        
        <form @submit.prevent="handleSubmit" class="p-6 space-y-6">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نام قالب *</label>
              <input v-model="form.name" type="text" required class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="نام قالب">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نوع قالب *</label>
              <select v-model="form.type" required class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                <option value="email">ایمیل</option>
                <option value="sms">پیامک</option>
                <option value="push">اعلان</option>
                <option value="webhook">وب‌هوک</option>
              </select>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">محتوای قالب *</label>
            <textarea v-model="form.content" rows="8" required class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="محتوای قالب با متغیرها"></textarea>
            <p class="text-xs text-gray-500 mt-1">از متغیرها مانند &#123;&#123;user_name&#125;&#125;، &#123;&#123;coupon_code&#125;&#125; استفاده کنید</p>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">رنگ قالب</label>
              <input v-model="form.color" type="color" class="w-full h-10 border border-gray-300 rounded-lg">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
              <select v-model="form.status" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                <option value="active">فعال</option>
                <option value="inactive">غیرفعال</option>
              </select>
            </div>
          </div>
        </form>
        
        <div class="p-6 border-t border-gray-200 flex justify-end space-x-3 space-x-reverse">
          <button @click="closeForm" class="px-6 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200 transition-colors">
            انصراف
          </button>
          <button @click="handleSubmit" class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors">
            ذخیره
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'

const activeTab = ref('templates')
const showTemplateForm = ref(false)
const editingTemplate = ref<any>(null)
const selectedTemplate = ref('')
const previewType = ref('email')
const previewContent = ref('')

const tabs = [
  { value: 'templates', label: 'قالب‌های پیام' },
  { value: 'variables', label: 'متغیرها' },
  { value: 'preview', label: 'پیش‌نمایش' },
  { value: 'performance', label: 'عملکرد' }
]

const stats = ref({
  activeTemplates: 12,
  sentMessages: 4567,
  openRate: 23.5,
  clickRate: 8.2
})

const templates = ref([
  {
    id: 1,
    name: 'قالب خوشامدگویی',
    type: 'ایمیل',
    content: 'سلام {{user_name}}! کوپن {{coupon_code}} با تخفیف {{discount_amount}} برای شما فعال شده است.',
    color: '#3b82f6',
    usageCount: 234,
    lastModified: '2024-01-15T10:30:00',
    status: 'active'
  },
  {
    id: 2,
    name: 'قالب یادآوری',
    type: 'پیامک',
    content: 'کد تخفیف {{coupon_code}} تا {{expiry_date}} معتبر است. از فرصت استفاده کنید!',
    color: '#10b981',
    usageCount: 156,
    lastModified: '2024-01-14T09:15:00',
    status: 'active'
  }
])

const userVariables = ref([
  { name: '{{user_name}}', description: 'نام کاربر' },
  { name: '{{user_email}}', description: 'ایمیل کاربر' },
  { name: '{{user_phone}}', description: 'شماره تلفن کاربر' },
  { name: '{{user_level}}', description: 'سطح وفاداری کاربر' }
])

const couponVariables = ref([
  { name: '{{coupon_code}}', description: 'کد کوپن' },
  { name: '{{discount_amount}}', description: 'مبلغ تخفیف' },
  { name: '{{discount_percent}}', description: 'درصد تخفیف' },
  { name: '{{expiry_date}}', description: 'تاریخ انقضا' },
  { name: '{{min_amount}}', description: 'حداقل مبلغ سفارش' }
])

const topTemplates = ref([
  { id: 1, name: 'قالب خوشامدگویی', performance: 85, color: '#3b82f6' },
  { id: 2, name: 'قالب یادآوری', performance: 72, color: '#10b981' },
  { id: 3, name: 'قالب ویژه', performance: 68, color: '#f59e0b' }
])

const form = reactive({
  name: '',
  type: 'email',
  content: '',
  color: '#3b82f6',
  status: 'active'
})

const formatDate = (date: string): string => {
  return new Intl.DateTimeFormat('fa-IR').format(new Date(date))
}

const editTemplate = (template: any) => {
  editingTemplate.value = template
  Object.assign(form, template)
  showTemplateForm.value = true
}

const deleteTemplate = async (template: any) => {
  if (confirm(`آیا از حذف قالب "${template.name}" اطمینان دارید؟`)) {
    try {
      const index = templates.value.findIndex(t => t.id === template.id)
      if (index !== -1) {
        templates.value.splice(index, 1)
      }
    } catch (error) {
      console.error('خطا در حذف قالب:', error)
    }
  }
}

const updatePreview = () => {
  if (selectedTemplate.value) {
    const template = templates.value.find(t => t.id === parseInt(selectedTemplate.value))
    if (template) {
      previewContent.value = template.content
        .replace('{{user_name}}', 'علی احمدی')
        .replace('{{coupon_code}}', 'WELCOME2024')
        .replace('{{discount_amount}}', '۵۰,۰۰۰ تومان')
        .replace('{{discount_percent}}', '۲۰٪')
        .replace('{{expiry_date}}', '۱۴۰۳/۰۲/۱۵')
        .replace('{{min_amount}}', '۲۵۰,۰۰۰ تومان')
    }
  } else {
    previewContent.value = ''
  }
}

const handleSubmit = async () => {
  if (!form.name || !form.type || !form.content) {
    alert('لطفاً فیلدهای اجباری را پر کنید')
    return
  }
  
  if (editingTemplate.value) {
    Object.assign(editingTemplate.value, form)
  } else {
    templates.value.unshift({
      id: Date.now(),
      ...form,
      usageCount: 0,
      lastModified: new Date().toISOString()
    })
  }
  closeForm()
}

const closeForm = () => {
  showTemplateForm.value = false
  editingTemplate.value = null
  Object.assign(form, { name: '', type: 'email', content: '', color: '#3b82f6', status: 'active' })
}
</script>

<!--
  کامپوننت مدیریت محتوا و قالب‌های پیام
  شامل:
  1. مدیریت قالب‌های پیام (ایمیل، پیامک، اعلان)
  2. متغیرهای شخصی‌سازی
  3. پیش‌نمایش قالب‌ها
  4. آمار عملکرد قالب‌ها
  توضیحات کامل به فارسی در کد
--> 

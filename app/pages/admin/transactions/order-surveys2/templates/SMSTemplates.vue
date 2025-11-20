<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="flex items-center justify-between">
      <div>
        <h3 class="text-lg font-semibold text-gray-800">قالب‌های SMS</h3>
        <p class="text-gray-600 text-sm">مدیریت قالب‌های پیام‌کش برای نظرسنجی</p>
      </div>
      <button 
        class="px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg text-sm transition-colors flex items-center space-x-2 space-x-reverse"
        @click="showCreateModal = true"
      >
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
        </svg>
        <span>قالب جدید</span>
      </button>
    </div>

    <!-- Template Categories -->
    <div class="bg-white rounded-lg border border-gray-200 px-4 py-4">
      <h4 class="text-md font-semibold text-gray-800 mb-3">دسته‌بندی قالب‌ها</h4>
      <div class="flex flex-wrap gap-2">
        <button 
          v-for="category in templateCategories"
          :key="category.id"
          :class="[
            'px-3 py-1 rounded-full text-sm font-medium transition-colors',
            selectedCategory === category.id 
              ? 'bg-blue-100 text-blue-800 border border-blue-200' 
              : 'bg-gray-100 text-gray-700 hover:bg-gray-200'
          ]"
          @click="selectedCategory = category.id"
        >
          {{ category.name }}
        </button>
      </div>
    </div>

    <!-- Templates Grid -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gapx-4 py-4">
      <div 
        v-for="template in filteredTemplates" 
        :key="template.id"
        class="bg-white rounded-lg border border-gray-200 px-4 py-4 hover:shadow-md transition-shadow"
      >
        <!-- Template Header -->
        <div class="flex items-start justify-between mb-4">
          <div>
            <h5 class="font-semibold text-gray-900">{{ template.name }}</h5>
            <p class="text-sm text-gray-500">{{ template.description }}</p>
          </div>
          <div class="flex items-center space-x-2 space-x-reverse">
            <span 
              :class="[
                'inline-flex items-center px-2 py-1 rounded-full text-xs font-medium',
                getStatusClass(template.status)
              ]"
            >
              {{ getStatusText(template.status) }}
            </span>
            <button 
              class="text-gray-400 hover:text-yellow-500 transition-colors"
              @click="toggleFavorite(template.id)"
            >
              <svg 
                :class="['w-4 h-4', template.isFavorite ? 'text-yellow-500 fill-current' : '']"
                fill="none" 
                stroke="currentColor" 
                viewBox="0 0 24 24"
              >
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z"></path>
              </svg>
            </button>
          </div>
        </div>

        <!-- Template Content -->
        <div class="mb-4">
          <div class="bg-gray-50 rounded-lg p-3 mb-3">
            <div class="text-sm text-gray-700 whitespace-pre-wrap">{{ template.content }}</div>
          </div>
          
          <!-- Variables -->
          <div v-if="template.variables.length > 0" class="mb-3">
            <div class="text-xs font-medium text-gray-700 mb-1">متغیرهای قابل استفاده:</div>
            <div class="flex flex-wrap gap-1">
              <span 
                v-for="variable in template.variables"
                :key="variable"
                class="inline-flex items-center px-2 py-1 bg-blue-100 text-blue-800 rounded text-xs font-medium"
              >
                {{ variable }}
              </span>
            </div>
          </div>

          <!-- Template Stats -->
          <div class="grid grid-cols-2 gap-3 text-xs">
            <div class="text-center">
              <div class="font-medium text-gray-900">{{ template.usageCount }}</div>
              <div class="text-gray-500">استفاده</div>
            </div>
            <div class="text-center">
              <div class="font-medium text-gray-900">{{ template.successRate }}%</div>
              <div class="text-gray-500">موفقیت</div>
            </div>
          </div>
        </div>

        <!-- Template Actions -->
        <div class="flex items-center justify-between pt-4 border-t border-gray-100">
          <div class="flex items-center space-x-2 space-x-reverse">
            <button 
              class="text-blue-600 hover:text-blue-800 transition-colors"
              title="ویرایش"
              @click="editTemplate(template.id)"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
              </svg>
            </button>
            
            <button 
              class="text-green-600 hover:text-green-800 transition-colors"
              title="کپی"
              @click="duplicateTemplate(template.id)"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"></path>
              </svg>
            </button>
            
            <button 
              class="text-purple-600 hover:text-purple-800 transition-colors"
              title="پیش‌نمایش"
              @click="previewTemplate(template.id)"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
              </svg>
            </button>
          </div>
          
          <button 
            v-if="canDeleteSMSTemplate"
            class="text-red-600 hover:text-red-800 transition-colors"
            title="حذف"
            @click="deleteTemplate(template.id)"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
            </svg>
          </button>
        </div>
      </div>
    </div>

    <!-- Create/Edit Template Modal -->
    <div 
      v-if="showCreateModal || showEditModal"
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50"
      @click="closeModal"
    >
      <div 
        class="bg-white rounded-lg px-4 py-4 w-full max-w-2xl mx-4 max-h-[90vh] overflow-y-auto"
        @click.stop
      >
        <div class="flex items-center justify-between mb-6">
          <h3 class="text-lg font-semibold text-gray-800">
            {{ showEditModal ? 'ویرایش قالب' : 'قالب جدید' }}
          </h3>
          <button 
            class="text-gray-400 hover:text-gray-600 transition-colors"
            @click="closeModal"
          >
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>

        <form class="space-y-4" @submit.prevent="saveTemplate">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">نام قالب</label>
            <input 
              v-model="templateForm.name"
              type="text"
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="مثال: نظرسنجی عمومی"
            >
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">توضیحات</label>
            <textarea 
              v-model="templateForm.description"
              rows="2"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 resize-none"
              placeholder="توضیح مختصر درباره قالب"
            ></textarea>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">دسته‌بندی</label>
            <select 
              v-model="templateForm.categoryId"
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="">انتخاب دسته‌بندی</option>
              <option 
                v-for="category in templateCategories"
                :key="category.id"
                :value="category.id"
              >
                {{ category.name }}
              </option>
            </select>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">محتوای پیام</label>
            <div class="mb-2">
              <div class="text-xs text-gray-500 mb-1">متغیرهای قابل استفاده:</div>
              <div class="flex flex-wrap gap-1 mb-2">
                <button 
                  v-for="variable in availableVariables"
                  :key="variable.key"
                  type="button"
                  class="inline-flex items-center px-2 py-1 bg-blue-100 text-blue-800 rounded text-xs font-medium hover:bg-blue-200 transition-colors"
                  @click="insertVariable(variable.key)"
                >
                  {{ variable.label }}
                </button>
              </div>
            </div>
            <textarea 
              v-model="templateForm.content"
              rows="6"
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 resize-none font-mono text-sm"
              placeholder="متن پیام خود را اینجا بنویسید..."
            ></textarea>
            <div class="text-xs text-gray-500 mt-1">
              تعداد کاراکتر: {{ templateForm.content.length }} / 160
            </div>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">وضعیت</label>
              <select 
                v-model="templateForm.status"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="active">فعال</option>
                <option value="inactive">غیرفعال</option>
                <option value="draft">پیش‌نویس</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">اولویت</label>
              <select 
                v-model="templateForm.priority"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="low">پایین</option>
                <option value="medium">متوسط</option>
                <option value="high">بالا</option>
              </select>
            </div>
          </div>

          <div class="flex items-center space-x-4 space-x-reverse pt-4 border-t border-gray-200">
            <button 
              type="submit"
              :disabled="saving"
              class="px-4 py-2 bg-blue-600 hover:bg-blue-700 disabled:bg-gray-400 text-white rounded-lg text-sm transition-colors flex items-center space-x-2 space-x-reverse"
            >
              <svg v-if="saving" class="animate-spin w-4 h-4" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              <span>{{ saving ? 'در حال ذخیره...' : 'ذخیره قالب' }}</span>
            </button>
            
            <button 
              type="button"
              class="px-4 py-2 bg-gray-300 hover:bg-gray-400 text-gray-700 rounded-lg text-sm transition-colors"
              @click="closeModal"
            >
              انصراف
            </button>
          </div>
        </form>
      </div>
    </div>

    <!-- Preview Modal -->
    <div 
      v-if="showPreviewModal"
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50"
      @click="showPreviewModal = false"
    >
      <div 
        class="bg-white rounded-lg px-4 py-4 w-full max-w-md mx-4"
        @click.stop
      >
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-800">پیش‌نمایش قالب</h3>
          <button 
            class="text-gray-400 hover:text-gray-600 transition-colors"
            @click="showPreviewModal = false"
          >
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>

        <div class="bg-gray-50 rounded-lg px-4 py-4 mb-4">
          <div class="text-sm text-gray-700 whitespace-pre-wrap">{{ previewContent }}</div>
        </div>

        <div class="text-xs text-gray-500 text-center">
          تعداد کاراکتر: {{ previewContent.length }} / 160
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const useAuth: () => { user: { id?: number; name?: string; email?: string } | null; hasPermission: (perm: string) => boolean }
</script>

<script setup lang="ts">
import { computed, reactive, ref } from 'vue';
// 
const { hasPermission } = useAuth()

const canDeleteSMSTemplate = computed(() => hasPermission('sms-template.delete'))

interface Template {
  id: number
  name: string
  description: string
  content: string
  categoryId: number
  status: string
  priority: string
  isFavorite: boolean
  usageCount: number
  successRate: number
  variables: string[]
  createdAt: string
  updatedAt: string
}

interface TemplateCategory {
  id: number
  name: string
  description: string
}

interface TemplateVariable {
  key: string
  label: string
  description: string
}

const emit = defineEmits<{
  'template-created': [template: Template]
  'template-updated': [template: Template]
  'template-deleted': [templateId: number]
  'template-duplicated': [templateId: number]
}>()

// Reactive data
const showCreateModal = ref(false)
const showEditModal = ref(false)
const showPreviewModal = ref(false)
const saving = ref(false)
const selectedCategory = ref<number | null>(null)
const editingTemplateId = ref<number | null>(null)
const previewContent = ref('')

const templateForm = reactive({
  name: '',
  description: '',
  content: '',
  categoryId: '',
  status: 'active',
  priority: 'medium'
})

const templateCategories = reactive<TemplateCategory[]>([
  { id: 1, name: 'نظرسنجی عمومی', description: 'قالب‌های عمومی نظرسنجی' },
  { id: 2, name: 'نظرسنجی محصول', description: 'قالب‌های مربوط به محصولات' },
  { id: 3, name: 'نظرسنجی خدمات', description: 'قالب‌های مربوط به خدمات' },
  { id: 4, name: 'یادآوری', description: 'قالب‌های یادآوری' },
  { id: 5, name: 'تشکر', description: 'قالب‌های تشکر' }
])

const availableVariables = reactive<TemplateVariable[]>([
  { key: '{customer_name}', label: 'نام مشتری', description: 'نام کامل مشتری' },
  { key: '{order_number}', label: 'شماره سفارش', description: 'شماره سفارش' },
  { key: '{product_name}', label: 'نام محصول', description: 'نام محصول سفارش شده' },
  { key: '{order_date}', label: 'تاریخ سفارش', description: 'تاریخ ثبت سفارش' },
  { key: '{delivery_date}', label: 'تاریخ تحویل', description: 'تاریخ تحویل سفارش' },
  { key: '{survey_link}', label: 'لینک نظرسنجی', description: 'لینک مستقیم نظرسنجی' },
  { key: '{company_name}', label: 'نام شرکت', description: 'نام شرکت شما' },
  { key: '{support_phone}', label: 'شماره پشتیبانی', description: 'شماره تماس پشتیبانی' }
])

const templates = reactive<Template[]>([
  {
    id: 1,
    name: 'نظرسنجی عمومی',
    description: 'قالب استاندارد برای نظرسنجی عمومی',
    content: 'سلام {customer_name} عزیز\nاز خرید شما تشکر می‌کنیم. لطفاً نظرات خود را درباره سفارش #{order_number} با ما به اشتراک بگذارید:\n{survey_link}\n{company_name}',
    categoryId: 1,
    status: 'active',
    priority: 'high',
    isFavorite: true,
    usageCount: 1250,
    successRate: 85,
    variables: ['{customer_name}', '{order_number}', '{survey_link}', '{company_name}'],
    createdAt: '2024-01-01T00:00:00Z',
    updatedAt: '2024-01-15T00:00:00Z'
  },
  {
    id: 2,
    name: 'نظرسنجی محصول',
    description: 'قالب تخصصی برای نظرسنجی محصولات',
    content: 'سلام {customer_name}\nامیدواریم از محصول {product_name} راضی باشید. نظرات شما برای ما مهم است:\n{survey_link}\n{company_name}',
    categoryId: 2,
    status: 'active',
    priority: 'medium',
    isFavorite: false,
    usageCount: 890,
    successRate: 92,
    variables: ['{customer_name}', '{product_name}', '{survey_link}', '{company_name}'],
    createdAt: '2024-01-05T00:00:00Z',
    updatedAt: '2024-01-10T00:00:00Z'
  },
  {
    id: 3,
    name: 'یادآوری نظرسنجی',
    description: 'قالب یادآوری برای مشتریانی که هنوز نظرسنجی نکرده‌اند',
    content: 'سلام {customer_name}\nهنوز نظرات شما را دریافت نکرده‌ایم. لطفاً چند دقیقه وقت بگذارید:\n{survey_link}\n{company_name}',
    categoryId: 4,
    status: 'active',
    priority: 'low',
    isFavorite: false,
    usageCount: 450,
    successRate: 78,
    variables: ['{customer_name}', '{survey_link}', '{company_name}'],
    createdAt: '2024-01-08T00:00:00Z',
    updatedAt: '2024-01-12T00:00:00Z'
  }
])

// Computed properties
const filteredTemplates = computed(() => {
  let filtered = templates
  
  if (selectedCategory.value) {
    filtered = filtered.filter(template => template.categoryId === selectedCategory.value)
  }
  
  return filtered
})

// Methods
const toggleFavorite = (templateId: number) => {
  const template = templates.find(t => t.id === templateId)
  if (template) {
    template.isFavorite = !template.isFavorite
  }
}

const editTemplate = (templateId: number) => {
  const template = templates.find(t => t.id === templateId)
  if (template) {
    editingTemplateId.value = templateId
    templateForm.name = template.name
    templateForm.description = template.description
    templateForm.content = template.content
    templateForm.categoryId = template.categoryId.toString()
    templateForm.status = template.status
    templateForm.priority = template.priority
    showEditModal.value = true
  }
}

const duplicateTemplate = (templateId: number) => {
  emit('template-duplicated', templateId)
}

const previewTemplate = (templateId: number) => {
  const template = templates.find(t => t.id === templateId)
  if (template) {
    previewContent.value = template.content
    showPreviewModal.value = true
  }
}

const deleteTemplate = (templateId: number) => {
  if (confirm('آیا مطمئن هستید که می‌خواهید این قالب را حذف کنید؟')) {
    const index = templates.findIndex(t => t.id === templateId)
    if (index !== -1) {
      templates.splice(index, 1)
      emit('template-deleted', templateId)
    }
  }
}

const insertVariable = (variable: string) => {
  const textarea = document.querySelector('textarea') as HTMLTextAreaElement
  if (textarea) {
    const start = textarea.selectionStart
    const end = textarea.selectionEnd
    const text = templateForm.content
    templateForm.content = text.substring(0, start) + variable + text.substring(end)
    
    // Set cursor position after inserted variable
    setTimeout(() => {
      textarea.focus()
      textarea.setSelectionRange(start + variable.length, start + variable.length)
    }, 0)
  }
}

const saveTemplate = async () => {
  saving.value = true
  try {
    await new Promise(resolve => setTimeout(resolve, 1000))
    
    if (showEditModal.value && editingTemplateId.value) {
      // Update existing template
      const index = templates.findIndex(t => t.id === editingTemplateId.value)
      if (index !== -1) {
        const updatedTemplate = {
          ...templates[index],
          name: templateForm.name,
          description: templateForm.description,
          content: templateForm.content,
          categoryId: parseInt(templateForm.categoryId),
          status: templateForm.status,
          priority: templateForm.priority,
          updatedAt: new Date().toISOString()
        }
        templates[index] = updatedTemplate
        emit('template-updated', updatedTemplate)
      }
    } else {
      // Create new template
      const newTemplate: Template = {
        id: Date.now(),
        name: templateForm.name,
        description: templateForm.description,
        content: templateForm.content,
        categoryId: parseInt(templateForm.categoryId),
        status: templateForm.status,
        priority: templateForm.priority,
        isFavorite: false,
        usageCount: 0,
        successRate: 0,
        variables: extractVariables(templateForm.content),
        createdAt: new Date().toISOString(),
        updatedAt: new Date().toISOString()
      }
      templates.push(newTemplate)
      emit('template-created', newTemplate)
    }
    
    closeModal()
  } finally {
    saving.value = false
  }
}

const closeModal = () => {
  showCreateModal.value = false
  showEditModal.value = false
  editingTemplateId.value = null
  resetForm()
}

const resetForm = () => {
  templateForm.name = ''
  templateForm.description = ''
  templateForm.content = ''
  templateForm.categoryId = ''
  templateForm.status = 'active'
  templateForm.priority = 'medium'
}

const extractVariables = (content: string): string[] => {
  const variableRegex = /\{([^}]+)\}/g
  const variables: string[] = []
  let match
  
  while ((match = variableRegex.exec(content)) !== null) {
    variables.push(match[0])
  }
  
  return variables
}

const getStatusClass = (status: string) => {
  const classes = {
    active: 'bg-green-100 text-green-800',
    inactive: 'bg-gray-100 text-gray-800',
    draft: 'bg-yellow-100 text-yellow-800'
  }
  return classes[status as keyof typeof classes] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status: string) => {
  const texts = {
    active: 'فعال',
    inactive: 'غیرفعال',
    draft: 'پیش‌نویس'
  }
  return texts[status as keyof typeof texts] || status
}
</script>

<style scoped>
.space-x-reverse > :not([hidden]) ~ :not([hidden]) {
  --tw-space-x-reverse: 1;
}
</style> 


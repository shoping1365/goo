<template>
  <div class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
    <div class="relative top-20 mx-auto p-5 border w-11/12 max-w-4xl shadow-lg rounded-md bg-white">
      <!-- Header -->
      <div class="flex justify-between items-center mb-6">
        <div>
          <h3 class="text-lg font-medium text-gray-900">ویرایش قالب پیام</h3>
          <p class="text-sm text-gray-600 mt-1">{{ template?.name }}</p>
        </div>
        <button 
          @click="$emit('close')"
          class="text-gray-400 hover:text-gray-600"
        >
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
        </button>
      </div>

      <!-- Content -->
      <div class="space-y-6">
        <!-- اطلاعات پایه -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">نام قالب</label>
            <input
              v-model="editedTemplate.name"
              type="text"
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">نوع</label>
            <select 
              v-model="editedTemplate.type"
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="email">ایمیل</option>
              <option value="sms">پیامک</option>
            </select>
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">توضیحات</label>
          <textarea
            v-model="editedTemplate.description"
            rows="2"
            class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
          ></textarea>
        </div>

        <!-- تب‌های ویرایش -->
        <div class="border-b border-gray-200">
          <nav class="-mb-px flex space-x-8 space-x-reverse">
            <button
              @click="activeTab = 'editor'"
              :class="{
                'border-blue-500 text-blue-600': activeTab === 'editor',
                'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300': activeTab !== 'editor'
              }"
              class="whitespace-nowrap py-2 px-1 border-b-2 font-medium text-sm"
            >
              ویرایشگر
            </button>
            <button
              @click="activeTab = 'preview'"
              :class="{
                'border-blue-500 text-blue-600': activeTab === 'preview',
                'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300': activeTab !== 'preview'
              }"
              class="whitespace-nowrap py-2 px-1 border-b-2 font-medium text-sm"
            >
              پیش‌نمایش
            </button>
            <button
              @click="activeTab = 'variables'"
              :class="{
                'border-blue-500 text-blue-600': activeTab === 'variables',
                'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300': activeTab !== 'variables'
              }"
              class="whitespace-nowrap py-2 px-1 border-b-2 font-medium text-sm"
            >
              متغیرها
            </button>
          </nav>
        </div>

        <!-- محتوای تب‌ها -->
        <div class="min-h-96">
          <!-- تب ویرایشگر -->
          <div v-if="activeTab === 'editor'" class="space-y-4">
            <div v-if="editedTemplate.type === 'email'">
              <div class="mb-4">
                <label class="block text-sm font-medium text-gray-700 mb-1">موضوع ایمیل</label>
                <input
                  v-model="editedTemplate.subject"
                  type="text"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                  placeholder="موضوع ایمیل"
                />
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">محتوای ایمیل</label>
                <div class="border border-gray-300 rounded-md">
                  <div class="bg-gray-50 px-3 py-2 border-b border-gray-300">
                    <div class="flex items-center space-x-4 space-x-reverse">
                      <button 
                        @click="insertVariable('bold')"
                        class="text-gray-600 hover:text-gray-800"
                        title="بولد"
                      >
                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 12h8a4 4 0 100-8H6v8zm0 0h8a4 4 0 110 8H6v-8z"></path>
                        </svg>
                      </button>
                      <button 
                        @click="insertVariable('italic')"
                        class="text-gray-600 hover:text-gray-800"
                        title="ایتالیک"
                      >
                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4"></path>
                        </svg>
                      </button>
                      <button 
                        @click="insertVariable('link')"
                        class="text-gray-600 hover:text-gray-800"
                        title="لینک"
                      >
                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1"></path>
                        </svg>
                      </button>
                      <div class="border-r border-gray-300 h-4"></div>
                      <button 
                        @click="insertVariable('name')"
                        class="text-xs bg-blue-100 text-blue-800 px-2 py-1 rounded hover:bg-blue-200"
                      >
                        نام
                      </button>
                      <button 
                        @click="insertVariable('amount')"
                        class="text-xs bg-blue-100 text-blue-800 px-2 py-1 rounded hover:bg-blue-200"
                      >
                        مبلغ
                      </button>
                      <button 
                        @click="insertVariable('code')"
                        class="text-xs bg-blue-100 text-blue-800 px-2 py-1 rounded hover:bg-blue-200"
                      >
                        کد کارت
                      </button>
                      <button 
                        @click="insertVariable('expiry')"
                        class="text-xs bg-blue-100 text-blue-800 px-2 py-1 rounded hover:bg-blue-200"
                      >
                        تاریخ انقضا
                      </button>
                    </div>
                  </div>
                  <textarea
                    v-model="editedTemplate.content"
                    rows="12"
                    class="w-full px-3 py-2 border-0 focus:outline-none focus:ring-0 resize-none"
                    placeholder="محتوای ایمیل را اینجا بنویسید..."
                  ></textarea>
                </div>
              </div>
            </div>

            <div v-else>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">متن پیامک</label>
                <div class="border border-gray-300 rounded-md">
                  <div class="bg-gray-50 px-3 py-2 border-b border-gray-300">
                    <div class="flex items-center space-x-4 space-x-reverse">
                      <span class="text-xs text-gray-600">{{ editedTemplate.content?.length || 0 }} کاراکتر</span>
                      <div class="border-r border-gray-300 h-4"></div>
                      <button 
                        @click="insertVariable('name')"
                        class="text-xs bg-blue-100 text-blue-800 px-2 py-1 rounded hover:bg-blue-200"
                      >
                        نام
                      </button>
                      <button 
                        @click="insertVariable('amount')"
                        class="text-xs bg-blue-100 text-blue-800 px-2 py-1 rounded hover:bg-blue-200"
                      >
                        مبلغ
                      </button>
                      <button 
                        @click="insertVariable('code')"
                        class="text-xs bg-blue-100 text-blue-800 px-2 py-1 rounded hover:bg-blue-200"
                      >
                        کد کارت
                      </button>
                      <button 
                        @click="insertVariable('expiry')"
                        class="text-xs bg-blue-100 text-blue-800 px-2 py-1 rounded hover:bg-blue-200"
                      >
                        تاریخ انقضا
                      </button>
                    </div>
                  </div>
                  <textarea
                    v-model="editedTemplate.content"
                    rows="6"
                    maxlength="160"
                    class="w-full px-3 py-2 border-0 focus:outline-none focus:ring-0 resize-none"
                    placeholder="متن پیامک را اینجا بنویسید..."
                  ></textarea>
                </div>
                <p class="text-xs text-gray-500 mt-1">حداکثر 160 کاراکتر</p>
              </div>
            </div>
          </div>

          <!-- تب پیش‌نمایش -->
          <div v-if="activeTab === 'preview'" class="space-y-4">
            <div class="bg-gray-50 rounded-lg p-6">
              <h4 class="text-sm font-medium text-gray-900 mb-2">پیش‌نمایش با داده‌های نمونه</h4>
              
              <div v-if="editedTemplate.type === 'email'" class="bg-white rounded-lg border border-gray-200 p-6">
                <div class="mb-2">
                  <span class="text-xs text-gray-500">از:</span>
                  <span class="text-sm font-medium">سیستم گیفت کارت &lt;noreply@giftcard.com&gt;</span>
                </div>
                <div class="mb-2">
                  <span class="text-xs text-gray-500">به:</span>
                  <span class="text-sm font-medium">علی احمدی &lt;ali@example.com&gt;</span>
                </div>
                <div class="mb-4">
                  <span class="text-xs text-gray-500">موضوع:</span>
                  <span class="text-sm font-medium">{{ previewSubject }}</span>
                </div>
                <div class="border-t border-gray-200 pt-4">
                  <div v-html="previewContent" class="text-sm text-gray-900 leading-relaxed"></div>
                </div>
              </div>

              <div v-else class="bg-white rounded-lg border border-gray-200 p-6">
                <div class="mb-2">
                  <span class="text-xs text-gray-500">از:</span>
                  <span class="text-sm font-medium">100000000</span>
                </div>
                <div class="mb-2">
                  <span class="text-xs text-gray-500">به:</span>
                  <span class="text-sm font-medium">09123456789</span>
                </div>
                <div class="border-t border-gray-200 pt-4">
                  <p class="text-sm text-gray-900">{{ previewContent }}</p>
                </div>
              </div>
            </div>
          </div>

          <!-- تب متغیرها -->
          <div v-if="activeTab === 'variables'" class="space-y-4">
            <div class="bg-gray-50 rounded-lg p-6">
              <h4 class="text-sm font-medium text-gray-900 mb-4">متغیرهای قابل استفاده</h4>
              
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div v-for="variable in availableVariables" :key="variable.key" class="bg-white rounded-lg border border-gray-200 p-3">
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="text-sm font-medium text-gray-900">{{ variable.name }}</p>
                      <p class="text-xs text-gray-600">{{ variable.description }}</p>
                    </div>
                    <button 
                      @click="insertVariable(variable.key)"
                      class="text-xs bg-blue-100 text-blue-800 px-2 py-1 rounded hover:bg-blue-200"
                    >
                      درج
                    </button>
                  </div>
                  <div class="mt-2">
                    <code class="text-xs bg-gray-100 px-2 py-1 rounded">{{ variable.example }}</code>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Footer -->
      <div class="flex justify-end space-x-3 space-x-reverse mt-6 pt-4 border-t border-gray-200">
        <button 
          @click="$emit('close')"
          class="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
        >
          انصراف
        </button>
        <button 
          @click="saveTemplate"
          :disabled="isSaving"
          class="px-4 py-2 text-sm font-medium text-white bg-blue-600 border border-transparent rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          <span v-if="isSaving">در حال ذخیره...</span>
          <span v-else>ذخیره قالب</span>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted, watch } from 'vue'

// Props
interface Template {
  id: number
  name: string
  description: string
  type: 'email' | 'sms'
  subject?: string
  content?: string
  variables: string[]
}

interface Props {
  template: Template | null
}

const props = defineProps<Props>()

// Emits
const emit = defineEmits<{
  close: []
  saved: [template: Template]
}>()

// Reactive data
const activeTab = ref('editor')
const isSaving = ref(false)

// متغیرهای قابل استفاده
const availableVariables = ref([
  {
    key: 'name',
    name: 'نام گیرنده',
    description: 'نام کامل گیرنده گیفت کارت',
    example: '{{name}}'
  },
  {
    key: 'sender_name',
    name: 'نام فرستنده',
    description: 'نام شخصی که کارت را ارسال کرده',
    example: '{{sender_name}}'
  },
  {
    key: 'amount',
    name: 'مبلغ کارت',
    description: 'مبلغ گیفت کارت به تومان',
    example: '{{amount}}'
  },
  {
    key: 'code',
    name: 'کد کارت',
    description: 'کد منحصر به فرد گیفت کارت',
    example: '{{code}}'
  },
  {
    key: 'expiry',
    name: 'تاریخ انقضا',
    description: 'تاریخ انقضای گیفت کارت',
    example: '{{expiry}}'
  },
  {
    key: 'remaining_amount',
    name: 'مبلغ باقی‌مانده',
    description: 'مبلغ باقی‌مانده در کارت',
    example: '{{remaining_amount}}'
  },
  {
    key: 'personal_message',
    name: 'پیام شخصی',
    description: 'پیام شخصی فرستنده',
    example: '{{personal_message}}'
  },
  {
    key: 'order_id',
    name: 'شماره سفارش',
    description: 'شماره سفارش خرید',
    example: '{{order_id}}'
  }
])

// قالب ویرایش شده
const editedTemplate = reactive<Template>({
  id: 0,
  name: '',
  description: '',
  type: 'email',
  subject: '',
  content: '',
  variables: []
})

// Computed properties
const previewSubject = computed(() => {
  if (!editedTemplate.subject) return ''
  return replaceVariables(editedTemplate.subject)
})

const previewContent = computed(() => {
  if (!editedTemplate.content) return ''
  return replaceVariables(editedTemplate.content)
})

// Methods
const replaceVariables = (text: string) => {
  let result = text
  
  // جایگزینی متغیرها با مقادیر نمونه
  const sampleData = {
    name: 'علی احمدی',
    sender_name: 'فاطمه محمدی',
    amount: '500,000 تومان',
    code: 'GC-BIRTHDAY-2024-001',
    expiry: '1403/12/29',
    remaining_amount: '250,000 تومان',
    personal_message: 'تولدت مبارک! امیدوارم این هدیه کوچک خوشت بیاد.',
    order_id: 'ORD-2024-001'
  }
  
  Object.entries(sampleData).forEach(([key, value]) => {
    result = result.replace(new RegExp(`{{${key}}}`, 'g'), value)
  })
  
  return result
}

const insertVariable = (variable: string) => {
  const variableText = `{{${variable}}}`
  
  if (activeTab.value === 'editor') {
    // درج متغیر در محتوای فعلی
    if (editedTemplate.content) {
      editedTemplate.content += variableText
    } else {
      editedTemplate.content = variableText
    }
  }
}

const saveTemplate = async () => {
  isSaving.value = true
  try {
    // شبیه‌سازی API call
    await new Promise(resolve => setTimeout(resolve, 1000))
    
    // استخراج متغیرهای استفاده شده
    const usedVariables: string[] = []
    const content = (editedTemplate.subject || '') + (editedTemplate.content || '')
    
    availableVariables.value.forEach(variable => {
      if (content.includes(`{{${variable.key}}}`)) {
        usedVariables.push(variable.key)
      }
    })
    
    editedTemplate.variables = usedVariables
    
    emit('saved', { ...editedTemplate })
  } catch (error) {
    console.error('خطا در ذخیره قالب:', error)
    alert('خطا در ذخیره قالب')
  } finally {
    isSaving.value = false
  }
}

// Lifecycle
onMounted(() => {
  if (props.template) {
    Object.assign(editedTemplate, props.template)
  }
})

watch(() => props.template, (newTemplate) => {
  if (newTemplate) {
    Object.assign(editedTemplate, newTemplate)
  }
})
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 

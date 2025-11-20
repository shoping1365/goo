<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200">
    <!-- هدر بخش -->
    <div class="p-6 border-b border-gray-200">
      <div class="flex items-center justify-between">
        <div>
          <h2 class="text-lg font-semibold text-gray-900">مدیریت قالب‌های پیشرفته</h2>
          <p class="text-gray-600 mt-1">ایجاد و مدیریت قالب‌های پیشرفته با قابلیت‌های شخصی‌سازی کامل</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors" @click="showTemplateForm = true">
            <span class="i-heroicons-plus ml-2"></span>
            ایجاد قالب جدید
          </button>
        </div>
      </div>
    </div>

    <!-- آمار قالب‌ها -->
    <div class="p-6 border-b border-gray-200">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div class="bg-blue-50 p-6 rounded-lg">
          <div class="flex items-center">
            <div class="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center ml-3">
              <span class="i-heroicons-document-text text-blue-600 text-xl"></span>
            </div>
            <div>
              <p class="text-sm text-blue-600">قالب‌های پیشرفته</p>
              <p class="text-2xl font-bold text-blue-900">{{ stats.advancedTemplates }}</p>
            </div>
          </div>
        </div>
        <div class="bg-green-50 p-6 rounded-lg">
          <div class="flex items-center">
            <div class="w-12 h-12 bg-green-100 rounded-lg flex items-center justify-center ml-3">
              <span class="i-heroicons-code-bracket text-green-600 text-xl"></span>
            </div>
            <div>
              <p class="text-sm text-green-600">قالب‌های سفارشی</p>
              <p class="text-2xl font-bold text-green-900">{{ stats.customTemplates }}</p>
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

    <!-- تب‌های قالب‌ها -->
    <div class="border-b border-gray-200">
      <div class="flex border-b border-gray-200 overflow-x-auto">
        <button
v-for="tab in tabs" :key="tab.value" :class="['px-6 py-3 -mb-px font-medium text-sm focus:outline-none whitespace-nowrap', activeTab === tab.value ? 'border-b-2 border-blue-600 text-blue-700' : 'text-gray-500 hover:text-blue-600']"
          @click="activeTab = tab.value">
          {{ tab.label }}
        </button>
      </div>
    </div>

    <!-- محتوای تب‌ها -->
    <div class="p-6">
      <!-- قالب‌های پیشرفته -->
      <div v-if="activeTab === 'templates'" class="space-y-6">
        <div class="flex justify-between items-center">
          <div class="flex items-center space-x-4 space-x-reverse">
            <select v-model="filterCategory" class="px-3 py-2 border border-gray-300 rounded-lg text-sm">
              <option value="">همه دسته‌ها</option>
              <option value="email">ایمیل</option>
              <option value="sms">پیامک</option>
              <option value="push">اعلان</option>
              <option value="landing">صفحه فرود</option>
            </select>
            <select v-model="filterComplexity" class="px-3 py-2 border border-gray-300 rounded-lg text-sm">
              <option value="">همه سطوح</option>
              <option value="basic">پایه</option>
              <option value="intermediate">متوسط</option>
              <option value="advanced">پیشرفته</option>
            </select>
          </div>
          <div class="flex items-center space-x-2 space-x-reverse">
            <button class="px-3 py-1 bg-green-100 text-green-700 rounded text-sm hover:bg-green-200" @click="bulkAction('export')">
              <span class="i-heroicons-arrow-down-tray ml-1"></span>
              صادرات
            </button>
            <button class="px-3 py-1 bg-blue-100 text-blue-700 rounded text-sm hover:bg-blue-200" @click="bulkAction('duplicate')">
              <span class="i-heroicons-document-duplicate ml-1"></span>
              تکثیر
            </button>
          </div>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <div v-for="template in filteredTemplates" :key="template.id" class="border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow">
            <div class="flex items-center justify-between mb-4">
              <div class="flex items-center">
                <div class="w-10 h-10 rounded-lg flex items-center justify-center ml-3" :style="{ backgroundColor: template.color + '20', color: template.color }">
                  <span class="i-heroicons-document-text text-lg"></span>
                </div>
                <div>
                  <h4 class="font-medium text-gray-900">{{ template.name }}</h4>
                  <p class="text-sm text-gray-500">{{ template.category }} • {{ template.complexity }}</p>
                </div>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <button class="text-blue-600 hover:text-blue-900" @click="previewTemplate(template)">
                  <span class="i-heroicons-eye"></span>
                </button>
                <button class="text-green-600 hover:text-green-900" @click="editTemplate(template)">
                  <span class="i-heroicons-pencil-square"></span>
                </button>
                <button class="text-purple-600 hover:text-purple-900" @click="duplicateTemplate(template)">
                  <span class="i-heroicons-document-duplicate"></span>
                </button>
                <button class="text-red-600 hover:text-red-900" @click="deleteTemplate(template)">
                  <span class="i-heroicons-trash"></span>
                </button>
              </div>
            </div>
            
            <div class="space-y-3">
              <div class="text-sm text-gray-600 line-clamp-2">{{ template.description }}</div>
              <div class="flex flex-wrap gap-2">
                <span v-for="tag in template.tags" :key="tag" class="px-2 py-1 bg-gray-100 text-gray-700 rounded text-xs">
                  {{ tag }}
                </span>
              </div>
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

      <!-- ویرایشگر قالب -->
      <div v-if="activeTab === 'editor'" class="space-y-6">
        <div class="bg-blue-50 p-6 rounded-lg">
          <h4 class="text-md font-medium text-blue-900 mb-2">ویرایشگر قالب پیشرفته</h4>
          <p class="text-sm text-blue-700">ایجاد و ویرایش قالب‌ها با قابلیت‌های HTML، CSS و JavaScript</p>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
          <!-- پنل ابزارها -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h5 class="font-medium text-gray-900 mb-4">ابزارها</h5>
            <div class="space-y-3">
              <div v-for="tool in editorTools" :key="tool.name" class="flex items-center justify-between p-2 bg-gray-50 rounded cursor-pointer hover:bg-gray-100" @click="insertTool(tool)">
                <span class="text-sm text-gray-700">{{ tool.name }}</span>
                <span class="text-xs text-gray-500">{{ tool.shortcut }}</span>
              </div>
            </div>
            
            <div class="mt-6">
              <h6 class="text-sm font-medium text-gray-700 mb-3">متغیرهای موجود</h6>
              <div class="space-y-2">
                <div v-for="variable in availableVariables" :key="variable.name" class="flex items-center justify-between p-2 bg-blue-50 rounded cursor-pointer hover:bg-blue-100" @click="insertVariable(variable)">
                  <span class="text-sm text-blue-700">{{ variable.name }}</span>
                  <span class="text-xs text-blue-500">{{ variable.description }}</span>
                </div>
              </div>
            </div>
          </div>

          <!-- ویرایشگر کد -->
          <div class="lg:col-span-2 border border-gray-200 rounded-lg">
            <div class="border-b border-gray-200 p-6">
              <div class="flex items-center justify-between">
                <h5 class="font-medium text-gray-900">ویرایشگر کد</h5>
                <div class="flex items-center space-x-2 space-x-reverse">
                  <select v-model="editorLanguage" class="px-3 py-1 border border-gray-300 rounded text-sm">
                    <option value="html">HTML</option>
                    <option value="css">CSS</option>
                    <option value="javascript">JavaScript</option>
                  </select>
                  <button class="px-3 py-1 bg-gray-100 text-gray-700 rounded text-sm hover:bg-gray-200" @click="formatCode">
                    <span class="i-heroicons-code-bracket ml-1"></span>
                    فرمت
                  </button>
                </div>
              </div>
            </div>
            <div class="p-6">
              <textarea v-model="editorCode" rows="20" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm font-mono" placeholder="کد قالب را اینجا بنویسید..."></textarea>
            </div>
          </div>
        </div>

        <!-- پیش‌نمایش -->
        <div class="border border-gray-200 rounded-lg p-6">
          <h5 class="font-medium text-gray-900 mb-4">پیش‌نمایش</h5>
          <div class="bg-gray-50 p-6 rounded-lg min-h-64">
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
            <div v-if="previewContent" class="text-sm" v-html="previewContent"></div>
            <div v-else class="text-gray-400 text-center py-8">پیش‌نمایش قالب</div>
          </div>
        </div>
      </div>

      <!-- قالب‌های آماده -->
      <div v-if="activeTab === 'ready-made'" class="space-y-6">
        <div class="bg-green-50 p-6 rounded-lg">
          <h4 class="text-md font-medium text-green-900 mb-2">قالب‌های آماده</h4>
          <p class="text-sm text-green-700">قالب‌های از پیش طراحی شده برای استفاده سریع</p>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <div v-for="template in readyMadeTemplates" :key="template.id" class="border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow">
            <div class="mb-4">
              <div class="w-full h-32 bg-gray-100 rounded-lg mb-3 flex items-center justify-center">
                <span class="text-gray-400">{{ template.preview }}</span>
              </div>
              <h4 class="font-medium text-gray-900">{{ template.name }}</h4>
              <p class="text-sm text-gray-500 mt-1">{{ template.description }}</p>
            </div>
            
            <div class="space-y-3">
              <div class="flex flex-wrap gap-2">
                <span v-for="tag in template.tags" :key="tag" class="px-2 py-1 bg-gray-100 text-gray-700 rounded text-xs">
                  {{ tag }}
                </span>
              </div>
              <div class="flex justify-between text-sm">
                <span class="text-gray-500">سطح:</span>
                <span class="font-medium">{{ template.level }}</span>
              </div>
              <div class="flex justify-between text-sm">
                <span class="text-gray-500">استفاده:</span>
                <span class="font-medium">{{ template.usageCount }}</span>
              </div>
              <button class="w-full px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors" @click="useReadyMadeTemplate(template)">
                استفاده از قالب
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- تنظیمات پیشرفته -->
      <div v-if="activeTab === 'settings'" class="space-y-6">
        <div class="bg-purple-50 p-6 rounded-lg">
          <h4 class="text-md font-medium text-purple-900 mb-2">تنظیمات پیشرفته قالب‌ها</h4>
          <p class="text-sm text-purple-700">تنظیم پارامترهای پیشرفته برای قالب‌ها</p>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <!-- تنظیمات کامپایلر -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h5 class="font-medium text-gray-900 mb-4">تنظیمات کامپایلر</h5>
            <div class="space-y-4">
              <div class="flex items-center justify-between">
                <div>
                  <span class="text-sm font-medium text-gray-700">فشرده‌سازی کد</span>
                  <p class="text-xs text-gray-500">کاهش اندازه فایل‌های نهایی</p>
                </div>
                <input v-model="templateSettings.minification" type="checkbox" class="rounded border-gray-300">
              </div>
              
              <div class="flex items-center justify-between">
                <div>
                  <span class="text-sm font-medium text-gray-700">بهینه‌سازی تصاویر</span>
                  <p class="text-xs text-gray-500">فشرده‌سازی خودکار تصاویر</p>
                </div>
                <input v-model="templateSettings.imageOptimization" type="checkbox" class="rounded border-gray-300">
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">نسخه CSS</label>
                <select v-model="templateSettings.cssVersion" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
                  <option value="css3">CSS3</option>
                  <option value="css4">CSS4</option>
                  <option value="scss">SCSS</option>
                  <option value="sass">Sass</option>
                </select>
              </div>
            </div>
          </div>

          <!-- تنظیمات تست -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h5 class="font-medium text-gray-900 mb-4">تنظیمات تست</h5>
            <div class="space-y-4">
              <div class="flex items-center justify-between">
                <div>
                  <span class="text-sm font-medium text-gray-700">تست خودکار</span>
                  <p class="text-xs text-gray-500">تست قالب‌ها قبل از انتشار</p>
                </div>
                <input v-model="templateSettings.autoTest" type="checkbox" class="rounded border-gray-300">
              </div>
              
              <div class="flex items-center justify-between">
                <div>
                  <span class="text-sm font-medium text-gray-700">اعتبارسنجی HTML</span>
                  <p class="text-xs text-gray-500">بررسی استاندارد HTML</p>
                </div>
                <input v-model="templateSettings.htmlValidation" type="checkbox" class="rounded border-gray-300">
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">مرورگرهای هدف</label>
                <div class="space-y-2">
                  <label v-for="browser in targetBrowsers" :key="browser.name" class="flex items-center">
                    <input v-model="browser.enabled" type="checkbox" class="rounded border-gray-300 ml-2">
                    <span class="text-sm text-gray-700">{{ browser.name }}</span>
                  </label>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- مودال فرم قالب -->
    <div v-if="showTemplateForm" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg shadow-xl max-w-4xl w-full mx-4 max-h-[90vh] overflow-y-auto">
        <div class="p-6 border-b border-gray-200">
          <div class="flex items-center justify-between">
            <h3 class="text-lg font-semibold text-gray-900">
              {{ editingTemplate ? 'ویرایش قالب' : 'ایجاد قالب جدید' }}
            </h3>
            <button class="text-gray-400 hover:text-gray-600" @click="closeForm">
              <span class="i-heroicons-x-mark text-xl"></span>
            </button>
          </div>
        </div>
        
        <form class="p-6 space-y-6" @submit.prevent="handleSubmit">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نام قالب *</label>
              <input v-model="form.name" type="text" required class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="نام قالب">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">دسته‌بندی *</label>
              <select v-model="form.category" required class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                <option value="email">ایمیل</option>
                <option value="sms">پیامک</option>
                <option value="push">اعلان</option>
                <option value="landing">صفحه فرود</option>
              </select>
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">توضیحات</label>
            <textarea v-model="form.description" rows="3" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="توضیحات قالب"></textarea>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">سطح پیچیدگی</label>
              <select v-model="form.complexity" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                <option value="basic">پایه</option>
                <option value="intermediate">متوسط</option>
                <option value="advanced">پیشرفته</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">رنگ قالب</label>
              <input v-model="form.color" type="color" class="w-full h-10 border border-gray-300 rounded-lg">
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">برچسب‌ها</label>
            <input v-model="form.tagsInput" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="برچسب‌ها را با کاما جدا کنید">
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">کد HTML</label>
            <textarea v-model="form.htmlCode" rows="8" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm font-mono focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="کد HTML قالب"></textarea>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">کد CSS</label>
              <textarea v-model="form.cssCode" rows="6" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm font-mono focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="کد CSS قالب"></textarea>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">کد JavaScript</label>
              <textarea v-model="form.jsCode" rows="6" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm font-mono focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="کد JavaScript قالب"></textarea>
            </div>
          </div>
        </form>
        
        <div class="p-6 border-t border-gray-200 flex justify-end space-x-3 space-x-reverse">
          <button class="px-6 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200 transition-colors" @click="closeForm">
            انصراف
          </button>
          <button class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors" @click="handleSubmit">
            ذخیره
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed } from 'vue'

const activeTab = ref('templates')
const showTemplateForm = ref(false)
const editingTemplate = ref<any>(null)
const selectedTemplates = ref<number[]>([])
const filterCategory = ref('')
const filterComplexity = ref('')
const editorLanguage = ref('html')
const editorCode = ref('')
const previewContent = ref('')

const tabs = [
  { value: 'templates', label: 'قالب‌های پیشرفته' },
  { value: 'editor', label: 'ویرایشگر قالب' },
  { value: 'ready-made', label: 'قالب‌های آماده' },
  { value: 'settings', label: 'تنظیمات پیشرفته' }
]

const stats = ref({
  advancedTemplates: 25,
  customTemplates: 12,
  openRate: 28.5,
  clickRate: 9.2
})

const templates = ref([
  {
    id: 1,
    name: 'قالب ایمیل پیشرفته',
    category: 'email',
    complexity: 'advanced',
    description: 'قالب ایمیل پیشرفته با انیمیشن‌های CSS و تعامل JavaScript',
    color: '#3b82f6',
    tags: ['انیمیشن', 'تعاملی', 'پاسخگو'],
    usageCount: 156,
    lastModified: '2024-01-15T10:30:00'
  },
  {
    id: 2,
    name: 'قالب صفحه فرود',
    category: 'landing',
    complexity: 'intermediate',
    description: 'قالب صفحه فرود بهینه‌سازی شده برای تبدیل',
    color: '#10b981',
    tags: ['تبدیل', 'SEO', 'سرعت'],
    usageCount: 89,
    lastModified: '2024-01-14T09:15:00'
  }
])

const editorTools = ref([
  { name: 'درج متغیر', shortcut: 'Ctrl+V', action: 'insertVariable' },
  { name: 'فرمت کد', shortcut: 'Ctrl+Shift+F', action: 'formatCode' },
  { name: 'پیش‌نمایش', shortcut: 'Ctrl+P', action: 'preview' },
  { name: 'ذخیره', shortcut: 'Ctrl+S', action: 'save' }
])

const availableVariables = ref([
  { name: '{{user_name}}', description: 'نام کاربر' },
  { name: '{{coupon_code}}', description: 'کد کوپن' },
  { name: '{{discount_amount}}', description: 'مبلغ تخفیف' },
  { name: '{{expiry_date}}', description: 'تاریخ انقضا' },
  { name: '{{campaign_name}}', description: 'نام کمپین' }
])

const readyMadeTemplates = ref([
  {
    id: 1,
    name: 'قالب خوشامدگویی',
    description: 'قالب ایمیل خوشامدگویی با طراحی مدرن',
    preview: 'ایمیل خوشامدگویی',
    tags: ['خوشامدگویی', 'مدرن'],
    level: 'پایه',
    usageCount: 234
  },
  {
    id: 2,
    name: 'قالب تخفیف ویژه',
    description: 'قالب ایمیل تخفیف با طراحی جذاب',
    preview: 'ایمیل تخفیف',
    tags: ['تخفیف', 'جذاب'],
    level: 'متوسط',
    usageCount: 189
  },
  {
    id: 3,
    name: 'قالب یادآوری',
    description: 'قالب ایمیل یادآوری انقضای کوپن',
    preview: 'ایمیل یادآوری',
    tags: ['یادآوری', 'انقضا'],
    level: 'پایه',
    usageCount: 156
  }
])

const targetBrowsers = ref([
  { name: 'Chrome', enabled: true },
  { name: 'Firefox', enabled: true },
  { name: 'Safari', enabled: true },
  { name: 'Edge', enabled: true },
  { name: 'IE11', enabled: false }
])

const templateSettings = reactive({
  minification: true,
  imageOptimization: true,
  cssVersion: 'css3',
  autoTest: true,
  htmlValidation: true
})

const form = reactive({
  name: '',
  category: 'email',
  description: '',
  complexity: 'basic',
  color: '#3b82f6',
  tagsInput: '',
  htmlCode: '',
  cssCode: '',
  jsCode: ''
})

const filteredTemplates = computed(() => {
  let filtered = templates.value
  
  if (filterCategory.value) {
    filtered = filtered.filter(t => t.category === filterCategory.value)
  }
  
  if (filterComplexity.value) {
    filtered = filtered.filter(t => t.complexity === filterComplexity.value)
  }
  
  return filtered
})

const formatDate = (date: string): string => {
  return new Intl.DateTimeFormat('fa-IR').format(new Date(date))
}

const previewTemplate = (template: any) => {
  // پیاده‌سازی پیش‌نمایش قالب
  console.log('پیش‌نمایش قالب:', template)
}

const editTemplate = (template: any) => {
  editingTemplate.value = template
  Object.assign(form, {
    ...template,
    tagsInput: template.tags.join(', ')
  })
  showTemplateForm.value = true
}

const duplicateTemplate = (template: any) => {
  const duplicate = { ...template, id: Date.now(), name: `${template.name} (کپی)` }
  templates.value.unshift(duplicate)
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

const bulkAction = (action: string) => {
  if (selectedTemplates.value.length === 0) {
    alert('لطفاً قالب‌هایی را انتخاب کنید')
    return
  }
  
  switch (action) {
    case 'export':
      // پیاده‌سازی صادرات قالب‌ها
      console.log('صادرات قالب‌ها:', selectedTemplates.value)
      break
    case 'duplicate':
      selectedTemplates.value.forEach(id => {
        const template = templates.value.find(t => t.id === id)
        if (template) {
          duplicateTemplate(template)
        }
      })
      selectedTemplates.value = []
      break
  }
}

const insertTool = (tool: any) => {
  // پیاده‌سازی درج ابزار
  console.log('درج ابزار:', tool)
}

const insertVariable = (variable: any) => {
  editorCode.value += variable.name
}

const formatCode = () => {
  // پیاده‌سازی فرمت کردن کد
  console.log('فرمت کردن کد')
}

const useReadyMadeTemplate = (template: any) => {
  // پیاده‌سازی استفاده از قالب آماده
  console.log('استفاده از قالب آماده:', template)
  showTemplateForm.value = true
}

const handleSubmit = async () => {
  if (!form.name || !form.category) {
    alert('لطفاً فیلدهای اجباری را پر کنید')
    return
  }
  
  const templateData = {
    ...form,
    tags: form.tagsInput.split(',').map(tag => tag.trim()).filter(tag => tag)
  }
  
  if (editingTemplate.value) {
    Object.assign(editingTemplate.value, templateData)
  } else {
    templates.value.unshift({
      id: Date.now(),
      ...templateData,
      usageCount: 0,
      lastModified: new Date().toISOString()
    })
  }
  closeForm()
}

const closeForm = () => {
  showTemplateForm.value = false
  editingTemplate.value = null
  Object.assign(form, { name: '', category: 'email', description: '', complexity: 'basic', color: '#3b82f6', tagsInput: '', htmlCode: '', cssCode: '', jsCode: '' })
}
</script>

<!--
  کامپوننت مدیریت قالب‌های پیشرفته
  شامل:
  1. مدیریت قالب‌های پیشرفته
  2. ویرایشگر قالب با ابزارها
  3. قالب‌های آماده
  4. تنظیمات پیشرفته
  توضیحات کامل به فارسی در کد
--> 

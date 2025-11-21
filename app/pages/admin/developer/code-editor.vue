<template>
  <div class="min-h-screen bg-gray-50 p-6">
    <div class="max-w-7xl mx-auto">
      <div class="mb-8">
        <h1 class="text-3xl font-bold text-gray-900 mb-2">ویرایشگر کد</h1>
        <p class="text-gray-600">ابزارهای ویرایش و دیباگ کد</p>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-4 gap-8">
        <div class="bg-white rounded-lg shadow">
          <div class="p-6 border-b border-gray-200">
            <h2 class="text-xl font-semibold text-gray-900">فایل‌ها</h2>
          </div>
          <div class="p-6">
            <div class="space-y-4">
              <div class="space-y-2">
                <div v-for="file in fileTree" :key="file.path" class="space-y-1">
                  <div 
                    :class="[
                      'flex items-center space-x-2 space-x-reverse p-2 rounded cursor-pointer hover:bg-gray-100',
                      selectedFile?.path === file.path ? 'bg-blue-50 border border-blue-200' : ''
                    ]"
                    @click="selectFile(file)"
                  >
                    <svg v-if="file.type === 'folder'" class="w-4 h-4 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-5l-2-2H5a2 2 0 00-2 2z"></path>
                    </svg>
                    <svg v-else class="w-4 h-4 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
                    </svg>
                    <span class="text-sm">{{ file.name }}</span>
                  </div>
                </div>
              </div>

              <div class="space-y-2 pt-4 border-t">
                <button 
                  class="w-full bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="createNewFile"
                >
                  فایل جدید
                </button>
                <button 
                  :disabled="!selectedFile"
                  class="w-full bg-green-600 hover:bg-green-700 disabled:bg-gray-400 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="saveFile"
                >
                  ذخیره
                </button>
                <button 
                  :disabled="!selectedFile"
                  class="w-full bg-purple-600 hover:bg-purple-700 disabled:bg-gray-400 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="formatCode"
                >
                  فرمت کد
                </button>
              </div>
            </div>
          </div>
        </div>

        <div class="lg:col-span-3 bg-white rounded-lg shadow">
          <div class="p-6 border-b border-gray-200">
            <div class="flex items-center justify-between">
              <h2 class="text-xl font-semibold text-gray-900">
                {{ selectedFile?.name || 'انتخاب فایل' }}
              </h2>
              <div class="flex items-center space-x-2 space-x-reverse">
                <select 
                  v-model="editorTheme"
                  class="border border-gray-300 rounded-lg px-3 py-1 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                  <option value="vs-dark">تاریک</option>
                  <option value="vs-light">روشن</option>
                  <option value="hc-black">تاریک بالا</option>
                </select>
                <select 
                  v-model="editorLanguage"
                  class="border border-gray-300 rounded-lg px-3 py-1 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                  <option value="javascript">JavaScript</option>
                  <option value="typescript">TypeScript</option>
                  <option value="html">HTML</option>
                  <option value="css">CSS</option>
                  <option value="json">JSON</option>
                  <option value="go">Go</option>
                  <option value="vue">Vue</option>
                </select>
              </div>
            </div>
          </div>
          <div class="p-6">
            <div v-if="selectedFile" class="border rounded-lg overflow-hidden">
              <textarea
                v-model="codeContent"
                class="w-full h-96 p-6 font-mono text-sm bg-gray-900 text-gray-100 focus:outline-none resize-none"
                :placeholder="getPlaceholder()"
                @input="onCodeChange"
              ></textarea>
            </div>
            <div v-else class="text-center py-12">
              <svg class="w-12 h-12 mx-auto text-gray-300 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
              </svg>
              <p class="text-gray-500">فایلی برای ویرایش انتخاب نشده است</p>
            </div>
          </div>
        </div>
      </div>

      <div class="mt-8 bg-white rounded-lg shadow">
        <div class="p-6 border-b border-gray-200">
          <div class="flex items-center justify-between">
            <h2 class="text-xl font-semibold text-gray-900">ترمینال</h2>
            <div class="flex items-center space-x-2 space-x-reverse">
              <button 
                class="bg-red-600 hover:bg-red-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                @click="clearTerminal"
              >
                پاک کردن
              </button>
              <button 
                class="bg-green-600 hover:bg-green-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                @click="executeCommand"
              >
                اجرا
              </button>
            </div>
          </div>
        </div>
        <div class="p-6">
          <div class="bg-gray-900 rounded-lg p-6 h-64 overflow-y-auto">
            <div class="space-y-2">
              <div v-for="(line, index) in terminalOutput" :key="index" class="text-green-400 font-mono text-sm">
                <span v-if="line.type === 'command'" class="text-blue-400">$ </span>
                <span v-if="line.type === 'output'" class="text-gray-300">> </span>
                <span v-if="line.type === 'error'" class="text-red-400">! </span>
                {{ line.content }}
              </div>
            </div>
            <div class="flex items-center mt-4">
              <span class="text-green-400 font-mono text-sm mr-2">$</span>
              <input
                v-model="currentCommand"
                type="text"
                class="flex-1 bg-transparent text-gray-100 font-mono text-sm focus:outline-none"
                placeholder="دستور را وارد کنید..."
                @keyup.enter="executeCommand"
              />
            </div>
          </div>
        </div>
      </div>

      <div class="mt-8 bg-white rounded-lg shadow">
        <div class="p-6 border-b border-gray-200">
          <h2 class="text-xl font-semibold text-gray-900">تحلیل کد</h2>
        </div>
        <div class="p-6">
          <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
            <div class="space-y-4">
              <h3 class="font-medium text-gray-900">خطاهای نحوی</h3>
              <div class="space-y-2">
                <div v-for="error in syntaxErrors" :key="error.id" class="flex items-center space-x-2 space-x-reverse p-3 bg-red-50 border border-red-200 rounded-lg">
                  <svg class="w-4 h-4 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
                  </svg>
                  <div>
                    <p class="text-sm font-medium text-red-800">{{ error.message }}</p>
                    <p class="text-xs text-red-600">خط {{ error.line }}</p>
                  </div>
                </div>
              </div>
            </div>

            <div class="space-y-4">
              <h3 class="font-medium text-gray-900">هشدارها</h3>
              <div class="space-y-2">
                <div v-for="warning in warnings" :key="warning.id" class="flex items-center space-x-2 space-x-reverse p-3 bg-yellow-50 border border-yellow-200 rounded-lg">
                  <svg class="w-4 h-4 text-yellow-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
                  </svg>
                  <div>
                    <p class="text-sm font-medium text-yellow-800">{{ warning.message }}</p>
                    <p class="text-xs text-yellow-600">خط {{ warning.line }}</p>
                  </div>
                </div>
              </div>
            </div>

            <div class="space-y-4">
              <h3 class="font-medium text-gray-900">معیارهای کد</h3>
              <div class="space-y-3">
                <div class="flex justify-between items-center">
                  <span class="text-sm text-gray-600">تعداد خطوط:</span>
                  <span class="text-sm font-medium">{{ codeMetrics.lines }}</span>
                </div>
                <div class="flex justify-between items-center">
                  <span class="text-sm text-gray-600">توابع:</span>
                  <span class="text-sm font-medium">{{ codeMetrics.functions }}</span>
                </div>
                <div class="flex justify-between items-center">
                  <span class="text-sm text-gray-600">متغیرها:</span>
                  <span class="text-sm font-medium">{{ codeMetrics.variables }}</span>
                </div>
                <div class="flex justify-between items-center">
                  <span class="text-sm text-gray-600">پیچیدگی:</span>
                  <span class="text-sm font-medium">{{ codeMetrics.complexity }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { reactive, ref, watch } from 'vue';

definePageMeta({
  layout: 'admin-main',
  middleware: ['developer-only']
});

// استفاده از useAuth برای چک کردن پرمیژن‌ها
// const { user, hasPermission } = useAuth()

const selectedFile = ref(null)
const codeContent = ref('')
const editorTheme = ref('vs-dark')
const editorLanguage = ref('javascript')
const currentCommand = ref('')
const terminalOutput = ref([])

const fileTree = ref([
  { name: 'src', type: 'folder', path: '/src' },
  { name: 'components', type: 'folder', path: '/src/components' },
  { name: 'App.vue', type: 'file', path: '/src/App.vue' },
  { name: 'main.js', type: 'file', path: '/src/main.js' },
  { name: 'utils', type: 'folder', path: '/src/utils' },
  { name: 'helpers.js', type: 'file', path: '/src/utils/helpers.js' },
  { name: 'api', type: 'folder', path: '/src/api' },
  { name: 'index.js', type: 'file', path: '/src/api/index.js' }
])

const syntaxErrors = ref([])
const warnings = ref([])
const codeMetrics = reactive({
  lines: 0,
  functions: 0,
  variables: 0,
  complexity: 0
})

function selectFile(file) {
  if (file.type === 'file') {
    selectedFile.value = file
    loadFileContent(file)
    analyzeCode()
  }
}

function loadFileContent(file) {
  const mockContents = {
    '/src/App.vue': '<template>\n  <div id="app">\n    <h1>{{ message }}</h1>\n  </div>\n</template>\n\n<script>\nexport default {\n  name: \'App\',\n  data() {\n    return {\n      message: \'Hello Vue!\'\n    }\n  }\n}\n<\/script>',
    '/src/main.js': 'import { createApp } from \'vue\'\nimport App from \'./App.vue\'\n\ncreateApp(App).mount(\'#app\')',
    '/src/utils/helpers.js': 'export function formatDate(date) {\n  return new Date(date).toLocaleDateString()\n}\n\nexport function calculateTotal(items) {\n  return items.reduce((sum, item) => sum + item.price, 0)\n}',
    '/src/api/index.js': 'import axios from \'axios\'\n\nconst api = axios.create({\n  baseURL: process.env.VUE_APP_API_URL\n})\n\nexport default api'
  }
  
  codeContent.value = mockContents[file.path] || '// فایل خالی'
  updateLanguage(file.name)
}

function updateLanguage(fileName) {
  const extension = fileName.split('.').pop()
  const languageMap = {
    'js': 'javascript',
    'ts': 'typescript',
    'vue': 'vue',
    'html': 'html',
    'css': 'css',
    'json': 'json',
    'go': 'go'
  }
  editorLanguage.value = languageMap[extension] || 'javascript'
}

function createNewFile() {
  const fileName = prompt('نام فایل را وارد کنید:')
  if (fileName) {
    const newFile = {
      name: fileName,
      type: 'file',
      path: `/src/${fileName}`
    }
    fileTree.value.push(newFile)
    selectFile(newFile)
  }
}

function saveFile() {
  if (selectedFile.value) {
    terminalOutput.value.push({
      type: 'command',
      content: `git add ${selectedFile.value.path}`
    })
    terminalOutput.value.push({
      type: 'output',
      content: `فایل ${selectedFile.value.name} با موفقیت ذخیره شد`
    })
  }
}

function formatCode() {
  if (selectedFile.value) {
    codeContent.value = codeContent.value.trim()
    terminalOutput.value.push({
      type: 'output',
      content: `کد ${selectedFile.value.name} فرمت شد`
    })
  }
}

function executeCommand() {
  if (currentCommand.value.trim()) {
    terminalOutput.value.push({
      type: 'command',
      content: currentCommand.value
    })
    
    const command = currentCommand.value.toLowerCase()
    if (command.includes('npm')) {
      terminalOutput.value.push({
        type: 'output',
        content: 'دستور npm اجرا شد'
      })
    } else if (command.includes('git')) {
      terminalOutput.value.push({
        type: 'output',
        content: 'دستور git اجرا شد'
      })
    } else {
      terminalOutput.value.push({
        type: 'error',
        content: 'دستور شناخته نشد'
      })
    }
    
    currentCommand.value = ''
  }
}

function clearTerminal() {
  terminalOutput.value = []
}

function analyzeCode() {
  const lines = codeContent.value.split('\n').length
  const functions = (codeContent.value.match(/function|=>/g) || []).length
  const variables = (codeContent.value.match(/const|let|var/g) || []).length
  const complexity = Math.min(10, Math.floor(lines / 10) + functions)
  
  codeMetrics.lines = lines
  codeMetrics.functions = functions
  codeMetrics.variables = variables
  codeMetrics.complexity = complexity
  
  syntaxErrors.value = []
  warnings.value = []
  
  if (lines > 50) {
    warnings.value.push({
      id: 1,
      message: 'فایل خیلی بزرگ است',
      line: 1
    })
  }
  
  if (complexity > 5) {
    warnings.value.push({
      id: 2,
      message: 'پیچیدگی کد بالا است',
      line: 1
    })
  }
}

function onCodeChange() {
  analyzeCode()
}

function getPlaceholder() {
  const placeholders = {
    javascript: '// کد JavaScript خود را اینجا بنویسید',
    typescript: '// کد TypeScript خود را اینجا بنویسید',
    html: '<!-- کد HTML خود را اینجا بنویسید -->',
    css: '/* کد CSS خود را اینجا بنویسید */',
    json: '{\n  // کد JSON خود را اینجا بنویسید\n}',
    go: '// کد Go خود را اینجا بنویسید',
    vue: '<template>\n  <!-- کد Vue خود را اینجا بنویسید -->\n</template>'
  }
  return placeholders[editorLanguage.value] || 'کد خود را اینجا بنویسید'
}

watch(codeContent, () => {
  analyzeCode()
})
</script> 

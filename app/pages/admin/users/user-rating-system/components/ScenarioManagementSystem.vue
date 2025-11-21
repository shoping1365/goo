<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="bg-white rounded-lg shadow px-4 py-4">
      <div class="flex items-center justify-between">
        <div>
          <h2 class="text-xl font-bold text-gray-900">سیستم مدیریت سناریوها</h2>
          <p class="mt-1 text-sm text-gray-500">تعریف و مدیریت سناریوهای مختلف امتیازدهی و رتبه‌بندی</p>
        </div>
        <div class="flex space-x-3 space-x-reverse">
          <button class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition-colors" @click="createNewScenario">
            ایجاد سناریو جدید
          </button>
          <button class="bg-green-600 text-white px-4 py-2 rounded-md hover:bg-green-700 transition-colors" @click="exportScenarios">
            خروجی سناریوها
          </button>
        </div>
      </div>
    </div>

    <!-- Scenario Templates -->
    <div class="bg-white rounded-lg shadow px-4 py-4">
      <h3 class="text-lg font-medium text-gray-900 mb-4">قالب‌های آماده سناریو</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gapx-4 py-4">
        <div v-for="template in scenarioTemplates" :key="template.id" class="border border-gray-200 rounded-lg px-4 py-4 hover:shadow-md transition-shadow">
          <div class="flex items-center justify-between mb-3">
            <h4 class="text-lg font-medium text-gray-900">{{ template.name }}</h4>
            <span :class="getTemplateTypeClass(template.type)" class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium">
              {{ getTemplateTypeText(template.type) }}
            </span>
          </div>
          <p class="text-sm text-gray-600 mb-4">{{ template.description }}</p>
          <div class="space-y-2 mb-4">
            <div class="flex justify-between text-sm">
              <span class="text-gray-500">سطح پیچیدگی:</span>
              <span class="text-gray-900">{{ template.complexity }}</span>
            </div>
            <div class="flex justify-between text-sm">
              <span class="text-gray-500">تعداد قوانین:</span>
              <span class="text-gray-900">{{ template.rulesCount }}</span>
            </div>
            <div class="flex justify-between text-sm">
              <span class="text-gray-500">زمان راه‌اندازی:</span>
              <span class="text-gray-900">{{ template.setupTime }}</span>
            </div>
          </div>
          <div class="flex space-x-2 space-x-reverse">
            <button class="bg-blue-600 text-white px-3 py-1 rounded-md text-sm hover:bg-blue-700" @click="useTemplate(template)">
              استفاده از قالب
            </button>
            <button class="bg-gray-600 text-white px-3 py-1 rounded-md text-sm hover:bg-gray-700" @click="previewTemplate(template)">
              پیش‌نمایش
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Active Scenarios -->
    <div class="bg-white rounded-lg shadow px-4 py-4">
      <h3 class="text-lg font-medium text-gray-900 mb-4">سناریوهای فعال</h3>
      
      <!-- Search and Filter -->
      <div class="flex flex-col sm:flex-row gapx-4 py-4 mb-6">
        <div class="flex-1">
          <input v-model="searchQuery" type="text" placeholder="جستجو در سناریوها..." class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <select v-model="filterStatus" class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
          <option value="">همه وضعیت‌ها</option>
          <option value="active">فعال</option>
          <option value="inactive">غیرفعال</option>
          <option value="draft">پیش‌نویس</option>
          <option value="testing">در حال تست</option>
        </select>
        <select v-model="filterType" class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
          <option value="">همه انواع</option>
          <option value="basic">پایه</option>
          <option value="advanced">پیشرفته</option>
          <option value="custom">سفارشی</option>
        </select>
      </div>

      <!-- Scenarios Table -->
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نام سناریو</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ ایجاد</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کاربران تحت تأثیر</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملکرد</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="scenario in paginatedScenarios" :key="scenario.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <div>
                  <div class="text-sm font-medium text-gray-900">{{ scenario.name }}</div>
                  <div class="text-sm text-gray-500">{{ scenario.description }}</div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getTypeClass(scenario.type)" class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium">
                  {{ getTypeText(scenario.type) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getStatusClass(scenario.status)" class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium">
                  {{ getStatusText(scenario.status) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ formatDate(scenario.createdAt) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ scenario.affectedUsers.toLocaleString() }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="w-16 bg-gray-200 rounded-full h-2 mr-2">
                    <div :style="{ width: scenario.performance + '%' }" class="bg-green-600 h-2 rounded-full"></div>
                  </div>
                  <span class="text-sm text-gray-900">{{ scenario.performance }}%</span>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3" @click="editScenario(scenario)">ویرایش</button>
                <button :class="scenario.status === 'active' ? 'text-red-600 hover:text-red-900' : 'text-green-600 hover:text-green-900'" class="ml-3" @click="toggleScenario(scenario)">
                  {{ scenario.status === 'active' ? 'غیرفعال' : 'فعال' }}
                </button>
                <button class="text-purple-600 hover:text-purple-900 ml-3" @click="duplicateScenario(scenario)">کپی</button>
                <button class="text-red-600 hover:text-red-900" @click="deleteScenario(scenario)">حذف</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Pagination -->
      <div class="flex items-center justify-between mt-6">
        <div class="flex items-center space-x-2 space-x-reverse">
          <span class="text-sm text-gray-700">نمایش {{ pagination.start }} تا {{ pagination.end }} از {{ pagination.total }} سناریو</span>
        </div>
        <div class="flex space-x-2 space-x-reverse">
          <button :disabled="currentPage === 1" class="px-3 py-1 border border-gray-300 rounded-md text-sm disabled:opacity-50" @click="previousPage">
            قبلی
          </button>
          <button v-for="page in visiblePages" :key="page" :class="page === currentPage ? 'bg-blue-600 text-white' : 'bg-white text-gray-700'" class="px-3 py-1 border border-gray-300 rounded-md text-sm" @click="goToPage(page)">
            {{ page }}
          </button>
          <button :disabled="currentPage === totalPages" class="px-3 py-1 border border-gray-300 rounded-md text-sm disabled:opacity-50" @click="nextPage">
            بعدی
          </button>
        </div>
      </div>
    </div>

    <!-- Scenario Performance Analytics -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gapx-4 py-4">
      <div class="bg-white rounded-lg shadow px-4 py-4">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-green-100 rounded-full flex items-center justify-center">
              <svg class="w-5 h-5 text-green-600" fill="currentColor" viewBox="0 0 20 20">
                <path d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-500">سناریوهای فعال</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.activeScenarios }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow px-4 py-4">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center">
              <svg class="w-5 h-5 text-blue-600" fill="currentColor" viewBox="0 0 20 20">
                <path d="M13 6a3 3 0 11-6 0 3 3 0 016 0zM18 8a2 2 0 11-4 0 2 2 0 014 0zM14 15a4 4 0 00-8 0v3h8v-3z"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-500">کاربران تحت تأثیر</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.totalAffectedUsers.toLocaleString() }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow px-4 py-4">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-yellow-100 rounded-full flex items-center justify-center">
              <svg class="w-5 h-5 text-yellow-600" fill="currentColor" viewBox="0 0 20 20">
                <path d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-500">میانگین عملکرد</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.avgPerformance }}%</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow px-4 py-4">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-purple-100 rounded-full flex items-center justify-center">
              <svg class="w-5 h-5 text-purple-600" fill="currentColor" viewBox="0 0 20 20">
                <path d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-500">سناریوهای تست</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.testingScenarios }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Scenario Creation Modal -->
    <div v-if="showCreateModal" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
      <div class="relative top-20 mx-auto p-5 border w-11/12 md:w-3/4 lg:w-1/2 shadow-lg rounded-md bg-white">
        <div class="mt-3">
          <h3 class="text-lg font-medium text-gray-900 mb-4">ایجاد سناریو جدید</h3>
          
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نام سناریو</label>
              <input v-model="newScenario.name" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">توضیحات</label>
              <textarea v-model="newScenario.description" rows="3" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"></textarea>
            </div>
            
            <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">نوع سناریو</label>
                <select v-model="newScenario.type" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
                  <option value="basic">پایه</option>
                  <option value="advanced">پیشرفته</option>
                  <option value="custom">سفارشی</option>
                </select>
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">اولویت</label>
                <select v-model="newScenario.priority" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
                  <option value="low">کم</option>
                  <option value="medium">متوسط</option>
                  <option value="high">زیاد</option>
                  <option value="critical">بحرانی</option>
                </select>
              </div>
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">قوانین امتیازدهی</label>
              <div class="space-y-2">
                <div v-for="(rule, index) in newScenario.rules" :key="index" class="flex items-center space-x-2 space-x-reverse">
                  <input v-model="rule.name" type="text" placeholder="نام قانون" class="flex-1 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
                  <input v-model="rule.score" type="number" placeholder="امتیاز" class="w-24 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
                  <button class="text-red-600 hover:text-red-900" @click="removeRule(index)">حذف</button>
                </div>
                <button class="text-blue-600 hover:text-blue-900 text-sm" @click="addRule">+ افزودن قانون</button>
              </div>
            </div>
          </div>
          
          <div class="flex justify-end space-x-3 space-x-reverse mt-6">
            <button class="px-4 py-2 border border-gray-300 rounded-md text-gray-700 hover:bg-gray-50" @click="closeCreateModal">
              انصراف
            </button>
            <button class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700" @click="saveScenario">
              ذخیره سناریو
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue';

interface Scenario {
  id: number;
  name: string;
  description: string;
  type: string;
  status?: string;
  createdAt?: string;
  affectedUsers?: number;
  performance?: number;
  complexity?: string;
  rulesCount?: number;
  setupTime?: string;
  [key: string]: unknown;
}

// Props and Emits definition
const props = defineProps<{
  scenarios?: Scenario[]
}>() // Defined props

defineEmits<{
  saveScenario: [scenario: Record<string, unknown>]
  updateScenario: [scenario: Record<string, unknown>]
  deleteScenario: [scenario: Record<string, unknown>]
  exportScenarios: [data: Record<string, unknown>]
}>()

// Reactive data
const searchQuery = ref('')
const filterStatus = ref('')
const filterType = ref('')
const currentPage = ref(1)
const itemsPerPage = 10
const showCreateModal = ref(false)

// New scenario form
const newScenario = ref({
  name: '',
  description: '',
  type: 'basic',
  priority: 'medium',
  rules: []
})

// Sample scenario templates
const scenarioTemplates = ref<Scenario[]>([
  {
    id: 1,
    name: 'سناریو پایه فروشگاه',
    description: 'سناریو استاندارد برای فروشگاه‌های آنلاین با امتیازدهی بر اساس خرید و نظرات',
    type: 'basic',
    complexity: 'ساده',
    rulesCount: 8,
    setupTime: '15 دقیقه'
  },
  {
    id: 2,
    name: 'سناریو پیشرفته وفاداری',
    description: 'سناریو پیشرفته با سیستم‌های پیچیده امتیازدهی و پاداش‌های چندلایه',
    type: 'advanced',
    complexity: 'متوسط',
    rulesCount: 15,
    setupTime: '45 دقیقه'
  },
  {
    id: 3,
    name: 'سناریو ارجاع و مشارکت',
    description: 'تمرکز بر سیستم ارجاع و مشارکت کاربران در محتوا',
    type: 'custom',
    complexity: 'پیشرفته',
    rulesCount: 12,
    setupTime: '30 دقیقه'
  },
  {
    id: 4,
    name: 'سناریو کیفیت و بازگشت',
    description: 'امتیازدهی بر اساس کیفیت نظرات و نرخ بازگشت کالا',
    type: 'advanced',
    complexity: 'متوسط',
    rulesCount: 10,
    setupTime: '25 دقیقه'
  },
  {
    id: 5,
    name: 'سناریو قدمت و وفاداری',
    description: 'پاداش برای کاربران قدیمی و وفادار با سیستم سطوح',
    type: 'basic',
    complexity: 'ساده',
    rulesCount: 6,
    setupTime: '20 دقیقه'
  },
  {
    id: 6,
    name: 'سناریو تداوم خرید',
    description: 'امتیازدهی بر اساس الگوی خرید و تداوم فعالیت',
    type: 'custom',
    complexity: 'پیشرفته',
    rulesCount: 14,
    setupTime: '35 دقیقه'
  }
])

// Sample scenarios data
const localScenarios = ref<Scenario[]>([
  {
    id: 1,
    name: 'سناریو پایه فروشگاه',
    description: 'سناریو استاندارد برای فروشگاه‌های آنلاین',
    type: 'basic',
    status: 'active',
    createdAt: '2024-01-15T10:30:00Z',
    affectedUsers: 1250,
    performance: 85
  },
  {
    id: 2,
    name: 'سناریو پیشرفته وفاداری',
    description: 'سیستم پیشرفته امتیازدهی و پاداش‌ها',
    type: 'advanced',
    status: 'testing',
    createdAt: '2024-01-14T15:45:00Z',
    affectedUsers: 850,
    performance: 72
  },
  {
    id: 3,
    name: 'سناریو ارجاع و مشارکت',
    description: 'تمرکز بر ارجاع و مشارکت کاربران',
    type: 'custom',
    status: 'active',
    createdAt: '2024-01-13T09:20:00Z',
    affectedUsers: 650,
    performance: 91
  }
])

const displayScenarios = computed(() => props.scenarios || localScenarios.value)

// Statistics
const stats = ref({
  activeScenarios: 8,
  totalAffectedUsers: 2750,
  avgPerformance: 82.5,
  testingScenarios: 3
})

// Pagination
const pagination = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage + 1
  const end = Math.min(currentPage.value * itemsPerPage, filteredScenarios.value.length)
  return {
    start,
    end,
    total: filteredScenarios.value.length
  }
})

const totalPages = computed(() => Math.ceil(filteredScenarios.value.length / itemsPerPage))

const visiblePages = computed(() => {
  const pages = []
  const start = Math.max(1, currentPage.value - 2)
  const end = Math.min(totalPages.value, currentPage.value + 2)
  
  for (let i = start; i <= end; i++) {
    pages.push(i)
  }
  return pages
})

// Filtered scenarios
const filteredScenarios = computed(() => {
  let filtered = displayScenarios.value

  if (searchQuery.value) {
    filtered = filtered.filter(scenario => 
      scenario.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      scenario.description.toLowerCase().includes(searchQuery.value.toLowerCase())
    )
  }

  if (filterStatus.value) {
    filtered = filtered.filter(scenario => scenario.status === filterStatus.value)
  }

  if (filterType.value) {
    filtered = filtered.filter(scenario => scenario.type === filterType.value)
  }

  return filtered
})

const paginatedScenarios = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage
  const end = start + itemsPerPage
  return filteredScenarios.value.slice(start, end)
})

// Methods
const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleDateString('fa-IR')
}

const getTemplateTypeClass = (type: string) => {
  switch (type) {
    case 'basic':
      return 'bg-green-100 text-green-800'
    case 'advanced':
      return 'bg-blue-100 text-blue-800'
    case 'custom':
      return 'bg-purple-100 text-purple-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getTemplateTypeText = (type: string) => {
  switch (type) {
    case 'basic':
      return 'پایه'
    case 'advanced':
      return 'پیشرفته'
    case 'custom':
      return 'سفارشی'
    default:
      return 'نامشخص'
  }
}

const getTypeClass = (type: string) => {
  switch (type) {
    case 'basic':
      return 'bg-green-100 text-green-800'
    case 'advanced':
      return 'bg-blue-100 text-blue-800'
    case 'custom':
      return 'bg-purple-100 text-purple-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getTypeText = (type: string) => {
  switch (type) {
    case 'basic':
      return 'پایه'
    case 'advanced':
      return 'پیشرفته'
    case 'custom':
      return 'سفارشی'
    default:
      return 'نامشخص'
  }
}

const getStatusClass = (status: string) => {
  switch (status) {
    case 'active':
      return 'bg-green-100 text-green-800'
    case 'inactive':
      return 'bg-red-100 text-red-800'
    case 'draft':
      return 'bg-yellow-100 text-yellow-800'
    case 'testing':
      return 'bg-blue-100 text-blue-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getStatusText = (status: string) => {
  switch (status) {
    case 'active':
      return 'فعال'
    case 'inactive':
      return 'غیرفعال'
    case 'draft':
      return 'پیش‌نویس'
    case 'testing':
      return 'در حال تست'
    default:
      return 'نامشخص'
  }
}

const createNewScenario = () => {
  showCreateModal.value = true
  newScenario.value = {
    name: '',
    description: '',
    type: 'basic',
    priority: 'medium',
    rules: []
  }
}

const closeCreateModal = () => {
  showCreateModal.value = false
}

const addRule = () => {
  newScenario.value.rules.push({
    name: '',
    score: 0
  })
}

const removeRule = (index: number) => {
  newScenario.value.rules.splice(index, 1)
}

const saveScenario = () => {
  // API call to save scenario
  closeCreateModal()
}

const useTemplate = (_template: Record<string, unknown>) => {
  // Use template to create new scenario
}

const previewTemplate = (_template: Record<string, unknown>) => {
  // Preview template details
}

const exportScenarios = () => {
  // Export scenarios to Excel
}

const editScenario = (_scenario: Scenario) => {
  // Edit scenario
}

const toggleScenario = (scenario: Scenario) => {
  scenario.status = scenario.status === 'active' ? 'inactive' : 'active'
  // API call to toggle scenario status
}

const duplicateScenario = (_scenario: Scenario) => {
  // Duplicate scenario
}

const deleteScenario = (_scenario: Scenario) => {
  // Delete scenario
}

const previousPage = () => {
  if (currentPage.value > 1) {
    currentPage.value--
  }
}

const nextPage = () => {
  if (currentPage.value < totalPages.value) {
    currentPage.value++
  }
}

const goToPage = (page: number) => {
  currentPage.value = page
}

// Lifecycle
// onMounted removed as it only contained console.log
</script> 

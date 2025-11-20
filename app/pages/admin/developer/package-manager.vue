<template>
  <div class="min-h-screen bg-gray-50 p-6">
    <div class="max-w-7xl mx-auto">
      <!-- Header -->
      <div class="mb-8">
        <h1 class="text-3xl font-bold text-gray-900 mb-2">مدیریت پکیج‌ها</h1>
        <p class="text-gray-600">ابزارهای مدیریت وابستگی‌ها و پکیج‌ها</p>
      </div>

      <!-- Stats Cards -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-blue-100 rounded-lg">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"></path>
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">کل پکیج‌ها</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.totalPackages }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-green-100 rounded-lg">
              <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">به‌روز</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.upToDate }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-yellow-100 rounded-lg">
              <svg class="w-6 h-6 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">نیاز به به‌روزرسانی</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.outdated }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-red-100 rounded-lg">
              <svg class="w-6 h-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">آسیب‌پذیر</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.vulnerable }}</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Package Manager Selection -->
      <div class="mb-8 bg-white rounded-lg shadow">
        <div class="p-6 border-b border-gray-200">
          <h2 class="text-xl font-semibold text-gray-900">انتخاب مدیر پکیج</h2>
        </div>
        <div class="p-6">
          <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
            <button 
              :class="[
                'flex items-center justify-center p-6 rounded-lg border-2 transition-colors',
                selectedManager === 'npm' ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-gray-300'
              ]"
              @click="selectPackageManager('npm')"
            >
              <div class="text-center">
                <div class="w-12 h-12 bg-red-500 rounded-lg flex items-center justify-center mx-auto mb-2">
                  <span class="text-white font-bold text-lg">npm</span>
                </div>
                <span class="text-sm font-medium text-gray-900">npm</span>
              </div>
            </button>

            <button 
              :class="[
                'flex items-center justify-center p-6 rounded-lg border-2 transition-colors',
                selectedManager === 'yarn' ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-gray-300'
              ]"
              @click="selectPackageManager('yarn')"
            >
              <div class="text-center">
                <div class="w-12 h-12 bg-blue-500 rounded-lg flex items-center justify-center mx-auto mb-2">
                  <span class="text-white font-bold text-lg">Y</span>
                </div>
                <span class="text-sm font-medium text-gray-900">Yarn</span>
              </div>
            </button>

            <button 
              :class="[
                'flex items-center justify-center p-6 rounded-lg border-2 transition-colors',
                selectedManager === 'pnpm' ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-gray-300'
              ]"
              @click="selectPackageManager('pnpm')"
            >
              <div class="text-center">
                <div class="w-12 h-12 bg-orange-500 rounded-lg flex items-center justify-center mx-auto mb-2">
                  <span class="text-white font-bold text-lg">p</span>
                </div>
                <span class="text-sm font-medium text-gray-900">pnpm</span>
              </div>
            </button>

            <button 
              :class="[
                'flex items-center justify-center p-6 rounded-lg border-2 transition-colors',
                selectedManager === 'go' ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-gray-300'
              ]"
              @click="selectPackageManager('go')"
            >
              <div class="text-center">
                <div class="w-12 h-12 bg-blue-600 rounded-lg flex items-center justify-center mx-auto mb-2">
                  <span class="text-white font-bold text-lg">Go</span>
                </div>
                <span class="text-sm font-medium text-gray-900">Go Modules</span>
              </div>
            </button>
          </div>
        </div>
      </div>

      <!-- Main Content Grid -->
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <!-- Package Search & Install -->
        <div class="bg-white rounded-lg shadow">
          <div class="p-6 border-b border-gray-200">
            <h2 class="text-xl font-semibold text-gray-900">جستجو و نصب</h2>
          </div>
          <div class="p-6">
            <div class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">جستجوی پکیج</label>
                <div class="flex items-center space-x-2 space-x-reverse">
                  <input 
                    v-model="searchQuery"
                    type="text"
                    class="flex-1 border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
                    placeholder="نام پکیج را وارد کنید..."
                    @input="searchPackages"
                  >
                  <button 
                    class="bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                    @click="searchPackages"
                  >
                    جستجو
                  </button>
                </div>
              </div>

              <!-- Search Results -->
              <div v-if="searchResults.length > 0">
                <label class="block text-sm font-medium text-gray-700 mb-2">نتایج جستجو</label>
                <div class="space-y-2 max-h-64 overflow-y-auto">
                                     <div v-for="pkg in searchResults" :key="pkg.name" class="flex items-center justify-between p-3 border rounded-lg">
                     <div>
                       <p class="font-medium text-gray-900">{{ pkg.name }}</p>
                       <p class="text-sm text-gray-500">{{ pkg.description }}</p>
                       <p class="text-xs text-gray-400">نسخه: {{ pkg.version }}</p>
                     </div>
                     <button 
                       class="bg-green-600 hover:bg-green-700 text-white font-medium py-1 px-3 rounded text-sm transition-colors"
                       @click="installPackage(pkg)"
                     >
                       نصب
                     </button>
                   </div>
                </div>
              </div>

              <!-- Quick Install -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">نصب سریع</label>
                <div class="space-y-2">
                  <button 
                    class="w-full bg-purple-600 hover:bg-purple-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                    @click="quickInstall('vue')"
                  >
                    نصب Vue.js
                  </button>
                  <button 
                    class="w-full bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                    @click="quickInstall('axios')"
                  >
                    نصب Axios
                  </button>
                  <button 
                    class="w-full bg-cyan-600 hover:bg-cyan-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                    @click="quickInstall('tailwindcss')"
                  >
                    نصب Tailwind CSS
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Installed Packages -->
        <div class="lg:col-span-2 bg-white rounded-lg shadow">
          <div class="p-6 border-b border-gray-200">
            <div class="flex items-center justify-between">
              <h2 class="text-xl font-semibold text-gray-900">پکیج‌های نصب شده</h2>
              <div class="flex items-center space-x-2 space-x-reverse">
                <button 
                  class="bg-yellow-600 hover:bg-yellow-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="updateAllPackages"
                >
                  به‌روزرسانی همه
                </button>
                <button 
                  class="bg-red-600 hover:bg-red-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="auditPackages"
                >
                  بررسی امنیت
                </button>
              </div>
            </div>
          </div>
          <div class="p-6">
            <div class="space-y-4">
                             <div v-for="pkg in installedPackages" :key="pkg.name" class="flex items-center justify-between p-6 border rounded-lg">
                 <div class="flex-1">
                   <div class="flex items-center space-x-3 space-x-reverse">
                     <span class="font-medium text-gray-900">{{ pkg.name }}</span>
                     <span class="text-sm text-gray-500">@{{ pkg.version }}</span>
                     <span v-if="pkg.latest !== pkg.version" class="text-xs bg-yellow-100 text-yellow-800 px-2 py-1 rounded">
                       نسخه جدید: {{ pkg.latest }}
                     </span>
                     <span v-if="pkg.vulnerable" class="text-xs bg-red-100 text-red-800 px-2 py-1 rounded">
                       آسیب‌پذیر
                     </span>
                   </div>
                   <p class="text-sm text-gray-500 mt-1">{{ pkg.description }}</p>
                 </div>
                 <div class="flex items-center space-x-2 space-x-reverse">
                   <button 
                     v-if="pkg.latest !== pkg.version"
                     class="text-yellow-600 hover:text-yellow-800 text-sm font-medium"
                     @click="updatePackage(pkg)"
                   >
                     به‌روزرسانی
                   </button>
                   <button 
                     class="text-red-600 hover:text-red-800 text-sm font-medium"
                     @click="removePackage(pkg)"
                   >
                     حذف
                   </button>
                   <button 
                     class="text-blue-600 hover:text-blue-800 text-sm font-medium"
                     @click="viewPackageDetails(pkg)"
                   >
                     جزئیات
                   </button>
                 </div>
               </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Package Scripts -->
      <div class="mt-8 bg-white rounded-lg shadow">
        <div class="p-6 border-b border-gray-200">
          <h2 class="text-xl font-semibold text-gray-900">اسکریپت‌های پکیج</h2>
        </div>
        <div class="p-6">
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
            <div v-for="script in packageScripts" :key="script.name" class="border rounded-lg p-6">
              <div class="flex items-center justify-between mb-2">
                <h3 class="font-medium text-gray-900">{{ script.name }}</h3>
                <span class="text-xs bg-gray-100 text-gray-600 px-2 py-1 rounded">{{ script.type }}</span>
              </div>
              <p class="text-sm text-gray-500 mb-3">{{ script.description }}</p>
              <button 
                class="w-full bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                @click="runScript(script)"
              >
                اجرا
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- Package Details Modal -->
      <div v-if="packageModal.show" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
        <div class="bg-white rounded-lg max-w-2xl max-h-4xl overflow-auto">
          <div class="p-6 border-b border-gray-200">
            <div class="flex items-center justify-between">
              <h3 class="text-lg font-semibold text-gray-900">{{ packageModal.package?.name }}</h3>
              <button class="text-gray-400 hover:text-gray-600" @click="packageModal.show = false">
                <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                </svg>
              </button>
            </div>
          </div>
          <div class="p-6">
            <div class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">توضیحات</label>
                <p class="text-gray-900">{{ packageModal.package?.description }}</p>
              </div>
              <div class="grid grid-cols-2 gap-6">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">نسخه فعلی</label>
                  <p class="text-gray-900">{{ packageModal.package?.version }}</p>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">آخرین نسخه</label>
                  <p class="text-gray-900">{{ packageModal.package?.latest }}</p>
                </div>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">وابستگی‌ها</label>
                <div class="space-y-1">
                  <span v-for="dep in packageModal.package?.dependencies" :key="dep" class="inline-block bg-gray-100 text-gray-700 px-2 py-1 rounded text-sm mr-2 mb-2">
                    {{ dep }}
                  </span>
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
import { ref, reactive } from 'vue'

definePageMeta({
  layout: 'admin-main',
  middleware: ['developer-only']
});

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user, hasPermission } = useAuth()

// Reactive data
const selectedManager = ref('npm')
const searchQuery = ref('')
const searchResults = ref([])
const packageModal = reactive({
  show: false,
  package: null
})

const stats = reactive({
  totalPackages: 45,
  upToDate: 32,
  outdated: 12,
  vulnerable: 1
})

const installedPackages = ref([
  {
    name: 'vue',
    version: '3.3.4',
    latest: '3.4.0',
    description: 'Progressive JavaScript framework',
    vulnerable: false
  },
  {
    name: 'axios',
    version: '1.6.0',
    latest: '1.6.0',
    description: 'Promise based HTTP client',
    vulnerable: false
  },
  {
    name: 'tailwindcss',
    version: '3.3.0',
    latest: '3.4.0',
    description: 'Utility-first CSS framework',
    vulnerable: false
  },
  {
    name: 'lodash',
    version: '4.17.20',
    latest: '4.17.21',
    description: 'JavaScript utility library',
    vulnerable: true
  }
])

const packageScripts = ref([
  {
    name: 'dev',
    type: 'development',
    description: 'اجرای سرور توسعه'
  },
  {
    name: 'build',
    type: 'production',
    description: 'ساخت پروژه برای تولید'
  },
  {
    name: 'test',
    type: 'testing',
    description: 'اجرای تست‌ها'
  },
  {
    name: 'lint',
    type: 'quality',
    description: 'بررسی کیفیت کد'
  }
])

// Methods
function selectPackageManager(manager) {
  selectedManager.value = manager
  alert(`مدیر پکیج به ${manager} تغییر یافت`)
}

function searchPackages() {
  if (!searchQuery.value.trim()) return
  
  // Simulate search results
  searchResults.value = [
    {
      name: searchQuery.value,
      description: 'A sample package description',
      version: '1.0.0'
    },
    {
      name: `${searchQuery.value}-utils`,
      description: 'Utility functions for ' + searchQuery.value,
      version: '2.1.0'
    }
  ]
}

function installPackage(pkg) {
  alert(`پکیج ${pkg.name} نصب شد`)
  searchResults.value = []
  searchQuery.value = ''
}

function quickInstall(packageName) {
  alert(`پکیج ${packageName} نصب شد`)
}

function updateAllPackages() {
  alert('همه پکیج‌ها به‌روزرسانی شدند')
}

function auditPackages() {
  alert('بررسی امنیت پکیج‌ها انجام شد')
}

function updatePackage(pkg) {
  alert(`پکیج ${pkg.name} به‌روزرسانی شد`)
}

function removePackage(pkg) {
  if (confirm(`آیا می‌خواهید پکیج ${pkg.name} را حذف کنید؟`)) {
    alert(`پکیج ${pkg.name} حذف شد`)
  }
}

function viewPackageDetails(pkg) {
  packageModal.package = {
    ...pkg,
    dependencies: ['vue', 'axios', 'lodash']
  }
  packageModal.show = true
}

function runScript(script) {
  alert(`اسکریپت ${script.name} اجرا شد`)
}
</script> 

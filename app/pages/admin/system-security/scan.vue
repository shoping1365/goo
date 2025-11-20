<template>
  <div class="p-6" dir="rtl">
    <div class="mb-6 bg-white rounded-lg shadow-md border border-gray-200 p-6">
      <h1 class="text-2xl font-bold text-gray-900 mb-2">اسکن</h1>
      <p class="text-gray-600">اسکن امنیتی سیستم و شناسایی آسیب‌پذیری‌ها</p>
    </div>

    <!-- آمار کلی -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
      <TemplateCard
        title="آخرین اسکن"
        :value="lastScanDate"
        variant="green"
      />
      <TemplateCard
        title="آسیب‌پذیری‌های بحرانی"
        :value="criticalVulnerabilities"
        variant="red"
      />
      <TemplateCard
        title="آسیب‌پذیری‌های متوسط"
        :value="mediumVulnerabilities"
        variant="yellow"
      />
      <TemplateCard
        title="آسیب‌پذیری‌های کم"
        :value="lowVulnerabilities"
        variant="blue"
      />
    </div>

    <!-- کنترل‌های اسکن -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
      <!-- کارت اسکن سریع -->
      <div class="bg-white rounded-2xl shadow-lg border border-gray-200 p-6">
        <div class="mb-4">
          <h3 class="text-lg font-bold text-gray-900">اسکن سریع</h3>
          <p class="text-sm text-gray-600">اسکن فوری فایل‌ها و پیکربندی‌های اصلی</p>
        </div>
        <button :disabled="isScanning" class="w-full px-4 py-3 bg-gradient-to-r from-cyan-400 to-blue-600 text-white rounded-lg hover:from-cyan-500 hover:to-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-cyan-500 shadow-md transition-all duration-200 hover:shadow-lg disabled:opacity-50 font-medium" @click="startQuickScan">
          {{ isScanning ? 'در حال اسکن...' : 'شروع اسکن سریع' }}
        </button>
      </div>
      
      <!-- کارت اسکن کامل -->
      <div class="bg-white rounded-2xl shadow-lg border border-gray-200 p-6">
        <div class="mb-4">
          <h3 class="text-lg font-bold text-gray-900">اسکن کامل</h3>
          <p class="text-sm text-gray-600">اسکن کامل تمام فایل‌ها و دیتابیس</p>
        </div>
        <button :disabled="isScanning" class="w-full px-4 py-3 bg-gradient-to-r from-teal-400 to-emerald-600 text-white rounded-lg hover:from-teal-500 hover:to-emerald-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-teal-500 shadow-md transition-all duration-200 hover:shadow-lg disabled:opacity-50 font-medium" @click="startFullScan">
          {{ isScanning ? 'در حال اسکن...' : 'شروع اسکن کامل' }}
        </button>
      </div>
      
      <!-- کارت اسکن سفارشی -->
      <div class="bg-white rounded-2xl shadow-lg border border-gray-200 p-6">
        <div class="mb-4">
          <h3 class="text-lg font-bold text-gray-900">اسکن سفارشی</h3>
          <p class="text-sm text-gray-600">انتخاب مسیرها و نوع اسکن</p>
        </div>
        <button :disabled="isScanning" class="w-full px-4 py-3 bg-gradient-to-r from-indigo-400 to-purple-600 text-white rounded-lg hover:from-indigo-500 hover:to-purple-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 shadow-md transition-all duration-200 hover:shadow-lg disabled:opacity-50 font-medium" @click="showCustomScanModal = true">
          اسکن سفارشی
        </button>
      </div>
    </div>

    <!-- پیشرفت اسکن -->
    <div v-if="isScanning" class="bg-white p-6 rounded-lg shadow-md border border-gray-200 mb-8">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">پیشرفت اسکن</h3>
      <div class="mb-4">
        <div class="flex justify-between text-sm text-gray-600 mb-2">
          <span>{{ scanProgress.currentFile }}</span>
          <span>{{ scanProgress.percentage }}%</span>
        </div>
        <div class="w-full bg-gray-200 rounded-full h-2">
          <div class="bg-blue-600 h-2 rounded-full transition-all duration-300" :style="{ width: scanProgress.percentage + '%' }"></div>
        </div>
      </div>
      <div class="text-sm text-gray-600">
        <p>فایل‌های اسکن شده: {{ scanProgress.scannedFiles }} از {{ scanProgress.totalFiles }}</p>
        <p>آسیب‌پذیری‌های یافت شده: {{ scanProgress.foundVulnerabilities }}</p>
      </div>
    </div>

    <!-- نتایج اسکن -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 mb-8">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-semibold text-gray-900">نتایج آخرین اسکن</h3>
      </div>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">سطح خطر</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع آسیب‌پذیری</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مسیر فایل</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">توضیحات</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="vulnerability in vulnerabilities" :key="vulnerability.id" :class="getVulnerabilityRowClass(vulnerability.severity)">
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getSeverityClass(vulnerability.severity)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ getSeverityText(vulnerability.severity) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ vulnerability.type }}</td>
              <td class="px-6 py-4 text-sm text-gray-900">{{ vulnerability.filePath }}</td>
              <td class="px-6 py-4 text-sm text-gray-900">{{ vulnerability.description }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-green-600 hover:text-green-900 ml-2" @click="fixVulnerability(vulnerability.id)">رفع خودکار</button>
                <button class="text-blue-600 hover:text-blue-900" @click="viewDetails(vulnerability.id)">جزئیات</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- تاریخچه اسکن‌ها -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-semibold text-gray-900">تاریخچه اسکن‌ها</h3>
      </div>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع اسکن</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مدت زمان</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">آسیب‌پذیری‌ها</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="scan in scanHistory" :key="scan.id">
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ scan.date }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ scan.type }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ scan.duration }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ scan.vulnerabilities }}</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getStatusClass(scan.status)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ getStatusText(scan.status) }}
                </span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- مودال اسکن سفارشی -->
    <div v-if="showCustomScanModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg p-6 w-full max-w-md">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">اسکن سفارشی</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">مسیرهای اسکن</label>
            <textarea v-model="customScan.paths" rows="3" placeholder="مسیرهای مورد نظر را وارد کنید (هر مسیر در یک خط)" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"></textarea>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع اسکن</label>
            <select v-model="customScan.type" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
              <option value="files">فایل‌ها</option>
              <option value="database">دیتابیس</option>
              <option value="network">شبکه</option>
              <option value="all">همه</option>
            </select>
          </div>
        </div>
        <div class="mt-6 flex justify-end space-x-3">
          <button class="px-4 py-2 bg-gray-200 text-gray-800 rounded-md hover:bg-gray-300" @click="showCustomScanModal = false">انصراف</button>
          <button class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700" @click="startCustomScan">شروع اسکن</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import TemplateCard from '@/components/common/TemplateCard.vue'
definePageMeta({ layout: 'admin-main' })

// استفاده از useAuth برای چک کردن پرمیژن‌ها
// Auth disabled
// داده‌های نمونه
const lastScanDate = ref('1402/10/15 10:30')
const criticalVulnerabilities = ref(2)
const mediumVulnerabilities = ref(8)
const lowVulnerabilities = ref(15)

const isScanning = ref(false)
const showCustomScanModal = ref(false)

const scanProgress = ref({
  percentage: 0,
  currentFile: '',
  scannedFiles: 0,
  totalFiles: 0,
  foundVulnerabilities: 0
})

const customScan = ref({
  paths: '',
  type: 'files'
})

const vulnerabilities = ref([
  {
    id: 1,
    severity: 'critical',
    type: 'SQL Injection',
    filePath: '/admin/users.php',
    description: 'امکان تزریق کد SQL در پارامتر user_id'
  },
  {
    id: 2,
    severity: 'critical',
    type: 'XSS',
    filePath: '/admin/comments.php',
    description: 'امکان تزریق کد JavaScript در فیلد نظر'
  },
  {
    id: 3,
    severity: 'medium',
    type: 'File Upload',
    filePath: '/admin/upload.php',
    description: 'عدم بررسی نوع فایل آپلود شده'
  },
  {
    id: 4,
    severity: 'low',
    type: 'Information Disclosure',
    filePath: '/config/database.php',
    description: 'نمایش اطلاعات حساس در خطاها'
  }
])

const scanHistory = ref([
  {
    id: 1,
    date: '1402/10/15 10:30',
    type: 'اسکن کامل',
    duration: '45 دقیقه',
    vulnerabilities: 25,
    status: 'completed'
  },
  {
    id: 2,
    date: '1402/10/14 15:20',
    type: 'اسکن سریع',
    duration: '5 دقیقه',
    vulnerabilities: 3,
    status: 'completed'
  },
  {
    id: 3,
    date: '1402/10/13 09:15',
    type: 'اسکن سفارشی',
    duration: '20 دقیقه',
    vulnerabilities: 12,
    status: 'failed'
  }
])

function getSeverityClass(severity) {
  switch (severity) {
    case 'critical':
      return 'bg-red-100 text-red-800'
    case 'medium':
      return 'bg-yellow-100 text-yellow-800'
    case 'low':
      return 'bg-blue-100 text-blue-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

function getSeverityText(severity) {
  switch (severity) {
    case 'critical':
      return 'بحرانی'
    case 'medium':
      return 'متوسط'
    case 'low':
      return 'کم'
    default:
      return 'نامشخص'
  }
}

function getVulnerabilityRowClass(severity) {
  if (severity === 'critical') {
    return 'bg-red-50'
  }
  return ''
}

function getStatusClass(status) {
  switch (status) {
    case 'completed':
      return 'bg-green-100 text-green-800'
    case 'failed':
      return 'bg-red-100 text-red-800'
    case 'running':
      return 'bg-blue-100 text-blue-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

function getStatusText(status) {
  switch (status) {
    case 'completed':
      return 'تکمیل شده'
    case 'failed':
      return 'ناموفق'
    case 'running':
      return 'در حال اجرا'
    default:
      return 'نامشخص'
  }
}

function startQuickScan() {
  isScanning.value = true
  scanProgress.value = {
    percentage: 0,
    currentFile: 'شروع اسکن...',
    scannedFiles: 0,
    totalFiles: 100,
    foundVulnerabilities: 0
  }
  
  // شبیه‌سازی اسکن
  const interval = setInterval(() => {
    scanProgress.value.percentage += 5
    scanProgress.value.scannedFiles += 5
    
    if (scanProgress.value.percentage >= 100) {
      clearInterval(interval)
      isScanning.value = false
      alert('اسکن سریع تکمیل شد')
    }
  }, 500)
}

function startFullScan() {
  isScanning.value = true
  scanProgress.value = {
    percentage: 0,
    currentFile: 'شروع اسکن کامل...',
    scannedFiles: 0,
    totalFiles: 1000,
    foundVulnerabilities: 0
  }
  
  // شبیه‌سازی اسکن کامل
  const interval = setInterval(() => {
    scanProgress.value.percentage += 2
    scanProgress.value.scannedFiles += 20
    
    if (scanProgress.value.percentage >= 100) {
      clearInterval(interval)
      isScanning.value = false
      alert('اسکن کامل تکمیل شد')
    }
  }, 1000)
}

function startCustomScan() {
  showCustomScanModal.value = false
  isScanning.value = true
  // منطق اسکن سفارشی
  alert('اسکن سفارشی شروع شد')
  setTimeout(() => {
    isScanning.value = false
  }, 3000)
}

function fixVulnerability(id) {
  // منطق رفع خودکار آسیب‌پذیری
  alert(`آسیب‌پذیری ${id} رفع شد`)
}

function viewDetails(id) {
  // منطق نمایش جزئیات
  alert(`جزئیات آسیب‌پذیری ${id} نمایش داده می‌شود`)
}
</script> 

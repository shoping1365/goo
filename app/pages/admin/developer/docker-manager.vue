<template>
  <div class="min-h-screen bg-gray-50 p-6">
    <div class="max-w-7xl mx-auto">
      <!-- Header -->
      <div class="mb-8">
        <h1 class="text-3xl font-bold text-gray-900 mb-2">مدیریت Docker</h1>
        <p class="text-gray-600">ابزارهای مدیریت کانتینرها و تصاویر</p>
      </div>

      <!-- Stats Cards -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-blue-100 rounded-lg">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">کانتینرهای فعال</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.runningContainers }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-green-100 rounded-lg">
              <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">تصاویر</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.images }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-purple-100 rounded-lg">
              <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">استفاده از CPU</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.cpuUsage }}%</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-yellow-100 rounded-lg">
              <svg class="w-6 h-6 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">استفاده از حافظه</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.memoryUsage }}%</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Main Content Grid -->
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <!-- Container Management -->
        <div class="bg-white rounded-lg shadow">
          <div class="p-6 border-b border-gray-200">
            <div class="flex items-center justify-between">
              <h2 class="text-xl font-semibold text-gray-900">مدیریت کانتینرها</h2>
              <button 
                @click="refreshContainers"
                class="bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
              >
                بروزرسانی
              </button>
            </div>
          </div>
          <div class="p-6">
            <div class="space-y-4">
              <div v-for="container in containers" :key="container.id" class="border rounded-lg p-6">
                <div class="flex items-center justify-between mb-2">
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <span class="font-medium text-gray-900">{{ container.name }}</span>
                    <span :class="[
                      'px-2 py-1 rounded text-xs font-medium',
                      container.status === 'running' ? 'bg-green-100 text-green-800' :
                      container.status === 'stopped' ? 'bg-red-100 text-red-800' :
                      'bg-yellow-100 text-yellow-800'
                    ]">
                      {{ container.status }}
                    </span>
                  </div>
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <button 
                      @click="startContainer(container)"
                      v-if="container.status === 'stopped'"
                      class="text-green-600 hover:text-green-800 text-sm font-medium"
                    >
                      شروع
                    </button>
                    <button 
                      @click="stopContainer(container)"
                      v-if="container.status === 'running'"
                      class="text-red-600 hover:text-red-800 text-sm font-medium"
                    >
                      توقف
                    </button>
                    <button 
                      @click="restartContainer(container)"
                      v-if="container.status === 'running'"
                      class="text-yellow-600 hover:text-yellow-800 text-sm font-medium"
                    >
                      راه‌اندازی مجدد
                    </button>
                    <button 
                      @click="removeContainer(container)"
                      class="text-gray-600 hover:text-gray-800 text-sm font-medium"
                    >
                      حذف
                    </button>
                  </div>
                </div>
                <div class="text-sm text-gray-500 space-y-1">
                  <p>تصویر: {{ container.image }}</p>
                  <p>پورت: {{ container.ports }}</p>
                  <p>حجم: {{ container.size }}</p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Image Management -->
        <div class="bg-white rounded-lg shadow">
          <div class="p-6 border-b border-gray-200">
            <h2 class="text-xl font-semibold text-gray-900">مدیریت تصاویر</h2>
          </div>
          <div class="p-6">
            <div class="space-y-4">
              <div v-for="image in images" :key="image.id" class="border rounded-lg p-6">
                <div class="flex items-center justify-between mb-2">
                  <div>
                    <p class="font-medium text-gray-900">{{ image.name }}</p>
                    <p class="text-sm text-gray-500">{{ image.tag }}</p>
                  </div>
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <button 
                      @click="runImage(image)"
                      class="text-blue-600 hover:text-blue-800 text-sm font-medium"
                    >
                      اجرا
                    </button>
                    <button 
                      @click="removeImage(image)"
                      class="text-red-600 hover:text-red-800 text-sm font-medium"
                    >
                      حذف
                    </button>
                  </div>
                </div>
                <div class="text-sm text-gray-500">
                  <p>حجم: {{ image.size }}</p>
                  <p>تاریخ ایجاد: {{ image.created }}</p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Docker Compose -->
        <div class="bg-white rounded-lg shadow">
          <div class="p-6 border-b border-gray-200">
            <h2 class="text-xl font-semibold text-gray-900">Docker Compose</h2>
          </div>
          <div class="p-6">
            <div class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">فایل docker-compose.yml</label>
                <textarea
                  v-model="composeFile"
                  rows="8"
                  class="w-full border border-gray-300 rounded-lg px-3 py-2 font-mono text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
                  placeholder="version: '3.8'&#10;services:&#10;  web:&#10;    image: nginx:latest&#10;    ports:&#10;      - '80:80'"
                ></textarea>
              </div>
              <div class="grid grid-cols-2 gap-2">
                <button 
                  @click="composeUp"
                  class="bg-green-600 hover:bg-green-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                >
                  اجرا (up)
                </button>
                <button 
                  @click="composeDown"
                  class="bg-red-600 hover:bg-red-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                >
                  توقف (down)
                </button>
                <button 
                  @click="composeBuild"
                  class="bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                >
                  ساخت (build)
                </button>
                <button 
                  @click="composeLogs"
                  class="bg-purple-600 hover:bg-purple-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                >
                  لاگ‌ها (logs)
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Container Logs -->
      <div class="mt-8 bg-white rounded-lg shadow">
        <div class="p-6 border-b border-gray-200">
          <div class="flex items-center justify-between">
            <h2 class="text-xl font-semibold text-gray-900">لاگ‌های کانتینر</h2>
            <div class="flex items-center space-x-2 space-x-reverse">
              <select 
                v-model="selectedContainer"
                class="border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
              >
                <option value="">انتخاب کانتینر</option>
                <option v-for="container in containers" :key="container.id" :value="container.id">
                  {{ container.name }}
                </option>
              </select>
              <button 
                @click="getLogs"
                :disabled="!selectedContainer"
                class="bg-blue-600 hover:bg-blue-700 disabled:bg-gray-400 text-white font-medium py-2 px-4 rounded-lg transition-colors"
              >
                دریافت لاگ‌ها
              </button>
            </div>
          </div>
        </div>
        <div class="p-6">
          <div class="bg-gray-900 rounded-lg p-6 h-64 overflow-y-auto">
            <pre class="text-green-400 font-mono text-sm">{{ containerLogs }}</pre>
          </div>
        </div>
      </div>

      <!-- Docker Stats -->
      <div class="mt-8 bg-white rounded-lg shadow">
        <div class="p-6 border-b border-gray-200">
          <h2 class="text-xl font-semibold text-gray-900">آمار سیستم</h2>
        </div>
        <div class="p-6">
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
            <div class="space-y-2">
              <h3 class="font-medium text-gray-900">استفاده از CPU</h3>
              <div class="w-full bg-gray-200 rounded-full h-2">
                <div class="bg-blue-600 h-2 rounded-full" :style="{ width: stats.cpuUsage + '%' }"></div>
              </div>
              <p class="text-sm text-gray-500">{{ stats.cpuUsage }}%</p>
            </div>
            <div class="space-y-2">
              <h3 class="font-medium text-gray-900">استفاده از حافظه</h3>
              <div class="w-full bg-gray-200 rounded-full h-2">
                <div class="bg-green-600 h-2 rounded-full" :style="{ width: stats.memoryUsage + '%' }"></div>
              </div>
              <p class="text-sm text-gray-500">{{ stats.memoryUsage }}%</p>
            </div>
            <div class="space-y-2">
              <h3 class="font-medium text-gray-900">استفاده از دیسک</h3>
              <div class="w-full bg-gray-200 rounded-full h-2">
                <div class="bg-yellow-600 h-2 rounded-full" :style="{ width: stats.diskUsage + '%' }"></div>
              </div>
              <p class="text-sm text-gray-500">{{ stats.diskUsage }}%</p>
            </div>
            <div class="space-y-2">
              <h3 class="font-medium text-gray-900">شبکه</h3>
              <div class="w-full bg-gray-200 rounded-full h-2">
                <div class="bg-purple-600 h-2 rounded-full" :style="{ width: stats.networkUsage + '%' }"></div>
              </div>
              <p class="text-sm text-gray-500">{{ stats.networkUsage }}%</p>
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
const selectedContainer = ref('')
const containerLogs = ref('')
const composeFile = ref(`version: '3.8'
services:
  web:
    image: nginx:latest
    ports:
      - '80:80'
  db:
    image: postgres:13
    environment:
      POSTGRES_DB: myapp
      POSTGRES_USER: user
      POSTGRES_PASSWORD: password`)

const stats = reactive({
  runningContainers: 3,
  images: 12,
  cpuUsage: 45,
  memoryUsage: 62,
  diskUsage: 28,
  networkUsage: 15
})

const containers = ref([
  {
    id: '1',
    name: 'web-server',
    status: 'running',
    image: 'nginx:latest',
    ports: '80:80',
    size: '2.1 MB'
  },
  {
    id: '2',
    name: 'database',
    status: 'running',
    image: 'postgres:13',
    ports: '5432:5432',
    size: '45.2 MB'
  },
  {
    id: '3',
    name: 'redis-cache',
    status: 'stopped',
    image: 'redis:alpine',
    ports: '6379:6379',
    size: '12.8 MB'
  }
])

const images = ref([
  {
    id: '1',
    name: 'nginx',
    tag: 'latest',
    size: '133 MB',
    created: '2 days ago'
  },
  {
    id: '2',
    name: 'postgres',
    tag: '13',
    size: '314 MB',
    created: '1 week ago'
  },
  {
    id: '3',
    name: 'redis',
    tag: 'alpine',
    size: '32.3 MB',
    created: '3 days ago'
  }
])

// Methods
function refreshContainers() {
  alert('کانتینرها بروزرسانی شدند')
}

function startContainer(container) {
  alert(`کانتینر ${container.name} شروع شد`)
  container.status = 'running'
}

function stopContainer(container) {
  alert(`کانتینر ${container.name} متوقف شد`)
  container.status = 'stopped'
}

function restartContainer(container) {
  alert(`کانتینر ${container.name} راه‌اندازی مجدد شد`)
}

function removeContainer(container) {
  if (confirm(`آیا می‌خواهید کانتینر ${container.name} را حذف کنید؟`)) {
    alert(`کانتینر ${container.name} حذف شد`)
  }
}

function runImage(image) {
  alert(`تصویر ${image.name}:${image.tag} اجرا شد`)
}

function removeImage(image) {
  if (confirm(`آیا می‌خواهید تصویر ${image.name}:${image.tag} را حذف کنید؟`)) {
    alert(`تصویر ${image.name}:${image.tag} حذف شد`)
  }
}

function composeUp() {
  alert('Docker Compose up اجرا شد')
}

function composeDown() {
  alert('Docker Compose down اجرا شد')
}

function composeBuild() {
  alert('Docker Compose build اجرا شد')
}

function composeLogs() {
  alert('لاگ‌های Docker Compose نمایش داده می‌شوند')
}

function getLogs() {
  if (selectedContainer.value) {
    containerLogs.value = `[2024-01-15 10:30:15] INFO: Container started
[2024-01-15 10:30:16] INFO: Application initialized
[2024-01-15 10:30:17] INFO: Database connection established
[2024-01-15 10:30:18] INFO: Server listening on port 80
[2024-01-15 10:30:19] INFO: Health check passed`
  }
}
</script> 

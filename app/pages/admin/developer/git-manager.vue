<template>
  <div class="min-h-screen bg-gray-50 p-6">
    <div class="max-w-7xl mx-auto">
      <!-- Header -->
      <div class="mb-8">
        <h1 class="text-3xl font-bold text-gray-900 mb-2">مدیریت مخزن گیت</h1>
        <p class="text-gray-600">ابزارهای مدیریت و کنترل نسخه</p>
      </div>

      <!-- Stats Cards -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-blue-100 rounded-lg">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 9l3 3-3 3m5 0h3M5 20h14a2 2 0 002-2V6a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">تعداد کامیت‌ها</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.commits }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-green-100 rounded-lg">
              <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"></path>
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">برچسب‌ها</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.tags }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-purple-100 rounded-lg">
              <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"></path>
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">شاخه‌ها</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.branches }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-yellow-100 rounded-lg">
              <svg class="w-6 h-6 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">آخرین کامیت</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.lastCommit }}</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Main Content Grid -->
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <!-- Repository Info -->
        <div class="bg-white rounded-lg shadow">
          <div class="p-6 border-b border-gray-200">
            <h2 class="text-xl font-semibold text-gray-900">اطلاعات مخزن</h2>
          </div>
          <div class="p-6">
            <div class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">آدرس مخزن</label>
                <div class="flex items-center space-x-2 space-x-reverse">
                  <input 
                    v-model="repositoryUrl"
                    type="text"
                    class="flex-1 border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
                    placeholder="https://github.com/user/repo.git"
                  >
                  <button 
                    class="bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                    @click="cloneRepository"
                  >
                    کلون
                  </button>
                </div>
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">شاخه فعلی</label>
                <select 
                  v-model="currentBranch"
                  class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
                  @change="switchBranch"
                >
                  <option v-for="branch in branches" :key="branch.name" :value="branch.name">
                    {{ branch.name }} {{ branch.current ? '(فعلی)' : '' }}
                  </option>
                </select>
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
                <div class="space-y-2">
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">فایل‌های تغییر یافته:</span>
                    <span class="text-sm font-medium text-red-600">{{ status.modified }}</span>
                  </div>
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">فایل‌های جدید:</span>
                    <span class="text-sm font-medium text-green-600">{{ status.added }}</span>
                  </div>
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">فایل‌های حذف شده:</span>
                    <span class="text-sm font-medium text-red-600">{{ status.deleted }}</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Commit History -->
        <div class="lg:col-span-2 bg-white rounded-lg shadow">
          <div class="p-6 border-b border-gray-200">
            <div class="flex items-center justify-between">
              <h2 class="text-xl font-semibold text-gray-900">تاریخچه کامیت‌ها</h2>
              <button 
                class="bg-green-600 hover:bg-green-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                @click="refreshCommits"
              >
                بروزرسانی
              </button>
            </div>
          </div>
          <div class="p-6">
            <div class="space-y-4">
              <div v-for="commit in commits" :key="commit.hash" class="border rounded-lg p-6 hover:bg-gray-50">
                <div class="flex items-start justify-between">
                  <div class="flex-1">
                    <div class="flex items-center space-x-2 space-x-reverse mb-2">
                      <span class="text-sm font-mono text-gray-500">{{ commit.hash.substring(0, 8) }}</span>
                      <span class="text-sm font-medium text-gray-900">{{ commit.message }}</span>
                    </div>
                    <div class="flex items-center space-x-4 space-x-reverse text-xs text-gray-500">
                      <span>{{ commit.author }}</span>
                      <span>{{ formatDate(commit.date) }}</span>
                      <span>{{ commit.files }} فایل</span>
                    </div>
                  </div>
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <button 
                      class="text-blue-600 hover:text-blue-800 text-sm font-medium"
                      @click="viewCommit(commit)"
                    >
                      مشاهده
                    </button>
                    <button 
                      class="text-red-600 hover:text-red-800 text-sm font-medium"
                      @click="revertCommit(commit)"
                    >
                      بازگشت
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Git Operations -->
      <div class="mt-8 bg-white rounded-lg shadow">
        <div class="p-6 border-b border-gray-200">
          <h2 class="text-xl font-semibold text-gray-900">عملیات گیت</h2>
        </div>
        <div class="p-6">
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
            <!-- Staging -->
            <div class="space-y-4">
              <h3 class="font-medium text-gray-900">مرحله‌بندی</h3>
              <div class="space-y-2">
                <button 
                  class="w-full bg-green-600 hover:bg-green-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="stageAll"
                >
                  اضافه کردن همه
                </button>
                <button 
                  class="w-full bg-gray-600 hover:bg-gray-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="unstageAll"
                >
                  حذف از مرحله
                </button>
              </div>
            </div>

            <!-- Committing -->
            <div class="space-y-4">
              <h3 class="font-medium text-gray-900">کامیت</h3>
              <div class="space-y-2">
                <input
                  v-model="commitMessage"
                  type="text"
                  placeholder="پیام کامیت..."
                  class="w-full border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                <button 
                  :disabled="!commitMessage"
                  class="w-full bg-blue-600 hover:bg-blue-700 disabled:bg-gray-400 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="commitChanges"
                >
                  کامیت
                </button>
              </div>
            </div>

            <!-- Pushing -->
            <div class="space-y-4">
              <h3 class="font-medium text-gray-900">ارسال</h3>
              <div class="space-y-2">
                <button 
                  class="w-full bg-purple-600 hover:bg-purple-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="pushChanges"
                >
                  ارسال تغییرات
                </button>
                <button 
                  class="w-full bg-indigo-600 hover:bg-indigo-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="pullChanges"
                >
                  دریافت تغییرات
                </button>
              </div>
            </div>

            <!-- Branching -->
            <div class="space-y-4">
              <h3 class="font-medium text-gray-900">شاخه‌ها</h3>
              <div class="space-y-2">
                <input
                  v-model="newBranchName"
                  type="text"
                  placeholder="نام شاخه جدید..."
                  class="w-full border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                <button 
                  :disabled="!newBranchName"
                  class="w-full bg-yellow-600 hover:bg-yellow-700 disabled:bg-gray-400 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="createBranch"
                >
                  ایجاد شاخه
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- File Changes -->
      <div class="mt-8 bg-white rounded-lg shadow">
        <div class="p-6 border-b border-gray-200">
          <h2 class="text-xl font-semibold text-gray-900">تغییرات فایل‌ها</h2>
        </div>
        <div class="p-6">
          <div class="space-y-4">
            <div v-for="file in changedFiles" :key="file.path" class="flex items-center justify-between p-6 border rounded-lg">
              <div class="flex items-center space-x-3 space-x-reverse">
                <span
:class="[
                  'px-2 py-1 rounded text-xs font-medium',
                  file.status === 'modified' ? 'bg-yellow-100 text-yellow-800' :
                  file.status === 'added' ? 'bg-green-100 text-green-800' :
                  'bg-red-100 text-red-800'
                ]">
                  {{ file.status }}
                </span>
                <span class="text-sm font-medium text-gray-900">{{ file.path }}</span>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <button 
                  v-if="!file.staged"
                  class="text-green-600 hover:text-green-800 text-sm font-medium"
                  @click="stageFile(file)"
                >
                  اضافه کردن
                </button>
                <button 
                  v-if="file.staged"
                  class="text-gray-600 hover:text-gray-800 text-sm font-medium"
                  @click="unstageFile(file)"
                >
                  حذف
                </button>
                <button 
                  class="text-blue-600 hover:text-blue-800 text-sm font-medium"
                  @click="viewDiff(file)"
                >
                  مشاهده تغییرات
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Diff Viewer Modal -->
      <div v-if="diffModal.show" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
        <div class="bg-white rounded-lg max-w-6xl max-h-4xl overflow-auto">
          <div class="p-6 border-b border-gray-200">
            <div class="flex items-center justify-between">
              <h3 class="text-lg font-semibold text-gray-900">تغییرات {{ diffModal.file?.path }}</h3>
              <button class="text-gray-400 hover:text-gray-600" @click="diffModal.show = false">
                <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                </svg>
              </button>
            </div>
          </div>
          <div class="p-6">
            <pre class="bg-gray-900 text-gray-100 p-6 rounded-lg overflow-x-auto text-sm">{{ diffModal.content }}</pre>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { reactive, ref } from 'vue';

definePageMeta({
  layout: 'admin-main',
  middleware: ['developer-only']
});

// استفاده از useAuth برای چک کردن پرمیژن‌ها
// const { user, hasPermission } = useAuth()

// Reactive data
const repositoryUrl = ref('https://github.com/user/my-project.git')
const currentBranch = ref('main')
const commitMessage = ref('')
const newBranchName = ref('')

const stats = reactive({
  commits: 156,
  tags: 8,
  branches: 12,
  lastCommit: '2h ago'
})

const status = reactive({
  modified: 3,
  added: 2,
  deleted: 1
})

const branches = ref([
  { name: 'main', current: true },
  { name: 'develop', current: false },
  { name: 'feature/user-auth', current: false },
  { name: 'hotfix/bug-fix', current: false }
])

const commits = ref([
  {
    hash: 'a1b2c3d4e5f6',
    message: 'Add user authentication feature',
    author: 'John Doe',
    date: '2024-01-15T10:30:00Z',
    files: 5
  },
  {
    hash: 'b2c3d4e5f6g7',
    message: 'Fix login form validation',
    author: 'Jane Smith',
    date: '2024-01-14T15:45:00Z',
    files: 2
  },
  {
    hash: 'c3d4e5f6g7h8',
    message: 'Update API documentation',
    author: 'Mike Johnson',
    date: '2024-01-13T09:20:00Z',
    files: 8
  }
])

const changedFiles = ref([
  { path: 'src/components/Login.vue', status: 'modified', staged: false },
  { path: 'src/utils/auth.js', status: 'modified', staged: true },
  { path: 'src/api/users.js', status: 'added', staged: false },
  { path: 'docs/README.md', status: 'modified', staged: true },
  { path: 'old-config.json', status: 'deleted', staged: false }
])

const diffModal = reactive({
  show: false,
  file: null,
  content: ''
})

// Methods
function cloneRepository() {
  alert(`مخزن ${repositoryUrl.value} کلون شد`)
}

function switchBranch() {
  alert(`شاخه به ${currentBranch.value} تغییر یافت`)
}

function refreshCommits() {
  alert('تاریخچه کامیت‌ها بروزرسانی شد')
}

function viewCommit(commit) {
  alert(`مشاهده کامیت ${commit.hash}`)
}

function revertCommit(commit) {
  if (confirm(`آیا می‌خواهید کامیت ${commit.hash} را بازگردانید؟`)) {
    alert('کامیت بازگردانده شد')
  }
}

function stageAll() {
  changedFiles.value.forEach(file => file.staged = true)
  alert('همه فایل‌ها به مرحله اضافه شدند')
}

function unstageAll() {
  changedFiles.value.forEach(file => file.staged = false)
  alert('همه فایل‌ها از مرحله حذف شدند')
}

function commitChanges() {
  if (commitMessage.value.trim()) {
    alert(`کامیت با پیام "${commitMessage.value}" ایجاد شد`)
    commitMessage.value = ''
  }
}

function pushChanges() {
  alert('تغییرات ارسال شدند')
}

function pullChanges() {
  alert('تغییرات دریافت شدند')
}

function createBranch() {
  if (newBranchName.value.trim()) {
    branches.value.push({ name: newBranchName.value, current: false })
    alert(`شاخه ${newBranchName.value} ایجاد شد`)
    newBranchName.value = ''
  }
}

function stageFile(file) {
  file.staged = true
}

function unstageFile(file) {
  file.staged = false
}

function viewDiff(file) {
  diffModal.file = file
  diffModal.content = `--- a/${file.path}
+++ b/${file.path}
@@ -1,3 +1,4 @@
 // Sample diff content
+// New line added
 // Another line
-// Removed line
+// Modified line`
  diffModal.show = true
}

function formatDate(dateString) {
  return new Date(dateString).toLocaleString('fa-IR')
}
</script> 

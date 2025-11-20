<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between">
        <div>
          <h2 class="text-xl font-bold text-gray-900">مدیریت پشتیبان‌گیری</h2>
          <p class="text-gray-600 mt-1">پشتیبان‌گیری، بازیابی و مدیریت نسخه‌های پشتیبان کیف پول</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'overview'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
            @click="activeTab = 'overview'"
          >
            نمای کلی
          </button>
          <button
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'backup'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
            @click="activeTab = 'backup'"
          >
            پشتیبان‌گیری
          </button>
          <button
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'restore'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
            @click="activeTab = 'restore'"
          >
            بازیابی
          </button>
          <button
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'settings'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
            @click="activeTab = 'settings'"
          >
            تنظیمات
          </button>
        </div>
      </div>
    </div>

    <!-- محتوای تب‌ها -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200">
      <!-- نمای کلی -->
      <div v-if="activeTab === 'overview'" class="p-6">
        <div class="space-y-6">
          <!-- آمار پشتیبان‌گیری -->
          <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
            <div class="bg-white border border-gray-200 rounded-lg p-6">
              <div class="flex items-center">
                <div class="w-8 h-8 bg-green-100 rounded-full flex items-center justify-center">
                  <svg class="w-4 h-4 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                  </svg>
                </div>
                <div class="mr-3">
                  <p class="text-sm font-medium text-gray-900">آخرین پشتیبان</p>
                  <p class="text-lg font-bold text-green-600">2 ساعت پیش</p>
                </div>
              </div>
            </div>
            <div class="bg-white border border-gray-200 rounded-lg p-6">
              <div class="flex items-center">
                <div class="w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center">
                  <svg class="w-4 h-4 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                  </svg>
                </div>
                <div class="mr-3">
                  <p class="text-sm font-medium text-gray-900">کل نسخه‌ها</p>
                  <p class="text-lg font-bold text-blue-600">24</p>
                </div>
              </div>
            </div>
            <div class="bg-white border border-gray-200 rounded-lg p-6">
              <div class="flex items-center">
                <div class="w-8 h-8 bg-purple-100 rounded-full flex items-center justify-center">
                  <svg class="w-4 h-4 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                </div>
                <div class="mr-3">
                  <p class="text-sm font-medium text-gray-900">پشتیبان بعدی</p>
                  <p class="text-lg font-bold text-purple-600">22:00</p>
                </div>
              </div>
            </div>
            <div class="bg-white border border-gray-200 rounded-lg p-6">
              <div class="flex items-center">
                <div class="w-8 h-8 bg-yellow-100 rounded-full flex items-center justify-center">
                  <svg class="w-4 h-4 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                </div>
                <div class="mr-3">
                  <p class="text-sm font-medium text-gray-900">وضعیت</p>
                  <p class="text-lg font-bold text-yellow-600">فعال</p>
                </div>
              </div>
            </div>
          </div>

          <!-- نمودار پشتیبان‌گیری -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">روند پشتیبان‌گیری (7 روز گذشته)</h3>
            <div class="h-64 flex items-end space-x-2 space-x-reverse overflow-x-auto">
              <div v-for="(day, index) in backupTrend" :key="index" class="flex-shrink-0 flex flex-col items-center min-w-16">
                <div
class="w-full bg-gray-200 rounded-t relative"
                     :style="{ height: getChartHeight(day.size) + 'px' }">
                  <div
class="w-full bg-gradient-to-t from-green-500 to-green-600 rounded-t transition-all duration-300 absolute bottom-0"
                       :style="{ height: getChartHeight(day.size) + 'px' }"></div>
                </div>
                <span class="text-xs text-gray-500 mt-1 text-center">{{ day.date }}</span>
                <span class="text-xs text-gray-400 mt-1 text-center">{{ formatFileSize(day.size) }}</span>
              </div>
            </div>
          </div>

          <!-- خلاصه پشتیبان‌های اخیر -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">پشتیبان‌های اخیر</h3>
            <div class="space-y-3">
              <div v-for="backup in recentBackups" :key="backup.id" class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
                <div class="flex items-center">
                  <div :class="getBackupStatusClass(backup.status)" class="w-8 h-8 rounded-full flex items-center justify-center ml-3">
                    <i :class="getBackupStatusIcon(backup.status)" class="text-white text-sm"></i>
                  </div>
                  <div>
                    <p class="font-medium text-gray-900">{{ backup.name }}</p>
                    <p class="text-sm text-gray-600">{{ backup.description }}</p>
                  </div>
                </div>
                <div class="text-left">
                  <p class="text-sm text-gray-500">{{ backup.time }}</p>
                  <span :class="getBackupTypeClass(backup.type)" class="text-xs px-2 py-1 rounded-full">
                    {{ backup.type }}
                  </span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- پشتیبان‌گیری -->
      <div v-if="activeTab === 'backup'" class="p-6">
        <div class="space-y-6">
          <!-- پشتیبان‌گیری دستی -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">پشتیبان‌گیری دستی</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <h4 class="font-medium text-gray-900 mb-3">نوع پشتیبان‌گیری</h4>
                <div class="space-y-3">
                  <div class="flex items-center">
                    <input v-model="manualBackup.type" type="radio" value="full" class="text-blue-600 focus:ring-blue-500">
                    <label class="mr-2 text-sm text-gray-700">پشتیبان کامل</label>
                  </div>
                  <div class="flex items-center">
                    <input v-model="manualBackup.type" type="radio" value="incremental" class="text-blue-600 focus:ring-blue-500">
                    <label class="mr-2 text-sm text-gray-700">پشتیبان افزایشی</label>
                  </div>
                  <div class="flex items-center">
                    <input v-model="manualBackup.type" type="radio" value="differential" class="text-blue-600 focus:ring-blue-500">
                    <label class="mr-2 text-sm text-gray-700">پشتیبان تفاضلی</label>
                  </div>
                </div>
              </div>
              <div>
                <h4 class="font-medium text-gray-900 mb-3">تنظیمات اضافی</h4>
                <div class="space-y-3">
                  <div class="flex items-center">
                    <input v-model="manualBackup.compress" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                    <label class="mr-2 text-sm text-gray-700">فشرده‌سازی</label>
                  </div>
                  <div class="flex items-center">
                    <input v-model="manualBackup.encrypt" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                    <label class="mr-2 text-sm text-gray-700">رمزگذاری</label>
                  </div>
                  <div class="flex items-center">
                    <input v-model="manualBackup.verify" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                    <label class="mr-2 text-sm text-gray-700">تأیید یکپارچگی</label>
                  </div>
                </div>
              </div>
            </div>
            <div class="mt-6">
              <button class="bg-blue-600 text-white px-6 py-2 rounded-lg hover:bg-blue-700 transition-colors" @click="createManualBackup">
                ایجاد پشتیبان
              </button>
            </div>
          </div>

          <!-- جدول پشتیبان‌ها -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">لیست پشتیبان‌ها</h3>
            <div class="overflow-x-auto">
              <table class="min-w-full divide-y divide-gray-200">
                <thead class="bg-gray-50">
                  <tr>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نام</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">اندازه</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
                  </tr>
                </thead>
                <tbody class="bg-white divide-y divide-gray-200">
                  <tr v-for="backup in backupList" :key="backup.id" class="hover:bg-gray-50">
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ backup.name }}</td>
                    <td class="px-6 py-4 whitespace-nowrap">
                      <span :class="getBackupTypeClass(backup.type)" class="px-2 py-1 text-xs rounded-full">
                        {{ backup.type }}
                      </span>
                    </td>
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ formatFileSize(backup.size) }}</td>
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ backup.date }}</td>
                    <td class="px-6 py-4 whitespace-nowrap">
                      <span :class="getBackupStatusClass(backup.status)" class="px-2 py-1 text-xs rounded-full">
                        {{ backup.status }}
                      </span>
                    </td>
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                      <div class="flex space-x-2 space-x-reverse">
                        <button class="text-blue-600 hover:text-blue-800">دانلود</button>
                        <button class="text-green-600 hover:text-green-800">بازیابی</button>
                        <button class="text-red-600 hover:text-red-800">حذف</button>
                      </div>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>

      <!-- بازیابی -->
      <div v-if="activeTab === 'restore'" class="p-6">
        <div class="space-y-6">
          <!-- انتخاب پشتیبان برای بازیابی -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">بازیابی از پشتیبان</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">انتخاب پشتیبان</label>
                <select v-model="restoreBackup.selectedBackup" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                  <option value="">انتخاب کنید</option>
                  <option v-for="backup in availableBackups" :key="backup.id" :value="backup.id">
                    {{ backup.name }} - {{ backup.date }}
                  </option>
                </select>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">نوع بازیابی</label>
                <select v-model="restoreBackup.type" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                  <option value="full">بازیابی کامل</option>
                  <option value="partial">بازیابی جزئی</option>
                  <option value="selective">بازیابی انتخابی</option>
                </select>
              </div>
            </div>
            <div class="mt-6">
              <button class="bg-green-600 text-white px-6 py-2 rounded-lg hover:bg-green-700 transition-colors" @click="startRestore">
                شروع بازیابی
              </button>
            </div>
          </div>

          <!-- آپلود فایل پشتیبان -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">آپلود فایل پشتیبان</h3>
            <div class="border-2 border-dashed border-gray-300 rounded-lg p-6 text-center">
              <svg class="mx-auto h-12 w-12 text-gray-400" stroke="currentColor" fill="none" viewBox="0 0 48 48">
                <path d="M28 8H12a4 4 0 00-4 4v20m32-12v8m0 0v8a4 4 0 01-4 4H12a4 4 0 01-4-4v-4m32-4l-3.172-3.172a4 4 0 00-5.656 0L28 28M8 32l9.172-9.172a4 4 0 015.656 0L28 28m0 0l4 4m4-24h8m-4-4v8m-12 4h.02" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
              </svg>
              <div class="mt-4">
                <label for="file-upload" class="cursor-pointer bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors">
                  انتخاب فایل
                </label>
                <input id="file-upload" type="file" class="hidden" accept=".zip,.tar,.gz" @change="handleFileUpload">
              </div>
              <p class="text-sm text-gray-600 mt-2">فایل‌های ZIP، TAR یا GZ تا 2GB</p>
            </div>
          </div>

          <!-- تاریخچه بازیابی -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">تاریخچه بازیابی</h3>
            <div class="space-y-3">
              <div v-for="restore in restoreHistory" :key="restore.id" class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
                <div class="flex items-center">
                  <div :class="getRestoreStatusClass(restore.status)" class="w-8 h-8 rounded-full flex items-center justify-center ml-3">
                    <i :class="getRestoreStatusIcon(restore.status)" class="text-white text-sm"></i>
                  </div>
                  <div>
                    <p class="font-medium text-gray-900">{{ restore.backupName }}</p>
                    <p class="text-sm text-gray-600">{{ restore.description }}</p>
                  </div>
                </div>
                <div class="text-left">
                  <p class="text-sm text-gray-500">{{ restore.time }}</p>
                  <span :class="getRestoreStatusClass(restore.status)" class="text-xs px-2 py-1 rounded-full">
                    {{ restore.status }}
                  </span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- تنظیمات -->
      <div v-if="activeTab === 'settings'" class="p-6">
        <div class="space-y-6">
          <!-- تنظیمات خودکار -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات پشتیبان‌گیری خودکار</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">فرکانس پشتیبان‌گیری</label>
                <select v-model="backupSettings.frequency" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                  <option value="hourly">ساعتی</option>
                  <option value="daily">روزانه</option>
                  <option value="weekly">هفتگی</option>
                  <option value="monthly">ماهانه</option>
                </select>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">زمان پشتیبان‌گیری</label>
                <input v-model="backupSettings.time" type="time" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">تعداد نسخه‌های نگهداری</label>
                <input v-model="backupSettings.retention" type="number" min="1" max="100" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">فشرده‌سازی</label>
                <select v-model="backupSettings.compression" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                  <option value="none">بدون فشرده‌سازی</option>
                  <option value="gzip">GZIP</option>
                  <option value="bzip2">BZIP2</option>
                  <option value="lzma">LZMA</option>
                </select>
              </div>
            </div>
          </div>

          <!-- تنظیمات ذخیره‌سازی -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات ذخیره‌سازی</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <h4 class="font-medium text-gray-900 mb-3">محل ذخیره</h4>
                <div class="space-y-3">
                  <div class="flex items-center">
                    <input v-model="backupSettings.storage" type="radio" value="local" class="text-blue-600 focus:ring-blue-500">
                    <label class="mr-2 text-sm text-gray-700">ذخیره محلی</label>
                  </div>
                  <div class="flex items-center">
                    <input v-model="backupSettings.storage" type="radio" value="cloud" class="text-blue-600 focus:ring-blue-500">
                    <label class="mr-2 text-sm text-gray-700">ذخیره ابری</label>
                  </div>
                  <div class="flex items-center">
                    <input v-model="backupSettings.storage" type="radio" value="both" class="text-blue-600 focus:ring-blue-500">
                    <label class="mr-2 text-sm text-gray-700">هر دو</label>
                  </div>
                </div>
              </div>
              <div>
                <h4 class="font-medium text-gray-900 mb-3">تنظیمات امنیت</h4>
                <div class="space-y-3">
                  <div class="flex items-center">
                    <input v-model="backupSettings.encrypt" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                    <label class="mr-2 text-sm text-gray-700">رمزگذاری پشتیبان‌ها</label>
                  </div>
                  <div class="flex items-center">
                    <input v-model="backupSettings.verify" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                    <label class="mr-2 text-sm text-gray-700">تأیید یکپارچگی</label>
                  </div>
                  <div class="flex items-center">
                    <input v-model="backupSettings.notify" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                    <label class="mr-2 text-sm text-gray-700">اعلان پس از پشتیبان‌گیری</label>
                  </div>
                </div>
              </div>
            </div>
            <div class="mt-6">
              <button class="bg-blue-600 text-white px-6 py-2 rounded-lg hover:bg-blue-700 transition-colors" @click="saveBackupSettings">
                ذخیره تنظیمات
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// تعریف reactive data
const activeTab = ref('overview')

// روند پشتیبان‌گیری (7 روز)
const backupTrend = ref([
  { date: 'شنبه', size: 1024 },
  { date: 'یکشنبه', size: 1536 },
  { date: 'دوشنبه', size: 1280 },
  { date: 'سه‌شنبه', size: 2048 },
  { date: 'چهارشنبه', size: 1792 },
  { date: 'پنج‌شنبه', size: 2304 },
  { date: 'جمعه', size: 2560 }
])

// پشتیبان‌های اخیر
const recentBackups = ref([
  {
    id: 1,
    name: 'پشتیبان کامل - 1402/10/15',
    description: 'پشتیبان کامل از تمام داده‌ها',
    time: '2 ساعت پیش',
    type: 'کامل',
    status: 'موفق'
  },
  {
    id: 2,
    name: 'پشتیبان افزایشی - 1402/10/14',
    description: 'پشتیبان افزایشی از تغییرات',
    time: '1 روز پیش',
    type: 'افزایشی',
    status: 'موفق'
  },
  {
    id: 3,
    name: 'پشتیبان تفاضلی - 1402/10/13',
    description: 'پشتیبان تفاضلی از آخرین پشتیبان کامل',
    time: '2 روز پیش',
    type: 'تفاضلی',
    status: 'در حال انجام'
  }
])

// پشتیبان‌گیری دستی
const manualBackup = ref({
  type: 'full',
  compress: true,
  encrypt: true,
  verify: true
})

// لیست پشتیبان‌ها
const backupList = ref([
  {
    id: 1,
    name: 'backup_full_20240131_140000.zip',
    type: 'کامل',
    size: 2048576,
    date: '1402/10/15 14:00',
    status: 'موفق'
  },
  {
    id: 2,
    name: 'backup_incremental_20240130_140000.zip',
    type: 'افزایشی',
    size: 512000,
    date: '1402/10/14 14:00',
    status: 'موفق'
  },
  {
    id: 3,
    name: 'backup_differential_20240129_140000.zip',
    type: 'تفاضلی',
    size: 1024000,
    date: '1402/10/13 14:00',
    status: 'موفق'
  }
])

// بازیابی
const restoreBackup = ref({
  selectedBackup: '',
  type: 'full'
})

// پشتیبان‌های موجود برای بازیابی
const availableBackups = ref([
  { id: 1, name: 'پشتیبان کامل - 1402/10/15', date: '1402/10/15 14:00' },
  { id: 2, name: 'پشتیبان افزایشی - 1402/10/14', date: '1402/10/14 14:00' },
  { id: 3, name: 'پشتیبان تفاضلی - 1402/10/13', date: '1402/10/13 14:00' }
])

// تاریخچه بازیابی
const restoreHistory = ref([
  {
    id: 1,
    backupName: 'پشتیبان کامل - 1402/10/15',
    description: 'بازیابی کامل سیستم',
    time: '1 روز پیش',
    status: 'موفق'
  },
  {
    id: 2,
    backupName: 'پشتیبان افزایشی - 1402/10/14',
    description: 'بازیابی جزئی داده‌ها',
    time: '3 روز پیش',
    status: 'ناموفق'
  }
])

// تنظیمات پشتیبان‌گیری
const backupSettings = ref({
  frequency: 'daily',
  time: '22:00',
  retention: 30,
  compression: 'gzip',
  storage: 'both',
  encrypt: true,
  verify: true,
  notify: true
})

// توابع کمکی
const getChartHeight = (value: number) => {
  const maxValue = Math.max(...backupTrend.value.map(item => item.size))
  const height = 200
  return (value / maxValue) * height
}

const formatFileSize = (bytes: number) => {
  if (bytes === 0) return '0 B'
  const k = 1024
  const sizes = ['B', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

const getBackupStatusClass = (status: string) => {
  const classes = {
    'موفق': 'bg-green-100 text-green-800',
    'در حال انجام': 'bg-yellow-100 text-yellow-800',
    'ناموفق': 'bg-red-100 text-red-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getBackupStatusIcon = (status: string) => {
  const icons = {
    'موفق': 'fas fa-check',
    'در حال انجام': 'fas fa-clock',
    'ناموفق': 'fas fa-times'
  }
  return icons[status] || 'fas fa-info-circle'
}

const getBackupTypeClass = (type: string) => {
  const classes = {
    'کامل': 'bg-blue-100 text-blue-800',
    'افزایشی': 'bg-green-100 text-green-800',
    'تفاضلی': 'bg-purple-100 text-purple-800'
  }
  return classes[type] || 'bg-gray-100 text-gray-800'
}

const getRestoreStatusClass = (status: string) => {
  const classes = {
    'موفق': 'bg-green-100 text-green-800',
    'ناموفق': 'bg-red-100 text-red-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getRestoreStatusIcon = (status: string) => {
  const icons = {
    'موفق': 'fas fa-check',
    'ناموفق': 'fas fa-times'
  }
  return icons[status] || 'fas fa-info-circle'
}

// توابع عملیاتی
const createManualBackup = () => {
  console.log('ایجاد پشتیبان دستی:', manualBackup.value)
}

const startRestore = () => {
  console.log('شروع بازیابی:', restoreBackup.value)
}

const handleFileUpload = (event: Event) => {
  const target = event.target as HTMLInputElement
  if (target.files && target.files[0]) {
    console.log('فایل انتخاب شد:', target.files[0].name)
  }
}

const saveBackupSettings = () => {
  console.log('ذخیره تنظیمات پشتیبان‌گیری:', backupSettings.value)
}
</script> 

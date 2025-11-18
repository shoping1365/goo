<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">تنظیمات پشتیبان‌گیری</h4>
        <p class="text-sm text-gray-600 mt-1">تنظیمات پشتیبان‌گیری و بازیابی داده‌های اتصال حسابداری</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          @click="createBackup"
          :disabled="isCreatingBackup"
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-green-600 hover:bg-green-700 disabled:bg-gray-400 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
        >
          <svg v-if="isCreatingBackup" class="w-4 h-4 ml-2 animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
          <svg v-else class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7H5a2 2 0 00-2 2v9a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-3m-1 4l-3 3m0 0l-3-3m3 3V4" />
          </svg>
          {{ isCreatingBackup ? 'در حال ایجاد...' : 'ایجاد پشتیبان' }}
        </button>
        <button
          @click="saveBackupSettings"
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
          </svg>
          ذخیره تنظیمات
        </button>
      </div>
    </div>

    <!-- آمار پشتیبان‌گیری -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div class="bg-blue-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-blue-600">{{ backupStats.totalBackups }}</div>
            <div class="text-sm text-blue-700">کل پشتیبان‌ها</div>
          </div>
          <div class="w-10 h-10 bg-blue-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7H5a2 2 0 00-2 2v9a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-3m-1 4l-3 3m0 0l-3-3m3 3V4" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-green-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-green-600">{{ backupStats.lastBackup }}</div>
            <div class="text-sm text-green-700">آخرین پشتیبان</div>
          </div>
          <div class="w-10 h-10 bg-green-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-yellow-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-yellow-600">{{ backupStats.totalSize }}</div>
            <div class="text-sm text-yellow-700">حجم کل</div>
          </div>
          <div class="w-10 h-10 bg-yellow-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 7v10c0 2.21 3.582 4 8 4s8-1.79 8-4V7M4 7c0 2.21 3.582 4 8 4s8-1.79 8-4M4 7c0-2.21 3.582-4 8-4s8 1.79 8 4" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-purple-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-purple-600">{{ backupStats.successRate }}</div>
            <div class="text-sm text-purple-700">نرخ موفقیت</div>
          </div>
          <div class="w-10 h-10 bg-purple-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات پشتیبان‌گیری -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- تنظیمات خودکار -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h5 class="text-md font-medium text-gray-900 mb-4">پشتیبان‌گیری خودکار</h5>
        <div class="space-y-4">
          <div>
            <label class="flex items-center">
              <input
                v-model="backupSettings.autoBackup"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">فعال کردن پشتیبان‌گیری خودکار</span>
            </label>
          </div>

          <div v-if="backupSettings.autoBackup">
            <label class="block text-sm font-medium text-gray-700 mb-2">فرکانس پشتیبان‌گیری</label>
            <select
              v-model="backupSettings.backupFrequency"
              class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="daily">روزانه</option>
              <option value="weekly">هفتگی</option>
              <option value="monthly">ماهانه</option>
            </select>
          </div>

          <div v-if="backupSettings.autoBackup">
            <label class="block text-sm font-medium text-gray-700 mb-2">زمان پشتیبان‌گیری</label>
            <input
              v-model="backupSettings.backupTime"
              type="time"
              class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">مدت نگهداری (روز)</label>
            <input
              v-model.number="backupSettings.retentionDays"
              type="number"
              min="1"
              max="365"
              class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
            <p class="text-xs text-gray-500 mt-1">تعداد روز نگهداری پشتیبان‌ها</p>
          </div>

          <div>
            <label class="flex items-center">
              <input
                v-model="backupSettings.compressBackup"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">فشرده‌سازی پشتیبان</span>
            </label>
          </div>

          <div>
            <label class="flex items-center">
              <input
                v-model="backupSettings.encryptBackup"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">رمزگذاری پشتیبان</span>
            </label>
          </div>
        </div>
      </div>

      <!-- تنظیمات محتوا -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h5 class="text-md font-medium text-gray-900 mb-4">محتویات پشتیبان</h5>
        <div class="space-y-4">
          <div>
            <label class="flex items-center">
              <input
                v-model="backupSettings.includeConnections"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">تنظیمات اتصالات</span>
            </label>
          </div>

          <div>
            <label class="flex items-center">
              <input
                v-model="backupSettings.includeMappings"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">نگاشت حساب‌ها</span>
            </label>
          </div>

          <div>
            <label class="flex items-center">
              <input
                v-model="backupSettings.includeTransactions"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">تراکنش‌های همگام‌سازی شده</span>
            </label>
          </div>

          <div>
            <label class="flex items-center">
              <input
                v-model="backupSettings.includeLogs"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">فایل‌های لاگ</span>
            </label>
          </div>

          <div>
            <label class="flex items-center">
              <input
                v-model="backupSettings.includeSettings"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">تنظیمات سیستم</span>
            </label>
          </div>

          <div>
            <label class="flex items-center">
              <input
                v-model="backupSettings.includeUsers"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">اطلاعات کاربران</span>
            </label>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر اندازه پشتیبان (MB)</label>
            <input
              v-model.number="backupSettings.maxBackupSize"
              type="number"
              min="10"
              max="10000"
              class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات ذخیره‌سازی -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- محل ذخیره -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h5 class="text-md font-medium text-gray-900 mb-4">محل ذخیره پشتیبان</h5>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع ذخیره‌سازی</label>
            <select
              v-model="backupSettings.storageType"
              class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="local">محلی</option>
              <option value="ftp">FTP</option>
              <option value="s3">Amazon S3</option>
              <option value="google">Google Drive</option>
              <option value="dropbox">Dropbox</option>
            </select>
          </div>

          <div v-if="backupSettings.storageType === 'local'">
            <label class="block text-sm font-medium text-gray-700 mb-2">مسیر ذخیره</label>
            <input
              v-model="backupSettings.localPath"
              type="text"
              class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="/path/to/backup"
            />
          </div>

          <div v-if="backupSettings.storageType === 'ftp'">
            <div class="grid grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">آدرس سرور</label>
                <input
                  v-model="backupSettings.ftpHost"
                  type="text"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                  placeholder="ftp.example.com"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">پورت</label>
                <input
                  v-model="backupSettings.ftpPort"
                  type="number"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                  placeholder="21"
                />
              </div>
            </div>
            <div class="grid grid-cols-2 gap-6 mt-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">نام کاربری</label>
                <input
                  v-model="backupSettings.ftpUsername"
                  type="text"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">رمز عبور</label>
                <input
                  v-model="backupSettings.ftpPassword"
                  type="password"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                />
              </div>
            </div>
          </div>

          <div v-if="backupSettings.storageType === 's3'">
            <div class="grid grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">Access Key</label>
                <input
                  v-model="backupSettings.s3AccessKey"
                  type="text"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">Secret Key</label>
                <input
                  v-model="backupSettings.s3SecretKey"
                  type="password"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                />
              </div>
            </div>
            <div class="mt-4">
              <label class="block text-sm font-medium text-gray-700 mb-2">نام Bucket</label>
              <input
                v-model="backupSettings.s3Bucket"
                type="text"
                class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="my-backup-bucket"
              />
            </div>
          </div>
        </div>
      </div>

      <!-- تنظیمات اعلان -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h5 class="text-md font-medium text-gray-900 mb-4">اعلان‌های پشتیبان‌گیری</h5>
        <div class="space-y-4">
          <div>
            <label class="flex items-center">
              <input
                v-model="backupSettings.notifyOnSuccess"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">اعلان پس از پشتیبان‌گیری موفق</span>
            </label>
          </div>

          <div>
            <label class="flex items-center">
              <input
                v-model="backupSettings.notifyOnFailure"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">اعلان در صورت خطا</span>
            </label>
          </div>

          <div>
            <label class="flex items-center">
              <input
                v-model="backupSettings.notifyOnStorageFull"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">اعلان پر شدن فضای ذخیره</span>
            </label>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">کانال‌های اعلان</label>
            <div class="space-y-2">
              <label class="flex items-center">
                <input
                  v-model="backupSettings.emailNotifications"
                  type="checkbox"
                  class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                />
                <span class="mr-2 text-sm text-gray-700">ایمیل</span>
              </label>
              <label class="flex items-center">
                <input
                  v-model="backupSettings.smsNotifications"
                  type="checkbox"
                  class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                />
                <span class="mr-2 text-sm text-gray-700">پیامک</span>
              </label>
              <label class="flex items-center">
                <input
                  v-model="backupSettings.pushNotifications"
                  type="checkbox"
                  class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                />
                <span class="mr-2 text-sm text-gray-700">اعلان‌های فوری</span>
              </label>
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">آدرس‌های ایمیل</label>
            <textarea
              v-model="backupSettings.notificationEmails"
              rows="3"
              class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="هر ایمیل در یک خط جداگانه"
            ></textarea>
          </div>
        </div>
      </div>
    </div>

    <!-- لیست پشتیبان‌ها -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">پشتیبان‌های موجود</h5>
      </div>
      <div class="overflow-x-auto">
        <table class="w-full text-sm">
          <thead>
            <tr class="border-b border-gray-200 bg-gray-50">
              <th class="text-right py-3 px-4 font-medium text-gray-600">نام فایل</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">تاریخ ایجاد</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">حجم</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">وضعیت</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">عملیات</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="backup in backups" :key="backup.id" class="border-b border-gray-100 hover:bg-gray-50">
              <td class="py-3 px-4">
                <div class="flex items-center space-x-3 space-x-reverse">
                  <svg class="w-5 h-5 text-blue-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7H5a2 2 0 00-2 2v9a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-3m-1 4l-3 3m0 0l-3-3m3 3V4" />
                  </svg>
                  <span class="font-medium text-gray-900">{{ backup.filename }}</span>
                </div>
              </td>
              <td class="py-3 px-4">
                <div class="text-sm text-gray-900">{{ backup.createdDate }}</div>
                <div class="text-xs text-gray-500">{{ backup.createdTime }}</div>
              </td>
              <td class="py-3 px-4">
                <span class="text-sm text-gray-900">{{ backup.size }}</span>
              </td>
              <td class="py-3 px-4">
                <span
                  class="px-2 py-1 rounded-full text-xs font-medium"
                  :class="getBackupStatusClass(backup.status)"
                >
                  {{ getBackupStatusLabel(backup.status) }}
                </span>
              </td>
              <td class="py-3 px-4">
                <div class="flex items-center space-x-2 space-x-reverse">
                  <button
                    @click="downloadBackup(backup)"
                    class="p-1 text-blue-600 hover:text-blue-800"
                    title="دانلود"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                    </svg>
                  </button>
                  <button
                    @click="restoreBackup(backup)"
                    class="p-1 text-green-600 hover:text-green-800"
                    title="بازیابی"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
                    </svg>
                  </button>
                  <button
                    @click="deleteBackup(backup)"
                    class="p-1 text-red-600 hover:text-red-800"
                    title="حذف"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                    </svg>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// وضعیت ایجاد پشتیبان
const isCreatingBackup = ref(false)

// آمار پشتیبان‌گیری
const backupStats = ref({
  totalBackups: 45,
  lastBackup: 'امروز ۰۲:۳۰',
  totalSize: '۲.۵ GB',
  successRate: 98.5
})

// تنظیمات پشتیبان‌گیری
const backupSettings = ref({
  autoBackup: true,
  backupFrequency: 'daily',
  backupTime: '02:00',
  retentionDays: 30,
  compressBackup: true,
  encryptBackup: true,
  includeConnections: true,
  includeMappings: true,
  includeTransactions: false,
  includeLogs: true,
  includeSettings: true,
  includeUsers: false,
  maxBackupSize: 500,
  storageType: 'local',
  localPath: '/backups/accounting',
  ftpHost: '',
  ftpPort: 21,
  ftpUsername: '',
  ftpPassword: '',
  s3AccessKey: '',
  s3SecretKey: '',
  s3Bucket: '',
  notifyOnSuccess: true,
  notifyOnFailure: true,
  notifyOnStorageFull: true,
  emailNotifications: true,
  smsNotifications: false,
  pushNotifications: true,
  notificationEmails: 'admin@example.com\nmanager@example.com'
})

// پشتیبان‌های موجود
const backups = ref([
  {
    id: 1,
    filename: 'backup_2024_01_15_02_30.zip',
    createdDate: '۱۴۰۲/۱۰/۲۵',
    createdTime: '۰۲:۳۰',
    size: '۱۵۰ MB',
    status: 'completed'
  },
  {
    id: 2,
    filename: 'backup_2024_01_14_02_30.zip',
    createdDate: '۱۴۰۲/۱۰/۲۴',
    createdTime: '۰۲:۳۰',
    size: '۱۴۸ MB',
    status: 'completed'
  },
  {
    id: 3,
    filename: 'backup_2024_01_13_02_30.zip',
    createdDate: '۱۴۰۲/۱۰/۲۳',
    createdTime: '۰۲:۳۰',
    size: '۱۴۵ MB',
    status: 'completed'
  },
  {
    id: 4,
    filename: 'backup_2024_01_12_02_30.zip',
    createdDate: '۱۴۰۲/۱۰/۲۲',
    createdTime: '۰۲:۳۰',
    size: '۱۴۲ MB',
    status: 'failed'
  }
])

// کلاس وضعیت پشتیبان
const getBackupStatusClass = (status: string) => {
  const classes = {
    completed: 'bg-green-100 text-green-700',
    failed: 'bg-red-100 text-red-700',
    in_progress: 'bg-yellow-100 text-yellow-700'
  }
  return classes[status] || 'bg-gray-100 text-gray-700'
}

// برچسب وضعیت پشتیبان
const getBackupStatusLabel = (status: string) => {
  const labels = {
    completed: 'تکمیل شده',
    failed: 'ناموفق',
    in_progress: 'در حال انجام'
  }
  return labels[status] || status
}

// ایجاد پشتیبان
const createBackup = async () => {
  try {
    isCreatingBackup.value = true
    // TODO: ایجاد پشتیبان
    console.log('ایجاد پشتیبان')
    
    // شبیه‌سازی ایجاد پشتیبان
    await new Promise(resolve => setTimeout(resolve, 3000))
    
    console.log('پشتیبان با موفقیت ایجاد شد')
  } catch (error) {
    console.error('خطا در ایجاد پشتیبان:', error)
  } finally {
    isCreatingBackup.value = false
  }
}

// ذخیره تنظیمات پشتیبان‌گیری
const saveBackupSettings = () => {
  // TODO: ذخیره تنظیمات پشتیبان‌گیری
  console.log('ذخیره تنظیمات پشتیبان‌گیری:', backupSettings.value)
}

// دانلود پشتیبان
const downloadBackup = (backup: any) => {
  // TODO: دانلود پشتیبان
  console.log('دانلود پشتیبان:', backup)
}

// بازیابی پشتیبان
const restoreBackup = (backup: any) => {
  // TODO: بازیابی پشتیبان
  console.log('بازیابی پشتیبان:', backup)
}

// حذف پشتیبان
const deleteBackup = (backup: any) => {
  // TODO: حذف پشتیبان
  console.log('حذف پشتیبان:', backup)
}
</script>

<!--
  کامپوننت تنظیمات پشتیبان‌گیری
  شامل:
  1. آمار پشتیبان‌گیری
  2. تنظیمات خودکار
  3. محتویات پشتیبان
  4. محل ذخیره
  5. اعلان‌ها
  6. لیست پشتیبان‌ها
  7. طراحی مدرن و کاملاً ریسپانسیو
--> 

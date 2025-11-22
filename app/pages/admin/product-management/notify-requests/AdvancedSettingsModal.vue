<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
</script>

<script setup lang="ts">
import { useRuntimeConfig } from 'nuxt/app'
import { useNotifyRequests } from '~/composables/useNotifyRequests'

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

const { showSettings, settingsTabs, activeSettingsTab, settings, saveAdvancedSettings } = useNotifyRequests()
const config = useRuntimeConfig()

// ذخیره تنظیمات پیشرفته در بک‌اند
const persistAdvanced = async () => {
  const base = config.public.goApiBase
  const payload = [
    { key: 'notify.message.template', value: settings.value.messageTemplate, category: 'notify', type: 'text' },
    { key: 'notify.user.can_cancel', value: String(settings.value.userCanCancel), category: 'notify', type: 'boolean' },
    { key: 'notify.user.can_edit', value: String(settings.value.userCanEdit), category: 'notify', type: 'boolean' },
    { key: 'notify.user.dashboard', value: String(settings.value.userDashboardEnabled), category: 'notify', type: 'boolean' }
  ]
  await $fetch(`${base}/api/admin/settings`, { method: 'PUT', body: payload })
}
</script>

<template>

  <!-- Settings Modal -->
  <div v-if="showSettings" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
    <div class="bg-white rounded-xl w-full max-w-6xl max-h-[90vh] mx-4 shadow-2xl flex overflow-hidden">
      <!-- Sidebar -->
      <div class="w-64 bg-gray-50 border-l border-gray-200">
        <div class="p-6 border-b border-gray-200">
          <h3 class="text-lg font-semibold text-gray-900">تنظیمات</h3>
          <p class="text-sm text-gray-600 mt-1">پیکربندی سیستم اطلاع‌رسانی</p>
        </div>

        <nav class="p-6 space-y-2">
          <button
              v-for="tab in settingsTabs"
              :key="tab.id"
              :class="[
                'w-full flex items-center gap-3 px-3 py-2 text-sm font-medium rounded-lg transition-all',
                activeSettingsTab === tab.id
                  ? 'bg-blue-100 text-blue-700 border-blue-200 border'
                  : 'text-gray-700 hover:bg-gray-100'
              ]"
              @click="activeSettingsTab = tab.id"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="tab.icon"></path>
            </svg>
            {{ tab.name }}
          </button>
        </nav>
      </div>

      <!-- Content -->
      <div class="flex-1 flex flex-col">
        <!-- Header -->
        <div class="p-6 border-b border-gray-200 flex items-center justify-between">
          <div>
            <h4 class="text-xl font-semibold text-gray-900">{{ settingsTabs.find(t => t.id === activeSettingsTab)?.name }}</h4>
            <p class="text-sm text-gray-600 mt-1">پیکربندی تنظیمات {{ settingsTabs.find(t => t.id === activeSettingsTab)?.name }}</p>
          </div>
          <button class="text-gray-400 hover:text-gray-600 transition-colors" @click="showSettings = false">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>

        <!-- Content Area -->
        <div class="flex-1 overflow-y-auto p-6">
          <!-- تب عمومی -->
          <div v-if="activeSettingsTab === 'general'" class="space-y-6">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">تاخیر ارسال (دقیقه)</label>
                <input
                    v-model="settings.delayMinutes"
                    type="number"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all"
                    placeholder="0"
                >
                <p class="text-xs text-gray-500 mt-1">زمان انتظار قبل از ارسال اطلاع‌رسانی</p>
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر تعداد ارسال روزانه</label>
                <input
                    v-model="settings.dailyLimit"
                    type="number"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all"
                    placeholder="1000"
                >
                <p class="text-xs text-gray-500 mt-1">محدودیت ارسال اطلاع‌رسانی در روز</p>
              </div>
            </div>

            <div class="border-t pt-6">
              <h5 class="text-lg font-medium text-gray-900 mb-4">انواع اطلاع‌رسانی</h5>
              <div class="space-y-4">
                <label class="flex items-center gap-3">
                  <input
                      v-model="settings.emailNotifications"
                      type="checkbox"
                      class="w-4 h-4 text-blue-600 rounded border-gray-300 focus:ring-blue-500"
                  >
                  <span class="text-sm font-medium text-gray-700">اطلاع‌رسانی ایمیل</span>
                </label>

                <label class="flex items-center gap-3">
                  <input
                      v-model="settings.smsNotifications"
                      type="checkbox"
                      class="w-4 h-4 text-blue-600 rounded border-gray-300 focus:ring-blue-500"
                  >
                  <span class="text-sm font-medium text-gray-700">اطلاع‌رسانی پیامک</span>
                </label>

                <label class="flex items-center gap-3">
                  <input
                      v-model="settings.pushNotifications"
                      type="checkbox"
                      class="w-4 h-4 text-blue-600 rounded border-gray-300 focus:ring-blue-500"
                  >
                  <span class="text-sm font-medium text-gray-700">اطلاع‌رسانی پوش</span>
                </label>
              </div>
            </div>
          </div>

          <!-- تب ظاهر -->
          <div v-else-if="activeSettingsTab === 'appearance'" class="space-y-6">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">تم رنگی</label>
                <select
                    v-model="settings.theme"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all"
                >
                  <option value="light">روشن</option>
                  <option value="dark">تیره</option>
                  <option value="auto">خودکار (بر اساس سیستم)</option>
                </select>
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">اندازه فونت</label>
                <select
                    v-model="settings.fontSize"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all"
                >
                  <option value="small">کوچک</option>
                  <option value="medium">متوسط</option>
                  <option value="large">بزرگ</option>
                </select>
              </div>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">رنگ اصلی</label>
              <div class="flex items-center gap-6">
                <input
                    v-model="settings.primaryColor"
                    type="color"
                    class="w-12 h-10 border border-gray-300 rounded-lg cursor-pointer"
                >
                <input
                    v-model="settings.primaryColor"
                    type="text"
                    class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all"
                    placeholder="#3B82F6"
                >
              </div>
              <p class="text-xs text-gray-500 mt-1">رنگ اصلی واسط کاربری</p>
            </div>

            <div class="border-t pt-6">
              <h5 class="text-lg font-medium text-gray-900 mb-4">پیش‌نمایش</h5>
              <div class="bg-gray-50 rounded-lg p-6 border-2 border-dashed border-gray-300">
                <div class="bg-white rounded-lg p-6 shadow-sm">
                  <div class="flex items-center gap-3 mb-3">
                    <div
                        class="w-4 h-4 rounded-full"
                        :style="{ backgroundColor: settings.primaryColor }"
                    ></div>
                    <span
                        :class="['font-medium', settings.fontSize === 'small' ? 'text-sm' : settings.fontSize === 'large' ? 'text-lg' : 'text-base']"
                    >
                        نمونه متن با تنظیمات انتخابی
                      </span>
                  </div>
                  <button
                      class="px-4 py-2 text-white rounded-lg font-medium"
                      :style="{ backgroundColor: settings.primaryColor }"
                  >
                    دکمه نمونه
                  </button>
                </div>
              </div>
            </div>
          </div>

          <!-- تب متون -->
          <div v-else-if="activeSettingsTab === 'content'" class="space-y-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">قالب پیام اطلاع‌رسانی</label>
              <textarea
                  v-model="settings.messageTemplate"
                  rows="4"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all"
                  placeholder="سلام {name}، محصول {product} که درخواست اطلاع‌رسانی کرده‌اید اکنون موجود است."
              ></textarea>
              <p class="text-xs text-gray-500 mt-1">از متغیرهای {name} و {product} استفاده کنید</p>
            </div>
          </div>

          <!-- تب ایمیل -->
          <div v-else-if="activeSettingsTab === 'email'" class="space-y-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">ارائه‌دهنده ایمیل</label>
              <select
                  v-model="settings.emailProvider"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all"
              >
                <option value="smtp">SMTP سفارشی</option>
                <option value="mailgun">Mailgun</option>
                <option value="sendgrid">SendGrid</option>
              </select>
            </div>

            <div v-if="settings.emailProvider === 'smtp'" class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">سرور SMTP</label>
                <input
                    v-model="settings.smtpHost"
                    type="text"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all"
                    placeholder="smtp.gmail.com"
                >
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">پورت</label>
                <input
                    v-model="settings.smtpPort"
                    type="number"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all"
                    placeholder="587"
                >
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">نام کاربری</label>
                <input
                    v-model="settings.smtpUsername"
                    type="text"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all"
                    placeholder="your-email@gmail.com"
                >
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">رمز عبور</label>
                <input
                    v-model="settings.smtpPassword"
                    type="password"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all"
                    placeholder="••••••••"
                >
              </div>
            </div>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">ایمیل فرستنده</label>
                <input
                    v-model="settings.emailFrom"
                    type="email"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all"
                    placeholder="noreply@example.com"
                >
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">نام فرستنده</label>
                <input
                    v-model="settings.emailFromName"
                    type="text"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all"
                    placeholder="فروشگاه ما"
                >
              </div>
            </div>

            <div class="border-t pt-6">
              <h5 class="text-lg font-medium text-gray-900 mb-4">تست اتصال</h5>
              <div class="bg-blue-50 rounded-lg p-6">
                <div class="flex items-center gap-3 mb-3">
                  <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                  </svg>
                  <span class="text-sm text-blue-800">برای اطمینان از صحت تنظیمات، یک ایمیل تست ارسال کنید</span>
                </div>
                <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors text-sm font-medium">
                  ارسال ایمیل تست
                </button>
              </div>
            </div>
          </div>

          <!-- تب پیامک -->
          <div v-else-if="activeSettingsTab === 'sms'" class="space-y-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">ارائه‌دهنده پیامک</label>
              <select
                  v-model="settings.smsProvider"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all"
              >
                <option value="kavenegar">کاوه‌نگار</option>
                <option value="melipayamak">ملی‌پیامک</option>
                <option value="farapayamak">فراپیامک</option>
              </select>
            </div>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">کلید API</label>
                <input
                    v-model="settings.smsApiKey"
                    type="password"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all"
                    placeholder="API Key"
                >
                <p class="text-xs text-gray-500 mt-1">کلید API از پنل ارائه‌دهنده پیامک</p>
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">شماره خط</label>
                <input
                    v-model="settings.smsLineNumber"
                    type="text"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all"
                    placeholder="10009988776655"
                >
                <p class="text-xs text-gray-500 mt-1">شماره خط پیامک اختصاصی</p>
              </div>
            </div>

            <div class="border-t pt-6">
              <h5 class="text-lg font-medium text-gray-900 mb-4">تست اتصال</h5>
              <div class="bg-green-50 rounded-lg p-6">
                <div class="flex items-center gap-3 mb-3">
                  <svg class="w-5 h-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
                  </svg>
                  <span class="text-sm text-green-800">برای اطمینان از صحت تنظیمات، یک پیامک تست ارسال کنید</span>
                </div>
                <div class="flex items-center gap-3">
                  <input
                      type="text"
                      placeholder="شماره موبایل تست"
                      class="flex-1 px-3 py-2 border border-green-300 rounded-lg focus:ring-2 focus:ring-green-500 text-sm"
                  >
                  <button class="bg-green-600 text-white px-4 py-2 rounded-lg hover:bg-green-700 transition-colors text-sm font-medium">
                    ارسال پیامک تست
                  </button>
                </div>
              </div>
            </div>
          </div>

          <!-- تب اطلاع‌رسانی خودکار -->
          <div v-else-if="activeSettingsTab === 'automation'" class="space-y-6">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div class="space-y-4">
                <label class="flex items-center gap-3">
                  <input
                      v-model="settings.autoNotifyStock"
                      type="checkbox"
                      class="w-4 h-4 text-blue-600 rounded border-gray-300 focus:ring-blue-500"
                  >
                  <span class="text-sm font-medium text-gray-700">اطلاع‌رسانی خودکار موجودی</span>
                </label>

                <label class="flex items-center gap-3">
                  <input
                      v-model="settings.autoNotifyDiscount"
                      type="checkbox"
                      class="w-4 h-4 text-blue-600 rounded border-gray-300 focus:ring-blue-500"
                  >
                  <span class="text-sm font-medium text-gray-700">اطلاع‌رسانی خودکار تخفیف</span>
                </label>
              </div>

              <div class="space-y-4">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">تاخیر اطلاع‌رسانی خودکار (دقیقه)</label>
                  <input
                      v-model="settings.autoNotifyDelay"
                      type="number"
                      class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all"
                      placeholder="5"
                  >
                </div>

                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر اطلاع‌رسانی خودکار روزانه</label>
                  <input
                      v-model="settings.maxAutoNotifications"
                      type="number"
                      class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all"
                      placeholder="100"
                  >
                </div>
              </div>
            </div>

            <div class="border-t pt-6">
              <h5 class="text-lg font-medium text-gray-900 mb-4">قوانین خودکار</h5>
              <div class="bg-yellow-50 rounded-lg p-6">
                <div class="flex items-start gap-3">
                  <svg class="w-5 h-5 text-yellow-600 mt-0.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L4.732 15.5c-.77.833.192 2.5 1.732 2.5z"></path>
                  </svg>
                  <div class="text-sm text-yellow-800">
                    <p class="font-medium mb-1">نکات مهم اطلاع‌رسانی خودکار:</p>
                    <ul class="list-disc list-inside space-y-1">
                      <li>اطلاع‌رسانی خودکار فقط در ساعات اداری فعال است</li>
                      <li>حداکثر 3 بار برای هر کاربر اطلاع‌رسانی ارسال می‌شود</li>
                      <li>فاصله زمانی بین اطلاع‌رسانی‌ها حداقل 24 ساعت است</li>
                    </ul>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- تب لاگ‌ها -->
          <div v-else-if="activeSettingsTab === 'logs'" class="space-y-6">
            <div class="flex items-center justify-between">
              <h5 class="text-lg font-medium text-gray-900">فعالیت‌های اخیر</h5>
              <button class="bg-red-500 text-white px-4 py-2 rounded-lg hover:bg-red-600 transition-colors text-sm font-medium">
                پاک کردن لاگ‌ها
              </button>
            </div>

            <div class="bg-white border border-gray-200 rounded-lg overflow-hidden">
              <div class="px-4 py-3 bg-gray-50 border-b border-gray-200">
                <div class="flex items-center justify-between">
                  <span class="text-sm font-medium text-gray-700">آخرین فعالیت‌ها</span>
                  <span class="text-xs text-gray-500">امروز</span>
                </div>
              </div>

              <div class="divide-y divide-gray-100">
                <div class="px-4 py-3 hover:bg-gray-50">
                  <div class="flex items-center gap-3">
                    <div class="w-2 h-2 bg-green-500 rounded-full"></div>
                    <div class="flex-1">
                      <p class="text-sm text-gray-900">اطلاع‌رسانی موجودی برای محصول "گوشی شیائومی" ارسال شد</p>
                      <p class="text-xs text-gray-500">2 دقیقه پیش</p>
                    </div>
                  </div>
                </div>

                <div class="px-4 py-3 hover:bg-gray-50">
                  <div class="flex items-center gap-3">
                    <div class="w-2 h-2 bg-blue-500 rounded-full"></div>
                    <div class="flex-1">
                      <p class="text-sm text-gray-900">درخواست جدید اطلاع‌رسانی تخفیف دریافت شد</p>
                      <p class="text-xs text-gray-500">5 دقیقه پیش</p>
                    </div>
                  </div>
                </div>

                <div class="px-4 py-3 hover:bg-gray-50">
                  <div class="flex items-center gap-3">
                    <div class="w-2 h-2 bg-red-500 rounded-full"></div>
                    <div class="flex-1">
                      <p class="text-sm text-gray-900">خطا در ارسال ایمیل اطلاع‌رسانی</p>
                      <p class="text-xs text-gray-500">10 دقیقه پیش</p>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
              <div class="bg-green-50 rounded-lg p-6">
                <div class="text-2xl font-bold text-green-600">156</div>
                <div class="text-sm text-green-800">اطلاع‌رسانی موفق امروز</div>
              </div>

              <div class="bg-red-50 rounded-lg p-6">
                <div class="text-2xl font-bold text-red-600">3</div>
                <div class="text-sm text-red-800">خطا در ارسال امروز</div>
              </div>

              <div class="bg-blue-50 rounded-lg p-6">
                <div class="text-2xl font-bold text-blue-600">28</div>
                <div class="text-sm text-blue-800">درخواست جدید امروز</div>
              </div>
            </div>
          </div>

          <!-- تب پنل کاربری -->
          <div v-else-if="activeSettingsTab === 'user-panel'" class="space-y-6">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div class="space-y-4">
                <label class="flex items-center gap-3">
                  <input
                      v-model="settings.userDashboardEnabled"
                      type="checkbox"
                      class="w-4 h-4 text-blue-600 rounded border-gray-300 focus:ring-blue-500"
                  >
                  <span class="text-sm font-medium text-gray-700">فعال‌سازی پنل کاربری</span>
                </label>

                <label class="flex items-center gap-3">
                  <input
                      v-model="settings.userCanCancel"
                      type="checkbox"
                      class="w-4 h-4 text-blue-600 rounded border-gray-300 focus:ring-blue-500"
                  >
                  <span class="text-sm font-medium text-gray-700">امکان لغو درخواست توسط کاربر</span>
                </label>

                <label class="flex items-center gap-3">
                  <input
                      v-model="settings.userCanEdit"
                      type="checkbox"
                      class="w-4 h-4 text-blue-600 rounded border-gray-300 focus:ring-blue-500"
                  >
                  <span class="text-sm font-medium text-gray-700">امکان ویرایش درخواست توسط کاربر</span>
                </label>

                <label class="flex items-center gap-3">
                  <input
                      v-model="settings.showUserHistory"
                      type="checkbox"
                      class="w-4 h-4 text-blue-600 rounded border-gray-300 focus:ring-blue-500"
                  >
                  <span class="text-sm font-medium text-gray-700">نمایش تاریخچه به کاربر</span>
                </label>
              </div>

              <div class="bg-gray-50 rounded-lg p-6">
                <h6 class="text-sm font-medium text-gray-900 mb-3">پیش‌نمایش پنل کاربری</h6>
                <div class="bg-white rounded-lg p-3 border border-gray-200 text-sm">
                  <div class="flex items-center justify-between mb-2">
                    <span class="text-gray-700">درخواست‌های من</span>
                    <span class="bg-blue-100 text-blue-800 px-2 py-1 rounded text-xs">3 مورد</span>
                  </div>
                  <div class="space-y-2">
                    <div class="flex items-center justify-between text-xs">
                      <span>گوشی شیائومی</span>
                      <span class="text-green-600">اطلاع‌رسانی شد</span>
                    </div>
                    <div class="flex items-center justify-between text-xs">
                      <span>لپ‌تاپ ایسوس</span>
                      <span class="text-yellow-600">در انتظار</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <div class="border-t pt-6">
              <h5 class="text-lg font-medium text-gray-900 mb-4">لینک پنل کاربری</h5>
              <div class="bg-blue-50 rounded-lg p-6">
                <div class="flex items-center gap-3">
                  <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.102m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1"></path>
                  </svg>
                  <div class="flex-1">
                    <p class="text-sm font-medium text-blue-900 mb-1">آدرس پنل کاربری:</p>
                    <code class="text-sm text-blue-700 bg-blue-100 px-2 py-1 rounded">
                      https://yoursite.com/user/notifications
                    </code>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- سایر تب‌ها -->
          <div v-else class="space-y-6">
            <div class="bg-gray-50 rounded-lg p-6 text-center">
              <svg class="w-16 h-16 text-gray-300 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="settingsTabs.find(t => t.id === activeSettingsTab)?.icon"></path>
              </svg>
              <h5 class="text-lg font-medium text-gray-900 mb-2">{{ settingsTabs.find(t => t.id === activeSettingsTab)?.name }}</h5>
              <p class="text-gray-600">این بخش به زودی اضافه خواهد شد</p>
            </div>
          </div>
        </div>

        <!-- Footer -->
        <div class="p-6 border-t border-gray-200 flex justify-end space-x-3 space-x-reverse">
          <button
              class="px-6 py-2 text-gray-700 border border-gray-300 rounded-lg hover:bg-gray-50 transition-all font-medium"
              @click="showSettings = false"
          >
            لغو
          </button>
          <button
              class="px-6 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 transition-all shadow-md hover:shadow-lg font-medium"
              @click="(async () => { await persistAdvanced(); saveAdvancedSettings() })()"
          >
            ذخیره تغییرات
          </button>
        </div>
      </div>
    </div>
  </div>

</template>

<style scoped>

</style>

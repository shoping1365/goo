<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex justify-between items-center">
        <div>
          <h2 class="text-xl font-bold text-gray-900">تنظیمات نوتیفیکیشن</h2>
          <p class="text-gray-600 mt-1">مدیریت تنظیمات اطلاع‌رسانی سیستم گیفت کارت</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button 
            @click="saveSettings"
            class="px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
          >
            ذخیره تنظیمات
          </button>
          <button 
            @click="resetSettings"
            class="px-4 py-2 bg-gray-600 text-white text-sm font-medium rounded-lg hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
          >
            بازنشانی
          </button>
        </div>
      </div>
    </div>

    <!-- تب‌های تنظیمات -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200">
      <div class="border-b border-gray-200">
        <nav class="-mb-px flex space-x-8 space-x-reverse" aria-label="Tabs">
          <button
            v-for="tab in tabs"
            :key="tab.id"
            @click="activeTab = tab.id"
            :class="{
              'border-blue-500 text-blue-600': activeTab === tab.id,
              'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300': activeTab !== tab.id
            }"
            class="whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm"
          >
            {{ tab.name }}
          </button>
        </nav>
      </div>

      <div class="p-6">
        <!-- تنظیمات عمومی -->
        <div v-if="activeTab === 'general'" class="space-y-6">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">فعال‌سازی سیستم نوتیفیکیشن</label>
              <div class="flex items-center">
                <input
                  v-model="settings.general.enabled"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label class="mr-2 text-sm text-gray-700">فعال</label>
              </div>
              <p class="text-xs text-gray-500 mt-1">فعال یا غیرفعال کردن کل سیستم نوتیفیکیشن</p>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نمایش نوتیفیکیشن در داشبورد</label>
              <div class="flex items-center">
                <input
                  v-model="settings.general.showInDashboard"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label class="mr-2 text-sm text-gray-700">نمایش</label>
              </div>
              <p class="text-xs text-gray-500 mt-1">نمایش نوتیفیکیشن‌ها در داشبورد ادمین</p>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر تعداد نوتیفیکیشن نمایشی</label>
              <input
                v-model.number="settings.general.maxDisplayCount"
                type="number"
                min="1"
                max="100"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              />
              <p class="text-xs text-gray-500 mt-1">حداکثر تعداد نوتیفیکیشن‌هایی که در لیست نمایش داده می‌شوند</p>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">مدت نگهداری نوتیفیکیشن‌ها (روز)</label>
              <input
                v-model.number="settings.general.retentionDays"
                type="number"
                min="1"
                max="365"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              />
              <p class="text-xs text-gray-500 mt-1">نوتیفیکیشن‌های قدیمی‌تر از این تعداد روز حذف می‌شوند</p>
            </div>
          </div>
        </div>

        <!-- تنظیمات خرید جدید -->
        <div v-if="activeTab === 'purchase'" class="space-y-6">
          <div class="bg-blue-50 border border-blue-200 rounded-lg p-6">
            <h3 class="text-sm font-medium text-blue-900 mb-2">اطلاع‌رسانی خرید جدید</h3>
            <p class="text-sm text-blue-700">تنظیمات مربوط به اطلاع‌رسانی هنگام خرید گیفت کارت جدید</p>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">فعال‌سازی اطلاع‌رسانی خرید</label>
              <div class="flex items-center">
                <input
                  v-model="settings.purchase.enabled"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label class="mr-2 text-sm text-gray-700">فعال</label>
              </div>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">حداقل مبلغ برای اطلاع‌رسانی</label>
              <input
                v-model.number="settings.purchase.minAmount"
                type="number"
                min="0"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              />
              <p class="text-xs text-gray-500 mt-1">فقط خریدهای بالاتر از این مبلغ اطلاع‌رسانی می‌شوند</p>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">اولویت پیش‌فرض</label>
              <select
                v-model="settings.purchase.defaultPriority"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="low">پایین</option>
                <option value="medium">متوسط</option>
                <option value="high">بالا</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">ارسال ایمیل</label>
              <div class="flex items-center">
                <input
                  v-model="settings.purchase.sendEmail"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label class="mr-2 text-sm text-gray-700">فعال</label>
              </div>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">ارسال پیامک</label>
              <div class="flex items-center">
                <input
                  v-model="settings.purchase.sendSMS"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label class="mr-2 text-sm text-gray-700">فعال</label>
              </div>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نمایش در داشبورد</label>
              <div class="flex items-center">
                <input
                  v-model="settings.purchase.showInDashboard"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label class="mr-2 text-sm text-gray-700">فعال</label>
              </div>
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">متغیرهای قالب پیام</label>
            <div class="bg-gray-50 rounded-lg p-6">
              <p class="text-sm text-gray-600 mb-2">متغیرهای قابل استفاده در قالب پیام:</p>
              <div class="grid grid-cols-2 md:grid-cols-4 gap-2 text-xs">
                <code class="bg-white px-2 py-1 rounded border">{{ '{card_code}' }}</code>
                <code class="bg-white px-2 py-1 rounded border">{{ '{amount}' }}</code>
                <code class="bg-white px-2 py-1 rounded border">{{ '{buyer_name}' }}</code>
                <code class="bg-white px-2 py-1 rounded border">{{ '{payment_method}' }}</code>
                <code class="bg-white px-2 py-1 rounded border">{{ '{purchase_date}' }}</code>
                <code class="bg-white px-2 py-1 rounded border">{{ '{expiry_date}' }}</code>
                <code class="bg-white px-2 py-1 rounded border">{{ '{recipient_name}' }}</code>
                <code class="bg-white px-2 py-1 rounded border">{{ '{card_type}' }}</code>
              </div>
            </div>
          </div>
        </div>

        <!-- تنظیمات هشدار انقضا -->
        <div v-if="activeTab === 'expiry'" class="space-y-6">
          <div class="bg-yellow-50 border border-yellow-200 rounded-lg p-6">
            <h3 class="text-sm font-medium text-yellow-900 mb-2">هشدار انقضای نزدیک</h3>
            <p class="text-sm text-yellow-700">تنظیمات مربوط به اطلاع‌رسانی کارت‌هایی که به زودی منقضی می‌شوند</p>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">فعال‌سازی هشدار انقضا</label>
              <div class="flex items-center">
                <input
                  v-model="settings.expiry.enabled"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label class="mr-2 text-sm text-gray-700">فعال</label>
              </div>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">روزهای قبل از انقضا</label>
              <input
                v-model.number="settings.expiry.daysBeforeExpiry"
                type="number"
                min="1"
                max="30"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              />
              <p class="text-xs text-gray-500 mt-1">چند روز قبل از انقضا هشدار ارسال شود</p>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">اولویت پیش‌فرض</label>
              <select
                v-model="settings.expiry.defaultPriority"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="low">پایین</option>
                <option value="medium">متوسط</option>
                <option value="high">بالا</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">حداقل مبلغ باقی‌مانده</label>
              <input
                v-model.number="settings.expiry.minRemainingAmount"
                type="number"
                min="0"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              />
              <p class="text-xs text-gray-500 mt-1">فقط کارت‌هایی با مبلغ باقی‌مانده بالاتر از این مقدار هشدار می‌گیرند</p>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">ارسال ایمیل</label>
              <div class="flex items-center">
                <input
                  v-model="settings.expiry.sendEmail"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label class="mr-2 text-sm text-gray-700">فعال</label>
              </div>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">ارسال پیامک</label>
              <div class="flex items-center">
                <input
                  v-model="settings.expiry.sendSMS"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label class="mr-2 text-sm text-gray-700">فعال</label>
              </div>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نمایش در داشبورد</label>
              <div class="flex items-center">
                <input
                  v-model="settings.expiry.showInDashboard"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label class="mr-2 text-sm text-gray-700">فعال</label>
              </div>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">تکرار هشدار</label>
              <select
                v-model="settings.expiry.repeatInterval"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="once">فقط یک بار</option>
                <option value="daily">روزانه</option>
                <option value="weekly">هفتگی</option>
              </select>
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">متغیرهای قالب پیام</label>
            <div class="bg-gray-50 rounded-lg p-6">
              <p class="text-sm text-gray-600 mb-2">متغیرهای قابل استفاده در قالب پیام:</p>
              <div class="grid grid-cols-2 md:grid-cols-4 gap-2 text-xs">
                <code class="bg-white px-2 py-1 rounded border">{{ '{card_code}' }}</code>
                <code class="bg-white px-2 py-1 rounded border">{{ '{remaining_amount}' }}</code>
                <code class="bg-white px-2 py-1 rounded border">{{ '{expiry_date}' }}</code>
                <code class="bg-white px-2 py-1 rounded border">{{ '{days_remaining}' }}</code>
                <code class="bg-white px-2 py-1 rounded border">{{ '{recipient_name}' }}</code>
                <code class="bg-white px-2 py-1 rounded border">{{ '{card_type}' }}</code>
                <code class="bg-white px-2 py-1 rounded border">{{ '{purchase_date}' }}</code>
                <code class="bg-white px-2 py-1 rounded border">{{ '{buyer_name}' }}</code>
              </div>
            </div>
          </div>
        </div>

        <!-- تنظیمات تراکنش‌های مشکوک -->
        <div v-if="activeTab === 'suspicious'" class="space-y-6">
          <div class="bg-red-50 border border-red-200 rounded-lg p-6">
            <h3 class="text-sm font-medium text-red-900 mb-2">گزارش تراکنش‌های مشکوک</h3>
            <p class="text-sm text-red-700">تنظیمات مربوط به شناسایی و اطلاع‌رسانی تراکنش‌های مشکوک</p>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">فعال‌سازی تشخیص تراکنش مشکوک</label>
              <div class="flex items-center">
                <input
                  v-model="settings.suspicious.enabled"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label class="mr-2 text-sm text-gray-700">فعال</label>
              </div>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">حد آستانه مبلغ مشکوک</label>
              <input
                v-model.number="settings.suspicious.amountThreshold"
                type="number"
                min="0"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              />
              <p class="text-xs text-gray-500 mt-1">تراکنش‌های بالاتر از این مبلغ مشکوک محسوب می‌شوند</p>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">اولویت پیش‌فرض</label>
              <select
                v-model="settings.suspicious.defaultPriority"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="low">پایین</option>
                <option value="medium">متوسط</option>
                <option value="high">بالا</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر تعداد تلاش ناموفق</label>
              <input
                v-model.number="settings.suspicious.maxFailedAttempts"
                type="number"
                min="1"
                max="10"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              />
              <p class="text-xs text-gray-500 mt-1">تعداد تلاش‌های ناموفق قبل از علامت‌گذاری به عنوان مشکوک</p>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">ارسال ایمیل</label>
              <div class="flex items-center">
                <input
                  v-model="settings.suspicious.sendEmail"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label class="mr-2 text-sm text-gray-700">فعال</label>
              </div>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">ارسال پیامک</label>
              <div class="flex items-center">
                <input
                  v-model="settings.suspicious.sendSMS"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label class="mr-2 text-sm text-gray-700">فعال</label>
              </div>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نمایش در داشبورد</label>
              <div class="flex items-center">
                <input
                  v-model="settings.suspicious.showInDashboard"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label class="mr-2 text-sm text-gray-700">فعال</label>
              </div>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">مسدود کردن خودکار</label>
              <div class="flex items-center">
                <input
                  v-model="settings.suspicious.autoBlock"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label class="mr-2 text-sm text-gray-700">فعال</label>
              </div>
              <p class="text-xs text-gray-500 mt-1">مسدود کردن خودکار کارت در صورت تشخیص تراکنش مشکوک</p>
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">معیارهای تشخیص تراکنش مشکوک</label>
            <div class="space-y-3">
              <div class="flex items-center">
                <input
                  v-model="settings.suspicious.criteria.largeAmount"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label class="mr-2 text-sm text-gray-700">مبلغ بالا</label>
              </div>
              <div class="flex items-center">
                <input
                  v-model="settings.suspicious.criteria.multipleAttempts"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label class="mr-2 text-sm text-gray-700">تلاش‌های متعدد</label>
              </div>
              <div class="flex items-center">
                <input
                  v-model="settings.suspicious.criteria.unusualTime"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label class="mr-2 text-sm text-gray-700">زمان غیرعادی</label>
              </div>
              <div class="flex items-center">
                <input
                  v-model="settings.suspicious.criteria.differentIP"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label class="mr-2 text-sm text-gray-700">IP متفاوت</label>
              </div>
              <div class="flex items-center">
                <input
                  v-model="settings.suspicious.criteria.rapidTransactions"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label class="mr-2 text-sm text-gray-700">تراکنش‌های سریع</label>
              </div>
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">متغیرهای قالب پیام</label>
            <div class="bg-gray-50 rounded-lg p-6">
              <p class="text-sm text-gray-600 mb-2">متغیرهای قابل استفاده در قالب پیام:</p>
              <div class="grid grid-cols-2 md:grid-cols-4 gap-2 text-xs">
                <code class="bg-white px-2 py-1 rounded border">{{ '{card_code}' }}</code>
                <code class="bg-white px-2 py-1 rounded border">{{ '{transaction_amount}' }}</code>
                <code class="bg-white px-2 py-1 rounded border">{{ '{transaction_time}' }}</code>
                <code class="bg-white px-2 py-1 rounded border">{{ '{ip_address}' }}</code>
                <code class="bg-white px-2 py-1 rounded border">{{ '{user_agent}' }}</code>
                <code class="bg-white px-2 py-1 rounded border">{{ '{suspicion_reason}' }}</code>
                <code class="bg-white px-2 py-1 rounded border">{{ '{attempt_count}' }}</code>
                <code class="bg-white px-2 py-1 rounded border">{{ '{location}' }}</code>
              </div>
            </div>
          </div>
        </div>

        <!-- تنظیمات پیشرفته -->
        <div v-if="activeTab === 'advanced'" class="space-y-6">
          <div class="bg-purple-50 border border-purple-200 rounded-lg p-6">
            <h3 class="text-sm font-medium text-purple-900 mb-2">تنظیمات پیشرفته</h3>
            <p class="text-sm text-purple-700">تنظیمات پیشرفته و تخصصی سیستم نوتیفیکیشن</p>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">فاصله زمانی بین نوتیفیکیشن‌ها (دقیقه)</label>
              <input
                v-model.number="settings.advanced.notificationInterval"
                type="number"
                min="1"
                max="60"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              />
              <p class="text-xs text-gray-500 mt-1">حداقل فاصله زمانی بین ارسال نوتیفیکیشن‌های مشابه</p>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر تعداد نوتیفیکیشن در روز</label>
              <input
                v-model.number="settings.advanced.maxDailyNotifications"
                type="number"
                min="1"
                max="1000"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              />
              <p class="text-xs text-gray-500 mt-1">حداکثر تعداد نوتیفیکیشن‌هایی که در روز ارسال می‌شوند</p>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">زمان شروع ارسال (ساعت)</label>
              <input
                v-model.number="settings.advanced.startTime"
                type="number"
                min="0"
                max="23"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              />
              <p class="text-xs text-gray-500 mt-1">ساعت شروع ارسال نوتیفیکیشن‌ها (24 ساعته)</p>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">زمان پایان ارسال (ساعت)</label>
              <input
                v-model.number="settings.advanced.endTime"
                type="number"
                min="0"
                max="23"
                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              />
              <p class="text-xs text-gray-500 mt-1">ساعت پایان ارسال نوتیفیکیشن‌ها (24 ساعته)</p>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">فعال‌سازی لاگ کامل</label>
              <div class="flex items-center">
                <input
                  v-model="settings.advanced.enableFullLogging"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label class="mr-2 text-sm text-gray-700">فعال</label>
              </div>
              <p class="text-xs text-gray-500 mt-1">ثبت تمام جزئیات نوتیفیکیشن‌ها در لاگ</p>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">فعال‌سازی تست خودکار</label>
              <div class="flex items-center">
                <input
                  v-model="settings.advanced.enableAutoTest"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label class="mr-2 text-sm text-gray-700">فعال</label>
              </div>
              <p class="text-xs text-gray-500 mt-1">تست خودکار سیستم نوتیفیکیشن هر 24 ساعت</p>
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">کاربران دریافت‌کننده نوتیفیکیشن</label>
            <div class="space-y-2">
              <div v-for="role in availableRoles" :key="role.id" class="flex items-center">
                <input
                  v-model="settings.advanced.notificationRecipients"
                  :value="role.id"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label class="mr-2 text-sm text-gray-700">{{ role.name }}</label>
              </div>
            </div>
            <p class="text-xs text-gray-500 mt-1">انتخاب کنید کدام نقش‌های کاربری نوتیفیکیشن دریافت کنند</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'

// Reactive data
const activeTab = ref('general')

// تب‌های تنظیمات
const tabs = [
  { id: 'general', name: 'تنظیمات عمومی' },
  { id: 'purchase', name: 'خرید جدید' },
  { id: 'expiry', name: 'هشدار انقضا' },
  { id: 'suspicious', name: 'تراکنش مشکوک' },
  { id: 'advanced', name: 'پیشرفته' }
]

// تنظیمات پیش‌فرض
const defaultSettings = {
  general: {
    enabled: true,
    showInDashboard: true,
    maxDisplayCount: 50,
    retentionDays: 30
  },
  purchase: {
    enabled: true,
    minAmount: 100000,
    defaultPriority: 'medium',
    sendEmail: true,
    sendSMS: false,
    showInDashboard: true
  },
  expiry: {
    enabled: true,
    daysBeforeExpiry: 7,
    defaultPriority: 'high',
    minRemainingAmount: 50000,
    sendEmail: true,
    sendSMS: true,
    showInDashboard: true,
    repeatInterval: 'daily'
  },
  suspicious: {
    enabled: true,
    amountThreshold: 1000000,
    defaultPriority: 'high',
    maxFailedAttempts: 3,
    sendEmail: true,
    sendSMS: true,
    showInDashboard: true,
    autoBlock: false,
    criteria: {
      largeAmount: true,
      multipleAttempts: true,
      unusualTime: true,
      differentIP: true,
      rapidTransactions: true
    }
  },
  advanced: {
    notificationInterval: 5,
    maxDailyNotifications: 100,
    startTime: 8,
    endTime: 22,
    enableFullLogging: false,
    enableAutoTest: true,
    notificationRecipients: ['admin', 'finance']
  }
}

// تنظیمات فعلی
const settings = reactive(JSON.parse(JSON.stringify(defaultSettings)))

// نقش‌های موجود
const availableRoles = [
  { id: 'admin', name: 'مدیر سیستم' },
  { id: 'finance', name: 'مدیر مالی' },
  { id: 'support', name: 'پشتیبانی' },
  { id: 'sales', name: 'فروش' }
]

// Methods
const saveSettings = async () => {
  try {
    console.log('تنظیمات در حال ذخیره...', settings)
    
    // اینجا API call برای ذخیره تنظیمات اضافه می‌شود
    // await $fetch('/api/admin/giftcard/notification-settings', {
    //   method: 'POST',
    //   body: settings
    // })
    
    // نمایش پیام موفقیت
    alert('تنظیمات با موفقیت ذخیره شد')
  } catch (error) {
    console.error('خطا در ذخیره تنظیمات:', error)
    alert('خطا در ذخیره تنظیمات')
  }
}

const resetSettings = () => {
  if (confirm('آیا مطمئن هستید که می‌خواهید تنظیمات را به حالت پیش‌فرض بازگردانید؟')) {
    Object.assign(settings, JSON.parse(JSON.stringify(defaultSettings)))
    console.log('تنظیمات به حالت پیش‌فرض بازگردانده شد')
  }
}

const loadSettings = async () => {
  try {
    console.log('در حال بارگذاری تنظیمات...')
    
    // اینجا API call برای دریافت تنظیمات اضافه می‌شود
    // const response = await $fetch('/api/admin/giftcard/notification-settings')
    // Object.assign(settings, response)
    
    console.log('تنظیمات بارگذاری شد')
  } catch (error) {
    console.error('خطا در بارگذاری تنظیمات:', error)
  }
}

// Lifecycle
onMounted(() => {
  loadSettings()
  console.log('Gift card notification settings component mounted')
})
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 

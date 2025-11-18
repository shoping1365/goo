<template>
  <div dir="rtl" class="min-h-screen flex flex-col">
    <!-- Developer Only Access -->
    <div v-if="!isDeveloper" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg p-8 text-center">
        <h2 class="text-2xl font-bold text-red-600 mb-4">دسترسی ممنوع</h2>
        <p class="text-gray-700">این بخش فقط برای توسعه‌دهندگان قابل دسترسی است.</p>
      </div>
    </div>
    <!-- تاپ‌بار -->
    <div class="flex items-center justify-between bg-white shadow p-6 mb-6 rounded-b-lg">
      <h1 class="text-2xl font-bold text-gray-800">مانیتورینگ دیتابیس</h1>
      <div class="flex gap-2 items-center">
        <div :class="[
          'px-6 py-2 rounded-lg text-white text-lg font-bold',
          connectionStatus === 'connected' ? 'bg-green-500' : 'bg-red-500'
        ]">
          وضعیت اتصال :
          <span>
            {{ connectionStatus === 'connected' ? 'متصل' : 'قطع' }}
          </span>
        </div>
        <button @click="testConnection" class="flex items-center gap-1 bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded-lg shadow">
          <span>تست مجدد</span>
          <i class="icon-refresh"></i>
        </button>
        <button @click="showSettings = true" class="flex items-center gap-1 bg-gray-500 hover:bg-gray-600 text-white px-4 py-2 rounded-lg shadow">
          <span>تنظیمات</span>
          <i class="icon-settings"></i>
        </button>
        <button @click="resetDatabase" class="flex items-center gap-1 bg-yellow-400 hover:bg-yellow-500 text-white px-4 py-2 rounded-lg shadow">
          <span>ریست دیتابیس</span>
          <i class="icon-reset"></i>
        </button>
      </div>
    </div>
    <!-- کارت‌های آمار -->
    <div class="grid grid-cols-1 md:grid-cols-5 gap-6 mb-6">
      <div class="rounded-xl p-6 text-center" style="background: #e6fcff;">
        <h3 class="text-sm font-bold mb-1" style="color: #00bcd4;">نام دیتابیس</h3>
        <p class="text-2xl font-extrabold" style="color: #0097a7;">{{ dbInfo?.database ?? '-' }}</p>
      </div>
      <div class="rounded-xl p-6 text-center" style="background: #f6e6ff;">
        <h3 class="text-sm font-bold mb-1" style="color: #a259e6;">زمان پاسخ</h3>
        <p class="text-2xl font-extrabold" style="color: #8e24aa;">{{ systemStats?.responseTime ?? '-' }}ms</p>
      </div>
      <div class="rounded-xl p-6 text-center" style="background: #e6fff6;">
        <h3 class="text-sm font-bold mb-1" style="color: #00c853;">حجم دیتابیس</h3>
        <p class="text-2xl font-extrabold" style="color: #388e3c;">{{ formatDbSize(systemStats?.databaseSize) }}</p>
      </div>
      <div class="rounded-xl p-6 text-center" style="background: #e6f2ff;">
        <h3 class="text-sm font-bold mb-1" style="color: #1976d2;">تعداد اتصالات فعال</h3>
        <p class="text-2xl font-extrabold" style="color: #1565c0;">{{ systemStats?.activeConnections ?? '-' }}</p>
      </div>
      <div class="rounded-xl p-6 text-center" style="background: #fff6e6;">
        <h3 class="text-sm font-bold mb-1" style="color: #ff9800;">مدت زمان اتصال</h3>
        <p class="text-2xl font-extrabold" style="color: #e65100;">{{ formatUptimeStr(dbUptime) }}</p>
      </div>
    </div>
    <!-- جدول کوئری‌های فعال -->
    <div class="bg-white rounded-lg shadow p-6 mb-6">
      <div class="flex items-center justify-between cursor-pointer" @click="activeQueriesExpanded = !activeQueriesExpanded">
        <h2 class="text-lg font-bold text-gray-800">کوئری‌های فعال</h2>
        <span class="text-2xl select-none text-gray-400 w-8 text-center inline-block">{{ activeQueriesExpanded ? '–' : '+' }}</span>
      </div>
      <div v-show="activeQueriesExpanded" class="mt-4">
        <div v-if="activeQueries.length > 5" class="overflow-y-auto max-h-80 min-h-0">
          <table class="min-w-full text-sm text-right">
            <thead>
              <tr class="bg-gray-100">
                <th class="px-4 py-2">کاربر</th>
                <th class="px-4 py-2">آدرس کلاینت</th>
                <th class="px-4 py-2">مدت اجرا</th>
                <th class="px-4 py-2">زمان اجرا</th>
                <th class="px-4 py-2">زمان توقف</th>
                <th class="px-4 py-2">وضعیت</th>
                <th class="px-4 py-2">کوئری</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="query in activeQueries" :key="query.pid">
                <td class="px-4 py-2">{{ query.usename }}</td>
                <td class="px-4 py-2">{{ query.client_addr }}</td>
                <td class="px-4 py-2">{{ query.duration }}</td>
                <td class="px-4 py-2">{{ formatStartTime(query) }}</td>
                <td class="px-4 py-2 text-gray-400">در حال اجرا</td>
                <td class="px-4 py-2">{{ query.state }}</td>
                <td class="px-4 py-2">{{ query.query }}</td>
              </tr>
              <tr v-if="!activeQueries.length">
                <td colspan="7" class="text-center text-gray-400 py-4">کوئری فعالی وجود ندارد</td>
              </tr>
            </tbody>
          </table>
        </div>
        <table v-else class="min-w-full text-sm text-right">
          <thead>
            <tr class="bg-gray-100">
              <th class="px-4 py-2">کاربر</th>
              <th class="px-4 py-2">آدرس کلاینت</th>
              <th class="px-4 py-2">مدت اجرا</th>
              <th class="px-4 py-2">زمان اجرا</th>
              <th class="px-4 py-2">زمان توقف</th>
              <th class="px-4 py-2">وضعیت</th>
              <th class="px-4 py-2">کوئری</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="query in activeQueries" :key="query.pid">
              <td class="px-4 py-2">{{ query.usename }}</td>
              <td class="px-4 py-2">{{ query.client_addr }}</td>
              <td class="px-4 py-2">{{ query.duration }}</td>
              <td class="px-4 py-2">{{ formatStartTime(query) }}</td>
              <td class="px-4 py-2 text-gray-400">در حال اجرا</td>
              <td class="px-4 py-2">{{ query.state }}</td>
              <td class="px-4 py-2">{{ query.query }}</td>
            </tr>
            <tr v-if="!activeQueries.length">
              <td colspan="7" class="text-center text-gray-400 py-4">کوئری فعالی وجود ندارد</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
    <!-- جدول کوئری‌های کند -->
    <div class="bg-white rounded-lg shadow p-6 mb-6">
      <div class="flex items-center justify-between cursor-pointer" @click="slowQueriesExpanded = !slowQueriesExpanded">
        <h2 class="text-lg font-bold text-gray-800">کوئری‌های کند</h2>
        <span class="text-2xl select-none text-gray-400 w-8 text-center inline-block">{{ slowQueriesExpanded ? '–' : '+' }}</span>
      </div>
      <div v-show="slowQueriesExpanded" class="mt-4">
        <div v-if="slowQueries.length > 5" class="overflow-y-auto max-h-80 min-h-0">
          <table class="min-w-full text-sm text-right">
            <thead>
              <tr class="bg-gray-100">
                <th class="px-4 py-2">تعداد اجرا</th>
                <th class="px-4 py-2">میانگین زمان</th>
                <th class="px-4 py-2">حداکثر زمان</th>
                <th class="px-4 py-2">مدت اجرا</th>
                <th class="px-4 py-2">کوئری</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="query in slowQueries" :key="query.queryid">
                <td class="px-4 py-2">{{ query.calls }}</td>
                <td class="px-4 py-2">{{ query.mean_time }}ms</td>
                <td class="px-4 py-2">{{ query.max_time }}ms</td>
                <td class="px-4 py-2">{{ query.total_exec_time ? query.total_exec_time + 'ms' : '-' }}</td>
                <td class="px-4 py-2">{{ query.query }}</td>
              </tr>
              <tr v-if="!slowQueries.length">
                <td colspan="5" class="text-center text-gray-400 py-4">کوئری کندی وجود ندارد یا افزونه فعال نیست</td>
              </tr>
            </tbody>
          </table>
        </div>
        <table v-else class="min-w-full text-sm text-right">
          <thead>
            <tr class="bg-gray-100">
              <th class="px-4 py-2">تعداد اجرا</th>
              <th class="px-4 py-2">میانگین زمان</th>
              <th class="px-4 py-2">حداکثر زمان</th>
              <th class="px-4 py-2">مدت اجرا</th>
              <th class="px-4 py-2">کوئری</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="query in slowQueries" :key="query.queryid">
              <td class="px-4 py-2">{{ query.calls }}</td>
              <td class="px-4 py-2">{{ query.mean_time }}ms</td>
              <td class="px-4 py-2">{{ query.max_time }}ms</td>
              <td class="px-4 py-2">{{ query.total_exec_time ? query.total_exec_time + 'ms' : '-' }}</td>
              <td class="px-4 py-2">{{ query.query }}</td>
            </tr>
            <tr v-if="!slowQueries.length">
              <td colspan="5" class="text-center text-gray-400 py-4">کوئری کندی وجود ندارد یا افزونه فعال نیست</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
    <!-- اجرای کوئری دلخواه -->
    <div class="bg-white rounded-lg shadow p-6 mb-6">
      <div class="flex items-center justify-between cursor-pointer" @click="customQueryExpanded = !customQueryExpanded">
        <h2 class="text-lg font-bold text-gray-800">اجرای کوئری دلخواه</h2>
        <span class="text-2xl select-none text-gray-400 w-8 text-center inline-block">{{ customQueryExpanded ? '–' : '+' }}</span>
      </div>
      <div v-show="customQueryExpanded" class="mt-4">
        <div class="p-6 rounded-lg">
          <div class="flex gap-2 mb-4 justify-start">
            <input v-model="query" type="text" placeholder="مثال: SELECT * FROM users LIMIT 5" class="flex-1 rounded-lg border-gray-300 focus:border-blue-500 focus:ring-1 focus:ring-blue-500 px-4 py-2 text-left bg-blue-50 ltr:text-left" dir="ltr" />
            <button @click="runQuery" :disabled="queryLoading" class="bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded-lg shadow disabled:opacity-50">اجرا</button>
            <button @click="resetQuery" class="bg-gray-200 hover:bg-gray-300 text-gray-700 px-4 py-2 rounded-lg">پاک‌سازی</button>
          </div>
          <div v-if="queryLoading" class="text-blue-600 mb-2">در حال اجرا...</div>
          <div v-if="queryError" class="bg-red-50 border border-red-200 rounded-lg p-6">
            <p class="text-red-700">{{ queryError }}</p>
          </div>
          <div v-if="queryResult && queryResult.length" class="mt-4">
            <div class="bg-white rounded-lg border border-gray-200 p-6 w-full">
              <div class="w-full overflow-x-auto">
                <table class="border-collapse min-w-full w-max">
                  <thead>
                    <tr class="bg-gray-100">
                      <th v-for="col in goodTableColumns" :key="col.field" class="px-4 py-2 border text-right text-sm font-medium text-gray-700">{{ col.label }}</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr v-for="(row, index) in queryResult" :key="index">
                      <td v-for="col in goodTableColumns" :key="col.field" class="px-4 py-2 border text-sm text-gray-800">{{ row[col.field] }}</td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- باتم‌بار -->
    <div v-if="showSettings" class="fixed inset-0 bg-black bg-opacity-40 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg shadow-lg p-8 w-full max-w-2xl relative">
        <button @click="showSettings = false" class="absolute left-4 top-6 text-gray-500 hover:text-red-500 text-xl">×</button>
        <div class="flex gap-2 mb-4 border-b pb-2">
          <button :class="['px-3 py-1 rounded font-bold transition-colors', settingsTab === 'db' ? 'bg-pink-200 text-pink-900' : 'bg-pink-50 text-pink-700 hover:bg-pink-100']" @click="settingsTab = 'db'">تنظیمات دیتابیس</button>
          <button :class="['px-3 py-1 rounded font-bold transition-colors', settingsTab === 'update' ? 'bg-teal-200 text-teal-900' : 'bg-teal-50 text-teal-700 hover:bg-teal-100']" @click="settingsTab = 'update'">تنظیمات بروزرسانی</button>
          <button :class="['px-3 py-1 rounded font-bold transition-colors', settingsTab === 'customQuery' ? 'bg-yellow-200 text-yellow-900' : 'bg-yellow-50 text-yellow-700 hover:bg-yellow-100']" @click="settingsTab = 'customQuery'">اجزای کوئری دلخواه</button>
          <button :class="['px-3 py-1 rounded font-bold transition-colors', settingsTab === 'pool' ? 'bg-indigo-200 text-indigo-900' : 'bg-indigo-50 text-indigo-700 hover:bg-indigo-100']" @click="settingsTab = 'pool'">کانکشن‌پول</button>
        </div>
        <div v-if="settingsTab === 'db'">
          <h2 class="text-lg font-bold mb-4">تنظیمات دیتابیس</h2>
          <div class="space-y-3">
            <div>
              <label class="block text-sm mb-1">نام دیتابیس</label>
              <input v-model="settings.database" class="w-full border rounded px-3 py-2" />
            </div>
            <div>
              <label class="block text-sm mb-1">نام کاربری</label>
              <input v-model="settings.user" class="w-full border rounded px-3 py-2" />
            </div>
            <div>
              <label class="block text-sm mb-1">رمز عبور</label>
              <input v-model="settings.password" type="password" class="w-full border rounded px-3 py-2" />
            </div>
            <div>
              <label class="block text-sm mb-1">آدرس سرور</label>
              <input v-model="settings.host" class="w-full border rounded px-3 py-2" />
            </div>
            <div>
              <label class="block text-sm mb-1">پورت</label>
              <input v-model="settings.port" class="w-full border rounded px-3 py-2" />
            </div>
          </div>
          <div class="flex justify-end gap-2 mt-6">
            <button @click="showSettings = false" class="bg-gray-200 hover:bg-gray-300 text-gray-700 font-bold px-4 py-2 rounded transition-colors">انصراف</button>
            <button @click="applyUpdateSettings" class="bg-green-200 hover:bg-green-300 text-green-900 font-bold px-4 py-2 rounded transition-colors">ذخیره</button>
          </div>
        </div>
        <div v-else-if="settingsTab === 'update'">
          <h2 class="text-lg font-bold mb-4">تنظیمات بروزرسانی</h2>
          <div class="space-y-3">
            <div class="flex items-center gap-2">
              <input type="checkbox" v-model="autoRefresh.enabled" id="auto-refresh-checkbox" />
              <label for="auto-refresh-checkbox" class="text-sm">بروزرسانی خودکار فعال باشد</label>
            </div>
            <div v-if="autoRefresh.enabled">
              <label class="block text-sm mb-1">فاصله بروزرسانی (ثانیه)</label>
              <input v-model.number="autoRefresh.interval" type="number" min="1" class="w-full border rounded px-3 py-2" />
            </div>
          </div>
          <div class="flex justify-end gap-2 mt-6">
            <button @click="showSettings = false" class="bg-gray-200 hover:bg-gray-300 text-gray-700 font-bold px-4 py-2 rounded transition-colors">انصراف</button>
            <button @click="applyUpdateSettings" class="bg-green-200 hover:bg-green-300 text-green-900 font-bold px-4 py-2 rounded transition-colors">ذخیره</button>
          </div>
        </div>
        <div v-else-if="settingsTab === 'customQuery'">
          <h2 class="text-lg font-bold mb-4">اجزای کوئری دلخواه</h2>
          <div class="space-y-3">
            <div class="flex items-center gap-2">
              <input type="checkbox" v-model="customQueryAuto.enabled" id="custom-query-auto-checkbox" />
              <label for="custom-query-auto-checkbox" class="text-sm">اجرای خودکار کوئری فعال باشد</label>
            </div>
            <div v-if="customQueryAuto.enabled">
              <label class="block text-sm mb-1">فاصله اجرا (ثانیه)</label>
              <input v-model.number="customQueryAuto.interval" type="number" min="1" class="w-full border rounded px-3 py-2" />
            </div>
            <div>
              <label class="block text-sm mb-1">کوئری فعلی</label>
              <div class="flex gap-2 items-center">
                <textarea v-model="query" class="w-full border rounded px-3 py-2 bg-gray-100 ltr:text-left" rows="2" dir="ltr"></textarea>
                <button @click="query = 'SELECT * FROM pg_stat_activity LIMIT 5'" class="bg-blue-100 hover:bg-blue-200 text-blue-800 font-bold px-3 py-2 rounded transition-colors">بازنشانی</button>
              </div>
            </div>
            <div class="flex justify-end gap-2 mt-6">
              <button @click="showSettings = false" class="bg-gray-200 hover:bg-gray-300 text-gray-700 font-bold px-4 py-2 rounded transition-colors">انصراف</button>
              <button @click="applyUpdateSettings" class="bg-green-200 hover:bg-green-300 text-green-900 font-bold px-4 py-2 rounded transition-colors">ذخیره</button>
            </div>
          </div>
          <div class="mt-6">
            <div class="flex items-center justify-between mb-2">
              <span class="font-bold">۱۰ تست آخر کوئری دلخواه</span>
              <div class="flex gap-2">
                <button @click="showCustomQueryLogs = !showCustomQueryLogs" class="bg-pink-100 hover:bg-pink-200 text-pink-800 font-bold px-3 py-2 rounded transition-colors text-sm">
                  {{ showCustomQueryLogs ? 'بستن لیست' : 'نمایش لیست' }}
                </button>
                <button v-if="showCustomQueryLogs" @click="fetchCustomQueryLogs" class="bg-yellow-100 hover:bg-yellow-200 text-yellow-800 font-bold px-3 py-2 rounded transition-colors text-sm">بروزرسانی</button>
              </div>
            </div>
            <div v-if="showCustomQueryLogs">
              <div v-if="customQueryLogsLoading" class="text-center py-4">در حال بارگذاری...</div>
              <div v-else class="overflow-x-auto max-w-full">
                <div class="bg-white rounded-lg border border-gray-200 p-6 w-full">
                  <div class="w-full overflow-x-auto">
                    <table class="border-collapse text-xs text-right min-w-full w-max">
                      <thead>
                        <tr class="bg-gray-100">
                          <th class="px-3 py-2 border text-right font-medium text-gray-700">زمان اجرا</th>
                          <th class="px-3 py-2 border text-right font-medium text-gray-700">مدت پاسخ (ms)</th>
                          <th class="px-3 py-2 border text-right font-medium text-gray-700">کوئری</th>
                          <th class="px-3 py-2 border text-right font-medium text-gray-700">نتیجه/خطا</th>
                        </tr>
                      </thead>
                      <tbody>
                        <tr v-for="(log, idx) in customQueryLogs" :key="log.id" :class="idx % 2 === 0 ? 'bg-gray-50' : 'bg-white'">
                          <td class="px-3 py-2 border">{{ formatDateTime(log.executed_at) }}</td>
                          <td class="px-3 py-2 border">{{ log.response_time_ms }}</td>
                          <td class="px-3 py-2 border ltr text-left" dir="ltr">{{ log.query }}</td>
                          <td class="px-3 py-2 border">{{ shortResult(log) }}</td>
                        </tr>
                        <tr v-if="!customQueryLogs.length">
                          <td colspan="4" class="text-center text-gray-400 py-4 border">لاگی وجود ندارد</td>
                        </tr>
                      </tbody>
                    </table>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div v-else-if="settingsTab === 'pool'">
          <h2 class="text-lg font-bold mb-4">تنظیمات کانکشن‌پول دیتابیس</h2>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm mb-1">تعداد هسته‌های اختصاص‌داده‌شده (Allocated Cores)</label>
              <input v-model.number="pool.allocatedCores" type="number" min="1" class="w-full border rounded px-3 py-2" />
              <p class="text-xs text-gray-500 mt-1">هرچه بیشتر باشد، می‌توان کانکشن‌های بیکار و باز بیشتری داشت؛ اما فشار روی DB هم بیشتر می‌شود.</p>
            </div>
            <div>
              <label class="block text-sm mb-1">حداکثر کانکشن‌های بیکار (Max Idle Conns)</label>
              <input v-model.number="pool.maxIdleConns" type="number" min="0" class="w-full border rounded px-3 py-2" />
              <p class="text-xs text-gray-500 mt-1">کانکشن‌های آماده برای استفاده مجدد. پیشنهاد: برابر با تعداد هسته‌ها.</p>
            </div>
            <div>
              <label class="block text-sm mb-1">حداکثر کانکشن‌های باز (Max Open Conns)</label>
              <input v-model.number="pool.maxOpenConns" type="number" min="0" class="w-full border rounded px-3 py-2" />
              <p class="text-xs text-gray-500 mt-1">حداکثر اتصالات هم‌زمان. پیشنهاد: ۲ × تعداد هسته‌ها.</p>
            </div>
            <div>
              <label class="block text-sm mb-1">حداکثر عمر اتصال (دقیقه) Conn Max Lifetime</label>
              <input v-model.number="pool.connMaxLifetimeMinutes" type="number" min="1" class="w-full border rounded px-3 py-2" />
              <p class="text-xs text-gray-500 mt-1">اتصالات قدیمی به‌صورت دوره‌ای بسته و جایگزین می‌شوند. پیشنهاد: ≤ 60 دقیقه.</p>
            </div>
            <div>
              <label class="block text-sm mb-1">حداکثر بیکاری اتصال (دقیقه) Conn Max Idle Time</label>
              <input v-model.number="pool.connMaxIdleTimeMinutes" type="number" min="0" class="w-full border rounded px-3 py-2" />
              <p class="text-xs text-gray-500 mt-1">اگر اتصال این مدت بیکار بماند، بسته می‌شود. پیشنهاد: ~30 دقیقه.</p>
            </div>
          </div>
          <div class="mt-4 p-3 bg-gray-50 rounded border text-sm">
            <div class="font-bold mb-2">پیشنهاد سیستم</div>
            <ul class="list-disc pr-6 space-y-1">
              <li>Max Idle Conns: {{ recommended.maxIdleConns }} (≈ تعداد هسته‌ها)</li>
              <li>Max Open Conns: {{ recommended.maxOpenConns }} (≈ ۲× هسته‌ها)</li>
              <li>Conn Max Lifetime: {{ recommended.connMaxLifetimeMinutes }} دقیقه</li>
              <li>Conn Max Idle Time: {{ recommended.connMaxIdleTimeMinutes }} دقیقه</li>
            </ul>
            <div class="mt-2 text-gray-600">هسته‌های سیستم: {{ systemInfo.cpuCores }}</div>
          </div>
          <div class="flex justify-end gap-2 mt-6">
            <button @click="loadPool()" class="bg-gray-200 hover:bg-gray-300 text-gray-700 font-bold px-4 py-2 rounded transition-colors">بازیابی</button>
            <button @click="savePool()" class="bg-indigo-600 hover:bg-indigo-700 text-white font-bold px-4 py-2 rounded transition-colors">اعمال تنظیمات</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, watch, computed } from 'vue';
import dayjs from 'dayjs/esm/index.js';

definePageMeta({
  layout: 'admin-main',
  middleware: ['developer-only']
});

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user, hasPermission } = useAuth()

const connectionStatus = ref('checking');
const dbInfo = ref({
  database: 'mydb',
  user: 'postgres',
  host: 'localhost',
  port: '9090'
});
const error = ref(null);
const isTesting = ref(false);
const showSettings = ref(false);
const settingsTab = ref('db');
const query = ref('SELECT * FROM pg_stat_activity LIMIT 5');
const queryResult = ref(null);
const queryError = ref(null);
const queryTime = ref(null);
const queryRows = ref(null);
const isQueryRunning = ref(false);
const systemStats = ref({
  responseTime: '-',
  databaseSize: '-',
  activeConnections: '-'
});
const activeQueries = ref([]);
const slowQueries = ref([]);
const dbUptime = ref('-');
let pollingInterval = null;

const autoRefresh = ref({
  enabled: true,
  interval: 1
});

const customQueryAuto = ref({
  enabled: false,
  interval: 10,
});
let customQueryInterval = null;

const settings = ref({
  database: '',
  user: '',
  password: '',
  host: '',
  port: ''
});

const customQueryLogs = ref([]);
const customQueryLogsLoading = ref(false);

const autoRefreshSettings = ref({
  enabled: true,
  interval: 1
});

const customQueryAutoSettings = ref({
  enabled: false,
  interval: 10,
});

const showCustomQueryLogs = ref(false);
const activeQueriesExpanded = ref(false);
const slowQueriesExpanded = ref(false);
const customQueryExpanded = ref(false);
const queryLoading = ref(false);

// Pool settings state
const pool = ref({
  allocatedCores: 0,
  maxIdleConns: 0,
  maxOpenConns: 0,
  connMaxLifetimeMinutes: 60,
  connMaxIdleTimeMinutes: 30,
})
const recommended = ref({
  maxIdleConns: 0,
  maxOpenConns: 0,
  connMaxLifetimeMinutes: 60,
  connMaxIdleTimeMinutes: 30,
})
const systemInfo = ref({ cpuCores: 0 })

// بررسی نقش developer
const isDeveloper = ref(false);

// احراز هویت غیرفعال شده است - کاربر developer در نظر گرفته می‌شود
onMounted(async () => {
  isDeveloper.value = true;
});

const goodTableColumns = computed(() => {
  if (!queryResult.value || !queryResult.value.length) return [];
  return Object.keys(queryResult.value[0]).map(key => ({
    label: key,
    field: key,
    tdClass: key === 'query'
      ? 'align-top whitespace-pre-line break-words min-w-[120px] max-w-[320px]'
      : 'align-top whitespace-nowrap',
    thClass: 'text-center'
  }));
});

async function loadAutoRefreshSettings() {
  try {
    const res = await fetch('/api/admin/system/auto-refresh-settings');
    if (res.ok) {
      const data = await res.json();
      if (typeof data === 'object' && data !== null) {
        autoRefresh.value = Object.assign(autoRefresh.value, data);
      }
    }
  } catch {}
}

async function saveAutoRefreshSettings() {
  await fetch('/api/admin/system/auto-refresh-settings', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(autoRefresh.value)
  });
}

async function loadCustomQueryAutoSettings() {
  try {
    const res = await fetch('/api/admin/system/custom-query-auto-settings');
    if (res.ok) {
      const data = await res.json();
      if (typeof data === 'object' && data !== null) {
        customQueryAuto.value = Object.assign(customQueryAuto.value, data);
      }
    }
  } catch {}
}

async function saveCustomQueryAutoSettings() {
  await fetch('/api/admin/system/custom-query-auto-settings', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(customQueryAuto.value)
  });
}

async function fetchDbMonitor() {
  try {
    const response = await fetch('/api/admin/system/db-monitor');
    if (!response.ok) {
      throw new Error(`Server responded with status: ${response.status}`);
    }
    const data = await response.json();
    activeQueries.value = data.activeQueries || [];
    slowQueries.value = data.slowQueries || [];
  } catch (e) {
    activeQueries.value = [];
    slowQueries.value = [];
  }
}

async function testConnection() {
  isTesting.value = true;
  error.value = null;
  try {
    const startTime = Date.now();
    const response = await fetch('/health/stats');
    const responseTime = Date.now() - startTime;
    const result = await response.json();
    if (result && result.status === 'ok') {
      connectionStatus.value = 'connected';
      systemStats.value.responseTime = responseTime;
      systemStats.value.databaseSize = result.data.database_size || '-';
      systemStats.value.activeConnections = result.data.active_connections || '-';
      dbUptime.value = result.data.uptime || '-';
      dbInfo.value = {
        database: 'mydb',
        user: 'postgres',
        host: 'localhost',
        port: '9090'
      };
    } else {
      connectionStatus.value = 'disconnected';
      error.value = result.message || 'خطا در دریافت اطلاعات';
      systemStats.value.responseTime = '-';
      systemStats.value.databaseSize = '-';
      systemStats.value.activeConnections = '-';
      dbUptime.value = '-';
    }
  } catch (err) {
    connectionStatus.value = 'disconnected';
    systemStats.value.responseTime = '-';
    systemStats.value.databaseSize = '-';
    systemStats.value.activeConnections = '-';
    dbUptime.value = '-';
    error.value = err.message;
  } finally {
    isTesting.value = false;
  }
}

function applyUpdateSettings() {
  showSettings.value = false;
  saveAutoRefreshSettings();
  saveCustomQueryAutoSettings();
  setupAutoRefresh();
}

function setupAutoRefresh() {
  if (pollingInterval) clearInterval(pollingInterval);
  if (autoRefresh.value.enabled) {
    pollingInterval = setInterval(() => {
      testConnection();
      fetchDbMonitor();
    }, autoRefresh.value.interval * 1000);
  }
}

function setupCustomQueryAuto() {
  if (customQueryInterval) clearInterval(customQueryInterval);
  if (customQueryAuto.value.enabled) {
    customQueryInterval = setInterval(() => {
      runQuery();
    }, customQueryAuto.value.interval * 1000);
  }
}

watch(() => customQueryAuto.value.enabled, setupCustomQueryAuto);
watch(() => customQueryAuto.value.interval, setupCustomQueryAuto);

watch(() => autoRefresh.value.enabled, setupAutoRefresh);
watch(() => autoRefresh.value.interval, setupAutoRefresh);

onMounted(() => {
  loadAutoRefreshSettings();
  loadCustomQueryAutoSettings();
  testConnection();
  fetchDbMonitor();
  setupAutoRefresh();
  setupCustomQueryAuto();
  fetchCustomQueryLogs();
  loadPool();
});

onUnmounted(() => {
  if (pollingInterval) clearInterval(pollingInterval);
  if (customQueryInterval) clearInterval(customQueryInterval);
});

watch(showSettings, (val) => {
  if (val && dbInfo.value) {
    settingsTab.value = 'db';
    settings.value.database = dbInfo.value.database || '';
    settings.value.user = dbInfo.value.user || '';
    settings.value.password = '';
    settings.value.host = dbInfo.value.host || '';
    settings.value.port = dbInfo.value.port || '';
  }
});

async function runQuery() {
  queryResult.value = null;
  queryError.value = null;
  queryTime.value = null;
  queryRows.value = null;
  queryLoading.value = true;
  const start = performance.now();
  try {
    const config = useRuntimeConfig()
    const apiBase = config.public.goApiBase
    if (!apiBase) {
      throw new Error('NUXT_PUBLIC_GO_API_BASE is not configured in .env file')
    }
    const res = await fetch(`${apiBase}/health/query`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ query: query.value })
    });
    const elapsed = performance.now() - start;
    queryTime.value = elapsed;
    if (!res.ok) {
      const err = await res.text();
      throw new Error(err);
    }
    const data = await res.json();
    queryResult.value = data.result;
    if (data && data.result && Array.isArray(data.result)) {
      queryRows.value = data.result.length;
    }
  } catch (e) {
    queryError.value = e.message || 'خطا در اجرای کوئری';
  } finally {
    queryLoading.value = false;
  }
}

async function saveSettings() {
  try {
    const dbUrl = `postgres://${settings.value.user}:${settings.value.password}@${settings.value.host}:${settings.value.port}/${settings.value.database}?sslmode=disable`;
    await fetch('/api/save-settings', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ databaseUrl: dbUrl })
    });
    showSettings.value = false;
    await testConnection();
  } catch (err) {
    error.value = err.message;
  }
}

function formatUptimeStr(str) {
  if (!str || str === '-') return '-';
  return str;
}

function formatDbSize(size) {
  if (size === undefined || size === null || isNaN(Number(size))) return '-';
  return Number(size).toFixed(2) + ' MB';
}

async function resetDatabase() {
  if (confirm('آیا از ریست دیتابیس مطمئن هستید؟')) {
    await fetch('/api/admin/system/reset-database', { method: 'POST' });
    alert('دیتابیس ریست شد!');
    await testConnection();
  }
}

async function fetchCustomQueryLogs() {
  customQueryLogsLoading.value = true;
  try {
    const res = await fetch('/api/admin/system/custom-query-logs');
    if (res.ok) {
      customQueryLogs.value = await res.json();
    } else {
      customQueryLogs.value = [];
    }
  } catch {
    customQueryLogs.value = [];
  } finally {
    customQueryLogsLoading.value = false;
  }
}

watch(settingsTab, (tab) => {
  if (tab === 'customQuery') fetchCustomQueryLogs();
});

function shortResult(log) {
  if (log.error) return '❌ ' + log.error;
  try {
    const parsed = JSON.parse(log.result);
    if (Array.isArray(parsed) && parsed.length > 0) {
      return '✅ ' + JSON.stringify(parsed[0]).slice(0, 80) + (JSON.stringify(parsed[0]).length > 80 ? '...' : '');
    } else if (Array.isArray(parsed)) {
      return '✅ نتیجه خالی';
    }
    return '✅ ' + JSON.stringify(parsed).slice(0, 80);
  } catch {
    return '✅';
  }
}

function formatDateTime(dt) {
  if (!dt) return '-';
  const d = new Date(dt);
  return d.toLocaleString('fa-IR');
}

function formatStartTime(query) {
  // Prefer query_start from backend if available
  if (query.query_start) {
    const d = typeof query.query_start === 'string' ? new Date(query.query_start) : query.query_start;
    if (!isNaN(d.getTime())) {
      // Format as Gregorian, English digits
      return dayjs(d).format('HH:mm:ss - YYYY/MM/DD');
    }
  }
  // Fallback: calculate from duration
  if (query.duration) {
    const now = new Date();
    const parts = query.duration.split(':');
    if (parts.length < 3) return '-';
    const hours = parseInt(parts[0], 10);
    const minutes = parseInt(parts[1], 10);
    const seconds = parseFloat(parts[2]);
    const totalMs = ((hours * 60 + minutes) * 60 + seconds) * 1000;
    const start = new Date(now.getTime() - totalMs);
    return dayjs(start).format('HH:mm:ss - YYYY/MM/DD');
  }
  return '-';
}

async function loadPool() {
  try {
    const res = await $fetch('/api/admin/system/db-pool')
    if (res && res.success) {
      const data = res.data || {}
      pool.value = Object.assign(pool.value, data.current || {})
      recommended.value = Object.assign(recommended.value, data.recommended || {})
      systemInfo.value = Object.assign(systemInfo.value, data.system || {})
      if (!pool.value.allocatedCores && systemInfo.value.cpuCores) {
        pool.value.allocatedCores = systemInfo.value.cpuCores
      }
    }
  } catch {}
}

async function savePool() {
  try {
    await $fetch('/api/admin/system/db-pool', { method: 'POST', body: pool.value })
    alert('تنظیمات اعمال شد')
  } catch (e) {
    alert('خطا در ذخیره تنظیمات: ' + (e?.message || e))
  }
}
</script>

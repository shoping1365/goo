<template>
  <div class="min-h-screen bg-gray-50 p-6">
    <div class="max-w-7xl mx-auto">
      <!-- Header -->
      <div class="mb-8">
        <div class="flex items-center justify-between">
          <div>
            <h1 class="text-3xl font-bold text-gray-900 mb-2">مدیریت ربات‌ها</h1>
            <p class="text-gray-600">نظارت و کنترل ربات‌ها و خزنده‌های وب</p>
          </div>
          <div class="flex items-center gap-6 space-x-3 space-x-reverse">
            <!-- Toggle: Bots enabled (master switch) -->
            <label class="flex items-center cursor-pointer select-none">
              <span class="ml-3 text-sm text-gray-600">فعال بودن سپر</span>
              <input type="checkbox" class="sr-only" :checked="botsEnabled" @change="toggleBotsEnabled">
              <div :class="['w-12 h-7 rounded-full p-1 transition', botsEnabled ? 'bg-green-500' : 'bg-gray-300']">
                <div :class="['w-5 h-5 bg-white rounded-full transition', botsEnabled ? 'translate-x-5' : 'translate-x-0']"></div>
              </div>
            </label>
            <!-- Toggle: Bot Shield allow_all/restricted -->
            <label class="flex items-center cursor-pointer select-none">
              <span class="ml-3 text-sm text-gray-600">سپر ربات</span>
              <input type="checkbox" class="sr-only" :checked="isBotShieldOn" @change="toggleBotShield">
              <div :class="['w-12 h-7 rounded-full p-1 transition', isBotShieldOn ? 'bg-green-500' : 'bg-gray-300']">
                <div :class="['w-5 h-5 bg-white rounded-full transition', isBotShieldOn ? 'translate-x-5' : 'translate-x-0']"></div>
              </div>
            </label>
            <button @click="refreshData" class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors">
              <svg class="w-5 h-5 inline ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
              </svg>
              بروزرسانی
            </button>
            <button @click="showAddBotModal = true" class="bg-green-600 text-white px-4 py-2 rounded-lg hover:bg-green-700 transition-colors">
              <svg class="w-5 h-5 inline ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
              </svg>
              افزودن ربات
            </button>
            <button @click="blockAllBots" class="bg-red-600 text-white px-4 py-2 rounded-lg hover:bg-red-700 transition-colors">
              <svg class="w-5 h-5 inline ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728L5.636 5.636m12.728 12.728L18.364 5.636M5.636 18.364l12.728-12.728"></path>
              </svg>
              مسدود کردن همه
            </button>
            <button @click="whitelistAllBots" class="bg-green-600 text-white px-4 py-2 rounded-lg hover:bg-green-700 transition-colors">
              <svg class="w-5 h-5 inline ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
              وایت‌لیست همه
            </button>
          </div>
        </div>
      </div>

      <!-- Stats Cards -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-3 bg-green-100 rounded-lg">
              <svg class="w-8 h-8 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
            <div class="mr-4">
              <p class="text-sm font-medium text-gray-600">ربات‌های مجاز</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.allowed }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-3 bg-red-100 rounded-lg">
              <svg class="w-8 h-8 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728L5.636 5.636m12.728 12.728L18.364 5.636M5.636 18.364l12.728-12.728"></path>
              </svg>
            </div>
            <div class="mr-4">
              <p class="text-sm font-medium text-gray-600">ربات‌های مسدود</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.blocked }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-3 bg-blue-100 rounded-lg">
              <svg class="w-8 h-8 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
              </svg>
            </div>
            <div class="mr-4">
              <p class="text-sm font-medium text-gray-600">درخواست‌های امروز</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.requestsToday }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-3 bg-yellow-100 rounded-lg">
              <svg class="w-8 h-8 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
            <div class="mr-4">
              <p class="text-sm font-medium text-gray-600">آخرین فعالیت</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.lastActivity }}</p>
            </div>
          </div>
        </div>
      </div>

      <!-- BotShield Settings -->
      <div class="bg-white rounded-lg shadow p-6 mb-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات سپر ربات</h3>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div class="space-y-3">
            <label class="flex items-center cursor-pointer select-none">
              <span class="ml-3 text-sm text-gray-600">فقط مانیتور (بدون مسدودسازی)</span>
              <input type="checkbox" class="sr-only" :checked="monitorOnly" @change="toggleMonitorOnly">
              <div :class="['w-12 h-7 rounded-full p-1 transition', monitorOnly ? 'bg-green-500' : 'bg-gray-300']">
                <div :class="['w-5 h-5 bg-white rounded-full transition', monitorOnly ? 'translate-x-5' : 'translate-x-0']"></div>
              </div>
            </label>

            <label class="flex items-center cursor-pointer select-none">
              <span class="ml-3 text-sm text-gray-600">عبور درخواست‌های دارای Authorization</span>
              <input type="checkbox" class="sr-only" :checked="skipAuthorized" @change="toggleSkipAuthorized">
              <div :class="['w-12 h-7 rounded-full p-1 transition', skipAuthorized ? 'bg-green-500' : 'bg-gray-300']">
                <div :class="['w-5 h-5 bg-white rounded-full transition', skipAuthorized ? 'translate-x-5' : 'translate-x-0']"></div>
              </div>
            </label>

            <label class="flex items-center cursor-pointer select-none">
              <span class="ml-3 text-sm text-gray-600">عبور IP های محلی (127.0.0.1, ::1, 192.168.x)</span>
              <input type="checkbox" class="sr-only" :checked="skipLocalIPs" @change="toggleSkipLocalIPs">
              <div :class="['w-12 h-7 rounded-full p-1 transition', skipLocalIPs ? 'bg-green-500' : 'bg-gray-300']">
                <div :class="['w-5 h-5 bg-white rounded-full transition', skipLocalIPs ? 'translate-x-5' : 'translate-x-0']"></div>
              </div>
            </label>
          </div>

          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">متدهای عبوری (CSV)</label>
              <div class="flex gap-2">
                <input v-model="skipMethodsCsv" type="text" class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="مثلاً: OPTIONS,HEAD"/>
                <button @click="saveSkipMethods" class="px-3 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">ذخیره</button>
              </div>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">مسیرهای استثناء (CSV از prefix)</label>
              <div class="flex gap-2">
                <input v-model="excludePathsCsv" type="text" class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="مثلاً: /_nuxt,/.nuxt,/assets,/public"/>
                <button @click="saveExcludePaths" class="px-3 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">ذخیره</button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Live Activity Feed -->
      <div class="bg-white rounded-lg shadow p-6 mb-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">فعالیت زنده ربات‌ها</h3>
          <div class="flex items-center space-x-2 space-x-reverse">
            <span class="text-sm text-gray-600">وضعیت:</span>
            <span class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium bg-green-100 text-green-800">
              <span class="w-2 h-2 bg-green-400 rounded-full ml-1"></span>
              آنلاین
            </span>
          </div>
        </div>
        <div class="space-y-3 max-h-64 overflow-y-auto">
          <div v-for="activity in liveActivities" :key="activity.id" class="flex items-center p-3 bg-gray-50 rounded-lg">
            <div class="flex-shrink-0">
              <div class="w-8 h-8 rounded-full bg-blue-100 flex items-center justify-center">
                <svg class="w-4 h-4 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.75 17L9 20l-1 1h8l-1-1-.75-3M3 13h18M5 17h14a2 2 0 002-2V5a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"></path>
                </svg>
              </div>
            </div>
            <div class="mr-3 flex-1">
              <p class="text-sm font-medium text-gray-900">{{ activity.botName }}</p>
              <p class="text-xs text-gray-500">{{ activity.action }} - {{ activity.page }}</p>
            </div>
            <div class="text-right">
              <p class="text-xs text-gray-500">{{ formatTime(activity.timestamp) }}</p>
              <span :class="getActivityStatusClass(activity.status)" class="inline-flex px-2 py-1 rounded-full text-xs font-medium">
                {{ getActivityStatusLabel(activity.status) }}
              </span>
            </div>
          </div>
        </div>
      </div>

      <!-- Filters -->
      <div class="bg-white rounded-lg shadow p-6 mb-6">
        <div class="flex flex-wrap gap-6">
          <div class="flex-1 min-w-64">
            <label class="block text-sm font-medium text-gray-700 mb-2">جستجو</label>
            <input
              v-model="searchQuery"
              type="text"
              placeholder="جستجو بر اساس نام یا User Agent..."
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
          </div>
          <div class="w-48">
            <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
            <select
              v-model="statusFilter"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="all">همه</option>
              <option value="allowed">مجاز</option>
              <option value="blocked">مسدود</option>
              <option value="monitoring">تحت نظارت</option>
            </select>
          </div>
          <div class="w-48">
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع</label>
            <select
              v-model="typeFilter"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="all">همه</option>
              <option value="search_engine">موتور جستجو</option>
              <option value="crawler">خزنده</option>
              <option value="general">ربات عمومی</option>
              <option value="malicious">بدخواه</option>
            </select>
          </div>
        </div>
      </div>

      <!-- Bots Table -->
      <div class="bg-white rounded-lg shadow overflow-hidden">
        <div class="px-6 py-4 border-b border-gray-200">
          <h3 class="text-lg font-semibold text-gray-900">لیست ربات‌ها</h3>
        </div>
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نام ربات</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">User Agent</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">آخرین بازدید</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="bot in filteredBots" :key="bot.id">
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="flex items-center">
                    <div class="">
                      <div class="h-10 w-10 rounded-full bg-gray-100 flex items-center justify-center">
                        <svg class="h-6 w-6 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.75 17L9 20l-1 1h8l-1-1-.75-3M3 13h18M5 17h14a2 2 0 002-2V5a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"></path>
                        </svg>
                      </div>
                    </div>
                    <div class="mr-4">
                      <div class="text-sm font-medium text-gray-900">{{ bot.name }}</div>
                      <div class="text-sm text-gray-500">{{ bot.ip }}</div>
                    </div>
                  </div>
                </td>
                <td class="px-6 py-4">
                  <div class="text-sm text-gray-900 max-w-xs truncate">{{ bot.userAgent }}</div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span :class="getTypeClass(bot.type)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                    {{ getTypeLabel(bot.type) }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span :class="getStatusClass(bot.status)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                    {{ getStatusLabel(bot.status) }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {{ formatDate(bot.lastVisit) }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                  <div class="flex space-x-2 space-x-reverse">
                    <button
                      v-if="bot.status === 'blocked'"
                      @click="allowBot(bot.id)"
                      class="text-green-600 hover:text-green-900"
                    >
                      آزادسازی
                    </button>
                    <button
                      v-else
                      @click="blockBot(bot.id)"
                      class="text-red-600 hover:text-red-900"
                    >
                      مسدود کردن
                    </button>
                    <button
                      @click="showWhitelistModal = true"
                      class="text-blue-600 hover:text-blue-900"
                    >
                      وایت‌لیست
                    </button>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>

    <!-- Add Bot Modal -->
    <div v-if="showAddBotModal" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
      <div class="relative top-20 mx-auto p-5 border w-full max-w-md sm:max-w-lg md:max-w-xl shadow-lg rounded-md bg-white">
        <div class="mt-3">
          <h3 class="text-lg font-medium text-gray-900 mb-4">افزودن ربات جدید</h3>
          <form @submit.prevent="addBot">
            <div class="mb-4">
              <label class="block text-sm font-medium text-gray-700 mb-2">نام ربات</label>
              <input
                v-model="newBot.name"
                type="text"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
            </div>
            <div class="mb-4">
              <label class="block text-sm font-medium text-gray-700 mb-2">User Agent</label>
              <input
                v-model="newBot.userAgent"
                type="text"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
            </div>
            <div class="mb-4">
              <label class="block text-sm font-medium text-gray-700 mb-2">نوع</label>
              <select
                v-model="newBot.type"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="search_engine">موتور جستجو</option>
                <option value="crawler">خزنده</option>
                <option value="general">ربات عمومی</option>
                <option value="malicious">بدخواه</option>
              </select>
            </div>
            <div class="flex justify-end space-x-3 space-x-reverse">
              <button
                type="button"
                @click="showAddBotModal = false"
                class="px-4 py-2 text-gray-700 bg-gray-200 rounded-lg hover:bg-gray-300"
              >
                انصراف
              </button>
              <button
                type="submit"
                class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700"
              >
                افزودن
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>

    <!-- Whitelist Modal -->
    <div v-if="showWhitelistModal" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
      <div class="relative top-20 mx-auto p-5 border w-full max-w-md sm:max-w-lg md:max-w-xl shadow-lg rounded-md bg-white">
        <div class="mt-3">
          <h3 class="text-lg font-medium text-gray-900 mb-4">افزودن به لیست سفید</h3>
          <form @submit.prevent="addToWhitelist">
            <div class="mb-4">
              <label class="block text-sm font-medium text-gray-700 mb-2">نوع شناسایی</label>
              <select
                v-model="whitelistForm.type"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="name">نام ربات</option>
                <option value="userAgent">User Agent</option>
                <option value="ip">آدرس IP</option>
                <option value="pattern">الگوی User Agent</option>
              </select>
            </div>
            <div class="mb-4">
              <label class="block text-sm font-medium text-gray-700 mb-2">مقدار</label>
              <input
                v-model="whitelistForm.value"
                type="text"
                required
                :placeholder="getWhitelistPlaceholder()"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
            </div>
            <div class="mb-4">
              <label class="block text-sm font-medium text-gray-700 mb-2">توضیحات (اختیاری)</label>
              <textarea
                v-model="whitelistForm.description"
                rows="3"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              ></textarea>
            </div>
            <div class="flex justify-end space-x-3 space-x-reverse">
              <button
                type="button"
                @click="showWhitelistModal = false"
                class="px-4 py-2 text-gray-700 bg-gray-200 rounded-lg hover:bg-gray-300"
              >
                انصراف
              </button>
              <button
                type="submit"
                class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700"
              >
                افزودن به لیست سفید
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>

    <!-- Block All Confirmation Modal -->
    <div v-if="showBlockAllModal" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
      <div class="relative top-20 mx-auto p-5 border w-full max-w-md sm:max-w-lg md:max-w-xl shadow-lg rounded-md bg-white">
        <div class="mt-3">
          <div class="mx-auto flex items-center justify-center h-12 w-12 rounded-full bg-red-100 mb-4">
            <svg class="h-6 w-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
            </svg>
          </div>
          <h3 class="text-lg font-medium text-gray-900 mb-2">تأیید مسدود کردن همه ربات‌ها</h3>
          <p class="text-sm text-gray-500 mb-4">
            آیا مطمئن هستید که می‌خواهید تمام ربات‌ها را مسدود کنید؟ این عمل شامل ربات‌های موتورهای جستجو نیز می‌شود.
          </p>
          <div class="flex justify-end space-x-3 space-x-reverse">
            <button
              @click="showBlockAllModal = false"
              class="px-4 py-2 text-gray-700 bg-gray-200 rounded-lg hover:bg-gray-300"
            >
              انصراف
            </button>
            <button
              @click="confirmBlockAll"
              class="px-4 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700"
            >
              مسدود کردن همه
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
declare const useFetch: <T = unknown>(url: string, options?: { key?: string; default?: () => T; method?: string; body?: unknown; query?: Record<string, string | number> }) => Promise<{ data: { value: T }; refresh: () => Promise<void> }>
declare const $fetch: <T = unknown>(url: string, options?: { method?: string; body?: unknown; query?: Record<string, string | number> }) => Promise<T>
</script>

<script setup lang="ts">
import { computed, onMounted, ref, watchEffect } from 'vue';

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// داده از API تنظیمات ربات‌ها (اتصال واقعی)
const { data: botConfig, refresh } = await useFetch('/api/admin/bots/config', { key: 'bots-config', default: () => ({ data: {} }) })
// master enable state
const botsEnabled = computed(() => {
  const raw = (botConfig.value?.data?.['bots.enabled'] || 'true').toString().toLowerCase()
  return !(raw === 'false' || raw === '0' || raw === 'no')
})
const toggleBotsEnabled = async () => {
  const next = botsEnabled.value ? 'false' : 'true'
  await $fetch('/api/admin/bots/config', { method: 'PUT', body: { 'bots.enabled': next } })
  await refresh()
}

// BotShield new controls
const monitorOnly = computed(() => {
  const raw = (botConfig.value?.data?.['bots.monitor_only'] || 'true').toString().toLowerCase()
  return !(raw === 'false' || raw === '0' || raw === 'no')
})
const skipAuthorized = computed(() => {
  const raw = (botConfig.value?.data?.['bots.skip_authorized'] || 'true').toString().toLowerCase()
  return !(raw === 'false' || raw === '0' || raw === 'no')
})
const skipLocalIPs = computed(() => {
  const raw = (botConfig.value?.data?.['bots.skip_local_ips'] || 'true').toString().toLowerCase()
  return !(raw === 'false' || raw === '0' || raw === 'no')
})

const skipMethodsCsv = ref('')
const excludePathsCsv = ref('')

watchEffect(() => {
  skipMethodsCsv.value = (botConfig.value?.data?.['bots.skip_methods'] || 'OPTIONS,HEAD').toString()
  excludePathsCsv.value = (botConfig.value?.data?.['bots.exclude_paths'] || '/_nuxt,/.nuxt,/assets,/public,/uploads,/static,/favicon.ico,/robots.txt,/health,/admin,/api/admin').toString()
})

const toggleMonitorOnly = async () => {
  const next = monitorOnly.value ? 'false' : 'true'
  await $fetch('/api/admin/bots/config', { method: 'PUT', body: { 'bots.monitor_only': next } })
  await refresh()
}
const toggleSkipAuthorized = async () => {
  const next = skipAuthorized.value ? 'false' : 'true'
  await $fetch('/api/admin/bots/config', { method: 'PUT', body: { 'bots.skip_authorized': next } })
  await refresh()
}
const toggleSkipLocalIPs = async () => {
  const next = skipLocalIPs.value ? 'false' : 'true'
  await $fetch('/api/admin/bots/config', { method: 'PUT', body: { 'bots.skip_local_ips': next } })
  await refresh()
}
const saveSkipMethods = async () => {
  await $fetch('/api/admin/bots/config', { method: 'PUT', body: { 'bots.skip_methods': skipMethodsCsv.value } })
  await refresh()
}
const saveExcludePaths = async () => {
  await $fetch('/api/admin/bots/config', { method: 'PUT', body: { 'bots.exclude_paths': excludePathsCsv.value } })
  await refresh()
}
// وضعیت سپر ربات (on = restricted, off = allow_all)
const isBotShieldOn = computed(() => {
  const mode = (botConfig.value?.data?.['bots.global_mode'] || '').toString().toLowerCase()
  return mode === 'restricted' || mode === 'block_all'
})

const toggleBotShield = async () => {
  const next = isBotShieldOn.value ? 'allow_all' : 'restricted'
  await $fetch('/api/admin/bots/config', { method: 'PUT', body: { 'bots.global_mode': next } })
  await refresh()
}
// تعریف متغیرهای reactive
const showAddBotModal = ref(false)
const showWhitelistModal = ref(false)
const showBlockAllModal = ref(false)
const searchQuery = ref('')
const statusFilter = ref('all')
const typeFilter = ref('all')

// آمار ربات‌ها
const stats = ref({
  allowed: 15,
  blocked: 8,
  requestsToday: 1247,
  lastActivity: '2 دقیقه پیش'
})

// فرم وایت‌لیست
const whitelistForm = ref({
  type: 'name',
  value: '',
  description: ''
})

// فرم افزودن ربات جدید
const newBot = ref({
  name: '',
  userAgent: '',
  type: 'general'
})

// فعالیت‌های زنده (نمونه)
const liveActivities = ref([
  {
    id: 1,
    botName: 'Googlebot',
    action: 'بازدید صفحه',
    page: '/product/123',
    status: 'allowed',
    timestamp: new Date(Date.now() - 30000)
  },
  {
    id: 2,
    botName: 'Bingbot',
    action: 'تلاش دسترسی',
    page: '/admin/login',
    status: 'blocked',
    timestamp: new Date(Date.now() - 60000)
  },
  {
    id: 3,
    botName: 'Unknown Bot',
    action: 'درخواست API',
    page: '/api/products',
    status: 'monitoring',
    timestamp: new Date(Date.now() - 90000)
  }
])

// تعریف انواع
type BotAPIResponse = {
  id?: number
  user_agent?: string
  name?: string
  type?: string
  status?: string
  ip?: string
  ip_address?: string
  created_at?: string | Date
  lastVisit?: Date | string
}

type Bot = {
  id: number
  name: string
  userAgent: string
  type: string
  status: string
  ip?: string
  lastVisit?: Date
}

// لیست ربات‌ها (دریافت از API)
const bots = ref<Bot[]>([
  {
    id: 1,
    name: 'Googlebot',
    userAgent: 'Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)',
    type: 'search_engine',
    status: 'allowed',
    ip: '66.249.66.1',
    lastVisit: new Date(Date.now() - 300000)
  },
  {
    id: 2,
    name: 'Bingbot',
    userAgent: 'Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)',
    type: 'search_engine',
    status: 'allowed',
    ip: '157.55.39.1',
    lastVisit: new Date(Date.now() - 600000)
  },
  {
    id: 3,
    name: 'Yandex',
    userAgent: 'Mozilla/5.0 (compatible; YandexBot/3.0; +http://yandex.com/bots)',
    type: 'search_engine',
    status: 'allowed',
    ip: '100.43.80.1',
    lastVisit: new Date(Date.now() - 900000)
  },
  {
    id: 4,
    name: 'Unknown Bot',
    userAgent: 'Mozilla/5.0 (compatible; UnknownBot/1.0)',
    type: 'malicious',
    status: 'blocked',
    ip: '192.168.1.100',
    lastVisit: new Date(Date.now() - 1200000)
  }
])

// محاسبه ربات‌های فیلتر شده
const filteredBots = computed(() => {
  return bots.value.filter(bot => {
    const matchesSearch = !searchQuery.value || 
      bot.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      bot.userAgent.toLowerCase().includes(searchQuery.value.toLowerCase())
    
    const matchesStatus = statusFilter.value === 'all' || bot.status === statusFilter.value
    const matchesType = typeFilter.value === 'all' || bot.type === typeFilter.value
    
    return matchesSearch && matchesStatus && matchesType
  })
})

// توابع
// همگام‌سازی لیست ربات‌ها از API با الگوهای پیش‌فرض
const syncBotsFromAPI = async () => {
  const response = await $fetch<BotAPIResponse[] | { data?: BotAPIResponse[] }>('/api/admin/bots', { query: { patterns: 'bot,crawler,spider' } })
  const data = Array.isArray(response) ? response : (response?.data || [])
  if (Array.isArray(data) && data.length > 0) {
    // map ساده برای تطبیق با جدول
    bots.value = data.map((row: BotAPIResponse) => ({
      id: row.id || 0,
      name: (row.user_agent || row.name || '').split('/')[0] || 'Bot',
      userAgent: row.user_agent || '',
      type: row.type || 'crawler',
      status: row.status || 'monitoring',
      ip: row.ip_address,
      lastVisit: new Date(row.created_at || Date.now())
    }))
  }
}
const refreshData = () => { refresh() }

const addBot = () => {
  const bot = {
    id: Date.now(),
    name: newBot.value.name,
    userAgent: newBot.value.userAgent,
    type: newBot.value.type,
    status: 'allowed',
    ip: '192.168.1.1',
    lastVisit: new Date()
  }
  
  bots.value.push(bot)
  showAddBotModal.value = false
  
  // ریست کردن فرم
  newBot.value = {
    name: '',
    userAgent: '',
    type: 'general'
  }
}

const blockBot = async (id) => {
  const bot = bots.value.find(b => b.id === id)
  if (!bot) return
  // افزودن UA به بلاک لیست
  const key = 'bots.blocklist.user_agents'
  const current = (botConfig.value?.data?.[key] || '').toString()
  const next = current ? `${current},${bot.userAgent}` : bot.userAgent
  await $fetch('/api/admin/bots/config', { method: 'PUT', body: { [key]: next } })
  await refresh()
  bot.status = 'blocked'
}

const allowBot = async (id) => {
  const bot = bots.value.find(b => b.id === id)
  if (!bot) return
  // حذف UA از بلاک لیست (فرانت حذف ساده؛ بک‌اند فعلاً overwrite کلید)
  const key = 'bots.blocklist.user_agents'
  const current = (botConfig.value?.data?.[key] || '').toString()
  const parts = current.split(',').map(s => s.trim()).filter(Boolean)
  const filtered = parts.filter(p => p !== bot.userAgent)
  await $fetch('/api/admin/bots/config', { method: 'PUT', body: { [key]: filtered.join(',') } })
  await refresh()
  bot.status = 'allowed'
}

const blockAllBots = () => { showBlockAllModal.value = true }

const confirmBlockAll = async () => {
  // حالت کلی را روی block_all می‌گذاریم
  await $fetch('/api/admin/bots/config', { method: 'PUT', body: { 'bots.global_mode': 'block_all' } })
  await refresh()
  showBlockAllModal.value = false
}

const whitelistAllBots = async () => {
  // حالت کلی را روی allow_all می‌گذاریم
  await $fetch('/api/admin/bots/config', { method: 'PUT', body: { 'bots.global_mode': 'allow_all' } })
  await refresh()
}

const addToWhitelist = async () => {
  const typeKeyMap = { name: 'bots.allowlist.user_agents', userAgent: 'bots.allowlist.user_agents', ip: 'bots.allowlist.ips', pattern: 'bots.allowlist.user_agents' }
  const key = typeKeyMap[whitelistForm.value.type]
  if (!key) { showWhitelistModal.value = false; return }
  const current = (botConfig.value?.data?.[key] || '').toString()
  const next = current ? `${current},${whitelistForm.value.value}` : whitelistForm.value.value
  await $fetch('/api/admin/bots/config', { method: 'PUT', body: { [key]: next } })
  await refresh()
  showWhitelistModal.value = false
  whitelistForm.value = { type: 'name', value: '', description: '' }
}

const getWhitelistPlaceholder = () => {
  const placeholders = {
    name: 'مثال: Googlebot',
    userAgent: 'مثال: Mozilla/5.0 (compatible; Googlebot/2.1)',
    ip: 'مثال: 66.249.66.1',
    pattern: 'مثال: *Googlebot*'
  }
  return placeholders[whitelistForm.value.type] || ''
}

const formatDate = (date: Date | string) => {
  const now = new Date()
  const dateObj = typeof date === 'string' ? new Date(date) : date
  const diff = now.getTime() - dateObj.getTime()
  const minutes = Math.floor(diff / 60000)
  
  if (minutes < 1) return 'کمتر از یک دقیقه پیش'
  if (minutes < 60) return `${minutes} دقیقه پیش`
  
  const hours = Math.floor(minutes / 60)
  if (hours < 24) return `${hours} ساعت پیش`
  
  const days = Math.floor(hours / 24)
  return `${days} روز پیش`
}

const formatTime = (date: Date | string) => {
  const now = new Date()
  const dateObj = typeof date === 'string' ? new Date(date) : date
  const diff = now.getTime() - dateObj.getTime()
  const minutes = Math.floor(diff / 60000)
  const seconds = Math.floor((diff % 60000) / 1000)

  if (minutes > 0) {
    return `${minutes} دقیقه و ${seconds} ثانیه پیش`
  } else {
    return `${seconds} ثانیه پیش`
  }
}

const getTypeClass = (type) => {
  const classes = {
    search_engine: 'bg-blue-100 text-blue-800',
    crawler: 'bg-green-100 text-green-800',
    general: 'bg-gray-100 text-gray-800',
    malicious: 'bg-red-100 text-red-800'
  }
  return classes[type] || 'bg-gray-100 text-gray-800'
}

const getTypeLabel = (type) => {
  const labels = {
    search_engine: 'موتور جستجو',
    crawler: 'خزنده',
    general: 'ربات عمومی',
    malicious: 'بدخواه'
  }
  return labels[type] || 'نامشخص'
}

const getStatusClass = (status) => {
  const classes = {
    allowed: 'bg-green-100 text-green-800',
    blocked: 'bg-red-100 text-red-800',
    monitoring: 'bg-yellow-100 text-yellow-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getStatusLabel = (status) => {
  const labels = {
    allowed: 'مجاز',
    blocked: 'مسدود',
    monitoring: 'تحت نظارت'
  }
  return labels[status] || 'نامشخص'
}

const getActivityStatusClass = (status) => {
  const classes = {
    allowed: 'bg-green-100 text-green-800',
    blocked: 'bg-red-100 text-red-800',
    monitoring: 'bg-yellow-100 text-yellow-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getActivityStatusLabel = (status) => {
  const labels = {
    allowed: 'مجاز',
    blocked: 'مسدود',
    monitoring: 'تحت نظارت'
  }
  return labels[status] || 'نامشخص'
}

// شبیه‌سازی فعالیت‌های زنده
onMounted(() => {
  setInterval(() => {
    const newActivity = {
      id: Date.now(),
      botName: ['Googlebot', 'Bingbot', 'Yandex', 'Unknown Bot'][Math.floor(Math.random() * 4)],
      action: ['بازدید صفحه', 'تلاش دسترسی', 'درخواست API'][Math.floor(Math.random() * 3)],
      page: ['/', '/product/123', '/blog/article-1', '/admin/login'][Math.floor(Math.random() * 4)],
      status: ['allowed', 'blocked', 'monitoring'][Math.floor(Math.random() * 3)],
      timestamp: new Date()
    }

    liveActivities.value.unshift(newActivity)

    // حذف فعالیت‌های قدیمی (بیش از 50 مورد)
    if (liveActivities.value.length > 50) {
      liveActivities.value = liveActivities.value.slice(0, 50)
    }
  }, 5000) // هر 5 ثانیه یک فعالیت جدید
  // بارگیری اولیه ربات‌ها از API
  syncBotsFromAPI()
})
</script> 

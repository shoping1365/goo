<template>
  <div dir="rtl" class="min-h-screen bg-gray-100">
    <!-- محتوای اصلی -->
    <div class="flex flex-row-reverse gap-6 p-6 h-screen">
      <!-- ستون اول - اطلاعات کاربر (فقط وقتی چت انتخاب شده) -->
      <div v-if="selectedSession" class="order-1 w-80 bg-white rounded-lg shadow-sm border border-gray-200">
        <!-- هدر اطلاعات کاربر -->
        <div class="p-6 border-b border-gray-200">
          <div class="flex items-center justify-between">
            <h3 class="text-lg font-semibold text-gray-900">اطلاعات کاربر</h3>
            <button class="p-2 text-gray-500 hover:bg-gray-100 rounded-lg transition-colors">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728L5.636 5.636m12.728 12.728L18.364 5.636M5.636 18.364l12.728-12.728"></path>
              </svg>
            </button>
          </div>
          <button class="mt-3 w-full px-3 py-2 text-sm text-red-600 border border-red-300 rounded-lg hover:bg-red-50 transition-colors flex items-center justify-center space-x-2 space-x-reverse">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728L5.636 5.636m12.728 12.728L18.364 5.636M5.636 18.364l12.728-12.728"></path>
            </svg>
            <span>انتقال به لیست سیاه</span>
          </button>
        </div>

        <!-- اطلاعات کاربر -->
        <div class="p-6">
          <!-- شناسه کاربر -->
          <div class="flex items-center space-x-3 space-x-reverse mb-6">
            <div class="text-2xl font-bold text-purple-600">{{ selectedSession.customer_id || selectedSession.customer_name || 'کاربر' }}</div>
            <div class="w-8 h-8 bg-purple-500 rounded-full flex items-center justify-center">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"></path>
              </svg>
            </div>
          </div>
          <div class="text-sm text-blue-600 mb-6 truncate">{{ selectedSession.url || 'بدون URL' }}</div>

          <!-- اطلاعات تماس -->
          <div class="space-y-4 mb-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">ایمیل :</label>
              <div class="flex items-center space-x-2 space-x-reverse">
                <input type="email" class="flex-1 px-3 py-2 border border-gray-300 rounded-lg text-sm" placeholder="ایمیل کاربر">
                <button class="p-2 text-gray-500 hover:bg-gray-100 rounded-lg">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"></path>
                  </svg>
                </button>
              </div>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">موبایل :</label>
              <div class="flex items-center space-x-2 space-x-reverse">
                <input type="tel" class="flex-1 px-3 py-2 border border-gray-300 rounded-lg text-sm" placeholder="شماره موبایل">
                <button class="p-2 text-gray-500 hover:bg-gray-100 rounded-lg">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"></path>
                  </svg>
                </button>
              </div>
            </div>
          </div>

          <!-- یادداشت -->
          <div class="mb-6">
            <label class="block text-sm font-medium text-gray-700 mb-2">یادداشت</label>
            <textarea 
              class="w-full px-3 py-2 border-2 border-yellow-300 rounded-lg text-sm resize-none" 
              rows="3" 
              placeholder="بنویسید و اینتر بزنید"
            ></textarea>
          </div>

          <!-- مشخصه‌های کاربر -->
          <div>
            <h4 class="text-sm font-medium text-gray-700 mb-3">مشخصه های کاربر</h4>
            <div class="space-y-2 text-sm">
              <div class="flex justify-between">
                <span class="text-gray-600">آخرین صفحه :</span>
                <span class="text-gray-900">-</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">حضور در سایت :</span>
                <span class="text-gray-900">-</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">آی پی :</span>
                <span class="text-gray-900">37.32.18.9</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">دستگاه :</span>
                <span class="text-gray-900">Chrome-Android</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">مرورگر :</span>
                <span class="text-gray-900">Chrome 138</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">در صفحه :</span>
                <span class="text-gray-900 text-xs truncate">https://iranxia.net/product/redm...</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">ارجاع :</span>
                <span class="text-gray-900 text-xs truncate">https://iranxia.net/product-cate...</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- ستون دوم - چت یا صفحه خوش‌آمدگویی -->
      <div v-if="selectedSession" class="order-2 flex-1 bg-white rounded-lg shadow-sm border border-gray-200 flex flex-col">
        <!-- هدر چت -->
        <div class="p-6 border-b border-gray-200 bg-gray-50">
          <div class="flex items-center justify-between">
            <div class="flex items-center space-x-3 space-x-reverse">
              <span class="text-lg font-semibold text-gray-900">کاربر : {{ selectedSession.customer_id || selectedSession.customer_name || 'کاربر' }}</span>
              <div class="w-8 h-8 bg-purple-500 rounded-full flex items-center justify-center">
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"></path>
                </svg>
              </div>
            </div>
            <div class="flex items-center space-x-2 space-x-reverse">
              <button class="px-3 py-1 text-sm text-gray-600 hover:bg-gray-200 rounded-lg transition-colors flex items-center space-x-1 space-x-reverse" @click="handleTransferSession">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
                </svg>
                <span>انتقال مکالمه</span>
              </button>
              <button class="px-3 py-1 text-sm text-gray-600 hover:bg-gray-200 rounded-lg transition-colors" @click="handleCloseSession">بستن مکالمه</button>
              <button class="p-1 text-red-500 hover:bg-red-50 rounded-lg transition-colors">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                </svg>
              </button>
            </div>
          </div>
        </div>

        <!-- محتوای چت -->
        <div class="flex-1 p-6 overflow-y-auto">
          <!-- تاریخ -->
          <div v-if="messages.length > 0" class="text-center text-sm text-gray-500 mb-4">
            {{ formatDate(messages[0].created_at) }}
          </div>
          
          <!-- پیام‌ها -->
          <div v-for="message in messages" :key="message.id" class="mb-4">
            <div v-if="message.sender_type === 'customer'" class="flex justify-end">
              <div class="max-w-xs bg-blue-500 text-white rounded-lg p-3">
                <div class="text-sm mb-2">{{ message.message }}</div>
                <div class="flex items-center justify-between text-xs opacity-75">
                  <span>{{ formatTime(message.created_at) }}</span>
                </div>
              </div>
            </div>
            <div v-else class="flex justify-start">
              <div class="max-w-xs bg-gray-200 text-gray-800 rounded-lg p-3">
                <div class="text-sm mb-2">{{ message.message }}</div>
                <div class="flex items-center justify-between text-xs opacity-75">
                  <span>{{ formatTime(message.created_at) }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- ورودی پیام -->
        <div class="p-6 border-t border-gray-200">
          <div class="flex items-center space-x-3 space-x-reverse">
            <input 
              id="file-input" 
              type="file" 
              class="hidden" 
              accept="image/*,video/*"
              @change="handleFileSelect"
            >
            <label for="file-input" class="p-2 text-gray-500 hover:bg-gray-100 rounded-lg transition-colors cursor-pointer">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.172 7l-6.586 6.586a2 2 0 102.828 2.828l6.414-6.586a4 4 0 00-5.656-5.656l-6.415 6.585a6 6 0 108.486 8.486L20.5 13"></path>
              </svg>
            </label>
            <input 
              v-model="newMessage"
              type="text" 
              placeholder="پیام خود را بنویسید..." 
              class="flex-1 px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-1 focus:ring-blue-400 focus:border-blue-400"
              @keyup.enter="handleSendMessage"
            >
            <button class="p-2 text-gray-500 hover:bg-gray-100 rounded-lg transition-colors">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11a7 7 0 01-7 7m0 0a7 7 0 01-7-7m7 7v4m0 0H8m4 0h4m-4-8a3 3 0 01-3-3V5a3 3 0 116 0v6a3 3 0 01-3 3z"></path>
              </svg>
            </button>
            <button class="w-10 h-10 bg-purple-600 text-white rounded-full flex items-center justify-center hover:bg-purple-700 transition-colors" @click="handleSendMessage">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8"></path>
              </svg>
            </button>
          </div>
        </div>
      </div>

      <!-- صفحه خوش‌آمدگویی (وقتی هیچ چتی انتخاب نشده) -->
      <div v-else class="order-2 flex-1 flex items-center justify-center">
        <div class="bg-white rounded-lg shadow-sm p-12 max-w-md w-full">
          <!-- ربات چت -->
          <div class="flex justify-center mb-8">
            <div class="w-24 h-24 bg-purple-600 rounded-full flex items-center justify-center relative">
              <!-- آنتن‌ها -->
              <div class="absolute -top-2 left-1/2 transform -translate-x-1/2">
                <div class="w-1 h-3 bg-white rounded-full"></div>
              </div>
              <div class="absolute -top-2 right-1/2 transform translate-x-1/2">
                <div class="w-1 h-3 bg-white rounded-full"></div>
              </div>
              <!-- صورت ربات -->
              <div class="w-16 h-16 bg-white rounded-full flex flex-col items-center justify-center">
                <!-- چشم‌ها -->
                <div class="flex space-x-2 space-x-reverse mb-1">
                  <div class="w-2 h-2 bg-purple-600 rounded-full"></div>
                  <div class="w-2 h-2 bg-purple-600 rounded-full"></div>
                </div>
                <!-- لبخند -->
                <div class="w-4 h-1 bg-purple-600 rounded-full"></div>
              </div>
            </div>
          </div>
          
          <!-- متن خوش‌آمدگویی -->
          <h1 class="text-2xl font-bold text-gray-900 text-center mb-4">خوش آمدید</h1>
          <p class="text-gray-600 text-center mb-8">شما به وب سایتهای زیر دسترسی دارید</p>
          
          <!-- فیلد URL -->
          <div class="bg-gray-100 border border-gray-300 rounded-lg px-4 py-3 flex items-center">
            <svg class="w-5 h-5 text-gray-500 ml-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1"></path>
            </svg>
            <span class="text-gray-700 text-sm">https://iranxia.net</span>
          </div>
        </div>
      </div>

      <!-- ستون سوم - سایدبار لیست چت‌ها -->
      <div class="order-3 w-full md:w-80 lg:w-96 ml-auto bg-white rounded-lg shadow-sm border border-gray-200">
        <!-- هدر بنفش سایدبار -->
        <div class="bg-purple-600 text-white p-6">
          <div class="flex items-center justify-between">
            <div class="flex items-center space-x-3 space-x-reverse">
              <div class="w-8 h-8 bg-white rounded-full flex items-center justify-center">
                <svg class="w-4 h-4 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"></path>
                </svg>
              </div>
              <div class="w-8 h-8 bg-white rounded-full flex items-center justify-center">
                <svg class="w-4 h-4 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 4.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"></path>
                </svg>
              </div>
            </div>
            
            <div class="flex items-center space-x-3 space-x-reverse">
              <div class="text-right">
                <div class="text-sm font-medium">آ. آذری</div>
                <div class="flex items-center space-x-2 space-x-reverse text-xs">
                  <span>آنلاین</span>
                  <div class="w-8 h-4 bg-gray-300 rounded-full relative">
                    <div class="w-3 h-3 bg-green-500 rounded-full absolute right-0.5 top-0.5"></div>
                  </div>
                </div>
              </div>
              <div class="w-10 h-10 bg-orange-400 rounded-lg flex items-center justify-center">
                <div class="w-6 h-6 bg-white rounded-full flex items-center justify-center">
                  <div class="w-4 h-4 bg-red-500 rounded-full"></div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- جستجو -->
        <div class="p-6 border-b border-gray-200">
          <div class="relative">
            <input
              v-model="searchQuery"
              type="text"
              placeholder="جست و جو در بین مکالمات"
              class="w-full px-4 py-2 pr-10 text-sm border border-gray-300 rounded-lg bg-gray-50 focus:bg-white focus:ring-1 focus:ring-blue-400 focus:border-blue-400"
            />
            <svg class="absolute left-3 top-2.5 w-4 h-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
            </svg>
          </div>
        </div>

        <!-- تب‌ها -->
        <div class="flex bg-white">
          <button
:class="[
                    'flex-1 px-3 py-3 text-center text-sm transition-colors',
                    currentTab === 'closed' ? 'bg-gray-100 text-gray-900' : 'text-gray-600 hover:bg-gray-50'
                  ]" 
                  @click="handleTabChange('closed')">
            مکالمه‌های دیگر
          </button>
          <button
:class="[
                    'flex-1 px-3 py-3 text-center text-sm transition-colors relative',
                    currentTab === 'waiting' ? 'bg-yellow-50 text-yellow-600 border-b-2 border-yellow-500' : 'text-gray-600 hover:bg-gray-50'
                  ]" 
                  @click="handleTabChange('waiting')">
            <span>در انتظار پاسخ</span>
            <div v-if="waitingSessionsCount > 0" class="absolute -top-1 -right-1 w-5 h-5 bg-white border border-yellow-500 rounded-full flex items-center justify-center">
              <span class="text-xs text-yellow-600 font-medium">{{ waitingSessionsCount }}</span>
            </div>
          </button>
          <button
:class="[
                    'flex-1 px-3 py-3 text-center text-sm transition-colors relative',
                    currentTab === 'active' ? 'bg-red-50 text-red-600 border-b-2 border-red-500' : 'text-gray-600 hover:bg-gray-50'
                  ]" 
                  @click="handleTabChange('active')">
            <span>در حال مکالمه</span>
            <div v-if="activeSessionsCount > 0" class="absolute -top-1 -right-1 w-5 h-5 bg-white border border-red-500 rounded-full flex items-center justify-center">
              <span class="text-xs text-red-600 font-medium">{{ activeSessionsCount }}</span>
            </div>
          </button>
        </div>

        <!-- فیلتر -->
        <div class="px-4 py-2 flex items-center justify-between text-xs text-gray-600 border-b border-gray-100">
          <div class="flex items-center space-x-1 space-x-reverse">
            <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
            </svg>
            <span>باقی مانده : {{ filteredSessions.length }}</span>
          </div>
          <div class="flex items-center space-x-1 space-x-reverse">
            <span>جدید ترین</span>
            <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
            </svg>
          </div>
        </div>

        <!-- Loading state -->
        <div v-if="isLoading" class="p-6 text-center text-gray-500">
          در حال بارگذاری...
        </div>

        <!-- Error state -->
        <div v-else-if="error" class="p-6 text-center text-red-500">
          {{ error }}
        </div>

        <!-- لیست چت‌ها -->
        <div v-else class="overflow-y-auto" style="height: calc(100% - 280px);">
          <div
v-for="session in filteredSessions" :key="session.id" 
               :class="[
                 'p-3 border-b border-gray-100 hover:bg-gray-50 cursor-pointer transition-colors',
                 selectedSession?.id === session.id ? 'bg-blue-50 border-blue-200' : ''
               ]"
               @click="handleSelectSession(session)">
            <div class="flex items-center justify-between mb-2">
              <div class="flex items-center space-x-2 space-x-reverse">
                <span class="text-sm text-gray-600">کاربر :</span>
                <span class="text-sm font-medium text-gray-900">{{ session.customer_id || session.customer_name || 'کاربر' }}</span>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <div v-if="session.unread_count > 0" class="w-5 h-5 bg-black text-white text-xs rounded-full flex items-center justify-center font-medium">
                  {{ session.unread_count }}
                </div>
                <div class="w-8 h-8 bg-purple-500 rounded-full flex items-center justify-center">
                  <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"></path>
                  </svg>
                </div>
              </div>
            </div>
            <div class="text-xs text-blue-600 mb-2 truncate">{{ session.url || 'بدون URL' }}</div>
            <div class="text-sm text-gray-800 pr-2">{{ session.last_message || 'بدون پیام' }}</div>
          </div>
        </div>

        <!-- دکمه بیشتر -->
        <div class="p-6 border-t border-gray-200 bg-white">
          <button class="w-full px-4 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200 transition-colors text-sm">
            بیشتر
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import type { Ref } from 'vue';

interface ChatSession {
  id: string | number;
  customer_id?: string;
  customer_name?: string;
  customer_phone?: string;
  url?: string;
  status: 'active' | 'waiting' | 'closed';
  last_message?: string;
  unread_count?: number;
  created_at?: string;
  updated_at?: string;
  [key: string]: unknown;
}

interface Message {
  id: string | number;
  session_id: string | number;
  sender_type: 'customer' | 'operator' | 'system';
  message: string;
  created_at: string;
  is_read?: boolean;
  [key: string]: unknown;
}

declare const definePageMeta: (meta: { layout?: string }) => void
declare const useLiveChat: () => {
  chatSessions: Ref<ChatSession[]>;
  selectedSession: Ref<ChatSession | null>;
  messages: Ref<Message[]>;
  isLoading: Ref<boolean>;
  error: Ref<string | null>;
  fetchChatSessions: (status?: string) => Promise<void>;
  fetchMessages: (sessionId: string | number) => Promise<void>;
  sendMessage: (sessionId: string | number, message: string, file?: File | null) => Promise<void>;
  updateSessionStatus: (sessionId: string | number, status: string) => Promise<void>;
  startPolling: () => void;
  stopPolling: () => void;
  selectSession: (session: ChatSession) => Promise<void>;
  markAsRead: (sessionId: string | number) => Promise<void>;
}
</script>

<script setup lang="ts">
import { computed, onMounted, onUnmounted, ref } from 'vue';

definePageMeta({
  layout: 'admin-main'
})

// Use the live chat composable
const {
  chatSessions,
  selectedSession,
  messages,
  isLoading,
  error,
  fetchChatSessions,
  fetchMessages: _fetchMessages,
  sendMessage,
  updateSessionStatus,
  // connectWebSocket,
  // disconnectWebSocket,
  // enableRealtime,
  startPolling,
  stopPolling,
  selectSession,
  markAsRead,
  // getRealTimeMetrics
} = useLiveChat()

const realTimeMetrics = ref(null)
const getRealTimeMetrics = async () => null

// Local state
const currentTab = ref('active') // active, waiting, closed
const searchQuery = ref('')
const newMessage = ref('')
const selectedFile = ref<File | null>(null)
// const realTimeMetrics = ref(null)

// Computed
const filteredSessions = computed(() => {
  let filtered = chatSessions.value.filter(session => {
    if (currentTab.value === 'active') return session.status === 'active'
    if (currentTab.value === 'waiting') return session.status === 'waiting'
    if (currentTab.value === 'closed') return session.status === 'closed'
    return true
  })
  
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    filtered = filtered.filter(session => 
      session.customer_name?.toLowerCase().includes(query) ||
      session.customer_phone?.includes(query) ||
      (session.customer_id && String(session.customer_id).includes(query)) ||
      session.last_message?.toLowerCase().includes(query)
    )
  }
  
  return filtered
})

const activeSessionsCount = computed(() => 
  chatSessions.value.filter(s => s.status === 'active').length
)

const waitingSessionsCount = computed(() => 
  chatSessions.value.filter(s => s.status === 'waiting').length
)

const _closedSessionsCount = computed(() => 
  chatSessions.value.filter(s => s.status === 'closed').length
)

// Methods
const handleSelectSession = async (session: ChatSession) => {
  await selectSession(session)
  await markAsRead(session.id)
}

const handleSendMessage = async () => {
  if (!newMessage.value.trim() && !selectedFile.value) return
  
  if (selectedSession.value) {
    try {
      await sendMessage(selectedSession.value.id, newMessage.value, selectedFile.value)
      newMessage.value = ''
      selectedFile.value = null
    } catch (err) {
      // console.error('Error sending message:', err)
    }
  }
}

const handleFileSelect = (event: Event) => {
  const target = event.target as HTMLInputElement
  if (target.files && target.files[0]) {
    selectedFile.value = target.files[0]
  }
}

const handleTabChange = (tab: string) => {
  currentTab.value = tab
  fetchChatSessions(tab === 'all' ? undefined : tab)
}

const handleCloseSession = async () => {
  if (!selectedSession.value) return
  
  try {
    await updateSessionStatus(selectedSession.value.id, 'closed')
    selectedSession.value = null
  } catch (err) {
    // console.error('Error closing session:', err)
  }
}

const handleTransferSession = async () => {
  // TODO: Implement session transfer functionality
  // console.log('Transfer session functionality to be implemented')
}

const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleDateString('fa-IR', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}

const formatTime = (dateString: string) => {
  return new Date(dateString).toLocaleTimeString('fa-IR', {
    hour: '2-digit',
    minute: '2-digit'
  })
}

// Auto-refresh metrics
let metricsInterval: NodeJS.Timeout | null = null

const startMetricsRefresh = () => {
  metricsInterval = setInterval(async () => {
    const metrics = await getRealTimeMetrics()
    if (metrics) {
      realTimeMetrics.value = metrics
    }
  }, 30000) // Refresh every 30 seconds
}

const stopMetricsRefresh = () => {
  if (metricsInterval) {
    clearInterval(metricsInterval)
    metricsInterval = null
  }
}

// Lifecycle
onMounted(async () => {
  await fetchChatSessions('active')
  // موقت: فقط Polling فعال است و وب‌سوکت غیرفعال می‌باشد
  startPolling()
  startMetricsRefresh()
})

onUnmounted(() => {
  stopPolling()
  stopMetricsRefresh()
})
</script>

<style scoped>
/* Custom scrollbar */
::-webkit-scrollbar {
  width: 6px;
}

::-webkit-scrollbar-track {
  background: #f1f1f1;
}

::-webkit-scrollbar-thumb {
  background: #c1c1c1;
  border-radius: 3px;
}

::-webkit-scrollbar-thumb:hover {
  background: #a8a8a8;
}
</style> 

<template>
  <div class="min-h-screen bg-gray-50">
    <!-- محتوای اصلی -->
    <div>
      <!-- هدر صفحه -->
      <header class="bg-white shadow-sm border-b border-gray-200">
        <div class="px-6 py-4">
          <div class="flex items-center justify-between">
            <div>
              <h1 class="text-2xl font-bold text-gray-900">مدیریت انبارها</h1>
              <p class="text-gray-600">سیستم جامع مدیریت موجودی و انبارداری</p>
            </div>
            <div class="flex items-center gap-6">
              <!-- انتخاب انبار -->
              <div class="relative">
                <select class="appearance-none bg-white border border-gray-300 rounded-lg px-4 py-2 pr-10 text-gray-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent">
                  <option>انبار مرکزی تهران</option>
                  <option>انبار اصفهان</option>
                  <option>انبار شیراز</option>
                  <option>انبار مشهد</option>
                </select>
                <svg class="absolute left-3 top-1/2 transform -translate-y-1/2 w-4 h-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                </svg>
              </div>
              
              <!-- فیلتر زمان -->
              <div class="relative">
                <select class="appearance-none bg-white border border-gray-300 rounded-lg px-4 py-2 pr-10 text-gray-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent">
                  <option>امروز</option>
                  <option>هفته جاری</option>
                  <option>ماه جاری</option>
                  <option>سال جاری</option>
                </select>
                <svg class="absolute left-3 top-1/2 transform -translate-y-1/2 w-4 h-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                </svg>
              </div>
              
              <!-- جستجو -->
              <div class="relative">
                <input type="text" placeholder="جستجو..." class="w-64 bg-white border border-gray-300 rounded-lg pl-10 pr-4 py-2 text-gray-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent">
                <svg class="absolute left-3 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                </svg>
              </div>
              
              <!-- اعلان‌ها -->
              <button class="relative p-2 text-gray-600 hover:text-gray-900 hover:bg-gray-100 rounded-lg transition-colors">
                <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-5 5v-5z" />
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 7h6m0 10v-3m-3 3h.01M9 17h.01M9 14h.01M12 14h.01M15 11h.01M12 11h.01M9 11h.01M7 21h10a2 2 0 002-2V5a2 2 0 00-2-2H7a2 2 0 00-2 2v14a2 2 0 002 2z" />
                </svg>
                <span class="absolute top-0 right-0 w-3 h-3 bg-red-500 rounded-full"></span>
              </button>
              
              <!-- پروفایل کاربر -->
              <button class="flex items-center gap-2 p-2 text-gray-600 hover:text-gray-900 hover:bg-gray-100 rounded-lg transition-colors">
                <div class="w-8 h-8 bg-gradient-to-r from-green-500 to-blue-600 rounded-full flex items-center justify-center">
                  <span class="text-white text-sm font-medium">AD</span>
                </div>
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                </svg>
              </button>
            </div>
          </div>
        </div>
      </header>

      <!-- محتوای اصلی -->
      <main class="px-6 py-6">
        <!-- تب‌های اصلی -->
        <div class="mb-6">
          <div class="border-b border-gray-200">
            <nav class="-mb-px flex space-x-8 space-x-reverse">
              <button 
                @click="activeMainTab = 'dashboard'"
                :class="[
                  'py-2 px-1 border-b-2 font-medium text-sm transition-colors',
                  activeMainTab === 'dashboard'
                    ? 'border-blue-500 text-blue-600'
                    : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
                ]"
              >
                داشبورد
              </button>
              <button 
                @click="activeMainTab = 'warehouses'"
                :class="[
                  'py-2 px-1 border-b-2 font-medium text-sm transition-colors',
                  activeMainTab === 'warehouses'
                    ? 'border-blue-500 text-blue-600'
                    : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
                ]"
              >
                انبارها
              </button>
            </nav>
          </div>
        </div>

        <!-- محتوای تب‌ها -->
        <div v-show="activeMainTab === 'dashboard'">
          <Dashboard />
        </div>
        
        <div v-show="activeMainTab === 'warehouses'">
          <WarehouseList />
        </div>
      </main>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string }) => void
</script>

<script setup lang="ts">
import { ref } from 'vue'
import Dashboard from './components/Dashboard.vue'
import WarehouseList from './components/WarehouseList.vue'

// تعریف layout صفحه
definePageMeta({
  layout: 'admin-main'
})

// مدیریت تب فعال
const activeMainTab = ref('dashboard')
</script>

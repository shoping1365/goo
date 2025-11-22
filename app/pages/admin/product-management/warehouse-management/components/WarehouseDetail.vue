<template>
  <div class="space-y-6">
    <!-- هدر اطلاعات کلی انبار -->
    <div class="bg-gradient-to-br from-blue-50 via-indigo-50 to-purple-50 rounded-3xl shadow-xl p-8 border border-blue-100">
      <div class="flex flex-col lg:flex-row lg:items-start lg:justify-between gap-6">
        <!-- اطلاعات اصلی -->
        <div class="space-y-4">
          <div class="flex items-center gap-6">
            <div class="w-16 h-16 bg-gradient-to-br from-blue-600 to-indigo-700 rounded-3xl flex items-center justify-center shadow-lg">
              <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
              </svg>
            </div>
            <div>
              <h1 class="text-4xl font-bold bg-gradient-to-r from-blue-600 to-indigo-700 bg-clip-text text-transparent">{{ warehouse.name }}</h1>
              <div class="flex items-center gap-6 mt-2">
                <span class="inline-flex items-center px-4 py-2 rounded-2xl text-sm font-bold bg-gradient-to-r from-blue-100 to-indigo-100 text-blue-800 border border-blue-200 shadow-sm">
                  {{ warehouse.code }}
                </span>
                <span 
                  :class="[
                    'inline-flex items-center px-4 py-2 rounded-2xl text-sm font-bold',
                    warehouse.status === 'فعال' 
                      ? 'bg-green-100 text-green-800 border border-green-200' 
                      : 'bg-red-100 text-red-800 border border-red-200'
                  ]"
                >
                  <span 
                    :class="[
                      'w-3 h-3 rounded-full mr-2',
                      warehouse.status === 'فعال' ? 'bg-green-400' : 'bg-red-400'
                    ]"
                  ></span>
                  {{ warehouse.status }}
                </span>
              </div>
            </div>
          </div>
          
          <!-- آدرس و موقعیت -->
          <div class="space-y-3">
            <div class="flex items-start gap-3">
              <svg class="w-5 h-5 text-blue-500 mt-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
              </svg>
              <div>
                <div class="text-lg font-bold text-gray-900">{{ warehouse.city }}</div>
                <div class="text-gray-600">{{ warehouse.address }}</div>
                <div class="text-sm text-gray-500 mt-1">مختصات: {{ warehouse.coordinates || '35.6892° N, 51.3890° E' }}</div>
              </div>
            </div>
          </div>
        </div>
        
        <!-- آمار و ظرفیت -->
        <div class="space-y-4">
          <div class="grid grid-cols-2 gap-6">
            <div class="bg-white/60 backdrop-blur-sm rounded-2xl p-6 shadow-lg border border-white/20">
              <div class="text-2xl font-bold text-gray-900">{{ warehouse.capacity.used }}</div>
              <div class="text-sm text-gray-600">ظرفیت استفاده شده</div>
            </div>
            <div class="bg-white/60 backdrop-blur-sm rounded-2xl p-6 shadow-lg border border-white/20">
              <div class="text-2xl font-bold text-gray-900">{{ warehouse.capacity.total }}</div>
              <div class="text-sm text-gray-600">ظرفیت کل</div>
            </div>
          </div>
          
          <!-- نوار پیشرفت ظرفیت -->
          <div class="w-64">
            <div class="flex items-center justify-between text-sm mb-2">
              <span class="text-gray-600">{{ Math.round((warehouse.capacity.used / warehouse.capacity.total) * 100) }}% پر شده</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-3 shadow-inner">
              <div 
                class="bg-gradient-to-r from-blue-500 to-indigo-600 h-3 rounded-full transition-all duration-500 shadow-lg" 
                :style="{ width: `${(warehouse.capacity.used / warehouse.capacity.total) * 100}%` }"
              ></div>
            </div>
          </div>
        </div>
      </div>
      
      <!-- مدیر انبار -->
      <div class="mt-6 pt-6 border-t border-blue-200">
        <div class="flex items-center gap-6">
          <div class="w-12 h-12 bg-gradient-to-br from-green-500 to-blue-600 rounded-2xl flex items-center justify-center shadow-lg">
            <span class="text-white text-lg font-bold">{{ warehouse.manager.initials }}</span>
          </div>
          <div>
            <div class="text-lg font-bold text-gray-900">مدیر انبار: {{ warehouse.manager.name }}</div>
            <div class="text-gray-600">{{ warehouse.manager.phone }}</div>
            <div class="text-sm text-gray-500">{{ warehouse.manager.email || 'manager@warehouse.com' }}</div>
          </div>
        </div>
      </div>
    </div>

    <!-- تب‌های اصلی -->
    <div class="bg-white/80 backdrop-blur-sm rounded-3xl shadow-2xl border border-white/20">
      <div class="border-b border-gray-200">
        <nav class="flex space-x-8 space-x-reverse px-8">
          <button 
            v-for="tab in tabs" 
            :key="tab.id"
            :class="[
              'py-4 px-2 border-b-2 font-medium text-sm transition-colors',
              activeTab === tab.id
                ? 'border-blue-500 text-blue-600'
                : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
            ]"
            @click="activeTab = tab.id"
          >
            <div class="flex items-center gap-2">
              <component :is="tab.icon" class="w-5 h-5" />
              {{ tab.name }}
            </div>
          </button>
        </nav>
      </div>
      
      <!-- محتوای تب‌ها -->
      <div class="p-8">
        <!-- تب موجودی کالاها -->
        <div v-show="activeTab === 'inventory'">
          <InventoryTab :warehouse="warehouse" />
        </div>
        
        <!-- تب ورودی کالا -->
        <div v-show="activeTab === 'inbound'">
          <InboundTab :warehouse="warehouse" />
        </div>
        
        <!-- تب خروجی کالا -->
        <div v-show="activeTab === 'outbound'">
          <OutboundTab :warehouse="warehouse" />
        </div>
        
        <!-- تب تاریخچه موجودی -->
        <div v-show="activeTab === 'history'">
          <HistoryTab :warehouse="warehouse" />
        </div>
        
        <!-- تب گزارش مغایرت -->
        <div v-show="activeTab === 'discrepancy'">
          <DiscrepancyTab :warehouse="warehouse" />
        </div>
        
        <!-- تب تنظیمات انبار -->
        <div v-show="activeTab === 'settings'">
          <SettingsTab :warehouse="warehouse" />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import InboundTab from './InboundTab.vue'

import InventoryTab from './InventoryTab.vue'

// تعریف تب‌ها
const tabs = [
  { id: 'inventory', name: 'موجودی کالاها', icon: 'InventoryIcon' },
  { id: 'inbound', name: 'ورودی کالا', icon: 'InboundIcon' },
  { id: 'outbound', name: 'خروجی کالا', icon: 'OutboundIcon' },
  { id: 'history', name: 'تاریخچه موجودی', icon: 'HistoryIcon' },
  { id: 'discrepancy', name: 'گزارش مغایرت', icon: 'DiscrepancyIcon' },
  { id: 'settings', name: 'تنظیمات انبار', icon: 'SettingsIcon' }
]

// تب فعال
const activeTab = ref('inventory')

// داده‌های نمونه انبار
const warehouse = ref({
  id: 1,
  name: 'انبار مرکزی تهران',
  code: 'WH-THR-001',
  type: 'انبار اصلی',
  city: 'تهران',
  address: 'خیابان ولیعصر، پلاک ۱۲۳',
  coordinates: '35.6892° N, 51.3890° E',
  manager: { 
    name: 'علی احمدی', 
    phone: '۰۹۱۲۳۴۵۶۷۸۹', 
    email: 'ali.ahmadi@company.com',
    initials: 'عا' 
  },
  skuCount: 1247,
  capacity: { used: 850, total: 1200 },
  status: 'فعال'
})

// آیکون‌های تب‌ها
// const InventoryIcon = () => h('svg', { class: 'w-5 h-5', fill: 'none', stroke: 'currentColor', viewBox: '0 0 24 24' }, [
//   h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', 'stroke-width': '2', d: 'M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4' })
// ])

// const InboundIcon = () => h('svg', { class: 'w-5 h-5', fill: 'none', stroke: 'currentColor', viewBox: '0 0 24 24' }, [
//   h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', 'stroke-width': '2', d: 'M7 16V4m0 0L3 8m4-4l4 4m6 0v12m0 0l4-4m-4 4l-4-4' })
// ])

// const OutboundIcon = () => h('svg', { class: 'w-5 h-5', fill: 'none', stroke: 'currentColor', viewBox: '0 0 24 24' }, [
//   h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', 'stroke-width': '2', d: 'M17 8V4m0 0l-4 4m4-4l4 4m-4 12v-4m0 0l4-4m-4 4l-4-4' })
// ])

// const HistoryIcon = () => h('svg', { class: 'w-5 h-5', fill: 'none', stroke: 'currentColor', viewBox: '0 0 24 24' }, [
//   h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', 'stroke-width': '2', d: 'M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z' })
// ])

// const DiscrepancyIcon = () => h('svg', { class: 'w-5 h-5', fill: 'none', stroke: 'currentColor', viewBox: '0 0 24 24' }, [
//   h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', 'stroke-width': '2', d: 'M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z' })
// ])

// const SettingsIcon = () => h('svg', { class: 'w-5 h-5', fill: 'none', stroke: 'currentColor', viewBox: '0 0 24 24' }, [
//   h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', 'stroke-width': '2', d: 'M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z' }),
//   h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', 'stroke-width': '2', d: 'M15 12a3 3 0 11-6 0 3 3 0 016 0z' })
// ])
</script>

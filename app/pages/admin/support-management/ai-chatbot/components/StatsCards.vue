<template>
  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
    <!-- کارت کل بات‌ها -->
    <div class="bg-white rounded-lg shadow-md p-6 border border-gray-200">
      <div class="flex items-center justify-between">
        <div>
          <p class="text-sm font-medium text-gray-600">کل بات‌ها</p>
          <p class="text-2xl font-bold text-gray-900">{{ totalBots }}</p>
        </div>
        <div class="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center">
          <span class="text-blue-600 text-xl">🤖</span>
        </div>
      </div>
    </div>

    <!-- کارت بات‌های فعال -->
    <div class="bg-white rounded-lg shadow-md p-6 border border-gray-200">
      <div class="flex items-center justify-between">
        <div>
          <p class="text-sm font-medium text-gray-600">بات‌های فعال</p>
          <p class="text-2xl font-bold text-green-600">{{ activeBots }}</p>
        </div>
        <div class="w-12 h-12 bg-green-100 rounded-lg flex items-center justify-center">
          <span class="text-green-600 text-xl">✅</span>
        </div>
      </div>
    </div>

    <!-- کارت کل مکالمات -->
    <div class="bg-white rounded-lg shadow-md p-6 border border-gray-200">
      <div class="flex items-center justify-between">
        <div>
          <p class="text-sm font-medium text-gray-600">کل مکالمات</p>
          <p class="text-2xl font-bold text-gray-900">{{ totalConversations }}</p>
        </div>
        <div class="w-12 h-12 bg-purple-100 rounded-lg flex items-center justify-center">
          <span class="text-purple-600 text-xl">💬</span>
        </div>
      </div>
    </div>

    <!-- کارت رضایت کاربران -->
    <div class="bg-white rounded-lg shadow-md p-6 border border-gray-200">
      <div class="flex items-center justify-between">
        <div>
          <p class="text-sm font-medium text-gray-600">رضایت کاربران</p>
          <p class="text-2xl font-bold text-orange-600">{{ satisfactionRate }}%</p>
        </div>
        <div class="w-12 h-12 bg-orange-100 rounded-lg flex items-center justify-center">
          <span class="text-orange-600 text-xl">⭐</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
// Props
const props = defineProps({
  totalBots: {
    type: Number,
    required: true
  },
  activeBots: {
    type: Number,
    required: true
  },
  totalConversations: {
    type: Number,
    required: true
  },
  satisfactionRate: {
    type: Number,
    required: true
  }
})
</script> 

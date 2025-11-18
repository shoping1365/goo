<template>
  <div class="bg-white rounded-lg shadow-md border border-gray-200">
    <div class="p-6 border-b border-gray-200">
      <h2 class="text-xl font-bold text-gray-900">لیست چت بات‌ها</h2>
    </div>
    
    <div class="overflow-x-auto">
      <table class="w-full">
        <thead>
          <tr class="border-b border-gray-200 bg-gray-50">
            <th class="text-right py-3 px-4 font-medium text-gray-900">نام بات</th>
            <th class="text-right py-3 px-4 font-medium text-gray-900">نوع</th>
            <th class="text-right py-3 px-4 font-medium text-gray-900">وضعیت</th>
            <th class="text-right py-3 px-4 font-medium text-gray-900">مکالمات امروز</th>
            <th class="text-right py-3 px-4 font-medium text-gray-900">دقت پاسخ</th>
            <th class="text-right py-3 px-4 font-medium text-gray-900">آخرین به‌روزرسانی</th>
            <th class="text-right py-3 px-4 font-medium text-gray-900">عملیات</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="bot in chatbots" :key="bot.id" class="border-b border-gray-100 hover:bg-gray-50">
            <td class="py-4 px-4">
              <div class="flex items-center space-x-3 space-x-reverse">
                <div class="w-10 h-10 bg-blue-100 rounded-full flex items-center justify-center">
                  <span class="text-blue-600 text-lg">{{ bot.icon }}</span>
                </div>
                <div>
                  <p class="font-medium text-gray-900">{{ bot.name }}</p>
                  <p class="text-sm text-gray-600">{{ bot.description }}</p>
                </div>
              </div>
            </td>
            <td class="py-4 px-4">
              <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium" 
                    :class="getTypeClass(bot.type)">
                {{ bot.type }}
              </span>
            </td>
            <td class="py-4 px-4">
              <div class="flex items-center space-x-2 space-x-reverse">
                <div class="relative">
                  <select v-model="bot.status" class="appearance-none bg-transparent border-none text-sm focus:outline-none">
                    <option value="active">فعال</option>
                    <option value="inactive">غیرفعال</option>
                    <option value="training">در حال آموزش</option>
                  </select>
                  <div class="absolute inset-y-0 left-0 flex items-center pointer-events-none">
                    <span class="text-gray-400">▼</span>
                  </div>
                </div>
                <div v-if="bot.status === 'active'" class="w-2 h-2 bg-green-500 rounded-full"></div>
                <div v-else-if="bot.status === 'training'" class="w-2 h-2 bg-yellow-500 rounded-full animate-pulse"></div>
              </div>
            </td>
            <td class="py-4 px-4 text-gray-600">{{ bot.todayConversations }}</td>
            <td class="py-4 px-4">
              <div class="flex items-center space-x-2 space-x-reverse">
                <div class="w-16 bg-gray-200 rounded-full h-2">
                  <div class="bg-green-500 h-2 rounded-full" :style="{ width: bot.accuracy + '%' }"></div>
                </div>
                <span class="text-sm text-gray-600">{{ bot.accuracy }}%</span>
              </div>
            </td>
            <td class="py-4 px-4 text-gray-600 text-sm">{{ bot.lastUpdate }}</td>
            <td class="py-4 px-4">
              <div class="flex items-center space-x-2 space-x-reverse">
                <button @click="$emit('edit', bot)" class="text-gray-500 hover:text-blue-600" title="ویرایش">
                  <span class="text-lg">✏️</span>
                </button>
                <button @click="$emit('test', bot)" class="text-gray-500 hover:text-green-600" title="تست">
                  <span class="text-lg">🧪</span>
                </button>
                <button @click="$emit('analytics', bot)" class="text-gray-500 hover:text-purple-600" title="آمار">
                  <span class="text-lg">📊</span>
                </button>
                <button @click="$emit('delete', bot.id)" class="text-gray-500 hover:text-red-600" title="حذف">
                  <span class="text-lg">🗑️</span>
                </button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- پیام خالی -->
    <div v-if="chatbots.length === 0" class="text-center py-12">
      <div class="text-gray-400 text-6xl mb-4">🤖</div>
      <p class="text-gray-500 font-medium">هیچ چت باتی تعریف نشده است</p>
      <p class="text-gray-400 text-sm mt-2">برای شروع، یک چت بات جدید ایجاد کنید</p>
    </div>
  </div>
</template>

<script setup>
// Props
const props = defineProps({
  chatbots: {
    type: Array,
    required: true
  }
})

// Emits
const emit = defineEmits(['edit', 'test', 'analytics', 'delete'])

// توابع
function getTypeClass(type) {
  const classes = {
    'پشتیبانی عمومی': 'bg-blue-100 text-blue-800',
    'فروش و مشاوره': 'bg-green-100 text-green-800',
    'سوالات متداول': 'bg-purple-100 text-purple-800',
    'رزرو و سفارش': 'bg-orange-100 text-orange-800',
    'سفارشی': 'bg-gray-100 text-gray-800'
  }
  return classes[type] || 'bg-gray-100 text-gray-800'
}
</script> 

<template>
  <div class="bg-white rounded-lg shadow p-6">
    <h2 class="text-xl font-semibold text-gray-900 mb-6">مدیریت محتوا</h2>
    <div class="space-y-6">
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">عنوان صفحه ارجاع</label>
        <input v-model="content.title" type="text" class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500">
      </div>
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">توضیحات برنامه ارجاع</label>
        <textarea v-model="content.description" rows="3" class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"></textarea>
      </div>
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">قوانین و شرایط</label>
        <textarea v-model="content.rules" rows="4" class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"></textarea>
      </div>
      <div class="flex justify-end">
        <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors" @click="saveContent">
          ذخیره محتوا
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const content = ref({
  title: 'برنامه ارجاع مشتریان',
  description: 'با دعوت دوستان خود، پاداش دریافت کنید!',
  rules: 'شرایط و قوانین برنامه ارجاع را اینجا وارد کنید.'
})

function saveContent() {
  // TODO: فراخوانی API برای ذخیره محتوا
  alert('محتوا با موفقیت ذخیره شد!')
}
</script> 

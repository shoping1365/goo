<template>
  <div class="bg-white rounded-lg shadow p-6">
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-xl font-semibold text-gray-900">مدیریت محتوا و راهنما</h2>
      <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors">
        ذخیره تغییرات
      </button>
    </div>
    <div class="space-y-6">
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">متن معرفی برنامه وفاداری</label>
        <textarea v-model="content.intro" rows="3" class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"></textarea>
      </div>
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">قوانین و مقررات</label>
        <textarea v-model="content.rules" rows="4" class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"></textarea>
      </div>
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">سوالات متداول</label>
        <textarea v-model="content.faq" rows="4" class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"></textarea>
      </div>
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">راهنمای استفاده</label>
        <textarea v-model="content.guide" rows="4" class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"></textarea>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const content = ref({
  intro: 'برنامه وفاداری ما به شما این امکان را می‌دهد تا با هر خرید امتیاز جمع‌آوری کنید و از مزایای ویژه بهره‌مند شوید.',
  rules: 'امتیازات فقط برای خریدهای ثبت‌شده و پرداخت‌شده تعلق می‌گیرد. هرگونه سوءاستفاده منجر به حذف امتیاز خواهد شد.',
  faq: 'چگونه امتیاز بگیرم؟ با خرید و فعالیت در سایت امتیاز دریافت می‌کنید. چگونه امتیاز را خرج کنم؟ امتیازها را می‌توانید برای دریافت تخفیف یا پاداش استفاده کنید.',
  guide: 'برای استفاده از برنامه وفاداری ابتدا وارد حساب کاربری خود شوید و سپس از بخش وفاداری، امتیازات و پاداش‌های خود را مدیریت کنید.'
})
</script> 

<template>
  <div class="bg-white rounded-2xl shadow-xl border border-gray-100 p-8 my-8 flex flex-col items-center">
    <button
      class="bg-gradient-to-r from-blue-600 to-purple-600 text-white px-10 py-4 rounded-xl text-xl font-bold shadow-lg hover:scale-105 transition-all"
      @click="submitSurvey"
    >
      ارسال نظرسنجی
    </button>
    <div v-if="submitted" class="mt-6 text-green-600 text-lg font-semibold">
      نظرسنجی شما با موفقیت ثبت شد!
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
const submitted = ref(false)
function submitSurvey() {
  // اینجا می‌توانید ارسال به سرور را اضافه کنید
  submitted.value = true
}
</script> 

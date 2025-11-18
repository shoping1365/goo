<template>
     <div class="bg-white rounded-2xl shadow-xl border border-gray-100 p-8 mb-8">
       <div class="flex items-center mb-6">
         <div class="p-3 bg-gradient-to-r from-emerald-400 to-teal-600 rounded-xl shadow-lg">
           <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
             <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"></path>
           </svg>
         </div>
         <div class="mr-4">
           <h2 class="text-2xl font-bold text-gray-900">نظرات آزاد</h2>
           <p class="text-gray-600 mt-1">نظر خود را به صورت آزاد درباره تجربه خرید ثبت کنید</p>
         </div>
       </div>
   
       <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
         <!-- ستون اول -->
         <div class="space-y-6">
           <!-- نظر کلی درباره تجربه خرید -->
           <div class="bg-gradient-to-r from-blue-50 to-indigo-50 rounded-xl p-6 border border-blue-100">
             <div class="flex items-center mb-4">
               <div class="p-2 bg-blue-500 rounded-lg mr-3">
                 <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                   <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                 </svg>
               </div>
               <h3 class="text-lg font-semibold text-gray-900">نظر کلی درباره تجربه خرید</h3>
             </div>
             
             <div class="relative">
               <textarea
                 v-model="comments.overallExperience"
                 rows="4"
                 placeholder="تجربه کلی خود را از خرید در این فروشگاه بنویسید..."
                 class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 resize-none transition-colors"
                 maxlength="500"
               ></textarea>
               <div class="absolute bottom-2 left-2 text-xs text-gray-400">
                 {{ comments.overallExperience.length }}/500
               </div>
             </div>
             
             <div class="mt-3 flex items-center space-x-2 space-x-reverse">
               <div class="flex space-x-1 space-x-reverse">
                 <span class="text-sm text-gray-600">مفید:</span>
                 <button
                   v-for="i in 5"
                   :key="i"
                   @click="helpfulRating.overall = i"
                   class="text-lg transition-colors"
                   :class="i <= helpfulRating.overall ? 'text-yellow-400' : 'text-gray-300 hover:text-yellow-300'"
                 >★</button>
               </div>
             </div>
           </div>
   
           <!-- نقاط قوت -->
           <div class="bg-gradient-to-r from-green-50 to-emerald-50 rounded-xl p-6 border border-green-100">
             <div class="flex items-center mb-4">
               <div class="p-2 bg-green-500 rounded-lg mr-3">
                 <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                   <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
                 </svg>
               </div>
               <h3 class="text-lg font-semibold text-gray-900">نقاط قوت</h3>
             </div>
             
             <div class="relative">
               <textarea
                 v-model="comments.strengths"
                 rows="4"
                 placeholder="چه چیزهایی در این فروشگاه خوب بود؟ نقاط مثبت را بنویسید..."
                 class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500 resize-none transition-colors"
                 maxlength="400"
               ></textarea>
               <div class="absolute bottom-2 left-2 text-xs text-gray-400">
                 {{ comments.strengths.length }}/400
               </div>
             </div>
             
             <!-- پیشنهادات سریع -->
             <div class="mt-3">
               <p class="text-sm text-gray-600 mb-2">پیشنهادات سریع:</p>
               <div class="flex flex-wrap gap-2">
                 <button
                   v-for="suggestion in strengthSuggestions"
                   :key="suggestion"
                   @click="addStrengthSuggestion(suggestion)"
                   class="px-3 py-1 bg-green-100 text-green-700 rounded-full text-sm hover:bg-green-200 transition-colors"
                 >
                   {{ suggestion }}
                 </button>
               </div>
             </div>
           </div>
         </div>
   
         <!-- ستون دوم -->
         <div class="space-y-6">
           <!-- نقاط ضعف -->
           <div class="bg-gradient-to-r from-red-50 to-pink-50 rounded-xl p-6 border border-red-100">
             <div class="flex items-center mb-4">
               <div class="p-2 bg-red-500 rounded-lg mr-3">
                 <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                   <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                 </svg>
               </div>
               <h3 class="text-lg font-semibold text-gray-900">نقاط ضعف</h3>
             </div>
             
             <div class="relative">
               <textarea
                 v-model="comments.weaknesses"
                 rows="4"
                 placeholder="چه چیزهایی نیاز به بهبود دارد؟ مشکلات را بنویسید..."
                 class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500 resize-none transition-colors"
                 maxlength="400"
               ></textarea>
               <div class="absolute bottom-2 left-2 text-xs text-gray-400">
                 {{ comments.weaknesses.length }}/400
               </div>
             </div>
             
             <!-- پیشنهادات سریع -->
             <div class="mt-3">
               <p class="text-sm text-gray-600 mb-2">پیشنهادات سریع:</p>
               <div class="flex flex-wrap gap-2">
                 <button
                   v-for="suggestion in weaknessSuggestions"
                   :key="suggestion"
                   @click="addWeaknessSuggestion(suggestion)"
                   class="px-3 py-1 bg-red-100 text-red-700 rounded-full text-sm hover:bg-red-200 transition-colors"
                 >
                   {{ suggestion }}
                 </button>
               </div>
             </div>
           </div>
   
           <!-- پیشنهادات بهبود -->
           <div class="bg-gradient-to-r from-purple-50 to-violet-50 rounded-xl p-6 border border-purple-100">
             <div class="flex items-center mb-4">
               <div class="p-2 bg-purple-500 rounded-lg mr-3">
                 <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                   <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"></path>
                 </svg>
               </div>
               <h3 class="text-lg font-semibold text-gray-900">پیشنهادات بهبود</h3>
             </div>
             
             <div class="relative">
               <textarea
                 v-model="comments.improvements"
                 rows="4"
                 placeholder="چه پیشنهاداتی برای بهبود خدمات دارید؟"
                 class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500 resize-none transition-colors"
                 maxlength="400"
               ></textarea>
               <div class="absolute bottom-2 left-2 text-xs text-gray-400">
                 {{ comments.improvements.length }}/400
               </div>
             </div>
             
             <!-- پیشنهادات سریع -->
             <div class="mt-3">
               <p class="text-sm text-gray-600 mb-2">پیشنهادات سریع:</p>
               <div class="flex flex-wrap gap-2">
                 <button
                   v-for="suggestion in improvementSuggestions"
                   :key="suggestion"
                   @click="addImprovementSuggestion(suggestion)"
                   class="px-3 py-1 bg-purple-100 text-purple-700 rounded-full text-sm hover:bg-purple-200 transition-colors"
                 >
                   {{ suggestion }}
                 </button>
               </div>
             </div>
           </div>
         </div>
       </div>
    </div>
</template>

<script setup>
import { ref } from 'vue'

const comments = ref({
  overallExperience: '',
  strengths: '',
  weaknesses: '',
  improvements: ''
})

const helpfulRating = ref({
  overall: 0
})

const strengthSuggestions = [
  'پشتیبانی عالی',
  'ارسال سریع',
  'بسته‌بندی مناسب',
  'قیمت مناسب',
  'تنوع محصولات'
]

const weaknessSuggestions = [
  'تاخیر در ارسال',
  'پاسخگویی ضعیف',
  'قیمت بالا',
  'عدم تنوع',
  'کیفیت پایین'
]

const improvementSuggestions = [
  'افزایش تنوع محصولات',
  'بهبود پشتیبانی',
  'کاهش زمان ارسال',
  'ارائه تخفیف‌های بیشتر',
  'بهبود بسته‌بندی'
]

function addStrengthSuggestion(suggestion) {
  if (!comments.value.strengths.includes(suggestion)) {
    comments.value.strengths += (comments.value.strengths ? '، ' : '') + suggestion
  }
}

function addWeaknessSuggestion(suggestion) {
  if (!comments.value.weaknesses.includes(suggestion)) {
    comments.value.weaknesses += (comments.value.weaknesses ? '، ' : '') + suggestion
  }
}

function addImprovementSuggestion(suggestion) {
  if (!comments.value.improvements.includes(suggestion)) {
    comments.value.improvements += (comments.value.improvements ? '، ' : '') + suggestion
  }
}
</script>

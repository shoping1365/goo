<template>
     <div class="validation-preview bg-white rounded-2xl shadow-xl border border-gray-100 p-8 mb-8">
       <div class="flex items-center mb-6">
         <div class="p-3 bg-gradient-to-r from-green-400 to-green-600 rounded-xl shadow-lg">
           <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
             <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
           </svg>
         </div>
         <div class="mr-4">
           <h2 class="text-2xl font-bold text-gray-900">پیش‌نمایش و بررسی نهایی</h2>
           <p class="text-gray-600 mt-1">پاسخ‌های خود را بررسی کنید و در صورت نیاز ویرایش کنید</p>
         </div>
       </div>
   
       <!-- Validation Summary -->
       <div class="mb-6">
         <div class="bg-blue-50 rounded-lg p-6 border border-blue-200">
           <div class="flex items-center space-x-3 space-x-reverse">
             <div class="p-2 bg-blue-500 rounded-lg">
               <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                 <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
               </svg>
             </div>
             <div>
               <h3 class="font-semibold text-blue-900">وضعیت تکمیل</h3>
               <p class="text-sm text-blue-700">{{ completedSections }}/{{ totalSections }} بخش تکمیل شده</p>
             </div>
           </div>
         </div>
       </div>
   
       <!-- Sections Preview -->
       <div class="space-y-6">
         <!-- Order Information -->
         <div class="preview-section bg-gray-50 rounded-lg p-6">
           <div class="flex items-center justify-between mb-3">
             <h3 class="text-lg font-semibold text-gray-900">اطلاعات سفارش</h3>
             <span class="px-2 py-1 bg-green-100 text-green-800 rounded-full text-sm">تکمیل شده</span>
           </div>
           <div class="grid grid-cols-1 md:grid-cols-2 gap-6 text-sm">
             <div>
               <span class="font-medium text-gray-700">شماره سفارش:</span>
               <span class="text-gray-900 mr-2">{{ orderInfo.orderNumber }}</span>
             </div>
             <div>
               <span class="font-medium text-gray-700">تاریخ سفارش:</span>
               <span class="text-gray-900 mr-2">{{ orderInfo.orderDate }}</span>
             </div>
           </div>
         </div>
   
         <!-- Overall Rating -->
         <div class="preview-section bg-gray-50 rounded-lg p-6">
           <div class="flex items-center justify-between mb-3">
             <h3 class="text-lg font-semibold text-gray-900">امتیاز کلی</h3>
             <span :class="getValidationClass('overallRating')" class="px-2 py-1 rounded-full text-sm">
               {{ getValidationText('overallRating') }}
             </span>
           </div>
           <div v-if="surveyData.overallRating > 0" class="flex items-center space-x-2 space-x-reverse">
             <div class="flex">
               <span v-for="i in 5" :key="i" class="text-yellow-400 text-xl">
                 {{ i <= surveyData.overallRating ? '★' : '☆' }}
               </span>
             </div>
             <span class="text-gray-700">{{ getRatingText(surveyData.overallRating) }}</span>
           </div>
           <div v-else class="text-red-600 text-sm">امتیاز کلی وارد نشده است</div>
         </div>
   
         <!-- Detailed Ratings -->
         <div class="preview-section bg-gray-50 rounded-lg p-6">
           <div class="flex items-center justify-between mb-3">
             <h3 class="text-lg font-semibold text-gray-900">امتیازات جزئی</h3>
             <span :class="getValidationClass('detailedRatings')" class="px-2 py-1 rounded-full text-sm">
               {{ getValidationText('detailedRatings') }}
             </span>
           </div>
           <div v-if="hasDetailedRatings" class="grid grid-cols-2 md:grid-cols-3 gap-3">
             <div v-for="(rating, key) in surveyData.detailedRatings" :key="key" class="text-sm">
               <span class="font-medium text-gray-700">{{ getCategoryName(key) }}:</span>
               <span class="text-gray-900 mr-2">{{ rating }}/5</span>
             </div>
           </div>
           <div v-else class="text-red-600 text-sm">امتیازات جزئی وارد نشده است</div>
         </div>
   
         <!-- Comments -->
         <div class="preview-section bg-gray-50 rounded-lg p-6">
           <div class="flex items-center justify-between mb-3">
             <h3 class="text-lg font-semibold text-gray-900">نظرات و پیشنهادات</h3>
             <span :class="getValidationClass('comments')" class="px-2 py-1 rounded-full text-sm">
               {{ getValidationText('comments') }}
             </span>
           </div>
           <div v-if="hasComments" class="space-y-2">
             <div v-for="(comment, key) in surveyData.comments" :key="key" class="text-sm">
               <span class="font-medium text-gray-700">{{ getCommentLabel(key) }}:</span>
               <p class="text-gray-900 mt-1">{{ comment || 'وارد نشده' }}</p>
             </div>
           </div>
           <div v-else class="text-yellow-600 text-sm">هیچ نظری وارد نشده است (اختیاری)</div>
         </div>
   
         <!-- Specialized Questions -->
         <div class="preview-section bg-gray-50 rounded-lg p-6">
           <div class="flex items-center justify-between mb-3">
             <h3 class="text-lg font-semibold text-gray-900">سوالات تخصصی</h3>
             <span :class="getValidationClass('specializedQuestions')" class="px-2 py-1 rounded-full text-sm">
               {{ getValidationText('specializedQuestions') }}
             </span>
           </div>
           <div v-if="hasSpecializedAnswers" class="space-y-2">
             <div v-for="(answer, key) in surveyData.specializedQuestions" :key="key" class="text-sm">
               <span class="font-medium text-gray-700">{{ getQuestionName(key) }}:</span>
               <span class="text-gray-900 mr-2">{{ getAnswerText(key, answer) }}</span>
             </div>
           </div>
           <div v-else class="text-red-600 text-sm">سوالات تخصصی پاسخ داده نشده است</div>
         </div>
   
         <!-- Optional Section -->
         <div class="preview-section bg-gray-50 rounded-lg p-6">
           <div class="flex items-center justify-between mb-3">
             <h3 class="text-lg font-semibold text-gray-900">بخش اختیاری</h3>
             <span class="px-2 py-1 bg-blue-100 text-blue-800 rounded-full text-sm">اختیاری</span>
           </div>
           <div v-if="hasOptionalData" class="space-y-2">
             <div v-if="surveyData.optional?.newsletter" class="text-sm">
               <span class="font-medium text-gray-700">عضویت در خبرنامه:</span>
               <span class="text-green-600 mr-2">بله</span>
             </div>
             <div v-if="surveyData.optional?.contactInfo" class="text-sm">
               <span class="font-medium text-gray-700">اطلاعات تماس:</span>
               <span class="text-gray-900 mr-2">{{ surveyData.optional.contactInfo }}</span>
             </div>
           </div>
           <div v-else class="text-gray-500 text-sm">هیچ اطلاعات اختیاری وارد نشده است</div>
         </div>
       </div>
   
       <!-- Action Buttons -->
       <div class="flex justify-between mt-8">
         <button 
           @click="editSurvey"
           class="px-6 py-3 bg-blue-600 text-white rounded-lg font-medium hover:bg-blue-700 transition-colors"
         >
           ویرایش نظرسنجی
         </button>
         
         <div class="flex space-x-3 space-x-reverse">
           <button 
             @click="saveDraft"
             class="px-6 py-3 bg-gray-500 text-white rounded-lg font-medium hover:bg-gray-600 transition-colors"
           >
             ذخیره پیش‌نویس
           </button>
           
           <button 
             @click="submitSurvey"
             :disabled="!isValid"
             class="px-6 py-3 bg-green-600 text-white rounded-lg font-medium hover:bg-green-700 transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
           >
             ارسال نهایی
           </button>
         </div>
       </div>
   
       <!-- Validation Errors -->
       <div v-if="validationErrors.length > 0" class="mt-6 bg-red-50 border border-red-200 rounded-lg p-6">
         <h4 class="font-semibold text-red-800 mb-2">خطاهای اعتبارسنجی:</h4>
         <ul class="list-disc list-inside space-y-1">
           <li v-for="error in validationErrors" :key="error" class="text-sm text-red-700">
             {{ error }}
           </li>
         </ul>
       </div>
     </div>
   </template>
   
   <script setup lang="ts">
   import { ref, computed } from 'vue'
   
   interface SurveyData {
     overallRating: number
     detailedRatings: Record<string, number>
     comments: Record<string, string>
     specializedQuestions: Record<string, string>
     optional?: {
       newsletter: boolean
       contactInfo: string
     }
   }
   
   interface OrderInfo {
     orderNumber: string
     orderDate: string
   }
   
   interface Props {
     surveyData: SurveyData
     orderInfo: OrderInfo
   }
   
   const props = defineProps<Props>()
   const emit = defineEmits(['edit', 'save-draft', 'submit'])
   
   const validationErrors = ref<string[]>([])
   
   // Computed properties
   const totalSections = 6
   const completedSections = computed(() => {
     let completed = 0
     if (props.surveyData.overallRating > 0) completed++
     if (hasDetailedRatings.value) completed++
     if (hasComments.value) completed++
     if (hasSpecializedAnswers.value) completed++
     if (props.orderInfo.orderNumber) completed++
     return completed
   })
   
   const hasDetailedRatings = computed(() => {
     return Object.values(props.surveyData.detailedRatings).some(rating => rating > 0)
   })
   
   const hasComments = computed(() => {
     return Object.values(props.surveyData.comments).some(comment => comment.trim().length > 0)
   })
   
   const hasSpecializedAnswers = computed(() => {
     return Object.values(props.surveyData.specializedQuestions).some(answer => answer)
   })
   
   const hasOptionalData = computed(() => {
     return props.surveyData.optional && (
       props.surveyData.optional.newsletter || 
       props.surveyData.optional.contactInfo
     )
   })
   
   const isValid = computed(() => {
     validationErrors.value = []
     
     if (props.surveyData.overallRating === 0) {
       validationErrors.value.push('امتیاز کلی الزامی است')
     }
     
     if (!hasDetailedRatings.value) {
       validationErrors.value.push('حداقل یک امتیاز جزئی الزامی است')
     }
     
     if (!hasSpecializedAnswers.value) {
       validationErrors.value.push('پاسخ به سوالات تخصصی الزامی است')
     }
     
     return validationErrors.value.length === 0
   })
   
   // Methods
   const getValidationClass = (section: string) => {
     switch (section) {
       case 'overallRating':
         return props.surveyData.overallRating > 0 ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'
       case 'detailedRatings':
         return hasDetailedRatings.value ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'
       case 'comments':
         return hasComments.value ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800'
       case 'specializedQuestions':
         return hasSpecializedAnswers.value ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'
       default:
         return 'bg-gray-100 text-gray-800'
     }
   }
   
   const getValidationText = (section: string) => {
     switch (section) {
       case 'overallRating':
         return props.surveyData.overallRating > 0 ? 'تکمیل شده' : 'ناقص'
       case 'detailedRatings':
         return hasDetailedRatings.value ? 'تکمیل شده' : 'ناقص'
       case 'comments':
         return hasComments.value ? 'تکمیل شده' : 'اختیاری'
       case 'specializedQuestions':
         return hasSpecializedAnswers.value ? 'تکمیل شده' : 'ناقص'
       default:
         return 'نامشخص'
     }
   }
   
   const getRatingText = (rating: number) => {
     switch (rating) {
       case 1: return 'خیلی بد'
       case 2: return 'بد'
       case 3: return 'متوسط'
       case 4: return 'خوب'
       case 5: return 'عالی'
       default: return 'امتیاز دهید'
     }
   }
   
   const getCategoryName = (key: string) => {
     const names = {
       productQuality: 'کیفیت محصول',
       pricing: 'قیمت‌گذاری',
       deliverySpeed: 'سرعت تحویل',
       packaging: 'بسته‌بندی',
       afterSales: 'پشتیبانی',
       userInterface: 'رابط کاربری'
     }
     return names[key] || key
   }
   
   const getCommentLabel = (key: string) => {
     const labels = {
       overallExperience: 'تجربه کلی',
       strengths: 'نقاط قوت',
       weaknesses: 'نقاط ضعف',
       improvements: 'پیشنهادات بهبود'
     }
     return labels[key] || key
   }
   
   const getQuestionName = (key: string) => {
     const names = {
       paymentSatisfaction: 'رضایت از پرداخت',
       shippingSatisfaction: 'رضایت از ارسال',
       communicationQuality: 'کیفیت ارتباطات',
       productInfoAccuracy: 'دقت اطلاعات محصول',
       supportResponseSpeed: 'سرعت پاسخگویی'
     }
     return names[key] || key
   }
   
   const getAnswerText = (key: string, value: string) => {
     const allOptions = {
       paymentSatisfaction: {
         very_satisfied: 'خیلی راضی',
         satisfied: 'راضی',
         neutral: 'متوسط',
         dissatisfied: 'ناراضی'
       },
       shippingSatisfaction: {
         very_satisfied: 'خیلی راضی',
         satisfied: 'راضی',
         neutral: 'متوسط',
         dissatisfied: 'ناراضی'
       },
       communicationQuality: {
         excellent: 'عالی',
         good: 'خوب',
         average: 'متوسط',
         poor: 'ضعیف'
       },
       productInfoAccuracy: {
         very_accurate: 'کاملاً دقیق',
         accurate: 'دقیق',
         average: 'متوسط',
         inaccurate: 'غیر دقیق'
       },
       supportResponseSpeed: {
         very_fast: 'خیلی سریع',
         fast: 'سریع',
         average: 'متوسط',
         slow: 'کند'
       }
     }
     
     return allOptions[key]?.[value] || value
   }
   
   const editSurvey = () => {
     emit('edit')
   }
   
   const saveDraft = () => {
     emit('save-draft', props.surveyData)
   }
   
   const submitSurvey = () => {
     if (isValid.value) {
       emit('submit', props.surveyData)
     }
   }
   </script>
   
   <style scoped>
   .validation-preview {
     transition: all 0.3s ease;
   }
   
   .preview-section {
     transition: all 0.2s ease;
   }
   
   .preview-section:hover {
     background-color: #f8fafc;
     transform: translateY(-1px);
   }
   
   /* Animation for validation status */
   @keyframes pulse {
     0%, 100% {
       opacity: 1;
     }
     50% {
       opacity: 0.7;
     }
   }
   
   .bg-red-100 {
     animation: pulse 2s infinite;
   }
   
   /* Smooth transitions */
   .transition-colors {
     transition: all 0.2s ease-in-out;
   }
   
   /* Focus styles */
   button:focus {
     outline: 2px solid #6366f1;
     outline-offset: 2px;
   }
   </style>

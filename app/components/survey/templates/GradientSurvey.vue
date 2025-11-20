<template>
     <div class="bg-gradient-to-br from-indigo-500 via-purple-500 to-pink-500 rounded-2xl shadow-2xl overflow-hidden">
       <div class="p-6 text-white text-center">
         <div class="w-16 h-16 bg-white bg-opacity-20 rounded-full mx-auto mb-4 flex items-center justify-center backdrop-blur-sm">
           <svg class="w-8 h-8" fill="currentColor" viewBox="0 0 20 20">
             <path d="M2 10a8 8 0 018-8v8h8a8 8 0 11-16 0z"></path>
             <path d="M12 2.252A8.014 8.014 0 0117.748 8H12V2.252z"></path>
           </svg>
         </div>
         <h2 class="text-2xl font-bold mb-2">نظرسنجی خرید</h2>
         <p class="text-indigo-100">تجربه شما برای بهبود خدمات</p>
       </div>
       
       <div class="p-6 bg-white bg-opacity-10 backdrop-blur-sm">
         <!-- Order Info -->
         <div class="bg-white bg-opacity-20 backdrop-blur-sm rounded-xl p-6 mb-6 border border-white border-opacity-30">
           <h3 class="font-semibold text-white mb-2">اطلاعات سفارش</h3>
           <div class="grid grid-cols-2 gap-6 text-sm">
             <div>
               <span class="text-indigo-200">محصول:</span>
               <span class="font-medium text-white mr-2">{{ order.productName }}</span>
             </div>
             <div>
               <span class="text-indigo-200">مبلغ:</span>
               <span class="font-medium text-white mr-2">{{ formatPrice(order.totalAmount) }}</span>
             </div>
           </div>
         </div>
   
         <!-- Rating -->
         <div class="mb-6">
           <h3 class="text-lg font-semibold text-white mb-4 text-center">امتیاز دهید</h3>
           <div class="flex justify-center space-x-3 space-x-reverse">
             <button 
               v-for="rating in 5" 
               :key="rating"
               :class="[
                 'w-10 h-10 rounded-lg text-white font-bold transition-all transform hover:scale-110',
                 selectedRating === rating 
                   ? 'bg-white bg-opacity-30 shadow-lg transform scale-110' 
                   : 'bg-white bg-opacity-20 hover:bg-opacity-30'
               ]"
               @click="selectedRating = rating"
             >
               {{ rating }}
             </button>
           </div>
           <div class="text-center mt-3 text-sm text-indigo-200">
             {{ getRatingText(selectedRating) }}
           </div>
         </div>
   
         <!-- Questions -->
         <div class="space-y-4 mb-6">
           <div>
             <label class="block text-sm font-medium text-white mb-2">کیفیت محصول:</label>
             <div class="grid grid-cols-2 gap-2">
               <button 
                 v-for="option in qualityOptions" 
                 :key="option.value"
                 :class="[
                   'px-3 py-2 rounded-lg text-sm transition-all',
                   qualityRating === option.value 
                     ? 'bg-white bg-opacity-30 text-white shadow-md' 
                     : 'bg-white bg-opacity-20 hover:bg-opacity-30 text-white'
                 ]"
                 @click="qualityRating = option.value"
               >
                 {{ option.label }}
               </button>
             </div>
           </div>
   
           <div>
             <label class="block text-sm font-medium text-white mb-2">سرعت ارسال:</label>
             <div class="grid grid-cols-3 gap-2">
               <button 
                 v-for="option in deliveryOptions" 
                 :key="option.value"
                 :class="[
                   'px-3 py-2 rounded-lg text-sm transition-all',
                   deliveryRating === option.value 
                     ? 'bg-white bg-opacity-30 text-white shadow-md' 
                     : 'bg-white bg-opacity-20 hover:bg-opacity-30 text-white'
                 ]"
                 @click="deliveryRating = option.value"
               >
                 {{ option.label }}
               </button>
             </div>
           </div>
         </div>
   
         <!-- Comment -->
         <div class="mb-6">
           <label class="block text-sm font-medium text-white mb-2">نظر شما:</label>
           <textarea 
             v-model="comment"
             class="w-full px-4 py-3 bg-white bg-opacity-20 backdrop-blur-sm border border-white border-opacity-30 rounded-xl text-white placeholder-white placeholder-opacity-70 focus:ring-2 focus:ring-white focus:border-white" 
             rows="3" 
             placeholder="تجربه خود را با ما به اشتراک بگذارید..."
           ></textarea>
         </div>
   
         <!-- Submit Button -->
         <button 
           :disabled="!selectedRating || submitting"
           class="w-full bg-white bg-opacity-20 hover:bg-opacity-30 disabled:bg-opacity-10 backdrop-blur-sm text-white font-semibold py-3 px-6 rounded-xl transition-all transform hover:scale-105 border border-white border-opacity-30 flex items-center justify-center space-x-2 space-x-reverse"
           @click="submitSurvey"
         >
           <svg v-if="submitting" class="animate-spin w-4 h-4" fill="none" viewBox="0 0 24 24">
             <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
             <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
           </svg>
           <span>{{ submitting ? 'در حال ارسال...' : 'ارسال نظر' }}</span>
         </button>
       </div>
     </div>
   </template>
   
   <script setup lang="ts">
   import { ref } from 'vue'
   
   interface Order {
     id: string
     customerName: string
     productName: string
     orderDate: string
     totalAmount: number
   }
   
   interface Props {
     order: Order
   }
   
   const props = defineProps<Props>()
   const emit = defineEmits<{
     submit: [data: any]
   }>()
   
   const selectedRating = ref(0)
   const qualityRating = ref('')
   const deliveryRating = ref('')
   const comment = ref('')
   const submitting = ref(false)
   
   const qualityOptions = [
     { value: 'excellent', label: 'عالی' },
     { value: 'good', label: 'خوب' },
     { value: 'average', label: 'متوسط' },
     { value: 'poor', label: 'ضعیف' }
   ]
   
   const deliveryOptions = [
     { value: 'fast', label: 'سریع' },
     { value: 'normal', label: 'معمولی' },
     { value: 'slow', label: 'کند' }
   ]
   
   const formatPrice = (price: number) => {
     return new Intl.NumberFormat('fa-IR').format(price) + ' تومان'
   }
   
   const getRatingText = (rating: number) => {
     const texts = {
       1: 'خیلی ضعیف',
       2: 'ضعیف',
       3: 'متوسط',
       4: 'خوب',
       5: 'عالی'
     }
     return texts[rating as keyof typeof texts] || ''
   }
   
   const submitSurvey = async () => {
     if (!selectedRating.value) return
     
     submitting.value = true
     try {
       await new Promise(resolve => setTimeout(resolve, 1000))
       
       emit('submit', {
         rating: selectedRating.value,
         qualityRating: qualityRating.value,
         deliveryRating: deliveryRating.value,
         comment: comment.value,
         submittedAt: new Date().toISOString()
       })
     } finally {
       submitting.value = false
     }
   }
   </script>
   
   <style scoped>
   .space-x-reverse > :not([hidden]) ~ :not([hidden]) {
     --tw-space-x-reverse: 1;
   }
   </style>

 <template>
     <div class="progress-bar bg-white rounded-xl shadow-lg border border-gray-200 p-6 mb-6">
       <div class="flex items-center justify-between mb-4">
         <h3 class="text-lg font-semibold text-gray-900">پیشرفت تکمیل نظرسنجی</h3>
         <span class="text-sm font-medium text-blue-600">{{ progress }}% تکمیل شده</span>
       </div>
   
       <!-- Progress Bar -->
       <div class="relative mb-6">
         <div class="w-full bg-gray-200 rounded-full h-3">
           <div 
             class="bg-gradient-to-r from-blue-500 to-purple-600 h-3 rounded-full transition-all duration-500 ease-out"
             :style="{ width: `${progress}%` }"
           ></div>
         </div>
         <div class="absolute -top-2 left-0 w-full flex justify-between">
           <div 
             v-for="(step, index) in steps" 
             :key="index"
             class="flex flex-col items-center"
           >
             <div 
               class="w-6 h-6 rounded-full border-2 flex items-center justify-center text-xs font-bold transition-all duration-300"
               :class="getStepClass(index)"
             >
               <span v-if="index < currentStep" class="text-white">✓</span>
               <span v-else-if="index === currentStep" class="text-blue-600">{{ index + 1 }}</span>
               <span v-else class="text-gray-400">{{ index + 1 }}</span>
             </div>
             <span 
               class="text-xs mt-1 text-center max-w-20"
               :class="index <= currentStep ? 'text-gray-900' : 'text-gray-400'"
             >
               {{ step.title }}
             </span>
           </div>
         </div>
       </div>
   
       <!-- Current Step Info -->
       <div class="bg-blue-50 rounded-lg p-6">
         <div class="flex items-center space-x-3 space-x-reverse">
           <div class="p-2 bg-blue-500 rounded-lg">
             <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
               <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
             </svg>
           </div>
           <div>
             <h4 class="font-medium text-blue-900">مرحله {{ currentStep + 1 }} از {{ steps.length }}</h4>
             <p class="text-sm text-blue-700">{{ currentStepInfo }}</p>
           </div>
         </div>
       </div>
   
       <!-- Step Navigation -->
       <div class="flex justify-between mt-4">
         <button 
           :disabled="currentStep === 0"
           class="px-4 py-2 bg-gray-100 text-gray-700 rounded-lg font-medium transition-colors disabled:opacity-50 disabled:cursor-not-allowed hover:bg-gray-200"
           @click="previousStep"
         >
           مرحله قبل
         </button>
         
         <div class="flex space-x-2 space-x-reverse">
           <button 
             v-for="(step, index) in steps" 
             :key="index"
             class="px-3 py-1 rounded-lg text-sm font-medium transition-colors"
             :class="getStepButtonClass(index)"
             @click="goToStep(index)"
           >
             {{ step.title }}
           </button>
         </div>
   
         <button 
           :disabled="currentStep === steps.length - 1"
           class="px-4 py-2 bg-blue-600 text-white rounded-lg font-medium transition-colors disabled:opacity-50 disabled:cursor-not-allowed hover:bg-blue-700"
           @click="nextStep"
         >
           مرحله بعد
         </button>
       </div>
   
       <!-- Estimated Time -->
       <div class="mt-4 text-center">
         <span class="text-sm text-gray-500">
           زمان تخمینی باقی‌مانده: {{ estimatedTimeRemaining }}
         </span>
       </div>
     </div>
   </template>
   
   <script setup lang="ts">
   import { ref, computed, watch } from 'vue'
   
   interface Step {
     title: string
     description: string
     completed: boolean
   }
   
   interface Props {
     currentStep?: number
     steps?: Step[]
   }
   
   const props = withDefaults(defineProps<Props>(), {
     currentStep: 0,
     steps: () => [
       { title: 'اطلاعات سفارش', description: 'بررسی جزئیات سفارش', completed: false },
       { title: 'امتیاز کلی', description: 'امتیازدهی کلی تجربه خرید', completed: false },
       { title: 'امتیاز جزئی', description: 'امتیازدهی جنبه‌های مختلف', completed: false },
       { title: 'نظرات', description: 'نظرات و پیشنهادات', completed: false },
       { title: 'سوالات تخصصی', description: 'سوالات تخصصی', completed: false },
       { title: 'بخش اختیاری', description: 'اطلاعات اضافی', completed: false },
       { title: 'پاداش', description: 'دریافت پاداش', completed: false },
       { title: 'ارسال', description: 'ارسال نهایی', completed: false }
     ]
   })
   
   const emit = defineEmits(['step-change', 'step-complete'])
   
   const currentStep = ref(props.currentStep)
   const steps = ref(props.steps)
   
   // Computed properties
   const progress = computed(() => {
     const completedSteps = steps.value.filter(step => step.completed).length
     return Math.round((completedSteps / steps.value.length) * 100)
   })
   
   const currentStepInfo = computed(() => {
     return steps.value[currentStep.value]?.description || ''
   })
   
   const estimatedTimeRemaining = computed(() => {
     const remainingSteps = steps.value.length - currentStep.value
     const estimatedMinutes = remainingSteps * 2 // 2 minutes per step
     if (estimatedMinutes < 60) {
       return `${estimatedMinutes} دقیقه`
     } else {
       const hours = Math.floor(estimatedMinutes / 60)
       const minutes = estimatedMinutes % 60
       return `${hours} ساعت و ${minutes} دقیقه`
     }
   })
   
   // Methods
   const getStepClass = (index: number) => {
     if (index < currentStep.value) {
       return 'bg-green-500 border-green-500' // Completed
     } else if (index === currentStep.value) {
       return 'bg-blue-500 border-blue-500' // Current
     } else {
       return 'bg-white border-gray-300' // Pending
     }
   }
   
   const getStepButtonClass = (index: number) => {
     if (index < currentStep.value) {
       return 'bg-green-100 text-green-800 hover:bg-green-200' // Completed
     } else if (index === currentStep.value) {
       return 'bg-blue-100 text-blue-800' // Current
     } else {
       return 'bg-gray-100 text-gray-600 hover:bg-gray-200' // Pending
     }
   }
   
   const nextStep = () => {
     if (currentStep.value < steps.value.length - 1) {
       currentStep.value++
       emit('step-change', currentStep.value)
     }
   }
   
   const previousStep = () => {
     if (currentStep.value > 0) {
       currentStep.value--
       emit('step-change', currentStep.value)
     }
   }
   
   const goToStep = (index: number) => {
     currentStep.value = index
     emit('step-change', index)
   }
   
   const completeCurrentStep = () => {
     if (steps.value[currentStep.value]) {
       steps.value[currentStep.value].completed = true
       emit('step-complete', currentStep.value)
     }
   }
   
   // Watch for changes
   watch(currentStep, (newStep) => {
     emit('step-change', newStep)
   })
   
   // Expose methods for parent component
   defineExpose({
     nextStep,
     previousStep,
     goToStep,
     completeCurrentStep,
     currentStep,
     progress
   })
   </script>
   
   <style scoped>
   .progress-bar {
     transition: all 0.3s ease;
   }
   
   .progress-bar:hover {
     box-shadow: 0 10px 25px rgba(0, 0, 0, 0.1);
   }
   
   /* Smooth transitions for progress bar */
   .bg-gradient-to-r {
     transition: width 0.5s ease-out;
   }
   
   /* Step button animations */
   button:not(:disabled):hover {
     transform: translateY(-1px);
     transition: transform 0.2s ease;
   }
   
   /* Pulse animation for current step */
   @keyframes pulse {
     0%, 100% {
       opacity: 1;
     }
     50% {
       opacity: 0.7;
     }
   }
   
   .bg-blue-500 {
     animation: pulse 2s infinite;
   }
   </style>

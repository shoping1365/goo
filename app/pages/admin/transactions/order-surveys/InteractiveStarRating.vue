<template>
     <div class="interactive-star-rating">
       <div class="stars-container" role="radiogroup" :aria-label="ariaLabel">
         <button
           v-for="i in maxStars"
           :key="i"
           type="button"
           :aria-checked="i <= rating"
           :aria-label="`${i} ستاره`"
           class="star-button"
           :class="getStarClass(i)"
           :disabled="disabled"
           @click="handleClick(i)"
           @mouseenter="handleHover(i)"
           @mouseleave="handleMouseLeave"
           @keydown="handleKeydown"
         >
           <svg 
             class="star-icon" 
             viewBox="0 0 24 24" 
             fill="currentColor"
             :class="getStarIconClass(i)"
           >
             <path d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z"/>
           </svg>
         </button>
       </div>
       
       <!-- Rating Text -->
       <div v-if="showText" class="rating-text">
         <span class="text-lg font-semibold" :class="getRatingTextColor()">
           {{ getRatingText() }}
         </span>
         <span v-if="showValue" class="text-sm text-gray-500 mr-2">
           ({{ rating }}/{{ maxStars }})
         </span>
       </div>
       
       <!-- Rating Description -->
       <div v-if="showDescription && rating > 0" class="rating-description">
         <p class="text-sm text-gray-600">{{ getRatingDescription() }}</p>
       </div>
     </div>
   </template>
   
   <script setup lang="ts">
   import { computed, ref, watch } from 'vue'
   
   interface Props {
     modelValue?: number
     maxStars?: number
     size?: 'sm' | 'md' | 'lg' | 'xl'
     showText?: boolean
     showValue?: boolean
     showDescription?: boolean
     disabled?: boolean
     readonly?: boolean
     ariaLabel?: string
     allowHalf?: boolean
   }
   
   const props = defineProps<Props>()
   
   const emit = defineEmits(['update:modelValue', 'change', 'hover'])
   
   const rating = ref(props.modelValue)
   const hoverRating = ref(0)
   
   // Computed
   const currentRating = computed(() => hoverRating.value || rating.value)
   
   // Methods
   const handleClick = (value: number) => {
     if (props.disabled || props.readonly) return
     
     const newRating = props.allowHalf && value === rating.value ? value - 0.5 : value
     rating.value = newRating
     emit('update:modelValue', newRating)
     emit('change', newRating)
   }
   
   const handleHover = (value: number) => {
     if (props.disabled || props.readonly) return
     hoverRating.value = value
     emit('hover', value)
   }
   
   const handleMouseLeave = () => {
     if (props.disabled || props.readonly) return
     hoverRating.value = 0
     emit('hover', 0)
   }
   
   const handleKeydown = (event: KeyboardEvent) => {
     if (props.disabled || props.readonly) return
     
     switch (event.key) {
       case 'ArrowRight':
       case 'ArrowUp':
         event.preventDefault()
         if (rating.value < props.maxStars) {
           handleClick(rating.value + 1)
         }
         break
       case 'ArrowLeft':
       case 'ArrowDown':
         event.preventDefault()
         if (rating.value > 0) {
           handleClick(rating.value - 1)
         }
         break
       case 'Enter':
       case ' ':
         event.preventDefault()
         break
     }
   }
   
   const getStarClass = () => {
     const baseClasses = [
       'star-button',
       'transition-all',
       'duration-200',
       'focus:outline-none',
       'focus:ring-2',
       'focus:ring-blue-500',
       'focus:ring-offset-2',
       'rounded-full'
     ]
     
     const sizeClasses = {
       sm: 'w-6 h-6',
       md: 'w-8 h-8',
       lg: 'w-10 h-10',
       xl: 'w-12 h-12'
     }
     
     const stateClasses = props.disabled 
       ? 'opacity-50 cursor-not-allowed' 
       : 'cursor-pointer hover:scale-110'
     
     return [...baseClasses, sizeClasses[props.size], stateClasses].join(' ')
   }
   
   const getStarIconClass = () => {
     const sizeClasses = {
       sm: 'w-5 h-5',
       md: 'w-6 h-6',
       lg: 'w-8 h-8',
       xl: 'w-10 h-10'
     }
     
     return sizeClasses[props.size]
   }
   
   const getRatingText = () => {
     const ratingValue = currentRating.value
     const texts = {
       0: 'امتیاز دهید',
       0.5: 'خیلی ضعیف',
       1: 'ضعیف',
       1.5: 'ضعیف تا متوسط',
       2: 'متوسط',
       2.5: 'متوسط تا خوب',
       3: 'خوب',
       3.5: 'خوب تا عالی',
       4: 'عالی',
       4.5: 'عالی تا فوق‌العاده',
       5: 'فوق‌العاده'
     }
     return texts[ratingValue] || 'امتیاز دهید'
   }
   
   const getRatingTextColor = () => {
     const ratingValue = currentRating.value
     if (ratingValue >= 4) return 'text-green-600'
     if (ratingValue >= 3) return 'text-blue-600'
     if (ratingValue >= 2) return 'text-yellow-600'
     return 'text-red-600'
   }
   
   const getRatingDescription = () => {
     const ratingValue = rating.value
     const descriptions = {
       1: 'تجربه شما با ما رضایت‌بخش نبوده است. لطفاً نظرات خود را با ما در میان بگذارید.',
       2: 'متأسفانه تجربه شما چندان رضایت‌بخش نبوده است. ما در حال بهبود خدمات خود هستیم.',
       3: 'متشکریم از نظر شما. ما در حال تلاش برای ارائه خدمات بهتر هستیم.',
       4: 'سپاسگزاریم از رضایت شما. امیدواریم همیشه بتوانیم رضایت شما را جلب کنیم.',
       5: 'بسیار متشکریم از رضایت کامل شما. این برای ما افتخار بزرگی است.'
     }
     return descriptions[Math.floor(ratingValue)] || ''
   }
   
   // Watch for prop changes
   watch(() => props.modelValue, (newValue) => {
     rating.value = newValue
   })
   
   // Expose methods and values
   defineExpose({
     rating,
     currentRating,
     setRating: (value: number) => {
       rating.value = value
       emit('update:modelValue', value)
     },
     reset: () => {
       rating.value = 0
       hoverRating.value = 0
       emit('update:modelValue', 0)
     }
   })
   </script>
   
   <style scoped>
   .interactive-star-rating {
     display: inline-flex;
     flex-direction: column;
     align-items: center;
     gap: 0.5rem;
   }
   
   .stars-container {
     display: flex;
     gap: 0.25rem;
     align-items: center;
   }
   
   .star-button {
     position: relative;
     display: flex;
     align-items: center;
     justify-content: center;
     background: none;
     border: none;
     padding: 0;
   }
   
   .star-icon {
     transition: all 0.2s ease-in-out;
   }
   
   /* Star colors based on rating */
   .star-button:nth-child(1) .star-icon { color: #fbbf24; }
   .star-button:nth-child(2) .star-icon { color: #fbbf24; }
   .star-button:nth-child(3) .star-icon { color: #fbbf24; }
   .star-button:nth-child(4) .star-icon { color: #fbbf24; }
   .star-button:nth-child(5) .star-icon { color: #fbbf24; }
   
   /* Unfilled stars */
   .star-button .star-icon {
     color: #d1d5db;
   }
   
   /* Filled stars */
   .star-button:has(.star-icon[style*="fill: currentColor"]) .star-icon,
   .star-button:hover ~ .star-button .star-icon {
     color: #d1d5db;
   }
   
   .star-button:hover .star-icon,
   .star-button:hover ~ .star-button .star-icon {
     color: #fbbf24;
   }
   
   /* Rating text animations */
   .rating-text {
     animation: fadeInUp 0.3s ease-out;
   }
   
   .rating-description {
     animation: fadeInUp 0.3s ease-out 0.1s both;
   }
   
   @keyframes fadeInUp {
     from {
       opacity: 0;
       transform: translateY(10px);
     }
     to {
       opacity: 1;
       transform: translateY(0);
     }
   }
   
   /* Pulse animation for high ratings */
   .star-button:nth-child(4) .star-icon,
   .star-button:nth-child(5) .star-icon {
     animation: pulse 2s infinite;
   }
   
   @keyframes pulse {
     0%, 100% {
       opacity: 1;
     }
     50% {
       opacity: 0.8;
     }
   }
   
   /* Focus styles for accessibility */
   .star-button:focus {
     outline: none;
     box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.5);
   }
   
   /* Disabled state */
   .star-button:disabled {
     cursor: not-allowed;
     opacity: 0.5;
   }
   
   .star-button:disabled:hover {
     transform: none;
   }
   </style>

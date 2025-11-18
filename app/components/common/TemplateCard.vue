<template>
  <div class="bg-white rounded-lg shadow-md border border-gray-200 p-1">
    <div
      :class="[
        'overflow-hidden shadow-sm rounded-md cursor-pointer border transition-all duration-200 hover:shadow-md hover:scale-105',
        variantClasses.bgGradient,
        variantClasses.borderColor
      ]"
      @click="$emit('click')"
    >
      <div class="p-3">
        <div class="flex items-center justify-between">
          <div class="flex-1">
            <dl>
              <dt :class="['text-xs font-medium truncate', variantClasses.titleColor]">{{ title }}</dt>
              <dd :class="['text-base font-bold', variantClasses.valueColor]">{{ value }}</dd>
            </dl>
          </div>
          <div class="flex-shrink-0 mr-3">
            <div :class="['w-8 h-8 rounded-lg flex items-center justify-center shadow-md', variantClasses.iconBg]">
              <slot name="icon">
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"></path>
                </svg>
              </slot>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
// کامپوننت کارت آماری با استایل‌های پیش‌فرض زیبا
// props:
// - title: عنوان کارت
// - value: مقدار آماری
// - variant: نوع رنگ (blue, purple, orange, green, red, yellow, cyan, indigo)
//
// مثال استفاده:
// <TemplateCard title="کل فایل‌ها" :value="totalFiles" variant="blue" @click="filterByType('all')">
//   <template #icon>...</template>
// </TemplateCard>

const props = defineProps({
  title: { type: String, required: true },
  value: { type: [String, Number], required: true },
  variant: { type: String, default: 'blue' }
})

// استایل‌های پیش‌فرض برای هر variant - رنگ‌های پاستیلی دقیق مثل تصویر
const variantClasses = computed(() => {
  const variants = {
    blue: {
      bgGradient: 'bg-gradient-to-br from-blue-50/90 to-blue-100/80',
      borderColor: 'border border-blue-100/60',
      titleColor: 'text-blue-600',
      valueColor: 'text-blue-800',
      iconBg: 'bg-gradient-to-br from-blue-300 to-blue-400'
    },
    purple: {
      bgGradient: 'bg-gradient-to-br from-purple-50/90 to-purple-100/80',
      borderColor: 'border border-purple-100/60',
      titleColor: 'text-purple-600',
      valueColor: 'text-purple-800',
      iconBg: 'bg-gradient-to-br from-purple-300 to-purple-400'
    },
    orange: {
      bgGradient: 'bg-gradient-to-br from-orange-50/90 to-orange-100/80',
      borderColor: 'border border-orange-100/60',
      titleColor: 'text-orange-600',
      valueColor: 'text-orange-800',
      iconBg: 'bg-gradient-to-br from-orange-300 to-orange-400'
    },
    green: {
      bgGradient: 'bg-gradient-to-br from-green-50/90 to-green-100/80',
      borderColor: 'border border-green-100/60',
      titleColor: 'text-green-600',
      valueColor: 'text-green-800',
      iconBg: 'bg-gradient-to-br from-green-300 to-green-400'
    },
    red: {
      bgGradient: 'bg-gradient-to-br from-red-50/90 to-red-100/80',
      borderColor: 'border border-red-100/60',
      titleColor: 'text-red-600',
      valueColor: 'text-red-800',
      iconBg: 'bg-gradient-to-br from-red-300 to-red-400'
    },
    yellow: {
      bgGradient: 'bg-gradient-to-br from-yellow-50/90 to-yellow-100/80',
      borderColor: 'border border-yellow-100/60',
      titleColor: 'text-yellow-600',
      valueColor: 'text-yellow-800',
      iconBg: 'bg-gradient-to-br from-yellow-300 to-yellow-400'
    },
    cyan: {
      bgGradient: 'bg-gradient-to-br from-cyan-50/90 to-cyan-100/80',
      borderColor: 'border border-cyan-100/60',
      titleColor: 'text-cyan-600',
      valueColor: 'text-cyan-800',
      iconBg: 'bg-gradient-to-br from-cyan-300 to-cyan-400'
    },
    indigo: {
      bgGradient: 'bg-gradient-to-br from-indigo-50/90 to-indigo-100/80',
      borderColor: 'border border-indigo-100/60',
      titleColor: 'text-indigo-600',
      valueColor: 'text-indigo-800',
      iconBg: 'bg-gradient-to-br from-indigo-300 to-indigo-400'
    }
  }
  
  return variants[props.variant] || variants.blue
})
</script>

<!--
توضیحات:
- این کامپوننت برای نمایش کارت آماری با استایل‌های پیش‌فرض زیبا استفاده می‌شود.
- فقط با دادن variant، تمام استایل‌ها اعمال می‌شود.
- آیکون را می‌توان با slot به نام icon جایگزین کرد.
- کدها کوتاه‌تر و استفاده آسان‌تر شده است.
--> 
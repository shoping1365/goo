<template>
  <button
    :class="[
      'inline-flex items-center font-medium rounded-lg focus:outline-none transition-all duration-200 shadow-md',
      bgGradient,
      textColor,
      borderColor,
      sizeClass,
      hoverClass,
      focusClass,
      customClass,
      { 'opacity-50 cursor-not-allowed': disabled }
    ]"
    :type="type"
    :disabled="disabled"
    @click="$emit('click', $event)"
  >
    <slot />
  </button>
</template>

<script setup lang="ts">
// کامپوننت دکمه تمپلیت با قابلیت تعیین رنگ و استایل
// props:
// - bgGradient: کلاس گرادینت یا رنگ پس‌زمینه
// - textColor: رنگ متن
// - borderColor: رنگ بردر
// - hoverClass: کلاس افکت hover
// - focusClass: کلاس افکت focus
// - size: اندازه دکمه (small, medium, large)
// - type: نوع دکمه (button, submit, ...)
//
// مثال استفاده:
// <TemplateButton bgGradient="bg-gradient-to-r from-blue-400 to-blue-600" textColor="text-white" borderColor="border border-blue-500" hoverClass="hover:from-blue-500 hover:to-blue-700" size="large">دکمه تست</TemplateButton>

defineEmits<{
  'click': [event: MouseEvent]
}>()

const props = defineProps({
  bgGradient: { type: String, default: '' },
  textColor: { type: String, default: '' },
  borderColor: { type: String, default: '' },
  hoverClass: { type: String, default: '' },
  focusClass: { type: String, default: '' },
  size: { type: String, default: 'medium' },
  type: { type: String, default: 'button' },
  customClass: { type: String, default: '' },
  disabled: { type: Boolean, default: false }
})

const sizeClass = computed(() => {
  switch (props.size) {
    case 'small':
      return 'px-3 py-2 text-sm'
    case 'large':
      return 'px-8 py-4 text-lg'
    default:
      return 'px-5 py-3 text-sm'
  }
})
</script>

<!--
توضیحات:
- این کامپوننت برای ساخت دکمه با رنگ و استایل دلخواه استفاده می‌شود.
- همه پارامترها و افکت‌ها قابل شخصی‌سازی است.
- مستندسازی کامل به فارسی انجام شده است.
--> 
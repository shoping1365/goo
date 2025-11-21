<template>
  <div class="footer-widget footer-widget--custom">
    <h4 v-if="viewModel.title" class="footer-widget__title">{{ viewModel.title }}</h4>
    <!-- 
      ⚠️ امنیت XSS: استفاده از v-html خطرناک است!
      
      این کد محتوای HTML را بدون sanitization نمایش می‌دهد که می‌تواند منجر به حملات XSS شود.
      
      ✅ راه حل صحیح:
      1. قبل از استفاده از v-html، محتوا را با کتابخانه sanitization (مثل DOMPurify) پاکسازی کنید
      2. یا از {{ }} به جای v-html استفاده کنید اگر HTML نیاز نیست
      3. محتوای کاربر را هرگز بدون sanitization در v-html قرار ندهید
      
      مثال صحیح:
      import DOMPurify from 'dompurify'
      const sanitizedHtml = computed(() => DOMPurify.sanitize(viewModel.value.html))
      <div v-html="sanitizedHtml"></div>
    -->
    <SanitizedHtml v-if="viewModel.html" :content="viewModel.html" class="footer-widget__content" />
    <slot v-else />
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import SanitizedHtml from '~/components/common/SanitizedHtml.vue'

const props = withDefaults(defineProps<{
  title?: string
  html?: string
}>(), {
  title: '',
  html: ''
})

const viewModel = computed(() => {
  const html = props.html || ''
  
  return {
    title: props.title?.trim() || '',
    html
  }
})
</script>

<style scoped>
.footer-widget {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.footer-widget__title {
  margin: 0;
  font-size: 0.95rem;
  font-weight: 600;
  color: var(--footer-widget-text-color, #111827);
}

.footer-widget__content :deep(*) {
  color: inherit;
  font-size: 0.9rem;
  line-height: 1.6;
}
</style>

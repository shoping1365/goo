<template>
  <!-- 
    ⚠️ امنیت XSS: استفاده از v-html خطرناک است!
    
    این کد محتوای HTML را بدون sanitization نمایش می‌دهد که می‌تواند منجر به حملات XSS شود.
    
    ✅ راه حل صحیح:
    1. قبل از استفاده از v-html، محتوا را با کتابخانه sanitization (مثل DOMPurify) پاکسازی کنید
    2. یا از {{ }} به جای v-html استفاده کنید اگر HTML نیاز نیست
    3. محتوای کاربر را هرگز بدون sanitization در v-html قرار ندهید
    
    مثال صحیح:
    import DOMPurify from 'dompurify'
    const sanitizedContent = computed(() => DOMPurify.sanitize(props.content))
    <div v-html="sanitizedContent"></div>
  -->
  <div class="prose prose-lg max-w-none text-right" v-html="content"></div>
</template>

<script setup lang="ts">
const props = defineProps({
  content: String
})
</script> 
<template>
  <div class="auth-layout-container">
    <!-- فقط محتوای صفحه، بدون Header/Footer -->
    <main class="auth-main-content">
      <slot />
    </main>
  </div>
</template>

<script setup lang="ts">
// Layout ساده برای صفحات احراز هویت
// بدون Header, Footer, و سایر component‌های اضافی
</script>

<style scoped>
.auth-layout-container {
  min-height: 100vh;
  width: 100%;
}

.auth-main-content {
  width: 100%;
  min-height: 100vh;
}
</style>

<template>
  <div class="section-card bg-white/95 rounded-2xl shadow-xl p-8 mb-8">
    <slot />
  </div>
</template>

<script setup>
// فقط یک wrapper ساده برای زیباسازی
</script>

<style scoped>
.section-card {
  background: rgba(255,255,255,0.97);
  border-radius: 2rem;
  box-shadow: 0 8px 40px #bdbdf544, 0 1.5px 8px #bdbdf522;
  padding: 2.5rem 2rem;
  margin-bottom: 2rem;
  min-height: 300px;
}
</style> 

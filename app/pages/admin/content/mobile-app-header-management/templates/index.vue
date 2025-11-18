<template>
  <div class="mobile-header-templates">
    <h2>انتخاب از نمونه‌های آماده</h2>
    <div class="templates-grid">
      <MinimalTemplate :header="{}" @click="selectTemplate('minimal')" />
      <BannerTemplate :header="{}" @click="selectTemplate('banner')" />
      <ClassicTemplate :header="{}" @click="selectTemplate('classic')" />
    </div>
  </div>
</template>

<script setup lang="ts">
import BannerTemplate from './banner-template.vue';
import ClassicTemplate from './classic-template.vue';
import MinimalTemplate from './minimal-template.vue';

// Emits
const emit = defineEmits<{
  templateSelected: [templateType: string]
}>()

// Methods
const selectTemplate = (templateType: string) => {
  emit('templateSelected', templateType)
}
</script>

<style scoped>
.mobile-header-templates {
  background: white;
  border-radius: 12px;
  padding: 30px;
  margin-bottom: 30px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

.mobile-header-templates h2 {
  font-size: 1.5rem;
  color: #374151;
  margin-bottom: 20px;
  text-align: center;
}

.templates-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 20px;
}

@media (max-width: 768px) {
  .templates-grid {
    grid-template-columns: 1fr;
  }
}
</style>


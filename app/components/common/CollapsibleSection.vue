<template>
  <div class="collapsible-section">
    <div class="header" @click="toggle">
      <h3>{{ title }}</h3>
      <span class="icon" :style="{ transform: isOpen ? 'rotate(180deg)' : 'rotate(0deg)' }">&#9660;</span>
    </div>
    <div v-if="isOpen" class="content">
      <slot></slot>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch, onMounted } from 'vue';

const props = defineProps({
  title: {
    type: String,
    required: true
  },
  initialOpen: {
    type: Boolean,
    default: false
  },
  persistKey: {
    type: String,
    default: ''
  }
});

const isOpen = ref(props.initialOpen);
const localStorageKey = `cs_${props.persistKey}`;

onMounted(() => {
  try {
    const saved = localStorage.getItem(localStorageKey);
    if (saved !== null) {
      isOpen.value = saved === '1';
    }
  } catch (e) {
    console.error("خطا در بارگذاری وضعیت بخش تاشو از localStorage:", e);
  }
});

watch(isOpen, (val) => {
  if (props.persistKey) {
    try {
      localStorage.setItem(localStorageKey, val ? '1' : '0');
    } catch (e) {
      console.error("خطا در ذخیره وضعیت بخش تاشو در localStorage:", e);
    }
  }
});

const toggle = () => {
  isOpen.value = !isOpen.value;
};
</script>

<style scoped>
/* استایل‌های پیش‌فرض برای استفاده عمومی */
.collapsible-section {
  border: 1px solid #444;
  margin-bottom: 10px;
  border-radius: 5px;
  overflow: hidden;
}

.header {
  background-color: #222d32;
  padding: 10px 15px;
  cursor: pointer;
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-weight: bold;
  color: white;
}

.header h3 {
  margin: 0;
  font-size: 1.1em;
  color: white;
}

.icon {
  font-size: 0.8em;
  transition: transform 0.3s ease;
  color: white;
}

.content {
  padding: 15px;
  border-top: 1px solid #444;
  background-color: #222d32;
  color: white;
}

/* استایل‌های مخصوص سایدبار ادمین */
:deep(.admin-sidebar) .collapsible-section {
  border: none;
  margin-bottom: 0;
  border-radius: 0;
  overflow: visible;
}

:deep(.admin-sidebar) .header {
  background-color: transparent;
  padding: 8px 12px;
  border: none;
  color: #d1d5db;
  font-weight: normal;
}

:deep(.admin-sidebar) .header h3 {
  font-size: 0.9em;
  color: #d1d5db;
}

:deep(.admin-sidebar) .icon {
  color: #d1d5db;
  font-size: 0.7em;
}

:deep(.admin-sidebar) .content {
  padding: 0;
  border-top: none;
  background-color: transparent;
  color: inherit;
}
</style>

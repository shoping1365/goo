<template>
  <div class="navigation-item-selector">
    <h4 class="section-title">انتخاب آیتم‌های ناوبری</h4>
    <p class="section-description">آیتم‌های مورد نظر برای نمایش در ناوبری موبایل را انتخاب کنید</p>
    
    <div class="items-grid">
      <div 
        v-for="item in navigationItems" 
        :key="item.id"
        class="item-card"
        :class="{ 'selected': selectedItems.includes(item.id) }"
        @click="toggleItem(item.id)"
      >
        <div class="item-icon">{{ item.icon }}</div>
        <div class="item-content">
          <h5 class="item-title">{{ item.title }}</h5>
          <p class="item-description">{{ item.description }}</p>
        </div>
        <div class="item-checkbox">
          <input 
            type="checkbox" 
            :checked="selectedItems.includes(item.id)"
            class="checkbox-input"
            @change="toggleItem(item.id)"
          />
        </div>
      </div>
    </div>

    <!-- نمایش آیتم‌های انتخاب شده -->
    <div v-if="selectedItems.length > 0" class="selected-items">
      <h5 class="selected-title">آیتم‌های انتخاب شده:</h5>
      <div class="selected-list">
        <span 
          v-for="itemId in selectedItems" 
          :key="itemId"
          class="selected-item"
        >
          {{ getItemById(itemId)?.title }}
        </span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'

// Props
const props = defineProps({
  modelValue: {
    type: Array,
    default: () => []
  }
})

// Emits
const emit = defineEmits(['update:modelValue'])

// Data
const selectedItems = ref([...(props.modelValue || [])])

// Static navigation items to prevent memory leaks
const navigationItems = [
  {
    id: 'home',
    title: 'خانه',
    description: 'صفحه اصلی سایت',
    icon: '🏠',
    template: 'Home',
    url: '/'
  },
  {
    id: 'categories',
    title: 'دسته‌ها',
    description: 'دسته‌بندی محصولات',
    icon: '📱',
    template: 'Categories',
    url: '/categories'
  },
  {
    id: 'cart',
    title: 'سبد خرید',
    description: 'محصولات انتخاب شده',
    icon: '🛒',
    template: 'Cart',
    url: '/cart'
  },
  {
    id: 'contact',
    title: 'تماس',
    description: 'اطلاعات تماس با ما',
    icon: '📞',
    template: 'Contact',
    url: '/contact'
  },
  {
    id: 'login',
    title: 'ورود',
    description: 'ورود به حساب کاربری',
    icon: '👤',
    template: 'Login',
    url: '/auth/login'
  },
  {
    id: 'dark-mode',
    title: 'حالت تاریک',
    description: 'تغییر تم سایت',
    icon: '🌙',
    template: 'DarkMode',
    url: '#'
  }
]

// Methods
const toggleItem = (itemId) => {
  const index = selectedItems.value.indexOf(itemId)
  if (index > -1) {
    selectedItems.value.splice(index, 1)
  } else {
    selectedItems.value.push(itemId)
  }
  emit('update:modelValue', [...selectedItems.value])
}

const getItemById = (itemId) => {
  return navigationItems.find(item => item.id === itemId)
}

// Watch for external changes - optimized to prevent memory leaks
watch(() => props.modelValue, (newValue) => {
  if (JSON.stringify(selectedItems.value) !== JSON.stringify(newValue || [])) {
    selectedItems.value = [...(newValue || [])]
  }
}, { immediate: true })
</script>

<style scoped>
.navigation-item-selector {
  background: #f9fafb;
  border-radius: 8px;
  padding: 20px;
  border: 1px solid #e5e7eb;
  margin-top: 16px;
}

.section-title {
  font-size: 1.1rem;
  font-weight: 600;
  color: #374151;
  margin: 0 0 8px 0;
  text-align: right;
}

.section-description {
  font-size: 0.9rem;
  color: #6b7280;
  margin: 0 0 20px 0;
  text-align: right;
  line-height: 1.5;
}

.items-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 12px;
  margin-bottom: 20px;
}

.item-card {
  background: white;
  border: 2px solid #e5e7eb;
  border-radius: 8px;
  padding: 16px;
  cursor: pointer;
  transition: all 0.2s ease;
  display: flex;
  align-items: center;
  gap: 12px;
}

.item-card:hover {
  border-color: #3b82f6;
  box-shadow: 0 2px 8px rgba(59, 130, 246, 0.1);
}

.item-card.selected {
  border-color: #3b82f6;
  background: #eff6ff;
}

.item-icon {
  font-size: 24px;
  width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #f3f4f6;
  border-radius: 8px;
  flex-shrink: 0;
}

.item-content {
  flex: 1;
}

.item-title {
  font-size: 1rem;
  font-weight: 600;
  color: #374151;
  margin: 0 0 4px 0;
  text-align: right;
}

.item-description {
  font-size: 0.8rem;
  color: #6b7280;
  margin: 0;
  text-align: right;
  line-height: 1.4;
}

.item-checkbox {
  flex-shrink: 0;
}

.checkbox-input {
  width: 18px;
  height: 18px;
  accent-color: #3b82f6;
  cursor: pointer;
}

.selected-items {
  background: white;
  border-radius: 8px;
  padding: 16px;
  border: 1px solid #d1d5db;
}

.selected-title {
  font-size: 0.9rem;
  font-weight: 600;
  color: #374151;
  margin: 0 0 12px 0;
  text-align: right;
}

.selected-list {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.selected-item {
  background: #3b82f6;
  color: white;
  padding: 4px 12px;
  border-radius: 16px;
  font-size: 0.8rem;
  font-weight: 500;
}

@media (max-width: 768px) {
  .items-grid {
    grid-template-columns: 1fr;
  }
  
  .item-card {
    padding: 12px;
  }
  
  .item-icon {
    width: 32px;
    height: 32px;
    font-size: 20px;
  }
}
</style>

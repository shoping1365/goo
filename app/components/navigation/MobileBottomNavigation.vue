<template>
  <!-- ناوبری پایین صفحه - فقط در موبایل و تبلت -->
  <nav v-if="mobileNavigations.length > 0" class="lg:hidden fixed bottom-0 left-0 right-0 bg-gray-100 border-t border-gray-300 z-50" dir="rtl">
    <div class="flex justify-around items-center py-2">
      <component 
        :is="getTemplateComponent(item.template)" 
        v-for="item in mobileNavigations"
        :key="item.id"
      />
    </div>
  </nav>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
// import { useRoute } from 'vue-router' // Currently unused
import Home from '~/pages/admin/content/mobile-app-navigation-management/template/Home.vue'
import Categories from '~/pages/admin/content/mobile-app-navigation-management/template/Categories.vue'
import Cart from '~/pages/admin/content/mobile-app-navigation-management/template/Cart.vue'
import Contact from '~/pages/admin/content/mobile-app-navigation-management/template/Contact.vue'
import Login from '~/pages/admin/content/mobile-app-navigation-management/template/Login.vue'
import DarkMode from '~/pages/admin/content/mobile-app-navigation-management/template/DarkMode.vue'

// const route = useRoute() // Currently unused
const mobileNavigations = ref([])
const loading = ref(false)

// تابع برای دریافت کامپوننت template
const getTemplateComponent = (templateName: string) => {
  const components = {
    'Home': Home,
    'Categories': Categories,
    'Cart': Cart,
    'Contact': Contact,
    'Login': Login,
    'DarkMode': DarkMode
  }
  return components[templateName] || Home
}

// تابع برای دریافت template بر اساس id
const getTemplateFromId = (id: string) => {
  const templateMap = {
    'home': 'Home',
    'categories': 'Categories',
    'cart': 'Cart',
    'contact': 'Contact',
    'login': 'Login',
    'dark-mode': 'DarkMode'
  }
  return templateMap[id] || 'Home'
}

// دریافت ناوبری‌های موبایل از API
const loadMobileNavigations = async () => {
  loading.value = true
  
  try {
    interface NavigationResponse {
      success: boolean
      data?: Array<{
        navigation_items?: string
        [key: string]: unknown
      }>
    }
    
    interface NavigationItem {
      id?: string | number
      type?: string
      label?: string
      icon?: string
      link?: string
      [key: string]: unknown
    }
    
    const data = await $fetch<NavigationResponse>('/api/mobile-app-navigation-settings')
    
    // استخراج آیتم‌های ناوبری از اولین ناوبری فعال
    if (data.success && data.data && data.data.length > 0) {
      const firstNavigation = data.data[0]
      if (firstNavigation.navigation_items) {
        try {
          const items = JSON.parse(firstNavigation.navigation_items) as Array<NavigationItem | string>
          mobileNavigations.value = items.map((item: NavigationItem | string, index: number) => {
            // اگر item یک string است (مثل "home", "cart")
            if (typeof item === 'string') {
              return {
                id: index + 1,
                template: getTemplateFromId(item)
              }
            }
            // اگر item یک object است
            return {
              id: index + 1,
              label: item.label || item.name || item.text || item.title || `آیتم ${index + 1}`,
              url: item.url || item.path || item.link || '/',
              icon: item.icon || item.iconName || 'home',
              template: item.template || getTemplateFromId(item.id) || 'Home'
            }
          })
        } catch (e) {
          console.error('Error parsing navigation items:', e)
        }
      }
    }
  } catch (error) {
    console.error('Error loading mobile navigations:', error)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  loadMobileNavigations()
})
</script>

<style scoped>
/* اضافه کردن padding-bottom به body فقط در موبایل و تبلت */
@media (max-width: 1023px) {
  :deep(body) {
    padding-bottom: 80px;
  }
}
</style>
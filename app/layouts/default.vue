<template>
  <div class="layout-container">
    <!-- Admin Panel Top Bar (only for authenticated users) - بالاترین جایگاه -->
    <AdminPanelTopBar v-if="isClient && !$route.path.startsWith('/admin')" />
    <Header />
    <!-- Mobile Header -->
    <MobileHeader v-if="isClient" />
    <main class="main-content">
      <slot />
    </main>
    <Footer v-if="!($route.path.startsWith('/account') || $route.path.startsWith('/checkout'))" />
    <!-- Mobile Bottom Navigation -->
    <MobileBottomNavigation v-if="isClient" />
  </div>
</template>

<script setup lang="ts">
import Header from '~/pages/admin/content/header-management/componets/Header.vue'
import Footer from '~/pages/admin/content/footer-management/componets/Footer.vue'
import AdminPanelTopBar from '~/components/admin/ui/AdminPanelTopBar.vue'
import MobileBottomNavigation from '~/components/navigation/MobileBottomNavigation.vue'
import MobileHeader from '~/components/widgets/MobileHeader.vue'
import { useRoute } from 'vue-router'
import { ref, onMounted } from 'vue'

const $route = useRoute()
const isClient = ref(false)

onMounted(() => {
  isClient.value = true
})
</script>

<style scoped>

</style> 

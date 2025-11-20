<template>
  <aside
:class="[
    'admin-sidebar bg-[#222d32] border-l border-gray-600 transform transition-transform duration-300 ease-in-out z-20',
    sidebarClasses
  ]">
    <div class="p-6 border-b border-gray-600">
      <h2 class="text-lg font-bold text-white">پنل مدیریت</h2>
    </div>

    <nav class="p-6 space-y-2 text-sm text-right">
      <NuxtLink
v-if="canShowMenu('/admin')" to="/admin" exact class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                active-class="bg-gray-700 font-semibold text-white">داشبورد</NuxtLink>

      <!-- مدیریت رسانه -->
      <CollapsibleSection
        v-if="canShowSection(['media_library.read','media_upload.read','media_image_compression.read','media_video_compression.read'])"
        title="مدیریت رسانه"
        :persist-key="'media-management'"
        :initial-open="isActivePrefix('/admin/media-management')"
      >
        <div class="collapsible-sub-links-wrapper">
          <NuxtLink
v-if="canShowMenu('/admin/media-management/library')" to="/admin/media-management/library" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                    active-class="bg-gray-700 font-semibold text-white">کتابخانه رسانه</NuxtLink>
          <NuxtLink
v-if="canShowMenu('/admin/media-management/upload')" to="/admin/media-management/upload" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                    active-class="bg-gray-700 font-semibold text-white">افزودن رسانه جدید</NuxtLink>
          <NuxtLink
v-if="canShowMenu('/admin/media-management/image-compression')" to="/admin/media-management/image-compression" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                    active-class="bg-gray-700 font-semibold text-white">تنظیمات تصاویر</NuxtLink>
        </div>
      </CollapsibleSection>

      <!-- مدیریت محتوا -->
      <CollapsibleSection
        v-if="canShowSection(['content_banners.read','content_posts.read','content_pages.read','content_menus.read','content_footer.read','content_header.read'])"
        title="مدیریت محتوا"
        :persist-key="'content-management'"
        :initial-open="isActivePrefix('/admin/content')"
      >
        <div class="collapsible-sub-links-wrapper">
          <NuxtLink
v-if="canShowMenu('/admin/content/banners')" to="/admin/content/banners" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                    active-class="bg-gray-700 font-semibold text-white">بنرها</NuxtLink>
          <NuxtLink
v-if="canShowMenu('/admin/content/posts')" to="/admin/content/posts" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                    active-class="bg-gray-700 font-semibold text-white">نوشته‌ها</NuxtLink>
          <NuxtLink
v-if="canShowMenu('/admin/content/pages')" to="/admin/content/pages" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                    active-class="bg-gray-700 font-semibold text-white">صفحات</NuxtLink>
          <NuxtLink
v-if="canShowMenu('/admin/content/menus')" to="/admin/content/menus" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                    active-class="bg-gray-700 font-semibold text-white">منوها</NuxtLink>
          <NuxtLink
v-if="canShowMenu('/admin/content/header')" to="/admin/content/header" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                    active-class="bg-gray-700 font-semibold text-white">هدر</NuxtLink>
          <NuxtLink
v-if="canShowMenu('/admin/content/footer')" to="/admin/content/footer" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                    active-class="bg-gray-700 font-semibold text-white">فوتر</NuxtLink>
        </div>
      </CollapsibleSection>

      <!-- مدیریت محصولات -->
      <CollapsibleSection
        v-if="canShowSection(['products.read','product_categories.read','product_brands.read','product_attributes.read','product_reviews.read','product_wishlists.read','product_notify_requests.read'])"
        title="مدیریت محصولات"
        :persist-key="'product-management'"
        :initial-open="isActivePrefix('/admin/product-management')"
      >
        <div class="collapsible-sub-links-wrapper">
          <NuxtLink
v-if="canShowMenu('/admin/product-management/products')" to="/admin/product-management/products" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                    active-class="bg-gray-700 font-semibold text-white">محصولات</NuxtLink>
          <NuxtLink
v-if="canShowMenu('/admin/product-management/product-categories')" to="/admin/product-management/product-categories" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                    active-class="bg-gray-700 font-semibold text-white">دسته‌بندی‌ها</NuxtLink>
          <NuxtLink
v-if="canShowMenu('/admin/product-management/brands')" to="/admin/product-management/brands" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                    active-class="bg-gray-700 font-semibold text-white">برندها</NuxtLink>
          <NuxtLink
v-if="canShowMenu('/admin/product-management/attributes')" to="/admin/product-management/attributes" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                    active-class="bg-gray-700 font-semibold text-white">ویژگی‌ها</NuxtLink>
          <NuxtLink
v-if="canShowMenu('/admin/product-management/reviews')" to="/admin/product-management/reviews" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                    active-class="bg-gray-700 font-semibold text-white">نظرات</NuxtLink>
        </div>
      </CollapsibleSection>

      <!-- سایر بخش‌ها -->
      <NuxtLink
v-if="canShowMenu('/admin/users')" to="/admin/users" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                active-class="bg-gray-700 font-semibold text-white">کاربران</NuxtLink>
      <NuxtLink
v-if="canShowMenu('/admin/transactions')" to="/admin/transactions" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                active-class="bg-gray-700 font-semibold text-white">تراکنش‌ها</NuxtLink>
      <NuxtLink
v-if="canShowMenu('/admin/notifications')" to="/admin/notifications" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                active-class="bg-gray-700 font-semibold text-white">اعلان‌ها</NuxtLink>
      <NuxtLink
v-if="canShowMenu('/admin/access-management')" to="/admin/access-management" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                active-class="bg-gray-700 font-semibold text-white">مدیریت دسترسی</NuxtLink>
      <NuxtLink
v-if="canShowMenu('/admin/settings')" to="/admin/settings" class="block px-4 py-2 rounded-md hover:bg-gray-700 text-white"
                active-class="bg-gray-700 font-semibold text-white">تنظیمات</NuxtLink>
    </nav>
  </aside>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import CollapsibleSection from '~/components/common/CollapsibleSection.vue'

// Props
interface Props {
  isOpen?: boolean
  isCollapsed?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  isOpen: true,
  isCollapsed: false
})

// Computed sidebar classes
const sidebarClasses = computed(() => {
  if (props.isCollapsed) {
    return 'w-0 md:w-16'
  }
  return props.isOpen ? 'w-64 translate-x-0' : 'translate-x-full w-0'
})

// Helper functions
const canShowMenu = (path: string) => {
  // TODO: Implement permission checking
  return true
}

const canShowSection = (permissions: string[]) => {
  // TODO: Implement section permission checking
  return true
}

const isActivePrefix = (prefix: string) => {
  if (import.meta.server) return false
  const route = useRoute()
  return route.path.startsWith(prefix)
}
</script>

<style scoped>
.admin-sidebar {
  position: fixed;
  left: 0;
  top: 3rem; /* زیر هدر */
  height: calc(100vh - 3rem);
  overflow-y: auto;
}

.collapsible-sub-links-wrapper {
  margin-right: 1rem;
  margin-top: 0.5rem;
}
</style>


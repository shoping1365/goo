<template>
  <div class="product-new-page min-h-screen p-6">
    <!-- دکمه‌های ذخیره - در بالای صفحه -->
    <div class="bg-white rounded-lg shadow-lg p-6 mb-6">
      <div class="flex justify-between items-center">
        <div class="text-xl font-bold text-gray-900">
          ویرایش محصول
        </div>
        
        <div class="flex gap-3">
          <button
            :disabled="isSaving"
            :class="[
              'inline-flex items-center px-6 py-2 rounded-lg text-white bg-gradient-to-r from-purple-500 to-indigo-600 shadow-md hover:shadow-lg hover:scale-105 transition-all duration-200 text-base font-semibold',
              isSaving 
                ? 'bg-gray-300 text-gray-500 cursor-not-allowed' 
                : ''
            ]"
            @click="previewProduct"
          >
            <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553 2.276A2 2 0 0121 14.118V17a2 2 0 01-2 2H5a2 2 0 01-2-2v-2.882a2 2 0 01.447-1.842L8 10m7 0V7a5 5 0 00-10 0v3m10 0a5 5 0 01-10 0" />
            </svg>
            پیش‌نمایش
          </button>

          <button
            :disabled="isSaving"
            :class="[
              'inline-flex items-center px-6 py-2 rounded-lg text-white bg-gradient-to-r from-blue-500 to-blue-600 shadow-md hover:shadow-lg hover:scale-105 transition-all duration-200 text-base font-semibold',
              isSaving 
                ? 'bg-gray-300 text-gray-500 cursor-not-allowed' 
                : ''
            ]"
            @click="saveAndContinueEditing"
          >
            <svg v-if="isSaving" class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            <svg v-else class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-1.414a2 2 0 00-2.828 0L11 14h3v3l-2.586-2.586z"></path>
            </svg>
            <span>{{ isSaving ? 'در حال ذخیره...' : 'ذخیره و ادامه ویرایش' }}</span>
          </button>

          <button
            :disabled="isSaving"
            :class="[
              'inline-flex items-center px-6 py-2 rounded-lg text-white bg-gradient-to-r from-green-500 to-emerald-600 shadow-md hover:shadow-lg hover:scale-105 transition-all duration-200 text-base font-semibold',
              isSaving 
                ? 'bg-gray-300 text-gray-500 cursor-not-allowed' 
                : ''
            ]"
            @click="saveProduct"
          >
            <svg v-if="isSaving" class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            <svg v-else class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
            </svg>
            {{ isSaving ? 'در حال ذخیره...' : 'ذخیره' }}
          </button>

          <!-- دکمه تنظیمات (چرخ‌دنده) -->
          <button
            type="button"
            class="px-3 py-2 rounded-lg text-sm font-medium flex items-center gap-2 shadow-lg hover:shadow-xl transition-all duration-200 bg-gradient-to-r from-pink-100 to-blue-100 text-blue-700"
            title="تنظیمات بخش‌ها"
            @click="openSettings"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" d="M12 15.5A3.5 3.5 0 1 0 12 8.5a3.5 3.5 0 0 0 0 7zm7.94-2.06a1 1 0 0 0 .26-1.09l-1-1.73a1 1 0 0 1 0-.94l1-1.73a1 1 0 0 0-.26-1.09l-2-2a1 1 0 0 0-1.09-.26l-1.73 1a1 1 0 0 1-.94 0l-1.73-1a1 1 0 0 0-1.09.26l-2 2a1 1 0 0 0-.26 1.09l1 1.73a1 1 0 0 1 0 .94l-1 1.73a1 1 0 0 0 .26 1.09l2 2a1 1 0 0 0 1.09.26l1.73-1a1 1 0 0 1 .94 0l1.73 1a1 1 0 0 0 1.09-.26l2-2z" />
            </svg>
          </button>
        </div>
      </div>
    </div>

    <!-- TopBar -->
    <div class="flex items-center justify-between bg-gray-50 px-6 py-4 mb-4 rounded-t shadow-sm">
      <div class="flex items-center gap-6">
        <h1 class="text-base font-normal text-gray-600">ویرایش محصول</h1>
      </div>
      <div class="flex items-center gap-6">
        <NuxtLink to="/admin/product-management/products" class="text-blue-600 hover:underline text-xs">بازگشت به لیست محصولات</NuxtLink>
      </div>
    </div>

    <!-- Tabs -->
    <div class="bg-gray-50 px-4 pt-2 flex gap-2 overflow-x-auto mb-6">
      <a 
        v-for="tab in tabs" 
        :key="tab.value" 
        href="#"
        :class="[
          'px-4 py-2 rounded-t text-sm transition cursor-pointer',
          activeTab === tab.value
            ? 'font-bold text-blue-700' 
            : 'hover:bg-gray-50 text-gray-600'
        ]"
        @click.prevent="setActiveTab(tab.value)"
      >
        {{ tab.label }}
      </a>
    </div>

    <!-- Tab Content -->
    <div class="tab-content">
      <template v-if="activeTab === 'images'">
        <ProductImagesTab default-media-category="products" />
      </template>
      <template v-else>
        <component :is="currentTabComponent" />
      </template>
    </div>

    <!-- Global Delete Confirm Modal for this page -->
    <DeleteConfirmModal ref="deleteModalRef" />
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
declare const navigateTo: (to: string) => Promise<void>
</script>

<script setup lang="ts">
import { computed, defineAsyncComponent, onMounted, onUnmounted, provide, ref } from 'vue'
import { useRoute } from 'vue-router'
import DeleteConfirmModal from '~/components/common/DeleteConfirmModal.vue'
import ProductImagesTab from '~/components/product/ProductImagesTab.vue'
import ProductSeoTab from '~/components/product/ProductSeoTab.vue'
import ProductVideoTab from '~/components/product/ProductVideoTab.vue'
import { useDeleteConfirmModal } from '~/composables/useDeleteConfirmModal'
import { useNotifier } from '~/composables/useNotifier'
import { useProductLink } from '~/composables/useProductLink'
import { useProductCreateStore } from '~/stores/productCreate'

// Meta
definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// Store
const pStore = useProductCreateStore()
const isSaving = computed(() => pStore.isSaving)
const { buildProductLink } = useProductLink()
// Delete confirm modal ref (project standard)
const { deleteModalRef } = useDeleteConfirmModal()

// Router for navigation
// const router = useRouter()

// Tabs Management
const activeTab = ref('info')

// تنظیمات بخش‌ها
const sectionSettings = ref({
  // بخش‌های اطلاعات محصول
  mainInfo: true,
  technicalInfo: true,
  displaySettings: true,
  scheduling: true,
  management: true,
  strengthsWeaknesses: true,
  
  // سایر تب‌ها
  pricing: true,
  inventory: true,
  shipping: true,
  images: true,
  variants: true,
  specs: true,
  faq: true,
  seo: true,
  related: true,
  complements: true,
  video: true
})

// بارگذاری تنظیمات از localStorage
function loadSectionSettings() {
  if (typeof window !== 'undefined') {
    const savedSettings = localStorage.getItem('productPageSections')
    if (savedSettings) {
      try {
        const parsed = JSON.parse(savedSettings)
        // Backward-compat: migrate related -> complements
        if (parsed && parsed.related !== undefined && parsed.complements === undefined) {
          parsed.complements = parsed.related
        }
        sectionSettings.value = { ...sectionSettings.value, ...parsed }
      } catch (error) {
        console.error('خطا در بارگذاری تنظیمات:', error)
      }
    }
  }
}

// فیلتر کردن تب‌ها بر اساس تنظیمات
const tabs = computed(() => {
  const allTabs = [
    { label: 'اطلاعات محصول', value: 'info' },
    { label: 'قیمت گذاری', value: 'pricing' },
    { label: 'مدیریت موجودی', value: 'inventory' },
    { label: 'حمل و نقل', value: 'shipping' },
    { label: 'مدیریت تصاویر', value: 'images' },
    { label: 'مدیریت تنوع کالایی', value: 'variants' },
    { label: 'مشخصات فنی محصول', value: 'specs' },
    { label: 'سوالات متداول', value: 'faq' },
    { label: 'سئو', value: 'seo' },
    { label: 'محصولات مکمل', value: 'complements' },
    { label: 'ویدیو', value: 'video' },
  ]
  
  return allTabs.filter(tab => {
    if (tab.value === 'info') {
      // برای تب اطلاعات محصول، حداقل یکی از بخش‌ها باید فعال باشد
      return sectionSettings.value.mainInfo || 
             sectionSettings.value.technicalInfo || 
             sectionSettings.value.displaySettings || 
             sectionSettings.value.scheduling || 
             sectionSettings.value.management || 
             sectionSettings.value.strengthsWeaknesses
    }
    return sectionSettings.value[tab.value]
  })
})

const tabComponents = {
  info: defineAsyncComponent(() => import('./info.vue')),
  pricing: defineAsyncComponent(() => import('./pricing.vue')),
  shipping: defineAsyncComponent(() => import('./shipping.vue')),
  inventory: defineAsyncComponent(() => import('./inventory.vue')),
  images: ProductImagesTab,
  variants: defineAsyncComponent(() => import('./variants.vue')),
  specs: defineAsyncComponent(() => import('./specs.vue')),
  faq: defineAsyncComponent(() => import('./faq.vue')),
  seo: ProductSeoTab,
  complements: defineAsyncComponent(() => import('./complements.vue')),
  related: defineAsyncComponent(() => import('./complements.vue')),
  video: ProductVideoTab
}

const currentTabComponent = computed(() => tabComponents[activeTab.value] || tabComponents.info)

const setActiveTab = (tabValue) => {
  activeTab.value = tabValue
}

// Notifier
const notifier = useNotifier()

// Lifecycle Hooks
onMounted(async () => {
  // بارگذاری تنظیمات بخش‌ها
  loadSectionSettings()

  const route = useRoute()
  const routeId = route.query.id as string | undefined

  if (routeId) {
    await pStore.loadCategories()
    await pStore.loadProductForEdit(routeId)
  } else {
    // اگر شناسه‌ای در query نبود، به لیست محصولات برگرد
    navigateTo('/admin/product-management/products')
    return
  }

  activeTab.value = 'info'
})

onUnmounted(() => {
  // Clean up when the user navigates away from the page.
  // This ensures that when they come back to "Add New", it's a fresh start.
  // فرانت حق نگهداری شناسه ندارد
  // Reset the store to avoid data leaks into other parts of the app
  pStore.$reset();
});

// Save Functions
async function saveProduct() {
  try {
    // If we are in "edit mode" (because of save & continue), we should update.
    // Otherwise, we create a new product.
    if (pStore.isEditMode && pStore.editingProductId) {
      const product = await pStore.updateProduct(pStore.editingProductId);
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      if (product && (product as any).id) {
        // نمایش پیام موفقیت با notifier
        notifier.success('تغییرات با موفقیت ذخیره شد')
        navigateTo('/admin/product-management/products')
      }
    } else {
      const product = await pStore.createProduct()
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      if (product && (product as any).id) {
        notifier.success('محصول با موفقیت ایجاد شد')
        navigateTo('/admin/product-management/products')
      }
    }
  } catch (error) {
    console.error('Error saving product:', error)
    interface ErrorWithResponse {
      response?: unknown
      skipToast?: boolean
      message?: string
    }
    const err = error as ErrorWithResponse
    if (err?.response || err?.skipToast) return
    notifier.error(err?.message || 'خطا در ذخیره محصول', 'خطا')
  }
}

async function saveAndContinueEditing() {
  try {
    const product = await pStore.saveAndContinueEditing()
    if (product && product.id) {
      notifier.success('محصول با موفقیت ذخیره و آماده ویرایش است')
      // Instead of redirecting, we just stay on the page.
      // The store is already updated with the new product's data and ID.
      // The page is now effectively in "edit" mode for the new product.
    }
  } catch (error) {
    console.error('Error saving and continuing:', error)
    interface ErrorWithResponse {
      response?: unknown
      skipToast?: boolean
      message?: string
    }
    const err = error as ErrorWithResponse
    if (err?.response || err?.skipToast) return
    notifier.error(err?.message || 'خطا در ذخیره محصول', 'خطا')
  }
}

async function previewProduct() {
  try {
    const productId = pStore.editingProductId
    if (!productId) {
      notifier.warning('شناسه محصول یافت نشد')
      return
    }
    const sku = (pStore.productForm.sku || '').toString().trim()
    const slug = (pStore.productForm.slug || '').toString().trim()
    const link = sku
      ? buildProductLink({ sku, slug })
      : buildProductLink(String(productId))
    window.open(link, '_blank')
  } catch (error) {
    notifier.error('خطا در پیش‌نمایش محصول')
  }
}

// Settings Modal
function openSettings() {
  import('~/components/common/SectionSettingsModal.vue').then(async mod => {
    const { createApp } = await import('vue')
    const container = document.createElement('div')
    document.body.appendChild(container)

    const app = createApp(mod.default, {
      show: true,
      sections: sectionSettings.value,
      sectionLabels: {},
      sectionGroups: [],
      'onUpdate:show': (val) => {
        if (!val) {
          app.unmount()
          container.remove()
        }
      },
      'onSections-updated': (newSettings) => {
        // به‌روزرسانی تنظیمات
        sectionSettings.value = { ...sectionSettings.value, ...newSettings }
        
        // بررسی اینکه تب فعال هنوز موجود است
        const availableTabs = tabs.value.map(tab => tab.value)
        if (!availableTabs.includes(activeTab.value)) {
          // اگر تب فعال حذف شده، به اولین تب موجود برو
          activeTab.value = availableTabs[0] || 'info'
        }
      }
    })
    app.mount(container)
  })
}

// Provide section settings to child components
provide('sectionSettings', sectionSettings)

// Provide product name to SEO tab
provide('pageTitle', computed(() => pStore.productForm.name))
</script>

<style>
/* Keep inputs blue on focus */
.product-new-page input:focus,
.product-new-page textarea:focus,
.product-new-page select:focus {
  border-color: #60a5fa;
  outline: none;
  box-shadow: 0 0 0 1px #60a5fa;
}
</style> 

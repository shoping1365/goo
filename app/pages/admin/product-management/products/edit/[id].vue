<template>
  <div class="product-new-page min-h-screen p-6">
    <div class="bg-white rounded-lg shadow-lg p-6 mb-6">
      <div class="flex justify-between items-center">
        <div class="text-xl font-bold text-gray-900">ویرایش محصول</div>
        <div class="flex gap-3">
          <button :disabled="isSaving" class="inline-flex items-center px-6 py-2 rounded-lg text-white bg-gradient-to-r from-purple-500 to-indigo-600 shadow-md hover:shadow-lg hover:scale-105 transition-all duration-200 text-base font-semibold" @click="previewProduct">
            پیش‌نمایش
          </button>
          <button :disabled="isSaving" class="inline-flex items-center px-6 py-2 rounded-lg text-white bg-gradient-to-r from-blue-500 to-blue-600 shadow-md hover:shadow-lg hover:scale-105 transition-all duration-200 text-base font-semibold" @click="saveAndContinueEditing">
            <span>{{ isSaving ? 'در حال ذخیره...' : 'ذخیره و ادامه ویرایش' }}</span>
          </button>
          <button :disabled="isSaving" class="inline-flex items-center px-6 py-2 rounded-lg text-white bg-gradient-to-r from-green-500 to-emerald-600 shadow-md hover:shadow-lg hover:scale-105 transition-all duration-200 text-base font-semibold" @click="saveProduct">
            {{ isSaving ? 'در حال ذخیره...' : 'ذخیره' }}
          </button>
        </div>
      </div>
    </div>

    <div class="flex items-center justify-between bg-gray-50 px-6 py-4 mb-4 rounded-t shadow-sm">
      <div class="flex items-center gap-6">
        <h1 class="text-base font-normal text-gray-600">ویرایش محصول</h1>
      </div>
      <div class="flex items-center gap-6">
        <NuxtLink to="/admin/product-management/products" class="text-blue-600 hover:underline text-xs">بازگشت به لیست محصولات</NuxtLink>
      </div>
    </div>

    <div class="bg-gray-50 px-4 pt-2 flex gap-2 overflow-x-auto mb-6">
      <a
v-for="tab in tabs" :key="tab.value" href="#" :class="['px-4 py-2 rounded-t text-sm transition cursor-pointer', activeTab === tab.value ? 'font-bold text-blue-700' : 'hover:bg-gray-50 text-gray-600']"
         @click.prevent="setActiveTab(tab.value)">
        {{ tab.label }}
      </a>
    </div>

    <div class="tab-content">
      <template v-if="activeTab === 'images'">
        <ProductImagesTab default-media-category="products" />
      </template>
      <template v-else-if="activeTab === 'seo'">
        <!-- تب SEO فقط وقتی محصول لود شده نمایش داده شود -->
        <ProductSeoTab
          v-if="pStore.editingProductId && pStore.productForm.name && pStore.productForm.sku"
          v-bind="currentTabProps"
          @update:slug="(v:string)=> pStore.productForm.slug = v"
        />
        <div v-else class="flex items-center justify-center py-12">
          <div class="text-center">
            <svg class="w-8 h-8 animate-spin text-blue-500 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
            </svg>
            <p class="text-gray-500">در حال بارگذاری اطلاعات محصول...</p>
            <p class="text-xs text-gray-400 mt-2">
              editingProductId: {{ pStore.editingProductId }}<br>
              name: {{ pStore.productForm.name }}<br>
              sku: {{ pStore.productForm.sku }}
            </p>
          </div>
        </div>
      </template>
      <template v-else>
        <component
          :is="currentTabComponent"
          v-bind="currentTabProps"
          @update:slug="(v:string)=> pStore.productForm.slug = v"
        />
      </template>
    </div>
  </div>
</template>

<script setup lang="ts">
// @ts-ignore: Nuxt macro
definePageMeta({ layout: 'admin-main' })

import { computed, defineAsyncComponent, onMounted, onUnmounted, provide, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import ProductImagesTab from '~/components/product/ProductImagesTab.vue'
import ProductSeoTab from '~/components/product/ProductSeoTab.vue'
import ProductVideoTab from '~/components/product/ProductVideoTab.vue'
import { useProductLink } from '~/composables/useProductLink'
import { useProductCreateStore } from '~/stores/productCreate'

const pStore = useProductCreateStore()
const isSaving = computed(() => pStore.isSaving)
const { buildProductLink } = useProductLink()

const router = useRouter()
const route = useRoute()

const activeTab = ref('info')

// فوراً productStore را provide کن
provide('productStore', pStore)

const sectionSettings = ref({
  mainInfo: true,
  technicalInfo: true,
  displaySettings: true,
  scheduling: true,
  management: true,
  strengthsWeaknesses: true,
  pricing: true,
  inventory: true,
  shipping: true,
  images: true,
  variants: true,
  specs: true,
  faq: true,
  seo: true,
  related: true,
  video: true
})

function loadSectionSettings() {
  if (typeof window !== 'undefined') {
    const saved = localStorage.getItem('productPageSections')
    if (saved) {
      try { sectionSettings.value = { ...sectionSettings.value, ...JSON.parse(saved) } } catch {}
    }
  }
}

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
      return sectionSettings.value.mainInfo || sectionSettings.value.technicalInfo || sectionSettings.value.displaySettings || sectionSettings.value.scheduling || sectionSettings.value.management || sectionSettings.value.strengthsWeaknesses
    }
    return sectionSettings.value[tab.value as keyof typeof sectionSettings.value]
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
  video: ProductVideoTab
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const currentTabComponent = computed(() => (tabComponents as any)[activeTab.value] || (tabComponents as any).info)
const currentTabProps = computed(() => {
  if (activeTab.value === 'seo') {
    return {
      slug: pStore.productForm.slug,
      defaultTitle: pStore.productForm.name,
      metaTitle: pStore.productForm.seo_title || pStore.productForm.name,
      initialMetaDescription: pStore.productForm.meta_description,
      url: pStore.productForm.url
    }
  }
  return {}
})
const setActiveTab = (tabValue: string) => { activeTab.value = tabValue }

onMounted(async () => {
  loadSectionSettings()
  
  const id = route.params.id as string | undefined
  if (!id) {
    // اگر شناسه‌ای در URL وجود ندارد به لیست محصولات برگرد
    router.push('/admin/product-management/products')
    return
  }
  await pStore.loadCategories()
  await pStore.loadBrands()
  await pStore.loadProductForEdit(id)
  
  activeTab.value = 'info'
})

onUnmounted(() => {
  pStore.$reset()
})

async function saveProduct() {
  try {
    if (pStore.isEditMode && pStore.editingProductId) {
      const product = await pStore.updateProduct(pStore.editingProductId)
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      if (product && (product as any).id) {
        alert('تغییرات با موفقیت ذخیره شد')
        router.push('/admin/product-management/products')
      }
    } else {
      const product = await pStore.createProduct()
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      if (product && (product as any).id) {
        alert('محصول با موفقیت ایجاد شد')
        router.push('/admin/product-management/products')
      }
    }
  } catch {
    alert('خطا در ذخیره محصول')
  }
}

async function saveAndContinueEditing() {
  try {
    const product = await pStore.saveAndContinueEditing()
    if (product && product.id) {
      alert('محصول با موفقیت ذخیره و آماده ویرایش است')
    }
  } catch {
    alert('خطا در ذخیره محصول')
  }
}

async function previewProduct() {
  try {
    const productId = pStore.editingProductId
    if (!productId) return
    const sku = (pStore.productForm.sku || '').toString().trim()
    const slug = (pStore.productForm.slug || '').toString().trim()
    const link = sku ? buildProductLink({ sku, slug }) : buildProductLink(String(productId))
    window.open(link, '_blank')
  } catch {}
}

provide('sectionSettings', sectionSettings)
provide('pageTitle', computed(() => pStore.productForm.name))
</script>

<style>
.product-new-page input:focus,
.product-new-page textarea:focus,
.product-new-page select:focus {
  border-color: #60a5fa;
  outline: none;
  box-shadow: 0 0 0 1px #60a5fa;
}
</style>


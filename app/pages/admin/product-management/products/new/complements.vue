<template>
  <!-- محصولات مرتبط و مکمل -->
  <div class="bg-gradient-to-br from-emerald-50 to-blue-50 rounded-2xl border border-emerald-200 shadow-lg overflow-hidden mb-8">
    <div class="bg-gradient-to-r from-emerald-600 to-blue-600 px-6 py-4 cursor-pointer" @click="toggleSection('relatedProducts')">
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-3">
          <div class="p-2 bg-white/20 rounded-xl">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              
            </svg>
          </div>
          <h3 class="text-xl font-bold text-white">محصولات مرتبط و مکمل</h3>
        </div>
        <div class="p-2 bg-white/20 rounded-lg">
          <span class="text-white font-bold text-lg">{{ sections.relatedProducts ? '−' : '+' }}</span>
        </div>
      </div>
    </div>

    <div v-show="sections.relatedProducts" class="p-8">
      <div class="grid grid-cols-1 gap-6">
        <!-- جستجو و افزودن محصول -->
        <div class="bg-white rounded-xl shadow-sm border border-emerald-100 p-6">
          <label class="flex items-center gap-2 text-lg font-semibold text-gray-900 mb-6">
            <div class="p-2 bg-emerald-100 rounded-lg">
              <svg class="w-5 h-5 text-emerald-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
              </svg>
            </div>
            جستجو و افزودن محصول مرتبط
          </label>

          <div class="grid grid-cols-1 md:grid-cols-12 gap-6">
            <div class="md:col-span-8">
              <input
v-model="q"
                type="text"
                class="w-full px-4 py-3 bg-white border border-gray-300 rounded-lg focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500 text-gray-900 transition-all duration-200"
                dir="rtl"
                placeholder="نام محصول، کد محصول یا دسته‌بندی را جستجو کنید..."
                @input="searchProducts" />
            </div>
            <div class="md:col-span-2">
              <select class="w-full px-4 py-3 bg-white border border-gray-300 rounded-lg focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500 text-gray-900 transition-all duration-200 appearance-none">
                <option value="all">همه دسته‌ها</option>
                <option value="electronics">الکترونیک</option>
                <option value="clothing">پوشاک</option>
                <option value="home">خانه و آشپزخانه</option>
              </select>
            </div>
            <div class="md:col-span-2">
              <button class="w-full bg-gradient-to-r from-emerald-600 to-blue-600 text-white rounded-lg px-6 py-3 font-semibold hover:from-emerald-700 hover:to-blue-700 transition-all duration-200 shadow-md hover:shadow-lg">
                جستجو
              </button>
            </div>
          </div>
        </div>

        <!-- نتایج جستجو -->
        <div class="bg-white rounded-xl shadow-sm border border-blue-100 p-6">
          <h4 class="text-lg font-semibold text-gray-800 mb-6">نتایج جستجو</h4>
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            <!-- کارت محصول نمونه -->
            <div v-for="product in results" :key="product.id" class="border border-gray-200 rounded-lg p-6 hover:shadow-md transition-all duration-200 hover:border-emerald-300">
              <div class="flex gap-3">
                <div class="w-16 h-16 bg-gradient-to-br from-gray-100 to-gray-200 rounded-lg flex items-center justify-center flex-shrink-0 overflow-hidden">
                  <img
                    v-if="product.thumbnail || product.image || product.main_image || (product.images && (product.images[0] && ('thumbnail' in product.images[0] ? product.images[0].thumbnail : product.images[0].image_url)))"
                    :src="String(product.thumbnail || product.image || product.main_image || (product.images && product.images[0] && ('thumbnail' in product.images[0] ? product.images[0].thumbnail : product.images[0].image_url)) || '')"
                    alt=""
                    class="w-full h-full object-cover"
                    @click="openPreview(product)"
                  />
                  <svg v-else class="w-6 h-6 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                  </svg>
                </div>
                <div class="flex-1 min-w-0">
                  <h5 class="text-sm font-medium text-gray-900 truncate">{{ product.name }}</h5>
                  <p class="text-xs text-gray-500 mb-2">دسته: {{ product.category }}</p>
                  <div class="flex items-center justify-between">
                    <span class="text-sm font-bold text-emerald-600">{{ product.price }} تومان</span>
                    <button class="text-xs bg-emerald-100 text-emerald-700 px-2 py-1 rounded-full hover:bg-emerald-200 transition-colors" @click="addRelated(product.id)">
                      افزودن
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- لیست محصولات انتخاب شده -->
        <div class="bg-white rounded-xl shadow-sm border border-purple-100 p-6">
          <div class="flex items-center justify-between mb-6">
            <h4 class="text-lg font-semibold text-gray-800">محصولات مرتبط انتخاب شده</h4>
            <span class="bg-purple-100 text-purple-800 px-3 py-1 rounded-full text-sm font-medium">{{ selected.length }} محصول</span>
          </div>

          <div class="overflow-hidden">
            <div class="space-y-4">
              <!-- آیتم محصول -->
              <div v-for="item in selected" :key="item.id" class="flex items-center gap-6 p-6 bg-gray-50 rounded-lg hover:bg-gray-100 transition-colors">
                <div class="w-12 h-12 bg-gradient-to-br from-purple-100 to-purple-200 rounded-lg flex items-center justify-center overflow-hidden">
                  <img
                    v-if="item.thumbnail || item.image_url || item.image || item.main_image"
                    :src="String(item.thumbnail || item.image_url || item.image || item.main_image || '')"
                    alt=""
                    class="w-full h-full object-cover"
                    @click="openPreview(item)"
                  />
                  <svg v-else class="w-5 h-5 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    
                  </svg>
                </div>
                <div class="flex-1">
                  <h5 class="font-medium text-gray-900">{{ item.name }}</h5>
                  <p class="text-sm text-gray-500">دسته: {{ item.category }}</p>
                </div>
                <div class="text-sm font-semibold text-purple-600">{{ item.price }} تومان</div>
                <div class="flex items-center gap-2">
                  <label class="text-xs text-gray-600">اولویت:</label>
                  <input type="number" class="w-16 px-2 py-1 border border-gray-300 rounded text-center text-sm focus:ring-2 focus:ring-purple-500 focus:border-purple-500" min="1" :value="item.priority" @input="item.priority = ($event.target as HTMLInputElement).value" />
                </div>
                <button class="p-2 text-red-500 hover:text-red-700 hover:bg-red-50 rounded-lg transition-colors" @click="removeRelated(item.id)">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <!-- Preview modal mount -->
  <ImagePreviewModal v-model="showPreview" :image="previewImage" />
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
</script>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import ImagePreviewModal from '~/components/media/ImagePreviewModal.vue'
import { useProductCreateStore } from '~/stores/productCreate'

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

interface ComplementProduct {
  id: number
  name: string
  sku?: string
  price?: number
  thumbnail?: string | null
  images?: { image_url: string }[]
  image?: string
  main_image?: string
  [key: string]: unknown
}

const store = useProductCreateStore()
const q = ref('')
const results = ref<ComplementProduct[]>([])
const selected = ref<ComplementProduct[]>([])

// Sections state
const sections = ref({
  relatedProducts: true
})

const toggleSection = (sectionName: string) => {
  if (sections.value.hasOwnProperty(sectionName)) {
    sections.value[sectionName] = !sections.value[sectionName]
  }
}

function toThumbnail(url?: string | null){
  if (!url) return null
  const dotIdx = url.lastIndexOf('.')
  if (dotIdx === -1) return url + '_thumbnail'
  return url.slice(0, dotIdx) + '_thumbnail' + url.slice(dotIdx)
}

async function searchProducts(){
  if (!q.value.trim()) { 
    results.value = []; 
    return 
  }
  try {
    const res = await $fetch<ComplementProduct[] | { data: ComplementProduct[] }>(`/api/admin/products?search=${encodeURIComponent(q.value)}&limit=10`)
    const arr = Array.isArray(res) ? res : (res?.data || [])
    results.value = arr.map((p) => {
      const firstImage = (p?.images && p.images.length > 0 && p.images[0]?.image_url) || p.image || p.main_image || p.thumbnail
      const thumb = toThumbnail(firstImage as string)
      return {
        id: p.id,
        name: p.name,
        category: (p?.category && typeof p.category === 'object' && 'name' in p.category ? String(p.category.name) : '') || '',
        price: p.price,
        sku: p.sku,
        slug: p.slug,
        image: firstImage,
        thumbnail: thumb || firstImage,
      }
    })
  } catch(e){ 
    console.error('Error searching products:', e)
    results.value = [] 
  }
}

async function addRelated(productId:number){
  if (!store.isEditMode || !store.editingProductId) {
    console.warn('Cannot add related product: not in edit mode')
    return
  }
  try {
    await $fetch(`/api/product-complements/${store.editingProductId}`, { 
      method:'POST', 
      body:{ complement_product_id: productId } 
    })
    await refreshSelected()
  } catch { 
    // Error adding related product
  }
}

async function removeRelated(productId:number){
  if (!store.isEditMode || !store.editingProductId) {
    console.warn('Cannot remove related product: not in edit mode')
    return
  }
  try {
    await $fetch(`/api/product-complements/${store.editingProductId}/${productId}`, { method:'DELETE' })
    await refreshSelected()
  } catch { 
    // Error removing related product
  }
}

async function refreshSelected(){
  if (!store.isEditMode || !store.editingProductId) {
    selected.value = []
    return
  }
  try {
    interface ComplementItem {
      id?: number | string
      image_url?: string
      image?: string
      main_image?: string
      thumbnail?: string
      [key: string]: unknown
    }
    interface ComplementsResponse {
      data?: ComplementItem[]
      [key: string]: unknown
    }
    const res = await $fetch<ComplementItem[] | ComplementsResponse>(`/api/product-complements/${store.editingProductId}`)
    const data = Array.isArray(res) ? res : (res?.data ?? [])
    const arr = Array.isArray(data) ? data : []
    selected.value = arr.map((it: ComplementItem) => {
      const base = it.image_url || it.image || it.main_image || it.thumbnail
      const thumb = toThumbnail(base as string | null)
      return { 
        id: it.id ?? 0,
        name: String(it.name ?? ''),
        category: (it.category && typeof it.category === 'object' && 'name' in it.category ? String(it.category.name) : String(it.category ?? '')),
        price: Number(it.price ?? 0),
        sku: it.sku ? String(it.sku) : undefined,
        slug: it.slug ? String(it.slug) : undefined,
        image: base ? String(base) : undefined,
        main_image: it.main_image ? String(it.main_image) : undefined,
        thumbnail: thumb || base ? String(thumb || base) : undefined,
        images: it.images
      } as ComplementProduct
    })
  } catch { selected.value = [] }
}

// Preview modal state
const showPreview = ref(false)
// eslint-disable-next-line @typescript-eslint/no-explicit-any
const previewImage = ref<any>(null)
// eslint-disable-next-line @typescript-eslint/no-explicit-any
function getOriginalUrl(obj:any){
  if(!obj) return null
  if(obj.images && obj.images[0]?.image_url) return obj.images[0].image_url
  return obj.image_url || obj.image || obj.main_image || obj.thumbnail || null
}
// eslint-disable-next-line @typescript-eslint/no-explicit-any
function openPreview(obj:any){
  const url = getOriginalUrl(obj)
  if(!url) return
  previewImage.value = { url, name: obj?.name || '' }
  showPreview.value = true
}

// Load selected products on mount
onMounted(() => {
  refreshSelected()
})
</script>

<style scoped>
/* Custom Select Arrow */
select {
  background-image: url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' fill='none' viewBox='0 0 20 20'%3e%3cpath stroke='%236b7280' stroke-linecap='round' stroke-linejoin='round' stroke-width='1.5' d='m6 8 4 4 4-4'/%3e%3c/svg%3e");
  background-position: right 0.5rem center;
  background-repeat: no-repeat;
  background-size: 1.5em 1.5em;
  padding-right: 2.5rem;
}

/* Hover transitions */
.transition-all {
  transition: all 0.3s ease;
}

/* Custom focus ring */
.focus\:ring-2:focus {
  --tw-ring-offset-shadow: var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color);
  --tw-ring-shadow: var(--tw-ring-inset) 0 0 0 calc(2px + var(--tw-ring-offset-width)) var(--tw-ring-color);
  box-shadow: var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000);
}
</style>


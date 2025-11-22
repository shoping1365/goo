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
              <select
                v-model="selectedCategory"
                class="w-full px-4 py-3 bg-white border border-gray-300 rounded-lg focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500 text-gray-900 transition-all duration-200 appearance-none"
                @change="searchProducts"
              >
                <option :value="''">همه دسته‌ها</option>
                <option v-for="cat in categories" :key="cat.id" :value="cat.id">{{ cat.name }}</option>
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
            <div v-for="product in results" :key="product.id" class="border border-gray-200 rounded-lg p-6 hover:shadow-md transition-all duration-200 hover:border-emerald-300">
              <div class="flex gap-3">
                <div class="w-16 h-16 bg-gradient-to-br from-gray-100 to-gray-200 rounded-lg flex items-center justify-center flex-shrink-0 overflow-hidden">
                  <img
                    v-if="product.thumbnail || product.image || product.main_image || (product.images && (product.images[0]?.thumbnail || product.images[0]?.image_url))"
                    :src="product.thumbnail || product.image || product.main_image || product.images[0]?.thumbnail || product.images[0]?.image_url"
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
                    <span class="text-sm font-bold text-emerald-600">{{ new Intl.NumberFormat('fa-IR').format(product.price) }} تومان</span>
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
              <div v-for="item in selected" :key="item.id" class="flex items-center gap-6 p-6 bg-gray-50 rounded-lg hover:bg-gray-100 transition-colors">
                <div class="w-12 h-12 bg-gradient-to-br from-purple-100 to-purple-200 rounded-lg flex items-center justify-center overflow-hidden">
                  <img
                    v-if="item.thumbnail || item.image_url || item.image || item.main_image"
                    :src="item.thumbnail || item.image_url || item.image || item.main_image"
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
                <div class="text-sm font-semibold text-purple-600">{{ new Intl.NumberFormat('fa-IR').format(item.price) }} تومان</div>
                <div class="flex items-center gap-2">
                  <label class="text-xs text-gray-600">اولویت:</label>
                  <input type="number" class="w-16 px-2 py-1 border border-gray-300 rounded text-center text-sm focus:ring-2 focus:ring-purple-500 focus:border-purple-500" min="1" :value="item.priority || 1" @input="onPriorityInput(item, $event)" />
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

  <!-- تنظیمات نمایش -->
  <div class="bg-gradient-to-br from-orange-50 to-pink-50 rounded-2xl border border-orange-200 shadow-lg overflow-hidden">
    <div class="bg-gradient-to-r from-orange-600 to-pink-600 px-6 py-4 cursor-pointer" @click="toggleSection('relatedSettings')">
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-3">
          <div class="p-2 bg-white/20 rounded-xl">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6V4m0 2a2 2 0 100 4m0-4a2 2 0 110 4m-6 8a2 2 0 100-4m0 4a2 2 0 100 4m0-4v2m0-6V4m6 6v10m6-2a2 2 0 100-4m0 4a2 2 0 100 4m0-4v2m0-6V4"></path>
            </svg>
          </div>
          <h3 class="text-xl font-bold text-white">تنظیمات نمایش محصولات مرتبط</h3>
        </div>
        <div class="p-2 bg-white/20 rounded-lg">
          <span class="text-white font-bold text-lg">{{ sections.relatedSettings ? '−' : '+' }}</span>
        </div>
      </div>
    </div>
    
    <div v-show="sections.relatedSettings" class="p-8">
      <div class="grid grid-cols-1 gap-6">
        <!-- تنظیمات اصلی -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div class="bg-white rounded-xl shadow-sm border border-orange-100 p-6">
            <label class="flex items-center gap-2 text-sm font-semibold text-gray-900 mb-4">
              <div class="p-2 bg-orange-100 rounded-lg">
                <svg class="w-4 h-4 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 20l4-16m2 16l4-16M6 9h14M4 15h14"></path>
                </svg>
              </div>
              حداکثر تعداد نمایش
            </label>
            <input type="number" class="w-full px-4 py-3 bg-white border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-orange-500 text-gray-900 transition-all duration-200" min="1" max="20" value="6" />
          </div>

          <div class="bg-white rounded-xl shadow-sm border border-pink-100 p-6">
            <label class="flex items-center gap-2 text-sm font-semibold text-gray-900 mb-4">
              <div class="p-2 bg-pink-100 rounded-lg">
                <svg class="w-4 h-4 text-pink-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4h13M3 8h9m-9 4h9m5-4v12m0 0l-4-4m4 4l4-4"></path>
                </svg>
              </div>
              نحوه مرتب‌سازی
            </label>
            <select class="w-full px-4 py-3 bg-white border border-gray-300 rounded-lg focus:ring-2 focus:ring-pink-500 focus:border-pink-500 text-gray-900 transition-all duration-200 appearance-none">
              <option value="manual">دستی (بر اساس اولویت)</option>
              <option value="price">قیمت</option>
              <option value="popularity">محبوبیت</option>
              <option value="newest">جدیدترین</option>
            </select>
          </div>
        </div>

        <!-- تنظیمات نمایش -->
        <div class="bg-white rounded-xl shadow-sm border border-slate-100 p-6">
          <h4 class="text-lg font-semibold text-gray-800 mb-6">گزینه‌های نمایش</h4>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <label class="flex items-center gap-3 p-6 rounded-lg border border-gray-200 hover:bg-gray-50 cursor-pointer transition-all duration-200">
              <input type="checkbox" class="w-5 h-5 text-orange-600 bg-gray-100 border-gray-300 rounded focus:ring-orange-500 focus:ring-2" checked />
              <div class="flex items-center gap-2">
                <div class="p-1 bg-orange-100 rounded">
                  <svg class="w-3 h-3 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4M7.835 4.697a3.42 3.42 0 001.946-.806 3.42 3.42 0 014.438 0 3.42 3.42 0 001.946.806 3.42 3.42 0 013.138 3.138 3.42 3.42 0 00.806 1.946 3.42 3.42 0 010 4.438 3.42 3.42 0 00-.806 1.946 3.42 3.42 0 01-3.138 3.138 3.42 3.42 0 00-1.946.806 3.42 3.42 0 01-4.438 0 3.42 3.42 0 00-1.946-.806 3.42 3.42 0 01-3.138-3.138 3.42 3.42 0 00-.806-1.946 3.42 3.42 0 010-4.438 3.42 3.42 0 00.806-1.946 3.42 3.42 0 013.138-3.138z"></path>
                  </svg>
                </div>
                <span class="text-sm text-gray-700 font-medium">نمایش در صفحه محصول</span>
              </div>
            </label>

            <label class="flex items-center gap-3 p-6 rounded-lg border border-gray-200 hover:bg-gray-50 cursor-pointer transition-all duration-200">
              <input type="checkbox" class="w-5 h-5 text-green-600 bg-gray-100 border-gray-300 rounded focus:ring-green-500 focus:ring-2" />
              <div class="flex items-center gap-2">
                <div class="p-1 bg-green-100 rounded">
                  <svg class="w-3 h-3 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2ل.4 2M7 13h10l4-8H5.4m0 0L7 13m0 0l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17M17 13v4a2 2 0 01-2 2H9a2 2 0 01-2-2v-4m8 0V9a2 2 0 00-2-2H9a2 2 0 00-2 2v4.01"></path>
                  </svg>
                </div>
                <span class="text-sm text-gray-700 font-medium">نمایش در سبد خرید</span>
              </div>
            </label>

            <label class="flex items-center gap-3 p-6 rounded-lg border border-gray-200 hover:bg-gray-50 cursor-pointer transition-all duration-200">
              <input type="checkbox" class="w-5 h-5 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 focus:ring-2" />
              <div class="flex items-center gap-2">
                <div class="p-1 bg-blue-100 rounded">
                  <svg class="w-3 h-3 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"></path>
                  </svg>
                </div>
                <span class="text-sm text-gray-700 font-medium">پیشنهاد خرید همراه</span>
              </div>
            </label>

            <label class="flex items-center gap-3 p-6 rounded-lg border border-gray-200 hover:bg-gray-50 cursor-pointer transition-all duration-200">
              <input type="checkbox" class="w-5 h-5 text-purple-600 bg-gray-100 border-gray-300 rounded focus:ring-purple-500 focus:ring-2" checked />
              <div class="flex items-center gap-2">
                <div class="p-1 bg-purple-100 rounded">
                  <svg class="w-3 h-3 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
                  </svg>
                </div>
                <span class="text-sm text-gray-700 font-medium">نمایش قیمت محصولات</span>
              </div>
            </label>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- Image preview modal -->
  <ImagePreviewModal v-model="showPreview" :image="previewImage" />
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
</script>

<script setup lang="ts">
import { onMounted, reactive, ref, watch } from 'vue';
import ImagePreviewModal from '~/components/media/ImagePreviewModal.vue';
import { useProductCreateStore } from '~/stores/productCreate';

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

interface ComplementProduct {
  id: number | string;
  name?: string;
  image_url?: string;
  image?: string;
  main_image?: string;
  thumbnail?: string;
  priority?: number;
  price?: number;
  sku?: string;
  slug?: string;
  category?: string | { name: string };
  images?: { image_url: string; thumbnail?: string }[];
  [key: string]: unknown;
}

interface PreviewImage {
  url: string;
  name: string;
}

interface Category {
  id: number | string;
  name: string;
  [key: string]: unknown;
}

const sections = reactive({
  relatedProducts: false,
  relatedSettings: false
})

const toggleSection = (section: keyof typeof sections) => {
  sections[section] = !sections[section]
}

const store = useProductCreateStore()
const q = ref('')
const results = ref<ComplementProduct[]>([])
const selected = ref<ComplementProduct[]>([])
const categories = ref<Category[]>([])
const selectedCategory = ref<string | number | ''>('')

function toThumbnail(url?: string | null){
  if (!url) return null
  const dotIdx = url.lastIndexOf('.')
  if (dotIdx === -1) return url + '_thumbnail'
  return url.slice(0, dotIdx) + '_thumbnail' + url.slice(dotIdx)
}

async function searchProducts(){
  if (!q.value.trim()) { results.value = []; return }
  try {
    const params = new URLSearchParams({
      search: q.value,
      limit: '10',
    })
    if (selectedCategory.value) params.set('category_id', String(selectedCategory.value))
    const res = await $fetch<ComplementProduct[] | { data: ComplementProduct[] }>(`/api/products?${params.toString()}`)
    const arr = Array.isArray(res) ? res : ('data' in res ? res.data : [])
    results.value = arr.map((p) => {
      const firstImage = (p?.images && p.images.length > 0 && p.images[0]?.image_url) || p.image || p.main_image || p.thumbnail
      const thumb = toThumbnail(firstImage as string)
      return {
        id: p.id,
        name: p.name,
        category: (p.category && typeof p.category === 'object' && 'name' in p.category) ? (p.category as { name: string }).name : '',
        price: p.price,
        sku: p.sku,
        slug: p.slug,
        image: firstImage,
        thumbnail: thumb || firstImage,
      }
    })
  } catch{ results.value = [] }
}

async function addRelated(productId){
  if (!store.isEditMode || !store.editingProductId) return
  try {
    await $fetch(`/api/product-complements/${store.editingProductId}`, { method:'POST', body:{ complement_product_id: productId } })
    await refreshSelected()
  } catch{ }
}

async function removeRelated(productId){
  if (!store.isEditMode || !store.editingProductId) return
  try {
    await $fetch(`/api/product-complements/${store.editingProductId}/${productId}`, { method:'DELETE' })
    await refreshSelected()
  } catch{ }
}

async function refreshSelected(){
  try {
    const res = await $fetch<{ data: ComplementProduct[] } | ComplementProduct[]>(`/api/product-complements/${store.editingProductId}`)
    const data = (res && 'data' in res) ? res.data : res
    const arr = Array.isArray(data) ? data : []
    selected.value = arr.map((it) => {
      const base = it.image_url || it.image || it.main_image || it.thumbnail
      const thumb = toThumbnail(base)
      return { ...it, thumbnail: thumb || base }
    })
  } catch{ selected.value = [] }
}

function onPriorityInput(item: ComplementProduct, evt: Event){
  const val = Number((evt.target as HTMLInputElement).value || 1)
  item.priority = val
  // ذخیره فوری اولویت در بک‌اند بدون تغییر استایل
  if (store.isEditMode && store.editingProductId && item?.id){
    // در نبود API اختصاصی update، از همان POST با UPSERT استفاده می‌کنیم
    $fetch(`/api/product-complements/${store.editingProductId}`, {
      method: 'POST',
      body: { complement_product_id: item.id, priority: val }
    }).catch(()=>{})
  }
}

// Preview modal state
const showPreview = ref(false)
const previewImage = ref<PreviewImage | null>(null)
function getOriginalUrl(obj: ComplementProduct){
  if(!obj) return null
  if(obj.images && obj.images[0]?.image_url) return obj.images[0].image_url
  return obj.image_url || obj.image || obj.main_image || obj.thumbnail || null
}
function openPreview(obj: ComplementProduct){
  const url = getOriginalUrl(obj)
  if(!url) return
  previewImage.value = { url, name: obj?.name || '' }
  showPreview.value = true
}

onMounted(async () => {
  try {
    const res = await $fetch<Category[]>('/api/admin/product-categories?all=1')
    categories.value = Array.isArray(res) ? res : []
  } catch { categories.value = [] }
  // اگر محصول در حالت ویرایش است، مکمل‌ها را لود کن
  if (store.isEditMode && store.editingProductId) {
    await refreshSelected()
  }
})

// وقتی شناسهٔ محصول برای ویرایش ست شد، لیست مکمل‌ها را لود کن
watch(() => store.editingProductId, async (val) => {
  if (val) {
    await refreshSelected()
  }
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


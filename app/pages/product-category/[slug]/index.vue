<template>
  <div class="space-y-6 max-w-screen-xl mx-auto px-2 pb-24">
    <!-- Top banner -->
    <div v-if="cat" class="bg-white rounded-lg shadow px-4 py-4 flex flex-col items-center text-center">
      <img v-if="catImage(cat)" :src="catImage(cat)" :alt="cat.name" class="w-full max-w-sm h-auto mb-4 rounded-lg object-cover" @error="onImgError($event)" />
      <h1 class="text-2xl font-bold text-gray-800">{{ cat.name }}</h1>
      <!-- eslint-disable-next-line vue/no-v-html -->
      <div v-if="cat.description" class="mt-4 text-gray-700 leading-relaxed" v-html="cat.description"></div>
    </div>
    
    <!-- Category not found message -->
    <div v-else class="bg-white rounded-lg shadow px-4 py-4 flex flex-col items-center text-center">
      <h1 class="text-2xl font-bold text-gray-800 mb-4">دسته‌بندی یافت نشد</h1>
      <p class="text-gray-600">دسته‌بندی با نام "{{ route.params.slug }}" در سیستم موجود نیست.</p>
      <NuxtLink to="/" class="mt-4 text-blue-600 hover:text-blue-800">بازگشت به صفحه اصلی</NuxtLink>
    </div>

    <!-- Content -->
    <div v-if="cat" class="grid grid-cols-1 lg:grid-cols-4 gapx-4 py-4">
      <!-- Main list -->
      <div class="lg:col-span-3 space-y-4">
        <!-- Header with count & sort -->
        <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between bg-white shadow rounded-lg px-4 py-4 text-sm">
          <p class="mb-2 sm:mb-0">در حال نمایش {{ filteredProducts.length }} نتیجه</p>
          <div class="flex gap-3 overflow-auto">
            <button v-for="o in sortOptions" :key="o.value" :class="[ 'px-3 py-1 rounded-full whitespace-nowrap', sortBy===o.value ? 'bg-teal-500 text-white' : 'bg-gray-100 text-gray-600 hover:bg-gray-200']" @click="setSort(o.value)">{{ o.label }}</button>
          </div>
        </div>

        <!-- Products grid -->
        <div v-if="loading" class="text-center py-20 bg-white rounded-lg shadow">در حال بارگذاری...</div>
        <div v-else>
          <div v-if="filteredSortedProducts.length" class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 gapx-4 py-4">
            <NuxtLink v-for="p in filteredSortedProducts" :key="p.id" :to="buildProductLink(p)" class="group block bg-white rounded-lg shadow hover:shadow-lg transition-all duration-200">
              <div class="relative w-full h-[230px] bg-gray-50 flex items-center justify-center rounded-t-lg overflow-hidden">
                <!-- discount badge -->
                <span v-if="discountPercent(p)" class="absolute top-2 right-2 bg-red-500 text-white text-xs px-1.5 py-0.5 rounded">{{ discountPercent(p) }}٪</span>
                <img :src="productThumbnail(p)" :alt="p.name" class="object-cover w-full h-full transition-transform duration-200 group-hover:scale-105" @error="onImgError($event)" />
              </div>
              <div class="p-3 text-right space-y-1">
                <h3 class="text-sm font-semibold text-gray-800 line-clamp-2 min-h-[3rem]">{{ p.name }}</h3>
                <div v-if="p.category && p.category.name" class="text-xs text-gray-500">
                  دسته‌بندی: <span class="font-medium text-gray-700">{{ p.category.name }}</span>
                </div>
                <div v-if="p.brand && p.brand.name" class="text-xs text-gray-500">
                  برند: <span class="font-medium text-gray-700">{{ p.brand.name }}</span>
                </div>
                <template v-if="p.stock_quantity===0">
                  <p class="text-red-500 text-sm font-semibold">ناموجود</p>
                </template>
                <template v-else>
                  <p v-if="displayPrice(p)" class="text-emerald-600 font-bold flex items-center gap-1">
                    {{ displayPrice(p) }} <span class="text-xs">تومان</span>
                  </p>
                  <p v-if="p.sale_price && p.sale_price>0" class="text-gray-400 text-xs line-through">{{ p.price.toLocaleString('fa-IR') }} تومان</p>
                </template>
              </div>
            </NuxtLink>
          </div>
          <p v-else class="text-gray-600 bg-white p-8 rounded-lg shadow text-center">محصولی یافت نشد.</p>
        </div>
      </div>

      <!-- Sidebar -->
      <aside class="space-y-6">
        <!-- Price filter -->
        <div class="bg-white px-4 py-4 rounded-lg shadow space-y-4 text-sm">
          <h3 class="font-bold mb-2 text-center">فیلتر براساس قیمت:</h3>
          <div class="flex flex-col gap-2">
            <label>حداقل: {{ priceMin.toLocaleString('fa-IR') }} تومان</label>
            <input v-model.number="priceMin" type="range" min="0" :max="priceMaxGlobal" />
            <label>حداکثر: {{ priceMax.toLocaleString('fa-IR') }} تومان</label>
            <input v-model.number="priceMax" type="range" min="0" :max="priceMaxGlobal" />
            <button class="bg-teal-600 text-white w-full py-1 rounded" @click="applyPriceFilter">فیلتر</button>
          </div>
        </div>
        <!-- Last visited -->
        <div class="bg-white px-4 py-4 rounded-lg shadow text-sm">
          <h3 class="font-bold mb-3 text-center">آخرین بازدید های شما</h3>
          <ul v-if="lastVisited.length" class="space-y-2">
            <li v-for="p in lastVisited" :key="p.id" class="flex items-center gap-2">
              <img :src="productThumbnail(p)" class="w-10 h-10 object-cover rounded" />
            <NuxtLink :to="buildProductLink(p)" class="flex-1 text-xs truncate hover:text-teal-600">{{ p.name }}</NuxtLink>
            </li>
          </ul>
          <p v-else class="text-center text-gray-500">بازدیدی ندارید</p>
        </div>
      </aside>
    </div>
  </div>
</template>

<script lang="ts">
declare const useRoute: () => { params: Record<string, string>; query: Record<string, string> }
declare const useHead: (head: { title?: string; meta?: Array<{ name?: string; content?: string }> }) => void
declare const navigateTo: (to: string) => Promise<void>
declare const definePageMeta: (meta: { layout?: string }) => void
</script>

<script setup lang="ts">
import { computed, onMounted, ref, watch, watchEffect } from 'vue';
import { useProductLink } from '~/composables/useProductLink';

const route = useRoute();
const { buildProductLink } = useProductLink()
const slug = route.params.slug

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const cat = ref<any>(null)
const pending = ref(false)
// eslint-disable-next-line @typescript-eslint/no-explicit-any
const error = ref<any>(null)
// eslint-disable-next-line @typescript-eslint/no-explicit-any
const products = ref<any[]>([])
const loading = ref(false)

const fetchCat = async () => {
  try {
    pending.value = true
    const isPreview = route.query.preview === '1' || route.query.preview === 'true'
    
    // ابتدا لیست تمام دسته‌بندی‌ها را دریافت کن
    interface Category {
      id?: number | string
      slug?: string
      name?: string
      [key: string]: unknown
    }
    interface CategoriesResponse {
      data?: Category[]
      [key: string]: unknown
    }
    const categories = await $fetch<Category[] | CategoriesResponse>(`/api/product-categories?all=1`)
    const list = Array.isArray(categories) ? categories : (categories?.data || [])
    
    // جستجو برای دسته‌بندی با slug
    const result = list.find((c: Category) => c.slug === slug)
    
    if (!result) {
      throw new Error('Category not found')
    }
    
    // بررسی کنید که آیا این دسته‌بندی فرعی است یا نه
    if (result && result.parent_id && result.parent_id !== 0) {
      // اگر دسته‌بندی فرعی است، به URL صحیح redirect کن
      navigateTo(`/product-category/${result.parent_slug}/${result.slug}`)
      return
    }
    
    cat.value = result
    error.value = null
  } catch (e) {
    console.error('Error fetching category:', e)
    cat.value = null
    error.value = e
  } finally {
    pending.value = false
  }
}

// اجرای fetchCat در هر دو سمت سرور و کلاینت
if (import.meta.server) {
  await fetchCat()
}
if (import.meta.client) {
  onMounted(fetchCat)
}

// SEO head tags
watchEffect(() => {
  if (!cat.value) return
  useHead({
    title: cat.value.meta_title || cat.value.name,
    meta: [
      { name: 'description', content: cat.value.meta_description || cat.value.description || '' },
      ...(cat.value.meta_keywords ? [{ name: 'keywords', content: cat.value.meta_keywords }] : [])
    ]
  })
})

// لود محصولات دسته‌بندی
watch(() => cat.value, async (newCat) => {
  if (newCat && newCat.id) {
    try {
      loading.value = true
      const response = await $fetch('/api/products/public')
      const allProducts = Array.isArray(response) ? response : []
      
      // فیلتر محصولات بر اساس دسته‌بندی
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      products.value = allProducts.filter((p: any) => {
        const pid = p.category_id != null ? String(p.category_id) : (p.category ? String(p.category.id) : '')
        return pid === String(newCat.id)
      })
    } catch (error) {
      console.error('❌ Error fetching products:', error)
      products.value = []
    } finally {
      loading.value = false
    }
  }
})

// Price filter state
const priceMin = ref(0)
const priceMaxGlobal = computed(() => {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const vals = products.value.map((p: any) => p.sale_price && p.sale_price>0 ? p.sale_price : p.price).filter((v:number)=>v && v>0)
  return vals.length ? Math.max(...vals) : 0
})
const priceMax = ref(0)
// once products loaded set priceMax
watchEffect(() => {
  if (!loading.value) priceMax.value = priceMaxGlobal.value
})
function applyPriceFilter(){ /* computed handles */ }

// Sorting
const sortOptions = [
  { label:'گرانترین', value:'priceDesc' },
  { label:'ارزانترین', value:'priceAsc' },
  { label:'جدیدترین', value:'new' }
]
const sortBy = ref('priceDesc')
const setSort = (v:string)=>{ sortBy.value = v }

// Computed lists
const filteredProducts = computed(()=>{
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  return products.value.filter((p: any)=>{
    const price = p.sale_price && p.sale_price>0 ? p.sale_price : p.price
    if(priceMax.value && price>priceMax.value) return false
    if(price < priceMin.value) return false
    return true
  })
})
const filteredSortedProducts = computed(()=>{
  const list = [...filteredProducts.value]
  switch(sortBy.value){
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    case 'priceAsc': list.sort((a: any,b: any)=>( (a.sale_price&&a.sale_price>0?a.sale_price:a.price) - (b.sale_price&&b.sale_price>0?b.sale_price:b.price) )); break
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    case 'priceDesc': list.sort((a: any,b: any)=>( (b.sale_price&&b.sale_price>0?b.sale_price:b.price) - (a.sale_price&&a.sale_price>0?a.sale_price:a.price) )); break
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    case 'new': list.sort((a: any,b: any)=> new Date(b.created_at).getTime() - new Date(a.created_at).getTime()); break
  }
  return list
})

// eslint-disable-next-line @typescript-eslint/no-explicit-any
function discountPercent(p:any){
  if(p.sale_price && p.sale_price>0 && p.price && p.price>p.sale_price){
    return Math.round((1 - p.sale_price/p.price)*100)
  }
  return 0
}

// Last visited (simple localStorage demo)
// eslint-disable-next-line @typescript-eslint/no-explicit-any
const lastVisited = ref<any[]>([])
onMounted(()=>{
  try {
    const ls = localStorage.getItem('lastVisitedProducts')
    if(ls){
      const arr = JSON.parse(ls)
      if(Array.isArray(arr)) lastVisited.value = arr.slice(0,5)
    }
  }catch(e){}
})
// update localStorage on product page separately (not here)

// eslint-disable-next-line @typescript-eslint/no-explicit-any
function productThumbnail(p:any){
  const toVariant = (url:string)=>{
    if(!url) return ''
    const idx = url.lastIndexOf('.')
    if(idx === -1) return url
    const base = url.substring(0,idx)
    const ext = url.substring(idx)
    const candidates = [`${base}_small${ext}`, `${base}_thumbnail${ext}`]
    // pick first candidate (assume exists), else original
    return candidates[0]
  }
  if (p.image) return toVariant(p.image)
  if (Array.isArray(p.images) && p.images.length) return toVariant(p.images[0].image_url || p.images[0].url)
  return '/statics/images/default-image_100.png'
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
function onImgError(e:any){
  e.target.src = '/statics/images/default-image_100.png'
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
function displayPrice(p:any){
  const price = p.sale_price && p.sale_price>0 ? p.sale_price : p.price
  return price && price>0 ? price.toLocaleString('fa-IR') : ''
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
function catImage(c:any){
  const url = c.image_url || c.banner_url || ''
  if(!url) return ''
  // try category variant
  const idx = url.lastIndexOf('.')
  if(idx===-1) return url
  return url.slice(0,idx)+"_category"+url.slice(idx)
}

definePageMeta({ layout: 'default' })
</script>
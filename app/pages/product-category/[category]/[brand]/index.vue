<template>
  <div>
    <div class="bg-white rounded-lg shadow px-4 py-4 mb-6">
      <h1 class="text-2xl font-bold text-gray-800">{{ categoryData?.name }}</h1>
      <div v-if="categoryData?.description" class="mt-4 text-gray-700 leading-relaxed" v-html="categoryData.description"></div>
    </div>
    <div class="bg-gray-50 rounded-lg shadow px-4 py-4">
      <h2 class="text-xl font-semibold mb-4">محصولات این دسته</h2>
      <div v-if="loading" class="text-center py-6">در حال بارگذاری...</div>
      <div v-else>
        <div v-if="products.length" class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 gapx-4 py-4">
          <NuxtLink
            v-for="p in products"
            :key="p.id"
            :to="`/product/sku-${p.sku || p.id}`"
            class="block bg-white rounded-lg shadow hover:shadow-lg transition-all duration-200 group"
          >
            <!-- Image -->
            <div class="w-full h-[300px] overflow-hidden rounded-t-lg flex items-center justify-center bg-gray-100">
              <img
                :src="productThumbnail(p)"
                @error="onImgError($event)"
                alt="img"
                class="object-cover w-full h-full transition-transform duration-200 group-hover:scale-105"
              />
            </div>
            <!-- Info -->
            <div class="p-3 text-right space-y-1">
              <h3 class="text-sm font-semibold text-gray-800 line-clamp-2 min-h-[3rem]">{{ p.name }}</h3>
              <p class="text-emerald-600 font-bold" v-if="displayPrice(p)">{{ displayPrice(p) }} <span class="text-xs">تومان</span></p>
              <p v-else class="text-gray-500 text-xs">قیمت ثبت نشده</p>
            </div>
          </NuxtLink>
        </div>
        <p v-else class="text-gray-600">محصولی در این دسته‌بندی ثبت نشده است.</p>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string }) => void
declare const useFetch: <T = unknown>(url: string, options?: { transform?: (data: any) => T }) => Promise<{ data: { value: T }; pending: { value: boolean }; error: { value: Error | null } }>
declare const useHead: (head: { title?: string; meta?: Array<{ name?: string; content?: string }> }) => void
declare const useRoute: () => { params: Record<string, string>; query: Record<string, string> }
declare const createError: (options: { statusCode: number; statusMessage: string }) => Error

interface Category {
  id: number
  name: string
  slug: string
  description?: string
  meta_title?: string
  meta_description?: string
  meta_keywords?: string
}
</script>

<script setup lang="ts">
// useFetch is auto-imported in Nuxt 3
import { computed, ref, watchEffect } from 'vue';

const route = useRoute()
const combinedSlug = `${route.params.category}/${route.params.brand}`

const { data: cat } = await useFetch<Category>(`/api/product-category/${combinedSlug}`)

if (!cat.value) {
  throw createError({ statusCode: 404, statusMessage: 'Category not found' })
}

// Computed property برای unwrap کردن reactive ref
const categoryData = computed(() => cat.value as Category | null)

// dynamic SEO meta
watchEffect(() => {
  if (!categoryData.value) return
  useHead({
    title: categoryData.value.meta_title || categoryData.value.name,
    meta: [
      { name: 'description', content: categoryData.value.meta_description || categoryData.value.description || '' },
      ...(categoryData.value.meta_keywords ? [{ name: 'keywords', content: categoryData.value.meta_keywords }] : [])
    ]
  })
})

const products = ref<any[]>([])
const loading = ref(true)
try {
  const { data } = await useFetch('/api/products')
  const all = (data.value || []) as any[]
  products.value = all.filter(p => {
    const pid = p.category_id != null ? String(p.category_id) : (p.category ? String(p.category.id) : '')
    if (categoryData.value && pid === String(categoryData.value.id)) return true
    if (categoryData.value && p.category && p.category.slug === categoryData.value.slug) return true
    return false
  })
} finally {
  loading.value = false
}

function productThumbnail(p:any){
  const toVariant = (url:string)=>{
    if(!url) return ''
    const idx = url.lastIndexOf('.')
    if(idx === -1) return url
    const base = url.substring(0,idx)
    const ext = url.substring(idx)
    const candidates = [`${base}_small${ext}`, `${base}_thumbnail${ext}`]
    return candidates[0]
  }
  if (p.image) return toVariant(p.image)
  if (Array.isArray(p.images) && p.images.length) return toVariant(p.images[0].image_url || p.images[0].url)
  return '/statics/images/default-image_100.png'
}

function onImgError(e:any){
  e.target.src = '/statics/images/default-image_100.png'
}

function displayPrice(p:any){
  const price = p.sale_price && p.sale_price>0 ? p.sale_price : p.price
  return price && price>0 ? price.toLocaleString('fa-IR') : ''
}

definePageMeta({ layout: 'default' })
</script> 
<template>
  <div class="space-y-6" dir="rtl">
    <div v-if="complements.length" class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
      <NuxtLink
        v-for="compProduct in complements"
        :key="compProduct.id"
        :to="`/product/sku-${compProduct.sku || compProduct.id}${compProduct.slug ? '/' + encodeURIComponent(compProduct.slug) : ''}`"
        class="bg-white rounded-lg border border-gray-200 overflow-hidden hover:shadow-lg transition-shadow group"
      >
        <div class="aspect-square bg-gray-100 overflow-hidden">
          <img
            :src="imageUrl(compProduct)"
            :alt="compProduct.name"
            class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300"
          />
        </div>
        <div class="p-6">
          <h3 class="font-medium text-gray-900 text-sm mb-2 line-clamp-2">{{ compProduct.name }}</h3>
          <div class="flex items-center justify-between">
            <div class="text-green-600 font-bold text-sm">
              {{ formatPrice(compProduct.price) }} تومان
            </div>
          </div>
          <div class="mt-3" @click.stop>
            <button
              v-if="!isInCart(compProduct.id) && isAvailable(compProduct)"
              class="w-full bg-green-600 hover:bg-green-700 text-white text-sm font-semibold rounded-md py-2"
              @click="add(compProduct.id)"
            >
              افزودن به سبد خرید
            </button>
            <button
              v-else-if="!isInCart(compProduct.id) && !isAvailable(compProduct)"
              disabled
              class="w-full bg-red-100 text-red-700 text-sm font-semibold rounded-md py-2 cursor-not-allowed"
            >
              ناموجود
            </button>
            <div v-else class="flex items-center gap-2">
              <button class="w-8 h-8 rounded-md bg-gray-100 hover:bg-gray-200 flex items-center justify-center" @click="decrement(compProduct.id)">−</button>
              <span class="min-w-[2rem] text-center">{{ getItemQuantity(compProduct.id) }}</span>
              <button class="w-8 h-8 rounded-md bg-gray-100 hover:bg-gray-200 flex items-center justify-center" @click="increment(compProduct.id)">+</button>
              <button class="ml-auto text-red-600 text-xs" @click="remove(compProduct.id)">حذف</button>
            </div>
          </div>
        </div>
      </NuxtLink>
    </div>

    <div v-else class="text-center py-12">
      <div class="text-gray-400 mb-4">
        <svg class="w-16 h-16 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          
        </svg>
      </div>
      <h3 class="text-lg font-medium text-gray-900 mb-2">محصول مکملی موجود نیست</h3>
      <p class="text-gray-500">در حال حاضر محصول مکملی برای نمایش وجود ندارد.</p>
    </div>
  </div>
  
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useCart } from '~/composables/useCart'

interface Product {
  id: number
  name: string
  main_image: string
  price: number
  rating?: number
  sku?: string
  slug?: string
}

interface Props {
  product: {
    complement_products?: Product[]
  }
}

const props = defineProps<Props>()

const complements = computed(() => props.product.complement_products || [])

// تابع برای تبدیل تصویر به thumbnail
function toThumbnail(url: string): string {
  if (!url) return '/default-product.jpg'
  
  const dotIdx = url.lastIndexOf('.')
  if (dotIdx === -1) return url + '_thumbnail.webp'
  
  const baseName = url.slice(0, dotIdx)
  const originalExt = url.slice(dotIdx)
  
  // اگر فرمت اصلی jpg یا jpeg است، thumbnail را با webp بساز
  if (originalExt.toLowerCase() === '.jpg' || originalExt.toLowerCase() === '.jpeg') {
    return baseName + '_thumbnail.webp'
  }
  
  // برای سایر فرمت‌ها، همان فرمت را حفظ کن
  return baseName + '_thumbnail' + originalExt
}

function imageUrl(p: Record<string, unknown>): string {
  const base = (p?.images && Array.isArray(p.images) && (p.images[0] as Record<string, unknown>)?.image_url) || p?.image_url || p?.image || p?.main_image || ''
  return toThumbnail(base as string)
}

const formatPrice = (price: number): string => {
  return new Intl.NumberFormat('fa-IR').format(price)
}

// Cart interactions
const { addToCart, updateCartItem, removeFromCart, isInCart, getItemQuantity } = useCart()

async function add(productId:number){
  // اطمینان از ایجاد session با GET cart (ست‌شدن کوکی توسط بک‌اند)
  try { await $fetch('/api/cart', { credentials: 'include' }) } catch {}
  await addToCart(productId, 1)
}
async function increment(productId:number){
  const qty = getItemQuantity(productId) + 1
  const item = cartItem(productId)
  if(item) await updateCartItem(item.id, qty)
}
async function decrement(productId:number){
  const qty = Math.max(1, getItemQuantity(productId) - 1)
  const item = cartItem(productId)
  if(item) await updateCartItem(item.id, qty)
}
async function remove(productId:number){
  const item = cartItem(productId)
  if(item) await removeFromCart(item.id)
}

function cartItem(productId:number){
  const { cartItems } = useCart()
  return cartItems.value.find((i: Record<string, unknown>)=> i.product_id === productId)
}

// موجود بودن برای خرید در کارت مکمل
function isAvailable(p: Record<string, unknown>){
  const price = Number(p?.price || 0)
  const track = p?.track_inventory === true
  const stock = Number(p?.stock_quantity ?? 0)
  const preorder = p?.allow_preorder === true
  // قانون: اگر قیمت 0 باشد یا (رهگیری موجودی فعال و موجودی 0 و پیش‌خرید غیرفعال)، ناموجود
  if (price <= 0) return false
  if (track && stock <= 0 && !preorder) return false
  return true
}
</script>

<style scoped>
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>



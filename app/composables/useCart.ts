// Composable برای مدیریت سبد خرید
// -------------------------
// Imports & Shared State
// -------------------------
import { readonly, ref } from 'vue';
import type { Product } from '~/types/widget';
// این متغیرها در سطح ماژول تعریف شده‌اند تا در تمام فراخوانی‌های
// useCart به‌صورت مشترک استفاده شوند و وضعیت سبد خرید در سراسر
// برنامه یکسان باقی بماند.
const cartItems = ref<(Product & { quantity: number; id: number; product_id: number; cart_id?: number; is_next?: boolean })[]>([])
const cartCount = ref<number>(0)
const cartTotal = ref<number>(0)
const loading = ref<boolean>(false)

interface ApiError {
  status?: number
  data?: {
    message?: string
  }
  [key: string]: unknown
}

export const useCart = () => {
  // اطمینان از ایجاد سشن سبد خرید
  const ensureSession = async () => {
    // @ts-ignore - Nuxt auto-imports
    const cartCookie = useCookie('session_id')
    if (cartCookie.value) return

    try {
      // ابتدا سعی می‌کنیم session ایجاد کنیم
      await $fetch('/api/cart/create-session', { method: 'POST', credentials: 'include' })
      return
    } catch {
      // اگر ایجاد session شکست خورد، سعی می‌کنیم از بک‌اند session بگیریم
      try {
        await $fetch('/api/cart', { method: 'GET', credentials: 'include' })
        return
      } catch (e2) {
        console.error('خطا در ایجاد سشن سبد خرید', e2)
      }
    }
  }

  // دریافت محتویات سبد خرید
  const fetchCart = async () => {
    try {
      loading.value = true
      const response = await $fetch<unknown>('/api/cart', { credentials: 'include' })

      // پشتیبانی از هر دو ساختار پاسخ (Go و Nuxt)
      const cartData = (response && typeof response === 'object')
        ? ((response as { data?: unknown }).data && typeof (response as { data?: unknown }).data === 'object' ? (response as { data: Record<string, unknown> }).data : response as Record<string, unknown>)
        : { items: [], total_items: 0, total_price: 0 }

      const itemsArr = Array.isArray(cartData.items) ? cartData.items : []

      cartItems.value = itemsArr
        .filter((item: unknown) => item && typeof item === 'object')
        .map((item: unknown) => {
          const it = item as Record<string, unknown>
          if (!it.id || !it.product_id) return null
          return {
            id: Number(it.id),
            cart_id: Number(it.cart_id) || 0,
            product_id: Number(it.product_id),
            quantity: Number(it.quantity) || 1,
            is_next: Boolean(it.is_next || it.IsNextPurchase || false),
            qty: Number(it.quantity) || 1,
            name: String(it.name || 'محصول بدون نام'),
            sku: String(it.sku || ''),
            image: String(it.image || '/default-product.svg'),
            price: Number(it.price) || 0,
            original_price: it.original_price ? Number(it.original_price) : undefined,
            discount_percentage: it.discount_percentage ? Number(it.discount_percentage) : undefined,
            rating: it.rating ? Number(it.rating) : undefined,
            rating_count: it.rating_count ? Number(it.rating_count) : undefined,
            features: Array.isArray(it.features) ? it.features as string[] : [],
            created_at: it.created_at ? String(it.created_at) : null,
            updated_at: it.updated_at ? String(it.updated_at) : null
          }
        })
        .filter(Boolean) as (Product & { quantity: number; id: number; product_id: number; cart_id?: number; is_next?: boolean })[]

      cartCount.value = Number(cartData.total_items) || 0
      cartTotal.value = Number(cartData.total_price) || 0
    } catch (error) {
      console.error('خطا در دریافت سبد خرید:', error)
      // در صورت خطا، سبد خرید را خالی می‌کنیم
      cartItems.value = []
      cartCount.value = 0
      cartTotal.value = 0
    } finally {
      loading.value = false
    }
  }

  // افزودن محصول به سبد خرید
  const addToCart = async (productId: number, quantity: number = 1, productData?: Product) => {
    // اطمینان از وجود سشن قبل از ارسال درخواست به بک‌اند
    await ensureSession()

    try {
      loading.value = true

      // اطمینان از صحت نوع داده‌ها
      const parsedProductId = parseInt(String(productId))
      const parsedQuantity = Math.max(1, parseInt(String(quantity))) // حداقل 1

      if (isNaN(parsedProductId) || parsedProductId <= 0) {
        return { success: false, message: 'شناسه محصول نامعتبر است' }
      }

      if (isNaN(parsedQuantity) || parsedQuantity <= 0) {
        return { success: false, message: 'تعداد محصول نامعتبر است' }
      }



      // دریافت CSRF token
      // @ts-ignore - Nuxt auto-imports
      const { getCSRFToken } = useCSRF()
      const csrfToken = await getCSRFToken()

      if (!csrfToken) {
        return { success: false, message: 'CSRF token در دسترس نیست' }
      }

      const response = await $fetch('/api/cart/add', {
        method: 'POST',
        headers: {
          'x-csrf-token': csrfToken
        },
        body: {
          product_id: parsedProductId,
          quantity: parsedQuantity,
          // ارسال داده‌های کامل محصول
          name: productData?.name || 'محصول بدون نام',
          price: productData?.price || 0,
          image: productData?.image || '/default-product.svg',
          original_price: productData?.original_price,
          discount_percentage: productData?.discount_percentage,
          rating: productData?.rating,
          rating_count: productData?.rating_count,
          sku: productData?.sku
        },
        credentials: 'include'
      })

      if (response && (response as { success: boolean; message?: string }).success) {
        // بعد از افزودن، سبد خرید را به‌روزرسانی می‌کنیم
        await fetchCart()
        return { success: true, message: (response as { message?: string }).message || 'محصول به سبد خرید اضافه شد' }
      }

      return { success: false, message: (response as { message?: string })?.message || 'خطا در افزودن به سبد خرید' }
    } catch (error) {
      console.error('خطا در افزودن به سبد خرید:', error)
      const apiError = error as ApiError
      console.error('Error details:', apiError.data) // جزئیات بیشتر

      // پردازش خطاهای مختلف
      if (apiError.status === 400) {
        return { success: false, message: apiError.data?.message || 'داده‌های ورودی نامعتبر است' }
      } else if (apiError.status === 404) {
        return { success: false, message: 'محصول یافت نشد' }
      } else {
        return { success: false, message: 'خطا در افزودن به سبد خرید' }
      }
    } finally {
      loading.value = false
    }
  }

      // انتقال آیتم به خرید بعدی
  const moveCartItemToNext = async (cartItemId: number) => {
    // @ts-ignore - Nuxt auto-imports
    const sessionId = useCookie<string>('session_id').value
    await $fetch<unknown>('/api/cart/move-next', { method: 'PUT', body: { cart_item_id: cartItemId, session_id: sessionId }, credentials: 'include' })
    await fetchCart()
  }

  // بازگرداندن آیتم از خرید بعدی به سبد
  const moveCartItemToCart = async (cartItemId: number) => {
    // @ts-ignore - Nuxt auto-imports
    const sessionId = useCookie<string>('session_id').value
    await $fetch<unknown>('/api/cart/move-cart', { method: 'PUT', body: { cart_item_id: cartItemId, session_id: sessionId }, credentials: 'include' })
    await fetchCart()
  }

  // به‌روزرسانی تعداد محصول در سبد خرید
  const updateCartItem = async (cartItemId: number, quantity: number) => {
    try {
      loading.value = true

      // دریافت CSRF token
      // @ts-ignore - Nuxt auto-imports
      const { getCSRFToken } = useCSRF()
      const csrfToken = await getCSRFToken()

      if (!csrfToken) {
        return { success: false, message: 'CSRF token در دسترس نیست' }
      }

      const response = await $fetch('/api/cart/update', {
        method: 'PUT',
        headers: {
          'x-csrf-token': csrfToken
        },
        body: {
          cart_item_id: cartItemId,
          quantity: quantity,
          // @ts-ignore - Nuxt auto-imports
          session_id: useCookie('session_id').value
        },
        credentials: 'include'
      })

      if ((response as { success: boolean; message?: string }).success) {
        await fetchCart()
        return { success: true, message: (response as { message?: string }).message }
      }

      return { success: false, message: (response as { message?: string }).message || 'خطا در به‌روزرسانی' }
    } catch (error) {
      console.error('خطا در به‌روزرسانی سبد خرید:', error)
      return { success: false, message: 'خطا در به‌روزرسانی سبد خرید' }
    } finally {
      loading.value = false
    }
  }

  // حذف محصول از سبد خرید
  const removeFromCart = async (cartItemId: number) => {
    try {
      loading.value = true

      // دریافت CSRF token
      // @ts-ignore - Nuxt auto-imports
      const { getCSRFToken } = useCSRF()
      const csrfToken = await getCSRFToken()

      if (!csrfToken) {
        return { success: false, message: 'CSRF token در دسترس نیست' }
      }

      const response = await $fetch('/api/cart/remove', {
        method: 'DELETE',
        headers: {
          'x-csrf-token': csrfToken
        },
        body: {
          cart_item_id: cartItemId,
          // @ts-ignore - Nuxt auto-imports
          session_id: useCookie('session_id').value
        },
        credentials: 'include'
      })

      if ((response as { success: boolean; message?: string }).success) {
        await fetchCart()
        return { success: true, message: (response as { message?: string }).message }
      }

      return { success: false, message: (response as { message?: string }).message || 'خطا در حذف' }
    } catch (error) {
      console.error('خطا در حذف از سبد خرید:', error)
      return { success: false, message: 'خطا در حذف از سبد خرید' }
    } finally {
      loading.value = false
    }
  }

  // پاک کردن کامل سبد خرید
  const clearCart = async () => {
    try {
      loading.value = true

      // دریافت CSRF token
      // @ts-ignore - Nuxt auto-imports
      const { getCSRFToken } = useCSRF()
      const csrfToken = await getCSRFToken()

      if (!csrfToken) {
        return { success: false, message: 'CSRF token در دسترس نیست' }
      }

      // @ts-ignore - Nuxt auto-imports
      const sessionId = useCookie('session_id').value
      if (sessionId) {
        await $fetch('/api/cart/clear', {
          method: 'DELETE',
          headers: {
            'x-csrf-token': csrfToken
          },
          body: { session_id: sessionId },
          credentials: 'include'
        })
      }
      return { success: true, message: 'سبد خرید پاک شد' }
    } catch (error) {
      console.error('خطا در پاک کردن سبد خرید:', error)
      return { success: false, message: 'خطا در پاک کردن سبد خرید' }
    } finally {
      loading.value = false
    }
  }

  // بررسی وجود محصول در سبد خرید
  const isInCart = (productId: number) => {
    return cartItems.value.some(item => item.product_id === productId)
  }

  // دریافت تعداد محصول در سبد خرید
  const getItemQuantity = (productId: number) => {
    const item = cartItems.value.find(item => item.product_id === productId)
    return item ? item.quantity : 0
  }

  return {
    // State
    cartItems: readonly(cartItems),
    cartCount: readonly(cartCount),
    cartTotal: readonly(cartTotal),
    loading: readonly(loading),

    // Methods
    fetchCart,
    addToCart,
    updateCartItem,
    removeFromCart,
    moveCartItemToNext,
    moveCartItemToCart,
    clearCart,
    isInCart,
    getItemQuantity
  }
} 
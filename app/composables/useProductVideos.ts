import { computed, readonly, ref } from 'vue'
import { useProductCreateStore } from '~/stores/productCreate'

export interface ProductVideo {
  id?: number
  product_id?: number
  title: string
  description: string
  video_url: string
  thumbnail_url?: string
  show_in_gallery: boolean
  autoplay: boolean
  show_controls: boolean
  sort_order: number
  lazy_load?: boolean
  created_at?: string
  updated_at?: string
}

export const useProductVideos = () => {
  const pStore = useProductCreateStore()

  const videos = ref<ProductVideo[]>([])
  const isLoading = ref(false)
  const error = ref<string | null>(null)

  // دریافت شناسه محصول فعلی
  const currentProductId = computed<number | null>(() => {
    const raw = pStore.editingProductId as unknown
    if (raw === null || raw === undefined) return null
    const asNumber = Number(raw)
    return Number.isFinite(asNumber) ? asNumber : null
  })

  // دریافت ویدیوهای محصول
  const fetchVideos = async (productId?: number) => {
    if (!productId && !currentProductId.value) {
      error.value = 'شناسه محصول موجود نیست'
      return
    }

    const id = productId || currentProductId.value
    isLoading.value = true
    error.value = null

    try {
      const apiUrl = `/api/product-videos/${id}`
      const resp = await fetch(apiUrl).then(res => res.json())

      // Accept both wrapped { success, data } and raw shapes like { data: [] } or []
      let list: ProductVideo[] = []
      if (Array.isArray(resp)) {
        list = resp as ProductVideo[]
      } else if (Array.isArray(resp?.data)) {
        list = resp.data as ProductVideo[]
      } else if (resp?.success && Array.isArray(resp?.data)) {
        list = resp.data as ProductVideo[]
      }

      // نرمال‌سازی مسیر ویدیوها برای سروینگ از public
      const normalize = (raw?: string) => {
        if (!raw) return ''
        let p = String(raw).replace(/\\/g, '/').replace(/\/public\//i, '/')
        p = p.replace(/^\/uploads\b/i, '/uploads')
        if (!/^https?:\/\//i.test(p) && !p.startsWith('/')) p = '/' + p
        if (!/\/uploads\//i.test(p) && /^\/(products|product-categories|brands|banners|customer)\//.test(p)) {
          p = '/uploads/media' + p
        }
        return p
      }
      videos.value = (list || []).map(v => ({ ...v, video_url: normalize(v.video_url), thumbnail_url: normalize(v.thumbnail_url) }))
    } catch (err) {
      const e = err as Error
      error.value = e?.message || 'خطا در دریافت ویدیوها'
    } finally {
      isLoading.value = false
    }
  }

  // افزودن ویدیو جدید
  const addVideo = async (videoData: Omit<ProductVideo, 'id' | 'product_id' | 'created_at' | 'updated_at'>) => {
    if (!currentProductId.value) {
      error.value = 'شناسه محصول موجود نیست'
      return false
    }

    isLoading.value = true
    error.value = null

    try {
      const resp = await fetch(`/api/product-videos/${currentProductId.value}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(videoData)
      }).then(res => res.json())

      // Accept both wrapped and raw responses
      const created = (resp && resp.success !== undefined) ? resp.data : resp
      if (created && (created.id || created.video_url)) {
        videos.value.push(created as ProductVideo)
        return true
      }
      error.value = 'خطا در افزودن ویدیو'
      return false
    } catch (err) {
      const e = err as Error
      console.error('خطا در افزودن ویدیو:', e)
      error.value = e.message || 'خطا در افزودن ویدیو'
      return false
    } finally {
      isLoading.value = false
    }
  }

  // ویرایش ویدیو
  const updateVideo = async (videoId: number, videoData: Partial<ProductVideo>) => {
    isLoading.value = true
    error.value = null

    try {
      const resp = await fetch(`/api/product-videos/${videoId}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(videoData)
      }).then(res => res.json())

      // Consider both wrapped and raw {message}
      const ok = (resp && resp.success !== false)
      if (ok) {
        const index = videos.value.findIndex(v => v.id === videoId)
        if (index !== -1) {
          videos.value[index] = { ...videos.value[index], ...videoData }
        }
        return true
      }
      error.value = 'خطا در ویرایش ویدیو'
      return false
    } catch (err) {
      const e = err as Error
      console.error('خطا در ویرایش ویدیو:', e)
      error.value = e.message || 'خطا در ویرایش ویدیو'
      return false
    } finally {
      isLoading.value = false
    }
  }

  // حذف ویدیو
  const deleteVideo = async (videoId: number) => {
    isLoading.value = true
    error.value = null

    try {
      const resp = await fetch(`/api/product-videos/${videoId}`, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json'
        }
      }).then(res => res.json())

      const ok = (resp && resp.success !== false)
      if (ok) {
        const index = videos.value.findIndex(v => v.id === videoId)
        if (index !== -1) {
          videos.value.splice(index, 1)
        }
        return true
      }
      error.value = 'خطا در حذف ویدیو'
      return false
    } catch (err) {
      const e = err as Error
      console.error('خطا در حذف ویدیو:', e)
      error.value = e.message || 'خطا در حذف ویدیو'
      return false
    } finally {
      isLoading.value = false
    }
  }

  // تغییر ترتیب ویدیوها
  const reorderVideos = async (videoIds: number[]) => {
    const updates = videoIds.map((id, index) => ({
      id,
      sort_order: index
    }))

    const promises = updates.map(update =>
      updateVideo(update.id, { sort_order: update.sort_order })
    )

    try {
      await Promise.all(promises)
      // بازخوانی لیست برای اطمینان از ترتیب صحیح
      await fetchVideos()
      return true
    } catch (err) {
      console.error('خطا در تغییر ترتیب ویدیوها:', err)
      return false
    }
  }

  // پاک کردن خطا
  const clearError = () => {
    error.value = null
  }

  return {
    videos: readonly(videos),
    isLoading: readonly(isLoading),
    error: readonly(error),
    currentProductId,
    fetchVideos,
    addVideo,
    updateVideo,
    deleteVideo,
    reorderVideos,
    clearError
  }
} 
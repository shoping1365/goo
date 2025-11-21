<template>
  <div class="product-videos-tab">
    <!-- ویدیو اصلی محصول -->
    <div v-if="product.video_url" class="main-video-section mb-8">
      <h3 class="text-xl font-bold mb-4 text-gray-900">ویدیو اصلی محصول</h3>
      <div class="bg-white rounded-lg shadow-md overflow-hidden">
        <ProductVideo 
          :video-url="normalize(product.video_url)"
          :poster-image="product.images?.[0]?.image_url || product.images?.[0]?.url"
          :title="product.name"
          :description="product.description"
          :autoplay="false"
          :show-controls="true"
          :show-in-gallery="true"
          :lazy-load="true"
        />
      </div>
    </div>
    
    <!-- لیست ویدیوهای محصول -->
    <div v-if="videosToShow.length > 0" class="additional-videos-section">
      <h3 class="text-xl font-bold mb-4 text-gray-900">
        {{ product.video_url ? 'ویدیوهای اضافی' : 'ویدیوهای محصول' }}
      </h3>
      
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div 
          v-for="video in videosToShow" 
          :key="video.id"
          class="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-lg transition-shadow duration-200"
        >
          <div class="video-thumbnail">
            <ProductVideo 
              :video-url="normalize(video.video_url)"
              :poster-image="video.thumbnail_url || product.images?.[0]?.image_url || product.images?.[0]?.url"
              :title="video.title"
              :description="video.description"
              :autoplay="video.autoplay"
              :show-controls="video.show_controls"
              :show-in-gallery="video.show_in_gallery"
              :lazy-load="video.lazy_load"
            />
          </div>
          
          <div class="p-6">
            <h4 class="font-semibold text-gray-900 mb-2">{{ video.title }}</h4>
            <p v-if="video.description" class="text-sm text-gray-600 mb-3 line-clamp-2">
              {{ video.description }}
            </p>
            
            <div class="flex items-center justify-between text-xs text-gray-500">
              <span v-if="video.duration">{{ formatDuration(video.duration) }}</span>
              <span v-if="video.created_at">{{ formatDate(video.created_at) }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
    
    <!-- حالت خالی -->
    <div v-else>
      <p>هیچ ویدیویی برای این محصول ثبت نشده است.</p>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import ProductVideo from '@/components/product/ProductVideo.vue'

const props = defineProps({
  product: { type: Object, required: true }
})

// ویدیوهای قابل نمایش (از خود محصول یا از API در صورت نبود)
const videos = ref(Array.isArray(props.product?.videos) ? props.product.videos : [])
const videosToShow = computed(() => {
  const list = Array.isArray(videos.value) ? videos.value : []
  // فقط ویدیوهایی که آدرس معتبر دارند
  return list.filter(v => !!normalize(v?.video_url))
})

async function fetchVideosForProduct(pid){
  if (!pid) return
  try {
    const res = await $fetch(`/api/product-videos/${pid}`)
    const list = Array.isArray(res?.data) ? res.data : (Array.isArray(res) ? res : [])
    videos.value = list
  } catch (_e) {
    // ignore
  }
}

onMounted(async () => {
  if (props.product?.id) {
    await fetchVideosForProduct(props.product.id)
  }
})

watch(() => props.product?.id, (newId, oldId) => {
  if (newId && newId !== oldId) fetchVideosForProduct(newId)
})

// نرمال‌سازی مسیر ویدیو برای سروینگ از public
const normalize = (raw) => {
  if (!raw) return ''
  let p = String(raw).replace(/\\/g,'/').replace(/\/public\//i,'/')
  // همسان‌سازی حروف کوچک/بزرگ برای پوشه uploads در ابتدای مسیر
  p = p.replace(/^\/uploads\b/i, '/uploads')
  if (!/^https?:\/\//i.test(p) && !p.startsWith('/')) p = '/' + p
  if (!/\/uploads\//i.test(p) && /^\/(products|product-categories|brands|banners|customer)\//.test(p)) {
    p = '/uploads/media' + p
  }
  return p
}

// تابع فرمت کردن مدت زمان ویدیو
const formatDuration = (seconds) => {
  if (!seconds) return ''
  
  const hours = Math.floor(seconds / 3600)
  const minutes = Math.floor((seconds % 3600) / 60)
  const secs = seconds % 60
  
  if (hours > 0) {
    return `${hours}:${minutes.toString().padStart(2, '0')}:${secs.toString().padStart(2, '0')}`
  } else {
    return `${minutes}:${secs.toString().padStart(2, '0')}`
  }
}

// تابع فرمت کردن تاریخ
const formatDate = (dateString) => {
  if (!dateString) return ''
  
  const date = new Date(dateString)
  return date.toLocaleDateString('fa-IR', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}
</script>

<style scoped>
.product-videos-tab {
  padding: 1rem 0;
}

.main-video-section {
  border-bottom: 1px solid #e5e7eb;
  padding-bottom: 2rem;
}

.video-thumbnail {
  position: relative;
}

.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

/* حالت خالی حذف شد */

/* Responsive styles */
@media (max-width: 768px) {
  .grid {
    grid-template-columns: 1fr;
  }
  
  .main-video-section {
    padding-bottom: 1.5rem;
  }
}
</style> 
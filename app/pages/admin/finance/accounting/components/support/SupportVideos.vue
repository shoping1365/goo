<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">ویدیوهای آموزشی</h4>
        <p class="text-sm text-gray-600 mt-1">آموزش‌های تصویری برای اتصال و استفاده از سیستم حسابداری</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          @click="openVideoLibrary"
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
          </svg>
          کتابخانه ویدیو
        </button>
      </div>
    </div>

    <!-- فیلترها -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <div class="flex flex-col md:flex-row gap-6">
        <div class="flex-1">
          <input
            v-model="searchQuery"
            type="text"
            placeholder="جستجو در ویدیوها..."
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
        </div>
        <div class="flex gap-2">
          <select
            v-model="selectedCategory"
            class="px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه دسته‌ها</option>
            <option value="connection">اتصال</option>
            <option value="configuration">پیکربندی</option>
            <option value="sync">همگام‌سازی</option>
            <option value="troubleshooting">عیب‌یابی</option>
          </select>
          <select
            v-model="selectedLevel"
            class="px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه سطوح</option>
            <option value="beginner">مبتدی</option>
            <option value="intermediate">متوسط</option>
            <option value="advanced">پیشرفته</option>
          </select>
        </div>
      </div>
    </div>

    <!-- ویدیوهای پیشنهادی -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">ویدیوهای پیشنهادی</h5>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div
          v-for="video in featuredVideos"
          :key="video.id"
          class="bg-gray-50 rounded-lg overflow-hidden hover:shadow-md transition-shadow"
        >
          <div class="relative">
            <img
              :src="video.thumbnail"
              :alt="video.title"
              class="w-full h-48 object-cover"
            />
            <div class="absolute inset-0 bg-black bg-opacity-20 flex items-center justify-center">
              <button
                @click="playVideo(video)"
                class="bg-white bg-opacity-90 rounded-full p-3 hover:bg-opacity-100 transition-all"
              >
                <svg class="w-6 h-6 text-gray-800" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M8 5v14l11-7z"/>
                </svg>
              </button>
            </div>
            <div class="absolute top-2 right-2">
              <span
                class="px-2 py-1 text-xs font-medium rounded-full"
                :class="getLevelClass(video.level)"
              >
                {{ getLevelLabel(video.level) }}
              </span>
            </div>
          </div>
          <div class="p-6">
            <h6 class="font-medium text-gray-900 mb-2">{{ video.title }}</h6>
            <p class="text-sm text-gray-600 mb-3">{{ video.description }}</p>
            <div class="flex items-center justify-between text-xs text-gray-500">
              <span>{{ video.duration }}</span>
              <span>{{ video.views }} بازدید</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- لیست کامل ویدیوها -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">همه ویدیوها</h5>
      </div>
      <div class="divide-y divide-gray-200">
        <div
          v-for="video in filteredVideos"
          :key="video.id"
          class="p-6 hover:bg-gray-50 transition-colors"
        >
          <div class="flex items-start space-x-4 space-x-reverse">
            <div class="flex-shrink-0">
              <img
                :src="video.thumbnail"
                :alt="video.title"
                class="w-32 h-20 object-cover rounded-lg"
              />
            </div>
            <div class="flex-1 min-w-0">
              <div class="flex items-center justify-between">
                <h6 class="text-sm font-medium text-gray-900 truncate">{{ video.title }}</h6>
                <div class="flex items-center space-x-2 space-x-reverse">
                  <span
                    class="px-2 py-1 text-xs font-medium rounded-full"
                    :class="getLevelClass(video.level)"
                  >
                    {{ getLevelLabel(video.level) }}
                  </span>
                  <span
                    class="px-2 py-1 text-xs font-medium rounded-full bg-blue-100 text-blue-800"
                  >
                    {{ getCategoryLabel(video.category) }}
                  </span>
                </div>
              </div>
              <p class="text-sm text-gray-600 mt-1">{{ video.description }}</p>
              <div class="flex items-center justify-between mt-3">
                <div class="flex items-center space-x-4 space-x-reverse text-xs text-gray-500">
                  <span>{{ video.duration }}</span>
                  <span>{{ video.views }} بازدید</span>
                  <span>{{ video.uploadDate }}</span>
                </div>
                <button
                  @click="playVideo(video)"
                  class="inline-flex items-center px-3 py-1 border border-transparent text-xs font-medium rounded-md text-blue-700 bg-blue-100 hover:bg-blue-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
                >
                  <svg class="w-3 h-3 ml-1" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M8 5v14l11-7z"/>
                  </svg>
                  پخش
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- مودال پخش ویدیو -->
    <div
      v-if="showVideoModal"
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50"
      @click="closeVideoModal"
    >
      <div class="bg-white rounded-lg p-6 max-w-4xl w-full mx-4" @click.stop>
        <div class="flex items-center justify-between mb-4">
          <h5 class="text-lg font-medium text-gray-900">{{ selectedVideo?.title }}</h5>
          <button
            @click="closeVideoModal"
            class="text-gray-400 hover:text-gray-600"
          >
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
        <div class="aspect-video bg-gray-900 rounded-lg flex items-center justify-center">
          <div class="text-white text-center">
            <svg class="w-16 h-16 mx-auto mb-4" fill="currentColor" viewBox="0 0 24 24">
              <path d="M8 5v14l11-7z"/>
            </svg>
            <p>پخش ویدیو: {{ selectedVideo?.title }}</p>
          </div>
        </div>
        <div class="mt-4">
          <p class="text-sm text-gray-600">{{ selectedVideo?.description }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'

// جستجو و فیلترها
const searchQuery = ref('')
const selectedCategory = ref('')
const selectedLevel = ref('')

// مودال ویدیو
const showVideoModal = ref(false)
const selectedVideo = ref(null)

// ویدیوهای پیشنهادی
const featuredVideos = ref([
  {
    id: 1,
    title: 'راهنمای کامل اتصال به هلو',
    description: 'آموزش گام‌به‌گام اتصال سیستم به نرم‌افزار هلو',
    thumbnail: '/images/video-thumb-1.jpg',
    duration: '۱۵:۳۰',
    views: 1250,
    level: 'beginner',
    category: 'connection'
  },
  {
    id: 2,
    title: 'پیکربندی پیشرفته همگام‌سازی',
    description: 'تنظیمات پیشرفته برای همگام‌سازی داده‌ها',
    thumbnail: '/images/video-thumb-2.jpg',
    duration: '۲۳:۱۵',
    views: 890,
    level: 'advanced',
    category: 'sync'
  },
  {
    id: 3,
    title: 'عیب‌یابی مشکلات رایج',
    description: 'راه‌حل مشکلات معمول در اتصال حسابداری',
    thumbnail: '/images/video-thumb-3.jpg',
    duration: '۱۸:۴۵',
    views: 2100,
    level: 'intermediate',
    category: 'troubleshooting'
  }
])

interface Video {
  id: number
  title: string
  description: string
  thumbnail: string
  duration: string
  views: number
  level: string
  category: string
  uploadDate?: string
}

// همه ویدیوها
const allVideos = ref<Video[]>([
  ...featuredVideos.value,
  {
    id: 4,
    title: 'تنظیم امنیت اتصال',
    description: 'آموزش تنظیمات امنیتی برای اتصال حسابداری',
    thumbnail: '/images/video-thumb-4.jpg',
    duration: '۱۲:۲۰',
    views: 750,
    level: 'intermediate',
    category: 'configuration',
    uploadDate: '۲ روز پیش'
  },
  {
    id: 5,
    title: 'بهینه‌سازی عملکرد',
    description: 'راه‌های بهبود سرعت و عملکرد اتصال',
    thumbnail: '/images/video-thumb-5.jpg',
    duration: '۲۰:۱۰',
    views: 1100,
    level: 'advanced',
    category: 'configuration',
    uploadDate: '۱ هفته پیش'
  }
])

// فیلتر کردن ویدیوها
const filteredVideos = computed(() => {
  return allVideos.value.filter(video => {
    if (searchQuery.value && !video.title.includes(searchQuery.value)) return false
    if (selectedCategory.value && video.category !== selectedCategory.value) return false
    if (selectedLevel.value && video.level !== selectedLevel.value) return false
    return true
  })
})

// کلاس سطح
const getLevelClass = (level: string) => {
  const classes = {
    beginner: 'bg-green-100 text-green-800',
    intermediate: 'bg-yellow-100 text-yellow-800',
    advanced: 'bg-red-100 text-red-800'
  }
  return classes[level] || 'bg-gray-100 text-gray-800'
}

// برچسب سطح
const getLevelLabel = (level: string) => {
  const labels = {
    beginner: 'مبتدی',
    intermediate: 'متوسط',
    advanced: 'پیشرفته'
  }
  return labels[level] || level
}

// برچسب دسته
const getCategoryLabel = (category: string) => {
  const labels = {
    connection: 'اتصال',
    configuration: 'پیکربندی',
    sync: 'همگام‌سازی',
    troubleshooting: 'عیب‌یابی'
  }
  return labels[category] || category
}

// پخش ویدیو
const playVideo = (video: any) => {
  selectedVideo.value = video
  showVideoModal.value = true
}

// بستن مودال
const closeVideoModal = () => {
  showVideoModal.value = false
  selectedVideo.value = null
}

// باز کردن کتابخانه ویدیو
const openVideoLibrary = () => {
  // TODO: باز کردن کتابخانه ویدیو
  console.log('باز کردن کتابخانه ویدیو')
}
</script>

<!--
  کامپوننت ویدیوهای آموزشی
  شامل:
  1. ویدیوهای پیشنهادی
  2. لیست کامل ویدیوها
  3. فیلتر و جستجو
  4. مودال پخش ویدیو
  5. طراحی مدرن و کاملاً ریسپانسیو
--> 

<template>
  <div class="icon-selector">
    <!-- Tabs for Icon/Image Selection -->
    <div class="mb-4">
      <div class="flex space-x-1 space-x-reverse">
        <button
          @click="activeTab = 'icon'"
          :class="[
            'flex-1 px-4 py-2 text-sm font-medium rounded-lg transition-colors',
            activeTab === 'icon'
              ? 'bg-blue-100 text-blue-700'
              : 'bg-gray-100 text-gray-700 hover:bg-gray-200'
          ]"
        >
          آیکون
        </button>
        <button
          @click="activeTab = 'image'"
          :class="[
            'flex-1 px-4 py-2 text-sm font-medium rounded-lg transition-colors',
            activeTab === 'image'
              ? 'bg-blue-100 text-blue-700'
              : 'bg-gray-100 text-gray-700 hover:bg-gray-200'
          ]"
        >
          تصویر
        </button>
      </div>
    </div>

    <!-- Icon Selection Tab -->
    <div v-if="activeTab === 'icon'">
      <!-- Icon Categories -->
      <div class="mb-4">
        <div class="flex space-x-2 space-x-reverse overflow-x-auto pb-2">
          <button
            v-for="category in iconCategories"
            :key="category.name"
            @click="selectedCategory = category.name"
            :class="[
              'px-3 py-2 rounded-lg text-sm font-medium transition-colors whitespace-nowrap',
              selectedCategory === category.name
                ? 'bg-blue-100 text-blue-700'
                : 'bg-gray-100 text-gray-700 hover:bg-gray-200'
            ]"
          >
            {{ category.label }}
          </button>
        </div>
      </div>

      <!-- Icons Grid -->
      <div class="grid grid-cols-6 gap-2 max-h-64 overflow-y-auto">
        <button
          v-for="icon in filteredIcons"
          :key="icon"
          @click="selectIcon(icon)"
          class="p-3 border border-gray-200 rounded-lg hover:bg-gray-50 transition-colors flex items-center justify-center"
          :class="{ 'bg-blue-100 border-blue-300': selectedIcon === icon }"
        >
          <i :class="icon" class="text-lg text-gray-600"></i>
        </button>
      </div>

      <!-- Icon Search -->
      <div class="mt-4">
        <input
          v-model="searchQuery"
          type="text"
          placeholder="جستجو در آیکون‌ها..."
          class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 text-sm"
        />
      </div>
    </div>

    <!-- Image Upload Tab -->
    <div v-if="activeTab === 'image'">
      <!-- Current Image Preview -->
      <div v-if="selectedImage" class="mb-4">
        <div class="relative inline-block">
          <img
            :src="selectedImage"
            alt="تصویر انتخاب شده"
            class="w-20 h-20 object-cover rounded-lg border border-gray-200"
          />
          <button
            @click="removeSelectedImage"
            class="absolute -top-2 -right-2 bg-red-500 text-white rounded-full w-6 h-6 flex items-center justify-center text-xs hover:bg-red-600 transition-colors"
          >
            ×
          </button>
        </div>
      </div>

      <!-- Image Upload Area -->
      <div class="mb-4">
        <div
          @click="triggerFileInput"
          @dragover.prevent
          @drop.prevent="handleFileDrop"
          class="border-2 border-dashed border-gray-300 rounded-lg p-6 text-center cursor-pointer hover:border-blue-400 transition-colors"
        >
          <div class="text-gray-500">
            <svg class="mx-auto h-12 w-12 mb-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12" />
            </svg>
            <p class="text-sm font-medium">برای آپلود کلیک کنید یا فایل را اینجا بکشید</p>
            <p class="text-xs text-gray-400 mt-1">PNG, JPG, GIF تا 2MB</p>
          </div>
        </div>
        <input
          ref="fileInput"
          type="file"
          accept="image/*"
          class="hidden"
          @change="handleFileSelect"
        />
      </div>

      <!-- Image URL Input -->
      <div class="mb-4">
        <label class="block text-sm font-medium text-gray-700 mb-2">یا آدرس تصویر را وارد کنید</label>
        <div class="flex space-x-2 space-x-reverse">
          <input
            v-model="imageUrl"
            type="url"
            placeholder="https://example.com/image.jpg"
            class="flex-1 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 text-sm"
          />
          <button
            @click="loadImageFromUrl"
            :disabled="!imageUrl"
            class="px-4 py-2 bg-blue-100 hover:bg-blue-200 text-blue-700 rounded-md text-sm font-medium transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
          >
            بارگذاری
          </button>
        </div>
      </div>

      <!-- Image Settings -->
      <div v-if="selectedImage" class="space-y-3">
        <h4 class="font-medium text-gray-900">تنظیمات تصویر</h4>
        
        <!-- Image Size -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">اندازه</label>
          <select
            v-model="imageSize"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 text-sm"
          >
            <option value="small">کوچک (16px)</option>
            <option value="medium">متوسط (24px)</option>
            <option value="large">بزرگ (32px)</option>
            <option value="custom">سفارشی</option>
          </select>
        </div>

        <!-- Custom Size -->
        <div v-if="imageSize === 'custom'">
          <label class="block text-sm font-medium text-gray-700 mb-1">اندازه سفارشی (px)</label>
          <input
            v-model="customImageSize"
            type="number"
            min="8"
            max="100"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 text-sm"
          />
        </div>

        <!-- Image Style -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">سبک</label>
          <select
            v-model="imageStyle"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 text-sm"
          >
            <option value="default">پیش‌فرض</option>
            <option value="rounded">گرد</option>
            <option value="circle">دایره</option>
            <option value="square">مربع</option>
          </select>
        </div>
      </div>
    </div>

    <!-- Actions -->
    <div class="mt-4 flex space-x-2 space-x-reverse">
      <button
        @click="confirmSelection"
        :disabled="!hasSelection"
        class="flex-1 bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-lg font-medium transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
      >
        انتخاب
      </button>
      <button
        @click="closeModal"
        class="px-4 py-2 border border-gray-300 text-gray-700 rounded-lg font-medium hover:bg-gray-50 transition-colors"
      >
        انصراف
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'

const props = defineProps({
  modelValue: {
    type: String,
    default: ''
  }
})

const emit = defineEmits(['update:modelValue', 'close'])

const activeTab = ref('icon')
const selectedIcon = ref(props.modelValue)
const selectedCategory = ref('all')
const searchQuery = ref('')

// Image related states
const selectedImage = ref('')
const imageUrl = ref('')
const imageSize = ref('medium')
const customImageSize = ref(24)
const imageStyle = ref('default')
const fileInput = ref(null)

const iconCategories = [
  { name: 'all', label: 'همه' },
  { name: 'navigation', label: 'ناوبری' },
  { name: 'social', label: 'شبکه‌های اجتماعی' },
  { name: 'business', label: 'کسب و کار' },
  { name: 'communication', label: 'ارتباطات' }
]

const allIcons = {
  navigation: [
    'fas fa-home',
    'fas fa-arrow-right',
    'fas fa-arrow-left',
    'fas fa-chevron-down',
    'fas fa-chevron-up',
    'fas fa-bars',
    'fas fa-times',
    'fas fa-search',
    'fas fa-map-marker-alt'
  ],
  social: [
    'fab fa-facebook',
    'fab fa-twitter',
    'fab fa-instagram',
    'fab fa-linkedin',
    'fab fa-youtube',
    'fab fa-telegram',
    'fab fa-whatsapp'
  ],
  business: [
    'fas fa-shopping-cart',
    'fas fa-store',
    'fas fa-box',
    'fas fa-truck',
    'fas fa-credit-card',
    'fas fa-wallet',
    'fas fa-chart-line',
    'fas fa-users'
  ],
  communication: [
    'fas fa-envelope',
    'fas fa-phone',
    'fas fa-comment',
    'fas fa-comments',
    'fas fa-bell',
    'fas fa-user',
    'fas fa-user-circle',
    'fas fa-cog'
  ]
}

const filteredIcons = computed(() => {
  let icons = []
  
  if (selectedCategory.value === 'all') {
    Object.values(allIcons).forEach(categoryIcons => {
      icons = icons.concat(categoryIcons)
    })
  } else {
    icons = allIcons[selectedCategory.value] || []
  }
  
  if (searchQuery.value) {
    icons = icons.filter(icon => 
      icon.toLowerCase().includes(searchQuery.value.toLowerCase())
    )
  }
  
  return icons
})

const hasSelection = computed(() => {
  if (activeTab.value === 'icon') {
    return !!selectedIcon.value
  } else {
    return !!selectedImage.value
  }
})

const selectIcon = (icon) => {
  selectedIcon.value = icon
  selectedImage.value = '' // Clear image when icon is selected
}

const triggerFileInput = () => {
  fileInput.value.click()
}

const handleFileSelect = (event) => {
  const file = event.target.files[0]
  if (file) {
    processImageFile(file)
  }
}

const handleFileDrop = (event) => {
  const file = event.dataTransfer.files[0]
  if (file && file.type.startsWith('image/')) {
    processImageFile(file)
  }
}

const processImageFile = (file) => {
  // Check file size (2MB limit)
  if (file.size > 2 * 1024 * 1024) {
    alert('حجم فایل نباید بیشتر از 2 مگابایت باشد')
    return
  }

  const reader = new FileReader()
  reader.onload = (e) => {
    selectedImage.value = e.target.result
    selectedIcon.value = '' // Clear icon when image is selected
  }
  reader.readAsDataURL(file)
}

const loadImageFromUrl = () => {
  if (imageUrl.value) {
    selectedImage.value = imageUrl.value
    selectedIcon.value = '' // Clear icon when image is selected
  }
}

const removeSelectedImage = () => {
  selectedImage.value = ''
  imageUrl.value = ''
}

const confirmSelection = () => {
  let result = ''
  
  if (activeTab.value === 'icon') {
    result = selectedIcon.value
  } else if (activeTab.value === 'image') {
    // Create image object with metadata
    const imageData = {
      type: 'image',
      url: selectedImage.value,
      size: imageSize.value === 'custom' ? customImageSize.value : imageSize.value,
      style: imageStyle.value
    }
    result = JSON.stringify(imageData)
  }
  
  emit('update:modelValue', result)
  emit('close')
}

const closeModal = () => {
  emit('close')
}
</script>

<style scoped>
.icon-selector {
  max-width: 500px;
  margin: 0 auto;
}
</style>
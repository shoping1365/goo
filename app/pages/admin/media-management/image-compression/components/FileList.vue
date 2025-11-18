<template>
  <div>
    <!-- Action Bar -->
    <div class="flex items-center gap-2 mb-4 select-none">
      <button @click="scanImages" class="bg-purple-100 text-purple-700 text-xs md:text-sm px-3 py-1 rounded-lg hover:bg-purple-200 transition-colors">اسکن</button>
      <button v-if="selectedImages.length>0 && selectedImages.length<images.length" @click="selectAllImages" class="bg-indigo-100 text-indigo-700 text-xs md:text-sm px-3 py-1 rounded-lg hover:bg-indigo-200 transition-colors">انتخاب همه</button>
      <button v-if="selectedImages.length>0" :disabled="isCompressing" @click="$emit('compress')" class="bg-pink-100 text-pink-700 text-xs md:text-sm px-3 py-1 rounded-lg hover:bg-pink-200 transition-colors disabled:opacity-50">فشرده‌سازی</button>
    </div>

    <!-- Images Grid -->
    <div v-if="images.length" class="grid grid-cols-3 md:grid-cols-6 lg:grid-cols-8 xl:grid-cols-10 gap-3">
      <div v-for="img in images" :key="img.id" :class="containerClass(img.id)" class="relative group cursor-pointer rounded-xl border-2 transition duration-200 overflow-hidden shadow-md">
        <!-- Preview -->
        <div class="aspect-square bg-gray-100 flex items-center justify-center relative" @click="openPreview(img)">
          <img :src="img.thumbnail||img.url" :alt="img.name" class="w-full h-full object-cover" />
        </div>
        <!-- Info -->
        <div class="p-3 bg-white" @click.stop="handleImageSelect(img.id,$event)">
          <p class="text-sm font-medium text-gray-900 truncate" :title="img.name">{{ img.name }}</p>
          <div class="flex items-center justify-between mt-1">
            <p class="text-xs text-gray-500">{{ formatFileSize(img.size) }}</p>
            <p class="text-xs text-gray-500">{{ img.dimensions }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// Auto-imported composable
// @ts-ignore
const { images, selectedImages, isCompressing, scanImages, selectAllImages, handleImageSelect, openCompareModal } = useImageCompression()

function containerClass(id:number){
  return selectedImages.includes(id) ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-gray-300'
}

function formatFileSize(bytes:number){
  if(!bytes && bytes!==0) return ''
  const sizes=['B','KB','MB','GB']
  let idx=0
  let val=bytes
  while(val>=1024 && idx<sizes.length-1){val/=1024;idx++}
  return `${val.toFixed( idx===0 ?0:1)} ${sizes[idx]}`
}

function openPreview(img:any){
  // در این نسخه فقط مقایسه ساده نمایش می‌دهیم
  openCompareModal({
    id: img.id,
    name: img.name,
    thumbnail: img.thumbnail||img.url,
    dimensions: img.dimensions||'',
    originalSize: img.size,
    compressedSize: img.size,
    reduction: 0,
    status: 'success',
    originalUrl: img.url,
    compressedUrl: img.url,
  } as any)
}
</script> 

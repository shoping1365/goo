<template>
  <div v-if="showCompareModal" class="fixed inset-0 bg-black bg-opacity-75 flex items-center justify-center z-50" @click="closeCompareModal">
    <div class="relative bg-white rounded-lg overflow-hidden shadow-lg w-[95vw] max-w-none max-h-[95vh] overflow-auto" @click.stop>
      <!-- Close btn -->
      <button class="absolute left-4 top-6 bg-white bg-opacity-80 rounded-full p-1 shadow hover:bg-opacity-100 transition-all" @click="closeCompareModal">
        <svg class="w-6 h-6 text-gray-700" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" /></svg>
      </button>
      <div class="flex flex-col sm:flex-row">
        <!-- Original -->
        <div class="flex-1 p-6 flex flex-col items-center border-b sm:border-b-0 sm:border-r">
          <span class="font-bold text-gray-700 mb-2">عکس اصلی</span>
          <img :src="compareModalData?.originalUrl" :alt="compareModalData?.name" class="w-full max-h-[80vh] object-contain" />
          <span v-if="compareModalData" class="text-sm text-gray-600 mt-2">{{ formatFileSize(compareModalData.originalSize) }}</span>
        </div>
        <!-- Compressed -->
        <div class="flex-1 p-6 flex flex-col items-center">
          <span class="font-bold text-gray-700 mb-2">عکس فشرده‌شده</span>
          <img :src="compareModalData?.compressedUrl" :alt="compareModalData?.name" class="w-full max-h-[80vh] object-contain" />
          <span v-if="compareModalData" class="text-sm text-gray-600 mt-2">{{ formatFileSize(compareModalData.compressedSize) }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useImageCompression } from '~/composables/useImageCompression'
const { showCompareModal, compareModalData, closeCompareModal } = useImageCompression()

function formatFileSize(bytes:number){
  if(!bytes&&bytes!==0) return ''
  const sizes=['B','KB','MB','GB']
  let idx=0
  let val=bytes
  while(val>=1024&&idx<sizes.length-1){val/=1024;idx++}
  return `${val.toFixed(idx===0?0:1)} ${sizes[idx]}`
}
</script> 

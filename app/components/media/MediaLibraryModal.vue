<template>
  <div v-if="modelValue" class="fixed inset-0 bg-black bg-opacity-60 z-50 flex items-center justify-center p-6">
    <div class="bg-white rounded-2xl shadow-2xl w-full max-w-[1200px] h-[80vh] flex flex-col overflow-hidden relative border border-gray-200">
      <!-- Header -->
      <div class="flex flex-row items-center p-6 border-b gap-6 bg-gradient-to-l from-blue-50 to-white">
        <!-- Title on right -->
        <h3 class="text-2xl font-extrabold text-blue-800 tracking-tight flex-shrink-0">کتابخانه رسانه</h3>
        <div class="flex-1"></div>
        <!-- Browse upload button on far left -->
        <div>
          <input ref="fileInput" type="file" multiple accept="image/*,video/*" class="hidden" @change="handleFileChange" />
          <button class="px-5 py-2 rounded-lg bg-blue-500 hover:bg-blue-600 text-white text-base flex items-center gap-2 font-semibold shadow-md transition-all duration-200" @click="triggerBrowse">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
            </svg>
            افزودن
          </button>
        </div>
      </div>
      <!-- Filters -->
      <div class="p-6 flex flex-wrap items-center gap-6 border-b text-base bg-gray-50">
        <select v-model="selectedCategory" class="input input-bordered bg-white rounded-lg px-4 py-2 border border-gray-300 focus:ring-2 focus:ring-blue-400">
          <option value="all">همه</option>
          <option value="library">کتابخانه</option>
          <option value="products">محصولات</option>
          <option value="product-categories">دسته‌بندی محصولات</option>
          <option value="brands">برندها</option>
          <option value="banners">بنرهای تبلیغاتی</option>
          <option value="customer">مشتریان</option>
        </select>
        <input v-model="search" placeholder="جستجو..." class="input input-bordered flex-1 min-w-[200px] rounded-lg px-4 py-2 border border-gray-300 focus:ring-2 focus:ring-blue-400" />
      </div>
      <!-- Content -->
      <div class="flex-1 overflow-auto p-6 media-scroll bg-white">
        <ImageSelector
          :images="filtered"
          :selected-images="internalSelected"
          :format-file-size="formatFileSize"
          @select="toggleSelect"
          @preview="onPreview"
        />
      </div>
      <!-- Footer -->
      <div class="p-6 border-t flex flex-row-reverse items-center gap-2 justify-between bg-gray-50">
        <!-- انتخاب button on right -->
        <button
          class="px-10 py-3 rounded-xl font-bold flex items-center gap-2 disabled:opacity-50 bg-emerald-400 text-emerald-900 hover:bg-emerald-500 shadow-md text-lg transition-all duration-200"
          :disabled="!internalSelected.length || isUploading"
          @click="confirm"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
          </svg>
          انتخاب<span v-if="internalSelected.length"> ({{ internalSelected.length }})</span>
        </button>
        <div v-if="isUploading" class="flex-1 flex items-center justify-center">
          <div class="w-full max-w-sm h-2 bg-gray-200 rounded overflow-hidden mx-4 relative">
            <div class="h-full bg-blue-500 transition-all absolute right-0 top-0" :style="{ width: uploadProgress + '%' }"></div>
          </div>
          <span class="text-xs text-gray-600 w-12">{{ uploadProgress }}%</span>
        </div>
        <!-- per-file progress small list (only unfinished) -->
        <div v-if="uploads.some(u=>u.progress<100)" class="flex flex-col gap-1 w-full max-w-xs mx-4 text-right" dir="rtl">
          <div v-for="u in uploads.filter(x=>x.progress<100)" :key="u.name" class="flex items-center gap-2">
            <span class="text-[10px] truncate flex-1">{{ u.name }}</span>
            <div class="h-1 bg-gray-200 rounded flex-1 relative">
              <div class="absolute right-0 top-0 h-full bg-green-500" :style="{ width: u.progress + '%' }"></div>
            </div>
            <span class="text-[10px] text-gray-600 w-6">{{ u.progress }}%</span>
          </div>
        </div>
        <!-- انصراف button on left -->
        <button
          class="px-8 py-3 rounded-xl font-bold flex items-center gap-2 bg-rose-300 text-rose-900 hover:bg-rose-400 shadow-md text-lg transition-all duration-200"
          :disabled="isUploading"
          @click="$emit('update:modelValue', false)"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
          انصراف
        </button>
      </div>
      <MediaPreviewModal 
        v-if="previewShow && preview" 
        v-model="previewShow" 
        :file="preview" 
        :can-delete="canDeleteMedia" 
        @delete="onFileDeleted" 
        @save="onFileSaved"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch, reactive, nextTick, triggerRef } from 'vue'
import { useRequestHeaders } from 'nuxt/app'
import { useAuth } from '~/composables/useAuth'

import ImageSelector from '~/components/media/ImageSelector.vue'

// نرمال‌سازی مسیر فایل‌ها برای نمایش صحیح
function normalizePath(p:string){
  // مطابق صفحه کتابخانه: فقط بک‌اسلش‌ها را به اسلش تبدیل و یک اسلش ابتدایی تضمین می‌کنیم
  const norm = (p || '').replace(/\\/g,'/')
  return norm.startsWith('/') ? norm : `/${norm}`
}
import MediaPreviewModal from '~/components/media/MediaPreviewModal.vue'

// استفاده از useAuth برای چک کردن پرمیژن‌ها
interface UseAuthReturn {
  hasPermission: (permission: string) => boolean
}

const { hasPermission } = useAuth() as UseAuthReturn

// Computed برای چک کردن پرمیژن حذف
const canDeleteMedia = computed(() => hasPermission('media_library.delete'))

interface MediaFile {
  id:number
  url:string
  thumbnail:string
  type:string
  name:string
  size:number
  category:string
  extension?: string
  file_path?: string
  mime_type?: string
  file_type?: string
  original_name?: string
  file_name?: string
  file_size?: number
  [key: string]: unknown
}

const props = defineProps<{ modelValue:boolean, modelSelected?:number[], defaultCategory?:string, fileType?:string, contextTitle?: string, maxSelect?: number }>()
const emit = defineEmits(['update:modelValue','update:selected','confirm','media-uploaded'])


const list = ref<MediaFile[]>([])
const selectedCategory = ref(props.defaultCategory || 'product-categories')
const search = ref('')

// uploads in progress {name, progress}
const uploads = reactive<{name:string; progress:number}[]>([])

const internalSelected = ref<number[]>(props.modelSelected ? [...props.modelSelected] : [])
watch(()=>props.modelSelected,val=>{ if(val) internalSelected.value=[...val]})
watch(selectedCategory, ()=>{ fetchList() })
watch(()=>props.fileType, ()=>{
  // پیش‌فرض همیشگی: دسته‌بندی محصولات (طبق درخواست)
  if(!props.defaultCategory){
    selectedCategory.value = 'product-categories'
  }
})

// Reset selection when modal is closed
watch(() => props.modelValue, (newVal) => {
  if (!newVal) {
    // Modal is closing, reset selection
    internalSelected.value = []
  }
})

function formatFileSize(bytes:number){ const k=1024; const sizes=['B','KB','MB','GB']; const i=Math.floor(Math.log(bytes)/Math.log(k)); return (bytes/Math.pow(k,i)).toFixed(1)+' '+sizes[i] }

async function fetchList(){
  const catQuery = (['products','customer','product-categories','brands','banners'].includes(selectedCategory.value)) ? `?category=${selectedCategory.value}` : ''
  const url = `/api/media/list${catQuery}`
  
  try {
    const headers = process.server ? useRequestHeaders(['cookie', 'authorization']) : undefined
    const response = await $fetch(url, {
      credentials: 'include',
      headers
    })
    interface MediaResponse {
      data?: MediaFile[]
    }
    
    const raw: MediaFile[] = Array.isArray(response) 
      ? response 
      : ((response as MediaResponse)?.data || [])
    
    interface MediaFileRaw {
      url?: string
      file_path?: string
      mime_type?: string
      file_type?: string
      original_name?: string
      file_name?: string
      size?: number
      file_size?: number
      category?: string
      [key: string]: unknown
    }
    
    list.value = raw.map((r: MediaFileRaw)=> { 
      const p = (r.url || r.file_path || '') as string
      let path = normalizePath(p)
      // سازگارسازی مسیرهای ویدیوهای products به مسیر سروینگ public
      path = path.replace(/\/public\//,'/')
      if(!/^https?:\/\//i.test(path)){
        if(!path.startsWith('/')) path = '/' + path
      }
      if(!/\/uploads\//.test(path) && /^\/(products|product-categories|brands|banners|customer|library)\//.test(path)){
        path = '/uploads/media' + path
      }
      const mime = (r.mime_type || r.file_type || '') as string
      const ext = (mime.split('/')?.pop()||'').toLowerCase()
      const isImg = mime.startsWith('image')||['webp','jpg','jpeg','png','gif'].includes(ext)
      const isVid = mime.startsWith('video')||['mp4','webm','mov','avi'].includes(ext)
      const isGif = ext==='gif'
      return { 
        id: r.id as number,
        url:path,
        thumbnail:isImg && !isGif ? path : '',
        type:isImg?'image':(isVid?'video':'other'),
        name:(r.original_name || r.file_name || '') as string,
        size:(r.size || r.file_size || 0) as number,
        extension: ext,
        category: (
          (r.category || '') as string ||
          (p.includes('/product-categories/') ? 'product-categories' :
          p.includes('/brands/') ? 'brands' :
          p.includes('/banners/') ? 'banners' :
          p.includes('/products/') ? 'products' :
          p.includes('/customer/') ? 'customer' : 'library')
        ) as string
      } as MediaFile
    })

  } catch (error) {
    console.error('خطا در fetchList:', error)
  }
}
fetchList()

const filtered = computed(()=>{

  let arr=list.value.filter(i=>{
    if (props.fileType === 'video') return i.type==='video';
    if (props.fileType === 'image') return i.type==='image';
    return i.type==='image' || i.type==='video';
  })

  // فیلتر بر اساس دسته‌بندی
  if(selectedCategory.value==='library') arr=arr.filter(i=>i.category==='library')
  else if(selectedCategory.value==='customer') arr=arr.filter(i=>i.category==='customer')
  else if(selectedCategory.value==='products') arr=arr.filter(i=>i.category==='products')
  else if(selectedCategory.value==='product-categories') arr=arr.filter(i=>i.category==='product-categories')
  else if(selectedCategory.value==='brands') arr=arr.filter(i=>i.category==='brands')
  else if(selectedCategory.value==='banners') arr=arr.filter(i=>i.category==='banners')
  // اگر 'all' انتخاب شده، فیلتر دسته‌بندی اعمال نکن
  else if(selectedCategory.value==='all') {

  }

  // فیلتر بر اساس جستجو
  if(search.value) arr=arr.filter(i=> i.name.toLowerCase().includes(search.value.toLowerCase()))

  return arr
})

function toggleSelect(id:number, event?: MouseEvent){
  // اگر ctrl/meta فشرده شده باشد، حالت انتخاب چندگانه است
  if(event && (event.ctrlKey || event.metaKey)) {
    // اگر maxSelect تنظیم شده باشد، امکان انتخاب چندگانه را محدود کن
    if(props.maxSelect && props.maxSelect > 1) {
      const idx = internalSelected.value.indexOf(id)
      if(idx > -1) {
        internalSelected.value.splice(idx, 1)
      } else if(internalSelected.value.length < props.maxSelect) {
        internalSelected.value.push(id)
      }
    } else {
      // اگر maxSelect تنظیم نشده یا 1 باشد، امکان انتخاب چندگانه را غیرفعال کن
      internalSelected.value = [id]
    }
  } else {
    // در غیر این صورت، فقط این آیتم را انتخاب کن (انتخاب قبلی را پاک کن)
    internalSelected.value = [id]
  }
}
function confirm(){ 
  emit('confirm', list.value.filter(i=>internalSelected.value.includes(i.id))); 
  emit('update:modelValue', false);
  // Reset selection after confirmation
  internalSelected.value = [];
}

const preview = ref(null)
const previewShow = ref(false)
function onPreview(file) {
  preview.value = JSON.parse(JSON.stringify(file))
  nextTick(() => {
    previewShow.value = false
    nextTick(()=>{ previewShow.value = true })
  })
}

// upload progress state
const isUploading = ref(false)
const uploadProgress = ref(0)

// upload browse logic
  const fileInput = ref<HTMLInputElement|null>(null)
function triggerBrowse(){ fileInput.value?.click() }

async function handleFileChange(e:Event){
  const files = (e.target as HTMLInputElement).files
  if(!files || files.length===0) return
  const filesArr = Array.from(files)
  if(filesArr.length===0) return

  // بررسی صحت ویدیوها قبل از آپلود
  for(const file of filesArr) {
    const isVideo = file.type.startsWith('video/')
    const isVideoExt = /\.(mp4|webm|mov|avi)$/i.test(file.name)
    if(isVideo || isVideoExt) {
      if(!(isVideo && isVideoExt)) {
        alert('فرمت یا نوع ویدیوی انتخابی معتبر نیست. فقط فایل‌های mp4, webm, mov, avi با type ویدیویی مجاز هستند.')
        return
      }
    }
  }

  // helper to build Media object
  interface MediaFileRaw {
    url?: string
    file_path?: string
    mime_type?: string
    file_type?: string
    original_name?: string
    file_name?: string
    size?: number
    file_size?: number
    category?: string
    [key: string]: unknown
  }
  
  function buildMedia(obj: MediaFileRaw): MediaFile {
    const p=(obj.url||obj.file_path||'') as string
    let path = normalizePath(p)
    path = path.replace(/\/public\//,'/')
    if(!/^https?:\/\//i.test(path)){
      if(!path.startsWith('/')) path = '/' + path
    }
    if(!/\/uploads\//.test(path) && /^\/(products|product-categories|brands|banners|customer|library)\//.test(path)){
      path = '/uploads/media' + path
    }
    const mime=(obj.mime_type||obj.file_type||'') as string
    const ext=(mime.split('/')?.pop()||'').toLowerCase()
    const isImg=mime.startsWith('image')||['webp','jpg','jpeg','png','gif'].includes(ext)
    const isVid=mime.startsWith('video')||['mp4','webm','mov','avi'].includes(ext)
    const isGif=ext==='gif'
    
    // اطمینان از اینکه category درست تنظیم شده
    let category = (obj.category || selectedCategory.value) as string
    if (category === 'all') {
      category = selectedCategory.value
    }
    
    const mediaFile: MediaFile = {
      id: (obj.id || 0) as number,
      url:path, 
      thumbnail:isImg && !isGif ? path : '', 
      type:isImg?'image':(isVid?'video':'other'), 
      name:(obj.original_name||obj.file_name||'') as string, 
      size:(obj.size||obj.file_size||0) as number, 
      category: category,
      extension: ext,
      file_path: obj.file_path,
      mime_type: obj.mime_type,
      file_type: obj.file_type,
      original_name: obj.original_name,
      file_name: obj.file_name,
      file_size: obj.file_size
    }
    return mediaFile
  }
  
  // reset states
  isUploading.value = true
  uploadProgress.value = 0

  // Reset selection after upload
  internalSelected.value = []

  const _completed = 0
  const _total = filesArr.length
  // completed and total are unused, but kept for potential future use

  const tasks = filesArr.map(file=>{
    const item = reactive({ name:file.name, progress:0 })
    uploads.push(item)

    return new Promise<void>((resolve)=>{
      const xhr = new XMLHttpRequest()
      xhr.open('POST','/api/media/upload')
  

      xhr.upload.onprogress = (e)=>{
        if(e.lengthComputable){
          item.progress = Math.round((e.loaded/e.total)*100)
          // overall progress as mean
          const sum = uploads.reduce((s,x)=>s+x.progress,0)
          uploadProgress.value = Math.round(sum / uploads.length)
        }
      }

      const fd=new FormData()
      fd.append('file', file)
      fd.append('category', selectedCategory.value)
      // پارامترهای اختیاری برای پردازش AI با تاخیر
      fd.append('context_title', props.contextTitle || document?.title || '')
      fd.append('context_lang', 'fa')
      fd.append('auto_ai_meta','1')

      xhr.onreadystatechange = ()=>{
        if(xhr.readyState===4){
          item.progress = 100
          // schedule removal of completed entry
          setTimeout(()=>{ const pos=uploads.indexOf(item); if(pos>-1) uploads.splice(pos,1) },0)

          if(xhr.status===200){
            try{
              const json = JSON.parse(xhr.responseText) as { success?: boolean; data?: unknown; files?: unknown[] }
              const payload = json.data || (Array.isArray(json.files)? json.files[0]: null)
              if(json.success && payload && typeof payload === 'object' && 'id' in payload){
                // insert minimal record immediately
                const temp = { ...buildMedia(payload as MediaFileRaw) }
                if(temp.type === 'image') {
                  temp.type = 'image'
                  if(!temp.thumbnail) temp.thumbnail = temp.url
                }

                // اضافه کردن فایل جدید به لیست
                list.value = [temp, ...list.value]

                // emit event برای اطلاع از آپلود فایل جدید
                emit('media-uploaded', temp)
                
                // اطمینان از اینکه لیست به‌روزرسانی شده
                // force reactivity update
                const currentList = [...list.value]
                list.value = []
                list.value = currentList

                // force trigger reactivity
                triggerRef(list)

                // scroll to top to reveal new item
                const el = document.querySelector('.media-scroll') as HTMLElement | null
                if(el) el.scrollTo({ top:0, behavior:'smooth' })

                // fetch complete record (with generated thumbnail) and update
                const payloadId = (payload as { id?: number | string }).id
                if (payloadId) {
                  $fetch(`/api/media/${payloadId}`).then((data)=>{
                    if(data && typeof data === 'object'){
                      const full = buildMedia(data as MediaFileRaw)
                      const p = list.value.findIndex(x=>x.id===full.id)
                      if(p>-1){
                        const newArr=[...list.value]
                        if(full.type === 'image') {
                          full.type = 'image'
                          if(!full.thumbnail) full.thumbnail=full.url
                        }
                        newArr[p]=full
                        list.value=newArr
                      }
                    }
                  }).catch((error: unknown) => {
                    console.error('خطا در دریافت فایل کامل:', error)
                  })
                }
              }
            } catch (error) {
              console.error('خطا در پردازش پاسخ:', error)
            }
          }

          // recompute overall
          const sum = uploads.reduce((s,x)=>s+x.progress,0)
          uploadProgress.value = uploads.length? Math.round(sum / uploads.length):100
          resolve()
        }
      }

      xhr.send(fd)
    })
  })

  await Promise.allSettled(tasks)
  // No need to refresh - we already have all media in local state
  // await fetchList() // Removed to prevent refresh issues
  isUploading.value = false
}

async function onFileDeleted(id:number){
  try {
    // حذف از سرور
    const response = await fetch(`/api/media/delete?id=${id}`, { method: 'DELETE' })
    if (!response.ok) {
      console.error('Failed to delete file from server')
      return
    }
    
    // حذف از لیست محلی
    const idx = list.value.findIndex(f=>f.id===id)
    if(idx>-1) list.value.splice(idx,1)
    
    // پاک کردن از انتخاب شده‌ها
    const selIdx = internalSelected.value.indexOf(id)
    if(selIdx>-1) internalSelected.value.splice(selIdx,1)
    
    // بستن preview
    previewShow.value = false
    preview.value = null
  } catch (error) {
    console.error('Error deleting file:', error)
  }
}

interface FileSavePayload {
  id: number
  [key: string]: unknown
}

function onFileSaved(payload: FileSavePayload) {
  const idx = list.value.findIndex(f=>f.id===payload.id)
  if(idx>-1){ list.value[idx] = { ...list.value[idx], ...payload } }
}
</script> 
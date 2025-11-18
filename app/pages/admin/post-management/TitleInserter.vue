<template>
  <div class="title-inserter">
    <!-- دکمه‌های درج عنوان و AI -->
    <div class="flex items-center gapx-4 py-4">
      <!-- دکمه درج عنوان -->
      <button 
        @click="showModal = true"
        class="inline-flex items-center px-4 py-2 border border-transparent rounded-lg text-sm font-medium text-white bg-gradient-to-r from-rose-400 to-pink-500 hover:from-rose-500 hover:to-pink-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-rose-500 transition-all duration-200 shadow-md hover:shadow-lg"
      >
        <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"></path>
        </svg>
        درج عنوان
      </button>

      <!-- دکمه AI -->
      <button 
        @click="handleAIClick"
        class="inline-flex items-center px-4 py-2 border border-transparent rounded-lg text-sm font-medium text-white bg-gradient-to-r from-blue-300 to-purple-400 hover:from-blue-400 hover:to-purple-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-400 transition-all duration-200 shadow-md hover:shadow-lg"
      >
        <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"></path>
        </svg>
        AI
      </button>
    </div>

    <!-- Modal درج عنوان -->
    <div v-if="showModal" class="fixed inset-0 flex items-center justify-center z-50">
      <div class="bg-gradient-to-br from-blue-50 to-indigo-100 rounded-lg w-full max-w-2xl mx-4 max-h-[90vh] overflow-hidden shadow-2xl border-2 border-blue-200">
        <div class="flex items-center justify-between px-4 py-4 border-b border-gray-200">
          <h3 class="text-lg font-semibold text-gray-900">افزودن عنوان جدید</h3>
          <button @click="closeModal" class="text-gray-400 hover:text-gray-600">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>
        
        <div class="px-4 py-4 overflow-y-auto max-h-[calc(90vh-120px)] bg-gradient-to-br from-blue-50 to-indigo-100">

          <!-- فرم افزودن عنوان -->
          <form @submit.prevent="addTitle">
            <div class="space-y-4">
              <!-- متن عنوان -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">عنوان بخش مقاله</label>
                <input 
                  v-model="titleForm.title"
                  type="text"
                  required
                  placeholder="مثال: مقدمه، روش تحقیق، نتایج، نتیجه‌گیری..."
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors text-right bg-white"
                  dir="rtl"
                  @keyup.enter="addTitle"
                >
              </div>
            </div>

            <!-- دکمه‌های عملیات -->
            <div class="flex items-center justify-end space-x-3 space-x-reverse mt-6">
              <button 
                type="button"
                @click="addTitle"
                :disabled="!titleForm.title.trim()"
                class="px-4 py-2 text-sm font-medium text-green-700 bg-green-100 border border-green-300 rounded-lg hover:bg-green-200 transition-colors disabled:opacity-50"
              >
                افزودن عنوان
              </button>
              <button 
                type="button"
                @click="closeModal"
                class="px-4 py-2 text-sm font-medium text-gray-700 bg-gray-100 border border-gray-300 rounded-lg hover:bg-gray-200 transition-colors"
              >
                بستن
              </button>
            </div>
          </form>

          <!-- لیست عناوین موجود -->
          <div v-if="extractedTitles.length > 0" class="mt-8">
            <h4 class="text-sm font-medium text-gray-900 mb-3">عناوین موجود در خلاصه نوشته</h4>
            <div class="space-y-2 max-h-40 overflow-y-auto">
              <div 
                v-for="(title, index) in extractedTitles" 
                :key="index"
                class="flex items-center justify-between p-3 bg-gray-50 border border-gray-200 rounded-lg"
              >
                <div class="flex items-center flex-1">
                  <span class="text-xs text-gray-500 bg-white px-2 py-1 rounded border ml-2">
                    {{ index + 1 }}
                  </span>
                  <div v-if="editingIndex !== index" class="flex-1 ml-3">
                    <span class="text-sm font-medium text-gray-900">{{ title }}</span>
                  </div>
                  <input 
                    v-else
                    v-model="editingTitle"
                    type="text"
                    class="flex-1 ml-3 px-2 py-1 text-sm border border-gray-300 rounded bg-white focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                    @keyup.enter="saveEdit(index)"
                    @keyup.esc="cancelEdit"
                    ref="editInput"
                  >
                </div>
                <div class="flex items-center space-x-1 space-x-reverse">
                  <button 
                    v-if="editingIndex !== index"
                    @click="startEdit(index, title)"
                    class="p-1 text-gray-400 hover:text-blue-600 transition-colors"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
                    </svg>
                  </button>
                  <button 
                    v-else
                    @click="saveEdit(index)"
                    class="p-1 text-gray-400 hover:text-green-600 transition-colors"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
                    </svg>
                  </button>
                  <button 
                    v-if="editingIndex !== index"
                    @click="removeTitle(index)"
                    class="p-1 text-gray-400 hover:text-red-600 transition-colors"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                    </svg>
                  </button>
                  <button 
                    v-else
                    @click="cancelEdit"
                    class="p-1 text-gray-400 hover:text-red-600 transition-colors"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                    </svg>
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, nextTick } from 'vue'

// Props
const props = defineProps<{
  excerpt: string
}>()

// Emits
const emit = defineEmits<{
  'update:excerpt': [excerpt: string]
}>()

// State
const showModal = ref(false)
const editingIndex = ref(-1)
const editingTitle = ref('')
const editInput = ref<HTMLInputElement>()

// فرم عنوان
const titleForm = reactive({
  title: ''
})

// استخراج عناوین از خلاصه نوشته
const extractedTitles = computed(() => {
  if (!props.excerpt) return []
  
  // تبدیل <br> به خط جدید و تقسیم بر اساس خط جدید
  const textWithNewlines = props.excerpt.replace(/<br\s*\/?>/gi, '\n')
  const lines = textWithNewlines.split('\n').map(line => line.trim()).filter(line => line.length > 0)
  
  // استخراج متن از HTML links
  const titles: string[] = []
  lines.forEach(line => {
    // اگر خط شامل HTML link است
    const linkMatch = line.match(/<a[^>]*>([^<]*)<\/a>/)
    if (linkMatch) {
      titles.push(linkMatch[1]) // متن داخل link
    } else if (line.length < 100 && !line.includes('http')) {
      // اگر خط ساده است و کوتاه
      titles.push(line)
    }
  })
  
  return titles
})

// افزودن عنوان جدید
const addTitle = () => {
  if (!titleForm.title.trim()) return

  const newTitle = titleForm.title.trim()
  const linkText = `<a href="#id-${extractedTitles.value.length + 1}" class="text-blue-600 hover:text-blue-800">${newTitle}</a>`
  
  // درج لینک در خلاصه نوشته (هر عنوان در یک خط جداگانه)
  let newExcerpt = props.excerpt || ''
  if (newExcerpt) {
    newExcerpt += `<br><br>${linkText}`
  } else {
    newExcerpt = linkText
  }
  
  // ارسال به والد
  emit('update:excerpt', newExcerpt)
  
  // ریست کردن فرم
  titleForm.title = ''
  
  // مودال باز می‌مونه تا کاربر بتونه چندین عنوان اضافه کنه
}

// شروع ویرایش عنوان
const startEdit = (index: number, title: string) => {
  editingIndex.value = index
  editingTitle.value = title
  nextTick(() => {
    editInput.value?.focus()
  })
}

// ذخیره ویرایش عنوان
const saveEdit = (index: number) => {
  if (!editingTitle.value.trim()) return
  
  const titles = extractedTitles.value
  const oldTitle = titles[index]
  const newTitle = editingTitle.value.trim()
  
  // جایگزینی عنوان در خلاصه نوشته
  let newExcerpt = props.excerpt
  const oldLinkPattern = new RegExp(`<a href="#id-${index + 1}"[^>]*>${oldTitle}</a>`, 'g')
  const newLinkText = `<a href="#id-${index + 1}" class="text-blue-600 hover:text-blue-800">${newTitle}</a>`
  newExcerpt = newExcerpt.replace(oldLinkPattern, newLinkText)
  
  emit('update:excerpt', newExcerpt)
  
  // خروج از حالت ویرایش
  editingIndex.value = -1
  editingTitle.value = ''
}

// لغو ویرایش
const cancelEdit = () => {
  editingIndex.value = -1
  editingTitle.value = ''
}

// حذف عنوان
const removeTitle = (index: number) => {
  if (confirm('آیا از حذف این عنوان اطمینان دارید؟')) {
    const titles = extractedTitles.value
    const titleToRemove = titles[index]
    
    // حذف عنوان از خلاصه نوشته
    let newExcerpt = props.excerpt
    const titlePattern = new RegExp(`<a href="#id-${index + 1}"[^>]*>${titleToRemove}</a>`, 'g')
    newExcerpt = newExcerpt.replace(titlePattern, '')
    
    // حذف خطوط خالی اضافی و <br> های اضافی
    newExcerpt = newExcerpt.replace(/<br\s*\/?>\s*<br\s*\/?>\s*<br\s*\/?>/gi, '<br><br>')
    newExcerpt = newExcerpt.replace(/\n\s*\n\s*\n/g, '\n\n')
    
    emit('update:excerpt', newExcerpt)
  }
}

// بستن مودال
const closeModal = () => {
  showModal.value = false
  titleForm.title = ''
  editingIndex.value = -1
  editingTitle.value = ''
}

// مدیریت کلیک روی دکمه AI
const handleAIClick = () => {
  // فعلاً فقط یک پیام نمایش می‌دیم
  // در آینده می‌تونیم AI integration اضافه کنیم
  alert('قابلیت AI به زودی اضافه خواهد شد! 🤖')
}
</script>

<style scoped>
.title-inserter {
  position: relative;
  display: inline-block;
}

/* استایل برای scrollbar */
.overflow-y-auto {
  scrollbar-width: thin;
  scrollbar-color: #cbd5e0 #f7fafc;
}

.overflow-y-auto::-webkit-scrollbar {
  width: 6px;
}

.overflow-y-auto::-webkit-scrollbar-track {
  background: #f7fafc;
}

.overflow-y-auto::-webkit-scrollbar-thumb {
  background-color: #cbd5e0;
  border-radius: 3px;
}

.overflow-y-auto::-webkit-scrollbar-thumb:hover {
  background-color: #a0aec0;
}
</style> 

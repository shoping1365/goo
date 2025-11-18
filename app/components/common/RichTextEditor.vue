<template>
  <Editor
    :model-value="modelValue"
    @update:modelValue="val => emit('update:modelValue', val)"
    :tinymce-script-src="scriptSrc"
    :init="editorConfig"
    :api-key="apiKey"
    class="rich-text-editor"
  />
  <!-- Media Library modal for image/video pick -->
  <MediaLibraryModal
    v-model="showMediaModal"
    :fileType="requestedFileType"
    :defaultCategory="'product-categories'"
    @confirm="onMediaPicked"
  />
</template>

<script setup lang="ts">
import Editor from '@tinymce/tinymce-vue'
import { computed, onMounted, ref } from 'vue'
import MediaLibraryModal from '@/components/media/MediaLibraryModal.vue'

const props = defineProps({
  modelValue: { type: String, default: '' },
  direction: { type: String, default: 'rtl' }, // rtl یا ltr
  lang: { type: String, default: 'fa' }, // fa یا en
  height: { type: Number, default: 800 },
  scriptSrc: { type: String, default: '/tinymce-assets/tinymce/tinymce.min.js' },
  apiKey: { type: String, default: '' },
})
const emit = defineEmits(['update:modelValue', 'h-tag-selected'])

// state for file picker → media library
const showMediaModal = ref(false)
const requestedFileType = ref<'image' | 'video' | 'file' | null>(null)
const pendingPicker = ref<{ cb: (url: string, meta?: any) => void; meta: any } | null>(null)
let activeEditor: any = null
// پیش‌فرض سراسری برای ادیتور: محصولات
const mediaDefaultCategory = ref<'library' | 'products' | 'customer' | 'product-categories' | 'brands' | 'banners'>('product-categories')

function onMediaPicked(files: any[]) {
  if (!files || files.length === 0) {
    showMediaModal.value = false
    requestedFileType.value = null
    pendingPicker.value = null
    return
  }
  const metaType = requestedFileType.value
  // اولویت با نوع درخواستی از ادیتور است
  const picked = (metaType === 'image')
    ? files.find((f: any) => f.type === 'image') || files[0]
    : (metaType === 'video')
      ? files.find((f: any) => f.type === 'video') || files[0]
      : files[0]

  try {
    if (picked) {
      if (pendingPicker.value) {
        // حالت دیالوگ داخلی TinyMCE
        if (metaType === 'image') {
          pendingPicker.value.cb(picked.url, { alt: picked.name })
        } else if (metaType === 'video') {
          pendingPicker.value.cb(picked.url, {})
        } else {
          pendingPicker.value.cb(picked.url)
        }
      } else if (activeEditor) {
        // درج مستقیم از طریق دکمه‌های سفارشی
        if (metaType === 'image') {
          activeEditor.insertContent(`<img src="${picked.url}" alt="${picked.name || ''}" />`)
        } else if (metaType === 'video') {
          activeEditor.insertContent(`<video controls preload="metadata" src="${picked.url}"></video>`) 
        } else {
          activeEditor.insertContent(`<a href="${picked.url}" rel="noopener" target="_blank">${picked.name || picked.url}</a>`) 
        }
      }
    }
  } finally {
    showMediaModal.value = false
    requestedFileType.value = null
    pendingPicker.value = null
  }
}

const editorConfig = computed(() => ({
  height: props.height,
  width: '100%',
  menubar: true,
  branding: false,
  promotion: false,
  directionality: props.direction,
  language: 'fa',
  language_url: '/tinymce-assets/tinymce/langs/fa.js',
  base_url: '/tinymce-assets/tinymce',
  suffix: '',
  license_key: 'gpl',
  plugins: [
    'advlist', 'autolink', 'lists', 'link', 'image', 'charmap', 'preview', 'anchor',
    'searchreplace', 'visualblocks', 'code', 'fullscreen', 'insertdatetime', 'media',
    'table', 'wordcount', 'emoticons', 'codesample', 'nonbreaking', 'directionality', 'visualchars'
  ],
  toolbar:
    'undo redo | blocks fontfamily fontsize | bold italic underline strikethrough forecolor backcolor | ' +
    'alignleft aligncenter alignright alignjustify | bullist numlist outdent indent | ' +
    'ltr rtl | link myimage myvideo table codesample | emoticons charmap | ' +
    'removeformat | fullscreen preview print | searchreplace visualblocks code',
  content_style:
    props.direction === 'rtl'
      ? `
        @font-face { font-family: 'IRANSansWeb'; src: url('/fonts/IRANSansWeb.woff2') format('woff2'); font-display: swap; }
        @font-face { font-family: 'IRANSansWebFaNum'; src: url('/fonts/IRANSansWebFaNum.woff2') format('woff2'); font-display: swap; }
        @font-face { font-family: 'Vazir'; src: url('/fonts/Vazir.woff2') format('woff2'); font-display: swap; }
        @font-face { font-family: 'Shabnam'; src: url('/fonts/Shabnam.woff2') format('woff2'); font-display: swap; }
        @font-face { font-family: 'Sahel'; src: url('/fonts/Sahel.woff2') format('woff2'); font-display: swap; }
        @font-face { font-family: 'Samim'; src: url('/fonts/Samim-Bold.woff') format('woff'); font-display: swap; }
        @font-face { font-family: 'Parastoo'; src: url('/fonts/Parastoo.woff2') format('woff2'); font-display: swap; }
        @font-face { font-family: 'Tanha'; src: url('/fonts/Tanha.woff2') format('woff2'); font-display: swap; }
        @font-face { font-family: 'Yekan'; src: url('/fonts/yekan-regular.woff') format('woff'); font-display: swap; }
        body { font-family: 'IRANSansWeb', 'IRANSansWebFaNum', 'Vazir', 'Shabnam', Tahoma, Arial; direction: rtl; font-size:13px; }
      `
      : "body { font-family: Arial, sans-serif; direction: ltr; font-size:13px; }",
  // تنظیم فونت‌های قابل انتخاب در dropdown
  font_family_formats: props.direction === 'rtl' 
    ? 'IRANSansWeb=IRANSansWeb; IRANSansWebFaNum=IRANSansWebFaNum; Vazir=Vazir; Shabnam=Shabnam; Sahel=Sahel; Samim=Samim; Parastoo=Parastoo; Tanha=Tanha; Yekan=Yekan; Arial=Arial; Tahoma=Tahoma'
    : 'Arial=Arial; Helvetica=Helvetica; Times New Roman=Times New Roman; Georgia=Georgia; Verdana=Verdana',
  // غیرفعال‌سازی Quickbars تا منوی شناور هنگام کلیک باز نشود
  quickbars_insert_toolbar: '',
  quickbars_selection_toolbar: '',
  toolbar_mode: 'sliding',
  // جلوگیری از مزاحمت منوی کانتکست TinyMCE
  contextmenu: false,
  // یکپارچه‌سازی با کتابخانه رسانه پروژه
  file_picker_types: 'image media',
  automatic_uploads: false,
  file_picker_callback: (cb: (url: string, meta?: any) => void, _value: any, meta: any) => {
    pendingPicker.value = { cb, meta }
    requestedFileType.value = meta?.filetype === 'image' ? 'image' : (meta?.filetype === 'media' ? 'video' : 'file')
    showMediaModal.value = true
  },
  autoresize_bottom_margin: 20,
  resize: true,
  setup: (editor: any) => {
    activeEditor = editor
    // دکمه‌های سفارشی برای درج از کتابخانه
    editor.ui.registry.addButton('myimage', {
      icon: 'image',
      tooltip: 'افزودن تصویر از کتابخانه',
      onAction: () => { requestedFileType.value = 'image'; pendingPicker.value = null; showMediaModal.value = true }
    })
    editor.ui.registry.addButton('myvideo', {
      icon: 'embed',
      tooltip: 'افزودن ویدیو از کتابخانه',
      onAction: () => { requestedFileType.value = 'video'; pendingPicker.value = null; showMediaModal.value = true }
    })
    // اضافه کردن event listener برای انتخاب تگ‌های h
    editor.on('NodeChange', (e: any) => {
      const node = e.element
      if (node && (node.nodeName === 'H1' || node.nodeName === 'H2' || node.nodeName === 'H3' || 
                   node.nodeName === 'H4' || node.nodeName === 'H5' || node.nodeName === 'H6')) {
        // ایجاد event برای انتخاب تگ h
        const event = new MouseEvent('click', {
          clientX: e.clientX || 0,
          clientY: e.clientY || 0
        })
        emit('h-tag-selected', event)
      }
    })
  }
}))

onMounted(() => {
  // حذف دکمه Explore trial اگر وجود داشته باشد
  setTimeout(() => {
    const btn = document.querySelector('.tox-promotion-button');
    if (btn) btn.remove();
  }, 500);
});
</script>

<style>
.rich-text-editor :deep(.tox-editor-container) {
  resize: both;
  overflow: auto ;
  min-width: 200px;
}
.rich-text-editor {
}
.rich-text-editor :deep(.tox-promotion),
.rich-text-editor :deep(.tox-promotion-button) {
  display: none;
}
</style>

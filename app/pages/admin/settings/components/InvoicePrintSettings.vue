<template>
  <div class="space-y-8">
    <!-- هدر بخش -->
    <div class="relative overflow-hidden bg-gradient-to-br from-purple-50 via-pink-50 to-red-50 rounded-2xl p-8">
      <div class="absolute top-0 right-0 w-32 h-32 bg-gradient-to-bl from-purple-200/30 to-pink-200/30 rounded-full -translate-y-16 translate-x-16"></div>
      <div class="absolute bottom-0 left-0 w-24 h-24 bg-gradient-to-tr from-red-200/30 to-orange-200/30 rounded-full translate-y-12 -translate-x-12"></div>
      
      <div class="relative z-10">
        <div class="flex items-center mb-4">
          <div class="w-12 h-12 bg-gradient-to-br from-purple-500 to-pink-600 rounded-xl flex items-center justify-center mr-4">
            <i class="i-heroicons-printer text-white text-xl"></i>
          </div>
          <div>
            <h2 class="text-3xl font-bold bg-gradient-to-r from-gray-800 to-gray-600 bg-clip-text text-transparent">پیکربندی فاکتور و چاپ</h2>
            <p class="text-gray-600 mt-1">تنظیمات مربوط به فاکتورها، قالب‌های چاپ و خروجی‌های سیستم</p>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات فاکتور -->
    <div class="bg-white rounded-2xl shadow-lg border border-gray-100 overflow-hidden">
      <div class="bg-gradient-to-r from-purple-500 to-pink-600 px-6 py-4">
        <div class="flex items-center">
          <i class="i-heroicons-document-text text-white text-xl mr-3"></i>
          <h3 class="text-xl font-bold text-white">تنظیمات فاکتور</h3>
        </div>
      </div>
      
      <div class="p-6 space-y-6">
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <div class="group">
            <label class="block text-sm font-semibold text-gray-700 mb-3 flex items-center">
              <i class="i-heroicons-building-office text-purple-500 mr-2"></i>
              نام شرکت در فاکتور
            </label>
            <input 
              v-model="invoiceSettings.companyName" 
              type="text" 
              class="w-full px-4 py-3 border-2 border-gray-200 rounded-xl focus:outline-none focus:border-purple-500 focus:ring-4 focus:ring-purple-100 transition-all duration-200 group-hover:border-gray-300"
              placeholder="نام شرکت شما"
            >
          </div>
          
          <div class="group">
            <label class="block text-sm font-semibold text-gray-700 mb-3 flex items-center">
              <i class="i-heroicons-identification text-purple-500 mr-2"></i>
              شماره اقتصادی
            </label>
            <input 
              v-model="invoiceSettings.taxNumber" 
              type="text" 
              class="w-full px-4 py-3 border-2 border-gray-200 rounded-xl focus:outline-none focus:border-purple-500 focus:ring-4 focus:ring-purple-100 transition-all duration-200 group-hover:border-gray-300"
              placeholder="شماره اقتصادی شرکت"
            >
          </div>
          
          <div class="group">
            <label class="block text-sm font-semibold text-gray-700 mb-3 flex items-center">
              <i class="i-heroicons-map-pin text-purple-500 mr-2"></i>
              آدرس شرکت
            </label>
            <textarea 
              v-model="invoiceSettings.companyAddress" 
              rows="3"
              class="w-full px-4 py-3 border-2 border-gray-200 rounded-xl focus:outline-none focus:border-purple-500 focus:ring-4 focus:ring-purple-100 transition-all duration-200 group-hover:border-gray-300"
              placeholder="آدرس کامل شرکت"
            ></textarea>
          </div>
          
          <div class="group">
            <label class="block text-sm font-semibold text-gray-700 mb-3 flex items-center">
              <i class="i-heroicons-phone text-purple-500 mr-2"></i>
              شماره تلفن شرکت
            </label>
            <input 
              v-model="invoiceSettings.companyPhone" 
              type="text" 
              class="w-full px-4 py-3 border-2 border-gray-200 rounded-xl focus:outline-none focus:border-purple-500 focus:ring-4 focus:ring-purple-100 transition-all duration-200 group-hover:border-gray-300"
              placeholder="شماره تلفن شرکت"
            >
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات چاپ -->
    <div class="bg-white rounded-2xl shadow-lg border border-gray-100 overflow-hidden">
      <div class="bg-gradient-to-r from-blue-500 to-indigo-600 px-6 py-4">
        <div class="flex items-center">
          <i class="i-heroicons-printer text-white text-xl mr-3"></i>
          <h3 class="text-xl font-bold text-white">تنظیمات چاپ</h3>
        </div>
      </div>
      
      <div class="p-6 space-y-6">
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <div class="group">
            <label class="block text-sm font-semibold text-gray-700 mb-3 flex items-center">
              <i class="i-heroicons-document text-blue-500 mr-2"></i>
              اندازه کاغذ پیش‌فرض
            </label>
            <select v-model="printSettings.defaultPaperSize" class="w-full px-4 py-3 border-2 border-gray-200 rounded-xl focus:outline-none focus:border-blue-500 focus:ring-4 focus:ring-blue-100 transition-all duration-200 group-hover:border-gray-300 appearance-none bg-white">
              <option value="a4">A4 (210 × 297 mm)</option>
              <option value="a5">A5 (148 × 210 mm)</option>
              <option value="letter">Letter (216 × 279 mm)</option>
              <option value="legal">Legal (216 × 356 mm)</option>
            </select>
          </div>
          
          <div class="group">
            <label class="block text-sm font-semibold text-gray-700 mb-3 flex items-center">
              <i class="i-heroicons-view-columns text-blue-500 mr-2"></i>
              جهت چاپ
            </label>
            <select v-model="printSettings.orientation" class="w-full px-4 py-3 border-2 border-gray-200 rounded-xl focus:outline-none focus:border-blue-500 focus:ring-4 focus:ring-blue-100 transition-all duration-200 group-hover:border-gray-300 appearance-none bg-white">
              <option value="portrait">عمودی (Portrait)</option>
              <option value="landscape">افقی (Landscape)</option>
            </select>
          </div>
          
          <div class="group">
            <label class="block text-sm font-semibold text-gray-700 mb-3 flex items-center">
              <i class="i-heroicons-font-family text-blue-500 mr-2"></i>
              فونت پیش‌فرض
            </label>
            <select v-model="printSettings.defaultFont" class="w-full px-4 py-3 border-2 border-gray-200 rounded-xl focus:outline-none focus:border-blue-500 focus:ring-4 focus:ring-blue-100 transition-all duration-200 group-hover:border-gray-300 appearance-none bg-white">
              <option value="iransans">ایران سنس</option>
              <option value="tahoma">تahoma</option>
              <option value="arial">Arial</option>
              <option value="times">Times New Roman</option>
            </select>
          </div>
          
          <div class="group">
            <label class="block text-sm font-semibold text-gray-700 mb-3 flex items-center">
              <i class="i-heroicons-arrows-pointing-out text-blue-500 mr-2"></i>
              اندازه فونت
            </label>
            <select v-model="printSettings.fontSize" class="w-full px-4 py-3 border-2 border-gray-200 rounded-xl focus:outline-none focus:border-blue-500 focus:ring-4 focus:ring-blue-100 transition-all duration-200 group-hover:border-gray-300 appearance-none bg-white">
              <option value="10">10pt</option>
              <option value="11">11pt</option>
              <option value="12">12pt</option>
              <option value="14">14pt</option>
              <option value="16">16pt</option>
            </select>
          </div>
        </div>
        
        <div class="space-y-4">
          <div class="flex items-center">
            <input 
              v-model="printSettings.showLogo" 
              type="checkbox" 
              class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
            >
            <label class="mr-3 text-sm font-medium text-gray-700">نمایش لوگو در فاکتور</label>
          </div>
          
          <div class="flex items-center">
            <input 
              v-model="printSettings.showQRCode" 
              type="checkbox" 
              class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
            >
            <label class="mr-3 text-sm font-medium text-gray-700">نمایش کد QR در فاکتور</label>
          </div>
          
          <div class="flex items-center">
            <input 
              v-model="printSettings.showBarcode" 
              type="checkbox" 
              class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
            >
            <label class="mr-3 text-sm font-medium text-gray-700">نمایش بارکد در فاکتور</label>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات قالب -->
    <div class="bg-white rounded-2xl shadow-lg border border-gray-100 overflow-hidden">
      <div class="bg-gradient-to-r from-green-500 to-emerald-600 px-6 py-4">
        <div class="flex items-center">
          <i class="i-heroicons-paint-brush text-white text-xl mr-3"></i>
          <h3 class="text-xl font-bold text-white">تنظیمات قالب</h3>
        </div>
      </div>
      
      <div class="p-6 space-y-6">
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <div class="group">
            <label class="block text-sm font-semibold text-gray-700 mb-3 flex items-center">
              <i class="i-heroicons-swatch text-green-500 mr-2"></i>
              رنگ اصلی قالب
            </label>
            <input 
              v-model="templateSettings.primaryColor" 
              type="color" 
              class="w-full h-12 border-2 border-gray-200 rounded-xl focus:outline-none focus:border-green-500 focus:ring-4 focus:ring-green-100 transition-all duration-200 group-hover:border-gray-300"
            >
          </div>
          
          <div class="group">
            <label class="block text-sm font-semibold text-gray-700 mb-3 flex items-center">
              <i class="i-heroicons-swatch text-green-500 mr-2"></i>
              رنگ ثانویه قالب
            </label>
            <input 
              v-model="templateSettings.secondaryColor" 
              type="color" 
              class="w-full h-12 border-2 border-gray-200 rounded-xl focus:outline-none focus:border-green-500 focus:ring-4 focus:ring-green-100 transition-all duration-200 group-hover:border-gray-300"
            >
          </div>
        </div>
        
        <div class="space-y-4">
          <div class="flex items-center">
            <input 
              v-model="templateSettings.showHeader" 
              type="checkbox" 
              class="w-4 h-4 text-green-600 border-gray-300 rounded focus:ring-green-500"
            >
            <label class="mr-3 text-sm font-medium text-gray-700">نمایش هدر در فاکتور</label>
          </div>
          
          <div class="flex items-center">
            <input 
              v-model="templateSettings.showFooter" 
              type="checkbox" 
              class="w-4 h-4 text-green-600 border-gray-300 rounded focus:ring-green-500"
            >
            <label class="mr-3 text-sm font-medium text-gray-700">نمایش فوتر در فاکتور</label>
          </div>
          
          <div class="flex items-center">
            <input 
              v-model="templateSettings.showPageNumbers" 
              type="checkbox" 
              class="w-4 h-4 text-green-600 border-gray-300 rounded focus:ring-green-500"
            >
            <label class="mr-3 text-sm font-medium text-gray-700">نمایش شماره صفحه</label>
          </div>
        </div>
      </div>
    </div>

    <!-- دکمه‌های عملیات -->
    <div class="bg-gradient-to-r from-gray-50 to-gray-100 rounded-2xl p-6">
      <div class="flex flex-col sm:flex-row justify-end space-y-3 sm:space-y-0 sm:space-x-3 sm:space-x-reverse">
        <button 
          type="button" 
          @click="resetInvoicePrintSettings"
          class="px-8 py-3 border-2 border-gray-300 text-gray-700 rounded-xl hover:bg-gray-50 hover:border-gray-400 transition-all duration-200 font-semibold flex items-center justify-center"
        >
          <i class="i-heroicons-arrow-path mr-2"></i>
          بازنشانی تنظیمات
        </button>
        <button 
          type="button" 
          @click="saveInvoicePrintSettings"
          :disabled="savingInvoicePrint"
          class="px-8 py-3 bg-gradient-to-r from-purple-500 to-pink-600 text-white rounded-xl hover:from-purple-600 hover:to-pink-700 disabled:opacity-50 transition-all duration-200 shadow-lg hover:shadow-xl font-semibold flex items-center justify-center"
        >
          <i v-if="savingInvoicePrint" class="i-heroicons-arrow-path animate-spin mr-2"></i>
          <i v-else class="i-heroicons-check mr-2"></i>
          {{ savingInvoicePrint ? 'در حال ذخیره...' : 'ذخیره تنظیمات فاکتور و چاپ' }}
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  invoiceSettings: {
    type: Object,
    required: true
  },
  printSettings: {
    type: Object,
    required: true
  },
  templateSettings: {
    type: Object,
    required: true
  },
  savingInvoicePrint: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['save', 'reset'])

const saveInvoicePrintSettings = () => {
  emit('save')
}

const resetInvoicePrintSettings = () => {
  emit('reset')
}
</script> 

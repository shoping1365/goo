<template>
  <div class="space-y-8">
    <h3 class="text-lg font-semibold text-gray-900">مدیریت قالب‌های کارت</h3>
    
    <!-- پیش‌نمایش زنده قالب فعال -->
    <div class="bg-white border rounded-lg px-4 py-4">
      <h4 class="text-md font-bold text-gray-800 mb-4">پیش‌نمایش قالب فعال</h4>
      <div class="flex flex-col md:flex-row gapx-4 py-4">
        <div class="flex-1">
          <div class="relative w-full max-w-md mx-auto">
            <img 
              :src="activeTemplate.image" 
              alt="قالب فعال" 
              class="w-full h-48 object-cover rounded-lg border shadow-lg"
            />
            <div 
              class="absolute inset-0 flex items-center justify-center"
              style="background: rgba(0,0,0,0.1);"
            >
              <div class="text-center text-white">
                <div class="text-2xl font-bold mb-2">گیفت کارت</div>
                <div class="text-lg">مبلغ: ۵۰۰,۰۰۰ تومان</div>
                <div class="text-sm mt-2">کد: GC-2024-001</div>
              </div>
            </div>
          </div>
        </div>
        <div class="flex-1">
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نام قالب فعال</label>
              <div class="text-lg font-semibold text-gray-900">{{ activeTemplate.name }}</div>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نوع قالب</label>
              <span
class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                :class="isDefaultTemplate ? 'bg-blue-100 text-blue-800' : 'bg-green-100 text-green-800'"
              >
                {{ isDefaultTemplate ? 'پیش‌فرض' : 'سفارشی' }}
              </span>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
              <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800">
                فعال
              </span>
            </div>
            <div class="pt-4">
              <button
                class="w-full px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
                @click="customizeActiveTemplate"
              >
                سفارشی‌سازی قالب فعال
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- لیست قالب‌ها -->
    <div>
      <h4 class="text-md font-bold text-gray-800 mb-4">قالب‌های پیش‌فرض</h4>
      <div class="grid grid-cols-2 md:grid-cols-4 gapx-4 py-4 mb-6">
        <div v-for="tpl in defaultTemplates" :key="tpl.id" class="bg-white border rounded-lg p-3 flex flex-col items-center">
          <img :src="tpl.image" alt="قالب پیش‌فرض" class="w-32 h-20 object-cover rounded mb-2 border" />
          <div class="text-sm text-gray-700 mb-1">{{ tpl.name }}</div>
          <button
            :class="tpl.id === selectedTemplateId ? 'bg-blue-600 text-white' : 'bg-gray-200 text-gray-700'"
            class="px-3 py-1 rounded text-xs font-medium mt-1"
            @click="selectTemplate(tpl)"
          >
            {{ tpl.id === selectedTemplateId ? 'قالب فعال' : 'انتخاب قالب' }}
          </button>
        </div>
      </div>
      <h4 class="text-md font-bold text-gray-800 mb-4">قالب‌های سفارشی</h4>
      <div class="grid grid-cols-2 md:grid-cols-4 gapx-4 py-4">
        <div v-for="tpl in customTemplates" :key="tpl.id" class="bg-white border rounded-lg p-3 flex flex-col items-center relative">
          <img :src="tpl.image" alt="قالب سفارشی" class="w-32 h-20 object-cover rounded mb-2 border" />
          <div class="text-sm text-gray-700 mb-1">{{ tpl.name }}</div>
          <div class="flex gap-2">
            <button class="text-blue-600 text-xs" @click="editTemplate(tpl)">ویرایش</button>
            <button class="text-red-600 text-xs" @click="deleteTemplate(tpl.id)">حذف</button>
          </div>
          <button
            :class="tpl.id === selectedTemplateId ? 'bg-blue-600 text-white' : 'bg-gray-200 text-gray-700'"
            class="px-3 py-1 rounded text-xs font-medium mt-2"
            @click="selectTemplate(tpl)"
          >
            {{ tpl.id === selectedTemplateId ? 'قالب فعال' : 'انتخاب قالب' }}
          </button>
        </div>
        <!-- افزودن قالب سفارشی جدید -->
        <div class="flex flex-col items-center justify-center border-2 border-dashed border-gray-300 rounded-lg px-4 py-4 cursor-pointer hover:bg-gray-50" @click="showAddModal = true">
          <svg class="w-10 h-10 text-gray-400 mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
          </svg>
          <span class="text-xs text-gray-500">افزودن قالب جدید</span>
        </div>
      </div>
    </div>
    <!-- مودال افزودن قالب سفارشی -->
    <div v-if="showAddModal" class="fixed inset-0 bg-gray-600 bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg px-4 py-4 w-full max-w-md relative">
        <button class="absolute top-2 left-2 text-gray-400 hover:text-gray-600" @click="showAddModal = false">
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
        <h4 class="text-md font-bold text-gray-800 mb-4">افزودن قالب سفارشی جدید</h4>
        <form class="space-y-4" @submit.prevent="handleAddTemplate">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نام قالب</label>
            <input v-model="newTemplate.name" type="text" required class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">آپلود تصویر قالب</label>
            <input type="file" accept="image/*" class="w-full" @change="onImageChange" />
            <img v-if="newTemplate.image" :src="newTemplate.image" alt="پیش‌نمایش" class="w-32 h-20 object-cover rounded mt-2 border" />
          </div>
          <div class="flex justify-end">
            <button type="submit" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">افزودن</button>
          </div>
        </form>
      </div>
    </div>
    <!-- ویرایشگر طراحی ساده -->
    <div v-if="showEditor" class="fixed inset-0 bg-gray-600 bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg px-4 py-4 w-full max-w-6xl relative max-h-[90vh] overflow-y-auto">
        <button class="absolute top-2 left-2 text-gray-400 hover:text-gray-600" @click="showEditor = false">
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
        <h4 class="text-md font-bold text-gray-800 mb-4">ویرایشگر طراحی قالب</h4>
        
        <!-- نوار ابزار -->
        <div class="flex gap-2 mb-4 pb-2 border-b">
          <button class="px-3 py-1 bg-blue-600 text-white text-xs rounded hover:bg-blue-700" @click="addNewLayer">
            افزودن لایه
          </button>
          <button class="px-3 py-1 bg-green-600 text-white text-xs rounded hover:bg-green-700" @click="duplicateLayer">
            تکثیر لایه
          </button>
          <button class="px-3 py-1 bg-red-600 text-white text-xs rounded hover:bg-red-700" @click="deleteLayer">
            حذف لایه
          </button>
          <button :disabled="!canUndo" class="px-3 py-1 bg-gray-600 text-white text-xs rounded hover:bg-gray-700 disabled:opacity-50" @click="undoAction">
            بازگشت
          </button>
          <button :disabled="!canRedo" class="px-3 py-1 bg-gray-600 text-white text-xs rounded hover:bg-gray-700 disabled:opacity-50" @click="redoAction">
            تکرار
          </button>
          <button class="px-3 py-1 bg-purple-600 text-white text-xs rounded hover:bg-purple-700" @click="exportTemplate">
            خروجی قالب
          </button>
          <button class="px-3 py-1 bg-orange-600 text-white text-xs rounded hover:bg-orange-700" @click="showTemplates">
            قالب‌های آماده
          </button>
        </div>
        
        <div class="flex flex-col lg:flex-row gapx-4 py-4">
          <!-- پنل تنظیمات -->
          <div class="lg:w-1/3 space-y-4">
            <!-- انتخاب لایه -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">لایه فعلی</label>
              <select v-model="currentLayerIndex" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent">
                <option v-for="(layer, index) in layers" :key="index" :value="index">
                  {{ layer.name }} ({{ layer.type }})
                </option>
              </select>
            </div>
            
            <!-- تنظیمات لایه فعلی -->
            <div v-if="currentLayer">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">نام لایه</label>
                <input v-model="currentLayer.name" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent" />
              </div>
              
              <div v-if="currentLayer.type === 'text'">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">متن</label>
                  <textarea v-model="currentLayer.text" rows="3" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"></textarea>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">فونت</label>
                  <select v-model="currentLayer.fontFamily" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent">
                    <option value="IRANSans">ایران سنس</option>
                    <option value="Vazir">وزیر</option>
                    <option value="Samim">صمیم</option>
                    <option value="Shabnam">شبنم</option>
                    <option value="Arial">Arial</option>
                  </select>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">رنگ</label>
                  <input v-model="currentLayer.color" type="color" class="w-full h-10 border border-gray-300 rounded-lg" />
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">اندازه</label>
                  <input v-model="currentLayer.fontSize" type="range" min="12" max="48" class="w-full" />
                  <span class="text-sm text-gray-600">{{ currentLayer.fontSize }}px</span>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">ضخامت</label>
                  <select v-model="currentLayer.fontWeight" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent">
                    <option value="normal">عادی</option>
                    <option value="bold">ضخیم</option>
                    <option value="bolder">خیلی ضخیم</option>
                    <option value="lighter">نازک</option>
                  </select>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">چرخش</label>
                  <input v-model="currentLayer.rotation" type="range" min="-180" max="180" class="w-full" />
                  <span class="text-sm text-gray-600">{{ currentLayer.rotation }}°</span>
                </div>
                
                <!-- افکت‌های متن -->
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">افکت‌های متن</label>
                  <div class="space-y-2">
                    <div class="flex items-center">
                      <input id="shadow" v-model="currentLayer.effects.shadow" type="checkbox" class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded" />
                      <label for="shadow" class="mr-2 block text-sm text-gray-900">سایه متن</label>
                    </div>
                    <div class="flex items-center">
                      <input id="outline" v-model="currentLayer.effects.outline" type="checkbox" class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded" />
                      <label for="outline" class="mr-2 block text-sm text-gray-900">حاشیه متن</label>
                    </div>
                    <div class="flex items-center">
                      <input id="gradient" v-model="currentLayer.effects.gradient" type="checkbox" class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded" />
                      <label for="gradient" class="mr-2 block text-sm text-gray-900">گرادیانت</label>
                    </div>
                  </div>
                </div>
                
                <div v-if="currentLayer.effects.shadow">
                  <label class="block text-sm font-medium text-gray-700 mb-2">رنگ سایه</label>
                  <input v-model="currentLayer.effects.shadowColor" type="color" class="w-full h-8 border border-gray-300 rounded-lg" />
                </div>
                <div v-if="currentLayer.effects.outline">
                  <label class="block text-sm font-medium text-gray-700 mb-2">رنگ حاشیه</label>
                  <input v-model="currentLayer.effects.outlineColor" type="color" class="w-full h-8 border border-gray-300 rounded-lg" />
                </div>
              </div>
              
              <div v-if="currentLayer.type === 'image'">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">آپلود تصویر</label>
                  <input type="file" accept="image/*" class="w-full" @change="onLayerImageChange" />
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">شفافیت</label>
                  <input v-model="currentLayer.opacity" type="range" min="0" max="100" class="w-full" />
                  <span class="text-sm text-gray-600">{{ currentLayer.opacity }}%</span>
                </div>
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">موقعیت</label>
                <div class="grid grid-cols-2 gap-2 text-sm">
                  <div>
                    <label class="block text-xs text-gray-600">X</label>
                    <input v-model="currentLayer.position.x" type="number" class="w-full px-2 py-1 border border-gray-300 rounded text-xs" />
                  </div>
                  <div>
                    <label class="block text-xs text-gray-600">Y</label>
                    <input v-model="currentLayer.position.y" type="number" class="w-full px-2 py-1 border border-gray-300 rounded text-xs" />
                  </div>
                </div>
              </div>
              
              <!-- تنظیمات پیشرفته -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">تنظیمات پیشرفته</label>
                <div class="space-y-2">
                  <div class="flex items-center">
                    <input id="locked" v-model="currentLayer.locked" type="checkbox" class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded" />
                    <label for="locked" class="mr-2 block text-sm text-gray-900">قفل لایه</label>
                  </div>
                  <div class="flex items-center">
                    <input id="visible" v-model="currentLayer.visible" type="checkbox" class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded" />
                    <label for="visible" class="mr-2 block text-sm text-gray-900">نمایش لایه</label>
                  </div>
                </div>
              </div>
            </div>
            
            <div class="pt-4">
              <button class="w-full px-4 py-2 bg-green-600 text-white text-sm font-medium rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2" @click="saveEditor">
                ذخیره تغییرات
              </button>
            </div>
          </div>
          
          <!-- ناحیه طراحی -->
          <div class="lg:w-2/3">
            <div class="flex justify-center">
              <div class="relative w-80 h-48 border-2 border-gray-300 rounded-lg bg-gray-50 overflow-hidden">
                <img :src="editorTemplate.image" class="absolute inset-0 w-full h-full object-cover" />
                
                <!-- لایه‌ها -->
                <div 
                  v-for="(layer, index) in layers" 
                  :key="index"
                  class="absolute cursor-move select-none"
                  :class="{ 'cursor-not-allowed': layer.locked, 'opacity-50': !layer.visible }"
                  :style="{
                    top: layer.position.y + 'px',
                    left: layer.position.x + 'px',
                    transform: `rotate(${layer.rotation || 0}deg)`,
                    opacity: layer.opacity ? layer.opacity / 100 : 1,
                    zIndex: index
                  }"
                  @mousedown="layer.locked ? null : startLayerDrag(index, $event)"
                >
                  <!-- لایه متن -->
                  <div 
                    v-if="layer.type === 'text'"
                    :style="{
                      color: layer.color,
                      fontSize: layer.fontSize + 'px',
                      fontFamily: layer.fontFamily,
                      fontWeight: layer.fontWeight || 'normal',
                      textShadow: layer.effects?.shadow ? `${layer.effects.shadowColor} 2px 2px 4px` : 'none',
                      textDecoration: layer.effects?.outline ? 'underline' : 'none',
                      textDecorationColor: layer.effects?.outlineColor,
                      background: layer.effects?.gradient ? `linear-gradient(to right, ${layer.color}, ${layer.effects.outlineColor})` : 'none',
                      WebkitBackgroundClip: layer.effects?.gradient ? 'text' : 'initial',
                      WebkitTextFillColor: layer.effects?.gradient ? 'transparent' : 'initial'
                    }"
                  >
                    {{ layer.text || 'متن نمونه' }}
                  </div>
                  
                  <!-- لایه تصویر -->
                  <img 
                    v-if="layer.type === 'image' && layer.imageUrl"
                    :src="layer.imageUrl" 
                    class="max-w-20 max-h-20 object-contain"
                  />
                  
                  <!-- نشانگر قفل -->
                  <div v-if="layer.locked" class="absolute -top-2 -right-2 w-4 h-4 bg-red-500 rounded-full flex items-center justify-center">
                    <svg class="w-2 h-2 text-white" fill="currentColor" viewBox="0 0 20 20">
                      <path fill-rule="evenodd" d="M5 9V7a5 5 0 0110 0v2a2 2 0 012 2v5a2 2 0 01-2 2H5a2 2 0 01-2-2v-5a2 2 0 012-2zm8-2v2H7V7a3 3 0 016 0z" clip-rule="evenodd" />
                    </svg>
                  </div>
                </div>
                
                <!-- راهنما -->
                <div class="absolute top-2 left-2 text-xs text-white bg-black bg-opacity-50 px-2 py-1 rounded">
                  لایه‌ها را بکشید تا جابجا کنید
                </div>
              </div>
            </div>
            
            <!-- پیش‌نمایش نهایی -->
            <div class="mt-4 text-center">
              <h5 class="text-sm font-medium text-gray-700 mb-2">پیش‌نمایش نهایی</h5>
              <div class="inline-block relative w-64 h-40 border rounded-lg overflow-hidden shadow-lg">
                <img :src="editorTemplate.image" class="w-full h-full object-cover" />
                <div 
                  v-for="(layer, index) in layers.filter(l => l.visible)" 
                  :key="index"
                  class="absolute"
                  :style="{
                    top: layer.position.y + 'px',
                    left: layer.position.x + 'px',
                    transform: `rotate(${layer.rotation || 0}deg)`,
                    opacity: layer.opacity ? layer.opacity / 100 : 1,
                    zIndex: index
                  }"
                >
                  <div 
                    v-if="layer.type === 'text'"
                    :style="{
                      color: layer.color,
                      fontSize: layer.fontSize + 'px',
                      fontFamily: layer.fontFamily,
                      fontWeight: layer.fontWeight || 'normal',
                      textShadow: layer.effects?.shadow ? `${layer.effects.shadowColor} 2px 2px 4px` : 'none',
                      textDecoration: layer.effects?.outline ? 'underline' : 'none',
                      textDecorationColor: layer.effects?.outlineColor,
                      background: layer.effects?.gradient ? `linear-gradient(to right, ${layer.color}, ${layer.effects.outlineColor})` : 'none',
                      WebkitBackgroundClip: layer.effects?.gradient ? 'text' : 'initial',
                      WebkitTextFillColor: layer.effects?.gradient ? 'transparent' : 'initial'
                    }"
                  >
                    {{ layer.text || 'متن نمونه' }}
                  </div>
                  <img 
                    v-if="layer.type === 'image' && layer.imageUrl"
                    :src="layer.imageUrl" 
                    class="max-w-16 max-h-16 object-contain"
                  />
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    
    <!-- مودال قالب‌های آماده -->
    <div v-if="showTemplatesModal" class="fixed inset-0 bg-gray-600 bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg px-4 py-4 w-full max-w-4xl max-h-[80vh] overflow-y-auto">
        <div class="flex justify-between items-center mb-4">
          <h4 class="text-lg font-bold text-gray-800">قالب‌های آماده</h4>
          <button class="text-gray-400 hover:text-gray-600" @click="showTemplatesModal = false">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
        <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gapx-4 py-4">
          <div 
            v-for="template in readyTemplates" 
            :key="template.id"
            class="border rounded-lg p-3 cursor-pointer hover:bg-gray-50"
            @click="applyReadyTemplate(template)"
          >
            <div class="relative w-full h-24 bg-gray-100 rounded mb-2">
              <div class="absolute inset-0 flex items-center justify-center text-xs text-gray-500">
                پیش‌نمایش
              </div>
            </div>
            <div class="text-sm font-medium text-gray-900">{{ template.name }}</div>
            <div class="text-xs text-gray-600">{{ template.description }}</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
const showAddModal = ref(false)
const showEditor = ref(false)
const showTemplatesModal = ref(false)
const selectedTemplateId = ref(1)
const defaultTemplates = ref([
  { id: 1, name: 'کلاسیک', image: '/statics/images/giftcard-default-1.png' },
  { id: 2, name: 'مدرن', image: '/statics/images/giftcard-default-2.png' }
])
const customTemplates = ref([
  // نمونه اولیه
])
const newTemplate = ref({ name: '', image: '' })

// متغیرهای ویرایشگر
const editorTemplate = ref({
  name: '',
  image: '/default-product.svg'
})

const layers = ref([
  {
    id: 1,
    name: 'متن اصلی',
    type: 'text',
    text: 'گیفت کارت',
    color: '#000000',
    fontSize: 24,
    fontFamily: 'IRANSans',
    fontWeight: 'bold',
    rotation: 0,
    position: { x: 50, y: 50 },
    opacity: 100,
    effects: {
      shadow: false,
      shadowColor: '#000000',
      outline: false,
      outlineColor: '#000000',
      gradient: false
    },
    locked: false,
    visible: true
  }
])

const currentLayerIndex = ref(0)
const currentLayer = computed(() => layers.value[currentLayerIndex.value])

// متغیرهای undo/redo
const undoStack = ref([])
const redoStack = ref([])
const canUndo = computed(() => undoStack.value.length > 0)
const canRedo = computed(() => redoStack.value.length > 0)

// قالب‌های آماده
const readyTemplates = ref([
  {
    id: 1,
    name: 'قالب تولد',
    description: 'برای مناسبت‌های تولد',
    layers: [
      {
        id: 1,
        name: 'عنوان',
        type: 'text',
        text: 'تولدت مبارک!',
        color: '#FF6B6B',
        fontSize: 28,
        fontFamily: 'IRANSans',
        fontWeight: 'bold',
        rotation: 0,
        position: { x: 40, y: 30 },
        opacity: 100,
        effects: {
          shadow: false,
          shadowColor: '#000000',
          outline: false,
          outlineColor: '#000000',
          gradient: false
        },
        locked: false,
        visible: true
      },
      {
        id: 2,
        name: 'متن فرعی',
        type: 'text',
        text: 'با آرزوی شادی و موفقیت',
        color: '#4ECDC4',
        fontSize: 16,
        fontFamily: 'Vazir',
        fontWeight: 'normal',
        rotation: 0,
        position: { x: 45, y: 70 },
        opacity: 100,
        effects: {
          shadow: false,
          shadowColor: '#000000',
          outline: false,
          outlineColor: '#000000',
          gradient: false
        },
        locked: false,
        visible: true
      }
    ]
  },
  {
    id: 2,
    name: 'قالب عروسی',
    description: 'برای مراسم عروسی',
    layers: [
      {
        id: 1,
        name: 'عنوان',
        type: 'text',
        text: 'عروسی مبارک',
        color: '#FFD93D',
        fontSize: 32,
        fontFamily: 'Samim',
        fontWeight: 'bold',
        rotation: 0,
        position: { x: 35, y: 25 },
        opacity: 100,
        effects: {
          shadow: false,
          shadowColor: '#000000',
          outline: false,
          outlineColor: '#000000',
          gradient: false
        },
        locked: false,
        visible: true
      },
      {
        id: 2,
        name: 'متن فرعی',
        type: 'text',
        text: 'با آرزوی زندگی سرشار از عشق',
        color: '#6C5CE7',
        fontSize: 14,
        fontFamily: 'Shabnam',
        fontWeight: 'normal',
        rotation: 0,
        position: { x: 40, y: 75 },
        opacity: 100,
        effects: {
          shadow: false,
          shadowColor: '#000000',
          outline: false,
          outlineColor: '#000000',
          gradient: false
        },
        locked: false,
        visible: true
      }
    ]
  },
  {
    id: 3,
    name: 'قالب سال نو',
    description: 'برای سال نو',
    layers: [
      {
        id: 1,
        name: 'عنوان',
        type: 'text',
        text: 'سال نو مبارک',
        color: '#00B894',
        fontSize: 30,
        fontFamily: 'IRANSans',
        fontWeight: 'bold',
        rotation: 0,
        position: { x: 30, y: 30 },
        opacity: 100,
        effects: {
          shadow: false,
          shadowColor: '#000000',
          outline: false,
          outlineColor: '#000000',
          gradient: false
        },
        locked: false,
        visible: true
      },
      {
        id: 2,
        name: 'متن فرعی',
        type: 'text',
        text: 'با آرزوی سالی پر از موفقیت',
        color: '#FDCB6E',
        fontSize: 16,
        fontFamily: 'Vazir',
        fontWeight: 'normal',
        rotation: 0,
        position: { x: 35, y: 70 },
        opacity: 100,
        effects: {
          shadow: false,
          shadowColor: '#000000',
          outline: false,
          outlineColor: '#000000',
          gradient: false
        },
        locked: false,
        visible: true
      }
    ]
  }
])

// متغیرهای drag & drop برای لایه‌ها
const isLayerDragging = ref(false)
const layerDragStartPos = ref({ x: 0, y: 0 })
const dragLayerIndex = ref(-1)

// محاسبه قالب فعال
const activeTemplate = computed(() => {
  const defaultTpl = defaultTemplates.value.find(t => t.id === selectedTemplateId.value)
  const customTpl = customTemplates.value.find(t => t.id === selectedTemplateId.value)
  return defaultTpl || customTpl || defaultTemplates.value[0]
})

// بررسی نوع قالب
const isDefaultTemplate = computed(() => {
  return defaultTemplates.value.some(t => t.id === selectedTemplateId.value)
})

function selectTemplate(tpl) {
  selectedTemplateId.value = tpl.id
}

function customizeActiveTemplate() {
  if (isDefaultTemplate.value) {
    // اگر قالب پیش‌فرض است، یک کپی سفارشی ایجاد کن
    const newCustomTemplate = {
      id: Date.now(),
      name: `${activeTemplate.value.name} - کپی`,
      image: activeTemplate.value.image
    }
    customTemplates.value.push(newCustomTemplate)
    selectedTemplateId.value = newCustomTemplate.id
  }
  editTemplate(activeTemplate.value)
}

function handleAddTemplate() {
  if (!newTemplate.value.name || !newTemplate.value.image) {
    alert('نام و تصویر قالب الزامی است.')
    return
  }
  const id = Date.now()
  customTemplates.value.push({ id, name: newTemplate.value.name, image: newTemplate.value.image })
  newTemplate.value = { name: '', image: '' }
  showAddModal.value = false
}

function onImageChange(e) {
  const file = e.target.files[0]
  if (!file) return
  const reader = new FileReader()
  reader.onload = (ev) => {
    newTemplate.value.image = ev.target.result
  }
  reader.readAsDataURL(file)
}

function editTemplate(tpl) {
  editorTemplate.value = { ...tpl }
  // تنظیم لایه پیش‌فرض
  layers.value = [
    {
      id: 1,
      name: 'متن اصلی',
      type: 'text',
      text: 'گیفت کارت',
      color: '#000000',
      fontSize: 24,
      fontFamily: 'IRANSans',
      fontWeight: 'bold',
      rotation: 0,
      position: { x: 50, y: 50 },
      opacity: 100,
      effects: {
        shadow: false,
        shadowColor: '#000000',
        outline: false,
        outlineColor: '#000000',
        gradient: false
      },
      locked: false,
      visible: true
    }
  ]
  currentLayerIndex.value = 0
  showEditor.value = true
}

function saveEditor() {
  // ذخیره قالب با لایه‌ها
  const idx = customTemplates.value.findIndex(t => t.id === editorTemplate.value.id)
  if (idx !== -1) {
    customTemplates.value[idx].name = editorTemplate.value.name
    customTemplates.value[idx].layers = JSON.parse(JSON.stringify(layers.value))
  } else {
    // ایجاد قالب جدید
    const newTemplate = {
      id: Date.now(),
      name: editorTemplate.value.name,
      image: editorTemplate.value.image,
      layers: JSON.parse(JSON.stringify(layers.value))
    }
    customTemplates.value.push(newTemplate)
  }
  showEditor.value = false
}

function deleteTemplate(id) {
  if (confirm('آیا از حذف این قالب مطمئن هستید؟')) {
    customTemplates.value = customTemplates.value.filter(t => t.id !== id)
  }
}

// Drag & Drop logic قدیمی حذف شده - حالا از سیستم لایه‌ها استفاده می‌شود

// متدهای ویرایشگر
const addNewLayer = () => {
  const newLayer = {
    id: Date.now(),
    name: `لایه جدید ${layers.value.length + 1}`,
    type: 'text',
    text: 'متن جدید',
    color: '#000000',
    fontSize: 18,
    fontFamily: 'IRANSans',
    fontWeight: 'normal',
    rotation: 0,
    position: { x: 100, y: 100 },
    opacity: 100,
    effects: {
      shadow: false,
      shadowColor: '#000000',
      outline: false,
      outlineColor: '#000000',
      gradient: false
    },
    locked: false,
    visible: true
  }
  layers.value.push(newLayer)
  currentLayerIndex.value = layers.value.length - 1
}

const duplicateLayer = () => {
  if (currentLayer.value) {
    const duplicated = {
      ...currentLayer.value,
      id: Date.now(),
      name: `${currentLayer.value.name} (کپی)`,
      position: {
        x: currentLayer.value.position.x + 20,
        y: currentLayer.value.position.y + 20
      }
    }
    layers.value.push(duplicated)
    currentLayerIndex.value = layers.value.length - 1
  }
}

const deleteLayer = () => {
  if (layers.value.length > 1) {
    layers.value.splice(currentLayerIndex.value, 1)
    if (currentLayerIndex.value >= layers.value.length) {
      currentLayerIndex.value = layers.value.length - 1
    }
  }
}

const exportTemplate = () => {
  const templateData = {
    name: editorTemplate.value.name,
    image: editorTemplate.value.image,
    layers: layers.value,
    exportDate: new Date().toISOString()
  }
  
  const dataStr = JSON.stringify(templateData, null, 2)
  const dataBlob = new Blob([dataStr], { type: 'application/json' })
  const url = URL.createObjectURL(dataBlob)
  
  const link = document.createElement('a')
  link.href = url
  link.download = `giftcard-template-${Date.now()}.json`
  link.click()
  
  URL.revokeObjectURL(url)
}

const showTemplates = () => {
  showTemplatesModal.value = true
}

const applyReadyTemplate = (template) => {
  layers.value = JSON.parse(JSON.stringify(template.layers))
  editorTemplate.value.name = template.name
  currentLayerIndex.value = 0
  showTemplatesModal.value = false
}

const onLayerImageChange = (event) => {
  const file = event.target.files[0]
  if (file && currentLayer.value) {
    const reader = new FileReader()
    reader.onload = (e) => {
      currentLayer.value.imageUrl = e.target.result
      currentLayer.value.type = 'image'
    }
    reader.readAsDataURL(file)
  }
}

const startLayerDrag = (layerIndex, event) => {
  isLayerDragging.value = true
  dragLayerIndex.value = layerIndex
  layerDragStartPos.value = {
    x: event.clientX - layers.value[layerIndex].position.x,
    y: event.clientY - layers.value[layerIndex].position.y
  }
  
  document.addEventListener('mousemove', onLayerDrag)
  document.addEventListener('mouseup', stopLayerDrag)
}

const onLayerDrag = (event) => {
  if (isLayerDragging.value && dragLayerIndex.value >= 0) {
    const layer = layers.value[dragLayerIndex.value]
    layer.position.x = event.clientX - layerDragStartPos.value.x
    layer.position.y = event.clientY - layerDragStartPos.value.y
  }
}

const stopLayerDrag = () => {
  isLayerDragging.value = false
  dragLayerIndex.value = -1
  document.removeEventListener('mousemove', onLayerDrag)
  document.removeEventListener('mouseup', stopLayerDrag)
}

// متدهای undo/redo
const undoAction = () => {
  if (canUndo.value) {
    const action = undoStack.value.pop()
    redoStack.value.push(action)
    // اعمال عملیات برای بازگرداندن
    if (action.type === 'add') {
      layers.value.splice(action.index, 1)
      currentLayerIndex.value = action.newIndex
    } else if (action.type === 'delete') {
      layers.value.splice(action.index, 0, action.layer)
      currentLayerIndex.value = action.newIndex
    } else if (action.type === 'move') {
      layers.value[action.index].position = action.oldPosition
      layers.value[action.index].rotation = action.oldRotation
    } else if (action.type === 'edit') {
      layers.value[action.index] = action.layer
    } else if (action.type === 'duplicate') {
      layers.value.splice(action.index, 0, action.layer)
      currentLayerIndex.value = action.newIndex
    }
  }
}

const redoAction = () => {
  if (canRedo.value) {
    const action = redoStack.value.pop()
    undoStack.value.push(action)
    // اعمال عملیات برای تکرار
    if (action.type === 'add') {
      layers.value.splice(action.index, 0, action.layer)
      currentLayerIndex.value = action.newIndex
    } else if (action.type === 'delete') {
      layers.value.splice(action.index, 1)
      currentLayerIndex.value = action.newIndex
    } else if (action.type === 'move') {
      layers.value[action.index].position = action.oldPosition
      layers.value[action.index].rotation = action.oldRotation
    } else if (action.type === 'edit') {
      layers.value[action.index] = action.layer
    } else if (action.type === 'duplicate') {
      layers.value.splice(action.index, 0, action.layer)
      currentLayerIndex.value = action.newIndex
    }
  }
}

// تابع برای ذخیره عملیات
const saveAction = (type, layer, index, newIndex = null) => {
  const action = {
    type: type,
    layer: layer,
    index: index,
    newIndex: newIndex
  }
  undoStack.value.push(action)
  redoStack.value = [] // هر عملیات جدید، تاریخچه را پاک می‌کند
}
</script>

<style scoped>
/* استایل‌های خاص مدیریت قالب */

/* استایل‌های افکت‌های متن */
.text-shadow {
  text-shadow: 2px 2px 4px rgba(0,0,0,0.5);
}

.text-outline {
  -webkit-text-stroke: 1px;
  -webkit-text-stroke-color: currentColor;
}

.text-gradient {
  background: linear-gradient(45deg, var(--gradient-start), var(--gradient-end));
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

/* بهبود عملکرد Drag & Drop */
.cursor-move {
  cursor: grab;
}

.cursor-move:active {
  cursor: grabbing;
}

/* انیمیشن‌های نرم */
.transition-all {
  transition: all 0.2s ease-in-out;
}
</style> 

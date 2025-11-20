<template>
  <div class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
    <div class="relative top-20 mx-auto p-5 border w-11/12 max-w-4xl shadow-lg rounded-md bg-white">
      <!-- هدر مودال -->
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-lg font-semibold text-gray-900">ویرایش تم گیفت کارت</h3>
        <button
          class="text-gray-400 hover:text-gray-600"
          @click="$emit('close')"
        >
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
        </button>
      </div>

      <!-- فرم ویرایش تم -->
      <form class="space-y-6" @submit.prevent="updateTheme">
        <!-- اطلاعات اصلی -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              نام تم <span class="text-red-500">*</span>
            </label>
            <input
              v-model="form.name"
              type="text"
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="مثال: تم تولد شاد"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              دسته‌بندی <span class="text-red-500">*</span>
            </label>
            <select
              v-model="form.category"
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            >
              <option value="">انتخاب کنید</option>
              <option value="تولد">تولد</option>
              <option value="عروسی">عروسی</option>
              <option value="سال نو">سال نو</option>
              <option value="شرکتی">شرکتی</option>
              <option value="تخفیف">تخفیف</option>
              <option value="عمومی">عمومی</option>
            </select>
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            توضیحات
          </label>
          <textarea
            v-model="form.description"
            rows="3"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="توضیحات مربوط به تم..."
          ></textarea>
        </div>

        <!-- آیکون و رنگ‌ها -->
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              آیکون
            </label>
            <div class="grid grid-cols-6 gap-2">
              <button
                v-for="icon in availableIcons"
                :key="icon"
                type="button"
                :class="{
                  'bg-blue-100 border-blue-500': form.icon === icon,
                  'bg-gray-100 border-gray-300': form.icon !== icon
                }"
                class="w-10 h-10 flex items-center justify-center text-xl border-2 rounded-lg hover:bg-gray-50"
                @click="form.icon = icon"
              >
                {{ icon }}
              </button>
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              رنگ اصلی
            </label>
            <div class="flex items-center space-x-2">
              <input
                v-model="form.primaryColor"
                type="color"
                class="w-12 h-10 border border-gray-300 rounded-lg"
              />
              <input
                v-model="form.primaryColor"
                type="text"
                class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="#FF6B6B"
              />
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              رنگ پس‌زمینه
            </label>
            <div class="flex items-center space-x-2">
              <input
                v-model="form.backgroundColor"
                type="color"
                class="w-12 h-10 border border-gray-300 rounded-lg"
              />
              <input
                v-model="form.backgroundColor"
                type="text"
                class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="#FFF5F5"
              />
            </div>
          </div>
        </div>

        <!-- تنظیمات فونت -->
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              فونت
            </label>
            <select
              v-model="form.fontFamily"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            >
              <option value="IRANSans">IRANSans</option>
              <option value="IRANSansWeb">IRANSansWeb</option>
              <option value="Tahoma">Tahoma</option>
              <option value="Arial">Arial</option>
            </select>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              اندازه فونت
            </label>
            <select
              v-model="form.fontSize"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            >
              <option value="12px">کوچک (12px)</option>
              <option value="14px">متوسط کوچک (14px)</option>
              <option value="16px">متوسط (16px)</option>
              <option value="18px">متوسط بزرگ (18px)</option>
              <option value="20px">بزرگ (20px)</option>
              <option value="24px">خیلی بزرگ (24px)</option>
            </select>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              رنگ متن
            </label>
            <div class="flex items-center space-x-2">
              <input
                v-model="form.textColor"
                type="color"
                class="w-12 h-10 border border-gray-300 rounded-lg"
              />
              <input
                v-model="form.textColor"
                type="text"
                class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="#2D3748"
              />
            </div>
          </div>
        </div>

        <!-- تنظیمات ظاهری -->
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              رنگ حاشیه
            </label>
            <div class="flex items-center space-x-2">
              <input
                v-model="form.borderColor"
                type="color"
                class="w-12 h-10 border border-gray-300 rounded-lg"
              />
              <input
                v-model="form.borderColor"
                type="text"
                class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="#FFE2E2"
              />
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              شعاع گوشه
            </label>
            <select
              v-model="form.borderRadius"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            >
              <option value="4px">کوچک (4px)</option>
              <option value="8px">متوسط کوچک (8px)</option>
              <option value="12px">متوسط (12px)</option>
              <option value="16px">متوسط بزرگ (16px)</option>
              <option value="20px">بزرگ (20px)</option>
              <option value="24px">خیلی بزرگ (24px)</option>
            </select>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              سایه
            </label>
            <select
              v-model="form.shadow"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            >
              <option value="none">بدون سایه</option>
              <option value="0 1px 3px rgba(0, 0, 0, 0.1)">سایه کوچک</option>
              <option value="0 4px 6px rgba(0, 0, 0, 0.1)">سایه متوسط</option>
              <option value="0 10px 15px rgba(0, 0, 0, 0.1)">سایه بزرگ</option>
              <option value="0 20px 25px rgba(0, 0, 0, 0.15)">سایه خیلی بزرگ</option>
            </select>
          </div>
        </div>

        <!-- انیمیشن -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            انیمیشن
          </label>
          <select
            v-model="form.animation"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          >
            <option value="none">بدون انیمیشن</option>
            <option value="fadeIn">ظهور تدریجی</option>
            <option value="slideIn">لغزش</option>
            <option value="bounce">پرش</option>
            <option value="pulse">ضربان</option>
            <option value="rotate">چرخش</option>
          </select>
        </div>

        <!-- CSS سفارشی -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            CSS سفارشی
          </label>
          <textarea
            v-model="form.customCSS"
            rows="4"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent font-mono text-sm"
            placeholder="/* کد CSS سفارشی خود را اینجا وارد کنید */"
          ></textarea>
          <p class="text-xs text-gray-500 mt-1">
            می‌توانید کد CSS سفارشی برای تنظیمات بیشتر اضافه کنید
          </p>
        </div>

        <!-- تنظیمات اضافی -->
        <div class="flex items-center space-x-4">
          <label class="flex items-center">
            <input
              v-model="form.isDefault"
              type="checkbox"
              class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
            />
            <span class="mr-2 text-sm text-gray-700">تم پیش‌فرض</span>
          </label>

          <label class="flex items-center">
            <input
              v-model="form.isActive"
              type="checkbox"
              class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
            />
            <span class="mr-2 text-sm text-gray-700">فعال</span>
          </label>
        </div>

        <!-- آمار تم -->
        <div class="bg-gray-50 p-6 rounded-lg">
          <h4 class="text-sm font-medium text-gray-700 mb-3">آمار تم</h4>
          <div class="grid grid-cols-2 md:grid-cols-4 gap-6">
            <div class="text-center">
              <p class="text-2xl font-bold text-blue-600">{{ theme.usageCount }}</p>
              <p class="text-xs text-gray-500">استفاده شده</p>
            </div>
            <div class="text-center">
              <p class="text-2xl font-bold text-green-600">{{ theme.rating }}</p>
              <p class="text-xs text-gray-500">امتیاز</p>
            </div>
            <div class="text-center">
              <p class="text-2xl font-bold text-purple-600">{{ theme.id }}</p>
              <p class="text-xs text-gray-500">شناسه</p>
            </div>
            <div class="text-center">
              <p class="text-2xl font-bold text-orange-600">{{ getStatusText(theme.status) }}</p>
              <p class="text-xs text-gray-500">وضعیت</p>
            </div>
          </div>
        </div>

        <!-- پیش‌نمایش -->
        <div class="border border-gray-200 rounded-lg p-6">
          <h4 class="text-sm font-medium text-gray-700 mb-3">پیش‌نمایش تم</h4>
          <div
            class="p-6 rounded-lg text-center"
            :style="{
              backgroundColor: form.backgroundColor,
              color: form.textColor,
              fontFamily: form.fontFamily,
              fontSize: form.fontSize,
              border: `2px solid ${form.borderColor}`,
              borderRadius: form.borderRadius,
              boxShadow: form.shadow
            }"
          >
            <div class="text-4xl mb-3">{{ form.icon || '🎁' }}</div>
            <h3 class="text-xl font-bold mb-2">{{ form.name || 'نام تم' }}</h3>
            <p class="text-sm opacity-80">{{ form.description || 'توضیحات تم' }}</p>
            <div class="mt-4 p-3 rounded" :style="{ backgroundColor: form.primaryColor, color: 'white' }">
              مبلغ: ۵۰,۰۰۰ تومان
            </div>
          </div>
        </div>

        <!-- دکمه‌های عملیات -->
        <div class="flex justify-end space-x-3 space-x-reverse">
          <button
            type="button"
            class="px-4 py-2 bg-gray-300 text-gray-700 text-sm font-medium rounded-lg hover:bg-gray-400 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
            @click="$emit('close')"
          >
            انصراف
          </button>
          <button
            type="submit"
            class="px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
          >
            بروزرسانی تم
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
// تعریف props و emits
const props = defineProps({
  theme: {
    type: Object,
    required: true
  }
})

const emit = defineEmits(['close', 'updated'])

// تعریف متغیرهای reactive
const form = ref({
  name: '',
  description: '',
  category: '',
  icon: '🎁',
  primaryColor: '#FF6B6B',
  backgroundColor: '#FFF5F5',
  textColor: '#2D3748',
  borderColor: '#FFE2E2',
  borderRadius: '12px',
  shadow: '0 4px 6px rgba(0, 0, 0, 0.1)',
  fontFamily: 'IRANSans',
  fontSize: '16px',
  animation: 'fadeIn',
  customCSS: '',
  isDefault: false,
  isActive: true
})

// آیکون‌های موجود
const availableIcons = [
  '🎁', '🎂', '💒', '🎊', '🎯', '🏢', '💝', '🎉', '🎈', '🎪',
  '🎨', '🎭', '🎪', '🎫', '🎬', '🎤', '🎧', '🎼', '🎹', '🎸',
  '🎺', '🎻', '🎷', '🥁', '🎮', '🎲', '🎳', '🎯', '🎱', '🎾',
  '🏀', '⚽', '🏈', '⚾', '🎾', '🏐', '🏉', '🎱', '🏓', '🏸',
  '🏊', '🏄', '🚴', '🏇', '🏂', '⛷️', '🏋️', '🤸', '🤺', '🤾'
]

// تابع دریافت متن وضعیت
const getStatusText = (status) => {
  const statusMap = {
    active: 'فعال',
    inactive: 'غیرفعال',
    pending: 'در انتظار'
  }
  return statusMap[status] || status
}

// مقداردهی اولیه فرم با داده‌های تم
onMounted(() => {
  if (props.theme) {
    form.value = {
      name: props.theme.name || '',
      description: props.theme.description || '',
      category: props.theme.category || '',
      icon: props.theme.icon || '🎁',
      primaryColor: props.theme.primaryColor || '#FF6B6B',
      backgroundColor: props.theme.backgroundColor || '#FFF5F5',
      textColor: props.theme.textColor || '#2D3748',
      borderColor: props.theme.borderColor || '#FFE2E2',
      borderRadius: props.theme.borderRadius || '12px',
      shadow: props.theme.shadow || '0 4px 6px rgba(0, 0, 0, 0.1)',
      fontFamily: props.theme.fontFamily || 'IRANSans',
      fontSize: props.theme.fontSize || '16px',
      animation: props.theme.animation || 'fadeIn',
      customCSS: props.theme.customCSS || '',
      isDefault: props.theme.isDefault || false,
      isActive: props.theme.status === 'active'
    }
  }
})

// تابع بروزرسانی تم
const updateTheme = () => {
  // اعتبارسنجی فرم
  if (!form.value.name || !form.value.category) {
    alert('لطفاً نام و دسته‌بندی تم را وارد کنید')
    return
  }

  // بروزرسانی تم
  const updatedTheme = {
    ...props.theme,
    ...form.value,
    status: form.value.isActive ? 'active' : 'inactive',
    updatedAt: new Date().toISOString()
  }

  // ارسال به کامپوننت والد
  emit('updated', updatedTheme)
}
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 

<template>
  <div class="bg-white rounded-lg shadow p-6">
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-xl font-semibold text-gray-900">مدیریت سطوح وفاداری</h2>
      <button 
        class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors"
        @click="showLevelForm = true; editingLevel = null"
      >
        افزودن سطح جدید
      </button>
    </div>

    <!-- لیست سطوح -->
    <div class="overflow-x-auto">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              سطح
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              نام
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              شرایط ارتقا
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              مزایا
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              تعداد اعضا
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              وضعیت
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              عملیات
            </th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-for="level in loyaltyLevels" :key="level.id">
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="flex items-center">
                <div :class="getLevelColor(level.name)" class="w-8 h-8 rounded-full flex items-center justify-center">
                  <span class="text-white font-semibold text-sm">{{ level.icon }}</span>
                </div>
              </div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="text-sm font-medium text-gray-900">{{ level.name }}</div>
              <div class="text-sm text-gray-500">{{ level.description }}</div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="text-sm text-gray-900">
                <div>حداقل خرید: {{ level.minPurchase }} تومان</div>
                <div>تعداد خرید: {{ level.minOrders }} بار</div>
              </div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="text-sm text-gray-900">
                <div>تخفیف: {{ level.discount }}%</div>
                <div>ارسال رایگان: {{ level.freeShipping ? 'بله' : 'خیر' }}</div>
              </div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="text-sm text-gray-900">{{ level.memberCount }}</div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <span :class="level.active ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'" class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full">
                {{ level.active ? 'فعال' : 'غیرفعال' }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
              <button 
                class="text-blue-600 hover:text-blue-900 ml-4"
                @click="editingLevel = level; showLevelForm = true"
              >
                ویرایش
              </button>
              <button 
                :class="level.active ? 'text-red-600 hover:text-red-900' : 'text-green-600 hover:text-green-900'"
                @click="toggleLevelStatus(level)"
              >
                {{ level.active ? 'غیرفعال' : 'فعال' }}
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- مودال فرم سطح -->
    <div v-if="showLevelForm" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
      <div class="relative top-20 mx-auto p-5 border w-full max-w-md sm:max-w-lg md:max-w-xl shadow-lg rounded-md bg-white">
        <div class="mt-3">
          <h3 class="text-lg font-medium text-gray-900 mb-4">
            {{ editingLevel ? 'ویرایش سطح' : 'افزودن سطح جدید' }}
          </h3>
          
          <form @submit.prevent="saveLevel">
            <div class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-gray-700">نام سطح</label>
                <input 
                  v-model="levelForm.name" 
                  type="text" 
                  class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                  required
                >
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700">توضیحات</label>
                <textarea 
                  v-model="levelForm.description" 
                  rows="2"
                  class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                ></textarea>
              </div>
              
              <div class="grid grid-cols-2 gap-6">
                <div>
                  <label class="block text-sm font-medium text-gray-700">حداقل خرید (تومان)</label>
                  <input 
                    v-model="levelForm.minPurchase" 
                    type="number" 
                    class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                    required
                  >
                </div>
                
                <div>
                  <label class="block text-sm font-medium text-gray-700">حداقل تعداد خرید</label>
                  <input 
                    v-model="levelForm.minOrders" 
                    type="number" 
                    class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                    required
                  >
                </div>
              </div>
              
              <div class="grid grid-cols-2 gap-6">
                <div>
                  <label class="block text-sm font-medium text-gray-700">درصد تخفیف</label>
                  <input 
                    v-model="levelForm.discount" 
                    type="number" 
                    min="0" 
                    max="100"
                    class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                    required
                  >
                </div>
                
                <div>
                  <label class="block text-sm font-medium text-gray-700">ارسال رایگان</label>
                  <select 
                    v-model="levelForm.freeShipping"
                    class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                  >
                    <option :value="false">خیر</option>
                    <option :value="true">بله</option>
                  </select>
                </div>
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700">آیکون</label>
                <input 
                  v-model="levelForm.icon" 
                  type="text" 
                  placeholder="🏆"
                  class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                >
              </div>
            </div>
            
            <div class="flex justify-end space-x-3 space-x-reverse mt-6">
              <button 
                type="button"
                class="bg-gray-300 text-gray-700 px-4 py-2 rounded-md hover:bg-gray-400 transition-colors"
                @click="showLevelForm = false"
              >
                انصراف
              </button>
              <button 
                type="submit"
                class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition-colors"
              >
                {{ editingLevel ? 'ویرایش' : 'افزودن' }}
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'

// سطوح وفاداری
const loyaltyLevels = ref([
  {
    id: 1,
    name: 'برنزی',
    description: 'سطح پایه برای مشتریان جدید',
    icon: '🥉',
    minPurchase: 0,
    minOrders: 0,
    discount: 5,
    freeShipping: false,
    memberCount: 850,
    active: true
  },
  {
    id: 2,
    name: 'نقره‌ای',
    description: 'سطح متوسط برای مشتریان وفادار',
    icon: '🥈',
    minPurchase: 500000,
    minOrders: 3,
    discount: 10,
    freeShipping: false,
    memberCount: 320,
    active: true
  },
  {
    id: 3,
    name: 'طلایی',
    description: 'سطح بالا برای مشتریان ویژه',
    icon: '🥇',
    minPurchase: 2000000,
    minOrders: 10,
    discount: 15,
    freeShipping: true,
    memberCount: 68,
    active: true
  },
  {
    id: 4,
    name: 'الماس',
    description: 'سطح VIP برای مشتریان برتر',
    icon: '💎',
    minPurchase: 5000000,
    minOrders: 25,
    discount: 20,
    freeShipping: true,
    memberCount: 12,
    active: true
  }
])

// State برای مودال
const showLevelForm = ref(false)
const editingLevel = ref(null)

// فرم سطح
const levelForm = reactive({
  name: '',
  description: '',
  icon: '',
  minPurchase: 0,
  minOrders: 0,
  discount: 0,
  freeShipping: false
})

// رنگ‌های سطوح
function getLevelColor(levelName: string) {
  const colors = {
    'برنزی': 'bg-yellow-600',
    'نقره‌ای': 'bg-gray-500',
    'طلایی': 'bg-yellow-500',
    'الماس': 'bg-blue-600'
  }
  return colors[levelName] || 'bg-gray-400'
}

// تغییر وضعیت سطح
function toggleLevelStatus(level: any) {
  level.active = !level.active
  // TODO: فراخوانی API برای تغییر وضعیت
}

// ذخیره سطح
function saveLevel() {
  if (editingLevel.value) {
    // ویرایش سطح موجود
    const index = loyaltyLevels.value.findIndex(l => l.id === editingLevel.value.id)
    if (index !== -1) {
      loyaltyLevels.value[index] = { ...editingLevel.value, ...levelForm }
    }
  } else {
    // افزودن سطح جدید
    const newLevel = {
      id: Date.now(),
      ...levelForm,
      memberCount: 0,
      active: true
    }
    loyaltyLevels.value.push(newLevel)
  }
  
  // پاک کردن فرم
  Object.assign(levelForm, {
    name: '',
    description: '',
    icon: '',
    minPurchase: 0,
    minOrders: 0,
    discount: 0,
    freeShipping: false
  })
  
  showLevelForm.value = false
  editingLevel.value = null
}
</script> 

<template>
  <div class="cart-root bg-[#f7fafd] min-h-[100vh] py-10" dir="rtl" style="font-family: 'IranYekan', 'Vazir', 'IRANSansX', 'Tahoma', sans-serif;">
    <div class="max-w-7xl mx-auto">
      <!-- تب سبد خرید و خرید بعدی -->
      <div class="flex justify-start border-b border-gray-200 mb-8 bg-white shadow-sm p-6 relative">
        <div class="flex">
          <button :class="['tab-btn', activeTab===0 ? 'active' : '']" @click="activeTab=0">
            سبد خرید <span v-if="currentCart?.length" class="tab-badge">{{ currentCart?.length }}</span>
          </button>
          <button :class="['tab-btn', activeTab===1 ? 'active' : '']" @click="activeTab=1">
            خرید بعدی <span v-if="nextItems?.length" class="tab-badge">{{ nextItems?.length }}</span>
          </button>
        </div>
        <hr class="absolute left-0 right-0 border-gray-300" style="height: 1px; background-color: #d1d5db; bottom: 14px;" />
      </div>

      <div v-if="activeTab===0">
        <!-- حالت سبد خرید خالی -->
        <div v-if="currentCart.length === 0" class="bg-white rounded-2xl shadow-sm p-16 text-center min-h-[300px] flex items-center justify-center">
          <div class="text-gray-400 text-lg">سبد خرید شما خالی است</div>
        </div>
        
        <!-- حالت سبد خرید دارای محصول -->
        <div v-else class="flex flex-col lg:flex-row-reverse gap-8">
          <!-- خلاصه سفارش و پنل اکشن (سمت راست) -->
          <div class="w-full lg:w-[350px] flex flex-col gap-6 mt-6 lg:mt-0">
            <div>
              <CartSummary :total="cartTotal" :count="currentCart.length" />
            </div>
          </div>

          <!-- لیست سبد خرید (سمت چپ) -->
          <div class="flex-1">
            <div class="bg-white shadow-sm p-6 mb-4 flex justify-between items-center relative border-b border-gray-200">
              <div class="text-[#1a2341] font-bold text-lg">سبد خرید</div>
              <!-- منوی سه‌نقطه‌ای اکشن -->
              <div class="dropdown inline-block relative">
                <button class="flex items-center justify-center w-10 h-10 rounded-full hover:bg-gray-100 focus:outline-none" @click="showActions = !showActions">
                  <svg class="w-6 h-6 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24"><circle cx="12" cy="6" r="1.5"/><circle cx="12" cy="12" r="1.5"/><circle cx="12" cy="18" r="1.5"/></svg>
                </button>
                <div v-if="showActions" class="absolute left-0 mt-2 w-48 bg-white rounded-2xl shadow-lg p-2 text-sm z-20">
                  <button class="flex items-center justify-between w-full px-3 py-2 hover:text-[#008eb2]" @click="moveAllToNext">
                    <span>انتقال همه به خرید بعدی</span>
                    <i class="ri-calendar-2-line"></i>
                  </button>
                  <button class="flex items-center justify-between w-full px-3 py-2 text-red-600 hover:text-red-700" @click="removeAll">
                    <span>حذف همه</span>
                    <i class="ri-delete-bin-6-line"></i>
                  </button>
                </div>
              </div>
            </div>
            <div class="bg-white rounded-2xl shadow-sm p-6">
              <CartList :items="currentCart" :is-next="false" @move-to-next="moveToNext" @remove="removeFromCart" @change-qty="changeQty" />
            </div>
            <hr class="my-4 border-gray-200" />
          </div>
        </div>
      </div>
      
      <div v-else>
        <!-- حالت خرید بعدی خالی -->
        <div v-if="nextItems.length === 0" class="bg-white rounded-2xl shadow-sm p-16 text-center min-h-[300px] flex items-center justify-center">
          <div class="text-gray-400 text-lg">خرید بعدی شما خالی است</div>
        </div>
        
        <!-- حالت خرید بعدی دارای محصول -->
        <div v-else class="flex flex-col lg:flex-row-reverse gap-8">
          <!-- لیست خرید بعدی -->
          <div class="flex-1">
            <div class="bg-white shadow-sm p-6 mb-4 flex justify-between items-center border-b border-gray-200">
              <div class="text-[#1a2341] font-bold text-lg">خرید بعدی</div>
              <button class="p-2 rounded-full hover:bg-gray-100">
                <svg class="w-6 h-6 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24"><circle cx="12" cy="6" r="1.5"/><circle cx="12" cy="12" r="1.5"/><circle cx="12" cy="18" r="1.5"/></svg>
              </button>
            </div>
            <div class="bg-white rounded-2xl shadow-sm p-6">
              <CartList :items="nextItems" :is-next="true" @move-to-cart="moveToCart" @remove="removeFromNext" />
            </div>
            <hr class="my-4 border-gray-200" />
          </div>
        </div>
      </div>
      
      <!-- اسلایدر محصولات پیشنهادی -->
      <div v-if="currentCart.length > 0" class="mt-12">
        <CartRecommendations />
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, onMounted, ref } from 'vue'
import CartList from './CartList.vue'
import CartRecommendations from './CartRecommendations.vue'
import CartSummary from './CartSummary.vue'

// استفاده از composable سبد خرید
const { 
  cartItems, 
  cartTotal, 
  fetchCart, 
  updateCartItem, 
  removeFromCart,
  moveCartItemToNext,
  moveCartItemToCart,
  clearCart
} = useCart()

const activeTab = ref(0)
const nextItems = computed(()=> cartItems.value.filter(i=>i.is_next))
const currentCart = computed(()=> cartItems.value.filter(i=>!i.is_next))
const showActions = ref(false)

// دریافت سبد خرید در ابتدای بارگذاری صفحه
onMounted(async () => {
  try {
    await fetchCart()
  } catch {
    // خطا در بارگذاری سبد خرید
  }
})

// انتقال به خرید بعدی
function moveToNext(item) {
  if (item && item.id) {
    moveCartItemToNext(item.id)
  }
}

// انتقال به سبد خرید
function moveToCart(item) {
  if (item && item.id) {
    moveCartItemToCart(item.id)
  }
}

// حذف از خرید بعدی
function removeFromNext(item) {
  if (item && item.id) {
    // حذف آیتم خرید بعدی نیز از دیتابیس
    removeFromCart(item.id)
  }
}

// تغییر تعداد
async function changeQty({item, val}) {
  if (item && item.id && val > 0) {
    try {
      const result = await updateCartItem(item.id, val)
      if (!result.success) {
        // نمایش پیام خطا
      }
    } catch {
      // خطا در تغییر تعداد
    }
  }
}

// انتقال همه کالاها به خرید بعدی
function moveAllToNext() {
  currentCart.value.forEach(it => moveCartItemToNext(it.id))
  showActions.value = false
}

// حذف همه کالاها از سبد خرید
async function removeAll() {
  try {
    const result = await clearCart()
    if (result.success) {
      await fetchCart() // به‌روزرسانی سبد خرید
    }
  } catch {
    // console.error('خطا در حذف همه آیتم‌ها:', error)
  }
  showActions.value = false
}
</script>

<style scoped>
.cart-root {
  font-family: 'IranYekan', 'Vazir', 'IRANSansX', 'Tahoma', sans-serif;
  background: linear-gradient(135deg, #f8fafc 0%, #e0e7ff 100%);
}
.tab-btn {
  position: relative;
  font-weight: 700;
  font-size: 1rem;
  padding: 0.75rem 1rem;
  background: none;
  border: none;
  outline: none;
  color: #64748b;
  cursor: pointer;
  border-bottom: 2px solid transparent;
  transition: all .18s;
  font-family: 'IranYekan', 'Vazir', 'IRANSansX', 'Tahoma', sans-serif;
}
.tab-btn.active {
  color: #7e22ce;
  border-bottom: 2px solid #7e22ce;
}
.tab-badge {
  background: #7e22ce;
  color: #fff;
  font-size: 0.7rem;
  border-radius: 0.25rem;
  padding: 0.1rem 0.4rem;
  margin-right: 0.3rem;
  font-family: 'IranYekan', 'Vazir', 'IRANSansX', 'Tahoma', sans-serif;
}
</style> 

<template>
  <div class="space-y-6">
    <!-- اطلاعات ناوگان -->
    <div class="px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-blue-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7h12m0 0l-4-4m4 4l-4 4m0 6H4m0 0l4 4m-4-4l4-4"></path>
        </svg>
        اطلاعات ناوگان
      </h3>

      <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
        <!-- راننده -->
        <div>
          <h4 class="text-md font-medium text-gray-700 mb-3">اطلاعات راننده</h4>
          <div class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">نام راننده</label>
              <input 
                v-model="advancedLogistics.driverName"
                type="text" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="نام راننده"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">شماره موبایل</label>
              <input 
                v-model="advancedLogistics.driverPhone"
                type="tel" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="09123456789"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">شماره پلاک</label>
              <input 
                v-model="advancedLogistics.vehiclePlate"
                type="text" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="12-345-67"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">وضعیت راننده</label>
              <select 
                v-model="advancedLogistics.driverStatus"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="available">در دسترس</option>
                <option value="busy">مشغول</option>
                <option value="offline">آفلاین</option>
                <option value="break">استراحت</option>
              </select>
            </div>
          </div>
        </div>

        <!-- وسیله نقلیه -->
        <div>
          <h4 class="text-md font-medium text-gray-700 mb-3">اطلاعات وسیله نقلیه</h4>
          <div class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">نوع وسیله نقلیه</label>
              <select 
                v-model="advancedLogistics.vehicleType"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="motorcycle">موتورسیکلت</option>
                <option value="car">خودرو</option>
                <option value="van">ون</option>
                <option value="truck">کامیون</option>
                <option value="bicycle">دوچرخه</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">ظرفیت بار (کیلوگرم)</label>
              <input 
                v-model="advancedLogistics.vehicleCapacity"
                type="number" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="100"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">وضعیت وسیله نقلیه</label>
              <select 
                v-model="advancedLogistics.vehicleStatus"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="active">فعال</option>
                <option value="maintenance">در تعمیر</option>
                <option value="fuel">سوخت</option>
                <option value="inactive">غیرفعال</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">آخرین سرویس</label>
              <input 
                v-model="advancedLogistics.lastServiceDate"
                type="date" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- پیگیری زنده -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-green-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 20l-5.447-2.724A1 1 0 013 16.382V5.618a1 1 0 011.447-.894L9 7m0 13l6-3m-6 3V7m6 10l4.553 2.276A1 1 0 0021 18.382V7.618a1 1 0 00-1.447-.894L15 4m0 13V4m-6 3l6-3"></path>
        </svg>
        پیگیری زنده
      </h3>

      <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
        <!-- موقعیت فعلی -->
        <div>
          <h4 class="text-md font-medium text-gray-700 mb-3">موقعیت فعلی</h4>
          <div class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">عرض جغرافیایی</label>
              <input 
                v-model="advancedLogistics.currentLatitude"
                type="number" 
                step="0.000001"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="35.6892"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">طول جغرافیایی</label>
              <input 
                v-model="advancedLogistics.currentLongitude"
                type="number" 
                step="0.000001"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="51.3890"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">آدرس فعلی</label>
              <textarea 
                v-model="advancedLogistics.currentAddress"
                rows="2"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="آدرس فعلی راننده..."
              ></textarea>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">آخرین به‌روزرسانی</label>
              <input 
                v-model="advancedLogistics.lastLocationUpdate"
                type="datetime-local" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
          </div>
        </div>

        <!-- مسیر و زمان‌بندی -->
        <div>
          <h4 class="text-md font-medium text-gray-700 mb-3">مسیر و زمان‌بندی</h4>
          <div class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">زمان تخمینی رسیدن</label>
              <input 
                v-model="advancedLogistics.estimatedArrival"
                type="datetime-local" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">فاصله باقی‌مانده (کیلومتر)</label>
              <input 
                v-model="advancedLogistics.remainingDistance"
                type="number" 
                step="0.1"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="5.2"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">سرعت فعلی (کیلومتر بر ساعت)</label>
              <input 
                v-model="advancedLogistics.currentSpeed"
                type="number" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="45"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">وضعیت ترافیک</label>
              <select 
                v-model="advancedLogistics.trafficStatus"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="clear">آزاد</option>
                <option value="moderate">متوسط</option>
                <option value="heavy">سنگین</option>
                <option value="congested">بسیار سنگین</option>
              </select>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- زمان‌بندی تحویل -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-purple-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
        </svg>
        زمان‌بندی تحویل
      </h3>

      <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
        <!-- بازه‌های زمانی -->
        <div>
          <h4 class="text-md font-medium text-gray-700 mb-3">بازه‌های زمانی تحویل</h4>
          <div class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">بازه زمانی انتخاب شده</label>
              <select 
                v-model="advancedLogistics.selectedTimeSlot"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="">انتخاب بازه زمانی</option>
                <option value="morning">صبح (8-12)</option>
                <option value="afternoon">ظهر (12-16)</option>
                <option value="evening">عصر (16-20)</option>
                <option value="night">شب (20-24)</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">زمان شروع بازه</label>
              <input 
                v-model="advancedLogistics.timeSlotStart"
                type="time" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">زمان پایان بازه</label>
              <input 
                v-model="advancedLogistics.timeSlotEnd"
                type="time" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">وضعیت بازه زمانی</label>
              <select 
                v-model="advancedLogistics.timeSlotStatus"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="available">در دسترس</option>
                <option value="booked">رزرو شده</option>
                <option value="full">پر</option>
                <option value="cancelled">لغو شده</option>
              </select>
            </div>
          </div>
        </div>

        <!-- تنظیمات تحویل -->
        <div>
          <h4 class="text-md font-medium text-gray-700 mb-3">تنظیمات تحویل</h4>
          <div class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">نوع تحویل</label>
              <select 
                v-model="advancedLogistics.deliveryType"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="door">تحویل درب منزل</option>
                <option value="office">تحویل در محل کار</option>
                <option value="neighbor">تحویل به همسایه</option>
                <option value="security">تحویل به نگهبان</option>
                <option value="pickup">تحویل حضوری</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">اولویت تحویل</label>
              <select 
                v-model="advancedLogistics.deliveryPriority"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="low">کم</option>
                <option value="normal">عادی</option>
                <option value="high">بالا</option>
                <option value="urgent">فوری</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">هزینه تحویل ویژه</label>
              <input 
                v-model="advancedLogistics.specialDeliveryFee"
                type="number" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="0"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">یادداشت تحویل</label>
              <textarea 
                v-model="advancedLogistics.deliveryNotes"
                rows="2"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="یادداشت‌های مربوط به تحویل..."
              ></textarea>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- مقررات و محدودیت‌ها -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-red-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
        </svg>
        مقررات و محدودیت‌ها
      </h3>

      <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
        <!-- محدودیت‌های منطقه‌ای -->
        <div>
          <h4 class="text-md font-medium text-gray-700 mb-3">محدودیت‌های منطقه‌ای</h4>
          <div class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">منطقه تحویل</label>
              <select 
                v-model="advancedLogistics.deliveryZone"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="zone1">منطقه 1 - مرکز شهر</option>
                <option value="zone2">منطقه 2 - شمال شهر</option>
                <option value="zone3">منطقه 3 - جنوب شهر</option>
                <option value="zone4">منطقه 4 - شرق شهر</option>
                <option value="zone5">منطقه 5 - غرب شهر</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">محدودیت ترافیکی</label>
              <select 
                v-model="advancedLogistics.trafficRestriction"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="none">بدون محدودیت</option>
                <option value="odd_even">زوج و فرد</option>
                <option value="environmental">محدودیت زیست محیطی</option>
                <option value="construction">محدودیت ساخت و ساز</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">ساعت محدودیت</label>
              <input 
                v-model="advancedLogistics.restrictionHours"
                type="text" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="6-22"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">وضعیت محدودیت</label>
              <select 
                v-model="advancedLogistics.restrictionStatus"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="active">فعال</option>
                <option value="inactive">غیرفعال</option>
                <option value="temporary">موقت</option>
              </select>
            </div>
          </div>
        </div>

        <!-- تنظیمات لجستیک -->
        <div>
          <h4 class="text-md font-medium text-gray-700 mb-3">تنظیمات لجستیک</h4>
          <div class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">مرکز توزیع</label>
              <select 
                v-model="advancedLogistics.distributionCenter"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="center1">مرکز توزیع 1 - تهران</option>
                <option value="center2">مرکز توزیع 2 - اصفهان</option>
                <option value="center3">مرکز توزیع 3 - مشهد</option>
                <option value="center4">مرکز توزیع 4 - شیراز</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">مسیر بهینه</label>
              <select 
                v-model="advancedLogistics.optimizedRoute"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="fastest">سریع‌ترین مسیر</option>
                <option value="shortest">کوتاه‌ترین مسیر</option>
                <option value="economic">مقرون به صرفه</option>
                <option value="scenic">مسیر زیبا</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">حداکثر توقف</label>
              <input 
                v-model="advancedLogistics.maxStops"
                type="number" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="3"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">یادداشت لجستیک</label>
              <textarea 
                v-model="advancedLogistics.logisticsNotes"
                rows="2"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="یادداشت‌های لجستیک..."
              ></textarea>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تاریخچه لجستیک -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-orange-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"></path>
        </svg>
        تاریخچه لجستیک
      </h3>
      
      <div class="space-y-3">
        <div 
          v-for="(history, index) in advancedLogistics.logisticsHistory" 
          :key="index"
          class="flex items-center p-3 bg-gray-50 rounded-lg"
        >
          <div class="w-3 h-3 rounded-full mr-3" :class="getHistoryStatusColor(history.status)"></div>
          <div class="flex-1">
            <div class="text-sm font-medium text-gray-900">{{ history.action }}</div>
            <div class="text-xs text-gray-500">{{ formatDate(history.date) }} - {{ history.location }}</div>
            <div v-if="history.details" class="text-xs text-gray-600 mt-1">{{ history.details }}</div>
          </div>
          <div class="text-right">
            <div class="text-xs text-gray-500">{{ formatTime(history.date) }}</div>
            <div class="text-xs text-gray-500">{{ history.status }}</div>
          </div>
        </div>
      </div>

      <!-- دکمه افزودن تاریخچه -->
      <div class="mt-4">
        <button 
          class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors flex items-center text-sm"
          @click="addLogisticsHistory"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
          </svg>
          افزودن تاریخچه لجستیک
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'

// تعریف props
const props = defineProps({
  modelValue: {
    type: Object,
    default: () => ({
      driverName: '',
      driverPhone: '',
      vehiclePlate: '',
      driverStatus: 'available',
      vehicleType: 'car',
      vehicleCapacity: 100,
      vehicleStatus: 'active',
      lastServiceDate: '',
      currentLatitude: 35.6892,
      currentLongitude: 51.3890,
      currentAddress: '',
      lastLocationUpdate: '',
      estimatedArrival: '',
      remainingDistance: 0,
      currentSpeed: 0,
      trafficStatus: 'clear',
      selectedTimeSlot: '',
      timeSlotStart: '',
      timeSlotEnd: '',
      timeSlotStatus: 'available',
      deliveryType: 'door',
      deliveryPriority: 'normal',
      specialDeliveryFee: 0,
      deliveryNotes: '',
      deliveryZone: 'zone1',
      trafficRestriction: 'none',
      restrictionHours: '',
      restrictionStatus: 'inactive',
      distributionCenter: 'center1',
      optimizedRoute: 'fastest',
      maxStops: 3,
      logisticsNotes: '',
      logisticsHistory: []
    })
  }
})

// تعریف emit
const emit = defineEmits(['update:modelValue'])

// متغیر محلی
const advancedLogistics = ref({ ...props.modelValue })

// توابع کمکی
const getHistoryStatusColor = (status) => {
  const colors = {
    'started': 'bg-blue-400',
    'in_progress': 'bg-yellow-400',
    'completed': 'bg-green-400',
    'delayed': 'bg-orange-400',
    'cancelled': 'bg-red-400'
  }
  return colors[status] || 'bg-gray-400'
}

const formatDate = (date) => {
  if (!date) return 'نامشخص'
  return new Date(date).toLocaleDateString('fa-IR')
}

const formatTime = (date) => {
  if (!date) return ''
  return new Date(date).toLocaleTimeString('fa-IR', { 
    hour: '2-digit', 
    minute: '2-digit' 
  })
}

const addLogisticsHistory = () => {
  const newEntry = {
    action: 'شروع تحویل',
    date: new Date().toISOString(),
    location: 'مرکز توزیع',
    status: 'started',
    details: 'شروع فرآیند تحویل سفارش'
  }
  advancedLogistics.value.logisticsHistory.unshift(newEntry)
}

// نظارت بر تغییرات
watch(advancedLogistics, (newValue) => {
  emit('update:modelValue', newValue)
}, { deep: true })
</script> 

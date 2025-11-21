<template>
  <div class="bg-white rounded-lg shadow">
    <div class="px-6 py-4 border-b border-gray-200">
      <div class="flex items-center justify-between">
        <h3 class="text-lg font-medium text-gray-900">سیستم امتیازدهی خودکار</h3>
        <div class="flex items-center space-x-3 space-x-reverse">
          <label class="flex items-center">
            <input v-model="isAutoScoringEnabled" type="checkbox" class="mr-2 rounded border-gray-300 text-blue-600 focus:ring-blue-500">
            <span class="text-sm text-gray-700">فعال‌سازی سیستم خودکار</span>
          </label>
          <button class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition-colors" @click="saveSettings">
            ذخیره تنظیمات
          </button>
        </div>
      </div>
    </div>

    <div class="px-4 py-4">
      <!-- فعالیت‌های پایه -->
      <div class="mb-8">
        <h4 class="text-md font-medium text-gray-900 mb-4">فعالیت‌های پایه</h4>
        <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
          <div class="space-y-4">
            <div class="flex items-center justify-between px-4 py-4 bg-gray-50 rounded-lg">
              <div>
                <h5 class="font-medium text-gray-900">ورود به سیستم</h5>
                <p class="text-sm text-gray-500">امتیاز برای هر ورود روزانه</p>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <input v-model="scoringSettings.loginScore" type="number" min="0" max="10" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
                <span class="text-sm text-gray-600">امتیاز</span>
              </div>
            </div>

            <div class="flex items-center justify-between px-4 py-4 bg-gray-50 rounded-lg">
              <div>
                <h5 class="font-medium text-gray-900">مدت زمان فعالیت</h5>
                <p class="text-sm text-gray-500">امتیاز برای هر ساعت فعالیت</p>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <input v-model="scoringSettings.activityTimeScore" type="number" min="0" max="5" step="0.1" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
                <span class="text-sm text-gray-600">امتیاز</span>
              </div>
            </div>

            <div class="flex items-center justify-between px-4 py-4 bg-gray-50 rounded-lg">
              <div>
                <h5 class="font-medium text-gray-900">تکمیل پروفایل</h5>
                <p class="text-sm text-gray-500">امتیاز برای تکمیل 100% پروفایل</p>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <input v-model="scoringSettings.profileCompletionScore" type="number" min="0" max="50" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
                <span class="text-sm text-gray-600">امتیاز</span>
              </div>
            </div>
          </div>

          <div class="space-y-4">
            <div class="flex items-center justify-between px-4 py-4 bg-gray-50 rounded-lg">
              <div>
                <h5 class="font-medium text-gray-900">خرید محصول</h5>
                <p class="text-sm text-gray-500">امتیاز برای هر خرید موفق</p>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <input v-model="scoringSettings.purchaseScore" type="number" min="0" max="50" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
                <span class="text-sm text-gray-600">امتیاز</span>
              </div>
            </div>

            <div class="flex items-center justify-between px-4 py-4 bg-gray-50 rounded-lg">
              <div>
                <h5 class="font-medium text-gray-900">ارجاع کاربر جدید</h5>
                <p class="text-sm text-gray-500">امتیاز برای هر ارجاع موفق</p>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <input v-model="scoringSettings.referralScore" type="number" min="0" max="100" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
                <span class="text-sm text-gray-600">امتیاز</span>
              </div>
            </div>

            <div class="flex items-center justify-between px-4 py-4 bg-gray-50 rounded-lg">
              <div>
                <h5 class="font-medium text-gray-900">مشارکت در نظرسنجی</h5>
                <p class="text-sm text-gray-500">امتیاز برای هر مشارکت</p>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <input v-model="scoringSettings.surveyScore" type="number" min="0" max="20" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
                <span class="text-sm text-gray-600">امتیاز</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- فعالیت‌های تعاملی -->
      <div class="mb-8">
        <h4 class="text-md font-medium text-gray-900 mb-4">فعالیت‌های تعاملی</h4>
        <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
          <div class="space-y-4">
            <div class="flex items-center justify-between px-4 py-4 bg-blue-50 rounded-lg">
              <div>
                <h5 class="font-medium text-gray-900">نظرگذاری</h5>
                <p class="text-sm text-gray-500">امتیاز برای هر نظر مفید</p>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <input v-model="scoringSettings.reviewScore" type="number" min="0" max="20" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
                <span class="text-sm text-gray-600">امتیاز</span>
              </div>
            </div>

            <div class="flex items-center justify-between px-4 py-4 bg-blue-50 rounded-lg">
              <div>
                <h5 class="font-medium text-gray-900">امتیازدهی به محصولات</h5>
                <p class="text-sm text-gray-500">امتیاز برای هر امتیازدهی</p>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <input v-model="scoringSettings.productRatingScore" type="number" min="0" max="10" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
                <span class="text-sm text-gray-600">امتیاز</span>
              </div>
            </div>

            <div class="flex items-center justify-between px-4 py-4 bg-blue-50 rounded-lg">
              <div>
                <h5 class="font-medium text-gray-900">پاسخ به سوالات</h5>
                <p class="text-sm text-gray-500">امتیاز برای هر پاسخ مفید</p>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <input v-model="scoringSettings.answerScore" type="number" min="0" max="15" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
                <span class="text-sm text-gray-600">امتیاز</span>
              </div>
            </div>
          </div>

          <div class="space-y-4">
            <div class="flex items-center justify-between px-4 py-4 bg-blue-50 rounded-lg">
              <div>
                <h5 class="font-medium text-gray-900">مشارکت در بحث‌ها</h5>
                <p class="text-sm text-gray-500">امتیاز برای هر مشارکت</p>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <input v-model="scoringSettings.discussionScore" type="number" min="0" max="10" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
                <span class="text-sm text-gray-600">امتیاز</span>
              </div>
            </div>

            <div class="flex items-center justify-between px-4 py-4 bg-blue-50 rounded-lg">
              <div>
                <h5 class="font-medium text-gray-900">بازخورد مثبت از دیگران</h5>
                <p class="text-sm text-gray-500">امتیاز برای هر بازخورد مثبت</p>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <input v-model="scoringSettings.positiveFeedbackScore" type="number" min="0" max="15" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
                <span class="text-sm text-gray-600">امتیاز</span>
              </div>
            </div>

            <div class="flex items-center justify-between px-4 py-4 bg-blue-50 rounded-lg">
              <div>
                <h5 class="font-medium text-gray-900">اشتراک‌گذاری محتوا</h5>
                <p class="text-sm text-gray-500">امتیاز برای هر اشتراک‌گذاری</p>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <input v-model="scoringSettings.shareScore" type="number" min="0" max="8" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
                <span class="text-sm text-gray-600">امتیاز</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- امتیازات منفی -->
      <div class="mb-8">
        <h4 class="text-md font-medium text-gray-900 mb-4">امتیازات منفی</h4>
        <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
          <div class="space-y-4">
            <div class="flex items-center justify-between px-4 py-4 bg-red-50 rounded-lg">
              <div>
                <h5 class="font-medium text-gray-900">رفتار نامناسب</h5>
                <p class="text-sm text-gray-500">امتیاز منفی برای هر گزارش</p>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <input v-model="scoringSettings.inappropriateBehaviorPenalty" type="number" min="0" max="50" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
                <span class="text-sm text-gray-600">امتیاز</span>
              </div>
            </div>

            <div class="flex items-center justify-between px-4 py-4 bg-red-50 rounded-lg">
              <div>
                <h5 class="font-medium text-gray-900">نظرات نامناسب</h5>
                <p class="text-sm text-gray-500">امتیاز منفی برای هر نظر نامناسب</p>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <input v-model="scoringSettings.inappropriateReviewPenalty" type="number" min="0" max="25" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
                <span class="text-sm text-gray-600">امتیاز</span>
              </div>
            </div>
          </div>

          <div class="space-y-4">
            <div class="flex items-center justify-between px-4 py-4 bg-red-50 rounded-lg">
              <div>
                <h5 class="font-medium text-gray-900">عدم فعالیت</h5>
                <p class="text-sm text-gray-500">امتیاز منفی برای هر ماه عدم فعالیت</p>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <input v-model="scoringSettings.inactivityPenalty" type="number" min="0" max="10" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
                <span class="text-sm text-gray-600">امتیاز</span>
              </div>
            </div>

            <div class="flex items-center justify-between px-4 py-4 bg-red-50 rounded-lg">
              <div>
                <h5 class="font-medium text-gray-900">تخلف از قوانین</h5>
                <p class="text-sm text-gray-500">امتیاز منفی برای هر تخلف</p>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <input v-model="scoringSettings.violationPenalty" type="number" min="0" max="100" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
                <span class="text-sm text-gray-600">امتیاز</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- محدودیت‌های روزانه -->
      <div class="mb-8">
        <h4 class="text-md font-medium text-gray-900 mb-4">محدودیت‌های روزانه</h4>
        <div class="grid grid-cols-1 md:grid-cols-3 gapx-4 py-4">
          <div class="flex items-center justify-between px-4 py-4 bg-yellow-50 rounded-lg">
            <div>
              <h5 class="font-medium text-gray-900">حداکثر امتیاز روزانه</h5>
              <p class="text-sm text-gray-500">محدودیت امتیاز کسب شده در روز</p>
            </div>
            <div class="flex items-center space-x-2 space-x-reverse">
              <input v-model="scoringSettings.dailyScoreLimit" type="number" min="0" max="500" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
              <span class="text-sm text-gray-600">امتیاز</span>
            </div>
          </div>

          <div class="flex items-center justify-between px-4 py-4 bg-yellow-50 rounded-lg">
            <div>
              <h5 class="font-medium text-gray-900">حداکثر نظر روزانه</h5>
              <p class="text-sm text-gray-500">محدودیت تعداد نظرات در روز</p>
            </div>
            <div class="flex items-center space-x-2 space-x-reverse">
              <input v-model="scoringSettings.dailyReviewLimit" type="number" min="0" max="50" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
              <span class="text-sm text-gray-600">عدد</span>
            </div>
          </div>

          <div class="flex items-center justify-between px-4 py-4 bg-yellow-50 rounded-lg">
            <div>
              <h5 class="font-medium text-gray-900">حداکثر ارجاع روزانه</h5>
              <p class="text-sm text-gray-500">محدودیت تعداد ارجاعات در روز</p>
            </div>
            <div class="flex items-center space-x-2 space-x-reverse">
              <input v-model="scoringSettings.dailyReferralLimit" type="number" min="0" max="20" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
              <span class="text-sm text-gray-600">عدد</span>
            </div>
          </div>
        </div>
      </div>

      <!-- پیش‌نمایش فرمول -->
      <div class="bg-gray-50 rounded-lg px-4 py-4">
        <h4 class="text-md font-medium text-gray-900 mb-4">فرمول محاسبه امتیاز نهایی</h4>
        <div class="bg-white rounded-lg px-4 py-4 border">
          <p class="text-sm text-gray-700 font-mono">
            امتیاز نهایی = ({{ scoringSettings.loginScore }} × تعداد ورود) + 
            ({{ scoringSettings.activityTimeScore }} × مدت فعالیت) + 
            ({{ scoringSettings.reviewScore }} × کیفیت نظرات) + 
            ({{ scoringSettings.referralScore }} × تعداد ارجاعات) + 
            ({{ scoringSettings.purchaseScore }} × خریدها) + 
            ({{ scoringSettings.positiveFeedbackScore }} × بازخورد مثبت) - 
            ({{ scoringSettings.inappropriateBehaviorPenalty }} × امتیاز منفی)
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';

interface ScoringSettings {
  loginScore: number;
  activityTimeScore: number;
  profileCompletionScore: number;
  purchaseScore: number;
  referralScore: number;
  surveyScore: number;
  reviewScore: number;
  productRatingScore: number;
  answerScore: number;
  discussionScore: number;
  positiveFeedbackScore: number;
  shareScore: number;
  inappropriateBehaviorPenalty: number;
  inappropriateReviewPenalty: number;
  inactivityPenalty: number;
  violationPenalty: number;
  dailyScoreLimit: number;
  dailyReviewLimit: number;
  dailyReferralLimit: number;
}

interface AutoScoringProps {
  isEnabled?: boolean;
  scoringRules?: Partial<ScoringSettings>;
}

// Props definition
const props = defineProps<{
  initialSettings?: AutoScoringProps
}>()

// Emits
const emit = defineEmits<{
  saveSettings: [settings: Record<string, unknown>]
}>()

// Reactive data
const isAutoScoringEnabled = ref(true)

const scoringSettings = ref({
  // فعالیت‌های پایه
  loginScore: 1,
  activityTimeScore: 0.5,
  profileCompletionScore: 10,
  purchaseScore: 10,
  referralScore: 20,
  surveyScore: 3,

  // فعالیت‌های تعاملی
  reviewScore: 5,
  productRatingScore: 2,
  answerScore: 3,
  discussionScore: 2,
  positiveFeedbackScore: 5,
  shareScore: 2,

  // امتیازات منفی
  inappropriateBehaviorPenalty: 10,
  inappropriateReviewPenalty: 5,
  inactivityPenalty: 1,
  violationPenalty: 50,

  // محدودیت‌های روزانه
  dailyScoreLimit: 100,
  dailyReviewLimit: 10,
  dailyReferralLimit: 5
})

// Methods
const saveSettings = () => {
  const settings = {
    isEnabled: isAutoScoringEnabled.value,
    scoringRules: scoringSettings.value
  }
  emit('saveSettings', settings)
}

// Lifecycle
onMounted(() => {
  if (props.initialSettings) {
    isAutoScoringEnabled.value = props.initialSettings.isEnabled || true
    scoringSettings.value = { ...scoringSettings.value, ...props.initialSettings.scoringRules }
  }
})
</script> 

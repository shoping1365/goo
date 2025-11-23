<template>
  <div v-if="hasAccess">
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="border-b border-gray-200 pb-4">
      <h3 class="text-lg font-medium text-gray-900">بهینه‌سازی و یادگیری مداوم</h3>
      <p class="mt-1 text-sm text-gray-600">یادگیری مداوم از داده‌های جدید و بهینه‌سازی خودکار الگوریتم‌ها</p>
    </div>

    <!-- یادگیری آنلاین -->
    <div class="bg-gray-50 rounded-lg p-6">
      <div class="flex items-center justify-between mb-4">
        <h4 class="text-md font-medium text-gray-900">یادگیری آنلاین</h4>
        <div class="flex items-center space-x-2 space-x-reverse">
          <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800">
            فعال
          </span>
          <button class="text-blue-600 hover:text-blue-800 text-sm">تنظیمات</button>
        </div>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div class="bg-white rounded-lg p-6 border border-gray-200 text-center">
          <div class="text-2xl font-bold text-blue-600 mb-1">۲۴/۷</div>
          <div class="text-sm text-gray-600">یادگیری مداوم</div>
        </div>
        
        <div class="bg-white rounded-lg p-6 border border-gray-200 text-center">
          <div class="text-2xl font-bold text-green-600 mb-1">۱۲,۴۵۶</div>
          <div class="text-sm text-gray-600">نمونه جدید امروز</div>
        </div>
        
        <div class="bg-white rounded-lg p-6 border border-gray-200 text-center">
          <div class="text-2xl font-bold text-purple-600 mb-1">۹۸.۲%</div>
          <div class="text-sm text-gray-600">دقت یادگیری</div>
        </div>
      </div>
      
      <div class="mt-4 bg-white rounded-lg p-6 border border-gray-200">
        <h5 class="text-sm font-medium text-gray-900 mb-3">آخرین به‌روزرسانی‌ها</h5>
        <div class="space-y-2">
          <div class="flex items-center justify-between text-sm">
            <span class="text-gray-700">مدل Collaborative Filtering</span>
            <span class="text-green-600 font-medium">+۲.۳% بهبود</span>
          </div>
          <div class="flex items-center justify-between text-sm">
            <span class="text-gray-700">الگوریتم Content-Based</span>
            <span class="text-green-600 font-medium">+۱.۸% بهبود</span>
          </div>
          <div class="flex items-center justify-between text-sm">
            <span class="text-gray-700">مدل Deep Learning</span>
            <span class="text-green-600 font-medium">+۳.۱% بهبود</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Multi-Armed Bandit -->
    <div class="bg-gray-50 rounded-lg p-6">
      <div class="flex items-center justify-between mb-4">
        <h4 class="text-md font-medium text-gray-900">Multi-Armed Bandit</h4>
        <button class="text-blue-600 hover:text-blue-800 text-sm">تنظیمات بهینه‌سازی</button>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="bg-white rounded-lg p-6 border border-gray-200">
          <h5 class="text-sm font-medium text-gray-900 mb-3">عملکرد الگوریتم‌ها</h5>
          <div class="space-y-3">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">Collaborative Filtering</span>
              <div class="flex items-center">
                <div class="w-16 bg-gray-200 rounded-full h-2 ml-2">
                  <div class="bg-blue-600 h-2 rounded-full" style="width: 87%"></div>
                </div>
                <span class="text-sm font-medium">۸۷%</span>
              </div>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">Content-Based</span>
              <div class="flex items-center">
                <div class="w-16 bg-gray-200 rounded-full h-2 ml-2">
                  <div class="bg-green-600 h-2 rounded-full" style="width: 82%"></div>
                </div>
                <span class="text-sm font-medium">۸۲%</span>
              </div>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">Hybrid Approach</span>
              <div class="flex items-center">
                <div class="w-16 bg-gray-200 rounded-full h-2 ml-2">
                  <div class="bg-purple-600 h-2 rounded-full" style="width: 91%"></div>
                </div>
                <span class="text-sm font-medium">۹۱%</span>
              </div>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">Deep Learning</span>
              <div class="flex items-center">
                <div class="w-16 bg-gray-200 rounded-full h-2 ml-2">
                  <div class="bg-yellow-600 h-2 rounded-full" style="width: 89%"></div>
                </div>
                <span class="text-sm font-medium">۸۹%</span>
              </div>
            </div>
          </div>
        </div>
        
        <div class="bg-white rounded-lg p-6 border border-gray-200">
          <h5 class="text-sm font-medium text-gray-900 mb-3">توزیع انتخاب الگوریتم</h5>
          <div class="space-y-3">
            <div class="flex items-center justify-between">
              <div class="flex items-center">
                <div class="w-3 h-3 bg-blue-500 rounded-full ml-2"></div>
                <span class="text-sm text-gray-700">Collaborative</span>
              </div>
              <span class="text-sm font-medium">۳۵.۲%</span>
            </div>
            
            <div class="flex items-center justify-between">
              <div class="flex items-center">
                <div class="w-3 h-3 bg-green-500 rounded-full ml-2"></div>
                <span class="text-sm text-gray-700">Content-Based</span>
              </div>
              <span class="text-sm font-medium">۲۸.۷%</span>
            </div>
            
            <div class="flex items-center justify-between">
              <div class="flex items-center">
                <div class="w-3 h-3 bg-purple-500 rounded-full ml-2"></div>
                <span class="text-sm text-gray-700">Hybrid</span>
              </div>
              <span class="text-sm font-medium">۲۴.۳%</span>
            </div>
            
            <div class="flex items-center justify-between">
              <div class="flex items-center">
                <div class="w-3 h-3 bg-yellow-500 rounded-full ml-2"></div>
                <span class="text-sm text-gray-700">Deep Learning</span>
              </div>
              <span class="text-sm font-medium">۱۱.۸%</span>
            </div>
          </div>
        </div>
      </div>
    </div>



    <!-- بهینه‌سازی Hyperparameter -->
    <div class="bg-gray-50 rounded-lg p-6">
      <div class="flex items-center justify-between mb-4">
        <h4 class="text-md font-medium text-gray-900">بهینه‌سازی Hyperparameter</h4>
        <button class="text-blue-600 hover:text-blue-800 text-sm">شروع بهینه‌سازی</button>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="bg-white rounded-lg p-6 border border-gray-200">
          <h5 class="text-sm font-medium text-gray-900 mb-3">پارامترهای بهینه</h5>
          <div class="space-y-3">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">Learning Rate</span>
              <span class="text-sm font-medium">۰.۰۰۱</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">Batch Size</span>
              <span class="text-sm font-medium">۶۴</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">Epochs</span>
              <span class="text-sm font-medium">۱۰۰</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">Dropout Rate</span>
              <span class="text-sm font-medium">۰.۲</span>
            </div>
          </div>
        </div>
        
        <div class="bg-white rounded-lg p-6 border border-gray-200">
          <h5 class="text-sm font-medium text-gray-900 mb-3">عملکرد بهینه‌سازی</h5>
          <div class="space-y-3">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">بهترین دقت</span>
              <span class="text-sm font-medium text-green-600">۹۴.۸%</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">زمان بهینه‌سازی</span>
              <span class="text-sm font-medium">۲ ساعت ۳۴ دقیقه</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">تعداد آزمایش</span>
              <span class="text-sm font-medium">۱۵۶</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-700">بهبود نسبت به قبل</span>
              <span class="text-sm font-medium text-green-600">+۶.۲%</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Ensemble Methods -->
    <div class="bg-gray-50 rounded-lg p-6">
      <div class="flex items-center justify-between mb-4">
        <h4 class="text-md font-medium text-gray-900">Ensemble Methods</h4>
        <button class="text-blue-600 hover:text-blue-800 text-sm">تنظیمات Ensemble</button>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div class="bg-white rounded-lg p-6 border border-gray-200 text-center">
          <div class="text-2xl font-bold text-blue-600 mb-1">۵</div>
          <div class="text-sm text-gray-600">الگوریتم ترکیب شده</div>
        </div>
        
        <div class="bg-white rounded-lg p-6 border border-gray-200 text-center">
          <div class="text-2xl font-bold text-green-600 mb-1">۹۶.۳%</div>
          <div class="text-sm text-gray-600">دقت Ensemble</div>
        </div>
        
        <div class="bg-white rounded-lg p-6 border border-gray-200 text-center">
          <div class="text-2xl font-bold text-purple-600 mb-1">+۴.۱%</div>
          <div class="text-sm text-gray-600">بهبود نسبت به بهترین</div>
        </div>
      </div>
      
      <div class="mt-4 bg-white rounded-lg p-6 border border-gray-200">
        <h5 class="text-sm font-medium text-gray-900 mb-3">وزن‌دهی الگوریتم‌ها</h5>
        <div class="space-y-3">
          <div class="flex items-center justify-between">
            <span class="text-sm text-gray-700">Collaborative Filtering</span>
            <div class="flex items-center">
              <div class="w-20 bg-gray-200 rounded-full h-2 ml-2">
                <div class="bg-blue-600 h-2 rounded-full" style="width: 30%"></div>
              </div>
              <span class="text-sm font-medium">۳۰%</span>
            </div>
          </div>
          
          <div class="flex items-center justify-between">
            <span class="text-sm text-gray-700">Content-Based</span>
            <div class="flex items-center">
              <div class="w-20 bg-gray-200 rounded-full h-2 ml-2">
                <div class="bg-green-600 h-2 rounded-full" style="width: 25%"></div>
              </div>
              <span class="text-sm font-medium">۲۵%</span>
            </div>
          </div>
          
          <div class="flex items-center justify-between">
            <span class="text-sm text-gray-700">Deep Learning</span>
            <div class="flex items-center">
              <div class="w-20 bg-gray-200 rounded-full h-2 ml-2">
                <div class="bg-purple-600 h-2 rounded-full" style="width: 25%"></div>
              </div>
              <span class="text-sm font-medium">۲۵%</span>
            </div>
          </div>
          
          <div class="flex items-center justify-between">
            <span class="text-sm text-gray-700">Matrix Factorization</span>
            <div class="flex items-center">
              <div class="w-20 bg-gray-200 rounded-full h-2 ml-2">
                <div class="bg-yellow-600 h-2 rounded-full" style="width: 20%"></div>
              </div>
              <span class="text-sm font-medium">۲۰%</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات پیشرفته -->
    <div class="bg-gray-50 rounded-lg p-6">
      <h4 class="text-md font-medium text-gray-900 mb-4">تنظیمات پیشرفته</h4>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="space-y-4">
          <div class="flex items-center justify-between">
            <span class="text-sm font-medium text-gray-900">یادگیری آنلاین</span>
            <button class="relative inline-flex h-6 w-11 items-center rounded-full bg-blue-600">
              <span class="inline-block h-4 w-4 transform rounded-full bg-white transition translate-x-6"></span>
            </button>
          </div>
          
          <div class="flex items-center justify-between">
            <span class="text-sm font-medium text-gray-900">Multi-Armed Bandit</span>
            <button class="relative inline-flex h-6 w-11 items-center rounded-full bg-blue-600">
              <span class="inline-block h-4 w-4 transform rounded-full bg-white transition translate-x-6"></span>
            </button>
          </div>
          

        </div>
        
        <div class="space-y-4">
          <div class="flex items-center justify-between">
            <span class="text-sm font-medium text-gray-900">بهینه‌سازی Hyperparameter</span>
            <button class="relative inline-flex h-6 w-11 items-center rounded-full bg-blue-600">
              <span class="inline-block h-4 w-4 transform rounded-full bg-white transition translate-x-6"></span>
            </button>
          </div>
          
          <div class="flex items-center justify-between">
            <span class="text-sm font-medium text-gray-900">Ensemble Methods</span>
            <button class="relative inline-flex h-6 w-11 items-center rounded-full bg-blue-600">
              <span class="inline-block h-4 w-4 transform rounded-full bg-white transition translate-x-6"></span>
            </button>
          </div>
          
          <div class="flex items-center justify-between">
            <span class="text-sm font-medium text-gray-900">یادگیری مداوم</span>
            <button class="relative inline-flex h-6 w-11 items-center rounded-full bg-blue-600">
              <span class="inline-block h-4 w-4 transform rounded-full bg-white transition translate-x-6"></span>
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
  </div>
  <div v-else class="text-center p-8 text-gray-500">
    <p>شما مجوز دسترسی به این بخش را ندارید.</p>
  </div>
</template>

<script setup lang="ts">
import { useAuth } from '@/composables/useAuth';
import { computed, onMounted, ref, watch } from 'vue';
import { useRouter } from 'vue-router';

const { user } = useAuth();
const router = useRouter();

const hasAccess = computed(() => {
  return ['admin', 'developer'].includes(user.value?.role || '');
});

watch(hasAccess, (newValue) => {
  if (!newValue) {
    router.push('/404');
  }
});

onMounted(() => {
  if (!hasAccess.value) {
    router.push('/404');
  }
});


// متغیرهای کامپوننت
const _onlineLearningStats = ref({
  continuousLearning: '24/7',
  newSamplesToday: 12456,
  learningAccuracy: 98.2
})

const _algorithmPerformance = ref([
  { name: 'Collaborative Filtering', performance: 87, color: 'blue' },
  { name: 'Content-Based', performance: 82, color: 'green' },
  { name: 'Hybrid Approach', performance: 91, color: 'purple' },
  { name: 'Deep Learning', performance: 89, color: 'yellow' }
])

const _algorithmDistribution = ref([
  { name: 'Collaborative', percentage: 35.2, color: 'blue' },
  { name: 'Content-Based', percentage: 28.7, color: 'green' },
  { name: 'Hybrid', percentage: 24.3, color: 'purple' },
  { name: 'Deep Learning', percentage: 11.8, color: 'yellow' }
])



const _optimalParameters = ref({
  learningRate: 0.001,
  batchSize: 64,
  epochs: 100,
  dropoutRate: 0.2
})

const _optimizationStats = ref({
  bestAccuracy: 94.8,
  optimizationTime: '2 ساعت 34 دقیقه',
  numberOfTrials: 156,
  improvement: 6.2
})

const _ensembleStats = ref({
  numberOfAlgorithms: 5,
  ensembleAccuracy: 96.3,
  improvementOverBest: 4.1
})

const _algorithmWeights = ref([
  { name: 'Collaborative Filtering', weight: 30, color: 'blue' },
  { name: 'Content-Based', weight: 25, color: 'green' },
  { name: 'Deep Learning', weight: 25, color: 'purple' },
  { name: 'Matrix Factorization', weight: 20, color: 'yellow' }
])

// توابع کامپوننت
</script> 

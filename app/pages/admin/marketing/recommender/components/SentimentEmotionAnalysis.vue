<template>
  <div class="space-y-8">
    <!-- هدر بخش -->
    <div class="flex justify-between items-center">
      <div>
        <h2 class="text-2xl font-bold text-gray-900">تحلیل احساسات و عواطف</h2>
        <p class="mt-1 text-sm text-gray-600">تحلیل احساسات نظرات، تشخیص عواطف و تحلیل روندها</p>
      </div>
      <button class="bg-pink-600 text-white px-4 py-2 rounded-lg hover:bg-pink-700 transition-colors">
        <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"></path>
        </svg>
        تحلیل جدید
      </button>
    </div>

    <!-- آمار تحلیل احساسات -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-6">
      <div class="bg-gradient-to-r from-pink-500 to-pink-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">نظرات مثبت</p>
            <p class="text-2xl font-bold">78.5%</p>
            <p class="text-xs opacity-75">از کل نظرات</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.828 14.828a4 4 0 01-5.656 0M9 10h1m4 0h1m-6 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-blue-500 to-blue-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">احساسات شناسایی</p>
            <p class="text-2xl font-bold">12,847</p>
            <p class="text-xs opacity-75">امروز</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-green-500 to-green-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">دقت تشخیص</p>
            <p class="text-2xl font-bold">94.2%</p>
            <p class="text-xs opacity-75">میانگین</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-purple-500 to-purple-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">روند مثبت</p>
            <p class="text-2xl font-bold">+12.3%</p>
            <p class="text-xs opacity-75">از ماه گذشته</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-yellow-500 to-yellow-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">برند احساسات</p>
            <p class="text-2xl font-bold">4.8/5</p>
            <p class="text-xs opacity-75">امتیاز کلی</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z"></path>
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- Sentiment Analysis -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-pink-100 rounded-lg">
              <svg class="w-6 h-6 text-pink-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.828 14.828a4 4 0 01-5.656 0M9 10h1m4 0h1m-6 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">Sentiment Analysis</h3>
          </div>
          <label class="relative inline-flex items-center cursor-pointer">
            <input type="checkbox" v-model="features.sentimentAnalysis.enabled" class="sr-only peer">
            <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-pink-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-pink-600"></div>
          </label>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع تحلیل</label>
            <select v-model="features.sentimentAnalysis.type" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-pink-500 focus:border-pink-500">
              <option value="binary">دوتایی (مثبت/منفی)</option>
              <option value="ternary">سه‌تایی (مثبت/خنثی/منفی)</option>
              <option value="fine_grained">دقیق (5 سطح)</option>
              <option value="continuous">پیوسته (0-1)</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">الگوریتم تحلیل</label>
            <select v-model="features.sentimentAnalysis.algorithm" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-pink-500 focus:border-pink-500">
              <option value="bert">BERT</option>
              <option value="lstm">LSTM</option>
              <option value="transformer">Transformer</option>
              <option value="rule_based">Rule-Based</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">آستانه اطمینان</label>
            <input type="range" v-model="features.sentimentAnalysis.confidenceThreshold" min="0.5" max="1" step="0.05" class="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer">
            <div class="flex justify-between text-xs text-gray-500 mt-1">
              <span>0.5</span>
              <span>{{ features.sentimentAnalysis.confidenceThreshold }}</span>
              <span>1.0</span>
            </div>
          </div>
          
          <div class="bg-gray-50 rounded-lg p-6">
            <h4 class="font-medium text-gray-900 mb-2">نتایج تحلیل</h4>
            <div class="space-y-2 text-sm">
              <div class="flex justify-between">
                <span>مثبت:</span>
                <span class="font-medium text-green-600">78.5%</span>
              </div>
              <div class="flex justify-between">
                <span>خنثی:</span>
                <span class="font-medium text-gray-600">15.2%</span>
              </div>
              <div class="flex justify-between">
                <span>منفی:</span>
                <span class="font-medium text-red-600">6.3%</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Emotion Detection -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-blue-100 rounded-lg">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">Emotion Detection</h3>
          </div>
          <label class="relative inline-flex items-center cursor-pointer">
            <input type="checkbox" v-model="features.emotionDetection.enabled" class="sr-only peer">
            <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
          </label>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">احساسات قابل تشخیص</label>
            <div class="space-y-2">
              <label class="flex items-center">
                <input type="checkbox" v-model="features.emotionDetection.emotions" value="joy" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                <span class="mr-2 text-sm text-gray-700">شادی</span>
              </label>
              <label class="flex items-center">
                <input type="checkbox" v-model="features.emotionDetection.emotions" value="sadness" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                <span class="mr-2 text-sm text-gray-700">غم</span>
              </label>
              <label class="flex items-center">
                <input type="checkbox" v-model="features.emotionDetection.emotions" value="anger" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                <span class="mr-2 text-sm text-gray-700">خشم</span>
              </label>
              <label class="flex items-center">
                <input type="checkbox" v-model="features.emotionDetection.emotions" value="fear" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                <span class="mr-2 text-sm text-gray-700">ترس</span>
              </label>
              <label class="flex items-center">
                <input type="checkbox" v-model="features.emotionDetection.emotions" value="surprise" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                <span class="mr-2 text-sm text-gray-700">تعجب</span>
              </label>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">دقت تشخیص احساسات</label>
            <div class="flex items-center space-x-4 space-x-reverse">
              <input type="range" v-model="features.emotionDetection.accuracy" min="0" max="100" class="flex-1 h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer">
              <span class="text-sm font-medium text-gray-900 w-12">{{ features.emotionDetection.accuracy }}%</span>
            </div>
          </div>
          
          <div class="bg-gray-50 rounded-lg p-6">
            <h4 class="font-medium text-gray-900 mb-2">احساسات غالب</h4>
            <div class="space-y-2 text-sm">
              <div class="flex justify-between">
                <span>شادی:</span>
                <span class="font-medium text-green-600">45.2%</span>
              </div>
              <div class="flex justify-between">
                <span>رضایت:</span>
                <span class="font-medium text-blue-600">28.7%</span>
              </div>
              <div class="flex justify-between">
                <span>تعجب:</span>
                <span class="font-medium text-yellow-600">15.3%</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Tone Analysis -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <div class="flex items-center">
          <div class="p-2 bg-green-100 rounded-lg">
            <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"></path>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-900 mr-3">Tone Analysis</h3>
        </div>
        <label class="relative inline-flex items-center cursor-pointer">
          <input type="checkbox" v-model="features.toneAnalysis.enabled" class="sr-only peer">
          <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-green-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-green-600"></div>
        </label>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">ویژگی‌های لحن</label>
            <div class="space-y-2">
              <label class="flex items-center">
                <input type="checkbox" v-model="features.toneAnalysis.characteristics" value="formality" class="rounded border-gray-300 text-green-600 focus:ring-green-500">
                <span class="mr-2 text-sm text-gray-700">رسمیت</span>
              </label>
              <label class="flex items-center">
                <input type="checkbox" v-model="features.toneAnalysis.characteristics" value="politeness" class="rounded border-gray-300 text-green-600 focus:ring-green-500">
                <span class="mr-2 text-sm text-gray-700">ادب</span>
              </label>
              <label class="flex items-center">
                <input type="checkbox" v-model="features.toneAnalysis.characteristics" value="enthusiasm" class="rounded border-gray-300 text-green-600 focus:ring-green-500">
                <span class="mr-2 text-sm text-gray-700">اشتیاق</span>
              </label>
              <label class="flex items-center">
                <input type="checkbox" v-model="features.toneAnalysis.characteristics" value="urgency" class="rounded border-gray-300 text-green-600 focus:ring-green-500">
                <span class="mr-2 text-sm text-gray-700">فوریت</span>
              </label>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">سطح تحلیل</label>
            <select v-model="features.toneAnalysis.level" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500">
              <option value="basic">پایه</option>
              <option value="intermediate">متوسط</option>
              <option value="advanced">پیشرفته</option>
              <option value="expert">متخصص</option>
            </select>
          </div>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">وزن‌های لحن</label>
            <div class="space-y-2">
              <div class="flex items-center justify-between">
                <span class="text-sm text-gray-700">رسمیت</span>
                <input type="range" v-model="features.toneAnalysis.weights.formality" min="0" max="100" class="w-24 h-2 bg-gray-200 rounded-lg">
                <span class="text-xs text-gray-500 w-8">{{ features.toneAnalysis.weights.formality }}%</span>
              </div>
              <div class="flex items-center justify-between">
                <span class="text-sm text-gray-700">ادب</span>
                <input type="range" v-model="features.toneAnalysis.weights.politeness" min="0" max="100" class="w-24 h-2 bg-gray-200 rounded-lg">
                <span class="text-xs text-gray-500 w-8">{{ features.toneAnalysis.weights.politeness }}%</span>
              </div>
              <div class="flex items-center justify-between">
                <span class="text-sm text-gray-700">اشتیاق</span>
                <input type="range" v-model="features.toneAnalysis.weights.enthusiasm" min="0" max="100" class="w-24 h-2 bg-gray-200 rounded-lg">
                <span class="text-xs text-gray-500 w-8">{{ features.toneAnalysis.weights.enthusiasm }}%</span>
              </div>
            </div>
          </div>
          
          <div class="bg-gray-50 rounded-lg p-6">
            <h4 class="font-medium text-gray-900 mb-2">لحن غالب</h4>
            <div class="text-center">
              <div class="text-2xl font-bold text-green-600 mb-1">دوستانه</div>
              <div class="text-sm text-gray-600">با سطح بالای اشتیاق</div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Brand Sentiment -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-yellow-100 rounded-lg">
              <svg class="w-6 h-6 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">Brand Sentiment</h3>
          </div>
          <label class="relative inline-flex items-center cursor-pointer">
            <input type="checkbox" v-model="features.brandSentiment.enabled" class="sr-only peer">
            <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-yellow-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-yellow-600"></div>
          </label>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">برندهای تحت نظارت</label>
            <div class="space-y-2 max-h-32 overflow-y-auto">
              <div v-for="brand in features.brandSentiment.brands" :key="brand.id" class="flex items-center justify-between p-2 bg-gray-50 rounded-lg">
                <span class="text-sm text-gray-700">{{ brand.name }}</span>
                <div class="flex items-center space-x-2 space-x-reverse">
                  <span class="text-xs px-2 py-1 rounded-full" :class="brand.sentimentClass">{{ brand.sentiment }}</span>
                  <button @click="removeBrand(brand.id)" class="text-red-600 hover:text-red-800">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                    </svg>
                  </button>
                </div>
              </div>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">افزودن برند جدید</label>
            <div class="flex space-x-2 space-x-reverse">
              <input type="text" v-model="newBrand" placeholder="نام برند" class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500">
              <button @click="addBrand" class="bg-yellow-600 text-white px-4 py-2 rounded-lg hover:bg-yellow-700 transition-colors">
                افزودن
              </button>
            </div>
          </div>
          
          <div class="bg-gray-50 rounded-lg p-6">
            <h4 class="font-medium text-gray-900 mb-2">امتیاز کلی برند</h4>
            <div class="text-center">
              <div class="text-3xl font-bold text-yellow-600 mb-2">4.8/5</div>
              <div class="text-sm text-gray-600">بر اساس 12,847 نظر</div>
            </div>
          </div>
        </div>
      </div>

      <!-- Trend Analysis -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-purple-100 rounded-lg">
              <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">Trend Analysis</h3>
          </div>
          <label class="relative inline-flex items-center cursor-pointer">
            <input type="checkbox" v-model="features.trendAnalysis.enabled" class="sr-only peer">
            <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-purple-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-purple-600"></div>
          </label>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">بازه زمانی تحلیل</label>
            <select v-model="features.trendAnalysis.timeframe" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500">
              <option value="daily">روزانه</option>
              <option value="weekly">هفتگی</option>
              <option value="monthly">ماهانه</option>
              <option value="quarterly">فصلی</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع روند</label>
            <div class="space-y-2">
              <label class="flex items-center">
                <input type="checkbox" v-model="features.trendAnalysis.trends" value="sentiment" class="rounded border-gray-300 text-purple-600 focus:ring-purple-500">
                <span class="mr-2 text-sm text-gray-700">احساسات</span>
              </label>
              <label class="flex items-center">
                <input type="checkbox" v-model="features.trendAnalysis.trends" value="emotions" class="rounded border-gray-300 text-purple-600 focus:ring-purple-500">
                <span class="mr-2 text-sm text-gray-700">عواطف</span>
              </label>
              <label class="flex items-center">
                <input type="checkbox" v-model="features.trendAnalysis.trends" value="topics" class="rounded border-gray-300 text-purple-600 focus:ring-purple-500">
                <span class="mr-2 text-sm text-gray-700">موضوعات</span>
              </label>
              <label class="flex items-center">
                <input type="checkbox" v-model="features.trendAnalysis.trends" value="keywords" class="rounded border-gray-300 text-purple-600 focus:ring-purple-500">
                <span class="mr-2 text-sm text-gray-700">کلمات کلیدی</span>
              </label>
            </div>
          </div>
          
          <div class="bg-gray-50 rounded-lg p-6">
            <h4 class="font-medium text-gray-900 mb-2">روندهای اخیر</h4>
            <div class="space-y-2 text-sm">
              <div class="flex justify-between">
                <span>افزایش احساسات مثبت:</span>
                <span class="font-medium text-green-600">+12.3%</span>
              </div>
              <div class="flex justify-between">
                <span>کاهش احساسات منفی:</span>
                <span class="font-medium text-red-600">-8.7%</span>
              </div>
              <div class="flex justify-between">
                <span>روند احساسات:</span>
                <span class="font-medium text-purple-600">صعودی</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- نمودار روند احساسات -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-lg font-semibold text-gray-900">نمودار روند احساسات</h3>
        <div class="flex space-x-2 space-x-reverse">
          <button class="px-3 py-1 text-xs bg-blue-100 text-blue-800 rounded-lg">7 روز</button>
          <button class="px-3 py-1 text-xs bg-gray-100 text-gray-600 rounded-lg">30 روز</button>
          <button class="px-3 py-1 text-xs bg-gray-100 text-gray-600 rounded-lg">90 روز</button>
        </div>
      </div>
      
      <div class="h-64 bg-gray-50 rounded-lg flex items-center justify-center">
        <div class="text-center">
          <svg class="w-16 h-16 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
          </svg>
          <p class="text-gray-500">نمودار روند احساسات</p>
          <p class="text-sm text-gray-400">در حال بارگذاری...</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// کامپوننت تحلیل احساسات و عواطف
import { ref, reactive } from 'vue'

// متغیرهای reactive برای تنظیمات
const features = reactive({
  sentimentAnalysis: {
    enabled: true,
    type: 'ternary',
    algorithm: 'bert',
    confidenceThreshold: 0.8
  },
  emotionDetection: {
    enabled: true,
    emotions: ['joy', 'sadness', 'anger', 'fear', 'surprise'],
    accuracy: 94
  },
  toneAnalysis: {
    enabled: true,
    characteristics: ['formality', 'politeness', 'enthusiasm'],
    level: 'advanced',
    weights: {
      formality: 60,
      politeness: 80,
      enthusiasm: 70
    }
  },
  brandSentiment: {
    enabled: true,
    brands: [
      { id: 1, name: 'برند اصلی', sentiment: 'مثبت', sentimentClass: 'bg-green-100 text-green-800' },
      { id: 2, name: 'برند فرعی', sentiment: 'خنثی', sentimentClass: 'bg-gray-100 text-gray-800' },
      { id: 3, name: 'برند جدید', sentiment: 'مثبت', sentimentClass: 'bg-green-100 text-green-800' }
    ]
  },
  trendAnalysis: {
    enabled: true,
    timeframe: 'weekly',
    trends: ['sentiment', 'emotions', 'topics']
  }
})

const newBrand = ref('')

// توابع مدیریت
const addBrand = () => {
  if (newBrand.value.trim()) {
    features.brandSentiment.brands.push({
      id: Date.now(),
      name: newBrand.value,
      sentiment: 'خنثی',
      sentimentClass: 'bg-gray-100 text-gray-800'
    })
    newBrand.value = ''
  }
}

const removeBrand = (id: number) => {
  const index = features.brandSentiment.brands.findIndex(brand => brand.id === id)
  if (index > -1) {
    features.brandSentiment.brands.splice(index, 1)
  }
}

const updateSettings = () => {
  console.log('تنظیمات تحلیل احساسات به‌روزرسانی شد')
}

const exportAnalysis = () => {
  console.log('تحلیل احساسات صادر شد')
}
</script> 

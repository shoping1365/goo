<template>
  <div class="space-y-8">
    <!-- هدر بخش -->
    <div class="flex justify-between items-center">
      <div>
        <h2 class="text-2xl font-bold text-gray-900">الگوریتم‌های یادگیری ماشین</h2>
        <p class="mt-1 text-sm text-gray-600">مدیریت و بهینه‌سازی الگوریتم‌های پیشنهاددهی</p>
      </div>
      <button class="bg-purple-600 text-white px-4 py-2 rounded-lg hover:bg-purple-700 transition-colors">
        <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"></path>
        </svg>
        آموزش مدل جدید
      </button>
    </div>

    <!-- آمار الگوریتم‌ها -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-6">
      <div class="bg-gradient-to-r from-blue-500 to-blue-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">Collaborative</p>
            <p class="text-2xl font-bold">87.3%</p>
            <p class="text-xs opacity-75">دقت</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-green-500 to-green-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">Content-Based</p>
            <p class="text-2xl font-bold">92.1%</p>
            <p class="text-xs opacity-75">دقت</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-purple-500 to-purple-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">Matrix Factorization</p>
            <p class="text-2xl font-bold">89.7%</p>
            <p class="text-xs opacity-75">دقت</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-yellow-500 to-yellow-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">Deep Learning</p>
            <p class="text-2xl font-bold">94.5%</p>
            <p class="text-xs opacity-75">دقت</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-red-500 to-red-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">Reinforcement</p>
            <p class="text-2xl font-bold">91.2%</p>
            <p class="text-xs opacity-75">دقت</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- الگوریتم‌های فعال -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <!-- Collaborative Filtering -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-blue-100 rounded-lg">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">Collaborative Filtering</h3>
          </div>
          <label class="relative inline-flex items-center cursor-pointer">
            <input v-model="algorithms.collaborative.enabled" type="checkbox" class="sr-only peer">
            <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
          </label>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع الگوریتم</label>
            <select v-model="algorithms.collaborative.type" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
              <option value="user_based">User-Based</option>
              <option value="item_based">Item-Based</option>
              <option value="hybrid">Hybrid</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تعداد همسایه‌ها</label>
            <input v-model="algorithms.collaborative.neighbors" type="number" min="5" max="100" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">وزن اهمیت</label>
            <div class="flex items-center space-x-4 space-x-reverse">
              <input v-model="algorithms.collaborative.weight" type="range" min="0" max="100" class="flex-1 h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer">
              <span class="text-sm font-medium text-gray-900 w-12">{{ algorithms.collaborative.weight }}%</span>
            </div>
          </div>
          
          <div class="flex justify-between text-sm text-gray-600">
            <span>دقت: 87.3%</span>
            <span>سرعت: 234ms</span>
          </div>
        </div>
      </div>

      <!-- Content-Based Filtering -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-green-100 rounded-lg">
              <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">Content-Based Filtering</h3>
          </div>
          <label class="relative inline-flex items-center cursor-pointer">
            <input v-model="algorithms.contentBased.enabled" type="checkbox" class="sr-only peer">
            <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-green-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-green-600"></div>
          </label>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">ویژگی‌های محصول</label>
            <div class="space-y-2">
              <label class="flex items-center">
                <input v-model="algorithms.contentBased.features" type="checkbox" value="category" class="rounded border-gray-300 text-green-600 focus:ring-green-500">
                <span class="mr-2 text-sm text-gray-700">دسته‌بندی</span>
              </label>
              <label class="flex items-center">
                <input v-model="algorithms.contentBased.features" type="checkbox" value="brand" class="rounded border-gray-300 text-green-600 focus:ring-green-500">
                <span class="mr-2 text-sm text-gray-700">برند</span>
              </label>
              <label class="flex items-center">
                <input v-model="algorithms.contentBased.features" type="checkbox" value="price" class="rounded border-gray-300 text-green-600 focus:ring-green-500">
                <span class="mr-2 text-sm text-gray-700">قیمت</span>
              </label>
              <label class="flex items-center">
                <input v-model="algorithms.contentBased.features" type="checkbox" value="tags" class="rounded border-gray-300 text-green-600 focus:ring-green-500">
                <span class="mr-2 text-sm text-gray-700">برچسب‌ها</span>
              </label>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">الگوریتم تشابه</label>
            <select v-model="algorithms.contentBased.similarity" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500">
              <option value="cosine">Cosine Similarity</option>
              <option value="euclidean">Euclidean Distance</option>
              <option value="pearson">Pearson Correlation</option>
              <option value="jaccard">Jaccard Similarity</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">وزن اهمیت</label>
            <div class="flex items-center space-x-4 space-x-reverse">
              <input v-model="algorithms.contentBased.weight" type="range" min="0" max="100" class="flex-1 h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer">
              <span class="text-sm font-medium text-gray-900 w-12">{{ algorithms.contentBased.weight }}%</span>
            </div>
          </div>
          
          <div class="flex justify-between text-sm text-gray-600">
            <span>دقت: 92.1%</span>
            <span>سرعت: 156ms</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Matrix Factorization و Deep Learning -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <!-- Matrix Factorization -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-purple-100 rounded-lg">
              <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">Matrix Factorization</h3>
          </div>
          <label class="relative inline-flex items-center cursor-pointer">
            <input v-model="algorithms.matrixFactorization.enabled" type="checkbox" class="sr-only peer">
            <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-purple-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-purple-600"></div>
          </label>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تعداد فاکتورها</label>
            <input v-model="algorithms.matrixFactorization.factors" type="number" min="10" max="500" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نرخ یادگیری</label>
            <input v-model="algorithms.matrixFactorization.learningRate" type="number" min="0.001" max="1" step="0.001" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تعداد epoch</label>
            <input v-model="algorithms.matrixFactorization.epochs" type="number" min="10" max="1000" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">وزن اهمیت</label>
            <div class="flex items-center space-x-4 space-x-reverse">
              <input v-model="algorithms.matrixFactorization.weight" type="range" min="0" max="100" class="flex-1 h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer">
              <span class="text-sm font-medium text-gray-900 w-12">{{ algorithms.matrixFactorization.weight }}%</span>
            </div>
          </div>
          
          <div class="flex justify-between text-sm text-gray-600">
            <span>دقت: 89.7%</span>
            <span>سرعت: 189ms</span>
          </div>
        </div>
      </div>

      <!-- Deep Learning -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-yellow-100 rounded-lg">
              <svg class="w-6 h-6 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">Deep Learning</h3>
          </div>
          <label class="relative inline-flex items-center cursor-pointer">
            <input v-model="algorithms.deepLearning.enabled" type="checkbox" class="sr-only peer">
            <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-yellow-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-yellow-600"></div>
          </label>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">معماری شبکه</label>
            <select v-model="algorithms.deepLearning.architecture" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500">
              <option value="neural_cf">Neural Collaborative Filtering</option>
              <option value="wide_deep">Wide & Deep</option>
              <option value="deepfm">DeepFM</option>
              <option value="autorec">AutoRec</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تعداد لایه‌ها</label>
            <input v-model="algorithms.deepLearning.layers" type="number" min="2" max="10" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">اندازه لایه</label>
            <input v-model="algorithms.deepLearning.layerSize" type="number" min="32" max="512" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">وزن اهمیت</label>
            <div class="flex items-center space-x-4 space-x-reverse">
              <input v-model="algorithms.deepLearning.weight" type="range" min="0" max="100" class="flex-1 h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer">
              <span class="text-sm font-medium text-gray-900 w-12">{{ algorithms.deepLearning.weight }}%</span>
            </div>
          </div>
          
          <div class="flex justify-between text-sm text-gray-600">
            <span>دقت: 94.5%</span>
            <span>سرعت: 312ms</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Reinforcement Learning -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <div class="flex items-center">
          <div class="p-2 bg-red-100 rounded-lg">
            <svg class="w-6 h-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-900 mr-3">Reinforcement Learning</h3>
        </div>
        <label class="relative inline-flex items-center cursor-pointer">
          <input v-model="algorithms.reinforcement.enabled" type="checkbox" class="sr-only peer">
          <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-red-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-red-600"></div>
        </label>
      </div>
      
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع الگوریتم</label>
            <select v-model="algorithms.reinforcement.type" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500">
              <option value="q_learning">Q-Learning</option>
              <option value="policy_gradient">Policy Gradient</option>
              <option value="actor_critic">Actor-Critic</option>
              <option value="multi_armed_bandit">Multi-Armed Bandit</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نرخ یادگیری</label>
            <input v-model="algorithms.reinforcement.learningRate" type="number" min="0.001" max="1" step="0.001" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">فاکتور تخفیف</label>
            <input v-model="algorithms.reinforcement.discountFactor" type="number" min="0.1" max="1" step="0.1" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500">
          </div>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">استراتژی اکتشاف</label>
            <select v-model="algorithms.reinforcement.explorationStrategy" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500">
              <option value="epsilon_greedy">Epsilon-Greedy</option>
              <option value="softmax">Softmax</option>
              <option value="ucb">Upper Confidence Bound</option>
              <option value="thompson">Thompson Sampling</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">وزن اهمیت</label>
            <div class="flex items-center space-x-4 space-x-reverse">
              <input v-model="algorithms.reinforcement.weight" type="range" min="0" max="100" class="flex-1 h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer">
              <span class="text-sm font-medium text-gray-900 w-12">{{ algorithms.reinforcement.weight }}%</span>
            </div>
          </div>
          
          <div class="flex justify-between text-sm text-gray-600">
            <span>دقت: 91.2%</span>
            <span>سرعت: 278ms</span>
          </div>
        </div>
      </div>
    </div>

    <!-- مقایسه الگوریتم‌ها -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-semibold text-gray-900">مقایسه عملکرد الگوریتم‌ها</h3>
      </div>
      <div class="p-6">
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">الگوریتم</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">دقت</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">سرعت</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">حافظه</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="algorithm in algorithmComparison" :key="algorithm.id">
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="flex items-center">
                    <div class="w-8 h-8 rounded-lg mr-3 flex items-center justify-center" :class="algorithm.bgColor">
                      <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="algorithm.icon"></path>
                      </svg>
                    </div>
                    <div>
                      <div class="text-sm font-medium text-gray-900">{{ algorithm.name }}</div>
                      <div class="text-sm text-gray-500">{{ algorithm.description }}</div>
                    </div>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  {{ algorithm.accuracy }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  {{ algorithm.speed }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  {{ algorithm.memory }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full" :class="algorithm.statusClass">
                    {{ algorithm.status }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                  <button class="text-indigo-600 hover:text-indigo-900 mr-3">تنظیمات</button>
                  <button class="text-green-600 hover:text-green-900 mr-3">آموزش</button>
                  <button class="text-red-600 hover:text-red-900">حذف</button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
</script>

<script setup lang="ts">
import { ref } from 'vue'

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})
// تنظیمات الگوریتم‌ها
const algorithms = ref({
  collaborative: {
    enabled: true,
    type: 'user_based',
    neighbors: 20,
    weight: 30
  },
  contentBased: {
    enabled: true,
    features: ['category', 'brand', 'price'],
    similarity: 'cosine',
    weight: 25
  },
  matrixFactorization: {
    enabled: true,
    factors: 100,
    learningRate: 0.01,
    epochs: 100,
    weight: 20
  },
  deepLearning: {
    enabled: true,
    architecture: 'neural_cf',
    layers: 3,
    layerSize: 128,
    weight: 15
  },
  reinforcement: {
    enabled: false,
    type: 'q_learning',
    learningRate: 0.1,
    discountFactor: 0.9,
    explorationStrategy: 'epsilon_greedy',
    weight: 10
  }
})

// مقایسه الگوریتم‌ها
const algorithmComparison = ref([
  {
    id: 1,
    name: 'Collaborative Filtering',
    description: 'پیشنهاد بر اساس رفتار کاربران مشابه',
    accuracy: '87.3%',
    speed: '234ms',
    memory: '45MB',
    status: 'فعال',
    statusClass: 'bg-green-100 text-green-800',
    bgColor: 'bg-blue-500',
    icon: 'M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z'
  },
  {
    id: 2,
    name: 'Content-Based Filtering',
    description: 'پیشنهاد بر اساس ویژگی‌های محصول',
    accuracy: '92.1%',
    speed: '156ms',
    memory: '32MB',
    status: 'فعال',
    statusClass: 'bg-green-100 text-green-800',
    bgColor: 'bg-green-500',
    icon: ''
  },
  {
    id: 3,
    name: 'Matrix Factorization',
    description: 'تجزیه ماتریس برای پیشنهادات دقیق‌تر',
    accuracy: '89.7%',
    speed: '189ms',
    memory: '78MB',
    status: 'فعال',
    statusClass: 'bg-green-100 text-green-800',
    bgColor: 'bg-purple-500',
    icon: 'M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z'
  },
  {
    id: 4,
    name: 'Deep Learning',
    description: 'شبکه‌های عصبی برای تشخیص الگوهای پیچیده',
    accuracy: '94.5%',
    speed: '312ms',
    memory: '156MB',
    status: 'فعال',
    statusClass: 'bg-green-100 text-green-800',
    bgColor: 'bg-yellow-500',
    icon: 'M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z'
  },
  {
    id: 5,
    name: 'Reinforcement Learning',
    description: 'یادگیری از بازخورد کاربران',
    accuracy: '91.2%',
    speed: '278ms',
    memory: '89MB',
    status: 'غیرفعال',
    statusClass: 'bg-gray-100 text-gray-800',
    bgColor: 'bg-red-500',
    icon: 'M13 10V3L4 14h7v7l9-11h-7z'
  }
])
</script> 

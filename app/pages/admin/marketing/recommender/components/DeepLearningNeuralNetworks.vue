<template>
  <div class="space-y-8">
    <!-- هدر بخش -->
    <div class="flex justify-between items-center">
      <div>
        <h2 class="text-2xl font-bold text-gray-900">یادگیری عمیق و شبکه‌های عصبی</h2>
        <p class="mt-1 text-sm text-gray-600">مدیریت مدل‌های پیشرفته یادگیری عمیق</p>
      </div>
      <button class="bg-purple-600 text-white px-4 py-2 rounded-lg hover:bg-purple-700 transition-colors">
        <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"></path>
        </svg>
        آموزش مدل جدید
      </button>
    </div>

    <!-- آمار مدل‌ها -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-6">
      <div class="bg-gradient-to-r from-purple-500 to-purple-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">مدل‌های فعال</p>
            <p class="text-2xl font-bold">12</p>
            <p class="text-xs opacity-75">در حال اجرا</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-blue-500 to-blue-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">دقت متوسط</p>
            <p class="text-2xl font-bold">94.7%</p>
            <p class="text-xs opacity-75">میانگین کلی</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-green-500 to-green-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">زمان آموزش</p>
            <p class="text-2xl font-bold">2.3h</p>
            <p class="text-xs opacity-75">میانگین</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-yellow-500 to-yellow-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">GPU استفاده</p>
            <p class="text-2xl font-bold">78%</p>
            <p class="text-xs opacity-75">ظرفیت</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-red-500 to-red-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">مدل‌های در انتظار</p>
            <p class="text-2xl font-bold">3</p>
            <p class="text-xs opacity-75">آماده آموزش</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- Neural Collaborative Filtering -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-purple-100 rounded-lg">
              <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">Neural Collaborative Filtering</h3>
          </div>
          <label class="relative inline-flex items-center cursor-pointer">
            <input v-model="models.neuralCF.enabled" type="checkbox" class="sr-only peer">
            <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-purple-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-purple-600"></div>
          </label>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تعداد لایه‌های پنهان</label>
            <input v-model="models.neuralCF.hiddenLayers" type="number" min="1" max="10" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">اندازه لایه‌ها</label>
            <input v-model="models.neuralCF.layerSize" type="number" min="32" max="512" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نرخ یادگیری</label>
            <input v-model="models.neuralCF.learningRate" type="number" min="0.001" max="1" step="0.001" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Dropout Rate</label>
            <input v-model="models.neuralCF.dropoutRate" type="number" min="0" max="0.9" step="0.1" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500">
          </div>
          
          <div class="flex justify-between text-sm text-gray-600">
            <span>دقت: 96.2%</span>
            <span>سرعت: 145ms</span>
          </div>
        </div>
      </div>

      <!-- Autoencoders -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-blue-100 rounded-lg">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">Autoencoders</h3>
          </div>
          <label class="relative inline-flex items-center cursor-pointer">
            <input v-model="models.autoencoders.enabled" type="checkbox" class="sr-only peer">
            <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
          </label>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع Autoencoder</label>
            <select v-model="models.autoencoders.type" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
              <option value="vanilla">Vanilla Autoencoder</option>
              <option value="variational">Variational Autoencoder</option>
              <option value="denoising">Denoising Autoencoder</option>
              <option value="sparse">Sparse Autoencoder</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">ابعاد فشرده</label>
            <input v-model="models.autoencoders.compressionRatio" type="number" min="2" max="100" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Regularization</label>
            <select v-model="models.autoencoders.regularization" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
              <option value="l1">L1 Regularization</option>
              <option value="l2">L2 Regularization</option>
              <option value="dropout">Dropout</option>
              <option value="none">None</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Reconstruction Loss</label>
            <select v-model="models.autoencoders.lossFunction" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
              <option value="mse">Mean Squared Error</option>
              <option value="mae">Mean Absolute Error</option>
              <option value="binary_crossentropy">Binary Crossentropy</option>
            </select>
          </div>
          
          <div class="flex justify-between text-sm text-gray-600">
            <span>دقت: 93.8%</span>
            <span>سرعت: 89ms</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Recurrent Neural Networks -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-green-100 rounded-lg">
              <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">Recurrent Neural Networks</h3>
          </div>
          <label class="relative inline-flex items-center cursor-pointer">
            <input v-model="models.rnn.enabled" type="checkbox" class="sr-only peer">
            <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-green-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-green-600"></div>
          </label>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع RNN</label>
            <select v-model="models.rnn.type" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500">
              <option value="lstm">LSTM</option>
              <option value="gru">GRU</option>
              <option value="simple">Simple RNN</option>
              <option value="bidirectional">Bidirectional LSTM</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تعداد لایه‌های RNN</label>
            <input v-model="models.rnn.layers" type="number" min="1" max="5" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">اندازه لایه</label>
            <input v-model="models.rnn.hiddenSize" type="number" min="32" max="512" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Sequence Length</label>
            <input v-model="models.rnn.sequenceLength" type="number" min="5" max="100" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500">
          </div>
          
          <div class="flex justify-between text-sm text-gray-600">
            <span>دقت: 95.1%</span>
            <span>سرعت: 234ms</span>
          </div>
        </div>
      </div>

      <!-- Convolutional Neural Networks -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-yellow-100 rounded-lg">
              <svg class="w-6 h-6 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">Convolutional Neural Networks</h3>
          </div>
          <label class="relative inline-flex items-center cursor-pointer">
            <input v-model="models.cnn.enabled" type="checkbox" class="sr-only peer">
            <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-yellow-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-yellow-600"></div>
          </label>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">معماری CNN</label>
            <select v-model="models.cnn.architecture" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500">
              <option value="resnet">ResNet</option>
              <option value="vgg">VGG</option>
              <option value="inception">Inception</option>
              <option value="custom">Custom</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تعداد فیلترها</label>
            <input v-model="models.cnn.filters" type="number" min="16" max="512" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">اندازه Kernel</label>
            <input v-model="models.cnn.kernelSize" type="number" min="3" max="7" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Pooling</label>
            <select v-model="models.cnn.pooling" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500">
              <option value="max">Max Pooling</option>
              <option value="average">Average Pooling</option>
              <option value="global">Global Pooling</option>
            </select>
          </div>
          
          <div class="flex justify-between text-sm text-gray-600">
            <span>دقت: 97.3%</span>
            <span>سرعت: 189ms</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Transformer Models -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <div class="flex items-center">
          <div class="p-2 bg-red-100 rounded-lg">
            <svg class="w-6 h-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4"></path>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-900 mr-3">Transformer Models</h3>
        </div>
        <label class="relative inline-flex items-center cursor-pointer">
          <input v-model="models.transformer.enabled" type="checkbox" class="sr-only peer">
          <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-red-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-red-600"></div>
        </label>
      </div>
      
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع Transformer</label>
            <select v-model="models.transformer.type" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500">
              <option value="bert">BERT</option>
              <option value="gpt">GPT</option>
              <option value="t5">T5</option>
              <option value="custom">Custom Transformer</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تعداد لایه‌ها</label>
            <input v-model="models.transformer.layers" type="number" min="2" max="24" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">اندازه Embedding</label>
            <input v-model="models.transformer.embeddingSize" type="number" min="128" max="1024" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تعداد Head</label>
            <input v-model="models.transformer.numHeads" type="number" min="4" max="16" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500">
          </div>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Maximum Sequence Length</label>
            <input v-model="models.transformer.maxLength" type="number" min="64" max="512" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Dropout Rate</label>
            <input v-model="models.transformer.dropoutRate" type="number" min="0" max="0.9" step="0.1" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Positional Encoding</label>
            <select v-model="models.transformer.positionalEncoding" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500">
              <option value="sinusoidal">Sinusoidal</option>
              <option value="learned">Learned</option>
              <option value="relative">Relative</option>
            </select>
          </div>
          
          <div class="flex justify-between text-sm text-gray-600">
            <span>دقت: 98.1%</span>
            <span>سرعت: 312ms</span>
          </div>
        </div>
      </div>
    </div>

    <!-- مدیریت مدل‌ها -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-semibold text-gray-900">مدیریت مدل‌های یادگیری عمیق</h3>
      </div>
      <div class="p-6">
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مدل</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">دقت</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">آخرین به‌روزرسانی</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="model in deepLearningModels" :key="model.id">
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="flex items-center">
                    <div class="w-8 h-8 rounded-lg mr-3 flex items-center justify-center" :class="model.bgColor">
                      <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="model.icon"></path>
                      </svg>
                    </div>
                    <div>
                      <div class="text-sm font-medium text-gray-900">{{ model.name }}</div>
                      <div class="text-sm text-gray-500">{{ model.description }}</div>
                    </div>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  {{ model.type }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  {{ model.accuracy }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full" :class="model.statusClass">
                    {{ model.status }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  {{ model.lastUpdate }}
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

<script setup lang="ts">
import { ref } from 'vue';
// تنظیمات مدل‌ها
const models = ref({
  neuralCF: {
    enabled: true,
    hiddenLayers: 3,
    layerSize: 128,
    learningRate: 0.001,
    dropoutRate: 0.3
  },
  autoencoders: {
    enabled: true,
    type: 'variational',
    compressionRatio: 10,
    regularization: 'l2',
    lossFunction: 'mse'
  },
  rnn: {
    enabled: true,
    type: 'lstm',
    layers: 2,
    hiddenSize: 128,
    sequenceLength: 20
  },
  cnn: {
    enabled: true,
    architecture: 'resnet',
    filters: 64,
    kernelSize: 3,
    pooling: 'max'
  },
  transformer: {
    enabled: false,
    type: 'bert',
    layers: 6,
    embeddingSize: 256,
    numHeads: 8,
    maxLength: 128,
    dropoutRate: 0.1,
    positionalEncoding: 'sinusoidal'
  }
})

// مدل‌های یادگیری عمیق
const deepLearningModels = ref([
  {
    id: 1,
    name: 'Neural CF Model',
    description: 'مدل فیلترینگ مشارکتی عصبی',
    type: 'Neural CF',
    accuracy: '96.2%',
    status: 'فعال',
    statusClass: 'bg-green-100 text-green-800',
    lastUpdate: '2 روز پیش',
    bgColor: 'bg-purple-500',
    icon: 'M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z'
  },
  {
    id: 2,
    name: 'VAE Model',
    description: 'مدل Autoencoder متغیر',
    type: 'VAE',
    accuracy: '93.8%',
    status: 'فعال',
    statusClass: 'bg-green-100 text-green-800',
    lastUpdate: '1 هفته پیش',
    bgColor: 'bg-blue-500',
    icon: 'M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15'
  },
  {
    id: 3,
    name: 'LSTM Model',
    description: 'مدل LSTM برای تحلیل توالی',
    type: 'LSTM',
    accuracy: '95.1%',
    status: 'فعال',
    statusClass: 'bg-green-100 text-green-800',
    lastUpdate: '3 روز پیش',
    bgColor: 'bg-green-500',
    icon: 'M13 10V3L4 14h7v7l9-11h-7z'
  },
  {
    id: 4,
    name: 'ResNet Model',
    description: 'مدل CNN برای تحلیل تصاویر',
    type: 'CNN',
    accuracy: '97.3%',
    status: 'فعال',
    statusClass: 'bg-green-100 text-green-800',
    lastUpdate: '5 روز پیش',
    bgColor: 'bg-yellow-500',
    icon: 'M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z'
  },
  {
    id: 5,
    name: 'BERT Model',
    description: 'مدل Transformer برای پردازش متون',
    type: 'Transformer',
    accuracy: '98.1%',
    status: 'در انتظار',
    statusClass: 'bg-yellow-100 text-yellow-800',
    lastUpdate: '1 روز پیش',
    bgColor: 'bg-red-500',
    icon: 'M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4'
  }
])
</script> 

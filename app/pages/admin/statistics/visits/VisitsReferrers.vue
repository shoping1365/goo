<template>
  <div class="bg-white rounded-xl shadow-lg p-6">
    <h3 class="text-lg font-bold text-gray-900 mb-4">منابع ترافیک</h3>
    
    <div class="space-y-4">
      <div v-for="referrer in referrers" :key="referrer.source" class="flex items-center justify-between">
        <div class="flex items-center gap-3">
          <div class="w-10 h-10 rounded-lg flex items-center justify-center" :class="getSourceClass(referrer.source)">
            <component :is="getSourceIcon(referrer.source)" class="w-5 h-5 text-white" />
          </div>
          <div>
            <div class="font-semibold text-gray-900">{{ referrer.source }}</div>
            <div class="text-sm text-gray-500">{{ referrer.visits.toLocaleString('fa-IR') }} بازدید</div>
          </div>
        </div>
        <div class="flex items-center gap-3">
          <div class="w-24">
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div 
                class="h-2 rounded-full transition-all duration-500"
                :class="getProgressClass(referrer.source)"
                :style="{ width: referrer.percentage + '%' }"
              ></div>
            </div>
          </div>
          <div class="text-sm font-semibold text-gray-700 w-12 text-left">
            {{ referrer.percentage }}%
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { h } from 'vue';

interface Referrer {
  source: string
  visits: number
  percentage: number
}

defineProps<{ referrers: Referrer[] }>()

const getSourceIcon = (source: string) => {
  const icons = {
    'Google': () => h('svg', { fill: 'currentColor', viewBox: '0 0 24 24' }, [
      h('path', { d: 'M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z' }),
      h('path', { d: 'M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z' })
    ]),
    'مستقیم': () => h('svg', { fill: 'currentColor', viewBox: '0 0 20 20' }, [
      h('path', { 'fill-rule': 'evenodd', d: 'M3 3a1 1 0 000 2v8a2 2 0 002 2h2.586l-1.293 1.293a1 1 0 101.414 1.414L10 15.414l2.293 2.293a1 1 0 001.414-1.414L12.414 15H15a2 2 0 002-2V5a1 1 0 100-2H3zm11.707 4.707a1 1 0 00-1.414-1.414L10 9.586 8.707 8.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z', 'clip-rule': 'evenodd' })
    ]),
    'Instagram': () => h('svg', { fill: 'currentColor', viewBox: '0 0 24 24' }, [
      h('path', { d: 'M12 2.163c3.204 0 3.584.012 4.85.07 3.252.148 4.771 1.691 4.919 4.919.058 1.265.069 1.645.069 4.849 0 3.205-.012 3.584-.069 4.849-.149 3.225-1.664 4.771-4.919 4.919-1.266.058-1.644.07-4.85.07-3.204 0-3.584-.012-4.849-.07-3.26-.149-4.771-1.699-4.919-4.92-.058-1.265-.07-1.644-.07-4.849 0-3.204.013-3.583.07-4.849.149-3.227 1.664-4.771 4.919-4.919 1.266-.057 1.645-.069 4.849-.069zm0-2.163c-3.259 0-3.667.014-4.947.072-4.358.2-6.78 2.618-6.98 6.98-.059 1.281-.073 1.689-.073 4.948 0 3.259.014 3.668.072 4.948.2 4.358 2.618 6.78 6.98 6.98 1.281.058 1.689.072 4.948.072 3.259 0 3.668-.014 4.948-.072 4.354-.2 6.782-2.618 6.979-6.98.059-1.28.073-1.689.073-4.948 0-3.259-.014-3.667-.072-4.947-.196-4.354-2.617-6.78-6.979-6.98-1.281-.059-1.69-.073-4.949-.073zm0 5.838c-3.403 0-6.162 2.759-6.162 6.162s2.759 6.163 6.162 6.163 6.162-2.759 6.162-6.163c0-3.403-2.759-6.162-6.162-6.162zm0 10.162c-2.209 0-4-1.79-4-4 0-2.209 1.791-4 4-4s4 1.791 4 4c0 2.21-1.791 4-4 4zm6.406-11.845c-.796 0-1.441.645-1.441 1.44s.645 1.44 1.441 1.44c.795 0 1.439-.645 1.439-1.44s-.644-1.44-1.439-1.44z' })
    ])
  }
  return icons[source] || icons['مستقیم']
}

const getSourceClass = (source: string) => {
  const classes = {
    'Google': 'bg-gradient-to-br from-red-500 to-red-600',
    'مستقیم': 'bg-gradient-to-br from-gray-500 to-gray-600',
    'Instagram': 'bg-gradient-to-br from-pink-500 to-purple-600',
    'Telegram': 'bg-gradient-to-br from-blue-400 to-blue-500',
    'Facebook': 'bg-gradient-to-br from-blue-600 to-blue-700',
    'Twitter': 'bg-gradient-to-br from-blue-400 to-blue-500'
  }
  return classes[source] || 'bg-gray-500'
}

const getProgressClass = (source: string) => {
  const classes = {
    'Google': 'bg-gradient-to-r from-red-400 to-red-600',
    'مستقیم': 'bg-gradient-to-r from-gray-400 to-gray-600',
    'Instagram': 'bg-gradient-to-r from-pink-400 to-purple-600',
    'Telegram': 'bg-gradient-to-r from-blue-300 to-blue-500',
    'Facebook': 'bg-gradient-to-r from-blue-500 to-blue-700',
    'Twitter': 'bg-gradient-to-r from-blue-300 to-blue-500'
  }
  return classes[source] || 'bg-gray-400'
}
</script>

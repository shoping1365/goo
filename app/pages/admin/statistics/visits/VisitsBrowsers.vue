<template>
  <div class="bg-white rounded-xl shadow-lg p-6">
    <h3 class="text-lg font-bold text-gray-900 mb-4">آمار مرورگرها</h3>
    
    <div class="space-y-4">
      <div v-for="browser in browsers" :key="browser.browser" class="flex items-center justify-between">
        <div class="flex items-center gap-3">
          <div class="w-10 h-10 rounded-lg flex items-center justify-center" :class="getBrowserClass(browser.browser)">
            <component :is="getBrowserIcon(browser.browser)" class="w-5 h-5 text-white" />
          </div>
          <div>
            <div class="font-semibold text-gray-900">{{ browser.browser }}</div>
            <div class="text-sm text-gray-500">{{ browser.visits.toLocaleString('fa-IR') }} بازدید</div>
          </div>
        </div>
        <div class="flex items-center gap-3">
          <div class="w-24">
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div 
                class="h-2 rounded-full transition-all duration-500"
                :class="getProgressClass(browser.browser)"
                :style="{ width: browser.percentage + '%' }"
              ></div>
            </div>
          </div>
          <div class="text-sm font-semibold text-gray-700 w-12 text-left">
            {{ browser.percentage }}%
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
</script>

<script setup lang="ts">
import { h } from 'vue';

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

interface Browser {
  browser: string
  visits: number
  percentage: number
}

defineProps<{
  browsers: Browser[]
}>()

const getBrowserIcon = (browser: string) => {
  const icons = {
    'Chrome': () => h('svg', { fill: 'currentColor', viewBox: '0 0 24 24' }, [
      h('path', { d: 'M12 0C5.373 0 0 5.373 0 12s5.373 12 12 12 12-5.373 12-12S18.627 0 12 0zm4.894 14.842c-.847 1.467-2.42 2.445-4.2 2.445-1.78 0-3.353-.978-4.2-2.445l2.1-3.638h4.2l2.1 3.638z' })
    ]),
    'Safari': () => h('svg', { fill: 'currentColor', viewBox: '0 0 24 24' }, [
      h('path', { d: 'M12 24C5.373 24 0 18.627 0 12S5.373 0 12 0s12 5.373 12 12-5.373 12-12 12zm0-1c6.075 0 11-4.925 11-11S18.075 1 12 1 1 5.925 1 12s4.925 11 11 11z' }),
      h('path', { d: 'M12 2.5c-.553 0-1 .447-1 1s.447 1 1 1 1-.447 1-1-.447-1-1-1z' })
    ]),
    'Firefox': () => h('svg', { fill: 'currentColor', viewBox: '0 0 24 24' }, [
      h('path', { d: 'M21.834 10.634c-.297-3.155-2.395-5.82-5.245-6.756C15.83 3.569 15.025 3.4 14.204 3.4c-2.835 0-5.387 1.616-6.659 4.216-.416.851-.644 1.794-.644 2.784 0 3.588 2.912 6.5 6.5 6.5s6.5-2.912 6.5-6.5c0-.359-.029-.711-.086-1.055-.039-.233-.08-.46-.127-.68' })
    ]),
    'Edge': () => h('svg', { fill: 'currentColor', viewBox: '0 0 24 24' }, [
      h('path', { d: 'M21.818 12.35c-.03 5.486-4.476 9.93-9.962 9.93-3.436 0-6.594-1.746-8.427-4.66C1.815 15.063.955 11.772 1.136 8.5c.181-3.272 1.788-6.293 4.407-8.281C7.162-1.269 9.545-1.816 12.136 1.07c1.181 1.314 1.907 3.028 2.045 4.831.138 1.803-.233 3.605-1.046 5.157' })
    ]),
    'Opera': () => h('svg', { fill: 'currentColor', viewBox: '0 0 24 24' }, [
      h('path', { d: 'M12 24C5.373 24 0 18.627 0 12S5.373 0 12 0s12 5.373 12 12-5.373 12-12 12zm0-22C6.477 2 2 6.477 2 12s4.477 10 10 10 10-4.477 10-10S17.523 2 12 2z' })
    ])
  }
  return icons[browser] || icons['Chrome']
}

const getBrowserClass = (browser: string) => {
  const classes = {
    'Chrome': 'bg-gradient-to-br from-red-500 to-yellow-500',
    'Safari': 'bg-gradient-to-br from-blue-500 to-blue-600',
    'Firefox': 'bg-gradient-to-br from-orange-500 to-red-500',
    'Edge': 'bg-gradient-to-br from-blue-600 to-green-500',
    'Opera': 'bg-gradient-to-br from-red-600 to-pink-500'
  }
  return classes[browser] || 'bg-gray-500'
}

const getProgressClass = (browser: string) => {
  const classes = {
    'Chrome': 'bg-gradient-to-r from-red-400 to-yellow-500',
    'Safari': 'bg-gradient-to-r from-blue-400 to-blue-600',
    'Firefox': 'bg-gradient-to-r from-orange-400 to-red-500',
    'Edge': 'bg-gradient-to-r from-blue-500 to-green-500',
    'Opera': 'bg-gradient-to-r from-red-500 to-pink-500'
  }
  return classes[browser] || 'bg-gray-400'
}
</script>

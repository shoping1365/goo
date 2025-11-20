<template>
  <div class="bg-white rounded-xl shadow-lg p-6">
    <h3 class="text-lg font-bold text-gray-900 mb-4">آمار دستگاه‌ها</h3>
    
    <div class="space-y-4">
      <div v-for="device in devices" :key="device.device" class="flex items-center justify-between">
        <div class="flex items-center gap-3">
          <div class="w-10 h-10 rounded-lg flex items-center justify-center" :class="getDeviceClass(device.device)">
            <component :is="getDeviceIcon(device.device)" class="w-5 h-5 text-white" />
          </div>
          <div>
            <div class="font-semibold text-gray-900">{{ device.device }}</div>
            <div class="text-sm text-gray-500">{{ device.visits.toLocaleString('fa-IR') }} بازدید</div>
          </div>
        </div>
        <div class="flex items-center gap-3">
          <div class="w-24">
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div 
                class="h-2 rounded-full transition-all duration-500"
                :class="getProgressClass(device.device)"
                :style="{ width: device.percentage + '%' }"
              ></div>
            </div>
          </div>
          <div class="text-sm font-semibold text-gray-700 w-12 text-left">
            {{ device.percentage }}%
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { h } from 'vue';

interface Device {
  device: string
  visits: number
  percentage: number
}

defineProps<{
  devices: Device[]
}>()

const getDeviceIcon = (device: string) => {
  const icons = {
    'موبایل': () => h('svg', { fill: 'currentColor', viewBox: '0 0 20 20' }, [
      h('path', { 'fill-rule': 'evenodd', d: 'M7 2a2 2 0 00-2 2v12a2 2 0 002 2h6a2 2 0 002-2V4a2 2 0 00-2-2H7zM8 4h4v10H8V4z', 'clip-rule': 'evenodd' })
    ]),
    'دسکتاپ': () => h('svg', { fill: 'currentColor', viewBox: '0 0 20 20' }, [
      h('path', { 'fill-rule': 'evenodd', d: 'M3 5a2 2 0 012-2h10a2 2 0 012 2v8a2 2 0 01-2 2h-2.22l.123.489.804.804A1 1 0 0113 18H7a1 1 0 01-.707-1.707l.804-.804L7.22 15H5a2 2 0 01-2-2V5zm5.771 7H5V5h10v7H8.771z', 'clip-rule': 'evenodd' })
    ]),
    'تبلت': () => h('svg', { fill: 'currentColor', viewBox: '0 0 20 20' }, [
      h('path', { 'fill-rule': 'evenodd', d: 'M6 2a2 2 0 00-2 2v12a2 2 0 002 2h8a2 2 0 002-2V4a2 2 0 00-2-2H6zm1 2h6v10H7V4z', 'clip-rule': 'evenodd' })
    ])
  }
  return icons[device] || icons['موبایل']
}

const getDeviceClass = (device: string) => {
  const classes = {
    'موبایل': 'bg-gradient-to-br from-blue-500 to-blue-600',
    'دسکتاپ': 'bg-gradient-to-br from-green-500 to-green-600',
    'تبلت': 'bg-gradient-to-br from-purple-500 to-purple-600'
  }
  return classes[device] || 'bg-gray-500'
}

const getProgressClass = (device: string) => {
  const classes = {
    'موبایل': 'bg-gradient-to-r from-blue-400 to-blue-600',
    'دسکتاپ': 'bg-gradient-to-r from-green-400 to-green-600',
    'تبلت': 'bg-gradient-to-r from-purple-400 to-purple-600'
  }
  return classes[device] || 'bg-gray-400'
}
</script>

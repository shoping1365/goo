<template>
  <div class="survey-container">
    <component :is="currentTemplate" :order="order" @submit="handleSubmit" />
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import ModernSurvey from './templates/ModernSurvey.vue'
import ElegantSurvey from './templates/ElegantSurvey.vue'
import ColorfulSurvey from './templates/ColorfulSurvey.vue'
import MinimalSurvey from './templates/MinimalSurvey.vue'
import GradientSurvey from './templates/GradientSurvey.vue'

interface Order {
  id: string
  customerName: string
  productName: string
  orderDate: string
  totalAmount: number
}

interface Props {
  order: Order
  template?: string
}

const props = defineProps<Props>()

const emit = defineEmits<{
  submit: [data: any]
}>()

const currentTemplate = computed(() => {
  switch (props.template) {
    case 'modern': return ModernSurvey
    case 'elegant': return ElegantSurvey
    case 'colorful': return ColorfulSurvey
    case 'minimal': return MinimalSurvey
    case 'gradient': return GradientSurvey
    default: return ModernSurvey
  }
})

const handleSubmit = (data: any) => {
  emit('submit', {
    orderId: props.order.id,
    ...data
  })
}
</script>

<style scoped>
.survey-container {
  max-width: 500px;
  margin: 0 auto;
  padding: 20px;
}
</style> 

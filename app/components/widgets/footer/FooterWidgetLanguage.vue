<template>
  <div class="footer-widget footer-widget--language">
    <label class="footer-widget__label" :for="selectId">{{ viewModel.label }}</label>
    <select :id="selectId" v-model="selected" class="footer-widget__select" @change="handleChange">
      <option v-for="language in viewModel.languages" :key="language.code" :value="language.code">
        {{ language.label }}
      </option>
    </select>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, watch } from 'vue'

type LanguageOption = {
  code: string
  label: string
  active?: boolean
}

const props = withDefaults(defineProps<{
  languages?: LanguageOption[]
  label?: string
}>(), {
  label: 'Language'
})

const emit = defineEmits<{ (e: 'change', payload: string): void }>()

const selectId = `footer-language-${Math.random().toString(36).slice(2, 9)}`

const normalizeLanguages = (list?: LanguageOption[]) => {
  if (!Array.isArray(list) || !list.length) {
    return [
      { code: 'en', label: 'English', active: true },
      { code: 'fa', label: 'Farsi', active: false }
    ]
  }
  return list.map((item) => ({
    code: item.code,
    label: item.label || item.code,
    active: Boolean(item.active)
  }))
}

const viewModel = computed(() => ({
  label: props.label || 'Language',
  languages: normalizeLanguages(props.languages)
}))

const selected = ref(viewModel.value.languages.find((item) => item.active)?.code || viewModel.value.languages[0]?.code || 'en')

watch(
  () => props.languages,
  () => {
    const active = viewModel.value.languages.find((item) => item.active)
    if (active) {
      selected.value = active.code
    }
  },
  { deep: true }
)

const handleChange = () => {
  emit('change', selected.value)
}
</script>

<style scoped>
.footer-widget {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.footer-widget__label {
  font-size: 0.875rem;
  font-weight: 600;
  color: var(--footer-widget-text-color, #111827);
}

.footer-widget__select {
  border: 1px solid var(--footer-widget-input-border, #d1d5db);
  border-radius: 0.5rem;
  padding: 0.5rem 0.75rem;
  font-size: 0.875rem;
}
</style>

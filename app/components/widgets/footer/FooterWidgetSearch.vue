<template>
  <form class="footer-widget footer-widget--search" :action="viewModel.action" method="get" @submit.prevent="handleSubmit">
    <label class="footer-widget__label" :for="inputId">{{ viewModel.label }}</label>
    <div class="footer-widget__field">
      <input
        :id="inputId"
        v-model="query"
        type="search"
        :name="viewModel.queryParam"
        :placeholder="viewModel.placeholder"
        class="footer-widget__input"
      />
      <button type="submit" class="footer-widget__button">{{ viewModel.buttonLabel }}</button>
    </div>
  </form>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'

const props = withDefaults(defineProps<{
  action?: string
  placeholder?: string
  buttonLabel?: string
  label?: string
  queryParam?: string
}>(), {
  action: '/search',
  placeholder: 'Search...',
  buttonLabel: 'Search',
  label: 'Search the catalog',
  queryParam: 'q'
})

const emit = defineEmits<{ (e: 'submit', payload: string): void }>()

const query = ref('')
const inputId = `footer-search-${Math.random().toString(36).slice(2, 9)}`
const viewModel = computed(() => ({
  action: props.action || '/search',
  placeholder: props.placeholder || 'Search...',
  buttonLabel: props.buttonLabel || 'Search',
  label: props.label || 'Search the catalog',
  queryParam: props.queryParam || 'q'
}))

const handleSubmit = () => {
  emit('submit', query.value)
  if (typeof window === 'undefined' || !viewModel.value.action) {
    return
  }
  const url = new URL(viewModel.value.action, window.location.origin)
  url.searchParams.set(viewModel.value.queryParam, query.value)
  window.location.href = url.toString()
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

.footer-widget__field {
  display: flex;
  align-items: stretch;
  gap: 0.5rem;
}

.footer-widget__input {
  flex: 1;
  min-width: 0;
  border: 1px solid var(--footer-widget-input-border, #d1d5db);
  border-radius: 0.5rem;
  padding: 0.5rem 0.75rem;
  font-size: 0.875rem;
}

.footer-widget__button {
  border: none;
  border-radius: 0.5rem;
  padding: 0 1rem;
  background: var(--footer-widget-button-bg, #2563eb);
  color: #fff;
  font-weight: 600;
  cursor: pointer;
}

.footer-widget__button:hover {
  background: var(--footer-widget-button-bg-hover, #1d4ed8);
}
</style>

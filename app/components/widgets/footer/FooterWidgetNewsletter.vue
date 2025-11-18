<template>
  <form class="footer-widget footer-widget--newsletter" :action="viewModel.action" method="post" @submit.prevent="handleSubmit">
    <h4 v-if="viewModel.title" class="footer-widget__title">{{ viewModel.title }}</h4>
    <p v-if="viewModel.description" class="footer-widget__description">{{ viewModel.description }}</p>
    <div class="footer-widget__field">
      <input
        :id="inputId"
        v-model="email"
        type="email"
        name="email"
        :placeholder="viewModel.placeholder"
        required
        class="footer-widget__input"
      />
      <button type="submit" class="footer-widget__button">{{ viewModel.buttonLabel }}</button>
    </div>
    <p
      v-if="feedback"
      class="footer-widget__feedback"
      :class="{ error: feedbackType === 'error' }"
    >
      {{ feedback }}
    </p>
  </form>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'

const props = withDefaults(defineProps<{
  title?: string
  description?: string
  placeholder?: string
  buttonLabel?: string
  action?: string
  successMessage?: string
  errorMessage?: string
}>(), {
  title: 'Stay in the loop',
  description: 'Subscribe to receive the latest updates.',
  placeholder: 'Enter your email',
  buttonLabel: 'Subscribe',
  action: '',
  successMessage: 'Thanks for subscribing!',
  errorMessage: 'Please enter a valid email address.'
})

const emit = defineEmits<{
  (e: 'submit', payload: string): void
}>()

const inputId = `footer-newsletter-${Math.random().toString(36).slice(2, 9)}`
const email = ref('')
const feedback = ref('')
const feedbackType = ref<'success' | 'error' | ''>('')

const viewModel = computed(() => ({
  title: props.title?.trim() || '',
  description: props.description?.trim() || '',
  placeholder: props.placeholder?.trim() || 'Enter your email',
  buttonLabel: props.buttonLabel?.trim() || 'Subscribe',
  action: props.action?.trim() || '',
  successMessage: props.successMessage || 'Thanks for subscribing!',
  errorMessage: props.errorMessage || 'Please enter a valid email address.'
}))

const validateEmail = (value: string) => {
  return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(value)
}

const handleSubmit = () => {
  if (!validateEmail(email.value)) {
    feedback.value = viewModel.value.errorMessage
    feedbackType.value = 'error'
    return
  }
  feedback.value = viewModel.value.successMessage
  feedbackType.value = 'success'
  emit('submit', email.value)
}
</script>

<style scoped>
.footer-widget {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.footer-widget__title {
  margin: 0;
  font-size: 0.95rem;
  font-weight: 600;
  color: var(--footer-widget-text-color, #111827);
}

.footer-widget__description {
  margin: 0;
  font-size: 0.85rem;
  color: var(--footer-widget-muted-color, #6b7280);
}

.footer-widget__field {
  display: flex;
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

.footer-widget__feedback {
  margin: 0;
  font-size: 0.8rem;
  color: var(--footer-widget-feedback-color, #10b981);
}

.footer-widget__feedback.error {
  color: var(--footer-widget-feedback-error-color, #ef4444);
}
</style>

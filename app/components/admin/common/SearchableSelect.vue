<template>
  <div ref="container" class="relative">
    <!-- Input field -->
    <input
      ref="inputRef"
      v-model="searchTerm"
      :placeholder="placeholder"
      class="block w-full pl-3 pr-8 py-1.5 border border-gray-300 rounded-md leading-5 bg-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-xs shadow-sm"
      dir="rtl"
      autocomplete="off"
      @focus="onFocus"
      @input="onInput"
      @keydown="onKeyDown"
      @click="onInputClick"
    />
    
    <!-- Dropdown arrow -->
    <div class="absolute inset-y-0 left-0 flex items-center pl-2">
      <button
        type="button"
        class="text-gray-400 hover:text-gray-600 focus:outline-none"
        tabindex="-1"
        @click="toggleDropdown"
      >
        <svg
          class="w-4 h-4 transition-transform duration-200"
          :class="{ 'rotate-180': showDropdown }"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
        >
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
        </svg>
      </button>
    </div>

    <!-- Clear button -->
    <div v-if="selectedValue || searchTerm" class="absolute inset-y-0 left-6 flex items-center pl-2">
      <button
        type="button"
        class="text-gray-400 hover:text-gray-600 focus:outline-none"
        tabindex="-1"
        @click="clearSelection"
      >
        <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
        </svg>
      </button>
    </div>

    <!-- Dropdown list rendered in body -->
    <teleport to="body">
      <ul
        v-show="showDropdown"
        :style="{
          position: 'fixed',
          top: dropdownPosition.top + 'px',
          left: dropdownPosition.left + 'px',
          width: dropdownPosition.width + 'px',
          zIndex: 2147483647
        }"
      >
        <li
          v-for="(option, index) in filteredOptions"
          :key="option.value"
          :class="[
            'px-3 py-2 text-xs cursor-pointer transition-colors',
            {
              'bg-blue-100 text-blue-900': highlightedIndex === index,
              'hover:bg-gray-100': highlightedIndex !== index
            }
          ]"
          class="text-right"
          @click="selectOption(option)"
          @mouseenter="highlightedIndex = index"
        >
          {{ option.label }}
        </li>
        <li v-if="filteredOptions.length === 0" class="px-3 py-2 text-xs text-gray-500 text-right">
          نتیجه‌ای یافت نشد
        </li>
      </ul>
    </teleport>
  </div>
</template>

<script setup>
import { ref, computed, watch, nextTick, onMounted, onUnmounted } from 'vue'

const props = defineProps({
  modelValue: {
    type: [String, Number],
    default: ''
  },
  options: {
    type: Array,
    required: true,
    default: () => []
  },
  placeholder: {
    type: String,
    default: 'انتخاب کنید...'
  },
  labelKey: {
    type: String,
    default: 'label'
  },
  valueKey: {
    type: String,
    default: 'value'
  }
})

const emit = defineEmits(['update:modelValue'])

// Refs
const container = ref(null)
const inputRef = ref(null)

// State
const searchTerm = ref('')
const showDropdown = ref(false)
const highlightedIndex = ref(-1)
const selectedValue = ref(props.modelValue)

// Dropdown position state
const dropdownPosition = ref({ top: 0, left: 0, width: 0 })

function updateDropdownPosition() {
  if (inputRef.value) {
    const rect = inputRef.value.getBoundingClientRect()
    dropdownPosition.value = {
      top: rect.bottom,
      left: rect.left,
      width: rect.width
    }
    console.log('🟢 dropdownPosition:', dropdownPosition.value)
  }
}

watch(showDropdown, (show) => {
  if (show) {
    nextTick(() => {
      updateDropdownPosition()
    })
  }
})

// Computed
const normalizedOptions = computed(() => {
  if (!props.options || !Array.isArray(props.options)) {
    return []
  }
  
  return props.options.map(option => {
    if (typeof option === 'object' && option !== null) {
      return {
        label: String(option[props.labelKey] || ''),
        value: option[props.valueKey]
      }
    }
    return {
      label: String(option),
      value: option
    }
  })
})

const filteredOptions = computed(() => {
  // DEBUG: Log filtered options
  console.log('🔎 [SearchableSelect] filteredOptions:', searchTerm.value, filteredOptionsRaw())
  if (!searchTerm.value || searchTerm.value.trim() === '') {
    return normalizedOptions.value
  }
  const term = searchTerm.value.toLowerCase()
  return normalizedOptions.value.filter(option =>
    option.label.toLowerCase().includes(term)
  )
})

function filteredOptionsRaw() {
  if (!searchTerm.value || searchTerm.value.trim() === '') {
    return normalizedOptions.value
  }
  const term = searchTerm.value.toLowerCase()
  return normalizedOptions.value.filter(option =>
    option.label.toLowerCase().includes(term)
  )
}

const selectedOption = computed(() => {
  return normalizedOptions.value.find(option => option.value == selectedValue.value)
})

// Methods
const onFocus = () => {
  showDropdown.value = true
}

const onInputClick = () => {
  if (!showDropdown.value) {
    showDropdown.value = true
  }
}

const onInput = () => {
  if (!showDropdown.value) {
    showDropdown.value = true
  }
  highlightedIndex.value = -1
  
  // If there's an exact match, select it
  const exactMatch = normalizedOptions.value.find(option => 
    option.label.toLowerCase() === searchTerm.value.toLowerCase()
  )
  
  if (exactMatch) {
    selectedValue.value = exactMatch.value
    emit('update:modelValue', exactMatch.value)
  } else if (!searchTerm.value) {
    // If search is cleared, clear selection
    selectedValue.value = ''
    emit('update:modelValue', '')
  }
}

const selectOption = (option) => {
  selectedValue.value = option.value
  searchTerm.value = option.label
  showDropdown.value = false
  highlightedIndex.value = -1
  emit('update:modelValue', option.value === '' ? '' : option.value)
  inputRef.value?.blur()
}

const clearSelection = () => {
  selectedValue.value = ''
  searchTerm.value = ''
  showDropdown.value = false
  emit('update:modelValue', '')
  inputRef.value?.focus()
}

const toggleDropdown = () => {
  console.log('🔄 Toggle dropdown clicked, current state:', showDropdown.value)
  showDropdown.value = !showDropdown.value
  console.log('🔄 New dropdown state:', showDropdown.value, 'Options count:', filteredOptions.value.length)
  if (showDropdown.value) {
    nextTick(() => {
      inputRef.value?.focus()
    })
  }
}

const onKeyDown = (event) => {
  switch (event.key) {
    case 'ArrowDown':
      event.preventDefault()
      if (!showDropdown.value) {
        showDropdown.value = true
        // Give the DOM time to render before setting highlight
        nextTick(() => {
          if (filteredOptions.value.length > 0) {
            highlightedIndex.value = 0
          }
        })
      } else {
        if (filteredOptions.value.length > 0) {
          if (highlightedIndex.value < filteredOptions.value.length - 1) {
            highlightedIndex.value++
          } else {
            highlightedIndex.value = 0
          }
        }
      }
      break
      
    case 'ArrowUp':
      event.preventDefault()
      if (!showDropdown.value) {
        showDropdown.value = true
        // Give the DOM time to render before setting highlight
        nextTick(() => {
          if (filteredOptions.value.length > 0) {
            highlightedIndex.value = filteredOptions.value.length - 1
          }
        })
      } else {
        if (filteredOptions.value.length > 0) {
          if (highlightedIndex.value > 0) {
            highlightedIndex.value--
          } else {
            highlightedIndex.value = filteredOptions.value.length - 1
          }
        }
      }
      break
      
    case 'Enter':
      event.preventDefault()
      if (showDropdown.value && highlightedIndex.value >= 0 && filteredOptions.value[highlightedIndex.value]) {
        selectOption(filteredOptions.value[highlightedIndex.value])
      }
      break
      
    case 'Escape':
      showDropdown.value = false
      inputRef.value?.blur()
      break

    case 'Tab':
      showDropdown.value = false
      break
  }
}

const handleClickOutside = (event) => {
  if (container.value && !container.value.contains(event.target)) {
    showDropdown.value = false
  }
}

// Watchers
watch(() => props.modelValue, (newValue) => {
  selectedValue.value = newValue
  if (newValue) {
    const option = normalizedOptions.value.find(opt => opt.value == newValue)
    searchTerm.value = option ? option.label : ''
  } else {
    searchTerm.value = ''
  }
}, { immediate: true })

// Watch options to update display when they change
watch(() => normalizedOptions.value, (newOptions) => {
  if (selectedValue.value) {
    const option = newOptions.find(opt => opt.value == selectedValue.value)
    if (option) {
      searchTerm.value = option.label
    }
  }
}, { immediate: true })

watch(showDropdown, (show) => {
  if (show) {
    highlightedIndex.value = -1
  }
})

// Lifecycle
onMounted(() => {
  document.addEventListener('click', handleClickOutside)
  console.log('🎯 SearchableSelect mounted with options:', props.options.length)
  console.log('🟣 inputRef onMounted:', inputRef.value)
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>

<style scoped>
/* استایل واضح برای dropdown */
ul {
  background: #fff;
  color: #222;
  font-size: 16px;
  border: 2px solid #1976d2;
  box-shadow: 0 8px 24px 0 rgba(0,0,0,0.18);
  margin: 0;
  padding: 0;
  max-height: 300px;
  overflow-y: auto;
}
li {
  background: #fff;
  color: #222;
  border-bottom: 1px solid #eee;
  min-height: 32px;
  display: block;
  list-style: none;
  padding: 8px 12px;
  cursor: pointer;
}
li:last-child {
  border-bottom: none;
}
</style> 

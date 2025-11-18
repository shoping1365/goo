import { defineNuxtPlugin } from 'nuxt/app'

export default defineNuxtPlugin(() => {
  if (import.meta.client) {
    // Override any global space blocking
    const forceEnableSpace = () => {
      // Find all inputs and textareas with data-hotkey-ignore
      const elements = document.querySelectorAll('input[data-hotkey-ignore], textarea[data-hotkey-ignore]')

      elements.forEach(element => {
        // Remove any existing space blocking handlers
        element.addEventListener('keydown', (e: KeyboardEvent) => {
          if (e.code === 'Space' || e.key === ' ' || e.keyCode === 32) {
            // Force allow space
            e.stopPropagation()
            e.stopImmediatePropagation()
          }
        }, { capture: true, passive: false })

        element.addEventListener('keypress', (e: KeyboardEvent) => {
          if (e.code === 'Space' || e.key === ' ' || e.keyCode === 32) {
            e.stopPropagation()
            e.stopImmediatePropagation()
          }
        }, { capture: true, passive: false })

        element.addEventListener('input', (e: InputEvent) => {
          // Ensure space characters are preserved
          if (e.inputType === 'insertText' && e.data === ' ') {
            e.stopPropagation()
            e.stopImmediatePropagation()
          }
        }, { capture: true })
      })
    }

    // Run on page load and when DOM changes
    forceEnableSpace()

    // Watch for new elements being added
    const observer = new MutationObserver(() => {
      forceEnableSpace()
    })

    observer.observe(document.body, {
      childList: true,
      subtree: true
    })

    // Additional safety: block any global space prevention
    document.addEventListener('keydown', (e: KeyboardEvent) => {
      if (e.code === 'Space' || e.key === ' ' || e.keyCode === 32) {
        const target = e.target as HTMLElement
        if (target && (target.tagName === 'INPUT' || target.tagName === 'TEXTAREA') && target.hasAttribute('data-hotkey-ignore')) {
          e.stopPropagation()
          e.stopImmediatePropagation()
        }
      }
    }, { capture: true })
  }
}) 
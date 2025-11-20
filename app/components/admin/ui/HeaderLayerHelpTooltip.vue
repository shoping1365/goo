<template>
  <div class="help-tooltip-container">
    <!-- دکمه راهنما -->
    <button
      class="help-button"
      type="button"
      :aria-label="'راهنمای نمایش هدر و لایه‌ها'"
      :aria-expanded="isTooltipVisible"
      @click="toggleTooltip"
    >
      <svg
        class="help-icon"
        width="16"
        height="16"
        viewBox="0 0 24 24"
        fill="currentColor"
      >
        <path
          d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm1 17h-2v-2h2v2zm2.07-7.75l-.9.92C13.45 12.9 13 13.5 13 15h-2v-.5c0-1.1.45-2.1 1.17-2.83l1.24-1.26c.37-.36.59-.86.59-1.41 0-1.1-.9-2-2-2s-2 .9-2 2H8c0-2.21 1.79-4 4-4s4 1.79 4 4c0 .88-.36 1.68-.93 2.25z"
        />
      </svg>
    </button>

    <!-- راهنمای تعاملی -->
    <div
      v-if="isTooltipVisible"
      class="tooltip-content"
      :class="{ 'tooltip-visible': isTooltipVisible }"
    >
      <div class="tooltip-header">
        <h4 class="tooltip-title">راهنمای نمایش هدر و لایه‌ها</h4>
        <button
          class="close-button"
          type="button"
          aria-label="بستن راهنما"
          @click="closeTooltip"
        >
          <svg width="16" height="16" viewBox="0 0 24 24" fill="currentColor">
            <path
              d="M19 6.41L17.59 5 12 10.59 6.41 5 5 6.41 10.59 12 5 17.59 6.41 19 12 13.41 17.59 19 19 17.59 13.41 12z"
            />
          </svg>
        </button>
      </div>

      <div class="tooltip-body">
        <div class="help-section">
          <h5 class="section-title">منطق نمایش:</h5>
          <ul class="help-list">
            <li class="help-item">
              <span class="item-icon exclude">✗</span>
              <span class="item-text">
                <strong>اگر در لیست مستثنی‌های هدر باشد</strong> → هدر لود نمی‌شود
              </span>
            </li>
            <li class="help-item">
              <span class="item-icon exclude">✗</span>
              <span class="item-text">
                <strong>اگر در لیست مستثنی‌های لایه باشد</strong> → لایه لود نمی‌شود
              </span>
            </li>
            <li class="help-item">
              <span class="item-icon include">✓</span>
              <span class="item-text">
                <strong>اگر در لیست صفحات خاص هدر باشد</strong> → هدر لود می‌شود
              </span>
            </li>
            <li class="help-item">
              <span class="item-icon include">✓</span>
              <span class="item-text">
                <strong>اگر در لیست صفحات خاص لایه باشد</strong> → لایه لود می‌شود
              </span>
            </li>
            <li class="help-item">
              <span class="item-icon default">●</span>
              <span class="item-text">
                <strong>اگر هیچ‌کدام نباشد</strong> → پیش‌فرض: هر دو لود می‌شوند
              </span>
            </li>
          </ul>
        </div>

        <div class="help-section">
          <h5 class="section-title">مثال‌های استفاده:</h5>
          <div class="examples">
            <div class="example-item">
              <h6>نمایش در تمام صفحات:</h6>
              <code>pageSelection: "all"</code>
            </div>
            <div class="example-item">
              <h6>مستثنی کردن صفحات:</h6>
              <code>pageSelection: "exclude"</code><br>
              <code>excludedPages: "/admin,/auth/login"</code>
            </div>
            <div class="example-item">
              <h6>نمایش فقط در صفحات خاص:</h6>
              <code>pageSelection: "specific"</code><br>
              <code>specificPages: "/,/about,/contact"</code>
            </div>
          </div>
        </div>

        <div class="help-section">
          <h5 class="section-title">نکات مهم:</h5>
          <ul class="notes-list">
            <li>تنظیمات هدر و لایه‌ها مستقل از هم عمل می‌کنند</li>
            <li>اولویت با مستثنی کردن است، سپس صفحات خاص</li>
            <li>آدرس‌ها باید دقیق باشند (مثل: /product/sku-123)</li>
            <li>می‌توانید چندین آدرس را با کاما جدا کنید</li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
const isTooltipVisible = ref(false)

const toggleTooltip = () => {
  isTooltipVisible.value = !isTooltipVisible.value
}

const closeTooltip = () => {
  isTooltipVisible.value = false
}

// بستن راهنما با کلیک خارج از آن
onMounted(() => {
  const handleClickOutside = (event) => {
    if (!event.target.closest('.help-tooltip-container')) {
      isTooltipVisible.value = false
    }
  }

  document.addEventListener('click', handleClickOutside)

  onUnmounted(() => {
    document.removeEventListener('click', handleClickOutside)
  })
})
</script>

<style scoped>
.help-tooltip-container {
  position: relative;
  display: inline-block;
}

.help-button {
  background: none;
  border: none;
  cursor: pointer;
  padding: 4px;
  border-radius: 50%;
  transition: all 0.2s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #6b7280;
}

.help-button:hover {
  background-color: #f3f4f6;
  color: #374151;
}

.help-icon {
  width: 16px;
  height: 16px;
}

.tooltip-content {
  position: absolute;
  top: 100%;
  right: 0;
  width: 400px;
  max-width: 90vw;
  background: white;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  box-shadow: 0 10px 25px -5px rgba(0, 0, 0, 0.1), 0 10px 10px -5px rgba(0, 0, 0, 0.04);
  z-index: 1000;
  opacity: 0;
  transform: translateY(-10px);
  transition: all 0.2s ease;
}

.tooltip-visible {
  opacity: 1;
  transform: translateY(0);
}

.tooltip-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px 20px;
  border-bottom: 1px solid #e5e7eb;
  background-color: #f9fafb;
  border-radius: 8px 8px 0 0;
}

.tooltip-title {
  margin: 0;
  font-size: 16px;
  font-weight: 600;
  color: #111827;
}

.close-button {
  background: none;
  border: none;
  cursor: pointer;
  padding: 4px;
  border-radius: 4px;
  color: #6b7280;
  transition: all 0.2s ease;
}

.close-button:hover {
  background-color: #e5e7eb;
  color: #374151;
}

.tooltip-body {
  padding: 20px;
  max-height: 60vh;
  overflow-y: auto;
}

.help-section {
  margin-bottom: 24px;
}

.help-section:last-child {
  margin-bottom: 0;
}

.section-title {
  margin: 0 0 12px 0;
  font-size: 14px;
  font-weight: 600;
  color: #374151;
  border-bottom: 1px solid #e5e7eb;
  padding-bottom: 8px;
}

.help-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

.help-item {
  display: flex;
  align-items: flex-start;
  margin-bottom: 12px;
  padding: 8px 0;
}

.item-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 20px;
  height: 20px;
  border-radius: 50%;
  margin-left: 12px;
  font-size: 12px;
  font-weight: bold;
  flex-shrink: 0;
}

.item-icon.exclude {
  background-color: #fee2e2;
  color: #dc2626;
}

.item-icon.include {
  background-color: #dcfce7;
  color: #16a34a;
}

.item-icon.default {
  background-color: #dbeafe;
  color: #2563eb;
}

.item-text {
  font-size: 13px;
  line-height: 1.5;
  color: #4b5563;
}

.examples {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.example-item {
  background-color: #f9fafb;
  border: 1px solid #e5e7eb;
  border-radius: 6px;
  padding: 12px;
}

.example-item h6 {
  margin: 0 0 8px 0;
  font-size: 13px;
  font-weight: 600;
  color: #374151;
}

.example-item code {
  background-color: #f3f4f6;
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 12px;
  color: #1f2937;
  font-family: 'Courier New', monospace;
}

.notes-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

.notes-list li {
  position: relative;
  padding-right: 20px;
  margin-bottom: 8px;
  font-size: 13px;
  color: #6b7280;
  line-height: 1.5;
}

.notes-list li:before {
  content: "•";
  position: absolute;
  right: 0;
  color: #9ca3af;
}

/* Responsive */
@media (max-width: 640px) {
  .tooltip-content {
    width: calc(100vw - 32px);
    right: -16px;
  }
  
  .tooltip-body {
    padding: 16px;
  }
  
  .help-section {
    margin-bottom: 20px;
  }
}
</style>

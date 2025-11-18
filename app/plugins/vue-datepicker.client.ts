import { defineNuxtPlugin } from 'nuxt/app'
import { VueDatePicker } from '@vuepic/vue-datepicker'
import '@vuepic/vue-datepicker/dist/main.css'

// پلاگین تنظیمات تاریخ برای Vue Datepicker
export default defineNuxtPlugin((nuxtApp) => {
     nuxtApp.vueApp.component('date-picker', VueDatePicker)
})

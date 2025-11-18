import { defineStore } from 'pinia'

// @ts-ignore - persist comes from pinia-plugin-persistedstate
export const useAdminThemeStore = defineStore('adminTheme', {
  state: () => ({
    font: 'iransans' as string
  }),
  persist: true,
}) 
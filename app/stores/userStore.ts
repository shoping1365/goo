import { defineStore } from 'pinia'

interface User {
  id: number
  name: string
  email: string
  phone?: string
  status: string
  created_at: string
  updated_at: string
}

interface UserState {
  users: User[]
  loading: boolean
  error: string | null
}

export const useUserStore = defineStore('user', {
  state: (): UserState => ({
    users: [],
    loading: false,
    error: null
  }),

  actions: {
    // دریافت لیست کاربران
    async fetchUsers() {
      this.loading = true
      this.error = null
      try {
        const response = await $fetch<unknown>('/api/users')
        this.users = Array.isArray(response) ? response : (response as { data?: unknown[] })?.data || []
      } catch {
        this.error = 'خطا در دریافت لیست کاربران'
      } finally {
        this.loading = false
      }
    },

    // ثبت کاربر جدید
    async registerUser(userData: Record<string, unknown>) {
      this.loading = true
      this.error = null
      try {
        const response = await $fetch<User>('/api/users', {
          method: 'POST',
          body: userData
        })
        // اضافه کردن کاربر جدید به لیست
        this.users.push(response)
        return response
      } catch (error) {
        this.error = 'خطا در ثبت کاربر'
        throw error
      } finally {
        this.loading = false
      }
    },

    // به‌روزرسانی کاربر
    async updateUser(id: number, userData: Record<string, unknown>) {
      this.loading = true
      this.error = null
      try {
        const response = await $fetch<User>(`/api/users/${id}`, {
          method: 'PUT',
          body: userData
        })
        // به‌روزرسانی کاربر در لیست
        const index = this.users.findIndex(user => user.id === id)
        if (index !== -1) {
          this.users[index] = response
        }
        return response
      } catch (error) {
        this.error = 'خطا در به‌روزرسانی کاربر'
        throw error
      } finally {
        this.loading = false
      }
    },

    // حذف کاربر
    async deleteUser(id: number) {
      this.loading = true
      this.error = null
      try {
        await $fetch(`/api/users/${id}`, {
          method: 'DELETE'
        })
        // حذف کاربر از لیست
        this.users = this.users.filter(user => user.id !== id)
      } catch (error) {
        this.error = 'خطا در حذف کاربر'
        throw error
      } finally {
        this.loading = false
      }
    },

    // اکسپورت به Excel
    async exportExcel() {
      try {
        // ایجاد فایل Excel با استفاده از داده‌های کاربران
        const data = this.users.map(user => ({
          'شناسه': user.id,
          'نام': user.name,
          'ایمیل': user.email,
          'تلفن': user.phone || '',
          'وضعیت': user.status,
          'تاریخ ثبت‌نام': new Date(user.created_at).toLocaleDateString('fa-IR'),
          'آخرین به‌روزرسانی': new Date(user.updated_at).toLocaleDateString('fa-IR')
        }))

        // تبدیل به CSV
        const headers = Object.keys(data[0] || {})
        const csvContent = [
          headers.join(','),
          ...data.map(row => headers.map((header: string) => `"${(row as Record<string, unknown>)[header]}"`).join(','))
        ].join('\n')

        // دانلود فایل
        const blob = new Blob([csvContent], { type: 'text/csv;charset=utf-8;' })
        const link = document.createElement('a')
        const url = URL.createObjectURL(blob)
        link.setAttribute('href', url)
        link.setAttribute('download', `users_${new Date().toISOString().split('T')[0]}.csv`)
        link.style.visibility = 'hidden'
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
      } catch {
        this.error = 'خطا در اکسپورت فایل'
      }
    }
  },

  getters: {
    // دریافت کاربر بر اساس ID
    getUserById: (state) => (id: number) => {
      return state.users.find(user => user.id === id)
    },

    // فیلتر کردن کاربران بر اساس وضعیت
    getUsersByStatus: (state) => (status: string) => {
      return state.users.filter(user => user.status === status)
    },

    // جستجو در کاربران
    searchUsers: (state) => (query: string) => {
      const searchTerm = query.toLowerCase()
      return state.users.filter(user =>
        user.name.toLowerCase().includes(searchTerm) ||
        user.email.toLowerCase().includes(searchTerm) ||
        (user.phone && user.phone.includes(searchTerm))
      )
    }
  }
}) 
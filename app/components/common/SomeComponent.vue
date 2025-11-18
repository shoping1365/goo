<template>
  <div>
    <button @click="fetchData">Fetch Data</button>
  </div>
</template>

<script>
export default {
  methods: {
    async fetchData() {
      try {
        const response = await this.$api.get('/some-endpoint');
        // پردازش داده‌ها
      } catch (error) {
        const msg = error.response?.data?.message || 'یک خطا رخ داده است.';
        if (this.isAdminOrDeveloper()) {
          alert(msg);
        }
      }
    },
    isAdminOrDeveloper() {
      // منطق برای بررسی اینکه آیا کاربر ادمین یا دولوپر است
      return true; // اینجا باید منطق واقعی را پیاده‌سازی کنید
    }
  }
}
</script> 
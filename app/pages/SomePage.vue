<template>
  <div>
    <h1>Some Page</h1>
    <button @click="loadData">Load Data</button>
  </div>
</template>

<script>
export default {
  methods: {
    async loadData() {
      try {
        await this.$api.get('/another-endpoint');
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
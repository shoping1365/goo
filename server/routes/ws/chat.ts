// server/routes/ws/chat.ts
// وب‌سوکت ساده برای جلوگیری از خطای ECONNRESET هنگام باز شدن صفحه Live-Chat
// این پیاده‌سازی فعلاً فقط اتصال را برقرار نگه می‌دارد و هر 30 ثانیه بسته می‌شود.
// در آینده می‌توانید منطق واقعی چت را جایگزین کنید.

import { defineWebSocketHandler } from 'h3'

export default defineWebSocketHandler({
  open(ws) {
    // اتصال موفق
  },
  message(ws, message) {
    try {
      const data = JSON.parse(message.toString())
      if (data.type === 'ping') {
        ws.send(JSON.stringify({ type: 'pong' }))
      }
    } catch (err) {
      // ignore malformed messages
    }
  },
  close() {
  },
})
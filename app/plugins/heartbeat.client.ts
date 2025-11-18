import { defineNuxtPlugin } from 'nuxt/app'

export default defineNuxtPlugin(() => {
  // در server-side، heartbeat را غیرفعال کن
  if (import.meta.server) return;

  let timer: ReturnType<typeof setInterval> | null = null;
  let backoffMs = 5_000; // شروع با 5 ثانیه
  const maxBackoffMs = 60_000; // سقف 60 ثانیه
  let isAuthenticated = false; // وضعیت احراز هویت

  const AUTH_STATE_STORAGE_KEY = 'auth-state'

  // بررسی اینکه آیا کاربر وارد شده است یا نه
  const checkAuthentication = () => {
    // بررسی مقدار ذخیره‌شده در localStorage (در صورت وجود)
    try {
      const authState = localStorage.getItem(AUTH_STATE_STORAGE_KEY)
      if (authState === 'logged-in') {
        return true
      }
      if (authState === 'logged-out') {
        return false
      }
    } catch {
      // در صورت خطا، مسیر بعدی بررسی می‌شود
    }

    // بررسی presence کوکی‌های HttpOnly
    const cookies = document.cookie.split(';')
    const accessToken = cookies.find(cookie => cookie.trim().startsWith('access_token='))?.split('=')[1]

    return Boolean(accessToken)
  }

  const sendHeartbeat = async () => {
    // فقط اگر کاربر وارد شده باشد، heartbeat بفرست
    if (!isAuthenticated) {
      return
    }

    try {
      await $fetch('/api/users/heartbeat', {
        method: 'PUT',
        credentials: 'include'
      });
      // موفقیت: به بازه ثابت برگرد
      backoffMs = 5_000;
      try { window.dispatchEvent(new CustomEvent('users:heartbeat-ok', { detail: { ts: Date.now() } })); } catch { }
    } catch (err) {
      // قطع بک‌اند یا خطای شبکه: به شکل backoff عمل کنیم تا از فشار جلوگیری شود
      backoffMs = Math.min(backoffMs * 2, maxBackoffMs);
      if (timer) clearInterval(timer);
      timer = setInterval(sendHeartbeat, backoffMs);
    }
  };

  // تابعی برای همگام‌سازی بازه بر اساس آستانه UI (حداقل 5s، حداکثر 60s)
  const getIntervalMs = () => {
    try {
      const saved = localStorage.getItem('users_online_threshold_sec');
      const sec = saved ? parseInt(saved, 10) : 10; // پیش‌فرض 10 ثانیه
      const clamped = Math.max(5, Math.min(60, sec));
      return clamped * 1000;
    } catch {
      return 30_000;
    }
  };

  const startTimer = () => {
    if (timer) clearInterval(timer);
    const interval = getIntervalMs();
    timer = setInterval(sendHeartbeat, interval);
  };

  const stopTimer = () => {
    if (timer) {
      clearInterval(timer);
      timer = null;
    }
  };

  let listenersAttached = false

  const onThresholdChange = () => {
    if (isAuthenticated) {
      backoffMs = 5_000;
      startTimer();
      sendHeartbeat();
    }
  }

  const onVisibilityChange = () => {
    if (document.visibilityState === 'visible' && isAuthenticated) {
      sendHeartbeat();
    }
  }

  const onBeforeUnload = () => {
    stopTimer();
  }

  const attachActiveListeners = () => {
    if (listenersAttached) {
      return
    }

    window.addEventListener('users_online_threshold_changed', onThresholdChange)
    document.addEventListener('visibilitychange', onVisibilityChange)
    window.addEventListener('beforeunload', onBeforeUnload)

    listenersAttached = true
  }

  const detachActiveListeners = () => {
    if (!listenersAttached) {
      return
    }

    window.removeEventListener('users_online_threshold_changed', onThresholdChange)
    document.removeEventListener('visibilitychange', onVisibilityChange)
    window.removeEventListener('beforeunload', onBeforeUnload)

    listenersAttached = false
  }

  // بررسی اولیه احراز هویت و شروع/توقف heartbeat
  const syncHeartbeatState = () => {
    const authenticated = checkAuthentication()

    updateHeartbeatState(authenticated)
  }

  const updateHeartbeatState = (authenticated: boolean) => {
    if (authenticated) {
      if (!isAuthenticated) {
        isAuthenticated = true
        attachActiveListeners()
        startTimer()
        sendHeartbeat()
      }
    } else {
      if (isAuthenticated) {
        stopTimer()
        detachActiveListeners()
      }
      isAuthenticated = false
    }
  }

  // بررسی اولیه
  syncHeartbeatState()

  // گوش دادن به رویدادهای login/logout
  const handleAuthStateEvent = (detailAuthenticated: boolean) => {
    updateHeartbeatState(detailAuthenticated)
  }

  window.addEventListener('storage', (e) => {
    if (e.key === AUTH_STATE_STORAGE_KEY) {
      if (e.newValue === 'logged-in') {
        handleAuthStateEvent(true)
      }
      if (e.newValue === 'logged-out') {
        handleAuthStateEvent(false)
      }
    }
  })

  window.addEventListener('auth:state-changed', (event: Event) => {
    const customEvent = event as CustomEvent<{ authenticated: boolean }>
    handleAuthStateEvent(customEvent.detail?.authenticated ?? false)
  })
}); 
import { ref, computed, onUnmounted } from 'vue'

export function useLiveChat() {
  // State
  const chatSessions = ref<any[]>([])
  const selectedSession = ref<any | null>(null)
  const messages = ref<any[]>([])
  const isLoading = ref(false)
  const error = ref<string | null>(null)

  // WebSocket
  let ws: WebSocket | null = null
  let wsReconnectDelayMs = 2000
  const wsReconnectDelayMaxMs = 30000

  // Toast util - Temporarily disabled to prevent JavaScript errors
  const notify = (msg: string, type: 'error' | 'success' | 'warning' | 'info' = 'error') => {
    // Temporarily use console instead of toast to avoid vue3-toastify errors
    console.log(`[${type.toUpperCase()}] ${msg}`)
    // TODO: Re-enable toast notifications after fixing vue3-toastify cleanup issues
    /*
    try {
      const toast = typeof window !== 'undefined' ? (window as any).$toast : null
      if (toast && typeof toast[type] === 'function') {
        toast[type](msg, { duration: 5000 })
      }
    } catch (error) {
      console.debug('Toast notification failed:', error)
    }
    */
  }

  // Fetch chat sessions
  const fetchChatSessions = async (status: string = 'active') => {
    isLoading.value = true
    error.value = null
    try {
      const response: any = await $fetch(`/api/admin/chat/sessions`, {
        query: { status }
      })
      chatSessions.value = response?.data || []
    } catch (e: any) {
      error.value = e?.message || 'خطا در دریافت لیست چت‌ها'
      notify(error.value)
    } finally {
      isLoading.value = false
    }
  }

  // Fetch messages for a session
  const fetchMessages = async (sessionId: number) => {
    isLoading.value = true
    error.value = null
    try {
      const response: any = await $fetch(`/api/admin/chat/sessions/${sessionId}/messages`)
      messages.value = response?.data || []
    } catch (e: any) {
      error.value = e?.message || 'خطا در دریافت پیام‌ها'
      notify(error.value)
    } finally {
      isLoading.value = false
    }
  }

  // Send message
  const sendMessage = async (sessionId: number, text: string, file?: File | null) => {
    try {
      const form = new FormData()
      if (text) form.append('message', text)
      if (file) form.append('file', file)
      const res: any = await ($fetch as any)(`/api/admin/chat/sessions/${sessionId}/messages`, {
        method: 'POST',
        body: form
      })
      const newMsg = res?.data
      if (newMsg) messages.value.push(newMsg)
      return newMsg
    } catch (e: any) {
      error.value = e?.message || 'خطا در ارسال پیام'
      notify(error.value)
      throw e
    }
  }

  // Update session status
  const updateSessionStatus = async (sessionId: number, status: string) => {
    await ($fetch as any)(`/api/admin/chat/sessions/${sessionId}`, { method: 'PUT', body: { status } })
    await fetchChatSessions('active')
  }

  // Select session
  const selectSession = async (session: any) => {
    selectedSession.value = session
    await fetchMessages(session.id)
  }

  // Mark as read
  const markAsRead = async (sessionId: number) => {
    try {
      await ($fetch as any)(`/api/admin/chat/sessions/${sessionId}/read`, { method: 'POST' })
      const idx = chatSessions.value.findIndex(s => s.id === sessionId)
      if (idx !== -1) chatSessions.value[idx].unread_count = 0
    } catch (e) {
      notify('خطا در علامت‌گذاری پیام‌ها به عنوان خوانده شده')
    }
  }

  // Real-time metrics
  const getRealTimeMetrics = async () => {
    try {
      const response: any = await $fetch(`/api/admin/chat/metrics`)
      return response
    } catch {
      return null
    }
  }

  // WebSocket connect with notifications + backoff
  const connectWebSocket = () => {
    try {
      const protocol = typeof window !== 'undefined' && window.location.protocol === 'https:' ? 'wss:' : 'ws:'
      const base = `${protocol}//${window.location.host}`
      const path = '/api/chat/admin/ws/'
      ws = new WebSocket(`${base}${path}`)

      ws.onopen = () => {
        wsReconnectDelayMs = 2000
        notify('اتصال وب‌سوکت برقرار شد', 'success')
        // در صورت فعال بودن Polling، متوقف شود
        stopPolling()
      }

      ws.onmessage = (event) => {
        try {
          const data = JSON.parse(event.data)
          handleWebSocketMessage(data)
        } catch { }
      }

      ws.onerror = (ev: any) => {
        import('~/utils/wsErrors')
          .then(({ parseWsError }) => {
            const msg = parseWsError(ev)
            notify(msg)
            error.value = msg
          })
          .catch(() => {
            notify('خطا در اتصال WebSocket')
            error.value = 'خطا در اتصال WebSocket'
          })
      }

      ws.onclose = (e: CloseEvent) => {
        import('~/utils/wsErrors')
          .then(({ parseWsClose }) => {
            const info = parseWsClose(e)
            const retryMsg = e.wasClean ? '' : ' تلاش برای اتصال مجدد...'
            notify(`${info.userMessage}.${retryMsg}`)
          })
          .catch(() => {
            notify('اتصال وب‌سوکت قطع شد')
          })
        setTimeout(() => {
          if (ws?.readyState === WebSocket.CLOSED) connectWebSocket()
        }, wsReconnectDelayMs)
        wsReconnectDelayMs = Math.min(wsReconnectDelayMs * 2, wsReconnectDelayMaxMs)

        // در زمان قطع، Polling فعال باشد تا UI زنده بماند
        startPolling('active', 5000)
      }
    } catch (err) {
      notify('خطا در اتصال WebSocket')
      error.value = 'خطا در اتصال WebSocket'
    }
  }

  // Polling fallback (for environments without WebSocket)
  let pollTimer: ReturnType<typeof setInterval> | null = null
  const startPolling = (status: string = 'active', intervalMs = 5000) => {
    // اجرای فوری برای اولین بار
    fetchChatSessions(status)
    if (selectedSession.value?.id) {
      fetchMessages(selectedSession.value.id)
    }
    // تنظیم بازه تکرار
    stopPolling()
    pollTimer = setInterval(() => {
      fetchChatSessions(status)
      if (selectedSession.value?.id) {
        fetchMessages(selectedSession.value.id)
      }
    }, intervalMs)
  }

  const stopPolling = () => {
    if (pollTimer) {
      clearInterval(pollTimer)
      pollTimer = null
    }
  }

  const handleWebSocketMessage = (data: any) => {
    switch (data?.type) {
      case 'new_message':
        if (selectedSession.value && (selectedSession.value.session_id === data?.data?.session_id || selectedSession.value.id === data?.data?.session_id)) {
          messages.value.push(data.data)
        }
        fetchChatSessions()
        break
      case 'message':
        fetchChatSessions()
        break
      case 'session_update': {
        const idx = chatSessions.value.findIndex(s => s.session_id === data?.session_id || s.id === data?.session_id)
        if (idx !== -1) chatSessions.value[idx] = data.session
        break
      }
      case 'new_session':
        chatSessions.value.unshift(data.session)
        break
    }
  }

  const disconnectWebSocket = () => {
    if (ws) {
      ws.close()
      ws = null
    }
  }

  onUnmounted(() => {
    disconnectWebSocket()
    stopPolling()
  })

  // فعال‌سازی Realtime: ابتدا تلاش برای WS، در صورت عدم برقراری، Polling به‌صورت خودکار شروع می‌شود
  const enableRealtime = (status: string = 'active') => {
    try {
      connectWebSocket()
    } catch { }
    // اگر ظرف 5 ثانیه WS باز نشد، Polling را شروع کن
    setTimeout(() => {
      if (!ws || ws.readyState !== WebSocket.OPEN) {
        startPolling(status, 5000)
      }
    }, 5000)
  }

  return {
    chatSessions,
    selectedSession,
    messages,
    isLoading,
    error,
    fetchChatSessions,
    fetchMessages,
    sendMessage,
    updateSessionStatus,
    connectWebSocket,
    disconnectWebSocket,
    enableRealtime,
    startPolling,
    stopPolling,
    selectSession,
    markAsRead,
    getRealTimeMetrics,
  }
}
import { eventHandler, sendRedirect } from 'h3'

// Redirect legacy /admin/chat/sessions to /api/admin/chat/sessions (Nitro route)
export default eventHandler((event) => {
  const url = event.path || event.node?.req?.url || ''
  if (url.startsWith('/admin/chat/sessions')) {
    const qs = url.includes('?') ? url.substring(url.indexOf('?')) : ''
    return sendRedirect(event, `/api/admin/chat/sessions${qs}`, 307)
  }
})



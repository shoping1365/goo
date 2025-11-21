import { beforeEach, describe, expect, it, vi } from 'vitest'
import { useLiveChat } from '../../app/composables/useLiveChat'

// Mock $fetch
const mockFetch = vi.fn()
vi.stubGlobal('$fetch', mockFetch)

describe('useLiveChat Composable', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  it('should initialize with default values', () => {
    const chat = useLiveChat()
    expect(chat.chatSessions.value).toEqual([])
    expect(chat.selectedSession.value).toBeNull()
    expect(chat.messages.value).toEqual([])
    expect(chat.isLoading.value).toBe(false)
    expect(chat.error.value).toBeNull()
  })

  describe('fetchChatSessions', () => {
    it('should fetch sessions successfully', async () => {
      const mockSessions = [{ id: 1, session_id: 'sess-1' }]
      mockFetch.mockResolvedValue({ data: mockSessions })
      
      const chat = useLiveChat()
      await chat.fetchChatSessions()
      
      expect(mockFetch).toHaveBeenCalledWith('/api/admin/chat/sessions', expect.objectContaining({
        query: { status: 'active' }
      }))
      expect(chat.chatSessions.value).toEqual(mockSessions)
      expect(chat.isLoading.value).toBe(false)
    })

    it('should handle fetch error', async () => {
      mockFetch.mockRejectedValue(new Error('API Error'))
      
      const chat = useLiveChat()
      await chat.fetchChatSessions()
      
      expect(chat.error.value).toBe('خطا در دریافت لیست چت‌ها')
      expect(chat.chatSessions.value).toEqual([])
    })
  })

  describe('fetchMessages', () => {
    it('should fetch messages successfully', async () => {
      const mockMessages = [{ id: 1, message: 'Hello' }]
      mockFetch.mockResolvedValue({ data: mockMessages })
      
      const chat = useLiveChat()
      await chat.fetchMessages(1)
      
      expect(mockFetch).toHaveBeenCalledWith('/api/admin/chat/sessions/1/messages')
      expect(chat.messages.value).toEqual(mockMessages)
    })
  })

  describe('sendMessage', () => {
    it('should send message successfully', async () => {
      const newMessage = { id: 2, message: 'New msg' }
      mockFetch.mockResolvedValue({ data: newMessage })
      
      const chat = useLiveChat()
      const result = await chat.sendMessage(1, 'New msg')
      
      expect(mockFetch).toHaveBeenCalledWith('/api/admin/chat/sessions/1/messages', expect.objectContaining({
        method: 'POST',
        body: expect.any(FormData)
      }))
      expect(result).toEqual(newMessage)
      expect(chat.messages.value).toContainEqual(newMessage)
    })

    it('should handle send error', async () => {
      mockFetch.mockRejectedValue(new Error('Send failed'))
      
      const chat = useLiveChat()
      await chat.sendMessage(1, 'msg')
      
      expect(chat.error.value).toBe('خطا در ارسال پیام')
    })
  })
})

import { ref } from 'vue'

interface SendOTPResponse {
  message: string
  expires_in: number
}

interface ApiError {
  data?: {
    error?: string
  }
}

export const useOTP = () => {
  const loading = ref(false)
  const error = ref('')
  const successMessage = ref('')
  const remainingTime = ref(0)

  // ارسال OTP
  const sendOTP = async (mobile: string) => {
    loading.value = true
    error.value = ''
    successMessage.value = ''

    try {
      const response = await $fetch<SendOTPResponse>('/api/auth/send-otp', {
        method: 'POST',
        body: { mobile },
      })

      successMessage.value = response.message
      remainingTime.value = response.expires_in

      // شمارش معکوس
      const interval = setInterval(() => {
        remainingTime.value--
        if (remainingTime.value <= 0) {
          clearInterval(interval)
        }
      }, 1000)

      return true
    } catch (err) {
      const apiError = err as ApiError
      error.value = apiError.data?.error || 'ارسال OTP ناموفق بود'
      return false
    } finally {
      loading.value = false
    }
  }

  // ارسال مجدد OTP
  const resendOTP = async (mobile: string) => {
    if (remainingTime.value > 0) {
      error.value = `لطفا ${remainingTime.value} ثانیه صبر کنید`
      return false
    }

    return await sendOTP(mobile)
  }

  const resetError = () => {
    error.value = ''
  }

  return {
    loading,
    error,
    successMessage,
    remainingTime,
    sendOTP,
    resendOTP,
    resetError,
  }
}

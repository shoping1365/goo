import { defineEventHandler } from 'h3'

export default defineEventHandler(() => ({
  status: 'success',
  data: {
    paperSize: 'A4',
    includeLogo: true,
    qrCode: true
  }
}))
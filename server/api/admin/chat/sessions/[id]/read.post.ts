import { defineEventHandler } from 'h3'
import { handleMarkSessionRead } from './_shared/handleMarkSessionRead'

export default defineEventHandler(handleMarkSessionRead)
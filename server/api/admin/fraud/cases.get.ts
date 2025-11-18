import { defineEventHandler, getQuery } from 'h3';
import { proxy } from "../../_utils/fetchProxy";

export default defineEventHandler(async (event) => {
  const q = getQuery(event)
  const queryRecord: Record<string, string> = {}
  for (const [key, value] of Object.entries(q)) {
    if (value !== undefined) {
      queryRecord[key] = String(value)
    }
  }
  const qs = new URLSearchParams(queryRecord).toString()
  return proxy(event, `/api/admin/fraud/cases${qs ? `?${qs}` : ""}`)
});


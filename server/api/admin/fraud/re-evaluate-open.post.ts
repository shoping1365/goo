import { defineEventHandler } from 'h3';
import { proxy } from "../../_utils/fetchProxy";

export default defineEventHandler(async (event) => {
  return proxy(event, "/api/admin/fraud/re-evaluate-open");
});


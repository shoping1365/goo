<script setup lang="ts">
import { ref } from 'vue';

const form = ref({
  full_name: '',
  username: '',
  mobile: '',
  role: 'customer',
});
const otp = ref('');
const step = ref(1);
const message = ref('');
const error = ref('');

async function sendOtp() {
  error.value = '';
  message.value = '';
  const res = await fetch('/api/auth/send-otp', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ mobile: form.value.mobile }),
  });
  const data = await res.json();
  if (data.success) {
    message.value = data.message;
    step.value = 2;
  } else {
    error.value = data.message || 'خطا در ارسال کد تایید';
  }
}

async function handleSubmit() {
  error.value = '';
  message.value = '';
  // First, verify OTP
  const verifyRes = await fetch('/api/auth/verify-otp', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ mobile: form.value.mobile, otp: otp.value }),
  });
  const verifyData = await verifyRes.json();
  if (!verifyData.success) {
    error.value = verifyData.message || 'کد تایید اشتباه است';
    return;
  }
  // If OTP is correct, register user
  const regRes = await fetch('/api/users/register', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(form.value),
  });
  const regData = await regRes.json();
  if (regRes.ok) {
    message.value = 'ثبت نام با موفقیت انجام شد!';
    step.value = 3;
  } else {
    error.value = regData.error || 'خطا در ثبت نام';
  }
}
</script> 

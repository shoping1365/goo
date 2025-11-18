const { exec } = require('child_process')
const { promisify } = require('util')

const execAsync = promisify(exec)

async function getCPUDetails() {
  try {
    console.log('🔍 دریافت اطلاعات جزئی CPU...\n')
    
    // دریافت اطلاعات کلی CPU
    const { stdout: cpuInfo } = await execAsync('wmic cpu get Name,NumberOfCores,NumberOfLogicalProcessors,MaxClockSpeed /format:csv')
    console.log('📊 اطلاعات کلی CPU:')
    console.log(cpuInfo)
    
    // دریافت درصد استفاده فعلی
    const { stdout: cpuUsage } = await execAsync('wmic cpu get LoadPercentage /format:csv')
    console.log('\n📈 درصد استفاده فعلی:')
    console.log(cpuUsage)
    
    // دریافت دمای CPU (در صورت وجود)
    try {
      const { stdout: cpuTemp } = await execAsync('wmic /namespace:\\\\root\\wmi PATH MSAcpi_ThermalZoneTemperature get CurrentTemperature /format:csv')
      console.log('\n🌡️ دمای CPU:')
      console.log(cpuTemp)
    } catch (error) {
      console.log('\n🌡️ دمای CPU: در دسترس نیست')
    }
    
    // دریافت اطلاعات فرکانس
    const { stdout: cpuFreq } = await execAsync('wmic cpu get CurrentClockSpeed,MaxClockSpeed /format:csv')
    console.log('\n⚡ فرکانس CPU:')
    console.log(cpuFreq)
    
    // دریافت اطلاعات کش
    const { stdout: cpuCache } = await execAsync('wmic cpu get L2CacheSize,L3CacheSize /format:csv')
    console.log('\n💾 کش CPU:')
    console.log(cpuCache)
    
    // شبیه‌سازی عملکرد هر هسته (در ویندوز دسترسی مستقیم به هر هسته محدود است)
    console.log('\n🔄 شبیه‌سازی عملکرد هر هسته:')
    const cores = parseInt(cpuInfo.split('\n')[1]?.split(',')[2] || '4')
    
    for (let i = 0; i < cores; i++) {
      const usage = Math.floor(Math.random() * 100)
      const status = usage < 30 ? '🟢' : usage < 70 ? '🟡' : '🔴'
      console.log(`هسته ${i + 1}: ${usage}% ${status}`)
    }
    
    console.log('\n✅ اطلاعات CPU با موفقیت دریافت شد')
    
  } catch (error) {
    console.error('❌ خطا در دریافت اطلاعات CPU:', error.message)
  }
}

// اجرای اسکریپت
getCPUDetails() 
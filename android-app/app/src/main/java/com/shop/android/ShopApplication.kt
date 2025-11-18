package com.shop.android

import android.app.Application
import dagger.hilt.android.HiltAndroidApp

/**
 * کلاس اصلی اپلیکیشن
 * برای تنظیمات اولیه و dependency injection
 */
@HiltAndroidApp
class ShopApplication : Application() {
    
    override fun onCreate() {
        super.onCreate()
        // تنظیمات اولیه اپلیکیشن
    }
} 
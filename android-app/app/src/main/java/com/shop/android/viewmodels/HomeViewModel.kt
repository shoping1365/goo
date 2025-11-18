package com.shop.android.viewmodels

import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.shop.android.models.Product
import com.shop.android.network.ApiService
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.launch
import javax.inject.Inject

/**
 * ViewModel برای صفحه اصلی
 * مدیریت داده‌های محصولات و وضعیت loading
 */
@HiltViewModel
class HomeViewModel @Inject constructor(
    private val apiService: ApiService
) : ViewModel() {

    private val _products = MutableLiveData<List<Product>>()
    val products: LiveData<List<Product>> = _products

    private val _isLoading = MutableLiveData<Boolean>()
    val isLoading: LiveData<Boolean> = _isLoading

    private val _error = MutableLiveData<String>()
    val error: LiveData<String> = _error

    /**
     * بارگذاری محصولات از API
     */
    fun loadProducts() {
        viewModelScope.launch {
            try {
                _isLoading.value = true
                val response = apiService.getProducts()
                
                if (response.isSuccessful) {
                    _products.value = response.body() ?: emptyList()
                } else {
                    _error.value = "خطا در بارگذاری محصولات"
                }
            } catch (e: Exception) {
                _error.value = "خطا در اتصال به سرور"
            } finally {
                _isLoading.value = false
            }
        }
    }

    /**
     * جستجو در محصولات
     */
    fun searchProducts(query: String) {
        viewModelScope.launch {
            try {
                _isLoading.value = true
                val response = apiService.getProducts(search = query)
                
                if (response.isSuccessful) {
                    _products.value = response.body() ?: emptyList()
                } else {
                    _error.value = "خطا در جستجو"
                }
            } catch (e: Exception) {
                _error.value = "خطا در اتصال به سرور"
            } finally {
                _isLoading.value = false
            }
        }
    }

    /**
     * فیلتر محصولات بر اساس دسته‌بندی
     */
    fun filterByCategory(category: String) {
        viewModelScope.launch {
            try {
                _isLoading.value = true
                val response = apiService.getProducts(category = category)
                
                if (response.isSuccessful) {
                    _products.value = response.body() ?: emptyList()
                } else {
                    _error.value = "خطا در فیلتر کردن"
                }
            } catch (e: Exception) {
                _error.value = "خطا در اتصال به سرور"
            } finally {
                _isLoading.value = false
            }
        }
    }
} 
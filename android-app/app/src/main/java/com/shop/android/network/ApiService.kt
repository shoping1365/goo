package com.shop.android.network

import com.shop.android.models.Product
import com.shop.android.models.Category
import com.shop.android.models.LoginRequest
import com.shop.android.models.LoginResponse
import retrofit2.Response
import retrofit2.http.*

/**
 * رابط API برای اتصال به بک‌اند Go
 * تمام endpoint های مورد نیاز اپلیکیشن
 */
interface ApiService {
    
    // محصولات
    @GET("products")
    suspend fun getProducts(
        @Query("page") page: Int = 1,
        @Query("limit") limit: Int = 20,
        @Query("category") category: String? = null,
        @Query("search") search: String? = null
    ): Response<List<Product>>
    
    @GET("products/{id}")
    suspend fun getProduct(@Path("id") id: Int): Response<Product>
    
    // دسته‌بندی‌ها
    @GET("categories")
    suspend fun getCategories(): Response<List<Category>>
    
    // احراز هویت
    @POST("auth/login")
    suspend fun login(@Body loginRequest: LoginRequest): Response<LoginResponse>
    
    @POST("auth/register")
    suspend fun register(@Body registerRequest: RegisterRequest): Response<LoginResponse>
    
    // سبد خرید
    @GET("cart")
    suspend fun getCart(): Response<CartResponse>
    
    @POST("cart/add")
    suspend fun addToCart(@Body cartRequest: CartRequest): Response<CartResponse>
    
    @DELETE("cart/{itemId}")
    suspend fun removeFromCart(@Path("itemId") itemId: Int): Response<CartResponse>
    
    // سفارشات
    @POST("orders")
    suspend fun createOrder(@Body orderRequest: OrderRequest): Response<OrderResponse>
    
    @GET("orders")
    suspend fun getOrders(): Response<List<Order>>
}

// مدل‌های درخواست
data class LoginRequest(
    val email: String,
    val password: String
)

data class RegisterRequest(
    val name: String,
    val email: String,
    val password: String,
    val phone: String
)

data class CartRequest(
    val productId: Int,
    val quantity: Int
)

data class OrderRequest(
    val items: List<CartItem>,
    val shippingAddress: String,
    val paymentMethod: String
)

// مدل‌های پاسخ
data class LoginResponse(
    val token: String,
    val user: User
)

data class CartResponse(
    val items: List<CartItem>,
    val total: Double
)

data class OrderResponse(
    val orderId: String,
    val status: String
)

// مدل‌های کمکی
data class User(
    val id: Int,
    val name: String,
    val email: String
)

data class CartItem(
    val id: Int,
    val product: Product,
    val quantity: Int
)

data class Order(
    val id: String,
    val items: List<CartItem>,
    val total: Double,
    val status: String,
    val createdAt: String
) 
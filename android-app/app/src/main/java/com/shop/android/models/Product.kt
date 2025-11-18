package com.shop.android.models

import com.google.gson.annotations.SerializedName

/**
 * مدل داده محصول
 * مطابق با API بک‌اند Go
 */
data class Product(
    @SerializedName("id")
    val id: Int,
    
    @SerializedName("name")
    val name: String,
    
    @SerializedName("description")
    val description: String,
    
    @SerializedName("price")
    val price: Double,
    
    @SerializedName("images")
    val images: List<String>,
    
    @SerializedName("category")
    val category: String,
    
    @SerializedName("brand")
    val brand: String,
    
    @SerializedName("stock")
    val stock: Int,
    
    @SerializedName("sku")
    val sku: String? = null,
    
    @SerializedName("rating")
    val rating: Double? = null,
    
    @SerializedName("reviews_count")
    val reviewsCount: Int? = null
) 
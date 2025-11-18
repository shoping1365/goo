package com.shop.android.models

import com.google.gson.annotations.SerializedName

/**
 * مدل داده دسته‌بندی
 */
data class Category(
    @SerializedName("id")
    val id: Int,
    
    @SerializedName("name")
    val name: String,
    
    @SerializedName("slug")
    val slug: String,
    
    @SerializedName("image")
    val image: String? = null,
    
    @SerializedName("description")
    val description: String? = null,
    
    @SerializedName("parent_id")
    val parentId: Int? = null
) 
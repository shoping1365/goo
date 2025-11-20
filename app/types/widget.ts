// انواع ابزارک‌ها
export type WidgetType =
  | 'main-slider-side-banner'
  | 'single-slider-side'
  | 'full-slider'

  | 'full-banner'

  | 'double-banner'
  | 'triple-banner'
  | 'quad-banner'
  | 'penta-banner'
  | 'brands-slider'
  | 'services-slider'
  | 'category-box'
  | 'blog-posts'
  | 'product-carousel'
  | 'mobile-header';

// وضعیت ابزارک
export type WidgetStatus = 'active' | 'inactive' | 'draft';

// صفحات قابل استفاده
export type WidgetPage = 'home' | 'category' | 'product' | 'about' | 'contact';

// رابط اصلی ابزارک
export interface Widget {
  id: number;
  code: string;
  title: string;
  description?: string;
  type: WidgetType;
  status: WidgetStatus;
  order: number;
  image?: string;
  link?: string;
  page: WidgetPage;
  show_on_mobile: boolean;
  config: WidgetConfig;
  created_at: string;
  updated_at: string;
}

// تنظیمات پایه ابزارک
export interface WidgetConfig {
  mobile_height?: number | 'auto'; // ارتفاع اسلایدر موبایل
  show_in_mobile?: boolean; // نمایش در موبایل
  padding?: string;
  margin?: string;
  custom_class?: string;
}

// تنظیمات اسلایدر
export interface SliderConfig extends WidgetConfig {
  wide_bg: boolean;
  bg_color: string;
  bg_enabled: boolean;
  banner_position?: 'left' | 'right';
  slider_banner_ratio?: string; // نسبت اسلایدر و بنر (مثل "7-5", "6-6", "5-7")
  height: number;
  padding_top: number;
  padding_bottom: number;
  margin_left: number;
  margin_right: number;
  autoplay_delay: number;
  autoplay_enabled: boolean;
  navigation_enabled: boolean;
  slides: SlideItem[];
  side_banners?: BannerItem[];
  mobile_slides?: SlideItem[];
  mobile_side_banners?: BannerItem[];
  // تنظیمات جدید اسلایدر تکی
  slider_count?: number;
  display_order?: 'asc' | 'desc';
  show_navigation?: boolean;
  show_pagination?: boolean;
  show_title?: boolean;
  show_description?: boolean;
  transition_duration?: number;
  navigation_style?: 'default' | 'minimal' | 'bold';
  pagination_style?: 'default' | 'small' | 'large';
  title_style?: 'default' | 'custom';
  loop_enabled?: boolean;
  pause_on_hover?: boolean;
  // تنظیمات ظاهری
  border_radius?: number;
  border_width?: number;
  border_color?: string;
  shadow_enabled?: boolean;
  shadow_style?: 'default' | 'large';
  navigation_size?: number;
  pagination_color?: string;
  title_background_opacity?: number;
  title_text_color?: string;
  title_font_size?: number;
  // Mobile settings - تنظیمات کامل موبایل
  mobile_height?: number | 'auto';
  mobile_padding_top?: number;
  mobile_padding_bottom?: number;
  mobile_margin_left?: number;
  mobile_margin_right?: number;
  mobile_show_slider?: boolean;
  mobile_bg_enabled?: boolean;
  mobile_bg_color?: string;
  mobile_wide_bg?: boolean;
  // Additional mobile settings for slider
  mobile_autoplay_enabled?: boolean;
  mobile_autoplay_delay?: number;
  mobile_loop_enabled?: boolean;
  mobile_show_navigation?: boolean;
  mobile_show_pagination?: boolean;
  mobile_show_title?: boolean;
  mobile_show_description?: boolean;
  mobile_vertical_display?: boolean;
  mobile_slider_width?: number | 'auto';
  mobile_banner_width?: number | 'auto';
  mobile_banner_height?: number | 'auto';
  mobile_banner_position?: 'top' | 'bottom';
  // Easy load feature
  easy_load_enabled?: boolean;
  // Desktop slider width settings
  slider_width?: number;
  center_width?: number;
  banner_width?: number;
  // نسبت‌های بنر برای بنرهای چندتایی
  banner1_ratio?: number;
  banner2_ratio?: number;
  banner3_ratio?: number;
  banner4_ratio?: number;
  banner5_ratio?: number;
}

// تنظیمات بنر
export interface BannerConfig extends WidgetConfig {
  layout?: 'single' | 'double' | 'triple' | 'quad';
  banners: BannerItem[];
  mobile_banners?: BannerItem[];  // آرایه جداگانه برای بنرهای موبایل
  // فیلدهای اضافی برای تنظیمات پیشرفته
  wide_bg?: boolean;
  bg_color?: string;
  bg_enabled?: boolean;
  height?: number;
  banner_width?: number;
  center_width?: number;
  border_radius?: number;
  gap?: number;
  display_order?: 'asc' | 'desc';
  layout_type?: 'horizontal' | 'vertical' | 'grid';
  // فیلدهای پایه که از WidgetConfig می‌آیند
  background_color?: string;
  padding?: string;
  margin?: string;
  custom_class?: string;
  // پدینگ و مارجین
  padding_top?: number;
  padding_bottom?: number;
  margin_left?: number;
  margin_right?: number;
  // تنظیمات نمایش
  show_title?: boolean;
  show_description?: boolean;
  // تنظیمات ظاهری
  shadow_enabled?: boolean;
  shadow_style?: 'default' | 'large';
  // تنظیمات موبایل
  mobile_vertical_display?: boolean;
  mobile_height?: number;
  mobile_banner_width?: number; // عرض بنر موبایل
  mobile_slider_width?: number; // عرض اسلایدر در موبایل
  mobile_padding_top?: number;
  mobile_padding_bottom?: number;
  mobile_margin_left?: number;
  mobile_margin_right?: number;
  // تنظیمات عکس موبایل
  mobile_image_mode?: 'auto' | 'separate';
  mobile_crop_width?: number; // عرض پیش‌فرض برای برش موبایل
  mobile_crop_height?: number; // ارتفاع پیش‌فرض برای برش موبایل

  mobile_cropped_image?: string; // URL عکس برش داده شده برای موبایل
  mobile_image?: string;
  // تنظیمات لیزی لود
  easy_load_enabled?: boolean;
  // نسبت‌های بنر برای بنرهای چندتایی
  banner1_ratio?: number;
  banner2_ratio?: number;
  banner3_ratio?: number;
  banner4_ratio?: number;
  banner5_ratio?: number;
}

// تنظیمات بنر تک تمام عرض
export interface FullBannerConfig extends WidgetConfig {
  title?: string;
  image?: string;
  mobile_image?: string;
  link?: string;
  openInNewTab?: boolean;
  banners?: BannerItem[];
  mobile_banners?: BannerItem[];  // آرایه جداگانه برای بنرهای موبایل
}

// تنظیمات بنر دو تایی
export interface DoubleBannerConfig extends WidgetConfig {
  banners: BannerItem[];
  mobile_banners?: BannerItem[];  // آرایه جداگانه برای بنرهای موبایل
}

// تنظیمات بنر سه تایی
export interface TripleBannerConfig extends WidgetConfig {
  banners: BannerItem[];
}

// تنظیمات بنر چهار تایی
export interface QuadBannerConfig extends WidgetConfig {
  banners: BannerItem[];
}

// تنظیمات محصولات
export interface ProductConfig extends WidgetConfig {
  category_id?: number;
  product_count: number;
  layout: 'grid' | 'list' | 'carousel';
  show_price: boolean;
  show_rating: boolean;
  show_discount: boolean;
}

// تنظیمات کاروسل محصولات
export interface ProductCarouselConfig extends WidgetConfig {
  title?: string;
  subtitle?: string;
  categoryId?: number | null;
  productCount: number;
  slidesPerView: number;
  showPrice: boolean;
  showRating: boolean;
  showDiscount: boolean;
  autoplay: boolean;
  autoplayDelay: number;
  wide_bg: boolean;
  bg_color: string;
  background_color?: string; // نگه‌داری برای سازگاری، اما توصیه به استفاده از bg_color
  padding?: string;
  margin?: string;
  showNavigation?: boolean;
  showIndicators?: boolean;
  showControls?: boolean;
  navigationStyle?: 'default' | 'minimal' | 'bold';
  navigationSize?: number;
  indicatorStyle?: 'default' | 'small' | 'large';
  indicatorColor?: string;
}

// تنظیمات دسته‌بندی
export interface CategoryConfig extends WidgetConfig {
  categories: CategoryItem[];
  layout: 'grid' | 'list';
  vertical_layout: boolean;
  show_product_count: boolean;
  columns: number;
  mobile_columns: number;
  box_width: 'full' | 'center';
  show_title: boolean;
  title: string;
  title_color: string;
  title_alignment?: 'left' | 'center' | 'right';
  card_style: 'default' | 'rounded' | 'shadow';
  image_size: 'small' | 'medium' | 'large';
  border_radius: 'none' | 'small' | 'medium' | 'large';
  text_alignment: 'left' | 'center' | 'right';
  background_color: string;
  show_background: boolean;
  full_width_background: boolean;
  show_border: boolean;
  border_color: string;
  border_width: 'thin' | 'medium' | 'thick';
  category_source: 'best_selling' | 'user_search' | 'user_interests' | 'manual';
  padding_top?: string;
  padding_bottom?: string;
}

// آیتم اسلایدر
export interface SlideItem {
  id?: number;
  title: string;
  description?: string;
  image: string;
  mobile_image?: string; // عکس جداگانه برای موبایل
  link?: string;
  openInNewTab?: boolean; // باز کردن لینک در صفحه جدید
  order: number;
  status: 'active' | 'inactive';
  showTitle?: boolean; // کنترل نمایش عنوان در اسلایدر
}

// آیتم بنر
export interface BannerItem {
  id?: number;
  title: string;
  description?: string;
  image: string;
  mobile_image?: string; // عکس جداگانه برای موبایل
  link?: string;
  openInNewTab?: boolean; // باز کردن لینک در صفحه جدید
  order: number;
  status: 'active' | 'inactive';
  showTitle?: boolean; // کنترل نمایش عنوان در بنر
}

// آیتم دسته‌بندی
export interface CategoryItem {
  id: number;
  title: string;
  description?: string;
  image: string;
  link: string;
  product_count?: number;
  // فیلدهای اضافی برای فرم ویرایش
  searchTerm?: string;
  showDropdown?: boolean;
  selectedCategory?: Record<string, unknown>;
}

// Request/Response types
export interface CreateWidgetRequest {
  title: string;
  description?: string;
  type: WidgetType;
  status?: WidgetStatus;
  page?: WidgetPage;
  show_on_mobile?: boolean;
  config?: WidgetConfig;
}

export interface UpdateWidgetRequest {
  title?: string;
  description?: string;
  type?: WidgetType;
  status?: WidgetStatus;
  order?: number;
  image?: string;
  link?: string;
  page?: WidgetPage;
  show_on_mobile?: boolean;
  config?: WidgetConfig;
}

export interface UpdateWidgetOrderRequest {
  items: Array<{
    id: number;
    order: number;
  }>;
}

export interface WidgetResponse {
  message: string;
  data: Widget;
}

export interface WidgetListResponse {
  message: string;
  data: Widget[];
}

// Form types for admin
export interface WidgetFormData {
  title: string;
  description: string;
  type: WidgetType;
  status: WidgetStatus;
  page: WidgetPage;
  config: WidgetConfig;
}

// Widget type labels (for UI)
export const WIDGET_TYPE_LABELS: Record<WidgetType, string> = {
  'main-slider-side-banner': 'اسلایدر اصلی و بنر کناری',
  'single-slider-side': 'اسلایدر تکی (بدون سایدبار)',
  'full-slider': 'اسلایدر تمام عرض',
  'full-banner': 'بنر تکی',
  'double-banner': 'بنر دوتایی',
  'triple-banner': 'بنر سه‌تایی',
  'quad-banner': 'بنر 4 تایی',
  'penta-banner': 'بنر پنج‌تایی',
  'brands-slider': 'اسلایدر برند ها',
  'services-slider': 'اسلایدر خدمات',
  'category-box': 'دسته بندی محصولات',
  'blog-posts': 'نوشته های وبلاگ',
  'product-carousel': 'کاروسل محصولات',
  'mobile-header': 'هدر موبایل',
};

// Widget status labels
export const WIDGET_STATUS_LABELS: Record<WidgetStatus, string> = {
  'active': 'فعال',
  'inactive': 'غیرفعال',
  'draft': 'پیش‌نویس',
};

// Widget page labels
export const WIDGET_PAGE_LABELS: Record<WidgetPage, string> = {
  'home': 'صفحه اصلی',
  'category': 'صفحه دسته‌بندی',
  'product': 'صفحه محصول',
  'about': 'درباره ما',
  'contact': 'تماس با ما',
};

// Validation helpers
export const validateWidget = (widget: Partial<Widget>): string[] => {
  const errors: string[] = [];

  if (!widget.title) {
    errors.push('عنوان ابزارک الزامی است');
  }

  if (!widget.type) {
    errors.push('نوع ابزارک الزامی است');
  }

  if (widget.type && !Object.keys(WIDGET_TYPE_LABELS).includes(widget.type)) {
    errors.push('نوع ابزارک نامعتبر است');
  }

  return errors;
};

// Helper functions
export const getWidgetTypeLabel = (type: WidgetType): string => {
  return WIDGET_TYPE_LABELS[type] || type;
};

export const getWidgetStatusLabel = (status: WidgetStatus): string => {
  return WIDGET_STATUS_LABELS[status] || status;
};

export const getWidgetPageLabel = (page: WidgetPage): string => {
  return WIDGET_PAGE_LABELS[page] || page;
};

export const generateWidgetCode = (type: WidgetType, count: number): string => {
  const typeMap: Record<WidgetType, string> = {
    'main-slider-side-banner': 'SLIDER_MAIN',
    'single-slider-side': 'SLIDER_SINGLE',
    'full-slider': 'SLIDER_FULL',
    'full-banner': 'BANNER_FULL',
    'double-banner': 'BANNER_DOUBLE',
    'triple-banner': 'BANNER_TRIPLE',
    'quad-banner': 'BANNER_QUAD',
    'penta-banner': 'BANNER_PENTA',
    'brands-slider': 'BRANDS_SLIDER',
    'services-slider': 'SERVICES_SLIDER',
    'category-box': 'CATEGORY_BOX',
    'blog-posts': 'BLOG_POSTS',
    'product-carousel': 'PRODUCT_CAROUSEL',
    'mobile-header': 'MOBILE_HEADER',
  };

  const prefix = typeMap[type] || 'WIDGET';
  return `${prefix}_${String(count).padStart(3, '0')}`;
};

// interface برای محصولات
export interface Product {
  id: number;
  name: string;
  image?: string;
  price: number;
  original_price?: number;
  discount_percentage?: number;
  rating?: number;
  rating_count?: number;
  sku?: string | number;
  slug?: string | null;
  seo_title?: string | null;
  features?: string[];
  stock_quantity?: number;
  in_stock?: boolean;
  call_for_price?: boolean;
  created_at?: string;
  updated_at?: string;
} 
export interface OrderItem {
  id: number;
  name: string;
  sku: string;
  quantity: number;
  image?: string;
  unitPrice: number;
  totalPrice: number;
}

export interface Order {
  id: number;
  orderNumber: string;
  date: string;
  status: string;
  paymentStatus: string;
  paymentMethod: string;
  finalAmount: number;
  totalAmount: number;
  userName?: string;
  userMobile?: string;
  userEmail?: string;
  userNationalCode?: string;
  userLastLoginIP?: string;
  shippingAddress?: string;
  billingAddress?: string;
  createdAt?: string; // Used in mapping
}

export interface OrderResponse {
  success: boolean;
  data: {
    orders: Record<string, unknown>[]; // Raw order data from API
  };
}

export interface OrderItemsResponse {
  success: boolean;
  data: {
    items: OrderItem[];
  };
}

import { NewOrderFormType, Order } from "@/types/order";
import axios from "axios";

export const getAllOrders = async () => {
  return await axios.get<Order[]>("http://localhost:8080/orders");
};

export const createNewOrder = async (data: NewOrderFormType) => {
  return await axios.post<Order>("http://localhost:8080/orders", data);
};

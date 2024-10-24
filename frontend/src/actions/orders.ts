"use server";

import { Order } from "@/types/order";
import axios from "axios";

export const getAllOrders = async () => {
  return await axios.get<Order[]>("http://localhost:8080/orders");
};

import { OrderProps } from "@/types/order";
import React from "react";

const OrderCards: React.FC<OrderProps> = ({ order }) => {
  return <div>{order.id}</div>;
};

export default OrderCards;

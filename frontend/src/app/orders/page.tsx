"use server";

import { createNewOrder, getAllOrders } from "@/actions/orders";
import React from "react";
import z from "zod";
import OrderCards from "./orderCards";
import NewOrderCard from "./NewOrderCard";

const Page = async () => {
  const orders = await getAllOrders();

  return (
    <div>
      <div>
        {orders.data.map((o, i) => (
          <OrderCards key={i} order={o} />
        ))}
      </div>
      <NewOrderCard />
    </div>
  );
};

Page.getInitialProps = async (ctx: any) => {
  return {};
};

export default Page;

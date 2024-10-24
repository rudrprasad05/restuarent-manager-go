"use server";

import { getAllOrders } from "@/actions/orders";
import React from "react";
import z from "zod";
import OrderCards from "./orderCards";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";

const page = async () => {
  const orders = await getAllOrders();

  const NewOrderForm = z.object({
    customer: z.string(),
    amount: z.string(),
    orderStatus: z.string(),
  });

  type NewOrderFormType = z.infer<typeof NewOrderForm>;

  const form = useForm<NewOrderFormType>({
    resolver: zodResolver(NewOrderForm),
    defaultValues: {
      customer: "",
      amount: "",
      orderStatus: "",
    },
  });

  const onSubmit = async (data: NewOrderFormType) => {
    // const res = await ChangePassword(token as string, data.password)
  };

  return (
    <div>
      <div>
        {orders.data.map((o, i) => (
          <OrderCards key={i} order={o} />
        ))}
      </div>

      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-3">
          <FormField
            control={form.control}
            name="customer"
            render={({ field }) => (
              <FormItem>
                <FormLabel>customer</FormLabel>
                <FormControl>
                  <Input
                    type="text"
                    autoComplete="off"
                    placeholder="enter customer"
                    {...field}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="amount"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Enter amount</FormLabel>
                <FormControl>
                  <Input
                    type="text"
                    autoComplete="off"
                    placeholder="enter amount"
                    {...field}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="orderStatus"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Enter status</FormLabel>
                <FormControl>
                  <Input
                    type="text"
                    autoComplete="off"
                    placeholder="enter orderStatus"
                    {...field}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <Button className="w-full" type="submit">
            Verify
          </Button>
        </form>
      </Form>
    </div>
  );
};

"use client";

import React from "react";
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
import { NewOrderForm, NewOrderFormType } from "@/types/order";
import { createNewOrder } from "@/actions/orders";

const NewOrderCard = () => {
  const form = useForm<NewOrderFormType>({
    resolver: zodResolver(NewOrderForm),
    defaultValues: {
      customer: "",
      amount: "",
      orderStatus: "",
    },
  });

  const onSubmit = async (data: NewOrderFormType) => {
    const res = await createNewOrder(data);
    console.log(res.data);
  };
  return (
    <div>
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

export default NewOrderCard;

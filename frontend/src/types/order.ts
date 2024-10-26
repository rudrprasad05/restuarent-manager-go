import { z } from "zod";

export interface Order {
  id: number; // Same as int in Go
  customer: string; // Same as string in Go
  amount: number; // Same as float64 in Go
  orderStatus: string; // Same as string in Go
}

export interface OrderProps {
  order: Order;
}

export const NewOrderForm = z.object({
  customer: z.string(),
  amount: z.string(),
  orderStatus: z.string(),
});

export type NewOrderFormType = z.infer<typeof NewOrderForm>;

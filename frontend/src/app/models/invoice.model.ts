import { Product } from "./product.model";

export interface Invoice {
  id?: number;
  products?: Array<Product>;
  quantities?: Array<number>;
  status?: string;
  CreatedAt?: Date;
}

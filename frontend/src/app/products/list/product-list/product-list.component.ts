import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { Invoice } from 'src/app/models/invoice.model';
import { Product } from 'src/app/models/product.model';
import { InvoiceService } from 'src/app/services/invoice.service';
import { ProductService } from 'src/app/services/product.service';
import { ProductFormComponent } from '../../form/product-form/product-form.component';
import { InvoiceListComponent } from '../invoice-list/invoice-list.component';

@Component({
  selector: 'product-list-component',
  templateUrl: './product-list.component.html',
  styleUrls: ['./product-list.component.scss']
})
export class ProductListComponent implements OnInit {
  products: Array<Product> = new Array<Product>();

  constructor(
    private productService: ProductService,
    private dialog: MatDialog,
    private dialogInvoices: MatDialog,
    private invoiceService: InvoiceService,
  ) { }

  ngOnInit() {
    this.loadProducts();
  }

  loadProducts() {
    this.productService.getProducts().subscribe(data => {
      this.products = data;
    });
  }

  editProduct(product: Product) {
    const dialogRef = this.dialog.open(ProductFormComponent, {
      data: product,
      width: '800px',
      height: '550px'
    });

    dialogRef.afterClosed().subscribe(() => {
      this.loadProducts();
    });
  }

  deleteProduct(id: number) {
    if (confirm('Tem certeza que deseja excluir este produto?')) {
      this.productService.deleteProduct(id).subscribe(() => this.loadProducts());
    }
  }

  addProduct() {
    const dialogRef = this.dialog.open(ProductFormComponent, {
      width: '800px',
      height: '550px'
    });

    dialogRef.afterClosed().subscribe(() => {
      this.loadProducts();
    });
  }

  requestInvoice(product: Product): void {
    const newInvoice: Invoice = {
      products: new Array<Product>(product),
      quantities: new Array<number>(10),
    }

    this.invoiceService.createInvoice(newInvoice).subscribe({
      next: () => { console.log('Criando...') },
      error: (err) => { console.error(err) },
    })
  }

  seeInvoices(): void {
    this.dialogInvoices.open(InvoiceListComponent, {
      width: '800px',
      height: '550px'
    })
  }
}

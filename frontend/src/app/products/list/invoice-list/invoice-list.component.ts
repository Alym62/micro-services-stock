import { Component, OnInit } from '@angular/core';
import { MatDialogRef } from '@angular/material/dialog';
import { Invoice } from 'src/app/models/invoice.model';
import { InvoiceService } from 'src/app/services/invoice.service';

@Component({
  selector: 'invoice-list-component',
  templateUrl: './invoice-list.component.html',
  styleUrls: ['./invoice-list.component.scss']
})
export class InvoiceListComponent implements OnInit {
  invoices: Array<Invoice> = new Array<Invoice>();

  constructor(
    private dialogRef: MatDialogRef<InvoiceListComponent>,
    private invoiceService: InvoiceService,
  ) { }

  ngOnInit() {
    this.loadProducts();
  }

  loadProducts() {
    this.invoiceService.getInvoices().subscribe(data => {
      this.invoices = data;
    });
  }

  close(): void {
    this.dialogRef.close();
  }
}

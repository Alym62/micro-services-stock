import { Component, Inject } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { Product } from 'src/app/models/product.model';
import { ProductService } from 'src/app/services/product.service';

@Component({
  selector: 'product-form-component',
  templateUrl: './product-form.component.html',
  styleUrls: ['./product-form.component.scss']
})
export class ProductFormComponent {
  productForm: FormGroup;

  constructor(
    private fb: FormBuilder,
    private productService: ProductService,
    private dialogRef: MatDialogRef<ProductFormComponent>,
    @Inject(MAT_DIALOG_DATA) public data: Product
  ) {
    this.productForm = this.fb.group({
      name: [data?.name || '', Validators.required],
      price: [data?.price || '', Validators.required],
      description: [data?.description || ''],
      quantity: [data?.quantity || ''],
    });
  }

  save() {
    const product = this.productForm.value;
    if (this.data?.id) {
      this.productService.updateProduct(this.data.id, product).subscribe(() => this.dialogRef.close());
    } else {
      this.productService.createProduct(product).subscribe(() => this.dialogRef.close());
    }
  }

  close(): void {
    this.dialogRef.close();
  }
}

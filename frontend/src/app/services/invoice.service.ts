import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Invoice } from '../models/invoice.model';

@Injectable({
  providedIn: 'root'
})
export class InvoiceService {
  private apiUrl = 'http://localhost:8080/api/v1/invoice';

  constructor(private http: HttpClient) { }

  getInvoices(): Observable<Array<Invoice>> {
    return this.http.get<Array<Invoice>>(`${this.apiUrl}/list`);
  }

  createInvoice(invoice: Invoice): Observable<Invoice> {
    return this.http.post<Invoice>(`${this.apiUrl}/create`, invoice);
  }
}

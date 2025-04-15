// src/modules/go-client/go-client.service.ts
import { Injectable } from '@nestjs/common';
import { HttpService } from '@nestjs/axios'; // Correct import
import { map } from 'rxjs/operators';

@Injectable()
export class GoClientService {
  constructor(private readonly httpService: HttpService) {}

  async getDataFromGoService(): Promise<any> {
    try {
      return this.httpService.get('http://localhost:8080/ping').pipe(
        map(response => response.data),
      ).toPromise();
    } catch (error) {
      console.error('Error fetching data from Go service:', error);
      throw error;
    }
  }
}
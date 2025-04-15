// src/modules/ping/ping.service.ts
import { Injectable } from '@nestjs/common';

@Injectable()
export class PingService {
  ping(): string {
    return 'pong';
  }
}
// src/modules/go-client/go-client.controller.ts
import { Controller, Get } from '@nestjs/common';
import { GoClientService } from './go-client.service';

@Controller('go-client')
export class GoClientController {
  constructor(private readonly goClientService: GoClientService) {}

  @Get('data')
  async getData(): Promise<any> {
    return this.goClientService.getDataFromGoService();
  }
}
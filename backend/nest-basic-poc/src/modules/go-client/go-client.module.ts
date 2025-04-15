// src/modules/go-client/go-client.module.ts
import { Module } from '@nestjs/common';
import { HttpModule } from '@nestjs/axios';
import { GoClientService } from './go-client.service';
import { GoClientController } from './go-client.controller';

@Module({
  imports: [HttpModule],
  controllers: [GoClientController], // Add controller
  providers: [GoClientService],
  exports: [GoClientService],
})
export class GoClientModule {}
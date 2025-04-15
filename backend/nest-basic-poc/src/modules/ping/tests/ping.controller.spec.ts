// src/modules/ping/ping.controller.spec.ts
import { Test, TestingModule } from '@nestjs/testing';
import { PingController } from '../ping.controller';
import { PingService } from '../ping.service';

describe('PingController', () => {
  let pingController: PingController;

  beforeEach(async () => {
    const app: TestingModule = await Test.createTestingModule({
      controllers: [PingController],
      providers: [PingService],
    }).compile();

    pingController = app.get<PingController>(PingController);
  });

  describe('ping', () => {
    it('should return "pong"', () => {
      expect(pingController.ping()).toBe('pong');
    });
  });
});
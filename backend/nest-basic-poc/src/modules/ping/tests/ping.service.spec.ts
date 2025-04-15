// src/modules/ping/ping.service.spec.ts
import { Test, TestingModule } from '@nestjs/testing';
import { PingService } from '../ping.service';

describe('PingService', () => {
  let pingService: PingService;

  beforeEach(async () => {
    const app: TestingModule = await Test.createTestingModule({
      providers: [PingService],
    }).compile();

    pingService = app.get<PingService>(PingService);
  });

  describe('ping', () => {
    it('should return "pong"', () => {
      expect(pingService.ping()).toBe('pong');
    });
  });
});
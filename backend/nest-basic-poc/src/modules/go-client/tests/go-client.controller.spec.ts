// src/modules/go-client/go-client.controller.spec.ts
import { Test, TestingModule } from '@nestjs/testing';
import { GoClientController } from '../go-client.controller';
import { GoClientService } from '../go-client.service';
import { HttpModule } from '@nestjs/axios';

describe('GoClientController', () => {
  let controller: GoClientController;
  let service: GoClientService;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      imports: [HttpModule],
      controllers: [GoClientController],
      providers: [GoClientService],
    }).compile();

    controller = module.get<GoClientController>(GoClientController);
    service = module.get<GoClientService>(GoClientService);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });

  it('should get data from Go service', async () => {
    const mockData = { message: 'Data from Go service' };
    jest.spyOn(service, 'getDataFromGoService').mockResolvedValue(mockData);

    const result = await controller.getData();
    expect(result).toEqual(mockData);
  });
});
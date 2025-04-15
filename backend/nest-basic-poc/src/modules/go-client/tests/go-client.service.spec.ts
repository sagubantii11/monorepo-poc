// src/modules/go-client/go-client.service.spec.ts
import { Test, TestingModule } from '@nestjs/testing';
import { GoClientService } from '../go-client.service';
import { HttpModule, HttpService } from '@nestjs/axios'; // Correct import
import { of } from 'rxjs';

describe('GoClientService', () => {
  let service: GoClientService;
  let httpService: HttpService;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      imports: [HttpModule],
      providers: [GoClientService],
    }).compile();

    service = module.get<GoClientService>(GoClientService);
    httpService = module.get<HttpService>(HttpService);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });

  it('should get data from Go service', async () => {
    const mockData = { message: 'Data from Go service' };
    jest.spyOn(httpService, 'get').mockReturnValue(of({ data: mockData }) as any);

    const result = await service.getDataFromGoService();
    expect(result).toEqual(mockData);
  });

  it('should handle errors from Go service', async () => {
    jest.spyOn(httpService, 'get').mockImplementation(() => {
      throw new Error('Go service error');
    });

    await expect(service.getDataFromGoService()).rejects.toThrow('Go service error');
  });
});
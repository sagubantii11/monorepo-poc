// src/modules/database/database.module.spec.ts
import { Test, TestingModule } from '@nestjs/testing';
import { DatabaseModule } from './database.module';
import { TypeOrmModule } from '@nestjs/typeorm';
import configuration from '../../config/configuration';
import { ConfigModule } from '@nestjs/config';

describe('DatabaseModule', () => {
  let module: TestingModule;

  beforeEach(async () => {
    module = await Test.createTestingModule({
      imports: [
        DatabaseModule,
        TypeOrmModule.forRootAsync({
          imports: [ConfigModule],
          useFactory: async (configService) => {
            return {
              type: 'sqlite',
              database: ':memory:',
              entities: [__dirname + '/../**/*.entity{.ts,.js}'],
              synchronize: true,
            };
          },
          inject: [ConfigModule],
        }),
        ConfigModule.forRoot({
          load: [configuration],
          isGlobal: true,
        }),
      ],
    }).compile();
  });

  it('should be defined', () => {
    expect(module).toBeDefined();
  });
});
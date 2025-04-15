// src/modules/auth/auth.service.spec.ts
import { Test, TestingModule } from '@nestjs/testing';
import { AuthService } from '../auth.service';
import { UnauthorizedException } from '@nestjs/common';

describe('AuthService', () => {
  let authService: AuthService;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [AuthService],
    }).compile();

    authService = module.get<AuthService>(AuthService);
  });

  describe('login', () => {
    it('should return an access token on successful login', async () => {
      expect(await authService.login('test', 'password')).toEqual({ access_token: 'fake_token' });
    });

    it('should throw UnauthorizedException on invalid credentials', async () => {
      await expect(authService.login('invalid', 'credentials')).rejects.toThrow(UnauthorizedException);
    });
  });
});
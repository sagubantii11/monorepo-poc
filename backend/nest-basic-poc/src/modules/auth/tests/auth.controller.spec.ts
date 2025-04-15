// src/modules/auth/auth.controller.spec.ts
import { Test, TestingModule } from '@nestjs/testing';
import { AuthController } from '../auth.controller';
import { AuthService } from '../auth.service';
import { UnauthorizedException } from '@nestjs/common';

describe('AuthController', () => {
  let authController: AuthController;
  let authService: AuthService;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [AuthController],
      providers: [AuthService],
    }).compile();

    authController = module.get<AuthController>(AuthController);
    authService = module.get<AuthService>(AuthService);
  });

  describe('login', () => {
    it('should return an access token on successful login', async () => {
      const result = { access_token: 'fake_token' };
      jest.spyOn(authService, 'login').mockResolvedValue(result);

      expect(await authController.login({ username: 'test', password: 'password' })).toBe(result);
    });

    it('should throw UnauthorizedException on invalid credentials', async () => {
      jest.spyOn(authService, 'login').mockRejectedValue(new UnauthorizedException('Invalid credentials'));

      await expect(
        authController.login({ username: 'invalid', password: 'credentials' }),
      ).rejects.toThrow(UnauthorizedException);
    });
  });
});
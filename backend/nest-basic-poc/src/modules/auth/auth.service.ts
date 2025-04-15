// src/modules/auth/auth.service.ts
import { Injectable, UnauthorizedException } from '@nestjs/common';

@Injectable()
export class AuthService {
  async login(username: string, password: string): Promise<any> {
    // In a real application, you would validate the username and password against a database.
    if (username === 'test' && password === 'password') {
      // In a real application, you would generate a JWT token.
      return { access_token: 'fake_token' };
    } else {
      throw new UnauthorizedException('Invalid credentials');
    }
  }
}
import { registerAs } from "@nestjs/config";

export default registerAs('config', () => ({
  port: process.env.PORT || 3000,
  database: {
    type: process.env.DATABASE_TYPE || 'postgres',
    host: process.env.DATABASE_HOST || 'localhost',
    port: process.env.DATABASE_PORT || 5432,
    username: process.env.DATABASE_USER || 'user',
    password: process.env.DATABASE_PASSWORD || 'password',
    database: process.env.DATABASE_NAME || 'mydb',
    synchronize: process.env.NODE_ENV !== 'production',
    autoLoadEntities: true,
  },
}));
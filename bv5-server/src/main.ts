import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import {
  utilities as nestWinstonModuleUtilities,
  WinstonModule,
} from 'nest-winston';
import {
  format as WinstonFormat,
  transports as WinstonTransports,
} from 'winston';

async function bootstrap() {
  const app = await NestFactory.create(AppModule, {
    logger: WinstonModule.createLogger({
      transports: [
        new WinstonTransports.Console({
          level: 'info',
          format: WinstonFormat.combine(
            WinstonFormat.timestamp(),
            WinstonFormat.ms(),
            WinstonFormat.errors({ stack: true }),
            WinstonFormat.colorize(),
            nestWinstonModuleUtilities.format.nestLike('BardView5'),
          ),
        }),
      ],
    }),
  });
  await app.listen(3000);
}

bootstrap();

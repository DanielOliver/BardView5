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
import { DocumentBuilder, SwaggerModule } from '@nestjs/swagger';
import { ApiPrefixV1 } from './globals';
import { ValidationPipe } from '@nestjs/common';

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
  app.useGlobalPipes(new ValidationPipe());

  const config = new DocumentBuilder()
    .setTitle('BardView5')
    .setDescription('You want to play some table top RPGs?')
    .setVersion('0.0.1')
    .build();
  const document = SwaggerModule.createDocument(app, config);
  SwaggerModule.setup(ApiPrefixV1, app, document, {
    customSiteTitle: 'BardView5 Swagger',
  });
  await app.listen(3000);
}

bootstrap();

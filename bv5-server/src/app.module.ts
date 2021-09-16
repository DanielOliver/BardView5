import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { UserService } from './user.service';
import { PrismaService } from './prisma.service';
import { ConfigModule } from '@nestjs/config';
import { OrmModule } from './orm/orm.module';

@Module({
  imports: [ConfigModule.forRoot(), OrmModule],
  controllers: [AppController],
  providers: [UserService, PrismaService],
})
export class AppModule {}

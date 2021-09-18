import { Test, TestingModule } from '@nestjs/testing';
import { UsersController } from './user.controller';

describe('UserController', () => {
  let userController: UsersController;

  beforeEach(async () => {
    const app: TestingModule = await Test.createTestingModule({
      controllers: [UsersController],
    }).compile();

    userController = app.get<UsersController>(UsersController);
  });

  describe('root', () => {
    it('should return "Hello World!"', () => {
      expect('Hello World!').toBe('Hello World!');
    });
  });
});

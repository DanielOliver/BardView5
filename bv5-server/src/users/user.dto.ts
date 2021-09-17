import { IsEmail, IsNotEmpty } from 'class-validator';

class UserBase {
  @IsEmail()
  email!: string;
  @IsNotEmpty()
  name!: string;
}

export class UserResponse extends UserBase {
  id!: string;

  public constructor(init?: Partial<UserResponse>) {
    super();

    this.email = init.email;
    this.name = init.name;
    this.id = init.id;
  }
}

export class UserCreationRequest extends UserBase {}

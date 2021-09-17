import { IsEmail, IsNotEmpty } from 'class-validator';

class UserBase {
  @IsEmail()
  email!: string;
  @IsNotEmpty()
  name!: string;
  tags: string[];
}

export class UserResponse extends UserBase {
  uid!: string;

  public constructor(init?: Partial<UserResponse>) {
    super();

    this.email = init.email;
    this.name = init.name;
    this.uid = init.uid;
    this.tags = init.tags;
  }
}

export class UserCreationRequest extends UserBase {}

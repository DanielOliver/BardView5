import { IsEmail, IsNotEmpty, MaxLength } from 'class-validator';

class UserBase {
  @IsEmail()
  email!: string;
  @IsNotEmpty()
  name!: string;
  @MaxLength(40, {
    each: true,
  })
  tags: string[];
}

export class UserResponse extends UserBase {
  uid!: string;
  isActive: boolean;

  public constructor(init?: Partial<UserResponse>) {
    super();

    this.email = init.email;
    this.name = init.name;
    this.uid = init.uid;
    this.tags = init.tags;
    this.isActive = init.isActive;
  }
}

export class UserCreationRequest extends UserBase {}

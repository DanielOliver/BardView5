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

  public constructor(init?: Partial<UserResponse>) {
    super();

    this.email = init.email;
    this.name = init.name;
    this.uid = init.uid;
    this.tags = init.tags;
  }
}

export class UserCreationRequest extends UserBase {}

export class RoleAssignmentResponse {
  uid!: string;
  createdAt!: Date;
  effectiveDate!: Date;
  endDate?: Date;
  isActive: boolean;
  roleUid!: string;
  userUid!: string;

  public constructor(init?: Partial<RoleAssignmentResponse>) {
    this.uid = init.uid;
    this.createdAt = init.createdAt;
    this.uid = init.uid;
    this.effectiveDate = init.effectiveDate;
    this.endDate = init.endDate;
    this.isActive = init.isActive;
    this.roleUid = init.roleUid;
    this.userUid = init.userUid;
  }
}

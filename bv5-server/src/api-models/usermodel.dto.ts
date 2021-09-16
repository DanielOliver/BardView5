export class UserModel {
  id!: number;
  email!: string;
  name?: string;

  public constructor(init?: Partial<UserModel>) {
    // Object.assign(this, init);
    this.email = init.email;
    this.name = init.name;
    this.id = init.id;
  }
}

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

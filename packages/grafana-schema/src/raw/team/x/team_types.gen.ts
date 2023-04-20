// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by:
//     kinds/gen.go
// Using jennies:
//     TSResourceJenny
//
// Run 'make gen-cue' from repository root to regenerate.

export enum Permission {
  Admin = 4,
  Editor = 2,
  Member = 0,
  Viewer = 1,
}

export interface Team {
  /**
   * AccessControl metadata associated with a given resource.
   */
  accessControl?: Record<string, boolean>;
  /**
   * AvatarUrl is the team's avatar URL.
   */
  avatarUrl?: string;
  /**
   * Email of the team.
   */
  email?: string;
  /**
   * MemberCount is the number of the team members.
   */
  memberCount: number;
  /**
   * Name of the team.
   */
  name: string;
  /**
   * OrgId is the ID of an organisation the team belongs to.
   */
  orgId: number;
  /**
   * TODO - it seems it's a team_member.permission, unlikely it should belong to the team kind
   */
  permission: Permission;
}

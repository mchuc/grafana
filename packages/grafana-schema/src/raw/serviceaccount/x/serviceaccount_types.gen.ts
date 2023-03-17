// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by:
//     kinds/gen.go
// Using jennies:
//     TSTypesJenny
//     LatestMajorsOrXJenny
//
// Run 'make gen-cue' from repository root to regenerate.

export interface ServiceAccount {
  spec: {
    /**
     * ID is the unique identifier of the service account in the database.
     */
    id: number;
    /**
     * OrgId is the ID of an organisation the service account belongs to.
     */
    orgId: number;
    /**
     * Name of the service account.
     */
    name: string;
    /**
     * Login of the service account.
     */
    login: string;
    /**
     * IsDisabled indicates if the service account is disabled.
     */
    isDisabled: boolean;
    /**
     * Role is the Grafana organization role of the service account which can be 'Viewer', 'Editor', 'Admin'.
     */
    role: OrgRole;
    /**
     * Tokens is the number of active tokens for the service account.
     * Tokens are used to authenticate the service account against Grafana.
     */
    tokens: number;
    /**
     * AvatarUrl is the service account's avatar URL. It allows the frontend to display a picture in front
     * of the service account.
     */
    avatarUrl: string;
    /**
     * AccessControl metadata associated with a given resource.
     */
    accessControl?: Record<string, boolean>;
    /**
     * Teams is a list of teams the service account belongs to.
     */
    teams?: Array<string>;
    /**
     * Created indicates when the service account was created.
     */
    created?: number;
    /**
     * Updated indicates when the service account was updated.
     */
    updated?: number;
  };
}

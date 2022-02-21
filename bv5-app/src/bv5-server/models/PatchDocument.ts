/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

/**
 * A JSONPatch document as defined by RFC 6902
 */
export type PatchDocument = Array<{
    /**
     * The operation to be performed
     */
    op: 'add' | 'remove' | 'replace' | 'move' | 'copy' | 'test';
    /**
     * A JSON-Pointer
     */
    path: string;
    /**
     * The value to be used within the operations.
     */
    value?: any;
    /**
     * A string containing a JSON Pointer value.
     */
    from?: string;
}>;